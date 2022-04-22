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
* `feature_set` - Flow or proxy inspection feature set. Valid values: `flow`, `proxy`.
* `comment` - Comment.
* `sip` - SIP. The structure of `sip` block is documented below.
* `sccp` - SCCP. The structure of `sccp` block is documented below.
* `msrp` - MSRP. The structure of `msrp` block is documented below.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `sip` block supports:

* `status` - Enable/disable SIP. Valid values: `disable`, `enable`.
* `rtp` - Enable/disable create pinholes for RTP traffic to traverse firewall. Valid values: `disable`, `enable`.
* `nat_port_range` - RTP NAT port range.
* `open_register_pinhole` - Enable/disable open pinhole for REGISTER Contact port. Valid values: `disable`, `enable`.
* `open_contact_pinhole` - Enable/disable open pinhole for non-REGISTER Contact port. Valid values: `disable`, `enable`.
* `strict_register` - Enable/disable only allow the registrar to connect. Valid values: `disable`, `enable`.
* `register_rate` - REGISTER request rate limit (per second, per policy).
* `register_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.
* `invite_rate` - INVITE request rate limit (per second, per policy).
* `invite_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.
* `max_dialogs` - Maximum number of concurrent calls/dialogs (per policy).
* `max_line_length` - Maximum SIP header line length (78-4096).
* `block_long_lines` - Enable/disable block requests with headers exceeding max-line-length. Valid values: `disable`, `enable`.
* `block_unknown` - Block unrecognized SIP requests (enabled by default). Valid values: `disable`, `enable`.
* `call_keepalive` - Continue tracking calls with no RTP for this many minutes.
* `block_ack` - Enable/disable block ACK requests. Valid values: `disable`, `enable`.
* `block_bye` - Enable/disable block BYE requests. Valid values: `disable`, `enable`.
* `block_cancel` - Enable/disable block CANCEL requests. Valid values: `disable`, `enable`.
* `block_info` - Enable/disable block INFO requests. Valid values: `disable`, `enable`.
* `block_invite` - Enable/disable block INVITE requests. Valid values: `disable`, `enable`.
* `block_message` - Enable/disable block MESSAGE requests. Valid values: `disable`, `enable`.
* `block_notify` - Enable/disable block NOTIFY requests. Valid values: `disable`, `enable`.
* `block_options` - Enable/disable block OPTIONS requests and no OPTIONS as notifying message for redundancy either. Valid values: `disable`, `enable`.
* `block_prack` - Enable/disable block prack requests. Valid values: `disable`, `enable`.
* `block_publish` - Enable/disable block PUBLISH requests. Valid values: `disable`, `enable`.
* `block_refer` - Enable/disable block REFER requests. Valid values: `disable`, `enable`.
* `block_register` - Enable/disable block REGISTER requests. Valid values: `disable`, `enable`.
* `block_subscribe` - Enable/disable block SUBSCRIBE requests. Valid values: `disable`, `enable`.
* `block_update` - Enable/disable block UPDATE requests. Valid values: `disable`, `enable`.
* `register_contact_trace` - Enable/disable trace original IP/port within the contact header of REGISTER requests. Valid values: `disable`, `enable`.
* `open_via_pinhole` - Enable/disable open pinhole for Via port. Valid values: `disable`, `enable`.
* `open_record_route_pinhole` - Enable/disable open pinhole for Record-Route port. Valid values: `disable`, `enable`.
* `rfc2543_branch` - Enable/disable support via branch compliant with RFC 2543. Valid values: `disable`, `enable`.
* `log_violations` - Enable/disable logging of SIP violations. Valid values: `disable`, `enable`.
* `log_call_summary` - Enable/disable logging of SIP call summary. Valid values: `disable`, `enable`.
* `nat_trace` - Enable/disable preservation of original IP in SDP i line. Valid values: `disable`, `enable`.
* `subscribe_rate` - SUBSCRIBE request rate limit (per second, per policy).
* `subscribe_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.
* `message_rate` - MESSAGE request rate limit (per second, per policy).
* `message_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.
* `notify_rate` - NOTIFY request rate limit (per second, per policy).
* `notify_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.
* `refer_rate` - REFER request rate limit (per second, per policy).
* `refer_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.
* `update_rate` - UPDATE request rate limit (per second, per policy).
* `update_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.
* `options_rate` - OPTIONS request rate limit (per second, per policy).
* `options_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.
* `ack_rate` - ACK request rate limit (per second, per policy).
* `ack_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.
* `prack_rate` - PRACK request rate limit (per second, per policy).
* `prack_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.
* `info_rate` - INFO request rate limit (per second, per policy).
* `info_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.
* `publish_rate` - PUBLISH request rate limit (per second, per policy).
* `publish_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.
* `bye_rate` - BYE request rate limit (per second, per policy).
* `bye_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.
* `cancel_rate` - CANCEL request rate limit (per second, per policy).
* `cancel_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.
* `preserve_override` - Override i line to preserve original IPS (default: append). Valid values: `disable`, `enable`.
* `no_sdp_fixup` - Enable/disable no SDP fix-up. Valid values: `disable`, `enable`.
* `contact_fixup` - Fixup contact anyway even if contact's IP:port doesn't match session's IP:port. Valid values: `disable`, `enable`.
* `max_idle_dialogs` - Maximum number established but idle dialogs to retain (per policy).
* `block_geo_red_options` - Enable/disable block OPTIONS requests, but OPTIONS requests still notify for redundancy. Valid values: `disable`, `enable`.
* `hosted_nat_traversal` - Hosted NAT Traversal (HNT). Valid values: `disable`, `enable`.
* `hnt_restrict_source_ip` - Enable/disable restrict RTP source IP to be the same as SIP source IP when HNT is enabled. Valid values: `disable`, `enable`.
* `max_body_length` - Maximum SIP message body length (0 meaning no limit).
* `unknown_header` - Action for unknown SIP header. Valid values: `discard`, `pass`, `respond`.
* `malformed_request_line` - Action for malformed request line. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_via` - Action for malformed VIA header. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_from` - Action for malformed From header. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_to` - Action for malformed To header. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_call_id` - Action for malformed Call-ID header. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_cseq` - Action for malformed CSeq header. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_rack` - Action for malformed RAck header. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_rseq` - Action for malformed RSeq header. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_contact` - Action for malformed Contact header. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_record_route` - Action for malformed Record-Route header. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_route` - Action for malformed Route header. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_expires` - Action for malformed Expires header. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_content_type` - Action for malformed Content-Type header. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_content_length` - Action for malformed Content-Length header. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_max_forwards` - Action for malformed Max-Forwards header. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_allow` - Action for malformed Allow header. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_p_asserted_identity` - Action for malformed P-Asserted-Identity header. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_no_require` - Action for malformed SIP messages without Require header. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_no_proxy_require` - Action for malformed SIP messages without Proxy-Require header. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_sdp_v` - Action for malformed SDP v line. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_sdp_o` - Action for malformed SDP o line. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_sdp_s` - Action for malformed SDP s line. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_sdp_i` - Action for malformed SDP i line. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_sdp_c` - Action for malformed SDP c line. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_sdp_b` - Action for malformed SDP b line. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_sdp_z` - Action for malformed SDP z line. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_sdp_k` - Action for malformed SDP k line. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_sdp_a` - Action for malformed SDP a line. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_sdp_t` - Action for malformed SDP t line. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_sdp_r` - Action for malformed SDP r line. Valid values: `discard`, `pass`, `respond`.
* `malformed_header_sdp_m` - Action for malformed SDP m line. Valid values: `discard`, `pass`, `respond`.
* `provisional_invite_expiry_time` - Expiry time for provisional INVITE (10 - 3600 sec).
* `ips_rtp` - Enable/disable allow IPS on RTP. Valid values: `disable`, `enable`.
* `ssl_mode` - SSL/TLS mode for encryption & decryption of traffic. Valid values: `off`, `full`.
* `ssl_send_empty_frags` - Send empty fragments to avoid attack on CBC IV (SSL 3.0 & TLS 1.0 only). Valid values: `enable`, `disable`.
* `ssl_client_renegotiation` - Allow/block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
* `ssl_algorithm` - Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
* `ssl_pfs` - SSL Perfect Forward Secrecy. Valid values: `require`, `deny`, `allow`.
* `ssl_min_version` - Lowest SSL/TLS version to negotiate.
* `ssl_max_version` - Highest SSL/TLS version to negotiate.
* `ssl_client_certificate` - Name of Certificate to offer to server if requested.
* `ssl_server_certificate` - Name of Certificate return to the client in every SSL connection.
* `ssl_auth_client` - Require a client certificate and authenticate it with the peer/peergrp.
* `ssl_auth_server` - Authenticate the server's certificate with the peer/peergrp.

The `sccp` block supports:

* `status` - Enable/disable SCCP. Valid values: `disable`, `enable`.
* `block_mcast` - Enable/disable block multicast RTP connections. Valid values: `disable`, `enable`.
* `verify_header` - Enable/disable verify SCCP header content. Valid values: `disable`, `enable`.
* `log_call_summary` - Enable/disable log summary of SCCP calls. Valid values: `disable`, `enable`.
* `log_violations` - Enable/disable logging of SCCP violations. Valid values: `disable`, `enable`.
* `max_calls` - Maximum calls per minute per SCCP client (max 65535).

The `msrp` block supports:

* `status` - Enable/disable MSRP. Valid values: `disable`, `enable`.
* `log_violations` - Enable/disable logging of MSRP violations. Valid values: `disable`, `enable`.
* `max_msg_size` - Maximum allowable MSRP message size (1-65535).
* `max_msg_size_action` - Action for violation of max-msg-size. Valid values: `pass`, `block`, `reset`, `monitor`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Voip Profile can be imported using any of these accepted formats:
```
$ terraform import fortios_voip_profile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_voip_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
