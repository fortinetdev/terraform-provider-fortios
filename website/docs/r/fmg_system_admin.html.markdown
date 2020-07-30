---
layout: "fortios"
page_title: "FortiOS: fortios_fmg_system_admin"
sidebar_current: "docs-fortios-fortimanager-resource-system-admin"
subcategory: "FortiManager"
description: |-
  Provides a resource to configure system admin setting for FortiManager.
---

# fortios_fmg_system_admin
This resource supports modifying system admin setting for FortiManager.

## Example Usage
```hcl
resource "fortios_fmg_system_admin" "test1" {
  http_port    = 80
  https_port   = 443
  idle_timeout = 20
}
```

## Argument Reference
The following arguments are supported:

* `http_port` - Http port.
* `https_port` - Https port.
* `idle_timeout` - Idle Timeout(1-480 minute).

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `http_port` - Http port.
* `https_port` - Https port.
* `idle_timeout` - Idle Timeout(1-480 minute).
