// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Global FortiAnalyzer settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogFortianalyzerSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogFortianalyzerSettingUpdate,
		Read:   resourceLogFortianalyzerSettingRead,
		Update: resourceLogFortianalyzerSettingUpdate,
		Delete: resourceLogFortianalyzerSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_archive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"alt_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"fallback_to_primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certificate_verification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"server_cert_ca": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"preshared_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"access_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hmac_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enc_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conn_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),
				Optional:     true,
				Computed:     true,
			},
			"monitor_keepalive_period": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),
				Optional:     true,
				Computed:     true,
			},
			"monitor_failure_retry_period": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 86400),
				Optional:     true,
				Computed:     true,
			},
			"mgmt_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"faz_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"certificate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"source_ip": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"__change_ip": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
			},
			"upload_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upload_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reliable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_log_rate": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100000),
				Optional:     true,
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
			},
			"vrf_select": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 511),
				Optional:     true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceLogFortianalyzerSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectLogFortianalyzerSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzerSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogFortianalyzerSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzerSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogFortianalyzerSetting")
	}

	return resourceLogFortianalyzerSettingRead(d, m)
}

func resourceLogFortianalyzerSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogFortianalyzerSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzerSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateLogFortianalyzerSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing LogFortianalyzerSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogFortianalyzerSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadLogFortianalyzerSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzerSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogFortianalyzerSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzerSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogFortianalyzerSettingStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingIpsArchive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingAltServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingFallbackToPrimary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingCertificateVerification(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingSerial(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenLogFortianalyzerSettingSerialName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenLogFortianalyzerSettingSerialName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingServerCertCa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingPresharedKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingAccessConfig(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingHmacAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingEncAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingConnTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenLogFortianalyzerSettingMonitorKeepalivePeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenLogFortianalyzerSettingMonitorFailureRetryPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenLogFortianalyzerSettingMgmtName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingFazType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenLogFortianalyzerSettingCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingChangeIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenLogFortianalyzerSettingUploadOption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingUploadInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingUploadDay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingUploadTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingReliable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingMaxLogRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenLogFortianalyzerSettingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerSettingVrfSelect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectLogFortianalyzerSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("status", flattenLogFortianalyzerSettingStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("ips_archive", flattenLogFortianalyzerSettingIpsArchive(o["ips-archive"], d, "ips_archive", sv)); err != nil {
		if !fortiAPIPatch(o["ips-archive"]) {
			return fmt.Errorf("Error reading ips_archive: %v", err)
		}
	}

	if err = d.Set("server", flattenLogFortianalyzerSettingServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("alt_server", flattenLogFortianalyzerSettingAltServer(o["alt-server"], d, "alt_server", sv)); err != nil {
		if !fortiAPIPatch(o["alt-server"]) {
			return fmt.Errorf("Error reading alt_server: %v", err)
		}
	}

	if err = d.Set("fallback_to_primary", flattenLogFortianalyzerSettingFallbackToPrimary(o["fallback-to-primary"], d, "fallback_to_primary", sv)); err != nil {
		if !fortiAPIPatch(o["fallback-to-primary"]) {
			return fmt.Errorf("Error reading fallback_to_primary: %v", err)
		}
	}

	if err = d.Set("certificate_verification", flattenLogFortianalyzerSettingCertificateVerification(o["certificate-verification"], d, "certificate_verification", sv)); err != nil {
		if !fortiAPIPatch(o["certificate-verification"]) {
			return fmt.Errorf("Error reading certificate_verification: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("serial", flattenLogFortianalyzerSettingSerial(o["serial"], d, "serial", sv)); err != nil {
			if !fortiAPIPatch(o["serial"]) {
				return fmt.Errorf("Error reading serial: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("serial"); ok {
			if err = d.Set("serial", flattenLogFortianalyzerSettingSerial(o["serial"], d, "serial", sv)); err != nil {
				if !fortiAPIPatch(o["serial"]) {
					return fmt.Errorf("Error reading serial: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_cert_ca", flattenLogFortianalyzerSettingServerCertCa(o["server-cert-ca"], d, "server_cert_ca", sv)); err != nil {
		if !fortiAPIPatch(o["server-cert-ca"]) {
			return fmt.Errorf("Error reading server_cert_ca: %v", err)
		}
	}

	if err = d.Set("preshared_key", flattenLogFortianalyzerSettingPresharedKey(o["preshared-key"], d, "preshared_key", sv)); err != nil {
		if !fortiAPIPatch(o["preshared-key"]) {
			return fmt.Errorf("Error reading preshared_key: %v", err)
		}
	}

	if err = d.Set("access_config", flattenLogFortianalyzerSettingAccessConfig(o["access-config"], d, "access_config", sv)); err != nil {
		if !fortiAPIPatch(o["access-config"]) {
			return fmt.Errorf("Error reading access_config: %v", err)
		}
	}

	if err = d.Set("hmac_algorithm", flattenLogFortianalyzerSettingHmacAlgorithm(o["hmac-algorithm"], d, "hmac_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["hmac-algorithm"]) {
			return fmt.Errorf("Error reading hmac_algorithm: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenLogFortianalyzerSettingEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["enc-algorithm"]) {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenLogFortianalyzerSettingSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-version"]) {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("conn_timeout", flattenLogFortianalyzerSettingConnTimeout(o["conn-timeout"], d, "conn_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["conn-timeout"]) {
			return fmt.Errorf("Error reading conn_timeout: %v", err)
		}
	}

	if err = d.Set("monitor_keepalive_period", flattenLogFortianalyzerSettingMonitorKeepalivePeriod(o["monitor-keepalive-period"], d, "monitor_keepalive_period", sv)); err != nil {
		if !fortiAPIPatch(o["monitor-keepalive-period"]) {
			return fmt.Errorf("Error reading monitor_keepalive_period: %v", err)
		}
	}

	if err = d.Set("monitor_failure_retry_period", flattenLogFortianalyzerSettingMonitorFailureRetryPeriod(o["monitor-failure-retry-period"], d, "monitor_failure_retry_period", sv)); err != nil {
		if !fortiAPIPatch(o["monitor-failure-retry-period"]) {
			return fmt.Errorf("Error reading monitor_failure_retry_period: %v", err)
		}
	}

	if err = d.Set("mgmt_name", flattenLogFortianalyzerSettingMgmtName(o["mgmt-name"], d, "mgmt_name", sv)); err != nil {
		if !fortiAPIPatch(o["mgmt-name"]) {
			return fmt.Errorf("Error reading mgmt_name: %v", err)
		}
	}

	if err = d.Set("faz_type", flattenLogFortianalyzerSettingFazType(o["faz-type"], d, "faz_type", sv)); err != nil {
		if !fortiAPIPatch(o["faz-type"]) {
			return fmt.Errorf("Error reading faz_type: %v", err)
		}
	}

	if err = d.Set("certificate", flattenLogFortianalyzerSettingCertificate(o["certificate"], d, "certificate", sv)); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenLogFortianalyzerSettingSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("__change_ip", flattenLogFortianalyzerSettingChangeIp(o["__change_ip"], d, "__change_ip", sv)); err != nil {
		if !fortiAPIPatch(o["__change_ip"]) {
			return fmt.Errorf("Error reading __change_ip: %v", err)
		}
	}

	if err = d.Set("upload_option", flattenLogFortianalyzerSettingUploadOption(o["upload-option"], d, "upload_option", sv)); err != nil {
		if !fortiAPIPatch(o["upload-option"]) {
			return fmt.Errorf("Error reading upload_option: %v", err)
		}
	}

	if err = d.Set("upload_interval", flattenLogFortianalyzerSettingUploadInterval(o["upload-interval"], d, "upload_interval", sv)); err != nil {
		if !fortiAPIPatch(o["upload-interval"]) {
			return fmt.Errorf("Error reading upload_interval: %v", err)
		}
	}

	if err = d.Set("upload_day", flattenLogFortianalyzerSettingUploadDay(o["upload-day"], d, "upload_day", sv)); err != nil {
		if !fortiAPIPatch(o["upload-day"]) {
			return fmt.Errorf("Error reading upload_day: %v", err)
		}
	}

	if err = d.Set("upload_time", flattenLogFortianalyzerSettingUploadTime(o["upload-time"], d, "upload_time", sv)); err != nil {
		if !fortiAPIPatch(o["upload-time"]) {
			return fmt.Errorf("Error reading upload_time: %v", err)
		}
	}

	if err = d.Set("reliable", flattenLogFortianalyzerSettingReliable(o["reliable"], d, "reliable", sv)); err != nil {
		if !fortiAPIPatch(o["reliable"]) {
			return fmt.Errorf("Error reading reliable: %v", err)
		}
	}

	if err = d.Set("priority", flattenLogFortianalyzerSettingPriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("max_log_rate", flattenLogFortianalyzerSettingMaxLogRate(o["max-log-rate"], d, "max_log_rate", sv)); err != nil {
		if !fortiAPIPatch(o["max-log-rate"]) {
			return fmt.Errorf("Error reading max_log_rate: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenLogFortianalyzerSettingInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenLogFortianalyzerSettingInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenLogFortianalyzerSettingVrfSelect(o["vrf-select"], d, "vrf_select", sv)); err != nil {
		if !fortiAPIPatch(o["vrf-select"]) {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func flattenLogFortianalyzerSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogFortianalyzerSettingStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingIpsArchive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingAltServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingFallbackToPrimary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingCertificateVerification(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingSerial(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandLogFortianalyzerSettingSerialName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLogFortianalyzerSettingSerialName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingServerCertCa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingPresharedKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingAccessConfig(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingHmacAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingEncAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingConnTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingMonitorKeepalivePeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingMonitorFailureRetryPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingMgmtName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingFazType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingChangeIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingUploadOption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingUploadInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingUploadDay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingUploadTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingReliable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingMaxLogRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerSettingVrfSelect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogFortianalyzerSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("ips_archive"); ok {
		if setArgNil {
			obj["ips-archive"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingIpsArchive(d, v, "ips_archive", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ips-archive"] = t
			}
		}
	}

	if v, ok := d.GetOk("server"); ok {
		if setArgNil {
			obj["server"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingServer(d, v, "server", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server"] = t
			}
		}
	} else if d.HasChange("server") {
		obj["server"] = nil
	}

	if v, ok := d.GetOk("alt_server"); ok {
		if setArgNil {
			obj["alt-server"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingAltServer(d, v, "alt_server", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["alt-server"] = t
			}
		}
	} else if d.HasChange("alt_server") {
		obj["alt-server"] = nil
	}

	if v, ok := d.GetOk("fallback_to_primary"); ok {
		if setArgNil {
			obj["fallback-to-primary"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingFallbackToPrimary(d, v, "fallback_to_primary", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fallback-to-primary"] = t
			}
		}
	}

	if v, ok := d.GetOk("certificate_verification"); ok {
		if setArgNil {
			obj["certificate-verification"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingCertificateVerification(d, v, "certificate_verification", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["certificate-verification"] = t
			}
		}
	}

	if v, ok := d.GetOk("serial"); ok || d.HasChange("serial") {
		if setArgNil {
			obj["serial"] = make([]struct{}, 0)
		} else {
			t, err := expandLogFortianalyzerSettingSerial(d, v, "serial", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["serial"] = t
			}
		}
	}

	if v, ok := d.GetOk("server_cert_ca"); ok {
		if setArgNil {
			obj["server-cert-ca"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingServerCertCa(d, v, "server_cert_ca", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server-cert-ca"] = t
			}
		}
	} else if d.HasChange("server_cert_ca") {
		obj["server-cert-ca"] = nil
	}

	if v, ok := d.GetOk("preshared_key"); ok {
		if setArgNil {
			obj["preshared-key"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingPresharedKey(d, v, "preshared_key", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["preshared-key"] = t
			}
		}
	} else if d.HasChange("preshared_key") {
		obj["preshared-key"] = nil
	}

	if v, ok := d.GetOk("access_config"); ok {
		if setArgNil {
			obj["access-config"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingAccessConfig(d, v, "access_config", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["access-config"] = t
			}
		}
	}

	if v, ok := d.GetOk("hmac_algorithm"); ok {
		if setArgNil {
			obj["hmac-algorithm"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingHmacAlgorithm(d, v, "hmac_algorithm", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hmac-algorithm"] = t
			}
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok {
		if setArgNil {
			obj["enc-algorithm"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingEncAlgorithm(d, v, "enc_algorithm", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["enc-algorithm"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok {
		if setArgNil {
			obj["ssl-min-proto-version"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingSslMinProtoVersion(d, v, "ssl_min_proto_version", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssl-min-proto-version"] = t
			}
		}
	}

	if v, ok := d.GetOk("conn_timeout"); ok {
		if setArgNil {
			obj["conn-timeout"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingConnTimeout(d, v, "conn_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["conn-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("monitor_keepalive_period"); ok {
		if setArgNil {
			obj["monitor-keepalive-period"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingMonitorKeepalivePeriod(d, v, "monitor_keepalive_period", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["monitor-keepalive-period"] = t
			}
		}
	}

	if v, ok := d.GetOk("monitor_failure_retry_period"); ok {
		if setArgNil {
			obj["monitor-failure-retry-period"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingMonitorFailureRetryPeriod(d, v, "monitor_failure_retry_period", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["monitor-failure-retry-period"] = t
			}
		}
	}

	if v, ok := d.GetOk("mgmt_name"); ok {
		if setArgNil {
			obj["mgmt-name"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingMgmtName(d, v, "mgmt_name", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mgmt-name"] = t
			}
		}
	} else if d.HasChange("mgmt_name") {
		obj["mgmt-name"] = nil
	}

	if v, ok := d.GetOkExists("faz_type"); ok {
		if setArgNil {
			obj["faz-type"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingFazType(d, v, "faz_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["faz-type"] = t
			}
		}
	} else if d.HasChange("faz_type") {
		obj["faz-type"] = nil
	}

	if v, ok := d.GetOk("certificate"); ok {
		if setArgNil {
			obj["certificate"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingCertificate(d, v, "certificate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["certificate"] = t
			}
		}
	} else if d.HasChange("certificate") {
		obj["certificate"] = nil
	}

	if v, ok := d.GetOk("source_ip"); ok {
		if setArgNil {
			obj["source-ip"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingSourceIp(d, v, "source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip"] = t
			}
		}
	} else if d.HasChange("source_ip") {
		obj["source-ip"] = nil
	}

	if v, ok := d.GetOkExists("__change_ip"); ok {
		if setArgNil {
			obj["__change_ip"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingChangeIp(d, v, "__change_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["__change_ip"] = t
			}
		}
	} else if d.HasChange("__change_ip") {
		obj["__change_ip"] = nil
	}

	if v, ok := d.GetOk("upload_option"); ok {
		if setArgNil {
			obj["upload-option"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingUploadOption(d, v, "upload_option", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload-option"] = t
			}
		}
	}

	if v, ok := d.GetOk("upload_interval"); ok {
		if setArgNil {
			obj["upload-interval"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingUploadInterval(d, v, "upload_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("upload_day"); ok {
		if setArgNil {
			obj["upload-day"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingUploadDay(d, v, "upload_day", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload-day"] = t
			}
		}
	} else if d.HasChange("upload_day") {
		obj["upload-day"] = nil
	}

	if v, ok := d.GetOk("upload_time"); ok {
		if setArgNil {
			obj["upload-time"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingUploadTime(d, v, "upload_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload-time"] = t
			}
		}
	} else if d.HasChange("upload_time") {
		obj["upload-time"] = nil
	}

	if v, ok := d.GetOk("reliable"); ok {
		if setArgNil {
			obj["reliable"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingReliable(d, v, "reliable", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["reliable"] = t
			}
		}
	}

	if v, ok := d.GetOk("priority"); ok {
		if setArgNil {
			obj["priority"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingPriority(d, v, "priority", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["priority"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("max_log_rate"); ok {
		if setArgNil {
			obj["max-log-rate"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingMaxLogRate(d, v, "max_log_rate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-log-rate"] = t
			}
		}
	} else if d.HasChange("max_log_rate") {
		obj["max-log-rate"] = nil
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		if setArgNil {
			obj["interface-select-method"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingInterfaceSelectMethod(d, v, "interface_select_method", sv)
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
			t, err := expandLogFortianalyzerSettingInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOkExists("vrf_select"); ok {
		if setArgNil {
			obj["vrf-select"] = nil
		} else {
			t, err := expandLogFortianalyzerSettingVrfSelect(d, v, "vrf_select", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vrf-select"] = t
			}
		}
	} else if d.HasChange("vrf_select") {
		obj["vrf-select"] = nil
	}

	return &obj, nil
}
