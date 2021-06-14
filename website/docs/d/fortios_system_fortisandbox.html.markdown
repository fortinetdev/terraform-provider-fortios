---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fortisandbox"
description: |-
  Get information on fortios system fortisandbox.
---

# Data Source: fortios_system_fortisandbox
Use this data source to get information on fortios system fortisandbox

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `status` - Enable/disable FortiSandbox.
* `forticloud` - Enable/disable FortiSandbox Cloud.
* `server` - IPv4 or IPv6 address of the remote FortiSandbox.
* `source_ip` - Source IP address for communications to FortiSandbox.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `interface` - Specify outgoing interface to reach server.
* `enc_algorithm` - Configure the level of SSL protection for secure communication with FortiSandbox.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `email` - Notifier email address.

