---
layout: "fortios"
page_title: "FortiOS: fortios_log_syslog_setting"
sidebar_current: "docs-fortios-resource-log-syslog-setting"
subcategory: "FortiGate"
description: |-
  Provides a resource to configure logging to remote Syslog logging servers.
---

# fortios_log_syslog_setting
Provides a resource to configure logging to remote Syslog logging servers.

## Example Usage
```hcl
resource "fortios_log_syslog_setting" "test2" {
  status    = "enable"
  server    = "2.2.2.2"
  mode      = "udp"
  port      = "514"
  facility  = "local7"
  source_ip = "10.2.2.199"
  format    = "csv"
}
```

## Argument Reference
The following arguments are supported:

* `status` - (Required) Enable/disable remote syslog logging.
* `server` - Address of remote syslog server.
* `mode` - Remote syslog logging over UDP/Reliable TCP.
* `port` - Server listen port.
* `facility` - Remote syslog facility.
* `source_ip` - Source IP address of syslog.
* `format` - Log format.

## Attributes Reference
The following attributes are exported:

* `status` - Enable/disable remote syslog logging.
* `server` - Address of remote syslog server.
* `mode` - Remote syslog logging over UDP/Reliable TCP.
* `port` - Server listen port.
* `facility` - Remote syslog facility.
* `source_ip` - Source IP address of syslog.
* `format` - Log format.

