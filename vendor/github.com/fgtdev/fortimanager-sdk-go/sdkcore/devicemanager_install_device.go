package fmgclient

import (
	"fmt"
)

type JSONDVMInstallDev struct {
	Name    string `json:"name"`
	Timeout int
}

func (c *FmgSDKClient) CreateDVMInstallDev(params *JSONDVMInstallDev) (err error) {
	defer c.Trace("CreateDVMInstallDev")()

	d := map[string]interface{}{
		"adom": "root",
		"scope": map[string]string{
			"name": params.Name,
			"vdom": "root",
		},
	}

	p := map[string]interface{}{
		"data": d,
		"url":  "/securityconsole/install/device",
	}

	result, err := c.Do("exec", p)

	if err != nil {
		err = fmt.Errorf("CreateDVMInstallDev failed: %s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	if data["task"] != nil {
		taskid := int(data["task"].(float64))
		err = c.QueryTask(taskid, params.Timeout)
		if err != nil {
			return err
		}
	} else {
		err = fmt.Errorf("cannot get the task id after executing the script")
		return
	}

	return
}
