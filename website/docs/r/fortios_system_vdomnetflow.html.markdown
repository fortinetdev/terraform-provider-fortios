---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomnetflow"
description: |-
  Configure NetFlow per VDOM.
---

# fortios_system_vdomnetflow
Configure NetFlow per VDOM.

## Example Usage

```hcl
resource "fortios_system_vdomnetflow" "trname" {
  collector_ip   = "0.0.0.0"
  collector_port = 2055
  source_ip      = "0.0.0.0"
  vdom_netflow   = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `vdom_netflow` - Enable/disable NetFlow per VDOM.
* `collector_ip` - NetFlow collector IP address.
* `collector_port` - NetFlow collector port number.
* `source_ip` - Source IP address for communication with the NetFlow agent.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System VdomNetflow can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_vdomnetflow.labelname SystemVdomNetflow
$ unset "FORTIOS_IMPORT_TABLE"
```
