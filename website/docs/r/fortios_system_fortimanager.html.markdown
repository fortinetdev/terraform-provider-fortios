---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fortimanager"
description: |-
  Configure FortiManager.
---

# fortios_system_fortimanager
Configure FortiManager.

Due to the limitations of the current system, the feature is temporarily unavailable. Please use the following resource configuration as an alternative.

## Example
```
resource "fortios_system_centralmanagement" "trname" {
  allow_monitor                 = "enable"
  allow_push_configuration      = "enable"
  allow_push_firmware           = "enable"
  allow_remote_firmware_upgrade = "enable"
  enc_algorithm                 = "high"
  fmg                           = "\"192.168.52.177\""
  include_default_servers       = "enable"
  mode                          = "normal"
  type                          = "fortimanager"
  vdom                          = "root"
}
```

