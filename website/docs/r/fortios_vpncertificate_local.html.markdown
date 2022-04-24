---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpncertificate_local"
description: |-
  Local keys and certificates.
---

# fortios_vpncertificate_local
Local keys and certificates.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `password` - Password as a PEM file.
* `comments` - Comment.
* `private_key` - PEM format key, encrypted with a password.
* `certificate` - PEM format certificate.
* `csr` - Certificate Signing Request.
* `state` - Certificate Signing Request State.
* `scep_url` - SCEP server URL.
* `range` - Either a global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
* `source` - Certificate source type.
* `auto_regenerate_days` - Number of days to wait before expiry of an updated local certificate is requested (0 = disabled).
* `auto_regenerate_days_warning` - Number of days to wait before an expiry warning message is generated (0 = disabled).
* `scep_password` - SCEP server challenge password for auto-regeneration.
* `ca_identifier` - CA identifier of the CA server for signing via SCEP.
* `name_encoding` - Name encoding method for auto-regeneration. Valid values: `printable`, `utf8`.
* `source_ip` - Source IP address for communications to the SCEP server.
* `ike_localid` - Local ID the FortiGate uses for authentication as a VPN client.
* `ike_localid_type` - IKE local ID type. Valid values: `asn1dn`, `fqdn`.
* `last_updated` - Time at which certificate was last updated.
* `enroll_protocol` - Certificate enrollment protocol.
* `cmp_server` - 'ADDRESS:PORT' for CMP server.
* `cmp_path` - Path location inside CMP server.
* `cmp_server_cert` - CMP server certificate.
* `cmp_regeneration_method` - CMP auto-regeneration method. Valid values: `keyupate`, `renewal`.
* `acme_ca_url` - The URL for the ACME CA server (Let's Encrypt is the default provider).
* `acme_domain` - A valid domain that resolves to this Fortigate.
* `acme_email` - Contact email address that is required by some CAs like LetsEncrypt.
* `acme_rsa_key_size` - Length of the RSA private key of the generated cert (Minimum 2048 bits).
* `acme_renew_window` - Beginning of the renewal window (in days before certificate expiration, 30 by default).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnCertificate Local can be imported using any of these accepted formats:
```
$ terraform import fortios_vpncertificate_local.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_vpncertificate_local.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
