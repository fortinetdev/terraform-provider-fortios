---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallssl_setting"
description: |-
  SSL proxy settings.
---

# fortios_firewallssl_setting
SSL proxy settings.

## Example Usage

```hcl
resource "fortios_firewallssl_setting" "trname" {
  abbreviate_handshake      = "enable"
  cert_cache_capacity       = 200
  cert_cache_timeout        = 10
  kxp_queue_threshold       = 16
  no_matching_cipher_action = "bypass"
  proxy_connect_timeout     = 30
  session_cache_capacity    = 500
  session_cache_timeout     = 20
  ssl_dh_bits               = "2048"
  ssl_queue_threshold       = 32
  ssl_send_empty_frags      = "enable"
}
```

## Argument Reference


The following arguments are supported:

* `proxy_connect_timeout` - (Required) Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
* `ssl_dh_bits` - (Required) Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048).
* `ssl_send_empty_frags` - (Required) Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only).
* `no_matching_cipher_action` - (Required) Bypass or drop the connection when no matching cipher is found.
* `cert_cache_capacity` - (Required) Maximum capacity of the host certificate cache (0 - 500, default = 200).
* `cert_cache_timeout` - (Required) Time limit to keep certificate cache (1 - 120 min, default = 10).
* `session_cache_capacity` - (Required) Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
* `session_cache_timeout` - (Required) Time limit to keep SSL session state (1 - 60 min, default = 20).
* `kxp_queue_threshold` - Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
* `ssl_queue_threshold` - Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
* `abbreviate_handshake` - Enable/disable use of SSL abbreviated handshake.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

FirewallSsl Setting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewallssl_setting.labelname FirewallSslSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
