---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_npu"
description: |-
  Configure NPU attributes.
---

# fortios_system_npu
Configure NPU attributes. Applies to FortiOS Version `7.0.4,7.2.13,8.0.0`.

## Argument Reference

The following arguments are supported:

* `dedicated_management_cpu` - Enable to dedicate one CPU for GUI and CLI connections when NPs are busy. Valid values: `enable`, `disable`.
* `dedicated_lacp_queue` - Enable/disable dedication of HIF queue #0 for LACP. Valid values: `enable`, `disable`.
* `dedicated_management_affinity` - Affinity setting for management daemons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
* `dos_options` - NPU DoS configurations. The structure of `dos_options` block is documented below.
* `policy_offload_level` - Configure firewall policy offload level. Valid values: `disable`, `dos-offload`.
* `napi_break_interval` -  NAPI break interval (default 0).
* `hpe` - Host protection engine configuration. The structure of `hpe` block is documented below.
* `fastpath` - Enable/disable NP6 offloading (also called fast path). Valid values: `disable`, `enable`.
* `capwap_offload` - Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions. Valid values: `enable`, `disable`.
* `vxlan_offload` - Enable/disable offloading vxlan. Valid values: `enable`, `disable`.
* `vxlan_mac_flapping_guard` - Enable/disable VxLAN MAC flapping guard. Valid values: `enable`, `disable`.
* `default_qos_type` - Set default QoS type. Valid values: `policing`, `shaping`, `policing-enhanced`.
* `default_ipsec_mcs_type` - Configure default IPSec MCS type. Valid values: `policing`, `shaping`.
* `shaping_stats` - Enable/disable NP7 traffic shaping statistics (default = disable). Valid values: `disable`, `enable`.
* `mcs_host_packet_tpe_shaping` - Enable/disable NPU shaping for host traffic with shaping profile on IPSec interface. Valid values: `enable`, `disable`.
* `gtp_support` - Enable/Disable NP7 GTP support Valid values: `enable`, `disable`.
* `per_session_accounting` - Set per-session accounting. Valid values: `traffic-log-only`, `disable`, `enable`.
* `session_acct_interval` - Session accounting update interval (default 5 sec). On FortiOS versions >= 8.0.0: 1 - 600 sec. On FortiOS versions 7.2.13: 1 - 10 sec.
* `max_session_timeout` - Maximum time interval for refreshing NPU-offloaded sessions (10 - 1000 sec, default 40 sec).
* `fp_anomaly` - IPv4/IPv6 anomaly protection. The structure of `fp_anomaly` block is documented below.
* `ip_reassembly` - IP reassebmly engine configuration. The structure of `ip_reassembly` block is documented below.
* `hash_tbl_spread` - Enable/disable hash table entry spread (default enabled). Valid values: `enable`, `disable`.
* `vlan_lookup_cache` - Enable/disable VLAN lookup cache. On FortiOS versions >= 8.0.0: default = disabled. On FortiOS versions 7.2.13: default enabled. Valid values: `enable`, `disable`.
* `ip_fragment_offload` - Enable/disable NP7 NPU IP fragment offload. Valid values: `disable`, `enable`.
* `htx_icmp_csum_chk` - Set HTX icmp csum checking mode. Valid values: `drop`, `pass`.
* `htab_msg_queue` - Set hash table message queue mode. Valid values: `data`, `idle`, `dedicated`.
* `htab_dedi_queue_nr` - Set the number of dedicate queue for hash table messages.
* `dsw_dts_profile` - Configure NPU DSW DTS profile. The structure of `dsw_dts_profile` block is documented below.
* `dsw_queue_dts_profile` - Configure NPU DSW Queue DTS profile. The structure of `dsw_queue_dts_profile` block is documented below.
* `npu_tcam` - Configure NPU TCAM policies. The structure of `npu_tcam` block is documented below.
* `np_queues` - Configure queue assignment on NP7. The structure of `np_queues` block is documented below.
* `sw_np_rate` - Bandwidth from switch to NP.
* `sw_np_rate_unit` - Unit for bandwidth from switch to NP. Valid values: `mbps`, `pps`.
* `sw_np_rate_burst` - Burst value for bandwidth from switch to NP.
* `double_level_mcast_offload` - Enable/disable double level mcast offload. Valid values: `enable`, `disable`.
* `mcast_denied_ses_offload` - Enable/disable offloading of multicast denied sessions. Valid values: `enable`, `disable`.
* `ipsec_enc_subengine_mask` - IPsec encryption subengine mask (0x1 - 0xff, default 0xff).
* `ipsec_dec_subengine_mask` - IPsec decryption subengine mask (0x1 - 0xff, default 0xff).
* `np6_cps_optimization_mode` - Enable/disable NP6 connection per second (CPS) optimization mode. Valid values: `enable`, `disable`.
* `sw_np_bandwidth` - Bandwidth from switch to NP. Valid values: `0G`, `2G`, `4G`, `5G`, `6G`.
* `strip_esp_padding` - Enable/disable stripping ESP padding. Valid values: `enable`, `disable`.
* `strip_clear_text_padding` - Enable/disable stripping clear text padding. Valid values: `enable`, `disable`.
* `ipsec_inbound_cache` - Enable/disable IPsec inbound cache for anti-replay. Valid values: `enable`, `disable`.
* `sse_backpressure` - Enable/disable sse backpressure. Valid values: `enable`, `disable`.
* `rdp_offload` - Enable/disable rdp offload. Valid values: `enable`, `disable`.
* `ipsec_over_vlink` - Enable/disable IPSEC over vlink. Valid values: `enable`, `disable`.
* `uesp_offload` - Enable/disable UDP-encapsulated ESP offload (default = disable). Valid values: `enable`, `disable`.
* `mcast_session_accounting` - Enable/disable traffic accounting for each multicast session through TAE counter. Valid values: `tpe-based`, `session-based`, `disable`.
* `ipsec_mtu_override` - Enable/disable NP6 IPsec MTU override. Valid values: `disable`, `enable`.
* `session_denied_offload` - Enable/disable offloading of denied sessions. Requires ses-denied-traffic to be set. Valid values: `disable`, `enable`.
* `qtm_buf_mode` - QTM channel configuration for packet buffer. Valid values: `6ch`, `4ch`.
* `ipsec_ob_np_sel` - IPsec NP selection for OB SA offloading. Valid values: `rr`, `Packet`, `Hash`.
* `max_receive_unit` - Set the maximum packet size for receive, larger packets will be silently dropped.
* `lag_hash_gre` - Set LAG hash for standard GRE. Valid values: `disable`, `gre_inner_l3`, `gre_inner_l4`, `gre_inner_l3l4`.
* `use_mse_oft` - Enable/disable use of MSE OFT. Valid values: `enable`, `disable`.
* `ike_port` - Configure additional IPsec ports for offloading. The structure of `ike_port` block is documented below.
* `priority_protocol` - Configure NPU priority protocol. The structure of `priority_protocol` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `dos_options` block supports:

