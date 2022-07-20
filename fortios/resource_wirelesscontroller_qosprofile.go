// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure WiFi quality of service (QoS) profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerQosProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerQosProfileCreate,
		Read:   resourceWirelessControllerQosProfileRead,
		Update: resourceWirelessControllerQosProfileUpdate,
		Delete: resourceWirelessControllerQosProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
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
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"uplink": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2097152),
				Optional:     true,
				Computed:     true,
			},
			"downlink": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2097152),
				Optional:     true,
				Computed:     true,
			},
			"uplink_sta": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2097152),
				Optional:     true,
				Computed:     true,
			},
			"downlink_sta": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2097152),
				Optional:     true,
				Computed:     true,
			},
			"burst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wmm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wmm_uapsd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"call_admission_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"call_capacity": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 60),
				Optional:     true,
				Computed:     true,
			},
			"bandwidth_admission_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bandwidth_capacity": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 600000),
				Optional:     true,
				Computed:     true,
			},
			"dscp_wmm_mapping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp_wmm_vo": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dscp_wmm_vi": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dscp_wmm_be": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dscp_wmm_bk": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"wmm_dscp_marking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wmm_vo_dscp": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"wmm_vi_dscp": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"wmm_be_dscp": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"wmm_bk_dscp": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWirelessControllerQosProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerQosProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerQosProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerQosProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerQosProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerQosProfile")
	}

	return resourceWirelessControllerQosProfileRead(d, m)
}

func resourceWirelessControllerQosProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerQosProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerQosProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerQosProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerQosProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerQosProfile")
	}

	return resourceWirelessControllerQosProfileRead(d, m)
}

func resourceWirelessControllerQosProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerQosProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerQosProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerQosProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWirelessControllerQosProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerQosProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerQosProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerQosProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerQosProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileUplink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileDownlink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileUplinkSta(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileDownlinkSta(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileBurst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileWmm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileWmmUapsd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileCallAdmissionControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileCallCapacity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileBandwidthAdmissionControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileBandwidthCapacity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileDscpWmmMapping(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileDscpWmmVo(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenWirelessControllerQosProfileDscpWmmVoId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWirelessControllerQosProfileDscpWmmVoId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileDscpWmmVi(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenWirelessControllerQosProfileDscpWmmViId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWirelessControllerQosProfileDscpWmmViId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileDscpWmmBe(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenWirelessControllerQosProfileDscpWmmBeId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWirelessControllerQosProfileDscpWmmBeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileDscpWmmBk(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenWirelessControllerQosProfileDscpWmmBkId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWirelessControllerQosProfileDscpWmmBkId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileWmmDscpMarking(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileWmmVoDscp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileWmmViDscp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileWmmBeDscp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileWmmBkDscp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerQosProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerQosProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerQosProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("uplink", flattenWirelessControllerQosProfileUplink(o["uplink"], d, "uplink", sv)); err != nil {
		if !fortiAPIPatch(o["uplink"]) {
			return fmt.Errorf("Error reading uplink: %v", err)
		}
	}

	if err = d.Set("downlink", flattenWirelessControllerQosProfileDownlink(o["downlink"], d, "downlink", sv)); err != nil {
		if !fortiAPIPatch(o["downlink"]) {
			return fmt.Errorf("Error reading downlink: %v", err)
		}
	}

	if err = d.Set("uplink_sta", flattenWirelessControllerQosProfileUplinkSta(o["uplink-sta"], d, "uplink_sta", sv)); err != nil {
		if !fortiAPIPatch(o["uplink-sta"]) {
			return fmt.Errorf("Error reading uplink_sta: %v", err)
		}
	}

	if err = d.Set("downlink_sta", flattenWirelessControllerQosProfileDownlinkSta(o["downlink-sta"], d, "downlink_sta", sv)); err != nil {
		if !fortiAPIPatch(o["downlink-sta"]) {
			return fmt.Errorf("Error reading downlink_sta: %v", err)
		}
	}

	if err = d.Set("burst", flattenWirelessControllerQosProfileBurst(o["burst"], d, "burst", sv)); err != nil {
		if !fortiAPIPatch(o["burst"]) {
			return fmt.Errorf("Error reading burst: %v", err)
		}
	}

	if err = d.Set("wmm", flattenWirelessControllerQosProfileWmm(o["wmm"], d, "wmm", sv)); err != nil {
		if !fortiAPIPatch(o["wmm"]) {
			return fmt.Errorf("Error reading wmm: %v", err)
		}
	}

	if err = d.Set("wmm_uapsd", flattenWirelessControllerQosProfileWmmUapsd(o["wmm-uapsd"], d, "wmm_uapsd", sv)); err != nil {
		if !fortiAPIPatch(o["wmm-uapsd"]) {
			return fmt.Errorf("Error reading wmm_uapsd: %v", err)
		}
	}

	if err = d.Set("call_admission_control", flattenWirelessControllerQosProfileCallAdmissionControl(o["call-admission-control"], d, "call_admission_control", sv)); err != nil {
		if !fortiAPIPatch(o["call-admission-control"]) {
			return fmt.Errorf("Error reading call_admission_control: %v", err)
		}
	}

	if err = d.Set("call_capacity", flattenWirelessControllerQosProfileCallCapacity(o["call-capacity"], d, "call_capacity", sv)); err != nil {
		if !fortiAPIPatch(o["call-capacity"]) {
			return fmt.Errorf("Error reading call_capacity: %v", err)
		}
	}

	if err = d.Set("bandwidth_admission_control", flattenWirelessControllerQosProfileBandwidthAdmissionControl(o["bandwidth-admission-control"], d, "bandwidth_admission_control", sv)); err != nil {
		if !fortiAPIPatch(o["bandwidth-admission-control"]) {
			return fmt.Errorf("Error reading bandwidth_admission_control: %v", err)
		}
	}

	if err = d.Set("bandwidth_capacity", flattenWirelessControllerQosProfileBandwidthCapacity(o["bandwidth-capacity"], d, "bandwidth_capacity", sv)); err != nil {
		if !fortiAPIPatch(o["bandwidth-capacity"]) {
			return fmt.Errorf("Error reading bandwidth_capacity: %v", err)
		}
	}

	if err = d.Set("dscp_wmm_mapping", flattenWirelessControllerQosProfileDscpWmmMapping(o["dscp-wmm-mapping"], d, "dscp_wmm_mapping", sv)); err != nil {
		if !fortiAPIPatch(o["dscp-wmm-mapping"]) {
			return fmt.Errorf("Error reading dscp_wmm_mapping: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dscp_wmm_vo", flattenWirelessControllerQosProfileDscpWmmVo(o["dscp-wmm-vo"], d, "dscp_wmm_vo", sv)); err != nil {
			if !fortiAPIPatch(o["dscp-wmm-vo"]) {
				return fmt.Errorf("Error reading dscp_wmm_vo: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dscp_wmm_vo"); ok {
			if err = d.Set("dscp_wmm_vo", flattenWirelessControllerQosProfileDscpWmmVo(o["dscp-wmm-vo"], d, "dscp_wmm_vo", sv)); err != nil {
				if !fortiAPIPatch(o["dscp-wmm-vo"]) {
					return fmt.Errorf("Error reading dscp_wmm_vo: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dscp_wmm_vi", flattenWirelessControllerQosProfileDscpWmmVi(o["dscp-wmm-vi"], d, "dscp_wmm_vi", sv)); err != nil {
			if !fortiAPIPatch(o["dscp-wmm-vi"]) {
				return fmt.Errorf("Error reading dscp_wmm_vi: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dscp_wmm_vi"); ok {
			if err = d.Set("dscp_wmm_vi", flattenWirelessControllerQosProfileDscpWmmVi(o["dscp-wmm-vi"], d, "dscp_wmm_vi", sv)); err != nil {
				if !fortiAPIPatch(o["dscp-wmm-vi"]) {
					return fmt.Errorf("Error reading dscp_wmm_vi: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dscp_wmm_be", flattenWirelessControllerQosProfileDscpWmmBe(o["dscp-wmm-be"], d, "dscp_wmm_be", sv)); err != nil {
			if !fortiAPIPatch(o["dscp-wmm-be"]) {
				return fmt.Errorf("Error reading dscp_wmm_be: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dscp_wmm_be"); ok {
			if err = d.Set("dscp_wmm_be", flattenWirelessControllerQosProfileDscpWmmBe(o["dscp-wmm-be"], d, "dscp_wmm_be", sv)); err != nil {
				if !fortiAPIPatch(o["dscp-wmm-be"]) {
					return fmt.Errorf("Error reading dscp_wmm_be: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dscp_wmm_bk", flattenWirelessControllerQosProfileDscpWmmBk(o["dscp-wmm-bk"], d, "dscp_wmm_bk", sv)); err != nil {
			if !fortiAPIPatch(o["dscp-wmm-bk"]) {
				return fmt.Errorf("Error reading dscp_wmm_bk: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dscp_wmm_bk"); ok {
			if err = d.Set("dscp_wmm_bk", flattenWirelessControllerQosProfileDscpWmmBk(o["dscp-wmm-bk"], d, "dscp_wmm_bk", sv)); err != nil {
				if !fortiAPIPatch(o["dscp-wmm-bk"]) {
					return fmt.Errorf("Error reading dscp_wmm_bk: %v", err)
				}
			}
		}
	}

	if err = d.Set("wmm_dscp_marking", flattenWirelessControllerQosProfileWmmDscpMarking(o["wmm-dscp-marking"], d, "wmm_dscp_marking", sv)); err != nil {
		if !fortiAPIPatch(o["wmm-dscp-marking"]) {
			return fmt.Errorf("Error reading wmm_dscp_marking: %v", err)
		}
	}

	if err = d.Set("wmm_vo_dscp", flattenWirelessControllerQosProfileWmmVoDscp(o["wmm-vo-dscp"], d, "wmm_vo_dscp", sv)); err != nil {
		if !fortiAPIPatch(o["wmm-vo-dscp"]) {
			return fmt.Errorf("Error reading wmm_vo_dscp: %v", err)
		}
	}

	if err = d.Set("wmm_vi_dscp", flattenWirelessControllerQosProfileWmmViDscp(o["wmm-vi-dscp"], d, "wmm_vi_dscp", sv)); err != nil {
		if !fortiAPIPatch(o["wmm-vi-dscp"]) {
			return fmt.Errorf("Error reading wmm_vi_dscp: %v", err)
		}
	}

	if err = d.Set("wmm_be_dscp", flattenWirelessControllerQosProfileWmmBeDscp(o["wmm-be-dscp"], d, "wmm_be_dscp", sv)); err != nil {
		if !fortiAPIPatch(o["wmm-be-dscp"]) {
			return fmt.Errorf("Error reading wmm_be_dscp: %v", err)
		}
	}

	if err = d.Set("wmm_bk_dscp", flattenWirelessControllerQosProfileWmmBkDscp(o["wmm-bk-dscp"], d, "wmm_bk_dscp", sv)); err != nil {
		if !fortiAPIPatch(o["wmm-bk-dscp"]) {
			return fmt.Errorf("Error reading wmm_bk_dscp: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerQosProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerQosProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileUplink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileDownlink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileUplinkSta(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileDownlinkSta(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileBurst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileWmm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileWmmUapsd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileCallAdmissionControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileCallCapacity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileBandwidthAdmissionControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileBandwidthCapacity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileDscpWmmMapping(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileDscpWmmVo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandWirelessControllerQosProfileDscpWmmVoId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerQosProfileDscpWmmVoId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileDscpWmmVi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandWirelessControllerQosProfileDscpWmmViId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerQosProfileDscpWmmViId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileDscpWmmBe(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandWirelessControllerQosProfileDscpWmmBeId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerQosProfileDscpWmmBeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileDscpWmmBk(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandWirelessControllerQosProfileDscpWmmBkId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerQosProfileDscpWmmBkId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileWmmDscpMarking(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileWmmVoDscp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileWmmViDscp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileWmmBeDscp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileWmmBkDscp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerQosProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandWirelessControllerQosProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandWirelessControllerQosProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOkExists("uplink"); ok {

		t, err := expandWirelessControllerQosProfileUplink(d, v, "uplink", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uplink"] = t
		}
	}

	if v, ok := d.GetOkExists("downlink"); ok {

		t, err := expandWirelessControllerQosProfileDownlink(d, v, "downlink", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["downlink"] = t
		}
	}

	if v, ok := d.GetOkExists("uplink_sta"); ok {

		t, err := expandWirelessControllerQosProfileUplinkSta(d, v, "uplink_sta", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uplink-sta"] = t
		}
	}

	if v, ok := d.GetOkExists("downlink_sta"); ok {

		t, err := expandWirelessControllerQosProfileDownlinkSta(d, v, "downlink_sta", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["downlink-sta"] = t
		}
	}

	if v, ok := d.GetOk("burst"); ok {

		t, err := expandWirelessControllerQosProfileBurst(d, v, "burst", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["burst"] = t
		}
	}

	if v, ok := d.GetOk("wmm"); ok {

		t, err := expandWirelessControllerQosProfileWmm(d, v, "wmm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm"] = t
		}
	}

	if v, ok := d.GetOk("wmm_uapsd"); ok {

		t, err := expandWirelessControllerQosProfileWmmUapsd(d, v, "wmm_uapsd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm-uapsd"] = t
		}
	}

	if v, ok := d.GetOk("call_admission_control"); ok {

		t, err := expandWirelessControllerQosProfileCallAdmissionControl(d, v, "call_admission_control", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-admission-control"] = t
		}
	}

	if v, ok := d.GetOkExists("call_capacity"); ok {

		t, err := expandWirelessControllerQosProfileCallCapacity(d, v, "call_capacity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-capacity"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_admission_control"); ok {

		t, err := expandWirelessControllerQosProfileBandwidthAdmissionControl(d, v, "bandwidth_admission_control", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-admission-control"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_capacity"); ok {

		t, err := expandWirelessControllerQosProfileBandwidthCapacity(d, v, "bandwidth_capacity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-capacity"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_mapping"); ok {

		t, err := expandWirelessControllerQosProfileDscpWmmMapping(d, v, "dscp_wmm_mapping", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-mapping"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_vo"); ok {

		t, err := expandWirelessControllerQosProfileDscpWmmVo(d, v, "dscp_wmm_vo", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-vo"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_vi"); ok {

		t, err := expandWirelessControllerQosProfileDscpWmmVi(d, v, "dscp_wmm_vi", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-vi"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_be"); ok {

		t, err := expandWirelessControllerQosProfileDscpWmmBe(d, v, "dscp_wmm_be", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-be"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_bk"); ok {

		t, err := expandWirelessControllerQosProfileDscpWmmBk(d, v, "dscp_wmm_bk", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-bk"] = t
		}
	}

	if v, ok := d.GetOk("wmm_dscp_marking"); ok {

		t, err := expandWirelessControllerQosProfileWmmDscpMarking(d, v, "wmm_dscp_marking", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm-dscp-marking"] = t
		}
	}

	if v, ok := d.GetOkExists("wmm_vo_dscp"); ok {

		t, err := expandWirelessControllerQosProfileWmmVoDscp(d, v, "wmm_vo_dscp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm-vo-dscp"] = t
		}
	}

	if v, ok := d.GetOkExists("wmm_vi_dscp"); ok {

		t, err := expandWirelessControllerQosProfileWmmViDscp(d, v, "wmm_vi_dscp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm-vi-dscp"] = t
		}
	}

	if v, ok := d.GetOkExists("wmm_be_dscp"); ok {

		t, err := expandWirelessControllerQosProfileWmmBeDscp(d, v, "wmm_be_dscp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm-be-dscp"] = t
		}
	}

	if v, ok := d.GetOkExists("wmm_bk_dscp"); ok {

		t, err := expandWirelessControllerQosProfileWmmBkDscp(d, v, "wmm_bk_dscp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm-bk-dscp"] = t
		}
	}

	return &obj, nil
}
