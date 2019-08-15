package fortimngclient

import (
	"fmt"
	"strconv"
)

type JSONFirewallSecurityPolicy struct {
	Name     string   `json:"name"`
	Action   string   `json:"action"`
	PolicyId string   `json:"policyid"`
	SrcAddr  []string `json:"srcaddr"`
	SrcIntf  []string `json:"srcintf"`
	DstAddr  []string `json:"dstaddr"`
	DstIntf  []string `json:"dstintf"`
	Service  []string `json:"service"`
	Schedule []string `json:"schedule"`
}

type FirewallSecurityPolicyInput struct {
	Policy      *JSONFirewallSecurityPolicy
	PackageName string
}

// Create and Update function
func (c *FortiMngClient) CreateUpdateFirewallSecurityPolicy(params *FirewallSecurityPolicyInput, method string) (id string, err error) {
	defer c.Trace("CreateUpdateFirewallSecurityPolicy")()

	p := map[string]interface{}{
		"data": params.Policy,
		"url":  "/pm/config/adom/root/pkg/" + params.PackageName + "/firewall/policy",
	}

	result, err := c.Do(method, p)

	if err != nil {
		err = fmt.Errorf("CreateUpdateFirewallSecurityPolicy failed: %s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	if data["policyid"] != nil {
		id = strconv.Itoa(int(data["policyid"].(float64)))
	} else {
		err = fmt.Errorf("Get policyid from response failed: %s", err)
		return
	}

	return
}

func (c *FortiMngClient) ReadFirewallSecurityPolicy(id, pkg_name string) (out *JSONFirewallSecurityPolicy, err error) {
	defer c.Trace("ReadFirewallSecurityPolicy")()

	p := map[string]interface{}{
		"url": "/pm/config/adom/root/pkg/" + pkg_name + "/firewall/policy/" + id,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadFirewallSecurityPolicy failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONFirewallSecurityPolicy{}
	if data["name"] != nil {
		out.Name = data["name"].(string)
	}
	if data["action"] != nil {
		out.Action = c.PolicyAction2Str(int(data["action"].(float64)))
	}
	if data["srcaddr"] != nil {
		m := c.InterfaceArray2StrArray(data["srcaddr"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.SrcAddr = append(out.SrcAddr, m[i])
		}
	}
	if data["srcintf"] != nil {
		m := c.InterfaceArray2StrArray(data["srcintf"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.SrcIntf = append(out.SrcIntf, m[i])
		}
	}
	if data["dstaddr"] != nil {
		m := c.InterfaceArray2StrArray(data["dstaddr"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.DstAddr = append(out.DstAddr, m[i])
		}
	}
	if data["dstintf"] != nil {
		m := c.InterfaceArray2StrArray(data["dstintf"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.DstIntf = append(out.DstIntf, m[i])
		}
	}
	if data["service"] != nil {
		m := c.InterfaceArray2StrArray(data["service"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.Service = append(out.Service, m[i])
		}
	}
	if data["schedule"] != nil {
		m := c.InterfaceArray2StrArray(data["schedule"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.Schedule = append(out.Schedule, m[i])
		}
	}

	return
}

func (c *FortiMngClient) DeleteFirewallSecurityPolicy(id, pkg_name string) (err error) {
	defer c.Trace("DeleteFirewallSecurityPolicy")()

	p := map[string]interface{}{
		"url": "/pm/config/adom/root/pkg/" + pkg_name + "/firewall/policy/" + id,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteFirewallSecurityPolicy failed :%s", err)
		return
	}

	return
}
