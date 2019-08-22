package fortimngclient

import (
	"fmt"
)

type JSONFirewallObjectVip struct {
	Name     string `json:"name"`
	Comment  string `json:"comment"`
	Type     string `json:"type"`
	ArpReply string `json:"arp-reply"`
	MappedIp string `json:"mappedip"`
	ExtIp    string `json:"extip"`
	ExtIntf  string `json:"extintf"`
}

// Create and Update function
func (c *FortiMngClient) CreateUpdateFirewallObjectVip(params *JSONFirewallObjectVip, method string) (err error) {
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

func (c *FortiMngClient) ReadFirewallObjectVip(name string) (out *JSONFirewallObjectVip, err error) {
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
		out.Type = c.FirewallObjectVipType2Str(int(data["type"].(float64)))
	}
	if data["comment"] != nil {
		out.Comment = data["comment"].(string)
	}
	if data["arp-reply"] != nil {
		out.ArpReply = c.ControlSwitch2Str(int(data["arp-reply"].(float64)))
	}
	if data["mappedip"] != nil {
		m := c.InterfaceArray2StrArray(data["mappedip"].([]interface{}))
		out.MappedIp = m[0]
	}
	if data["extip"] != nil {
		m := c.InterfaceArray2StrArray(data["extip"].([]interface{}))
		out.ExtIp = m[0]
	}
	if data["extintf"] != nil {
		m := c.InterfaceArray2StrArray(data["extintf"].([]interface{}))
		out.ExtIntf = m[0]
	}

	return

}

func (c *FortiMngClient) DeleteFirewallObjectVip(name string) (err error) {
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
