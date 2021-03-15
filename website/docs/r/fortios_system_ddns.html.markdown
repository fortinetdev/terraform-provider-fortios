---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ddns"
description: |-
  Configure DDNS.
---

# fortios_system_ddns
Configure DDNS.

## Example Usage

```hcl
resource "fortios_system_ddns" "trname" {
  bound_ip        = "0.0.0.0"
  clear_text      = "disable"
  ddns_auth       = "disable"
  ddns_domain     = "www.s.com"
  ddns_password   = "ewewcd"
  ddns_server     = "tzo.com"
  ddns_server_ip  = "0.0.0.0"
  ddns_ttl        = 300
  ddns_username   = "sie2ae"
  ddnsid          = 1
  ssl_certificate = "Fortinet_Factory"
  update_interval = 300
  use_public_ip   = "disable"
  monitor_interface {
    interface_name = "port2"
  }
}
```

## Argument Reference

The following arguments are supported:

* `ddnsid` - DDNS ID.
* `ddns_server` - (Required) Select a DDNS service provider. Valid values: `dyndns.org`, `dyns.net`, `tzo.com`, `vavic.com`, `dipdns.net`, `now.net.cn`, `dhs.org`, `easydns.com`, `genericDDNS`, `FortiGuardDDNS`, `noip.com`.
* `ddns_server_ip` - Generic DDNS server IP.
* `ddns_zone` - Zone of your domain name (for example, DDNS.com).
* `ddns_ttl` - Time-to-live for DDNS packets.
* `ddns_auth` - Enable/disable TSIG authentication for your DDNS server. Valid values: `disable`, `tsig`.
* `ddns_keyname` - DDNS update key name.
* `ddns_key` - DDNS update key (base 64 encoding).
* `ddns_domain` - Your fully qualified domain name (for example, yourname.DDNS.com).
* `ddns_username` - DDNS user name.
* `ddns_sn` - DDNS Serial Number.
* `ddns_password` - DDNS password.
* `use_public_ip` - Enable/disable use of public IP address. Valid values: `disable`, `enable`.
* `update_interval` - DDNS update interval (60 - 2592000 sec, default = 300).
* `clear_text` - Enable/disable use of clear text connections. Valid values: `disable`, `enable`.
* `ssl_certificate` - Name of local certificate for SSL connections.
* `bound_ip` - Bound IP address.
* `monitor_interface` - (Required) Monitored interface. The structure of `monitor_interface` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `monitor_interface` block supports:

* `interface_name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{ddnsid}}.

## Import

System Ddns can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_ddns.labelname {{ddnsid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
