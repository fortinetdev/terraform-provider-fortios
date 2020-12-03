---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_iptranslation"
description: |-
  Configure firewall IP-translation.
---

# fortios_firewall_iptranslation
Configure firewall IP-translation.

## Example Usage

```hcl
resource "fortios_firewall_iptranslation" "trname" {
  endip       = "2.2.2.2"
  map_startip = "0.0.0.0"
  startip     = "1.1.1.1"
  transid     = 1
  type        = "SCTP"
}
```

## Argument Reference

The following arguments are supported:

* `transid` - (Required) IP translation ID.
* `type` - IP translation type (option: SCTP).
* `startip` - (Required) First IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
* `endip` - (Required) Final IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
* `map_startip` - (Required) Address to be used as the starting point for translation in the range (format xxx.xxx.xxx.xxx, default: 0.0.0.0).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{transid}}.

## Import

Firewall IpTranslation can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_iptranslation.labelname {{transid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
