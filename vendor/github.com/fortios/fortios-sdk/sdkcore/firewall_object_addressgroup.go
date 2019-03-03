package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	// "strconv"
)

// JSONFirewallObjectAddressGroup contains the parameters for Create and Update API function
type JSONFirewallObjectAddressGroup struct {
	Name    string     `json:"name"`
	Member  MultValues `json:"member"`
	Comment string     `json:"comment"`
}

// JSONCreateFirewallObjectAddressGroupOutput contains the output results for Create API function
type JSONCreateFirewallObjectAddressGroupOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateFirewallObjectAddressGroupOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateFirewallObjectAddressGroupOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateFirewallObjectAddressGroup API operation for FortiOS creates a new firewall address group for firewall policies.
// Returns the index value of the firewall address group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - addrgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallObjectAddressGroup(params *JSONFirewallObjectAddressGroup) (output *JSONCreateFirewallObjectAddressGroupOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/addrgrp"
	output = &JSONCreateFirewallObjectAddressGroupOutput{}
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
			if result["status"] != "success" {
				err = fmt.Errorf("cannot get the right response")
				return
			}
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

// UpdateFirewallObjectAddressGroup API operation for FortiOS updates the specified firewall address group for firewall policies.
// Returns the index value of the firewall address group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - addrgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallObjectAddressGroup(params *JSONFirewallObjectAddressGroup, mkey string) (output *JSONUpdateFirewallObjectAddressGroupOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/addrgrp"
	path += "/" + mkey
	output = &JSONUpdateFirewallObjectAddressGroupOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
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
				err = fmt.Errorf("cannot get the right response")
				return
			}
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

// DeleteFirewallObjectAddressGroup API operation for FortiOS deletes the specified firewall address group for firewall policies.
// Returns error for service API and SDK errors.
// See the firewall - addrgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallObjectAddressGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/addrgrp"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return
		}

		if result["status"] != "success" {
			err = fmt.Errorf("cannot get the right response")
			return
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

// ReadFirewallObjectAddressGroup API operation for FortiOS gets the firewall address group for firewall policies
// with the specified index value.
// Returns the requested firewall address group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - addrgrp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallObjectAddressGroup(mkey string) (output *JSONFirewallObjectAddressGroup, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/addrgrp"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("FOS-fortios reading response: %s", string(body))

	output = &JSONFirewallObjectAddressGroup{}
	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return
		}
		
		if result["status"] != "success" {
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
		if mapTmp["comment"] != nil {
			output.Comment = mapTmp["comment"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
