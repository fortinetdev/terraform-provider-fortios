---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_proxyaddress"
description: |-
  Web proxy address configuration.
---

# fortios_firewall_proxyaddress
Web proxy address configuration.

## Example Usage

```hcl
resource "fortios_firewall_proxyaddress" "trname" {
  case_sensitivity = "disable"
  color            = 2
  name             = "proxyaddress1"
  referrer         = "enable"
  type             = "host-regex"
  visibility       = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Address name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `type` - Proxy address type.
* `host` - Address object for the host.
* `host_regex` - Host name as a regular expression.
* `path` - URL path as a regular expression.
* `query` - Match the query part of the URL as a regular expression.
* `referrer` - Enable/disable use of referrer field in the HTTP header to match the address. Valid values: `enable`, `disable`.
* `category` - FortiGuard category ID. The structure of `category` block is documented below.
* `method` - HTTP request methods to be used. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`.
* `ua` - Names of browsers to be used as user agent. Valid values: `chrome`, `ms`, `firefox`, `safari`, `other`.
* `header_name` - Name of HTTP header.
* `header` - HTTP header name as a regular expression.
* `case_sensitivity` - Enable to make the pattern case sensitive. Valid values: `disable`, `enable`.
* `header_group` - HTTP header group. The structure of `header_group` block is documented below.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.
* `comment` - Optional comments.
* `application` - SaaS application. The structure of `application` block is documented below.
* `visibility` - Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `category` block supports:

* `id` - Fortiguard category id.

The `header_group` block supports:

* `id` - ID.
* `header_name` - HTTP header.
* `header` - HTTP header regular expression.
* `case_sensitivity` - Case sensitivity in pattern. Valid values: `disable`, `enable`.

The `tagging` block supports:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block supports:

* `name` - Tag name.

The `application` block supports:

* `name` - SaaS applicaton name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall ProxyAddress can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_proxyaddress.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_proxyaddress.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
