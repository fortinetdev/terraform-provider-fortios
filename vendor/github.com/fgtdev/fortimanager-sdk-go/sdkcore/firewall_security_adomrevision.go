package fortimngclient

import (
	"fmt"
	"strconv"
)

type JSONFirewallSecurityAdomRevision struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	CreatedBy   string `json:"created_by"`
	Locked      int    `json:"locked"`
	Version     string `json:"version"`
}

// Create and Update function
func (c *FortiMngClient) CreateUpdateFirewallSecurityAdomRevision(params *JSONFirewallSecurityAdomRevision, method string) (version string, err error) {
	defer c.Trace("CreateUpdateFirewallSecurityAdomRevision")()

	p := map[string]interface{}{
		"data": params,
		"url":  "/dvmdb/adom/root/revision",
	}

	result, err := c.Do(method, p)

	if err != nil {
		err = fmt.Errorf("CreateUpdateFirewallSecurityAdomRevision failed: %s", err)
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

func (c *FortiMngClient) ReadFirewallSecurityAdomRevision(version string) (out *JSONFirewallSecurityAdomRevision, err error) {
	defer c.Trace("ReadFirewallSecurityAdomRevision")()

	p := map[string]interface{}{
		"url": "/dvmdb/adom/root/revision/" + version,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadFirewallSecurityAdomRevision failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONFirewallSecurityAdomRevision{}
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

func (c *FortiMngClient) DeleteFirewallSecurityAdomRevision(version string) (err error) {
	defer c.Trace("DeleteFirewallSecurityAdomRevision")()

	p := map[string]interface{}{
		"url": "/dvmdb/adom/root/revision/" + version,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteFirewallSecurityAdomRevision failed :%s", err)
		return
	}

	return
}
