---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logtacacsaccounting3_filter"
description: |-
  Settings for TACACS+ accounting events filter.
---

# fortios_logtacacsaccounting3_filter
Settings for TACACS+ accounting events filter. Applies to FortiOS Version `>= 7.0.2`.

## Argument Reference

The following arguments are supported:

* `login_audit` - Enable/disable TACACS+ accounting for login events audit. Valid values: `enable`, `disable`.
* `config_change_audit` - Enable/disable TACACS+ accounting for configuration change events audit. Valid values: `enable`, `disable`.
* `cli_cmd_audit` - Enable/disable TACACS+ accounting for CLI commands audit. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogTacacsAccounting3 Filter can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_logtacacsaccounting3_filter.labelname LogTacacsAccounting3Filter
$ unset "FORTIOS_IMPORT_TABLE"
```
