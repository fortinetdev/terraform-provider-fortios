package fmgclient

import (
	"fmt"

	"github.com/fgtdev/fortimanager-sdk-go/util"
)

type JSONFirewallObjectIppool struct {
	Name           string `json:"name"`
	Comment        string `json:"comments"`
	Type           string `json:"type"`
	ArpIntf        string `json:"arp-intf"`
	ArpReply       string `json:"arp-reply"`
	AssociatedIntf string `json:"associated-interface"`
	StartIp        string `json:"startip"`
	EndIp          string `json:"endip"`
}

// Create and Update function
func (c *FmgSDKClient) CreateUpdateFirewallObjectIppool(params *JSONFirewallObjectIppool, method string) (err error) {
	defer c.Trace("CreateUpdateFirewallObjectIppool")()

	p := map[string]interface{}{
		"data": params,
		"url":  "/pm/config/adom/root/obj/firewall/ippool",
	}

	_, err = c.Do(method, p)

	if err != nil {
		err = fmt.Errorf("CreateUpdateFirewallObjectIppool failed: %s", err)
		return
	}

	return
}

func (c *FmgSDKClient) ReadFirewallObjectIppool(name string) (out *JSONFirewallObjectIppool, err error) {
	defer c.Trace("ReadFirewallObjectIppool")()

	p := map[string]interface{}{
		"url": "/pm/config/adom/root/obj/firewall/ippool/" + name,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadFirewallObjectIppool failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONFirewallObjectIppool{}
	if data["name"] != nil {
		out.Name = data["name"].(string)
	}
	if data["type"] != nil {
		out.Type = util.FirewallObjectIppoolType2Str(int(data["type"].(float64)))
	}
	if data["comments"] != nil {
		out.Comment = data["comments"].(string)
	}
	if data["startip"] != nil {
		out.StartIp = data["startip"].(string)
	}
	if data["endip"] != nil {
		out.EndIp = data["endip"].(string)
	}
	if data["arp-intf"] != nil {
		m := util.InterfaceArray2StrArray(data["arp-intf"].([]interface{}))
		// only 1 item is allowed here
		out.ArpIntf = m[0]
	}
	if data["associated-interface"] != nil {
		m := util.InterfaceArray2StrArray(data["associated-interface"].([]interface{}))
		// only 1 item is allowed here
		out.ArpIntf = m[0]
	}
	if data["arp-reply"] != nil {
		out.ArpReply = util.ControlSwitch2Str(int(data["arp-reply"].(float64)))
	}

	return

}

func (c *FmgSDKClient) DeleteFirewallObjectIppool(name string) (err error) {
	defer c.Trace("DeleteFirewallObjectIppool")()

	p := map[string]interface{}{
		"url": "/pm/config/adom/root/obj/firewall/ippool/" + name,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteFirewallObjectIppool failed :%s", err)
		return
	}

	return
}
