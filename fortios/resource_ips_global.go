// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPS global parameter.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceIpsGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpsGlobalUpdate,
		Read:   resourceIpsGlobalRead,
		Update: resourceIpsGlobalUpdate,
		Delete: resourceIpsGlobalDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"fail_open": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"database": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_submit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"anomaly_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_limit_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"intelligent_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"socket_size": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),
				Optional:     true,
				Computed:     true,
			},
			"engine_count": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"sync_session_ttl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skype_client_public_ipaddr": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"deep_app_insp_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"deep_app_insp_db_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"exclude_signatures": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"packet_log_queue_depth": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(128, 4096),
				Optional:     true,
				Computed:     true,
			},
			"ngfw_max_scan_range": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tls_active_probe": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"vdom": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
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
					},
				},
			},
		},
	}
}

func resourceIpsGlobalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectIpsGlobal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating IpsGlobal resource while getting object: %v", err)
	}

	o, err := c.UpdateIpsGlobal(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating IpsGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("IpsGlobal")
	}

	return resourceIpsGlobalRead(d, m)
}

func resourceIpsGlobalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteIpsGlobal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting IpsGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsGlobalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadIpsGlobal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading IpsGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIpsGlobal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading IpsGlobal resource from API: %v", err)
	}
	return nil
}

