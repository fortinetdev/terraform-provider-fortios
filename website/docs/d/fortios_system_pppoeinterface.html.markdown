---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_pppoeinterface"
description: |-
  Get information on an fortios system pppoeinterface.
---

# Data Source: fortios_system_pppoeinterface
Use this data source to get information on an fortios system pppoeinterface

## Argument Reference

* `name` - (Required) Specify the name of the desired system pppoeinterface.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Name of the PPPoE interface.
* `dial_on_demand` - Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface.
* `ipv6` - Enable/disable IPv6 Control Protocol (IPv6CP).
* `device` - Name for the physical interface.
* `username` - User name.
* `password` - Enter the password.
* `pppoe_egress_cos` - CoS in VLAN tag for outgoing PPPoE/PPP packets.
* `auth_type` - PPP authentication type to use.
* `ipunnumbered` - PPPoE unnumbered IP.
* `pppoe_unnumbered_negotiate` - Enable/disable PPPoE unnumbered negotiation.
* `idle_timeout` - PPPoE auto disconnect after idle timeout (0-4294967295 sec).
* `multilink` - Enable/disable PPP multilink support.
* `mrru` - PPP MRRU (296 - 65535, default = 1500).
* `disc_retry_timeout` - PPPoE discovery init timeout value in (0-4294967295 sec).
* `padt_retry_timeout` - PPPoE terminate timeout value in (0-4294967295 sec).
* `service_name` - PPPoE service name.
* `ac_name` - PPPoE AC name.
* `lcp_echo_interval` - PPPoE LCP echo interval in (0-4294967295 sec, default = 5).
* `lcp_max_echo_fails` - Maximum missed LCP echo messages before disconnect (0-4294967295, default = 3).

