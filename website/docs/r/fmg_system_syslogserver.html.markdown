---
layout: "fortios"
page_title: "FortiOS: fortios_fmg_system_syslogserver"
sidebar_current: "docs-fortios-fortimanager-resource-system-syslogserver"
subcategory: "FortiManager"
description: |-
  Provides a resource to configure system syslog server for FortiManager.
---

# fortios_fmg_system_syslogserver
This resource supports Create/Delete system syslog server for FortiManager.

## Example Usage
```hc
resource "fortios_fmg_system_syslogserver" "test1" {
  name = "test-syslog"
  ip   = "1.1.1.1"
  port = 99
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) Syslog server name.
* `ip` - Ipaddress of the syslog server.
* `port` - Port of the syslog server.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `name` - Syslog server name.
* `ip` - Ipaddress of the syslog server.
* `port` - Port of the syslog server.
