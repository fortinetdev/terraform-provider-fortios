---
layout: "fortios"
page_title: "FortiOS: fortios_fortimanager_devicemanager_install_policypackage"
sidebar_current: "docs-fortios-fortimanager-resource-devicemanager-install-policypackage"
description: |-
  Provides a resource to install devicemanager policy package from FortiManager to the related FortiGate
---

# fortios_fortimanager_devicemanager_install_policypackage
This resource supports installing devicemanager policy package from FortiManager to the related FortiGate

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

resource "fortios_fortimanager_devicemanager_install_policypackage" "test1" {
	package_name = "test-pkg1"
	timeout = 5
}
```

## Argument Reference
The following arguments are supported:

* `package_name` - (Required) The installation package name.
* `timeout` - Timeout for installing the package to the target, default: 3 minutes.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `package_name` - The installation package name.
* `timeout` - Timeout for installing the package to the target, default: 3 minutes.
