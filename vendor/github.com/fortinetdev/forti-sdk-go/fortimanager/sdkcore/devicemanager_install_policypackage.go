package fmgclient

import (
	"fmt"
)

// JSONDVMInstallPolicyPackage contains params for installing policy package
type JSONDVMInstallPolicyPackage struct {
	Name    string `json:"name"`
	Adom    string `json:"adom"`
	Timeout int
}

// CreateDVMInstallPolicyPackage is for installing policy package to the related device
// Input:
//   @params: infor needed
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) CreateDVMInstallPolicyPackage(params *JSONDVMInstallPolicyPackage) (err error) {
	defer c.Trace("CreateDVMInstallPolicyPackage")()

	d := map[string]interface{}{
		"adom": params.Adom,
		"pkg":  params.Name,
	}

	p := map[string]interface{}{
		"data": d,
		"url":  "/securityconsole/install/package",
	}

	result, err := c.Do("exec", p)

	if err != nil {
		err = fmt.Errorf("CreateDVMInstallPolicyPackage failed: %s", err)
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
		err = fmt.Errorf("cannot get the task id after installing the target")
		return
	}

	return
}
