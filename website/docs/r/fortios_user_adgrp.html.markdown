---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_adgrp"
description: |-
  Configure FSSO groups.
---

# fortios_user_adgrp
Configure FSSO groups.

## Example Usage

```hcl
resource "fortios_user_fsso" "trname1" {
  name       = "fssos1"
  port       = 32381
  port2      = 8000
  port3      = 8000
  port4      = 8000
  port5      = 8000
  server     = "1.1.1.1"
  source_ip  = "0.0.0.0"
  source_ip6 = "::"
}

resource "fortios_user_adgrp" "trname" {
  name        = "user_adgrp1"
  server_name = fortios_user_fsso.trname1.name
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `fabric_object` - Security Fabric global object setting. Valid values: `enable`, `disable`.
* `fabric_force_sync` - Enable/disable forced synchronization of configuration objects from the root FortiGate unit to the downstream devices.  Configuration conflict check is skipped. Valid values: `enable`, `disable`.
* `fabric_object_source` - Source of truth for fabric object. Valid values: `member`, `local`, `root`.
* `server_name` - FSSO agent name.
* `connector_source` - FSSO connector source.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.
* `fosid` - Group ID.

## Import

User Adgrp can be imported using any of these accepted formats:
```
$ terraform import fortios_user_adgrp.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_adgrp.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
