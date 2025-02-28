---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_ldbmonitor"
description: |-
  Configure server load balancing health monitors.
---

# fortios_firewall_ldbmonitor
Configure server load balancing health monitors.

## Example Usage

```hcl
resource "fortios_firewall_ldbmonitor" "trname" {
  http_max_redirects = 0
  interval           = 10
  name               = "ldbmonitor1"
  port               = 0
  retry              = 3
  timeout            = 2
  type               = "ping"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Monitor name.
* `type` - (Required) Select the Monitor type used by the health check monitor to check the health of the server. On FortiOS versions 6.2.0: PING | TCP | HTTP. On FortiOS versions 6.2.4-7.0.0: PING | TCP | HTTP | HTTPS. On FortiOS versions >= 7.0.1: PING | TCP | HTTP | HTTPS | DNS.
* `interval` - Time between health checks (default = 10). On FortiOS versions 6.2.0-7.0.17: 5 - 65635 sec. On FortiOS versions >= 7.2.0: 5 - 65535 sec.
* `timeout` - Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
* `retry` - Number health check attempts before the server is considered down (1 - 255, default = 3).
* `port` - Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (default = 0). On FortiOS versions 6.2.0-7.0.17: 0 - 65635. On FortiOS versions >= 7.2.0: 0 - 65535.
* `src_ip` - Source IP for ldb-monitor.
* `http_get` - URL used to send a GET request to check the health of an HTTP server.
* `http_match` - String to match the value expected in response to an HTTP-GET request.
* `http_max_redirects` - The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
* `dns_protocol` - Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
* `dns_request_domain` - Fully qualified domain name to resolve for the DNS probe.
* `dns_match_ip` - Response IP expected from DNS server.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall LdbMonitor can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_ldbmonitor.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_ldbmonitor.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
