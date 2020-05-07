package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fgtdev/fortios-sdk-go/util"
)

// JSONVPNIPsecPhase2Interface contains the parameters for Create and Update API function
type JSONVPNIPsecPhase2Interface struct {
	Name        string `json:"name"`
	Phase1name  string `json:"phase1name"`
	Proposal    string `json:"proposal"`
	Comments    string `json:"comments"`
	SrcAddrType string `json:"src-addr-type"`
	SrcStartIP  string `json:"src-start-ip,omitempty"`
	SrcEndIP    string `json:"src-end-ip,omitempty"`
	SrcSubnet   string `json:"src-subnet,omitempty"`
	DstAddrType string `json:"dst-addr-type"`
	SrcName     string `json:"src-name"`
	DstName     string `json:"dst-name"`
	DstStartIP  string `json:"dst-start-ip,omitempty"`
	DstEndIP    string `json:"dst-end-ip,omitempty"`
	DstSubnet   string `json:"dst-subnet,omitempty"`
}

// JSONCreateVPNIPsecPhase2InterfaceOutput contains the output results for Create API function
type JSONCreateVPNIPsecPhase2InterfaceOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateVPNIPsecPhase2InterfaceOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateVPNIPsecPhase2InterfaceOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateVPNIPsecPhase2Interface API operation for FortiOS creates a new a new phase 2 definition for a route-based (interface mode) IPsec VPN tunnel.
// Returns the index value of the phase2-interface setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn - ipsec phase2-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateVPNIPsecPhase2Interface(params *JSONVPNIPsecPhase2Interface) (output *JSONCreateVPNIPsecPhase2InterfaceOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/vpn.ipsec/phase2-interface"
	output = &JSONCreateVPNIPsecPhase2InterfaceOutput{}
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

// UpdateVPNIPsecPhase2Interface API operation for FortiOS updates the specified phase2-interface setting.
// Returns the index value of the phase2-interface setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn - ipsec phase2-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateVPNIPsecPhase2Interface(params *JSONVPNIPsecPhase2Interface, mkey string) (output *JSONUpdateVPNIPsecPhase2InterfaceOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.ipsec/phase2-interface"
	path += "/" + EscapeURLString(mkey)
	output = &JSONUpdateVPNIPsecPhase2InterfaceOutput{}
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

// DeleteVPNIPsecPhase2Interface API operation for FortiOS deletes the specified phase2-interface setting.
// Returns error for service API and SDK errors.
// See the vpn - ipsec phase2-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteVPNIPsecPhase2Interface(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.ipsec/phase2-interface"
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

// ReadVPNIPsecPhase2Interface API operation for FortiOS gets the phase2-interface setting
// with the specified index value.
// Returns the requested phase2-interface setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the vpn - ipsec phase2-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadVPNIPsecPhase2Interface(mkey string) (output *JSONVPNIPsecPhase2Interface, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.ipsec/phase2-interface"
	path += "/" + EscapeURLString(mkey)

	output = &JSONVPNIPsecPhase2Interface{}

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
		if mapTmp["phase1name"] != nil {
			output.Phase1name = mapTmp["phase1name"].(string)
		}
		if mapTmp["proposal"] != nil {
			output.Proposal = mapTmp["proposal"].(string)
		}
		if mapTmp["comments"] != nil {
			output.Comments = mapTmp["comments"].(string)
		}
		if mapTmp["src-addr-type"] != nil {
			output.SrcAddrType = mapTmp["src-addr-type"].(string)
		}
		if mapTmp["src-start-ip"] != nil {
			output.SrcStartIP = mapTmp["src-start-ip"].(string)
		}
		if mapTmp["src-end-ip"] != nil {
			output.SrcEndIP = mapTmp["src-end-ip"].(string)
		}
		if mapTmp["src-subnet"] != nil {
			output.SrcSubnet = mapTmp["src-subnet"].(string)
		}
		if mapTmp["dst-addr-type"] != nil {
			output.DstAddrType = mapTmp["dst-addr-type"].(string)
		}
		if mapTmp["src-name"] != nil {
			output.SrcName = mapTmp["src-name"].(string)
		}
		if mapTmp["dst-name"] != nil {
			output.DstName = mapTmp["dst-name"].(string)
		}
		if mapTmp["dst-start-ip"] != nil {
			output.DstStartIP = mapTmp["dst-start-ip"].(string)
		}
		if mapTmp["dst-end-ip"] != nil {
			output.DstEndIP = mapTmp["dst-end-ip"].(string)
		}
		if mapTmp["dst-subnet"] != nil {
			output.DstSubnet = mapTmp["dst-subnet"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
