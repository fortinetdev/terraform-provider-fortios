package fortimngclient

import (
	"fmt"
)

type SystemLicenseFortiCare struct {
	Target           string
	RegistrationCode string
}

func (c *FortiMngClient) AddSystemLicenseFortiCare(params *SystemLicenseFortiCare) (err error) {
	defer c.Trace("AddSystemLicenseFortiCare")()

	d := map[string]interface{}{
		"action": "post",
		"payload": map[string]string{
			"registration_code": params.RegistrationCode,
		},
		"resource": "/api/v2/monitor/registration/forticare/add-license",
		"target":   []string{"/adom/root/device/" + params.Target},
	}

	p := map[string]interface{}{
		"data": d,
		"url":  "/sys/proxy/json",
	}

	result, err := c.Do("exec", p)
	if err != nil {
		err = fmt.Errorf("AddSystemLicenseFortiCare failed: %s", err)
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
