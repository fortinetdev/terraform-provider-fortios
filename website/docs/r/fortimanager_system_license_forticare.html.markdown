---
layout: "fortios"
page_title: "FortiOS: fortios_fmg_system_license_forticare"
sidebar_current: "docs-fortios-fortimanager-resource-system-license-forticare"
description: |-
  Provides a resource to upload FortiCare registration code to FortiGate through FortiManager.
---

# fortios_fmg_system_license_forticare
This resource supports uploading FortiCare registration code to FortiGate through FortiManager.

## Example Usage
```hcl
resource "fortios_fmg_system_license_forticare" "test1" {
  target            = "fortigate-test"
  registration_code = "jn3t3Nw7qckQzt955Htkfj5hwQ6aaa"
}
```

## Argument Reference
The following arguments are supported:

* `target` - (Required) Target name, which is managed by FortiManager.
* `registration_code` - (Required) Registration code.
* `adom` - ADOM that the target device belongs to. default is 'root'.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `target` - Target name, which is managed by FortiManager.
* `registration_code` - Registration code.
* `adom` - ADOM name.
