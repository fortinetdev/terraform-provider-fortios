---
layout: "fortios"
page_title: "FortiOS: fortios_firewall_object_servicecategory"
sidebar_current: "docs-fortios-resource-firewall-object-servicecategory"
description: |-
  Provides a resource to configure firewall service categories of FortiOS.
---

# fortios_firewall_object_servicecategory
Provides a resource to configure firewall service category of FortiOS.

## Example Usage
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_firewall_object_servicecategory" "test_category_name" {
	name = "testcategory"
	comment = "comment"
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) Category name.
* `comment` - Comment.

## Attributes Reference
The following attributes are exported:

* `id` - The ID of the firewall service item.
* `name` -Category name.
* `comment` - Comment.

