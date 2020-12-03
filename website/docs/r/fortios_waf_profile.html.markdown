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

* `name` - (Required) WAF Profile name.
* `external` - Disable/Enable external HTTP Inspection.
* `extended_log` - Enable/disable extended logging.
* `signature` - WAF signatures. The structure of `signature` block is documented below.
* `constraint` - WAF HTTP protocol restrictions. The structure of `constraint` block is documented below.
* `method` - Method restriction. The structure of `method` block is documented below.
* `address_list` - Black address list and white address list. The structure of `address_list` block is documented below.
* `url_access` - URL access list The structure of `url_access` block is documented below.
* `comment` - Comment.

The `signature` block supports:

* `main_class` - Main signature class. The structure of `main_class` block is documented below.
* `disabled_sub_class` - Disabled signature subclasses. The structure of `disabled_sub_class` block is documented below.
* `disabled_signature` - Disabled signatures The structure of `disabled_signature` block is documented below.
* `credit_card_detection_threshold` - The minimum number of Credit cards to detect violation.
* `custom_signature` - Custom signature. The structure of `custom_signature` block is documented below.

The `main_class` block supports:

* `id` - Main signature class ID.
* `status` - Status.
* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.

The `disabled_sub_class` block supports:

* `id` - Signature subclass ID.

The `disabled_signature` block supports:

* `id` - Signature ID.

The `custom_signature` block supports:

* `name` - Signature name.
* `status` - Status.
* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.
* `direction` - Traffic direction.
* `case_sensitivity` - Case sensitivity in pattern.
* `pattern` - Match pattern.
* `target` - Match HTTP target.

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

* `status` - Enable/disable the constraint.
* `length` - Length of HTTP header in bytes (0 to 2147483647).
* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.

The `content_length` block supports:

* `status` - Enable/disable the constraint.
* `length` - Length of HTTP content in bytes (0 to 2147483647).
* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.

The `param_length` block supports:

* `status` - Enable/disable the constraint.
* `length` - Maximum length of parameter in URL, HTTP POST request or HTTP body in bytes (0 to 2147483647).
* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.

The `line_length` block supports:

* `status` - Enable/disable the constraint.
* `length` - Length of HTTP line in bytes (0 to 2147483647).
* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.

The `url_param_length` block supports:

* `status` - Enable/disable the constraint.
* `length` - Maximum length of URL parameter in bytes (0 to 2147483647).
* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.

The `version` block supports:

* `status` - Enable/disable the constraint.
* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.

The `method` block supports:

* `status` - Enable/disable the constraint.
* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.

The `hostname` block supports:

* `status` - Enable/disable the constraint.
* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.

The `malformed` block supports:

* `status` - Enable/disable the constraint.
* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.

The `max_cookie` block supports:

* `status` - Enable/disable the constraint.
* `max_cookie` - Maximum number of cookies in HTTP request (0 to 2147483647).
* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.

The `max_header_line` block supports:

* `status` - Enable/disable the constraint.
* `max_header_line` - Maximum number HTTP header lines (0 to 2147483647).
* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.

The `max_url_param` block supports:

* `status` - Enable/disable the constraint.
* `max_url_param` - Maximum number of parameters in URL (0 to 2147483647).
* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.

The `max_range_segment` block supports:

* `status` - Enable/disable the constraint.
* `max_range_segment` - Maximum number of range segments in HTTP range line (0 to 2147483647).
* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.

The `exception` block supports:

* `id` - Exception ID.
* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match.
* `address` - Host address.
* `header_length` - HTTP header length in request.
* `content_length` - HTTP content length in request.
* `param_length` - Maximum length of parameter in URL, HTTP POST request or HTTP body.
* `line_length` - HTTP line length in request.
* `url_param_length` - Maximum length of parameter in URL.
* `version` - Enable/disable HTTP version check.
* `method` - Enable/disable HTTP method check.
* `hostname` - Enable/disable hostname check.
* `malformed` - Enable/disable malformed HTTP request check.
* `max_cookie` - Maximum number of cookies in HTTP request.
* `max_header_line` - Maximum number of HTTP header line.
* `max_url_param` - Maximum number of parameters in URL.
* `max_range_segment` - Maximum number of range segments in HTTP range line.

The `method` block supports:

* `status` - Status.
* `log` - Enable/disable logging.
* `severity` - Severity.
* `default_allowed_methods` - Methods.
* `method_policy` - HTTP method policy. The structure of `method_policy` block is documented below.

The `method_policy` block supports:

* `id` - HTTP method policy ID.
* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match.
* `address` - Host address.
* `allowed_methods` - Allowed Methods.

The `address_list` block supports:

* `status` - Status.
* `blocked_log` - Enable/disable logging on blocked addresses.
* `severity` - Severity.
* `trusted_address` - Trusted address. The structure of `trusted_address` block is documented below.
* `blocked_address` - Blocked address. The structure of `blocked_address` block is documented below.

The `trusted_address` block supports:

* `name` - Address name.

The `blocked_address` block supports:

* `name` - Address name.

The `url_access` block supports:

* `id` - URL access ID.
* `address` - Host address.
* `action` - Action.
* `log` - Enable/disable logging.
* `severity` - Severity.
* `access_pattern` - URL access pattern. The structure of `access_pattern` block is documented below.

The `access_pattern` block supports:

* `id` - URL access pattern ID.
* `srcaddr` - Source address.
* `pattern` - URL pattern.
* `regex` - Enable/disable regular expression based pattern match.
* `negate` - Enable/disable match negation.


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
