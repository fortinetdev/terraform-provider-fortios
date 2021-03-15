---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dedicatedmgmt"
description: |-
  Configure dedicated management.
---

# fortios_system_dedicatedmgmt
Configure dedicated management.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable dedicated management. Valid values: `enable`, `disable`.
* `interface` - Dedicated management interface.
* `default_gateway` - Default gateway for dedicated management interface.
* `dhcp_server` - Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
* `dhcp_netmask` - DHCP netmask.
* `dhcp_start_ip` - DHCP start IP for dedicated management.
* `dhcp_end_ip` - DHCP end IP for dedicated management.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System DedicatedMgmt can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_dedicatedmgmt.labelname SystemDedicatedMgmt
$ unset "FORTIOS_IMPORT_TABLE"
```
