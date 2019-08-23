package fortimngclient

import (
	"fmt"
)

type JSONSystemDNSSetting struct {
	Primary   string `json:"primary,omitempty"`
	Secondary string `json:"secondary,omitempty"`
}

func (c *FortiMngClient) SetSystemDNSSetting(params *JSONSystemDNSSetting) (err error) {
	defer c.Trace("SetSystemDNSSetting")()

	p := map[string]interface{}{
		"data": *params,
		"url":  "/cli/global/system/dns",
	}

	_, err = c.Do("set", p)

	if err != nil {
		err = fmt.Errorf("SetSystemDNSSetting failed: %s", err)
		return
	}

	return
}

func (c *FortiMngClient) ReadSystemDNSSetting() (out *JSONSystemDNSSetting, err error) {
	defer c.Trace("ReadSystemDNSSetting")()

	p := map[string]interface{}{
		"url": "/cli/global/system/dns",
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadSystemDNSSetting failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONSystemDNSSetting{}
	if data["primary"] != nil {
		out.Primary = data["primary"].(string)
	}
	if data["secondary"] != nil {
		out.Secondary = data["secondary"].(string)
	}

	return
}
