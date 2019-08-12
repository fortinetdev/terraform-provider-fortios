package fortimngclient

import (
	"fmt"
)

type JSONSysAdminUser struct {
	UserId      string `json:"userid"`
	Passwd      string `json:"password"`
	Description string `json:"description"`
	UserType    string `json:"user_type"`
	ProfileId   string `json:"profileid"`
	RpcPermit   string `json:"rpc-permit"`
	Trusthost1  string `json:"trusthost1"`
	Trusthost2  string `json:"trusthost2"`
}

// Create and Update function
func (c *FortiMngClient) CreateUpdateSystemAdminUser(params *JSONSysAdminUser, method string) (err error) {
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

func (c *FortiMngClient) ReadSystemAdminUser(id string) (out *JSONSysAdminUser, err error) {
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
		out.UserType = c.UserType2Str(int(data["user_type"].(float64)))
	}

	if data["profileid"] != nil {
		out.ProfileId = data["profileid"].(string)
	}

	if data["rpc_permit"] != nil {
		out.RpcPermit = c.AccessRight2Str(int(data["rpc_permit"].(float64)))
	}
	// []string here, should be modified.....
	if data["trusthost1"] != nil {
		hosts := data["trusthost1"].([]interface{})
		out.Trusthost1 = hosts[0].(string) + " " + hosts[1].(string)
	}
	if data["trusthost2"] != nil {
		hosts := data["trusthost2"].([]interface{})
		out.Trusthost2 = hosts[0].(string) + " " + hosts[1].(string)
	}
	// for password: encrypted? ...

	return
}

func (c *FortiMngClient) DeleteSystemAdminUser(id string) (err error) {
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
