---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_dnstranslation"
description: |-
  Configure DNS translation.
---

# fortios_firewall_dnstranslation
Configure DNS translation.

## Example Usage

```hcl
resource "fortios_firewall_dnstranslation" "trname" {
  dst     = "2.2.2.2"
  fosid   = 1
  netmask = "255.0.0.0"
  src     = "1.1.1.1"
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - ID.
* `src` - IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
* `dst` - IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
* `netmask` - If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall Dnstranslation can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_dnstranslation.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
