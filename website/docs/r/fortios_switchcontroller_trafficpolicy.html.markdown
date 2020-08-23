---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_trafficpolicy"
description: |-
  Configure FortiSwitch traffic policy.
---

# fortios_switchcontroller_trafficpolicy
Configure FortiSwitch traffic policy.

## Example Usage

```hcl
resource "fortios_switchcontroller_trafficpolicy" "trname" {
  guaranteed_bandwidth = 0
  guaranteed_burst     = 0
  maximum_burst        = 0
  name                 = "1"
  policer_status       = "enable"
  type                 = "ingress"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Traffic policy name.
* `description` - Description of the traffic policy.
* `policer_status` - Enable/disable policer config on the traffic policy.
* `guaranteed_bandwidth` - Guaranteed bandwidth in kbps (max value = 524287000).
* `guaranteed_burst` - Guaranteed burst size in bytes (max value = 4294967295).
* `maximum_burst` - Maximum burst size in bytes (max value = 4294967295).
* `type` - Configure type of policy(ingress/egress).
* `cos` - COS queue(0 - 7), or unset to disable.
* `fosid` - FSW Policer id


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController TrafficPolicy can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_trafficpolicy.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
