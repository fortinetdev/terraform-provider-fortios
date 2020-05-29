package fmgclient

import (
	"fmt"

	"github.com/fortinetdev/forti-sdk-go/fortimanager/util"
)

// JSONSystemNetworkInterface contains the params for updating network interface setting
type JSONSystemNetworkInterface struct {
	Name          string   `json:"name"`
	Ip            string   `json:"ip,omitempty"`
	Status        string   `json:"status,omitempty"`
	Description   string   `json:"description,omitempty"`
	AllowAccess   []string `json:"allowaccess,omitempty"`
	ServiceAccess []string `json:"serviceaccess,omitempty"`
}

// UpdateSystemNetworkInterface is for updating network interface setting
// Input:
//   @params: infor needed
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) UpdateSystemNetworkInterface(params *JSONSystemNetworkInterface) (err error) {
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

// ReadSystemNetworkInterface is for reading the specific network interface setting
// Input:
//   @name: network interface name
// Output:
//   @out: network interface infor
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) ReadSystemNetworkInterface(name string) (out *JSONSystemNetworkInterface, err error) {
	defer c.Trace("ReadSystemNetworkInterface")()

	p := map[string]interface{}{
		"url": "/cli/global/system/interface/" + name,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadSystemNetworkInterface failed: %s", err)
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
		out.Status = util.IntfStatus2Str(int(data["status"].(float64)))
	}

	return
}
