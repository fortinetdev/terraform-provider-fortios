package fmgclient

import (
	"fmt"
	"strconv"

	"github.com/fgtdev/fortimanager-sdk-go/util"
)

// JSONFirewallSecurityPolicy contains the params for creating firewall security policy
type JSONFirewallSecurityPolicy struct {
	Name                 string   `json:"name"`
	Action               string   `json:"action"`
	PolicyId             string   `json:"policyid"`
	SrcAddr              []string `json:"srcaddr"`
	SrcIntf              []string `json:"srcintf"`
	DstAddr              []string `json:"dstaddr"`
	DstIntf              []string `json:"dstintf"`
	Service              []string `json:"service"`
	Schedule             []string `json:"schedule"`
	InternetService      string   `json:"internet-service"`
	InternetServiceID    []string `json:"internet-service-id"`
	InternetServiceSrc   string   `json:"internet-service-src"`
	InternetServiceSrcID []string `json:"internet-service-src-id"`
	Users                []string `json:"users"`
	Groups               []string `json:"groups"`
	Fsso                 string   `json:"fsso"`
	Rsso                 string   `json:"rsso"`
	Logtraffic           string   `json:"logtraffic"`
	LogtrafficStart      string   `json:"logtraffic-start"`
	CapturePacket        string   `json:"capture-packet"`
	Comments             string   `json:"comments"`
	NAT                  string   `json:"nat"`
	IpPool               string   `json:"ippool"`
	PoolName             []string `json:"poolname"`
	FixedPort            string   `json:"fixedport"`
	VpnTunnel            []string `json:"vpntunnel"`
	Inbound              string   `json:"inbound"`
	UTMStatus            string   `json:"utm-status"`
	ProfileType          string   `json:"profile-type"`
	// profile_type: single
	AvProfile              []string `json:"av-profile"`
	WebFilterProfile       []string `json:"webfilter-profile"`
	ApplicationList        []string `json:"application-list"`
	IpsSensor              []string `json:"ips-sensor"`
	WafProfile             []string `json:"waf-profile"`
	DnsFilterProfile       []string `json:"dnsfilter-profile"`
	ProfileProtocolOptions []string `json:"profile-protocol-options"`
	// profile_type: group
	ProfileGroup         []string `json:"profile-group"`
	TrafficShaper        []string `json:"traffic-shaper"`
	TrafficShaperReverse []string `json:"traffic-shaper-reverse"`
	PerIpShaper          []string `json:"per-ip-shaper"`
}

// FirewallSecurityPolicyInput contains the policy and the related package info
type FirewallSecurityPolicyInput struct {
	Policy      *JSONFirewallSecurityPolicy
	PackageName string
}

