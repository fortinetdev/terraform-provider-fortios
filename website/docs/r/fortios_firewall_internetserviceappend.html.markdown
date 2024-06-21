---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetserviceappend"
description: |-
  Configure additional port mappings for Internet Services.
---

# fortios_firewall_internetserviceappend
Configure additional port mappings for Internet Services. Applies to FortiOS Version `6.2.4,6.2.6,6.4.1,6.4.2,6.4.10,6.4.11,6.4.12,6.4.13,6.4.14,6.4.15,7.0.0,7.0.1,7.0.2,7.0.3,7.0.4,7.0.5,7.0.6,7.0.7,7.0.8,7.0.9,7.0.10,7.0.11,7.0.12,7.0.13,7.0.14,7.0.15,7.2.0,7.2.1,7.2.2,7.2.3,7.2.4,7.2.6,7.2.7,7.2.8,7.4.0,7.4.1,7.4.2,7.4.3,7.4.4`.

## Argument Reference

The following arguments are supported:

* `addr_mode` - Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`, `both`.
* `match_port` - Matching TCP/UDP/SCTP destination port (1 to 65535).
* `append_port` - Appending TCP/UDP/SCTP destination port (1 to 65535).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall InternetServiceAppend can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_internetserviceappend.labelname FirewallInternetServiceAppend

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_internetserviceappend.labelname FirewallInternetServiceAppend
$ unset "FORTIOS_IMPORT_TABLE"
```
