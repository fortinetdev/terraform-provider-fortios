// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure global attributes.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemGlobal() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemGlobalRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"language": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_allow_incompatible_fabric_fgt": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_replacement_message_groups": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_local_out": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_certificates": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_custom_language": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_wireless_opensecurity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_app_detection_sdwan": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_display_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_fortigate_cloud_sandbox": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_fortisandbox_cloud": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_firmware_upgrade_warning": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_firmware_upgrade_setup_warning": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_lines_per_page": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_https_ssl_versions": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_https_ssl_ciphersuites": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_https_ssl_banned_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admintimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_console_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ssd_trim_freq": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssd_trim_hour": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ssd_trim_min": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ssd_trim_weekday": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssd_trim_date": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_concurrent": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_lockout_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_lockout_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"refresh": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"failtime": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"purdue_level": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"daily_restart": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"restart_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wad_restart_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wad_restart_start_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wad_restart_end_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wad_p2s_max_body_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"radius_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"speedtestd_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"speedtestd_ctrl_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_login_max": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"remoteauthtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ldapconntimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"batch_cmdb": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_dlpstat_memory": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"multi_factor_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"autorun_log_fsck": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"traffic_priority": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"traffic_priority_level": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"quic_congestion_control_algo": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"quic_max_datagram_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"quic_udp_payload_size_shaping_per_cid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"quic_ack_thresold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"quic_pmtud": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"quic_tls_handshake_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"anti_replay": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"send_pmtu_icmp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"honor_df": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pmtu_discovery": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_switch_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"split_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"revision_image_auto_backup": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"revision_backup_on_logout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_allow_default_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_forticare_registration_setup_warning": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_auto_upgrade_setup_warning": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_workflow_management": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_cdn_usage": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"strong_crypto": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_cbc_cipher": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_hmac_md5": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_kex_sha1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_mac_weak": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_static_key_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_kex_algo": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_enc_algo": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_mac_algo": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_hostkey_algo": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_hostkey_override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_hostkey_password": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"ssh_hostkey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"snat_route_change": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_snat_route_change": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"speedtest_server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cli_audit_log": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dh_params": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fds_statistics": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fds_statistics_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"multicast_forward": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mc_ttl_notchange": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"asymroute": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tcp_option": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lldp_transmission": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lldp_reception": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxy_auth_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"proxy_keep_alive_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxy_re_authentication_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"proxy_re_authentication_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxy_auth_lifetime": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxy_auth_lifetime_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"proxy_resource_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxy_cert_use_mgmt_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sys_perf_log_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"check_protocol_header": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vip_arp_range": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"reset_sessionless_tcp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_traffic_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_allow_traffic_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"strict_dirty_session_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tcp_halfclose_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tcp_halfopen_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tcp_timewait_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tcp_rst_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"udp_idle_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"block_session_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ip_src_port_range": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pre_login_banner": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"post_login_banner": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tftp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"av_failopen": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"av_failopen_session": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory_use_threshold_extreme": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"memory_use_threshold_red": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"memory_use_threshold_green": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ip_fragment_mem_thresholds": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ip_fragment_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ipv6_fragment_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cpu_use_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"log_single_cpu_high": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"check_reset_range": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_report": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom_admin": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"long_vdom_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"edit_vdom_prompt": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_sport": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_https_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_hsts_max_age": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_ssh_password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_restrict_local": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_ssh_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_ssh_grace_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_ssh_v1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_telnet": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_telnet_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin_forticloud_sso_login": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_forticloud_sso_default_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_service_source_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_maintainer": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_https_pki_required": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wifi_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_lease_backup_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"wifi_ca_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_http_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"auth_https_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"auth_ike_saml_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"auth_keepalive": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"policy_auth_concurrent": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"auth_session_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"clt_cert_req": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fortiservice_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"endpoint_control_portal_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"endpoint_control_fds_access": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tp_mc_skip_policy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cfg_save": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cfg_revert_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"reboot_upon_config_restore": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_scp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_rating_result_submission": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_rating_run_on_schedule": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wireless_controller": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wireless_controller_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"fortiextender_data_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"fortiextender": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"extender_controller_reserved_network": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fortiextender_discovery_lockdown": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fortiextender_vlan_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fortiextender_provision_on_authorization": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"switch_controller": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"switch_controller_reserved_network": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dnsproxy_worker_count": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"url_filter_count": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"httpd_max_worker_count": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"proxy_worker_count": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"scanunit_count": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"proxy_hardware_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxy_kxp_hardware_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxy_cipher_hardware_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fgd_alert_subscription": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipsec_hmac_offload": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_accept_dad": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ipv6_allow_anycast_probe": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_allow_multicast_probe": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_allow_local_in_silent_drop": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_allow_local_in_slient_drop": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"csr_ca_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wimax_4g_usb": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cert_chain_max": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"sslvpn_max_worker_count": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"sslvpn_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vpn_ems_sn_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sslvpn_web_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sslvpn_ems_sn_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sslvpn_kxp_hardware_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sslvpn_cipher_hardware_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sslvpn_plugin_version_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"two_factor_ftk_expiry": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"two_factor_email_expiry": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"two_factor_sms_expiry": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"two_factor_fac_expiry": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"two_factor_ftm_expiry": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"per_user_bal": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"per_user_bwl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_server_count": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"virtual_server_hardware_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wad_worker_count": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"wad_worker_dev_cache": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"wad_csvc_cs_count": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"wad_csvc_db_count": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"wad_source_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wad_memory_change_granularity": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"login_timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_conflict_detection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"miglogd_children": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"log_daemon_cpu_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"special_file_23_support": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"log_uuid_policy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"log_uuid_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"log_ssl_connection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_rest_api_cache": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"rest_api_key_url_query": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_cdn_domain_override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_fortiguard_resource_fetch": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"arp_max_entry": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ha_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bfd_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cmdbsvr_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"av_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wad_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ips_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"miglog_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"syslog_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url_filter_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"router_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ndp_max_entry": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"br_fdb_max_entry": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_route_cache_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ipsec_qat_offload": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipsec_round_robin": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipsec_asic_offload": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipsec_soft_dec_async": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ike_embryonic_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"device_idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"user_device_store_max_devices": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"user_device_store_max_device_mem": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"user_device_store_max_users": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"user_device_store_max_unified_mem": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"device_identification_active_scan_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"compliance_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"compliance_check_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_device_latitude": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_device_longitude": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_data_encryption": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_auth_extension_device": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_theme": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_date_format": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_date_time_source": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"igmp_state_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cloud_communication": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fec_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ipsec_ha_seqjump_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"fortitoken_cloud": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fortitoken_cloud_push_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fortitoken_cloud_region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fortitoken_cloud_sync_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"faz_disk_buffer_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"irq_time_accounting": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fortiipam_integration": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"management_port_use_admin_sport": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"forticonverter_integration": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"forticonverter_config_upload": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"internet_service_database": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"internet_service_download_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"geoip_full_db": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"early_tcp_npu_session": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"npu_neighbor_update": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"delay_tcp_npu_session": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface_subnet_usage": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sflowd_max_children_num": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"fortigslb_integration": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_history_password_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"auth_session_auto_backup": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_session_auto_backup_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"scim_https_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"scim_http_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"scim_server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"application_bandwidth_tracking": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tls_session_cache": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemGlobalRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "SystemGlobal"

	o, err := c.ReadSystemGlobal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemGlobal: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemGlobal(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemGlobal from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemGlobalLanguage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiAllowIncompatibleFabricFgt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiReplacementMessageGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiLocalOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiCertificates(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiCustomLanguage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiWirelessOpensecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiAppDetectionSdwan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiDisplayHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiFortigateCloudSandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiFortisandboxCloud(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiFirmwareUpgradeWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiFirmwareUpgradeSetupWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiLinesPerPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminHttpsSslVersions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminHttpsSslCiphersuites(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminHttpsSslBannedCiphers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdmintimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminConsoleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSsdTrimFreq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSsdTrimHour(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSsdTrimMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSsdTrimWeekday(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSsdTrimDate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminConcurrent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminLockoutThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminLockoutDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalRefresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalFailtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalPurdueLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDailyRestart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalRestartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalWadRestartMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalWadRestartStartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalWadRestartEndTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalWadP2SMaxBodySize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalRadiusPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSpeedtestdServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSpeedtestdCtrlPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminLoginMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalRemoteauthtimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalLdapconntimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalBatchCmdb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalMaxDlpstatMemory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalMultiFactorAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAutorunLogFsck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTimezone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTrafficPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTrafficPriorityLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalQuicCongestionControlAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalQuicMaxDatagramSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalQuicUdpPayloadSizeShapingPerCid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalQuicAckThresold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalQuicPmtud(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalQuicTlsHandshakeTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAntiReplay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSendPmtuIcmp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalHonorDf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalPmtuDiscovery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalVirtualSwitchVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSplitPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalRevisionImageAutoBackup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalRevisionBackupOnLogout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalManagementVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiAllowDefaultHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiForticareRegistrationSetupWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiAutoUpgradeSetupWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiWorkflowManagement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiCdnUsage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAlias(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalStrongCrypto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSshCbcCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSshHmacMd5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSshKexSha1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSshMacWeak(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSslStaticKeyCiphers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSshKexAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSshEncAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSshMacAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSshHostkeyAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSshHostkeyOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSshHostkeyPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSshHostkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSnatRouteChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpv6SnatRouteChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSpeedtestServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalCliAuditLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDhParams(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalFdsStatistics(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalFdsStatisticsPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalMulticastForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalMcTtlNotchange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAsymroute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTcpOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalLldpTransmission(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalLldpReception(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalProxyAuthTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalProxyKeepAliveMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalProxyReAuthenticationTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalProxyReAuthenticationMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalProxyAuthLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalProxyAuthLifetimeTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalProxyResourceMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalProxyCertUseMgmtVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSysPerfLogInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalCheckProtocolHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalVipArpRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalResetSessionlessTcp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAllowTrafficRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpv6AllowTrafficRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalStrictDirtySessionCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTcpHalfcloseTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTcpHalfopenTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTcpTimewaitTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTcpRstTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalUdpIdleTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalBlockSessionTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpSrcPortRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalPreLoginBanner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalPostLoginBanner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTftp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAvFailopen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAvFailopenSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalMemoryUseThresholdExtreme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalMemoryUseThresholdRed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalMemoryUseThresholdGreen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpFragmentMemThresholds(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpFragmentTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpv6FragmentTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalCpuUseThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalLogSingleCpuHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalCheckResetRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalUpgradeReport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalVdomMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalVdomAdmin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalLongVdomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalEditVdomPrompt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminSport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminHttpsRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminHstsMaxAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminSshPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminRestrictLocal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminSshPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminSshGraceTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminSshV1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminTelnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminTelnetPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminForticloudSsoLogin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminForticloudSsoDefaultProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDefaultServiceSourcePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminMaintainer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalUserServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminHttpsPkiRequired(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalWifiCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDhcpLeaseBackupInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalWifiCaCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAuthHttpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAuthHttpsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAuthIkeSamlPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAuthKeepalive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalPolicyAuthConcurrent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAuthSessionLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAuthCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalCltCertReq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalFortiservicePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalEndpointControlPortalPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalEndpointControlFdsAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTpMcSkipPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalCfgSave(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalCfgRevertTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalRebootUponConfigRestore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAdminScp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSecurityRatingResultSubmission(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSecurityRatingRunOnSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalWirelessController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalWirelessControllerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalFortiextenderDataPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalFortiextender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalExtenderControllerReservedNetwork(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemGlobalFortiextenderDiscoveryLockdown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalFortiextenderVlanMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalFortiextenderProvisionOnAuthorization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSwitchController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSwitchControllerReservedNetwork(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemGlobalDnsproxyWorkerCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalUrlFilterCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalHttpdMaxWorkerCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalProxyWorkerCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalScanunitCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalProxyHardwareAcceleration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalProxyKxpHardwareAcceleration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalProxyCipherHardwareAcceleration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalFgdAlertSubscription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpsecHmacOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpv6AcceptDad(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpv6AllowAnycastProbe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpv6AllowMulticastProbe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpv6AllowLocalInSilentDrop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpv6AllowLocalInSlientDrop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalCsrCaAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalWimax4GUsb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalCertChainMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSslvpnMaxWorkerCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSslvpnAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalVpnEmsSnCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSslvpnWebMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSslvpnEmsSnCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSslvpnKxpHardwareAcceleration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSslvpnCipherHardwareAcceleration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSslvpnPluginVersionCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTwoFactorFtkExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTwoFactorEmailExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTwoFactorSmsExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTwoFactorFacExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTwoFactorFtmExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalPerUserBal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalPerUserBwl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalVirtualServerCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalVirtualServerHardwareAcceleration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalWadWorkerCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalWadWorkerDevCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalWadCsvcCsCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalWadCsvcDbCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalWadSourceAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalWadMemoryChangeGranularity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalLoginTimestamp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpConflictDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalMiglogdChildren(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalLogDaemonCpuThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSpecialFile23Support(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalLogUuidPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalLogUuidAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalLogSslConnection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiRestApiCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalRestApiKeyUrlQuery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiCdnDomainOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiFortiguardResourceFetch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalArpMaxEntry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalHaAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalBfdAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalCmdbsvrAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAvAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalWadAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpsAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalMiglogAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSyslogAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalUrlFilterAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalRouterAffinity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalNdpMaxEntry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalBrFdbMaxEntry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalMaxRouteCacheSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpsecQatOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpsecRoundRobin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpsecAsicOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpsecSoftDecAsync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIkeEmbryonicLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDeviceIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalUserDeviceStoreMaxDevices(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalUserDeviceStoreMaxDeviceMem(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalUserDeviceStoreMaxUsers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalUserDeviceStoreMaxUnifiedMem(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDeviceIdentificationActiveScanDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalComplianceCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalComplianceCheckTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiDeviceLatitude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiDeviceLongitude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalPrivateDataEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAutoAuthExtensionDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiTheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiDateFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGuiDateTimeSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIgmpStateLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalCloudCommunication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalFecPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIpsecHaSeqjumpRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalFortitokenCloud(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalFortitokenCloudPushStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalFortitokenCloudRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalFortitokenCloudSyncInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalFazDiskBufferSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalIrqTimeAccounting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalFortiipamIntegration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalManagementIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalManagementPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalManagementPortUseAdminSport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalForticonverterIntegration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalForticonverterConfigUpload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalInternetServiceDatabase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalInternetServiceDownloadList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenSystemGlobalInternetServiceDownloadListId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemGlobalInternetServiceDownloadListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalGeoipFullDb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalEarlyTcpNpuSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalNpuNeighborUpdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalDelayTcpNpuSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalInterfaceSubnetUsage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalSflowdMaxChildrenNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalFortigslbIntegration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalUserHistoryPasswordThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAuthSessionAutoBackup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalAuthSessionAutoBackupInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalScimHttpsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalScimHttpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalScimServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalApplicationBandwidthTracking(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemGlobalTlsSessionCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemGlobal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("language", dataSourceFlattenSystemGlobalLanguage(o["language"], d, "language")); err != nil {
		if !fortiAPIPatch(o["language"]) {
			return fmt.Errorf("Error reading language: %v", err)
		}
	}

	if err = d.Set("gui_allow_incompatible_fabric_fgt", dataSourceFlattenSystemGlobalGuiAllowIncompatibleFabricFgt(o["gui-allow-incompatible-fabric-fgt"], d, "gui_allow_incompatible_fabric_fgt")); err != nil {
		if !fortiAPIPatch(o["gui-allow-incompatible-fabric-fgt"]) {
			return fmt.Errorf("Error reading gui_allow_incompatible_fabric_fgt: %v", err)
		}
	}

	if err = d.Set("gui_ipv6", dataSourceFlattenSystemGlobalGuiIpv6(o["gui-ipv6"], d, "gui_ipv6")); err != nil {
		if !fortiAPIPatch(o["gui-ipv6"]) {
			return fmt.Errorf("Error reading gui_ipv6: %v", err)
		}
	}

	if err = d.Set("gui_replacement_message_groups", dataSourceFlattenSystemGlobalGuiReplacementMessageGroups(o["gui-replacement-message-groups"], d, "gui_replacement_message_groups")); err != nil {
		if !fortiAPIPatch(o["gui-replacement-message-groups"]) {
			return fmt.Errorf("Error reading gui_replacement_message_groups: %v", err)
		}
	}

	if err = d.Set("gui_local_out", dataSourceFlattenSystemGlobalGuiLocalOut(o["gui-local-out"], d, "gui_local_out")); err != nil {
		if !fortiAPIPatch(o["gui-local-out"]) {
			return fmt.Errorf("Error reading gui_local_out: %v", err)
		}
	}

	if err = d.Set("gui_certificates", dataSourceFlattenSystemGlobalGuiCertificates(o["gui-certificates"], d, "gui_certificates")); err != nil {
		if !fortiAPIPatch(o["gui-certificates"]) {
			return fmt.Errorf("Error reading gui_certificates: %v", err)
		}
	}

	if err = d.Set("gui_custom_language", dataSourceFlattenSystemGlobalGuiCustomLanguage(o["gui-custom-language"], d, "gui_custom_language")); err != nil {
		if !fortiAPIPatch(o["gui-custom-language"]) {
			return fmt.Errorf("Error reading gui_custom_language: %v", err)
		}
	}

	if err = d.Set("gui_wireless_opensecurity", dataSourceFlattenSystemGlobalGuiWirelessOpensecurity(o["gui-wireless-opensecurity"], d, "gui_wireless_opensecurity")); err != nil {
		if !fortiAPIPatch(o["gui-wireless-opensecurity"]) {
			return fmt.Errorf("Error reading gui_wireless_opensecurity: %v", err)
		}
	}

	if err = d.Set("gui_app_detection_sdwan", dataSourceFlattenSystemGlobalGuiAppDetectionSdwan(o["gui-app-detection-sdwan"], d, "gui_app_detection_sdwan")); err != nil {
		if !fortiAPIPatch(o["gui-app-detection-sdwan"]) {
			return fmt.Errorf("Error reading gui_app_detection_sdwan: %v", err)
		}
	}

	if err = d.Set("gui_display_hostname", dataSourceFlattenSystemGlobalGuiDisplayHostname(o["gui-display-hostname"], d, "gui_display_hostname")); err != nil {
		if !fortiAPIPatch(o["gui-display-hostname"]) {
			return fmt.Errorf("Error reading gui_display_hostname: %v", err)
		}
	}

	if err = d.Set("gui_fortigate_cloud_sandbox", dataSourceFlattenSystemGlobalGuiFortigateCloudSandbox(o["gui-fortigate-cloud-sandbox"], d, "gui_fortigate_cloud_sandbox")); err != nil {
		if !fortiAPIPatch(o["gui-fortigate-cloud-sandbox"]) {
			return fmt.Errorf("Error reading gui_fortigate_cloud_sandbox: %v", err)
		}
	}

	if err = d.Set("gui_fortisandbox_cloud", dataSourceFlattenSystemGlobalGuiFortisandboxCloud(o["gui-fortisandbox-cloud"], d, "gui_fortisandbox_cloud")); err != nil {
		if !fortiAPIPatch(o["gui-fortisandbox-cloud"]) {
			return fmt.Errorf("Error reading gui_fortisandbox_cloud: %v", err)
		}
	}

	if err = d.Set("gui_firmware_upgrade_warning", dataSourceFlattenSystemGlobalGuiFirmwareUpgradeWarning(o["gui-firmware-upgrade-warning"], d, "gui_firmware_upgrade_warning")); err != nil {
		if !fortiAPIPatch(o["gui-firmware-upgrade-warning"]) {
			return fmt.Errorf("Error reading gui_firmware_upgrade_warning: %v", err)
		}
	}

	if err = d.Set("gui_firmware_upgrade_setup_warning", dataSourceFlattenSystemGlobalGuiFirmwareUpgradeSetupWarning(o["gui-firmware-upgrade-setup-warning"], d, "gui_firmware_upgrade_setup_warning")); err != nil {
		if !fortiAPIPatch(o["gui-firmware-upgrade-setup-warning"]) {
			return fmt.Errorf("Error reading gui_firmware_upgrade_setup_warning: %v", err)
		}
	}

	if err = d.Set("gui_lines_per_page", dataSourceFlattenSystemGlobalGuiLinesPerPage(o["gui-lines-per-page"], d, "gui_lines_per_page")); err != nil {
		if !fortiAPIPatch(o["gui-lines-per-page"]) {
			return fmt.Errorf("Error reading gui_lines_per_page: %v", err)
		}
	}

	if err = d.Set("admin_https_ssl_versions", dataSourceFlattenSystemGlobalAdminHttpsSslVersions(o["admin-https-ssl-versions"], d, "admin_https_ssl_versions")); err != nil {
		if !fortiAPIPatch(o["admin-https-ssl-versions"]) {
			return fmt.Errorf("Error reading admin_https_ssl_versions: %v", err)
		}
	}

	if err = d.Set("admin_https_ssl_ciphersuites", dataSourceFlattenSystemGlobalAdminHttpsSslCiphersuites(o["admin-https-ssl-ciphersuites"], d, "admin_https_ssl_ciphersuites")); err != nil {
		if !fortiAPIPatch(o["admin-https-ssl-ciphersuites"]) {
			return fmt.Errorf("Error reading admin_https_ssl_ciphersuites: %v", err)
		}
	}

	if err = d.Set("admin_https_ssl_banned_ciphers", dataSourceFlattenSystemGlobalAdminHttpsSslBannedCiphers(o["admin-https-ssl-banned-ciphers"], d, "admin_https_ssl_banned_ciphers")); err != nil {
		if !fortiAPIPatch(o["admin-https-ssl-banned-ciphers"]) {
			return fmt.Errorf("Error reading admin_https_ssl_banned_ciphers: %v", err)
		}
	}

	if err = d.Set("admintimeout", dataSourceFlattenSystemGlobalAdmintimeout(o["admintimeout"], d, "admintimeout")); err != nil {
		if !fortiAPIPatch(o["admintimeout"]) {
			return fmt.Errorf("Error reading admintimeout: %v", err)
		}
	}

	if err = d.Set("admin_console_timeout", dataSourceFlattenSystemGlobalAdminConsoleTimeout(o["admin-console-timeout"], d, "admin_console_timeout")); err != nil {
		if !fortiAPIPatch(o["admin-console-timeout"]) {
			return fmt.Errorf("Error reading admin_console_timeout: %v", err)
		}
	}

	if err = d.Set("ssd_trim_freq", dataSourceFlattenSystemGlobalSsdTrimFreq(o["ssd-trim-freq"], d, "ssd_trim_freq")); err != nil {
		if !fortiAPIPatch(o["ssd-trim-freq"]) {
			return fmt.Errorf("Error reading ssd_trim_freq: %v", err)
		}
	}

	if err = d.Set("ssd_trim_hour", dataSourceFlattenSystemGlobalSsdTrimHour(o["ssd-trim-hour"], d, "ssd_trim_hour")); err != nil {
		if !fortiAPIPatch(o["ssd-trim-hour"]) {
			return fmt.Errorf("Error reading ssd_trim_hour: %v", err)
		}
	}

	if err = d.Set("ssd_trim_min", dataSourceFlattenSystemGlobalSsdTrimMin(o["ssd-trim-min"], d, "ssd_trim_min")); err != nil {
		if !fortiAPIPatch(o["ssd-trim-min"]) {
			return fmt.Errorf("Error reading ssd_trim_min: %v", err)
		}
	}

	if err = d.Set("ssd_trim_weekday", dataSourceFlattenSystemGlobalSsdTrimWeekday(o["ssd-trim-weekday"], d, "ssd_trim_weekday")); err != nil {
		if !fortiAPIPatch(o["ssd-trim-weekday"]) {
			return fmt.Errorf("Error reading ssd_trim_weekday: %v", err)
		}
	}

	if err = d.Set("ssd_trim_date", dataSourceFlattenSystemGlobalSsdTrimDate(o["ssd-trim-date"], d, "ssd_trim_date")); err != nil {
		if !fortiAPIPatch(o["ssd-trim-date"]) {
			return fmt.Errorf("Error reading ssd_trim_date: %v", err)
		}
	}

	if err = d.Set("admin_concurrent", dataSourceFlattenSystemGlobalAdminConcurrent(o["admin-concurrent"], d, "admin_concurrent")); err != nil {
		if !fortiAPIPatch(o["admin-concurrent"]) {
			return fmt.Errorf("Error reading admin_concurrent: %v", err)
		}
	}

	if err = d.Set("admin_lockout_threshold", dataSourceFlattenSystemGlobalAdminLockoutThreshold(o["admin-lockout-threshold"], d, "admin_lockout_threshold")); err != nil {
		if !fortiAPIPatch(o["admin-lockout-threshold"]) {
			return fmt.Errorf("Error reading admin_lockout_threshold: %v", err)
		}
	}

	if err = d.Set("admin_lockout_duration", dataSourceFlattenSystemGlobalAdminLockoutDuration(o["admin-lockout-duration"], d, "admin_lockout_duration")); err != nil {
		if !fortiAPIPatch(o["admin-lockout-duration"]) {
			return fmt.Errorf("Error reading admin_lockout_duration: %v", err)
		}
	}

	if err = d.Set("refresh", dataSourceFlattenSystemGlobalRefresh(o["refresh"], d, "refresh")); err != nil {
		if !fortiAPIPatch(o["refresh"]) {
			return fmt.Errorf("Error reading refresh: %v", err)
		}
	}

	if err = d.Set("interval", dataSourceFlattenSystemGlobalInterval(o["interval"], d, "interval")); err != nil {
		if !fortiAPIPatch(o["interval"]) {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("failtime", dataSourceFlattenSystemGlobalFailtime(o["failtime"], d, "failtime")); err != nil {
		if !fortiAPIPatch(o["failtime"]) {
			return fmt.Errorf("Error reading failtime: %v", err)
		}
	}

	if err = d.Set("purdue_level", dataSourceFlattenSystemGlobalPurdueLevel(o["purdue-level"], d, "purdue_level")); err != nil {
		if !fortiAPIPatch(o["purdue-level"]) {
			return fmt.Errorf("Error reading purdue_level: %v", err)
		}
	}

	if err = d.Set("daily_restart", dataSourceFlattenSystemGlobalDailyRestart(o["daily-restart"], d, "daily_restart")); err != nil {
		if !fortiAPIPatch(o["daily-restart"]) {
			return fmt.Errorf("Error reading daily_restart: %v", err)
		}
	}

	if err = d.Set("restart_time", dataSourceFlattenSystemGlobalRestartTime(o["restart-time"], d, "restart_time")); err != nil {
		if !fortiAPIPatch(o["restart-time"]) {
			return fmt.Errorf("Error reading restart_time: %v", err)
		}
	}

	if err = d.Set("wad_restart_mode", dataSourceFlattenSystemGlobalWadRestartMode(o["wad-restart-mode"], d, "wad_restart_mode")); err != nil {
		if !fortiAPIPatch(o["wad-restart-mode"]) {
			return fmt.Errorf("Error reading wad_restart_mode: %v", err)
		}
	}

	if err = d.Set("wad_restart_start_time", dataSourceFlattenSystemGlobalWadRestartStartTime(o["wad-restart-start-time"], d, "wad_restart_start_time")); err != nil {
		if !fortiAPIPatch(o["wad-restart-start-time"]) {
			return fmt.Errorf("Error reading wad_restart_start_time: %v", err)
		}
	}

	if err = d.Set("wad_restart_end_time", dataSourceFlattenSystemGlobalWadRestartEndTime(o["wad-restart-end-time"], d, "wad_restart_end_time")); err != nil {
		if !fortiAPIPatch(o["wad-restart-end-time"]) {
			return fmt.Errorf("Error reading wad_restart_end_time: %v", err)
		}
	}

	if err = d.Set("wad_p2s_max_body_size", dataSourceFlattenSystemGlobalWadP2SMaxBodySize(o["wad-p2s-max-body-size"], d, "wad_p2s_max_body_size")); err != nil {
		if !fortiAPIPatch(o["wad-p2s-max-body-size"]) {
			return fmt.Errorf("Error reading wad_p2s_max_body_size: %v", err)
		}
	}

	if err = d.Set("radius_port", dataSourceFlattenSystemGlobalRadiusPort(o["radius-port"], d, "radius_port")); err != nil {
		if !fortiAPIPatch(o["radius-port"]) {
			return fmt.Errorf("Error reading radius_port: %v", err)
		}
	}

	if err = d.Set("speedtestd_server_port", dataSourceFlattenSystemGlobalSpeedtestdServerPort(o["speedtestd-server-port"], d, "speedtestd_server_port")); err != nil {
		if !fortiAPIPatch(o["speedtestd-server-port"]) {
			return fmt.Errorf("Error reading speedtestd_server_port: %v", err)
		}
	}

	if err = d.Set("speedtestd_ctrl_port", dataSourceFlattenSystemGlobalSpeedtestdCtrlPort(o["speedtestd-ctrl-port"], d, "speedtestd_ctrl_port")); err != nil {
		if !fortiAPIPatch(o["speedtestd-ctrl-port"]) {
			return fmt.Errorf("Error reading speedtestd_ctrl_port: %v", err)
		}
	}

	if err = d.Set("admin_login_max", dataSourceFlattenSystemGlobalAdminLoginMax(o["admin-login-max"], d, "admin_login_max")); err != nil {
		if !fortiAPIPatch(o["admin-login-max"]) {
			return fmt.Errorf("Error reading admin_login_max: %v", err)
		}
	}

	if err = d.Set("remoteauthtimeout", dataSourceFlattenSystemGlobalRemoteauthtimeout(o["remoteauthtimeout"], d, "remoteauthtimeout")); err != nil {
		if !fortiAPIPatch(o["remoteauthtimeout"]) {
			return fmt.Errorf("Error reading remoteauthtimeout: %v", err)
		}
	}

	if err = d.Set("ldapconntimeout", dataSourceFlattenSystemGlobalLdapconntimeout(o["ldapconntimeout"], d, "ldapconntimeout")); err != nil {
		if !fortiAPIPatch(o["ldapconntimeout"]) {
			return fmt.Errorf("Error reading ldapconntimeout: %v", err)
		}
	}

	if err = d.Set("batch_cmdb", dataSourceFlattenSystemGlobalBatchCmdb(o["batch-cmdb"], d, "batch_cmdb")); err != nil {
		if !fortiAPIPatch(o["batch-cmdb"]) {
			return fmt.Errorf("Error reading batch_cmdb: %v", err)
		}
	}

	if err = d.Set("max_dlpstat_memory", dataSourceFlattenSystemGlobalMaxDlpstatMemory(o["max-dlpstat-memory"], d, "max_dlpstat_memory")); err != nil {
		if !fortiAPIPatch(o["max-dlpstat-memory"]) {
			return fmt.Errorf("Error reading max_dlpstat_memory: %v", err)
		}
	}

	if err = d.Set("multi_factor_authentication", dataSourceFlattenSystemGlobalMultiFactorAuthentication(o["multi-factor-authentication"], d, "multi_factor_authentication")); err != nil {
		if !fortiAPIPatch(o["multi-factor-authentication"]) {
			return fmt.Errorf("Error reading multi_factor_authentication: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", dataSourceFlattenSystemGlobalSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-version"]) {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("autorun_log_fsck", dataSourceFlattenSystemGlobalAutorunLogFsck(o["autorun-log-fsck"], d, "autorun_log_fsck")); err != nil {
		if !fortiAPIPatch(o["autorun-log-fsck"]) {
			return fmt.Errorf("Error reading autorun_log_fsck: %v", err)
		}
	}

	if err = d.Set("dst", dataSourceFlattenSystemGlobalDst(o["dst"], d, "dst")); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("timezone", dataSourceFlattenSystemGlobalTimezone(o["timezone"], d, "timezone")); err != nil {
		if !fortiAPIPatch(o["timezone"]) {
			return fmt.Errorf("Error reading timezone: %v", err)
		}
	}

	if err = d.Set("traffic_priority", dataSourceFlattenSystemGlobalTrafficPriority(o["traffic-priority"], d, "traffic_priority")); err != nil {
		if !fortiAPIPatch(o["traffic-priority"]) {
			return fmt.Errorf("Error reading traffic_priority: %v", err)
		}
	}

	if err = d.Set("traffic_priority_level", dataSourceFlattenSystemGlobalTrafficPriorityLevel(o["traffic-priority-level"], d, "traffic_priority_level")); err != nil {
		if !fortiAPIPatch(o["traffic-priority-level"]) {
			return fmt.Errorf("Error reading traffic_priority_level: %v", err)
		}
	}

	if err = d.Set("quic_congestion_control_algo", dataSourceFlattenSystemGlobalQuicCongestionControlAlgo(o["quic-congestion-control-algo"], d, "quic_congestion_control_algo")); err != nil {
		if !fortiAPIPatch(o["quic-congestion-control-algo"]) {
			return fmt.Errorf("Error reading quic_congestion_control_algo: %v", err)
		}
	}

	if err = d.Set("quic_max_datagram_size", dataSourceFlattenSystemGlobalQuicMaxDatagramSize(o["quic-max-datagram-size"], d, "quic_max_datagram_size")); err != nil {
		if !fortiAPIPatch(o["quic-max-datagram-size"]) {
			return fmt.Errorf("Error reading quic_max_datagram_size: %v", err)
		}
	}

	if err = d.Set("quic_udp_payload_size_shaping_per_cid", dataSourceFlattenSystemGlobalQuicUdpPayloadSizeShapingPerCid(o["quic-udp-payload-size-shaping-per-cid"], d, "quic_udp_payload_size_shaping_per_cid")); err != nil {
		if !fortiAPIPatch(o["quic-udp-payload-size-shaping-per-cid"]) {
			return fmt.Errorf("Error reading quic_udp_payload_size_shaping_per_cid: %v", err)
		}
	}

	if err = d.Set("quic_ack_thresold", dataSourceFlattenSystemGlobalQuicAckThresold(o["quic-ack-thresold"], d, "quic_ack_thresold")); err != nil {
		if !fortiAPIPatch(o["quic-ack-thresold"]) {
			return fmt.Errorf("Error reading quic_ack_thresold: %v", err)
		}
	}

	if err = d.Set("quic_pmtud", dataSourceFlattenSystemGlobalQuicPmtud(o["quic-pmtud"], d, "quic_pmtud")); err != nil {
		if !fortiAPIPatch(o["quic-pmtud"]) {
			return fmt.Errorf("Error reading quic_pmtud: %v", err)
		}
	}

	if err = d.Set("quic_tls_handshake_timeout", dataSourceFlattenSystemGlobalQuicTlsHandshakeTimeout(o["quic-tls-handshake-timeout"], d, "quic_tls_handshake_timeout")); err != nil {
		if !fortiAPIPatch(o["quic-tls-handshake-timeout"]) {
			return fmt.Errorf("Error reading quic_tls_handshake_timeout: %v", err)
		}
	}

	if err = d.Set("anti_replay", dataSourceFlattenSystemGlobalAntiReplay(o["anti-replay"], d, "anti_replay")); err != nil {
		if !fortiAPIPatch(o["anti-replay"]) {
			return fmt.Errorf("Error reading anti_replay: %v", err)
		}
	}

	if err = d.Set("send_pmtu_icmp", dataSourceFlattenSystemGlobalSendPmtuIcmp(o["send-pmtu-icmp"], d, "send_pmtu_icmp")); err != nil {
		if !fortiAPIPatch(o["send-pmtu-icmp"]) {
			return fmt.Errorf("Error reading send_pmtu_icmp: %v", err)
		}
	}

	if err = d.Set("honor_df", dataSourceFlattenSystemGlobalHonorDf(o["honor-df"], d, "honor_df")); err != nil {
		if !fortiAPIPatch(o["honor-df"]) {
			return fmt.Errorf("Error reading honor_df: %v", err)
		}
	}

	if err = d.Set("pmtu_discovery", dataSourceFlattenSystemGlobalPmtuDiscovery(o["pmtu-discovery"], d, "pmtu_discovery")); err != nil {
		if !fortiAPIPatch(o["pmtu-discovery"]) {
			return fmt.Errorf("Error reading pmtu_discovery: %v", err)
		}
	}

	if err = d.Set("virtual_switch_vlan", dataSourceFlattenSystemGlobalVirtualSwitchVlan(o["virtual-switch-vlan"], d, "virtual_switch_vlan")); err != nil {
		if !fortiAPIPatch(o["virtual-switch-vlan"]) {
			return fmt.Errorf("Error reading virtual_switch_vlan: %v", err)
		}
	}

	if err = d.Set("split_port", dataSourceFlattenSystemGlobalSplitPort(o["split-port"], d, "split_port")); err != nil {
		if !fortiAPIPatch(o["split-port"]) {
			return fmt.Errorf("Error reading split_port: %v", err)
		}
	}

	if err = d.Set("revision_image_auto_backup", dataSourceFlattenSystemGlobalRevisionImageAutoBackup(o["revision-image-auto-backup"], d, "revision_image_auto_backup")); err != nil {
		if !fortiAPIPatch(o["revision-image-auto-backup"]) {
			return fmt.Errorf("Error reading revision_image_auto_backup: %v", err)
		}
	}

	if err = d.Set("revision_backup_on_logout", dataSourceFlattenSystemGlobalRevisionBackupOnLogout(o["revision-backup-on-logout"], d, "revision_backup_on_logout")); err != nil {
		if !fortiAPIPatch(o["revision-backup-on-logout"]) {
			return fmt.Errorf("Error reading revision_backup_on_logout: %v", err)
		}
	}

	if err = d.Set("management_vdom", dataSourceFlattenSystemGlobalManagementVdom(o["management-vdom"], d, "management_vdom")); err != nil {
		if !fortiAPIPatch(o["management-vdom"]) {
			return fmt.Errorf("Error reading management_vdom: %v", err)
		}
	}

	if err = d.Set("hostname", dataSourceFlattenSystemGlobalHostname(o["hostname"], d, "hostname")); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("gui_allow_default_hostname", dataSourceFlattenSystemGlobalGuiAllowDefaultHostname(o["gui-allow-default-hostname"], d, "gui_allow_default_hostname")); err != nil {
		if !fortiAPIPatch(o["gui-allow-default-hostname"]) {
			return fmt.Errorf("Error reading gui_allow_default_hostname: %v", err)
		}
	}

	if err = d.Set("gui_forticare_registration_setup_warning", dataSourceFlattenSystemGlobalGuiForticareRegistrationSetupWarning(o["gui-forticare-registration-setup-warning"], d, "gui_forticare_registration_setup_warning")); err != nil {
		if !fortiAPIPatch(o["gui-forticare-registration-setup-warning"]) {
			return fmt.Errorf("Error reading gui_forticare_registration_setup_warning: %v", err)
		}
	}

	if err = d.Set("gui_auto_upgrade_setup_warning", dataSourceFlattenSystemGlobalGuiAutoUpgradeSetupWarning(o["gui-auto-upgrade-setup-warning"], d, "gui_auto_upgrade_setup_warning")); err != nil {
		if !fortiAPIPatch(o["gui-auto-upgrade-setup-warning"]) {
			return fmt.Errorf("Error reading gui_auto_upgrade_setup_warning: %v", err)
		}
	}

	if err = d.Set("gui_workflow_management", dataSourceFlattenSystemGlobalGuiWorkflowManagement(o["gui-workflow-management"], d, "gui_workflow_management")); err != nil {
		if !fortiAPIPatch(o["gui-workflow-management"]) {
			return fmt.Errorf("Error reading gui_workflow_management: %v", err)
		}
	}

	if err = d.Set("gui_cdn_usage", dataSourceFlattenSystemGlobalGuiCdnUsage(o["gui-cdn-usage"], d, "gui_cdn_usage")); err != nil {
		if !fortiAPIPatch(o["gui-cdn-usage"]) {
			return fmt.Errorf("Error reading gui_cdn_usage: %v", err)
		}
	}

	if err = d.Set("alias", dataSourceFlattenSystemGlobalAlias(o["alias"], d, "alias")); err != nil {
		if !fortiAPIPatch(o["alias"]) {
			return fmt.Errorf("Error reading alias: %v", err)
		}
	}

	if err = d.Set("strong_crypto", dataSourceFlattenSystemGlobalStrongCrypto(o["strong-crypto"], d, "strong_crypto")); err != nil {
		if !fortiAPIPatch(o["strong-crypto"]) {
			return fmt.Errorf("Error reading strong_crypto: %v", err)
		}
	}

	if err = d.Set("ssh_cbc_cipher", dataSourceFlattenSystemGlobalSshCbcCipher(o["ssh-cbc-cipher"], d, "ssh_cbc_cipher")); err != nil {
		if !fortiAPIPatch(o["ssh-cbc-cipher"]) {
			return fmt.Errorf("Error reading ssh_cbc_cipher: %v", err)
		}
	}

	if err = d.Set("ssh_hmac_md5", dataSourceFlattenSystemGlobalSshHmacMd5(o["ssh-hmac-md5"], d, "ssh_hmac_md5")); err != nil {
		if !fortiAPIPatch(o["ssh-hmac-md5"]) {
			return fmt.Errorf("Error reading ssh_hmac_md5: %v", err)
		}
	}

	if err = d.Set("ssh_kex_sha1", dataSourceFlattenSystemGlobalSshKexSha1(o["ssh-kex-sha1"], d, "ssh_kex_sha1")); err != nil {
		if !fortiAPIPatch(o["ssh-kex-sha1"]) {
			return fmt.Errorf("Error reading ssh_kex_sha1: %v", err)
		}
	}

	if err = d.Set("ssh_mac_weak", dataSourceFlattenSystemGlobalSshMacWeak(o["ssh-mac-weak"], d, "ssh_mac_weak")); err != nil {
		if !fortiAPIPatch(o["ssh-mac-weak"]) {
			return fmt.Errorf("Error reading ssh_mac_weak: %v", err)
		}
	}

	if err = d.Set("ssl_static_key_ciphers", dataSourceFlattenSystemGlobalSslStaticKeyCiphers(o["ssl-static-key-ciphers"], d, "ssl_static_key_ciphers")); err != nil {
		if !fortiAPIPatch(o["ssl-static-key-ciphers"]) {
			return fmt.Errorf("Error reading ssl_static_key_ciphers: %v", err)
		}
	}

	if err = d.Set("ssh_kex_algo", dataSourceFlattenSystemGlobalSshKexAlgo(o["ssh-kex-algo"], d, "ssh_kex_algo")); err != nil {
		if !fortiAPIPatch(o["ssh-kex-algo"]) {
			return fmt.Errorf("Error reading ssh_kex_algo: %v", err)
		}
	}

	if err = d.Set("ssh_enc_algo", dataSourceFlattenSystemGlobalSshEncAlgo(o["ssh-enc-algo"], d, "ssh_enc_algo")); err != nil {
		if !fortiAPIPatch(o["ssh-enc-algo"]) {
			return fmt.Errorf("Error reading ssh_enc_algo: %v", err)
		}
	}

	if err = d.Set("ssh_mac_algo", dataSourceFlattenSystemGlobalSshMacAlgo(o["ssh-mac-algo"], d, "ssh_mac_algo")); err != nil {
		if !fortiAPIPatch(o["ssh-mac-algo"]) {
			return fmt.Errorf("Error reading ssh_mac_algo: %v", err)
		}
	}

	if err = d.Set("ssh_hostkey_algo", dataSourceFlattenSystemGlobalSshHostkeyAlgo(o["ssh-hostkey-algo"], d, "ssh_hostkey_algo")); err != nil {
		if !fortiAPIPatch(o["ssh-hostkey-algo"]) {
			return fmt.Errorf("Error reading ssh_hostkey_algo: %v", err)
		}
	}

	if err = d.Set("ssh_hostkey_override", dataSourceFlattenSystemGlobalSshHostkeyOverride(o["ssh-hostkey-override"], d, "ssh_hostkey_override")); err != nil {
		if !fortiAPIPatch(o["ssh-hostkey-override"]) {
			return fmt.Errorf("Error reading ssh_hostkey_override: %v", err)
		}
	}

	if err = d.Set("ssh_hostkey", dataSourceFlattenSystemGlobalSshHostkey(o["ssh-hostkey"], d, "ssh_hostkey")); err != nil {
		if !fortiAPIPatch(o["ssh-hostkey"]) {
			return fmt.Errorf("Error reading ssh_hostkey: %v", err)
		}
	}

	if err = d.Set("snat_route_change", dataSourceFlattenSystemGlobalSnatRouteChange(o["snat-route-change"], d, "snat_route_change")); err != nil {
		if !fortiAPIPatch(o["snat-route-change"]) {
			return fmt.Errorf("Error reading snat_route_change: %v", err)
		}
	}

	if err = d.Set("ipv6_snat_route_change", dataSourceFlattenSystemGlobalIpv6SnatRouteChange(o["ipv6-snat-route-change"], d, "ipv6_snat_route_change")); err != nil {
		if !fortiAPIPatch(o["ipv6-snat-route-change"]) {
			return fmt.Errorf("Error reading ipv6_snat_route_change: %v", err)
		}
	}

	if err = d.Set("speedtest_server", dataSourceFlattenSystemGlobalSpeedtestServer(o["speedtest-server"], d, "speedtest_server")); err != nil {
		if !fortiAPIPatch(o["speedtest-server"]) {
			return fmt.Errorf("Error reading speedtest_server: %v", err)
		}
	}

	if err = d.Set("cli_audit_log", dataSourceFlattenSystemGlobalCliAuditLog(o["cli-audit-log"], d, "cli_audit_log")); err != nil {
		if !fortiAPIPatch(o["cli-audit-log"]) {
			return fmt.Errorf("Error reading cli_audit_log: %v", err)
		}
	}

	if err = d.Set("dh_params", dataSourceFlattenSystemGlobalDhParams(o["dh-params"], d, "dh_params")); err != nil {
		if !fortiAPIPatch(o["dh-params"]) {
			return fmt.Errorf("Error reading dh_params: %v", err)
		}
	}

	if err = d.Set("fds_statistics", dataSourceFlattenSystemGlobalFdsStatistics(o["fds-statistics"], d, "fds_statistics")); err != nil {
		if !fortiAPIPatch(o["fds-statistics"]) {
			return fmt.Errorf("Error reading fds_statistics: %v", err)
		}
	}

	if err = d.Set("fds_statistics_period", dataSourceFlattenSystemGlobalFdsStatisticsPeriod(o["fds-statistics-period"], d, "fds_statistics_period")); err != nil {
		if !fortiAPIPatch(o["fds-statistics-period"]) {
			return fmt.Errorf("Error reading fds_statistics_period: %v", err)
		}
	}

	if err = d.Set("multicast_forward", dataSourceFlattenSystemGlobalMulticastForward(o["multicast-forward"], d, "multicast_forward")); err != nil {
		if !fortiAPIPatch(o["multicast-forward"]) {
			return fmt.Errorf("Error reading multicast_forward: %v", err)
		}
	}

	if err = d.Set("mc_ttl_notchange", dataSourceFlattenSystemGlobalMcTtlNotchange(o["mc-ttl-notchange"], d, "mc_ttl_notchange")); err != nil {
		if !fortiAPIPatch(o["mc-ttl-notchange"]) {
			return fmt.Errorf("Error reading mc_ttl_notchange: %v", err)
		}
	}

	if err = d.Set("asymroute", dataSourceFlattenSystemGlobalAsymroute(o["asymroute"], d, "asymroute")); err != nil {
		if !fortiAPIPatch(o["asymroute"]) {
			return fmt.Errorf("Error reading asymroute: %v", err)
		}
	}

	if err = d.Set("tcp_option", dataSourceFlattenSystemGlobalTcpOption(o["tcp-option"], d, "tcp_option")); err != nil {
		if !fortiAPIPatch(o["tcp-option"]) {
			return fmt.Errorf("Error reading tcp_option: %v", err)
		}
	}

	if err = d.Set("lldp_transmission", dataSourceFlattenSystemGlobalLldpTransmission(o["lldp-transmission"], d, "lldp_transmission")); err != nil {
		if !fortiAPIPatch(o["lldp-transmission"]) {
			return fmt.Errorf("Error reading lldp_transmission: %v", err)
		}
	}

	if err = d.Set("lldp_reception", dataSourceFlattenSystemGlobalLldpReception(o["lldp-reception"], d, "lldp_reception")); err != nil {
		if !fortiAPIPatch(o["lldp-reception"]) {
			return fmt.Errorf("Error reading lldp_reception: %v", err)
		}
	}

	if err = d.Set("proxy_auth_timeout", dataSourceFlattenSystemGlobalProxyAuthTimeout(o["proxy-auth-timeout"], d, "proxy_auth_timeout")); err != nil {
		if !fortiAPIPatch(o["proxy-auth-timeout"]) {
			return fmt.Errorf("Error reading proxy_auth_timeout: %v", err)
		}
	}

	if err = d.Set("proxy_keep_alive_mode", dataSourceFlattenSystemGlobalProxyKeepAliveMode(o["proxy-keep-alive-mode"], d, "proxy_keep_alive_mode")); err != nil {
		if !fortiAPIPatch(o["proxy-keep-alive-mode"]) {
			return fmt.Errorf("Error reading proxy_keep_alive_mode: %v", err)
		}
	}

	if err = d.Set("proxy_re_authentication_time", dataSourceFlattenSystemGlobalProxyReAuthenticationTime(o["proxy-re-authentication-time"], d, "proxy_re_authentication_time")); err != nil {
		if !fortiAPIPatch(o["proxy-re-authentication-time"]) {
			return fmt.Errorf("Error reading proxy_re_authentication_time: %v", err)
		}
	}

	if err = d.Set("proxy_re_authentication_mode", dataSourceFlattenSystemGlobalProxyReAuthenticationMode(o["proxy-re-authentication-mode"], d, "proxy_re_authentication_mode")); err != nil {
		if !fortiAPIPatch(o["proxy-re-authentication-mode"]) {
			return fmt.Errorf("Error reading proxy_re_authentication_mode: %v", err)
		}
	}

	if err = d.Set("proxy_auth_lifetime", dataSourceFlattenSystemGlobalProxyAuthLifetime(o["proxy-auth-lifetime"], d, "proxy_auth_lifetime")); err != nil {
		if !fortiAPIPatch(o["proxy-auth-lifetime"]) {
			return fmt.Errorf("Error reading proxy_auth_lifetime: %v", err)
		}
	}

	if err = d.Set("proxy_auth_lifetime_timeout", dataSourceFlattenSystemGlobalProxyAuthLifetimeTimeout(o["proxy-auth-lifetime-timeout"], d, "proxy_auth_lifetime_timeout")); err != nil {
		if !fortiAPIPatch(o["proxy-auth-lifetime-timeout"]) {
			return fmt.Errorf("Error reading proxy_auth_lifetime_timeout: %v", err)
		}
	}

	if err = d.Set("proxy_resource_mode", dataSourceFlattenSystemGlobalProxyResourceMode(o["proxy-resource-mode"], d, "proxy_resource_mode")); err != nil {
		if !fortiAPIPatch(o["proxy-resource-mode"]) {
			return fmt.Errorf("Error reading proxy_resource_mode: %v", err)
		}
	}

	if err = d.Set("proxy_cert_use_mgmt_vdom", dataSourceFlattenSystemGlobalProxyCertUseMgmtVdom(o["proxy-cert-use-mgmt-vdom"], d, "proxy_cert_use_mgmt_vdom")); err != nil {
		if !fortiAPIPatch(o["proxy-cert-use-mgmt-vdom"]) {
			return fmt.Errorf("Error reading proxy_cert_use_mgmt_vdom: %v", err)
		}
	}

	if err = d.Set("sys_perf_log_interval", dataSourceFlattenSystemGlobalSysPerfLogInterval(o["sys-perf-log-interval"], d, "sys_perf_log_interval")); err != nil {
		if !fortiAPIPatch(o["sys-perf-log-interval"]) {
			return fmt.Errorf("Error reading sys_perf_log_interval: %v", err)
		}
	}

	if err = d.Set("check_protocol_header", dataSourceFlattenSystemGlobalCheckProtocolHeader(o["check-protocol-header"], d, "check_protocol_header")); err != nil {
		if !fortiAPIPatch(o["check-protocol-header"]) {
			return fmt.Errorf("Error reading check_protocol_header: %v", err)
		}
	}

	if err = d.Set("vip_arp_range", dataSourceFlattenSystemGlobalVipArpRange(o["vip-arp-range"], d, "vip_arp_range")); err != nil {
		if !fortiAPIPatch(o["vip-arp-range"]) {
			return fmt.Errorf("Error reading vip_arp_range: %v", err)
		}
	}

	if err = d.Set("reset_sessionless_tcp", dataSourceFlattenSystemGlobalResetSessionlessTcp(o["reset-sessionless-tcp"], d, "reset_sessionless_tcp")); err != nil {
		if !fortiAPIPatch(o["reset-sessionless-tcp"]) {
			return fmt.Errorf("Error reading reset_sessionless_tcp: %v", err)
		}
	}

	if err = d.Set("allow_traffic_redirect", dataSourceFlattenSystemGlobalAllowTrafficRedirect(o["allow-traffic-redirect"], d, "allow_traffic_redirect")); err != nil {
		if !fortiAPIPatch(o["allow-traffic-redirect"]) {
			return fmt.Errorf("Error reading allow_traffic_redirect: %v", err)
		}
	}

	if err = d.Set("ipv6_allow_traffic_redirect", dataSourceFlattenSystemGlobalIpv6AllowTrafficRedirect(o["ipv6-allow-traffic-redirect"], d, "ipv6_allow_traffic_redirect")); err != nil {
		if !fortiAPIPatch(o["ipv6-allow-traffic-redirect"]) {
			return fmt.Errorf("Error reading ipv6_allow_traffic_redirect: %v", err)
		}
	}

	if err = d.Set("strict_dirty_session_check", dataSourceFlattenSystemGlobalStrictDirtySessionCheck(o["strict-dirty-session-check"], d, "strict_dirty_session_check")); err != nil {
		if !fortiAPIPatch(o["strict-dirty-session-check"]) {
			return fmt.Errorf("Error reading strict_dirty_session_check: %v", err)
		}
	}

	if err = d.Set("tcp_halfclose_timer", dataSourceFlattenSystemGlobalTcpHalfcloseTimer(o["tcp-halfclose-timer"], d, "tcp_halfclose_timer")); err != nil {
		if !fortiAPIPatch(o["tcp-halfclose-timer"]) {
			return fmt.Errorf("Error reading tcp_halfclose_timer: %v", err)
		}
	}

	if err = d.Set("tcp_halfopen_timer", dataSourceFlattenSystemGlobalTcpHalfopenTimer(o["tcp-halfopen-timer"], d, "tcp_halfopen_timer")); err != nil {
		if !fortiAPIPatch(o["tcp-halfopen-timer"]) {
			return fmt.Errorf("Error reading tcp_halfopen_timer: %v", err)
		}
	}

	if err = d.Set("tcp_timewait_timer", dataSourceFlattenSystemGlobalTcpTimewaitTimer(o["tcp-timewait-timer"], d, "tcp_timewait_timer")); err != nil {
		if !fortiAPIPatch(o["tcp-timewait-timer"]) {
			return fmt.Errorf("Error reading tcp_timewait_timer: %v", err)
		}
	}

	if err = d.Set("tcp_rst_timer", dataSourceFlattenSystemGlobalTcpRstTimer(o["tcp-rst-timer"], d, "tcp_rst_timer")); err != nil {
		if !fortiAPIPatch(o["tcp-rst-timer"]) {
			return fmt.Errorf("Error reading tcp_rst_timer: %v", err)
		}
	}

	if err = d.Set("udp_idle_timer", dataSourceFlattenSystemGlobalUdpIdleTimer(o["udp-idle-timer"], d, "udp_idle_timer")); err != nil {
		if !fortiAPIPatch(o["udp-idle-timer"]) {
			return fmt.Errorf("Error reading udp_idle_timer: %v", err)
		}
	}

	if err = d.Set("block_session_timer", dataSourceFlattenSystemGlobalBlockSessionTimer(o["block-session-timer"], d, "block_session_timer")); err != nil {
		if !fortiAPIPatch(o["block-session-timer"]) {
			return fmt.Errorf("Error reading block_session_timer: %v", err)
		}
	}

	if err = d.Set("ip_src_port_range", dataSourceFlattenSystemGlobalIpSrcPortRange(o["ip-src-port-range"], d, "ip_src_port_range")); err != nil {
		if !fortiAPIPatch(o["ip-src-port-range"]) {
			return fmt.Errorf("Error reading ip_src_port_range: %v", err)
		}
	}

	if err = d.Set("pre_login_banner", dataSourceFlattenSystemGlobalPreLoginBanner(o["pre-login-banner"], d, "pre_login_banner")); err != nil {
		if !fortiAPIPatch(o["pre-login-banner"]) {
			return fmt.Errorf("Error reading pre_login_banner: %v", err)
		}
	}

	if err = d.Set("post_login_banner", dataSourceFlattenSystemGlobalPostLoginBanner(o["post-login-banner"], d, "post_login_banner")); err != nil {
		if !fortiAPIPatch(o["post-login-banner"]) {
			return fmt.Errorf("Error reading post_login_banner: %v", err)
		}
	}

	if err = d.Set("tftp", dataSourceFlattenSystemGlobalTftp(o["tftp"], d, "tftp")); err != nil {
		if !fortiAPIPatch(o["tftp"]) {
			return fmt.Errorf("Error reading tftp: %v", err)
		}
	}

	if err = d.Set("av_failopen", dataSourceFlattenSystemGlobalAvFailopen(o["av-failopen"], d, "av_failopen")); err != nil {
		if !fortiAPIPatch(o["av-failopen"]) {
			return fmt.Errorf("Error reading av_failopen: %v", err)
		}
	}

	if err = d.Set("av_failopen_session", dataSourceFlattenSystemGlobalAvFailopenSession(o["av-failopen-session"], d, "av_failopen_session")); err != nil {
		if !fortiAPIPatch(o["av-failopen-session"]) {
			return fmt.Errorf("Error reading av_failopen_session: %v", err)
		}
	}

	if err = d.Set("memory_use_threshold_extreme", dataSourceFlattenSystemGlobalMemoryUseThresholdExtreme(o["memory-use-threshold-extreme"], d, "memory_use_threshold_extreme")); err != nil {
		if !fortiAPIPatch(o["memory-use-threshold-extreme"]) {
			return fmt.Errorf("Error reading memory_use_threshold_extreme: %v", err)
		}
	}

	if err = d.Set("memory_use_threshold_red", dataSourceFlattenSystemGlobalMemoryUseThresholdRed(o["memory-use-threshold-red"], d, "memory_use_threshold_red")); err != nil {
		if !fortiAPIPatch(o["memory-use-threshold-red"]) {
			return fmt.Errorf("Error reading memory_use_threshold_red: %v", err)
		}
	}

	if err = d.Set("memory_use_threshold_green", dataSourceFlattenSystemGlobalMemoryUseThresholdGreen(o["memory-use-threshold-green"], d, "memory_use_threshold_green")); err != nil {
		if !fortiAPIPatch(o["memory-use-threshold-green"]) {
			return fmt.Errorf("Error reading memory_use_threshold_green: %v", err)
		}
	}

	if err = d.Set("ip_fragment_mem_thresholds", dataSourceFlattenSystemGlobalIpFragmentMemThresholds(o["ip-fragment-mem-thresholds"], d, "ip_fragment_mem_thresholds")); err != nil {
		if !fortiAPIPatch(o["ip-fragment-mem-thresholds"]) {
			return fmt.Errorf("Error reading ip_fragment_mem_thresholds: %v", err)
		}
	}

	if err = d.Set("ip_fragment_timeout", dataSourceFlattenSystemGlobalIpFragmentTimeout(o["ip-fragment-timeout"], d, "ip_fragment_timeout")); err != nil {
		if !fortiAPIPatch(o["ip-fragment-timeout"]) {
			return fmt.Errorf("Error reading ip_fragment_timeout: %v", err)
		}
	}

	if err = d.Set("ipv6_fragment_timeout", dataSourceFlattenSystemGlobalIpv6FragmentTimeout(o["ipv6-fragment-timeout"], d, "ipv6_fragment_timeout")); err != nil {
		if !fortiAPIPatch(o["ipv6-fragment-timeout"]) {
			return fmt.Errorf("Error reading ipv6_fragment_timeout: %v", err)
		}
	}

	if err = d.Set("cpu_use_threshold", dataSourceFlattenSystemGlobalCpuUseThreshold(o["cpu-use-threshold"], d, "cpu_use_threshold")); err != nil {
		if !fortiAPIPatch(o["cpu-use-threshold"]) {
			return fmt.Errorf("Error reading cpu_use_threshold: %v", err)
		}
	}

	if err = d.Set("log_single_cpu_high", dataSourceFlattenSystemGlobalLogSingleCpuHigh(o["log-single-cpu-high"], d, "log_single_cpu_high")); err != nil {
		if !fortiAPIPatch(o["log-single-cpu-high"]) {
			return fmt.Errorf("Error reading log_single_cpu_high: %v", err)
		}
	}

	if err = d.Set("check_reset_range", dataSourceFlattenSystemGlobalCheckResetRange(o["check-reset-range"], d, "check_reset_range")); err != nil {
		if !fortiAPIPatch(o["check-reset-range"]) {
			return fmt.Errorf("Error reading check_reset_range: %v", err)
		}
	}

	if err = d.Set("upgrade_report", dataSourceFlattenSystemGlobalUpgradeReport(o["upgrade-report"], d, "upgrade_report")); err != nil {
		if !fortiAPIPatch(o["upgrade-report"]) {
			return fmt.Errorf("Error reading upgrade_report: %v", err)
		}
	}

	if err = d.Set("vdom_mode", dataSourceFlattenSystemGlobalVdomMode(o["vdom-mode"], d, "vdom_mode")); err != nil {
		if !fortiAPIPatch(o["vdom-mode"]) {
			return fmt.Errorf("Error reading vdom_mode: %v", err)
		}
	}

	if err = d.Set("vdom_admin", dataSourceFlattenSystemGlobalVdomAdmin(o["vdom-admin"], d, "vdom_admin")); err != nil {
		if !fortiAPIPatch(o["vdom-admin"]) {
			return fmt.Errorf("Error reading vdom_admin: %v", err)
		}
	}

	if err = d.Set("long_vdom_name", dataSourceFlattenSystemGlobalLongVdomName(o["long-vdom-name"], d, "long_vdom_name")); err != nil {
		if !fortiAPIPatch(o["long-vdom-name"]) {
			return fmt.Errorf("Error reading long_vdom_name: %v", err)
		}
	}

	if err = d.Set("edit_vdom_prompt", dataSourceFlattenSystemGlobalEditVdomPrompt(o["edit-vdom-prompt"], d, "edit_vdom_prompt")); err != nil {
		if !fortiAPIPatch(o["edit-vdom-prompt"]) {
			return fmt.Errorf("Error reading edit_vdom_prompt: %v", err)
		}
	}

	if err = d.Set("admin_port", dataSourceFlattenSystemGlobalAdminPort(o["admin-port"], d, "admin_port")); err != nil {
		if !fortiAPIPatch(o["admin-port"]) {
			return fmt.Errorf("Error reading admin_port: %v", err)
		}
	}

	if err = d.Set("admin_sport", dataSourceFlattenSystemGlobalAdminSport(o["admin-sport"], d, "admin_sport")); err != nil {
		if !fortiAPIPatch(o["admin-sport"]) {
			return fmt.Errorf("Error reading admin_sport: %v", err)
		}
	}

	if err = d.Set("admin_host", dataSourceFlattenSystemGlobalAdminHost(o["admin-host"], d, "admin_host")); err != nil {
		if !fortiAPIPatch(o["admin-host"]) {
			return fmt.Errorf("Error reading admin_host: %v", err)
		}
	}

	if err = d.Set("admin_https_redirect", dataSourceFlattenSystemGlobalAdminHttpsRedirect(o["admin-https-redirect"], d, "admin_https_redirect")); err != nil {
		if !fortiAPIPatch(o["admin-https-redirect"]) {
			return fmt.Errorf("Error reading admin_https_redirect: %v", err)
		}
	}

	if err = d.Set("admin_hsts_max_age", dataSourceFlattenSystemGlobalAdminHstsMaxAge(o["admin-hsts-max-age"], d, "admin_hsts_max_age")); err != nil {
		if !fortiAPIPatch(o["admin-hsts-max-age"]) {
			return fmt.Errorf("Error reading admin_hsts_max_age: %v", err)
		}
	}

	if err = d.Set("admin_ssh_password", dataSourceFlattenSystemGlobalAdminSshPassword(o["admin-ssh-password"], d, "admin_ssh_password")); err != nil {
		if !fortiAPIPatch(o["admin-ssh-password"]) {
			return fmt.Errorf("Error reading admin_ssh_password: %v", err)
		}
	}

	if err = d.Set("admin_restrict_local", dataSourceFlattenSystemGlobalAdminRestrictLocal(o["admin-restrict-local"], d, "admin_restrict_local")); err != nil {
		if !fortiAPIPatch(o["admin-restrict-local"]) {
			return fmt.Errorf("Error reading admin_restrict_local: %v", err)
		}
	}

	if err = d.Set("admin_ssh_port", dataSourceFlattenSystemGlobalAdminSshPort(o["admin-ssh-port"], d, "admin_ssh_port")); err != nil {
		if !fortiAPIPatch(o["admin-ssh-port"]) {
			return fmt.Errorf("Error reading admin_ssh_port: %v", err)
		}
	}

	if err = d.Set("admin_ssh_grace_time", dataSourceFlattenSystemGlobalAdminSshGraceTime(o["admin-ssh-grace-time"], d, "admin_ssh_grace_time")); err != nil {
		if !fortiAPIPatch(o["admin-ssh-grace-time"]) {
			return fmt.Errorf("Error reading admin_ssh_grace_time: %v", err)
		}
	}

	if err = d.Set("admin_ssh_v1", dataSourceFlattenSystemGlobalAdminSshV1(o["admin-ssh-v1"], d, "admin_ssh_v1")); err != nil {
		if !fortiAPIPatch(o["admin-ssh-v1"]) {
			return fmt.Errorf("Error reading admin_ssh_v1: %v", err)
		}
	}

	if err = d.Set("admin_telnet", dataSourceFlattenSystemGlobalAdminTelnet(o["admin-telnet"], d, "admin_telnet")); err != nil {
		if !fortiAPIPatch(o["admin-telnet"]) {
			return fmt.Errorf("Error reading admin_telnet: %v", err)
		}
	}

	if err = d.Set("admin_telnet_port", dataSourceFlattenSystemGlobalAdminTelnetPort(o["admin-telnet-port"], d, "admin_telnet_port")); err != nil {
		if !fortiAPIPatch(o["admin-telnet-port"]) {
			return fmt.Errorf("Error reading admin_telnet_port: %v", err)
		}
	}

	if err = d.Set("admin_forticloud_sso_login", dataSourceFlattenSystemGlobalAdminForticloudSsoLogin(o["admin-forticloud-sso-login"], d, "admin_forticloud_sso_login")); err != nil {
		if !fortiAPIPatch(o["admin-forticloud-sso-login"]) {
			return fmt.Errorf("Error reading admin_forticloud_sso_login: %v", err)
		}
	}

	if err = d.Set("admin_forticloud_sso_default_profile", dataSourceFlattenSystemGlobalAdminForticloudSsoDefaultProfile(o["admin-forticloud-sso-default-profile"], d, "admin_forticloud_sso_default_profile")); err != nil {
		if !fortiAPIPatch(o["admin-forticloud-sso-default-profile"]) {
			return fmt.Errorf("Error reading admin_forticloud_sso_default_profile: %v", err)
		}
	}

	if err = d.Set("default_service_source_port", dataSourceFlattenSystemGlobalDefaultServiceSourcePort(o["default-service-source-port"], d, "default_service_source_port")); err != nil {
		if !fortiAPIPatch(o["default-service-source-port"]) {
			return fmt.Errorf("Error reading default_service_source_port: %v", err)
		}
	}

	if err = d.Set("admin_maintainer", dataSourceFlattenSystemGlobalAdminMaintainer(o["admin-maintainer"], d, "admin_maintainer")); err != nil {
		if !fortiAPIPatch(o["admin-maintainer"]) {
			return fmt.Errorf("Error reading admin_maintainer: %v", err)
		}
	}

	if err = d.Set("admin_server_cert", dataSourceFlattenSystemGlobalAdminServerCert(o["admin-server-cert"], d, "admin_server_cert")); err != nil {
		if !fortiAPIPatch(o["admin-server-cert"]) {
			return fmt.Errorf("Error reading admin_server_cert: %v", err)
		}
	}

	if err = d.Set("user_server_cert", dataSourceFlattenSystemGlobalUserServerCert(o["user-server-cert"], d, "user_server_cert")); err != nil {
		if !fortiAPIPatch(o["user-server-cert"]) {
			return fmt.Errorf("Error reading user_server_cert: %v", err)
		}
	}

	if err = d.Set("admin_https_pki_required", dataSourceFlattenSystemGlobalAdminHttpsPkiRequired(o["admin-https-pki-required"], d, "admin_https_pki_required")); err != nil {
		if !fortiAPIPatch(o["admin-https-pki-required"]) {
			return fmt.Errorf("Error reading admin_https_pki_required: %v", err)
		}
	}

	if err = d.Set("wifi_certificate", dataSourceFlattenSystemGlobalWifiCertificate(o["wifi-certificate"], d, "wifi_certificate")); err != nil {
		if !fortiAPIPatch(o["wifi-certificate"]) {
			return fmt.Errorf("Error reading wifi_certificate: %v", err)
		}
	}

	if err = d.Set("dhcp_lease_backup_interval", dataSourceFlattenSystemGlobalDhcpLeaseBackupInterval(o["dhcp-lease-backup-interval"], d, "dhcp_lease_backup_interval")); err != nil {
		if !fortiAPIPatch(o["dhcp-lease-backup-interval"]) {
			return fmt.Errorf("Error reading dhcp_lease_backup_interval: %v", err)
		}
	}

	if err = d.Set("wifi_ca_certificate", dataSourceFlattenSystemGlobalWifiCaCertificate(o["wifi-ca-certificate"], d, "wifi_ca_certificate")); err != nil {
		if !fortiAPIPatch(o["wifi-ca-certificate"]) {
			return fmt.Errorf("Error reading wifi_ca_certificate: %v", err)
		}
	}

	if err = d.Set("auth_http_port", dataSourceFlattenSystemGlobalAuthHttpPort(o["auth-http-port"], d, "auth_http_port")); err != nil {
		if !fortiAPIPatch(o["auth-http-port"]) {
			return fmt.Errorf("Error reading auth_http_port: %v", err)
		}
	}

	if err = d.Set("auth_https_port", dataSourceFlattenSystemGlobalAuthHttpsPort(o["auth-https-port"], d, "auth_https_port")); err != nil {
		if !fortiAPIPatch(o["auth-https-port"]) {
			return fmt.Errorf("Error reading auth_https_port: %v", err)
		}
	}

	if err = d.Set("auth_ike_saml_port", dataSourceFlattenSystemGlobalAuthIkeSamlPort(o["auth-ike-saml-port"], d, "auth_ike_saml_port")); err != nil {
		if !fortiAPIPatch(o["auth-ike-saml-port"]) {
			return fmt.Errorf("Error reading auth_ike_saml_port: %v", err)
		}
	}

	if err = d.Set("auth_keepalive", dataSourceFlattenSystemGlobalAuthKeepalive(o["auth-keepalive"], d, "auth_keepalive")); err != nil {
		if !fortiAPIPatch(o["auth-keepalive"]) {
			return fmt.Errorf("Error reading auth_keepalive: %v", err)
		}
	}

	if err = d.Set("policy_auth_concurrent", dataSourceFlattenSystemGlobalPolicyAuthConcurrent(o["policy-auth-concurrent"], d, "policy_auth_concurrent")); err != nil {
		if !fortiAPIPatch(o["policy-auth-concurrent"]) {
			return fmt.Errorf("Error reading policy_auth_concurrent: %v", err)
		}
	}

	if err = d.Set("auth_session_limit", dataSourceFlattenSystemGlobalAuthSessionLimit(o["auth-session-limit"], d, "auth_session_limit")); err != nil {
		if !fortiAPIPatch(o["auth-session-limit"]) {
			return fmt.Errorf("Error reading auth_session_limit: %v", err)
		}
	}

	if err = d.Set("auth_cert", dataSourceFlattenSystemGlobalAuthCert(o["auth-cert"], d, "auth_cert")); err != nil {
		if !fortiAPIPatch(o["auth-cert"]) {
			return fmt.Errorf("Error reading auth_cert: %v", err)
		}
	}

	if err = d.Set("clt_cert_req", dataSourceFlattenSystemGlobalCltCertReq(o["clt-cert-req"], d, "clt_cert_req")); err != nil {
		if !fortiAPIPatch(o["clt-cert-req"]) {
			return fmt.Errorf("Error reading clt_cert_req: %v", err)
		}
	}

	if err = d.Set("fortiservice_port", dataSourceFlattenSystemGlobalFortiservicePort(o["fortiservice-port"], d, "fortiservice_port")); err != nil {
		if !fortiAPIPatch(o["fortiservice-port"]) {
			return fmt.Errorf("Error reading fortiservice_port: %v", err)
		}
	}

	if err = d.Set("endpoint_control_portal_port", dataSourceFlattenSystemGlobalEndpointControlPortalPort(o["endpoint-control-portal-port"], d, "endpoint_control_portal_port")); err != nil {
		if !fortiAPIPatch(o["endpoint-control-portal-port"]) {
			return fmt.Errorf("Error reading endpoint_control_portal_port: %v", err)
		}
	}

	if err = d.Set("endpoint_control_fds_access", dataSourceFlattenSystemGlobalEndpointControlFdsAccess(o["endpoint-control-fds-access"], d, "endpoint_control_fds_access")); err != nil {
		if !fortiAPIPatch(o["endpoint-control-fds-access"]) {
			return fmt.Errorf("Error reading endpoint_control_fds_access: %v", err)
		}
	}

	if err = d.Set("tp_mc_skip_policy", dataSourceFlattenSystemGlobalTpMcSkipPolicy(o["tp-mc-skip-policy"], d, "tp_mc_skip_policy")); err != nil {
		if !fortiAPIPatch(o["tp-mc-skip-policy"]) {
			return fmt.Errorf("Error reading tp_mc_skip_policy: %v", err)
		}
	}

	if err = d.Set("cfg_save", dataSourceFlattenSystemGlobalCfgSave(o["cfg-save"], d, "cfg_save")); err != nil {
		if !fortiAPIPatch(o["cfg-save"]) {
			return fmt.Errorf("Error reading cfg_save: %v", err)
		}
	}

	if err = d.Set("cfg_revert_timeout", dataSourceFlattenSystemGlobalCfgRevertTimeout(o["cfg-revert-timeout"], d, "cfg_revert_timeout")); err != nil {
		if !fortiAPIPatch(o["cfg-revert-timeout"]) {
			return fmt.Errorf("Error reading cfg_revert_timeout: %v", err)
		}
	}

	if err = d.Set("reboot_upon_config_restore", dataSourceFlattenSystemGlobalRebootUponConfigRestore(o["reboot-upon-config-restore"], d, "reboot_upon_config_restore")); err != nil {
		if !fortiAPIPatch(o["reboot-upon-config-restore"]) {
			return fmt.Errorf("Error reading reboot_upon_config_restore: %v", err)
		}
	}

	if err = d.Set("admin_scp", dataSourceFlattenSystemGlobalAdminScp(o["admin-scp"], d, "admin_scp")); err != nil {
		if !fortiAPIPatch(o["admin-scp"]) {
			return fmt.Errorf("Error reading admin_scp: %v", err)
		}
	}

	if err = d.Set("security_rating_result_submission", dataSourceFlattenSystemGlobalSecurityRatingResultSubmission(o["security-rating-result-submission"], d, "security_rating_result_submission")); err != nil {
		if !fortiAPIPatch(o["security-rating-result-submission"]) {
			return fmt.Errorf("Error reading security_rating_result_submission: %v", err)
		}
	}

	if err = d.Set("security_rating_run_on_schedule", dataSourceFlattenSystemGlobalSecurityRatingRunOnSchedule(o["security-rating-run-on-schedule"], d, "security_rating_run_on_schedule")); err != nil {
		if !fortiAPIPatch(o["security-rating-run-on-schedule"]) {
			return fmt.Errorf("Error reading security_rating_run_on_schedule: %v", err)
		}
	}

	if err = d.Set("wireless_controller", dataSourceFlattenSystemGlobalWirelessController(o["wireless-controller"], d, "wireless_controller")); err != nil {
		if !fortiAPIPatch(o["wireless-controller"]) {
			return fmt.Errorf("Error reading wireless_controller: %v", err)
		}
	}

	if err = d.Set("wireless_controller_port", dataSourceFlattenSystemGlobalWirelessControllerPort(o["wireless-controller-port"], d, "wireless_controller_port")); err != nil {
		if !fortiAPIPatch(o["wireless-controller-port"]) {
			return fmt.Errorf("Error reading wireless_controller_port: %v", err)
		}
	}

	if err = d.Set("fortiextender_data_port", dataSourceFlattenSystemGlobalFortiextenderDataPort(o["fortiextender-data-port"], d, "fortiextender_data_port")); err != nil {
		if !fortiAPIPatch(o["fortiextender-data-port"]) {
			return fmt.Errorf("Error reading fortiextender_data_port: %v", err)
		}
	}

	if err = d.Set("fortiextender", dataSourceFlattenSystemGlobalFortiextender(o["fortiextender"], d, "fortiextender")); err != nil {
		if !fortiAPIPatch(o["fortiextender"]) {
			return fmt.Errorf("Error reading fortiextender: %v", err)
		}
	}

	if err = d.Set("extender_controller_reserved_network", dataSourceFlattenSystemGlobalExtenderControllerReservedNetwork(o["extender-controller-reserved-network"], d, "extender_controller_reserved_network")); err != nil {
		if !fortiAPIPatch(o["extender-controller-reserved-network"]) {
			return fmt.Errorf("Error reading extender_controller_reserved_network: %v", err)
		}
	}

	if err = d.Set("fortiextender_discovery_lockdown", dataSourceFlattenSystemGlobalFortiextenderDiscoveryLockdown(o["fortiextender-discovery-lockdown"], d, "fortiextender_discovery_lockdown")); err != nil {
		if !fortiAPIPatch(o["fortiextender-discovery-lockdown"]) {
			return fmt.Errorf("Error reading fortiextender_discovery_lockdown: %v", err)
		}
	}

	if err = d.Set("fortiextender_vlan_mode", dataSourceFlattenSystemGlobalFortiextenderVlanMode(o["fortiextender-vlan-mode"], d, "fortiextender_vlan_mode")); err != nil {
		if !fortiAPIPatch(o["fortiextender-vlan-mode"]) {
			return fmt.Errorf("Error reading fortiextender_vlan_mode: %v", err)
		}
	}

	if err = d.Set("fortiextender_provision_on_authorization", dataSourceFlattenSystemGlobalFortiextenderProvisionOnAuthorization(o["fortiextender-provision-on-authorization"], d, "fortiextender_provision_on_authorization")); err != nil {
		if !fortiAPIPatch(o["fortiextender-provision-on-authorization"]) {
			return fmt.Errorf("Error reading fortiextender_provision_on_authorization: %v", err)
		}
	}

	if err = d.Set("switch_controller", dataSourceFlattenSystemGlobalSwitchController(o["switch-controller"], d, "switch_controller")); err != nil {
		if !fortiAPIPatch(o["switch-controller"]) {
			return fmt.Errorf("Error reading switch_controller: %v", err)
		}
	}

	if err = d.Set("switch_controller_reserved_network", dataSourceFlattenSystemGlobalSwitchControllerReservedNetwork(o["switch-controller-reserved-network"], d, "switch_controller_reserved_network")); err != nil {
		if !fortiAPIPatch(o["switch-controller-reserved-network"]) {
			return fmt.Errorf("Error reading switch_controller_reserved_network: %v", err)
		}
	}

	if err = d.Set("dnsproxy_worker_count", dataSourceFlattenSystemGlobalDnsproxyWorkerCount(o["dnsproxy-worker-count"], d, "dnsproxy_worker_count")); err != nil {
		if !fortiAPIPatch(o["dnsproxy-worker-count"]) {
			return fmt.Errorf("Error reading dnsproxy_worker_count: %v", err)
		}
	}

	if err = d.Set("url_filter_count", dataSourceFlattenSystemGlobalUrlFilterCount(o["url-filter-count"], d, "url_filter_count")); err != nil {
		if !fortiAPIPatch(o["url-filter-count"]) {
			return fmt.Errorf("Error reading url_filter_count: %v", err)
		}
	}

	if err = d.Set("httpd_max_worker_count", dataSourceFlattenSystemGlobalHttpdMaxWorkerCount(o["httpd-max-worker-count"], d, "httpd_max_worker_count")); err != nil {
		if !fortiAPIPatch(o["httpd-max-worker-count"]) {
			return fmt.Errorf("Error reading httpd_max_worker_count: %v", err)
		}
	}

	if err = d.Set("proxy_worker_count", dataSourceFlattenSystemGlobalProxyWorkerCount(o["proxy-worker-count"], d, "proxy_worker_count")); err != nil {
		if !fortiAPIPatch(o["proxy-worker-count"]) {
			return fmt.Errorf("Error reading proxy_worker_count: %v", err)
		}
	}

	if err = d.Set("scanunit_count", dataSourceFlattenSystemGlobalScanunitCount(o["scanunit-count"], d, "scanunit_count")); err != nil {
		if !fortiAPIPatch(o["scanunit-count"]) {
			return fmt.Errorf("Error reading scanunit_count: %v", err)
		}
	}

	if err = d.Set("proxy_hardware_acceleration", dataSourceFlattenSystemGlobalProxyHardwareAcceleration(o["proxy-hardware-acceleration"], d, "proxy_hardware_acceleration")); err != nil {
		if !fortiAPIPatch(o["proxy-hardware-acceleration"]) {
			return fmt.Errorf("Error reading proxy_hardware_acceleration: %v", err)
		}
	}

	if err = d.Set("proxy_kxp_hardware_acceleration", dataSourceFlattenSystemGlobalProxyKxpHardwareAcceleration(o["proxy-kxp-hardware-acceleration"], d, "proxy_kxp_hardware_acceleration")); err != nil {
		if !fortiAPIPatch(o["proxy-kxp-hardware-acceleration"]) {
			return fmt.Errorf("Error reading proxy_kxp_hardware_acceleration: %v", err)
		}
	}

	if err = d.Set("proxy_cipher_hardware_acceleration", dataSourceFlattenSystemGlobalProxyCipherHardwareAcceleration(o["proxy-cipher-hardware-acceleration"], d, "proxy_cipher_hardware_acceleration")); err != nil {
		if !fortiAPIPatch(o["proxy-cipher-hardware-acceleration"]) {
			return fmt.Errorf("Error reading proxy_cipher_hardware_acceleration: %v", err)
		}
	}

	if err = d.Set("fgd_alert_subscription", dataSourceFlattenSystemGlobalFgdAlertSubscription(o["fgd-alert-subscription"], d, "fgd_alert_subscription")); err != nil {
		if !fortiAPIPatch(o["fgd-alert-subscription"]) {
			return fmt.Errorf("Error reading fgd_alert_subscription: %v", err)
		}
	}

	if err = d.Set("ipsec_hmac_offload", dataSourceFlattenSystemGlobalIpsecHmacOffload(o["ipsec-hmac-offload"], d, "ipsec_hmac_offload")); err != nil {
		if !fortiAPIPatch(o["ipsec-hmac-offload"]) {
			return fmt.Errorf("Error reading ipsec_hmac_offload: %v", err)
		}
	}

	if err = d.Set("ipv6_accept_dad", dataSourceFlattenSystemGlobalIpv6AcceptDad(o["ipv6-accept-dad"], d, "ipv6_accept_dad")); err != nil {
		if !fortiAPIPatch(o["ipv6-accept-dad"]) {
			return fmt.Errorf("Error reading ipv6_accept_dad: %v", err)
		}
	}

	if err = d.Set("ipv6_allow_anycast_probe", dataSourceFlattenSystemGlobalIpv6AllowAnycastProbe(o["ipv6-allow-anycast-probe"], d, "ipv6_allow_anycast_probe")); err != nil {
		if !fortiAPIPatch(o["ipv6-allow-anycast-probe"]) {
			return fmt.Errorf("Error reading ipv6_allow_anycast_probe: %v", err)
		}
	}

	if err = d.Set("ipv6_allow_multicast_probe", dataSourceFlattenSystemGlobalIpv6AllowMulticastProbe(o["ipv6-allow-multicast-probe"], d, "ipv6_allow_multicast_probe")); err != nil {
		if !fortiAPIPatch(o["ipv6-allow-multicast-probe"]) {
			return fmt.Errorf("Error reading ipv6_allow_multicast_probe: %v", err)
		}
	}

	if err = d.Set("ipv6_allow_local_in_silent_drop", dataSourceFlattenSystemGlobalIpv6AllowLocalInSilentDrop(o["ipv6-allow-local-in-silent-drop"], d, "ipv6_allow_local_in_silent_drop")); err != nil {
		if !fortiAPIPatch(o["ipv6-allow-local-in-silent-drop"]) {
			return fmt.Errorf("Error reading ipv6_allow_local_in_silent_drop: %v", err)
		}
	}

	if err = d.Set("ipv6_allow_local_in_slient_drop", dataSourceFlattenSystemGlobalIpv6AllowLocalInSlientDrop(o["ipv6-allow-local-in-slient-drop"], d, "ipv6_allow_local_in_slient_drop")); err != nil {
		if !fortiAPIPatch(o["ipv6-allow-local-in-slient-drop"]) {
			return fmt.Errorf("Error reading ipv6_allow_local_in_slient_drop: %v", err)
		}
	}

	if err = d.Set("csr_ca_attribute", dataSourceFlattenSystemGlobalCsrCaAttribute(o["csr-ca-attribute"], d, "csr_ca_attribute")); err != nil {
		if !fortiAPIPatch(o["csr-ca-attribute"]) {
			return fmt.Errorf("Error reading csr_ca_attribute: %v", err)
		}
	}

	if err = d.Set("wimax_4g_usb", dataSourceFlattenSystemGlobalWimax4GUsb(o["wimax-4g-usb"], d, "wimax_4g_usb")); err != nil {
		if !fortiAPIPatch(o["wimax-4g-usb"]) {
			return fmt.Errorf("Error reading wimax_4g_usb: %v", err)
		}
	}

	if err = d.Set("cert_chain_max", dataSourceFlattenSystemGlobalCertChainMax(o["cert-chain-max"], d, "cert_chain_max")); err != nil {
		if !fortiAPIPatch(o["cert-chain-max"]) {
			return fmt.Errorf("Error reading cert_chain_max: %v", err)
		}
	}

	if err = d.Set("sslvpn_max_worker_count", dataSourceFlattenSystemGlobalSslvpnMaxWorkerCount(o["sslvpn-max-worker-count"], d, "sslvpn_max_worker_count")); err != nil {
		if !fortiAPIPatch(o["sslvpn-max-worker-count"]) {
			return fmt.Errorf("Error reading sslvpn_max_worker_count: %v", err)
		}
	}

	if err = d.Set("sslvpn_affinity", dataSourceFlattenSystemGlobalSslvpnAffinity(o["sslvpn-affinity"], d, "sslvpn_affinity")); err != nil {
		if !fortiAPIPatch(o["sslvpn-affinity"]) {
			return fmt.Errorf("Error reading sslvpn_affinity: %v", err)
		}
	}

	if err = d.Set("vpn_ems_sn_check", dataSourceFlattenSystemGlobalVpnEmsSnCheck(o["vpn-ems-sn-check"], d, "vpn_ems_sn_check")); err != nil {
		if !fortiAPIPatch(o["vpn-ems-sn-check"]) {
			return fmt.Errorf("Error reading vpn_ems_sn_check: %v", err)
		}
	}

	if err = d.Set("sslvpn_web_mode", dataSourceFlattenSystemGlobalSslvpnWebMode(o["sslvpn-web-mode"], d, "sslvpn_web_mode")); err != nil {
		if !fortiAPIPatch(o["sslvpn-web-mode"]) {
			return fmt.Errorf("Error reading sslvpn_web_mode: %v", err)
		}
	}

	if err = d.Set("sslvpn_ems_sn_check", dataSourceFlattenSystemGlobalSslvpnEmsSnCheck(o["sslvpn-ems-sn-check"], d, "sslvpn_ems_sn_check")); err != nil {
		if !fortiAPIPatch(o["sslvpn-ems-sn-check"]) {
			return fmt.Errorf("Error reading sslvpn_ems_sn_check: %v", err)
		}
	}

	if err = d.Set("sslvpn_kxp_hardware_acceleration", dataSourceFlattenSystemGlobalSslvpnKxpHardwareAcceleration(o["sslvpn-kxp-hardware-acceleration"], d, "sslvpn_kxp_hardware_acceleration")); err != nil {
		if !fortiAPIPatch(o["sslvpn-kxp-hardware-acceleration"]) {
			return fmt.Errorf("Error reading sslvpn_kxp_hardware_acceleration: %v", err)
		}
	}

	if err = d.Set("sslvpn_cipher_hardware_acceleration", dataSourceFlattenSystemGlobalSslvpnCipherHardwareAcceleration(o["sslvpn-cipher-hardware-acceleration"], d, "sslvpn_cipher_hardware_acceleration")); err != nil {
		if !fortiAPIPatch(o["sslvpn-cipher-hardware-acceleration"]) {
			return fmt.Errorf("Error reading sslvpn_cipher_hardware_acceleration: %v", err)
		}
	}

	if err = d.Set("sslvpn_plugin_version_check", dataSourceFlattenSystemGlobalSslvpnPluginVersionCheck(o["sslvpn-plugin-version-check"], d, "sslvpn_plugin_version_check")); err != nil {
		if !fortiAPIPatch(o["sslvpn-plugin-version-check"]) {
			return fmt.Errorf("Error reading sslvpn_plugin_version_check: %v", err)
		}
	}

	if err = d.Set("two_factor_ftk_expiry", dataSourceFlattenSystemGlobalTwoFactorFtkExpiry(o["two-factor-ftk-expiry"], d, "two_factor_ftk_expiry")); err != nil {
		if !fortiAPIPatch(o["two-factor-ftk-expiry"]) {
			return fmt.Errorf("Error reading two_factor_ftk_expiry: %v", err)
		}
	}

	if err = d.Set("two_factor_email_expiry", dataSourceFlattenSystemGlobalTwoFactorEmailExpiry(o["two-factor-email-expiry"], d, "two_factor_email_expiry")); err != nil {
		if !fortiAPIPatch(o["two-factor-email-expiry"]) {
			return fmt.Errorf("Error reading two_factor_email_expiry: %v", err)
		}
	}

	if err = d.Set("two_factor_sms_expiry", dataSourceFlattenSystemGlobalTwoFactorSmsExpiry(o["two-factor-sms-expiry"], d, "two_factor_sms_expiry")); err != nil {
		if !fortiAPIPatch(o["two-factor-sms-expiry"]) {
			return fmt.Errorf("Error reading two_factor_sms_expiry: %v", err)
		}
	}

	if err = d.Set("two_factor_fac_expiry", dataSourceFlattenSystemGlobalTwoFactorFacExpiry(o["two-factor-fac-expiry"], d, "two_factor_fac_expiry")); err != nil {
		if !fortiAPIPatch(o["two-factor-fac-expiry"]) {
			return fmt.Errorf("Error reading two_factor_fac_expiry: %v", err)
		}
	}

	if err = d.Set("two_factor_ftm_expiry", dataSourceFlattenSystemGlobalTwoFactorFtmExpiry(o["two-factor-ftm-expiry"], d, "two_factor_ftm_expiry")); err != nil {
		if !fortiAPIPatch(o["two-factor-ftm-expiry"]) {
			return fmt.Errorf("Error reading two_factor_ftm_expiry: %v", err)
		}
	}

	if err = d.Set("per_user_bal", dataSourceFlattenSystemGlobalPerUserBal(o["per-user-bal"], d, "per_user_bal")); err != nil {
		if !fortiAPIPatch(o["per-user-bal"]) {
			return fmt.Errorf("Error reading per_user_bal: %v", err)
		}
	}

	if err = d.Set("per_user_bwl", dataSourceFlattenSystemGlobalPerUserBwl(o["per-user-bwl"], d, "per_user_bwl")); err != nil {
		if !fortiAPIPatch(o["per-user-bwl"]) {
			return fmt.Errorf("Error reading per_user_bwl: %v", err)
		}
	}

	if err = d.Set("virtual_server_count", dataSourceFlattenSystemGlobalVirtualServerCount(o["virtual-server-count"], d, "virtual_server_count")); err != nil {
		if !fortiAPIPatch(o["virtual-server-count"]) {
			return fmt.Errorf("Error reading virtual_server_count: %v", err)
		}
	}

	if err = d.Set("virtual_server_hardware_acceleration", dataSourceFlattenSystemGlobalVirtualServerHardwareAcceleration(o["virtual-server-hardware-acceleration"], d, "virtual_server_hardware_acceleration")); err != nil {
		if !fortiAPIPatch(o["virtual-server-hardware-acceleration"]) {
			return fmt.Errorf("Error reading virtual_server_hardware_acceleration: %v", err)
		}
	}

	if err = d.Set("wad_worker_count", dataSourceFlattenSystemGlobalWadWorkerCount(o["wad-worker-count"], d, "wad_worker_count")); err != nil {
		if !fortiAPIPatch(o["wad-worker-count"]) {
			return fmt.Errorf("Error reading wad_worker_count: %v", err)
		}
	}

	if err = d.Set("wad_worker_dev_cache", dataSourceFlattenSystemGlobalWadWorkerDevCache(o["wad-worker-dev-cache"], d, "wad_worker_dev_cache")); err != nil {
		if !fortiAPIPatch(o["wad-worker-dev-cache"]) {
			return fmt.Errorf("Error reading wad_worker_dev_cache: %v", err)
		}
	}

	if err = d.Set("wad_csvc_cs_count", dataSourceFlattenSystemGlobalWadCsvcCsCount(o["wad-csvc-cs-count"], d, "wad_csvc_cs_count")); err != nil {
		if !fortiAPIPatch(o["wad-csvc-cs-count"]) {
			return fmt.Errorf("Error reading wad_csvc_cs_count: %v", err)
		}
	}

	if err = d.Set("wad_csvc_db_count", dataSourceFlattenSystemGlobalWadCsvcDbCount(o["wad-csvc-db-count"], d, "wad_csvc_db_count")); err != nil {
		if !fortiAPIPatch(o["wad-csvc-db-count"]) {
			return fmt.Errorf("Error reading wad_csvc_db_count: %v", err)
		}
	}

	if err = d.Set("wad_source_affinity", dataSourceFlattenSystemGlobalWadSourceAffinity(o["wad-source-affinity"], d, "wad_source_affinity")); err != nil {
		if !fortiAPIPatch(o["wad-source-affinity"]) {
			return fmt.Errorf("Error reading wad_source_affinity: %v", err)
		}
	}

	if err = d.Set("wad_memory_change_granularity", dataSourceFlattenSystemGlobalWadMemoryChangeGranularity(o["wad-memory-change-granularity"], d, "wad_memory_change_granularity")); err != nil {
		if !fortiAPIPatch(o["wad-memory-change-granularity"]) {
			return fmt.Errorf("Error reading wad_memory_change_granularity: %v", err)
		}
	}

	if err = d.Set("login_timestamp", dataSourceFlattenSystemGlobalLoginTimestamp(o["login-timestamp"], d, "login_timestamp")); err != nil {
		if !fortiAPIPatch(o["login-timestamp"]) {
			return fmt.Errorf("Error reading login_timestamp: %v", err)
		}
	}

	if err = d.Set("ip_conflict_detection", dataSourceFlattenSystemGlobalIpConflictDetection(o["ip-conflict-detection"], d, "ip_conflict_detection")); err != nil {
		if !fortiAPIPatch(o["ip-conflict-detection"]) {
			return fmt.Errorf("Error reading ip_conflict_detection: %v", err)
		}
	}

	if err = d.Set("miglogd_children", dataSourceFlattenSystemGlobalMiglogdChildren(o["miglogd-children"], d, "miglogd_children")); err != nil {
		if !fortiAPIPatch(o["miglogd-children"]) {
			return fmt.Errorf("Error reading miglogd_children: %v", err)
		}
	}

	if err = d.Set("log_daemon_cpu_threshold", dataSourceFlattenSystemGlobalLogDaemonCpuThreshold(o["log-daemon-cpu-threshold"], d, "log_daemon_cpu_threshold")); err != nil {
		if !fortiAPIPatch(o["log-daemon-cpu-threshold"]) {
			return fmt.Errorf("Error reading log_daemon_cpu_threshold: %v", err)
		}
	}

	if err = d.Set("special_file_23_support", dataSourceFlattenSystemGlobalSpecialFile23Support(o["special-file-23-support"], d, "special_file_23_support")); err != nil {
		if !fortiAPIPatch(o["special-file-23-support"]) {
			return fmt.Errorf("Error reading special_file_23_support: %v", err)
		}
	}

	if err = d.Set("log_uuid_policy", dataSourceFlattenSystemGlobalLogUuidPolicy(o["log-uuid-policy"], d, "log_uuid_policy")); err != nil {
		if !fortiAPIPatch(o["log-uuid-policy"]) {
			return fmt.Errorf("Error reading log_uuid_policy: %v", err)
		}
	}

	if err = d.Set("log_uuid_address", dataSourceFlattenSystemGlobalLogUuidAddress(o["log-uuid-address"], d, "log_uuid_address")); err != nil {
		if !fortiAPIPatch(o["log-uuid-address"]) {
			return fmt.Errorf("Error reading log_uuid_address: %v", err)
		}
	}

	if err = d.Set("log_ssl_connection", dataSourceFlattenSystemGlobalLogSslConnection(o["log-ssl-connection"], d, "log_ssl_connection")); err != nil {
		if !fortiAPIPatch(o["log-ssl-connection"]) {
			return fmt.Errorf("Error reading log_ssl_connection: %v", err)
		}
	}

	if err = d.Set("gui_rest_api_cache", dataSourceFlattenSystemGlobalGuiRestApiCache(o["gui-rest-api-cache"], d, "gui_rest_api_cache")); err != nil {
		if !fortiAPIPatch(o["gui-rest-api-cache"]) {
			return fmt.Errorf("Error reading gui_rest_api_cache: %v", err)
		}
	}

	if err = d.Set("rest_api_key_url_query", dataSourceFlattenSystemGlobalRestApiKeyUrlQuery(o["rest-api-key-url-query"], d, "rest_api_key_url_query")); err != nil {
		if !fortiAPIPatch(o["rest-api-key-url-query"]) {
			return fmt.Errorf("Error reading rest_api_key_url_query: %v", err)
		}
	}

	if err = d.Set("gui_cdn_domain_override", dataSourceFlattenSystemGlobalGuiCdnDomainOverride(o["gui-cdn-domain-override"], d, "gui_cdn_domain_override")); err != nil {
		if !fortiAPIPatch(o["gui-cdn-domain-override"]) {
			return fmt.Errorf("Error reading gui_cdn_domain_override: %v", err)
		}
	}

	if err = d.Set("gui_fortiguard_resource_fetch", dataSourceFlattenSystemGlobalGuiFortiguardResourceFetch(o["gui-fortiguard-resource-fetch"], d, "gui_fortiguard_resource_fetch")); err != nil {
		if !fortiAPIPatch(o["gui-fortiguard-resource-fetch"]) {
			return fmt.Errorf("Error reading gui_fortiguard_resource_fetch: %v", err)
		}
	}

	if err = d.Set("arp_max_entry", dataSourceFlattenSystemGlobalArpMaxEntry(o["arp-max-entry"], d, "arp_max_entry")); err != nil {
		if !fortiAPIPatch(o["arp-max-entry"]) {
			return fmt.Errorf("Error reading arp_max_entry: %v", err)
		}
	}

	if err = d.Set("ha_affinity", dataSourceFlattenSystemGlobalHaAffinity(o["ha-affinity"], d, "ha_affinity")); err != nil {
		if !fortiAPIPatch(o["ha-affinity"]) {
			return fmt.Errorf("Error reading ha_affinity: %v", err)
		}
	}

	if err = d.Set("bfd_affinity", dataSourceFlattenSystemGlobalBfdAffinity(o["bfd-affinity"], d, "bfd_affinity")); err != nil {
		if !fortiAPIPatch(o["bfd-affinity"]) {
			return fmt.Errorf("Error reading bfd_affinity: %v", err)
		}
	}

	if err = d.Set("cmdbsvr_affinity", dataSourceFlattenSystemGlobalCmdbsvrAffinity(o["cmdbsvr-affinity"], d, "cmdbsvr_affinity")); err != nil {
		if !fortiAPIPatch(o["cmdbsvr-affinity"]) {
			return fmt.Errorf("Error reading cmdbsvr_affinity: %v", err)
		}
	}

	if err = d.Set("av_affinity", dataSourceFlattenSystemGlobalAvAffinity(o["av-affinity"], d, "av_affinity")); err != nil {
		if !fortiAPIPatch(o["av-affinity"]) {
			return fmt.Errorf("Error reading av_affinity: %v", err)
		}
	}

	if err = d.Set("wad_affinity", dataSourceFlattenSystemGlobalWadAffinity(o["wad-affinity"], d, "wad_affinity")); err != nil {
		if !fortiAPIPatch(o["wad-affinity"]) {
			return fmt.Errorf("Error reading wad_affinity: %v", err)
		}
	}

	if err = d.Set("ips_affinity", dataSourceFlattenSystemGlobalIpsAffinity(o["ips-affinity"], d, "ips_affinity")); err != nil {
		if !fortiAPIPatch(o["ips-affinity"]) {
			return fmt.Errorf("Error reading ips_affinity: %v", err)
		}
	}

	if err = d.Set("miglog_affinity", dataSourceFlattenSystemGlobalMiglogAffinity(o["miglog-affinity"], d, "miglog_affinity")); err != nil {
		if !fortiAPIPatch(o["miglog-affinity"]) {
			return fmt.Errorf("Error reading miglog_affinity: %v", err)
		}
	}

	if err = d.Set("syslog_affinity", dataSourceFlattenSystemGlobalSyslogAffinity(o["syslog-affinity"], d, "syslog_affinity")); err != nil {
		if !fortiAPIPatch(o["syslog-affinity"]) {
			return fmt.Errorf("Error reading syslog_affinity: %v", err)
		}
	}

	if err = d.Set("url_filter_affinity", dataSourceFlattenSystemGlobalUrlFilterAffinity(o["url-filter-affinity"], d, "url_filter_affinity")); err != nil {
		if !fortiAPIPatch(o["url-filter-affinity"]) {
			return fmt.Errorf("Error reading url_filter_affinity: %v", err)
		}
	}

	if err = d.Set("router_affinity", dataSourceFlattenSystemGlobalRouterAffinity(o["router-affinity"], d, "router_affinity")); err != nil {
		if !fortiAPIPatch(o["router-affinity"]) {
			return fmt.Errorf("Error reading router_affinity: %v", err)
		}
	}

	if err = d.Set("ndp_max_entry", dataSourceFlattenSystemGlobalNdpMaxEntry(o["ndp-max-entry"], d, "ndp_max_entry")); err != nil {
		if !fortiAPIPatch(o["ndp-max-entry"]) {
			return fmt.Errorf("Error reading ndp_max_entry: %v", err)
		}
	}

	if err = d.Set("br_fdb_max_entry", dataSourceFlattenSystemGlobalBrFdbMaxEntry(o["br-fdb-max-entry"], d, "br_fdb_max_entry")); err != nil {
		if !fortiAPIPatch(o["br-fdb-max-entry"]) {
			return fmt.Errorf("Error reading br_fdb_max_entry: %v", err)
		}
	}

	if err = d.Set("max_route_cache_size", dataSourceFlattenSystemGlobalMaxRouteCacheSize(o["max-route-cache-size"], d, "max_route_cache_size")); err != nil {
		if !fortiAPIPatch(o["max-route-cache-size"]) {
			return fmt.Errorf("Error reading max_route_cache_size: %v", err)
		}
	}

	if err = d.Set("ipsec_qat_offload", dataSourceFlattenSystemGlobalIpsecQatOffload(o["ipsec-qat-offload"], d, "ipsec_qat_offload")); err != nil {
		if !fortiAPIPatch(o["ipsec-qat-offload"]) {
			return fmt.Errorf("Error reading ipsec_qat_offload: %v", err)
		}
	}

	if err = d.Set("ipsec_round_robin", dataSourceFlattenSystemGlobalIpsecRoundRobin(o["ipsec-round-robin"], d, "ipsec_round_robin")); err != nil {
		if !fortiAPIPatch(o["ipsec-round-robin"]) {
			return fmt.Errorf("Error reading ipsec_round_robin: %v", err)
		}
	}

	if err = d.Set("ipsec_asic_offload", dataSourceFlattenSystemGlobalIpsecAsicOffload(o["ipsec-asic-offload"], d, "ipsec_asic_offload")); err != nil {
		if !fortiAPIPatch(o["ipsec-asic-offload"]) {
			return fmt.Errorf("Error reading ipsec_asic_offload: %v", err)
		}
	}

	if err = d.Set("ipsec_soft_dec_async", dataSourceFlattenSystemGlobalIpsecSoftDecAsync(o["ipsec-soft-dec-async"], d, "ipsec_soft_dec_async")); err != nil {
		if !fortiAPIPatch(o["ipsec-soft-dec-async"]) {
			return fmt.Errorf("Error reading ipsec_soft_dec_async: %v", err)
		}
	}

	if err = d.Set("ike_embryonic_limit", dataSourceFlattenSystemGlobalIkeEmbryonicLimit(o["ike-embryonic-limit"], d, "ike_embryonic_limit")); err != nil {
		if !fortiAPIPatch(o["ike-embryonic-limit"]) {
			return fmt.Errorf("Error reading ike_embryonic_limit: %v", err)
		}
	}

	if err = d.Set("device_idle_timeout", dataSourceFlattenSystemGlobalDeviceIdleTimeout(o["device-idle-timeout"], d, "device_idle_timeout")); err != nil {
		if !fortiAPIPatch(o["device-idle-timeout"]) {
			return fmt.Errorf("Error reading device_idle_timeout: %v", err)
		}
	}

	if err = d.Set("user_device_store_max_devices", dataSourceFlattenSystemGlobalUserDeviceStoreMaxDevices(o["user-device-store-max-devices"], d, "user_device_store_max_devices")); err != nil {
		if !fortiAPIPatch(o["user-device-store-max-devices"]) {
			return fmt.Errorf("Error reading user_device_store_max_devices: %v", err)
		}
	}

	if err = d.Set("user_device_store_max_device_mem", dataSourceFlattenSystemGlobalUserDeviceStoreMaxDeviceMem(o["user-device-store-max-device-mem"], d, "user_device_store_max_device_mem")); err != nil {
		if !fortiAPIPatch(o["user-device-store-max-device-mem"]) {
			return fmt.Errorf("Error reading user_device_store_max_device_mem: %v", err)
		}
	}

	if err = d.Set("user_device_store_max_users", dataSourceFlattenSystemGlobalUserDeviceStoreMaxUsers(o["user-device-store-max-users"], d, "user_device_store_max_users")); err != nil {
		if !fortiAPIPatch(o["user-device-store-max-users"]) {
			return fmt.Errorf("Error reading user_device_store_max_users: %v", err)
		}
	}

	if err = d.Set("user_device_store_max_unified_mem", dataSourceFlattenSystemGlobalUserDeviceStoreMaxUnifiedMem(o["user-device-store-max-unified-mem"], d, "user_device_store_max_unified_mem")); err != nil {
		if !fortiAPIPatch(o["user-device-store-max-unified-mem"]) {
			return fmt.Errorf("Error reading user_device_store_max_unified_mem: %v", err)
		}
	}

	if err = d.Set("device_identification_active_scan_delay", dataSourceFlattenSystemGlobalDeviceIdentificationActiveScanDelay(o["device-identification-active-scan-delay"], d, "device_identification_active_scan_delay")); err != nil {
		if !fortiAPIPatch(o["device-identification-active-scan-delay"]) {
			return fmt.Errorf("Error reading device_identification_active_scan_delay: %v", err)
		}
	}

	if err = d.Set("compliance_check", dataSourceFlattenSystemGlobalComplianceCheck(o["compliance-check"], d, "compliance_check")); err != nil {
		if !fortiAPIPatch(o["compliance-check"]) {
			return fmt.Errorf("Error reading compliance_check: %v", err)
		}
	}

	if err = d.Set("compliance_check_time", dataSourceFlattenSystemGlobalComplianceCheckTime(o["compliance-check-time"], d, "compliance_check_time")); err != nil {
		if !fortiAPIPatch(o["compliance-check-time"]) {
			return fmt.Errorf("Error reading compliance_check_time: %v", err)
		}
	}

	if err = d.Set("gui_device_latitude", dataSourceFlattenSystemGlobalGuiDeviceLatitude(o["gui-device-latitude"], d, "gui_device_latitude")); err != nil {
		if !fortiAPIPatch(o["gui-device-latitude"]) {
			return fmt.Errorf("Error reading gui_device_latitude: %v", err)
		}
	}

	if err = d.Set("gui_device_longitude", dataSourceFlattenSystemGlobalGuiDeviceLongitude(o["gui-device-longitude"], d, "gui_device_longitude")); err != nil {
		if !fortiAPIPatch(o["gui-device-longitude"]) {
			return fmt.Errorf("Error reading gui_device_longitude: %v", err)
		}
	}

	if err = d.Set("private_data_encryption", dataSourceFlattenSystemGlobalPrivateDataEncryption(o["private-data-encryption"], d, "private_data_encryption")); err != nil {
		if !fortiAPIPatch(o["private-data-encryption"]) {
			return fmt.Errorf("Error reading private_data_encryption: %v", err)
		}
	}

	if err = d.Set("auto_auth_extension_device", dataSourceFlattenSystemGlobalAutoAuthExtensionDevice(o["auto-auth-extension-device"], d, "auto_auth_extension_device")); err != nil {
		if !fortiAPIPatch(o["auto-auth-extension-device"]) {
			return fmt.Errorf("Error reading auto_auth_extension_device: %v", err)
		}
	}

	if err = d.Set("gui_theme", dataSourceFlattenSystemGlobalGuiTheme(o["gui-theme"], d, "gui_theme")); err != nil {
		if !fortiAPIPatch(o["gui-theme"]) {
			return fmt.Errorf("Error reading gui_theme: %v", err)
		}
	}

	if err = d.Set("gui_date_format", dataSourceFlattenSystemGlobalGuiDateFormat(o["gui-date-format"], d, "gui_date_format")); err != nil {
		if !fortiAPIPatch(o["gui-date-format"]) {
			return fmt.Errorf("Error reading gui_date_format: %v", err)
		}
	}

	if err = d.Set("gui_date_time_source", dataSourceFlattenSystemGlobalGuiDateTimeSource(o["gui-date-time-source"], d, "gui_date_time_source")); err != nil {
		if !fortiAPIPatch(o["gui-date-time-source"]) {
			return fmt.Errorf("Error reading gui_date_time_source: %v", err)
		}
	}

	if err = d.Set("igmp_state_limit", dataSourceFlattenSystemGlobalIgmpStateLimit(o["igmp-state-limit"], d, "igmp_state_limit")); err != nil {
		if !fortiAPIPatch(o["igmp-state-limit"]) {
			return fmt.Errorf("Error reading igmp_state_limit: %v", err)
		}
	}

	if err = d.Set("cloud_communication", dataSourceFlattenSystemGlobalCloudCommunication(o["cloud-communication"], d, "cloud_communication")); err != nil {
		if !fortiAPIPatch(o["cloud-communication"]) {
			return fmt.Errorf("Error reading cloud_communication: %v", err)
		}
	}

	if err = d.Set("fec_port", dataSourceFlattenSystemGlobalFecPort(o["fec-port"], d, "fec_port")); err != nil {
		if !fortiAPIPatch(o["fec-port"]) {
			return fmt.Errorf("Error reading fec_port: %v", err)
		}
	}

	if err = d.Set("ipsec_ha_seqjump_rate", dataSourceFlattenSystemGlobalIpsecHaSeqjumpRate(o["ipsec-ha-seqjump-rate"], d, "ipsec_ha_seqjump_rate")); err != nil {
		if !fortiAPIPatch(o["ipsec-ha-seqjump-rate"]) {
			return fmt.Errorf("Error reading ipsec_ha_seqjump_rate: %v", err)
		}
	}

	if err = d.Set("fortitoken_cloud", dataSourceFlattenSystemGlobalFortitokenCloud(o["fortitoken-cloud"], d, "fortitoken_cloud")); err != nil {
		if !fortiAPIPatch(o["fortitoken-cloud"]) {
			return fmt.Errorf("Error reading fortitoken_cloud: %v", err)
		}
	}

	if err = d.Set("fortitoken_cloud_push_status", dataSourceFlattenSystemGlobalFortitokenCloudPushStatus(o["fortitoken-cloud-push-status"], d, "fortitoken_cloud_push_status")); err != nil {
		if !fortiAPIPatch(o["fortitoken-cloud-push-status"]) {
			return fmt.Errorf("Error reading fortitoken_cloud_push_status: %v", err)
		}
	}

	if err = d.Set("fortitoken_cloud_region", dataSourceFlattenSystemGlobalFortitokenCloudRegion(o["fortitoken-cloud-region"], d, "fortitoken_cloud_region")); err != nil {
		if !fortiAPIPatch(o["fortitoken-cloud-region"]) {
			return fmt.Errorf("Error reading fortitoken_cloud_region: %v", err)
		}
	}

	if err = d.Set("fortitoken_cloud_sync_interval", dataSourceFlattenSystemGlobalFortitokenCloudSyncInterval(o["fortitoken-cloud-sync-interval"], d, "fortitoken_cloud_sync_interval")); err != nil {
		if !fortiAPIPatch(o["fortitoken-cloud-sync-interval"]) {
			return fmt.Errorf("Error reading fortitoken_cloud_sync_interval: %v", err)
		}
	}

	if err = d.Set("faz_disk_buffer_size", dataSourceFlattenSystemGlobalFazDiskBufferSize(o["faz-disk-buffer-size"], d, "faz_disk_buffer_size")); err != nil {
		if !fortiAPIPatch(o["faz-disk-buffer-size"]) {
			return fmt.Errorf("Error reading faz_disk_buffer_size: %v", err)
		}
	}

	if err = d.Set("irq_time_accounting", dataSourceFlattenSystemGlobalIrqTimeAccounting(o["irq-time-accounting"], d, "irq_time_accounting")); err != nil {
		if !fortiAPIPatch(o["irq-time-accounting"]) {
			return fmt.Errorf("Error reading irq_time_accounting: %v", err)
		}
	}

	if err = d.Set("fortiipam_integration", dataSourceFlattenSystemGlobalFortiipamIntegration(o["fortiipam-integration"], d, "fortiipam_integration")); err != nil {
		if !fortiAPIPatch(o["fortiipam-integration"]) {
			return fmt.Errorf("Error reading fortiipam_integration: %v", err)
		}
	}

	if err = d.Set("management_ip", dataSourceFlattenSystemGlobalManagementIp(o["management-ip"], d, "management_ip")); err != nil {
		if !fortiAPIPatch(o["management-ip"]) {
			return fmt.Errorf("Error reading management_ip: %v", err)
		}
	}

	if err = d.Set("management_port", dataSourceFlattenSystemGlobalManagementPort(o["management-port"], d, "management_port")); err != nil {
		if !fortiAPIPatch(o["management-port"]) {
			return fmt.Errorf("Error reading management_port: %v", err)
		}
	}

	if err = d.Set("management_port_use_admin_sport", dataSourceFlattenSystemGlobalManagementPortUseAdminSport(o["management-port-use-admin-sport"], d, "management_port_use_admin_sport")); err != nil {
		if !fortiAPIPatch(o["management-port-use-admin-sport"]) {
			return fmt.Errorf("Error reading management_port_use_admin_sport: %v", err)
		}
	}

	if err = d.Set("forticonverter_integration", dataSourceFlattenSystemGlobalForticonverterIntegration(o["forticonverter-integration"], d, "forticonverter_integration")); err != nil {
		if !fortiAPIPatch(o["forticonverter-integration"]) {
			return fmt.Errorf("Error reading forticonverter_integration: %v", err)
		}
	}

	if err = d.Set("forticonverter_config_upload", dataSourceFlattenSystemGlobalForticonverterConfigUpload(o["forticonverter-config-upload"], d, "forticonverter_config_upload")); err != nil {
		if !fortiAPIPatch(o["forticonverter-config-upload"]) {
			return fmt.Errorf("Error reading forticonverter_config_upload: %v", err)
		}
	}

	if err = d.Set("internet_service_database", dataSourceFlattenSystemGlobalInternetServiceDatabase(o["internet-service-database"], d, "internet_service_database")); err != nil {
		if !fortiAPIPatch(o["internet-service-database"]) {
			return fmt.Errorf("Error reading internet_service_database: %v", err)
		}
	}

	if err = d.Set("internet_service_download_list", dataSourceFlattenSystemGlobalInternetServiceDownloadList(o["internet-service-download-list"], d, "internet_service_download_list")); err != nil {
		if !fortiAPIPatch(o["internet-service-download-list"]) {
			return fmt.Errorf("Error reading internet_service_download_list: %v", err)
		}
	}

	if err = d.Set("geoip_full_db", dataSourceFlattenSystemGlobalGeoipFullDb(o["geoip-full-db"], d, "geoip_full_db")); err != nil {
		if !fortiAPIPatch(o["geoip-full-db"]) {
			return fmt.Errorf("Error reading geoip_full_db: %v", err)
		}
	}

	if err = d.Set("early_tcp_npu_session", dataSourceFlattenSystemGlobalEarlyTcpNpuSession(o["early-tcp-npu-session"], d, "early_tcp_npu_session")); err != nil {
		if !fortiAPIPatch(o["early-tcp-npu-session"]) {
			return fmt.Errorf("Error reading early_tcp_npu_session: %v", err)
		}
	}

	if err = d.Set("npu_neighbor_update", dataSourceFlattenSystemGlobalNpuNeighborUpdate(o["npu-neighbor-update"], d, "npu_neighbor_update")); err != nil {
		if !fortiAPIPatch(o["npu-neighbor-update"]) {
			return fmt.Errorf("Error reading npu_neighbor_update: %v", err)
		}
	}

	if err = d.Set("delay_tcp_npu_session", dataSourceFlattenSystemGlobalDelayTcpNpuSession(o["delay-tcp-npu-session"], d, "delay_tcp_npu_session")); err != nil {
		if !fortiAPIPatch(o["delay-tcp-npu-session"]) {
			return fmt.Errorf("Error reading delay_tcp_npu_session: %v", err)
		}
	}

	if err = d.Set("interface_subnet_usage", dataSourceFlattenSystemGlobalInterfaceSubnetUsage(o["interface-subnet-usage"], d, "interface_subnet_usage")); err != nil {
		if !fortiAPIPatch(o["interface-subnet-usage"]) {
			return fmt.Errorf("Error reading interface_subnet_usage: %v", err)
		}
	}

	if err = d.Set("sflowd_max_children_num", dataSourceFlattenSystemGlobalSflowdMaxChildrenNum(o["sflowd-max-children-num"], d, "sflowd_max_children_num")); err != nil {
		if !fortiAPIPatch(o["sflowd-max-children-num"]) {
			return fmt.Errorf("Error reading sflowd_max_children_num: %v", err)
		}
	}

	if err = d.Set("fortigslb_integration", dataSourceFlattenSystemGlobalFortigslbIntegration(o["fortigslb-integration"], d, "fortigslb_integration")); err != nil {
		if !fortiAPIPatch(o["fortigslb-integration"]) {
			return fmt.Errorf("Error reading fortigslb_integration: %v", err)
		}
	}

	if err = d.Set("user_history_password_threshold", dataSourceFlattenSystemGlobalUserHistoryPasswordThreshold(o["user-history-password-threshold"], d, "user_history_password_threshold")); err != nil {
		if !fortiAPIPatch(o["user-history-password-threshold"]) {
			return fmt.Errorf("Error reading user_history_password_threshold: %v", err)
		}
	}

	if err = d.Set("auth_session_auto_backup", dataSourceFlattenSystemGlobalAuthSessionAutoBackup(o["auth-session-auto-backup"], d, "auth_session_auto_backup")); err != nil {
		if !fortiAPIPatch(o["auth-session-auto-backup"]) {
			return fmt.Errorf("Error reading auth_session_auto_backup: %v", err)
		}
	}

	if err = d.Set("auth_session_auto_backup_interval", dataSourceFlattenSystemGlobalAuthSessionAutoBackupInterval(o["auth-session-auto-backup-interval"], d, "auth_session_auto_backup_interval")); err != nil {
		if !fortiAPIPatch(o["auth-session-auto-backup-interval"]) {
			return fmt.Errorf("Error reading auth_session_auto_backup_interval: %v", err)
		}
	}

	if err = d.Set("scim_https_port", dataSourceFlattenSystemGlobalScimHttpsPort(o["scim-https-port"], d, "scim_https_port")); err != nil {
		if !fortiAPIPatch(o["scim-https-port"]) {
			return fmt.Errorf("Error reading scim_https_port: %v", err)
		}
	}

	if err = d.Set("scim_http_port", dataSourceFlattenSystemGlobalScimHttpPort(o["scim-http-port"], d, "scim_http_port")); err != nil {
		if !fortiAPIPatch(o["scim-http-port"]) {
			return fmt.Errorf("Error reading scim_http_port: %v", err)
		}
	}

	if err = d.Set("scim_server_cert", dataSourceFlattenSystemGlobalScimServerCert(o["scim-server-cert"], d, "scim_server_cert")); err != nil {
		if !fortiAPIPatch(o["scim-server-cert"]) {
			return fmt.Errorf("Error reading scim_server_cert: %v", err)
		}
	}

	if err = d.Set("application_bandwidth_tracking", dataSourceFlattenSystemGlobalApplicationBandwidthTracking(o["application-bandwidth-tracking"], d, "application_bandwidth_tracking")); err != nil {
		if !fortiAPIPatch(o["application-bandwidth-tracking"]) {
			return fmt.Errorf("Error reading application_bandwidth_tracking: %v", err)
		}
	}

	if err = d.Set("tls_session_cache", dataSourceFlattenSystemGlobalTlsSessionCache(o["tls-session-cache"], d, "tls_session_cache")); err != nil {
		if !fortiAPIPatch(o["tls-session-cache"]) {
			return fmt.Errorf("Error reading tls_session_cache: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
