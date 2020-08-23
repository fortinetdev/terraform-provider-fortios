---
layout: "fortios"
page_title: "FortiOS: fortios_system_vdom_setting"
sidebar_current: "docs-fortios-resource-system-vdom-setting"
subcategory: "FortiGate OldVersion"
description: |-
  Provides a resource to configure VDOM of FortiOS. The API user of the token for this feature should have a super admin profile, It can be set in CLI while GUI does not allow.
---

# fortios_system_vdom_setting
Provides a resource to configure VDOM of FortiOS. The API user of the token for this feature should have a super admin profile, It can be set in CLI while GUI does not allow.

~> **Warning:** The resource will be deprecated and replaced by `fortios_system_vdom`.

## Example Usage
```hcl
resource "fortios_system_vdom_setting" "test2" {
  name       = "aa1122"
  short_name = "aa1122"
  temporary  = 0
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) VDOM name.
* `short_name` - VDOM short name.
* `temporary` - Temporary.

## Attributes Reference
The following attributes are exported:

* `id` - The ID of the VDOM.
* `name` - VDOM name.
* `short_name` - VDOM short name.
* `temporary` - Temporary.


