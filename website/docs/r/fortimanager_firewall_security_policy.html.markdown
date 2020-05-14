---
layout: "fortios"
page_title: "FortiOS: fortios_fmg_firewall_security_policy"
sidebar_current: "docs-fortios-fortimanager-resource-firewall-security-policy"
description: |-
  Provides a resource to configure firewall security policy on FortiManager which could be installed to the FortiGate later
---

# fortios_fmg_firewall_security_policy
This resource supports Create/Read/Update/Delete firewall security policy on FortiManager which could be installed to the FortiGate later


## Example Usage
```hcl
resource "fortios_fmg_firewall_security_policy" "test1" {
	name = "policy-test"
	srcaddr = ["all"]
	srcintf = ["any"]
	dstaddr = ["all"]
	dstintf = ["any"]
	service = ["ALL"]
	action = "accept"
	schedule = ["always"]
	internet_service = "enable"
	internet_service_id = ["Alibaba.Web", "AOL.Web"]
	internet_service_src = "disable"
	users = ["guest"]
	groups = ["Guest-group"]
	rsso = "disable"
	fsso = "enable"
	logtraffic = "all"
	logtraffic_start = "enable"
	capture_packet = "enable"
	nat = "enable"
	ippool = "enable"
	fixedport = "enable"
	poolname = ["test1"]
	utm_status = "enable"
	profile_type = "single"
	av_profile = ["g-default"]
	dnsfilter_profile = ["default"]
	traffic_shaper = ["high-priority"]
	comments = "policy test"
	package_name = "dvm-test"
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) Policy name.
* `srcaddr` - (Required) Source address and adress group names.
* `srcintf` - (Required) Incoming interface.
* `dstaddr` - (Required) Destination address and adress group names.
* `dstintf` - (Required) Outgoing interface.
* `service` - (Required) Service and service group names.
* `action` - (Required) Policy action, default is deny. Enum: [allow, deny, ipsec].
* `schedule` - (Required) Schedule name.
* `internet_service` - Enable/disable use of Destination Internet Services for this policy.
* `internet_service_id` - Destination Internet Service ID.
* `internet_service_src` - Enable/disable use of Source Internet Services for this policy.
* `internet_service_src_id` - Source Internet Service ID.
* `users` - Names of individual users that can authenticate with this policy.
* `groups` - Names of user groups that can authenticate with this policy.
* `fsso` - Enable/disable Fortinet Single Sign-On.
* `rsso` - Enable/disable RADIUS Single Sign-On.
* `logtraffic` - Enable or disable logging. Enum: [disable, all, utm]
* `logtraffic_start` - Record logs when a session starts and ends. Enum: [disable, enable]
* `capture_packet` - Enable/disable capture packets.
* `comments` - Comments.
* `nat` - Enable/disable source NAT.
* `ippool` - Enable/disable to use IP Pools for source NAT.
* `poolname` - IP Pool names.
* `fixedport` - Enable/disable to prevent source NAT from changing a session's source port.
* `vpntunnel` - Policy-based IPsec VPN: name of the IPsec VPN Phase 1.
* `inbound` - Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN. Enum: [disable, enable]
* `utm_status` - Enable/disable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only. Enum: [single, group]
* `av_profile` - Name of an existing Antivirus profile.
* `webfilter_profile` - Name of an existing Web filter profile.
* `application_list` - Name of an existing Application list.
* `ips_sensor` - Name of an existing IPS sensor.
* `waf_profile` - Name of an existing Web application firewall profile.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `profile_group` - Name of profile group.
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `per_ip_shaper` - Per-IP traffic shaper.
* `package_name` - The package name which the policy will be added to.
* `adom` - ADOM name. default is 'root'.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `name` - Policy name.
* `srcaddr` - Source address and adress group names.
* `srcintf` - Incoming interface.
* `dstaddr` - Destination address and adress group names.
* `dstintf` - Outgoing interface.
* `service` - Service and service group names.
* `action` - Policy action.
* `schedule` - Schedule name.
* `internet_service` - Enable/disable use of Destination Internet Services for this policy.
* `internet_service_id` - Destination Internet Service ID.
* `internet_service_src` - Enable/disable use of Source Internet Services for this policy.
* `internet_service_src_id` - Source Internet Service ID.
* `users` - Names of individual users that can authenticate with this policy.
* `groups` - Names of user groups that can authenticate with this policy.
* `fsso` - Enable/disable Fortinet Single Sign-On.
* `rsso` - Enable/disable RADIUS Single Sign-On.
* `logtraffic` - Enable or disable logging.
* `logtraffic_start` - Record logs when a session starts and ends.
* `capture_packet` - Enable/disable capture packets.
* `comments` - Comments.
* `nat` - Enable/disable source NAT.
* `ippool` - Enable/disable to use IP Pools for source NAT.
* `poolname` - IP Pool names.
* `fixedport` - Enable/disable to prevent source NAT from changing a session's source port.
* `vpntunnel` - Policy-based IPsec VPN: name of the IPsec VPN Phase 1.
* `inbound` - Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN.
* `utm_status` - Enable/disable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only.
* `av_profile` - Name of an existing Antivirus profile.
* `webfilter_profile` - Name of an existing Web filter profile.
* `application_list` - Name of an existing Application list.
* `ips_sensor` - Name of an existing IPS sensor.
* `waf_profile` - Name of an existing Web application firewall profile.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `profile_group` - Name of profile group.
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `per_ip_shaper` - Per-IP traffic shaper.
* `package_name` - The package name which the policy will be added to.
* `adom` - ADOM name.
