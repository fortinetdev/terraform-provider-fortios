package fortimngclient

import (
	"fmt"
)

type JSONSystemNTP struct {
	Id           int    `json:"id"`
	Server       string `json:"server"`
	Status       string `json:"status,omitempty"`
	SyncInterval int    `json:"sync_interval,omitempty"`
}

func (c *FortiMngClient) SetSystemNTP(params *JSONSystemNTP) (err error) {
	defer c.Trace("SetSystemNTP")()

	d := map[string]interface{}{
		"ntpserver": map[string]interface{}{
			"id":     1,
			"server": params.Server,
		},
		"status":        params.Status,
		"sync_interval": params.SyncInterval,
	}

	p := map[string]interface{}{
		"data": d,
		"url":  "/cli/global/system/ntp",
	}

	_, err = c.Do("set", p)

	if err != nil {
		err = fmt.Errorf("SetSystemNTP failed: %s", err)
		return
	}

	return
}

func (c *FortiMngClient) ReadSystemNTP() (out *JSONSystemNTP, err error) {
	defer c.Trace("ReadSystemNTP")()

	p := map[string]interface{}{
		"url": "/cli/global/system/ntp",
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadSystemNTP failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONSystemNTP{}
	if data["status"] != nil {
		out.Status = c.ControlSwitch2Str(int(data["status"].(float64)))
	}
	if data["sync_interval"] != nil {
		out.SyncInterval = int(data["sync_interval"].(float64))
	}

	ntpSrv := (data["ntpserver"].([]interface{}))[0].(map[string]interface{})
	if ntpSrv["server"] != nil {
		out.Server = ntpSrv["server"].(string)
	}

	return
}
