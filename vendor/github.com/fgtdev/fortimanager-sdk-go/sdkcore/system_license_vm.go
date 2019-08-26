package fortimngclient

import (
	"fmt"
)

type JSONSystemLicenseVM struct {
	Target      string `json:"target"`
	FileContent string `json:"file_content"`
}

func (c *FortiMngClient) AddSystemLicenseVM(params *JSONSystemLicenseVM) (err error) {
	defer c.Trace("AddSystemLicenseVM")()

	d := map[string]interface{}{
		"action": "POST",
		"payload": map[string]string{
			"file_content": params.FileContent,
		},
		"resource": "/api/v2/monitor/system/vmlicense/upload",
		"target":   params.Target,
	}

	p := map[string]interface{}{
		"data": d,
		"url":  "/sys/proxy/json",
	}

	_, err = c.Do("exec", p)

	if err != nil {
		err = fmt.Errorf("AddSystemLicenseVM failed: %s", err)
		return
	}

	return
}
