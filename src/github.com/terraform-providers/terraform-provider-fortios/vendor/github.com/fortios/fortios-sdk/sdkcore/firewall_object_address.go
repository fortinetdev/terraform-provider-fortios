package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	// "strconv"
)

//JSONFirewallObjectAddressCommon contains ... need to comment completely
type JSONFirewallObjectAddressCommon struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Comment string `json:"comment"`
}

//JSONFirewallObjectAddressIPRange contains ... need to comment completely
type JSONFirewallObjectAddressIPRange struct {
	StartIP string `json:"start-ip"`
	EndIP   string `json:"end-ip"`
}

//JSONFirewallObjectAddressCountry contains ... need to comment completely
type JSONFirewallObjectAddressCountry struct {
	Country string `json:"country"`
}

//JSONFirewallObjectAddressFqdn contains ... need to comment completely
type JSONFirewallObjectAddressFqdn struct {
	Fqdn string `json:"fqdn"`
}

//JSONFirewallObjectAddressIPMask contains ... need to comment completely
type JSONFirewallObjectAddressIPMask struct {
	Subnet string `json:"subnet"`
}

//JSONFirewallObjectAddress contains ... need to comment completely
type JSONFirewallObjectAddress struct {
	*JSONFirewallObjectAddressCommon
	*JSONFirewallObjectAddressIPRange
	*JSONFirewallObjectAddressCountry
	*JSONFirewallObjectAddressFqdn
	*JSONFirewallObjectAddressIPMask
}

//JSONCreateFirewallObjectAddressOutput contains ... need to comment completely
type JSONCreateFirewallObjectAddressOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       float64 `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//JSONUpdateFirewallObjectAddressOutput contains ... need to comment completely
//Attention: The RESTful API changed the Mkey type from float64 in CREATE to string in UPDATE!
type JSONUpdateFirewallObjectAddressOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//CreateFirewallObjectAddress will send ... need to comment completely
func (c *FortiSDKClient) CreateFirewallObjectAddress(params *JSONFirewallObjectAddress) (output *JSONUpdateFirewallObjectAddressOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/address"
	output = &JSONUpdateFirewallObjectAddressOutput{}
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

//UpdateFirewallObjectAddress will send ... need to comment completely
func (c *FortiSDKClient) UpdateFirewallObjectAddress(params *JSONFirewallObjectAddress, mkey string) (output *JSONUpdateFirewallObjectAddressOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/address"
	path += "/" + mkey
	output = &JSONUpdateFirewallObjectAddressOutput{}
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

//DeleteFirewallObjectAddress will send ... need to comment completely
func (c *FortiSDKClient) DeleteFirewallObjectAddress(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/address"
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

//ReadFirewallObjectAddress will send ... need to comment completely
func (c *FortiSDKClient) ReadFirewallObjectAddress(mkey string) (output *JSONFirewallObjectAddress, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/address"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("shengh.............fortios reading response: %s", string(body))

	j1 := JSONFirewallObjectAddressCommon{}
	j2 := JSONFirewallObjectAddressIPRange{}
	j3 := JSONFirewallObjectAddressCountry{}
	j4 := JSONFirewallObjectAddressFqdn{}
	j5 := JSONFirewallObjectAddressIPMask{}

	output = &JSONFirewallObjectAddress{
		JSONFirewallObjectAddressCommon:  &j1,
		JSONFirewallObjectAddressIPRange: &j2,
		JSONFirewallObjectAddressCountry: &j3,
		JSONFirewallObjectAddressFqdn:    &j4,
		JSONFirewallObjectAddressIPMask:  &j5,
	}

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
		if mapTmp["type"] != nil {
			output.Type = mapTmp["type"].(string)
		}
		if mapTmp["subnet"] != nil {
			output.Subnet = mapTmp["subnet"].(string)
		}
		if mapTmp["start-ip"] != nil {
			output.StartIP = mapTmp["start-ip"].(string)
		}
		if mapTmp["end-ip"] != nil {
			output.EndIP = mapTmp["end-ip"].(string)
		}
		if mapTmp["fqdn"] != nil {
			output.Fqdn = mapTmp["fqdn"].(string)
		}
		if mapTmp["country"] != nil {
			output.Country = mapTmp["country"].(string)
		}
		if mapTmp["comment"] != nil {
			output.Comment = mapTmp["comment"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
