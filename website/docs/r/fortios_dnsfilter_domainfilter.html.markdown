---
subcategory: "FortiGate Dnsfilter"
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
* `entries` - (Required) DNS domain filter entries.

The `entries` block supports:

* `id` - Id.
* `domain` - Domain entries to be filtered.
* `type` - DNS domain filter type.
* `action` - Action to take for domain filter matches.
* `status` - Enable/disable this domain filter.


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
