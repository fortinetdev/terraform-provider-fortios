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

* `name` - Debug URL name.
* `url_pattern` - (Required) URL exemption pattern.
* `status` - Enable/disable this URL exemption. Valid values: `enable`, `disable`.
* `exact` - Enable/disable matching the exact path. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy DebugUrl can be imported using any of these accepted formats:
```
$ terraform import fortios_webproxy_debugurl.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_webproxy_debugurl.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
