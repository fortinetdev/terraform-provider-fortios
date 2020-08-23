---
subcategory: "FortiGate Vpn"
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
* `unavail_action` - Action when server is unavailable (revoke the certificate or ignore the result of the check).
* `source_ip` - Source IP address for communications to the OCSP server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnCertificate OcspServer can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_vpncertificate_ocspserver.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
