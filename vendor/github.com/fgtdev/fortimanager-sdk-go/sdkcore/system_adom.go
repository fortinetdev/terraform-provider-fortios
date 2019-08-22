package fortimngclient

import (
	"fmt"
	"strconv"
)

type JSONSystemAdom struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	State       string `json:"state"`
}

// Create and Update function
func (c *FortiMngClient) CreateUpdateSystemAdom(params *JSONSystemAdom, method string) (err error) {
	defer c.Trace("CreateUpdateSystemAdom")()

	p := map[string]interface{}{
		"data": params,
		"url":  "/dvmdb/adom",
	}

	_, err = c.Do(method, p)

	if err != nil {
		err = fmt.Errorf("CreateUpdateSystemAdom failed: %s", err)
		return
	}

	return
}

func (c *FortiMngClient) ReadSystemAdom(name string) (out *JSONSystemAdom, err error) {
	defer c.Trace("ReadSystemAdom")()

	p := map[string]interface{}{
		"url": "/dvmdb/adom/" + name,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadSystemAdom failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONSystemAdom{}
	if data["name"] != nil {
		out.Name = data["name"].(string)
	}
	if data["desc"] != nil {
		out.Description = data["desc"].(string)
	}
	if data["state"] != nil {
		out.State = strconv.Itoa(int(data["state"].(float64)))
	}

	return

}

func (c *FortiMngClient) DeleteSystemAdom(name string) (err error) {
	defer c.Trace("DeleteSystemAdom")()

	p := map[string]interface{}{
		"url": "/dvmdb/adom/" + name,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteSystemAdom failed :%s", err)
		return
	}

	return
}
