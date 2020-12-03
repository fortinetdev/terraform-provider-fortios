// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure WiFi quality of service (QoS) profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"

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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
				ForceNew:     true,
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
		},
	}
}

func resourceWirelessControllerQosProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerQosProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerQosProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerQosProfile(obj)

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

	obj, err := getObjectWirelessControllerQosProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerQosProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerQosProfile(obj, mkey)
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

	err := c.DeleteWirelessControllerQosProfile(mkey)
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

	o, err := c.ReadWirelessControllerQosProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerQosProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerQosProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerQosProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerQosProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileUplink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileDownlink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileUplinkSta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileDownlinkSta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileBurst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileWmm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileWmmUapsd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileCallAdmissionControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileCallCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileBandwidthAdmissionControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileBandwidthCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileDscpWmmMapping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileDscpWmmVo(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWirelessControllerQosProfileDscpWmmVoId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerQosProfileDscpWmmVoId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileDscpWmmVi(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWirelessControllerQosProfileDscpWmmViId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerQosProfileDscpWmmViId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileDscpWmmBe(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWirelessControllerQosProfileDscpWmmBeId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerQosProfileDscpWmmBeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerQosProfileDscpWmmBk(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWirelessControllerQosProfileDscpWmmBkId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerQosProfileDscpWmmBkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerQosProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerQosProfileName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerQosProfileComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("uplink", flattenWirelessControllerQosProfileUplink(o["uplink"], d, "uplink")); err != nil {
		if !fortiAPIPatch(o["uplink"]) {
			return fmt.Errorf("Error reading uplink: %v", err)
		}
	}

	if err = d.Set("downlink", flattenWirelessControllerQosProfileDownlink(o["downlink"], d, "downlink")); err != nil {
		if !fortiAPIPatch(o["downlink"]) {
			return fmt.Errorf("Error reading downlink: %v", err)
		}
	}

	if err = d.Set("uplink_sta", flattenWirelessControllerQosProfileUplinkSta(o["uplink-sta"], d, "uplink_sta")); err != nil {
		if !fortiAPIPatch(o["uplink-sta"]) {
			return fmt.Errorf("Error reading uplink_sta: %v", err)
		}
	}

	if err = d.Set("downlink_sta", flattenWirelessControllerQosProfileDownlinkSta(o["downlink-sta"], d, "downlink_sta")); err != nil {
		if !fortiAPIPatch(o["downlink-sta"]) {
			return fmt.Errorf("Error reading downlink_sta: %v", err)
		}
	}

	if err = d.Set("burst", flattenWirelessControllerQosProfileBurst(o["burst"], d, "burst")); err != nil {
		if !fortiAPIPatch(o["burst"]) {
			return fmt.Errorf("Error reading burst: %v", err)
		}
	}

	if err = d.Set("wmm", flattenWirelessControllerQosProfileWmm(o["wmm"], d, "wmm")); err != nil {
		if !fortiAPIPatch(o["wmm"]) {
			return fmt.Errorf("Error reading wmm: %v", err)
		}
	}

	if err = d.Set("wmm_uapsd", flattenWirelessControllerQosProfileWmmUapsd(o["wmm-uapsd"], d, "wmm_uapsd")); err != nil {
		if !fortiAPIPatch(o["wmm-uapsd"]) {
			return fmt.Errorf("Error reading wmm_uapsd: %v", err)
		}
	}

	if err = d.Set("call_admission_control", flattenWirelessControllerQosProfileCallAdmissionControl(o["call-admission-control"], d, "call_admission_control")); err != nil {
		if !fortiAPIPatch(o["call-admission-control"]) {
			return fmt.Errorf("Error reading call_admission_control: %v", err)
		}
	}

	if err = d.Set("call_capacity", flattenWirelessControllerQosProfileCallCapacity(o["call-capacity"], d, "call_capacity")); err != nil {
		if !fortiAPIPatch(o["call-capacity"]) {
			return fmt.Errorf("Error reading call_capacity: %v", err)
		}
	}

	if err = d.Set("bandwidth_admission_control", flattenWirelessControllerQosProfileBandwidthAdmissionControl(o["bandwidth-admission-control"], d, "bandwidth_admission_control")); err != nil {
		if !fortiAPIPatch(o["bandwidth-admission-control"]) {
			return fmt.Errorf("Error reading bandwidth_admission_control: %v", err)
		}
	}

	if err = d.Set("bandwidth_capacity", flattenWirelessControllerQosProfileBandwidthCapacity(o["bandwidth-capacity"], d, "bandwidth_capacity")); err != nil {
		if !fortiAPIPatch(o["bandwidth-capacity"]) {
			return fmt.Errorf("Error reading bandwidth_capacity: %v", err)
		}
	}

	if err = d.Set("dscp_wmm_mapping", flattenWirelessControllerQosProfileDscpWmmMapping(o["dscp-wmm-mapping"], d, "dscp_wmm_mapping")); err != nil {
		if !fortiAPIPatch(o["dscp-wmm-mapping"]) {
			return fmt.Errorf("Error reading dscp_wmm_mapping: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dscp_wmm_vo", flattenWirelessControllerQosProfileDscpWmmVo(o["dscp-wmm-vo"], d, "dscp_wmm_vo")); err != nil {
			if !fortiAPIPatch(o["dscp-wmm-vo"]) {
				return fmt.Errorf("Error reading dscp_wmm_vo: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dscp_wmm_vo"); ok {
			if err = d.Set("dscp_wmm_vo", flattenWirelessControllerQosProfileDscpWmmVo(o["dscp-wmm-vo"], d, "dscp_wmm_vo")); err != nil {
				if !fortiAPIPatch(o["dscp-wmm-vo"]) {
					return fmt.Errorf("Error reading dscp_wmm_vo: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dscp_wmm_vi", flattenWirelessControllerQosProfileDscpWmmVi(o["dscp-wmm-vi"], d, "dscp_wmm_vi")); err != nil {
			if !fortiAPIPatch(o["dscp-wmm-vi"]) {
				return fmt.Errorf("Error reading dscp_wmm_vi: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dscp_wmm_vi"); ok {
			if err = d.Set("dscp_wmm_vi", flattenWirelessControllerQosProfileDscpWmmVi(o["dscp-wmm-vi"], d, "dscp_wmm_vi")); err != nil {
				if !fortiAPIPatch(o["dscp-wmm-vi"]) {
					return fmt.Errorf("Error reading dscp_wmm_vi: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dscp_wmm_be", flattenWirelessControllerQosProfileDscpWmmBe(o["dscp-wmm-be"], d, "dscp_wmm_be")); err != nil {
			if !fortiAPIPatch(o["dscp-wmm-be"]) {
				return fmt.Errorf("Error reading dscp_wmm_be: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dscp_wmm_be"); ok {
			if err = d.Set("dscp_wmm_be", flattenWirelessControllerQosProfileDscpWmmBe(o["dscp-wmm-be"], d, "dscp_wmm_be")); err != nil {
				if !fortiAPIPatch(o["dscp-wmm-be"]) {
					return fmt.Errorf("Error reading dscp_wmm_be: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dscp_wmm_bk", flattenWirelessControllerQosProfileDscpWmmBk(o["dscp-wmm-bk"], d, "dscp_wmm_bk")); err != nil {
			if !fortiAPIPatch(o["dscp-wmm-bk"]) {
				return fmt.Errorf("Error reading dscp_wmm_bk: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dscp_wmm_bk"); ok {
			if err = d.Set("dscp_wmm_bk", flattenWirelessControllerQosProfileDscpWmmBk(o["dscp-wmm-bk"], d, "dscp_wmm_bk")); err != nil {
				if !fortiAPIPatch(o["dscp-wmm-bk"]) {
					return fmt.Errorf("Error reading dscp_wmm_bk: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerQosProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerQosProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileUplink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileDownlink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileUplinkSta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileDownlinkSta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileBurst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileWmm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileWmmUapsd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileCallAdmissionControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileCallCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileBandwidthAdmissionControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileBandwidthCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileDscpWmmMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileDscpWmmVo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerQosProfileDscpWmmVoId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerQosProfileDscpWmmVoId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileDscpWmmVi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerQosProfileDscpWmmViId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerQosProfileDscpWmmViId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileDscpWmmBe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerQosProfileDscpWmmBeId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerQosProfileDscpWmmBeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerQosProfileDscpWmmBk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerQosProfileDscpWmmBkId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerQosProfileDscpWmmBkId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerQosProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerQosProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWirelessControllerQosProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("uplink"); ok {
		t, err := expandWirelessControllerQosProfileUplink(d, v, "uplink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uplink"] = t
		}
	}

	if v, ok := d.GetOk("downlink"); ok {
		t, err := expandWirelessControllerQosProfileDownlink(d, v, "downlink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["downlink"] = t
		}
	}

	if v, ok := d.GetOk("uplink_sta"); ok {
		t, err := expandWirelessControllerQosProfileUplinkSta(d, v, "uplink_sta")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uplink-sta"] = t
		}
	}

	if v, ok := d.GetOk("downlink_sta"); ok {
		t, err := expandWirelessControllerQosProfileDownlinkSta(d, v, "downlink_sta")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["downlink-sta"] = t
		}
	}

	if v, ok := d.GetOk("burst"); ok {
		t, err := expandWirelessControllerQosProfileBurst(d, v, "burst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["burst"] = t
		}
	}

	if v, ok := d.GetOk("wmm"); ok {
		t, err := expandWirelessControllerQosProfileWmm(d, v, "wmm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm"] = t
		}
	}

	if v, ok := d.GetOk("wmm_uapsd"); ok {
		t, err := expandWirelessControllerQosProfileWmmUapsd(d, v, "wmm_uapsd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm-uapsd"] = t
		}
	}

	if v, ok := d.GetOk("call_admission_control"); ok {
		t, err := expandWirelessControllerQosProfileCallAdmissionControl(d, v, "call_admission_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-admission-control"] = t
		}
	}

	if v, ok := d.GetOk("call_capacity"); ok {
		t, err := expandWirelessControllerQosProfileCallCapacity(d, v, "call_capacity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-capacity"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_admission_control"); ok {
		t, err := expandWirelessControllerQosProfileBandwidthAdmissionControl(d, v, "bandwidth_admission_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-admission-control"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_capacity"); ok {
		t, err := expandWirelessControllerQosProfileBandwidthCapacity(d, v, "bandwidth_capacity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-capacity"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_mapping"); ok {
		t, err := expandWirelessControllerQosProfileDscpWmmMapping(d, v, "dscp_wmm_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-mapping"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_vo"); ok {
		t, err := expandWirelessControllerQosProfileDscpWmmVo(d, v, "dscp_wmm_vo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-vo"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_vi"); ok {
		t, err := expandWirelessControllerQosProfileDscpWmmVi(d, v, "dscp_wmm_vi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-vi"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_be"); ok {
		t, err := expandWirelessControllerQosProfileDscpWmmBe(d, v, "dscp_wmm_be")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-be"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_bk"); ok {
		t, err := expandWirelessControllerQosProfileDscpWmmBk(d, v, "dscp_wmm_bk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-bk"] = t
		}
	}

	return &obj, nil
}
