---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_sslsshprofile"
description: |-
  Configure SSL/SSH protocol options.
---

# fortios_firewall_sslsshprofile
Configure SSL/SSH protocol options.

## Example Usage

```hcl
resource "fortios_firewall_sslsshprofile" "t1" {
  name = "test1"

  ssl {
    inspect_all = "disable"
  }

  https {
    ports = "443 127 422 392"
  }
  ftps {
    ports = 990
  }
  imaps {
    ports = "993 1123"
  }
  pop3s {
    ports = 995
  }
  smtps {
    ports = 465
  }
}


resource "fortios_firewall_sslsshprofile" "t2" {
  name = "test2"

  ssl {
    inspect_all = "deep-inspection"
  }

  https {
    ports = 443
  }
}

```

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
* `dot` - Configure DNS over TLS options. The structure of `dot` block is documented below.
* `allowlist` - Enable/disable exempting servers by FortiGuard allowlist. Valid values: `enable`, `disable`.
* `block_blocklisted_certificates` - Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blocklist. Valid values: `disable`, `enable`.
* `whitelist` - Enable/disable exempting servers by FortiGuard whitelist. Valid values: `enable`, `disable`.
* `block_blacklisted_certificates` - Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blacklist. Valid values: `disable`, `enable`.
* `ssl_exempt` - Servers to exempt from SSL inspection. The structure of `ssl_exempt` block is documented below.
* `ech_outer_sni` - ClientHelloOuter SNIs to be blocked. The structure of `ech_outer_sni` block is documented below.
* `server_cert_mode` - Re-sign or replace the server's certificate. Valid values: `re-sign`, `replace`.
* `use_ssl_server` - Enable/disable the use of SSL server table for SSL offloading. Valid values: `disable`, `enable`.
* `caname` - CA certificate used by SSL Inspection.
* `untrusted_caname` - Untrusted CA certificate used by SSL Inspection.
* `server_cert` - Certificate used by SSL Inspection to replace server certificate.
* `ssl_server` - SSL servers. The structure of `ssl_server` block is documented below.
* `ssl_exemption_ip_rating` - Enable/disable IP based URL rating. Valid values: `enable`, `disable`.
* `ssl_anomaly_log` - Enable/disable logging SSL anomalies. Valid values: `disable`, `enable`.
* `ssl_exemption_log` - Enable/disable logging SSL exemptions. Valid values: `disable`, `enable`.
* `ssl_anomalies_log` - Enable/disable logging SSL anomalies. Valid values: `disable`, `enable`.
* `ssl_exemptions_log` - Enable/disable logging SSL exemptions. Valid values: `disable`, `enable`.
* `ssl_negotiation_log` - Enable/disable logging SSL negotiation. Valid values: `disable`, `enable`.
* `ssl_server_cert_log` - Enable/disable logging of server certificate information. Valid values: `disable`, `enable`.
* `ssl_handshake_log` - Enable/disable logging of TLS handshakes. Valid values: `disable`, `enable`.
* `rpc_over_https` - Enable/disable inspection of RPC over HTTPS. Valid values: `enable`, `disable`.
* `mapi_over_https` - Enable/disable inspection of MAPI over HTTPS. Valid values: `enable`, `disable`.
* `supported_alpn` - Configure ALPN option. Valid values: `http1-1`, `http2`, `all`, `none`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `ssl` block supports:

* `inspect_all` - Level of SSL inspection. Valid values: `disable`, `certificate-inspection`, `deep-inspection`.
* `client_certificate` - Action based on received client certificate. Valid values: `bypass`, `inspect`, `block`.
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported.
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow`, `block`.
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow`, `block`.
* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow`, `block`, `ignore`.
* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow`, `block`, `ignore`.
* `client_cert_request` - Action based on client certificate request. Valid values: `bypass`, `inspect`, `block`.
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported. Valid values: `bypass`, `inspect`, `block`.
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate. Valid values: `allow`, `block`.
* `untrusted_server_cert` - Allow, ignore, or block the untrusted SSL session server certificate. Valid values: `allow`, `block`, `ignore`.
* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow`, `block`, `ignore`.
* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow`, `block`, `ignore`.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `enable`, `strict`, `disable`.
* `cert_probe_failure` - Action based on certificate probe failure. Valid values: `allow`, `block`.
* `encrypted_client_hello` - Block/allow session based on existence of encrypted-client-hello. Valid values: `allow`, `block`.
* `min_allowed_ssl_version` - Minimum SSL version to be allowed. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

