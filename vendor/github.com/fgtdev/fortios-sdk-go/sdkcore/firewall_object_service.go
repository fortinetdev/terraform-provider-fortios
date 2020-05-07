package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/fgtdev/fortios-sdk-go/util"
)

// JSONFirewallObjectServiceCommon contains the General parameters for Create and Update API function
type JSONFirewallObjectServiceCommon struct {
	Name           string `json:"name"`
	Category       string `json:"category"`
	Protocol       string `json:"protocol"`
	Comment        string `json:"comment"`
	ProtocolNumber string `json:"protocol-number,omitempty"`
	Icmptype       string `json:"icmptype,omitempty"`
	Icmpcode       string `json:"icmpcode,omitempty"`
	TCPPortrange   string `json:"tcp-portrange,omitempty"`
	UDPPortrange   string `json:"udp-portrange,omitempty"`
	SctpPortrange  string `json:"sctp-portrange,omitempty"`
	SessionTTL     string `json:"session-ttl,omitempty"`
}

// JSONFirewallObjectServiceFqdn contains the FQDN parameters for Create and Update API function
type JSONFirewallObjectServiceFqdn struct {
	Fqdn string `json:"fqdn,omitempty"`
}

// JSONFirewallObjectServiceIprange contains the IP Range parameters for Create and Update API function
type JSONFirewallObjectServiceIprange struct {
	Iprange string `json:"iprange,omitempty"`
}

// JSONFirewallObjectService contains the parameters for Create and Update API function
type JSONFirewallObjectService struct {
	*JSONFirewallObjectServiceCommon
	*JSONFirewallObjectServiceFqdn
	*JSONFirewallObjectServiceIprange
}

// JSONCreateFirewallObjectServiceOutput contains the output results for Create API function
type JSONCreateFirewallObjectServiceOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateFirewallObjectServiceOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateFirewallObjectServiceOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateFirewallObjectService API operation for FortiOS creates a new firewall service.
// Returns the index value of the firewall service and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewal - service chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFirewallObjectService(params *JSONFirewallObjectService) (output *JSONCreateFirewallObjectServiceOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/firewall.service/custom"
	output = &JSONCreateFirewallObjectServiceOutput{}
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

// UpdateFirewallObjectService API operation for FortiOS updates the specified firewall service.
// Returns the index value of the firewall service and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewal - service chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFirewallObjectService(params *JSONFirewallObjectService, mkey string) (output *JSONUpdateFirewallObjectServiceOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/firewall.service/custom"
	path += "/" + EscapeURLString(mkey)
	output = &JSONUpdateFirewallObjectServiceOutput{}
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

// DeleteFirewallObjectService API operation for FortiOS deletes the specified firewall service.
// Returns error for service API and SDK errors.
// See the firewal - service chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallObjectService(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall.service/custom"
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

// ReadFirewallObjectService API operation for FortiOS gets the firewall service
// with the specified index value.
// Returns the requested firewall service value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the firewal - service chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFirewallObjectService(mkey string) (output *JSONFirewallObjectService, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/firewall.service/custom"
	path += "/" + EscapeURLString(mkey)

	j1 := JSONFirewallObjectServiceCommon{}
	j2 := JSONFirewallObjectServiceFqdn{}
	j3 := JSONFirewallObjectServiceIprange{}

	output = &JSONFirewallObjectService{
		JSONFirewallObjectServiceCommon:  &j1,
		JSONFirewallObjectServiceFqdn:    &j2,
		JSONFirewallObjectServiceIprange: &j3,
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
		if mapTmp["category"] != nil {
			output.Category = mapTmp["category"].(string)
		}
		if mapTmp["protocol"] != nil {
			output.Protocol = mapTmp["protocol"].(string)
		}
		if mapTmp["fqdn"] != nil {
			output.Fqdn = mapTmp["fqdn"].(string)
		}
		if mapTmp["iprange"] != nil {
			output.Iprange = mapTmp["iprange"].(string)
		}
		if mapTmp["comment"] != nil {
			output.Comment = mapTmp["comment"].(string)
		}
		if mapTmp["protocol-number"] != nil {
			output.ProtocolNumber = strconv.Itoa(int(mapTmp["protocol-number"].(float64)))
		}

		if output.Protocol == "ICMP" {
			if mapTmp["icmptype"] != nil {
				if mapTmp["icmptype"] != "" {
					output.Icmptype = strconv.Itoa(int(mapTmp["icmptype"].(float64)))
				} else {
					output.Icmptype = ""
				}
			}
			if mapTmp["icmpcode"] != nil {
				if mapTmp["icmpcode"] != "" {
					output.Icmpcode = strconv.Itoa(int(mapTmp["icmpcode"].(float64)))
				} else {
					output.Icmpcode = ""
				}
			}
		} else {
			if mapTmp["icmptype"] != nil {
				output.Icmptype = mapTmp["icmptype"].(string)
			}
			if mapTmp["icmpcode"] != nil {
				output.Icmpcode = mapTmp["icmpcode"].(string)
			}
		}
		if mapTmp["tcp-portrange"] != nil {
			output.TCPPortrange = mapTmp["tcp-portrange"].(string)
		}
		if mapTmp["udp-portrange"] != nil {
			output.UDPPortrange = mapTmp["udp-portrange"].(string)
		}
		if mapTmp["sctp-portrange"] != nil {
			output.SctpPortrange = mapTmp["sctp-portrange"].(string)
		}
		if mapTmp["session-ttl"] != nil {
			output.SessionTTL = strconv.Itoa(int(mapTmp["session-ttl"].(float64)))
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
