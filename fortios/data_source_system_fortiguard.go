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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemFortiguard() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemFortiguardRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_account_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"load_balance_servers": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"auto_join_forticloud": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_server_location": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sandbox_region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_ffdb": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_uwdb": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_extdb": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_build_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"persistent_connection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fortiguard_anycast": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fortiguard_anycast_source": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"antispam_force_off": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"antispam_cache": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"antispam_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"antispam_cache_mpercent": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"antispam_license": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"antispam_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"antispam_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"outbreak_prevention_force_off": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"outbreak_prevention_cache": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"outbreak_prevention_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"outbreak_prevention_cache_mpercent": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"outbreak_prevention_license": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"outbreak_prevention_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"outbreak_prevention_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"webfilter_force_off": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"webfilter_cache": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"webfilter_cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"webfilter_license": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"webfilter_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"webfilter_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"sdns_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sdns_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"anycast_sdns_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"anycast_sdns_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"sdns_options": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxy_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxy_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"proxy_username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxy_password": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"videofilter_license": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"videofilter_expiration": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ddns_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ddns_server_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ddns_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemFortiguardRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "SystemFortiguard"

	o, err := c.ReadSystemFortiguard(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemFortiguard: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemFortiguard(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemFortiguard from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemFortiguardProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardServiceAccountId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardLoadBalanceServers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAutoJoinForticloud(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardUpdateServerLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardSandboxRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardUpdateFfdb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardUpdateUwdb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardUpdateExtdb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardUpdateBuildProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardPersistentConnection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardFortiguardAnycast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardFortiguardAnycastSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAntispamForceOff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAntispamCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAntispamCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAntispamCacheMpercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAntispamLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAntispamExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAntispamTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardOutbreakPreventionForceOff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardOutbreakPreventionCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardOutbreakPreventionCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardOutbreakPreventionCacheMpercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardOutbreakPreventionLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardOutbreakPreventionExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardOutbreakPreventionTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardWebfilterForceOff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardWebfilterCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardWebfilterCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardWebfilterLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardWebfilterExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardWebfilterTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardSdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardSdnsServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAnycastSdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardAnycastSdnsServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardSdnsOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardProxyServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardProxyServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardProxyUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardProxyPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardVideofilterLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardVideofilterExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardDdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardDdnsServerIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardDdnsServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFortiguardInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemFortiguard(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("protocol", dataSourceFlattenSystemFortiguardProtocol(o["protocol"], d, "protocol")); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("port", dataSourceFlattenSystemFortiguardPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("service_account_id", dataSourceFlattenSystemFortiguardServiceAccountId(o["service-account-id"], d, "service_account_id")); err != nil {
		if !fortiAPIPatch(o["service-account-id"]) {
			return fmt.Errorf("Error reading service_account_id: %v", err)
		}
	}

	if err = d.Set("load_balance_servers", dataSourceFlattenSystemFortiguardLoadBalanceServers(o["load-balance-servers"], d, "load_balance_servers")); err != nil {
		if !fortiAPIPatch(o["load-balance-servers"]) {
			return fmt.Errorf("Error reading load_balance_servers: %v", err)
		}
	}

	if err = d.Set("auto_join_forticloud", dataSourceFlattenSystemFortiguardAutoJoinForticloud(o["auto-join-forticloud"], d, "auto_join_forticloud")); err != nil {
		if !fortiAPIPatch(o["auto-join-forticloud"]) {
			return fmt.Errorf("Error reading auto_join_forticloud: %v", err)
		}
	}

	if err = d.Set("update_server_location", dataSourceFlattenSystemFortiguardUpdateServerLocation(o["update-server-location"], d, "update_server_location")); err != nil {
		if !fortiAPIPatch(o["update-server-location"]) {
			return fmt.Errorf("Error reading update_server_location: %v", err)
		}
	}

	if err = d.Set("sandbox_region", dataSourceFlattenSystemFortiguardSandboxRegion(o["sandbox-region"], d, "sandbox_region")); err != nil {
		if !fortiAPIPatch(o["sandbox-region"]) {
			return fmt.Errorf("Error reading sandbox_region: %v", err)
		}
	}

	if err = d.Set("update_ffdb", dataSourceFlattenSystemFortiguardUpdateFfdb(o["update-ffdb"], d, "update_ffdb")); err != nil {
		if !fortiAPIPatch(o["update-ffdb"]) {
			return fmt.Errorf("Error reading update_ffdb: %v", err)
		}
	}

	if err = d.Set("update_uwdb", dataSourceFlattenSystemFortiguardUpdateUwdb(o["update-uwdb"], d, "update_uwdb")); err != nil {
		if !fortiAPIPatch(o["update-uwdb"]) {
			return fmt.Errorf("Error reading update_uwdb: %v", err)
		}
	}

	if err = d.Set("update_extdb", dataSourceFlattenSystemFortiguardUpdateExtdb(o["update-extdb"], d, "update_extdb")); err != nil {
		if !fortiAPIPatch(o["update-extdb"]) {
			return fmt.Errorf("Error reading update_extdb: %v", err)
		}
	}

	if err = d.Set("update_build_proxy", dataSourceFlattenSystemFortiguardUpdateBuildProxy(o["update-build-proxy"], d, "update_build_proxy")); err != nil {
		if !fortiAPIPatch(o["update-build-proxy"]) {
			return fmt.Errorf("Error reading update_build_proxy: %v", err)
		}
	}

	if err = d.Set("persistent_connection", dataSourceFlattenSystemFortiguardPersistentConnection(o["persistent-connection"], d, "persistent_connection")); err != nil {
		if !fortiAPIPatch(o["persistent-connection"]) {
			return fmt.Errorf("Error reading persistent_connection: %v", err)
		}
	}

	if err = d.Set("vdom", dataSourceFlattenSystemFortiguardVdom(o["vdom"], d, "vdom")); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("fortiguard_anycast", dataSourceFlattenSystemFortiguardFortiguardAnycast(o["fortiguard-anycast"], d, "fortiguard_anycast")); err != nil {
		if !fortiAPIPatch(o["fortiguard-anycast"]) {
			return fmt.Errorf("Error reading fortiguard_anycast: %v", err)
		}
	}

	if err = d.Set("fortiguard_anycast_source", dataSourceFlattenSystemFortiguardFortiguardAnycastSource(o["fortiguard-anycast-source"], d, "fortiguard_anycast_source")); err != nil {
		if !fortiAPIPatch(o["fortiguard-anycast-source"]) {
			return fmt.Errorf("Error reading fortiguard_anycast_source: %v", err)
		}
	}

	if err = d.Set("antispam_force_off", dataSourceFlattenSystemFortiguardAntispamForceOff(o["antispam-force-off"], d, "antispam_force_off")); err != nil {
		if !fortiAPIPatch(o["antispam-force-off"]) {
			return fmt.Errorf("Error reading antispam_force_off: %v", err)
		}
	}

	if err = d.Set("antispam_cache", dataSourceFlattenSystemFortiguardAntispamCache(o["antispam-cache"], d, "antispam_cache")); err != nil {
		if !fortiAPIPatch(o["antispam-cache"]) {
			return fmt.Errorf("Error reading antispam_cache: %v", err)
		}
	}

	if err = d.Set("antispam_cache_ttl", dataSourceFlattenSystemFortiguardAntispamCacheTtl(o["antispam-cache-ttl"], d, "antispam_cache_ttl")); err != nil {
		if !fortiAPIPatch(o["antispam-cache-ttl"]) {
			return fmt.Errorf("Error reading antispam_cache_ttl: %v", err)
		}
	}

	if err = d.Set("antispam_cache_mpercent", dataSourceFlattenSystemFortiguardAntispamCacheMpercent(o["antispam-cache-mpercent"], d, "antispam_cache_mpercent")); err != nil {
		if !fortiAPIPatch(o["antispam-cache-mpercent"]) {
			return fmt.Errorf("Error reading antispam_cache_mpercent: %v", err)
		}
	}

	if err = d.Set("antispam_license", dataSourceFlattenSystemFortiguardAntispamLicense(o["antispam-license"], d, "antispam_license")); err != nil {
		if !fortiAPIPatch(o["antispam-license"]) {
			return fmt.Errorf("Error reading antispam_license: %v", err)
		}
	}

	if err = d.Set("antispam_expiration", dataSourceFlattenSystemFortiguardAntispamExpiration(o["antispam-expiration"], d, "antispam_expiration")); err != nil {
		if !fortiAPIPatch(o["antispam-expiration"]) {
			return fmt.Errorf("Error reading antispam_expiration: %v", err)
		}
	}

	if err = d.Set("antispam_timeout", dataSourceFlattenSystemFortiguardAntispamTimeout(o["antispam-timeout"], d, "antispam_timeout")); err != nil {
		if !fortiAPIPatch(o["antispam-timeout"]) {
			return fmt.Errorf("Error reading antispam_timeout: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_force_off", dataSourceFlattenSystemFortiguardOutbreakPreventionForceOff(o["outbreak-prevention-force-off"], d, "outbreak_prevention_force_off")); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-force-off"]) {
			return fmt.Errorf("Error reading outbreak_prevention_force_off: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_cache", dataSourceFlattenSystemFortiguardOutbreakPreventionCache(o["outbreak-prevention-cache"], d, "outbreak_prevention_cache")); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-cache"]) {
			return fmt.Errorf("Error reading outbreak_prevention_cache: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_cache_ttl", dataSourceFlattenSystemFortiguardOutbreakPreventionCacheTtl(o["outbreak-prevention-cache-ttl"], d, "outbreak_prevention_cache_ttl")); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-cache-ttl"]) {
			return fmt.Errorf("Error reading outbreak_prevention_cache_ttl: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_cache_mpercent", dataSourceFlattenSystemFortiguardOutbreakPreventionCacheMpercent(o["outbreak-prevention-cache-mpercent"], d, "outbreak_prevention_cache_mpercent")); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-cache-mpercent"]) {
			return fmt.Errorf("Error reading outbreak_prevention_cache_mpercent: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_license", dataSourceFlattenSystemFortiguardOutbreakPreventionLicense(o["outbreak-prevention-license"], d, "outbreak_prevention_license")); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-license"]) {
			return fmt.Errorf("Error reading outbreak_prevention_license: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_expiration", dataSourceFlattenSystemFortiguardOutbreakPreventionExpiration(o["outbreak-prevention-expiration"], d, "outbreak_prevention_expiration")); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-expiration"]) {
			return fmt.Errorf("Error reading outbreak_prevention_expiration: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_timeout", dataSourceFlattenSystemFortiguardOutbreakPreventionTimeout(o["outbreak-prevention-timeout"], d, "outbreak_prevention_timeout")); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-timeout"]) {
			return fmt.Errorf("Error reading outbreak_prevention_timeout: %v", err)
		}
	}

	if err = d.Set("webfilter_force_off", dataSourceFlattenSystemFortiguardWebfilterForceOff(o["webfilter-force-off"], d, "webfilter_force_off")); err != nil {
		if !fortiAPIPatch(o["webfilter-force-off"]) {
			return fmt.Errorf("Error reading webfilter_force_off: %v", err)
		}
	}

	if err = d.Set("webfilter_cache", dataSourceFlattenSystemFortiguardWebfilterCache(o["webfilter-cache"], d, "webfilter_cache")); err != nil {
		if !fortiAPIPatch(o["webfilter-cache"]) {
			return fmt.Errorf("Error reading webfilter_cache: %v", err)
		}
	}

	if err = d.Set("webfilter_cache_ttl", dataSourceFlattenSystemFortiguardWebfilterCacheTtl(o["webfilter-cache-ttl"], d, "webfilter_cache_ttl")); err != nil {
		if !fortiAPIPatch(o["webfilter-cache-ttl"]) {
			return fmt.Errorf("Error reading webfilter_cache_ttl: %v", err)
		}
	}

	if err = d.Set("webfilter_license", dataSourceFlattenSystemFortiguardWebfilterLicense(o["webfilter-license"], d, "webfilter_license")); err != nil {
		if !fortiAPIPatch(o["webfilter-license"]) {
			return fmt.Errorf("Error reading webfilter_license: %v", err)
		}
	}

	if err = d.Set("webfilter_expiration", dataSourceFlattenSystemFortiguardWebfilterExpiration(o["webfilter-expiration"], d, "webfilter_expiration")); err != nil {
		if !fortiAPIPatch(o["webfilter-expiration"]) {
			return fmt.Errorf("Error reading webfilter_expiration: %v", err)
		}
	}

	if err = d.Set("webfilter_timeout", dataSourceFlattenSystemFortiguardWebfilterTimeout(o["webfilter-timeout"], d, "webfilter_timeout")); err != nil {
		if !fortiAPIPatch(o["webfilter-timeout"]) {
			return fmt.Errorf("Error reading webfilter_timeout: %v", err)
		}
	}

	if err = d.Set("sdns_server_ip", dataSourceFlattenSystemFortiguardSdnsServerIp(o["sdns-server-ip"], d, "sdns_server_ip")); err != nil {
		if !fortiAPIPatch(o["sdns-server-ip"]) {
			return fmt.Errorf("Error reading sdns_server_ip: %v", err)
		}
	}

	if err = d.Set("sdns_server_port", dataSourceFlattenSystemFortiguardSdnsServerPort(o["sdns-server-port"], d, "sdns_server_port")); err != nil {
		if !fortiAPIPatch(o["sdns-server-port"]) {
			return fmt.Errorf("Error reading sdns_server_port: %v", err)
		}
	}

	if err = d.Set("anycast_sdns_server_ip", dataSourceFlattenSystemFortiguardAnycastSdnsServerIp(o["anycast-sdns-server-ip"], d, "anycast_sdns_server_ip")); err != nil {
		if !fortiAPIPatch(o["anycast-sdns-server-ip"]) {
			return fmt.Errorf("Error reading anycast_sdns_server_ip: %v", err)
		}
	}

	if err = d.Set("anycast_sdns_server_port", dataSourceFlattenSystemFortiguardAnycastSdnsServerPort(o["anycast-sdns-server-port"], d, "anycast_sdns_server_port")); err != nil {
		if !fortiAPIPatch(o["anycast-sdns-server-port"]) {
			return fmt.Errorf("Error reading anycast_sdns_server_port: %v", err)
		}
	}

	if err = d.Set("sdns_options", dataSourceFlattenSystemFortiguardSdnsOptions(o["sdns-options"], d, "sdns_options")); err != nil {
		if !fortiAPIPatch(o["sdns-options"]) {
			return fmt.Errorf("Error reading sdns_options: %v", err)
		}
	}

	if err = d.Set("source_ip", dataSourceFlattenSystemFortiguardSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip6", dataSourceFlattenSystemFortiguardSourceIp6(o["source-ip6"], d, "source_ip6")); err != nil {
		if !fortiAPIPatch(o["source-ip6"]) {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("proxy_server_ip", dataSourceFlattenSystemFortiguardProxyServerIp(o["proxy-server-ip"], d, "proxy_server_ip")); err != nil {
		if !fortiAPIPatch(o["proxy-server-ip"]) {
			return fmt.Errorf("Error reading proxy_server_ip: %v", err)
		}
	}

	if err = d.Set("proxy_server_port", dataSourceFlattenSystemFortiguardProxyServerPort(o["proxy-server-port"], d, "proxy_server_port")); err != nil {
		if !fortiAPIPatch(o["proxy-server-port"]) {
			return fmt.Errorf("Error reading proxy_server_port: %v", err)
		}
	}

	if err = d.Set("proxy_username", dataSourceFlattenSystemFortiguardProxyUsername(o["proxy-username"], d, "proxy_username")); err != nil {
		if !fortiAPIPatch(o["proxy-username"]) {
			return fmt.Errorf("Error reading proxy_username: %v", err)
		}
	}

	if err = d.Set("videofilter_license", dataSourceFlattenSystemFortiguardVideofilterLicense(o["videofilter-license"], d, "videofilter_license")); err != nil {
		if !fortiAPIPatch(o["videofilter-license"]) {
			return fmt.Errorf("Error reading videofilter_license: %v", err)
		}
	}

	if err = d.Set("videofilter_expiration", dataSourceFlattenSystemFortiguardVideofilterExpiration(o["videofilter-expiration"], d, "videofilter_expiration")); err != nil {
		if !fortiAPIPatch(o["videofilter-expiration"]) {
			return fmt.Errorf("Error reading videofilter_expiration: %v", err)
		}
	}

	if err = d.Set("ddns_server_ip", dataSourceFlattenSystemFortiguardDdnsServerIp(o["ddns-server-ip"], d, "ddns_server_ip")); err != nil {
		if !fortiAPIPatch(o["ddns-server-ip"]) {
			return fmt.Errorf("Error reading ddns_server_ip: %v", err)
		}
	}

	if err = d.Set("ddns_server_ip6", dataSourceFlattenSystemFortiguardDdnsServerIp6(o["ddns-server-ip6"], d, "ddns_server_ip6")); err != nil {
		if !fortiAPIPatch(o["ddns-server-ip6"]) {
			return fmt.Errorf("Error reading ddns_server_ip6: %v", err)
		}
	}

	if err = d.Set("ddns_server_port", dataSourceFlattenSystemFortiguardDdnsServerPort(o["ddns-server-port"], d, "ddns_server_port")); err != nil {
		if !fortiAPIPatch(o["ddns-server-port"]) {
			return fmt.Errorf("Error reading ddns_server_port: %v", err)
		}
	}

	if err = d.Set("interface_select_method", dataSourceFlattenSystemFortiguardInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemFortiguardInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemFortiguardFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
