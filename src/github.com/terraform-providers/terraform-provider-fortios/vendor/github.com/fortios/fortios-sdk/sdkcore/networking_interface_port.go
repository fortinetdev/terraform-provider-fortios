package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

//JSONNetworkingInterfacePort contains ... need to comment completely
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

//JSONCreateNetworkingInterfacePortOutput contains ... need to comment completely
type JSONCreateNetworkingInterfacePortOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//JSONUpdateNetworkingInterfacePortOutput contains ... need to comment completely
//Attention: The RESTful API changed the Mkey type from float64 in CREATE to string in UPDATE!
type JSONUpdateNetworkingInterfacePortOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//CreateNetworkingInterfacePort will send ... need to comment completely
func (c *FortiSDKClient) CreateNetworkingInterfacePort(params *JSONNetworkingInterfacePort) (output *JSONCreateNetworkingInterfacePortOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/interface"
	output = &JSONCreateNetworkingInterfacePortOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	//log.Printf("shengh.............fortios resquest1: %s", locJSON)

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	//log.Printf("shengh.............fortios response1: %s", string(body))

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

//UpdateNetworkingInterfacePort will send ... need to comment completely
func (c *FortiSDKClient) UpdateNetworkingInterfacePort(params *JSONNetworkingInterfacePort, mkey string) (output *JSONUpdateNetworkingInterfacePortOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/interface"
	path += "/" + mkey
	output = &JSONUpdateNetworkingInterfacePortOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("shengh.............fortios resquest2: %s", locJSON)

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("shengh.............fortios response2: %s", string(body))

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

//DeleteNetworkingInterfacePort will send ... need to comment completely
func (c *FortiSDKClient) DeleteNetworkingInterfacePort(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/interface"
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

//ReadNetworkingInterfacePort will send ... need to comment completely
func (c *FortiSDKClient) ReadNetworkingInterfacePort(mkey string) (output *JSONNetworkingInterfacePort, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/interface"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("shengh.............fortios reading response: %s", string(body))

	output = &JSONNetworkingInterfacePort{}
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
