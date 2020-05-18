package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// JSONFirewallObjectVip contains the parameters for Create and Update API function
type JSONFirewallObjectVip struct {
	Name        string        `json:"name"`
	Comment     string        `json:"comment"`
	Extip       string        `json:"extip"`
	Mappedip    VIPMultValues `json:"mappedip"`
	Extintf     string        `json:"extintf,omitempty"`
	Portforward string        `json:"portforward,omitempty"`
	Protocol    string        `json:"protocol,omitempty"`
	Extport     string        `json:"extport,omitempty"`
	Mappedport  string        `json:"mappedport,omitempty"`
}

// JSONCreateFirewallObjectVipOutput contains the output results for Create API function
type JSONCreateFirewallObjectVipOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateFirewallObjectVipOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateFirewallObjectVipOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// VIPMultValue contains the output results for Read API function
type VIPMultValue struct {
	Range string `json:"range"`
}

// VIPMultValues contains the output results for Read API function
type VIPMultValues []VIPMultValue

// CreateFirewallObjectVip API operation for FortiOS creates a new firewall virtual IP.
// Returns the index value of the firewall virtual IP and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vip chapter in the FortiOS Handbook - CLI Reference.
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
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()
	log.Printf("FOS-fortios response: %v", result)

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}

		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(string)
		} else {
			err = fmt.Errorf("Failed to get mkey")
			return
		}
	}

	return
}

// UpdateFirewallObjectVip API operation for FortiOS updates the specified firewall virtual IP.
// Returns the index value of the firewall virtual IP and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vip chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallObjectVip(params *JSONFirewallObjectVip, mkey string) (output *JSONUpdateFirewallObjectVipOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/vip"
	path += "/" + EscapeURLString(mkey)
	output = &JSONUpdateFirewallObjectVipOutput{}
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
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

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

// DeleteFirewallObjectVip API operation for FortiOS deletes the specified firewall virtual IP.
// Returns error for service API and SDK errors.
// See the firewall - vip chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallObjectVip(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/vip"
	path += "/" + EscapeURLString(mkey)

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	err = fortiAPIErrorFormat(result, string(body))

	return
}

// ReadFirewallObjectVip API operation for FortiOS gets the firewall virtual IP
// with the specified index value.
// Returns the requested firewall virtual IP value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - vip chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallObjectVip(mkey string) (output *JSONFirewallObjectVip, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/vip"
	path += "/" + EscapeURLString(mkey)

	output = &JSONFirewallObjectVip{}

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios reading response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

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
		if mapTmp["protocol"] != nil {
			output.Protocol = mapTmp["protocol"].(string)
		}
		if mapTmp["extport"] != nil {
			output.Extport = mapTmp["extport"].(string)
		}
		if mapTmp["mappedport"] != nil {
			output.Mappedport = mapTmp["mappedport"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
