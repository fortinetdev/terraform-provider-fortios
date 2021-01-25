---
subcategory: "FortiGate IPS"
layout: "fortios"
page_title: "FortiOS: fortios_ips_rulesettings"
description: |-
  Configure IPS rule setting.
---

# fortios_ips_rulesettings
Configure IPS rule setting.

## Argument Reference


The following arguments are supported:

* `fosid` - Rule ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Ips RuleSettings can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_ips_rulesettings.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
