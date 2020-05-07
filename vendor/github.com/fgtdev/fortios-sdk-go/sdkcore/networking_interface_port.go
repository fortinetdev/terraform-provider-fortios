package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

)

// JSONNetworkingInterfacePort contains the parameters for Create and Update API function
type JSONNetworkingInterfacePort struct {
	// Portname          string `json:"portname"`
	Ipf                  string `json:"ip,omitempty"`
	Alias                string `json:"alias,omitempty"`
	Status               string `json:"status,omitempty"`
	DeviceIdentification string `json:"device-identification,omitempty"`
	TCPMss               string `json:"tcp-mss,omitempty"`
	Speed                string `json:"speed,omitempty"`
	MtuOverride          string `json:"mtu-override,omitempty"`
	Mtu                  string `json:"mtu,omitempty"`
	Role                 string `json:"role,omitempty"`
	Allowaccess          string `json:"allowaccess,omitempty"`
	Mode                 string `json:"mode,omitempty"`
	DNSServerOverride    string `json:"dns-server-override,omitempty"`
	Defaultgw            string `json:"defaultgw,omitempty"`
	Distance             string `json:"distance,omitempty"`
	Description          string `json:"description"`
	Type                 string `json:"type"`
	Interface            string `json:"interface,omitempty"`
	Name                 string `json:"name"`
	Vdom                 string `json:"vdom,omitempty"`
	Vlanid               string `json:"vlanid,omitempty"`
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
	//log.Printf("FOS-fortios response1: %s", string(body))

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
	log.Printf("FOS-fortios response2: %s", string(body))

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

// DeleteNetworkingInterfacePort API operation for FortiOS deletes the specified interface.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteNetworkingInterfacePort(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/interface"
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
