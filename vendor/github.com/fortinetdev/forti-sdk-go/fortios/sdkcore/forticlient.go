package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/fortinetdev/forti-sdk-go/fortios/auth"
	"github.com/fortinetdev/forti-sdk-go/fortios/config"
	"github.com/fortinetdev/forti-sdk-go/fortios/request"
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
	Fv      string
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

func escapeURLString(v string) string { // doesn't support "<>()"'#"
	return strings.Replace(url.QueryEscape(v), "+", "%20", -1)
}

// NewClient initializes a new global plugin client
// It returns the created client object
func NewClient(auth *auth.Auth, client *http.Client) (*FortiSDKClient, error) {
	c := &FortiSDKClient{}

	c.Config.Auth = auth
	c.Config.HTTPCon = client
	c.Config.FwTarget = auth.Hostname

	vsave := client.Timeout
	client.Timeout = time.Second * 30
	v, err := c.GetDeviceVersion()
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	client.Timeout = vsave
	c.Fv = v

	return c, nil
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
	err = req.Send2(2, true)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request, %s", err)
		return "", err
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body, %s", err)
		return "", err
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		if result["status"] == nil {
			err = fmt.Errorf("no status in response")
			return "", err
		}

		if result["status"] != "success" {
			err = fmt.Errorf("wrong status - %v", result["status"])
			return "", err
		}

		if result["version"] == nil {
			err = fmt.Errorf("no version in response")
			return "", err
		}

		return result["version"].(string), err
	}

	return "", err
}

func fortiAPIHttpStatus404Checking(result map[string]interface{}) (b404 bool) {
	b404 = false

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

				if result["cli_error"] != nil {
					err = fmt.Errorf(err.Error() + "\nCli response: \n%v", result["cli_error"])
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
