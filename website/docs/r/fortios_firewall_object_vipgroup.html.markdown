---
layout: "fortios"
page_title: "FortiOS: fortios_firewall_object_vipgroup"
sidebar_current: "docs-fortios-resource-firewall-object-vipgroup"
subcategory: "FortiGate"
description: |-
  Provides a resource to configure virtual IP groups of FortiOS.
---

# fortios_firewall_object_vipgroup
Provides a resource to configure virtual IP groups of FortiOS.

## Example Usage
```hcl
resource "fortios_firewall_object_vipgroup" "v11" {
  name      = "1fdsafd11a"
  interface = "port3"
  comments  = "comments"
  member    = ["vip1", "vip3"]
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) VIP group name.
* `interface` - Interface name.
* `member` - (Required) Member VIP objects of the group.
* `comments` - Comment.

## Attributes Reference
The following attributes are exported:

* `id` - The ID of the virtual IP groups item.
* `name` - VIP group name.
* `interface` - Interface name.
* `member` - Member VIP objects of the group.
* `comments` - Comment.
