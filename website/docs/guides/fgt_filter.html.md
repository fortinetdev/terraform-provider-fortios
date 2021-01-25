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

* `==` - Pattern must be identical to the value.
* `=*` - Pattern must be identical to the value.
* `!=` - Pattern does not match the value.
* `!*` - Pattern does not match the value.
* `=@` - Pattern found within value.
* `!@` - Pattern not found within value.
* `<=` - Value must be less than or equal to pattern.
* `<` - Value must be less than pattern.
* `>=` - Value must be greater than or equal to pattern.
* `>` - Value must be greater than pattern.


### Examples:

```HCL
# To display firewall policies with the schedule "always", use:
data "fortios_firewall_policy_list" sample1 {
  filter = "schedule==always"
}

output sample1_output {
  value = data.fortios_firewall_policy_list.sample1.policyidlist
}

# To display interfaces that do not have the name "port1", use:
data "fortios_system_interface_list" sample2 {
  filter = "name!=port1"
}

output sample2_output {
  value = data.fortios_system_interface_list.sample2.namelist
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
data "fortios_firewall_policy_list" sample3 {
  filter = "schedule==always,schedule==once"
}

output sample3_output {
  value = data.fortios_firewall_policy_list.sample3.policyidlist
}

# To display firewall policies with a schedule of "always" AND an action of "accept", use:
data "fortios_firewall_policy_list" sample4 {
  filter = "schedule==always&action==accept"
}

output sample4_output {
  value = data.fortios_firewall_policy_list.sample4.namelist
}

# To display firewall policies with a schedule of "always" AND an action of either "accept" or "deny", use:
data "fortios_firewall_policy_list" sample5 {
  filter = "schedule==always&action==accept,action==deny"
}

output sample5_output {
  value = data.fortios_firewall_policy_list.sample5.namelist
}

# To display firewall policies with policyid > 3, use:
data "fortios_firewall_policy_list" sample6 {
  filter = "policyid>3"
}

output sample6 {
  value = data.fortios_firewall_policy_list.sample6.namelist
}

```
