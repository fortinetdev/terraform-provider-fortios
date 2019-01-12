package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

//JSONSystemSettingGlobal contains ... need to comment completely
type JSONSystemSettingGlobal struct {
	Admintimeout string `json:"admintimeout"`
	Timezone     string `json:"timezone"`
	Hostname     string `json:"hostname"`
	AdminSport   string `json:"admin-sport"`
	AdminSSHPort string `json:"admin-ssh-port"`
}

//JSONCreateSystemSettingGlobalOutput contains ... need to comment completely
type JSONCreateSystemSettingGlobalOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//JSONUpdateSystemSettingGlobalOutput contains ... need to comment completely
//Attention: The RESTful API changed the Mkey type from float64 in CREATE to string in UPDATE!
type JSONUpdateSystemSettingGlobalOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//CreateSystemSettingGlobal will send ... need to comment completely
func (c *FortiSDKClient) CreateSystemSettingGlobal(params *JSONSystemSettingGlobal) (output *JSONCreateSystemSettingGlobalOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/global"
	output = &JSONCreateSystemSettingGlobalOutput{}
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

//UpdateSystemSettingGlobal will send ... need to comment completely
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

//DeleteSystemSettingGlobal will send ... need to comment completely
func (c *FortiSDKClient) DeleteSystemSettingGlobal(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/global"
	// path += "/" + mkey

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

//ReadSystemSettingGlobal will send ... need to comment completely
func (c *FortiSDKClient) ReadSystemSettingGlobal(mkey string) (output *JSONSystemSettingGlobal, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/global"
	// path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("shengh.............fortios reading response: %s", string(body))

	output = &JSONSystemSettingGlobal{}
	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return
		}

		mapTmp := (result["results"].(map[string]interface{}))

		if mapTmp == nil {
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

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
