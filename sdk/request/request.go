package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-fortios/sdk/config"
)

// Request describes the request to FortiOS service
type Request struct {
	Config       config.Config
	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	Path         string
	Params       interface{}
	Data         *bytes.Buffer
}

// Cookies contains the CSRF Token and Cookie by login using username and password
type Cookies struct {
	CSRFToken string
	Cookie    string
}

// New creates reqeust object with http method, path, params and data,
// It will save the http request, path, etc. for the next operations
// such as sending data, getting response, etc.
// It returns the created request object to the gobal plugin client.
func New(c config.Config, method string, path string, params interface{}, data *bytes.Buffer) *Request {
	var h *http.Request

	if data == nil {
		h, _ = http.NewRequest(method, "", nil)
	} else {
		h, _ = http.NewRequest(method, "", data)
	}

	r := &Request{
		Config:      c,
		Path:        path,
		HTTPRequest: h,
		Params:      params,
		Data:        data,
	}
	return r
}

// Build Request header

// Build Request Sign/Login Info

// Send request data to FortiOS.
// If errors are encountered, it returns the error.
func (r *Request) Send(headers *map[string]string) error {
	return r.Send2(15, false, headers)
}

// Send2 request data to FortiOS with bIgnoreVdom.
// If errors are encountered, it returns the error.
func (r *Request) Send2(retries int, ignvdom bool, headers *map[string]string) error {
	var err error
	var cookies *Cookies
	var login_method string
	token := r.Config.Auth.Token
	if headers != nil && len(*headers) != 0 {
		for hName, hValue := range *headers {
			if hValue != "" {
				r.HTTPRequest.Header.Set(hName, hValue)
			}
		}
	}
	if r.Config.Auth.Token == "" {
		token, cookies, err = r.LoginToken()
		if cookies != nil {
			r.HTTPRequest.Header.Set("X-CSRFTOKEN", cookies.CSRFToken)
			r.HTTPRequest.Header.Set("Cookie", cookies.Cookie)
		}

		login_method = "token"
		if err != nil {
			cookies, err = r.LoginSession()
			if err != nil {
				log.Printf("[ERROR] Failed to login: %v", err)
				return err
			}
			r.HTTPRequest.Header.Set("X-CSRFTOKEN", cookies.CSRFToken)
			r.HTTPRequest.Header.Set("Cookie", cookies.Cookie)
			login_method = "session"
		}
	}

	r.HTTPRequest.Header.Set("Content-Type", "application/json")
	r.HTTPRequest.Header.Set("accept", "application/json")
	if token != "" {
		r.HTTPRequest.Header.Set("Authorization", "Bearer "+token)
	}
	vdom := r.Config.Auth.Vdom
	if ignvdom == true {
		vdom = ""
	}
	u := r.buildURL(vdom, "")

	r.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		return err
	}

	retry := 0
	for {
		//Send
		rsp, errdo := r.Config.HTTPCon.Do(r.HTTPRequest)
		r.HTTPResponse = rsp
		if errdo != nil {
			if strings.Contains(errdo.Error(), "x509: ") {
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}

			if retry > retries {
				err = fmt.Errorf("lost connection to firewall with error: %v", filterapikey(errdo.Error()))
				break
			}
			time.Sleep(time.Second)
			log.Printf("Error found: %v, will resend again %s, %d", filterapikey(errdo.Error()), u, retry)

			retry++

		} else {
			break
		}
	}

	if login_method == "session" {
		errLogout := r.LogoutSession(cookies)
		if errLogout != nil {
			log.Printf("[WARNING] Issue occurs when logout session: %v", errLogout)
		}
	} else if login_method == "token" {
		errLogout := r.LogoutToken(token, cookies)
		if errLogout != nil {
			log.Printf("[WARNING] Issue occurs when logout token: %v", errLogout)
		}
	}

	return err
}

