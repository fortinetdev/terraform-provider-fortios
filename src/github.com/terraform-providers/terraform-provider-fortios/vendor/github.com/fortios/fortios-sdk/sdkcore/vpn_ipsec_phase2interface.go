package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	// "strconv"
)

//JSONVPNIPsecPhase2Interface contains ... need to comment completely
type JSONVPNIPsecPhase2Interface struct {
	Name        string `json:"name"`
	Phase1name  string `json:"phase1name"`
	Proposal    string `json:"proposal"`
	Comments    string `json:"comments"`
	SrcAddrType string `json:"src-addr-type"`
	SrcStartIP  string `json:"src-start-ip"`
	SrcEndIP    string `json:"src-end-ip"`
	SrcSubnet   string `json:"src-subnet"`
	DstAddrType string `json:"dst-addr-type"`
	SrcName     string `json:"src-name"`
	DstName     string `json:"dst-name"`
	DstStartIP  string `json:"dst-start-ip"`
	DstEndIP    string `json:"dst-end-ip"`
	DstSubnet   string `json:"dst-subnet"`
}

//JSONCreateVPNIPsecPhase2InterfaceOutput contains ... need to comment completely
type JSONCreateVPNIPsecPhase2InterfaceOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//JSONUpdateVPNIPsecPhase2InterfaceOutput contains ... need to comment completely
//Attention: The RESTful API changed the Mkey type from float64 in CREATE to string in UPDATE!
type JSONUpdateVPNIPsecPhase2InterfaceOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

//CreateVPNIPsecPhase2Interface will send ... need to comment completely
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
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)

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

//UpdateVPNIPsecPhase2Interface will send ... need to comment completely
func (c *FortiSDKClient) UpdateVPNIPsecPhase2Interface(params *JSONVPNIPsecPhase2Interface, mkey string) (output *JSONUpdateVPNIPsecPhase2InterfaceOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/vpn.ipsec/phase2-interface"
	path += "/" + mkey
	output = &JSONUpdateVPNIPsecPhase2InterfaceOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("shengh.............fortios response: %s", string(body))

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

//DeleteVPNIPsecPhase2Interface will send ... need to comment completely
func (c *FortiSDKClient) DeleteVPNIPsecPhase2Interface(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/vpn.ipsec/phase2-interface"
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

//ReadVPNIPsecPhase2Interface will send ... need to comment completely
func (c *FortiSDKClient) ReadVPNIPsecPhase2Interface(mkey string) (output *JSONVPNIPsecPhase2Interface, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/vpn.ipsec/phase2-interface"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("shengh.............fortios reading response: %s", string(body))

	output = &JSONVPNIPsecPhase2Interface{}
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