func flattenIpsGlobalFailOpen(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsGlobalDatabase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsGlobalTrafficSubmit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsGlobalAnomalyMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsGlobalSessionLimitMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsGlobalIntelligentMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsGlobalSocketSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsGlobalEngineCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsGlobalSyncSessionTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsGlobalSkypeClientPublicIpaddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsGlobalDeepAppInspTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsGlobalDeepAppInspDbLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsGlobalExcludeSignatures(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsGlobalPacketLogQueueDepth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsGlobalNgfwMaxScanRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsGlobalTlsActiveProbe(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "interface_select_method"
	if _, ok := i["interface-select-method"]; ok {

		result["interface_select_method"] = flattenIpsGlobalTlsActiveProbeInterfaceSelectMethod(i["interface-select-method"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "interface"
	if _, ok := i["interface"]; ok {

		result["interface"] = flattenIpsGlobalTlsActiveProbeInterface(i["interface"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vdom"
	if _, ok := i["vdom"]; ok {

		result["vdom"] = flattenIpsGlobalTlsActiveProbeVdom(i["vdom"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "source_ip"
	if _, ok := i["source-ip"]; ok {

		result["source_ip"] = flattenIpsGlobalTlsActiveProbeSourceIp(i["source-ip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "source_ip6"
	if _, ok := i["source-ip6"]; ok {

		result["source_ip6"] = flattenIpsGlobalTlsActiveProbeSourceIp6(i["source-ip6"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenIpsGlobalTlsActiveProbeInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsGlobalTlsActiveProbeInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsGlobalTlsActiveProbeVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsGlobalTlsActiveProbeSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsGlobalTlsActiveProbeSourceIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectIpsGlobal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fail_open", flattenIpsGlobalFailOpen(o["fail-open"], d, "fail_open", sv)); err != nil {
		if !fortiAPIPatch(o["fail-open"]) {
			return fmt.Errorf("Error reading fail_open: %v", err)
		}
	}

	if err = d.Set("database", flattenIpsGlobalDatabase(o["database"], d, "database", sv)); err != nil {
		if !fortiAPIPatch(o["database"]) {
			return fmt.Errorf("Error reading database: %v", err)
		}
	}

	if err = d.Set("traffic_submit", flattenIpsGlobalTrafficSubmit(o["traffic-submit"], d, "traffic_submit", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-submit"]) {
			return fmt.Errorf("Error reading traffic_submit: %v", err)
		}
	}

	if err = d.Set("anomaly_mode", flattenIpsGlobalAnomalyMode(o["anomaly-mode"], d, "anomaly_mode", sv)); err != nil {
		if !fortiAPIPatch(o["anomaly-mode"]) {
			return fmt.Errorf("Error reading anomaly_mode: %v", err)
		}
	}

	if err = d.Set("session_limit_mode", flattenIpsGlobalSessionLimitMode(o["session-limit-mode"], d, "session_limit_mode", sv)); err != nil {
		if !fortiAPIPatch(o["session-limit-mode"]) {
			return fmt.Errorf("Error reading session_limit_mode: %v", err)
		}
	}

	if err = d.Set("intelligent_mode", flattenIpsGlobalIntelligentMode(o["intelligent-mode"], d, "intelligent_mode", sv)); err != nil {
		if !fortiAPIPatch(o["intelligent-mode"]) {
			return fmt.Errorf("Error reading intelligent_mode: %v", err)
		}
	}

	if err = d.Set("socket_size", flattenIpsGlobalSocketSize(o["socket-size"], d, "socket_size", sv)); err != nil {
		if !fortiAPIPatch(o["socket-size"]) {
			return fmt.Errorf("Error reading socket_size: %v", err)
		}
	}

	if err = d.Set("engine_count", flattenIpsGlobalEngineCount(o["engine-count"], d, "engine_count", sv)); err != nil {
		if !fortiAPIPatch(o["engine-count"]) {
			return fmt.Errorf("Error reading engine_count: %v", err)
		}
	}

	if err = d.Set("sync_session_ttl", flattenIpsGlobalSyncSessionTtl(o["sync-session-ttl"], d, "sync_session_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["sync-session-ttl"]) {
			return fmt.Errorf("Error reading sync_session_ttl: %v", err)
		}
	}

	if err = d.Set("skype_client_public_ipaddr", flattenIpsGlobalSkypeClientPublicIpaddr(o["skype-client-public-ipaddr"], d, "skype_client_public_ipaddr", sv)); err != nil {
		if !fortiAPIPatch(o["skype-client-public-ipaddr"]) {
			return fmt.Errorf("Error reading skype_client_public_ipaddr: %v", err)
		}
	}

	if err = d.Set("deep_app_insp_timeout", flattenIpsGlobalDeepAppInspTimeout(o["deep-app-insp-timeout"], d, "deep_app_insp_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["deep-app-insp-timeout"]) {
			return fmt.Errorf("Error reading deep_app_insp_timeout: %v", err)
		}
	}

	if err = d.Set("deep_app_insp_db_limit", flattenIpsGlobalDeepAppInspDbLimit(o["deep-app-insp-db-limit"], d, "deep_app_insp_db_limit", sv)); err != nil {
		if !fortiAPIPatch(o["deep-app-insp-db-limit"]) {
			return fmt.Errorf("Error reading deep_app_insp_db_limit: %v", err)
		}
	}

	if err = d.Set("exclude_signatures", flattenIpsGlobalExcludeSignatures(o["exclude-signatures"], d, "exclude_signatures", sv)); err != nil {
		if !fortiAPIPatch(o["exclude-signatures"]) {
			return fmt.Errorf("Error reading exclude_signatures: %v", err)
		}
	}

	if err = d.Set("packet_log_queue_depth", flattenIpsGlobalPacketLogQueueDepth(o["packet-log-queue-depth"], d, "packet_log_queue_depth", sv)); err != nil {
		if !fortiAPIPatch(o["packet-log-queue-depth"]) {
			return fmt.Errorf("Error reading packet_log_queue_depth: %v", err)
		}
	}

	if err = d.Set("ngfw_max_scan_range", flattenIpsGlobalNgfwMaxScanRange(o["ngfw-max-scan-range"], d, "ngfw_max_scan_range", sv)); err != nil {
		if !fortiAPIPatch(o["ngfw-max-scan-range"]) {
			return fmt.Errorf("Error reading ngfw_max_scan_range: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("tls_active_probe", flattenIpsGlobalTlsActiveProbe(o["tls-active-probe"], d, "tls_active_probe", sv)); err != nil {
			if !fortiAPIPatch(o["tls-active-probe"]) {
				return fmt.Errorf("Error reading tls_active_probe: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tls_active_probe"); ok {
			if err = d.Set("tls_active_probe", flattenIpsGlobalTlsActiveProbe(o["tls-active-probe"], d, "tls_active_probe", sv)); err != nil {
				if !fortiAPIPatch(o["tls-active-probe"]) {
					return fmt.Errorf("Error reading tls_active_probe: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenIpsGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandIpsGlobalFailOpen(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalDatabase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalTrafficSubmit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalAnomalyMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalSessionLimitMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalIntelligentMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalSocketSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalEngineCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalSyncSessionTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalSkypeClientPublicIpaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalDeepAppInspTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalDeepAppInspDbLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalExcludeSignatures(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalPacketLogQueueDepth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalNgfwMaxScanRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalTlsActiveProbe(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "interface_select_method"
	if _, ok := d.GetOk(pre_append); ok {

		result["interface-select-method"], _ = expandIpsGlobalTlsActiveProbeInterfaceSelectMethod(d, i["interface_select_method"], pre_append, sv)
	}
	pre_append = pre + ".0." + "interface"
	if _, ok := d.GetOk(pre_append); ok {

		result["interface"], _ = expandIpsGlobalTlsActiveProbeInterface(d, i["interface"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vdom"
	if _, ok := d.GetOk(pre_append); ok {

		result["vdom"], _ = expandIpsGlobalTlsActiveProbeVdom(d, i["vdom"], pre_append, sv)
	}
	pre_append = pre + ".0." + "source_ip"
	if _, ok := d.GetOk(pre_append); ok {

		result["source-ip"], _ = expandIpsGlobalTlsActiveProbeSourceIp(d, i["source_ip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "source_ip6"
	if _, ok := d.GetOk(pre_append); ok {

		result["source-ip6"], _ = expandIpsGlobalTlsActiveProbeSourceIp6(d, i["source_ip6"], pre_append, sv)
	}

	return result, nil
}

func expandIpsGlobalTlsActiveProbeInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalTlsActiveProbeInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalTlsActiveProbeVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalTlsActiveProbeSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalTlsActiveProbeSourceIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectIpsGlobal(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fail_open"); ok {

		t, err := expandIpsGlobalFailOpen(d, v, "fail_open", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-open"] = t
		}
	}

	if v, ok := d.GetOk("database"); ok {

		t, err := expandIpsGlobalDatabase(d, v, "database", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database"] = t
		}
	}

	if v, ok := d.GetOk("traffic_submit"); ok {

		t, err := expandIpsGlobalTrafficSubmit(d, v, "traffic_submit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-submit"] = t
		}
	}

	if v, ok := d.GetOk("anomaly_mode"); ok {

		t, err := expandIpsGlobalAnomalyMode(d, v, "anomaly_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anomaly-mode"] = t
		}
	}

	if v, ok := d.GetOk("session_limit_mode"); ok {

		t, err := expandIpsGlobalSessionLimitMode(d, v, "session_limit_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-limit-mode"] = t
		}
	}

	if v, ok := d.GetOk("intelligent_mode"); ok {

		t, err := expandIpsGlobalIntelligentMode(d, v, "intelligent_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intelligent-mode"] = t
		}
	}

	if v, ok := d.GetOkExists("socket_size"); ok {

		t, err := expandIpsGlobalSocketSize(d, v, "socket_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["socket-size"] = t
		}
	}

	if v, ok := d.GetOkExists("engine_count"); ok {

		t, err := expandIpsGlobalEngineCount(d, v, "engine_count", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["engine-count"] = t
		}
	}

	if v, ok := d.GetOk("sync_session_ttl"); ok {

		t, err := expandIpsGlobalSyncSessionTtl(d, v, "sync_session_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-session-ttl"] = t
		}
	}

	if v, ok := d.GetOk("skype_client_public_ipaddr"); ok {

		t, err := expandIpsGlobalSkypeClientPublicIpaddr(d, v, "skype_client_public_ipaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["skype-client-public-ipaddr"] = t
		}
	}

	if v, ok := d.GetOkExists("deep_app_insp_timeout"); ok {

		t, err := expandIpsGlobalDeepAppInspTimeout(d, v, "deep_app_insp_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deep-app-insp-timeout"] = t
		}
	}

	if v, ok := d.GetOkExists("deep_app_insp_db_limit"); ok {

		t, err := expandIpsGlobalDeepAppInspDbLimit(d, v, "deep_app_insp_db_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deep-app-insp-db-limit"] = t
		}
	}

	if v, ok := d.GetOk("exclude_signatures"); ok {

		t, err := expandIpsGlobalExcludeSignatures(d, v, "exclude_signatures", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclude-signatures"] = t
		}
	}

	if v, ok := d.GetOk("packet_log_queue_depth"); ok {

		t, err := expandIpsGlobalPacketLogQueueDepth(d, v, "packet_log_queue_depth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-log-queue-depth"] = t
		}
	}

	if v, ok := d.GetOkExists("ngfw_max_scan_range"); ok {

		t, err := expandIpsGlobalNgfwMaxScanRange(d, v, "ngfw_max_scan_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ngfw-max-scan-range"] = t
		}
	}

	if v, ok := d.GetOk("tls_active_probe"); ok {

		t, err := expandIpsGlobalTlsActiveProbe(d, v, "tls_active_probe", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-active-probe"] = t
		}
	}

	return &obj, nil
}
