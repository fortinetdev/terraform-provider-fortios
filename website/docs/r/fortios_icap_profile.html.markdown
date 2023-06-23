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
* `comment` - Comment.
* `request` - Enable/disable whether an HTTP request is passed to an ICAP server. Valid values: `disable`, `enable`.
* `response` - Enable/disable whether an HTTP response is passed to an ICAP server. Valid values: `disable`, `enable`.
* `file_transfer` - Configure the file transfer protocols to pass transferred files to an ICAP server as REQMOD. Valid values: `ssh`, `ftp`.
* `streaming_content_bypass` - Enable/disable bypassing of ICAP server for streaming content. Valid values: `disable`, `enable`.
* `n204_size_limit` - 204 response size limit to be saved by ICAP client in megabytes (1 - 10, default = 1 MB).
* `n204_response` - Enable/disable allowance of 204 response from ICAP server. Valid values: `disable`, `enable`.
* `preview` - Enable/disable preview of data to ICAP server. Valid values: `disable`, `enable`.
* `preview_data_length` - Preview data length to be sent to ICAP server.
* `request_server` - ICAP server to use for an HTTP request.
* `response_server` - ICAP server to use for an HTTP response.
* `file_transfer_server` - ICAP server to use for a file transfer.
* `request_failure` - Action to take if the ICAP server cannot be contacted when processing an HTTP request. Valid values: `error`, `bypass`.
* `response_failure` - Action to take if the ICAP server cannot be contacted when processing an HTTP response. Valid values: `error`, `bypass`.
* `file_transfer_failure` - Action to take if the ICAP server cannot be contacted when processing a file transfer. Valid values: `error`, `bypass`.
* `request_path` - Path component of the ICAP URI that identifies the HTTP request processing service.
* `response_path` - Path component of the ICAP URI that identifies the HTTP response processing service.
* `file_transfer_path` - Path component of the ICAP URI that identifies the file transfer processing service.
* `methods` - The allowed HTTP methods that will be sent to ICAP server for further processing.
* `response_req_hdr` - Enable/disable addition of req-hdr for ICAP response modification (respmod) processing. Valid values: `disable`, `enable`.
* `respmod_default_action` - Default action to ICAP response modification (respmod) processing. Valid values: `forward`, `bypass`.
* `icap_block_log` - Enable/disable UTM log when infection found (default = disable). Valid values: `disable`, `enable`.
* `chunk_encap` - Enable/disable chunked encapsulation (default = disable). Valid values: `disable`, `enable`.
* `extension_feature` - Enable/disable ICAP extension features. Valid values: `scan-progress`.
* `scan_progress_interval` - Scan progress interval value.
* `timeout` - Time (in seconds) that ICAP client waits for the response from ICAP server.
* `icap_headers` - Configure ICAP forwarded request headers. The structure of `icap_headers` block is documented below.
* `respmod_forward_rules` - ICAP response mode forward rules. The structure of `respmod_forward_rules` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
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
$ terraform import fortios_icap_profile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_icap_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
