---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_routerbgp_neighborlist"
description: |-
  Provides a list of fortios_routerbgp_neighbor.
---

# Data Source: fortios_routerbgp_neighborlist
Provides a list of `fortios_routerbgp_neighbor`.

## Example Usage

```hcl
data "fortios_routerbgp_neighborlist" sample1 {
}

output output1 {
  value = data.fortios_routerbgp_neighborlist.sample1.iplist
}
```

## Argument Reference

* `filter` - (Optional) A filter used to scope the list. See [Filter results of datasource](https://registry.terraform.io/providers/fortinetdev/fortios/latest/docs/guides/fgt_filter).

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `iplist` -  A list of the `fortios_routerbgp_neighbor`.
