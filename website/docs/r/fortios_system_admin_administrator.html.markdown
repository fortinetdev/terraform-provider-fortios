---
layout: "fortios"
page_title: "FortiOS: fortios_system_admin_administrator"
sidebar_current: "docs-fortios-resource-system-admin-administrator"
subcategory: "FortiGate OldVersion"
description: |-
  Provides a resource to configure administrator accounts of FortiOS.
---

# fortios_system_admin_administrator
Provides a resource to configure administrator accounts of FortiOS.

!> **Warning:** The resource will be deprecated and replaced by new resource `fortios_system_admin`, we recommend that you use the new resource.

## Example Usage
```hcl
resource "fortios_system_admin_administrator" "admintest" {
  name       = "testadminacc"
  password   = "cc37331AC1"
  trusthost1 = "1.1.1.0 255.255.255.0"
  trusthost2 = "2.2.2.0 255.255.255.0"
  accprofile = "3d3"
  vdom       = ["root"]
  comments   = "comments"
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) User name.
* `password` - (Required) Admin user password.
* `trusthostN` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit.
* `vdom` - Virtual domain(s) that the administrator can access.
* `accprofile` - (Required) Access profile for this administrator. Access profiles control administrator access to FortiGate features.
* `comments` - Comment.

## Attributes Reference
The following attributes are exported:

* `id` - The ID of the administrator account item.
* `name` - User name.
* `password` - Admin user password.
* `trusthostN` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit.
* `vdom` - Virtual domain(s) that the administrator can access.
* `accprofile` - Access profile for this administrator. Access profiles control administrator access to FortiGate features.
* `comments` - Comment.