The `https` block supports:

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `status` - Configure protocol inspection status. Valid values: `disable`, `certificate-inspection`, `deep-inspection`.
* `quic` - QUIC inspection status. On FortiOS versions 7.4.1: default = disable. On FortiOS versions >= 7.4.2: default = inspect.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable`, `disable`.
* `client_certificate` - Action based on received client certificate. Valid values: `bypass`, `inspect`, `block`.
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported.
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow`, `block`.
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow`, `block`.
* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow`, `block`, `ignore`.
* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow`, `block`, `ignore`.
* `client_cert_request` - Action based on client certificate request. Valid values: `bypass`, `inspect`, `block`.
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported. Valid values: `bypass`, `inspect`, `block`.
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate. Valid values: `allow`, `block`.
* `untrusted_server_cert` - Allow, ignore, or block the untrusted SSL session server certificate. Valid values: `allow`, `block`, `ignore`.
* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow`, `block`, `ignore`.
* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow`, `block`, `ignore`.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `enable`, `strict`, `disable`.
* `cert_probe_failure` - Action based on certificate probe failure. Valid values: `allow`, `block`.
* `encrypted_client_hello` - Block/allow session based on existence of encrypted-client-hello. Valid values: `allow`, `block`.
* `min_allowed_ssl_version` - Minimum SSL version to be allowed. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

The `ftps` block supports:

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `status` - Configure protocol inspection status. Valid values: `disable`, `deep-inspection`.
* `client_certificate` - Action based on received client certificate. Valid values: `bypass`, `inspect`, `block`.
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported.
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow`, `block`.
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow`, `block`.
* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow`, `block`, `ignore`.
* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow`, `block`, `ignore`.
* `client_cert_request` - Action based on client certificate request. Valid values: `bypass`, `inspect`, `block`.
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported. Valid values: `bypass`, `inspect`, `block`.
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate. Valid values: `allow`, `block`.
* `untrusted_server_cert` - Allow, ignore, or block the untrusted SSL session server certificate. Valid values: `allow`, `block`, `ignore`.
* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow`, `block`, `ignore`.
* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow`, `block`, `ignore`.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `enable`, `strict`, `disable`.
* `min_allowed_ssl_version` - Minimum SSL version to be allowed. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

The `imaps` block supports:

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `status` - Configure protocol inspection status. Valid values: `disable`, `deep-inspection`.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable`, `disable`.
* `client_certificate` - Action based on received client certificate. Valid values: `bypass`, `inspect`, `block`.
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported.
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow`, `block`.
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow`, `block`.
* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow`, `block`, `ignore`.
* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow`, `block`, `ignore`.
* `client_cert_request` - Action based on client certificate request. Valid values: `bypass`, `inspect`, `block`.
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported. Valid values: `bypass`, `inspect`, `block`.
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate. Valid values: `allow`, `block`.
* `untrusted_server_cert` - Allow, ignore, or block the untrusted SSL session server certificate. Valid values: `allow`, `block`, `ignore`.
* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow`, `block`, `ignore`.
* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow`, `block`, `ignore`.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `enable`, `strict`, `disable`.

The `pop3s` block supports:

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `status` - Configure protocol inspection status. Valid values: `disable`, `deep-inspection`.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable`, `disable`.
* `client_certificate` - Action based on received client certificate. Valid values: `bypass`, `inspect`, `block`.
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported.
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow`, `block`.
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow`, `block`.
* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow`, `block`, `ignore`.
* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow`, `block`, `ignore`.
* `client_cert_request` - Action based on client certificate request. Valid values: `bypass`, `inspect`, `block`.
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported. Valid values: `bypass`, `inspect`, `block`.
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate. Valid values: `allow`, `block`.
* `untrusted_server_cert` - Allow, ignore, or block the untrusted SSL session server certificate. Valid values: `allow`, `block`, `ignore`.
* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow`, `block`, `ignore`.
* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow`, `block`, `ignore`.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `enable`, `strict`, `disable`.

The `smtps` block supports:

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `status` - Configure protocol inspection status. Valid values: `disable`, `deep-inspection`.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable`, `disable`.
* `client_certificate` - Action based on received client certificate. Valid values: `bypass`, `inspect`, `block`.
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported.
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow`, `block`.
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow`, `block`.
* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow`, `block`, `ignore`.
* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow`, `block`, `ignore`.
* `client_cert_request` - Action based on client certificate request. Valid values: `bypass`, `inspect`, `block`.
* `unsupported_ssl` - Action based on the SSL encryption used being unsupported. Valid values: `bypass`, `inspect`, `block`.
* `invalid_server_cert` - Allow or block the invalid SSL session server certificate. Valid values: `allow`, `block`.
* `untrusted_server_cert` - Allow, ignore, or block the untrusted SSL session server certificate. Valid values: `allow`, `block`, `ignore`.
* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow`, `block`, `ignore`.
* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow`, `block`, `ignore`.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `enable`, `strict`, `disable`.

