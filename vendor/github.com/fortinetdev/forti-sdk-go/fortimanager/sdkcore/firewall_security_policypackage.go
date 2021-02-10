package fmgclient

import (
	"fmt"

	"github.com/fortinetdev/forti-sdk-go/fortimanager/util"
)

// JSONFirewallSecurityPolicyPackage contains the params for creating firewall policy package
type JSONFirewallSecurityPolicyPackage struct {
	Name           string `json:"name"`
	Vdom           string `json:"vdom"`
	Target         string `json:"name"`
	InspectionMode string `json:"inspection-mode"`
}

// CreateUpdateFirewallSecurityPolicyPackage is for creating/updating the firewall security policy package
// Input:
//   @params: infor needed
//   @adom: adom
//   @method: operation method, "add" or "update"
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) CreateUpdateFirewallSecurityPolicyPackage(params *JSONFirewallSecurityPolicyPackage, method, adom string) (err error) {
	defer c.Trace("CreateUpdateFirewallSecurityPolicyPackage")()


	d := map[string]interface{}{
		"name": params.Name,
		"package settings": map[string]string{
			"inspection-mode": params.InspectionMode,
		},
		"type": "pkg",
	}

	v := make(map[string]string)

	if params.Target != "" {
		v["name"] = params.Target
	}

	if params.Vdom != "" {
		v["vdom"] = params.Vdom
	}

	if params.Target != "" || params.Vdom != "" {
		d["scope member"] = v
	}

	p := map[string]interface{}{
		"data": d,
		"url":  "/pm/pkg/adom/" + adom,
	}

	_, err = c.Do(method, p)

	if err != nil {
		return fmt.Errorf("AddFirewallSecurityPolicyPackage failed: %s", err)
	}

	return
}

// DeleteFirewallSecurityPolicyPackage is for deleting the specific firewall security policy package
// Input:
//   @adom: adom
//   @name: policy package name
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) DeleteFirewallSecurityPolicyPackage(adom, name string) (err error) {
	defer c.Trace("DeleteFirewallSecurityPolicyPackage")()

	p := map[string]interface{}{
		"url": "/pm/pkg/adom/" + adom + "/" + name,
	}

	_, err = c.Do("delete", p)

	if err != nil {
		return fmt.Errorf("DeleteFirewallSecurityPolicyPackage failed: %s", err)
	}

	return
}

// ReadFirewallSecurityPolicyPackage is for reading the specific firewall security policy package
// Input:
//   @adom: adom
//   @name: policy package name
// Output:
//   @out: policy package infor
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) ReadFirewallSecurityPolicyPackage(adom, name string) (out *JSONFirewallSecurityPolicyPackage, err error) {
	defer c.Trace("ReadFirewallSecurityPolicyPackage")()

	p := map[string]interface{}{
		"url": "/pm/pkg/adom/" + adom + "/" + name,
	}

	result, err := c.Do("get", p)

	if err != nil {
		err = fmt.Errorf("ReadFirewallSecurityPolicyPackage failed: %s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONFirewallSecurityPolicyPackage{}
	if data["name"] != nil {
		out.Name = data["name"].(string)
	}

	if data["scope member"] != nil {
		sm := (data["scope member"].([]interface{}))[0].(map[string]interface{})
		if sm["name"] != nil {
			out.Target = sm["name"].(string)
		}
	}

	ps := data["package settings"].(map[string]interface{})
	if ps["inspection-mode"] != nil {
		out.InspectionMode = util.InspectionMode2Str(int(ps["inspection-mode"].(float64)))
	}

	return
}
