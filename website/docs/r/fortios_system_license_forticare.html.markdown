---
layout: "fortios"
page_title: "FortiOS: fortios_system_license_forticare"
sidebar_current: "docs-fortios-resource-system-license-forticare"
subcategory: "FortiGate System"
description: |-
  Provides a resource to add a FortiCare license for FortiOS.
---

# fortios_system_license_forticare
Provides a resource to add a FortiCare license for FortiOS.

## Example Usage
```hcl
resource "fortios_system_license_forticare" "test2" {
  registration_code = "license"
}
```

## Argument Reference
The following arguments are supported:

* `registration_code` - (Required) Registration code.
