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
* `fqdn` - Forward server Fully Qualified Domain Name (FQDN).
* `port` - Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
* `healthcheck` - Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally.
* `monitor` - URL for forward server health check monitoring (default = http://www.google.com).
* `server_down_option` - Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination.
* `comment` - Comment.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy ForwardServer can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_webproxy_forwardserver.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
