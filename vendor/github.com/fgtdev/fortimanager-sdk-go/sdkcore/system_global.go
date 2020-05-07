package fmgclient

import (
	"fmt"

	"github.com/fgtdev/fortimanager-sdk-go/util"
)

// JSONSystemGlobal contains the params for updating system global setting
type JSONSystemGlobal struct {
	Hostname   string `json:"hostname,omitempty"`
	FazStatus  string `json:"faz-status,omitempty"`
	AdomStatus string `json:"adom-status,omitempty"`
	AdomMode   string `json:"adom-mode,omitempty"`
	TimeZone   string `json:"timezone,omitempty"`
}

// SetSystemGlobal is for updating the system global setting
// Input:
//   @params: infor needed
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) SetSystemGlobal(params *JSONSystemGlobal) (err error) {
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

// ReadSystemGlobal is for reading the system global setting
// Output:
//   @out: system global infor
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) ReadSystemGlobal() (out *JSONSystemGlobal, err error) {
	defer c.Trace("ReadSystemGlobal")()

	p := map[string]interface{}{
		"url": "/cli/global/system/global",
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadSystemGlobal failed: %s", err)
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
		out.FazStatus = util.ControlSwitch2Str(int(data["faz-status"].(float64)))
	}
	if data["adom-status"] != nil {
		out.AdomStatus = util.ControlSwitch2Str(int(data["adom-status"].(float64)))
	}
	if data["adom-mode"] != nil {
		out.AdomMode = util.AdomMode2Str(int(data["adom-mode"].(float64)))
	}
	if data["timezone"] != nil {
		out.TimeZone = util.TimeZone2Str(int(data["timezone"].(float64)))
	}

	return

}
