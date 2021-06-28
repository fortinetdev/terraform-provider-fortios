---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_certificate_download"
description: |-
  Get information on fortios system certificate.
---

# Data Source: fortios_system_certificate_download
Use this data source to get information on an fortios system certificate

## Example Usage

```hcl
data "fortios_system_certificate_download" "example" {
  name = "Fortinet_Factory"
  type = "local-cer"
}

output "debug" {
  value = data.fortios_system_certificate_download.example
}
```

## Argument Reference

* `name` - (Required) Name of certificate.
* `type` - (Required) Type of certificate (local-cer|remote-cer|local-ca|remote-ca|local-csr|crl).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Name of certificate.
* `type` - Type of certificate.
* `certificate` - Certificate in PEM format.
* `certificate_details` - Certificate details. The structure of `certificate_details` block is documented below.

The `certificate_details` block contains:

* `signature_algorithm` - The algorithm used to sign the certificate.
* `public_key_algorithm` - Tag category.
* `serial_number` - Number that uniquely identifies the certificate with the CA's system. The `format`
    function can be used to convert this base 10 number into other bases, such as hex.
* `is_ca` - `true` if this certificate is a ca certificate.
* `is_valid` - `true` if the certificate is within valid time period.
* `version` - The version the certificate is in.
* `issuer` - Who verified and signed the certificate, roughly following RFC2253.
* `subject` - The entity the certificate belongs to, roughly following RFC2253.
* `not_before` - The time after which the certificate is valid, as an RFC3339 timestamp.
* `not_after` - The time until which the certificate is invalid, as an RFC3339 timestamp.
* `sha1_fingerprint` - The SHA1 fingerprint of the public key of the certificate.