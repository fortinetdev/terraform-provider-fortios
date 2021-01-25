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

* `event` - Enable/disable event logging.
* `system` - Enable/disable system event logging.
* `vpn` - Enable/disable VPN event logging.
* `user` - Enable/disable user authentication event logging.
* `router` - Enable/disable router event logging.
* `wireless_activity` - Enable/disable wireless event logging.
* `wan_opt` - Enable/disable WAN optimization event logging.
* `endpoint` - Enable/disable endpoint event logging.
* `ha` - Enable/disable ha event logging.
* `compliance_check` - Enable/disable PCI DSS compliance check logging.
* `security_rating` - Enable/disable Security Rating result logging.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log Eventfilter can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_log_eventfilter.labelname LogEventfilter
$ unset "FORTIOS_IMPORT_TABLE"
```