* `npu_dos_meter_mode` - Set DoS meter NPU offloading mode. Valid values: `global`, `local`.
* `npu_dos_tpe_mode` - Enable/disable insertion of DoS meter ID to session table. Valid values: `enable`, `disable`.

The `hpe` block supports:

* `all_protocol` - Maximum packet rate of each host queue except high priority traffic(1K - 32M pps, default = 400K pps), set 0 to disable.
* `tcpsyn_max` - Maximum TCP SYN packet rate (1K - 40M pps, default = 32K pps).
* `tcpsyn_ack_max` - Maximum TCP carries SYN and ACK flags packet rate (1K - 32M pps, default = 40K pps).
* `tcpfin_rst_max` - Maximum TCP carries FIN or RST flags packet rate (1K - 32M pps, default = 40K pps).
* `tcp_max` - Maximum TCP packet rate (1K - 32M pps, default = 40K pps).
* `udp_max` - Maximum UDP packet rate (1K - 32M pps, default = 40K pps).
* `icmp_max` - Maximum ICMP packet rate (1K - 32M pps, default = 5K pps).
* `sctp_max` - Maximum SCTP packet rate (1K - 32M pps, default = 5K pps).
* `esp_max` - Maximum ESP packet rate (1K - 32M pps, default = 5K pps).
* `ip_frag_max` - Maximum fragmented IP packet rate (1K - 32M pps, default = 5K pps).
* `ip_others_max` - Maximum IP packet rate for other packets (packet types that cannot be set with other options) (1K - 32G pps, default = 5K pps).
* `arp_max` - Maximum ARP packet rate (1K - 32M pps, default = 5K pps). Entry is valid when ARP is removed from high-priority traffic.
* `l2_others_max` - Maximum L2 packet rate for L2 packets that are not ARP packets (1K - 32M pps, default = 5K pps).
* `high_priority` - Maximum packet rate for high priority traffic packets (1K - 32M pps, default = 400K pps).
* `enable_shaper` - Enable/Disable NPU Host Protection Engine (HPE) for packet type shaper. Valid values: `disable`, `enable`.

