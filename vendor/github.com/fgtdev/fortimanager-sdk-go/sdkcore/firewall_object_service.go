package fortimngclient

import (
	"fmt"
)

type JSONFirewallObjectService struct {
	Name          string   `json:"name"`
	Comment       string   `json:"comment"`
	Category      string   `json:"category"`
	Protocol      string   `json:"protocol"`
	Proxy         string   `json:"proxy"`
	Fqdn          string   `json:"fqdn"`
	Iprange       string   `json:"iprange"`
	TcpPortRange  []string `json:"tcp-portrange"`
	UdpPortRange  []string `json:"udp-portrange"`
	SctpPortRange []string `json:"sctp-portrange"`
	IcmpCode      int      `json:"icmpcode"`
	IcmpType      int      `json:"icmptype"`
	ProtocolNum   int      `json:"protocol-number"`
}

// Create and Update function
func (c *FortiMngClient) CreateUpdateFirewallObjectService(params *JSONFirewallObjectService, method string) (err error) {
	defer c.Trace("CreateUpdateFirewallObjectService")()

	p := map[string]interface{}{
		"data": params,
		"url":  "/pm/config/adom/root/obj/firewall/service/custom",
	}

	_, err = c.Do(method, p)

	if err != nil {
		err = fmt.Errorf("CreateUpdateFirewallObjectService failed: %s", err)
		return
	}

	return
}

func (c *FortiMngClient) ReadFirewallObjectService(name string) (out *JSONFirewallObjectService, err error) {
	defer c.Trace("ReadFirewallObjectService")()

	p := map[string]interface{}{
		"url": "/pm/config/adom/root/obj/firewall/service/custom/" + name,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadFirewallObjectService failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONFirewallObjectService{}
	if data["name"] != nil {
		out.Name = data["name"].(string)
	}
	if data["protocol"] != nil {
		out.Protocol = c.FirewallObjectProtocol2Str(int(data["protocol"].(float64)))
	}
	if data["comment"] != nil {
		out.Comment = data["comment"].(string)
	}
	if data["category"] != nil {
		m := c.InterfaceArray2StrArray(data["category"].([]interface{}))
		// only 1 item is allowed here
		out.Category = m[0]
	}
	if data["proxy"] != nil {
		out.Proxy = c.ControlSwitch2Str(int(data["proxy"].(float64)))
	}
	if data["fqdn"] != nil {
		out.Fqdn = data["fqdn"].(string)
	}
	if data["iprange"] != nil {
		out.Iprange = data["iprange"].(string)
	}
	if data["tcp-portrange"] != nil {
		out.TcpPortRange = c.InterfaceArray2StrArray(data["tcp-portrange"].([]interface{}))
	}
	if data["udp-portrange"] != nil {
		out.UdpPortRange = c.InterfaceArray2StrArray(data["udp-portrange"].([]interface{}))
	}
	if data["sctp-portrange"] != nil {
		out.SctpPortRange = c.InterfaceArray2StrArray(data["sctp-portrange"].([]interface{}))
	}
	if data["icmpcode"] != nil {
		out.IcmpCode = int(data["icmpcode"].(float64))
	}
	if data["icmptype"] != nil {
		out.IcmpType = int(data["icmptype"].(float64))
	}
	if data["protocol-number"] != nil {
		out.ProtocolNum = int(data["protocol-number"].(float64))
	}

	return
}

func (c *FortiMngClient) DeleteFirewallObjectService(name string) (err error) {
	defer c.Trace("DeleteFirewallObjectService")()

	p := map[string]interface{}{
		"url": "/pm/config/adom/root/obj/firewall/service/custom/" + name,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteFirewallObjectService failed :%s", err)
		return
	}

	return
}
