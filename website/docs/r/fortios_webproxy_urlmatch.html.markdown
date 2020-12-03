---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_urlmatch"
description: |-
  Exempt URLs from web proxy forwarding and caching.
---

# fortios_webproxy_urlmatch
Exempt URLs from web proxy forwarding and caching.

## Example Usage

```hcl
resource "fortios_webproxy_forwardserver" "trname2" {
  addr_type          = "fqdn"
  healthcheck        = "disable"
  ip                 = "0.0.0.0"
  monitor            = "http://www.google.com"
  name               = "test2"
  port               = 3128
  server_down_option = "block"
}

resource "fortios_webproxy_urlmatch" "trname" {
  cache_exemption = "disable"
  forward_server  = fortios_webproxy_forwardserver.trname2.name
  name            = "1"
  status          = "enable"
  url_pattern     = "/examples/servlet/*Servlet"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Configure a name for the URL to be exempted.
* `status` - Enable/disable exempting the URLs matching the URL pattern from web proxy forwarding and caching.
* `url_pattern` - (Required) URL pattern to be exempted from web proxy forwarding and caching.
* `forward_server` - Forward server name.
* `cache_exemption` - Enable/disable exempting this URL pattern from caching.
* `comment` - Comment.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy UrlMatch can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_webproxy_urlmatch.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
