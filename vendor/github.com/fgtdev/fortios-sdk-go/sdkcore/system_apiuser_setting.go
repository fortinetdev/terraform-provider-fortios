package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fgtdev/fortios-sdk-go/util"
)

// JSONSystemAPIUserSetting contains the parameters for Create and Update API function
type JSONSystemAPIUserSetting struct {
	Name       string            `json:"name"`
	Accprofile string            `json:"accprofile"`
	Vdom       MultValues        `json:"vdom"`
	Trusthost  APIUserMultValues `json:"trusthost"`
	Comments   string            `json:"comments"`
}

// JSONCreateSystemAPIUserSettingOutput contains the output results for Create API function
type JSONCreateSystemAPIUserSettingOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateSystemAPIUserSettingOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateSystemAPIUserSettingOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// APIUserMultValue contains the output results for Read API function
type APIUserMultValue struct {
	TypeF         string `json:"type"`
	Ipv4TrustHost string `json:"ipv4-trusthost"`
}

// APIUserMultValues contains the output results for Read API function
type APIUserMultValues []APIUserMultValue

// CreateSystemAPIUserSetting API operation for FortiOS creates a new API user.
// Returns the index value of the API user and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - api-user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAPIUserSetting(params *JSONSystemAPIUserSetting) (output *JSONCreateSystemAPIUserSettingOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/api-user"
	output = &JSONCreateSystemAPIUserSettingOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	e, req := c.NewRequest(HTTPMethod, path, nil, bytes)
	if e != nil {
		err = fmt.Errorf("new request error %s", e)
		return
	}
	
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}
		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(string)
		}
		if result["status"] != nil {
			if result["status"] != "success" {
				if result["error"] != nil {
					err = fmt.Errorf("status is %s and error no is %.0f", result["status"], result["error"])
				} else {
					err = fmt.Errorf("status is %s and error no is not found", result["status"])
				}

				if result["http_status"] != nil {
					err = fmt.Errorf("%s, details: %s", err, util.HttpStatus2Str(int(result["http_status"].(float64))))
				} else {
					err = fmt.Errorf("%s, and http_status no is not found", err)
				}

				return
			}
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get status from the response")
			return
		}
		if result["http_status"] != nil {
			output.HTTPStatus = result["http_status"].(float64)
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

// UpdateSystemAPIUserSetting API operation for FortiOS updates the specified API user.
// Returns the index value of the API user and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - api-user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAPIUserSetting(params *JSONSystemAPIUserSetting, mkey string) (output *JSONUpdateSystemAPIUserSettingOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/api-user"
	path += "/" + EscapeURLString(mkey)
	output = &JSONUpdateSystemAPIUserSettingOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	e, req := c.NewRequest(HTTPMethod, path, nil, bytes)
	if e != nil {
		err = fmt.Errorf("new request error %s", e)
		return
	}
	
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}
		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(string)
		}
		if result["status"] != nil {
			if result["status"] != "success" {
				if result["error"] != nil {
					err = fmt.Errorf("status is %s and error no is %.0f", result["status"], result["error"])
				} else {
					err = fmt.Errorf("status is %s and error no is not found", result["status"])
				}

				if result["http_status"] != nil {
					err = fmt.Errorf("%s, details: %s", err, util.HttpStatus2Str(int(result["http_status"].(float64))))
				} else {
					err = fmt.Errorf("%s, and http_status no is not found", err)
				}

				return
			}
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get status from the response")
			return
		}
		if result["http_status"] != nil {
			output.HTTPStatus = result["http_status"].(float64)
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

// DeleteSystemAPIUserSetting API operation for FortiOS deletes the specified API user.
// Returns error for service API and SDK errors.
// See the system - api-user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAPIUserSetting(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/api-user"
	path += "/" + EscapeURLString(mkey)

	e, req := c.NewRequest(HTTPMethod, path, nil, nil)
	if e != nil {
		err = fmt.Errorf("new request error %s", e)
		return
	}
	
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get status from the response")
			return
		}

		if result["status"] != "success" {
			if result["error"] != nil {
				err = fmt.Errorf("status is %s and error no is %.0f", result["status"], result["error"])
			} else {
				err = fmt.Errorf("status is %s and error no is not found", result["status"])
			}

			if result["http_status"] != nil {
				err = fmt.Errorf("%s, details: %s", err, util.HttpStatus2Str(int(result["http_status"].(float64))))
			} else {
				err = fmt.Errorf("%s, and http_status no is not found", err)
			}

			return
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

// ReadSystemAPIUserSetting API operation for FortiOS gets the API user
// with the specified index value.
// Returns the requested API user value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - api-user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAPIUserSetting(mkey string) (output *JSONSystemAPIUserSetting, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/api-user"
	path += "/" + EscapeURLString(mkey)

	output = &JSONSystemAPIUserSetting{}

	e, req := c.NewRequest(HTTPMethod, path, nil, nil)
	if e != nil {
		err = fmt.Errorf("new request error %s", e)
		return
	}
	
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios reading response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["http_status"] == nil {
			err = fmt.Errorf("cannot get http_status from the response")
			return
		}

		if result["http_status"] == 404.0 {
			output = nil
			return
		}

		if result["status"] == nil {
			err = fmt.Errorf("cannot get status from the response")
			return
		}

		if result["status"] != "success" {
			if result["error"] != nil {
				err = fmt.Errorf("status is %s and error no is %.0f", result["status"], result["error"])
			} else {
				err = fmt.Errorf("status is %s and error no is not found", result["status"])
			}

			if result["http_status"] != nil {
				err = fmt.Errorf("%s, details: %s", err, util.HttpStatus2Str(int(result["http_status"].(float64))))
			} else {
				err = fmt.Errorf("%s, and http_status no is not found", err)
			}

			return
		}

		mapTmp := (result["results"].([]interface{}))[0].(map[string]interface{})

		if mapTmp == nil {
			err = fmt.Errorf("cannot get the results from the response")
			return
		}

		if mapTmp["name"] != nil {
			output.Name = mapTmp["name"].(string)
		}
		if mapTmp["accprofile"] != nil {
			output.Accprofile = mapTmp["accprofile"].(string)
		}

		if mapTmp["vdom"] != nil {
			member := mapTmp["vdom"].([]interface{})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					MultValue{
						Name: c["name"].(string),
					})
			}
			output.Vdom = members
		}

		/*
			if mapTmp["vdom"] != nil {

				member := mapTmp["vdom"].([]interface{})

				output.Vdom = "root"

				for _, v := range member {
					c := v.(map[string]interface{})

					output.Vdom = c["name"].(string)

					// the break here because of the vdom in response becomes array, but vdom in request is string
					break
				}
			}*/

		if mapTmp["trusthost"] != nil {
			member := mapTmp["trusthost"].([]interface{})

			var members []APIUserMultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					APIUserMultValue{
						TypeF:         c["type"].(string),
						Ipv4TrustHost: c["ipv4-trusthost"].(string),
					})
			}
			output.Trusthost = members
		}
		if mapTmp["comments"] != nil {
			output.Comments = mapTmp["comments"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
	}

	return
}
