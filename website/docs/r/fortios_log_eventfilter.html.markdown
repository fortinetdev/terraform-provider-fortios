---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_log_eventfilter"
description: |-
  Configure log event filters.
---

# fortios_log_eventfilter
Configure log event filters.

## Example Usage

```hcl
resource "fortios_log_eventfilter" "trname" {
  compliance_check  = "enable"
  endpoint          = "enable"
  event             = "enable"
  ha                = "enable"
  router            = "enable"
  security_rating   = "enable"
  system            = "enable"
  user              = "enable"
  vpn               = "enable"
  wan_opt           = "enable"
  wireless_activity = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `event` - Enable/disable event logging. Valid values: `enable`, `disable`.
* `system` - Enable/disable system event logging. Valid values: `enable`, `disable`.
* `vpn` - Enable/disable VPN event logging. Valid values: `enable`, `disable`.
* `user` - Enable/disable user authentication event logging. Valid values: `enable`, `disable`.
* `router` - Enable/disable router event logging. Valid values: `enable`, `disable`.
* `wireless_activity` - Enable/disable wireless event logging. Valid values: `enable`, `disable`.
* `wan_opt` - Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
* `endpoint` - Enable/disable endpoint event logging. Valid values: `enable`, `disable`.
* `ha` - Enable/disable ha event logging. Valid values: `enable`, `disable`.
* `compliance_check` - Enable/disable PCI DSS compliance check logging. Valid values: `enable`, `disable`.
* `security_rating` - Enable/disable Security Rating result logging. Valid values: `enable`, `disable`.
* `fortiextender` - Enable/disable FortiExtender logging. Valid values: `enable`, `disable`.
* `connector` - Enable/disable SDN connector logging. Valid values: `enable`, `disable`.
* `sdwan` - Enable/disable SD-WAN logging. Valid values: `enable`, `disable`.
* `cifs` - Enable/disable CIFS logging. Valid values: `enable`, `disable`.
* `switch_controller` - Enable/disable Switch-Controller logging. Valid values: `enable`, `disable`.
* `rest_api` - Enable/disable REST API logging. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log Eventfilter can be imported using any of these accepted formats:
```
$ terraform import fortios_log_eventfilter.labelname LogEventfilter

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_log_eventfilter.labelname LogEventfilter
$ unset "FORTIOS_IMPORT_TABLE"
```
