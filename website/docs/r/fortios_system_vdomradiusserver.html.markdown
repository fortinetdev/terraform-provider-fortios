---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomradiusserver"
description: |-
  Configure a RADIUS server to use as a RADIUS Single Sign On (RSSO) server for this VDOM.
---

# fortios_system_vdomradiusserver
Configure a RADIUS server to use as a RADIUS Single Sign On (RSSO) server for this VDOM.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the VDOM that you are adding the RADIUS server to.
* `status` - Enable/disable the RSSO RADIUS server for this VDOM.
* `radius_server_vdom` - (Required) Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System VdomRadiusServer can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_vdomradiusserver.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