// CreateUpdateFirewallSecurityPolicy is for creating/updating the firewall security policy
// Input:
//   @params: infor needed
//   @method: operation method, "add" or "update"
//   @adom: adom
// Output:
//   @id: policy id
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) CreateUpdateFirewallSecurityPolicy(params *FirewallSecurityPolicyInput, method, adom string) (id string, err error) {
	defer c.Trace("CreateUpdateFirewallSecurityPolicy")()

	p := map[string]interface{}{
		"data": params.Policy,
		"url":  "/pm/config/adom/" + adom + "/pkg/" + params.PackageName + "/firewall/policy",
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

// ReadFirewallSecurityPolicy is for reading the specific firewall security policy
// Input:
//   @adom: adom
//   @id: policy id
//   @pkg_name: policy package name
// Output:
//   @out: policy infor
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) ReadFirewallSecurityPolicy(adom, id, pkg_name string) (out *JSONFirewallSecurityPolicy, err error) {
	defer c.Trace("ReadFirewallSecurityPolicy")()

	p := map[string]interface{}{
		"url": "/pm/config/adom/" + adom + "/pkg/" + pkg_name + "/firewall/policy/" + id,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadFirewallSecurityPolicy failed: %s", err)
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
	if data["comments"] != nil {
		out.Comments = data["comments"].(string)
	}
	if data["action"] != nil {
		out.Action = util.PolicyAction2Str(int(data["action"].(float64)))
	}
	if data["srcaddr"] != nil {
		m := util.InterfaceArray2StrArray(data["srcaddr"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.SrcAddr = append(out.SrcAddr, m[i])
		}
	}
	if data["srcintf"] != nil {
		m := util.InterfaceArray2StrArray(data["srcintf"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.SrcIntf = append(out.SrcIntf, m[i])
		}
	}
	if data["dstaddr"] != nil {
		m := util.InterfaceArray2StrArray(data["dstaddr"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.DstAddr = append(out.DstAddr, m[i])
		}
	}
	if data["dstintf"] != nil {
		m := util.InterfaceArray2StrArray(data["dstintf"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.DstIntf = append(out.DstIntf, m[i])
		}
	}
	if data["service"] != nil {
		m := util.InterfaceArray2StrArray(data["service"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.Service = append(out.Service, m[i])
		}
	}
	if data["schedule"] != nil {
		m := util.InterfaceArray2StrArray(data["schedule"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.Schedule = append(out.Schedule, m[i])
		}
	}
	if data["internet-service"] != nil {
		out.InternetService = util.ControlSwitch2Str(int(data["internet-service"].(float64)))
	}
	if data["internet-service-id"] != nil {
		m := util.InterfaceArray2StrArray(data["internet-service-id"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.InternetServiceID = append(out.InternetServiceID, m[i])
		}
	}
	if data["internet-service-src"] != nil {
		out.InternetServiceSrc = util.ControlSwitch2Str(int(data["internet-service-src"].(float64)))
	}
	if data["internet-service-src-id"] != nil {
		m := util.InterfaceArray2StrArray(data["internet-service-src-id"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.InternetServiceSrcID = append(out.InternetServiceSrcID, m[i])
		}
	}
	if data["users"] != nil {
		m := util.InterfaceArray2StrArray(data["users"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.Users = append(out.Users, m[i])
		}
	}
	if data["groups"] != nil {
		m := util.InterfaceArray2StrArray(data["groups"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.Groups = append(out.Groups, m[i])
		}
	}
	if data["rsso"] != nil {
		out.Rsso = util.ControlSwitch2Str(int(data["rsso"].(float64)))
	}
	if data["fsso"] != nil {
		out.Fsso = util.ControlSwitch2Str(int(data["fsso"].(float64)))
	}
	if data["logtraffic"] != nil {
		out.Logtraffic = util.LogTraffic2Str(int(data["logtraffic"].(float64)))
	}
	if data["logtraffic-start"] != nil {
		out.LogtrafficStart = util.ControlSwitch2Str(int(data["logtraffic-start"].(float64)))
	}
	if data["capture-packet"] != nil {
		out.CapturePacket = util.ControlSwitch2Str(int(data["capture-packet"].(float64)))
	}
	if data["nat"] != nil {
		out.NAT = util.ControlSwitch2Str(int(data["nat"].(float64)))
	}
	if data["ippool"] != nil {
		out.IpPool = util.ControlSwitch2Str(int(data["ippool"].(float64)))
	}
	if data["poolname"] != nil {
		m := util.InterfaceArray2StrArray(data["poolname"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.PoolName = append(out.PoolName, m[i])
		}
	}
	if data["fixedport"] != nil {
		out.FixedPort = util.ControlSwitch2Str(int(data["fixedport"].(float64)))
	}
	if data["vpntunnel"] != nil {
		m := util.InterfaceArray2StrArray(data["vpntunnel"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.VpnTunnel = append(out.VpnTunnel, m[i])
		}
	}
	if data["inbound"] != nil {
		out.Inbound = util.ControlSwitch2Str(int(data["inbound"].(float64)))
	}
	if data["utm-status"] != nil {
		out.UTMStatus = util.ControlSwitch2Str(int(data["utm-status"].(float64)))
	}
	if data["profile-type"] != nil {
		out.ProfileType = util.ProfileType2Str(int(data["profile-type"].(float64)))
	}
	if data["av-profile"] != nil {
		m := util.InterfaceArray2StrArray(data["av-profile"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.AvProfile = append(out.AvProfile, m[i])
		}
	}
	if data["webfilter-profile"] != nil {
		m := util.InterfaceArray2StrArray(data["webfilter-profile"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.WebFilterProfile = append(out.WebFilterProfile, m[i])
		}
	}
	if data["application-list"] != nil {
		m := util.InterfaceArray2StrArray(data["application-list"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.ApplicationList = append(out.ApplicationList, m[i])
		}
	}
	if data["ips-sensor"] != nil {
		m := util.InterfaceArray2StrArray(data["ips-sensor"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.IpsSensor = append(out.IpsSensor, m[i])
		}
	}
	if data["waf-profile"] != nil {
		m := util.InterfaceArray2StrArray(data["waf-profile"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.WafProfile = append(out.WafProfile, m[i])
		}
	}
	if data["dnsfilter-profile"] != nil {
		m := util.InterfaceArray2StrArray(data["dnsfilter-profile"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.DnsFilterProfile = append(out.DnsFilterProfile, m[i])
		}
	}
	if data["profile-protocol-options"] != nil {
		m := util.InterfaceArray2StrArray(data["profile-protocol-options"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.ProfileProtocolOptions = append(out.ProfileProtocolOptions, m[i])
		}
	}
	if data["profile-group"] != nil {
		m := util.InterfaceArray2StrArray(data["profile-group"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.ProfileGroup = append(out.ProfileGroup, m[i])
		}
	}
	if data["traffic-shaper"] != nil {
		m := util.InterfaceArray2StrArray(data["traffic-shaper"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.TrafficShaper = append(out.TrafficShaper, m[i])
		}
	}
	if data["traffic-shaper-reverse"] != nil {
		m := util.InterfaceArray2StrArray(data["traffic-shaper-reverse"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.TrafficShaperReverse = append(out.TrafficShaperReverse, m[i])
		}
	}
	if data["per-ip-shaper"] != nil {
		m := util.InterfaceArray2StrArray(data["per-ip-shaper"].([]interface{}))
		for i := 0; i < len(m); i++ {
			out.PerIpShaper = append(out.PerIpShaper, m[i])
		}

	}

	return
}

// DeleteFirewallSecurityPolicy is for deleting the specific firewall security policy
// Input:
//   @adom: adom
//   @id: policy id
//   @pkg_name: policy package name
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) DeleteFirewallSecurityPolicy(adom, id, pkg_name string) (err error) {
	defer c.Trace("DeleteFirewallSecurityPolicy")()

	p := map[string]interface{}{
		"url": "/pm/config/adom/" + adom + "/pkg/" + pkg_name + "/firewall/policy/" + id,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteFirewallSecurityPolicy failed: %s", err)
		return
	}

	return
}
