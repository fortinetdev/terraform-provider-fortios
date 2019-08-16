package fortimngclient

import (
	"fmt"
	"strconv"
)

type JSONSystemSyslog struct {
	Name string `json:"name"`
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

// Create and Update function
func (c *FortiMngClient) CreateUpdateSystemSyslog(params *JSONSystemSyslog, method string) (err error) {
	defer c.Trace("CreateUpdateSystemSyslog")()

	p := map[string]interface{}{
		"data": *params,
		"url":  "/cli/global/system/syslog",
	}

	_, err = c.Do(method, p)

	if err != nil {
		err = fmt.Errorf("CreateUpdateSystemSyslog failed: %s", err)
		return
	}

	return
}

func (c *FortiMngClient) ReadSystemSyslog(name string) (out *JSONSystemSyslog, err error) {
	defer c.Trace("ReadSystemSyslog")()

	p := map[string]interface{}{
		"url": "/cli/global/system/syslog/" + name,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadSystemSyslog failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONSystemSyslog{}
	if data["name"] != nil {
		out.Name = data["name"].(string)
	}
	if data["ip"] != nil {
		out.Ip = data["ip"].(string)
	}
	if data["port"] != nil {
		out.Port = strconv.Itoa(int(data["port"].(float64)))
	}

	return
}

func (c *FortiMngClient) DeleteSystemSyslog(name string) (err error) {
	defer c.Trace("DeleteSystemSyslog")()

	p := map[string]interface{}{
		"url": "/cli/global/system/syslog/" + name,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteSystemSyslog failed :%s", err)
		return
	}

	return
}
