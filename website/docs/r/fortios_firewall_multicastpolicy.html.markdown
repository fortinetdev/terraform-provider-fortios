---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_multicastpolicy"
description: |-
  Configure multicast NAT policies.
---

# fortios_firewall_multicastpolicy
Configure multicast NAT policies.

## Example Usage

```hcl
resource "fortios_firewall_multicastpolicy" "trname" {
  action     = "accept"
  dnat       = "0.0.0.0"
  dstintf    = "port4"
  end_port   = 65535
  fosid      = 1
  logtraffic = "enable"
  protocol   = 0
  snat       = "disable"
  snat_ip    = "0.0.0.0"
  srcintf    = "port3"
  start_port = 1
  status     = "enable"

  dstaddr {
    name = "all"
  }

  srcaddr {
    name = "all"
  }
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - Policy ID.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `name` - Policy name.
* `comments` - Comment.
* `status` - Enable/disable this policy. Valid values: `enable`, `disable`.
* `logtraffic` - Enable/disable logging traffic accepted by this policy.
* `srcintf` - (Required) Source interface name.
* `dstintf` - (Required) Destination interface name.
* `srcaddr` - (Required) Source address objects. The structure of `srcaddr` block is documented below.
* `dstaddr` - (Required) Destination address objects. The structure of `dstaddr` block is documented below.
* `snat` - Enable/disable substitution of the outgoing interface IP address for the original source IP address (called source NAT or SNAT). Valid values: `enable`, `disable`.
* `snat_ip` - IPv4 address to be used as the source address for NATed traffic.
* `dnat` - IPv4 DNAT address used for multicast destination addresses.
* `action` - Accept or deny traffic matching the policy. Valid values: `accept`, `deny`.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
* `end_port` -  Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
* `utm_status` - Enable to add an IPS security profile to the policy. Valid values: `enable`, `disable`.
* `ips_sensor` - Name of an existing IPS sensor.
* `auto_asic_offload` - Enable/disable offloading policy traffic for hardware acceleration. Valid values: `enable`, `disable`.
* `traffic_shaper` - Traffic shaper to apply to traffic forwarded by the multicast policy.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `srcaddr` block supports:

* `name` - Source address objects.

The `dstaddr` block supports:

* `name` - Destination address objects.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall MulticastPolicy can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_multicastpolicy.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_multicastpolicy.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
