package fortimngclient

import (
	"fmt"
)

type JSONFirewallSecurityPolicyPackage struct {
	Name   string `json:"name"`
	Target string `json:"name"`
}

func (c *FortiMngClient) CreateUpdateFirewallSecurityPolicyPackage(params *JSONFirewallSecurityPolicyPackage, method string) (err error) {
	defer c.Trace("CreateUpdateFirewallSecurityPolicyPackage")()

	d := map[string]interface{}{
		"name": params.Name,
		"scope member": map[string]string{
			"name": params.Target,
			"vdom": "root",
		},
		"type": "pkg",
	}

	p := map[string]interface{}{
		"data": d,
		"url":  "/pm/pkg/adom/root",
	}

	_, err = c.Do(method, p)

	if err != nil {
		return fmt.Errorf("AddFirewallSecurityPolicyPackage failed: %s", err)
	}

	return
}

func (c *FortiMngClient) DeleteFirewallSecurityPolicyPackage(name string) (err error) {
	defer c.Trace("DeleteFirewallSecurityPolicyPackage")()

	p := map[string]interface{}{
		"url": "/pm/pkg/adom/root/" + name,
	}

	_, err = c.Do("delete", p)

	if err != nil {
		return fmt.Errorf("DeleteFirewallSecurityPolicyPackage failed: %s", err)
	}

	return
}

func (c *FortiMngClient) ReadFirewallSecurityPolicyPackage(name string) (out *JSONFirewallSecurityPolicyPackage, err error) {
	defer c.Trace("ReadFirewallSecurityPolicyPackage")()

	p := map[string]interface{}{
		"url": "/pm/pkg/adom/root/" + name,
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

	sm := (data["scope member"].([]interface{}))[0].(map[string]interface{})
	if sm["name"] != nil {
		out.Target = sm["name"].(string)
	}

	return
}
