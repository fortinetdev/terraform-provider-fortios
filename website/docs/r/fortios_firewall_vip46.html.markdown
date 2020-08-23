---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_vip46"
description: |-
  Configure IPv4 to IPv6 virtual IPs.
---

# fortios_firewall_vip46
Configure IPv4 to IPv6 virtual IPs.

## Example Usage

```hcl
resource "fortios_firewall_vip46" "trname" {
  arp_reply   = "enable"
  color       = 0
  extip       = "10.202.1.200"
  extport     = "0-65535"
  fosid       = 0
  ldb_method  = "static"
  mappedip    = "2001:1:1:2::200"
  mappedport  = "0-65535"
  name        = "vip46s1"
  portforward = "disable"
  protocol    = "tcp"
  type        = "static-nat"
}
```

## Argument Reference

The following arguments are supported:

* `name` - VIP46 name.
* `fosid` - Custom defined id.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `comment` - Comment.
* `type` - VIP type: static NAT or server load balance.
* `src_filter` - Source IP filter (x.x.x.x/x).
* `extip` - (Required) Start-external-IP [-end-external-IP].
* `mappedip` - (Required) Start-mapped-IP [-end mapped-IP].
* `arp_reply` - Enable ARP reply.
* `portforward` - Enable port forwarding.
* `protocol` - Mapped port protocol.
* `extport` - (Required) External service port.
* `mappedport` - Mapped service port.
* `color` - Color of icon on the GUI.
* `ldb_method` - Load balance method.
* `server_type` - Server type.
* `realservers` - Real servers.
* `monitor` - Health monitors.

The `src_filter` block supports:

* `range` - Src-filter range.

The `realservers` block supports:

* `id` - Real server ID.
* `ip` - Mapped server IPv6.
* `port` - Mapped server port.
* `status` - Server administrative status.
* `weight` - weight
* `holddown_interval` - Hold down interval.
* `healthcheck` - Per server health check.
* `max_connections` - Maximum number of connections allowed to server.
* `monitor` - Health monitors.
* `client_ip` - Restrict server to a client IP in this range.

The `monitor` block supports:

* `name` - Health monitor name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Vip46 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_vip46.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
