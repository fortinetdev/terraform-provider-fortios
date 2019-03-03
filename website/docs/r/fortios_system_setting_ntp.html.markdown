---
layout: "fortios"
page_title: "FortiOS: fortios_system_setting_ntp"
sidebar_current: "docs-fortios-resource-system-setting-ntp"
description: |-
  Provides a resource to configure Network Time Protocol (NTP) servers of FortiOS.
---

# fortios_system_setting_ntp
Provides a resource to configure Network Time Protocol (NTP) servers of FortiOS.

## Example Usage
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_system_setting_ntp" "test2" {
	type = "custom"
	ntpserver = ["1.1.1.1", "3.3.3.3"]
	ntpsync = "disable"
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
