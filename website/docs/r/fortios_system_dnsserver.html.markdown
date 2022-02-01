---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dnsserver"
description: |-
  Configure DNS servers.
---

# fortios_system_dnsserver
Configure DNS servers.

## Example Usage

```hcl
resource "fortios_system_dnsserver" "trname" {
  dnsfilter_profile = "default"
  mode              = "forward-only"
  name              = "port3"
}
```

## Argument Reference

The following arguments are supported:

* `name` - DNS server name.
* `mode` - DNS server mode. Valid values: `recursive`, `non-recursive`, `forward-only`.
* `dnsfilter_profile` - DNS filter profile.
* `doh` - DNS over HTTPS. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System DnsServer can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_dnsserver.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
