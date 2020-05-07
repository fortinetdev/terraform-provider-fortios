package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fgtdev/fortios-sdk-go/util"
)

// JSONFirewallObjectVip contains the parameters for Create and Update API function
type JSONFirewallObjectVip struct {
	Name        string        `json:"name"`
	Comment     string        `json:"comment"`
	Extip       string        `json:"extip"`
	Mappedip    VIPMultValues `json:"mappedip"`
	Extintf     string        `json:"extintf"`
	Portforward string        `json:"portforward"`
	Protocol    string        `json:"protocol"`
	Extport     string        `json:"extport"`
	Mappedport  string        `json:"mappedport"`
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
	log.Printf("FOS-fortios response: %v", result)

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

// DeleteFirewallObjectVip API operation for FortiOS deletes the specified firewall virtual IP.
// Returns error for service API and SDK errors.
// See the firewall - vip chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFirewallObjectVip(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/firewall/vip"
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
