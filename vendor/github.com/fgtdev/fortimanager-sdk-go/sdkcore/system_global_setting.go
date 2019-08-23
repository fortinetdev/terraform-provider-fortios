package fortimngclient

import (
	"fmt"
)

type JSONSystemGlobalSetting struct {
	Hostname   string `json:"hostname"`
	FazStatus  string `json:"faz-status"`
	AdomStatus string `json:"adom-status"`
	/*
		AdomMode string `json:"adom-mode"`
		TimeZone string `json:"timezone"`
		Usg string `json:"usg"`
	*/
}

func (c *FortiMngClient) SetSystemGlobalSetting(params *JSONSystemGlobalSetting) (err error) {
	defer c.Trace("SetSystemGlobalSetting")()

	p := map[string]interface{}{
		"data": *params,
		"url":  "/cli/global/system/global",
	}

	_, err = c.Do("set", p)

	if err != nil {
		err = fmt.Errorf("SetSystemGlobalSetting failed: %s", err)
		return
	}

	return
}

func (c *FortiMngClient) ReadSystemGlobalSetting() (out *JSONSystemGlobalSetting, err error) {
	defer c.Trace("ReadSystemGlobalSetting")()

	p := map[string]interface{}{
		"url": "/cli/global/system/global",
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadSystemGlobalSetting failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONSystemGlobalSetting{}
	if data["hostname"] != nil {
		out.Hostname = data["hostname"].(string)
	}
	if data["faz-status"] != nil {
		out.FazStatus = c.ControlSwitch2Str(int(data["faz-status"].(float64)))
	}
	if data["adom-status"] != nil {
		out.AdomStatus = c.ControlSwitch2Str(int(data["adom-status"].(float64)))
	}

	return

}