// Send3 request data to FortiOS with custom vdom.
// If errors are encountered, it returns the error.
func (r *Request) Send3(vdomparam string, headers *map[string]string) error {
	retries := 15

	var err error
	var cookies *Cookies
	var login_method string
	token := r.Config.Auth.Token

	if headers != nil && len(*headers) != 0 {
		for hName, hValue := range *headers {
			if hValue != "" {
				r.HTTPRequest.Header.Set(hName, hValue)
			}
		}
	}

	if r.Config.Auth.Token == "" {
		token, cookies, err = r.LoginToken()
		if cookies != nil {
			r.HTTPRequest.Header.Set("X-CSRFTOKEN", cookies.CSRFToken)
			r.HTTPRequest.Header.Set("Cookie", cookies.Cookie)
		}
		login_method = "token"

		if err != nil {
			cookies, err = r.LoginSession()
			if err != nil {
				log.Printf("[ERROR] Failed to login: %v", err)
				return err
			}
			r.HTTPRequest.Header.Set("X-CSRFTOKEN", cookies.CSRFToken)
			r.HTTPRequest.Header.Set("Cookie", cookies.Cookie)
			login_method = "session"
		}
	}
	r.HTTPRequest.Header.Set("Content-Type", "application/json")
	if token != "" {
		r.HTTPRequest.Header.Set("Authorization", "Bearer "+token)
	}
	vdom := vdomparam
	if vdom == "" {
		vdom = r.Config.Auth.Vdom
	}
	u := r.buildURL(vdom, "")
	r.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Logged in successfully - Send3 URL: %s, %+v", u, r.HTTPRequest.Header)
	retry := 0
	for {
		//Send
		rsp, errdo := r.Config.HTTPCon.Do(r.HTTPRequest)
		r.HTTPResponse = rsp
		if errdo != nil {
			if strings.Contains(errdo.Error(), "x509: ") {
				err = fmt.Errorf("[ERROR] Error found: %v", filterapikey(errdo.Error()))
				break
			}

			if retry > retries {
				err = fmt.Errorf("[ERROR] lost connection to firewall with error: %v", filterapikey(errdo.Error()))
				break
			}
			time.Sleep(time.Second)
			log.Printf("[ERROR] Error found: %v, will resend again %s, %d", filterapikey(errdo.Error()), u, retry)

			retry++

		} else {
			break
		}
	}

	if login_method == "session" {
		errLogout := r.LogoutSession(cookies)
		if errLogout != nil {
			log.Printf("[WARNING] Issue occurs when logout session: %v", errLogout)
		}
	} else if login_method == "token" {
		errLogout := r.LogoutToken(token, cookies)
		if errLogout != nil {
			log.Printf("[WARNING] Issue occurs when logout token: %v", errLogout)
		}
	}

	return err
}

// checkValid check whether given credential is valid.
// If errors are encountered, it returns the error.
func (r *Request) CheckValid() error {
	retries := 15

	var err error
	var cookies *Cookies
	token := r.Config.Auth.Token

	if r.Config.Auth.Token == "" {
		token, cookies, err = r.LoginToken()
		if cookies != nil {
			r.HTTPRequest.Header.Set("X-CSRFTOKEN", cookies.CSRFToken)
			r.HTTPRequest.Header.Set("Cookie", cookies.Cookie)
		}
		if err != nil {
			log.Printf("[DEBUG] LoginToken error, try loginSession: %v", err)
			cookies, err = r.LoginSession()
			if err != nil {
				log.Printf("[ERROR] Failed to login: %v", err)
				return err
			}
			r.HTTPRequest.Header.Set("X-CSRFTOKEN", cookies.CSRFToken)
			r.HTTPRequest.Header.Set("Cookie", cookies.Cookie)
			// logout session
			errLogout := r.LogoutSession(cookies)
			if errLogout != nil {
				log.Printf("[WARNING] Issue occurs when logout session: %v", errLogout)
			}
			return nil
		}

		// logout token
		errLogout := r.LogoutToken(token, cookies)
		if errLogout != nil {
			log.Printf("[WARNING] Issue occurs when logout token: %v", errLogout)
		}
		return nil
	}
	r.HTTPRequest.Header.Set("Content-Type", "application/json")
	u := r.buildURL("", "")
	r.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		return err
	}

	retry := 0
	for {
		//Send
		rsp, errdo := r.Config.HTTPCon.Do(r.HTTPRequest)
		r.HTTPResponse = rsp
		if errdo != nil {
			if strings.Contains(errdo.Error(), "x509: ") {
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}

			if retry > retries {
				err = fmt.Errorf("lost connection to firewall with error: %v", filterapikey(errdo.Error()))
				break
			}
			time.Sleep(time.Second)
			log.Printf("Error found: %v, will resend again %s, %d", filterapikey(errdo.Error()), u, retry)

			retry++

		} else {
			break
		}
	}

	return err
}

