---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_forwardserver"
description: |-
  Configure forward-server addresses.
---

# fortios_webproxy_forwardserver
Configure forward-server addresses.

## Example Usage

```hcl
resource "fortios_webproxy_forwardserver" "trname" {
  addr_type          = "fqdn"
  healthcheck        = "disable"
  ip                 = "0.0.0.0"
  monitor            = "http://www.google.com"
  name               = "1"
  port               = 3128
  server_down_option = "block"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Server name.
* `addr_type` - Address type of the forwarding proxy server: IP or FQDN.
* `ip` - Forward proxy server IP address.
* `ipv6` - Forward proxy server IPv6 address.
* `fqdn` - Forward server Fully Qualified Domain Name (FQDN).
* `port` - Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server. Set to -1 means unset this variable. CLI output may have different value on different FortiOS version.
* `healthcheck` - Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally. Valid values: `disable`, `enable`.
* `monitor` - URL for forward server health check monitoring (default = www.google.com).
* `server_down_option` - Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
* `username` - HTTP authentication user name.
* `password` - HTTP authentication password.
* `comment` - Comment.
* `masquerade` - Enable/disable use of the IP address of the outgoing interface as the client IP address (default = enable) Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy ForwardServer can be imported using any of these accepted formats:
```
$ terraform import fortios_webproxy_forwardserver.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_webproxy_forwardserver.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
