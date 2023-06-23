---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrolleracl_ingress"
description: |-
  Configure ingress ACL policies to be applied on managed FortiSwitch ports.
---

# fortios_switchcontrolleracl_ingress
Configure ingress ACL policies to be applied on managed FortiSwitch ports. Applies to FortiOS Version `>= 7.4.0`.

## Argument Reference

The following arguments are supported:

* `fosid` - ACL ID.
* `description` - Description for the ACL policy.
* `action` - ACL actions. The structure of `action` block is documented below.
* `classifier` - ACL classifiers. The structure of `classifier` block is documented below.
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `action` block supports:

* `drop` - Enable/disable drop. Valid values: `enable`, `disable`.
* `count` - Enable/disable count. Valid values: `enable`, `disable`.

The `classifier` block supports:

* `dst_ip_prefix` - Destination IP address to be matched.
* `dst_mac` - Destination MAC address to be matched.
* `src_ip_prefix` - Source IP address to be matched.
* `src_mac` - Source MAC address to be matched.
* `vlan` - VLAN ID to be matched.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

SwitchControllerAcl Ingress can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontrolleracl_ingress.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontrolleracl_ingress.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
