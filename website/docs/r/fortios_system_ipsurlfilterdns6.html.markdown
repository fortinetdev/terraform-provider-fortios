---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipsurlfilterdns6"
description: |-
  Configure IPS URL filter IPv6 DNS servers.
---

# fortios_system_ipsurlfilterdns6
Configure IPS URL filter IPv6 DNS servers.

## Argument Reference

The following arguments are supported:

* `address6` - IPv6 address of DNS server.
* `status` - Enable/disable this server for IPv6 DNS queries. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{address6}}.

## Import

System IpsUrlfilterDns6 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_ipsurlfilterdns6.labelname {{address6}}
$ unset "FORTIOS_IMPORT_TABLE"
```
