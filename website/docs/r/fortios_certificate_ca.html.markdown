---
subcategory: "FortiGate Certificate"
layout: "fortios"
page_title: "FortiOS: fortios_certificate_ca"
description: |-
  CA certificate.
---

# fortios_certificate_ca
CA certificate.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `ca` - (Required) CA certificate as a PEM file.
* `range` - Either global or VDOM IP address range for the CA certificate.
* `source` - CA certificate source type.
* `trusted` - Enable/disable as a trusted CA.
* `scep_url` - URL of the SCEP server.
* `auto_update_days` - Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
* `auto_update_days_warning` - Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
* `source_ip` - Source IP address for communications to the SCEP server.
* `last_updated` - Time at which CA was last updated.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Certificate Ca can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_certificate_ca.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
