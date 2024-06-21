---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_global"
description: |-
  Configure global attributes.
---

# fortios_system_global
Configure global attributes.

## Example Usage

```hcl
resource "fortios_system_global" "trname" {
  admin_sport = 443
  alias       = "FGVM02TM20003062"
  hostname    = "ste11"
  timezone    = "04"
}
```

## Argument Reference

The following arguments are supported:

* `language` - GUI display language. Valid values: `english`, `french`, `spanish`, `portuguese`, `japanese`, `trach`, `simch`, `korean`.
* `gui_allow_incompatible_fabric_fgt` - Enable/disable Allow FGT with incompatible firmware to be treated as compatible in security fabric on the GUI. May cause unexpected error. Valid values: `enable`, `disable`.
* `gui_ipv6` - Enable/disable IPv6 settings on the GUI. Valid values: `enable`, `disable`.
* `gui_replacement_message_groups` - Enable/disable replacement message groups on the GUI. Valid values: `enable`, `disable`.
* `gui_local_out` - Enable/disable Local-out traffic on the GUI. Valid values: `enable`, `disable`.
* `gui_certificates` - Enable/disable the System > Certificate GUI page, allowing you to add and configure certificates from the GUI. Valid values: `enable`, `disable`.
* `gui_custom_language` - Enable/disable custom languages in GUI. Valid values: `enable`, `disable`.
* `gui_wireless_opensecurity` - Enable/disable wireless open security option on the GUI. Valid values: `enable`, `disable`.
* `gui_app_detection_sdwan` - Enable/disable Allow app-detection based SD-WAN. Valid values: `enable`, `disable`.
* `gui_display_hostname` - Enable/disable displaying the FortiGate's hostname on the GUI login page. Valid values: `enable`, `disable`.
* `gui_fortigate_cloud_sandbox` - Enable/disable displaying FortiGate Cloud Sandbox on the GUI. Valid values: `enable`, `disable`.
* `gui_fortisandbox_cloud` - Enable/disable displaying FortiSandbox Cloud on the GUI. Valid values: `enable`, `disable`.
* `gui_firmware_upgrade_warning` - Enable/disable the firmware upgrade warning on the GUI. Valid values: `enable`, `disable`.
* `gui_firmware_upgrade_setup_warning` - Enable/disable the firmware upgrade warning on GUI setup wizard. Valid values: `enable`, `disable`.
* `gui_lines_per_page` - Number of lines to display per page for web administration.
* `admin_https_ssl_versions` - Allowed TLS versions for web administration.
* `admin_https_ssl_ciphersuites` - Select one or more TLS 1.3 ciphersuites to enable. Does not affect ciphers in TLS 1.2 and below. At least one must be enabled. To disable all, remove TLS1.3 from admin-https-ssl-versions. Valid values: `TLS-AES-128-GCM-SHA256`, `TLS-AES-256-GCM-SHA384`, `TLS-CHACHA20-POLY1305-SHA256`, `TLS-AES-128-CCM-SHA256`, `TLS-AES-128-CCM-8-SHA256`.
* `admin_https_ssl_banned_ciphers` - Select one or more cipher technologies that cannot be used in GUI HTTPS negotiations. Only applies to TLS 1.2 and below. Valid values: `RSA`, `DHE`, `ECDHE`, `DSS`, `ECDSA`, `AES`, `AESGCM`, `CAMELLIA`, `3DES`, `SHA1`, `SHA256`, `SHA384`, `STATIC`, `CHACHA20`, `ARIA`, `AESCCM`.
* `admintimeout` - Number of minutes before an idle administrator session times out (default = 5). A shorter idle timeout is more secure. On FortiOS versions 6.2.0-6.2.6: 5 - 480 minutes (8 hours). On FortiOS versions >= 6.4.0: 1 - 480 minutes (8 hours).
* `admin_console_timeout` - Console login timeout that overrides the admin timeout value (15 - 300 seconds, default = 0, which disables the timeout).
* `ssd_trim_freq` - How often to run SSD Trim (default = weekly). SSD Trim prevents SSD drive data loss by finding and isolating errors. Valid values: `never`, `hourly`, `daily`, `weekly`, `monthly`.
* `ssd_trim_hour` - Hour of the day on which to run SSD Trim (0 - 23, default = 1).
* `ssd_trim_min` - Minute of the hour on which to run SSD Trim (0 - 59, 60 for random).
* `ssd_trim_weekday` - Day of week to run SSD Trim. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
* `ssd_trim_date` - Date within a month to run ssd trim.
* `admin_concurrent` - Enable/disable concurrent administrator logins. Use policy-auth-concurrent for firewall authenticated users. Valid values: `enable`, `disable`.
* `admin_lockout_threshold` - Number of failed login attempts before an administrator account is locked out for the admin-lockout-duration.
* `admin_lockout_duration` - Amount of time in seconds that an administrator account is locked out after reaching the admin-lockout-threshold for repeated failed login attempts.
* `refresh` - Statistics refresh interval second(s) in GUI.
* `interval` - Dead gateway detection interval.
* `failtime` - Fail-time for server lost.
* `purdue_level` - Purdue Level of this FortiGate. Valid values: `1`, `1.5`, `2`, `2.5`, `3`, `3.5`, `4`, `5`, `5.5`.
* `daily_restart` - Enable/disable daily restart of FortiGate unit. Use the restart-time option to set the time of day for the restart. Valid values: `enable`, `disable`.
* `restart_time` - Daily restart time (hh:mm).
* `wad_restart_mode` - WAD worker restart mode (default = none). Valid values: `none`, `time`, `memory`.
* `wad_restart_start_time` - WAD workers daily restart time (hh:mm).
* `wad_restart_end_time` - WAD workers daily restart end time (hh:mm).
* `radius_port` - RADIUS service port number.
* `speedtestd_server_port` - Speedtest server port number.
* `speedtestd_ctrl_port` - Speedtest server controller port number.
* `admin_login_max` - Maximum number of administrators who can be logged in at the same time (1 - 100, default = 100)
* `remoteauthtimeout` - Number of seconds that the FortiGate waits for responses from remote RADIUS, LDAP, or TACACS+ authentication servers. (default = 5). On FortiOS versions 6.2.0-6.2.6: 0-300 sec, 0 means no timeout. On FortiOS versions >= 6.4.0: 1-300 sec.
* `ldapconntimeout` - Global timeout for connections with remote LDAP servers in milliseconds (1 - 300000, default 500).
* `batch_cmdb` - Enable/disable batch mode, allowing you to enter a series of CLI commands that will execute as a group once they are loaded. Valid values: `enable`, `disable`.
* `max_dlpstat_memory` - Maximum DLP stat memory (0 - 4294967295).
* `multi_factor_authentication` - Enforce all login methods to require an additional authentication factor (default = optional). Valid values: `optional`, `mandatory`.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default = TLSv1.2).
* `autorun_log_fsck` - Enable/disable automatic log partition check after ungraceful shutdown. Valid values: `enable`, `disable`.
* `dst` - Enable/disable daylight saving time. Valid values: `enable`, `disable`.
* `timezone` - Number corresponding to your time zone from 00 to 86. Enter set timezone ? to view the list of time zones and the numbers that represent them.
* `traffic_priority` - Choose Type of Service (ToS) or Differentiated Services Code Point (DSCP) for traffic prioritization in traffic shaping. Valid values: `tos`, `dscp`.
* `traffic_priority_level` - Default system-wide level of priority for traffic prioritization. Valid values: `low`, `medium`, `high`.
* `quic_congestion_control_algo` - QUIC congestion control algorithm (default = cubic). Valid values: `cubic`, `bbr`, `bbr2`, `reno`.
* `quic_max_datagram_size` - Maximum transmit datagram size (1200 - 1500, default = 1500).
* `quic_udp_payload_size_shaping_per_cid` - Enable/disable UDP payload size shaping per connection ID (default = enable). Valid values: `enable`, `disable`.
* `quic_ack_thresold` - Maximum number of unacknowledged packets before sending ACK (2 - 5, default = 3).
* `quic_pmtud` - Enable/disable path MTU discovery (default = enable). Valid values: `enable`, `disable`.
* `quic_tls_handshake_timeout` - Time-to-live (TTL) for TLS handshake in seconds (1 - 60, default = 5).
* `anti_replay` - Level of checking for packet replay and TCP sequence checking. Valid values: `disable`, `loose`, `strict`.
* `send_pmtu_icmp` - Enable/disable sending of path maximum transmission unit (PMTU) - ICMP destination unreachable packet and to support PMTUD protocol on your network to reduce fragmentation of packets. Valid values: `enable`, `disable`.
* `honor_df` - Enable/disable honoring of Don't-Fragment (DF) flag. Valid values: `enable`, `disable`.
* `pmtu_discovery` - Enable/disable path MTU discovery. Valid values: `enable`, `disable`.
* `virtual_switch_vlan` - Enable/disable virtual switch VLAN. Valid values: `enable`, `disable`.
* `split_port` - Split port(s) to multiple 10Gbps ports.
* `revision_image_auto_backup` - Enable/disable back-up of the latest configuration revision after the firmware is upgraded. Valid values: `enable`, `disable`.
* `revision_backup_on_logout` - Enable/disable back-up of the latest configuration revision when an administrator logs out of the CLI or GUI. Valid values: `enable`, `disable`.
* `management_vdom` - Management virtual domain name.
* `hostname` - FortiGate unit's hostname. Most models will truncate names longer than 24 characters. Some models support hostnames up to 35 characters.
* `gui_allow_default_hostname` - Enable/disable the GUI warning about using a default hostname Valid values: `enable`, `disable`.
* `gui_forticare_registration_setup_warning` - Enable/disable the FortiCare registration setup warning on the GUI. Valid values: `enable`, `disable`.
* `gui_auto_upgrade_setup_warning` - Enable/disable the automatic patch upgrade setup prompt on the GUI. Valid values: `enable`, `disable`.
* `gui_workflow_management` - Enable/disable Workflow management features on the GUI. Valid values: `enable`, `disable`.
* `gui_cdn_usage` - Enable/disable Load GUI static files from a CDN. Valid values: `enable`, `disable`.
* `alias` - Alias for your FortiGate unit.
* `strong_crypto` - Enable to use strong encryption and only allow strong ciphers and digest for HTTPS/SSH/TLS/SSL functions. Valid values: `enable`, `disable`.
* `ssh_cbc_cipher` - Enable/disable CBC cipher for SSH access. Valid values: `enable`, `disable`.
* `ssh_hmac_md5` - Enable/disable HMAC-MD5 for SSH access. Valid values: `enable`, `disable`.
* `ssh_kex_sha1` - Enable/disable SHA1 key exchange for SSH access. Valid values: `enable`, `disable`.
* `ssh_mac_weak` - Enable/disable HMAC-SHA1 and UMAC-64-ETM for SSH access. Valid values: `enable`, `disable`.
* `ssl_static_key_ciphers` - Enable/disable static key ciphers in SSL/TLS connections (e.g. AES128-SHA, AES256-SHA, AES128-SHA256, AES256-SHA256). Valid values: `enable`, `disable`.
* `ssh_kex_algo` - Select one or more SSH kex algorithms.
* `ssh_enc_algo` - Select one or more SSH ciphers. Valid values: `chacha20-poly1305@openssh.com`, `aes128-ctr`, `aes192-ctr`, `aes256-ctr`, `arcfour256`, `arcfour128`, `aes128-cbc`, `3des-cbc`, `blowfish-cbc`, `cast128-cbc`, `aes192-cbc`, `aes256-cbc`, `arcfour`, `rijndael-cbc@lysator.liu.se`, `aes128-gcm@openssh.com`, `aes256-gcm@openssh.com`.
* `ssh_mac_algo` - Select one or more SSH MAC algorithms. Valid values: `hmac-md5`, `hmac-md5-etm@openssh.com`, `hmac-md5-96`, `hmac-md5-96-etm@openssh.com`, `hmac-sha1`, `hmac-sha1-etm@openssh.com`, `hmac-sha2-256`, `hmac-sha2-256-etm@openssh.com`, `hmac-sha2-512`, `hmac-sha2-512-etm@openssh.com`, `hmac-ripemd160`, `hmac-ripemd160@openssh.com`, `hmac-ripemd160-etm@openssh.com`, `umac-64@openssh.com`, `umac-128@openssh.com`, `umac-64-etm@openssh.com`, `umac-128-etm@openssh.com`.
* `ssh_hostkey_algo` - Select one or more SSH hostkey algorithms.
* `ssh_hostkey_override` - Enable/disable SSH host key override in SSH daemon. Valid values: `disable`, `enable`.
* `ssh_hostkey_password` - Password for ssh-hostkey.
* `ssh_hostkey` - Config SSH host key.
* `snat_route_change` - Enable/disable the ability to change the static NAT route. Valid values: `enable`, `disable`.
* `speedtest_server` - Enable/disable speed test server. Valid values: `enable`, `disable`.
* `cli_audit_log` - Enable/disable CLI audit log. Valid values: `enable`, `disable`.
* `dh_params` - Number of bits to use in the Diffie-Hellman exchange for HTTPS/SSH protocols. Valid values: `1024`, `1536`, `2048`, `3072`, `4096`, `6144`, `8192`.
* `fds_statistics` - Enable/disable sending IPS, Application Control, and AntiVirus data to FortiGuard. This data is used to improve FortiGuard services and is not shared with external parties and is protected by Fortinet's privacy policy. Valid values: `enable`, `disable`.
* `fds_statistics_period` - FortiGuard statistics collection period in minutes. (1 - 1440 min (1 min to 24 hours), default = 60).
* `multicast_forward` - Enable/disable multicast forwarding. Valid values: `enable`, `disable`.
* `mc_ttl_notchange` - Enable/disable no modification of multicast TTL. Valid values: `enable`, `disable`.
* `asymroute` - Enable/disable asymmetric route. Valid values: `enable`, `disable`.
* `tcp_option` - Enable SACK, timestamp and MSS TCP options. Valid values: `enable`, `disable`.
* `lldp_transmission` - Enable/disable Link Layer Discovery Protocol (LLDP) transmission. Valid values: `enable`, `disable`.
* `lldp_reception` - Enable/disable Link Layer Discovery Protocol (LLDP) reception. Valid values: `enable`, `disable`.
* `proxy_auth_timeout` - Authentication timeout in minutes for authenticated users (1 - 300 min, default = 10).
* `proxy_keep_alive_mode` - Control if users must re-authenticate after a session is closed, traffic has been idle, or from the point at which the user was authenticated. Valid values: `session`, `traffic`, `re-authentication`.
* `proxy_re_authentication_time` - The time limit that users must re-authenticate if proxy-keep-alive-mode is set to re-authenticate (1  - 86400 sec, default=30s.
* `proxy_re_authentication_mode` - Control if users must re-authenticate after a session is closed, traffic has been idle, or from the point at which the user was first created. Valid values: `session`, `traffic`, `absolute`.
* `proxy_auth_lifetime` - Enable/disable authenticated users lifetime control.  This is a cap on the total time a proxy user can be authenticated for after which re-authentication will take place. Valid values: `enable`, `disable`.
* `proxy_auth_lifetime_timeout` - Lifetime timeout in minutes for authenticated users (5  - 65535 min, default=480 (8 hours)).
* `proxy_resource_mode` - Enable/disable use of the maximum memory usage on the FortiGate unit's proxy processing of resources, such as block lists, allow lists, and external resources. Valid values: `enable`, `disable`.
* `proxy_cert_use_mgmt_vdom` - Enable/disable using management VDOM to send requests. Valid values: `enable`, `disable`.
* `sys_perf_log_interval` - Time in minutes between updates of performance statistics logging. (1 - 15 min, default = 5, 0 = disabled).
* `check_protocol_header` - Level of checking performed on protocol headers. Strict checking is more thorough but may affect performance. Loose checking is ok in most cases. Valid values: `loose`, `strict`.
* `vip_arp_range` - Controls the number of ARPs that the FortiGate sends for a Virtual IP (VIP) address range. Valid values: `unlimited`, `restricted`.
* `reset_sessionless_tcp` - Action to perform if the FortiGate receives a TCP packet but cannot find a corresponding session in its session table. NAT/Route mode only. Valid values: `enable`, `disable`.
* `allow_traffic_redirect` - Disable to allow traffic to be routed back on a different interface. Valid values: `enable`, `disable`.
* `ipv6_allow_traffic_redirect` - Disable to prevent IPv6 traffic with same local ingress and egress interface from being forwarded without policy check. Valid values: `enable`, `disable`.
* `strict_dirty_session_check` - Enable to check the session against the original policy when revalidating. This can prevent dropping of redirected sessions when web-filtering and authentication are enabled together. If this option is enabled, the FortiGate unit deletes a session if a routing or policy change causes the session to no longer match the policy that originally allowed the session. Valid values: `enable`, `disable`.
* `tcp_halfclose_timer` - Number of seconds the FortiGate unit should wait to close a session after one peer has sent a FIN packet but the other has not responded (1 - 86400 sec (1 day), default = 120).
* `tcp_halfopen_timer` - Number of seconds the FortiGate unit should wait to close a session after one peer has sent an open session packet but the other has not responded (1 - 86400 sec (1 day), default = 10).
* `tcp_timewait_timer` - Length of the TCP TIME-WAIT state in seconds (1 - 300 sec, default = 1).
* `tcp_rst_timer` - Length of the TCP CLOSE state in seconds (5 - 300 sec, default = 5).
* `udp_idle_timer` - UDP connection session timeout. This command can be useful in managing CPU and memory resources (1 - 86400 seconds (1 day), default = 60).
* `block_session_timer` - Duration in seconds for blocked sessions (1 - 300 sec  (5 minutes), default = 30).
* `ip_src_port_range` - IP source port range used for traffic originating from the FortiGate unit.
* `pre_login_banner` - Enable/disable displaying the administrator access disclaimer message on the login page before an administrator logs in. Valid values: `enable`, `disable`.
* `post_login_banner` - Enable/disable displaying the administrator access disclaimer message after an administrator successfully logs in. Valid values: `disable`, `enable`.
* `tftp` - Enable/disable TFTP. Valid values: `enable`, `disable`.
* `av_failopen` - Set the action to take if the FortiGate is running low on memory or the proxy connection limit has been reached. Valid values: `pass`, `off`, `one-shot`.
* `av_failopen_session` - When enabled and a proxy for a protocol runs out of room in its session table, that protocol goes into failopen mode and enacts the action specified by av-failopen. Valid values: `enable`, `disable`.
* `memory_use_threshold_extreme` - Threshold at which memory usage is considered extreme (new sessions are dropped) (% of total RAM, default = 95).
* `memory_use_threshold_red` - Threshold at which memory usage forces the FortiGate to enter conserve mode (% of total RAM, default = 88).
* `memory_use_threshold_green` - Threshold at which memory usage forces the FortiGate to exit conserve mode (% of total RAM, default = 82).
* `ip_fragment_mem_thresholds` - Maximum memory (MB) used to reassemble IPv4/IPv6 fragments.
* `cpu_use_threshold` - Threshold at which CPU usage is reported. (% of total CPU, default = 90).
* `log_single_cpu_high` - Enable/disable logging the event of a single CPU core reaching CPU usage threshold. Valid values: `enable`, `disable`.
* `check_reset_range` - Configure ICMP error message verification. You can either apply strict RST range checking or disable it. Valid values: `strict`, `disable`.
* `vdom_mode` - Enable/disable support for split/multiple virtual domains (VDOMs). Valid values: `no-vdom`, `split-vdom`, `multi-vdom`.
* `vdom_admin` - Enable/disable support for multiple virtual domains (VDOMs). Valid values: `enable`, `disable`.
* `long_vdom_name` - Enable/disable long VDOM name support. Valid values: `enable`, `disable`.
* `edit_vdom_prompt` - Enable/disable edit new VDOM prompt. Valid values: `enable`, `disable`.
* `admin_port` - Administrative access port for HTTP. (1 - 65535, default = 80).
* `admin_sport` - Administrative access port for HTTPS. (1 - 65535, default = 443).
* `admin_host` - Administrative host for HTTP and HTTPS. When set, will be used in lieu of the client's Host header for any redirection.
* `admin_https_redirect` - Enable/disable redirection of HTTP administration access to HTTPS. Valid values: `enable`, `disable`.
* `admin_hsts_max_age` - HTTPS Strict-Transport-Security header max-age in seconds. A value of 0 will reset any HSTS records in the browser.When admin-https-redirect is disabled the header max-age will be 0.
* `admin_ssh_password` - Enable/disable password authentication for SSH admin access. Valid values: `enable`, `disable`.
* `admin_restrict_local` - Enable/disable local admin authentication restriction when remote authenticator is up and running. (default = disable) Valid values: `enable`, `disable`.
* `admin_ssh_port` - Administrative access port for SSH. (1 - 65535, default = 22).
* `admin_ssh_grace_time` - Maximum time in seconds permitted between making an SSH connection to the FortiGate unit and authenticating (10 - 3600 sec (1 hour), default 120).
* `admin_ssh_v1` - Enable/disable SSH v1 compatibility. Valid values: `enable`, `disable`.
* `admin_telnet` - Enable/disable TELNET service. Valid values: `enable`, `disable`.
* `admin_telnet_port` - Administrative access port for TELNET. (1 - 65535, default = 23).
* `admin_forticloud_sso_login` - Enable/disable FortiCloud admin login via SSO. Valid values: `enable`, `disable`.
* `admin_forticloud_sso_default_profile` - Override access profile.
* `default_service_source_port` - Default service source port range. (default=1-65535)
* `admin_maintainer` - Enable/disable maintainer administrator login. When enabled, the maintainer account can be used to log in from the console after a hard reboot. The password is "bcpb" followed by the FortiGate unit serial number. You have limited time to complete this login. Valid values: `enable`, `disable`.
* `admin_server_cert` - Server certificate that the FortiGate uses for HTTPS administrative connections.
* `user_server_cert` - Certificate to use for https user authentication.
* `admin_https_pki_required` - Enable/disable admin login method. Enable to force administrators to provide a valid certificate to log in if PKI is enabled. Disable to allow administrators to log in with a certificate or password. Valid values: `enable`, `disable`.
* `wifi_certificate` - Certificate to use for WiFi authentication.
* `dhcp_lease_backup_interval` - DHCP leases backup interval in seconds (10 - 3600, default = 60).
* `wifi_ca_certificate` - CA certificate that verifies the WiFi certificate.
* `auth_http_port` - User authentication HTTP port. (1 - 65535). On FortiOS versions 6.2.0-6.2.6: default = 80. On FortiOS versions >= 6.4.0: default = 1000.
* `auth_https_port` - User authentication HTTPS port. (1 - 65535). On FortiOS versions 6.2.0-6.2.6: default = 443. On FortiOS versions >= 6.4.0: default = 1003.
* `auth_ike_saml_port` - User IKE SAML authentication port (0 - 65535, default = 1001).
* `auth_keepalive` - Enable to prevent user authentication sessions from timing out when idle. Valid values: `enable`, `disable`.
* `policy_auth_concurrent` - Number of concurrent firewall use logins from the same user (1 - 100, default = 0 means no limit).
* `auth_session_limit` - Action to take when the number of allowed user authenticated sessions is reached. Valid values: `block-new`, `logout-inactive`.
* `auth_cert` - Server certificate that the FortiGate uses for HTTPS firewall authentication connections.
* `clt_cert_req` - Enable/disable requiring administrators to have a client certificate to log into the GUI using HTTPS. Valid values: `enable`, `disable`.
* `fortiservice_port` - FortiService port (1 - 65535, default = 8013). Used by FortiClient endpoint compliance. Older versions of FortiClient used a different port.
* `endpoint_control_portal_port` - Endpoint control portal port (1 - 65535).
* `endpoint_control_fds_access` - Enable/disable access to the FortiGuard network for non-compliant endpoints. Valid values: `enable`, `disable`.
* `tp_mc_skip_policy` - Enable/disable skip policy check and allow multicast through. Valid values: `enable`, `disable`.
* `cfg_save` - Configuration file save mode for CLI changes. Valid values: `automatic`, `manual`, `revert`.
* `cfg_revert_timeout` - Time-out for reverting to the last saved configuration. (10 - 4294967295 seconds, default = 600).
* `reboot_upon_config_restore` - Enable/disable reboot of system upon restoring configuration. Valid values: `enable`, `disable`.
* `admin_scp` - Enable/disable using SCP to download the system configuration. You can use SCP as an alternative method for backing up the configuration. Valid values: `enable`, `disable`.
* `security_rating_result_submission` - Enable/disable the submission of Security Rating results to FortiGuard. Valid values: `enable`, `disable`.
* `security_rating_run_on_schedule` - Enable/disable scheduled runs of Security Rating. Valid values: `enable`, `disable`.
* `wireless_controller` - Enable/disable the wireless controller feature to use the FortiGate unit to manage FortiAPs. Valid values: `enable`, `disable`.
* `wireless_controller_port` - Port used for the control channel in wireless controller mode (wireless-mode is ac). The data channel port is the control channel port number plus one (1024 - 49150, default = 5246).
* `fortiextender_data_port` - FortiExtender data port (1024 - 49150, default = 25246).
* `fortiextender` - Enable/disable FortiExtender. Valid values: `enable`, `disable`.
* `extender_controller_reserved_network` - Configure reserved network subnet for managed LAN extension FortiExtenders. This is available when the extender daemon is running.
* `fortiextender_discovery_lockdown` - Enable/disable FortiExtender CAPWAP lockdown. Valid values: `disable`, `enable`.
* `fortiextender_vlan_mode` - Enable/disable FortiExtender VLAN mode. Valid values: `enable`, `disable`.
* `fortiextender_provision_on_authorization` - Enable/disable automatic provisioning of latest FortiExtender firmware on authorization. Valid values: `enable`, `disable`.
* `switch_controller` - Enable/disable switch controller feature. Switch controller allows you to manage FortiSwitch from the FortiGate itself. Valid values: `disable`, `enable`.
* `switch_controller_reserved_network` - Enable reserved network subnet for controlled switches. This is available when the switch controller is enabled.
* `dnsproxy_worker_count` - DNS proxy worker count.
* `url_filter_count` - URL filter daemon count.
* `proxy_worker_count` - Proxy worker count.
* `scanunit_count` - Number of scanunits. The range and the default depend on the number of CPUs. Only available on FortiGate units with multiple CPUs.
* `proxy_hardware_acceleration` - Enable/disable email proxy hardware acceleration. Valid values: `disable`, `enable`.
* `proxy_kxp_hardware_acceleration` - Enable/disable using the content processor to accelerate KXP traffic. Valid values: `disable`, `enable`.
* `proxy_cipher_hardware_acceleration` - Enable/disable using content processor (CP8 or CP9) hardware acceleration to encrypt and decrypt IPsec and SSL traffic. Valid values: `disable`, `enable`.
* `fgd_alert_subscription` - Type of alert to retrieve from FortiGuard. Valid values: `advisory`, `latest-threat`, `latest-virus`, `latest-attack`, `new-antivirus-db`, `new-attack-db`.
* `ipsec_hmac_offload` - Enable/disable offloading (hardware acceleration) of HMAC processing for IPsec VPN. Valid values: `enable`, `disable`.
* `ipv6_accept_dad` - Enable/disable acceptance of IPv6 Duplicate Address Detection (DAD).
* `ipv6_allow_anycast_probe` - Enable/disable IPv6 address probe through Anycast. Valid values: `enable`, `disable`.
* `ipv6_allow_multicast_probe` - Enable/disable IPv6 address probe through Multicast. Valid values: `enable`, `disable`.
* `ipv6_allow_local_in_silent_drop` - Enable/disable silent drop of IPv6 local-in traffic. Valid values: `enable`, `disable`.
* `ipv6_allow_local_in_slient_drop` - Enable/disable silent drop of IPv6 local-in traffic. Valid values: `enable`, `disable`.
* `csr_ca_attribute` - Enable/disable the CA attribute in certificates. Some CA servers reject CSRs that have the CA attribute. Valid values: `enable`, `disable`.
* `wimax_4g_usb` - Enable/disable comparability with WiMAX 4G USB devices. Valid values: `enable`, `disable`.
* `cert_chain_max` - Maximum number of certificates that can be traversed in a certificate chain.
* `sslvpn_max_worker_count` - Maximum number of SSL VPN processes. Upper limit for this value is the number of CPUs and depends on the model.
* `vpn_ems_sn_check` - Enable/disable verification of EMS serial number in SSL-VPN and IPsec VPN connection. Valid values: `enable`, `disable`.
* `sslvpn_web_mode` - Enable/disable SSL-VPN web mode. Valid values: `enable`, `disable`.
* `sslvpn_ems_sn_check` - Enable/disable verification of EMS serial number in SSL-VPN connection. Valid values: `enable`, `disable`.
* `sslvpn_kxp_hardware_acceleration` - Enable/disable SSL VPN KXP hardware acceleration. Valid values: `enable`, `disable`.
* `sslvpn_cipher_hardware_acceleration` - Enable/disable SSL VPN hardware acceleration. Valid values: `enable`, `disable`.
* `sslvpn_plugin_version_check` - Enable/disable checking browser's plugin version by SSL VPN. Valid values: `enable`, `disable`.
* `two_factor_ftk_expiry` - FortiToken authentication session timeout (60 - 600 sec (10 minutes), default = 60).
* `two_factor_email_expiry` - Email-based two-factor authentication session timeout (30 - 300 seconds (5 minutes), default = 60).
* `two_factor_sms_expiry` - SMS-based two-factor authentication session timeout (30 - 300 sec, default = 60).
* `two_factor_fac_expiry` - FortiAuthenticator token authentication session timeout (10 - 3600 seconds (1 hour), default = 60).
* `two_factor_ftm_expiry` - FortiToken Mobile session timeout (1 - 168 hours (7 days), default = 72).
* `per_user_bal` - Enable/disable per-user block/allow list filter. Valid values: `enable`, `disable`.
* `per_user_bwl` - Enable/disable per-user black/white list filter. Valid values: `enable`, `disable`.
* `virtual_server_count` - Maximum number of virtual server processes to create. The maximum is the number of CPU cores. This is not available on single-core CPUs.
* `virtual_server_hardware_acceleration` - Enable/disable virtual server hardware acceleration. Valid values: `disable`, `enable`.
* `wad_worker_count` - Number of explicit proxy WAN optimization daemon (WAD) processes. By default WAN optimization, explicit proxy, and web caching is handled by all of the CPU cores in a FortiGate unit.
* `wad_csvc_cs_count` - Number of concurrent WAD-cache-service object-cache processes.
* `wad_csvc_db_count` - Number of concurrent WAD-cache-service byte-cache processes.
* `wad_source_affinity` - Enable/disable dispatching traffic to WAD workers based on source affinity. Valid values: `disable`, `enable`.
* `wad_memory_change_granularity` - Minimum percentage change in system memory usage detected by the wad daemon prior to adjusting TCP window size for any active connection.
* `login_timestamp` - Enable/disable login time recording. Valid values: `enable`, `disable`.
* `miglogd_children` - Number of logging (miglogd) processes to be allowed to run. Higher number can reduce performance; lower number can slow log processing time.
* `special_file_23_support` - Enable/disable IPS detection of HIBUN format files when using Data Leak Protection. Valid values: `disable`, `enable`.
* `log_uuid_policy` - Enable/disable insertion of policy UUIDs to traffic logs. Valid values: `enable`, `disable`.
* `log_uuid_address` - Enable/disable insertion of address UUIDs to traffic logs. Valid values: `enable`, `disable`.
* `log_ssl_connection` - Enable/disable logging of SSL connection events. Valid values: `enable`, `disable`.
* `gui_rest_api_cache` - Enable/disable REST API result caching on FortiGate. Valid values: `enable`, `disable`.
* `gui_cdn_domain_override` - Domain of CDN server.
* `gui_fortiguard_resource_fetch` - Enable/disable retrieving static GUI resources from FortiGuard. Disabling it will improve GUI load time for air-gapped environments. Valid values: `enable`, `disable`.
* `arp_max_entry` - Maximum number of dynamically learned MAC addresses that can be added to the ARP table (131072 - 2147483647, default = 131072).
* `ha_affinity` - Affinity setting for HA daemons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
* `bfd_affinity` - Affinity setting for BFD daemon (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
* `cmdbsvr_affinity` - Affinity setting for cmdbsvr (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
* `av_affinity` - Affinity setting for AV scanning (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
* `wad_affinity` - Affinity setting for wad (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
* `ips_affinity` - Affinity setting for IPS (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx; allowed CPUs must be less than total number of IPS engine daemons).
* `miglog_affinity` - Affinity setting for logging. On FortiOS versions 6.2.0-7.2.3: 64-bit hexadecimal value in the format of xxxxxxxxxxxxxxxx. On FortiOS versions >= 7.2.4: hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx.
* `syslog_affinity` - Affinity setting for syslog (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
* `url_filter_affinity` - URL filter CPU affinity.
* `ndp_max_entry` - Maximum number of NDP table entries (set to 65,536 or higher; if set to 0, kernel holds 65,536 entries).
* `br_fdb_max_entry` - Maximum number of bridge forwarding database (FDB) entries.
* `max_route_cache_size` - Maximum number of IP route cache entries (0 - 2147483647).
* `ipsec_qat_offload` - Enable/disable QAT offloading (Intel QuickAssist) for IPsec VPN traffic. QuickAssist can accelerate IPsec encryption and decryption. Valid values: `enable`, `disable`.
* `ipsec_round_robin` - Enable/disable round-robin redistribution to multiple CPUs for IPsec VPN traffic. Valid values: `enable`, `disable`.
* `ipsec_asic_offload` - Enable/disable ASIC offloading (hardware acceleration) for IPsec VPN traffic. Hardware acceleration can offload IPsec VPN sessions and accelerate encryption and decryption. Valid values: `enable`, `disable`.
* `ipsec_soft_dec_async` - Enable/disable software decryption asynchronization (using multiple CPUs to do decryption) for IPsec VPN traffic. Valid values: `enable`, `disable`.
* `ike_embryonic_limit` - Maximum number of IPsec tunnels to negotiate simultaneously.
* `device_idle_timeout` - Time in seconds that a device must be idle to automatically log the device user out. (30 - 31536000 sec (30 sec to 1 year), default = 300).
* `user_device_store_max_devices` - Maximum number of devices allowed in user device store.
* `user_device_store_max_users` - Maximum number of users allowed in user device store.
* `user_device_store_max_unified_mem` - Maximum unified memory allowed in user device store.
* `device_identification_active_scan_delay` - Number of seconds to passively scan a device before performing an active scan. (20 - 3600 sec, (20 sec to 1 hour), default = 90).
* `compliance_check` - Enable/disable global PCI DSS compliance check. Valid values: `enable`, `disable`.
* `compliance_check_time` - Time of day to run scheduled PCI DSS compliance checks.
* `gui_device_latitude` - Add the latitude of the location of this FortiGate to position it on the Threat Map.
* `gui_device_longitude` - Add the longitude of the location of this FortiGate to position it on the Threat Map.
* `private_data_encryption` - Enable/disable private data encryption using an AES 128-bit key. Valid values: `disable`, `enable`.
* `auto_auth_extension_device` - Enable/disable automatic authorization of dedicated Fortinet extension devices. Valid values: `enable`, `disable`.
* `gui_theme` - Color scheme for the administration GUI.
* `gui_date_format` - Default date format used throughout GUI. Valid values: `yyyy/MM/dd`, `dd/MM/yyyy`, `MM/dd/yyyy`, `yyyy-MM-dd`, `dd-MM-yyyy`, `MM-dd-yyyy`.
* `gui_date_time_source` - Source from which the FortiGate GUI uses to display date and time entries. Valid values: `system`, `browser`.
* `igmp_state_limit` - Maximum number of IGMP memberships (96 - 64000, default = 3200).
* `cloud_communication` - Enable/disable all cloud communication. Valid values: `enable`, `disable`.
* `fec_port` - Local UDP port for Forward Error Correction (49152 - 65535).
* `ipsec_ha_seqjump_rate` - ESP jump ahead rate (1G - 10G pps equivalent).
* `fortitoken_cloud` - Enable/disable FortiToken Cloud service. Valid values: `enable`, `disable`.
* `fortitoken_cloud_push_status` - Enable/disable FTM push service of FortiToken Cloud. Valid values: `enable`, `disable`.
* `fortitoken_cloud_sync_interval` - Interval in which to clean up remote users in FortiToken Cloud (0 - 336 hours (14 days), default = 24, disable = 0).
* `faz_disk_buffer_size` - Maximum disk buffer size to temporarily store logs destined for FortiAnalyzer. To be used in the event that FortiAnalyzer is unavailalble.
* `irq_time_accounting` - Configure CPU IRQ time accounting mode. Valid values: `auto`, `force`.
* `fortiipam_integration` - Enable/disable integration with the FortiIPAM cloud service. Valid values: `enable`, `disable`.
* `management_ip` - Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
* `management_port` - Overriding port for management connection (Overrides admin port).
* `management_port_use_admin_sport` - Enable/disable use of the admin-sport setting for the management port. If disabled, FortiGate will allow user to specify management-port. Valid values: `enable`, `disable`.
* `forticonverter_integration` - Enable/disable FortiConverter integration service. Valid values: `enable`, `disable`.
* `forticonverter_config_upload` - Enable/disable config upload to FortiConverter. Valid values: `once`, `disable`.
* `internet_service_database` - Configure which Internet Service database size to download from FortiGuard and use.
* `internet_service_download_list` - Configure which on-demand Internet Service IDs are to be downloaded. The structure of `internet_service_download_list` block is documented below.
* `early_tcp_npu_session` - Enable/disable early TCP NPU session. Valid values: `enable`, `disable`.
* `npu_neighbor_update` - Enable/disable sending of probing packets to update neighbors for offloaded sessions. Valid values: `enable`, `disable`.
* `interface_subnet_usage` - Enable/disable allowing use of interface-subnet setting in firewall addresses (default = enable). Valid values: `disable`, `enable`.
* `sflowd_max_children_num` - Maximum number of sflowd child processes allowed to run.
* `fortigslb_integration` - Enable/disable integration with the FortiGSLB cloud service. Valid values: `disable`, `enable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `internet_service_download_list` block supports:

* `id` - Internet Service ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Global can be imported using any of these accepted formats:
```
$ terraform import fortios_system_global.labelname SystemGlobal

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_global.labelname SystemGlobal
$ unset "FORTIOS_IMPORT_TABLE"
```
