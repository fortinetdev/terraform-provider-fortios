---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ftmpush"
description: |-
  Get information on fortios system ftmpush.
---

# Data Source: fortios_system_ftmpush
Use this data source to get information on fortios system ftmpush

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `proxy` - Enable/disable communication to the proxy server in FortiGuard configuration.
* `server_port` - Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).
* `server_cert` - Name of the server certificate to be used for SSL (default = Fortinet_Factory).
* `server_ip` - IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).
* `server` - IPv4 address or domain name of FortiToken Mobile push services server.
* `status` - Enable/disable the use of FortiToken Mobile push services.

