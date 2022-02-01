---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_remotelog"
description: |-
  Configure logging by FortiSwitch device to a remote syslog server.
---

# fortios_switchcontroller_remotelog
Configure logging by FortiSwitch device to a remote syslog server. Applies to FortiOS Version `>= 6.2.4`.

## Argument Reference

The following arguments are supported:

* `name` - Remote log name.
* `status` - Enable/disable logging by FortiSwitch device to a remote syslog server. Valid values: `enable`, `disable`.
* `server` - IPv4 address of the remote syslog server.
* `port` - Remote syslog server listening port.
* `severity` - Severity of logs to be transferred to remote log server. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
* `csv` - Enable/disable comma-separated value (CSV) strings. Valid values: `enable`, `disable`.
* `facility` - Facility to log to remote syslog server. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController RemoteLog can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_remotelog.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
