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
* `server_type` - Address type of the DDNS server. Valid values: `ipv4`, `ipv6`.
* `ddns_server_addr` - Generic DDNS server IP/FQDN list. The structure of `ddns_server_addr` block is documented below.
* `ddns_server_ip` - Generic DDNS server IP.
* `ddns_zone` - Zone of your domain name (for example, DDNS.com).
* `ddns_ttl` - Time-to-live for DDNS packets.
* `ddns_auth` - Enable/disable TSIG authentication for your DDNS server. Valid values: `disable`, `tsig`.
* `ddns_keyname` - DDNS update key name.
* `ddns_key` - DDNS update key (base 64 encoding).
* `ddns_domain` - Your fully qualified domain name. For example, yourname.ddns.com.
* `ddns_username` - DDNS user name.
* `ddns_sn` - DDNS Serial Number.
* `ddns_password` - DDNS password.
* `use_public_ip` - Enable/disable use of public IP address. Valid values: `disable`, `enable`.
* `addr_type` - Address type of interface address in DDNS update. Valid values: `ipv4`, `ipv6`.
* `update_interval` - DDNS update interval, 60 - 2592000 sec. On FortiOS versions 6.2.0-7.0.3: default = 300. On FortiOS versions >= 7.0.4: 0 means default.
* `clear_text` - Enable/disable use of clear text connections. Valid values: `disable`, `enable`.
* `ssl_certificate` - Name of local certificate for SSL connections.
* `bound_ip` - Bound IP address.
* `monitor_interface` - (Required) Monitored interface. The structure of `monitor_interface` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `ddns_server_addr` block supports:

* `addr` - IP address or FQDN of the server.

The `monitor_interface` block supports:

* `interface_name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{ddnsid}}.

## Import

System Ddns can be imported using any of these accepted formats:
```
$ terraform import fortios_system_ddns.labelname {{ddnsid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_ddns.labelname {{ddnsid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
