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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemGlobalUpdate,
		Read:   resourceSystemGlobalRead,
		Update: resourceSystemGlobalUpdate,
		Delete: resourceSystemGlobalDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"language": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_replacement_message_groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_local_out": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_certificates": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_custom_language": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_wireless_opensecurity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_display_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_fortigate_cloud_sandbox": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_fortisandbox_cloud": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_firmware_upgrade_warning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_firmware_upgrade_setup_warning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_lines_per_page": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(20, 1000),
				Optional:     true,
				Computed:     true,
			},
			"admin_https_ssl_versions": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_https_ssl_ciphersuites": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_https_ssl_banned_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admintimeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 480),
				Optional:     true,
				Computed:     true,
			},
			"admin_console_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(15, 300),
				Optional:     true,
				Computed:     true,
			},
			"ssd_trim_freq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssd_trim_hour": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 23),
				Optional:     true,
				Computed:     true,
			},
			"ssd_trim_min": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 60),
				Optional:     true,
				Computed:     true,
			},
			"ssd_trim_weekday": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssd_trim_date": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 31),
				Optional:     true,
				Computed:     true,
			},
			"admin_concurrent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_lockout_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),
				Optional:     true,
				Computed:     true,
			},
			"admin_lockout_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"refresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"failtime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"daily_restart": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restart_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"admin_login_max": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional:     true,
				Computed:     true,
			},
			"remoteauthtimeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: intBetweenWithZero(1, 300),
				Optional:     true,
				Computed:     true,
			},
			"ldapconntimeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300000),
				Optional:     true,
				Computed:     true,
			},
			"batch_cmdb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_dlpstat_memory": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"multi_factor_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"autorun_log_fsck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_priority_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"anti_replay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_pmtu_icmp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"honor_df": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pmtu_discovery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtual_switch_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"split_port": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"revision_image_auto_backup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"revision_backup_on_logout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_vdom": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"hostname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"gui_allow_default_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_forticare_registration_setup_warning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_cdn_usage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alias": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"strong_crypto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_cbc_cipher": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_hmac_md5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_kex_sha1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_mac_weak": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_static_key_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_kex_algo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_enc_algo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_mac_algo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snat_route_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"speedtest_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cli_audit_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dh_params": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fds_statistics": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fds_statistics_period": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1440),
				Optional:     true,
				Computed:     true,
			},
			"multicast_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mc_ttl_notchange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"asymroute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lldp_transmission": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lldp_reception": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_auth_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),
				Optional:     true,
				Computed:     true,
			},
			"proxy_re_authentication_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_auth_lifetime": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_auth_lifetime_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 65535),
				Optional:     true,
				Computed:     true,
			},
			"proxy_resource_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_cert_use_mgmt_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sys_perf_log_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"check_protocol_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_arp_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reset_sessionless_tcp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_traffic_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_allow_traffic_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strict_dirty_session_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_halfclose_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 86400),
				Optional:     true,
				Computed:     true,
			},
			"tcp_halfopen_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 86400),
				Optional:     true,
				Computed:     true,
			},
			"tcp_timewait_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 300),
				Optional:     true,
				Computed:     true,
			},
			"tcp_rst_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 300),
				Optional:     true,
				Computed:     true,
			},
			"udp_idle_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 86400),
				Optional:     true,
				Computed:     true,
			},
			"block_session_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),
				Optional:     true,
				Computed:     true,
			},
			"ip_src_port_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pre_login_banner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"post_login_banner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tftp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_failopen": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_failopen_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"memory_use_threshold_extreme": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(70, 97),
				Optional:     true,
				Computed:     true,
			},
			"memory_use_threshold_red": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(70, 97),
				Optional:     true,
				Computed:     true,
			},
			"memory_use_threshold_green": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(70, 97),
				Optional:     true,
				Computed:     true,
			},
			"cpu_use_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(50, 99),
				Optional:     true,
				Computed:     true,
			},
			"check_reset_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom_admin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"long_vdom_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"edit_vdom_prompt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"admin_sport": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"admin_https_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_hsts_max_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin_ssh_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_restrict_local": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_ssh_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"admin_ssh_grace_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 3600),
				Optional:     true,
				Computed:     true,
			},
			"admin_ssh_v1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_telnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_telnet_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"admin_forticloud_sso_login": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_service_source_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_maintainer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_server_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"user_server_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"admin_https_pki_required": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_certificate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"wifi_ca_certificate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"auth_http_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"auth_https_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"auth_keepalive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_auth_concurrent": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),
				Optional:     true,
				Computed:     true,
			},
			"auth_session_limit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"clt_cert_req": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiservice_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"endpoint_control_portal_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"endpoint_control_fds_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tp_mc_skip_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cfg_save": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cfg_revert_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reboot_upon_config_restore": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_scp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_rating_result_submission": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_rating_run_on_schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wireless_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wireless_controller_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1024, 49150),
				Optional:     true,
				Computed:     true,
			},
			"fortiextender_data_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1024, 49150),
				Optional:     true,
				Computed:     true,
			},
			"fortiextender": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extender_controller_reserved_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiextender_discovery_lockdown": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiextender_vlan_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_reserved_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsproxy_worker_count": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"url_filter_count": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 4),
				Optional:     true,
				Computed:     true,
			},
			"proxy_worker_count": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"scanunit_count": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"proxy_hardware_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_kxp_hardware_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_cipher_hardware_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgd_alert_subscription": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_hmac_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_accept_dad": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2),
				Optional:     true,
				Computed:     true,
			},
			"ipv6_allow_anycast_probe": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"csr_ca_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wimax_4g_usb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert_chain_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sslvpn_max_worker_count": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),
				Optional:     true,
				Computed:     true,
			},
			"sslvpn_ems_sn_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslvpn_kxp_hardware_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslvpn_cipher_hardware_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslvpn_plugin_version_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"two_factor_ftk_expiry": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 600),
				Optional:     true,
				Computed:     true,
			},
			"two_factor_email_expiry": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 300),
				Optional:     true,
				Computed:     true,
			},
			"two_factor_sms_expiry": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 300),
				Optional:     true,
				Computed:     true,
			},
			"two_factor_fac_expiry": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 3600),
				Optional:     true,
				Computed:     true,
			},
			"two_factor_ftm_expiry": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 168),
				Optional:     true,
				Computed:     true,
			},
			"per_user_bal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"per_user_bwl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtual_server_count": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"virtual_server_hardware_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wad_worker_count": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"wad_csvc_cs_count": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1),
				Optional:     true,
				Computed:     true,
			},
			"wad_csvc_db_count": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"wad_source_affinity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wad_memory_change_granularity": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 25),
				Optional:     true,
				Computed:     true,
			},
			"login_timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"miglogd_children": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"special_file_23_support": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_uuid_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_uuid_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_ssl_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_rest_api_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arp_max_entry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ha_affinity": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"cmdbsvr_affinity": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"av_affinity": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"wad_affinity": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"ips_affinity": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"miglog_affinity": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),
				Optional:     true,
				Computed:     true,
			},
			"url_filter_affinity": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"ndp_max_entry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"br_fdb_max_entry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_route_cache_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipsec_asic_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_soft_dec_async": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike_embryonic_limit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(50, 20000),
				Optional:     true,
				Computed:     true,
			},
			"device_idle_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 31536000),
				Optional:     true,
				Computed:     true,
			},
			"user_device_store_max_devices": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10526, 240381),
				Optional:     true,
				Computed:     true,
			},
			"user_device_store_max_users": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10526, 240381),
				Optional:     true,
				Computed:     true,
			},
			"user_device_store_max_unified_mem": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(21052743, 1682668748),
				Optional:     true,
				Computed:     true,
			},
			"device_identification_active_scan_delay": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(20, 3600),
				Optional:     true,
				Computed:     true,
			},
			"compliance_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"compliance_check_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_device_latitude": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),
				Optional:     true,
				Computed:     true,
			},
			"gui_device_longitude": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),
				Optional:     true,
				Computed:     true,
			},
			"private_data_encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_auth_extension_device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_theme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_date_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_date_time_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"igmp_state_limit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(96, 128000),
				Optional:     true,
				Computed:     true,
			},
			"cloud_communication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fec_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(49152, 65535),
				Optional:     true,
				Computed:     true,
			},
			"ipsec_ha_seqjump_rate": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),
				Optional:     true,
				Computed:     true,
			},
			"fortitoken_cloud": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"faz_disk_buffer_size": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 214748364),
				Optional:     true,
				Computed:     true,
			},
			"irq_time_accounting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiipam_integration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_ip": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"management_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"management_port_use_admin_sport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_database": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemGlobalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemGlobal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemGlobal resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemGlobal(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemGlobal")
	}

	return resourceSystemGlobalRead(d, m)
}

func resourceSystemGlobalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemGlobal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemGlobalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemGlobal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemGlobal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemGlobal resource from API: %v", err)
	}
	return nil
}