func filterapikey(v string) string { // !!! need check whether need this or not
	re, _ := regexp.Compile("access_token=.*?\"")
	res := re.ReplaceAllString(v, "access_token=***************\"")

	return res
}

func (r *Request) buildURL(vdom, spvar string) string {
	u := "https://"
	u += r.Config.FwTarget
	u += r.Path

	param_list := []string{}
	if vdom != "" {
		if vdom == "global" {
			param_list = append(param_list, "scope=global")
		} else {
			param_list = append(param_list, "vdom="+vdom)
		}
	}
	if spvar != "" {
		param_list = append(param_list, spvar)
	}

	if len(param_list) > 0 {
		u += "?" + strings.Join(param_list, "&")
	}

	return u
}

// SendWithSpecialParams sends request data to FortiOS with special URL paramaters.
// If errors are encountered, it returns the error.
func (r *Request) SendWithSpecialParams(s, vdomparam string) error {
	var err error
	var cookies *Cookies
	var login_method string
	token := r.Config.Auth.Token
	if r.Config.Auth.Token == "" {
		token, cookies, err = r.LoginToken()
		if cookies != nil {
			r.HTTPRequest.Header.Set("X-CSRFTOKEN", cookies.CSRFToken)
			r.HTTPRequest.Header.Set("Cookie", cookies.Cookie)
		}
		login_method = "token"
		if err != nil {
			cookies, err = r.LoginSession()
			if err != nil {
				log.Printf("[ERROR] Failed to login: %v", err)
				return err
			}
			r.HTTPRequest.Header.Set("X-CSRFTOKEN", cookies.CSRFToken)
			r.HTTPRequest.Header.Set("Cookie", cookies.Cookie)
			login_method = "session"
		}
	}
	r.HTTPRequest.Header.Set("Content-Type", "application/json")
	if token != "" {
		r.HTTPRequest.Header.Set("Authorization", "Bearer "+token)
	}
	vdom := vdomparam
	if vdom == "" {
		vdom = r.Config.Auth.Vdom
	}
	u := r.buildURL(vdom, s)

	r.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		return err
	}

	retry := 0
	for {
		//Send
		rsp, errdo := r.Config.HTTPCon.Do(r.HTTPRequest)
		r.HTTPResponse = rsp
		if errdo != nil {
			if strings.Contains(errdo.Error(), "x509: ") {
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}

			if retry > 15 {
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}
			time.Sleep(time.Duration(1) * time.Second)
			log.Printf("Error found: %v, will resend again %s, %d", filterapikey(errdo.Error()), u, retry)

			retry++

		} else {
			break
		}
	}

	if login_method == "session" {
		errLogout := r.LogoutSession(cookies)
		if errLogout != nil {
			log.Printf("[WARNING] Issue occurs when logout session: %v", errLogout)
		}
	} else if login_method == "token" {
		errLogout := r.LogoutToken(token, cookies)
		if errLogout != nil {
			log.Printf("[WARNING] Issue occurs when logout token: %v", errLogout)
		}
	}

	return err
}

