---
layout: "fortios"
page_title: "FortiOS: fortios_firewall_object_servicegroup"
sidebar_current: "docs-fortios-resource-firewall-object-servicegroup"
description: |-
  Provides a resource to configure firewall service group of FortiOS.
---

# fortios_firewall_object_servicegroup
Provides a resource to configure firewall service group of FortiOS.

## Example Usage
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_firewall_object_servicegroup" "v11" {
	name = "1fdsafd11a"
	comment = "fdsafdsa"
	member = ["DCE-RPC", "DNS", "HTTPS"]
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) Service group name.
* `member` - (Required) Service objects contained within the group.
* `comment` - Comment.

## Attributes Reference
The following attributes are exported:

* `id` - The ID of the firewall service group item.
* `name` - Service group name.
* `member` - Service objects contained within the group.
* `comment` - Comment.

