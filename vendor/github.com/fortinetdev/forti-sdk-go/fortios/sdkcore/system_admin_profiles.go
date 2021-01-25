package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// JSONSystemAdminProfiles contains the parameters for Create and Update API function
type JSONSystemAdminProfiles struct {
	Name                 string `json:"name"`
	Scope                string `json:"scope,omitempty"`
	Comments             string `json:"comments"`
	Secfabgrp            string `json:"secfabgrp,omitempty"`
	Ftviewgrp            string `json:"ftviewgrp,omitempty"`
	Authgrp              string `json:"authgrp,omitempty"`
	Sysgrp               string `json:"sysgrp,omitempty"`
	Netgrp               string `json:"netgrp,omitempty"`
	Loggrp               string `json:"loggrp,omitempty"`
	Fwgrp                string `json:"fwgrp,omitempty"`
	Vpngrp               string `json:"vpngrp,omitempty"`
	Utmgrp               string `json:"utmgrp,omitempty"`
	Wanoptgrp            string `json:"wanoptgrp,omitempty"`
	Wifi                 string `json:"wifi,omitempty"`
	AdmintimeoutOverride string `json:"admintimeout-override,omitempty"`
}

// JSONCreateSystemAdminProfilesOutput contains the output results for Create API function
type JSONCreateSystemAdminProfilesOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateSystemAdminProfilesOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateSystemAdminProfilesOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateSystemAdminProfiles API operation for FortiOS creates a new access profile
// Returns the index value of the access profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - accprofile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAdminProfiles(params *JSONSystemAdminProfiles) (output *JSONCreateSystemAdminProfilesOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/accprofile"
	output = &JSONCreateSystemAdminProfilesOutput{}
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

// UpdateSystemAdminProfiles API operation for FortiOS updates the specified access profile
// Returns the index value of the access profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - accprofile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdminProfiles(params *JSONSystemAdminProfiles, mkey string) (output *JSONUpdateSystemAdminProfilesOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/accprofile"
	path += "/" + EscapeURLString(mkey)
	output = &JSONUpdateSystemAdminProfilesOutput{}
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

// DeleteSystemAdminProfiles API operation for FortiOS deletes the specified access profile
// Returns error for service API and SDK errors.
// See the system - accprofile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdminProfiles(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/accprofile"
	path += "/" + EscapeURLString(mkey)

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
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	return
}

// ReadSystemAdminProfiles API operation for FortiOS gets the access profile
// with the specified index value.
// Returns the requested access profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - accprofile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAdminProfiles(mkey string) (output *JSONSystemAdminProfiles, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/accprofile"
	path += "/" + EscapeURLString(mkey)

	output = &JSONSystemAdminProfiles{}

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
		mapTmp := (result["results"].([]interface{}))[0].(map[string]interface{})

		if mapTmp == nil {
			err = fmt.Errorf("cannot get the results from the response")
			return
		}

		if mapTmp["name"] != nil {
			output.Name = mapTmp["name"].(string)
		}
		if mapTmp["scope"] != nil {
			output.Scope = mapTmp["scope"].(string)
		}
		if mapTmp["comments"] != nil {
			output.Comments = mapTmp["comments"].(string)
		}
		if mapTmp["secfabgrp"] != nil {
			output.Secfabgrp = mapTmp["secfabgrp"].(string)
		}
		if mapTmp["ftviewgrp"] != nil {
			output.Ftviewgrp = mapTmp["ftviewgrp"].(string)
		}
		if mapTmp["authgrp"] != nil {
			output.Authgrp = mapTmp["authgrp"].(string)
		}
		if mapTmp["sysgrp"] != nil {
			output.Sysgrp = mapTmp["sysgrp"].(string)
		}
		if mapTmp["netgrp"] != nil {
			output.Netgrp = mapTmp["netgrp"].(string)
		}
		if mapTmp["loggrp"] != nil {
			output.Loggrp = mapTmp["loggrp"].(string)
		}
		if mapTmp["fwgrp"] != nil {
			output.Fwgrp = mapTmp["fwgrp"].(string)
		}
		if mapTmp["vpngrp"] != nil {
			output.Vpngrp = mapTmp["vpngrp"].(string)
		}
		if mapTmp["utmgrp"] != nil {
			output.Utmgrp = mapTmp["utmgrp"].(string)
		}
		if mapTmp["wanoptgrp"] != nil {
			output.Wanoptgrp = mapTmp["wanoptgrp"].(string)
		}
		if mapTmp["wifi"] != nil {
			output.Wifi = mapTmp["wifi"].(string)
		}
		if mapTmp["admintimeout-override"] != nil {
			output.AdmintimeoutOverride = mapTmp["admintimeout-override"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
