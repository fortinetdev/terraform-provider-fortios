package fmgclient

import (
	"fmt"

	"github.com/fortinetdev/forti-sdk-go/fortimanager/util"
)

// JSONFirewallObjectIppool contains the params for creating firewall object ippool
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

// CreateUpdateFirewallObjectIppool is for creating/updating the firewall object ippool
// Input:
//   @params: infor needed
//   @method: operation method, "add" or "update"
//   @adom: adom
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) CreateUpdateFirewallObjectIppool(params *JSONFirewallObjectIppool, method, adom string) (err error) {
	defer c.Trace("CreateUpdateFirewallObjectIppool")()

	p := map[string]interface{}{
		"data": params,
		"url":  "/pm/config/adom/" + adom + "/obj/firewall/ippool",
	}

	_, err = c.Do(method, p)

	if err != nil {
		err = fmt.Errorf("CreateUpdateFirewallObjectIppool failed: %s", err)
		return
	}

	return
}

// ReadFirewallObjectIppool is for reading the specific firewall object ippool
// Input:
//   @name: firewall object ippool name
//   @adom: adom
// Output:
//   @out: firewall object ippool infor
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) ReadFirewallObjectIppool(adom, name string) (out *JSONFirewallObjectIppool, err error) {
	defer c.Trace("ReadFirewallObjectIppool")()

	p := map[string]interface{}{
		"url": "/pm/config/adom/" + adom + "/obj/firewall/ippool/" + name,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadFirewallObjectIppool failed: %s", err)
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

// DeleteFirewallObjectIppool is for deleting the specific firewall object ippool
// Input:
//   @adom: adom
//   @name: firewall object ippool name
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) DeleteFirewallObjectIppool(adom, name string) (err error) {
	defer c.Trace("DeleteFirewallObjectIppool")()

	p := map[string]interface{}{
		"url": "/pm/config/adom/" + adom + "/obj/firewall/ippool/" + name,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteFirewallObjectIppool failed: %s", err)
		return
	}

	return
}
