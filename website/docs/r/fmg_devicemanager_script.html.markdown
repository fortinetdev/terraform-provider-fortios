---
layout: "fortios"
page_title: "FortiOS: fortios_fmg_devicemanager_script"
sidebar_current: "docs-fortios-fortimanager-resource-devicemanager-script"
subcategory: "FortiManager"
description: |-
  Provides a resource to configure devicemanager script for FortiManager.
---

# fortios_fmg_devicemanager_script
This resource supports Create/Read/Update/Delete devicemanager script for FortiManager.

## Example Usage
```hcl
resource "fortios_fmg_devicemanager_script" "test1" {
  name        = "config-intf3"
  description = "description"
  content     = "config system interface \n edit port3 \n\t set vdom \"root\"\n\t set ip 10.7.0.200 255.255.0.0 \n\t set allowaccess ping http https\n\t next \n end"
  target      = "remote_device"
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) Script name.
* `description` - Description.
* `content` - Script content, only cli script is supported now
* `target` - Script target, Enum: ["device_database", "remote_device", "adom_database"]
* `adom` - ADOM name. default is 'root'.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `name` - Script name.
* `description` - Description.
* `content` - Script content.
* `target` - Script target.
* `adom` - ADOM name.
