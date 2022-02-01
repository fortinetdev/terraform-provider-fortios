---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fortimanager"
description: |-
  Configure FortiManager.
---

# fortios_system_fortimanager
Configure FortiManager. Applies to FortiOS Version `<= 7.0.1`.

By design considerations, the feature is using the fortios_system_centralmanagement resource as documented below.

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

