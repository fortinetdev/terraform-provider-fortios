---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpncertificate_ca"
description: |-
  CA certificate.
---

# fortios_vpncertificate_ca
CA certificate.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `ca` - (Required) CA certificate as a PEM file.
* `range` - Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
* `source` - CA certificate source type.
* `ssl_inspection_trusted` - Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
* `trusted` - Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
* `scep_url` - URL of the SCEP server.
* `auto_update_days` - Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
* `auto_update_days_warning` - Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
* `source_ip` - Source IP address for communications to the SCEP server.
* `ca_identifier` - CA identifier of the SCEP server.
* `obsolete` - Enable/disable this CA as obsoleted. Valid values: `disable`, `enable`.
* `last_updated` - Time at which CA was last updated.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnCertificate Ca can be imported using any of these accepted formats:
```
$ terraform import fortios_vpncertificate_ca.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_vpncertificate_ca.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
