package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fgtdev/fortios-sdk-go/util"
)

// JSONFirewallObjectAddressCommon contains the General parameters for Create and Update API function
type JSONFirewallObjectAddressCommon struct {
	Name              string `json:"name"`
	Type              string `json:"type"`
	Comment           string `json:"comment"`
	AssociatedIntf    string `json:"associated-interface"`
	ShowInAddressList string `json:"visibility"`
	AllowRouting      string `json:"allow-routing"`
}

// JSONFirewallObjectAddressIPRange contains the IP Range parameters for Create and Update API function
type JSONFirewallObjectAddressIPRange struct {
	StartIP string `json:"start-ip"`
	EndIP   string `json:"end-ip"`
}

// JSONFirewallObjectAddressCountry contains the Country parameters for Create and Update API function
type JSONFirewallObjectAddressCountry struct {
	Country string `json:"country"`
}

// JSONFirewallObjectAddressFqdn contains the FQDN parameters for Create and Update API function
type JSONFirewallObjectAddressFqdn struct {
	Fqdn string `json:"fqdn"`
}

// JSONFirewallObjectAddressIPMask contains the Subnet parameters for Create and Update API function
type JSONFirewallObjectAddressIPMask struct {
	Subnet string `json:"subnet"`
}

// JSONFirewallObjectAddress contains the parameters for Create and Update API function
type JSONFirewallObjectAddress struct {
	*JSONFirewallObjectAddressCommon
	*JSONFirewallObjectAddressIPRange
	*JSONFirewallObjectAddressCountry
	*JSONFirewallObjectAddressFqdn
	*JSONFirewallObjectAddressIPMask
}

// JSONCreateFirewallObjectAddressOutput contains the output results for Create API function
type JSONCreateFirewallObjectAddressOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateFirewallObjectAddressOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateFirewallObjectAddressOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateFirewallObjectAddress API operation for FortiOS creates a new firewall address for firewall policies.
// Returns the index value of the firewall address and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - address chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallObjectAddress(params *JSONFirewallObjectAddress) (output *JSONCreateFirewallObjectAddressOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall/address"
	output = &JSONCreateFirewallObjectAddressOutput{}
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

	if result != nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}
		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(string)
		}
		if result["status"] != nil {
			if result["status"] != "success" {
				if result["error"] != nil {
					err = fmt.Errorf("status is %s and error no is %.0f", result["status"], result["error"])
				} else {
					err = fmt.Errorf("status is %s and error no is not found", result["status"])
				}

				if result["http_status"] != nil {
					err = fmt.Errorf("%s, details: %s", err, util.HttpStatus2Str(int(result["http_status"].(float64))))
				} else {
					err = fmt.Errorf("%s, and http_status no is not found", err)
				}

				return
			}
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get status from the response")
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

// UpdateFirewallObjectAddress API operation for FortiOS updates the specified firewall address for firewall policies.
// Returns the index value of the firewall address and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - address chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallObjectAddress(params *JSONFirewallObjectAddress, mkey string) (output *JSONUpdateFirewallObjectAddressOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall/address"
	path += "/" + EscapeURLString(mkey)
	output = &JSONUpdateFirewallObjectAddressOutput{}
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

	if result != nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}
		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(string)
		}
		if result["status"] != nil {
			if result["status"] != "success" {
				if result["error"] != nil {
					err = fmt.Errorf("status is %s and error no is %.0f", result["status"], result["error"])
				} else {
					err = fmt.Errorf("status is %s and error no is not found", result["status"])
				}

				if result["http_status"] != nil {
					err = fmt.Errorf("%s, details: %s", err, util.HttpStatus2Str(int(result["http_status"].(float64))))
				} else {
					err = fmt.Errorf("%s, and http_status no is not found", err)
				}

				return
			}
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get status from the response")
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

// DeleteFirewallObjectAddress API operation for FortiOS deletes the specified firewall address for firewall policies.
// Returns error for service API and SDK errors.
// See the firewall - address chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallObjectAddress(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/address"
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

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get status from the response")
			return
		}

		if result["status"] != "success" {
			if result["error"] != nil {
				err = fmt.Errorf("status is %s and error no is %.0f", result["status"], result["error"])
			} else {
				err = fmt.Errorf("status is %s and error no is not found", result["status"])
			}

			if result["http_status"] != nil {
				err = fmt.Errorf("%s, details: %s", err, util.HttpStatus2Str(int(result["http_status"].(float64))))
			} else {
				err = fmt.Errorf("%s, and http_status no is not found", err)
			}

			return
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

// ReadFirewallObjectAddress API operation for FortiOS gets the firewall address for firewall policies
// with the specified index value.
// Returns the requested firewall addresses value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewall - address chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallObjectAddress(mkey string) (output *JSONFirewallObjectAddress, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall/address"
	path += "/" + EscapeURLString(mkey)

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

	if result != nil {
		if result["http_status"] == nil {
			err = fmt.Errorf("cannot get http_status from the response")
			return
		}

		if result["http_status"] == 404.0 {
			output = nil
			return
		}

		if result["status"] == nil {
			err = fmt.Errorf("cannot get status from the response")
			return
		}

		if result["status"] != "success" {
			if result["error"] != nil {
				err = fmt.Errorf("status is %s and error no is %.0f", result["status"], result["error"])
			} else {
				err = fmt.Errorf("status is %s and error no is not found", result["status"])
			}

			if result["http_status"] != nil {
				err = fmt.Errorf("%s, details: %s", err, util.HttpStatus2Str(int(result["http_status"].(float64))))
			} else {
				err = fmt.Errorf("%s, and http_status no is not found", err)
			}

			return
		}

		mapTmp := (result["results"].([]interface{}))[0].(map[string]interface{})

		if mapTmp == nil {
			err = fmt.Errorf("cannot get the results from the response")
			return
		}

		if mapTmp["name"] != nil {
			output.Name = mapTmp["name"].(string)
		}
		if mapTmp["type"] != nil {
			output.Type = mapTmp["type"].(string)
		} else {
			err = fmt.Errorf("cannot get the right response, type doesn't exist.")
			return
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
		if mapTmp["associated-interface"] != nil {
			output.AssociatedIntf = mapTmp["associated-interface"].(string)
		}
		if mapTmp["visibility"] != nil {
			output.ShowInAddressList = mapTmp["visibility"].(string)
		}
		if mapTmp["allow-routing"] != nil {
			output.AllowRouting = mapTmp["allow-routing"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
