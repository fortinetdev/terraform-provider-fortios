package fortimngclient

import (
	"fmt"
)

type JSONSystemGlobal struct {
	Hostname   string `json:"hostname,omitempty"`
	FazStatus  string `json:"faz-status,omitempty"`
	AdomStatus string `json:"adom-status,omitempty"`
	AdomMode   string `json:"adom-mode,omitempty"`
	TimeZone   string `json:"timezone,omitempty"`
}

func (c *FortiMngClient) SetSystemGlobal(params *JSONSystemGlobal) (err error) {
	defer c.Trace("SetSystemGlobal")()

	p := map[string]interface{}{
		"data": *params,
		"url":  "/cli/global/system/global",
	}

	_, err = c.Do("set", p)

	if err != nil {
		err = fmt.Errorf("SetSystemGlobal failed: %s", err)
		return
	}

	return
}

func (c *FortiMngClient) ReadSystemGlobal() (out *JSONSystemGlobal, err error) {
	defer c.Trace("ReadSystemGlobal")()

	p := map[string]interface{}{
		"url": "/cli/global/system/global",
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadSystemGlobal failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONSystemGlobal{}
	if data["hostname"] != nil {
		out.Hostname = data["hostname"].(string)
	}
	if data["faz-status"] != nil {
		out.FazStatus = c.ControlSwitch2Str(int(data["faz-status"].(float64)))
	}
	if data["adom-status"] != nil {
		out.AdomStatus = c.ControlSwitch2Str(int(data["adom-status"].(float64)))
	}
	if data["adom-mode"] != nil {
		out.AdomMode = c.AdomMode2Str(int(data["adom-mode"].(float64)))
	}
	if data["timezone"] != nil {
		out.TimeZone = c.TimeZone2Str(int(data["timezone"].(float64)))
	}

	return

}
