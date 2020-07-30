---
layout: "fortios"
page_title: "FortiOS: fortios_fmg_firewall_object_vip"
sidebar_current: "docs-fortios-fortimanager-resource-firewall-object-vip"
subcategory: "FortiManager"
description: |-
  Provides a resource to configure firewall object virtual ip for FortiManager.
---

# fortios_fmg_firewall_object_vip
This resource supports Create/Read/Update/Delete firewall object virtual ip for FortiManager.

## Example Usage
```hcl
resource "fortios_fmg_firewall_object_vip" "test1" {
  name           = "fov-test1"
  comment        = "test obj vip"
  type           = "static-nat"
  arp_reply      = "enable"
  ext_ip         = "2.2.2.2"
  ext_intf       = "any"
  mapped_ip      = "1.1.1.1"
  config_default = "enable"
}

resource "fortios_fmg_firewall_object_vip" "test2" {
  name           = "fov-test2"
  comment        = "test obj vip"
  type           = "fqdn"
  arp_reply      = "disable"
  ext_ip         = "2.2.2.2-2.2.2.100"
  mapped_addr    = "update.microsoft.com"
  config_default = "enable"
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) Virtual IP name.
* `comment` - Comments.
* `type` -  Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]
* `ext_ip` - IP address or address range on the external interface that you want to map to an address or address range on the destination network.
* `ext_intf` - Interface connected to the source network that receives the packets that will be forwarded to the destination network.
* `arp_reply` - Enable to respond to ARP requests for this virtual IP address. Enabled by default.
* `mapped_ip` - Mapped Ip address.
* `config_default` - Enable/Disable config default value. enabled by default.
* `mapped_addr` - Mapped FQDN address name.
* `adom` - ADOM name. default is 'root'.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `name` - Virtual IP name.
* `comment` - Comments.
* `type` -  Virtual IP type.
* `ext_ip` - IP address or address range on the external interface that you want to map to an address or address range on the destination network.
* `ext_intf` - Interface connected to the source network that receives the packets that will be forwarded to the destination network.
* `arp_reply` - Enable to respond to ARP requests for this virtual IP address.
* `mapped_ip` - Mapped Ip address.
* `config_default` - Enable/Disable config default value.
* `mapped_addr` - Mapped FQDN address name.
* `adom` - ADOM name.