func flattenSystemGlobalLanguage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiIpv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiReplacementMessageGroups(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiLocalOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiCertificates(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiCustomLanguage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiWirelessOpensecurity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiDisplayHostname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiFortigateCloudSandbox(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiFortisandboxCloud(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiFirmwareUpgradeWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiFirmwareUpgradeSetupWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiLinesPerPage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminHttpsSslVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminHttpsSslCiphersuites(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminHttpsSslBannedCiphers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdmintimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminConsoleTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSsdTrimFreq(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSsdTrimHour(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSsdTrimMin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSsdTrimWeekday(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSsdTrimDate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminConcurrent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminLockoutThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminLockoutDuration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalRefresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalFailtime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDailyRestart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalRestartTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalRadiusPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminLoginMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalRemoteauthtimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalLdapconntimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalBatchCmdb(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalMaxDlpstatMemory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalMultiFactorAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAutorunLogFsck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTimezone(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTrafficPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTrafficPriorityLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAntiReplay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSendPmtuIcmp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalHonorDf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalPmtuDiscovery(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalVirtualSwitchVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSplitPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalRevisionImageAutoBackup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalRevisionBackupOnLogout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalManagementVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalHostname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiAllowDefaultHostname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiForticareRegistrationSetupWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiCdnUsage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAlias(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalStrongCrypto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSshCbcCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSshHmacMd5(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSshKexSha1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSshMacWeak(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSslStaticKeyCiphers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSshKexAlgo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSshEncAlgo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSshMacAlgo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSnatRouteChange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSpeedtestServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalCliAuditLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDhParams(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalFdsStatistics(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalFdsStatisticsPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalMulticastForward(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalMcTtlNotchange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAsymroute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTcpOption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalLldpTransmission(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalLldpReception(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalProxyAuthTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalProxyReAuthenticationMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalProxyAuthLifetime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalProxyAuthLifetimeTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalProxyResourceMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalProxyCertUseMgmtVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSysPerfLogInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalCheckProtocolHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalVipArpRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalResetSessionlessTcp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAllowTrafficRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalIpv6AllowTrafficRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalStrictDirtySessionCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTcpHalfcloseTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTcpHalfopenTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTcpTimewaitTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTcpRstTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalUdpIdleTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalBlockSessionTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalIpSrcPortRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalPreLoginBanner(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalPostLoginBanner(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTftp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAvFailopen(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAvFailopenSession(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalMemoryUseThresholdExtreme(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalMemoryUseThresholdRed(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalMemoryUseThresholdGreen(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalCpuUseThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalCheckResetRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalVdomMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalVdomAdmin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalLongVdomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalEditVdomPrompt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminSport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminHttpsRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminHstsMaxAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminSshPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminRestrictLocal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminSshPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminSshGraceTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminSshV1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminTelnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminTelnetPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminForticloudSsoLogin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDefaultServiceSourcePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminMaintainer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalUserServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminHttpsPkiRequired(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalWifiCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalWifiCaCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAuthHttpPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAuthHttpsPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAuthKeepalive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalPolicyAuthConcurrent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAuthSessionLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAuthCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalCltCertReq(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalFortiservicePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalEndpointControlPortalPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalEndpointControlFdsAccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTpMcSkipPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalCfgSave(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalCfgRevertTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalRebootUponConfigRestore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAdminScp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSecurityRatingResultSubmission(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSecurityRatingRunOnSchedule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalWirelessController(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalWirelessControllerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalFortiextenderDataPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalFortiextender(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalExtenderControllerReservedNetwork(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalFortiextenderDiscoveryLockdown(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalFortiextenderVlanMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSwitchController(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSwitchControllerReservedNetwork(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemGlobalDnsproxyWorkerCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalUrlFilterCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalProxyWorkerCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalScanunitCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalProxyHardwareAcceleration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalProxyKxpHardwareAcceleration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalProxyCipherHardwareAcceleration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalFgdAlertSubscription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalIpsecHmacOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalIpv6AcceptDad(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalIpv6AllowAnycastProbe(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalCsrCaAttribute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalWimax4GUsb(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalCertChainMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSslvpnMaxWorkerCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSslvpnEmsSnCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSslvpnKxpHardwareAcceleration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSslvpnCipherHardwareAcceleration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSslvpnPluginVersionCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTwoFactorFtkExpiry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTwoFactorEmailExpiry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTwoFactorSmsExpiry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTwoFactorFacExpiry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalTwoFactorFtmExpiry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalPerUserBal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalPerUserBwl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalVirtualServerCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalVirtualServerHardwareAcceleration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalWadWorkerCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalWadCsvcCsCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalWadCsvcDbCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalWadSourceAffinity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalWadMemoryChangeGranularity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalLoginTimestamp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalMiglogdChildren(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalSpecialFile23Support(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalLogUuidPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalLogUuidAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalLogSslConnection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiRestApiCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalArpMaxEntry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalHaAffinity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalCmdbsvrAffinity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAvAffinity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalWadAffinity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalIpsAffinity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalMiglogAffinity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalUrlFilterAffinity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalNdpMaxEntry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalBrFdbMaxEntry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalMaxRouteCacheSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalIpsecAsicOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalIpsecSoftDecAsync(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalIkeEmbryonicLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDeviceIdleTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalUserDeviceStoreMaxDevices(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalUserDeviceStoreMaxUsers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalUserDeviceStoreMaxUnifiedMem(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalDeviceIdentificationActiveScanDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalComplianceCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalComplianceCheckTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiDeviceLatitude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiDeviceLongitude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalPrivateDataEncryption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalAutoAuthExtensionDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiTheme(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiDateFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalGuiDateTimeSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalIgmpStateLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalCloudCommunication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalFecPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalIpsecHaSeqjumpRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalFortitokenCloud(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalFazDiskBufferSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalIrqTimeAccounting(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalFortiipamIntegration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalManagementIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalManagementPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalManagementPortUseAdminSport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGlobalInternetServiceDatabase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemGlobal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("language", flattenSystemGlobalLanguage(o["language"], d, "language", sv)); err != nil {
		if !fortiAPIPatch(o["language"]) {
			return fmt.Errorf("Error reading language: %v", err)
		}
	}

	if err = d.Set("gui_ipv6", flattenSystemGlobalGuiIpv6(o["gui-ipv6"], d, "gui_ipv6", sv)); err != nil {
		if !fortiAPIPatch(o["gui-ipv6"]) {
			return fmt.Errorf("Error reading gui_ipv6: %v", err)
		}
	}

	if err = d.Set("gui_replacement_message_groups", flattenSystemGlobalGuiReplacementMessageGroups(o["gui-replacement-message-groups"], d, "gui_replacement_message_groups", sv)); err != nil {
		if !fortiAPIPatch(o["gui-replacement-message-groups"]) {
			return fmt.Errorf("Error reading gui_replacement_message_groups: %v", err)
		}
	}

	if err = d.Set("gui_local_out", flattenSystemGlobalGuiLocalOut(o["gui-local-out"], d, "gui_local_out", sv)); err != nil {
		if !fortiAPIPatch(o["gui-local-out"]) {
			return fmt.Errorf("Error reading gui_local_out: %v", err)
		}
	}

	if err = d.Set("gui_certificates", flattenSystemGlobalGuiCertificates(o["gui-certificates"], d, "gui_certificates", sv)); err != nil {
		if !fortiAPIPatch(o["gui-certificates"]) {
			return fmt.Errorf("Error reading gui_certificates: %v", err)
		}
	}

	if err = d.Set("gui_custom_language", flattenSystemGlobalGuiCustomLanguage(o["gui-custom-language"], d, "gui_custom_language", sv)); err != nil {
		if !fortiAPIPatch(o["gui-custom-language"]) {
			return fmt.Errorf("Error reading gui_custom_language: %v", err)
		}
	}

	if err = d.Set("gui_wireless_opensecurity", flattenSystemGlobalGuiWirelessOpensecurity(o["gui-wireless-opensecurity"], d, "gui_wireless_opensecurity", sv)); err != nil {
		if !fortiAPIPatch(o["gui-wireless-opensecurity"]) {
			return fmt.Errorf("Error reading gui_wireless_opensecurity: %v", err)
		}
	}

	if err = d.Set("gui_display_hostname", flattenSystemGlobalGuiDisplayHostname(o["gui-display-hostname"], d, "gui_display_hostname", sv)); err != nil {
		if !fortiAPIPatch(o["gui-display-hostname"]) {
			return fmt.Errorf("Error reading gui_display_hostname: %v", err)
		}
	}

	if err = d.Set("gui_fortigate_cloud_sandbox", flattenSystemGlobalGuiFortigateCloudSandbox(o["gui-fortigate-cloud-sandbox"], d, "gui_fortigate_cloud_sandbox", sv)); err != nil {
		if !fortiAPIPatch(o["gui-fortigate-cloud-sandbox"]) {
			return fmt.Errorf("Error reading gui_fortigate_cloud_sandbox: %v", err)
		}
	}

	if err = d.Set("gui_fortisandbox_cloud", flattenSystemGlobalGuiFortisandboxCloud(o["gui-fortisandbox-cloud"], d, "gui_fortisandbox_cloud", sv)); err != nil {
		if !fortiAPIPatch(o["gui-fortisandbox-cloud"]) {
			return fmt.Errorf("Error reading gui_fortisandbox_cloud: %v", err)
		}
	}

	if err = d.Set("gui_firmware_upgrade_warning", flattenSystemGlobalGuiFirmwareUpgradeWarning(o["gui-firmware-upgrade-warning"], d, "gui_firmware_upgrade_warning", sv)); err != nil {
		if !fortiAPIPatch(o["gui-firmware-upgrade-warning"]) {
			return fmt.Errorf("Error reading gui_firmware_upgrade_warning: %v", err)
		}
	}

	if err = d.Set("gui_firmware_upgrade_setup_warning", flattenSystemGlobalGuiFirmwareUpgradeSetupWarning(o["gui-firmware-upgrade-setup-warning"], d, "gui_firmware_upgrade_setup_warning", sv)); err != nil {
		if !fortiAPIPatch(o["gui-firmware-upgrade-setup-warning"]) {
			return fmt.Errorf("Error reading gui_firmware_upgrade_setup_warning: %v", err)
		}
	}

	if err = d.Set("gui_lines_per_page", flattenSystemGlobalGuiLinesPerPage(o["gui-lines-per-page"], d, "gui_lines_per_page", sv)); err != nil {
		if !fortiAPIPatch(o["gui-lines-per-page"]) {
			return fmt.Errorf("Error reading gui_lines_per_page: %v", err)
		}
	}

	if err = d.Set("admin_https_ssl_versions", flattenSystemGlobalAdminHttpsSslVersions(o["admin-https-ssl-versions"], d, "admin_https_ssl_versions", sv)); err != nil {
		if !fortiAPIPatch(o["admin-https-ssl-versions"]) {
			return fmt.Errorf("Error reading admin_https_ssl_versions: %v", err)
		}
	}

	if err = d.Set("admin_https_ssl_ciphersuites", flattenSystemGlobalAdminHttpsSslCiphersuites(o["admin-https-ssl-ciphersuites"], d, "admin_https_ssl_ciphersuites", sv)); err != nil {
		if !fortiAPIPatch(o["admin-https-ssl-ciphersuites"]) {
			return fmt.Errorf("Error reading admin_https_ssl_ciphersuites: %v", err)
		}
	}

	if err = d.Set("admin_https_ssl_banned_ciphers", flattenSystemGlobalAdminHttpsSslBannedCiphers(o["admin-https-ssl-banned-ciphers"], d, "admin_https_ssl_banned_ciphers", sv)); err != nil {
		if !fortiAPIPatch(o["admin-https-ssl-banned-ciphers"]) {
			return fmt.Errorf("Error reading admin_https_ssl_banned_ciphers: %v", err)
		}
	}

	if err = d.Set("admintimeout", flattenSystemGlobalAdmintimeout(o["admintimeout"], d, "admintimeout", sv)); err != nil {
		if !fortiAPIPatch(o["admintimeout"]) {
			return fmt.Errorf("Error reading admintimeout: %v", err)
		}
	}

	if err = d.Set("admin_console_timeout", flattenSystemGlobalAdminConsoleTimeout(o["admin-console-timeout"], d, "admin_console_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["admin-console-timeout"]) {
			return fmt.Errorf("Error reading admin_console_timeout: %v", err)
		}
	}

	if err = d.Set("ssd_trim_freq", flattenSystemGlobalSsdTrimFreq(o["ssd-trim-freq"], d, "ssd_trim_freq", sv)); err != nil {
		if !fortiAPIPatch(o["ssd-trim-freq"]) {
			return fmt.Errorf("Error reading ssd_trim_freq: %v", err)
		}
	}

	if err = d.Set("ssd_trim_hour", flattenSystemGlobalSsdTrimHour(o["ssd-trim-hour"], d, "ssd_trim_hour", sv)); err != nil {
		if !fortiAPIPatch(o["ssd-trim-hour"]) {
			return fmt.Errorf("Error reading ssd_trim_hour: %v", err)
		}
	}

	if err = d.Set("ssd_trim_min", flattenSystemGlobalSsdTrimMin(o["ssd-trim-min"], d, "ssd_trim_min", sv)); err != nil {
		if !fortiAPIPatch(o["ssd-trim-min"]) {
			return fmt.Errorf("Error reading ssd_trim_min: %v", err)
		}
	}

	if err = d.Set("ssd_trim_weekday", flattenSystemGlobalSsdTrimWeekday(o["ssd-trim-weekday"], d, "ssd_trim_weekday", sv)); err != nil {
		if !fortiAPIPatch(o["ssd-trim-weekday"]) {
			return fmt.Errorf("Error reading ssd_trim_weekday: %v", err)
		}
	}

	if err = d.Set("ssd_trim_date", flattenSystemGlobalSsdTrimDate(o["ssd-trim-date"], d, "ssd_trim_date", sv)); err != nil {
		if !fortiAPIPatch(o["ssd-trim-date"]) {
			return fmt.Errorf("Error reading ssd_trim_date: %v", err)
		}
	}

	if err = d.Set("admin_concurrent", flattenSystemGlobalAdminConcurrent(o["admin-concurrent"], d, "admin_concurrent", sv)); err != nil {
		if !fortiAPIPatch(o["admin-concurrent"]) {
			return fmt.Errorf("Error reading admin_concurrent: %v", err)
		}
	}

	if err = d.Set("admin_lockout_threshold", flattenSystemGlobalAdminLockoutThreshold(o["admin-lockout-threshold"], d, "admin_lockout_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["admin-lockout-threshold"]) {
			return fmt.Errorf("Error reading admin_lockout_threshold: %v", err)
		}
	}

	if err = d.Set("admin_lockout_duration", flattenSystemGlobalAdminLockoutDuration(o["admin-lockout-duration"], d, "admin_lockout_duration", sv)); err != nil {
		if !fortiAPIPatch(o["admin-lockout-duration"]) {
			return fmt.Errorf("Error reading admin_lockout_duration: %v", err)
		}
	}

	if err = d.Set("refresh", flattenSystemGlobalRefresh(o["refresh"], d, "refresh", sv)); err != nil {
		if !fortiAPIPatch(o["refresh"]) {
			return fmt.Errorf("Error reading refresh: %v", err)
		}
	}

	if err = d.Set("interval", flattenSystemGlobalInterval(o["interval"], d, "interval", sv)); err != nil {
		if !fortiAPIPatch(o["interval"]) {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("failtime", flattenSystemGlobalFailtime(o["failtime"], d, "failtime", sv)); err != nil {
		if !fortiAPIPatch(o["failtime"]) {
			return fmt.Errorf("Error reading failtime: %v", err)
		}
	}

	if err = d.Set("daily_restart", flattenSystemGlobalDailyRestart(o["daily-restart"], d, "daily_restart", sv)); err != nil {
		if !fortiAPIPatch(o["daily-restart"]) {
			return fmt.Errorf("Error reading daily_restart: %v", err)
		}
	}

	if err = d.Set("restart_time", flattenSystemGlobalRestartTime(o["restart-time"], d, "restart_time", sv)); err != nil {
		if !fortiAPIPatch(o["restart-time"]) {
			return fmt.Errorf("Error reading restart_time: %v", err)
		}
	}

	if err = d.Set("radius_port", flattenSystemGlobalRadiusPort(o["radius-port"], d, "radius_port", sv)); err != nil {
		if !fortiAPIPatch(o["radius-port"]) {
			return fmt.Errorf("Error reading radius_port: %v", err)
		}
	}

	if err = d.Set("admin_login_max", flattenSystemGlobalAdminLoginMax(o["admin-login-max"], d, "admin_login_max", sv)); err != nil {
		if !fortiAPIPatch(o["admin-login-max"]) {
			return fmt.Errorf("Error reading admin_login_max: %v", err)
		}
	}

	if err = d.Set("remoteauthtimeout", flattenSystemGlobalRemoteauthtimeout(o["remoteauthtimeout"], d, "remoteauthtimeout", sv)); err != nil {
		if !fortiAPIPatch(o["remoteauthtimeout"]) {
			return fmt.Errorf("Error reading remoteauthtimeout: %v", err)
		}
	}

	if err = d.Set("ldapconntimeout", flattenSystemGlobalLdapconntimeout(o["ldapconntimeout"], d, "ldapconntimeout", sv)); err != nil {
		if !fortiAPIPatch(o["ldapconntimeout"]) {
			return fmt.Errorf("Error reading ldapconntimeout: %v", err)
		}
	}

	if err = d.Set("batch_cmdb", flattenSystemGlobalBatchCmdb(o["batch-cmdb"], d, "batch_cmdb", sv)); err != nil {
		if !fortiAPIPatch(o["batch-cmdb"]) {
			return fmt.Errorf("Error reading batch_cmdb: %v", err)
		}
	}

	if err = d.Set("max_dlpstat_memory", flattenSystemGlobalMaxDlpstatMemory(o["max-dlpstat-memory"], d, "max_dlpstat_memory", sv)); err != nil {
		if !fortiAPIPatch(o["max-dlpstat-memory"]) {
			return fmt.Errorf("Error reading max_dlpstat_memory: %v", err)
		}
	}

	if err = d.Set("multi_factor_authentication", flattenSystemGlobalMultiFactorAuthentication(o["multi-factor-authentication"], d, "multi_factor_authentication", sv)); err != nil {
		if !fortiAPIPatch(o["multi-factor-authentication"]) {
			return fmt.Errorf("Error reading multi_factor_authentication: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenSystemGlobalSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-version"]) {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("autorun_log_fsck", flattenSystemGlobalAutorunLogFsck(o["autorun-log-fsck"], d, "autorun_log_fsck", sv)); err != nil {
		if !fortiAPIPatch(o["autorun-log-fsck"]) {
			return fmt.Errorf("Error reading autorun_log_fsck: %v", err)
		}
	}

	if err = d.Set("dst", flattenSystemGlobalDst(o["dst"], d, "dst", sv)); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("timezone", flattenSystemGlobalTimezone(o["timezone"], d, "timezone", sv)); err != nil {
		if !fortiAPIPatch(o["timezone"]) {
			return fmt.Errorf("Error reading timezone: %v", err)
		}
	}

	if err = d.Set("traffic_priority", flattenSystemGlobalTrafficPriority(o["traffic-priority"], d, "traffic_priority", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-priority"]) {
			return fmt.Errorf("Error reading traffic_priority: %v", err)
		}
	}

	if err = d.Set("traffic_priority_level", flattenSystemGlobalTrafficPriorityLevel(o["traffic-priority-level"], d, "traffic_priority_level", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-priority-level"]) {
			return fmt.Errorf("Error reading traffic_priority_level: %v", err)
		}
	}

	if err = d.Set("anti_replay", flattenSystemGlobalAntiReplay(o["anti-replay"], d, "anti_replay", sv)); err != nil {
		if !fortiAPIPatch(o["anti-replay"]) {
			return fmt.Errorf("Error reading anti_replay: %v", err)
		}
	}

	if err = d.Set("send_pmtu_icmp", flattenSystemGlobalSendPmtuIcmp(o["send-pmtu-icmp"], d, "send_pmtu_icmp", sv)); err != nil {
		if !fortiAPIPatch(o["send-pmtu-icmp"]) {
			return fmt.Errorf("Error reading send_pmtu_icmp: %v", err)
		}
	}

	if err = d.Set("honor_df", flattenSystemGlobalHonorDf(o["honor-df"], d, "honor_df", sv)); err != nil {
		if !fortiAPIPatch(o["honor-df"]) {
			return fmt.Errorf("Error reading honor_df: %v", err)
		}
	}

	if err = d.Set("pmtu_discovery", flattenSystemGlobalPmtuDiscovery(o["pmtu-discovery"], d, "pmtu_discovery", sv)); err != nil {
		if !fortiAPIPatch(o["pmtu-discovery"]) {
			return fmt.Errorf("Error reading pmtu_discovery: %v", err)
		}
	}

	if err = d.Set("virtual_switch_vlan", flattenSystemGlobalVirtualSwitchVlan(o["virtual-switch-vlan"], d, "virtual_switch_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["virtual-switch-vlan"]) {
			return fmt.Errorf("Error reading virtual_switch_vlan: %v", err)
		}
	}

	if err = d.Set("split_port", flattenSystemGlobalSplitPort(o["split-port"], d, "split_port", sv)); err != nil {
		if !fortiAPIPatch(o["split-port"]) {
			return fmt.Errorf("Error reading split_port: %v", err)
		}
	}

	if err = d.Set("revision_image_auto_backup", flattenSystemGlobalRevisionImageAutoBackup(o["revision-image-auto-backup"], d, "revision_image_auto_backup", sv)); err != nil {
		if !fortiAPIPatch(o["revision-image-auto-backup"]) {
			return fmt.Errorf("Error reading revision_image_auto_backup: %v", err)
		}
	}

	if err = d.Set("revision_backup_on_logout", flattenSystemGlobalRevisionBackupOnLogout(o["revision-backup-on-logout"], d, "revision_backup_on_logout", sv)); err != nil {
		if !fortiAPIPatch(o["revision-backup-on-logout"]) {
			return fmt.Errorf("Error reading revision_backup_on_logout: %v", err)
		}
	}

	if err = d.Set("management_vdom", flattenSystemGlobalManagementVdom(o["management-vdom"], d, "management_vdom", sv)); err != nil {
		if !fortiAPIPatch(o["management-vdom"]) {
			return fmt.Errorf("Error reading management_vdom: %v", err)
		}
	}

	if err = d.Set("hostname", flattenSystemGlobalHostname(o["hostname"], d, "hostname", sv)); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("gui_allow_default_hostname", flattenSystemGlobalGuiAllowDefaultHostname(o["gui-allow-default-hostname"], d, "gui_allow_default_hostname", sv)); err != nil {
		if !fortiAPIPatch(o["gui-allow-default-hostname"]) {
			return fmt.Errorf("Error reading gui_allow_default_hostname: %v", err)
		}
	}

	if err = d.Set("gui_forticare_registration_setup_warning", flattenSystemGlobalGuiForticareRegistrationSetupWarning(o["gui-forticare-registration-setup-warning"], d, "gui_forticare_registration_setup_warning", sv)); err != nil {
		if !fortiAPIPatch(o["gui-forticare-registration-setup-warning"]) {
			return fmt.Errorf("Error reading gui_forticare_registration_setup_warning: %v", err)
		}
	}

	if err = d.Set("gui_cdn_usage", flattenSystemGlobalGuiCdnUsage(o["gui-cdn-usage"], d, "gui_cdn_usage", sv)); err != nil {
		if !fortiAPIPatch(o["gui-cdn-usage"]) {
			return fmt.Errorf("Error reading gui_cdn_usage: %v", err)
		}
	}

	if err = d.Set("alias", flattenSystemGlobalAlias(o["alias"], d, "alias", sv)); err != nil {
		if !fortiAPIPatch(o["alias"]) {
			return fmt.Errorf("Error reading alias: %v", err)
		}
	}

	if err = d.Set("strong_crypto", flattenSystemGlobalStrongCrypto(o["strong-crypto"], d, "strong_crypto", sv)); err != nil {
		if !fortiAPIPatch(o["strong-crypto"]) {
			return fmt.Errorf("Error reading strong_crypto: %v", err)
		}
	}

	if err = d.Set("ssh_cbc_cipher", flattenSystemGlobalSshCbcCipher(o["ssh-cbc-cipher"], d, "ssh_cbc_cipher", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-cbc-cipher"]) {
			return fmt.Errorf("Error reading ssh_cbc_cipher: %v", err)
		}
	}

	if err = d.Set("ssh_hmac_md5", flattenSystemGlobalSshHmacMd5(o["ssh-hmac-md5"], d, "ssh_hmac_md5", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-hmac-md5"]) {
			return fmt.Errorf("Error reading ssh_hmac_md5: %v", err)
		}
	}

	if err = d.Set("ssh_kex_sha1", flattenSystemGlobalSshKexSha1(o["ssh-kex-sha1"], d, "ssh_kex_sha1", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-kex-sha1"]) {
			return fmt.Errorf("Error reading ssh_kex_sha1: %v", err)
		}
	}

	if err = d.Set("ssh_mac_weak", flattenSystemGlobalSshMacWeak(o["ssh-mac-weak"], d, "ssh_mac_weak", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-mac-weak"]) {
			return fmt.Errorf("Error reading ssh_mac_weak: %v", err)
		}
	}

	if err = d.Set("ssl_static_key_ciphers", flattenSystemGlobalSslStaticKeyCiphers(o["ssl-static-key-ciphers"], d, "ssl_static_key_ciphers", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-static-key-ciphers"]) {
			return fmt.Errorf("Error reading ssl_static_key_ciphers: %v", err)
		}
	}

	if err = d.Set("ssh_kex_algo", flattenSystemGlobalSshKexAlgo(o["ssh-kex-algo"], d, "ssh_kex_algo", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-kex-algo"]) {
			return fmt.Errorf("Error reading ssh_kex_algo: %v", err)
		}
	}

	if err = d.Set("ssh_enc_algo", flattenSystemGlobalSshEncAlgo(o["ssh-enc-algo"], d, "ssh_enc_algo", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-enc-algo"]) {
			return fmt.Errorf("Error reading ssh_enc_algo: %v", err)
		}
	}

	if err = d.Set("ssh_mac_algo", flattenSystemGlobalSshMacAlgo(o["ssh-mac-algo"], d, "ssh_mac_algo", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-mac-algo"]) {
			return fmt.Errorf("Error reading ssh_mac_algo: %v", err)
		}
	}

	if err = d.Set("snat_route_change", flattenSystemGlobalSnatRouteChange(o["snat-route-change"], d, "snat_route_change", sv)); err != nil {
		if !fortiAPIPatch(o["snat-route-change"]) {
			return fmt.Errorf("Error reading snat_route_change: %v", err)
		}
	}

	if err = d.Set("speedtest_server", flattenSystemGlobalSpeedtestServer(o["speedtest-server"], d, "speedtest_server", sv)); err != nil {
		if !fortiAPIPatch(o["speedtest-server"]) {
			return fmt.Errorf("Error reading speedtest_server: %v", err)
		}
	}

	if err = d.Set("cli_audit_log", flattenSystemGlobalCliAuditLog(o["cli-audit-log"], d, "cli_audit_log", sv)); err != nil {
		if !fortiAPIPatch(o["cli-audit-log"]) {
			return fmt.Errorf("Error reading cli_audit_log: %v", err)
		}
	}

	if err = d.Set("dh_params", flattenSystemGlobalDhParams(o["dh-params"], d, "dh_params", sv)); err != nil {
		if !fortiAPIPatch(o["dh-params"]) {
			return fmt.Errorf("Error reading dh_params: %v", err)
		}
	}

	if err = d.Set("fds_statistics", flattenSystemGlobalFdsStatistics(o["fds-statistics"], d, "fds_statistics", sv)); err != nil {
		if !fortiAPIPatch(o["fds-statistics"]) {
			return fmt.Errorf("Error reading fds_statistics: %v", err)
		}
	}

	if err = d.Set("fds_statistics_period", flattenSystemGlobalFdsStatisticsPeriod(o["fds-statistics-period"], d, "fds_statistics_period", sv)); err != nil {
		if !fortiAPIPatch(o["fds-statistics-period"]) {
			return fmt.Errorf("Error reading fds_statistics_period: %v", err)
		}
	}

	if err = d.Set("multicast_forward", flattenSystemGlobalMulticastForward(o["multicast-forward"], d, "multicast_forward", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-forward"]) {
			return fmt.Errorf("Error reading multicast_forward: %v", err)
		}
	}

	if err = d.Set("mc_ttl_notchange", flattenSystemGlobalMcTtlNotchange(o["mc-ttl-notchange"], d, "mc_ttl_notchange", sv)); err != nil {
		if !fortiAPIPatch(o["mc-ttl-notchange"]) {
			return fmt.Errorf("Error reading mc_ttl_notchange: %v", err)
		}
	}

	if err = d.Set("asymroute", flattenSystemGlobalAsymroute(o["asymroute"], d, "asymroute", sv)); err != nil {
		if !fortiAPIPatch(o["asymroute"]) {
			return fmt.Errorf("Error reading asymroute: %v", err)
		}
	}

	if err = d.Set("tcp_option", flattenSystemGlobalTcpOption(o["tcp-option"], d, "tcp_option", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-option"]) {
			return fmt.Errorf("Error reading tcp_option: %v", err)
		}
	}

	if err = d.Set("lldp_transmission", flattenSystemGlobalLldpTransmission(o["lldp-transmission"], d, "lldp_transmission", sv)); err != nil {
		if !fortiAPIPatch(o["lldp-transmission"]) {
			return fmt.Errorf("Error reading lldp_transmission: %v", err)
		}
	}

	if err = d.Set("lldp_reception", flattenSystemGlobalLldpReception(o["lldp-reception"], d, "lldp_reception", sv)); err != nil {
		if !fortiAPIPatch(o["lldp-reception"]) {
			return fmt.Errorf("Error reading lldp_reception: %v", err)
		}
	}

	if err = d.Set("proxy_auth_timeout", flattenSystemGlobalProxyAuthTimeout(o["proxy-auth-timeout"], d, "proxy_auth_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-auth-timeout"]) {
			return fmt.Errorf("Error reading proxy_auth_timeout: %v", err)
		}
	}

	if err = d.Set("proxy_re_authentication_mode", flattenSystemGlobalProxyReAuthenticationMode(o["proxy-re-authentication-mode"], d, "proxy_re_authentication_mode", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-re-authentication-mode"]) {
			return fmt.Errorf("Error reading proxy_re_authentication_mode: %v", err)
		}
	}

	if err = d.Set("proxy_auth_lifetime", flattenSystemGlobalProxyAuthLifetime(o["proxy-auth-lifetime"], d, "proxy_auth_lifetime", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-auth-lifetime"]) {
			return fmt.Errorf("Error reading proxy_auth_lifetime: %v", err)
		}
	}

	if err = d.Set("proxy_auth_lifetime_timeout", flattenSystemGlobalProxyAuthLifetimeTimeout(o["proxy-auth-lifetime-timeout"], d, "proxy_auth_lifetime_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-auth-lifetime-timeout"]) {
			return fmt.Errorf("Error reading proxy_auth_lifetime_timeout: %v", err)
		}
	}

	if err = d.Set("proxy_resource_mode", flattenSystemGlobalProxyResourceMode(o["proxy-resource-mode"], d, "proxy_resource_mode", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-resource-mode"]) {
			return fmt.Errorf("Error reading proxy_resource_mode: %v", err)
		}
	}

	if err = d.Set("proxy_cert_use_mgmt_vdom", flattenSystemGlobalProxyCertUseMgmtVdom(o["proxy-cert-use-mgmt-vdom"], d, "proxy_cert_use_mgmt_vdom", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-cert-use-mgmt-vdom"]) {
			return fmt.Errorf("Error reading proxy_cert_use_mgmt_vdom: %v", err)
		}
	}

	if err = d.Set("sys_perf_log_interval", flattenSystemGlobalSysPerfLogInterval(o["sys-perf-log-interval"], d, "sys_perf_log_interval", sv)); err != nil {
		if !fortiAPIPatch(o["sys-perf-log-interval"]) {
			return fmt.Errorf("Error reading sys_perf_log_interval: %v", err)
		}
	}

	if err = d.Set("check_protocol_header", flattenSystemGlobalCheckProtocolHeader(o["check-protocol-header"], d, "check_protocol_header", sv)); err != nil {
		if !fortiAPIPatch(o["check-protocol-header"]) {
			return fmt.Errorf("Error reading check_protocol_header: %v", err)
		}
	}

	if err = d.Set("vip_arp_range", flattenSystemGlobalVipArpRange(o["vip-arp-range"], d, "vip_arp_range", sv)); err != nil {
		if !fortiAPIPatch(o["vip-arp-range"]) {
			return fmt.Errorf("Error reading vip_arp_range: %v", err)
		}
	}

	if err = d.Set("reset_sessionless_tcp", flattenSystemGlobalResetSessionlessTcp(o["reset-sessionless-tcp"], d, "reset_sessionless_tcp", sv)); err != nil {
		if !fortiAPIPatch(o["reset-sessionless-tcp"]) {
			return fmt.Errorf("Error reading reset_sessionless_tcp: %v", err)
		}
	}

	if err = d.Set("allow_traffic_redirect", flattenSystemGlobalAllowTrafficRedirect(o["allow-traffic-redirect"], d, "allow_traffic_redirect", sv)); err != nil {
		if !fortiAPIPatch(o["allow-traffic-redirect"]) {
			return fmt.Errorf("Error reading allow_traffic_redirect: %v", err)
		}
	}

	if err = d.Set("ipv6_allow_traffic_redirect", flattenSystemGlobalIpv6AllowTrafficRedirect(o["ipv6-allow-traffic-redirect"], d, "ipv6_allow_traffic_redirect", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-allow-traffic-redirect"]) {
			return fmt.Errorf("Error reading ipv6_allow_traffic_redirect: %v", err)
		}
	}

	if err = d.Set("strict_dirty_session_check", flattenSystemGlobalStrictDirtySessionCheck(o["strict-dirty-session-check"], d, "strict_dirty_session_check", sv)); err != nil {
		if !fortiAPIPatch(o["strict-dirty-session-check"]) {
			return fmt.Errorf("Error reading strict_dirty_session_check: %v", err)
		}
	}

	if err = d.Set("tcp_halfclose_timer", flattenSystemGlobalTcpHalfcloseTimer(o["tcp-halfclose-timer"], d, "tcp_halfclose_timer", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-halfclose-timer"]) {
			return fmt.Errorf("Error reading tcp_halfclose_timer: %v", err)
		}
	}

	if err = d.Set("tcp_halfopen_timer", flattenSystemGlobalTcpHalfopenTimer(o["tcp-halfopen-timer"], d, "tcp_halfopen_timer", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-halfopen-timer"]) {
			return fmt.Errorf("Error reading tcp_halfopen_timer: %v", err)
		}
	}

	if err = d.Set("tcp_timewait_timer", flattenSystemGlobalTcpTimewaitTimer(o["tcp-timewait-timer"], d, "tcp_timewait_timer", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-timewait-timer"]) {
			return fmt.Errorf("Error reading tcp_timewait_timer: %v", err)
		}
	}

	if err = d.Set("tcp_rst_timer", flattenSystemGlobalTcpRstTimer(o["tcp-rst-timer"], d, "tcp_rst_timer", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-rst-timer"]) {
			return fmt.Errorf("Error reading tcp_rst_timer: %v", err)
		}
	}

	if err = d.Set("udp_idle_timer", flattenSystemGlobalUdpIdleTimer(o["udp-idle-timer"], d, "udp_idle_timer", sv)); err != nil {
		if !fortiAPIPatch(o["udp-idle-timer"]) {
			return fmt.Errorf("Error reading udp_idle_timer: %v", err)
		}
	}

	if err = d.Set("block_session_timer", flattenSystemGlobalBlockSessionTimer(o["block-session-timer"], d, "block_session_timer", sv)); err != nil {
		if !fortiAPIPatch(o["block-session-timer"]) {
			return fmt.Errorf("Error reading block_session_timer: %v", err)
		}
	}

	if err = d.Set("ip_src_port_range", flattenSystemGlobalIpSrcPortRange(o["ip-src-port-range"], d, "ip_src_port_range", sv)); err != nil {
		if !fortiAPIPatch(o["ip-src-port-range"]) {
			return fmt.Errorf("Error reading ip_src_port_range: %v", err)
		}
	}

	if err = d.Set("pre_login_banner", flattenSystemGlobalPreLoginBanner(o["pre-login-banner"], d, "pre_login_banner", sv)); err != nil {
		if !fortiAPIPatch(o["pre-login-banner"]) {
			return fmt.Errorf("Error reading pre_login_banner: %v", err)
		}
	}

	if err = d.Set("post_login_banner", flattenSystemGlobalPostLoginBanner(o["post-login-banner"], d, "post_login_banner", sv)); err != nil {
		if !fortiAPIPatch(o["post-login-banner"]) {
			return fmt.Errorf("Error reading post_login_banner: %v", err)
		}
	}

	if err = d.Set("tftp", flattenSystemGlobalTftp(o["tftp"], d, "tftp", sv)); err != nil {
		if !fortiAPIPatch(o["tftp"]) {
			return fmt.Errorf("Error reading tftp: %v", err)
		}
	}

	if err = d.Set("av_failopen", flattenSystemGlobalAvFailopen(o["av-failopen"], d, "av_failopen", sv)); err != nil {
		if !fortiAPIPatch(o["av-failopen"]) {
			return fmt.Errorf("Error reading av_failopen: %v", err)
		}
	}

	if err = d.Set("av_failopen_session", flattenSystemGlobalAvFailopenSession(o["av-failopen-session"], d, "av_failopen_session", sv)); err != nil {
		if !fortiAPIPatch(o["av-failopen-session"]) {
			return fmt.Errorf("Error reading av_failopen_session: %v", err)
		}
	}

	if err = d.Set("memory_use_threshold_extreme", flattenSystemGlobalMemoryUseThresholdExtreme(o["memory-use-threshold-extreme"], d, "memory_use_threshold_extreme", sv)); err != nil {
		if !fortiAPIPatch(o["memory-use-threshold-extreme"]) {
			return fmt.Errorf("Error reading memory_use_threshold_extreme: %v", err)
		}
	}

	if err = d.Set("memory_use_threshold_red", flattenSystemGlobalMemoryUseThresholdRed(o["memory-use-threshold-red"], d, "memory_use_threshold_red", sv)); err != nil {
		if !fortiAPIPatch(o["memory-use-threshold-red"]) {
			return fmt.Errorf("Error reading memory_use_threshold_red: %v", err)
		}
	}

	if err = d.Set("memory_use_threshold_green", flattenSystemGlobalMemoryUseThresholdGreen(o["memory-use-threshold-green"], d, "memory_use_threshold_green", sv)); err != nil {
		if !fortiAPIPatch(o["memory-use-threshold-green"]) {
			return fmt.Errorf("Error reading memory_use_threshold_green: %v", err)
		}
	}

	if err = d.Set("cpu_use_threshold", flattenSystemGlobalCpuUseThreshold(o["cpu-use-threshold"], d, "cpu_use_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["cpu-use-threshold"]) {
			return fmt.Errorf("Error reading cpu_use_threshold: %v", err)
		}
	}

	if err = d.Set("check_reset_range", flattenSystemGlobalCheckResetRange(o["check-reset-range"], d, "check_reset_range", sv)); err != nil {
		if !fortiAPIPatch(o["check-reset-range"]) {
			return fmt.Errorf("Error reading check_reset_range: %v", err)
		}
	}

	if err = d.Set("vdom_mode", flattenSystemGlobalVdomMode(o["vdom-mode"], d, "vdom_mode", sv)); err != nil {
		if !fortiAPIPatch(o["vdom-mode"]) {
			return fmt.Errorf("Error reading vdom_mode: %v", err)
		}
	}

	if err = d.Set("vdom_admin", flattenSystemGlobalVdomAdmin(o["vdom-admin"], d, "vdom_admin", sv)); err != nil {
		if !fortiAPIPatch(o["vdom-admin"]) {
			return fmt.Errorf("Error reading vdom_admin: %v", err)
		}
	}

	if err = d.Set("long_vdom_name", flattenSystemGlobalLongVdomName(o["long-vdom-name"], d, "long_vdom_name", sv)); err != nil {
		if !fortiAPIPatch(o["long-vdom-name"]) {
			return fmt.Errorf("Error reading long_vdom_name: %v", err)
		}
	}

	if err = d.Set("edit_vdom_prompt", flattenSystemGlobalEditVdomPrompt(o["edit-vdom-prompt"], d, "edit_vdom_prompt", sv)); err != nil {
		if !fortiAPIPatch(o["edit-vdom-prompt"]) {
			return fmt.Errorf("Error reading edit_vdom_prompt: %v", err)
		}
	}

	if err = d.Set("admin_port", flattenSystemGlobalAdminPort(o["admin-port"], d, "admin_port", sv)); err != nil {
		if !fortiAPIPatch(o["admin-port"]) {
			return fmt.Errorf("Error reading admin_port: %v", err)
		}
	}

	if err = d.Set("admin_sport", flattenSystemGlobalAdminSport(o["admin-sport"], d, "admin_sport", sv)); err != nil {
		if !fortiAPIPatch(o["admin-sport"]) {
			return fmt.Errorf("Error reading admin_sport: %v", err)
		}
	}

	if err = d.Set("admin_https_redirect", flattenSystemGlobalAdminHttpsRedirect(o["admin-https-redirect"], d, "admin_https_redirect", sv)); err != nil {
		if !fortiAPIPatch(o["admin-https-redirect"]) {
			return fmt.Errorf("Error reading admin_https_redirect: %v", err)
		}
	}

	if err = d.Set("admin_hsts_max_age", flattenSystemGlobalAdminHstsMaxAge(o["admin-hsts-max-age"], d, "admin_hsts_max_age", sv)); err != nil {
		if !fortiAPIPatch(o["admin-hsts-max-age"]) {
			return fmt.Errorf("Error reading admin_hsts_max_age: %v", err)
		}
	}

	if err = d.Set("admin_ssh_password", flattenSystemGlobalAdminSshPassword(o["admin-ssh-password"], d, "admin_ssh_password", sv)); err != nil {
		if !fortiAPIPatch(o["admin-ssh-password"]) {
			return fmt.Errorf("Error reading admin_ssh_password: %v", err)
		}
	}

	if err = d.Set("admin_restrict_local", flattenSystemGlobalAdminRestrictLocal(o["admin-restrict-local"], d, "admin_restrict_local", sv)); err != nil {
		if !fortiAPIPatch(o["admin-restrict-local"]) {
			return fmt.Errorf("Error reading admin_restrict_local: %v", err)
		}
	}

	if err = d.Set("admin_ssh_port", flattenSystemGlobalAdminSshPort(o["admin-ssh-port"], d, "admin_ssh_port", sv)); err != nil {
		if !fortiAPIPatch(o["admin-ssh-port"]) {
			return fmt.Errorf("Error reading admin_ssh_port: %v", err)
		}
	}

	if err = d.Set("admin_ssh_grace_time", flattenSystemGlobalAdminSshGraceTime(o["admin-ssh-grace-time"], d, "admin_ssh_grace_time", sv)); err != nil {
		if !fortiAPIPatch(o["admin-ssh-grace-time"]) {
			return fmt.Errorf("Error reading admin_ssh_grace_time: %v", err)
		}
	}

	if err = d.Set("admin_ssh_v1", flattenSystemGlobalAdminSshV1(o["admin-ssh-v1"], d, "admin_ssh_v1", sv)); err != nil {
		if !fortiAPIPatch(o["admin-ssh-v1"]) {
			return fmt.Errorf("Error reading admin_ssh_v1: %v", err)
		}
	}

	if err = d.Set("admin_telnet", flattenSystemGlobalAdminTelnet(o["admin-telnet"], d, "admin_telnet", sv)); err != nil {
		if !fortiAPIPatch(o["admin-telnet"]) {
			return fmt.Errorf("Error reading admin_telnet: %v", err)
		}
	}

	if err = d.Set("admin_telnet_port", flattenSystemGlobalAdminTelnetPort(o["admin-telnet-port"], d, "admin_telnet_port", sv)); err != nil {
		if !fortiAPIPatch(o["admin-telnet-port"]) {
			return fmt.Errorf("Error reading admin_telnet_port: %v", err)
		}
	}

	if err = d.Set("admin_forticloud_sso_login", flattenSystemGlobalAdminForticloudSsoLogin(o["admin-forticloud-sso-login"], d, "admin_forticloud_sso_login", sv)); err != nil {
		if !fortiAPIPatch(o["admin-forticloud-sso-login"]) {
			return fmt.Errorf("Error reading admin_forticloud_sso_login: %v", err)
		}
	}

	if err = d.Set("default_service_source_port", flattenSystemGlobalDefaultServiceSourcePort(o["default-service-source-port"], d, "default_service_source_port", sv)); err != nil {
		if !fortiAPIPatch(o["default-service-source-port"]) {
			return fmt.Errorf("Error reading default_service_source_port: %v", err)
		}
	}

	if err = d.Set("admin_maintainer", flattenSystemGlobalAdminMaintainer(o["admin-maintainer"], d, "admin_maintainer", sv)); err != nil {
		if !fortiAPIPatch(o["admin-maintainer"]) {
			return fmt.Errorf("Error reading admin_maintainer: %v", err)
		}
	}

	if err = d.Set("admin_server_cert", flattenSystemGlobalAdminServerCert(o["admin-server-cert"], d, "admin_server_cert", sv)); err != nil {
		if !fortiAPIPatch(o["admin-server-cert"]) {
			return fmt.Errorf("Error reading admin_server_cert: %v", err)
		}
	}

	if err = d.Set("user_server_cert", flattenSystemGlobalUserServerCert(o["user-server-cert"], d, "user_server_cert", sv)); err != nil {
		if !fortiAPIPatch(o["user-server-cert"]) {
			return fmt.Errorf("Error reading user_server_cert: %v", err)
		}
	}

	if err = d.Set("admin_https_pki_required", flattenSystemGlobalAdminHttpsPkiRequired(o["admin-https-pki-required"], d, "admin_https_pki_required", sv)); err != nil {
		if !fortiAPIPatch(o["admin-https-pki-required"]) {
			return fmt.Errorf("Error reading admin_https_pki_required: %v", err)
		}
	}

	if err = d.Set("wifi_certificate", flattenSystemGlobalWifiCertificate(o["wifi-certificate"], d, "wifi_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["wifi-certificate"]) {
			return fmt.Errorf("Error reading wifi_certificate: %v", err)
		}
	}

	if err = d.Set("wifi_ca_certificate", flattenSystemGlobalWifiCaCertificate(o["wifi-ca-certificate"], d, "wifi_ca_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["wifi-ca-certificate"]) {
			return fmt.Errorf("Error reading wifi_ca_certificate: %v", err)
		}
	}

	if err = d.Set("auth_http_port", flattenSystemGlobalAuthHttpPort(o["auth-http-port"], d, "auth_http_port", sv)); err != nil {
		if !fortiAPIPatch(o["auth-http-port"]) {
			return fmt.Errorf("Error reading auth_http_port: %v", err)
		}
	}

	if err = d.Set("auth_https_port", flattenSystemGlobalAuthHttpsPort(o["auth-https-port"], d, "auth_https_port", sv)); err != nil {
		if !fortiAPIPatch(o["auth-https-port"]) {
			return fmt.Errorf("Error reading auth_https_port: %v", err)
		}
	}

	if err = d.Set("auth_keepalive", flattenSystemGlobalAuthKeepalive(o["auth-keepalive"], d, "auth_keepalive", sv)); err != nil {
		if !fortiAPIPatch(o["auth-keepalive"]) {
			return fmt.Errorf("Error reading auth_keepalive: %v", err)
		}
	}

	if err = d.Set("policy_auth_concurrent", flattenSystemGlobalPolicyAuthConcurrent(o["policy-auth-concurrent"], d, "policy_auth_concurrent", sv)); err != nil {
		if !fortiAPIPatch(o["policy-auth-concurrent"]) {
			return fmt.Errorf("Error reading policy_auth_concurrent: %v", err)
		}
	}

	if err = d.Set("auth_session_limit", flattenSystemGlobalAuthSessionLimit(o["auth-session-limit"], d, "auth_session_limit", sv)); err != nil {
		if !fortiAPIPatch(o["auth-session-limit"]) {
			return fmt.Errorf("Error reading auth_session_limit: %v", err)
		}
	}

	if err = d.Set("auth_cert", flattenSystemGlobalAuthCert(o["auth-cert"], d, "auth_cert", sv)); err != nil {
		if !fortiAPIPatch(o["auth-cert"]) {
			return fmt.Errorf("Error reading auth_cert: %v", err)
		}
	}

	if err = d.Set("clt_cert_req", flattenSystemGlobalCltCertReq(o["clt-cert-req"], d, "clt_cert_req", sv)); err != nil {
		if !fortiAPIPatch(o["clt-cert-req"]) {
			return fmt.Errorf("Error reading clt_cert_req: %v", err)
		}
	}

	if err = d.Set("fortiservice_port", flattenSystemGlobalFortiservicePort(o["fortiservice-port"], d, "fortiservice_port", sv)); err != nil {
		if !fortiAPIPatch(o["fortiservice-port"]) {
			return fmt.Errorf("Error reading fortiservice_port: %v", err)
		}
	}

	if err = d.Set("endpoint_control_portal_port", flattenSystemGlobalEndpointControlPortalPort(o["endpoint-control-portal-port"], d, "endpoint_control_portal_port", sv)); err != nil {
		if !fortiAPIPatch(o["endpoint-control-portal-port"]) {
			return fmt.Errorf("Error reading endpoint_control_portal_port: %v", err)
		}
	}

	if err = d.Set("endpoint_control_fds_access", flattenSystemGlobalEndpointControlFdsAccess(o["endpoint-control-fds-access"], d, "endpoint_control_fds_access", sv)); err != nil {
		if !fortiAPIPatch(o["endpoint-control-fds-access"]) {
			return fmt.Errorf("Error reading endpoint_control_fds_access: %v", err)
		}
	}

	if err = d.Set("tp_mc_skip_policy", flattenSystemGlobalTpMcSkipPolicy(o["tp-mc-skip-policy"], d, "tp_mc_skip_policy", sv)); err != nil {
		if !fortiAPIPatch(o["tp-mc-skip-policy"]) {
			return fmt.Errorf("Error reading tp_mc_skip_policy: %v", err)
		}
	}

	if err = d.Set("cfg_save", flattenSystemGlobalCfgSave(o["cfg-save"], d, "cfg_save", sv)); err != nil {
		if !fortiAPIPatch(o["cfg-save"]) {
			return fmt.Errorf("Error reading cfg_save: %v", err)
		}
	}

	if err = d.Set("cfg_revert_timeout", flattenSystemGlobalCfgRevertTimeout(o["cfg-revert-timeout"], d, "cfg_revert_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["cfg-revert-timeout"]) {
			return fmt.Errorf("Error reading cfg_revert_timeout: %v", err)
		}
	}

	if err = d.Set("reboot_upon_config_restore", flattenSystemGlobalRebootUponConfigRestore(o["reboot-upon-config-restore"], d, "reboot_upon_config_restore", sv)); err != nil {
		if !fortiAPIPatch(o["reboot-upon-config-restore"]) {
			return fmt.Errorf("Error reading reboot_upon_config_restore: %v", err)
		}
	}

	if err = d.Set("admin_scp", flattenSystemGlobalAdminScp(o["admin-scp"], d, "admin_scp", sv)); err != nil {
		if !fortiAPIPatch(o["admin-scp"]) {
			return fmt.Errorf("Error reading admin_scp: %v", err)
		}
	}

	if err = d.Set("security_rating_result_submission", flattenSystemGlobalSecurityRatingResultSubmission(o["security-rating-result-submission"], d, "security_rating_result_submission", sv)); err != nil {
		if !fortiAPIPatch(o["security-rating-result-submission"]) {
			return fmt.Errorf("Error reading security_rating_result_submission: %v", err)
		}
	}

	if err = d.Set("security_rating_run_on_schedule", flattenSystemGlobalSecurityRatingRunOnSchedule(o["security-rating-run-on-schedule"], d, "security_rating_run_on_schedule", sv)); err != nil {
		if !fortiAPIPatch(o["security-rating-run-on-schedule"]) {
			return fmt.Errorf("Error reading security_rating_run_on_schedule: %v", err)
		}
	}

	if err = d.Set("wireless_controller", flattenSystemGlobalWirelessController(o["wireless-controller"], d, "wireless_controller", sv)); err != nil {
		if !fortiAPIPatch(o["wireless-controller"]) {
			return fmt.Errorf("Error reading wireless_controller: %v", err)
		}
	}

	if err = d.Set("wireless_controller_port", flattenSystemGlobalWirelessControllerPort(o["wireless-controller-port"], d, "wireless_controller_port", sv)); err != nil {
		if !fortiAPIPatch(o["wireless-controller-port"]) {
			return fmt.Errorf("Error reading wireless_controller_port: %v", err)
		}
	}

	if err = d.Set("fortiextender_data_port", flattenSystemGlobalFortiextenderDataPort(o["fortiextender-data-port"], d, "fortiextender_data_port", sv)); err != nil {
		if !fortiAPIPatch(o["fortiextender-data-port"]) {
			return fmt.Errorf("Error reading fortiextender_data_port: %v", err)
		}
	}

	if err = d.Set("fortiextender", flattenSystemGlobalFortiextender(o["fortiextender"], d, "fortiextender", sv)); err != nil {
		if !fortiAPIPatch(o["fortiextender"]) {
			return fmt.Errorf("Error reading fortiextender: %v", err)
		}
	}

	if err = d.Set("extender_controller_reserved_network", flattenSystemGlobalExtenderControllerReservedNetwork(o["extender-controller-reserved-network"], d, "extender_controller_reserved_network", sv)); err != nil {
		if !fortiAPIPatch(o["extender-controller-reserved-network"]) {
			return fmt.Errorf("Error reading extender_controller_reserved_network: %v", err)
		}
	}

	if err = d.Set("fortiextender_discovery_lockdown", flattenSystemGlobalFortiextenderDiscoveryLockdown(o["fortiextender-discovery-lockdown"], d, "fortiextender_discovery_lockdown", sv)); err != nil {
		if !fortiAPIPatch(o["fortiextender-discovery-lockdown"]) {
			return fmt.Errorf("Error reading fortiextender_discovery_lockdown: %v", err)
		}
	}

	if err = d.Set("fortiextender_vlan_mode", flattenSystemGlobalFortiextenderVlanMode(o["fortiextender-vlan-mode"], d, "fortiextender_vlan_mode", sv)); err != nil {
		if !fortiAPIPatch(o["fortiextender-vlan-mode"]) {
			return fmt.Errorf("Error reading fortiextender_vlan_mode: %v", err)
		}
	}

	if err = d.Set("switch_controller", flattenSystemGlobalSwitchController(o["switch-controller"], d, "switch_controller", sv)); err != nil {
		if !fortiAPIPatch(o["switch-controller"]) {
			return fmt.Errorf("Error reading switch_controller: %v", err)
		}
	}

	if err = d.Set("switch_controller_reserved_network", flattenSystemGlobalSwitchControllerReservedNetwork(o["switch-controller-reserved-network"], d, "switch_controller_reserved_network", sv)); err != nil {
		if !fortiAPIPatch(o["switch-controller-reserved-network"]) {
			return fmt.Errorf("Error reading switch_controller_reserved_network: %v", err)
		}
	}

	if err = d.Set("dnsproxy_worker_count", flattenSystemGlobalDnsproxyWorkerCount(o["dnsproxy-worker-count"], d, "dnsproxy_worker_count", sv)); err != nil {
		if !fortiAPIPatch(o["dnsproxy-worker-count"]) {
			return fmt.Errorf("Error reading dnsproxy_worker_count: %v", err)
		}
	}

	if err = d.Set("url_filter_count", flattenSystemGlobalUrlFilterCount(o["url-filter-count"], d, "url_filter_count", sv)); err != nil {
		if !fortiAPIPatch(o["url-filter-count"]) {
			return fmt.Errorf("Error reading url_filter_count: %v", err)
		}
	}

	if err = d.Set("proxy_worker_count", flattenSystemGlobalProxyWorkerCount(o["proxy-worker-count"], d, "proxy_worker_count", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-worker-count"]) {
			return fmt.Errorf("Error reading proxy_worker_count: %v", err)
		}
	}

	if err = d.Set("scanunit_count", flattenSystemGlobalScanunitCount(o["scanunit-count"], d, "scanunit_count", sv)); err != nil {
		if !fortiAPIPatch(o["scanunit-count"]) {
			return fmt.Errorf("Error reading scanunit_count: %v", err)
		}
	}

	if err = d.Set("proxy_hardware_acceleration", flattenSystemGlobalProxyHardwareAcceleration(o["proxy-hardware-acceleration"], d, "proxy_hardware_acceleration", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-hardware-acceleration"]) {
			return fmt.Errorf("Error reading proxy_hardware_acceleration: %v", err)
		}
	}

	if err = d.Set("proxy_kxp_hardware_acceleration", flattenSystemGlobalProxyKxpHardwareAcceleration(o["proxy-kxp-hardware-acceleration"], d, "proxy_kxp_hardware_acceleration", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-kxp-hardware-acceleration"]) {
			return fmt.Errorf("Error reading proxy_kxp_hardware_acceleration: %v", err)
		}
	}

	if err = d.Set("proxy_cipher_hardware_acceleration", flattenSystemGlobalProxyCipherHardwareAcceleration(o["proxy-cipher-hardware-acceleration"], d, "proxy_cipher_hardware_acceleration", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-cipher-hardware-acceleration"]) {
			return fmt.Errorf("Error reading proxy_cipher_hardware_acceleration: %v", err)
		}
	}

	if err = d.Set("fgd_alert_subscription", flattenSystemGlobalFgdAlertSubscription(o["fgd-alert-subscription"], d, "fgd_alert_subscription", sv)); err != nil {
		if !fortiAPIPatch(o["fgd-alert-subscription"]) {
			return fmt.Errorf("Error reading fgd_alert_subscription: %v", err)
		}
	}

	if err = d.Set("ipsec_hmac_offload", flattenSystemGlobalIpsecHmacOffload(o["ipsec-hmac-offload"], d, "ipsec_hmac_offload", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-hmac-offload"]) {
			return fmt.Errorf("Error reading ipsec_hmac_offload: %v", err)
		}
	}

	if err = d.Set("ipv6_accept_dad", flattenSystemGlobalIpv6AcceptDad(o["ipv6-accept-dad"], d, "ipv6_accept_dad", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-accept-dad"]) {
			return fmt.Errorf("Error reading ipv6_accept_dad: %v", err)
		}
	}

	if err = d.Set("ipv6_allow_anycast_probe", flattenSystemGlobalIpv6AllowAnycastProbe(o["ipv6-allow-anycast-probe"], d, "ipv6_allow_anycast_probe", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-allow-anycast-probe"]) {
			return fmt.Errorf("Error reading ipv6_allow_anycast_probe: %v", err)
		}
	}

	if err = d.Set("csr_ca_attribute", flattenSystemGlobalCsrCaAttribute(o["csr-ca-attribute"], d, "csr_ca_attribute", sv)); err != nil {
		if !fortiAPIPatch(o["csr-ca-attribute"]) {
			return fmt.Errorf("Error reading csr_ca_attribute: %v", err)
		}
	}

	if err = d.Set("wimax_4g_usb", flattenSystemGlobalWimax4GUsb(o["wimax-4g-usb"], d, "wimax_4g_usb", sv)); err != nil {
		if !fortiAPIPatch(o["wimax-4g-usb"]) {
			return fmt.Errorf("Error reading wimax_4g_usb: %v", err)
		}
	}

	if err = d.Set("cert_chain_max", flattenSystemGlobalCertChainMax(o["cert-chain-max"], d, "cert_chain_max", sv)); err != nil {
		if !fortiAPIPatch(o["cert-chain-max"]) {
			return fmt.Errorf("Error reading cert_chain_max: %v", err)
		}
	}

	if err = d.Set("sslvpn_max_worker_count", flattenSystemGlobalSslvpnMaxWorkerCount(o["sslvpn-max-worker-count"], d, "sslvpn_max_worker_count", sv)); err != nil {
		if !fortiAPIPatch(o["sslvpn-max-worker-count"]) {
			return fmt.Errorf("Error reading sslvpn_max_worker_count: %v", err)
		}
	}

	if err = d.Set("sslvpn_ems_sn_check", flattenSystemGlobalSslvpnEmsSnCheck(o["sslvpn-ems-sn-check"], d, "sslvpn_ems_sn_check", sv)); err != nil {
		if !fortiAPIPatch(o["sslvpn-ems-sn-check"]) {
			return fmt.Errorf("Error reading sslvpn_ems_sn_check: %v", err)
		}
	}

	if err = d.Set("sslvpn_kxp_hardware_acceleration", flattenSystemGlobalSslvpnKxpHardwareAcceleration(o["sslvpn-kxp-hardware-acceleration"], d, "sslvpn_kxp_hardware_acceleration", sv)); err != nil {
		if !fortiAPIPatch(o["sslvpn-kxp-hardware-acceleration"]) {
			return fmt.Errorf("Error reading sslvpn_kxp_hardware_acceleration: %v", err)
		}
	}

	if err = d.Set("sslvpn_cipher_hardware_acceleration", flattenSystemGlobalSslvpnCipherHardwareAcceleration(o["sslvpn-cipher-hardware-acceleration"], d, "sslvpn_cipher_hardware_acceleration", sv)); err != nil {
		if !fortiAPIPatch(o["sslvpn-cipher-hardware-acceleration"]) {
			return fmt.Errorf("Error reading sslvpn_cipher_hardware_acceleration: %v", err)
		}
	}

	if err = d.Set("sslvpn_plugin_version_check", flattenSystemGlobalSslvpnPluginVersionCheck(o["sslvpn-plugin-version-check"], d, "sslvpn_plugin_version_check", sv)); err != nil {
		if !fortiAPIPatch(o["sslvpn-plugin-version-check"]) {
			return fmt.Errorf("Error reading sslvpn_plugin_version_check: %v", err)
		}
	}

	if err = d.Set("two_factor_ftk_expiry", flattenSystemGlobalTwoFactorFtkExpiry(o["two-factor-ftk-expiry"], d, "two_factor_ftk_expiry", sv)); err != nil {
		if !fortiAPIPatch(o["two-factor-ftk-expiry"]) {
			return fmt.Errorf("Error reading two_factor_ftk_expiry: %v", err)
		}
	}

	if err = d.Set("two_factor_email_expiry", flattenSystemGlobalTwoFactorEmailExpiry(o["two-factor-email-expiry"], d, "two_factor_email_expiry", sv)); err != nil {
		if !fortiAPIPatch(o["two-factor-email-expiry"]) {
			return fmt.Errorf("Error reading two_factor_email_expiry: %v", err)
		}
	}

	if err = d.Set("two_factor_sms_expiry", flattenSystemGlobalTwoFactorSmsExpiry(o["two-factor-sms-expiry"], d, "two_factor_sms_expiry", sv)); err != nil {
		if !fortiAPIPatch(o["two-factor-sms-expiry"]) {
			return fmt.Errorf("Error reading two_factor_sms_expiry: %v", err)
		}
	}

	if err = d.Set("two_factor_fac_expiry", flattenSystemGlobalTwoFactorFacExpiry(o["two-factor-fac-expiry"], d, "two_factor_fac_expiry", sv)); err != nil {
		if !fortiAPIPatch(o["two-factor-fac-expiry"]) {
			return fmt.Errorf("Error reading two_factor_fac_expiry: %v", err)
		}
	}

	if err = d.Set("two_factor_ftm_expiry", flattenSystemGlobalTwoFactorFtmExpiry(o["two-factor-ftm-expiry"], d, "two_factor_ftm_expiry", sv)); err != nil {
		if !fortiAPIPatch(o["two-factor-ftm-expiry"]) {
			return fmt.Errorf("Error reading two_factor_ftm_expiry: %v", err)
		}
	}

	if err = d.Set("per_user_bal", flattenSystemGlobalPerUserBal(o["per-user-bal"], d, "per_user_bal", sv)); err != nil {
		if !fortiAPIPatch(o["per-user-bal"]) {
			return fmt.Errorf("Error reading per_user_bal: %v", err)
		}
	}

	if err = d.Set("per_user_bwl", flattenSystemGlobalPerUserBwl(o["per-user-bwl"], d, "per_user_bwl", sv)); err != nil {
		if !fortiAPIPatch(o["per-user-bwl"]) {
			return fmt.Errorf("Error reading per_user_bwl: %v", err)
		}
	}

	if err = d.Set("virtual_server_count", flattenSystemGlobalVirtualServerCount(o["virtual-server-count"], d, "virtual_server_count", sv)); err != nil {
		if !fortiAPIPatch(o["virtual-server-count"]) {
			return fmt.Errorf("Error reading virtual_server_count: %v", err)
		}
	}

	if err = d.Set("virtual_server_hardware_acceleration", flattenSystemGlobalVirtualServerHardwareAcceleration(o["virtual-server-hardware-acceleration"], d, "virtual_server_hardware_acceleration", sv)); err != nil {
		if !fortiAPIPatch(o["virtual-server-hardware-acceleration"]) {
			return fmt.Errorf("Error reading virtual_server_hardware_acceleration: %v", err)
		}
	}

	if err = d.Set("wad_worker_count", flattenSystemGlobalWadWorkerCount(o["wad-worker-count"], d, "wad_worker_count", sv)); err != nil {
		if !fortiAPIPatch(o["wad-worker-count"]) {
			return fmt.Errorf("Error reading wad_worker_count: %v", err)
		}
	}

	if err = d.Set("wad_csvc_cs_count", flattenSystemGlobalWadCsvcCsCount(o["wad-csvc-cs-count"], d, "wad_csvc_cs_count", sv)); err != nil {
		if !fortiAPIPatch(o["wad-csvc-cs-count"]) {
			return fmt.Errorf("Error reading wad_csvc_cs_count: %v", err)
		}
	}

	if err = d.Set("wad_csvc_db_count", flattenSystemGlobalWadCsvcDbCount(o["wad-csvc-db-count"], d, "wad_csvc_db_count", sv)); err != nil {
		if !fortiAPIPatch(o["wad-csvc-db-count"]) {
			return fmt.Errorf("Error reading wad_csvc_db_count: %v", err)
		}
	}

	if err = d.Set("wad_source_affinity", flattenSystemGlobalWadSourceAffinity(o["wad-source-affinity"], d, "wad_source_affinity", sv)); err != nil {
		if !fortiAPIPatch(o["wad-source-affinity"]) {
			return fmt.Errorf("Error reading wad_source_affinity: %v", err)
		}
	}

	if err = d.Set("wad_memory_change_granularity", flattenSystemGlobalWadMemoryChangeGranularity(o["wad-memory-change-granularity"], d, "wad_memory_change_granularity", sv)); err != nil {
		if !fortiAPIPatch(o["wad-memory-change-granularity"]) {
			return fmt.Errorf("Error reading wad_memory_change_granularity: %v", err)
		}
	}

	if err = d.Set("login_timestamp", flattenSystemGlobalLoginTimestamp(o["login-timestamp"], d, "login_timestamp", sv)); err != nil {
		if !fortiAPIPatch(o["login-timestamp"]) {
			return fmt.Errorf("Error reading login_timestamp: %v", err)
		}
	}

	if err = d.Set("miglogd_children", flattenSystemGlobalMiglogdChildren(o["miglogd-children"], d, "miglogd_children", sv)); err != nil {
		if !fortiAPIPatch(o["miglogd-children"]) {
			return fmt.Errorf("Error reading miglogd_children: %v", err)
		}
	}

	if err = d.Set("special_file_23_support", flattenSystemGlobalSpecialFile23Support(o["special-file-23-support"], d, "special_file_23_support", sv)); err != nil {
		if !fortiAPIPatch(o["special-file-23-support"]) {
			return fmt.Errorf("Error reading special_file_23_support: %v", err)
		}
	}

	if err = d.Set("log_uuid_policy", flattenSystemGlobalLogUuidPolicy(o["log-uuid-policy"], d, "log_uuid_policy", sv)); err != nil {
		if !fortiAPIPatch(o["log-uuid-policy"]) {
			return fmt.Errorf("Error reading log_uuid_policy: %v", err)
		}
	}

	if err = d.Set("log_uuid_address", flattenSystemGlobalLogUuidAddress(o["log-uuid-address"], d, "log_uuid_address", sv)); err != nil {
		if !fortiAPIPatch(o["log-uuid-address"]) {
			return fmt.Errorf("Error reading log_uuid_address: %v", err)
		}
	}

	if err = d.Set("log_ssl_connection", flattenSystemGlobalLogSslConnection(o["log-ssl-connection"], d, "log_ssl_connection", sv)); err != nil {
		if !fortiAPIPatch(o["log-ssl-connection"]) {
			return fmt.Errorf("Error reading log_ssl_connection: %v", err)
		}
	}

	if err = d.Set("gui_rest_api_cache", flattenSystemGlobalGuiRestApiCache(o["gui-rest-api-cache"], d, "gui_rest_api_cache", sv)); err != nil {
		if !fortiAPIPatch(o["gui-rest-api-cache"]) {
			return fmt.Errorf("Error reading gui_rest_api_cache: %v", err)
		}
	}

	if err = d.Set("arp_max_entry", flattenSystemGlobalArpMaxEntry(o["arp-max-entry"], d, "arp_max_entry", sv)); err != nil {
		if !fortiAPIPatch(o["arp-max-entry"]) {
			return fmt.Errorf("Error reading arp_max_entry: %v", err)
		}
	}

	if err = d.Set("ha_affinity", flattenSystemGlobalHaAffinity(o["ha-affinity"], d, "ha_affinity", sv)); err != nil {
		if !fortiAPIPatch(o["ha-affinity"]) {
			return fmt.Errorf("Error reading ha_affinity: %v", err)
		}
	}

	if err = d.Set("cmdbsvr_affinity", flattenSystemGlobalCmdbsvrAffinity(o["cmdbsvr-affinity"], d, "cmdbsvr_affinity", sv)); err != nil {
		if !fortiAPIPatch(o["cmdbsvr-affinity"]) {
			return fmt.Errorf("Error reading cmdbsvr_affinity: %v", err)
		}
	}

	if err = d.Set("av_affinity", flattenSystemGlobalAvAffinity(o["av-affinity"], d, "av_affinity", sv)); err != nil {
		if !fortiAPIPatch(o["av-affinity"]) {
			return fmt.Errorf("Error reading av_affinity: %v", err)
		}
	}

	if err = d.Set("wad_affinity", flattenSystemGlobalWadAffinity(o["wad-affinity"], d, "wad_affinity", sv)); err != nil {
		if !fortiAPIPatch(o["wad-affinity"]) {
			return fmt.Errorf("Error reading wad_affinity: %v", err)
		}
	}

	if err = d.Set("ips_affinity", flattenSystemGlobalIpsAffinity(o["ips-affinity"], d, "ips_affinity", sv)); err != nil {
		if !fortiAPIPatch(o["ips-affinity"]) {
			return fmt.Errorf("Error reading ips_affinity: %v", err)
		}
	}

	if err = d.Set("miglog_affinity", flattenSystemGlobalMiglogAffinity(o["miglog-affinity"], d, "miglog_affinity", sv)); err != nil {
		if !fortiAPIPatch(o["miglog-affinity"]) {
			return fmt.Errorf("Error reading miglog_affinity: %v", err)
		}
	}

	if err = d.Set("url_filter_affinity", flattenSystemGlobalUrlFilterAffinity(o["url-filter-affinity"], d, "url_filter_affinity", sv)); err != nil {
		if !fortiAPIPatch(o["url-filter-affinity"]) {
			return fmt.Errorf("Error reading url_filter_affinity: %v", err)
		}
	}

	if err = d.Set("ndp_max_entry", flattenSystemGlobalNdpMaxEntry(o["ndp-max-entry"], d, "ndp_max_entry", sv)); err != nil {
		if !fortiAPIPatch(o["ndp-max-entry"]) {
			return fmt.Errorf("Error reading ndp_max_entry: %v", err)
		}
	}

	if err = d.Set("br_fdb_max_entry", flattenSystemGlobalBrFdbMaxEntry(o["br-fdb-max-entry"], d, "br_fdb_max_entry", sv)); err != nil {
		if !fortiAPIPatch(o["br-fdb-max-entry"]) {
			return fmt.Errorf("Error reading br_fdb_max_entry: %v", err)
		}
	}

	if err = d.Set("max_route_cache_size", flattenSystemGlobalMaxRouteCacheSize(o["max-route-cache-size"], d, "max_route_cache_size", sv)); err != nil {
		if !fortiAPIPatch(o["max-route-cache-size"]) {
			return fmt.Errorf("Error reading max_route_cache_size: %v", err)
		}
	}

	if err = d.Set("ipsec_asic_offload", flattenSystemGlobalIpsecAsicOffload(o["ipsec-asic-offload"], d, "ipsec_asic_offload", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-asic-offload"]) {
			return fmt.Errorf("Error reading ipsec_asic_offload: %v", err)
		}
	}

	if err = d.Set("ipsec_soft_dec_async", flattenSystemGlobalIpsecSoftDecAsync(o["ipsec-soft-dec-async"], d, "ipsec_soft_dec_async", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-soft-dec-async"]) {
			return fmt.Errorf("Error reading ipsec_soft_dec_async: %v", err)
		}
	}

	if err = d.Set("ike_embryonic_limit", flattenSystemGlobalIkeEmbryonicLimit(o["ike-embryonic-limit"], d, "ike_embryonic_limit", sv)); err != nil {
		if !fortiAPIPatch(o["ike-embryonic-limit"]) {
			return fmt.Errorf("Error reading ike_embryonic_limit: %v", err)
		}
	}

	if err = d.Set("device_idle_timeout", flattenSystemGlobalDeviceIdleTimeout(o["device-idle-timeout"], d, "device_idle_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["device-idle-timeout"]) {
			return fmt.Errorf("Error reading device_idle_timeout: %v", err)
		}
	}

	if err = d.Set("user_device_store_max_devices", flattenSystemGlobalUserDeviceStoreMaxDevices(o["user-device-store-max-devices"], d, "user_device_store_max_devices", sv)); err != nil {
		if !fortiAPIPatch(o["user-device-store-max-devices"]) {
			return fmt.Errorf("Error reading user_device_store_max_devices: %v", err)
		}
	}

	if err = d.Set("user_device_store_max_users", flattenSystemGlobalUserDeviceStoreMaxUsers(o["user-device-store-max-users"], d, "user_device_store_max_users", sv)); err != nil {
		if !fortiAPIPatch(o["user-device-store-max-users"]) {
			return fmt.Errorf("Error reading user_device_store_max_users: %v", err)
		}
	}

	if err = d.Set("user_device_store_max_unified_mem", flattenSystemGlobalUserDeviceStoreMaxUnifiedMem(o["user-device-store-max-unified-mem"], d, "user_device_store_max_unified_mem", sv)); err != nil {
		if !fortiAPIPatch(o["user-device-store-max-unified-mem"]) {
			return fmt.Errorf("Error reading user_device_store_max_unified_mem: %v", err)
		}
	}

	if err = d.Set("device_identification_active_scan_delay", flattenSystemGlobalDeviceIdentificationActiveScanDelay(o["device-identification-active-scan-delay"], d, "device_identification_active_scan_delay", sv)); err != nil {
		if !fortiAPIPatch(o["device-identification-active-scan-delay"]) {
			return fmt.Errorf("Error reading device_identification_active_scan_delay: %v", err)
		}
	}

	if err = d.Set("compliance_check", flattenSystemGlobalComplianceCheck(o["compliance-check"], d, "compliance_check", sv)); err != nil {
		if !fortiAPIPatch(o["compliance-check"]) {
			return fmt.Errorf("Error reading compliance_check: %v", err)
		}
	}

	if err = d.Set("compliance_check_time", flattenSystemGlobalComplianceCheckTime(o["compliance-check-time"], d, "compliance_check_time", sv)); err != nil {
		if !fortiAPIPatch(o["compliance-check-time"]) {
			return fmt.Errorf("Error reading compliance_check_time: %v", err)
		}
	}

	if err = d.Set("gui_device_latitude", flattenSystemGlobalGuiDeviceLatitude(o["gui-device-latitude"], d, "gui_device_latitude", sv)); err != nil {
		if !fortiAPIPatch(o["gui-device-latitude"]) {
			return fmt.Errorf("Error reading gui_device_latitude: %v", err)
		}
	}

	if err = d.Set("gui_device_longitude", flattenSystemGlobalGuiDeviceLongitude(o["gui-device-longitude"], d, "gui_device_longitude", sv)); err != nil {
		if !fortiAPIPatch(o["gui-device-longitude"]) {
			return fmt.Errorf("Error reading gui_device_longitude: %v", err)
		}
	}

	if err = d.Set("private_data_encryption", flattenSystemGlobalPrivateDataEncryption(o["private-data-encryption"], d, "private_data_encryption", sv)); err != nil {
		if !fortiAPIPatch(o["private-data-encryption"]) {
			return fmt.Errorf("Error reading private_data_encryption: %v", err)
		}
	}

	if err = d.Set("auto_auth_extension_device", flattenSystemGlobalAutoAuthExtensionDevice(o["auto-auth-extension-device"], d, "auto_auth_extension_device", sv)); err != nil {
		if !fortiAPIPatch(o["auto-auth-extension-device"]) {
			return fmt.Errorf("Error reading auto_auth_extension_device: %v", err)
		}
	}

	if err = d.Set("gui_theme", flattenSystemGlobalGuiTheme(o["gui-theme"], d, "gui_theme", sv)); err != nil {
		if !fortiAPIPatch(o["gui-theme"]) {
			return fmt.Errorf("Error reading gui_theme: %v", err)
		}
	}

	if err = d.Set("gui_date_format", flattenSystemGlobalGuiDateFormat(o["gui-date-format"], d, "gui_date_format", sv)); err != nil {
		if !fortiAPIPatch(o["gui-date-format"]) {
			return fmt.Errorf("Error reading gui_date_format: %v", err)
		}
	}

	if err = d.Set("gui_date_time_source", flattenSystemGlobalGuiDateTimeSource(o["gui-date-time-source"], d, "gui_date_time_source", sv)); err != nil {
		if !fortiAPIPatch(o["gui-date-time-source"]) {
			return fmt.Errorf("Error reading gui_date_time_source: %v", err)
		}
	}

	if err = d.Set("igmp_state_limit", flattenSystemGlobalIgmpStateLimit(o["igmp-state-limit"], d, "igmp_state_limit", sv)); err != nil {
		if !fortiAPIPatch(o["igmp-state-limit"]) {
			return fmt.Errorf("Error reading igmp_state_limit: %v", err)
		}
	}

	if err = d.Set("cloud_communication", flattenSystemGlobalCloudCommunication(o["cloud-communication"], d, "cloud_communication", sv)); err != nil {
		if !fortiAPIPatch(o["cloud-communication"]) {
			return fmt.Errorf("Error reading cloud_communication: %v", err)
		}
	}

	if err = d.Set("fec_port", flattenSystemGlobalFecPort(o["fec-port"], d, "fec_port", sv)); err != nil {
		if !fortiAPIPatch(o["fec-port"]) {
			return fmt.Errorf("Error reading fec_port: %v", err)
		}
	}

	if err = d.Set("ipsec_ha_seqjump_rate", flattenSystemGlobalIpsecHaSeqjumpRate(o["ipsec-ha-seqjump-rate"], d, "ipsec_ha_seqjump_rate", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-ha-seqjump-rate"]) {
			return fmt.Errorf("Error reading ipsec_ha_seqjump_rate: %v", err)
		}
	}

	if err = d.Set("fortitoken_cloud", flattenSystemGlobalFortitokenCloud(o["fortitoken-cloud"], d, "fortitoken_cloud", sv)); err != nil {
		if !fortiAPIPatch(o["fortitoken-cloud"]) {
			return fmt.Errorf("Error reading fortitoken_cloud: %v", err)
		}
	}

	if err = d.Set("faz_disk_buffer_size", flattenSystemGlobalFazDiskBufferSize(o["faz-disk-buffer-size"], d, "faz_disk_buffer_size", sv)); err != nil {
		if !fortiAPIPatch(o["faz-disk-buffer-size"]) {
			return fmt.Errorf("Error reading faz_disk_buffer_size: %v", err)
		}
	}

	if err = d.Set("irq_time_accounting", flattenSystemGlobalIrqTimeAccounting(o["irq-time-accounting"], d, "irq_time_accounting", sv)); err != nil {
		if !fortiAPIPatch(o["irq-time-accounting"]) {
			return fmt.Errorf("Error reading irq_time_accounting: %v", err)
		}
	}

	if err = d.Set("fortiipam_integration", flattenSystemGlobalFortiipamIntegration(o["fortiipam-integration"], d, "fortiipam_integration", sv)); err != nil {
		if !fortiAPIPatch(o["fortiipam-integration"]) {
			return fmt.Errorf("Error reading fortiipam_integration: %v", err)
		}
	}

	if err = d.Set("management_ip", flattenSystemGlobalManagementIp(o["management-ip"], d, "management_ip", sv)); err != nil {
		if !fortiAPIPatch(o["management-ip"]) {
			return fmt.Errorf("Error reading management_ip: %v", err)
		}
	}

	if err = d.Set("management_port", flattenSystemGlobalManagementPort(o["management-port"], d, "management_port", sv)); err != nil {
		if !fortiAPIPatch(o["management-port"]) {
			return fmt.Errorf("Error reading management_port: %v", err)
		}
	}

	if err = d.Set("management_port_use_admin_sport", flattenSystemGlobalManagementPortUseAdminSport(o["management-port-use-admin-sport"], d, "management_port_use_admin_sport", sv)); err != nil {
		if !fortiAPIPatch(o["management-port-use-admin-sport"]) {
			return fmt.Errorf("Error reading management_port_use_admin_sport: %v", err)
		}
	}

	if err = d.Set("internet_service_database", flattenSystemGlobalInternetServiceDatabase(o["internet-service-database"], d, "internet_service_database", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service-database"]) {
			return fmt.Errorf("Error reading internet_service_database: %v", err)
		}
	}

	return nil
}

func flattenSystemGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemGlobalLanguage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiIpv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiReplacementMessageGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiLocalOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiCertificates(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiCustomLanguage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiWirelessOpensecurity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiDisplayHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiFortigateCloudSandbox(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiFortisandboxCloud(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiFirmwareUpgradeWarning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiFirmwareUpgradeSetupWarning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiLinesPerPage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminHttpsSslVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminHttpsSslCiphersuites(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminHttpsSslBannedCiphers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdmintimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminConsoleTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSsdTrimFreq(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSsdTrimHour(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSsdTrimMin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSsdTrimWeekday(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSsdTrimDate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminConcurrent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminLockoutThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminLockoutDuration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRefresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFailtime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDailyRestart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRestartTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRadiusPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminLoginMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRemoteauthtimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLdapconntimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalBatchCmdb(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMaxDlpstatMemory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMultiFactorAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAutorunLogFsck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTimezone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTrafficPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTrafficPriorityLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAntiReplay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSendPmtuIcmp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHonorDf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPmtuDiscovery(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalVirtualSwitchVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSplitPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRevisionImageAutoBackup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRevisionBackupOnLogout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalManagementVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiAllowDefaultHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiForticareRegistrationSetupWarning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiCdnUsage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAlias(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalStrongCrypto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSshCbcCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSshHmacMd5(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSshKexSha1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSshMacWeak(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslStaticKeyCiphers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSshKexAlgo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSshEncAlgo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSshMacAlgo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSnatRouteChange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSpeedtestServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCliAuditLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDhParams(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFdsStatistics(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFdsStatisticsPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMulticastForward(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMcTtlNotchange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAsymroute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTcpOption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLldpTransmission(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLldpReception(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyAuthTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyReAuthenticationMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyAuthLifetime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyAuthLifetimeTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyResourceMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyCertUseMgmtVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSysPerfLogInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCheckProtocolHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalVipArpRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalResetSessionlessTcp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAllowTrafficRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpv6AllowTrafficRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalStrictDirtySessionCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTcpHalfcloseTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTcpHalfopenTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTcpTimewaitTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTcpRstTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalUdpIdleTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalBlockSessionTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpSrcPortRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPreLoginBanner(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPostLoginBanner(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTftp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAvFailopen(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAvFailopenSession(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMemoryUseThresholdExtreme(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMemoryUseThresholdRed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMemoryUseThresholdGreen(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCpuUseThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCheckResetRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalVdomMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalVdomAdmin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLongVdomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalEditVdomPrompt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminSport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminHttpsRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminHstsMaxAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminSshPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminRestrictLocal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminSshPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminSshGraceTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminSshV1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminTelnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminTelnetPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminForticloudSsoLogin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDefaultServiceSourcePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminMaintainer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalUserServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminHttpsPkiRequired(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWifiCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWifiCaCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAuthHttpPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAuthHttpsPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAuthKeepalive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPolicyAuthConcurrent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAuthSessionLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAuthCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCltCertReq(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortiservicePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalEndpointControlPortalPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalEndpointControlFdsAccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTpMcSkipPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCfgSave(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCfgRevertTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalRebootUponConfigRestore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAdminScp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSecurityRatingResultSubmission(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSecurityRatingRunOnSchedule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWirelessController(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWirelessControllerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortiextenderDataPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortiextender(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalExtenderControllerReservedNetwork(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortiextenderDiscoveryLockdown(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortiextenderVlanMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSwitchController(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSwitchControllerReservedNetwork(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDnsproxyWorkerCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalUrlFilterCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyWorkerCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalScanunitCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyHardwareAcceleration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyKxpHardwareAcceleration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalProxyCipherHardwareAcceleration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFgdAlertSubscription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpsecHmacOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpv6AcceptDad(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpv6AllowAnycastProbe(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCsrCaAttribute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWimax4GUsb(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCertChainMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslvpnMaxWorkerCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslvpnEmsSnCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslvpnKxpHardwareAcceleration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslvpnCipherHardwareAcceleration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSslvpnPluginVersionCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTwoFactorFtkExpiry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTwoFactorEmailExpiry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTwoFactorSmsExpiry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTwoFactorFacExpiry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalTwoFactorFtmExpiry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPerUserBal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPerUserBwl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalVirtualServerCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalVirtualServerHardwareAcceleration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWadWorkerCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWadCsvcCsCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWadCsvcDbCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWadSourceAffinity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWadMemoryChangeGranularity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLoginTimestamp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMiglogdChildren(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalSpecialFile23Support(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLogUuidPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLogUuidAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalLogSslConnection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiRestApiCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalArpMaxEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalHaAffinity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCmdbsvrAffinity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAvAffinity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalWadAffinity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpsAffinity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMiglogAffinity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalUrlFilterAffinity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalNdpMaxEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalBrFdbMaxEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalMaxRouteCacheSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpsecAsicOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpsecSoftDecAsync(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIkeEmbryonicLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDeviceIdleTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalUserDeviceStoreMaxDevices(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalUserDeviceStoreMaxUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalUserDeviceStoreMaxUnifiedMem(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalDeviceIdentificationActiveScanDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalComplianceCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalComplianceCheckTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiDeviceLatitude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiDeviceLongitude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalPrivateDataEncryption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalAutoAuthExtensionDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiTheme(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiDateFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalGuiDateTimeSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIgmpStateLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalCloudCommunication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFecPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIpsecHaSeqjumpRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortitokenCloud(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFazDiskBufferSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalIrqTimeAccounting(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalFortiipamIntegration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalManagementIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalManagementPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalManagementPortUseAdminSport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGlobalInternetServiceDatabase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemGlobal(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("language"); ok {

		t, err := expandSystemGlobalLanguage(d, v, "language", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["language"] = t
		}
	}

	if v, ok := d.GetOk("gui_ipv6"); ok {

		t, err := expandSystemGlobalGuiIpv6(d, v, "gui_ipv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-ipv6"] = t
		}
	}

	if v, ok := d.GetOk("gui_replacement_message_groups"); ok {

		t, err := expandSystemGlobalGuiReplacementMessageGroups(d, v, "gui_replacement_message_groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-replacement-message-groups"] = t
		}
	}

	if v, ok := d.GetOk("gui_local_out"); ok {

		t, err := expandSystemGlobalGuiLocalOut(d, v, "gui_local_out", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-local-out"] = t
		}
	}

	if v, ok := d.GetOk("gui_certificates"); ok {

		t, err := expandSystemGlobalGuiCertificates(d, v, "gui_certificates", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-certificates"] = t
		}
	}

	if v, ok := d.GetOk("gui_custom_language"); ok {

		t, err := expandSystemGlobalGuiCustomLanguage(d, v, "gui_custom_language", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-custom-language"] = t
		}
	}

	if v, ok := d.GetOk("gui_wireless_opensecurity"); ok {

		t, err := expandSystemGlobalGuiWirelessOpensecurity(d, v, "gui_wireless_opensecurity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-wireless-opensecurity"] = t
		}
	}

	if v, ok := d.GetOk("gui_display_hostname"); ok {

		t, err := expandSystemGlobalGuiDisplayHostname(d, v, "gui_display_hostname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-display-hostname"] = t
		}
	}

	if v, ok := d.GetOk("gui_fortigate_cloud_sandbox"); ok {

		t, err := expandSystemGlobalGuiFortigateCloudSandbox(d, v, "gui_fortigate_cloud_sandbox", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-fortigate-cloud-sandbox"] = t
		}
	}

	if v, ok := d.GetOk("gui_fortisandbox_cloud"); ok {

		t, err := expandSystemGlobalGuiFortisandboxCloud(d, v, "gui_fortisandbox_cloud", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-fortisandbox-cloud"] = t
		}
	}

	if v, ok := d.GetOk("gui_firmware_upgrade_warning"); ok {

		t, err := expandSystemGlobalGuiFirmwareUpgradeWarning(d, v, "gui_firmware_upgrade_warning", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-firmware-upgrade-warning"] = t
		}
	}

	if v, ok := d.GetOk("gui_firmware_upgrade_setup_warning"); ok {

		t, err := expandSystemGlobalGuiFirmwareUpgradeSetupWarning(d, v, "gui_firmware_upgrade_setup_warning", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-firmware-upgrade-setup-warning"] = t
		}
	}

	if v, ok := d.GetOk("gui_lines_per_page"); ok {

		t, err := expandSystemGlobalGuiLinesPerPage(d, v, "gui_lines_per_page", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-lines-per-page"] = t
		}
	}

	if v, ok := d.GetOk("admin_https_ssl_versions"); ok {

		t, err := expandSystemGlobalAdminHttpsSslVersions(d, v, "admin_https_ssl_versions", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-https-ssl-versions"] = t
		}
	}

	if v, ok := d.GetOk("admin_https_ssl_ciphersuites"); ok {

		t, err := expandSystemGlobalAdminHttpsSslCiphersuites(d, v, "admin_https_ssl_ciphersuites", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-https-ssl-ciphersuites"] = t
		}
	}

	if v, ok := d.GetOk("admin_https_ssl_banned_ciphers"); ok {

		t, err := expandSystemGlobalAdminHttpsSslBannedCiphers(d, v, "admin_https_ssl_banned_ciphers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-https-ssl-banned-ciphers"] = t
		}
	}

	if v, ok := d.GetOk("admintimeout"); ok {

		t, err := expandSystemGlobalAdmintimeout(d, v, "admintimeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admintimeout"] = t
		}
	}

	if v, ok := d.GetOk("admin_console_timeout"); ok {

		t, err := expandSystemGlobalAdminConsoleTimeout(d, v, "admin_console_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-console-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssd_trim_freq"); ok {

		t, err := expandSystemGlobalSsdTrimFreq(d, v, "ssd_trim_freq", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssd-trim-freq"] = t
		}
	}

	if v, ok := d.GetOkExists("ssd_trim_hour"); ok {

		t, err := expandSystemGlobalSsdTrimHour(d, v, "ssd_trim_hour", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssd-trim-hour"] = t
		}
	}

	if v, ok := d.GetOkExists("ssd_trim_min"); ok {

		t, err := expandSystemGlobalSsdTrimMin(d, v, "ssd_trim_min", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssd-trim-min"] = t
		}
	}

	if v, ok := d.GetOk("ssd_trim_weekday"); ok {

		t, err := expandSystemGlobalSsdTrimWeekday(d, v, "ssd_trim_weekday", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssd-trim-weekday"] = t
		}
	}

	if v, ok := d.GetOk("ssd_trim_date"); ok {

		t, err := expandSystemGlobalSsdTrimDate(d, v, "ssd_trim_date", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssd-trim-date"] = t
		}
	}

	if v, ok := d.GetOk("admin_concurrent"); ok {

		t, err := expandSystemGlobalAdminConcurrent(d, v, "admin_concurrent", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-concurrent"] = t
		}
	}

	if v, ok := d.GetOk("admin_lockout_threshold"); ok {

		t, err := expandSystemGlobalAdminLockoutThreshold(d, v, "admin_lockout_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-lockout-threshold"] = t
		}
	}

	if v, ok := d.GetOk("admin_lockout_duration"); ok {

		t, err := expandSystemGlobalAdminLockoutDuration(d, v, "admin_lockout_duration", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-lockout-duration"] = t
		}
	}

	if v, ok := d.GetOkExists("refresh"); ok {

		t, err := expandSystemGlobalRefresh(d, v, "refresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["refresh"] = t
		}
	}

	if v, ok := d.GetOkExists("interval"); ok {

		t, err := expandSystemGlobalInterval(d, v, "interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOkExists("failtime"); ok {

		t, err := expandSystemGlobalFailtime(d, v, "failtime", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failtime"] = t
		}
	}

	if v, ok := d.GetOk("daily_restart"); ok {

		t, err := expandSystemGlobalDailyRestart(d, v, "daily_restart", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["daily-restart"] = t
		}
	}

	if v, ok := d.GetOk("restart_time"); ok {

		t, err := expandSystemGlobalRestartTime(d, v, "restart_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restart-time"] = t
		}
	}

	if v, ok := d.GetOk("radius_port"); ok {

		t, err := expandSystemGlobalRadiusPort(d, v, "radius_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-port"] = t
		}
	}

	if v, ok := d.GetOk("admin_login_max"); ok {

		t, err := expandSystemGlobalAdminLoginMax(d, v, "admin_login_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-login-max"] = t
		}
	}

	if v, ok := d.GetOkExists("remoteauthtimeout"); ok {

		t, err := expandSystemGlobalRemoteauthtimeout(d, v, "remoteauthtimeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remoteauthtimeout"] = t
		}
	}

	if v, ok := d.GetOk("ldapconntimeout"); ok {

		t, err := expandSystemGlobalLdapconntimeout(d, v, "ldapconntimeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldapconntimeout"] = t
		}
	}

	if v, ok := d.GetOk("batch_cmdb"); ok {

		t, err := expandSystemGlobalBatchCmdb(d, v, "batch_cmdb", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["batch-cmdb"] = t
		}
	}

	if v, ok := d.GetOk("max_dlpstat_memory"); ok {

		t, err := expandSystemGlobalMaxDlpstatMemory(d, v, "max_dlpstat_memory", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-dlpstat-memory"] = t
		}
	}

	if v, ok := d.GetOk("multi_factor_authentication"); ok {

		t, err := expandSystemGlobalMultiFactorAuthentication(d, v, "multi_factor_authentication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multi-factor-authentication"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok {

		t, err := expandSystemGlobalSslMinProtoVersion(d, v, "ssl_min_proto_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("autorun_log_fsck"); ok {

		t, err := expandSystemGlobalAutorunLogFsck(d, v, "autorun_log_fsck", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["autorun-log-fsck"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok {

		t, err := expandSystemGlobalDst(d, v, "dst", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("timezone"); ok {

		t, err := expandSystemGlobalTimezone(d, v, "timezone", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timezone"] = t
		}
	}

	if v, ok := d.GetOk("traffic_priority"); ok {

		t, err := expandSystemGlobalTrafficPriority(d, v, "traffic_priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-priority"] = t
		}
	}

	if v, ok := d.GetOk("traffic_priority_level"); ok {

		t, err := expandSystemGlobalTrafficPriorityLevel(d, v, "traffic_priority_level", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-priority-level"] = t
		}
	}

	if v, ok := d.GetOk("anti_replay"); ok {

		t, err := expandSystemGlobalAntiReplay(d, v, "anti_replay", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anti-replay"] = t
		}
	}

	if v, ok := d.GetOk("send_pmtu_icmp"); ok {

		t, err := expandSystemGlobalSendPmtuIcmp(d, v, "send_pmtu_icmp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-pmtu-icmp"] = t
		}
	}

	if v, ok := d.GetOk("honor_df"); ok {

		t, err := expandSystemGlobalHonorDf(d, v, "honor_df", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["honor-df"] = t
		}
	}

	if v, ok := d.GetOk("pmtu_discovery"); ok {

		t, err := expandSystemGlobalPmtuDiscovery(d, v, "pmtu_discovery", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmtu-discovery"] = t
		}
	}

	if v, ok := d.GetOk("virtual_switch_vlan"); ok {

		t, err := expandSystemGlobalVirtualSwitchVlan(d, v, "virtual_switch_vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-switch-vlan"] = t
		}
	}

	if v, ok := d.GetOk("split_port"); ok {

		t, err := expandSystemGlobalSplitPort(d, v, "split_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-port"] = t
		}
	}

	if v, ok := d.GetOk("revision_image_auto_backup"); ok {

		t, err := expandSystemGlobalRevisionImageAutoBackup(d, v, "revision_image_auto_backup", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["revision-image-auto-backup"] = t
		}
	}

	if v, ok := d.GetOk("revision_backup_on_logout"); ok {

		t, err := expandSystemGlobalRevisionBackupOnLogout(d, v, "revision_backup_on_logout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["revision-backup-on-logout"] = t
		}
	}

	if v, ok := d.GetOk("management_vdom"); ok {

		t, err := expandSystemGlobalManagementVdom(d, v, "management_vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["management-vdom"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok {

		t, err := expandSystemGlobalHostname(d, v, "hostname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("gui_allow_default_hostname"); ok {

		t, err := expandSystemGlobalGuiAllowDefaultHostname(d, v, "gui_allow_default_hostname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-allow-default-hostname"] = t
		}
	}

	if v, ok := d.GetOk("gui_forticare_registration_setup_warning"); ok {

		t, err := expandSystemGlobalGuiForticareRegistrationSetupWarning(d, v, "gui_forticare_registration_setup_warning", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-forticare-registration-setup-warning"] = t
		}
	}

	if v, ok := d.GetOk("gui_cdn_usage"); ok {

		t, err := expandSystemGlobalGuiCdnUsage(d, v, "gui_cdn_usage", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-cdn-usage"] = t
		}
	}

	if v, ok := d.GetOk("alias"); ok {

		t, err := expandSystemGlobalAlias(d, v, "alias", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alias"] = t
		}
	}

	if v, ok := d.GetOk("strong_crypto"); ok {

		t, err := expandSystemGlobalStrongCrypto(d, v, "strong_crypto", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strong-crypto"] = t
		}
	}

	if v, ok := d.GetOk("ssh_cbc_cipher"); ok {

		t, err := expandSystemGlobalSshCbcCipher(d, v, "ssh_cbc_cipher", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-cbc-cipher"] = t
		}
	}

	if v, ok := d.GetOk("ssh_hmac_md5"); ok {

		t, err := expandSystemGlobalSshHmacMd5(d, v, "ssh_hmac_md5", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-hmac-md5"] = t
		}
	}

	if v, ok := d.GetOk("ssh_kex_sha1"); ok {

		t, err := expandSystemGlobalSshKexSha1(d, v, "ssh_kex_sha1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-kex-sha1"] = t
		}
	}

	if v, ok := d.GetOk("ssh_mac_weak"); ok {

		t, err := expandSystemGlobalSshMacWeak(d, v, "ssh_mac_weak", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-mac-weak"] = t
		}
	}

	if v, ok := d.GetOk("ssl_static_key_ciphers"); ok {

		t, err := expandSystemGlobalSslStaticKeyCiphers(d, v, "ssl_static_key_ciphers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-static-key-ciphers"] = t
		}
	}

	if v, ok := d.GetOk("ssh_kex_algo"); ok {

		t, err := expandSystemGlobalSshKexAlgo(d, v, "ssh_kex_algo", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-kex-algo"] = t
		}
	}

	if v, ok := d.GetOk("ssh_enc_algo"); ok {

		t, err := expandSystemGlobalSshEncAlgo(d, v, "ssh_enc_algo", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-enc-algo"] = t
		}
	}

	if v, ok := d.GetOk("ssh_mac_algo"); ok {

		t, err := expandSystemGlobalSshMacAlgo(d, v, "ssh_mac_algo", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-mac-algo"] = t
		}
	}

	if v, ok := d.GetOk("snat_route_change"); ok {

		t, err := expandSystemGlobalSnatRouteChange(d, v, "snat_route_change", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["snat-route-change"] = t
		}
	}

	if v, ok := d.GetOk("speedtest_server"); ok {

		t, err := expandSystemGlobalSpeedtestServer(d, v, "speedtest_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["speedtest-server"] = t
		}
	}

	if v, ok := d.GetOk("cli_audit_log"); ok {

		t, err := expandSystemGlobalCliAuditLog(d, v, "cli_audit_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cli-audit-log"] = t
		}
	}

	if v, ok := d.GetOk("dh_params"); ok {

		t, err := expandSystemGlobalDhParams(d, v, "dh_params", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-params"] = t
		}
	}

	if v, ok := d.GetOk("fds_statistics"); ok {

		t, err := expandSystemGlobalFdsStatistics(d, v, "fds_statistics", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fds-statistics"] = t
		}
	}

	if v, ok := d.GetOk("fds_statistics_period"); ok {

		t, err := expandSystemGlobalFdsStatisticsPeriod(d, v, "fds_statistics_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fds-statistics-period"] = t
		}
	}

	if v, ok := d.GetOk("multicast_forward"); ok {

		t, err := expandSystemGlobalMulticastForward(d, v, "multicast_forward", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-forward"] = t
		}
	}

	if v, ok := d.GetOk("mc_ttl_notchange"); ok {

		t, err := expandSystemGlobalMcTtlNotchange(d, v, "mc_ttl_notchange", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mc-ttl-notchange"] = t
		}
	}

	if v, ok := d.GetOk("asymroute"); ok {

		t, err := expandSystemGlobalAsymroute(d, v, "asymroute", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["asymroute"] = t
		}
	}

	if v, ok := d.GetOk("tcp_option"); ok {

		t, err := expandSystemGlobalTcpOption(d, v, "tcp_option", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-option"] = t
		}
	}

	if v, ok := d.GetOk("lldp_transmission"); ok {

		t, err := expandSystemGlobalLldpTransmission(d, v, "lldp_transmission", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-transmission"] = t
		}
	}

	if v, ok := d.GetOk("lldp_reception"); ok {

		t, err := expandSystemGlobalLldpReception(d, v, "lldp_reception", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-reception"] = t
		}
	}

	if v, ok := d.GetOk("proxy_auth_timeout"); ok {

		t, err := expandSystemGlobalProxyAuthTimeout(d, v, "proxy_auth_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-auth-timeout"] = t
		}
	}

	if v, ok := d.GetOk("proxy_re_authentication_mode"); ok {

		t, err := expandSystemGlobalProxyReAuthenticationMode(d, v, "proxy_re_authentication_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-re-authentication-mode"] = t
		}
	}

	if v, ok := d.GetOk("proxy_auth_lifetime"); ok {

		t, err := expandSystemGlobalProxyAuthLifetime(d, v, "proxy_auth_lifetime", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-auth-lifetime"] = t
		}
	}

	if v, ok := d.GetOk("proxy_auth_lifetime_timeout"); ok {

		t, err := expandSystemGlobalProxyAuthLifetimeTimeout(d, v, "proxy_auth_lifetime_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-auth-lifetime-timeout"] = t
		}
	}

	if v, ok := d.GetOk("proxy_resource_mode"); ok {

		t, err := expandSystemGlobalProxyResourceMode(d, v, "proxy_resource_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-resource-mode"] = t
		}
	}

	if v, ok := d.GetOk("proxy_cert_use_mgmt_vdom"); ok {

		t, err := expandSystemGlobalProxyCertUseMgmtVdom(d, v, "proxy_cert_use_mgmt_vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-cert-use-mgmt-vdom"] = t
		}
	}

	if v, ok := d.GetOkExists("sys_perf_log_interval"); ok {

		t, err := expandSystemGlobalSysPerfLogInterval(d, v, "sys_perf_log_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sys-perf-log-interval"] = t
		}
	}

	if v, ok := d.GetOk("check_protocol_header"); ok {

		t, err := expandSystemGlobalCheckProtocolHeader(d, v, "check_protocol_header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-protocol-header"] = t
		}
	}

	if v, ok := d.GetOk("vip_arp_range"); ok {

		t, err := expandSystemGlobalVipArpRange(d, v, "vip_arp_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip-arp-range"] = t
		}
	}

	if v, ok := d.GetOk("reset_sessionless_tcp"); ok {

		t, err := expandSystemGlobalResetSessionlessTcp(d, v, "reset_sessionless_tcp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reset-sessionless-tcp"] = t
		}
	}

	if v, ok := d.GetOk("allow_traffic_redirect"); ok {

		t, err := expandSystemGlobalAllowTrafficRedirect(d, v, "allow_traffic_redirect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-traffic-redirect"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_allow_traffic_redirect"); ok {

		t, err := expandSystemGlobalIpv6AllowTrafficRedirect(d, v, "ipv6_allow_traffic_redirect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-allow-traffic-redirect"] = t
		}
	}

	if v, ok := d.GetOk("strict_dirty_session_check"); ok {

		t, err := expandSystemGlobalStrictDirtySessionCheck(d, v, "strict_dirty_session_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strict-dirty-session-check"] = t
		}
	}

	if v, ok := d.GetOk("tcp_halfclose_timer"); ok {

		t, err := expandSystemGlobalTcpHalfcloseTimer(d, v, "tcp_halfclose_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-halfclose-timer"] = t
		}
	}

	if v, ok := d.GetOk("tcp_halfopen_timer"); ok {

		t, err := expandSystemGlobalTcpHalfopenTimer(d, v, "tcp_halfopen_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-halfopen-timer"] = t
		}
	}

	if v, ok := d.GetOkExists("tcp_timewait_timer"); ok {

		t, err := expandSystemGlobalTcpTimewaitTimer(d, v, "tcp_timewait_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-timewait-timer"] = t
		}
	}

	if v, ok := d.GetOk("tcp_rst_timer"); ok {

		t, err := expandSystemGlobalTcpRstTimer(d, v, "tcp_rst_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-rst-timer"] = t
		}
	}

	if v, ok := d.GetOk("udp_idle_timer"); ok {

		t, err := expandSystemGlobalUdpIdleTimer(d, v, "udp_idle_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-idle-timer"] = t
		}
	}

	if v, ok := d.GetOk("block_session_timer"); ok {

		t, err := expandSystemGlobalBlockSessionTimer(d, v, "block_session_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-session-timer"] = t
		}
	}

	if v, ok := d.GetOk("ip_src_port_range"); ok {

		t, err := expandSystemGlobalIpSrcPortRange(d, v, "ip_src_port_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-src-port-range"] = t
		}
	}

	if v, ok := d.GetOk("pre_login_banner"); ok {

		t, err := expandSystemGlobalPreLoginBanner(d, v, "pre_login_banner", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pre-login-banner"] = t
		}
	}

	if v, ok := d.GetOk("post_login_banner"); ok {

		t, err := expandSystemGlobalPostLoginBanner(d, v, "post_login_banner", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["post-login-banner"] = t
		}
	}

	if v, ok := d.GetOk("tftp"); ok {

		t, err := expandSystemGlobalTftp(d, v, "tftp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tftp"] = t
		}
	}

	if v, ok := d.GetOk("av_failopen"); ok {

		t, err := expandSystemGlobalAvFailopen(d, v, "av_failopen", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-failopen"] = t
		}
	}

	if v, ok := d.GetOk("av_failopen_session"); ok {

		t, err := expandSystemGlobalAvFailopenSession(d, v, "av_failopen_session", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-failopen-session"] = t
		}
	}

	if v, ok := d.GetOk("memory_use_threshold_extreme"); ok {

		t, err := expandSystemGlobalMemoryUseThresholdExtreme(d, v, "memory_use_threshold_extreme", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["memory-use-threshold-extreme"] = t
		}
	}

	if v, ok := d.GetOk("memory_use_threshold_red"); ok {

		t, err := expandSystemGlobalMemoryUseThresholdRed(d, v, "memory_use_threshold_red", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["memory-use-threshold-red"] = t
		}
	}

	if v, ok := d.GetOk("memory_use_threshold_green"); ok {

		t, err := expandSystemGlobalMemoryUseThresholdGreen(d, v, "memory_use_threshold_green", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["memory-use-threshold-green"] = t
		}
	}

	if v, ok := d.GetOk("cpu_use_threshold"); ok {

		t, err := expandSystemGlobalCpuUseThreshold(d, v, "cpu_use_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cpu-use-threshold"] = t
		}
	}

	if v, ok := d.GetOk("check_reset_range"); ok {

		t, err := expandSystemGlobalCheckResetRange(d, v, "check_reset_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-reset-range"] = t
		}
	}

	if v, ok := d.GetOk("vdom_mode"); ok {

		t, err := expandSystemGlobalVdomMode(d, v, "vdom_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom-mode"] = t
		}
	}

	if v, ok := d.GetOk("vdom_admin"); ok {

		t, err := expandSystemGlobalVdomAdmin(d, v, "vdom_admin", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom-admin"] = t
		}
	}

	if v, ok := d.GetOk("long_vdom_name"); ok {

		t, err := expandSystemGlobalLongVdomName(d, v, "long_vdom_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["long-vdom-name"] = t
		}
	}

	if v, ok := d.GetOk("edit_vdom_prompt"); ok {

		t, err := expandSystemGlobalEditVdomPrompt(d, v, "edit_vdom_prompt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["edit-vdom-prompt"] = t
		}
	}

	if v, ok := d.GetOk("admin_port"); ok {

		t, err := expandSystemGlobalAdminPort(d, v, "admin_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-port"] = t
		}
	}

	if v, ok := d.GetOk("admin_sport"); ok {

		t, err := expandSystemGlobalAdminSport(d, v, "admin_sport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-sport"] = t
		}
	}

	if v, ok := d.GetOk("admin_https_redirect"); ok {

		t, err := expandSystemGlobalAdminHttpsRedirect(d, v, "admin_https_redirect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-https-redirect"] = t
		}
	}

	if v, ok := d.GetOkExists("admin_hsts_max_age"); ok {

		t, err := expandSystemGlobalAdminHstsMaxAge(d, v, "admin_hsts_max_age", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-hsts-max-age"] = t
		}
	}

	if v, ok := d.GetOk("admin_ssh_password"); ok {

		t, err := expandSystemGlobalAdminSshPassword(d, v, "admin_ssh_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-ssh-password"] = t
		}
	}

	if v, ok := d.GetOk("admin_restrict_local"); ok {

		t, err := expandSystemGlobalAdminRestrictLocal(d, v, "admin_restrict_local", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-restrict-local"] = t
		}
	}

	if v, ok := d.GetOk("admin_ssh_port"); ok {

		t, err := expandSystemGlobalAdminSshPort(d, v, "admin_ssh_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-ssh-port"] = t
		}
	}

	if v, ok := d.GetOk("admin_ssh_grace_time"); ok {

		t, err := expandSystemGlobalAdminSshGraceTime(d, v, "admin_ssh_grace_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-ssh-grace-time"] = t
		}
	}

	if v, ok := d.GetOk("admin_ssh_v1"); ok {

		t, err := expandSystemGlobalAdminSshV1(d, v, "admin_ssh_v1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-ssh-v1"] = t
		}
	}

	if v, ok := d.GetOk("admin_telnet"); ok {

		t, err := expandSystemGlobalAdminTelnet(d, v, "admin_telnet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-telnet"] = t
		}
	}

	if v, ok := d.GetOk("admin_telnet_port"); ok {

		t, err := expandSystemGlobalAdminTelnetPort(d, v, "admin_telnet_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-telnet-port"] = t
		}
	}

	if v, ok := d.GetOk("admin_forticloud_sso_login"); ok {

		t, err := expandSystemGlobalAdminForticloudSsoLogin(d, v, "admin_forticloud_sso_login", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-forticloud-sso-login"] = t
		}
	}

	if v, ok := d.GetOk("default_service_source_port"); ok {

		t, err := expandSystemGlobalDefaultServiceSourcePort(d, v, "default_service_source_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-service-source-port"] = t
		}
	}

	if v, ok := d.GetOk("admin_maintainer"); ok {

		t, err := expandSystemGlobalAdminMaintainer(d, v, "admin_maintainer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-maintainer"] = t
		}
	}

	if v, ok := d.GetOk("admin_server_cert"); ok {

		t, err := expandSystemGlobalAdminServerCert(d, v, "admin_server_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-server-cert"] = t
		}
	}

	if v, ok := d.GetOk("user_server_cert"); ok {

		t, err := expandSystemGlobalUserServerCert(d, v, "user_server_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-server-cert"] = t
		}
	}

	if v, ok := d.GetOk("admin_https_pki_required"); ok {

		t, err := expandSystemGlobalAdminHttpsPkiRequired(d, v, "admin_https_pki_required", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-https-pki-required"] = t
		}
	}

	if v, ok := d.GetOk("wifi_certificate"); ok {

		t, err := expandSystemGlobalWifiCertificate(d, v, "wifi_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-certificate"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ca_certificate"); ok {

		t, err := expandSystemGlobalWifiCaCertificate(d, v, "wifi_ca_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ca-certificate"] = t
		}
	}

	if v, ok := d.GetOk("auth_http_port"); ok {

		t, err := expandSystemGlobalAuthHttpPort(d, v, "auth_http_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-http-port"] = t
		}
	}

	if v, ok := d.GetOk("auth_https_port"); ok {

		t, err := expandSystemGlobalAuthHttpsPort(d, v, "auth_https_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-https-port"] = t
		}
	}

	if v, ok := d.GetOk("auth_keepalive"); ok {

		t, err := expandSystemGlobalAuthKeepalive(d, v, "auth_keepalive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-keepalive"] = t
		}
	}

	if v, ok := d.GetOkExists("policy_auth_concurrent"); ok {

		t, err := expandSystemGlobalPolicyAuthConcurrent(d, v, "policy_auth_concurrent", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-auth-concurrent"] = t
		}
	}

	if v, ok := d.GetOk("auth_session_limit"); ok {

		t, err := expandSystemGlobalAuthSessionLimit(d, v, "auth_session_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-session-limit"] = t
		}
	}

	if v, ok := d.GetOk("auth_cert"); ok {

		t, err := expandSystemGlobalAuthCert(d, v, "auth_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-cert"] = t
		}
	}

	if v, ok := d.GetOk("clt_cert_req"); ok {

		t, err := expandSystemGlobalCltCertReq(d, v, "clt_cert_req", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clt-cert-req"] = t
		}
	}

	if v, ok := d.GetOk("fortiservice_port"); ok {

		t, err := expandSystemGlobalFortiservicePort(d, v, "fortiservice_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiservice-port"] = t
		}
	}

	if v, ok := d.GetOk("endpoint_control_portal_port"); ok {

		t, err := expandSystemGlobalEndpointControlPortalPort(d, v, "endpoint_control_portal_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endpoint-control-portal-port"] = t
		}
	}

	if v, ok := d.GetOk("endpoint_control_fds_access"); ok {

		t, err := expandSystemGlobalEndpointControlFdsAccess(d, v, "endpoint_control_fds_access", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endpoint-control-fds-access"] = t
		}
	}

	if v, ok := d.GetOk("tp_mc_skip_policy"); ok {

		t, err := expandSystemGlobalTpMcSkipPolicy(d, v, "tp_mc_skip_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tp-mc-skip-policy"] = t
		}
	}

	if v, ok := d.GetOk("cfg_save"); ok {

		t, err := expandSystemGlobalCfgSave(d, v, "cfg_save", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cfg-save"] = t
		}
	}

	if v, ok := d.GetOk("cfg_revert_timeout"); ok {

		t, err := expandSystemGlobalCfgRevertTimeout(d, v, "cfg_revert_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cfg-revert-timeout"] = t
		}
	}

	if v, ok := d.GetOk("reboot_upon_config_restore"); ok {

		t, err := expandSystemGlobalRebootUponConfigRestore(d, v, "reboot_upon_config_restore", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reboot-upon-config-restore"] = t
		}
	}

	if v, ok := d.GetOk("admin_scp"); ok {

		t, err := expandSystemGlobalAdminScp(d, v, "admin_scp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-scp"] = t
		}
	}

	if v, ok := d.GetOk("security_rating_result_submission"); ok {

		t, err := expandSystemGlobalSecurityRatingResultSubmission(d, v, "security_rating_result_submission", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-rating-result-submission"] = t
		}
	}

	if v, ok := d.GetOk("security_rating_run_on_schedule"); ok {

		t, err := expandSystemGlobalSecurityRatingRunOnSchedule(d, v, "security_rating_run_on_schedule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-rating-run-on-schedule"] = t
		}
	}

	if v, ok := d.GetOk("wireless_controller"); ok {

		t, err := expandSystemGlobalWirelessController(d, v, "wireless_controller", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wireless-controller"] = t
		}
	}

	if v, ok := d.GetOk("wireless_controller_port"); ok {

		t, err := expandSystemGlobalWirelessControllerPort(d, v, "wireless_controller_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wireless-controller-port"] = t
		}
	}

	if v, ok := d.GetOk("fortiextender_data_port"); ok {

		t, err := expandSystemGlobalFortiextenderDataPort(d, v, "fortiextender_data_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiextender-data-port"] = t
		}
	}

	if v, ok := d.GetOk("fortiextender"); ok {

		t, err := expandSystemGlobalFortiextender(d, v, "fortiextender", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiextender"] = t
		}
	}

	if v, ok := d.GetOk("extender_controller_reserved_network"); ok {

		t, err := expandSystemGlobalExtenderControllerReservedNetwork(d, v, "extender_controller_reserved_network", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extender-controller-reserved-network"] = t
		}
	}

	if v, ok := d.GetOk("fortiextender_discovery_lockdown"); ok {

		t, err := expandSystemGlobalFortiextenderDiscoveryLockdown(d, v, "fortiextender_discovery_lockdown", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiextender-discovery-lockdown"] = t
		}
	}

	if v, ok := d.GetOk("fortiextender_vlan_mode"); ok {

		t, err := expandSystemGlobalFortiextenderVlanMode(d, v, "fortiextender_vlan_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiextender-vlan-mode"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller"); ok {

		t, err := expandSystemGlobalSwitchController(d, v, "switch_controller", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_reserved_network"); ok {

		t, err := expandSystemGlobalSwitchControllerReservedNetwork(d, v, "switch_controller_reserved_network", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-reserved-network"] = t
		}
	}

	if v, ok := d.GetOk("dnsproxy_worker_count"); ok {

		t, err := expandSystemGlobalDnsproxyWorkerCount(d, v, "dnsproxy_worker_count", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnsproxy-worker-count"] = t
		}
	}

	if v, ok := d.GetOk("url_filter_count"); ok {

		t, err := expandSystemGlobalUrlFilterCount(d, v, "url_filter_count", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-filter-count"] = t
		}
	}

	if v, ok := d.GetOkExists("proxy_worker_count"); ok {

		t, err := expandSystemGlobalProxyWorkerCount(d, v, "proxy_worker_count", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-worker-count"] = t
		}
	}

	if v, ok := d.GetOkExists("scanunit_count"); ok {

		t, err := expandSystemGlobalScanunitCount(d, v, "scanunit_count", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scanunit-count"] = t
		}
	}

	if v, ok := d.GetOk("proxy_hardware_acceleration"); ok {

		t, err := expandSystemGlobalProxyHardwareAcceleration(d, v, "proxy_hardware_acceleration", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-hardware-acceleration"] = t
		}
	}

	if v, ok := d.GetOk("proxy_kxp_hardware_acceleration"); ok {

		t, err := expandSystemGlobalProxyKxpHardwareAcceleration(d, v, "proxy_kxp_hardware_acceleration", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-kxp-hardware-acceleration"] = t
		}
	}

	if v, ok := d.GetOk("proxy_cipher_hardware_acceleration"); ok {

		t, err := expandSystemGlobalProxyCipherHardwareAcceleration(d, v, "proxy_cipher_hardware_acceleration", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-cipher-hardware-acceleration"] = t
		}
	}

	if v, ok := d.GetOk("fgd_alert_subscription"); ok {

		t, err := expandSystemGlobalFgdAlertSubscription(d, v, "fgd_alert_subscription", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgd-alert-subscription"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_hmac_offload"); ok {

		t, err := expandSystemGlobalIpsecHmacOffload(d, v, "ipsec_hmac_offload", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-hmac-offload"] = t
		}
	}

	if v, ok := d.GetOkExists("ipv6_accept_dad"); ok {

		t, err := expandSystemGlobalIpv6AcceptDad(d, v, "ipv6_accept_dad", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-accept-dad"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_allow_anycast_probe"); ok {

		t, err := expandSystemGlobalIpv6AllowAnycastProbe(d, v, "ipv6_allow_anycast_probe", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-allow-anycast-probe"] = t
		}
	}

	if v, ok := d.GetOk("csr_ca_attribute"); ok {

		t, err := expandSystemGlobalCsrCaAttribute(d, v, "csr_ca_attribute", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["csr-ca-attribute"] = t
		}
	}

	if v, ok := d.GetOk("wimax_4g_usb"); ok {

		t, err := expandSystemGlobalWimax4GUsb(d, v, "wimax_4g_usb", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wimax-4g-usb"] = t
		}
	}

	if v, ok := d.GetOk("cert_chain_max"); ok {

		t, err := expandSystemGlobalCertChainMax(d, v, "cert_chain_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-chain-max"] = t
		}
	}

	if v, ok := d.GetOkExists("sslvpn_max_worker_count"); ok {

		t, err := expandSystemGlobalSslvpnMaxWorkerCount(d, v, "sslvpn_max_worker_count", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-max-worker-count"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_ems_sn_check"); ok {

		t, err := expandSystemGlobalSslvpnEmsSnCheck(d, v, "sslvpn_ems_sn_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-ems-sn-check"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_kxp_hardware_acceleration"); ok {

		t, err := expandSystemGlobalSslvpnKxpHardwareAcceleration(d, v, "sslvpn_kxp_hardware_acceleration", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-kxp-hardware-acceleration"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_cipher_hardware_acceleration"); ok {

		t, err := expandSystemGlobalSslvpnCipherHardwareAcceleration(d, v, "sslvpn_cipher_hardware_acceleration", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-cipher-hardware-acceleration"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn_plugin_version_check"); ok {

		t, err := expandSystemGlobalSslvpnPluginVersionCheck(d, v, "sslvpn_plugin_version_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn-plugin-version-check"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_ftk_expiry"); ok {

		t, err := expandSystemGlobalTwoFactorFtkExpiry(d, v, "two_factor_ftk_expiry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-ftk-expiry"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_email_expiry"); ok {

		t, err := expandSystemGlobalTwoFactorEmailExpiry(d, v, "two_factor_email_expiry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-email-expiry"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_sms_expiry"); ok {

		t, err := expandSystemGlobalTwoFactorSmsExpiry(d, v, "two_factor_sms_expiry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-sms-expiry"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_fac_expiry"); ok {

		t, err := expandSystemGlobalTwoFactorFacExpiry(d, v, "two_factor_fac_expiry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-fac-expiry"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_ftm_expiry"); ok {

		t, err := expandSystemGlobalTwoFactorFtmExpiry(d, v, "two_factor_ftm_expiry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-ftm-expiry"] = t
		}
	}

	if v, ok := d.GetOk("per_user_bal"); ok {

		t, err := expandSystemGlobalPerUserBal(d, v, "per_user_bal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-user-bal"] = t
		}
	}

	if v, ok := d.GetOk("per_user_bwl"); ok {

		t, err := expandSystemGlobalPerUserBwl(d, v, "per_user_bwl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-user-bwl"] = t
		}
	}

	if v, ok := d.GetOkExists("virtual_server_count"); ok {

		t, err := expandSystemGlobalVirtualServerCount(d, v, "virtual_server_count", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-server-count"] = t
		}
	}

	if v, ok := d.GetOk("virtual_server_hardware_acceleration"); ok {

		t, err := expandSystemGlobalVirtualServerHardwareAcceleration(d, v, "virtual_server_hardware_acceleration", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-server-hardware-acceleration"] = t
		}
	}

	if v, ok := d.GetOkExists("wad_worker_count"); ok {

		t, err := expandSystemGlobalWadWorkerCount(d, v, "wad_worker_count", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wad-worker-count"] = t
		}
	}

	if v, ok := d.GetOk("wad_csvc_cs_count"); ok {

		t, err := expandSystemGlobalWadCsvcCsCount(d, v, "wad_csvc_cs_count", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wad-csvc-cs-count"] = t
		}
	}

	if v, ok := d.GetOkExists("wad_csvc_db_count"); ok {

		t, err := expandSystemGlobalWadCsvcDbCount(d, v, "wad_csvc_db_count", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wad-csvc-db-count"] = t
		}
	}

	if v, ok := d.GetOk("wad_source_affinity"); ok {

		t, err := expandSystemGlobalWadSourceAffinity(d, v, "wad_source_affinity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wad-source-affinity"] = t
		}
	}

	if v, ok := d.GetOk("wad_memory_change_granularity"); ok {

		t, err := expandSystemGlobalWadMemoryChangeGranularity(d, v, "wad_memory_change_granularity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wad-memory-change-granularity"] = t
		}
	}

	if v, ok := d.GetOk("login_timestamp"); ok {

		t, err := expandSystemGlobalLoginTimestamp(d, v, "login_timestamp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-timestamp"] = t
		}
	}

	if v, ok := d.GetOkExists("miglogd_children"); ok {

		t, err := expandSystemGlobalMiglogdChildren(d, v, "miglogd_children", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["miglogd-children"] = t
		}
	}

	if v, ok := d.GetOk("special_file_23_support"); ok {

		t, err := expandSystemGlobalSpecialFile23Support(d, v, "special_file_23_support", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["special-file-23-support"] = t
		}
	}

	if v, ok := d.GetOk("log_uuid_policy"); ok {

		t, err := expandSystemGlobalLogUuidPolicy(d, v, "log_uuid_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-uuid-policy"] = t
		}
	}

	if v, ok := d.GetOk("log_uuid_address"); ok {

		t, err := expandSystemGlobalLogUuidAddress(d, v, "log_uuid_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-uuid-address"] = t
		}
	}

	if v, ok := d.GetOk("log_ssl_connection"); ok {

		t, err := expandSystemGlobalLogSslConnection(d, v, "log_ssl_connection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-ssl-connection"] = t
		}
	}

	if v, ok := d.GetOk("gui_rest_api_cache"); ok {

		t, err := expandSystemGlobalGuiRestApiCache(d, v, "gui_rest_api_cache", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-rest-api-cache"] = t
		}
	}

	if v, ok := d.GetOk("arp_max_entry"); ok {

		t, err := expandSystemGlobalArpMaxEntry(d, v, "arp_max_entry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-max-entry"] = t
		}
	}

	if v, ok := d.GetOk("ha_affinity"); ok {

		t, err := expandSystemGlobalHaAffinity(d, v, "ha_affinity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-affinity"] = t
		}
	}

	if v, ok := d.GetOk("cmdbsvr_affinity"); ok {

		t, err := expandSystemGlobalCmdbsvrAffinity(d, v, "cmdbsvr_affinity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmdbsvr-affinity"] = t
		}
	}

	if v, ok := d.GetOk("av_affinity"); ok {

		t, err := expandSystemGlobalAvAffinity(d, v, "av_affinity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-affinity"] = t
		}
	}

	if v, ok := d.GetOk("wad_affinity"); ok {

		t, err := expandSystemGlobalWadAffinity(d, v, "wad_affinity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wad-affinity"] = t
		}
	}

	if v, ok := d.GetOk("ips_affinity"); ok {

		t, err := expandSystemGlobalIpsAffinity(d, v, "ips_affinity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-affinity"] = t
		}
	}

	if v, ok := d.GetOk("miglog_affinity"); ok {

		t, err := expandSystemGlobalMiglogAffinity(d, v, "miglog_affinity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["miglog-affinity"] = t
		}
	}

	if v, ok := d.GetOk("url_filter_affinity"); ok {

		t, err := expandSystemGlobalUrlFilterAffinity(d, v, "url_filter_affinity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-filter-affinity"] = t
		}
	}

	if v, ok := d.GetOk("ndp_max_entry"); ok {

		t, err := expandSystemGlobalNdpMaxEntry(d, v, "ndp_max_entry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ndp-max-entry"] = t
		}
	}

	if v, ok := d.GetOk("br_fdb_max_entry"); ok {

		t, err := expandSystemGlobalBrFdbMaxEntry(d, v, "br_fdb_max_entry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["br-fdb-max-entry"] = t
		}
	}

	if v, ok := d.GetOkExists("max_route_cache_size"); ok {

		t, err := expandSystemGlobalMaxRouteCacheSize(d, v, "max_route_cache_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-route-cache-size"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_asic_offload"); ok {

		t, err := expandSystemGlobalIpsecAsicOffload(d, v, "ipsec_asic_offload", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-asic-offload"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_soft_dec_async"); ok {

		t, err := expandSystemGlobalIpsecSoftDecAsync(d, v, "ipsec_soft_dec_async", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-soft-dec-async"] = t
		}
	}

	if v, ok := d.GetOk("ike_embryonic_limit"); ok {

		t, err := expandSystemGlobalIkeEmbryonicLimit(d, v, "ike_embryonic_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-embryonic-limit"] = t
		}
	}

	if v, ok := d.GetOk("device_idle_timeout"); ok {

		t, err := expandSystemGlobalDeviceIdleTimeout(d, v, "device_idle_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("user_device_store_max_devices"); ok {

		t, err := expandSystemGlobalUserDeviceStoreMaxDevices(d, v, "user_device_store_max_devices", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-device-store-max-devices"] = t
		}
	}

	if v, ok := d.GetOk("user_device_store_max_users"); ok {

		t, err := expandSystemGlobalUserDeviceStoreMaxUsers(d, v, "user_device_store_max_users", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-device-store-max-users"] = t
		}
	}

	if v, ok := d.GetOk("user_device_store_max_unified_mem"); ok {

		t, err := expandSystemGlobalUserDeviceStoreMaxUnifiedMem(d, v, "user_device_store_max_unified_mem", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-device-store-max-unified-mem"] = t
		}
	}

	if v, ok := d.GetOk("device_identification_active_scan_delay"); ok {

		t, err := expandSystemGlobalDeviceIdentificationActiveScanDelay(d, v, "device_identification_active_scan_delay", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-identification-active-scan-delay"] = t
		}
	}

	if v, ok := d.GetOk("compliance_check"); ok {

		t, err := expandSystemGlobalComplianceCheck(d, v, "compliance_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compliance-check"] = t
		}
	}

	if v, ok := d.GetOk("compliance_check_time"); ok {

		t, err := expandSystemGlobalComplianceCheckTime(d, v, "compliance_check_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compliance-check-time"] = t
		}
	}

	if v, ok := d.GetOk("gui_device_latitude"); ok {

		t, err := expandSystemGlobalGuiDeviceLatitude(d, v, "gui_device_latitude", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-device-latitude"] = t
		}
	}

	if v, ok := d.GetOk("gui_device_longitude"); ok {

		t, err := expandSystemGlobalGuiDeviceLongitude(d, v, "gui_device_longitude", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-device-longitude"] = t
		}
	}

	if v, ok := d.GetOk("private_data_encryption"); ok {

		t, err := expandSystemGlobalPrivateDataEncryption(d, v, "private_data_encryption", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-data-encryption"] = t
		}
	}

	if v, ok := d.GetOk("auto_auth_extension_device"); ok {

		t, err := expandSystemGlobalAutoAuthExtensionDevice(d, v, "auto_auth_extension_device", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-auth-extension-device"] = t
		}
	}

	if v, ok := d.GetOk("gui_theme"); ok {

		t, err := expandSystemGlobalGuiTheme(d, v, "gui_theme", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-theme"] = t
		}
	}

	if v, ok := d.GetOk("gui_date_format"); ok {

		t, err := expandSystemGlobalGuiDateFormat(d, v, "gui_date_format", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-date-format"] = t
		}
	}

	if v, ok := d.GetOk("gui_date_time_source"); ok {

		t, err := expandSystemGlobalGuiDateTimeSource(d, v, "gui_date_time_source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-date-time-source"] = t
		}
	}

	if v, ok := d.GetOk("igmp_state_limit"); ok {

		t, err := expandSystemGlobalIgmpStateLimit(d, v, "igmp_state_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-state-limit"] = t
		}
	}

	if v, ok := d.GetOk("cloud_communication"); ok {

		t, err := expandSystemGlobalCloudCommunication(d, v, "cloud_communication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cloud-communication"] = t
		}
	}

	if v, ok := d.GetOk("fec_port"); ok {

		t, err := expandSystemGlobalFecPort(d, v, "fec_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-port"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_ha_seqjump_rate"); ok {

		t, err := expandSystemGlobalIpsecHaSeqjumpRate(d, v, "ipsec_ha_seqjump_rate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-ha-seqjump-rate"] = t
		}
	}

	if v, ok := d.GetOk("fortitoken_cloud"); ok {

		t, err := expandSystemGlobalFortitokenCloud(d, v, "fortitoken_cloud", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortitoken-cloud"] = t
		}
	}

	if v, ok := d.GetOkExists("faz_disk_buffer_size"); ok {

		t, err := expandSystemGlobalFazDiskBufferSize(d, v, "faz_disk_buffer_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-disk-buffer-size"] = t
		}
	}

	if v, ok := d.GetOk("irq_time_accounting"); ok {

		t, err := expandSystemGlobalIrqTimeAccounting(d, v, "irq_time_accounting", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["irq-time-accounting"] = t
		}
	}

	if v, ok := d.GetOk("fortiipam_integration"); ok {

		t, err := expandSystemGlobalFortiipamIntegration(d, v, "fortiipam_integration", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiipam-integration"] = t
		}
	}

	if v, ok := d.GetOk("management_ip"); ok {

		t, err := expandSystemGlobalManagementIp(d, v, "management_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["management-ip"] = t
		}
	}

	if v, ok := d.GetOk("management_port"); ok {

		t, err := expandSystemGlobalManagementPort(d, v, "management_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["management-port"] = t
		}
	}

	if v, ok := d.GetOk("management_port_use_admin_sport"); ok {

		t, err := expandSystemGlobalManagementPortUseAdminSport(d, v, "management_port_use_admin_sport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["management-port-use-admin-sport"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_database"); ok {

		t, err := expandSystemGlobalInternetServiceDatabase(d, v, "internet_service_database", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-database"] = t
		}
	}

	return &obj, nil
}
