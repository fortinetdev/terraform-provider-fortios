---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dnsdatabase"
description: |-
  Configure DNS databases.
---

# fortios_system_dnsdatabase
Configure DNS databases.

## Example Usage

```hcl
resource "fortios_system_dnsdatabase" "trname" {
  authoritative = "enable"
  contact       = "hostmaster"
  domain        = "s.com"
  ip_master     = "0.0.0.0"
  name          = "1"
  primary_name  = "dns"
  source_ip     = "0.0.0.0"
  status        = "enable"
  ttl           = 86400
  type          = "master"
  view          = "shadow"
  forwarder     = "\"9.9.9.9\" \"3.3.3.3\" "
  dns_entry {
    type     = "MX"
    ttl      = 3
    hostname = "sghsgh.com"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Zone name.
* `status` - Enable/disable this DNS zone. Valid values: `enable`, `disable`.
* `domain` - (Required) Domain name.
* `allow_transfer` - DNS zone transfer IP address list.
* `type` - (Required) Zone type (primary to manage entries directly, secondary to import entries from other zones).
* `view` - (Required) Zone view (public to serve public clients, shadow to serve internal clients).
* `ip_primary` - IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.
* `ip_master` - IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.
* `primary_name` - Domain name of the default DNS server for this zone.
* `contact` - Email address of the administrator for this zone. You can specify only the username, such as admin or the full email address, such as admin@test.com When using only a username, the domain of the email will be this zone.
* `ttl` - (Required) Default time-to-live value for the entries of this DNS zone (0 - 2147483647 sec, default = 86400).
* `authoritative` - (Required) Enable/disable authoritative zone. Valid values: `enable`, `disable`.
* `forwarder` - DNS zone forwarder IP address list.
* `forwarder6` - Forwarder IPv6 address.
* `source_ip` - Source IP for forwarding to DNS server.
* `source_ip6` - IPv6 source IP address for forwarding to DNS server.
* `source_ip_interface` - IP address of the specified interface as the source IP address.
* `rr_max` - Maximum number of resource records (10 - 65536, 0 means infinite).
* `dns_entry` - DNS entry. The structure of `dns_entry` block is documented below.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `dns_entry` block supports:

* `id` - DNS entry ID.
* `status` - Enable/disable resource record status. Valid values: `enable`, `disable`.
* `type` - Resource record type. Valid values: `A`, `NS`, `CNAME`, `MX`, `AAAA`, `PTR`, `PTR_V6`.
* `ttl` - Time-to-live for this entry (0 to 2147483647 sec, default = 0).
* `preference` - DNS entry preference (0 - 65535, highest preference = 0, default = 10).
* `ip` - IPv4 address of the host.
* `ipv6` - IPv6 address of the host.
* `hostname` - Name of the host.
* `canonical_name` - Canonical name of the host.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System DnsDatabase can be imported using any of these accepted formats:
```
$ terraform import fortios_system_dnsdatabase.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_dnsdatabase.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
