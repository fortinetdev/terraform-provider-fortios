package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

// JSONSystemSettingGlobal contains the parameters for Create and Update API function
type JSONSystemSettingGlobal struct {
	Admintimeout string `json:"admintimeout,omitempty"`
	Timezone     string `json:"timezone,omitempty"`
	Hostname     string `json:"hostname"`
	AdminSport   string `json:"admin-sport,omitempty"`
	AdminSSHPort string `json:"admin-ssh-port,omitempty"`
	AdminScp     string `json:"admin-scp,omitempty"`
}

// JSONCreateSystemSettingGlobalOutput contains the output results for Create API function
type JSONCreateSystemSettingGlobalOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateSystemSettingGlobalOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateSystemSettingGlobalOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateSystemSettingGlobal API operation for FortiOS
func (c *FortiSDKClient) CreateSystemSettingGlobal(params *JSONSystemSettingGlobal) (output *JSONCreateSystemSettingGlobalOutput, err error) {
	// HTTPMethod := "POST"
	// path := "/api/v2/cmdb/system/global"
	// output = &JSONCreateSystemSettingGlobalOutput{}
	// locJSON, err := json.Marshal(params)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// bytes := bytes.NewBuffer(locJSON)
	// req := c.NewRequest(HTTPMethod, path, nil, bytes)
	// err = req.Send()

	// body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	// req.HTTPResponse.Body.Close() //#

	// var result map[string]interface{}
	// json.Unmarshal([]byte(string(body)), &result)

	// if result != nil {
	// 	if result["vdom"] != nil {
	// 		output.Vdom = result["vdom"].(string)
	// 	}
	// 	if result["mkey"] != nil {
	// 		output.Mkey = result["mkey"].(string)
	// 	}
	// 	if result["status"] != nil {
	// 		if result["status"] != "success" {
	// 			err = fmt.Errorf("cannot get the right response")
	// 			return
	// 		}
	// 		output.Status = result["status"].(string)
	// 	} else {
	// 		err = fmt.Errorf("cannot get the right response")
	// 		return
	// 	}
	// 	if result["http_status"] != nil {
	// 		output.HTTPStatus = result["http_status"].(float64)
	// 	}
	// } else {
	// 	err = fmt.Errorf("cannot get the right response")
	// 	return
	// }

	return
}

// UpdateSystemSettingGlobal API operation for FortiOS configures global settings that affect FortiGate systems and configurations.
// Returns the execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSettingGlobal(params *JSONSystemSettingGlobal, mkey string) (output *JSONUpdateSystemSettingGlobalOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/global"
	// path += "/" + mkey
	output = &JSONUpdateSystemSettingGlobalOutput{}
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
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}

		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(string)
		}
	}

	return
}

// DeleteSystemSettingGlobal API operation for FortiOS
func (c *FortiSDKClient) DeleteSystemSettingGlobal(mkey string) (err error) {
	// HTTPMethod := "DELETE"
	// path := "/api/v2/cmdb/system/global"
	// // path += "/" + mkey

	// req := c.NewRequest(HTTPMethod, path, nil, nil)
	// err = req.Send()

	// body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	// req.HTTPResponse.Body.Close() //#

	// log.Printf("FOS-fortios response: %s", string(body))

	// var result map[string]interface{}
	// json.Unmarshal([]byte(string(body)), &result)

	// if result != nil {
	// 	if result["status"] == nil {
	// 		err = fmt.Errorf("cannot get the right response")
	// 		return
	// 	}
	//
	//	if result["status"] != "success" {
	//		err = fmt.Errorf("cannot get the right response")
	//		return
	//	}
	//
	// } else {
	// 	err = fmt.Errorf("cannot get the right response")
	// 	return
	// }

	return
}

// ReadSystemSettingGlobal API operation for FortiOS gets the global settings that affect FortiGate systems and configurations.
// Returns the requested global settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSettingGlobal(mkey string) (output *JSONSystemSettingGlobal, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/global"
	// path += "/" + mkey

	output = &JSONSystemSettingGlobal{}

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios reading response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	if fortiAPIHttpStatus404Checking(result) == true {
		output = nil
		return
	}

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		mapTmp := (result["results"].(map[string]interface{}))

		if mapTmp == nil {
			err = fmt.Errorf("cannot get the results from the response")
			return
		}

		if mapTmp["admintimeout"] != nil {
			output.Admintimeout = strconv.Itoa(int(mapTmp["admintimeout"].(float64)))
		}
		if mapTmp["timezone"] != nil {
			output.Timezone = mapTmp["timezone"].(string)
		}
		if mapTmp["hostname"] != nil {
			output.Hostname = mapTmp["hostname"].(string)
		}
		if mapTmp["admin-sport"] != nil {
			output.AdminSport = strconv.Itoa(int(mapTmp["admin-sport"].(float64)))
		}
		if mapTmp["admin-ssh-port"] != nil {
			output.AdminSSHPort = strconv.Itoa(int(mapTmp["admin-ssh-port"].(float64)))
		}
		if mapTmp["admin-scp"] != nil {
			output.AdminScp = mapTmp["admin-scp"].(string)
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
