package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

// JSONSystemPasswordPolicy contains the General parameters for Create and Update API function
type JSONSystemPasswordPolicy struct {
	Status             string `json:"status"`
	ApplyTo            string `json:"apply-to"`
	MinimunLength      string `json:"minimum-length"`
	MinLowerCaseLetter string `json:"min-lower-case-letter"`
	MinUpperCaseLetter string `json:"min-upper-case-letter"`
	MinNonAlphanumeric string `json:"min-non-alphanumeric"`
	MinNumber          string `json:"min-number"`
	Change4Characters  string `json:"change-4-characters"`
	ExpireStatus       string `json:"expire-status"`
	ExpireDay          string `json:"expire-day"`
	ReusePassword      string `json:"reuse-password"`
}

// JSONCreateSystemPasswordPolicyOutput contains the output results for Update API function
type JSONCreateSystemPasswordPolicyOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateSystemPasswordPolicyOutput contains the output results for Update API function
type JSONUpdateSystemPasswordPolicyOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateSystemPasswordPolicy API operation for FortiOS
func (c *FortiSDKClient) CreateSystemPasswordPolicy(params *JSONSystemPasswordPolicy) (output *JSONCreateSystemPasswordPolicyOutput, err error) {
	return
}

// UpdateSystemPasswordPolicy API operation for FortiOS
func (c *FortiSDKClient) UpdateSystemPasswordPolicy(params *JSONSystemPasswordPolicy, mkey string) (output *JSONUpdateSystemPasswordPolicyOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/password-policy"
	// path += "/" + mkey
	output = &JSONUpdateSystemPasswordPolicyOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}
	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request")
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body")
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
					err = fmt.Errorf("%s and http_status no is %.0f", err, result["http_status"])
				} else {
					err = fmt.Errorf("%s and and http_status no is not found", err)
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

// DeleteSystemPasswordPolicy API operation for FortiOS
func (c *FortiSDKClient) DeleteSystemPasswordPolicy(mkey string) (err error) {
	return
}

// ReadSystemPasswordPolicy API operation for FortiOS
func (c *FortiSDKClient) ReadSystemPasswordPolicy(mkey string) (output *JSONSystemPasswordPolicy, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/password-policy"
	output = &JSONSystemPasswordPolicy{}

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request")
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body")
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
				err = fmt.Errorf("%s and http_status no is %.0f", err, result["http_status"])
			} else {
				err = fmt.Errorf("%s and and http_status no is not found", err)
			}

			return
		}

		mapTmp := (result["results"].(map[string]interface{}))

		if mapTmp == nil {
			err = fmt.Errorf("cannot get the results from the response")
			return
		}
		if mapTmp["status"] != nil {
			output.Status = mapTmp["status"].(string)
		}
		if mapTmp["apply-to"] != nil {
			output.ApplyTo = mapTmp["apply-to"].(string)
		}
		if mapTmp["minimum-length"] != nil {
			output.MinimunLength = strconv.Itoa(int(mapTmp["minimum-length"].(float64)))
		}
		if mapTmp["min-lower-case-letter"] != nil {
			output.MinLowerCaseLetter = strconv.Itoa(int(mapTmp["min-lower-case-letter"].(float64)))
		}
		if mapTmp["min-upper-case-letter"] != nil {
			output.MinUpperCaseLetter = strconv.Itoa(int(mapTmp["min-upper-case-letter"].(float64)))
		}
		if mapTmp["min-non-alphanumeric"] != nil {
			output.MinNonAlphanumeric = strconv.Itoa(int(mapTmp["min-non-alphanumeric"].(float64)))
		}
		if mapTmp["min-number"] != nil {
			output.MinNumber = strconv.Itoa(int(mapTmp["min-number"].(float64)))
		}
		if mapTmp["change-4-characters"] != nil {
			output.Change4Characters = mapTmp["change-4-characters"].(string)
		}
		if mapTmp["expire-status"] != nil {
			output.ExpireStatus = mapTmp["expire-status"].(string)
		}
		if mapTmp["expire-day"] != nil {
			output.ExpireDay = strconv.Itoa(int(mapTmp["expire-day"].(float64)))
		}
		if mapTmp["reuse-password"] != nil {
			output.ReusePassword = mapTmp["reuse-password"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
