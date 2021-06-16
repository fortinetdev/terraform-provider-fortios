---
subcategory: "FortiGate DNSFilter"
layout: "fortios"
page_title: "FortiOS: fortios_dnsfilter_domainfilter"
description: |-
  Configure DNS domain filters.
---

# fortios_dnsfilter_domainfilter
Configure DNS domain filters.

## Example Usage

```hcl
resource "fortios_dnsfilter_domainfilter" "trname" {
  fosid = 1
  name  = "test"

  entries {
    action = "block"
    domain = "bac.com"
    id     = 1
    status = "enable"
    type   = "simple"
  }
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID.
* `name` - (Required) Name of table.
* `comment` - Optional comments.
* `entries` - DNS domain filter entries. The structure of `entries` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `entries` block supports:

* `id` - Id.
* `domain` - Domain entries to be filtered.
* `type` - DNS domain filter type. Valid values: `simple`, `regex`, `wildcard`.
* `action` - Action to take for domain filter matches. Valid values: `block`, `allow`, `monitor`.
* `status` - Enable/disable this domain filter. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Dnsfilter DomainFilter can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_dnsfilter_domainfilter.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
