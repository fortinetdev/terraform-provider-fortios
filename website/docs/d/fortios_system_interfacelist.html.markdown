---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_interfacelist"
description: |-
  Provides a list of fortios_system_interface.
---

# Data Source: fortios_system_interfacelist
Provides a list of `fortios_system_interface`.

## Example Usage

```hcl
data "fortios_system_interfacelist" sample1 {
  filter = "name!=port1"
}

output output1 {
  value = data.fortios_system_interfacelist.sample2.namelist
}
```

## Argument Reference

* `filter` - (Optional) A filter used to scope the list. See [Filter results of datasource](https://registry.terraform.io/providers/fortinetdev/fortios/latest/docs/guides/fgt_filter).

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `namelist` -  A list of the `fortios_system_interface`.
