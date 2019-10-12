package fmgclient

import (
	"fmt"
)

// JSONSystemAdmin contains the params for updating system admin setting
type JSONSystemAdmin struct {
	HttpPort    int `json:"http_port,omitempty"`
	HttpsPort   int `json:"https_port,omitempty"`
	IdleTimeout int `json:"idle_timeout,omitempty"`
}

// SetSystemAdmin is for updating the system admin setting
// Input:
//   @params: infor needed
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) SetSystemAdmin(params *JSONSystemAdmin) (err error) {
	defer c.Trace("SetSystemAdmin")()

	p := map[string]interface{}{
		"data": *params,
		"url":  "/cli/global/system/admin/setting",
	}

	_, err = c.Do("set", p)

	if err != nil {
		err = fmt.Errorf("SetSystemAdmin failed: %s", err)
		return
	}

	return
}

// ReadSystemAdmin is for reading the system admin setting
// Output:
//   @out: system admin infor
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) ReadSystemAdmin() (out *JSONSystemAdmin, err error) {
	defer c.Trace("ReadSystemAdmin")()

	p := map[string]interface{}{
		"url": "/cli/global/system/admin/setting",
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadSystemAdmin failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONSystemAdmin{}
	if data["http_port"] != nil {
		out.HttpPort = int(data["http_port"].(float64))
	}
	if data["https_port"] != nil {
		out.HttpsPort = int(data["https_port"].(float64))
	}
	if data["idle_timeout"] != nil {
		out.IdleTimeout = int(data["idle_timeout"].(float64))
	}

	return
}
