// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure FortiGuard services.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			"ddns_server_ip": &schema.Schema{
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
		},
	}
}

func resourceSystemFortiguardUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemFortiguard(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemFortiguard resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemFortiguard(obj, mkey)
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

	err := c.DeleteSystemFortiguard(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFortiguard resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFortiguardRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemFortiguard(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemFortiguard resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFortiguard(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFortiguard resource from API: %v", err)
	}
	return nil
}

func flattenSystemFortiguardPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardServiceAccountId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardLoadBalanceServers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAutoJoinForticloud(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardUpdateServerLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardSandboxRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamForceOff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamCacheMpercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardAntispamTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionForceOff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionCacheMpercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardOutbreakPreventionTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterForceOff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterExpiration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardWebfilterTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardSdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardSdnsServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardDdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiguardDdnsServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemFortiguard(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("port", flattenSystemFortiguardPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("service_account_id", flattenSystemFortiguardServiceAccountId(o["service-account-id"], d, "service_account_id")); err != nil {
		if !fortiAPIPatch(o["service-account-id"]) {
			return fmt.Errorf("Error reading service_account_id: %v", err)
		}
	}

	if err = d.Set("load_balance_servers", flattenSystemFortiguardLoadBalanceServers(o["load-balance-servers"], d, "load_balance_servers")); err != nil {
		if !fortiAPIPatch(o["load-balance-servers"]) {
			return fmt.Errorf("Error reading load_balance_servers: %v", err)
		}
	}

	if err = d.Set("auto_join_forticloud", flattenSystemFortiguardAutoJoinForticloud(o["auto-join-forticloud"], d, "auto_join_forticloud")); err != nil {
		if !fortiAPIPatch(o["auto-join-forticloud"]) {
			return fmt.Errorf("Error reading auto_join_forticloud: %v", err)
		}
	}

	if err = d.Set("update_server_location", flattenSystemFortiguardUpdateServerLocation(o["update-server-location"], d, "update_server_location")); err != nil {
		if !fortiAPIPatch(o["update-server-location"]) {
			return fmt.Errorf("Error reading update_server_location: %v", err)
		}
	}

	if err = d.Set("sandbox_region", flattenSystemFortiguardSandboxRegion(o["sandbox-region"], d, "sandbox_region")); err != nil {
		if !fortiAPIPatch(o["sandbox-region"]) {
			return fmt.Errorf("Error reading sandbox_region: %v", err)
		}
	}

	if err = d.Set("antispam_force_off", flattenSystemFortiguardAntispamForceOff(o["antispam-force-off"], d, "antispam_force_off")); err != nil {
		if !fortiAPIPatch(o["antispam-force-off"]) {
			return fmt.Errorf("Error reading antispam_force_off: %v", err)
		}
	}

	if err = d.Set("antispam_cache", flattenSystemFortiguardAntispamCache(o["antispam-cache"], d, "antispam_cache")); err != nil {
		if !fortiAPIPatch(o["antispam-cache"]) {
			return fmt.Errorf("Error reading antispam_cache: %v", err)
		}
	}

	if err = d.Set("antispam_cache_ttl", flattenSystemFortiguardAntispamCacheTtl(o["antispam-cache-ttl"], d, "antispam_cache_ttl")); err != nil {
		if !fortiAPIPatch(o["antispam-cache-ttl"]) {
			return fmt.Errorf("Error reading antispam_cache_ttl: %v", err)
		}
	}

	if err = d.Set("antispam_cache_mpercent", flattenSystemFortiguardAntispamCacheMpercent(o["antispam-cache-mpercent"], d, "antispam_cache_mpercent")); err != nil {
		if !fortiAPIPatch(o["antispam-cache-mpercent"]) {
			return fmt.Errorf("Error reading antispam_cache_mpercent: %v", err)
		}
	}

	if err = d.Set("antispam_license", flattenSystemFortiguardAntispamLicense(o["antispam-license"], d, "antispam_license")); err != nil {
		if !fortiAPIPatch(o["antispam-license"]) {
			return fmt.Errorf("Error reading antispam_license: %v", err)
		}
	}

	if err = d.Set("antispam_expiration", flattenSystemFortiguardAntispamExpiration(o["antispam-expiration"], d, "antispam_expiration")); err != nil {
		if !fortiAPIPatch(o["antispam-expiration"]) {
			return fmt.Errorf("Error reading antispam_expiration: %v", err)
		}
	}

	if err = d.Set("antispam_timeout", flattenSystemFortiguardAntispamTimeout(o["antispam-timeout"], d, "antispam_timeout")); err != nil {
		if !fortiAPIPatch(o["antispam-timeout"]) {
			return fmt.Errorf("Error reading antispam_timeout: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_force_off", flattenSystemFortiguardOutbreakPreventionForceOff(o["outbreak-prevention-force-off"], d, "outbreak_prevention_force_off")); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-force-off"]) {
			return fmt.Errorf("Error reading outbreak_prevention_force_off: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_cache", flattenSystemFortiguardOutbreakPreventionCache(o["outbreak-prevention-cache"], d, "outbreak_prevention_cache")); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-cache"]) {
			return fmt.Errorf("Error reading outbreak_prevention_cache: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_cache_ttl", flattenSystemFortiguardOutbreakPreventionCacheTtl(o["outbreak-prevention-cache-ttl"], d, "outbreak_prevention_cache_ttl")); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-cache-ttl"]) {
			return fmt.Errorf("Error reading outbreak_prevention_cache_ttl: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_cache_mpercent", flattenSystemFortiguardOutbreakPreventionCacheMpercent(o["outbreak-prevention-cache-mpercent"], d, "outbreak_prevention_cache_mpercent")); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-cache-mpercent"]) {
			return fmt.Errorf("Error reading outbreak_prevention_cache_mpercent: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_license", flattenSystemFortiguardOutbreakPreventionLicense(o["outbreak-prevention-license"], d, "outbreak_prevention_license")); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-license"]) {
			return fmt.Errorf("Error reading outbreak_prevention_license: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_expiration", flattenSystemFortiguardOutbreakPreventionExpiration(o["outbreak-prevention-expiration"], d, "outbreak_prevention_expiration")); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-expiration"]) {
			return fmt.Errorf("Error reading outbreak_prevention_expiration: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention_timeout", flattenSystemFortiguardOutbreakPreventionTimeout(o["outbreak-prevention-timeout"], d, "outbreak_prevention_timeout")); err != nil {
		if !fortiAPIPatch(o["outbreak-prevention-timeout"]) {
			return fmt.Errorf("Error reading outbreak_prevention_timeout: %v", err)
		}
	}

	if err = d.Set("webfilter_force_off", flattenSystemFortiguardWebfilterForceOff(o["webfilter-force-off"], d, "webfilter_force_off")); err != nil {
		if !fortiAPIPatch(o["webfilter-force-off"]) {
			return fmt.Errorf("Error reading webfilter_force_off: %v", err)
		}
	}

	if err = d.Set("webfilter_cache", flattenSystemFortiguardWebfilterCache(o["webfilter-cache"], d, "webfilter_cache")); err != nil {
		if !fortiAPIPatch(o["webfilter-cache"]) {
			return fmt.Errorf("Error reading webfilter_cache: %v", err)
		}
	}

	if err = d.Set("webfilter_cache_ttl", flattenSystemFortiguardWebfilterCacheTtl(o["webfilter-cache-ttl"], d, "webfilter_cache_ttl")); err != nil {
		if !fortiAPIPatch(o["webfilter-cache-ttl"]) {
			return fmt.Errorf("Error reading webfilter_cache_ttl: %v", err)
		}
	}

	if err = d.Set("webfilter_license", flattenSystemFortiguardWebfilterLicense(o["webfilter-license"], d, "webfilter_license")); err != nil {
		if !fortiAPIPatch(o["webfilter-license"]) {
			return fmt.Errorf("Error reading webfilter_license: %v", err)
		}
	}

	if err = d.Set("webfilter_expiration", flattenSystemFortiguardWebfilterExpiration(o["webfilter-expiration"], d, "webfilter_expiration")); err != nil {
		if !fortiAPIPatch(o["webfilter-expiration"]) {
			return fmt.Errorf("Error reading webfilter_expiration: %v", err)
		}
	}

	if err = d.Set("webfilter_timeout", flattenSystemFortiguardWebfilterTimeout(o["webfilter-timeout"], d, "webfilter_timeout")); err != nil {
		if !fortiAPIPatch(o["webfilter-timeout"]) {
			return fmt.Errorf("Error reading webfilter_timeout: %v", err)
		}
	}

	if err = d.Set("sdns_server_ip", flattenSystemFortiguardSdnsServerIp(o["sdns-server-ip"], d, "sdns_server_ip")); err != nil {
		if !fortiAPIPatch(o["sdns-server-ip"]) {
			return fmt.Errorf("Error reading sdns_server_ip: %v", err)
		}
	}

	if err = d.Set("sdns_server_port", flattenSystemFortiguardSdnsServerPort(o["sdns-server-port"], d, "sdns_server_port")); err != nil {
		if !fortiAPIPatch(o["sdns-server-port"]) {
			return fmt.Errorf("Error reading sdns_server_port: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemFortiguardSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenSystemFortiguardSourceIp6(o["source-ip6"], d, "source_ip6")); err != nil {
		if !fortiAPIPatch(o["source-ip6"]) {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("ddns_server_ip", flattenSystemFortiguardDdnsServerIp(o["ddns-server-ip"], d, "ddns_server_ip")); err != nil {
		if !fortiAPIPatch(o["ddns-server-ip"]) {
			return fmt.Errorf("Error reading ddns_server_ip: %v", err)
		}
	}

	if err = d.Set("ddns_server_port", flattenSystemFortiguardDdnsServerPort(o["ddns-server-port"], d, "ddns_server_port")); err != nil {
		if !fortiAPIPatch(o["ddns-server-port"]) {
			return fmt.Errorf("Error reading ddns_server_port: %v", err)
		}
	}

	return nil
}

func flattenSystemFortiguardFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemFortiguardPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardServiceAccountId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardLoadBalanceServers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAutoJoinForticloud(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardUpdateServerLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSandboxRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamForceOff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamCacheMpercent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamLicense(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamExpiration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardAntispamTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionForceOff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionCacheMpercent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionLicense(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionExpiration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardOutbreakPreventionTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterForceOff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterLicense(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterExpiration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardWebfilterTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSdnsServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSdnsServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardSourceIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardDdnsServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiguardDdnsServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFortiguard(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("port"); ok {
		t, err := expandSystemFortiguardPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("service_account_id"); ok {
		t, err := expandSystemFortiguardServiceAccountId(d, v, "service_account_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-account-id"] = t
		}
	}

	if v, ok := d.GetOk("load_balance_servers"); ok {
		t, err := expandSystemFortiguardLoadBalanceServers(d, v, "load_balance_servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balance-servers"] = t
		}
	}

	if v, ok := d.GetOk("auto_join_forticloud"); ok {
		t, err := expandSystemFortiguardAutoJoinForticloud(d, v, "auto_join_forticloud")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-join-forticloud"] = t
		}
	}

	if v, ok := d.GetOk("update_server_location"); ok {
		t, err := expandSystemFortiguardUpdateServerLocation(d, v, "update_server_location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-server-location"] = t
		}
	}

	if v, ok := d.GetOk("sandbox_region"); ok {
		t, err := expandSystemFortiguardSandboxRegion(d, v, "sandbox_region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sandbox-region"] = t
		}
	}

	if v, ok := d.GetOk("antispam_force_off"); ok {
		t, err := expandSystemFortiguardAntispamForceOff(d, v, "antispam_force_off")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-force-off"] = t
		}
	}

	if v, ok := d.GetOk("antispam_cache"); ok {
		t, err := expandSystemFortiguardAntispamCache(d, v, "antispam_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-cache"] = t
		}
	}

	if v, ok := d.GetOk("antispam_cache_ttl"); ok {
		t, err := expandSystemFortiguardAntispamCacheTtl(d, v, "antispam_cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("antispam_cache_mpercent"); ok {
		t, err := expandSystemFortiguardAntispamCacheMpercent(d, v, "antispam_cache_mpercent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-cache-mpercent"] = t
		}
	}

	if v, ok := d.GetOkExists("antispam_license"); ok {
		t, err := expandSystemFortiguardAntispamLicense(d, v, "antispam_license")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-license"] = t
		}
	}

	if v, ok := d.GetOkExists("antispam_expiration"); ok {
		t, err := expandSystemFortiguardAntispamExpiration(d, v, "antispam_expiration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-expiration"] = t
		}
	}

	if v, ok := d.GetOk("antispam_timeout"); ok {
		t, err := expandSystemFortiguardAntispamTimeout(d, v, "antispam_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antispam-timeout"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_force_off"); ok {
		t, err := expandSystemFortiguardOutbreakPreventionForceOff(d, v, "outbreak_prevention_force_off")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-force-off"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_cache"); ok {
		t, err := expandSystemFortiguardOutbreakPreventionCache(d, v, "outbreak_prevention_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-cache"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_cache_ttl"); ok {
		t, err := expandSystemFortiguardOutbreakPreventionCacheTtl(d, v, "outbreak_prevention_cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_cache_mpercent"); ok {
		t, err := expandSystemFortiguardOutbreakPreventionCacheMpercent(d, v, "outbreak_prevention_cache_mpercent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-cache-mpercent"] = t
		}
	}

	if v, ok := d.GetOkExists("outbreak_prevention_license"); ok {
		t, err := expandSystemFortiguardOutbreakPreventionLicense(d, v, "outbreak_prevention_license")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-license"] = t
		}
	}

	if v, ok := d.GetOkExists("outbreak_prevention_expiration"); ok {
		t, err := expandSystemFortiguardOutbreakPreventionExpiration(d, v, "outbreak_prevention_expiration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-expiration"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_timeout"); ok {
		t, err := expandSystemFortiguardOutbreakPreventionTimeout(d, v, "outbreak_prevention_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-timeout"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_force_off"); ok {
		t, err := expandSystemFortiguardWebfilterForceOff(d, v, "webfilter_force_off")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-force-off"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_cache"); ok {
		t, err := expandSystemFortiguardWebfilterCache(d, v, "webfilter_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-cache"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_cache_ttl"); ok {
		t, err := expandSystemFortiguardWebfilterCacheTtl(d, v, "webfilter_cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-cache-ttl"] = t
		}
	}

	if v, ok := d.GetOkExists("webfilter_license"); ok {
		t, err := expandSystemFortiguardWebfilterLicense(d, v, "webfilter_license")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-license"] = t
		}
	}

	if v, ok := d.GetOkExists("webfilter_expiration"); ok {
		t, err := expandSystemFortiguardWebfilterExpiration(d, v, "webfilter_expiration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-expiration"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_timeout"); ok {
		t, err := expandSystemFortiguardWebfilterTimeout(d, v, "webfilter_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-timeout"] = t
		}
	}

	if v, ok := d.GetOk("sdns_server_ip"); ok {
		t, err := expandSystemFortiguardSdnsServerIp(d, v, "sdns_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdns-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("sdns_server_port"); ok {
		t, err := expandSystemFortiguardSdnsServerPort(d, v, "sdns_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdns-server-port"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandSystemFortiguardSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok {
		t, err := expandSystemFortiguardSourceIp6(d, v, "source_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server_ip"); ok {
		t, err := expandSystemFortiguardDdnsServerIp(d, v, "ddns_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server_port"); ok {
		t, err := expandSystemFortiguardDdnsServerPort(d, v, "ddns_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server-port"] = t
		}
	}

	return &obj, nil
}
