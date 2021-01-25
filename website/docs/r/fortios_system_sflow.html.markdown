---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sflow"
description: |-
  Configure sFlow.
---

# fortios_system_sflow
Configure sFlow.

## Example Usage

```hcl
resource "fortios_system_sflow" "trname" {
  collector_ip   = "0.0.0.0"
  collector_port = 6343
  source_ip      = "0.0.0.0"
}
```

## Argument Reference


The following arguments are supported:

* `collector_ip` - (Required) IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
* `collector_port` - UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
* `source_ip` - Source IP address for sFlow agent.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Sflow can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_sflow.labelname SystemSflow
$ unset "FORTIOS_IMPORT_TABLE"
```
