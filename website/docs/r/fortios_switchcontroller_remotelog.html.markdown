---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_remotelog"
description: |-
  Configure logging by FortiSwitch device to a remote syslog server.
---

# fortios_switchcontroller_remotelog
Configure logging by FortiSwitch device to a remote syslog server.

## Argument Reference

The following arguments are supported:

* `name` - Remote log name.
* `status` - Enable/disable logging by FortiSwitch device to a remote syslog server.
* `server` - IPv4 address of the remote syslog server.
* `port` - Remote syslog server listening port.
* `severity` - Severity of logs to be transferred to remote log server.
* `csv` - Enable/disable comma-separated value (CSV) strings.
* `facility` - Facility to log to remote syslog server.


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
