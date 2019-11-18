---
layout: "fortios"
page_title: "FortiOS: fortios_fortimanager_system_license_vm"
sidebar_current: "docs-fortios-fortimanager-resource-system-license-vm"
description: |-
  Provides a resource to upload VM license to FortiGate through FortiManager.
---

# fortios_fortimanager_system_license_vm
This resource supports uploading VM license to FortiGate through FortiManager.

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

resource "fortios_fortimanager_system_license_vm" "test1" {
	target = "fortigate-test"
	file_content = "XXX" // your license file content
}
```

## Argument Reference
The following arguments are supported:

* `target` - (Required) Target name, which is managed by FortiManager.
* `file_content` - (Required) The license file, it needs to be base64 encoded.
* `adom` - ADOM that the target device belongs to. default is 'root'.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `target` - Target name, which is managed by FortiManager.
* `file_content` - The license file content, it needs to be base64 encoded.
* `adom` - ADOM name.
