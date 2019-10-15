---
layout: "fortios"
page_title: "FortiOS: fortios_system_license_vdom"
sidebar_current: "docs-fortios-resource-system-license-vdom"
description: |-
  Provides a resource to add a VDOM license for FortiOS.
---

# fortios_system_license_vdom
Provides a resource to add a VDOM license for FortiOS.

## Example Usage
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
}

resource "fortios_system_license_vdom" "test2" {
	license = "license"
}
```

## Argument Reference
The following arguments are supported:

* `license` - (Required) Registration code.
