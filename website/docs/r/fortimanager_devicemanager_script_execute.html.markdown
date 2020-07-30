---
layout: "fortios"
page_title: "FortiOS: fortios_fmg_devicemanager_script_execute"
sidebar_current: "docs-fortios-fortimanager-resource-devicemanager-script-execute"
subcategory: "FortiManager"
description: |-
  Provides a resource to execute devicemanager script on FortiManager.
---

# fortios_fmg_devicemanager_script_execute
This resource supports executing devicemanager script on Fortimanager.

## Example Usage
```hcl
resource "fortios_fmg_devicemanager_script_execute" "test3" {
  script_name    = "config-intf3"
  target_devname = "devname"
  timeout        = 5
}
```

## Argument Reference
The following arguments are supported:

* `script_name` - (Required) Script name.
* `target_devname` - (Required) Target device name, which the script will be installed.
* `timeout` - Timeout(minute) for executing the script, default is 3 minutes.
* `adom` - Source ADOM name. default is 'root'
* `vdom` - Vdom of managed device. default is 'root'

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `script_name` - Script name.
* `target_devname` - Target device name, which the script should be installed.
* `timeout` - Timeout(minute) for executing the script, default is 3 minutes.
* `adom` - Source ADOM name.
* `vdom` - Vdom of managed device.
