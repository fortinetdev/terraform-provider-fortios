---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_profileprotocoloptions"
description: |-
  Get information on an fortios firewall profileprotocoloptions.
---

# Data Source: fortios_firewall_profileprotocoloptions
Use this data source to get information on an fortios firewall profileprotocoloptions

## Argument Reference

* `name` - (Required) Specify the name of the desired firewall profileprotocoloptions.

## Attribute Reference

The following attributes are exported:

* `name` - Name.
* `comment` - Optional comments.
* `feature_set` - Flow/proxy feature set.
* `replacemsg_group` - Name of the replacement message group to be used
* `oversize_log` - Enable/disable logging for antivirus oversize file blocking.
* `switching_protocols_log` - Enable/disable logging for HTTP/HTTPS switching protocols.
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
* `rpc_over_http` - Enable/disable inspection of RPC over HTTP.

The `http` block contains:

* `ports` - Ports to scan for content (1 - 65535, default = 80).
* `status` - Enable/disable the active status of scanning for this protocol.
* `inspect_all` - Enable/disable the inspection of all ports for the protocol.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before).
* `options` - One or more options that can be applied to the session.
* `comfort_interval` - Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).
* `comfort_amount` - Amount of data to send in a transmission for client comforting (1 - 10240 bytes, default = 1).
* `range_block` - Enable/disable blocking of partial downloads.
* `http_policy` - Enable/disable HTTP policy check.
* `strip_x_forwarded_for` - Enable/disable stripping of HTTP X-Forwarded-For header.
* `post_lang` - ID codes for character sets to be used to convert to UTF-8 for banned words and DLP on HTTP posts (maximum of 5 character sets).
* `fortinet_bar` - Enable/disable Fortinet bar on HTML content.
* `fortinet_bar_port` - Port for use by Fortinet Bar (1 - 65535, default = 8011).
* `streaming_content_bypass` - Enable/disable bypassing of streaming content from buffering.
* `switching_protocols` - Bypass from scanning, or block a connection that attempts to switch protocol.
* `unknown_http_version` - How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1.
* `tunnel_non_http` - Configure how to process non-HTTP traffic when a profile configured for HTTP traffic accepts a non-HTTP session. Can occur if an application sends non-HTTP traffic using an HTTP destination port.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned (MB, 0 = unlimited (default).  Stream-based uncompression used only under certain conditions.).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `block_page_status_code` - Code number returned for blocked HTTP pages (non-FortiGuard only) (100 - 599, default = 403).
* `retry_count` - Number of attempts to retry HTTP connection (0 - 100, default = 0).
* `tcp_window_type` - Specify type of TCP window to use for this protocol.
* `tcp_window_minimum` - Minimum dynamic TCP window size (default = 128KB).
* `tcp_window_maximum` - Maximum dynamic TCP window size (default = 8MB).
* `tcp_window_size` - Set TCP static window size (default = 256KB).
* `ssl_offloaded` - SSL decryption and encryption performed by an external device.

The `ftp` block contains:

* `ports` - Ports to scan for content (1 - 65535, default = 21).
* `status` - Enable/disable the active status of scanning for this protocol.
* `inspect_all` - Enable/disable the inspection of all ports for the protocol.
* `options` - One or more options that can be applied to the session.
* `comfort_interval` - Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).
* `comfort_amount` - Amount of data to send in a transmission for client comforting (1 - 10240 bytes, default = 1).
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned (MB, 0 = unlimited (default).  Stream-based uncompression used only under certain conditions.).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `tcp_window_type` - TCP window type to use for this protocol.
* `tcp_window_minimum` - Minimum dynamic TCP window size.
* `tcp_window_maximum` - Maximum dynamic TCP window size.
* `tcp_window_size` - Set TCP static window size.
* `ssl_offloaded` - SSL decryption and encryption performed by an external device.

The `imap` block contains:

* `ports` - Ports to scan for content (1 - 65535, default = 143).
* `status` - Enable/disable the active status of scanning for this protocol.
* `inspect_all` - Enable/disable the inspection of all ports for the protocol.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before).
* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `ssl_offloaded` - SSL decryption and encryption performed by an external device.

The `mapi` block contains:

* `ports` - Ports to scan for content (1 - 65535, default = 135).
* `status` - Enable/disable the active status of scanning for this protocol.
* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.

The `pop3` block contains:

* `ports` - Ports to scan for content (1 - 65535, default = 110).
* `status` - Enable/disable the active status of scanning for this protocol.
* `inspect_all` - Enable/disable the inspection of all ports for the protocol.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before).
* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `ssl_offloaded` - SSL decryption and encryption performed by an external device.

The `smtp` block contains:

* `ports` - Ports to scan for content (1 - 65535, default = 25).
* `status` - Enable/disable the active status of scanning for this protocol.
* `inspect_all` - Enable/disable the inspection of all ports for the protocol.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before).
* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `server_busy` - Enable/disable SMTP server busy when server not available.
* `ssl_offloaded` - SSL decryption and encryption performed by an external device.

The `nntp` block contains:

* `ports` - Ports to scan for content (1 - 65535, default = 119).
* `status` - Enable/disable the active status of scanning for this protocol.
* `inspect_all` - Enable/disable the inspection of all ports for the protocol.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before).
* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.

The `ssh` block contains:

* `options` - One or more options that can be applied to the session.
* `comfort_interval` - Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).
* `comfort_amount` - Amount of data to send in a transmission for client comforting (1 - 65535 bytes, default = 1).
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned (MB, 0 = unlimited (default).  Stream-based uncompression used only under certain conditions.).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `tcp_window_type` - TCP window type to use for this protocol.
* `tcp_window_minimum` - Minimum dynamic TCP window size.
* `tcp_window_maximum` - Maximum dynamic TCP window size.
* `tcp_window_size` - Set TCP static window size.
* `ssl_offloaded` - SSL decryption and encryption performed by an external device.

The `dns` block contains:

* `ports` - Ports to scan for content (1 - 65535, default = 53).
* `status` - Enable/disable the active status of scanning for this protocol.

The `cifs` block contains:

* `ports` - Ports to scan for content (1 - 65535, default = 445).
* `status` - Enable/disable the active status of scanning for this protocol.
* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `tcp_window_type` - Specify type of TCP window to use for this protocol.
* `tcp_window_minimum` - Minimum dynamic TCP window size (default = 128KB).
* `tcp_window_maximum` - Maximum dynamic TCP window size (default = 8MB).
* `tcp_window_size` - Set TCP static window size (default = 256KB).
* `server_credential_type` - CIFS server credential type.
* `domain_controller` - Domain for which to decrypt CIFS traffic.
* `server_keytab` - Server keytab. The structure of `server_keytab` block is documented below.

The `server_keytab` block contains:

* `principal` - Service principal.  For example, "host/cifsserver.example.com@example.com".
* `keytab` - Base64 encoded keytab file containing credential of the server.

The `mail_signature` block contains:

* `status` - Enable/disable adding an email signature to SMTP email messages as they pass through the FortiGate.
* `signature` - Email signature to be added to outgoing email (if the signature contains spaces, enclose with quotation marks).

