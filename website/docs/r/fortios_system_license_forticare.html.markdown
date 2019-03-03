---
layout: "fortios"
page_title: "FortiOS: fortios_system_license_forticare"
sidebar_current: "docs-fortios-resource-system-license-forticare"
description: |-
  Provides a resource to add a FortiCare license for FortiOS.
---

# fortios_system_license_forticare
Provides a resource to add a FortiCare license for FortiOS.

## Example Usage
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_system_license_forticare" "test2" {
	registration_code = "fdasfdsafdsafdsa"
}
```

## Argument Reference
The following arguments are supported:
* `registration_code` - (Required) Registration code.
