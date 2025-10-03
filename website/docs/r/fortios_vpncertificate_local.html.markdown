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
* `private_key_retain` - Enable/disable retention of private key during SCEP renewal (default = disable). Valid values: `enable`, `disable`.
* `cmp_server` - Address and port for CMP server (format = address:port).
* `cmp_path` - Path location inside CMP server.
* `cmp_server_cert` - CMP server certificate.
* `cmp_regeneration_method` - CMP auto-regeneration method. Valid values: `keyupate`, `renewal`.
* `acme_ca_url` - The URL for the ACME CA server (Let's Encrypt is the default provider).
* `acme_domain` - A valid domain that resolves to this Fortigate.
* `acme_email` - Contact email address that is required by some CAs like LetsEncrypt.
* `acme_eab_key_id` - External Account Binding Key ID (optional setting).
* `acme_eab_key_hmac` - External Account Binding HMAC Key (URL-encoded base64).
* `acme_rsa_key_size` - Length of the RSA private key of the generated cert (Minimum 2048 bits).
* `acme_renew_window` - Beginning of the renewal window (in days before certificate expiration, 30 by default).
* `est_server` - Address and port for EST server (e.g. https://example.com:1234).
* `est_ca_id` - CA identifier of the CA server for signing via EST.
* `est_http_username` - HTTP Authentication username for signing via EST.
* `est_http_password` - HTTP Authentication password for signing via EST.
* `est_client_cert` - Certificate used to authenticate this FortiGate to EST server.
* `est_server_cert` - EST server's certificate must be verifiable by this certificate to be authenticated.
* `est_srp_username` - EST SRP authentication username.
* `est_srp_password` - EST SRP authentication password.
* `est_regeneration_method` - EST behavioral options during re-enrollment. Valid values: `create-new-key`, `use-existing-key`.
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
