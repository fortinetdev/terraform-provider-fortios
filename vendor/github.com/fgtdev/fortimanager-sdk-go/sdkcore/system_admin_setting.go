package fortimngclient

import (
	"fmt"
)

type JSONSystemAdminSetting struct {
	HttpPort    int `json:"http_port"`
	HttpsPort   int `json:"https_port"`
	IdleTimeout int `json:"idle_timeout"`
}

func (c *FortiMngClient) SetSystemAdminSetting(params *JSONSystemAdminSetting) (err error) {
	defer c.Trace("SetSystemAdminSetting")()

	p := map[string]interface{}{
		"data": *params,
		"url":  "/cli/global/system/admin/setting",
	}

	_, err = c.Do("set", p)

	if err != nil {
		err = fmt.Errorf("SetSystemAdminSetting failed: %s", err)
		return
	}

	return
}

func (c *FortiMngClient) ReadSystemAdminSetting() (out *JSONSystemAdminSetting, err error) {
	defer c.Trace("ReadSystemAdminSetting")()

	p := map[string]interface{}{
		"url": "/cli/global/system/admin/setting",
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadSystemAdminSetting failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONSystemAdminSetting{}
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
