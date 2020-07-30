---
layout: "fortios"
page_title: "FortiOS: fortios_firewall_object_servicecategory"
sidebar_current: "docs-fortios-resource-firewall-object-servicecategory"
subcategory: "FortiGate"
description: |-
  Provides a resource to configure firewall service categories of FortiOS.
---

# fortios_firewall_object_servicecategory
Provides a resource to configure firewall service category of FortiOS.

## Example Usage
```hcl
resource "fortios_firewall_object_servicecategory" "test_category_name" {
  name    = "testcategory"
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