// Login FortiOS using username and password in token mode, and return Cookies.
// If errors are encountered, it returns the error.
func (r *Request) LoginToken() (string, *Cookies, error) {
	var err error
	var token string
	var cookies *Cookies
	data := make(map[string]interface{})
	data["username"] = r.Config.Auth.Username
	data["secretkey"] = r.Config.Auth.Password
	data["password"] = r.Config.Auth.Password
	data["request_key"] = true
	data["ack_pre_disclaimer"] = true
	data["ack_post_disclaimer"] = true

	log.Printf("[DEBUG] Trying LoginToken")
	locJSON, err := json.Marshal(data)
	if err != nil {
		log.Printf("[ERROR] Encoding body data failed.")
		return token, cookies, err
	}

	bodyBytes := bytes.NewBuffer(locJSON)

	req, _ := http.NewRequest("POST", "", bodyBytes)
	req.Header.Set("Content-Type", "application/json")
	u := "https://"
	u += r.Config.FwTarget
	u += "/api/v2/authentication"
	req.URL, err = url.Parse(u)
	if err != nil {
		err = fmt.Errorf("Could not parse URL: %s", err)
		return token, cookies, err
	}

	rsp, err := r.Config.HTTPCon.Do(req)
	if err != nil {
		if strings.Contains(err.Error(), "x509: ") {
			err = fmt.Errorf("HTTP request error: %v", err)
			return token, cookies, err
		}
	}

	log.Printf("[DEBUG] LoginToken HTTP response: %+v", rsp)

	if rsp == nil {
		err = fmt.Errorf("Host is unreachable. HTTP response is nil.")
		return token, cookies, err
	}

	if rsp.Header == nil {
		err = fmt.Errorf("HTTP response header is nil.")
		return token, cookies, err
	}

	if rsp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Login failed with status code: %d", rsp.StatusCode)
		return token, cookies, err
	}

	body, err := ioutil.ReadAll(rsp.Body)
	rsp.Body.Close()

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body, %s", err)
		return token, cookies, err
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	if sc, ok := result["status"]; ok && (int(sc.(float64)) == 5 || int(sc.(float64)) == 6) {
		if session_key, ok := result["session_key"]; ok {
			token = session_key.(string)
		}
		if token == "" {
			err = fmt.Errorf("Login failed: cannot get token with response status %s.", result["status"])
		}
	} else if sc, ok := result["status_code"]; ok && (int(sc.(float64)) == 5 || int(sc.(float64)) == 6) {
		if session_key, ok := result["session_key"]; ok {
			token = session_key.(string)
		}
		if token == "" {
			err = fmt.Errorf("Login failed: cannot get token with response status code %s.", result["status_code"])
		}
	}

	if token == "" {
		cookies, err = r.ParseCookies(rsp)
		if err != nil {
			log.Printf("[ERROR] Parsing cookies failed, %v, %+v", err, rsp.Header)
			err = fmt.Errorf("LoginToken failed: %+v.", result)
		}
	}

	return token, cookies, err
}

// Login FortiOS using username and password in session mode, and return Cookies.
// If errors are encountered, it returns the error.
func (r *Request) LoginSession() (*Cookies, error) {
	var err error

	data := "username="
	data += r.Config.Auth.Username
	data += "&secretkey="
	data += r.Config.Auth.Password
	data += "&ajax=1"

	log.Printf("[DEBUG] Trying LoginSession")
	bodyBytes := bytes.NewBufferString(data)

	req, _ := http.NewRequest("POST", "", bodyBytes)
	req.Header.Set("Content-Type", "application/json")
	u := "https://"
	u += r.Config.FwTarget
	u += "/logincheck"
	req.URL, err = url.Parse(u)
	if err != nil {
		err = fmt.Errorf("Could not parse URL: %s", err)
		return nil, err
	}

	rsp, err := r.Config.HTTPCon.Do(req)
	if err != nil {
		if strings.Contains(err.Error(), "x509: ") {
			err = fmt.Errorf("HTTP request error: %v", err)
			return nil, err
		}
	}

	if rsp == nil {
		err = fmt.Errorf("Host is unreachable. HTTP response is nil.")
		return nil, err
	}

	if rsp.Header == nil {
		err = fmt.Errorf("HTTP response header is nil.")
		return nil, err
	}

	return r.ParseCookies(rsp)
}

