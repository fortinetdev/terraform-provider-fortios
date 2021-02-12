---
subcategory: ""
layout: "fortios"
page_title: "To filter results of datasource"
description: |-
  Filter results of datasource
---

# Filter results of datasource

Filter results of the list type datasources.

## Filter
**Syntax**: `[key][operator][pattern]`

The following filter operators are supported:

* `==` - Pattern must be identical to the value, case-sensitive.
* `=*` - Pattern must be identical to the value, case-insensitive.
* `!=` - Pattern does not match the value, case-sensitive.
* `!*` - Pattern does not match the value, case-insensitive.
* `=@` - Pattern found within value, case-insensitive.
* `!@` - Pattern not found within value, case-insensitive.
* `<=` - Value must be less than or equal to pattern.
* `<` - Value must be less than pattern.
* `>=` - Value must be greater than or equal to pattern.
* `>` - Value must be greater than pattern.


### Examples:

```HCL
# To display firewall policies with the schedule "always", use:
data "fortios_firewall_policylist" sample1 {
  filter = "schedule==always"
}

output sample1_output {
  value = data.fortios_firewall_policylist.sample1.policyidlist
}

# To display interfaces that do not have the name "port1", use:
data "fortios_system_interfacelist" sample2 {
  filter = "name!=port1"
}

output sample2_output {
  value = data.fortios_system_interfacelist.sample2.namelist
}

```

## Combination

To create a complex query, filters can be combined as follows:

* `Logical OR` - Separate filters using commas `,`.
* `Logical AND` - Separate filters using ampersands `&`.
* `Combining AND and OR` - Separate filters using commas `,` and ampersands `&`.

### Examples:
```HCL
# To display firewall policies using the "always" schedule OR the "once" schedule, use:
data "fortios_firewall_policylist" sample3 {
  filter = "schedule==always,schedule==once"
}

output sample3_output {
  value = data.fortios_firewall_policylist.sample3.policyidlist
}

# To display firewall policies with a schedule of "always" AND an action of "accept", use:
data "fortios_firewall_policylist" sample4 {
  filter = "schedule==always&action==accept"
}

output sample4_output {
  value = data.fortios_firewall_policylist.sample4.policyidlist
}

# To display firewall policies with a schedule of "always" AND an action of either "accept" or "deny", use:
data "fortios_firewall_policylist" sample5 {
  filter = "schedule==always&action==accept,action==deny"
}

output sample5_output {
  value = data.fortios_firewall_policylist.sample5.policyidlist
}

# To display firewall policies with policyid > 3, use:
data "fortios_firewall_policylist" sample6 {
  filter = "policyid>3"
}

output sample6 {
  value = data.fortios_firewall_policylist.sample6.policyidlist
}

# To display firewall addresses with name!=google-play and create an addrgrp object with the list, use:
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
