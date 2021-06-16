---
subcategory: "FortiGate WAF"
layout: "fortios"
page_title: "FortiOS: fortios_waf_profile"
description: |-
  Web application firewall configuration.
---

# fortios_waf_profile
Web application firewall configuration.

## Example Usage

```hcl
resource "fortios_waf_profile" "trname" {
  extended_log = "disable"
  external     = "disable"
  name         = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - WAF Profile name.
* `external` - Disable/Enable external HTTP Inspection. Valid values: `disable`, `enable`.
* `extended_log` - Enable/disable extended logging. Valid values: `enable`, `disable`.
* `signature` - WAF signatures. The structure of `signature` block is documented below.
* `constraint` - WAF HTTP protocol restrictions. The structure of `constraint` block is documented below.
* `method` - Method restriction. The structure of `method` block is documented below.
* `address_list` - Black address list and white address list. The structure of `address_list` block is documented below.
* `url_access` - URL access list The structure of `url_access` block is documented below.
* `comment` - Comment.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `signature` block supports:

* `main_class` - Main signature class. The structure of `main_class` block is documented below.
* `disabled_sub_class` - Disabled signature subclasses. The structure of `disabled_sub_class` block is documented below.
* `disabled_signature` - Disabled signatures The structure of `disabled_signature` block is documented below.
* `credit_card_detection_threshold` - The minimum number of Credit cards to detect violation.
* `custom_signature` - Custom signature. The structure of `custom_signature` block is documented below.

The `main_class` block supports:

* `id` - Main signature class ID.
* `status` - Status. Valid values: `enable`, `disable`.
* `action` - Action. Valid values: `allow`, `block`, `erase`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `severity` - Severity. Valid values: `high`, `medium`, `low`.

The `disabled_sub_class` block supports:

* `id` - Signature subclass ID.

The `disabled_signature` block supports:

* `id` - Signature ID.

The `custom_signature` block supports:

* `name` - Signature name.
* `status` - Status. Valid values: `enable`, `disable`.
* `action` - Action. Valid values: `allow`, `block`, `erase`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `severity` - Severity. Valid values: `high`, `medium`, `low`.
* `direction` - Traffic direction. Valid values: `request`, `response`.
* `case_sensitivity` - Case sensitivity in pattern. Valid values: `disable`, `enable`.
* `pattern` - Match pattern.
* `target` - Match HTTP target. Valid values: `arg`, `arg-name`, `req-body`, `req-cookie`, `req-cookie-name`, `req-filename`, `req-header`, `req-header-name`, `req-raw-uri`, `req-uri`, `resp-body`, `resp-hdr`, `resp-status`.

The `constraint` block supports:

* `header_length` - HTTP header length in request. The structure of `header_length` block is documented below.
* `content_length` - HTTP content length in request. The structure of `content_length` block is documented below.
* `param_length` - Maximum length of parameter in URL, HTTP POST request or HTTP body. The structure of `param_length` block is documented below.
* `line_length` - HTTP line length in request. The structure of `line_length` block is documented below.
* `url_param_length` - Maximum length of parameter in URL. The structure of `url_param_length` block is documented below.
* `version` - Enable/disable HTTP version check. The structure of `version` block is documented below.
* `method` - Enable/disable HTTP method check. The structure of `method` block is documented below.
* `hostname` - Enable/disable hostname check. The structure of `hostname` block is documented below.
* `malformed` - Enable/disable malformed HTTP request check. The structure of `malformed` block is documented below.
* `max_cookie` - Maximum number of cookies in HTTP request. The structure of `max_cookie` block is documented below.
* `max_header_line` - Maximum number of HTTP header line. The structure of `max_header_line` block is documented below.
* `max_url_param` - Maximum number of parameters in URL. The structure of `max_url_param` block is documented below.
* `max_range_segment` - Maximum number of range segments in HTTP range line. The structure of `max_range_segment` block is documented below.
* `exception` - HTTP constraint exception. The structure of `exception` block is documented below.

The `header_length` block supports:

* `status` - Enable/disable the constraint. Valid values: `enable`, `disable`.
* `length` - Length of HTTP header in bytes (0 to 2147483647).
* `action` - Action. Valid values: `allow`, `block`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `severity` - Severity. Valid values: `high`, `medium`, `low`.

The `content_length` block supports:

* `status` - Enable/disable the constraint. Valid values: `enable`, `disable`.
* `length` - Length of HTTP content in bytes (0 to 2147483647).
* `action` - Action. Valid values: `allow`, `block`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `severity` - Severity. Valid values: `high`, `medium`, `low`.

The `param_length` block supports:

* `status` - Enable/disable the constraint. Valid values: `enable`, `disable`.
* `length` - Maximum length of parameter in URL, HTTP POST request or HTTP body in bytes (0 to 2147483647).
* `action` - Action. Valid values: `allow`, `block`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `severity` - Severity. Valid values: `high`, `medium`, `low`.

The `line_length` block supports:

* `status` - Enable/disable the constraint. Valid values: `enable`, `disable`.
* `length` - Length of HTTP line in bytes (0 to 2147483647).
* `action` - Action. Valid values: `allow`, `block`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `severity` - Severity. Valid values: `high`, `medium`, `low`.

The `url_param_length` block supports:

* `status` - Enable/disable the constraint. Valid values: `enable`, `disable`.
* `length` - Maximum length of URL parameter in bytes (0 to 2147483647).
* `action` - Action. Valid values: `allow`, `block`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `severity` - Severity. Valid values: `high`, `medium`, `low`.

The `version` block supports:

* `status` - Enable/disable the constraint. Valid values: `enable`, `disable`.
* `action` - Action. Valid values: `allow`, `block`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `severity` - Severity. Valid values: `high`, `medium`, `low`.

The `method` block supports:

* `status` - Enable/disable the constraint. Valid values: `enable`, `disable`.
* `action` - Action. Valid values: `allow`, `block`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `severity` - Severity. Valid values: `high`, `medium`, `low`.

The `hostname` block supports:

* `status` - Enable/disable the constraint. Valid values: `enable`, `disable`.
* `action` - Action. Valid values: `allow`, `block`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `severity` - Severity. Valid values: `high`, `medium`, `low`.

The `malformed` block supports:

* `status` - Enable/disable the constraint. Valid values: `enable`, `disable`.
* `action` - Action. Valid values: `allow`, `block`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `severity` - Severity. Valid values: `high`, `medium`, `low`.

The `max_cookie` block supports:

* `status` - Enable/disable the constraint. Valid values: `enable`, `disable`.
* `max_cookie` - Maximum number of cookies in HTTP request (0 to 2147483647).
* `action` - Action. Valid values: `allow`, `block`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `severity` - Severity. Valid values: `high`, `medium`, `low`.

The `max_header_line` block supports:

* `status` - Enable/disable the constraint. Valid values: `enable`, `disable`.
* `max_header_line` - Maximum number HTTP header lines (0 to 2147483647).
* `action` - Action. Valid values: `allow`, `block`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `severity` - Severity. Valid values: `high`, `medium`, `low`.

The `max_url_param` block supports:

* `status` - Enable/disable the constraint. Valid values: `enable`, `disable`.
* `max_url_param` - Maximum number of parameters in URL (0 to 2147483647).
* `action` - Action. Valid values: `allow`, `block`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `severity` - Severity. Valid values: `high`, `medium`, `low`.

The `max_range_segment` block supports:

* `status` - Enable/disable the constraint. Valid values: `enable`, `disable`.
* `max_range_segment` - Maximum number of range segments in HTTP range line (0 to 2147483647).
* `action` - Action. Valid values: `allow`, `block`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `severity` - Severity. Valid values: `high`, `medium`, `low`.

The `exception` block supports:

* `id` - Exception ID.
* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match. Valid values: `enable`, `disable`.
* `address` - Host address.
* `header_length` - HTTP header length in request. Valid values: `enable`, `disable`.
* `content_length` - HTTP content length in request. Valid values: `enable`, `disable`.
* `param_length` - Maximum length of parameter in URL, HTTP POST request or HTTP body. Valid values: `enable`, `disable`.
* `line_length` - HTTP line length in request. Valid values: `enable`, `disable`.
* `url_param_length` - Maximum length of parameter in URL. Valid values: `enable`, `disable`.
* `version` - Enable/disable HTTP version check. Valid values: `enable`, `disable`.
* `method` - Enable/disable HTTP method check. Valid values: `enable`, `disable`.
* `hostname` - Enable/disable hostname check. Valid values: `enable`, `disable`.
* `malformed` - Enable/disable malformed HTTP request check. Valid values: `enable`, `disable`.
* `max_cookie` - Maximum number of cookies in HTTP request. Valid values: `enable`, `disable`.
* `max_header_line` - Maximum number of HTTP header line. Valid values: `enable`, `disable`.
* `max_url_param` - Maximum number of parameters in URL. Valid values: `enable`, `disable`.
* `max_range_segment` - Maximum number of range segments in HTTP range line. Valid values: `enable`, `disable`.

The `method` block supports:

* `status` - Status. Valid values: `enable`, `disable`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `severity` - Severity. Valid values: `high`, `medium`, `low`.
* `default_allowed_methods` - Methods. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`, `others`.
* `method_policy` - HTTP method policy. The structure of `method_policy` block is documented below.

