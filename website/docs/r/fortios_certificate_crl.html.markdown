---
subcategory: "FortiGate Certificate"
layout: "fortios"
page_title: "FortiOS: fortios_certificate_crl"
description: |-
  Certificate Revocation List as a PEM file.
---

# fortios_certificate_crl
Certificate Revocation List as a PEM file.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `crl` - Certificate Revocation List as a PEM file.
* `range` - Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
* `source` - Certificate source type.
* `update_vdom` - VDOM for CRL update.
* `ldap_server` - LDAP server name for CRL auto-update.
* `ldap_username` - LDAP server user name.
* `ldap_password` - LDAP server user password.
* `http_url` - HTTP server URL for CRL auto-update.
* `scep_url` - SCEP server URL for CRL auto-update.
* `scep_cert` - Local certificate for SCEP communication for CRL auto-update.
* `update_interval` - Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
* `source_ip` - Source IP address for communications to a HTTP or SCEP CA server.
* `last_updated` - Time at which CRL was last updated.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Certificate Crl can be imported using any of these accepted formats:
```
$ terraform import fortios_certificate_crl.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_certificate_crl.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
