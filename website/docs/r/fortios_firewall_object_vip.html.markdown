---
layout: "fortios"
page_title: "FortiOS: fortios_firewall_object_vip"
sidebar_current: "docs-fortios-resource-firewall-object-vip"
subcategory: "FortiGate OldVersion"
description: |-
  Provides a resource to configure firewall virtual IPs (VIPs) of FortiOS.
---

# fortios_firewall_object_vip
Provides a resource to configure firewall virtual IPs (VIPs) of FortiOS.

!> **Warning:** The resource will be deprecated and replaced by new resource `fortios_firewall_vip`, we recommend that you use the new resource.

## Example Usage
```hcl
resource "fortios_firewall_object_vip" "v11" {
  name        = "dfa"
  comment     = "fdsafdsafds"
  extip       = "11.1.1.1-21.1.1.1"
  mappedip    = ["22.2.2.2-32.2.2.2"]
  extintf     = "port3"
  portforward = "enable"
  protocol    = "tcp"
  extport     = "2-3"
  mappedport  = "4-5"
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) Virtual IP name.
* `extip` - (Required) IP address or address range on the external interface that you want to map to an address or address range on the
destination network.
* `mappedip` - (Required) IP address or address range on the destination network to which the external IP address is mapped.
* `extintf` - Interface connected to the source network that receives the packets that will be forwarded to the destination network.
* `portforward` - Enable/disable port forwarding.
* `protocol` - Protocol to use when forwarding packets.
* `extport` - Incoming port number range that you want to map to a port number range on the destination network.
* `mappedport` - Port number range on the destination network to which the external port number range is mapped.
* `comment` - Comment.

## Attributes Reference
The following attributes are exported:

* `id` - The ID of the firewall virtual IPs item.
* `name` - Virtual IP name.
* `extip` - IP address or address range on the external interface that you want to map to an address or address range on the
destination network.
* `mappedip` - IP address or address range on the destination network to which the external IP address is mapped.
* `extintf` - Interface connected to the source network that receives the packets that will be forwarded to the destination network.
* `portforward` - Enable/disable port forwarding.
* `protocol` - Protocol to use when forwarding packets.
* `extport` - Incoming port number range that you want to map to a port number range on the destination network.
* `mappedport` - Port number range on the destination network to which the external port number range is mapped.
* `comment` - Comment.

