package fmgclient

import (
	"fmt"
)

// JSONSystemLicenseVM contains the params for uploading vm license
type JSONSystemLicenseVM struct {
	Target      string
	Adom        string
	FileContent string
}

// AddSystemLicenseVM is for uploading vm license
// Input:
//   @params: infor needed
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) AddSystemLicenseVM(params *JSONSystemLicenseVM) (err error) {
	defer c.Trace("AddSystemLicenseVM")()

	d := map[string]interface{}{
		"action": "post",
		"payload": map[string]string{
			"file_content": params.FileContent,
		},
		"resource": "/api/v2/monitor/system/vmlicense/upload",
		"target":   []string{"/adom/" + params.Adom + "/device/" + params.Target},
	}

	p := map[string]interface{}{
		"data": d,
		"url":  "/sys/proxy/json",
	}

	result, err := c.Do("exec", p)
	if err != nil {
		err = fmt.Errorf("AddSystemLicenseVM failed: %s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].([]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	status := data[0].(map[string]interface{})["status"].(map[string]interface{})
	code := uint(status["code"].(float64))
	msg := status["message"].(string)
	if code != 0 || msg != "OK" {
		err = fmt.Errorf("status not right: code is %d, message is %s", code, msg)
		return
	}

	return
}
