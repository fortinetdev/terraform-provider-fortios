---
subcategory: "FortiGate VoIP"
layout: "fortios"
page_title: "FortiOS: fortios_voip_profile"
description: |-
  Configure VoIP profiles.
---

# fortios_voip_profile
Configure VoIP profiles.

## Example Usage

```hcl
resource "fortios_voip_profile" "trname" {
  comment = "test"
  name    = "1"

  sccp {
    block_mcast      = "disable"
    log_call_summary = "disable"
    log_violations   = "disable"
    max_calls        = 0
    status           = "enable"
    verify_header    = "disable"
  }

  sip {
    ack_rate                       = 0
    bye_rate                       = 0
    call_keepalive                 = 0
    cancel_rate                    = 0
    contact_fixup                  = "enable"
    hnt_restrict_source_ip         = "disable"
    hosted_nat_traversal           = "disable"
    info_rate                      = 0
    invite_rate                    = 0
    ips_rtp                        = "enable"
    log_call_summary               = "enable"
    log_violations                 = "disable"
    max_body_length                = 0
    max_dialogs                    = 0
    max_idle_dialogs               = 0
    max_line_length                = 998
    message_rate                   = 0
    nat_trace                      = "enable"
    no_sdp_fixup                   = "disable"
    notify_rate                    = 0
    open_contact_pinhole           = "enable"
    open_record_route_pinhole      = "enable"
    open_register_pinhole          = "enable"
    open_via_pinhole               = "disable"
    options_rate                   = 0
    prack_rate                     = 0
    preserve_override              = "disable"
    provisional_invite_expiry_time = 210
    publish_rate                   = 0
    refer_rate                     = 0
    register_contact_trace         = "disable"
    register_rate                  = 0
    rfc2543_branch                 = "disable"
    rtp                            = "enable"
    ssl_algorithm                  = "high"
    ssl_client_renegotiation       = "allow"
    ssl_max_version                = "tls-1.2"
    ssl_min_version                = "tls-1.1"
    ssl_mode                       = "off"
    ssl_pfs                        = "allow"
    ssl_send_empty_frags           = "enable"
    status                         = "enable"
    strict_register                = "enable"
    subscribe_rate                 = 0
    unknown_header                 = "pass"
    update_rate                    = 0
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Profile name.
* `comment` - Comment.
* `sip` - SIP. The structure of `sip` block is documented below.
* `sccp` - SCCP. The structure of `sccp` block is documented below.

The `sip` block supports:

* `status` - Enable/disable SIP.
* `rtp` - Enable/disable create pinholes for RTP traffic to traverse firewall.
* `nat_port_range` - RTP NAT port range.
* `open_register_pinhole` - Enable/disable open pinhole for REGISTER Contact port.
* `open_contact_pinhole` - Enable/disable open pinhole for non-REGISTER Contact port.
* `strict_register` - Enable/disable only allow the registrar to connect.
* `register_rate` - REGISTER request rate limit (per second, per policy).
* `invite_rate` - INVITE request rate limit (per second, per policy).
* `max_dialogs` - Maximum number of concurrent calls/dialogs (per policy).
* `max_line_length` - Maximum SIP header line length (78-4096).
* `block_long_lines` - Enable/disable block requests with headers exceeding max-line-length.
* `block_unknown` - Block unrecognized SIP requests (enabled by default).
* `call_keepalive` - Continue tracking calls with no RTP for this many minutes.
* `block_ack` - Enable/disable block ACK requests.
* `block_bye` - Enable/disable block BYE requests.
* `block_cancel` - Enable/disable block CANCEL requests.
* `block_info` - Enable/disable block INFO requests.
* `block_invite` - Enable/disable block INVITE requests.
* `block_message` - Enable/disable block MESSAGE requests.
* `block_notify` - Enable/disable block NOTIFY requests.
* `block_options` - Enable/disable block OPTIONS requests and no OPTIONS as notifying message for redundancy either.
* `block_prack` - Enable/disable block prack requests.
* `block_publish` - Enable/disable block PUBLISH requests.
* `block_refer` - Enable/disable block REFER requests.
* `block_register` - Enable/disable block REGISTER requests.
* `block_subscribe` - Enable/disable block SUBSCRIBE requests.
* `block_update` - Enable/disable block UPDATE requests.
* `register_contact_trace` - Enable/disable trace original IP/port within the contact header of REGISTER requests.
* `open_via_pinhole` - Enable/disable open pinhole for Via port.
* `open_record_route_pinhole` - Enable/disable open pinhole for Record-Route port.
* `rfc2543_branch` - Enable/disable support via branch compliant with RFC 2543.
* `log_violations` - Enable/disable logging of SIP violations.
* `log_call_summary` - Enable/disable logging of SIP call summary.
* `nat_trace` - Enable/disable preservation of original IP in SDP i line.
* `subscribe_rate` - SUBSCRIBE request rate limit (per second, per policy).
* `message_rate` - MESSAGE request rate limit (per second, per policy).
* `notify_rate` - NOTIFY request rate limit (per second, per policy).
* `refer_rate` - REFER request rate limit (per second, per policy).
* `update_rate` - UPDATE request rate limit (per second, per policy).
* `options_rate` - OPTIONS request rate limit (per second, per policy).
* `ack_rate` - ACK request rate limit (per second, per policy).
* `prack_rate` - PRACK request rate limit (per second, per policy).
* `info_rate` - INFO request rate limit (per second, per policy).
* `publish_rate` - PUBLISH request rate limit (per second, per policy).
* `bye_rate` - BYE request rate limit (per second, per policy).
* `cancel_rate` - CANCEL request rate limit (per second, per policy).
* `preserve_override` - Override i line to preserve original IPS (default: append).
* `no_sdp_fixup` - Enable/disable no SDP fix-up.
* `contact_fixup` - Fixup contact anyway even if contact's IP:port doesn't match session's IP:port.
* `max_idle_dialogs` - Maximum number established but idle dialogs to retain (per policy).
* `block_geo_red_options` - Enable/disable block OPTIONS requests, but OPTIONS requests still notify for redundancy.
* `hosted_nat_traversal` - Hosted NAT Traversal (HNT).
* `hnt_restrict_source_ip` - Enable/disable restrict RTP source IP to be the same as SIP source IP when HNT is enabled.
* `max_body_length` - Maximum SIP message body length (0 meaning no limit).
* `unknown_header` - Action for unknown SIP header.
* `malformed_request_line` - Action for malformed request line.
* `malformed_header_via` - Action for malformed VIA header.
* `malformed_header_from` - Action for malformed From header.
* `malformed_header_to` - Action for malformed To header.
* `malformed_header_call_id` - Action for malformed Call-ID header.
* `malformed_header_cseq` - Action for malformed CSeq header.
* `malformed_header_rack` - Action for malformed RAck header.
* `malformed_header_rseq` - Action for malformed RSeq header.
* `malformed_header_contact` - Action for malformed Contact header.
* `malformed_header_record_route` - Action for malformed Record-Route header.
* `malformed_header_route` - Action for malformed Route header.
* `malformed_header_expires` - Action for malformed Expires header.
* `malformed_header_content_type` - Action for malformed Content-Type header.
* `malformed_header_content_length` - Action for malformed Content-Length header.
* `malformed_header_max_forwards` - Action for malformed Max-Forwards header.
* `malformed_header_allow` - Action for malformed Allow header.
* `malformed_header_p_asserted_identity` - Action for malformed P-Asserted-Identity header.
* `malformed_header_sdp_v` - Action for malformed SDP v line.
* `malformed_header_sdp_o` - Action for malformed SDP o line.
* `malformed_header_sdp_s` - Action for malformed SDP s line.
* `malformed_header_sdp_i` - Action for malformed SDP i line.
* `malformed_header_sdp_c` - Action for malformed SDP c line.
* `malformed_header_sdp_b` - Action for malformed SDP b line.
* `malformed_header_sdp_z` - Action for malformed SDP z line.
* `malformed_header_sdp_k` - Action for malformed SDP k line.
* `malformed_header_sdp_a` - Action for malformed SDP a line.
* `malformed_header_sdp_t` - Action for malformed SDP t line.
* `malformed_header_sdp_r` - Action for malformed SDP r line.
* `malformed_header_sdp_m` - Action for malformed SDP m line.
* `provisional_invite_expiry_time` - Expiry time for provisional INVITE (10 - 3600 sec).
* `ips_rtp` - Enable/disable allow IPS on RTP.
* `ssl_mode` - SSL/TLS mode for encryption & decryption of traffic.
* `ssl_send_empty_frags` - Send empty fragments to avoid attack on CBC IV (SSL 3.0 & TLS 1.0 only).
* `ssl_client_renegotiation` - Allow/block client renegotiation by server.
* `ssl_algorithm` - Relative strength of encryption algorithms accepted in negotiation.
* `ssl_pfs` - SSL Perfect Forward Secrecy.
* `ssl_min_version` - Lowest SSL/TLS version to negotiate.
* `ssl_max_version` - Highest SSL/TLS version to negotiate.
* `ssl_client_certificate` - Name of Certificate to offer to server if requested.
* `ssl_server_certificate` - Name of Certificate return to the client in every SSL connection.
* `ssl_auth_client` - Require a client certificate and authenticate it with the peer/peergrp.
* `ssl_auth_server` - Authenticate the server's certificate with the peer/peergrp.

The `sccp` block supports:

* `status` - Enable/disable SCCP.
* `block_mcast` - Enable/disable block multicast RTP connections.
* `verify_header` - Enable/disable verify SCCP header content.
* `log_call_summary` - Enable/disable log summary of SCCP calls.
* `log_violations` - Enable/disable logging of SCCP violations.
* `max_calls` - Maximum calls per minute per SCCP client (max 65535).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Voip Profile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_voip_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
