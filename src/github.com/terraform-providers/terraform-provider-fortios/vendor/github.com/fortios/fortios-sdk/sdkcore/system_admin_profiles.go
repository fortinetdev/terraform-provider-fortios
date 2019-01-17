package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

//JSONSystemAdminProfiles contains ... need to comment completely
type JSONSystemAdminProfiles struct {
	Name                 string `json:"name"`
	Scope                string `json:"scope"`
	Comments             string `json:"comments"`
	Secfabgrp            string `json:"secfabgrp"`
	Ftviewgrp            string `json:"ftviewgrp"`
	Authgrp              string `json:"authgrp"`
	Sysgrp               string `json:"sysgrp"`
	Netgrp               string `json:"netgrp"`
	Loggrp               string `json:"loggrp"`
	Fwgrp                string `json:"fwgrp"`
	Vpngrp               string `json:"vpngrp"`
	Utmgrp               string `json:"utmgrp"`
	Wanoptgrp            string `json:"wanoptgrp"`
	Wifi                 string `json:"wifi"`
	AdmintimeoutOverride string `json:"admintimeout-override"`
}

//JSONCreateSystemAdminProfilesOutput contains ... need to comment completely
type JSONCreateSystemAdminProfilesOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       float64 `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//JSONUpdateSystemAdminProfilesOutput contains ... need to comment completely
//Attention: The RESTful API changed the Mkey type from float64 in CREATE to string in UPDATE!
type JSONUpdateSystemAdminProfilesOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//CreateSystemAdminProfiles will send ... need to comment completely
func (c *FortiSDKClient) CreateSystemAdminProfiles(params *JSONSystemAdminProfiles) (output *JSONUpdateSystemAdminProfilesOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/accprofile"
	output = &JSONUpdateSystemAdminProfilesOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)

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
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get the right response")
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

//UpdateSystemAdminProfiles will send ... need to comment completely
func (c *FortiSDKClient) UpdateSystemAdminProfiles(params *JSONSystemAdminProfiles, mkey string) (output *JSONUpdateSystemAdminProfilesOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/accprofile"
	path += "/" + mkey
	output = &JSONUpdateSystemAdminProfilesOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("shengh.............fortios response: %s", string(body))

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
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get the right response")
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

//DeleteSystemAdminProfiles will send ... need to comment completely
func (c *FortiSDKClient) DeleteSystemAdminProfiles(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/accprofile"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("shengh.............fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

//ReadSystemAdminProfiles will send ... need to comment completely
func (c *FortiSDKClient) ReadSystemAdminProfiles(mkey string) (output *JSONSystemAdminProfiles, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/accprofile"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("shengh.............fortios reading response: %s", string(body))

	output = &JSONSystemAdminProfiles{}
	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return
		}

		mapTmp := (result["results"].([]interface{}))[0].(map[string]interface{})

		if mapTmp == nil {
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
