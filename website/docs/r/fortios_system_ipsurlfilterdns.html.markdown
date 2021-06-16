---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipsurlfilterdns"
description: |-
  Configure IPS URL filter DNS servers.
---

# fortios_system_ipsurlfilterdns
Configure IPS URL filter DNS servers.

## Argument Reference

The following arguments are supported:

* `address` - DNS server IP address.
* `status` - Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
* `ipv6_capability` - Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{address}}.

## Import

System IpsUrlfilterDns can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_ipsurlfilterdns.labelname {{address}}
$ unset "FORTIOS_IMPORT_TABLE"
```
