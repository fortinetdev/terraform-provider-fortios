---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_proxyaddress"
description: |-
  Get information on an fortios firewall proxyaddress.
---

# Data Source: fortios_firewall_proxyaddress
Use this data source to get information on an fortios firewall proxyaddress

## Argument Reference

* `name` - (Required) Specify the name of the desired firewall proxyaddress.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Address name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `type` - Proxy address type.
* `host` - Address object for the host.
* `host_regex` - Host name as a regular expression.
* `path` - URL path as a regular expression.
* `query` - Match the query part of the URL as a regular expression.
* `referrer` - Enable/disable use of referrer field in the HTTP header to match the address.
* `category` - FortiGuard category ID. The structure of `category` block is documented below.
* `method` - HTTP request methods to be used.
* `ua` - Names of browsers to be used as user agent.
* `header_name` - Name of HTTP header.
* `header` - HTTP header name as a regular expression.
* `case_sensitivity` - Enable to make the pattern case sensitive.
* `header_group` - HTTP header group. The structure of `header_group` block is documented below.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.
* `comment` - Optional comments.
* `visibility` - Enable/disable visibility of the object in the GUI.

The `category` block contains:

* `id` - Fortiguard category id.

The `header_group` block contains:

* `id` - ID.
* `header_name` - HTTP header.
* `header` - HTTP header regular expression.
* `case_sensitivity` - Case sensitivity in pattern.

The `tagging` block contains:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.

