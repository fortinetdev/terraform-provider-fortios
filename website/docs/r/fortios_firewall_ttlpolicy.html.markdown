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
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

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
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_ttlpolicy.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
