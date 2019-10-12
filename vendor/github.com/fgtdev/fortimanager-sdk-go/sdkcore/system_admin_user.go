package fmgclient

import (
	"fmt"

	"github.com/fgtdev/fortimanager-sdk-go/util"
)

// JSONSysAdminUser contains the params for creating system admin user
type JSONSysAdminUser struct {
	UserId      string `json:"userid"`
	Passwd      string `json:"password"`
	Description string `json:"description"`
	UserType    string `json:"user_type"`
	ProfileId   string `json:"profileid"`
	RpcPermit   string `json:"rpc-permit"`
	Trusthost1  string `json:"trusthost1"`
	Trusthost2  string `json:"trusthost2"`
	Trusthost3  string `json:"trusthost3"`
}

// CreateUpdateSystemAdminUser is for creating or updating the system admin user
// Input:
//   @params: infor needed
//   @method: operation method, "add" or "update"
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) CreateUpdateSystemAdminUser(params *JSONSysAdminUser, method string) (err error) {
	defer c.Trace("CreateUpdateSystemAdminUser")()

	p := map[string]interface{}{
		"data": *params,
		"url":  "/cli/global/system/admin/user",
	}

	_, err = c.Do(method, p)

	if err != nil {
		err = fmt.Errorf("CreateUpdateSystemAdminUser failed: %s", err)
		return
	}

	return
}

// ReadSystemAdminUser is for reading the specific system admin user
// Input:
//   @id: admin user name
// Output:
//   @out: admin user infor
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) ReadSystemAdminUser(id string) (out *JSONSysAdminUser, err error) {
	defer c.Trace("ReadSystemAdminUser")()

	p := map[string]interface{}{
		"url": "/cli/global/system/admin/user/" + id,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadSystemAdminUser failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONSysAdminUser{}
	if data["userid"] != nil {
		out.UserId = data["userid"].(string)
	}
	if data["description"] != nil {
		out.Description = data["description"].(string)
	}
	if data["user_type"] != nil {
		out.UserType = util.UserType2Str(int(data["user_type"].(float64)))
	}
	if data["profileid"] != nil {
		out.ProfileId = data["profileid"].(string)
	}
	if data["rpc-permit"] != nil {
		out.RpcPermit = util.RpcPermit2Str(int(data["rpc-permit"].(float64)))
	}
	if data["trusthost1"] != nil {
		hosts := data["trusthost1"].([]interface{})
		out.Trusthost1 = hosts[0].(string) + " " + hosts[1].(string)
	}
	if data["trusthost2"] != nil {
		hosts := data["trusthost2"].([]interface{})
		out.Trusthost2 = hosts[0].(string) + " " + hosts[1].(string)
	}
	if data["trusthost3"] != nil {
		hosts := data["trusthost3"].([]interface{})
		out.Trusthost3 = hosts[0].(string) + " " + hosts[1].(string)
	}

	return
}

// DeleteSystemAdminUser is for deleting the specific system admin user
// Input:
//   @id: admin user name
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) DeleteSystemAdminUser(id string) (err error) {
	defer c.Trace("DeleteSystemAdminUser")()

	p := map[string]interface{}{
		"url": "/cli/global/system/admin/user/" + id,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteSystemAdminUser failed :%s", err)
		return
	}

	return
}
