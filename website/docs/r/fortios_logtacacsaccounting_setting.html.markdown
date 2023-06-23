---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logtacacsaccounting_setting"
description: |-
  Settings for TACACS+ accounting.
---

# fortios_logtacacsaccounting_setting
Settings for TACACS+ accounting. Applies to FortiOS Version `>= 7.0.2`.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable TACACS+ accounting. Valid values: `enable`, `disable`.
* `server` - Address of TACACS+ server.
* `server_key` - Key to access the TACACS+ server.
* `source_ip` - Source IP address for communication to TACACS+ server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogTacacsAccounting Setting can be imported using any of these accepted formats:
```
$ terraform import fortios_logtacacsaccounting_setting.labelname LogTacacsAccountingSetting

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_logtacacsaccounting_setting.labelname LogTacacsAccountingSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
