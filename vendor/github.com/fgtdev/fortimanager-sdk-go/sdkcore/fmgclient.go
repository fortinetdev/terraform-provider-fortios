package fmgclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fgtdev/fortimanager-sdk-go/auth"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Request represents a JSON-RPC request sent by a client.
type Request struct {
	Id      uint64         `json:"id"`
	Method  string         `json:"method"`
	Session string         `json:"session"`
	Params  [1]interface{} `json:"params"`
}

// FmgSDKClient is for communicating with FortiManager
type FmgSDKClient struct {
	Ipaddr string
	User   string
	Passwd string
	Debug  string
	Client *http.Client
	Auth   *auth.AuthBlob
}

// NewClient is for creating new client
// Input:
//   @ip: ipaddress of fortimanager
//   @user: username of fortimanager
//   @passwd: passwd of fortimanager
//   @client: http client used
// Output:
//   @FmgSDKClient: fmg client
func NewClient(ip, user, passwd string, client *http.Client) *FmgSDKClient {
	d := os.Getenv("TRACEDEBUG")

	fmgsdkclient := FmgSDKClient{
		Ipaddr: ip,
		User:   user,
		Passwd: passwd,
		Client: client,
		Debug:  d,
	}
	auth := auth.FmgAuthInit(fmgsdkclient.Login, fmgsdkclient.Logout)
	fmgsdkclient.Auth = auth
	return &fmgsdkclient
}

// Execute is for sending the http request to FortiManager
// Input:
//   @req: http request
// Output:
//   @result: http response result
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) Execute(req *Request) (result map[string]interface{}, err error) {
	j, _ := json.Marshal(req)

	if c.Debug == "ON" || c.Debug == "on" {
		log.Printf("[TRACEDEBUG] ==> request: %s", j)
	}

	httpResp, err := c.Client.Post("http://"+c.Ipaddr+"/jsonrpc", "application/json", bytes.NewBuffer(j))
	if err != nil {
		err = fmt.Errorf("Login failed: %s", err)
		return
	}
	defer httpResp.Body.Close()

	// check response
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}

	if c.Debug == "ON" || c.Debug == "on" {
		log.Printf("[TRACEDEBUG] ==> result: %s", body)
	}

	result = map[string]interface{}{}
	json.Unmarshal([]byte(string(body)), &result)

	if result != nil {
		if len(result) == 0 {
			err = fmt.Errorf("No result got, details:\n%s", body)
			return
		}

		if id := uint64(result["id"].(float64)); id != req.Id {
			err = fmt.Errorf("id not match, should be 1, but is %d", id)
			return
		}

		if result["result"] != nil {
			status := (result["result"].([]interface{}))[0].(map[string]interface{})["status"].(map[string]interface{})

			c := uint(status["code"].(float64))
			m := status["message"].(string)
			if c != 0 || m != "OK" {
				err = fmt.Errorf("status not right: code is %d, message is %s", c, m)
				return
			}
		} else {
			err = fmt.Errorf("can't get response status: %s", err)
			return
		}
	}

	return
}

// Login is for logging in
// Output:
//   @session: login session
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) Login() (session string, err error) {
	params := map[string]interface{}{
		"data": map[string]string{
			"user":   c.User,
			"passwd": c.Passwd,
		},
		"url": "/sys/login/user",
	}

	req := &Request{
		Id:     1,
		Method: "exec",
		Params: [1]interface{}{params},
	}

	result, err := c.Execute(req)
	if err != nil {
		return "", fmt.Errorf("login failed:%s", err)
	}

	session = result["session"].(string)
	return
}

// Logout is for logging out
// Input:
//   @s: login session
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) Logout(s string) (err error) {
	params := map[string]interface{}{
		"url": "/sys/logout",
	}

	req := &Request{
		Id:      1,
		Method:  "exec",
		Params:  [1]interface{}{params},
		Session: s,
	}

	_, err = c.Execute(req)
	if err != nil {
		err = fmt.Errorf("logout failed:%s", err)
		return
	}
	return
}

// Do is for executing the related request
// Input:
//   @method: operation method
//   @params: request infor needed
// Output:
//   @output: result infor got back
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) Do(method string, params map[string]interface{}) (output map[string]interface{}, err error) {
	//session, err := c.Login()
	session, err := c.Auth.GetToken(c.Ipaddr, false)
	if err != nil {
		return nil, fmt.Errorf("Executing failed", err)
	}
	//defer c.Logout(session)
	defer c.Auth.PutToken(session)
	req := &Request{
		Id:      1,
		Method:  method,
		Params:  [1]interface{}{params},
		Session: session,
	}

	retry := false
	reqOut, reqErr := c.Execute(req)
	// RETRY IN CASE WE USE A EXPIRED SESSION TOKEN
	if reqErr != nil && reqOut["result"] != nil {
		status := (reqOut["result"].([]interface{}))[0].(map[string]interface{})["status"].(map[string]interface{})
		status_code := int(status["code"].(float64))
		if status_code == -11 {
			retry = true
		}
	}
	if retry == true {
		session, tokenErr := c.Auth.GetToken(c.Ipaddr, true)
		if tokenErr != nil {
			return nil, fmt.Errorf("Executing failed", tokenErr)
		}
		defer c.Auth.PutToken(session)
		req = &Request{
			Id:      1,
			Method:  method,
			Params:  [1]interface{}{params},
			Session: session,
		}
		reqOut, reqErr = c.Execute(req)
	}
	return reqOut, reqErr
}

// Trace is for debugging
// Input:
//   @s: function name to be traced
// Output:
//   @func()
func (f *FmgSDKClient) Trace(s string) func() {
	if f.Debug == "ON" {
		log.Printf("[TRACEDEBUG] -> Enter %s <-", s)
		return func() { log.Printf("[TRACEDEBUG]    -> Leave %s <-", s) }
	}

	return func() {}
}
