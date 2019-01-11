package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	// "strconv"
)

//JSONFirewallObjectVip contains ... need to comment completely
type JSONFirewallObjectVip struct {
	Name        string        `json:"name"`
	Comment     string        `json:"comment"`
	Extip       string        `json:"extip"`
	Mappedip    VIPMultValues `json:"mappedip"`
	Extintf     string        `json:"extintf"`
	Portforward string        `json:"portforward"`
}

//JSONCreateFirewallObjectVipOutput contains ... need to comment completely
type JSONCreateFirewallObjectVipOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//JSONUpdateFirewallObjectVipOutput contains ... need to comment completely
//Attention: The RESTful API changed the Mkey type from float64 in CREATE to string in UPDATE!
type JSONUpdateFirewallObjectVipOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//VIPMultValue contains ... need to comment completely
type VIPMultValue struct {
	Range string `json:"range"`
}

//VIPMultValues contains ... need to comment completely
type VIPMultValues []VIPMultValue

//CreateFirewallObjectVip will send ... need to comment completely
func (c *FortiSDKClient) CreateFirewallObjectVip(params *JSONFirewallObjectVip) (output *JSONCreateFirewallObjectVipOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/vip"
	output = &JSONCreateFirewallObjectVipOutput{}
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
	log.Printf("shengh1.............fortios response: %v", result)

	if result != nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}
		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(string)
		} else {
			err = fmt.Errorf("cannot get the right response")
			return
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

//UpdateFirewallObjectVip will send ... need to comment completely
func (c *FortiSDKClient) UpdateFirewallObjectVip(params *JSONFirewallObjectVip, mkey string) (output *JSONUpdateFirewallObjectVipOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/vip"
	path += "/" + mkey
	output = &JSONUpdateFirewallObjectVipOutput{}
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

//DeleteFirewallObjectVip will send ... need to comment completely
func (c *FortiSDKClient) DeleteFirewallObjectVip(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/vip"
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

//ReadFirewallObjectVip will send ... need to comment completely
func (c *FortiSDKClient) ReadFirewallObjectVip(mkey string) (output *JSONFirewallObjectVip, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/vip"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("shengh.............fortios reading response: %s", string(body))

	output = &JSONFirewallObjectVip{}
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
		if mapTmp["comment"] != nil {
			output.Comment = mapTmp["comment"].(string)
		}
		if mapTmp["extip"] != nil {
			output.Extip = mapTmp["extip"].(string)
		}
		if mapTmp["mappedip"] != nil {
			member := mapTmp["mappedip"].([]interface{})

			var members []VIPMultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					VIPMultValue{
						Range: c["range"].(string),
					})
			}
			output.Mappedip = members
		}
		if mapTmp["extintf"] != nil {
			output.Extintf = mapTmp["extintf"].(string)
		}
		if mapTmp["portforward"] != nil {
			output.Portforward = mapTmp["portforward"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
