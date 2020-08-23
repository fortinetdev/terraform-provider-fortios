---
layout: "fortios"
page_title: "FortiOS: fortios_firewall_object_servicegroup"
sidebar_current: "docs-fortios-resource-firewall-object-servicegroup"
subcategory: "FortiGate OldVersion"
description: |-
  Provides a resource to configure firewall service group of FortiOS.
---

# fortios_firewall_object_servicegroup
Provides a resource to configure firewall service group of FortiOS.

~> **Warning:** The resource will be deprecated and replaced by `fortios_firewallservice_group`.

## Example Usage
```hcl
resource "fortios_firewall_object_servicegroup" "v11" {
  name    = "1fdsafd11a"
  comment = "fdsafdsa"
  member  = ["DCE-RPC", "DNS", "HTTPS"]
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

