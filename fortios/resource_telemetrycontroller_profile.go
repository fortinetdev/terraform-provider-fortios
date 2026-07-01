// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiTelemetry profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceTelemetryControllerProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceTelemetryControllerProfileCreate,
		Read:   resourceTelemetryControllerProfileRead,
		Update: resourceTelemetryControllerProfileUpdate,
		Delete: resourceTelemetryControllerProfileDelete,

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
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"app_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
						"monitor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(60000, 86400000),
							Optional:     true,
							Computed:     true,
						},
						"sla": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sla_factor": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"experience_score_threshold": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 10),
										Optional:     true,
										Computed:     true,
									},
									"failure_rate_threshold": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 100),
										Optional:     true,
										Computed:     true,
									},
									"latency_threshold": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 10000000),
										Optional:     true,
										Computed:     true,
									},
									"jitter_threshold": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 10000000),
										Optional:     true,
										Computed:     true,
									},
									"packet_loss_threshold": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 100),
										Optional:     true,
										Computed:     true,
									},
									"ttfb_threshold": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 10000000),
										Optional:     true,
										Computed:     true,
									},
									"atdt_threshold": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 10000000),
										Optional:     true,
										Computed:     true,
									},
									"tcp_rtt_threshold": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 10000000),
										Optional:     true,
										Computed:     true,
									},
									"dns_time_threshold": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 10000000),
										Optional:     true,
										Computed:     true,
									},
									"tls_time_threshold": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 10000000),
										Optional:     true,
										Computed:     true,
									},
									"app_throughput_threshold": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 10000),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
					},
				},
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

func resourceTelemetryControllerProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectTelemetryControllerProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating TelemetryControllerProfile resource while getting object: %v", err)
	}

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadTelemetryControllerProfile(mkey, vdomparam)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateTelemetryControllerProfile(obj, mkey, vdomparam)
			if err != nil {
				return fmt.Errorf("Error updating TelemetryControllerProfile resource: %v", err)
			}
		}
	}

	if !existing {
		o, err = c.CreateTelemetryControllerProfile(obj, vdomparam)

		if err != nil {
			return fmt.Errorf("Error creating TelemetryControllerProfile resource: %v", err)
		}
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("TelemetryControllerProfile")
	}

	return resourceTelemetryControllerProfileRead(d, m)
}

func resourceTelemetryControllerProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectTelemetryControllerProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating TelemetryControllerProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateTelemetryControllerProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating TelemetryControllerProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("TelemetryControllerProfile")
	}

	return resourceTelemetryControllerProfileRead(d, m)
}

func resourceTelemetryControllerProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteTelemetryControllerProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting TelemetryControllerProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceTelemetryControllerProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadTelemetryControllerProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading TelemetryControllerProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectTelemetryControllerProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading TelemetryControllerProfile resource from API: %v", err)
	}
	return nil
}

func flattenTelemetryControllerProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenTelemetryControllerProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenTelemetryControllerProfileApplication(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenTelemetryControllerProfileApplicationId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["app-name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "app_name"
			}
			tmp["app_name"] = flattenTelemetryControllerProfileApplicationAppName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["monitor"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
			}
			tmp["monitor"] = flattenTelemetryControllerProfileApplicationMonitor(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["interval"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
			}
			tmp["interval"] = flattenTelemetryControllerProfileApplicationInterval(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["sla"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
			}
			tmp["sla"] = flattenTelemetryControllerProfileApplicationSla(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenTelemetryControllerProfileApplicationId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenTelemetryControllerProfileApplicationAppName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenTelemetryControllerProfileApplicationMonitor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenTelemetryControllerProfileApplicationInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenTelemetryControllerProfileApplicationSla(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "sla_factor"
	if _, ok := i["sla-factor"]; ok {
		result["sla_factor"] = flattenTelemetryControllerProfileApplicationSlaSlaFactor(i["sla-factor"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "experience_score_threshold"
	if _, ok := i["experience-score-threshold"]; ok {
		result["experience_score_threshold"] = flattenTelemetryControllerProfileApplicationSlaExperienceScoreThreshold(i["experience-score-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "failure_rate_threshold"
	if _, ok := i["failure-rate-threshold"]; ok {
		result["failure_rate_threshold"] = flattenTelemetryControllerProfileApplicationSlaFailureRateThreshold(i["failure-rate-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "latency_threshold"
	if _, ok := i["latency-threshold"]; ok {
		result["latency_threshold"] = flattenTelemetryControllerProfileApplicationSlaLatencyThreshold(i["latency-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "jitter_threshold"
	if _, ok := i["jitter-threshold"]; ok {
		result["jitter_threshold"] = flattenTelemetryControllerProfileApplicationSlaJitterThreshold(i["jitter-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "packet_loss_threshold"
	if _, ok := i["packet-loss-threshold"]; ok {
		result["packet_loss_threshold"] = flattenTelemetryControllerProfileApplicationSlaPacketLossThreshold(i["packet-loss-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ttfb_threshold"
	if _, ok := i["ttfb-threshold"]; ok {
		result["ttfb_threshold"] = flattenTelemetryControllerProfileApplicationSlaTtfbThreshold(i["ttfb-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "atdt_threshold"
	if _, ok := i["atdt-threshold"]; ok {
		result["atdt_threshold"] = flattenTelemetryControllerProfileApplicationSlaAtdtThreshold(i["atdt-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_rtt_threshold"
	if _, ok := i["tcp-rtt-threshold"]; ok {
		result["tcp_rtt_threshold"] = flattenTelemetryControllerProfileApplicationSlaTcpRttThreshold(i["tcp-rtt-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dns_time_threshold"
	if _, ok := i["dns-time-threshold"]; ok {
		result["dns_time_threshold"] = flattenTelemetryControllerProfileApplicationSlaDnsTimeThreshold(i["dns-time-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tls_time_threshold"
	if _, ok := i["tls-time-threshold"]; ok {
		result["tls_time_threshold"] = flattenTelemetryControllerProfileApplicationSlaTlsTimeThreshold(i["tls-time-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "app_throughput_threshold"
	if _, ok := i["app-throughput-threshold"]; ok {
		result["app_throughput_threshold"] = flattenTelemetryControllerProfileApplicationSlaAppThroughputThreshold(i["app-throughput-threshold"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenTelemetryControllerProfileApplicationSlaSlaFactor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenTelemetryControllerProfileApplicationSlaExperienceScoreThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenTelemetryControllerProfileApplicationSlaFailureRateThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenTelemetryControllerProfileApplicationSlaLatencyThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenTelemetryControllerProfileApplicationSlaJitterThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenTelemetryControllerProfileApplicationSlaPacketLossThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenTelemetryControllerProfileApplicationSlaTtfbThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenTelemetryControllerProfileApplicationSlaAtdtThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenTelemetryControllerProfileApplicationSlaTcpRttThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenTelemetryControllerProfileApplicationSlaDnsTimeThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenTelemetryControllerProfileApplicationSlaTlsTimeThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenTelemetryControllerProfileApplicationSlaAppThroughputThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectTelemetryControllerProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenTelemetryControllerProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenTelemetryControllerProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("application", flattenTelemetryControllerProfileApplication(o["application"], d, "application", sv)); err != nil {
			if !fortiAPIPatch(o["application"]) {
				return fmt.Errorf("Error reading application: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("application"); ok {
			if err = d.Set("application", flattenTelemetryControllerProfileApplication(o["application"], d, "application", sv)); err != nil {
				if !fortiAPIPatch(o["application"]) {
					return fmt.Errorf("Error reading application: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenTelemetryControllerProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandTelemetryControllerProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerProfileApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandTelemetryControllerProfileApplicationId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "app_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["app-name"], _ = expandTelemetryControllerProfileApplicationAppName(d, i["app_name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["app-name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["monitor"], _ = expandTelemetryControllerProfileApplicationMonitor(d, i["monitor"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interval"], _ = expandTelemetryControllerProfileApplicationInterval(d, i["interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sla"], _ = expandTelemetryControllerProfileApplicationSla(d, i["sla"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["sla"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandTelemetryControllerProfileApplicationId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerProfileApplicationAppName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerProfileApplicationMonitor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerProfileApplicationInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerProfileApplicationSla(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "sla_factor"
	if _, ok := d.GetOk(pre_append); ok {
		result["sla-factor"], _ = expandTelemetryControllerProfileApplicationSlaSlaFactor(d, i["sla_factor"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["sla-factor"] = nil
	}
	pre_append = pre + ".0." + "experience_score_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["experience-score-threshold"], _ = expandTelemetryControllerProfileApplicationSlaExperienceScoreThreshold(d, i["experience_score_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "failure_rate_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["failure-rate-threshold"], _ = expandTelemetryControllerProfileApplicationSlaFailureRateThreshold(d, i["failure_rate_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "latency_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["latency-threshold"], _ = expandTelemetryControllerProfileApplicationSlaLatencyThreshold(d, i["latency_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "jitter_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["jitter-threshold"], _ = expandTelemetryControllerProfileApplicationSlaJitterThreshold(d, i["jitter_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "packet_loss_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["packet-loss-threshold"], _ = expandTelemetryControllerProfileApplicationSlaPacketLossThreshold(d, i["packet_loss_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ttfb_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["ttfb-threshold"], _ = expandTelemetryControllerProfileApplicationSlaTtfbThreshold(d, i["ttfb_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "atdt_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["atdt-threshold"], _ = expandTelemetryControllerProfileApplicationSlaAtdtThreshold(d, i["atdt_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tcp_rtt_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-rtt-threshold"], _ = expandTelemetryControllerProfileApplicationSlaTcpRttThreshold(d, i["tcp_rtt_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dns_time_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["dns-time-threshold"], _ = expandTelemetryControllerProfileApplicationSlaDnsTimeThreshold(d, i["dns_time_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tls_time_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["tls-time-threshold"], _ = expandTelemetryControllerProfileApplicationSlaTlsTimeThreshold(d, i["tls_time_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "app_throughput_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["app-throughput-threshold"], _ = expandTelemetryControllerProfileApplicationSlaAppThroughputThreshold(d, i["app_throughput_threshold"], pre_append, sv)
	}

	return result, nil
}

func expandTelemetryControllerProfileApplicationSlaSlaFactor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerProfileApplicationSlaExperienceScoreThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerProfileApplicationSlaFailureRateThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerProfileApplicationSlaLatencyThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerProfileApplicationSlaJitterThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerProfileApplicationSlaPacketLossThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerProfileApplicationSlaTtfbThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerProfileApplicationSlaAtdtThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerProfileApplicationSlaTcpRttThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerProfileApplicationSlaDnsTimeThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerProfileApplicationSlaTlsTimeThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerProfileApplicationSlaAppThroughputThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectTelemetryControllerProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandTelemetryControllerProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandTelemetryControllerProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandTelemetryControllerProfileApplication(d, v, "application", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	return &obj, nil
}
