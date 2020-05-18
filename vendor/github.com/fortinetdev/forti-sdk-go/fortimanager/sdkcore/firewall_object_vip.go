package fmgclient

import (
	"fmt"

	"github.com/fortinetdev/forti-sdk-go/fortimanager/util"
)

// JSONFirewallObjectVip contains the params for creating firewall object virtual ip
type JSONFirewallObjectVip struct {
	Name          string `json:"name"`
	Comment       string `json:"comment"`
	Type          string `json:"type"`
	ArpReply      string `json:"arp-reply"`
	MappedIp      string `json:"mappedip"`
	ExtIp         string `json:"extip"`
	ExtIntf       string `json:"extintf"`
	ConfigDefault string `json:"_if_no_default"`
	MappedAddr    string `json:"mapped-addr"`
}

// CreateUpdateFirewallObjectVip is for creating/updating the firewall object virtual ip
// Input:
//   @params: infor needed
//   @method: operation method, "add" or "update"
//   @adom: adom
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) CreateUpdateFirewallObjectVip(params *JSONFirewallObjectVip, method, adom string) (err error) {
	defer c.Trace("CreateUpdateFirewallObjectVip")()

	p := map[string]interface{}{
		"data": params,
		"url":  "/pm/config/adom/" + adom + "/obj/firewall/vip",
	}

	_, err = c.Do(method, p)

	if err != nil {
		err = fmt.Errorf("CreateUpdateFirewallObjectVip failed: %s", err)
		return
	}

	return
}

// ReadFirewallObjectVip is for reading the specific firewall object virtual ip
// Input:
//   @name: firewall object virtual ip name
//   @adom: adom
// Output:
//   @out: firewall object virtual ip infor
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) ReadFirewallObjectVip(adom, name string) (out *JSONFirewallObjectVip, err error) {
	defer c.Trace("ReadFirewallObjectVip")()

	p := map[string]interface{}{
		"url": "/pm/config/adom/" + adom + "/obj/firewall/vip/" + name,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadFirewallObjectVip failed: %s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONFirewallObjectVip{}
	if data["name"] != nil {
		out.Name = data["name"].(string)
	}
	if data["type"] != nil {
		out.Type = util.FirewallObjectVipType2Str(int(data["type"].(float64)))
	}
	if data["comment"] != nil {
		out.Comment = data["comment"].(string)
	}
	if data["arp-reply"] != nil {
		out.ArpReply = util.ControlSwitch2Str(int(data["arp-reply"].(float64)))
	}
	if data["mappedip"] != nil {
		m := util.InterfaceArray2StrArray(data["mappedip"].([]interface{}))
		out.MappedIp = m[0]
	}
	if data["extip"] != nil {
		m := util.InterfaceArray2StrArray(data["extip"].([]interface{}))
		out.ExtIp = m[0]
	}
	if data["extintf"] != nil {
		m := util.InterfaceArray2StrArray(data["extintf"].([]interface{}))
		out.ExtIntf = m[0]
	}
	if data["_if_no_default"] != nil {
		v := util.ControlSwitch2Str(int(data["_if_no_default"].(float64)))
		if v == "enable" {
			out.ConfigDefault = "disable"
		} else {
			out.ConfigDefault = "enable"
		}
	}
	if data["mapped-addr"] != nil {
		m := util.InterfaceArray2StrArray(data["mapped-addr"].([]interface{}))
		out.MappedAddr = m[0]
	}

	return

}

// DeleteFirewallObjectVip is for deleting the specific firewall object virtual ip
// Input:
//   @adom: adom
//   @name: firewall object virtual ip name
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) DeleteFirewallObjectVip(adom, name string) (err error) {
	defer c.Trace("DeleteFirewallObjectVip")()

	p := map[string]interface{}{
		"url": "/pm/config/adom/" + adom + "/obj/firewall/vip/" + name,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteFirewallObjectVip failed: %s", err)
		return
	}

	return
}
