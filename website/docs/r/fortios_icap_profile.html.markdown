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
* `request` - Enable/disable whether an HTTP request is passed to an ICAP server.
* `response` - Enable/disable whether an HTTP response is passed to an ICAP server.
* `streaming_content_bypass` - Enable/disable bypassing of ICAP server for streaming content.
* `request_server` - ICAP server to use for an HTTP request.
* `response_server` - ICAP server to use for an HTTP response.
* `request_failure` - Action to take if the ICAP server cannot be contacted when processing an HTTP request.
* `response_failure` - Action to take if the ICAP server cannot be contacted when processing an HTTP response.
* `request_path` - Path component of the ICAP URI that identifies the HTTP request processing service.
* `response_path` - Path component of the ICAP URI that identifies the HTTP response processing service.
* `methods` - The allowed HTTP methods that will be sent to ICAP server for further processing.
* `response_req_hdr` - Enable/disable addition of req-hdr for ICAP response modification (respmod) processing.
* `icap_headers` - Configure ICAP forwarded request headers. The structure of `icap_headers` block is documented below.

The `icap_headers` block supports:

* `id` - HTTP forwarded header ID.
* `name` - HTTP forwarded header name.
* `content` - HTTP header content.
* `base64_encoding` - Enable/disable use of base64 encoding of HTTP content.


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
