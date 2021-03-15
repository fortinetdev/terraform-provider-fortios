---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_profileprotocoloptions"
description: |-
  Configure protocol options.
---

# fortios_firewall_profileprotocoloptions
Configure protocol options.

## Example Usage

```hcl
resource "fortios_firewall_profileprotocoloptions" "trname" {
  name                    = "sp"
  oversize_log            = "disable"
  rpc_over_http           = "disable"
  switching_protocols_log = "disable"

  dns {
    ports  = 53
    status = "enable"
  }

  ftp {
    comfort_amount              = 1
    comfort_interval            = 10
    inspect_all                 = "disable"
    options                     = "splice"
    oversize_limit              = 10
    ports                       = 21
    scan_bzip2                  = "enable"
    status                      = "enable"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }

  http {
    block_page_status_code      = 403
    comfort_amount              = 1
    comfort_interval            = 10
    fortinet_bar                = "disable"
    fortinet_bar_port           = 8011
    http_policy                 = "disable"
    inspect_all                 = "disable"
    oversize_limit              = 10
    ports                       = 80
    range_block                 = "disable"
    retry_count                 = 0
    scan_bzip2                  = "enable"
    status                      = "enable"
    streaming_content_bypass    = "enable"
    strip_x_forwarded_for       = "disable"
    switching_protocols         = "bypass"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }

  imap {
    inspect_all                 = "disable"
    options                     = "fragmail"
    oversize_limit              = 10
    ports                       = 143
    scan_bzip2                  = "enable"
    status                      = "enable"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }

  mail_signature {
    status = "disable"
  }

  mapi {
    options                     = "fragmail"
    oversize_limit              = 10
    ports                       = 135
    scan_bzip2                  = "enable"
    status                      = "enable"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }

  nntp {
    inspect_all                 = "disable"
    options                     = "splice"
    oversize_limit              = 10
    ports                       = 119
    scan_bzip2                  = "enable"
    status                      = "enable"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }

  pop3 {
    inspect_all                 = "disable"
    options                     = "fragmail"
    oversize_limit              = 10
    ports                       = 110
    scan_bzip2                  = "enable"
    status                      = "enable"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }

  smtp {
    inspect_all                 = "disable"
    options                     = "fragmail splice"
    oversize_limit              = 10
    ports                       = 25
    scan_bzip2                  = "enable"
    server_busy                 = "disable"
    status                      = "enable"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `comment` - Optional comments.
* `feature_set` - Flow/proxy feature set. Valid values: `flow`, `proxy`.
* `replacemsg_group` - Name of the replacement message group to be used
* `oversize_log` - Enable/disable logging for antivirus oversize file blocking. Valid values: `disable`, `enable`.
* `switching_protocols_log` - Enable/disable logging for HTTP/HTTPS switching protocols. Valid values: `disable`, `enable`.
* `http` - Configure HTTP protocol options. The structure of `http` block is documented below.
* `ftp` - Configure FTP protocol options. The structure of `ftp` block is documented below.
* `imap` - Configure IMAP protocol options. The structure of `imap` block is documented below.
* `mapi` - Configure MAPI protocol options. The structure of `mapi` block is documented below.
* `pop3` - Configure POP3 protocol options. The structure of `pop3` block is documented below.
* `smtp` - Configure SMTP protocol options. The structure of `smtp` block is documented below.
* `nntp` - Configure NNTP protocol options. The structure of `nntp` block is documented below.
* `ssh` - Configure SFTP and SCP protocol options. The structure of `ssh` block is documented below.
* `dns` - Configure DNS protocol options. The structure of `dns` block is documented below.
* `cifs` - Configure CIFS protocol options. The structure of `cifs` block is documented below.
* `mail_signature` - Configure Mail signature. The structure of `mail_signature` block is documented below.
* `rpc_over_http` - Enable/disable inspection of RPC over HTTP. Valid values: `enable`, `disable`.

The `http` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 80).
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `enable`, `disable`.
* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `enable`, `disable`.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable`, `disable`.
* `options` - One or more options that can be applied to the session. Valid values: `clientcomfort`, `servercomfort`, `oversize`, `chunkedbypass`.
* `comfort_interval` - Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).
* `comfort_amount` - Amount of data to send in a transmission for client comforting (1 - 10240 bytes, default = 1).
* `range_block` - Enable/disable blocking of partial downloads. Valid values: `disable`, `enable`.
* `http_policy` - Enable/disable HTTP policy check. Valid values: `disable`, `enable`.
* `strip_x_forwarded_for` - Enable/disable stripping of HTTP X-Forwarded-For header. Valid values: `disable`, `enable`.
* `post_lang` - ID codes for character sets to be used to convert to UTF-8 for banned words and DLP on HTTP posts (maximum of 5 character sets). Valid values: `jisx0201`, `jisx0208`, `jisx0212`, `gb2312`, `ksc5601-ex`, `euc-jp`, `sjis`, `iso2022-jp`, `iso2022-jp-1`, `iso2022-jp-2`, `euc-cn`, `ces-gbk`, `hz`, `ces-big5`, `euc-kr`, `iso2022-jp-3`, `iso8859-1`, `tis620`, `cp874`, `cp1252`, `cp1251`.
* `fortinet_bar` - Enable/disable Fortinet bar on HTML content. Valid values: `enable`, `disable`.
* `fortinet_bar_port` - Port for use by Fortinet Bar (1 - 65535, default = 8011).
* `streaming_content_bypass` - Enable/disable bypassing of streaming content from buffering. Valid values: `enable`, `disable`.
* `switching_protocols` - Bypass from scanning, or block a connection that attempts to switch protocol. Valid values: `bypass`, `block`.
* `unknown_http_version` - How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1. Valid values: `reject`, `tunnel`, `best-effort`.
* `tunnel_non_http` - Configure how to process non-HTTP traffic when a profile configured for HTTP traffic accepts a non-HTTP session. Can occur if an application sends non-HTTP traffic using an HTTP destination port. Valid values: `enable`, `disable`.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned (MB, 0 = unlimited (default).  Stream-based uncompression used only under certain conditions.).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `enable`, `disable`.
* `block_page_status_code` - Code number returned for blocked HTTP pages (non-FortiGuard only) (100 - 599, default = 403).
* `retry_count` - Number of attempts to retry HTTP connection (0 - 100, default = 0).
* `tcp_window_type` - Specify type of TCP window to use for this protocol. Valid values: `system`, `static`, `dynamic`.
* `tcp_window_minimum` - Minimum dynamic TCP window size (default = 128KB).
* `tcp_window_maximum` - Maximum dynamic TCP window size (default = 8MB).
* `tcp_window_size` - Set TCP static window size (default = 256KB).
* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.

