---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logtacacsaccounting2_setting"
description: |-
  Settings for TACACS+ accounting.
---

# fortios_logtacacsaccounting2_setting
Settings for TACACS+ accounting. Applies to FortiOS Version `>= 7.0.2`.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable TACACS+ accounting. Valid values: `enable`, `disable`.
* `server` - Address of TACACS+ server.
* `server_key` - Key to access the TACACS+ server.
* `source_ip` - Source IP address for communication to TACACS+ server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogTacacsAccounting2 Setting can be imported using any of these accepted formats:
```
$ terraform import fortios_logtacacsaccounting2_setting.labelname LogTacacsAccounting2Setting

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_logtacacsaccounting2_setting.labelname LogTacacsAccounting2Setting
$ unset "FORTIOS_IMPORT_TABLE"
```
