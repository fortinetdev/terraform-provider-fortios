package fortimngclient

import (
	"fmt"
)

type JSONSystemNetworkInterface struct {
	Name          string   `json:"name"`
	Ip            string   `json:"ip,omitempty"`
	Status        string   `json:"status,omitempty"`
	Description   string   `json:"description,omitempty"`
	AllowAccess   []string `json:"allowaccess,omitempty"`
	ServiceAccess []string `json:"serviceaccess,omitempty"`
}

func (c *FortiMngClient) UpdateSystemNetworkInterface(params *JSONSystemNetworkInterface) (err error) {
	defer c.Trace("UpdateSystemNetworkInterface")()

	p := map[string]interface{}{
		"data": *params,
		"url":  "/cli/global/system/interface",
	}

	_, err = c.Do("update", p)

	if err != nil {
		err = fmt.Errorf("UpdateSystemNetworkInterface failed: %s", err)
		return
	}

	return
}

func (c *FortiMngClient) ReadSystemNetworkInterface(name string) (out *JSONSystemNetworkInterface, err error) {
	defer c.Trace("ReadSystemNetworkInterface")()

	p := map[string]interface{}{
		"url": "/cli/global/system/interface/" + name,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadSystemNetworkInterface failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONSystemNetworkInterface{}
	if data["name"] != nil {
		out.Name = data["name"].(string)
	}
	if data["description"] != nil {
		out.Description = data["description"].(string)
	}
	if data["ip"] != nil {
		ipaddr := data["ip"].([]interface{})
		out.Ip = ipaddr[0].(string) + " " + ipaddr[1].(string)
	}
	if data["status"] != nil {
		out.Status = c.IntfStatus2Str(int(data["status"].(float64)))
	}
	/*
		if data["allowaccess"] != nil {
			out.AllowAccess = c.Tmp(strconv.Itoa(int(data["allowaccess"].(float64))))
		}
		if data["serviceaccess"] != nil {
			out.ServiceAccess = c.Tmp1(strconv.Itoa(int(data["serviceaccess"].(float64))))
		}
	*/

	return
}