The `ftp` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 21).
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `enable`, `disable`.
* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `enable`, `disable`.
* `options` - One or more options that can be applied to the session. Valid values: `clientcomfort`, `oversize`, `splice`, `bypass-rest-command`, `bypass-mode-command`.
* `comfort_interval` - Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).
* `comfort_amount` - Amount of data to send in a transmission for client comforting (1 - 10240 bytes, default = 1).
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned (MB, 0 = unlimited (default).  Stream-based uncompression used only under certain conditions.).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `enable`, `disable`.
* `tcp_window_type` - TCP window type to use for this protocol. Valid values: `system`, `static`, `dynamic`.
* `tcp_window_minimum` - Minimum dynamic TCP window size.
* `tcp_window_maximum` - Maximum dynamic TCP window size.
* `tcp_window_size` - Set TCP static window size.
* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.

The `imap` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 143).
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `enable`, `disable`.
* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `enable`, `disable`.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable`, `disable`.
* `options` - One or more options that can be applied to the session. Valid values: `fragmail`, `oversize`.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `enable`, `disable`.
* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.

The `mapi` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 135).
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `enable`, `disable`.
* `options` - One or more options that can be applied to the session. Valid values: `fragmail`, `oversize`.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `enable`, `disable`.

The `pop3` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 110).
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `enable`, `disable`.
* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `enable`, `disable`.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable`, `disable`.
* `options` - One or more options that can be applied to the session. Valid values: `fragmail`, `oversize`.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `enable`, `disable`.
* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.

The `smtp` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 25).
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `enable`, `disable`.
* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `enable`, `disable`.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable`, `disable`.
* `options` - One or more options that can be applied to the session. Valid values: `fragmail`, `oversize`, `splice`.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `enable`, `disable`.
* `server_busy` - Enable/disable SMTP server busy when server not available. Valid values: `enable`, `disable`.
* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.

The `nntp` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 119).
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `enable`, `disable`.
* `inspect_all` - Enable/disable the inspection of all ports for the protocol. Valid values: `enable`, `disable`.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable`, `disable`.
* `options` - One or more options that can be applied to the session. Valid values: `oversize`, `splice`.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `enable`, `disable`.

The `ssh` block supports:

* `options` - One or more options that can be applied to the session. Valid values: `oversize`, `clientcomfort`, `servercomfort`.
* `comfort_interval` - Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).
* `comfort_amount` - Amount of data to send in a transmission for client comforting (1 - 65535 bytes, default = 1).
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned (MB, 0 = unlimited (default).  Stream-based uncompression used only under certain conditions.).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `enable`, `disable`.
* `tcp_window_type` - TCP window type to use for this protocol. Valid values: `system`, `static`, `dynamic`.
* `tcp_window_minimum` - Minimum dynamic TCP window size.
* `tcp_window_maximum` - Maximum dynamic TCP window size.
* `tcp_window_size` - Set TCP static window size.
* `ssl_offloaded` - SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.

The `dns` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 53).
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `enable`, `disable`.

The `cifs` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 445).
* `status` - Enable/disable the active status of scanning for this protocol. Valid values: `enable`, `disable`.
* `options` - One or more options that can be applied to the session. Valid values: `oversize`.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. Valid values: `enable`, `disable`.
* `tcp_window_type` - Specify type of TCP window to use for this protocol. Valid values: `system`, `static`, `dynamic`.
* `tcp_window_minimum` - Minimum dynamic TCP window size (default = 128KB).
* `tcp_window_maximum` - Maximum dynamic TCP window size (default = 8MB).
* `tcp_window_size` - Set TCP static window size (default = 256KB).
* `server_credential_type` - CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
* `domain_controller` - Domain for which to decrypt CIFS traffic.
* `server_keytab` - Server keytab. The structure of `server_keytab` block is documented below.

The `server_keytab` block supports:

* `principal` - Service principal.  For example, "host/cifsserver.example.com@example.com".
* `keytab` - Base64 encoded keytab file containing credential of the server.

The `mail_signature` block supports:

* `status` - Enable/disable adding an email signature to SMTP email messages as they pass through the FortiGate. Valid values: `disable`, `enable`.
* `signature` - Email signature to be added to outgoing email (if the signature contains spaces, enclose with quotation marks).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall ProfileProtocolOptions can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_profileprotocoloptions.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
