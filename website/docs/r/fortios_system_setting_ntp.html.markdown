---
layout: "fortios"
page_title: "FortiOS: fortios_system_setting_ntp"
sidebar_current: "docs-fortios-resource-system-setting-ntp"
subcategory: "FortiGate"
description: |-
  Provides a resource to configure Network Time Protocol (NTP) servers of FortiOS.
---

# fortios_system_setting_ntp
Provides a resource to configure Network Time Protocol (NTP) servers of FortiOS.

## Example Usage
```hcl
resource "fortios_system_setting_ntp" "test2" {
  type      = "custom"
  ntpserver = ["1.1.1.1", "3.3.3.3"]
  ntpsync   = "disable"
}
```

## Argument Reference
The following arguments are supported:

* `type` - (Required) Use the FortiGuard NTP server or any other available NTP Server.
* `ntpserver` - Configure the FortiGate to connect to any available third-party NTP server.
* `ntpsync` - Enable/disable setting the FortiGate system time by synchronizing with an NTP Server.

## Attributes Reference
The following attributes are exported:

* `type` - Use the FortiGuard NTP server or any other available NTP Server.
* `ntpserver` - Configure the FortiGate to connect to any available third-party NTP server.
* `ntpsync` - Enable/disable setting the FortiGate system time by synchronizing with an NTP Server.
