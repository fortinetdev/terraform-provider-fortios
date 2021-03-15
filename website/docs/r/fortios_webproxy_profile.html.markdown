---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_profile"
description: |-
  Configure web proxy profiles.
---

# fortios_webproxy_profile
Configure web proxy profiles.

## Example Usage

```hcl
resource "fortios_webproxy_profile" "trname" {
  header_client_ip              = "pass"
  header_front_end_https        = "pass"
  header_via_request            = "add"
  header_via_response           = "pass"
  header_x_authenticated_groups = "pass"
  header_x_authenticated_user   = "pass"
  header_x_forwarded_for        = "pass"
  log_header_change             = "disable"
  name                          = "1"
  strip_encoding                = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Profile name.
* `header_client_ip` - Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
* `header_via_request` - Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
* `header_via_response` - Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
* `header_x_forwarded_for` - Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
* `header_front_end_https` - Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
* `header_x_authenticated_user` - Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
* `header_x_authenticated_groups` - Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
* `strip_encoding` - Enable/disable stripping unsupported encoding from the request header. Valid values: `enable`, `disable`.
* `log_header_change` - Enable/disable logging HTTP header changes. Valid values: `enable`, `disable`.
* `headers` - Configure HTTP forwarded requests headers. The structure of `headers` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `headers` block supports:

* `id` - HTTP forwarded header id.
* `name` - HTTP forwarded header name.
* `dstaddr` - Destination address and address group names. The structure of `dstaddr` block is documented below.
* `dstaddr6` - Destination address and address group names (IPv6). The structure of `dstaddr6` block is documented below.
* `action` - Action when the HTTP header is forwarded. Valid values: `add-to-request`, `add-to-response`, `remove-from-request`, `remove-from-response`.
* `content` - HTTP header content.
* `base64_encoding` - Enable/disable use of base64 encoding of HTTP content. Valid values: `disable`, `enable`.
* `add_option` - Configure options to append content to existing HTTP header or add new HTTP header. Valid values: `append`, `new-on-not-found`, `new`.
* `protocol` - Configure protocol(s) to take add-option action on (HTTP, HTTPS, or both). Valid values: `https`, `http`.

The `dstaddr` block supports:

* `name` - Address name.

The `dstaddr6` block supports:

* `name` - Address name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy Profile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_webproxy_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
