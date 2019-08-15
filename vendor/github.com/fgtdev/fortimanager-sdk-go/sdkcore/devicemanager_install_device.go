package fortimngclient

import (
	"fmt"
)

type JSONDVMInstallDev struct {
	Name string `json:"name"`
}

func (c *FortiMngClient) CreateDVMInstallDev(params *JSONDVMInstallDev) (err error) {
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

	_, err = c.Do("exec", p)

	if err != nil {
		err = fmt.Errorf("CreateDVMInstallDev failed: %s", err)
		return
	}

	return
}
