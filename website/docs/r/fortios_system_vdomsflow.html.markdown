---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomsflow"
description: |-
  Configure sFlow per VDOM to add or change the IP address and UDP port that FortiGate sFlow agents in this VDOM use to send sFlow datagrams to an sFlow collector.
---

# fortios_system_vdomsflow
Configure sFlow per VDOM to add or change the IP address and UDP port that FortiGate sFlow agents in this VDOM use to send sFlow datagrams to an sFlow collector.

## Example Usage

```hcl
resource "fortios_system_vdomsflow" "trname" {
  collector_ip   = "0.0.0.0"
  collector_port = 6343
  source_ip      = "0.0.0.0"
  vdom_sflow     = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `vdom_sflow` - Enable/disable the sFlow configuration for the current VDOM.
* `collector_ip` - IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
* `collector_port` - UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
* `source_ip` - Source IP address for sFlow agent.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System VdomSflow can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_vdomsflow.labelname SystemVdomSflow
$ unset "FORTIOS_IMPORT_TABLE"
```
