---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_staticlist"
description: |-
  Provides a list of fortios_router_static.
---

# Data Source: fortios_router_staticlist
Provides a list of `fortios_router_static`.

## Example Usage

```hcl
data "fortios_router_staticlist" sample1 {
  filter="seq_num>1"
}

output output1 {
  value = data.fortios_router_staticlist.sample1.seq_numlist
}
```

## Argument Reference

* `filter` - (Optional) A filter used to scope the list. See [Filter results of datasource](https://registry.terraform.io/providers/fortinetdev/fortios/latest/docs/guides/fgt_filter).

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `seq_numlist` -  A list of the `fortios_router_static`.
