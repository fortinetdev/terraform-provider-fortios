---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_policylist"
description: |-
  Provides a list of fortios_firewall_policy.
---

# Data Source: fortios_firewall_policylist
Provides a list of `fortios_firewall_policy`.

## Example Usage

```hcl
data "fortios_firewall_policylist" sample1 {
}

output output1 {
  value = data.fortios_firewall_policylist.sample1
}

data "fortios_firewall_policylist" sample2 {
  filter = "schedule==always&action==accept,action==deny"
}

output sample2_output {
  value = data.fortios_firewall_policylist.sample2.policyidlist
}
```

## Argument Reference

* `filter` - (Optional) A filter used to scope the list. See [Filter results of datasource](https://registry.terraform.io/providers/fortinetdev/fortios/latest/docs/guides/fgt_filter).

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `policyidlist` -  A list of the `fortios_firewall_policy`.
