package fmgclient

import (
	"fmt"
)

// JSONDVMDeviceCreate contains params for adding device to devicemanager
type JSONDVMDeviceCreate struct {
	UserId   string `json:"adm_usr"`
	Passwd   string `json:"adm_pass"`
	Ipaddr   string `json:"ip"`
	Name     string `json:"name"`
	MgmtMode string `json:"mgmt_mode"`
}

// JSONDVMDeviceDel contains params for deleting device from devicemanager
type JSONDVMDeviceDel struct {
	Name string `json:"device"`
	Adom string `json:"adom"`
}

// CreateDVMDevice add device to devicemanager
// Input:
//   @params: infor needed to add the device
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) CreateDVMDevice(adom string, params *JSONDVMDeviceCreate) (err error) {
	defer c.Trace("CreateDVMDevice")()

	data := map[string]interface{}{
		"adom":   adom,
		"device": params,
	}

	p := map[string]interface{}{
		"data": data,
		"url":  "/dvm/cmd/add/device",
	}

	_, err = c.Do("exec", p)

	if err != nil {
		return fmt.Errorf("CreateDVMDevice failed: %s", err)
	}

	return
}

// DeleteDVMDevice delelte device from devicemanager
// Input:
//   @params: infor needed to delete the device
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) DeleteDVMDevice(params *JSONDVMDeviceDel) (err error) {
	defer c.Trace("DeleteDVMDevice")()

	p := map[string]interface{}{
		"data": params,
		"url":  "/dvm/cmd/del/device",
	}

	_, err = c.Do("exec", p)

	if err != nil {
		return fmt.Errorf("DeleteDVMDevice failed: %s", err)
	}

	return
}
