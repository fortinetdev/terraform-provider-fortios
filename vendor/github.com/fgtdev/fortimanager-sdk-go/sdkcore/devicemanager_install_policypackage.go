package fortimngclient

import (
	"fmt"
)

type JSONDVMInstallPolicyPackage struct {
	Name string `json:"name"`
}

func (c *FortiMngClient) CreateDVMInstallPolicyPackage(params *JSONDVMInstallPolicyPackage) (err error) {
	defer c.Trace("CreateDVMInstallPolicyPackage")()

	d := map[string]interface{}{
		"adom": "root",
		"pkg":  params.Name,
	}

	p := map[string]interface{}{
		"data": d,
		"url":  "/securityconsole/install/package",
	}

	_, err = c.Do("exec", p)

	if err != nil {
		err = fmt.Errorf("CreateDVMInstallPolicyPackage failed: %s", err)
		return
	}

	return
}
