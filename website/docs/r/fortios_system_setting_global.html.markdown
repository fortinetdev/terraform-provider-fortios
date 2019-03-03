---
layout: "fortios"
page_title: "FortiOS: fortios_system_setting_global"
sidebar_current: "docs-fortios-resource-system-setting-global"
description: |-
  Provides a resource to configure options related to the overall operation of FortiOS.
---

# fortios_system_setting_global
Provides a resource to configure options related to the overall operation of FortiOS.

## Example Usage
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_system_setting_global" "test1" {
	admintimeout = 65
	timezone = "04"
	hostname = "mytestFortiGate"
	admin_sport = 443
	admin_ssh_port = 22
}
```

## Argument Reference
The following arguments are supported:
* `hostname` - (Required) FortiGate unit's hostname.
* `admintimeout` - Number of minutes before an idle administrator session time out.
* `timezone` - Number corresponding to your time zone from 00 to 86.
* `admin_sport` - Administrative access port for HTTPS.
* `admin_ssh_port` - Administrative access port for SSH.

## Attributes Reference
The following attributes are exported:
* `admintimeout` - Number of minutes before an idle administrator session time out.
* `timezone` - Number corresponding to your time zone from 00 to 86.
* `hostname` - FortiGate unit's hostname.
* `admin_sport` - Administrative access port for HTTPS.
* `admin_ssh_port` - Administrative access port for SSH.
