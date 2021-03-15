---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemautoupdate_pushupdate"
description: |-
  Configure push updates.
---

# fortios_systemautoupdate_pushupdate
Configure push updates.

## Example Usage

```hcl
resource "fortios_systemautoupdate_pushupdate" "trname" {
  address  = "0.0.0.0"
  override = "disable"
  port     = 9443
  status   = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - (Required) Enable/disable push updates. Valid values: `enable`, `disable`.
* `override` - (Required) Enable/disable push update override server. Valid values: `enable`, `disable`.
* `address` - (Required) Push update override server.
* `port` - (Required) Push update override port. (Do not overlap with other service ports)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SystemAutoupdate PushUpdate can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_systemautoupdate_pushupdate.labelname SystemAutoupdatePushUpdate
$ unset "FORTIOS_IMPORT_TABLE"
```
