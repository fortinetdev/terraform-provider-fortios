---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_addresslist"
description: |-
  Provides a list of fortios_firewall_address.
---

# Data Source: fortios_firewall_addresslist
Provides a list of `fortios_firewall_address`.

## Example Usage

```hcl
data "fortios_firewall_addresslist" sample7 {
  filter = "name!=google-play"
}

output sample7 {
  value = data.fortios_firewall_addresslist.sample7.namelist
}

resource "fortios_firewall_addrgrp" grp1 {
  name = "addrgrptest1"

  dynamic_sort_subtable = "true"

  dynamic "member" {
    for_each = toset(data.fortios_firewall_addresslist.sample7.namelist)
    content {
      name = member.key
    }
  }
}
```

## Argument Reference

* `filter` - (Optional) A filter used to scope the list. See [Filter results of datasource](https://registry.terraform.io/providers/fortinetdev/fortios/latest/docs/guides/fgt_filter).

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `namelist` -  A list of the `fortios_firewall_address`.
