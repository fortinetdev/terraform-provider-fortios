package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/fgtdev/fortios-sdk-go/auth"
	"github.com/fgtdev/fortios-sdk-go/config"
	"github.com/fgtdev/fortios-sdk-go/request"
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
	Config  config.Config
	Retries int
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

// EscapeURLString escapes the string so it can be safely placed inside a URL query
func EscapeURLString(v string) string { // doesn't support "<>()"'#"
	return strings.Replace(url.QueryEscape(v), "+", "%20", -1)
}

// NewClient initializes a new global plugin client
// It returns the created client object
func NewClient(auth *auth.Auth, client *http.Client) *FortiSDKClient {
	c := &FortiSDKClient{}

	c.Config.Auth = auth
	c.Config.HTTPCon = client
	c.Config.FwTarget = auth.Hostname

	return c
}

// NewRequest creates the request to FortiOS for the client
// and return it to the client
func (c *FortiSDKClient) NewRequest(method string, path string, params interface{}, data *bytes.Buffer) *request.Request {
	return request.New(c.Config, method, path, params, data)
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

func fortiAPIHttpStatus404Checking(result map[string]interface{}) (b404 bool) {
	b404 = false;

	if result != nil {
		if result["http_status"] != nil && result["http_status"] == 404.0 {
				b404 = true
				return
		}
	}

	return
}

func fortiAPIErrorFormat(result map[string]interface{}, body string) (err error) {
	if result != nil {
		if result["status"] != nil {
			if result["status"] == "success" {
				err = nil
				return
			}

			if result["http_status"] != nil {
				// 200	OK: Request returns successful
				if result["http_status"] == 400.0 {
					err = fmt.Errorf("Bad Request - Request cannot be processed by the API (%.0f)", result["http_status"])
				} else if result["http_status"] == 401.0 {
					err = fmt.Errorf("Not Authorized - Request without successful login session (%.0f)", result["http_status"])
				} else if result["http_status"] == 403.0 {
					err = fmt.Errorf("Forbidden - Request is missing CSRF token or administrator is missing access profile permissions (%.0f)", result["http_status"])
				} else if result["http_status"] == 404.0 {
					err = fmt.Errorf("Resource Not Found - Unable to find the specified resource (%.0f)", result["http_status"])
				} else if result["http_status"] == 405.0 {
					err = fmt.Errorf("Method Not Allowed - Specified HTTP method is not allowed for this resource (%.0f)", result["http_status"])
				} else if result["http_status"] == 413.0 {
					err = fmt.Errorf("Request Entity Too Large - Request cannot be processed due to large entity (%.0f)", result["http_status"])
				} else if result["http_status"] == 424.0 {
					err = fmt.Errorf("Failed Dependency - Fail dependency can be duplicate resource, missing required parameter, missing required attribute, invalid attribute value (%.0f)", result["http_status"])
				} else if result["http_status"] == 429.0 {
					err = fmt.Errorf("Access temporarily blocked - Maximum failed authentications reached. The offended source is temporarily blocked for certain amount of time (%.0f)", result["http_status"])
				} else if result["http_status"] == 500.0 {
					err = fmt.Errorf("Internal Server Error - Internal error when processing the request (%.0f)", result["http_status"])
				} else {
					err = fmt.Errorf("Unknow Error (%.0f)", result["http_status"])
				}

				return
			}

			err = fmt.Errorf("\n%v", body)
			return

		}
		err = fmt.Errorf("\n%v", body)
		return
	}

	// Authorization Required, etc. | Attention: scalable here
	err = fmt.Errorf("\n%v", body)
	return
}

//Build input data by sdk
