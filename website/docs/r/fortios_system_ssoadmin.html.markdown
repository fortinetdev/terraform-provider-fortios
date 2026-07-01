---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ssoadmin"
description: |-
  Configure SSO admin users.
---

# fortios_system_ssoadmin
Configure SSO admin users. Applies to FortiOS Version `>= 6.2.4`.

## Example Usage

```hcl
resource "fortios_system_ssoadmin" "trname" {
  accprofile = "super_admin"
  name       = "admin1"

  vdom {
    name = "root"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - SSO admin name.
* `accprofile` - (Required) SSO admin user access profile.
* `gui_theme_type` - Use predefined themes or custom themes. Valid values: `predefined`, `custom`.
* `gui_theme` - Predefined theme that overrides the default FortiGate theme. Valid values: `jade`, `neutrino`, `mariner`, `graphite`, `melongene`, `jet-stream`, `security-fabric`, `retro`, `dark-matter`, `onyx`, `eclipse`, `none`.
* `gui_custom_theme` - Custom theme that overrides the default FortiGate theme.
* `gui_llm_provider` - Select the LLM provider. Valid values: `fortiai`, `openai`.
* `openai_api_key` - OpenAI API key.
* `openai_api_key_part2` - OpenAI API key part 2 for excess length.
* `openai_model` - OpenAI model.
* `openai_project_id` - OpenAI project ID.
* `openai_org_id` - OpenAI organization ID.
* `vdom` - Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
* `gui_ignore_release_overview_version` - The FortiOS version to ignore release overview prompt for.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.

The `vdom` block supports:

* `name` - Virtual domain name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SsoAdmin can be imported using any of these accepted formats:
```
$ terraform import fortios_system_ssoadmin.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_ssoadmin.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
