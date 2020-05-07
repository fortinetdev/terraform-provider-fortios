package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/fgtdev/fortios-sdk-go/util"
)

// JSONLogSyslogSetting contains the parameters for Create and Update API function
type JSONLogSyslogSetting struct {
	Status   string `json:"status"`
	Server   string `json:"server,omitempty"`
	Mode     string `json:"mode,omitempty"`
	Port     string `json:"port,omitempty"`
	Facility string `json:"facility,omitempty"`
	SourceIP string `json:"source-ip,omitempty"`
	Format   string `json:"format,omitempty"`
}

// JSONCreateLogSyslogSettingOutput contains the output results for Create API function
type JSONCreateLogSyslogSettingOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateLogSyslogSettingOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateLogSyslogSettingOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateLogSyslogSetting API operation for FortiOS creates a new remote Syslog logging server.
// Returns the index value of the remote Syslog logging server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - syslogd setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateLogSyslogSetting(params *JSONLogSyslogSetting) (output *JSONCreateLogSyslogSettingOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/log.syslogd/setting"
	output = &JSONCreateLogSyslogSettingOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
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

// UpdateLogSyslogSetting API operation for FortiOS updates the specified remote Syslog logging server.
// Returns the index value of the remote Syslog logging server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - syslogd setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogSetting(params *JSONLogSyslogSetting, mkey string) (output *JSONUpdateLogSyslogSettingOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd/setting"
	// path += "/" + mkey
	output = &JSONUpdateLogSyslogSettingOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("FOS-fortios params: %v", params)
	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
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

// DeleteLogSyslogSetting API operation for FortiOS deletes the specified remote Syslog logging server.
// Returns error for service API and SDK errors.
// See the log - syslogd setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogSetting(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/log.syslogd/setting"
	// path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
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

// ReadLogSyslogSetting API operation for FortiOS gets the remote Syslog logging server
// with the specified index value.
// Returns the requested remote Syslog logging server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - syslogd setting chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogSetting(mkey string) (output *JSONLogSyslogSetting, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd/setting"
	// path += "/" + mkey

	output = &JSONLogSyslogSetting{}

	req := c.NewRequest(HTTPMethod, path, nil, nil)
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

		mapTmp := (result["results"].(map[string]interface{}))

		if mapTmp == nil {
			err = fmt.Errorf("cannot get the results from the response")
			return
		}

		if mapTmp["status"] != nil {
			output.Status = mapTmp["status"].(string)
		}
		if mapTmp["server"] != nil {
			output.Server = mapTmp["server"].(string)
		}
		if mapTmp["mode"] != nil {
			output.Mode = mapTmp["mode"].(string)
		}
		if mapTmp["port"] != nil {
			output.Port = strconv.Itoa(int(mapTmp["port"].(float64)))
		}
		if mapTmp["facility"] != nil {
			output.Facility = mapTmp["facility"].(string)
		}
		if mapTmp["source-ip"] != nil {
			output.SourceIP = mapTmp["source-ip"].(string)
		}
		if mapTmp["format"] != nil {
			output.Format = mapTmp["format"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
