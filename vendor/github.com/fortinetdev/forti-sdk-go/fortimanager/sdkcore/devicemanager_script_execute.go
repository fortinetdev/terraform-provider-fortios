package fmgclient

import (
	"fmt"
)

// JSONDVMScriptExecute contains params for executing script
type JSONDVMScriptExecute struct {
	ScriptName    string `json:"script"`
	TargetDevName string `json:"name"`
	Package       string `json:"package"`
	Vdom          string `json:"vdom"`
	Timeout       int
}

// ExecuteDVMScript is for executing the script
// Input:
//   @adom: adom
//   @params: infor needed
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) ExecuteDVMScript(adom string, params *JSONDVMScriptExecute) (err error) {
	defer c.Trace("ExecuteDVMScript")()

	d := map[string]interface{}{
		"adom": adom,
		"scope": map[string]string{
			"name": params.TargetDevName,
			"vdom": params.Vdom,
		},
		"script": params.ScriptName,
		"package": params.Package,
	}

	p := map[string]interface{}{
		"data": d,
		"url":  "/dvmdb/adom/" + adom + "/script/execute",
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
