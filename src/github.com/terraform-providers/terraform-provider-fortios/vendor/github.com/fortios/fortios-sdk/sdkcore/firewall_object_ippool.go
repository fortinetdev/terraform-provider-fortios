package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	// "strconv"
)

//JSONFirewallObjectIPPool contains ... need to comment completely
type JSONFirewallObjectIPPool struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Startip  string `json:"startip"`
	Endip    string `json:"endip"`
	ArpReply string `json:"arp-reply"`
	Comments string `json:"comments"`
}

//JSONCreateFirewallObjectIPPoolOutput contains ... need to comment completely
type JSONCreateFirewallObjectIPPoolOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//JSONUpdateFirewallObjectIPPoolOutput contains ... need to comment completely
//Attention: The RESTful API changed the Mkey type from float64 in CREATE to string in UPDATE!
type JSONUpdateFirewallObjectIPPoolOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//CreateFirewallObjectIPPool will send ... need to comment completely
func (c *FortiSDKClient) CreateFirewallObjectIPPool(params *JSONFirewallObjectIPPool) (output *JSONCreateFirewallObjectIPPoolOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/ippool"
	output = &JSONCreateFirewallObjectIPPoolOutput{}
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

//UpdateFirewallObjectIPPool will send ... need to comment completely
func (c *FortiSDKClient) UpdateFirewallObjectIPPool(params *JSONFirewallObjectIPPool, mkey string) (output *JSONUpdateFirewallObjectIPPoolOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/ippool"
	path += "/" + mkey
	output = &JSONUpdateFirewallObjectIPPoolOutput{}
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

//DeleteFirewallObjectIPPool will send ... need to comment completely
func (c *FortiSDKClient) DeleteFirewallObjectIPPool(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/ippool"
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

//ReadFirewallObjectIPPool will send ... need to comment completely
func (c *FortiSDKClient) ReadFirewallObjectIPPool(mkey string) (output *JSONFirewallObjectIPPool, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/ippool"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("shengh.............fortios reading response: %s", string(body))

	output = &JSONFirewallObjectIPPool{}
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
		if mapTmp["startip"] != nil {
			output.Startip = mapTmp["startip"].(string)
		}
		if mapTmp["endip"] != nil {
			output.Endip = mapTmp["endip"].(string)
		}
		if mapTmp["arp-reply"] != nil {
			output.ArpReply = mapTmp["arp-reply"].(string)
		}
		if mapTmp["comments"] != nil {
			output.Comments = mapTmp["comments"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
