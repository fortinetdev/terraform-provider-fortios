---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpncertificate_setting"
description: |-
  VPN certificate setting.
---

# fortios_vpncertificate_setting
VPN certificate setting.

## Example Usage

```hcl
resource "fortios_vpncertificate_setting" "trname" {
  certname_dsa1024      = "Fortinet_SSL_DSA1024"
  certname_dsa2048      = "Fortinet_SSL_DSA2048"
  certname_ecdsa256     = "Fortinet_SSL_ECDSA256"
  certname_ecdsa384     = "Fortinet_SSL_ECDSA384"
  certname_rsa1024      = "Fortinet_SSL_RSA1024"
  certname_rsa2048      = "Fortinet_SSL_RSA2048"
  check_ca_cert         = "enable"
  check_ca_chain        = "disable"
  cmp_save_extra_certs  = "disable"
  cn_match              = "substring"
  ocsp_option           = "server"
  ocsp_status           = "disable"
  ssl_min_proto_version = "default"
  strict_crl_check      = "disable"
  strict_ocsp_check     = "disable"
  subject_match         = "substring"
}
```

## Argument Reference

The following arguments are supported:

* `ocsp_status` - Enable/disable receiving certificates using the OCSP.
* `ocsp_option` - Specify whether the OCSP URL is from certificate or configured OCSP server.
* `ocsp_default_server` - Default OCSP server.
* `check_ca_cert` - Enable/disable verification of the user certificate and pass authentication if any CA in the chain is trusted (default = enable).
* `check_ca_chain` - Enable/disable verification of the entire certificate chain and pass authentication only if the chain is complete and all of the CAs in the chain are trusted (default = disable).
* `subject_match` - When searching for a matching certificate, control how to find matches in the certificate subject name.
* `cn_match` - When searching for a matching certificate, control how to find matches in the cn attribute of the certificate subject name.
* `strict_crl_check` - Enable/disable strict mode CRL checking.
* `strict_ocsp_check` - Enable/disable strict mode OCSP checking.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `cmp_save_extra_certs` - Enable/disable saving extra certificates in CMP mode.
* `certname_rsa1024` - (Required) 1024 bit RSA key certificate for re-signing server certificates for SSL inspection.
* `certname_rsa2048` - (Required) 2048 bit RSA key certificate for re-signing server certificates for SSL inspection.
* `certname_dsa1024` - (Required) 1024 bit DSA key certificate for re-signing server certificates for SSL inspection.
* `certname_dsa2048` - (Required) 2048 bit DSA key certificate for re-signing server certificates for SSL inspection.
* `certname_ecdsa256` - (Required) 256 bit ECDSA key certificate for re-signing server certificates for SSL inspection.
* `certname_ecdsa384` - (Required) 384 bit ECDSA key certificate for re-signing server certificates for SSL inspection.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

VpnCertificate Setting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_vpncertificate_setting.labelname VpnCertificateSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
