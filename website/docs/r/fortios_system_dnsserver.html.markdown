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
* `mode` - DNS server mode.
* `dnsfilter_profile` - DNS filter profile.


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
