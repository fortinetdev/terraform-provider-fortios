package fmgclient

import (
	"fmt"
)

// JSONSystemSyslogServer contains the params for creating or updating syslog server
type JSONSystemSyslogServer struct {
	Name string `json:"name"`
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}

// CreateUpdateSystemSyslogServer is for creating or updating syslog server
// Input:
//   @params: infor needed
//   @method: operation method, "add" or "update"
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) CreateUpdateSystemSyslogServer(params *JSONSystemSyslogServer, method string) (err error) {
	defer c.Trace("CreateUpdateSystemSyslogServer")()

	p := map[string]interface{}{
		"data": *params,
		"url":  "/cli/global/system/syslog",
	}

	_, err = c.Do(method, p)

	if err != nil {
		err = fmt.Errorf("CreateUpdateSystemSyslogServer failed: %s", err)
		return
	}

	return
}

// ReadSystemSyslogServer is for reading the specific syslog server
// Input:
//   @name: syslog server name
// Output:
//   @out: syslog server infor
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) ReadSystemSyslogServer(name string) (out *JSONSystemSyslogServer, err error) {
	defer c.Trace("ReadSystemSyslogServer")()

	p := map[string]interface{}{
		"url": "/cli/global/system/syslog/" + name,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadSystemSyslogServer failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONSystemSyslogServer{}
	if data["name"] != nil {
		out.Name = data["name"].(string)
	}
	if data["ip"] != nil {
		out.Ip = data["ip"].(string)
	}
	if data["port"] != nil {
		out.Port = int(data["port"].(float64))
	}

	return
}

// DeleteSystemSyslogServer is for deleting the specific syslog server
// Input:
//   @name: syslog server name
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) DeleteSystemSyslogServer(name string) (err error) {
	defer c.Trace("DeleteSystemSyslogServer")()

	p := map[string]interface{}{
		"url": "/cli/global/system/syslog/" + name,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteSystemSyslogServer failed :%s", err)
		return
	}

	return
}
