---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_policy46"
description: |-
  Get information on an fortios firewall policy46.
---

# Data Source: fortios_firewall_policy46
Use this data source to get information on an fortios firewall policy46

## Argument Reference

* `policyid` - (Required) Specify the policyid of the desired firewall policy46.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `permit_any_host` - Enable/disable allowing any host.
* `policyid` - Policy ID.
* `name` - Policy name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `srcintf` - Source interface name.
* `dstintf` - Destination interface name.
* `srcaddr` - Source address objects. The structure of `srcaddr` block is documented below.
* `dstaddr` - Destination address objects. The structure of `dstaddr` block is documented below.
* `action` - Accept or deny traffic matching the policy.
* `status` - Enable/disable this policy.
* `schedule` - Schedule name.
* `service` - Service name. The structure of `service` block is documented below.
* `logtraffic` - Enable/disable traffic logging for this policy.
* `logtraffic_start` - Record logs when a session starts and ends.
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `per_ip_shaper` - Per IP traffic shaper.
* `fixedport` - Enable/disable fixed port for this policy.
* `tcp_mss_sender` - TCP Maximum Segment Size value of sender (0 - 65535, default = 0).
* `tcp_mss_receiver` - TCP Maximum Segment Size value of receiver (0 - 65535, default = 0)
* `comments` - Comment.
* `ippool` - Enable/disable use of IP Pools for source NAT.
* `poolname` - IP Pool names. The structure of `poolname` block is documented below.

The `srcaddr` block contains:

* `name` - Address name.

The `dstaddr` block contains:

* `name` - Address name.

The `service` block contains:

* `name` - Service name.

The `poolname` block contains:

* `name` - IP pool name.

