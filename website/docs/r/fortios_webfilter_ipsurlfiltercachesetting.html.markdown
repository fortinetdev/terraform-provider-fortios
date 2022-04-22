---
subcategory: "FortiGate WebFilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_ipsurlfiltercachesetting"
description: |-
  Configure IPS URL filter cache settings.
---

# fortios_webfilter_ipsurlfiltercachesetting
Configure IPS URL filter cache settings.

## Example Usage

```hcl
resource "fortios_webfilter_ipsurlfiltercachesetting" "trname" {
  dns_retry_interval = 0
  extended_ttl       = 0
}
```

## Argument Reference

The following arguments are supported:

* `dns_retry_interval` - Retry interval. Refresh DNS faster than TTL to capture multiple IPs for hosts. 0 means use DNS server's TTL only.
* `extended_ttl` - Extend time to live beyond reported by DNS. 0 means use DNS server's TTL
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Webfilter IpsUrlfilterCacheSetting can be imported using any of these accepted formats:
```
$ terraform import fortios_webfilter_ipsurlfiltercachesetting.labelname WebfilterIpsUrlfilterCacheSetting

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_webfilter_ipsurlfiltercachesetting.labelname WebfilterIpsUrlfilterCacheSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
