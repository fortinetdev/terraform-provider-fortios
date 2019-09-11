package fortimngclient

import (
	"fmt"
)

type JSONFirewallObjectAddress struct {
	Name           string `json:"name"`
	Type           string `json:"type"`
	Comment        string `json:"comment"`
	Fqdn           string `json:"fqdn"`
	AssociatedIntf string `json:"associated-interface"`
	Subnet         string `json:"subnet"`
	StartIp        string `json:"start-ip"`
	EndIp          string `json:"end-ip"`
	AllowRouting   string `json:"allow-routing"`
}

// Create and Update function
func (c *FortiMngClient) CreateUpdateFirewallObjectAddress(params *JSONFirewallObjectAddress, method string) (err error) {
	defer c.Trace("CreateUpdateFirewallObjectAddress")()

	p := map[string]interface{}{
		"data": params,
		"url":  "/pm/config/adom/root/obj/firewall/address",
	}

	_, err = c.Do(method, p)

	if err != nil {
		err = fmt.Errorf("CreateUpdateFirewallObjectAddress failed: %s", err)
		return
	}

	return
}

func (c *FortiMngClient) ReadFirewallObjectAddress(name string) (out *JSONFirewallObjectAddress, err error) {
	defer c.Trace("ReadFirewallObjectAddress")()

	p := map[string]interface{}{
		"url": "/pm/config/adom/root/obj/firewall/address/" + name,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadFirewallObjectAddress failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONFirewallObjectAddress{}
	if data["name"] != nil {
		out.Name = data["name"].(string)
	}
	if data["type"] != nil {
		out.Type = c.FirewallObjectAddrType2Str(int(data["type"].(float64)))
	}
	if data["comment"] != nil {
		out.Comment = data["comment"].(string)
	}
	if data["fqdn"] != nil {
		out.Fqdn = data["fqdn"].(string)
	}

	if data["associated-interface"] != nil {
		m := c.InterfaceArray2StrArray(data["associated-interface"].([]interface{}))
		// only 1 item is allowed here
		out.AssociatedIntf = m[0]
	}
	if data["subnet"] != nil {
		sn := data["subnet"].([]interface{})
		out.Subnet = sn[0].(string) + " " + sn[1].(string)
	}
	if data["start-ip"] != nil {
		out.StartIp = data["start-ip"].(string)
	}
	if data["end-ip"] != nil {
		out.EndIp = data["end-ip"].(string)
	}
	if data["allow-routing"] != nil {
		out.AllowRouting = c.ControlSwitch2Str(int(data["allow-routing"].(float64)))
	}

	return
}

func (c *FortiMngClient) DeleteFirewallObjectAddress(name string) (err error) {
	defer c.Trace("DeleteFirewallObjectAddress")()

	p := map[string]interface{}{
		"url": "/pm/config/adom/root/obj/firewall/address/" + name,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteFirewallObjectAddress failed :%s", err)
		return
	}

	return
}