func (r *Request) ParseCookies(rsp *http.Response) (*Cookies, error) {
	csrfToken := ""
	cookie := ""
	err := error(nil)
	if setCookie, ok := rsp.Header["Set-Cookie"]; ok {
		for _, item := range setCookie {
			// reg := regexp.MustCompile(`ccsrf[_]{0,1}token.*?=(.*)`)
			reg := regexp.MustCompile(`^(ccsrf_token|ccsrftoken)[^=]*=(?:"([^"]+)"|([^;]+))`)
			full_cookie := strings.Split(item, ";")[0]
			match := reg.FindStringSubmatch(full_cookie)
			if match != nil {
				if len(match) > 3 && match[2] == "" {
					csrfToken = match[3]
				} else {
					csrfToken = match[2]
				}
			}

			if cookie == "" {
				cookie = full_cookie
			} else {
				cookie = fmt.Sprintf("%s; %s", cookie, full_cookie)
			}
		}

		switch csrfToken {
		case "":
			err = fmt.Errorf("Could not get CSRF Token from HTTP request response.")
			return nil, err
		case "0%260":
			err = fmt.Errorf("Username or password not valid.")
			return nil, err
		}
		if cookie == "" {
			err = fmt.Errorf("Could not get cookie from HTTP request response.")
			return nil, err
		}
	} else {
		err = fmt.Errorf("Logincheck response do not contains Set-Cookie [%+v]", rsp.Header)
		return nil, err
	}

	ck := &Cookies{
		CSRFToken: csrfToken,
		Cookie:    cookie,
	}
	return ck, nil
}

// Logout current session.
// If errors are encountered, it returns the error.
func (r *Request) LogoutSession(cookies *Cookies) error {
	var err error

	if cookies.CSRFToken == "" {
		log.Printf("[INFO] CSRF Token is empty, no need to logout.")
		return nil
	}

	req, _ := http.NewRequest("POST", "", nil)
	req.Header.Set("X-CSRFTOKEN", cookies.CSRFToken)
	req.Header.Set("Cookie", cookies.Cookie)
	req.Header.Set("Content-Type", "application/json")
	u := "https://"
	u += r.Config.FwTarget
	u += "/logout"
	req.URL, err = url.Parse(u)
	if err != nil {
		err = fmt.Errorf("Could not parse URL: %s", err)
		return err
	}

	resp, err := r.Config.HTTPCon.Do(req)
	if err != nil {
		if strings.Contains(err.Error(), "x509: ") {
			err = fmt.Errorf("HTTP request error: %v", err)
			return err
		}
	}

	log.Printf("[DEBUG] Logged out session successfully: %d", resp.StatusCode)

	return nil
}

// Logout current token based authentication.
// If errors are encountered, it returns the error.
func (r *Request) LogoutToken(token string, cookies *Cookies) error {
	var err error

	data := "session_key="
	data += token

	bodyBytes := bytes.NewBufferString(data)

	req, _ := http.NewRequest("DELETE", "", bodyBytes)
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token) // !!! check whether logout or not
	}

	if cookies != nil {
		req.Header.Set("X-CSRFTOKEN", cookies.CSRFToken)
		req.Header.Set("Cookie", cookies.Cookie)
	}

	u := "https://"
	u += r.Config.FwTarget
	u += "/api/v2/authentication"
	req.URL, err = url.Parse(u)
	if err != nil {
		err = fmt.Errorf("Could not parse URL: %s", err)
		return err
	}

	resp, err := r.Config.HTTPCon.Do(req)
	if err != nil {
		if strings.Contains(err.Error(), "x509: ") {
			err = fmt.Errorf("HTTP request error: %v", err)
			return err
		}
	}

	log.Printf("[DEBUG] Logged out token successfully: %d, req header %+v, resp: %+v", resp.StatusCode, req.Header, resp)
	return nil
}
