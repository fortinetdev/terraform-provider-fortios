---
layout: "fortios"
page_title: "FortiOS: fortios_fmg_devicemanager_install_policypackage"
sidebar_current: "docs-fortios-fortimanager-resource-devicemanager-install-policypackage"
subcategory: "FortiManager"
description: |-
  Provides a resource to install devicemanager policy package from FortiManager to the related FortiGate
---

# fortios_fmg_devicemanager_install_policypackage
This resource supports installing devicemanager policy package from FortiManager to the related FortiGate

## Example Usage
```hcl
resource "fortios_fmg_devicemanager_install_policypackage" "test1" {
  package_name = "test-pkg1"
  timeout      = 5
}
```

## Argument Reference
The following arguments are supported:

* `package_name` - (Required) The installation package name.
* `timeout` - Timeout for installing the package to the target, default: 3 minutes.
* `adom` - Source ADOM name. default is 'root'

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `package_name` - The installation package name.
* `timeout` - Timeout for installing the package to the target, default: 3 minutes.
* `adom` - Source ADOM name.
