---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_DoSpolicy"
description: |-
  Get information on an fortios firewall DoSpolicy.
---

# Data Source: fortios_firewall_DoSpolicy
Use this data source to get information on an fortios firewall DoSpolicy

## Argument Reference

* `policyid` - (Required) Specify the policyid of the desired firewall DoSpolicy.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `policyid` - Policy ID.
* `status` - Enable/disable this policy.
* `name` - Policy name.
* `comments` - Comment.
* `interface` - Incoming interface name from available interfaces.
* `custom_tags` - Custom tags. The structure of `custom_tags` block is documented below.
* `srcaddr` - Source address name from available addresses. The structure of `srcaddr` block is documented below.
* `dstaddr` - Destination address name from available addresses. The structure of `dstaddr` block is documented below.
* `service` - Service object from available options. The structure of `service` block is documented below.
* `anomaly` - Anomaly name. The structure of `anomaly` block is documented below.

The `custom_tags` block contains:

* `name` - Names of custom tags used with this policy.

The `srcaddr` block contains:

* `name` - Service name.

The `dstaddr` block contains:

* `name` - Address name.

The `service` block contains:

* `name` - Service name.

The `anomaly` block contains:

* `name` - Anomaly name.
* `status` - Enable/disable this anomaly.
* `log` - Enable/disable anomaly logging.
* `action` - Action taken when the threshold is reached.
* `synproxy_ttl` - Determine Time to live (TTL) value for packets replied by syn proxy module.
* `synproxy_tos` - Determine TCP differentiated services code point value (type of service).
* `synproxy_tcp_mss` - Determine TCP maximum segment size (MSS) value for packets replied by syn proxy module.
* `synproxy_tcp_sack` - enable/disable TCP selective acknowledage (SACK) for packets replied by syn proxy module.
* `synproxy_tcp_timestamp` - enable/disable TCP timestamp option for packets replied by syn proxy module.
* `synproxy_tcp_window` - Determine TCP Window size for packets replied by syn proxy module.
* `synproxy_tcp_windowscale` - Determine TCP window scale option value for packets replied by syn proxy module.
* `quarantine` - Quarantine method.
* `quarantine_expiry` - Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.
* `quarantine_log` - Enable/disable quarantine logging.
* `threshold` - Anomaly threshold. Number of detected instances per minute that triggers the anomaly action.
* `thresholddefault` - Number of detected instances per minute which triggers action (1 - 2147483647, default = 1000). Note that each anomaly has a different threshold value assigned to it.

