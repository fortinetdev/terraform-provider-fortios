---
layout: "fortios"
page_title: "FortiOS: fortios_fmg_system_adom"
sidebar_current: "docs-fortios-fortimanager-resource-system-adom"
description: |-
  Provides a resource to configure system adom for FortiManager.
---

# fortios_fmg_system_adom
This resource supports Create/Read/Update/Delete system adom for FortiManager.

## Example Usage
```hcl
resource "fortios_fmg_system_adom" "test1" {
  name                                              = "sys_adom_test"
  type                                              = "FortiCarrier"
  central_management_vpn                            = false
  central_management_fortiap                        = true
  central_management_sdwan                          = false
  mode                                              = "Normal"
  perform_policy_check_before_every_install         = true
  action_when_conflicts_occur_during_policy_check   = "Continue"
  auto_push_policy_packages_when_device_back_online = "Enable"
  status                                            = 1
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) Administrative Domain name.
* `type` - Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
* `central_management_vpn` - True or False.
* `central_management_fortiap` - True or False.
* `central_management_sdwan` - True or False.
* `mode` - Adom mode: Normal or Backup.
* `perform_policy_check_before_every_install` - True or False.
* `action_when_conflicts_occur_during_policy_check` - True or False.
* `auto_push_policy_packages_when_device_back_online` - True or False.
* `status` - Adom status. 0 means off and 1 means on.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `name` - Administrative Domain name.
* `type` - Domain type.
* `central_management_vpn` - True or False.
* `central_management_fortiap` - True or False.
* `central_management_sdwan` - True or False.
* `mode` - Adom mode: Normal or Backup.
* `perform_policy_check_before_every_install` - True or False.
* `action_when_conflicts_occur_during_policy_check` - True or False.
* `auto_push_policy_packages_when_device_back_online` - True or False.
* `status` - Adom status.
