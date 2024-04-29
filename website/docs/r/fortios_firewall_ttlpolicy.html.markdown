---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_ttlpolicy"
description: |-
  Configure TTL policies.
---

# fortios_firewall_ttlpolicy
Configure TTL policies.

## Example Usage

```hcl
resource "fortios_firewall_ttlpolicy" "trname" {
  action   = "accept"
  fosid    = 1
  schedule = "always"
  srcintf  = "port3"
  status   = "enable"
  ttl      = "23"

  service {
    name = "ALL"
  }

  srcaddr {
    name = "all"
  }
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID.
* `status` - Enable/disable this TTL policy. Valid values: `enable`, `disable`.
* `action` - Action to be performed on traffic matching this policy (default = deny). Valid values: `accept`, `deny`.
* `srcintf` - (Required) Source interface name from available interfaces.
* `srcaddr` - (Required) Source address object(s) from available options. Separate multiple names with a space. The structure of `srcaddr` block is documented below.
* `service` - (Required) Service object(s) from available options. Separate multiple names with a space. The structure of `service` block is documented below.
* `schedule` - (Required) Schedule object from available options.
* `ttl` - (Required) Value/range to match against the packet's Time to Live value (format: ttl[ - ttl_high], 1 - 255).
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `srcaddr` block supports:

* `name` - Address name.

The `service` block supports:

* `name` - Service name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall TtlPolicy can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_ttlpolicy.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_ttlpolicy.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
