---
layout: "fortios"
page_title: "FortiOS: fortios_system_license_vdom"
sidebar_current: "docs-fortios-resource-system-license-vdom"
subcategory: "FortiGate System"
description: |-
  Provides a resource to add a VDOM license for FortiOS.
---

# fortios_system_license_vdom
Provides a resource to add a VDOM license for FortiOS.

## Example Usage
```hcl
resource "fortios_system_license_vdom" "test2" {
  license = "license"
}
```

## Argument Reference
The following arguments are supported:

* `license` - (Required) Registration code.