The `fp_anomaly` block supports:

* `tcp_syn_fin` - TCP SYN flood SYN/FIN flag set anomalies.  Valid values: `allow`, `drop`, `trap-to-host`.
* `tcp_fin_noack` - TCP SYN flood with FIN flag set without ACK setting anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `tcp_fin_only` - TCP SYN flood with only FIN flag set anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `tcp_no_flag` - TCP SYN flood with no flag set anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `tcp_syn_data` - TCP SYN flood packets with data anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `tcp_winnuke` - TCP WinNuke anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `tcp_land` - TCP land anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `udp_land` - UDP land anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `icmp_land` - ICMP land anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `icmp_frag` - Layer 3 fragmented packets that could be part of layer 4 ICMP anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv4_land` - Land anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv4_proto_err` - Invalid layer 4 protocol anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv4_unknopt` - Unknown option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv4_optrr` - Record route option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv4_optssrr` - Strict source record route option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv4_optlsrr` - Loose source record route option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv4_optstream` - Stream option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv4_optsecurity` - Security option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv4_opttimestamp` - Timestamp option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv4_csum_err` - Invalid IPv4 IP checksum anomalies. Valid values: `drop`, `trap-to-host`.
* `tcp_csum_err` - Invalid IPv4 TCP checksum anomalies.
* `udp_csum_err` - Invalid IPv4 UDP checksum anomalies.
* `udplite_csum_err` - Invalid IPv4 UDP-Lite checksum anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `icmp_csum_err` - Invalid IPv4 ICMP checksum anomalies.
* `gre_csum_err` - Invalid IPv4 GRE checksum anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `sctp_csum_err` - Invalid IPv4 SCTP checksum anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv6_land` - Land anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv6_proto_err` - Layer 4 invalid protocol anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv6_unknopt` - Unknown option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv6_saddr_err` - Source address as multicast anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv6_daddr_err` - Destination address as unspecified or loopback address anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv6_optralert` - Router alert option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv6_optjumbo` - Jumbo options anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv6_opttunnel` - Tunnel encapsulation limit option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv6_opthomeaddr` - Home address option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv6_optnsap` - Network service access point address option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv6_optendpid` - End point identification anomalies. Valid values: `allow`, `drop`, `trap-to-host`.
* `ipv6_optinvld` - Invalid option anomalies.Invalid option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

The `ip_reassembly` block supports:

* `min_timeout` - Minimum timeout value for IP reassembly (5 us - 600,000,000 us).
* `max_timeout` - Maximum timeout value for IP reassembly (5 us - 600,000,000 us).
* `status` - Set IP reassembly processing status. Valid values: `disable`, `enable`.

The `dsw_dts_profile` block supports:

* `profile_id` - Set NPU DSW DTS profile profile id.
* `min_limit` - Set NPU DSW DTS profile min-limt.
* `step` - Set NPU DSW DTS profile step.
* `action` - Set NPU DSW DTS profile action. Valid values: `wait`, `drop`, `drop_tmr_0`, `drop_tmr_1`, `enque`, `enque_0`, `enque_1`.

The `dsw_queue_dts_profile` block supports:

* `name` - Name.
* `iport` - Set NPU DSW DTS in port. Valid values: `eif0`, `eif1`, `eif2`, `eif3`, `eif4`, `eif5`, `eif6`, `eif7`, `htx0`, `htx1`, `sse0`, `sse1`, `sse2`, `sse3`, `rlt`, `dfr`, `ipseci`, `ipseco`, `ipti`, `ipto`, `vep0`, `vep2`, `vep4`, `vep6`, `ivs`, `l2ti1`, `l2to`, `l2ti0`, `ple`, `spath`, `qtm`.
* `oport` - Set NPU DSW DTS out port. Valid values: `eif0`, `eif1`, `eif2`, `eif3`, `eif4`, `eif5`, `eif6`, `eif7`, `hrx`, `sse0`, `sse1`, `sse2`, `sse3`, `rlt`, `dfr`, `ipseci`, `ipseco`, `ipti`, `ipto`, `vep0`, `vep2`, `vep4`, `vep6`, `ivs`, `l2ti1`, `l2to`, `l2ti0`, `ple`, `sync`, `nss`, `tsk`, `qtm`.
* `profile_id` - Set NPU DSW DTS profile ID.
* `queue_select` - Set NPU DSW DTS queue ID select (0 - reset to default).

The `npu_tcam` block supports:

* `name` - NPU TCAM policies name.
* `type` - TCAM policy type. Valid values: `L2_src_tc`, `L2_tgt_tc`, `L2_src_mir`, `L2_tgt_mir`, `L2_src_act`, `L2_tgt_act`, `IPv4_src_tc`, `IPv4_tgt_tc`, `IPv4_src_mir`, `IPv4_tgt_mir`, `IPv4_src_act`, `IPv4_tgt_act`, `IPv6_src_tc`, `IPv6_tgt_tc`, `IPv6_src_mir`, `IPv6_tgt_mir`, `IPv6_src_act`, `IPv6_tgt_act`.
* `oid` - NPU TCAM OID.
* `vid` - NPU TCAM VID.
* `data` - Data fields of TCAM. The structure of `data` block is documented below.
* `mask` - Mask fields of TCAM. The structure of `mask` block is documented below.
* `mir_act` - Mirror action of TCAM. The structure of `mir_act` block is documented below.
* `pri_act` - Priority action of TCAM. The structure of `pri_act` block is documented below.
* `sact` - Source action of TCAM. The structure of `sact` block is documented below.
* `tact` - Target action of TCAM. The structure of `tact` block is documented below.

The `data` block supports:

* `gen_buf_cnt` - tcam data gen info buffer count.
* `gen_pri` - tcam data gen info priority.
* `gen_pri_v` - tcam data gen info priority valid. Valid values: `valid`, `invalid`.
* `gen_iv` - tcam data gen info iv. Valid values: `valid`, `invalid`.
* `gen_tv` - tcam data gen info tv. Valid values: `valid`, `invalid`.
* `gen_pkt_ctrl` - tcam data gen info packet control.
* `gen_l3_flags` - tcam data gen info L3 flags.
* `gen_l4_flags` - tcam data gen info L4 flags.
* `vdid` - tcam data vdom id.
* `tp` - tcam data target port.
* `tgt_updt` - tcam data target port update. Valid values: `enable`, `disable`.
* `smac_change` - tcam data source MAC change. Valid values: `enable`, `disable`.
* `ext_tag` - tcam data extension tag. Valid values: `enable`, `disable`.
* `tgt_v` - tcam data target valid. Valid values: `valid`, `invalid`.
* `tvid` - tcam data target vid.
* `tgt_cfi` - tcam data target cfi. Valid values: `enable`, `disable`.
* `tgt_prio` - tcam data target priority.
* `sp` - tcam data source port.
* `src_updt` - tcam data source update. Valid values: `enable`, `disable`.
* `slink` - tcam data sublink.
* `svid` - tcam data source vid.
* `src_cfi` - tcam data source cfi. Valid values: `enable`, `disable`.
* `src_prio` - tcam data source priority.
* `srcmac` - tcam data src macaddr.
* `dstmac` - tcam data dst macaddr.
* `ethertype` - tcam data ethertype.
* `ipver` - tcam data ip header version.
* `ihl` - tcam data ipv4 IHL.
* `ip4_id` - tcam data ipv4 id.
* `srcip` - tcam data src ipv4 address.
* `dstip` - tcam data dst ipv4 address.
* `ip6_fl` - tcam data ipv6 flow label.
* `srcipv6` - tcam data src ipv6 address.
* `dstipv6` - tcam data dst ipv6 address.
* `ttl` - tcam data ip ttl.
* `protocol` - tcam data ip protocol.
* `tos` - tcam data ip tos.
* `frag_off` - tcam data ip flag fragment offset.
* `mf` - tcam data ip flag mf. Valid values: `enable`, `disable`.
* `df` - tcam data ip flag df. Valid values: `enable`, `disable`.
* `srcport` - tcam data L4 src port.
* `dstport` - tcam data L4 dst port.
* `tcp_fin` - tcam data tcp flag fin. Valid values: `enable`, `disable`.
* `tcp_syn` - tcam data tcp flag syn. Valid values: `enable`, `disable`.
* `tcp_rst` - tcam data tcp flag rst. Valid values: `enable`, `disable`.
* `tcp_push` - tcam data tcp flag push. Valid values: `enable`, `disable`.
* `tcp_ack` - tcam data tcp flag ack. Valid values: `enable`, `disable`.
* `tcp_urg` - tcam data tcp flag urg. Valid values: `enable`, `disable`.
* `tcp_ece` - tcam data tcp flag ece. Valid values: `enable`, `disable`.
* `tcp_cwr` - tcam data tcp flag cwr. Valid values: `enable`, `disable`.
* `l4_wd8` - tcam data L4 word8.
* `l4_wd9` - tcam data L4 word9.
* `l4_wd10` - tcam data L4 word10.
* `l4_wd11` - tcam data L4 word11.

The `mask` block supports:

* `gen_buf_cnt` - tcam mask gen info buffer count.
* `gen_pri` - tcam mask gen info priority.
* `gen_pri_v` - tcam mask gen info priority valid. Valid values: `valid`, `invalid`.
* `gen_iv` - tcam mask gen info iv. Valid values: `valid`, `invalid`.
* `gen_tv` - tcam mask gen info tv. Valid values: `valid`, `invalid`.
* `gen_pkt_ctrl` - tcam mask gen info packet control.
* `gen_l3_flags` - tcam mask gen info L3 flags.
* `gen_l4_flags` - tcam mask gen info L4 flags.
* `vdid` - tcam mask vdom id.
* `tp` - tcam mask target port.
* `tgt_updt` - tcam mask target port update. Valid values: `enable`, `disable`.
* `smac_change` - tcam mask source MAC change. Valid values: `enable`, `disable`.
* `ext_tag` - tcam mask extension tag. Valid values: `enable`, `disable`.
* `tgt_v` - tcam mask target valid. Valid values: `valid`, `invalid`.
* `tvid` - tcam mask target vid.
* `tgt_cfi` - tcam mask target cfi. Valid values: `enable`, `disable`.
* `tgt_prio` - tcam mask target priority.
* `sp` - tcam mask source port.
* `src_updt` - tcam mask source update. Valid values: `enable`, `disable`.
* `slink` - tcam mask sublink.
* `svid` - tcam mask source vid.
* `src_cfi` - tcam mask source cfi. Valid values: `enable`, `disable`.
* `src_prio` - tcam mask source priority.
* `srcmac` - tcam mask src macaddr.
* `dstmac` - tcam mask dst macaddr.
* `ethertype` - tcam mask ethertype.
* `ipver` - tcam mask ip header version.
* `ihl` - tcam mask ipv4 IHL.
* `ip4_id` - tcam mask ipv4 id.
* `srcip` - tcam mask src ipv4 address.
* `dstip` - tcam mask dst ipv4 address.
* `ip6_fl` - tcam mask ipv6 flow label.
* `srcipv6` - tcam mask src ipv6 address.
* `dstipv6` - tcam mask dst ipv6 address.
* `ttl` - tcam mask ip ttl.
* `protocol` - tcam mask ip protocol.
* `tos` - tcam mask ip tos.
* `frag_off` - tcam data ip flag fragment offset.
* `mf` - tcam mask ip flag mf. Valid values: `enable`, `disable`.
* `df` - tcam mask ip flag df. Valid values: `enable`, `disable`.
* `srcport` - tcam mask L4 src port.
* `dstport` - tcam mask L4 dst port.
* `tcp_fin` - tcam mask tcp flag fin. Valid values: `enable`, `disable`.
* `tcp_syn` - tcam mask tcp flag syn. Valid values: `enable`, `disable`.
* `tcp_rst` - tcam mask tcp flag rst. Valid values: `enable`, `disable`.
* `tcp_push` - tcam mask tcp flag push. Valid values: `enable`, `disable`.
* `tcp_ack` - tcam mask tcp flag ack. Valid values: `enable`, `disable`.
* `tcp_urg` - tcam mask tcp flag urg. Valid values: `enable`, `disable`.
* `tcp_ece` - tcam mask tcp flag ece. Valid values: `enable`, `disable`.
* `tcp_cwr` - tcam mask tcp flag cwr. Valid values: `enable`, `disable`.
* `l4_wd8` - tcam mask L4 word8.
* `l4_wd9` - tcam mask L4 word9.
* `l4_wd10` - tcam mask L4 word10.
* `l4_wd11` - tcam mask L4 word11.

The `mir_act` block supports:

* `vlif` - tcam mirror action vlif.

The `pri_act` block supports:

* `priority` - tcam priority action priority.
* `weight` - tcam priority action weight.

The `sact` block supports:

* `fwd_lif_v` - Enable to set sact fwd-lif. Valid values: `enable`, `disable`.
* `fwd_lif` - tcam sact fwd-lif.
* `fwd_tvid_v` - Enable to set sact fwd-vid. Valid values: `enable`, `disable`.
* `fwd_tvid` - tcam sact fwd-tvid.
* `df_lif_v` - Enable to set sact df-lif. Valid values: `enable`, `disable`.
* `df_lif` - tcam sact df-lif.
* `act_v` - Enable to set sact act. Valid values: `enable`, `disable`.
* `act` - tcam sact act.
* `pleen_v` - Enable to set sact pleen. Valid values: `enable`, `disable`.
* `pleen` - tcam sact pleen.
* `icpen_v` - Enable to set sact icpen. Valid values: `enable`, `disable`.
* `icpen` - tcam sact icpen.
* `vdm_v` - Enable to set sact vdm. Valid values: `enable`, `disable`.
* `vdm` - tcam sact vdm.
* `learn_v` - Enable to set sact learn. Valid values: `enable`, `disable`.
* `learn` - tcam sact learn.
* `rfsh_v` - Enable to set sact rfsh. Valid values: `enable`, `disable`.
* `rfsh` - tcam sact rfsh.
* `fwd_v` - Enable to set sact fwd. Valid values: `enable`, `disable`.
* `fwd` - tcam sact fwd.
* `x_mode_v` - Enable to set sact x-mode. Valid values: `enable`, `disable`.
* `x_mode` - tcam sact x-mode.
* `promis_v` - Enable to set sact promis. Valid values: `enable`, `disable`.
* `promis` - tcam sact promis.
* `bmproc_v` - Enable to set sact bmproc. Valid values: `enable`, `disable`.
* `bmproc` - tcam sact bmproc.
* `mac_id_v` - Enable to set sact mac-id. Valid values: `enable`, `disable`.
* `mac_id` - tcam sact mac-id.
* `dosen_v` - Enable to set sact dosen. Valid values: `enable`, `disable`.
* `dosen` - tcam sact dosen.
* `dfr_v` - Enable to set sact dfr. Valid values: `enable`, `disable`.
* `dfr` - tcam sact dfr.
* `m_srh_ctrl_v` - Enable to set sact m-srh-ctrl. Valid values: `enable`, `disable`.
* `m_srh_ctrl` - tcam sact m-srh-ctrl.
* `tpe_id_v` - Enable to set sact tpe-id. Valid values: `enable`, `disable`.
* `tpe_id` - tcam sact tpe-id.
* `vdom_id_v` - Enable to set sact vdom-id. Valid values: `enable`, `disable`.
* `vdom_id` - tcam sact vdom-id.
* `mss_v` - Enable to set sact mss. Valid values: `enable`, `disable`.
* `mss` - tcam sact mss.
* `tp_smchk_v` - Enable to set sact tp mode. Valid values: `enable`, `disable`.
* `tp_smchk` - tcam sact tp mode.
* `etype_pid_v` - Enable to set sact etype-pid. Valid values: `enable`, `disable`.
* `etype_pid` - tcam sact etype-pid.
* `frag_proc_v` - Enable to set sact frag-proc. Valid values: `enable`, `disable`.
* `frag_proc` - tcam sact frag-proc.
* `espff_proc_v` - Enable to set sact espff-proc. Valid values: `enable`, `disable`.
* `espff_proc` - tcam sact espff-proc.
* `prio_pid_v` - Enable to set sact prio-pid. Valid values: `enable`, `disable`.
* `prio_pid` - tcam sact prio-pid.
* `igmp_mld_snp_v` - Enable to set sact igmp-mld-snp. Valid values: `enable`, `disable`.
* `igmp_mld_snp` - tcam sact igmp-mld-snp.
* `smac_skip_v` - Enable to set sact smac-skip. Valid values: `enable`, `disable`.
* `smac_skip` - tcam sact smac-skip.
* `dmac_skip_v` - Enable to set sact dmac-skip. Valid values: `enable`, `disable`.
* `dmac_skip` - tcam sact dmac-skip.

The `tact` block supports:

* `act_v` - Enable to set tact act. Valid values: `enable`, `disable`.
* `act` - tcam tact act.
* `mtuv4_v` - Enable to set tact mtuv4. Valid values: `enable`, `disable`.
* `mtuv4` - tcam tact mtuv4.
* `mtuv6_v` - Enable to set tact mtuv6. Valid values: `enable`, `disable`.
* `mtuv6` - tcam tact mtuv6.
* `mac_id_v` - Enable to set tact mac-id. Valid values: `enable`, `disable`.
* `mac_id` - tcam tact mac-id.
* `slif_act_v` - Enable to set tact slif-act. Valid values: `enable`, `disable`.
* `slif_act` - tcam tact slif-act.
* `tlif_act_v` - Enable to set tact tlif-act. Valid values: `enable`, `disable`.
* `tlif_act` - tcam tact tlif-act.
* `tgtv_act_v` - Enable to set tact tgtv-act. Valid values: `enable`, `disable`.
* `tgtv_act` - tcam tact tgtv-act.
* `tpeid_v` - Enable to set tact tpeid. Valid values: `enable`, `disable`.
* `tpeid` - tcam tact tpeid.
* `v6fe_v` - Enable to set tact v6fe. Valid values: `enable`, `disable`.
* `v6fe` - tcam tact v6fe.
* `xlt_vid_v` - Enable to set tact xlt-vid. Valid values: `enable`, `disable`.
* `xlt_vid` - tcam tact xlt-vid.
* `xlt_lif_v` - Enable to set tact xlt-lif. Valid values: `enable`, `disable`.
* `xlt_lif` - tcam tact xlt-lif.
* `mss_t_v` - Enable to set tact mss. Valid values: `enable`, `disable`.
* `mss_t` - tcam tact mss.
* `lnkid_v` - Enable to set tact lnkid. Valid values: `enable`, `disable`.
* `lnkid` - tcam tact lnkid.
* `sublnkid_v` - Enable to set tact sublnkid. Valid values: `enable`, `disable`.
* `sublnkid` - tcam tact sublnkid.
* `fmtuv4_s_v` - Enable to set tact fmtuv4-s. Valid values: `enable`, `disable`.
* `fmtuv4_s` - tcam tact fmtuv4-s.
* `fmtuv6_s_v` - Enable to set tact fmtuv6-s. Valid values: `enable`, `disable`.
* `fmtuv6_s` - tcam tact fmtuv6-s.
* `vep_en_v` - Enable to set tact vep-en. Valid values: `enable`, `disable`.
* `vep_en` - tcam tact vep_en.
* `vep_slid_v` - Enable to set tact vep-slid. Valid values: `enable`, `disable`.
* `vep_slid` - tcam tact vep_slid.

The `np_queues` block supports:

* `profile` - Configure a NP7 class profile. The structure of `profile` block is documented below.
* `ethernet_type` - Configure a NP7 QoS Ethernet Type. The structure of `ethernet_type` block is documented below.
* `ip_protocol` - Configure a NP7 QoS IP Protocol. The structure of `ip_protocol` block is documented below.
* `ip_service` - Configure a NP7 QoS IP Service. The structure of `ip_service` block is documented below.
* `scheduler` - Configure a NP7 QoS Scheduler. The structure of `scheduler` block is documented below.
* `custom_etype_lookup` - Enable/Disable np-queue lookup for custom Ethernet Types. Valid values: `enable`, `disable`.

The `profile` block supports:

* `id` - Profile ID.
* `type` - Profile type. Valid values: `cos`, `dscp`.
* `weight` - Class weight.
* `cos0` - Queue number of CoS 0. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `cos1` - Queue number of CoS 1. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `cos2` - Queue number of CoS 2. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `cos3` - Queue number of CoS 3. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `cos4` - Queue number of CoS 4. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `cos5` - Queue number of CoS 5. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `cos6` - Queue number of CoS 6. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `cos7` - Queue number of CoS 7. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp0` - Queue number of DSCP 0. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp1` - Queue number of DSCP 1. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp2` - Queue number of DSCP 2. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp3` - Queue number of DSCP 3. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp4` - Queue number of DSCP 4. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp5` - Queue number of DSCP 5. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp6` - Queue number of DSCP 6. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp7` - Queue number of DSCP 7. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp8` - Queue number of DSCP 8. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp9` - Queue number of DSCP 9. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp10` - Queue number of DSCP 10. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp11` - Queue number of DSCP 11. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp12` - Queue number of DSCP 12. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp13` - Queue number of DSCP 13. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp14` - Queue number of DSCP 14. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp15` - Queue number of DSCP 15. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp16` - Queue number of DSCP 16. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp17` - Queue number of DSCP 17. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp18` - Queue number of DSCP 18. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp19` - Queue number of DSCP 19. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp20` - Queue number of DSCP 20. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp21` - Queue number of DSCP 21. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp22` - Queue number of DSCP 22. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp23` - Queue number of DSCP 23. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp24` - Queue number of DSCP 24. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp25` - Queue number of DSCP 25. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp26` - Queue number of DSCP 26. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp27` - Queue number of DSCP 27. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp28` - Queue number of DSCP 28. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp29` - Queue number of DSCP 29. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp30` - Queue number of DSCP 30. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp31` - Queue number of DSCP 31. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp32` - Queue number of DSCP 32. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp33` - Queue number of DSCP 33. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp34` - Queue number of DSCP 34. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp35` - Queue number of DSCP 35. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp36` - Queue number of DSCP 36. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp37` - Queue number of DSCP 37. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp38` - Queue number of DSCP 38. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp39` - Queue number of DSCP 39. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp40` - Queue number of DSCP 40. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp41` - Queue number of DSCP 41. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp42` - Queue number of DSCP 42. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp43` - Queue number of DSCP 43. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp44` - Queue number of DSCP 44. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp45` - Queue number of DSCP 45. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp46` - Queue number of DSCP 46. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp47` - Queue number of DSCP 47. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp48` - Queue number of DSCP 48. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp49` - Queue number of DSCP 49. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp50` - Queue number of DSCP 50. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp51` - Queue number of DSCP 51. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp52` - Queue number of DSCP 52. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp53` - Queue number of DSCP 53. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp54` - Queue number of DSCP 54. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp55` - Queue number of DSCP 55. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp56` - Queue number of DSCP 56. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp57` - Queue number of DSCP 57. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp58` - Queue number of DSCP 58. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp59` - Queue number of DSCP 59. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp60` - Queue number of DSCP 60. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp61` - Queue number of DSCP 61. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp62` - Queue number of DSCP 62. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.
* `dscp63` - Queue number of DSCP 63. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

