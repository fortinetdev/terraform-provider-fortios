package fmgclient

import (
	"fmt"
)

// JSONDVMInstallDev contains the params for installing device
type JSONDVMInstallDev struct {
	Name    string `json:"name"`
	Vdom    string `json:"vdom"`
	Timeout int
}

// CreateDVMInstallDev is for installing scripts to the related device
// Input:
//   @params: infor needed
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) CreateDVMInstallDev(adom string, params *JSONDVMInstallDev) (err error) {
	defer c.Trace("CreateDVMInstallDev")()

	d := map[string]interface{}{
		"adom": adom,
		"scope": map[string]string{
			"name": params.Name,
			"vdom": params.Vdom,
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
