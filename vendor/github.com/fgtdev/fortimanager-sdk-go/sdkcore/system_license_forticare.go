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
		"action": "POST",
		"payload": map[string]string{
			"registration_code": params.RegistrationCode,
		},
		"resource": "/api/v2/monitor/registration/forticare/add-license",
		"target":   params.Target,
	}

	p := map[string]interface{}{
		"data": d,
		"url":  "/sys/proxy/json",
	}

	_, err = c.Do("exec", p)

	if err != nil {
		err = fmt.Errorf("AddSystemLicenseFortiCare failed: %s", err)
		return
	}

	return
}
