---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_emailserver"
description: |-
  Get information on fortios system emailserver.
---

# Data Source: fortios_system_emailserver
Use this data source to get information on fortios system emailserver

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `type` - Use FortiGuard Message service or custom email server.
* `reply_to` - Reply-To email address.
* `server` - SMTP server IP address or hostname.
* `port` - SMTP server port.
* `source_ip` - SMTP server IPv4 source IP.
* `source_ip6` - SMTP server IPv6 source IP.
* `authenticate` - Enable/disable authentication.
* `validate_server` - Enable/disable validation of server certificate.
* `username` - SMTP server user name for authentication.
* `password` - SMTP server user password for authentication.
* `security` - Connection security used by the email server.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).