The `ethernet_type` block supports:

* `name` - Ethernet Type Name.
* `type` - Ethernet Type.
* `queue` - Queue Number.
* `weight` - Class Weight.

The `ip_protocol` block supports:

* `name` - IP Protocol Name.
* `protocol` - IP Protocol.
* `queue` - Queue Number.
* `weight` - Class Weight.

The `ip_service` block supports:

* `name` - IP service name.
* `protocol` - IP protocol.
* `sport` - Source port.
* `dport` - Destination port.
* `queue` - Queue number.
* `weight` - Class weight.

The `scheduler` block supports:

* `name` - Scheduler name.
* `mode` - Scheduler mode. Valid values: `none`, `priority`, `round-robin`.

The `ike_port` block supports:

* `port` - Port.

The `priority_protocol` block supports:

* `bgp` - Enable/disable NPU BGP priority protocol. Valid values: `enable`, `disable`.
* `slbc` - Enable/disable NPU SLBC priority protocol. Valid values: `enable`, `disable`.
* `bfd` - Enable/disable NPU BFD priority protocol. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Npu can be imported using any of these accepted formats:
```
$ terraform import fortios_system_npu.labelname SystemNpu

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_npu.labelname SystemNpu
$ unset "FORTIOS_IMPORT_TABLE"
```
