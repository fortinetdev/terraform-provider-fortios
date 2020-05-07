package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fgtdev/fortios-sdk-go/util"
)

// JSONVPNIPsecPhase1Interface contains the parameters for Create and Update API function
type JSONVPNIPsecPhase1Interface struct {
	Name                string     `json:"name"`
	Type                string     `json:"type"`
	Interface           string     `json:"interface"`
	Peertype            string     `json:"peertype"`
	Proposal            string     `json:"proposal"`
	Comments            string     `json:"comments"`
	WizardType          string     `json:"wizard-type"`
	RemoteGw            string     `json:"remote-gw"`
	Psksecret           string     `json:"psksecret"`
	Certificate         MultValues `json:"certificate"`
	Peerid              string     `json:"peerid"`
	Peer                string     `json:"peer"`
	Peergrp             string     `json:"peergrp"`
	IPv4SplitInclude    string     `json:"ipv4-split-include"`
	SplitIncludeService string     `json:"split-include-service"`
	IPv4SplitExclude    string     `json:"ipv4-split-exclude"`
	ModeCfg             string     `json:"mode-cfg"`
	Authmethod          string     `json:"authmethod"`
	AuthmethodRemote    string     `json:"authmethod-remote"`
}

// JSONCreateVPNIPsecPhase1InterfaceOutput contains the output results for Create API function
type JSONCreateVPNIPsecPhase1InterfaceOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateVPNIPsecPhase1InterfaceOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateVPNIPsecPhase1InterfaceOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateVPNIPsecPhase1Interface API operation for FortiOS creates a new phase 1 definition for a route-based (interface mode) IPsec VPN tunnel.
// Returns the index value of the phase1-interface setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn - ipsec phase1-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVPNIPsecPhase1Interface(params *JSONVPNIPsecPhase1Interface) (output *JSONCreateVPNIPsecPhase1InterfaceOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.ipsec/phase1-interface"
	output = &JSONCreateVPNIPsecPhase1InterfaceOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	e, req := c.NewRequest(HTTPMethod, path, nil, bytes)
	if e != nil {
		err = fmt.Errorf("new request error %s", e)
		return
	}
	
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

// UpdateVPNIPsecPhase1Interface API operation for FortiOS updates the specified phase1-interface setting.
// Returns the index value of the phase1-interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn - ipsec phase1-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVPNIPsecPhase1Interface(params *JSONVPNIPsecPhase1Interface, mkey string) (output *JSONUpdateVPNIPsecPhase1InterfaceOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.ipsec/phase1-interface"
	path += "/" + EscapeURLString(mkey)
	output = &JSONUpdateVPNIPsecPhase1InterfaceOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	e, req := c.NewRequest(HTTPMethod, path, nil, bytes)
	if e != nil {
		err = fmt.Errorf("new request error %s", e)
		return
	}
	
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

// DeleteVPNIPsecPhase1Interface API operation for FortiOS deletes the specified phase1-interface setting.
// Returns error for service API and SDK errors.
// See the vpn - ipsec phase1-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVPNIPsecPhase1Interface(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.ipsec/phase1-interface"
	path += "/" + EscapeURLString(mkey)

	e, req := c.NewRequest(HTTPMethod, path, nil, nil)
	if e != nil {
		err = fmt.Errorf("new request error %s", e)
		return
	}
	
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

// ReadVPNIPsecPhase1Interface API operation for FortiOS gets the phase1-interface setting
// with the specified index value.
// Returns the requested phase1-interface setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn - ipsec phase1-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVPNIPsecPhase1Interface(mkey string) (output *JSONVPNIPsecPhase1Interface, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.ipsec/phase1-interface"
	path += "/" + EscapeURLString(mkey)

	output = &JSONVPNIPsecPhase1Interface{}

	e, req := c.NewRequest(HTTPMethod, path, nil, nil)
	if e != nil {
		err = fmt.Errorf("new request error %s", e)
		return
	}
	
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
		}
		if mapTmp["interface"] != nil {
			output.Interface = mapTmp["interface"].(string)
		}
		if mapTmp["peertype"] != nil {
			output.Peertype = mapTmp["peertype"].(string)
		}
		if mapTmp["proposal"] != nil {
			output.Proposal = mapTmp["proposal"].(string)
		}
		if mapTmp["comments"] != nil {
			output.Comments = mapTmp["comments"].(string)
		}
		if mapTmp["wizard-type"] != nil {
			output.WizardType = mapTmp["wizard-type"].(string)
		}
		if mapTmp["remote-gw"] != nil {
			output.RemoteGw = mapTmp["remote-gw"].(string)
		}
		if mapTmp["psksecret"] != nil {
			output.Psksecret = mapTmp["psksecret"].(string)
		}
		if mapTmp["certificate"] != nil {
			member := mapTmp["certificate"].([]interface{})

			var members []MultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					MultValue{
						Name: c["name"].(string),
					})
			}
			output.Certificate = members
		}
		if mapTmp["peerid"] != nil {
			output.Peerid = mapTmp["peerid"].(string)
		}
		if mapTmp["peer"] != nil {
			output.Peer = mapTmp["peer"].(string)
		}
		if mapTmp["peergrp"] != nil {
			output.Peergrp = mapTmp["peergrp"].(string)
		}
		if mapTmp["ipv4-split-include"] != nil {
			output.IPv4SplitInclude = mapTmp["ipv4-split-include"].(string)
		}
		if mapTmp["split-include-service"] != nil {
			output.SplitIncludeService = mapTmp["split-include-service"].(string)
		}
		if mapTmp["ipv4-split-exclude"] != nil {
			output.IPv4SplitExclude = mapTmp["ipv4-split-exclude"].(string)
		}
		if mapTmp["mode-cfg"] != nil {
			output.ModeCfg = mapTmp["mode-cfg"].(string)
		}
		if mapTmp["authmethod"] != nil {
			output.Authmethod = mapTmp["authmethod"].(string)
		}
		if mapTmp["authmethod-remote"] != nil {
			output.AuthmethodRemote = mapTmp["authmethod-remote"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
