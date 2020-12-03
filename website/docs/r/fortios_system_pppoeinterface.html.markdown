---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_pppoeinterface"
description: |-
  Configure the PPPoE interfaces.
---

# fortios_system_pppoeinterface
Configure the PPPoE interfaces.

## Example Usage

```hcl
resource "fortios_system_pppoeinterface" "trname" {
  auth_type                  = "auto"
  device                     = "port3"
  dial_on_demand             = "disable"
  disc_retry_timeout         = 1
  idle_timeout               = 0
  ipunnumbered               = "0.0.0.0"
  ipv6                       = "disable"
  lcp_echo_interval          = 5
  lcp_max_echo_fails         = 3
  name                       = "14"
  padt_retry_timeout         = 1
  pppoe_unnumbered_negotiate = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the PPPoE interface.
* `dial_on_demand` - Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface.
* `ipv6` - Enable/disable IPv6 Control Protocol (IPv6CP).
* `device` - (Required) Name for the physical interface.
* `username` - User name.
* `password` - Enter the password.
* `auth_type` - PPP authentication type to use.
* `ipunnumbered` - PPPoE unnumbered IP.
* `pppoe_unnumbered_negotiate` - Enable/disable PPPoE unnumbered negotiation.
* `idle_timeout` - PPPoE auto disconnect after idle timeout (0-4294967295 sec).
* `disc_retry_timeout` - PPPoE discovery init timeout value in (0-4294967295 sec).
* `padt_retry_timeout` - PPPoE terminate timeout value in (0-4294967295 sec).
* `service_name` - PPPoE service name.
* `ac_name` - PPPoE AC name.
* `lcp_echo_interval` - PPPoE LCP echo interval in (0-4294967295 sec, default = 5).
* `lcp_max_echo_fails` - Maximum missed LCP echo messages before disconnect (0-4294967295, default = 3).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System PppoeInterface can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_pppoeinterface.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
