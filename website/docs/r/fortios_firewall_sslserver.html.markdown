---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_sslserver"
description: |-
  Configure SSL servers.
---

# fortios_firewall_sslserver
Configure SSL servers.

## Example Usage

```hcl
resource "fortios_firewall_sslserver" "trname" {
  add_header_x_forwarded_proto = "enable"
  ip                           = "1.1.1.1"
  mapped_port                  = 2234
  name                         = "sslserver1"
  port                         = 32321
  ssl_algorithm                = "high"
  ssl_cert                     = "Fortinet_CA_SSL"
  ssl_client_renegotiation     = "allow"
  ssl_dh_bits                  = "2048"
  ssl_max_version              = "tls-1.2"
  ssl_min_version              = "tls-1.1"
  ssl_mode                     = "half"
  ssl_send_empty_frags         = "enable"
  url_rewrite                  = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `name` - Server name.
* `ip` - (Required) IPv4 address of the SSL server.
* `port` - (Required) Server service port (1 - 65535, default = 443).
* `ssl_mode` - SSL/TLS mode for encryption and decryption of traffic.
* `add_header_x_forwarded_proto` - Enable/disable adding an X-Forwarded-Proto header to forwarded requests.
* `mapped_port` - Mapped server service port (1 - 65535, default = 80).
* `ssl_cert` - (Required) Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
* `ssl_dh_bits` - Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048).
* `ssl_algorithm` - Relative strength of encryption algorithms accepted in negotiation.
* `ssl_client_renegotiation` - Allow or block client renegotiation by server.
* `ssl_min_version` - Lowest SSL/TLS version to negotiate.
* `ssl_max_version` - Highest SSL/TLS version to negotiate.
* `ssl_send_empty_frags` - Enable/disable sending empty fragments to avoid attack on CBC IV.
* `url_rewrite` - Enable/disable rewriting the URL.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall SslServer can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_sslserver.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
