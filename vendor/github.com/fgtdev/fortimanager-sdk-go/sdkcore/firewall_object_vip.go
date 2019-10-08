package fmgclient

import (
	"fmt"

	"github.com/fgtdev/fortimanager-sdk-go/util"
)

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

// Create and Update function
func (c *FmgSDKClient) CreateUpdateFirewallObjectVip(params *JSONFirewallObjectVip, method string) (err error) {
	defer c.Trace("CreateUpdateFirewallObjectVip")()

	p := map[string]interface{}{
		"data": params,
		"url":  "/pm/config/adom/root/obj/firewall/vip",
	}

	_, err = c.Do(method, p)

	if err != nil {
		err = fmt.Errorf("CreateUpdateFirewallObjectVip failed: %s", err)
		return
	}

	return
}

func (c *FmgSDKClient) ReadFirewallObjectVip(name string) (out *JSONFirewallObjectVip, err error) {
	defer c.Trace("ReadFirewallObjectVip")()

	p := map[string]interface{}{
		"url": "/pm/config/adom/root/obj/firewall/vip/" + name,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadFirewallObjectVip failed :%s", err)
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

func (c *FmgSDKClient) DeleteFirewallObjectVip(name string) (err error) {
	defer c.Trace("DeleteFirewallObjectVip")()

	p := map[string]interface{}{
		"url": "/pm/config/adom/root/obj/firewall/vip/" + name,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteFirewallObjectVip failed :%s", err)
		return
	}

	return
}
