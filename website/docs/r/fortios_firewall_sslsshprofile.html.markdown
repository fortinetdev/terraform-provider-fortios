---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_sslsshprofile"
description: |-
  Configure SSL/SSH protocol options.
---

# fortios_firewall_sslsshprofile
Configure SSL/SSH protocol options.

## Argument Reference


The following arguments are supported:

* `name` - (Required) Name.
* `comment` - Optional comments.
* `ssl` - Configure SSL options. The structure of `ssl` block is documented below.
* `https` - Configure HTTPS options. The structure of `https` block is documented below.
* `ftps` - Configure FTPS options. The structure of `ftps` block is documented below.
* `imaps` - Configure IMAPS options. The structure of `imaps` block is documented below.
* `pop3s` - Configure POP3S options. The structure of `pop3s` block is documented below.
* `smtps` - Configure SMTPS options. The structure of `smtps` block is documented below.
* `ssh` - Configure SSH options. The structure of `ssh` block is documented below.
* `whitelist` - Enable/disable exempting servers by FortiGuard whitelist.
* `ssl_exempt` - Servers to exempt from SSL inspection. The structure of `ssl_exempt` block is documented below.
* `server_cert_mode` - Re-sign or replace the server's certificate.
* `use_ssl_server` - Enable/disable the use of SSL server table for SSL offloading.
* `caname` - CA certificate used by SSL Inspection.
* `untrusted_caname` - Untrusted CA certificate used by SSL Inspection.
* `server_cert` - Certificate used by SSL Inspection to replace server certificate.
* `ssl_server` - SSL servers. The structure of `ssl_server` block is documented below.
* `ssl_anomalies_log` - Enable/disable logging SSL anomalies.
* `ssl_exemptions_log` - Enable/disable logging SSL exemptions.
* `rpc_over_https` - Enable/disable inspection of RPC over HTTPS.
* `mapi_over_https` - Enable/disable inspection of MAPI over HTTPS.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `ssl` block supports:

* `inspect_all` - Level of SSL inspection.
* `client_cert_request` - Action based on client certificate request.
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported.
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate.
* `untrusted_server_cert` - Allow, ignore, or block the untrusted SSL session server certificate.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.

The `https` block supports:

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `status` - Configure protocol inspection status.
* `client_cert_request` - Action based on client certificate request.
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported.
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate.
* `untrusted_server_cert` - Allow, ignore, or block the untrusted SSL session server certificate.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.

The `ftps` block supports:

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `status` - Configure protocol inspection status.
* `client_cert_request` - Action based on client certificate request.
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported.
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate.
* `untrusted_server_cert` - Allow, ignore, or block the untrusted SSL session server certificate.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.

The `imaps` block supports:

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `status` - Configure protocol inspection status.
* `client_cert_request` - Action based on client certificate request.
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported.
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate.
* `untrusted_server_cert` - Allow, ignore, or block the untrusted SSL session server certificate.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.

The `pop3s` block supports:

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `status` - Configure protocol inspection status.
* `client_cert_request` - Action based on client certificate request.
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported.
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate.
* `untrusted_server_cert` - Allow, ignore, or block the untrusted SSL session server certificate.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.

The `smtps` block supports:

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `status` - Configure protocol inspection status.
* `client_cert_request` - Action based on client certificate request.
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported.
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate.
* `untrusted_server_cert` - Allow, ignore, or block the untrusted SSL session server certificate.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.

The `ssh` block supports:

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `status` - Configure protocol inspection status.
* `inspect_all` - Level of SSL inspection.
* `unsupported_version` - Action based on SSH version being unsupported.
* `ssh_policy_check` - Enable/disable SSH policy check.
* `ssh_tun_policy_check` - Enable/disable SSH tunnel policy check.
* `ssh_algorithm` - Relative strength of encryption algorithms accepted during negotiation.

The `ssl_exempt` block supports:

* `id` - ID number.
* `type` - Type of address object (IPv4 or IPv6) or FortiGuard category.
* `fortiguard_category` - FortiGuard category ID.
* `address` - IPv4 address object.
* `address6` - IPv6 address object.
* `wildcard_fqdn` - Exempt servers by wildcard FQDN.
* `regex` - Exempt servers by regular expression.

The `ssl_server` block supports:

* `id` - SSL server ID.
* `ip` - IPv4 address of the SSL server.
* `https_client_cert_request` - Action based on client certificate request during the HTTPS handshake.
* `smtps_client_cert_request` - Action based on client certificate request during the SMTPS handshake.
* `pop3s_client_cert_request` - Action based on client certificate request during the POP3S handshake.
* `imaps_client_cert_request` - Action based on client certificate request during the IMAPS handshake.
* `ftps_client_cert_request` - Action based on client certificate request during the FTPS handshake.
* `ssl_other_client_cert_request` - Action based on client certificate request during an SSL protocol handshake.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall SslSshProfile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_sslsshprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
