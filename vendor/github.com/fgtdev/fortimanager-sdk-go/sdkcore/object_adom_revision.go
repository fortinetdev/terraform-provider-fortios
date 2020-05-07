package fmgclient

import (
	"fmt"
	"strconv"
)

// JSONObjectAdomRevision contains the params for creating obj adom revision
type JSONObjectAdomRevision struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	CreatedBy   string `json:"created_by"`
	Locked      int    `json:"locked"`
	Version     string `json:"version"`
}

// CreateUpdateObjectAdomRevision is for creating/updating the obj adom revision
// Input:
//   @params: infor needed
//   @method: operation method, "add" or "update"
//   @adom: adom
// Output:
//   @version: adom revision version
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) CreateUpdateObjectAdomRevision(params *JSONObjectAdomRevision, method, adom string) (version string, err error) {
	defer c.Trace("CreateUpdateObjectAdomRevision")()

	p := map[string]interface{}{
		"data": params,
		"url":  "/dvmdb/adom/" + adom + "/revision",
	}

	result, err := c.Do(method, p)

	if err != nil {
		err = fmt.Errorf("CreateUpdateObjectAdomRevision failed: %s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	if data["version"] != nil {
		version = strconv.Itoa(int(data["version"].(float64)))
	} else {
		err = fmt.Errorf("can't get version from response")
	}

	return
}

// ReadObjectAdomRevision is for reading the specific obj adom revision
// Input:
//   @version: adom revision version
//   @adom: adom
// Output:
//   @out: adom revision infor
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) ReadObjectAdomRevision(adom, version string) (out *JSONObjectAdomRevision, err error) {
	defer c.Trace("ReadObjectAdomRevision")()

	p := map[string]interface{}{
		"url": "/dvmdb/adom/" + adom + "/revision/" + version,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadObjectAdomRevision failed: %s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONObjectAdomRevision{}
	if data["name"] != nil {
		out.Name = data["name"].(string)
	}
	if data["locked"] != nil {
		out.Locked = int(data["locked"].(float64))
	}
	if data["desc"] != nil {
		out.Description = data["desc"].(string)
	}
	if data["created_by"] != nil {
		out.CreatedBy = data["created_by"].(string)
	}

	return

}

// DeleteObjectAdomRevision is for deleting the specific obj adom revision
// Input:
//   @adom: adom
//   @version: adom revision version
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) DeleteObjectAdomRevision(adom, version string) (err error) {
	defer c.Trace("DeleteObjectAdomRevision")()

	p := map[string]interface{}{
		"url": "/dvmdb/adom/" + adom + "/revision/" + version,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteObjectAdomRevision failed: %s", err)
		return
	}

	return
}
