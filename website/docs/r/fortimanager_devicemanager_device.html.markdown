---
layout: "fortios"
page_title: "FortiOS: fortios_fortimanager_devicemanager_device"
sidebar_current: "docs-fortios-fortimanager-resource-devicemanager-device"
description: |-
  Provides a resource to add/delete online FortiGate to/from FortiManager
---

# fortios_fortimanager_devicemanager_device
This resource supports adding/deleting online FortiGate to/from FortiManager

## Example Usage
```hcl
provider "fortios" {
	hostname = "192.168.88.100"
	username = "APIUser"
	passwd = "admin"
	product = "fortimanager"
	insecure = false
	cabundlefile = "/path/yourCA.crt"
}

resource "fortios_fortimanager_devicemanager_device" "test1" {
	userid = "admin"
	password = ""
	ipaddr = "192.168.88.101"
	device_name  = "FGVM64-test"
}
```

## Argument Reference
The following arguments are supported:

* `userid` - (Required) User name.
* `password` - Password.
* `ipaddr` - (Required) Fortigate's ipaddress.
* `device_name` - (Required) Fortigate's device name.
* `adom` - Name or ID of the ADOM where the command is to be executed on.

## Attributes Reference
The following attributes are exported:
* `id` - The resource id.
* `userid` -  User name.
* `password` - Password.
* `ipaddr` -  Fortigate's ipaddress.
* `device_name` - Fortigate's device name.
* `adom` - Name or ID of the ADOM where the command is to be executed on.
