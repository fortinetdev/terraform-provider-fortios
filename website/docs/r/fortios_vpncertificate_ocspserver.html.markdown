---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpncertificate_ocspserver"
description: |-
  OCSP server configuration.
---

# fortios_vpncertificate_ocspserver
OCSP server configuration.

## Example Usage

```hcl
resource "fortios_vpncertificate_ocspserver" "trname" {
  cert           = "ACCVRAIZ1"
  name           = "ocspservers1"
  source_ip      = "0.0.0.0"
  unavail_action = "revoke"
  url            = "www.tetserv.com"
}
```

## Argument Reference

The following arguments are supported:

* `name` - OCSP server entry name.
* `url` - OCSP server URL.
* `cert` - OCSP server certificate.
* `secondary_url` - Secondary OCSP server URL.
* `secondary_cert` - Secondary OCSP server certificate.
* `unavail_action` - Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
* `source_ip` - Source IP address for communications to the OCSP server.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnCertificate OcspServer can be imported using any of these accepted formats:
```
$ terraform import fortios_vpncertificate_ocspserver.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_vpncertificate_ocspserver.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
