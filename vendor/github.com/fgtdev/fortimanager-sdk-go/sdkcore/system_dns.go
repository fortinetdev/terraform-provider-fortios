package fmgclient

import (
	"fmt"
)

type JSONSystemDNS struct {
	Primary   string `json:"primary,omitempty"`
	Secondary string `json:"secondary,omitempty"`
}

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

func (c *FmgSDKClient) ReadSystemDNS() (out *JSONSystemDNS, err error) {
	defer c.Trace("ReadSystemDNS")()

	p := map[string]interface{}{
		"url": "/cli/global/system/dns",
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadSystemDNS failed :%s", err)
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
