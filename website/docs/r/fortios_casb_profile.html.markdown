---
subcategory: "FortiGate Casb"
layout: "fortios"
page_title: "FortiOS: fortios_casb_profile"
description: |-
  Configure CASB profile.
---

# fortios_casb_profile
Configure CASB profile. Applies to FortiOS Version `>= 7.4.1`.

## Argument Reference

The following arguments are supported:

* `name` - CASB profile name.
* `saas_application` - CASB profile SaaS application. The structure of `saas_application` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `saas_application` block supports:

* `name` - CASB profile SaaS application name.
* `status` - Enable/disable setting. Valid values: `enable`, `disable`.
* `safe_search` - Enable/disable safe search. Valid values: `enable`, `disable`.
* `safe_search_control` - CASB profile safe search control. The structure of `safe_search_control` block is documented below.
* `tenant_control` - Enable/disable tenant control. Valid values: `enable`, `disable`.
* `tenant_control_tenants` - CASB profile tenant control tenants. The structure of `tenant_control_tenants` block is documented below.
* `domain_control` - Enable/disable domain control. Valid values: `enable`, `disable`.
* `domain_control_domains` - CASB profile domain control domains. The structure of `domain_control_domains` block is documented below.
* `log` - Enable/disable log settings. Valid values: `enable`, `disable`.
* `access_rule` - CASB profile access rule. The structure of `access_rule` block is documented below.
* `custom_control` - CASB profile custom control. The structure of `custom_control` block is documented below.

The `safe_search_control` block supports:

* `name` - Safe search control name.

The `tenant_control_tenants` block supports:

* `name` - Tenant control tenants name.

The `domain_control_domains` block supports:

* `name` - Domain control domain name.

The `access_rule` block supports:

* `name` - CASB access rule activity name.
* `action` - CASB access rule action. Valid values: `bypass`, `block`, `monitor`.
* `bypass` - CASB bypass options. Valid values: `av`, `dlp`, `web-filter`, `file-filter`, `video-filter`.

The `custom_control` block supports:

* `name` - CASB custom control user activity name.
* `option` - CASB custom control option. The structure of `option` block is documented below.

The `option` block supports:

* `name` - CASB custom control option name.
* `user_input` - CASB custom control user input. The structure of `user_input` block is documented below.

The `user_input` block supports:

* `value` - user input value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Casb Profile can be imported using any of these accepted formats:
```
$ terraform import fortios_casb_profile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_casb_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
