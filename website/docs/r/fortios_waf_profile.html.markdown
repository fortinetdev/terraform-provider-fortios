---
subcategory: "FortiGate Waf"
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
* `external` - Disable/Enable external HTTP Inspection.
* `extended_log` - Enable/disable extended logging.
* `signature` - WAF signatures.
* `constraint` - WAF HTTP protocol restrictions.
* `method` - Method restriction.
* `address_list` - Black address list and white address list.
* `url_access` - URL access list
* `comment` - Comment.

The `signature` block supports:

* `main_class` - Main signature class.
* `disabled_sub_class` - Disabled signature subclasses.
* `disabled_signature` - Disabled signatures
* `credit_card_detection_threshold` - The minimum number of Credit cards to detect violation.
* `custom_signature` - Custom signature.

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
* `exception` - HTTP constraint exception.

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
* `method_policy` - HTTP method policy.

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
* `trusted_address` - Trusted address.
* `blocked_address` - Blocked address.

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
* `access_pattern` - URL access pattern.

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
