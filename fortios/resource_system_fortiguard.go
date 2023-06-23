// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiGuard services.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFortiguard() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFortiguardUpdate,
		Read:   resourceSystemFortiguardRead,
		Update: resourceSystemFortiguardUpdate,
		Delete: resourceSystemFortiguardDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_account_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 50),
				Optional:     true,
				Computed:     true,
			},
			"load_balance_servers": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 266),
				Optional:     true,
				Computed:     true,
			},
			"auto_join_forticloud": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_server_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sandbox_region": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"sandbox_inline_scan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_ffdb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_uwdb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_dldb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_extdb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_build_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistent_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"auto_firmware_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_firmware_upgrade_day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_firmware_upgrade_delay": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 14),
				Optional:     true,
				Computed:     true,
			},
			"auto_firmware_upgrade_start_hour": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 23),
				Optional:     true,
				Computed:     true,
			},
			"auto_firmware_upgrade_end_hour": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 23),
				Optional:     true,
				Computed:     true,
			},
			"fds_license_expiring_days": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional:     true,
				Computed:     true,
			},
			"fortiguard_anycast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiguard_anycast_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"antispam_force_off": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"antispam_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"antispam_cache_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 86400),
				Optional:     true,
				Computed:     true,
			},
			"antispam_cache_mpermille": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 150),
				Optional:     true,
				Computed:     true,
			},
			"antispam_cache_mpercent": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 15),
				Optional:     true,
				Computed:     true,
			},
			"antispam_license": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"antispam_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"antispam_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),
				Required:     true,
			},
			"outbreak_prevention_force_off": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"outbreak_prevention_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"outbreak_prevention_cache_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 86400),
				Optional:     true,
				Computed:     true,
			},
			"outbreak_prevention_cache_mpermille": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 150),
				Optional:     true,
				Computed:     true,
			},
			"outbreak_prevention_cache_mpercent": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 15),
				Optional:     true,
				Computed:     true,
			},
			"outbreak_prevention_license": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"outbreak_prevention_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"outbreak_prevention_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),
				Required:     true,
			},
			"webfilter_force_off": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_cache_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 86400),
				Optional:     true,
				Computed:     true,
			},
			"webfilter_license": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"webfilter_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"webfilter_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),
				Required:     true,
			},
			"sdns_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdns_server_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"anycast_sdns_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"anycast_sdns_server_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"sdns_options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_server_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"proxy_username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"proxy_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"videofilter_license": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"videofilter_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ddns_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_server_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_server_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemFortiguardUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemFortiguard(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemFortiguard resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemFortiguard(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemFortiguard resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemFortiguard")
	}

	return resourceSystemFortiguardRead(d, m)
}

func resourceSystemFortiguardDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemFortiguard(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemFortiguard resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFortiguard(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemFortiguard resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFortiguardRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemFortiguard(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemFortiguard resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFortiguard(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemFortiguard resource from API: %v", err)
	}
	return nil
}

func flattenSystemFortiguardProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardServiceAccountId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardLoadBalanceServers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAutoJoinForticloud(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardUpdateServerLocation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardSandboxRegion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardSandboxInlineScan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardUpdateFfdb(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardUpdateUwdb(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardUpdateDldb(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardUpdateExtdb(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardUpdateBuildProxy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardPersistentConnection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAutoFirmwareUpgrade(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAutoFirmwareUpgradeDay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAutoFirmwareUpgradeDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAutoFirmwareUpgradeStartHour(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAutoFirmwareUpgradeEndHour(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardFdsLicenseExpiringDays(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardFortiguardAnycast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardFortiguardAnycastSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamForceOff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamCacheTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamCacheMpermille(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamCacheMpercent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamLicense(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamExpiration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionForceOff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionCacheTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionCacheMpermille(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionCacheMpercent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionLicense(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionExpiration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterForceOff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterCacheTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterLicense(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterExpiration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardSdnsServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardSdnsServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAnycastSdnsServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardAnycastSdnsServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardSdnsOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardSourceIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardProxyServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardProxyServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardProxyUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardProxyPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardVideofilterLicense(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardVideofilterExpiration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardDdnsServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardDdnsServerIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardDdnsServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortiguardInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemFortiguard(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("protocol", flattenSystemFortiguardProtocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemFortiguardPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("service_account_id", flattenSystemFortiguardServiceAccountId(o["service-account-id"], d, "service_account_id", sv)); err != nil {
		if !fortiAPIPatch(o["service-account-id"]) {
			return fmt.Errorf("Error reading service_account_id: %v", err)
		}
	}

	if err = d.Set("load_balance_servers", flattenSystemFortiguardLoadBalanceServers(o["load-balance-servers"], d, "load_balance_servers", sv)); err != nil {
		if !fortiAPIPatch(o["load-balance-servers"]) {
			return fmt.Errorf("Error reading load_balance_servers: %v", err)
		}
	}

	if err = d.Set("auto_join_forticloud", flattenSystemFortiguardAutoJoinForticloud(o["auto-join-forticloud"], d, "auto_join_forticloud", sv)); err != nil {
		if !fortiAPIPatch(o["auto-join-forticloud"]) {
			return fmt.Errorf("Error reading auto_join_forticloud: %v", err)
		}
	}

	if err = d.Set("update_server_location", flattenSystemFortiguardUpdateServerLocation(o["update-server-location"], d, "update_server_location", sv)); err != nil {
		if !fortiAPIPatch(o["update-server-location"]) {
			return fmt.Errorf("Error reading update_server_location: %v", err)
		}
	}

	if err = d.Set("sandbox_region", flattenSystemFortiguardSandboxRegion(o["sandbox-region"], d, "sandbox_region", sv)); err != nil {
		if !fortiAPIPatch(o["sandbox-region"]) {
			return fmt.Errorf("Error reading sandbox_region: %v", err)
		}
	}

	if err = d.Set("sandbox_inline_scan", flattenSystemFortiguardSandboxInlineScan(o["sandbox-inline-scan"], d, "sandbox_inline_scan", sv)); err != nil {
		if !fortiAPIPatch(o["sandbox-inline-scan"]) {
			return fmt.Errorf("Error reading sandbox_inline_scan: %v", err)
		}
	}

	if err = d.Set("update_ffdb", flattenSystemFortiguardUpdateFfdb(o["update-ffdb"], d, "update_ffdb", sv)); err != nil {
		if !fortiAPIPatch(o["update-ffdb"]) {
			return fmt.Errorf("Error reading update_ffdb: %v", err)
		}
	}

	if err = d.Set("update_uwdb", flattenSystemFortiguardUpdateUwdb(o["update-uwdb"], d, "update_uwdb", sv)); err != nil {
		if !fortiAPIPatch(o["update-uwdb"]) {
			return fmt.Errorf("Error reading update_uwdb: %v", err)
		}
	}

	if err = d.Set("update_dldb", flattenSystemFortiguardUpdateDldb(o["update-dldb"], d, "update_dldb", sv)); err != nil {
		if !fortiAPIPatch(o["update-dldb"]) {
			return fmt.Errorf("Error reading update_dldb: %v", err)
		}
	}

	if err = d.Set("update_extdb", flattenSystemFortiguardUpdateExtdb(o["update-extdb"], d, "update_extdb", sv)); err != nil {
		if !fortiAPIPatch(o["update-extdb"]) {
			return fmt.Errorf("Error reading update_extdb: %v", err)
		}
	}

	if err = d.Set("update_build_proxy", flattenSystemFortiguardUpdateBuildProxy(o["update-build-proxy"], d, "update_build_proxy", sv)); err != nil {
		if !fortiAPIPatch(o["update-build-proxy"]) {
			return fmt.Errorf("Error reading update_build_proxy: %v", err)
		}
	}

	if err = d.Set("persistent_connection", flattenSystemFortiguardPersistentConnection(o["persistent-connection"], d, "persistent_connection", sv)); err != nil {
		if !fortiAPIPatch(o["persistent-connection"]) {
			return fmt.Errorf("Error reading persistent_connection: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemFortiguardVdom(o["vdom"], d, "vdom", sv)); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("auto_firmware_upgrade", flattenSystemFortiguardAutoFirmwareUpgrade(o["auto-firmware-upgrade"], d, "auto_firmware_upgrade", sv)); err != nil {
		if !fortiAPIPatch(o["auto-firmware-upgrade"]) {
			return fmt.Errorf("Error reading auto_firmware_upgrade: %v", err)
		}
	}

	if err = d.Set("auto_firmware_upgrade_day", flattenSystemFortiguardAutoFirmwareUpgradeDay(o["auto-firmware-upgrade-day"], d, "auto_firmware_upgrade_day", sv)); err != nil {
		if !fortiAPIPatch(o["auto-firmware-upgrade-day"]) {
			return fmt.Errorf("Error reading auto_firmware_upgrade_day: %v", err)
		}
	}

	if err = d.Set("auto_firmware_upgrade_delay", flattenSystemFortiguardAutoFirmwareUpgradeDelay(o["auto-firmware-upgrade-delay"], d, "auto_firmware_upgrade_delay", sv)); err != nil {
		if !fortiAPIPatch(o["auto-firmware-upgrade-delay"]) {
			return fmt.Errorf("Error reading auto_firmware_upgrade_delay: %v", err)
		}
	}

	if err = d.Set("auto_firmware_upgrade_start_hour", flattenSystemFortiguardAutoFirmwareUpgradeStartHour(o["auto-firmware-upgrade-start-hour"], d, "auto_firmware_upgrade_start_hour", sv)); err != nil {
		if !fortiAPIPatch(o["auto-firmware-upgrade-start-hour"]) {
			return fmt.Errorf("Error reading auto_firmware_upgrade_start_hour: %v", err)
		}
	}

	if err = d.Set("auto_firmware_upgrade_end_hour", flattenSystemFortiguardAutoFirmwareUpgradeEndHour(o["auto-firmware-upgrade-end-hour"], d, "auto_firmware_upgrade_end_hour", sv)); err != nil {
		if !fortiAPIPatch(o["auto-firmware-upgrade-end-hour"]) {
			return fmt.Errorf("Error reading auto_firmware_upgrade_end_hour: %v", err)
		}
	}

	if err = d.Set("fds_license_expiring_days", flattenSystemFortiguardFdsLicenseExpiringDays(o["FDS-license-expiring-days"], d, "fds_license_expiring_days", sv)); err != nil {
		if !fortiAPIPatch(o["FDS-license-expiring-days"]) {
			return fmt.Errorf("Error reading fds_license_expiring_days: %v", err)
		}
	}

	if err = d.Set("fortiguard_anycast", flattenSystemFortiguardFortiguardAnycast(o["fortiguard-anycast"], d, "fortiguard_anycast", sv)); err != nil {
		if !fortiAPIPatch(o["fortiguard-anycast"]) {
			return fmt.Errorf("Error reading fortiguard_anycast: %v", err)
		}
	}

	if err = d.Set("fortiguard_anycast_source", flattenSystemFortiguardFortiguardAnycastSource(o["fortiguard-anycast-source"], d, "fortiguard_anycast_source", sv)); err != nil {
		if !fortiAPIPatch(o["fortiguard-anycast-source"]) {
			return fmt.Errorf("Error reading fortiguard_anycast_source: %v", err)
		}
	}

	if err = d.Set("antispam_force_off", flattenSystemFortiguardAntispamForceOff(o["antispam-force-off"], d, "antispam_force_off", sv)); err != nil {
		if !fortiAPIPatch(o["antispam-force-off"]) {
			return fmt.Errorf("Error reading antispam_force_off: %v", err)
		}
	}

	if err = d.Set("antispam_cache", flattenSystemFortiguardAntispamCache(o["antispam-cache"], d, "antispam_cache", sv)); err != nil {
		if !fortiAPIPatch(o["antispam-cache"]) {
			return fmt.Errorf("Error reading antispam_cache: %v", err)
		}
	}

	if err = d.Set("antispam_cache_ttl", flattenSystemFortiguardAntispamCacheTtl(o["antispam-cache-ttl"], d, "antispam_cache_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["antispam-cache-ttl"]) {
			return fmt.Errorf("Error reading antispam_cache_ttl: %v", err)
		}
	}

	if err = d.Set("antispam_cache_mpermille", flattenSystemFortiguardAntispamCacheMpermille(o["antispam-cache-mpermille"], d, "antispam_cache_mpermille", sv)); err != nil {
		if !fortiAPIPatch(o["antispam-cache-mpermille"]) {
			return fmt.Errorf("Error reading antispam_cache_mpermille: %v", err)
		}
	}

	if err = d.Set("antispam_cache_mpercent", flattenSystemFortiguardAntispamCacheMpercent(o["antispam-cache-mpercent"], d, "antispam_cache_mpercent", sv)); err != nil {
		if !fortiAPIPatch(o["antispam-cache-mpercent"]) {
			return fmt.Errorf("Error reading antispam_cache_mpercent: %v", err)
		}
	}

	if err = d.Set("antispam_license", flattenSystemFortiguardAntispamLicense(o["antispam-license"], d, "antispam_license", sv)); err != nil {
		if !fortiAPIPatch(o["antispam-license"]) {
			return fmt.Errorf("Error reading antispam_license: %v", err)
		}
	}

	if err = d.Set("antispam_expiration", flattenSystemFortiguardAntispamExpiration(o["antispam-expiration"], d, "antispam_expiration", sv)); err != nil {
		if !fortiAPIPatch(o["antispam-expiration"]) {
			return fmt.Errorf("Error reading antispam_expiration: %v", err)
		}
	}

	if err = d.Set("antispam_timeout", flattenSystemFortiguardAntispamTimeout(o["antispam-timeout"], d, "antispam_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["antispam-timeout"]) {
			return fmt.Errorf("Error reading antispam_timeout: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_force_off", flattenSystemFortiguardOutbreakPreventionForceOff(o["outbreak-prevention-force-off"], d, "outbreak_prevention_force_off", sv)); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-force-off"]) {
			return fmt.Errorf("Error reading outbreak_prevention_force_off: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_cache", flattenSystemFortiguardOutbreakPreventionCache(o["outbreak-prevention-cache"], d, "outbreak_prevention_cache", sv)); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-cache"]) {
			return fmt.Errorf("Error reading outbreak_prevention_cache: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_cache_ttl", flattenSystemFortiguardOutbreakPreventionCacheTtl(o["outbreak-prevention-cache-ttl"], d, "outbreak_prevention_cache_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-cache-ttl"]) {
			return fmt.Errorf("Error reading outbreak_prevention_cache_ttl: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_cache_mpermille", flattenSystemFortiguardOutbreakPreventionCacheMpermille(o["outbreak-prevention-cache-mpermille"], d, "outbreak_prevention_cache_mpermille", sv)); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-cache-mpermille"]) {
			return fmt.Errorf("Error reading outbreak_prevention_cache_mpermille: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_cache_mpercent", flattenSystemFortiguardOutbreakPreventionCacheMpercent(o["outbreak-prevention-cache-mpercent"], d, "outbreak_prevention_cache_mpercent", sv)); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-cache-mpercent"]) {
			return fmt.Errorf("Error reading outbreak_prevention_cache_mpercent: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_license", flattenSystemFortiguardOutbreakPreventionLicense(o["outbreak-prevention-license"], d, "outbreak_prevention_license", sv)); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-license"]) {
			return fmt.Errorf("Error reading outbreak_prevention_license: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_expiration", flattenSystemFortiguardOutbreakPreventionExpiration(o["outbreak-prevention-expiration"], d, "outbreak_prevention_expiration", sv)); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-expiration"]) {
			return fmt.Errorf("Error reading outbreak_prevention_expiration: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_timeout", flattenSystemFortiguardOutbreakPreventionTimeout(o["outbreak-prevention-timeout"], d, "outbreak_prevention_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-timeout"]) {
			return fmt.Errorf("Error reading outbreak_prevention_timeout: %v", err)
		}
	}

	if err = d.Set("webfilter_force_off", flattenSystemFortiguardWebfilterForceOff(o["webfilter-force-off"], d, "webfilter_force_off", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-force-off"]) {
			return fmt.Errorf("Error reading webfilter_force_off: %v", err)
		}
	}

	if err = d.Set("webfilter_cache", flattenSystemFortiguardWebfilterCache(o["webfilter-cache"], d, "webfilter_cache", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-cache"]) {
			return fmt.Errorf("Error reading webfilter_cache: %v", err)
		}
	}

	if err = d.Set("webfilter_cache_ttl", flattenSystemFortiguardWebfilterCacheTtl(o["webfilter-cache-ttl"], d, "webfilter_cache_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-cache-ttl"]) {
			return fmt.Errorf("Error reading webfilter_cache_ttl: %v", err)
		}
	}

	if err = d.Set("webfilter_license", flattenSystemFortiguardWebfilterLicense(o["webfilter-license"], d, "webfilter_license", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-license"]) {
			return fmt.Errorf("Error reading webfilter_license: %v", err)
		}
	}

	if err = d.Set("webfilter_expiration", flattenSystemFortiguardWebfilterExpiration(o["webfilter-expiration"], d, "webfilter_expiration", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-expiration"]) {
			return fmt.Errorf("Error reading webfilter_expiration: %v", err)
		}
	}

	if err = d.Set("webfilter_timeout", flattenSystemFortiguardWebfilterTimeout(o["webfilter-timeout"], d, "webfilter_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-timeout"]) {
			return fmt.Errorf("Error reading webfilter_timeout: %v", err)
		}
	}

	if err = d.Set("sdns_server_ip", flattenSystemFortiguardSdnsServerIp(o["sdns-server-ip"], d, "sdns_server_ip", sv)); err != nil {
		if !fortiAPIPatch(o["sdns-server-ip"]) {
			return fmt.Errorf("Error reading sdns_server_ip: %v", err)
		}
	}

	if err = d.Set("sdns_server_port", flattenSystemFortiguardSdnsServerPort(o["sdns-server-port"], d, "sdns_server_port", sv)); err != nil {
		if !fortiAPIPatch(o["sdns-server-port"]) {
			return fmt.Errorf("Error reading sdns_server_port: %v", err)
		}
	}

	if err = d.Set("anycast_sdns_server_ip", flattenSystemFortiguardAnycastSdnsServerIp(o["anycast-sdns-server-ip"], d, "anycast_sdns_server_ip", sv)); err != nil {
		if !fortiAPIPatch(o["anycast-sdns-server-ip"]) {
			return fmt.Errorf("Error reading anycast_sdns_server_ip: %v", err)
		}
	}

	if err = d.Set("anycast_sdns_server_port", flattenSystemFortiguardAnycastSdnsServerPort(o["anycast-sdns-server-port"], d, "anycast_sdns_server_port", sv)); err != nil {
		if !fortiAPIPatch(o["anycast-sdns-server-port"]) {
			return fmt.Errorf("Error reading anycast_sdns_server_port: %v", err)
		}
	}

	if err = d.Set("sdns_options", flattenSystemFortiguardSdnsOptions(o["sdns-options"], d, "sdns_options", sv)); err != nil {
		if !fortiAPIPatch(o["sdns-options"]) {
			return fmt.Errorf("Error reading sdns_options: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemFortiguardSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenSystemFortiguardSourceIp6(o["source-ip6"], d, "source_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip6"]) {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("proxy_server_ip", flattenSystemFortiguardProxyServerIp(o["proxy-server-ip"], d, "proxy_server_ip", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-server-ip"]) {
			return fmt.Errorf("Error reading proxy_server_ip: %v", err)
		}
	}

	if err = d.Set("proxy_server_port", flattenSystemFortiguardProxyServerPort(o["proxy-server-port"], d, "proxy_server_port", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-server-port"]) {
			return fmt.Errorf("Error reading proxy_server_port: %v", err)
		}
	}

	if err = d.Set("proxy_username", flattenSystemFortiguardProxyUsername(o["proxy-username"], d, "proxy_username", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-username"]) {
			return fmt.Errorf("Error reading proxy_username: %v", err)
		}
	}

	if err = d.Set("videofilter_license", flattenSystemFortiguardVideofilterLicense(o["videofilter-license"], d, "videofilter_license", sv)); err != nil {
		if !fortiAPIPatch(o["videofilter-license"]) {
			return fmt.Errorf("Error reading videofilter_license: %v", err)
		}
	}

	if err = d.Set("videofilter_expiration", flattenSystemFortiguardVideofilterExpiration(o["videofilter-expiration"], d, "videofilter_expiration", sv)); err != nil {
		if !fortiAPIPatch(o["videofilter-expiration"]) {
			return fmt.Errorf("Error reading videofilter_expiration: %v", err)
		}
	}

	if err = d.Set("ddns_server_ip", flattenSystemFortiguardDdnsServerIp(o["ddns-server-ip"], d, "ddns_server_ip", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-server-ip"]) {
			return fmt.Errorf("Error reading ddns_server_ip: %v", err)
		}
	}

	if err = d.Set("ddns_server_ip6", flattenSystemFortiguardDdnsServerIp6(o["ddns-server-ip6"], d, "ddns_server_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-server-ip6"]) {
			return fmt.Errorf("Error reading ddns_server_ip6: %v", err)
		}
	}

	if err = d.Set("ddns_server_port", flattenSystemFortiguardDdnsServerPort(o["ddns-server-port"], d, "ddns_server_port", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-server-port"]) {
			return fmt.Errorf("Error reading ddns_server_port: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenSystemFortiguardInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemFortiguardInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func flattenSystemFortiguardFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemFortiguardProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardServiceAccountId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardLoadBalanceServers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAutoJoinForticloud(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardUpdateServerLocation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSandboxRegion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSandboxInlineScan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardUpdateFfdb(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardUpdateUwdb(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardUpdateDldb(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardUpdateExtdb(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardUpdateBuildProxy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardPersistentConnection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAutoFirmwareUpgrade(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAutoFirmwareUpgradeDay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAutoFirmwareUpgradeDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAutoFirmwareUpgradeStartHour(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAutoFirmwareUpgradeEndHour(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardFdsLicenseExpiringDays(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardFortiguardAnycast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardFortiguardAnycastSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamForceOff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamCacheTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamCacheMpermille(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamCacheMpercent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamLicense(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamExpiration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionForceOff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionCacheTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionCacheMpermille(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionCacheMpercent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionLicense(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionExpiration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterForceOff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterCacheTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterLicense(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterExpiration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSdnsServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSdnsServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAnycastSdnsServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAnycastSdnsServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSdnsOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSourceIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardProxyServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardProxyServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardProxyUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardProxyPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardVideofilterLicense(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardVideofilterExpiration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardDdnsServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardDdnsServerIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardDdnsServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFortiguard(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("protocol"); ok {
		if setArgNil {
			obj["protocol"] = nil
		} else {
			t, err := expandSystemFortiguardProtocol(d, v, "protocol", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["protocol"] = t
			}
		}
	}

	if v, ok := d.GetOk("port"); ok {
		if setArgNil {
			obj["port"] = nil
		} else {
			t, err := expandSystemFortiguardPort(d, v, "port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["port"] = t
			}
		}
	}

	if v, ok := d.GetOk("service_account_id"); ok {
		if setArgNil {
			obj["service-account-id"] = nil
		} else {
			t, err := expandSystemFortiguardServiceAccountId(d, v, "service_account_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["service-account-id"] = t
			}
		}
	}

	if v, ok := d.GetOk("load_balance_servers"); ok {
		if setArgNil {
			obj["load-balance-servers"] = nil
		} else {
			t, err := expandSystemFortiguardLoadBalanceServers(d, v, "load_balance_servers", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["load-balance-servers"] = t
			}
		}
	}

	if v, ok := d.GetOk("auto_join_forticloud"); ok {
		if setArgNil {
			obj["auto-join-forticloud"] = nil
		} else {
			t, err := expandSystemFortiguardAutoJoinForticloud(d, v, "auto_join_forticloud", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-join-forticloud"] = t
			}
		}
	}

	if v, ok := d.GetOk("update_server_location"); ok {
		if setArgNil {
			obj["update-server-location"] = nil
		} else {
			t, err := expandSystemFortiguardUpdateServerLocation(d, v, "update_server_location", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["update-server-location"] = t
			}
		}
	}

	if v, ok := d.GetOk("sandbox_region"); ok {
		if setArgNil {
			obj["sandbox-region"] = nil
		} else {
			t, err := expandSystemFortiguardSandboxRegion(d, v, "sandbox_region", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sandbox-region"] = t
			}
		}
	}

	if v, ok := d.GetOk("sandbox_inline_scan"); ok {
		if setArgNil {
			obj["sandbox-inline-scan"] = nil
		} else {
			t, err := expandSystemFortiguardSandboxInlineScan(d, v, "sandbox_inline_scan", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sandbox-inline-scan"] = t
			}
		}
	}

	if v, ok := d.GetOk("update_ffdb"); ok {
		if setArgNil {
			obj["update-ffdb"] = nil
		} else {
			t, err := expandSystemFortiguardUpdateFfdb(d, v, "update_ffdb", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["update-ffdb"] = t
			}
		}
	}

	if v, ok := d.GetOk("update_uwdb"); ok {
		if setArgNil {
			obj["update-uwdb"] = nil
		} else {
			t, err := expandSystemFortiguardUpdateUwdb(d, v, "update_uwdb", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["update-uwdb"] = t
			}
		}
	}

	if v, ok := d.GetOk("update_dldb"); ok {
		if setArgNil {
			obj["update-dldb"] = nil
		} else {
			t, err := expandSystemFortiguardUpdateDldb(d, v, "update_dldb", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["update-dldb"] = t
			}
		}
	}

	if v, ok := d.GetOk("update_extdb"); ok {
		if setArgNil {
			obj["update-extdb"] = nil
		} else {
			t, err := expandSystemFortiguardUpdateExtdb(d, v, "update_extdb", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["update-extdb"] = t
			}
		}
	}

	if v, ok := d.GetOk("update_build_proxy"); ok {
		if setArgNil {
			obj["update-build-proxy"] = nil
		} else {
			t, err := expandSystemFortiguardUpdateBuildProxy(d, v, "update_build_proxy", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["update-build-proxy"] = t
			}
		}
	}

	if v, ok := d.GetOk("persistent_connection"); ok {
		if setArgNil {
			obj["persistent-connection"] = nil
		} else {
			t, err := expandSystemFortiguardPersistentConnection(d, v, "persistent_connection", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["persistent-connection"] = t
			}
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		if setArgNil {
			obj["vdom"] = nil
		} else {
			t, err := expandSystemFortiguardVdom(d, v, "vdom", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vdom"] = t
			}
		}
	}

	if v, ok := d.GetOk("auto_firmware_upgrade"); ok {
		if setArgNil {
			obj["auto-firmware-upgrade"] = nil
		} else {
			t, err := expandSystemFortiguardAutoFirmwareUpgrade(d, v, "auto_firmware_upgrade", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-firmware-upgrade"] = t
			}
		}
	}

	if v, ok := d.GetOk("auto_firmware_upgrade_day"); ok {
		if setArgNil {
			obj["auto-firmware-upgrade-day"] = nil
		} else {
			t, err := expandSystemFortiguardAutoFirmwareUpgradeDay(d, v, "auto_firmware_upgrade_day", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-firmware-upgrade-day"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("auto_firmware_upgrade_delay"); ok {
		if setArgNil {
			obj["auto-firmware-upgrade-delay"] = nil
		} else {
			t, err := expandSystemFortiguardAutoFirmwareUpgradeDelay(d, v, "auto_firmware_upgrade_delay", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-firmware-upgrade-delay"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("auto_firmware_upgrade_start_hour"); ok {
		if setArgNil {
			obj["auto-firmware-upgrade-start-hour"] = nil
		} else {
			t, err := expandSystemFortiguardAutoFirmwareUpgradeStartHour(d, v, "auto_firmware_upgrade_start_hour", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-firmware-upgrade-start-hour"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("auto_firmware_upgrade_end_hour"); ok {
		if setArgNil {
			obj["auto-firmware-upgrade-end-hour"] = nil
		} else {
			t, err := expandSystemFortiguardAutoFirmwareUpgradeEndHour(d, v, "auto_firmware_upgrade_end_hour", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-firmware-upgrade-end-hour"] = t
			}
		}
	}

	if v, ok := d.GetOk("fds_license_expiring_days"); ok {
		if setArgNil {
			obj["FDS-license-expiring-days"] = nil
		} else {
			t, err := expandSystemFortiguardFdsLicenseExpiringDays(d, v, "fds_license_expiring_days", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["FDS-license-expiring-days"] = t
			}
		}
	}

	if v, ok := d.GetOk("fortiguard_anycast"); ok {
		if setArgNil {
			obj["fortiguard-anycast"] = nil
		} else {
			t, err := expandSystemFortiguardFortiguardAnycast(d, v, "fortiguard_anycast", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fortiguard-anycast"] = t
			}
		}
	}

	if v, ok := d.GetOk("fortiguard_anycast_source"); ok {
		if setArgNil {
			obj["fortiguard-anycast-source"] = nil
		} else {
			t, err := expandSystemFortiguardFortiguardAnycastSource(d, v, "fortiguard_anycast_source", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fortiguard-anycast-source"] = t
			}
		}
	}

	if v, ok := d.GetOk("antispam_force_off"); ok {
		if setArgNil {
			obj["antispam-force-off"] = nil
		} else {
			t, err := expandSystemFortiguardAntispamForceOff(d, v, "antispam_force_off", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["antispam-force-off"] = t
			}
		}
	}

	if v, ok := d.GetOk("antispam_cache"); ok {
		if setArgNil {
			obj["antispam-cache"] = nil
		} else {
			t, err := expandSystemFortiguardAntispamCache(d, v, "antispam_cache", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["antispam-cache"] = t
			}
		}
	}

	if v, ok := d.GetOk("antispam_cache_ttl"); ok {
		if setArgNil {
			obj["antispam-cache-ttl"] = nil
		} else {
			t, err := expandSystemFortiguardAntispamCacheTtl(d, v, "antispam_cache_ttl", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["antispam-cache-ttl"] = t
			}
		}
	}

	if v, ok := d.GetOk("antispam_cache_mpermille"); ok {
		if setArgNil {
			obj["antispam-cache-mpermille"] = nil
		} else {
			t, err := expandSystemFortiguardAntispamCacheMpermille(d, v, "antispam_cache_mpermille", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["antispam-cache-mpermille"] = t
			}
		}
	}

	if v, ok := d.GetOk("antispam_cache_mpercent"); ok {
		if setArgNil {
			obj["antispam-cache-mpercent"] = nil
		} else {
			t, err := expandSystemFortiguardAntispamCacheMpercent(d, v, "antispam_cache_mpercent", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["antispam-cache-mpercent"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("antispam_license"); ok {
		if setArgNil {
			obj["antispam-license"] = nil
		} else {
			t, err := expandSystemFortiguardAntispamLicense(d, v, "antispam_license", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["antispam-license"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("antispam_expiration"); ok {
		if setArgNil {
			obj["antispam-expiration"] = nil
		} else {
			t, err := expandSystemFortiguardAntispamExpiration(d, v, "antispam_expiration", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["antispam-expiration"] = t
			}
		}
	}

	if v, ok := d.GetOk("antispam_timeout"); ok {
		if setArgNil {
			obj["antispam-timeout"] = nil
		} else {
			t, err := expandSystemFortiguardAntispamTimeout(d, v, "antispam_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["antispam-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_force_off"); ok {
		if setArgNil {
			obj["outbreak-prevention-force-off"] = nil
		} else {
			t, err := expandSystemFortiguardOutbreakPreventionForceOff(d, v, "outbreak_prevention_force_off", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["outbreak-prevention-force-off"] = t
			}
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_cache"); ok {
		if setArgNil {
			obj["outbreak-prevention-cache"] = nil
		} else {
			t, err := expandSystemFortiguardOutbreakPreventionCache(d, v, "outbreak_prevention_cache", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["outbreak-prevention-cache"] = t
			}
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_cache_ttl"); ok {
		if setArgNil {
			obj["outbreak-prevention-cache-ttl"] = nil
		} else {
			t, err := expandSystemFortiguardOutbreakPreventionCacheTtl(d, v, "outbreak_prevention_cache_ttl", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["outbreak-prevention-cache-ttl"] = t
			}
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_cache_mpermille"); ok {
		if setArgNil {
			obj["outbreak-prevention-cache-mpermille"] = nil
		} else {
			t, err := expandSystemFortiguardOutbreakPreventionCacheMpermille(d, v, "outbreak_prevention_cache_mpermille", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["outbreak-prevention-cache-mpermille"] = t
			}
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_cache_mpercent"); ok {
		if setArgNil {
			obj["outbreak-prevention-cache-mpercent"] = nil
		} else {
			t, err := expandSystemFortiguardOutbreakPreventionCacheMpercent(d, v, "outbreak_prevention_cache_mpercent", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["outbreak-prevention-cache-mpercent"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("outbreak_prevention_license"); ok {
		if setArgNil {
			obj["outbreak-prevention-license"] = nil
		} else {
			t, err := expandSystemFortiguardOutbreakPreventionLicense(d, v, "outbreak_prevention_license", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["outbreak-prevention-license"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("outbreak_prevention_expiration"); ok {
		if setArgNil {
			obj["outbreak-prevention-expiration"] = nil
		} else {
			t, err := expandSystemFortiguardOutbreakPreventionExpiration(d, v, "outbreak_prevention_expiration", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["outbreak-prevention-expiration"] = t
			}
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_timeout"); ok {
		if setArgNil {
			obj["outbreak-prevention-timeout"] = nil
		} else {
			t, err := expandSystemFortiguardOutbreakPreventionTimeout(d, v, "outbreak_prevention_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["outbreak-prevention-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("webfilter_force_off"); ok {
		if setArgNil {
			obj["webfilter-force-off"] = nil
		} else {
			t, err := expandSystemFortiguardWebfilterForceOff(d, v, "webfilter_force_off", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["webfilter-force-off"] = t
			}
		}
	}

	if v, ok := d.GetOk("webfilter_cache"); ok {
		if setArgNil {
			obj["webfilter-cache"] = nil
		} else {
			t, err := expandSystemFortiguardWebfilterCache(d, v, "webfilter_cache", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["webfilter-cache"] = t
			}
		}
	}

	if v, ok := d.GetOk("webfilter_cache_ttl"); ok {
		if setArgNil {
			obj["webfilter-cache-ttl"] = nil
		} else {
			t, err := expandSystemFortiguardWebfilterCacheTtl(d, v, "webfilter_cache_ttl", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["webfilter-cache-ttl"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("webfilter_license"); ok {
		if setArgNil {
			obj["webfilter-license"] = nil
		} else {
			t, err := expandSystemFortiguardWebfilterLicense(d, v, "webfilter_license", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["webfilter-license"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("webfilter_expiration"); ok {
		if setArgNil {
			obj["webfilter-expiration"] = nil
		} else {
			t, err := expandSystemFortiguardWebfilterExpiration(d, v, "webfilter_expiration", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["webfilter-expiration"] = t
			}
		}
	}

	if v, ok := d.GetOk("webfilter_timeout"); ok {
		if setArgNil {
			obj["webfilter-timeout"] = nil
		} else {
			t, err := expandSystemFortiguardWebfilterTimeout(d, v, "webfilter_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["webfilter-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("sdns_server_ip"); ok {
		if setArgNil {
			obj["sdns-server-ip"] = nil
		} else {
			t, err := expandSystemFortiguardSdnsServerIp(d, v, "sdns_server_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sdns-server-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("sdns_server_port"); ok {
		if setArgNil {
			obj["sdns-server-port"] = nil
		} else {
			t, err := expandSystemFortiguardSdnsServerPort(d, v, "sdns_server_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sdns-server-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("anycast_sdns_server_ip"); ok {
		if setArgNil {
			obj["anycast-sdns-server-ip"] = nil
		} else {
			t, err := expandSystemFortiguardAnycastSdnsServerIp(d, v, "anycast_sdns_server_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["anycast-sdns-server-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("anycast_sdns_server_port"); ok {
		if setArgNil {
			obj["anycast-sdns-server-port"] = nil
		} else {
			t, err := expandSystemFortiguardAnycastSdnsServerPort(d, v, "anycast_sdns_server_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["anycast-sdns-server-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("sdns_options"); ok {
		if setArgNil {
			obj["sdns-options"] = nil
		} else {
			t, err := expandSystemFortiguardSdnsOptions(d, v, "sdns_options", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sdns-options"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		if setArgNil {
			obj["source-ip"] = nil
		} else {
			t, err := expandSystemFortiguardSourceIp(d, v, "source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok {
		if setArgNil {
			obj["source-ip6"] = nil
		} else {
			t, err := expandSystemFortiguardSourceIp6(d, v, "source_ip6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip6"] = t
			}
		}
	}

	if v, ok := d.GetOk("proxy_server_ip"); ok {
		if setArgNil {
			obj["proxy-server-ip"] = nil
		} else {
			t, err := expandSystemFortiguardProxyServerIp(d, v, "proxy_server_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["proxy-server-ip"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("proxy_server_port"); ok {
		if setArgNil {
			obj["proxy-server-port"] = nil
		} else {
			t, err := expandSystemFortiguardProxyServerPort(d, v, "proxy_server_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["proxy-server-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("proxy_username"); ok {
		if setArgNil {
			obj["proxy-username"] = nil
		} else {
			t, err := expandSystemFortiguardProxyUsername(d, v, "proxy_username", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["proxy-username"] = t
			}
		}
	}

	if v, ok := d.GetOk("proxy_password"); ok {
		if setArgNil {
			obj["proxy-password"] = nil
		} else {
			t, err := expandSystemFortiguardProxyPassword(d, v, "proxy_password", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["proxy-password"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("videofilter_license"); ok {
		if setArgNil {
			obj["videofilter-license"] = nil
		} else {
			t, err := expandSystemFortiguardVideofilterLicense(d, v, "videofilter_license", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["videofilter-license"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("videofilter_expiration"); ok {
		if setArgNil {
			obj["videofilter-expiration"] = nil
		} else {
			t, err := expandSystemFortiguardVideofilterExpiration(d, v, "videofilter_expiration", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["videofilter-expiration"] = t
			}
		}
	}

	if v, ok := d.GetOk("ddns_server_ip"); ok {
		if setArgNil {
			obj["ddns-server-ip"] = nil
		} else {
			t, err := expandSystemFortiguardDdnsServerIp(d, v, "ddns_server_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ddns-server-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("ddns_server_ip6"); ok {
		if setArgNil {
			obj["ddns-server-ip6"] = nil
		} else {
			t, err := expandSystemFortiguardDdnsServerIp6(d, v, "ddns_server_ip6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ddns-server-ip6"] = t
			}
		}
	}

	if v, ok := d.GetOk("ddns_server_port"); ok {
		if setArgNil {
			obj["ddns-server-port"] = nil
		} else {
			t, err := expandSystemFortiguardDdnsServerPort(d, v, "ddns_server_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ddns-server-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		if setArgNil {
			obj["interface-select-method"] = nil
		} else {
			t, err := expandSystemFortiguardInterfaceSelectMethod(d, v, "interface_select_method", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface-select-method"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		if setArgNil {
			obj["interface"] = nil
		} else {
			t, err := expandSystemFortiguardInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	}

	return &obj, nil
}