The `method_policy` block supports:

* `id` - HTTP method policy ID.
* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match. Valid values: `enable`, `disable`.
* `address` - Host address.
* `allowed_methods` - Allowed Methods. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`, `others`.

The `address_list` block supports:

* `status` - Status. Valid values: `enable`, `disable`.
* `blocked_log` - Enable/disable logging on blocked addresses. Valid values: `enable`, `disable`.
* `severity` - Severity. Valid values: `high`, `medium`, `low`.
* `trusted_address` - Trusted address. The structure of `trusted_address` block is documented below.
* `blocked_address` - Blocked address. The structure of `blocked_address` block is documented below.

The `trusted_address` block supports:

* `name` - Address name.

The `blocked_address` block supports:

* `name` - Address name.

The `url_access` block supports:

* `id` - URL access ID.
* `address` - Host address.
* `action` - Action. Valid values: `bypass`, `permit`, `block`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `severity` - Severity. Valid values: `high`, `medium`, `low`.
* `access_pattern` - URL access pattern. The structure of `access_pattern` block is documented below.

The `access_pattern` block supports:

* `id` - URL access pattern ID.
* `srcaddr` - Source address.
* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match. Valid values: `enable`, `disable`.
* `negate` - Enable/disable match negation. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Waf Profile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_waf_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
