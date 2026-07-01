---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_proxyaddress6"
description: |-
  Configure web proxy address6.
---

# fortios_firewall_proxyaddress6
Configure web proxy address6. Applies to FortiOS Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `name` - Address6 name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `type` - Proxy address6 type. Valid values: `host-regex`, `url`, `category`, `method`, `ua`, `header`, `src-advanced`, `dst-advanced`, `saas`, `llm-server`.
* `host` - Address6 object for the host.
* `host_regex` - Host name as a regular expression.
* `path` - URL path as a regular expression.
* `query` - Match the query part of the URL as a regular expression.
* `referrer` - Enable/disable use of referrer field in the HTTP header to match the address. Valid values: `enable`, `disable`.
* `category` - FortiGuard category ID. The structure of `category` block is documented below.
* `method` - HTTP request methods to be used. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`, `update`, `patch`, `other`.
* `ua` - Names of browsers to be used as user agent. Valid values: `chrome`, `ms`, `firefox`, `safari`, `ie`, `edge`, `other`.
* `ua_min_ver` - Minimum version of the user agent specified in dotted notation. For example, use 90.0.1 with the ua field set to "chrome" to require Google Chrome's minimum version must be 90.0.1.
* `ua_max_ver` - Maximum version of the user agent specified in dotted notation. For example, use 120 with the ua field set to "chrome" to require Google Chrome's maximum version must be 120.
* `header_name` - Name of HTTP header.
* `header` - HTTP header value as a regular expression.
* `case_sensitivity` - Enable to make the pattern case sensitive. Valid values: `disable`, `enable`.
* `header_group` - HTTP header group. The structure of `header_group` block is documented below.
* `llm_servers` - LLM Proxy server names. The structure of `llm_servers` block is documented below.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.
* `comment` - Optional comments.
* `application` - SaaS application. The structure of `application` block is documented below.
* `display_with` - Display object with first tag, all tags, or just the icon. Valid values: `all-tags`, `first-tag-only`, `icon-and-color`.
* `custom_tags` - Custom tags. The structure of `custom_tags` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.

The `category` block supports:

* `id` - FortiGuard category ID.

The `header_group` block supports:

* `id` - ID.
* `header_name` - HTTP header.
* `header` - HTTP header regular expression.
* `case_sensitivity` - Case sensitivity in pattern. Valid values: `disable`, `enable`.

The `llm_servers` block supports:

* `name` - Server name.

The `tagging` block supports:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block supports:

* `name` - Tag name.

The `application` block supports:

* `name` - SaaS application name.

The `custom_tags` block supports:

* `name` - Names of custom tags used with this address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall ProxyAddress6 can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_proxyaddress6.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_proxyaddress6.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
