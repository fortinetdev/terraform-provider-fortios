---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_centralsnatmap"
description: |-
  Get information on an fortios firewall centralsnatmap.
---

# Data Source: fortios_firewall_centralsnatmap
Use this data source to get information on an fortios firewall centralsnatmap

## Argument Reference

* `policyid` - (Required) Specify the policyid of the desired firewall centralsnatmap.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `policyid` - Policy ID.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `status` - Enable/disable the active status of this policy.
* `type` - IPv4/IPv6 source NAT.
* `orig_addr` - Original address. The structure of `orig_addr` block is documented below.
* `orig_addr6` - IPv6 Original address. The structure of `orig_addr6` block is documented below.
* `srcintf` - Source interface name from available interfaces. The structure of `srcintf` block is documented below.
* `dst_addr` - Destination address name from available addresses. The structure of `dst_addr` block is documented below.
* `dst_addr6` - IPv6 Destination address. The structure of `dst_addr6` block is documented below.
* `dstintf` - Destination interface name from available interfaces. The structure of `dstintf` block is documented below.
* `nat_ippool` - Name of the IP pools to be used to translate addresses from available IP Pools. The structure of `nat_ippool` block is documented below.
* `nat_ippool6` - IPv6 pools to be used for source NAT. The structure of `nat_ippool6` block is documented below.
* `protocol` - Integer value for the protocol type (0 - 255).
* `orig_port` - Original TCP port (0 to 65535).
* `nat_port` - Translated port or port range (0 to 65535).
* `dst_port` - Destination port or port range (1 to 65535, 0 means any port).
* `nat` - Enable/disable source NAT.
* `nat46` - Enable/disable NAT46.
* `nat64` - Enable/disable NAT64.
* `comments` - Comment.

The `orig_addr` block contains:

* `name` - Address name.

The `orig_addr6` block contains:

* `name` - Address name.

The `srcintf` block contains:

* `name` - Interface name.

The `dst_addr` block contains:

* `name` - Address name.

The `dst_addr6` block contains:

* `name` - Address name.

The `dstintf` block contains:

* `name` - Interface name.

The `nat_ippool` block contains:

* `name` - IP pool name.

The `nat_ippool6` block contains:

* `name` - IPv6 pool name.

