---
layout: "fortios"
page_title: "FortiOS: fortios_firewall_object_ippool"
sidebar_current: "docs-fortios-resource-firewall-object-ippool"
subcategory: "FortiGate OldVersion"
description: |-
  Provides a resource to configure IPv4 IP address pools of FortiOS.
---

# fortios_firewall_object_ippool
Provides a resource to configure IPv4 IP address pools of FortiOS.

!> **Warning:** The resource will be deprecated and replaced by new resource `fortios_firewall_ippool`, we recommend that you use the new resource.

## Example Usage for Overload Ippool
```hcl
resource "fortios_firewall_object_ippool" "s1" {
  name      = "ddd"
  type      = "overload"
  startip   = "11.0.0.0"
  endip     = "22.0.0.0"
  arp_reply = "enable"
  comments  = "fdsaf"
}
```

## Example Usage for One-to-one Ippool
```hcl
resource "fortios_firewall_object_ippool" "s2" {
  name      = "dd22d"
  type      = "one-to-one"
  startip   = "121.0.0.0"
  endip     = "222.0.0.0"
  arp_reply = "enable"
  comments  = "fdsaf"
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) IP pool name.
* `type` - (Required) IP pool type(Support overload and one-to-one).
* `startip` - (Required) First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
* `endip` - (Required) Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
* `arp_reply` - Enable/disable replying to ARP requests when an IP Pool is added to a policy.
* `comments` - Comment.

## Attributes Reference
The following attributes are exported:

* `id` - The ID of the IP pool item.
* `name` - IP pool name.
* `type` - IP pool type(Support overload and one-to-one).
* `startip` - First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
* `endip` - Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).
* `arp_reply` - Enable/disable replying to ARP requests when an IP Pool is added to a policy.
* `comments` - Comment.
