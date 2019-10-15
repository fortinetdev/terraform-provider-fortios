---
layout: "fortios"
page_title: "FortiOS: fortios_firewall_object_vipgroup"
sidebar_current: "docs-fortios-resource-firewall-object-vipgroup"
description: |-
  Provides a resource to configure virtual IP groups of FortiOS.
---

# fortios_firewall_object_vipgroup
Provides a resource to configure virtual IP groups of FortiOS.

## Example Usage
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_firewall_object_vipgroup" "v11" {
	name = "1fdsafd11a"
	interface = "port3"
	comments = "comments"
	member = ["vip1", "vip3"]
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
