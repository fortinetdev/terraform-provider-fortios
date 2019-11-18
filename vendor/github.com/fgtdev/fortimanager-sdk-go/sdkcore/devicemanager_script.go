package fmgclient

import (
	"fmt"

	"github.com/fgtdev/fortimanager-sdk-go/util"
)

// JSONDVMScript contains the params for creating the script
type JSONDVMScript struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	Content     string `json:"content"`
	Target      string `json:"target"`
	Type        string `json:"type"`
}

// CreateUpdateDVMScript is for creating/updating the script
// Input:
//   @params: infor needed
//   @method: operation method, "add" or "update"
//   @adom: adom
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) CreateUpdateDVMScript(params *JSONDVMScript, method, adom string) (err error) {
	defer c.Trace("CreateUpdateDVMScript")()

	p := map[string]interface{}{
		"data": *params,
		"url":  "/dvmdb/adom/" + adom + "/script",
	}

	_, err = c.Do(method, p)

	if err != nil {
		err = fmt.Errorf("CreateUpdateDVMScript failed: %s", err)
		return
	}

	return
}

// ReadDVMScript is for reading the script infor
// Input:
//   @adom: adom
//   @id: script name
// Output:
//   @out: script infor
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) ReadDVMScript(adom, id string) (out *JSONDVMScript, err error) {
	defer c.Trace("ReadDVMScript")()

	p := map[string]interface{}{
		"url": "/dvmdb/adom/" + adom + "/script/" + id,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadDVMScript failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONDVMScript{}
	if data["name"] != nil {
		out.Name = data["name"].(string)
	}
	if data["desc"] != nil {
		out.Description = data["desc"].(string)
	}
	if data["content"] != nil {
		out.Content = data["content"].(string)
	}
	if data["target"] != nil {
		out.Target = util.ScriptTarget2Str(int(data["target"].(float64)))
	}

	return
}

// DeleteDVMScript is for deleting the related script
// Input:
//   @adom: adom
//   @id: script name
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) DeleteDVMScript(adom, id string) (err error) {
	defer c.Trace("DeleteDVMScript")()

	p := map[string]interface{}{
		"url": "/dvmdb/adom/" + adom + "/script/" + id,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteDVMScript failed :%s", err)
		return
	}

	return
}
