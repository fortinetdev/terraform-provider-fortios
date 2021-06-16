---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetserviceappend"
description: |-
  Configure additional port mappings for Internet Services.
---

# fortios_firewall_internetserviceappend
Configure additional port mappings for Internet Services.

## Argument Reference

The following arguments are supported:

* `match_port` - Matching TCP/UDP/SCTP destination port (1 to 65535).
* `append_port` - Appending TCP/UDP/SCTP destination port (1 to 65535).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall InternetServiceAppend can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_internetserviceappend.labelname FirewallInternetServiceAppend
$ unset "FORTIOS_IMPORT_TABLE"
```
