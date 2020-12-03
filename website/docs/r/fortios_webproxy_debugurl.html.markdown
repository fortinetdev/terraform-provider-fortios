---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_debugurl"
description: |-
  Configure debug URL addresses.
---

# fortios_webproxy_debugurl
Configure debug URL addresses.

## Example Usage

```hcl
resource "fortios_webproxy_debugurl" "trname" {
  exact       = "enable"
  name        = "1"
  status      = "enable"
  url_pattern = "/examples/servlet/*Servlet"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Debug URL name.
* `url_pattern` - (Required) URL exemption pattern.
* `status` - Enable/disable this URL exemption.
* `exact` - Enable/disable matching the exact path.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy DebugUrl can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_webproxy_debugurl.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
