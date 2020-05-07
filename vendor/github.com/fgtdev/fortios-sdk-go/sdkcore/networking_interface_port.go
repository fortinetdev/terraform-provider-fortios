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

// JSONNetworkingInterfacePort contains the parameters for Create and Update API function
type JSONNetworkingInterfacePort struct {
	// Portname          string `json:"portname"`
	Ipf                  string `json:"ip"`
	Alias                string `json:"alias"`
	Status               string `json:"status"`
	DeviceIdentification string `json:"device-identification"`
	TCPMss               string `json:"tcp-mss"`
	Speed                string `json:"speed"`
	MtuOverride          string `json:"mtu-override"`
	Mtu                  string `json:"mtu"`
	Role                 string `json:"role"`
	Allowaccess          string `json:"allowaccess"`
	Mode                 string `json:"mode"`
	DNSServerOverride    string `json:"dns-server-override"`
	Defaultgw            string `json:"defaultgw"`
	Distance             string `json:"distance"`
	Description          string `json:"description"`
	Type                 string `json:"type"`
	Interface            string `json:"interface"`
	Name                 string `json:"name"`
	Vdom                 string `json:"vdom"`
	Vlanid               string `json:"vlanid"`
}

// JSONCreateNetworkingInterfacePortOutput contains the output results for Create API function
type JSONCreateNetworkingInterfacePortOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateNetworkingInterfacePortOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateNetworkingInterfacePortOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateNetworkingInterfacePort API operation for FortiOS creates a new interface.
// Returns the index value of the interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateNetworkingInterfacePort(params *JSONNetworkingInterfacePort) (output *JSONCreateNetworkingInterfacePortOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/interface"
	output = &JSONCreateNetworkingInterfacePortOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	//log.Printf("FOS-fortios resquest1: %s", locJSON)

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
	//log.Printf("FOS-fortios response1: %s", string(body))

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

// UpdateNetworkingInterfacePort API operation for FortiOS updates the specified interface.
// Returns the index value of the interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateNetworkingInterfacePort(params *JSONNetworkingInterfacePort, mkey string) (output *JSONUpdateNetworkingInterfacePortOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/interface"
	path += "/" + EscapeURLString(mkey)
	output = &JSONUpdateNetworkingInterfacePortOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("FOS-fortios resquest2: %s", locJSON)

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
	log.Printf("FOS-fortios response2: %s", string(body))

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

// DeleteNetworkingInterfacePort API operation for FortiOS deletes the specified interface.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteNetworkingInterfacePort(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/interface"
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

// ReadNetworkingInterfacePort API operation for FortiOS gets the interface
// with the specified index value.
// Returns the requested interface value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadNetworkingInterfacePort(mkey string) (output *JSONNetworkingInterfacePort, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/interface"
	path += "/" + EscapeURLString(mkey)

	output = &JSONNetworkingInterfacePort{}

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

		// if mapTmp["portname"] != nil {
		// 	output.Portname = mapTmp["portname"].(string)
		// }
		if mapTmp["ip"] != nil {
			output.Ipf = mapTmp["ip"].(string)
		}
		if mapTmp["alias"] != nil {
			output.Alias = mapTmp["alias"].(string)
		}
		if mapTmp["status"] != nil {
			output.Status = mapTmp["status"].(string)
		}
		if mapTmp["device-identification"] != nil {
			output.DeviceIdentification = mapTmp["device-identification"].(string)
		}
		if mapTmp["tcp-mss"] != nil {
			output.TCPMss = strconv.Itoa(int(mapTmp["tcp-mss"].(float64)))
		}
		if mapTmp["speed"] != nil {
			output.Speed = mapTmp["speed"].(string)
		}
		if mapTmp["mtu-override"] != nil {
			output.MtuOverride = mapTmp["mtu-override"].(string)
		}
		if mapTmp["mtu"] != nil {
			output.Mtu = strconv.Itoa(int(mapTmp["mtu"].(float64)))
		}
		if mapTmp["role"] != nil {
			output.Role = mapTmp["role"].(string)
		}
		if mapTmp["allowaccess"] != nil {
			output.Allowaccess = mapTmp["allowaccess"].(string)
		}
		if mapTmp["mode"] != nil {
			output.Mode = mapTmp["mode"].(string)
		}
		if mapTmp["dns-server-override"] != nil {
			output.DNSServerOverride = mapTmp["dns-server-override"].(string)
		}
		if mapTmp["defaultgw"] != nil {
			output.Defaultgw = mapTmp["defaultgw"].(string)
		}
		if mapTmp["distance"] != nil {
			output.Distance = strconv.Itoa(int(mapTmp["distance"].(float64)))
		}
		if mapTmp["description"] != nil {
			output.Description = mapTmp["description"].(string)
		}
		if mapTmp["type"] != nil {
			output.Type = mapTmp["type"].(string)
		}
		if mapTmp["interface"] != nil {
			output.Interface = mapTmp["interface"].(string)
		}
		if mapTmp["name"] != nil {
			output.Name = mapTmp["name"].(string)
		}
		if mapTmp["vdom"] != nil {
			output.Vdom = mapTmp["vdom"].(string)
		}
		if mapTmp["vlanid"] != nil {
			output.Vlanid = strconv.Itoa(int(mapTmp["vlanid"].(float64)))
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
