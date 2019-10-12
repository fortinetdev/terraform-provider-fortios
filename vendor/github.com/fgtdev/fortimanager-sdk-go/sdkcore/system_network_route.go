package fmgclient

import (
	"fmt"
)

// JSONSysNetworkRoute contains the params for creating or updating network route
type JSONSysNetworkRoute struct {
	Device  string `json:"device"`
	Dst     string `json:"dst"`
	Gateway string `json:"gateway"`
	SeqNum  string `json:"seq_num"`
}

// CreateUpdateSystemNetworkRoute is for creating or updating network route
// Input:
//   @params: infor needed
//   @method: operation method, "add" or "update"
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) CreateUpdateSystemNetworkRoute(params *JSONSysNetworkRoute, method string) (err error) {
	defer c.Trace("CreateUpdateSystemNetworkRoute")()

	p := map[string]interface{}{
		"data": *params,
		"url":  "/cli/global/system/route",
	}

	_, err = c.Do(method, p)

	if err != nil {
		err = fmt.Errorf("CreateUpdateSystemNetworkRoute failed: %s", err)
		return
	}

	return
}

// ReadSystemNetworkRoute is for reading the specific network route
// Input:
//   @id: network route id
// Output:
//   @out: network route infor
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) ReadSystemNetworkRoute(id string) (out *JSONSysNetworkRoute, err error) {
	defer c.Trace("ReadSystemNetworkRoute")()

	p := map[string]interface{}{
		"url": "/cli/global/system/route/" + id,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadSystemNetworkRoute failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONSysNetworkRoute{}
	if data["device"] != nil {
		out.Device = data["device"].(string)
	}
	if data["gateway"] != nil {
		out.Gateway = data["gateway"].(string)
	}
	if data["dst"] != nil {
		ip := data["dst"].([]interface{})
		out.Dst = ip[0].(string) + " " + ip[1].(string)
	}

	return
}

// DeleteSystemNetworkRoute is for deleting the specific network route
// Input:
//   @id: network route id
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) DeleteSystemNetworkRoute(id string) (err error) {
	defer c.Trace("DeleteSystemNetworkRoute")()

	p := map[string]interface{}{
		"url": "/cli/global/system/route/" + id,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteSystemNetworkRoute failed :%s", err)
		return
	}

	return
}
