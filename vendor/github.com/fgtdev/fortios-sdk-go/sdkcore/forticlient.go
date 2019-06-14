package forticlient

import (
	"net/http"
	"github.com/fgtdev/fortios-sdk-go/config"
	"github.com/fgtdev/fortios-sdk-go/auth"
	"github.com/fgtdev/fortios-sdk-go/request"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	// "strconv"
)

// MultValue describes the nested structure in the results
type MultValue struct {
	Name string `json:"name"`
}

// MultValues describes the nested structure in the results
type MultValues []MultValue

// FortiSDKClient describes the global FortiOS plugin client instance
type FortiSDKClient struct {
	Config   config.Config
	Retries  int
}

// ExtractString extracts strings from result and put them into a string array,
// and return the string array
func ExtractString(members []MultValue) []string {
	vs := make([]string, 0, len(members))
	for _, v := range members {
		c := v.Name
		vs = append(vs, c)
	}
	return vs
}


// NewClient initializes a new global plugin client
// It returns the created client object
func NewClient(auth *auth.Auth, client *http.Client) *FortiSDKClient{
	c := &FortiSDKClient{ }

	c.Config.Auth = auth
	c.Config.HTTPCon = client
	c.Config.FwTarget = auth.Hostname

	return c
}

// NewRequest creates the request to FortiOS for the client
// and return it to the client
func (c *FortiSDKClient) NewRequest(method string, path string, params interface{}, data *bytes.Buffer) *request.Request {
	return request.New(c.Config, method, path, params, data);
}

// GetDeviceVersion gets the version of FortiOS
// It returns version as string
func (c *FortiSDKClient) GetDeviceVersion() (version string, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/global"

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("FOS-fortios reading response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return "", err
		}

		if result["status"] != "success" {
			err = fmt.Errorf("cannot get the right response")
			return "", err
		}

		if result["version"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return "", err
		}

		return result["version"].(string), err
	}

	err = fmt.Errorf("cannot get the right response")
	return "", err
}

//Build input data by sdk