The `ssh` block supports:

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `status` - Configure protocol inspection status. Valid values: `disable`, `deep-inspection`.
* `inspect_all` - Level of SSL inspection. Valid values: `disable`, `deep-inspection`.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable`, `disable`.
* `unsupported_version` - Action based on SSH version being unsupported. Valid values: `bypass`, `block`.
* `ssh_policy_check` - Enable/disable SSH policy check. Valid values: `disable`, `enable`.
* `ssh_tun_policy_check` - Enable/disable SSH tunnel policy check. Valid values: `disable`, `enable`.
* `ssh_algorithm` - Relative strength of encryption algorithms accepted during negotiation. Valid values: `compatible`, `high-encryption`.

The `dot` block supports:

* `status` - Configure protocol inspection status. Valid values: `disable`, `deep-inspection`.
* `quic` - QUIC inspection status. On FortiOS versions 7.4.1: default = disable. On FortiOS versions >= 7.4.2: default = inspect.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable`, `disable`.
* `client_certificate` - Action based on received client certificate. Valid values: `bypass`, `inspect`, `block`.
* `unsupported_ssl_version` - Action based on the SSL version used being unsupported.
* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow`, `block`.
* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow`, `block`.
* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow`, `block`, `ignore`.
* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow`, `block`, `ignore`.
* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA. Valid values: `allow`, `block`, `ignore`.
* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow`, `block`, `ignore`.
* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow`, `block`, `ignore`.
* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `enable`, `strict`, `disable`.

The `ssl_exempt` block supports:

* `id` - ID number.
* `type` - Type of address object (IPv4 or IPv6) or FortiGuard category. Valid values: `fortiguard-category`, `address`, `address6`, `wildcard-fqdn`, `regex`.
* `fortiguard_category` - FortiGuard category ID.
* `address` - IPv4 address object.
* `address6` - IPv6 address object.
* `wildcard_fqdn` - Exempt servers by wildcard FQDN.
* `regex` - Exempt servers by regular expression.

The `ech_outer_sni` block supports:

* `name` - ClientHelloOuter SNI name.
* `sni` - ClientHelloOuter SNI to be blocked.

The `ssl_server` block supports:

* `id` - SSL server ID.
* `ip` - IPv4 address of the SSL server.
* `https_client_certificate` - Action based on received client certificate during the HTTPS handshake. Valid values: `bypass`, `inspect`, `block`.
* `smtps_client_certificate` - Action based on received client certificate during the SMTPS handshake. Valid values: `bypass`, `inspect`, `block`.
* `pop3s_client_certificate` - Action based on received client certificate during the POP3S handshake. Valid values: `bypass`, `inspect`, `block`.
* `imaps_client_certificate` - Action based on received client certificate during the IMAPS handshake. Valid values: `bypass`, `inspect`, `block`.
* `ftps_client_certificate` - Action based on received client certificate during the FTPS handshake. Valid values: `bypass`, `inspect`, `block`.
* `ssl_other_client_certificate` - Action based on received client certificate during an SSL protocol handshake. Valid values: `bypass`, `inspect`, `block`.
* `https_client_cert_request` - Action based on client certificate request during the HTTPS handshake. Valid values: `bypass`, `inspect`, `block`.
* `smtps_client_cert_request` - Action based on client certificate request during the SMTPS handshake. Valid values: `bypass`, `inspect`, `block`.
* `pop3s_client_cert_request` - Action based on client certificate request during the POP3S handshake. Valid values: `bypass`, `inspect`, `block`.
* `imaps_client_cert_request` - Action based on client certificate request during the IMAPS handshake. Valid values: `bypass`, `inspect`, `block`.
* `ftps_client_cert_request` - Action based on client certificate request during the FTPS handshake. Valid values: `bypass`, `inspect`, `block`.
* `ssl_other_client_cert_request` - Action based on client certificate request during an SSL protocol handshake. Valid values: `bypass`, `inspect`, `block`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall SslSshProfile can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_sslsshprofile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_sslsshprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
