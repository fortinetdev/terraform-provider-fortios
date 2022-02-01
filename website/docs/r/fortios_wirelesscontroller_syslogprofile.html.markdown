---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_syslogprofile"
description: |-
  Configure Wireless Termination Points (WTP) system log server profile.
---

# fortios_wirelesscontroller_syslogprofile
Configure Wireless Termination Points (WTP) system log server profile. Applies to FortiOS Version `>= 7.0.2`.

## Argument Reference

The following arguments are supported:

* `name` - WTP system log server profile name.
* `comment` - Comment.
* `server_status` - Enable/disable FortiAP units to send log messages to a syslog server (default = enable). Valid values: `enable`, `disable`.
* `server_addr_type` - Syslog server address type (default = IP) Valid values: `fqdn`, `ip`.
* `server_fqdn` - FQDN of syslog server that FortiAP units send log messages to.
* `server_ip` - IP address of syslog server that FortiAP units send log messages to.
* `server_port` - Port number of syslog server that FortiAP units send log messages to (default = 514).
* `log_level` - Lowest level of log messages that FortiAP units send to this server (default = information) Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debugging`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController SyslogProfile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_syslogprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
