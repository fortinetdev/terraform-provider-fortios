---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_isolatorserver"
description: |-
  Configure forward-server addresses.
---

# fortios_webproxy_isolatorserver
Configure forward-server addresses. Applies to FortiOS Version `>= 7.6.1`.

## Argument Reference

The following arguments are supported:

* `name` - Server name.
* `addr_type` - Address type of the forwarding proxy server: IP or FQDN. Valid values: `ip`, `ipv6`, `fqdn`.
* `ip` - Forward proxy server IP address.
* `ipv6` - Forward proxy server IPv6 address.
* `fqdn` - Forward server Fully Qualified Domain Name (FQDN).
* `port` - Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server. Set to -1 means unset this variable. CLI output may have different value on different FortiOS version.
* `comment` - Comment.
* `masquerade` - Enable/disable use of the IP address of the outgoing interface as the client IP address (default = enable) Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy IsolatorServer can be imported using any of these accepted formats:
```
$ terraform import fortios_webproxy_isolatorserver.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_webproxy_isolatorserver.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
