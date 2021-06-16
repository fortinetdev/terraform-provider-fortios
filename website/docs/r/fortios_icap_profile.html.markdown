---
subcategory: "FortiGate ICAP"
layout: "fortios"
page_title: "FortiOS: fortios_icap_profile"
description: |-
  Configure ICAP profiles.
---

# fortios_icap_profile
Configure ICAP profiles.

## Example Usage

```hcl
resource "fortios_icap_profile" "trname" {
  methods                  = "delete get head options post put trace other"
  name                     = "1"
  request                  = "disable"
  request_failure          = "error"
  response                 = "disable"
  response_failure         = "error"
  response_req_hdr         = "disable"
  streaming_content_bypass = "disable"

  icap_headers {
    base64_encoding = "disable"
    content         = "$user"
    name            = "X-Authenticated-User"
  }
}
```

## Argument Reference

The following arguments are supported:

* `replacemsg_group` - Replacement message group.
* `name` - ICAP profile name.
* `request` - Enable/disable whether an HTTP request is passed to an ICAP server. Valid values: `disable`, `enable`.
* `response` - Enable/disable whether an HTTP response is passed to an ICAP server. Valid values: `disable`, `enable`.
* `streaming_content_bypass` - Enable/disable bypassing of ICAP server for streaming content. Valid values: `disable`, `enable`.
* `preview` - Enable/disable preview of data to ICAP server. Valid values: `disable`, `enable`.
* `preview_data_length` - Preview data length to be sent to ICAP server.
* `request_server` - ICAP server to use for an HTTP request.
* `response_server` - ICAP server to use for an HTTP response.
* `request_failure` - Action to take if the ICAP server cannot be contacted when processing an HTTP request. Valid values: `error`, `bypass`.
* `response_failure` - Action to take if the ICAP server cannot be contacted when processing an HTTP response. Valid values: `error`, `bypass`.
* `request_path` - Path component of the ICAP URI that identifies the HTTP request processing service.
* `response_path` - Path component of the ICAP URI that identifies the HTTP response processing service.
* `methods` - The allowed HTTP methods that will be sent to ICAP server for further processing. Valid values: `delete`, `get`, `head`, `options`, `post`, `put`, `trace`, `other`.
* `response_req_hdr` - Enable/disable addition of req-hdr for ICAP response modification (respmod) processing. Valid values: `disable`, `enable`.
* `respmod_default_action` - Default action to ICAP response modification (respmod) processing. Valid values: `forward`, `bypass`.
* `icap_headers` - Configure ICAP forwarded request headers. The structure of `icap_headers` block is documented below.
* `respmod_forward_rules` - ICAP response mode forward rules. The structure of `respmod_forward_rules` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `icap_headers` block supports:

* `id` - HTTP forwarded header ID.
* `name` - HTTP forwarded header name.
* `content` - HTTP header content.
* `base64_encoding` - Enable/disable use of base64 encoding of HTTP content. Valid values: `disable`, `enable`.

The `respmod_forward_rules` block supports:

* `name` - Address name.
* `host` - Address object for the host.
* `header_group` - HTTP header group. The structure of `header_group` block is documented below.
* `action` - Action to be taken for ICAP server. Valid values: `forward`, `bypass`.
* `http_resp_status_code` - HTTP response status code. The structure of `http_resp_status_code` block is documented below.

The `header_group` block supports:

* `id` - ID.
* `header_name` - HTTP header.
* `header` - HTTP header regular expression.
* `case_sensitivity` - Enable/disable case sensitivity when matching header. Valid values: `disable`, `enable`.

The `http_resp_status_code` block supports:

* `code` - HTTP response status code.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Icap Profile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_icap_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
