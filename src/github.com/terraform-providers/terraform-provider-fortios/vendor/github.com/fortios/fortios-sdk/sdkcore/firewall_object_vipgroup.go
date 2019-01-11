package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	// "strconv"
)

//JSONFirewallObjectVipGroup contains ... need to comment completely
type JSONFirewallObjectVipGroup struct {
	Name      string     `json:"name"`
	Comments  string     `json:"comments"`
	Interface string     `json:"interface"`
	Member    MultValues `json:"member"`
}

//JSONCreateFirewallObjectVipGroupOutput contains ... need to comment completely
type JSONCreateFirewallObjectVipGroupOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//JSONUpdateFirewallObjectVipGroupOutput contains ... need to comment completely
//Attention: The RESTful API changed the Mkey type from float64 in CREATE to string in UPDATE!
type JSONUpdateFirewallObjectVipGroupOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//CreateFirewallObjectVipGroup will send ... need to comment completely
func (c *FortiSDKClient) CreateFirewallObjectVipGroup(params *JSONFirewallObjectVipGroup) (output *JSONCreateFirewallObjectVipGroupOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/vipgrp"
	output = &JSONCreateFirewallObjectVipGroupOutput{}
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

//UpdateFirewallObjectVipGroup will send ... need to comment completely
func (c *FortiSDKClient) UpdateFirewallObjectVipGroup(params *JSONFirewallObjectVipGroup, mkey string) (output *JSONUpdateFirewallObjectVipGroupOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/vipgrp"
	path += "/" + mkey
	output = &JSONUpdateFirewallObjectVipGroupOutput{}
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

//DeleteFirewallObjectVipGroup will send ... need to comment completely
func (c *FortiSDKClient) DeleteFirewallObjectVipGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/vipgrp"
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

//ReadFirewallObjectVipGroup will send ... need to comment completely
func (c *FortiSDKClient) ReadFirewallObjectVipGroup(mkey string) (output *JSONFirewallObjectVipGroup, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/vipgrp"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("shengh.............fortios reading response: %s", string(body))

	output = &JSONFirewallObjectVipGroup{}
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
		if mapTmp["comments"] != nil {
			output.Comments = mapTmp["comments"].(string)
		}
		if mapTmp["interface"] != nil {
			output.Interface = mapTmp["interface"].(string)
		}
		if mapTmp["member"] != nil {
			member := mapTmp["member"].([]interface{})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					MultValue{
						Name: c["name"].(string),
					})
			}
			output.Member = members
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
