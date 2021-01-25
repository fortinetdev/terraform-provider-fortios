package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// JSONFirewallObjectIPPool contains the parameters for Create and Update API function
type JSONFirewallObjectIPPool struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Startip  string `json:"startip"`
	Endip    string `json:"endip"`
	ArpReply string `json:"arp-reply,omitempty"`
	Comments string `json:"comments"`
}

// JSONCreateFirewallObjectIPPoolOutput contains the output results for Create API function
type JSONCreateFirewallObjectIPPoolOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateFirewallObjectIPPoolOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateFirewallObjectIPPoolOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateFirewallObjectIPPool API operation for FortiOS creates a new IP address pool.
// Returns the index value of the IP address pool and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ippool chapter in the FortiOS Handbook - CLI Reference.
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

// UpdateFirewallObjectIPPool API operation for FortiOS updates the specified IP address pool.
// Returns the index value of the IP address pool and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ippool chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallObjectIPPool(params *JSONFirewallObjectIPPool, mkey string) (output *JSONUpdateFirewallObjectIPPoolOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/ippool"
	path += "/" + EscapeURLString(mkey)
	output = &JSONUpdateFirewallObjectIPPoolOutput{}
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

// DeleteFirewallObjectIPPool API operation for FortiOS deletes the specified IP address pool.
// Returns error for service API and SDK errors.
// See the firewall - ippool chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallObjectIPPool(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/ippool"
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

// ReadFirewallObjectIPPool API operation for FortiOS gets the IP address pool
// with the specified index value.
// Returns the requested IP address pool value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - ippool chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallObjectIPPool(mkey string) (output *JSONFirewallObjectIPPool, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/ippool"
	path += "/" + EscapeURLString(mkey)

	output = &JSONFirewallObjectIPPool{}

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
			err = fmt.Errorf("The set of return values is empty")
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

		return
	}

	return
}
