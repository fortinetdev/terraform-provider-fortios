package fmgclient

import (
	"fmt"
)

// JSONSystemDNS contains the params for updating DNS setting
type JSONSystemDNS struct {
	Primary   string `json:"primary,omitempty"`
	Secondary string `json:"secondary,omitempty"`
}

// SetSystemDNS is for updating DNS setting
// Input:
//   @params: infor needed
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) SetSystemDNS(params *JSONSystemDNS) (err error) {
	defer c.Trace("SetSystemDNS")()

	p := map[string]interface{}{
		"data": *params,
		"url":  "/cli/global/system/dns",
	}

	_, err = c.Do("set", p)

	if err != nil {
		err = fmt.Errorf("SetSystemDNS failed: %s", err)
		return
	}

	return
}

// ReadSystemDNS is for reading DNS setting
// Output:
//   @out: DNS infor
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) ReadSystemDNS() (out *JSONSystemDNS, err error) {
	defer c.Trace("ReadSystemDNS")()

	p := map[string]interface{}{
		"url": "/cli/global/system/dns",
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadSystemDNS failed: %s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONSystemDNS{}
	if data["primary"] != nil {
		out.Primary = data["primary"].(string)
	}
	if data["secondary"] != nil {
		out.Secondary = data["secondary"].(string)
	}

	return
}
