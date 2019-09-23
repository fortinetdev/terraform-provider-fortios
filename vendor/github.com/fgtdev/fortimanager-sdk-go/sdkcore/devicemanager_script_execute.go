package fmgclient

import (
	"fmt"
)

type JSONDVMScriptExecute struct {
	ScriptName    string `json:"script"`
	TargetDevName string `json:"name"`
	Timeout       int
}

// Create and Update function
func (c *FmgSDKClient) ExecuteDVMScript(params *JSONDVMScriptExecute) (err error) {
	defer c.Trace("ExecuteDVMScript")()

	d := map[string]interface{}{
		"adom": "root",
		"scope": map[string]string{
			"vdom": "root",
			"name": params.TargetDevName,
		},
		"script": params.ScriptName,
	}

	p := map[string]interface{}{
		"data": d,
		"url":  "/dvmdb/adom/root/script/execute",
	}

	result, err := c.Do("exec", p)

	if err != nil {
		err = fmt.Errorf("ExecuteDVMScript failed: %s", err)
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
