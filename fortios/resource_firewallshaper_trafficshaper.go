// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure shared traffic shaper.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallShaperTrafficShaper() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallShaperTrafficShaperCreate,
		Read:   resourceFirewallShaperTrafficShaperRead,
		Update: resourceFirewallShaperTrafficShaperUpdate,
		Delete: resourceFirewallShaperTrafficShaperDelete,

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
				Optional:     true,
				Computed:     true,
			},
			"guaranteed_bandwidth": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 80000000),
				Optional:     true,
				Computed:     true,
			},
			"maximum_bandwidth": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 80000000),
				Optional:     true,
				Computed:     true,
			},
			"bandwidth_unit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"per_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp_marking_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exceed_bandwidth": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 80000000),
				Optional:     true,
				Computed:     true,
			},
			"exceed_dscp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maximum_dscp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cos_marking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cos_marking_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exceed_cos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maximum_cos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"overhead": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),
				Optional:     true,
				Computed:     true,
			},
			"exceed_class_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallShaperTrafficShaperCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallShaperTrafficShaper(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallShaperTrafficShaper resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallShaperTrafficShaper(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallShaperTrafficShaper resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallShaperTrafficShaper")
	}

	return resourceFirewallShaperTrafficShaperRead(d, m)
}

func resourceFirewallShaperTrafficShaperUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallShaperTrafficShaper(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallShaperTrafficShaper resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallShaperTrafficShaper(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallShaperTrafficShaper resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallShaperTrafficShaper")
	}

	return resourceFirewallShaperTrafficShaperRead(d, m)
}

func resourceFirewallShaperTrafficShaperDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallShaperTrafficShaper(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallShaperTrafficShaper resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallShaperTrafficShaperRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallShaperTrafficShaper(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallShaperTrafficShaper resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallShaperTrafficShaper(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallShaperTrafficShaper resource from API: %v", err)
	}
	return nil
}

func flattenFirewallShaperTrafficShaperName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperGuaranteedBandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperMaximumBandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperBandwidthUnit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperPerPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperDiffserv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperDiffservcode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperDscpMarkingMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperExceedBandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperExceedDscp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperMaximumDscp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperCosMarking(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperCosMarkingMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperCos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperExceedCos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperMaximumCos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperOverhead(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperExceedClassId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallShaperTrafficShaper(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenFirewallShaperTrafficShaperName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("guaranteed_bandwidth", flattenFirewallShaperTrafficShaperGuaranteedBandwidth(o["guaranteed-bandwidth"], d, "guaranteed_bandwidth", sv)); err != nil {
		if !fortiAPIPatch(o["guaranteed-bandwidth"]) {
			return fmt.Errorf("Error reading guaranteed_bandwidth: %v", err)
		}
	}

	if err = d.Set("maximum_bandwidth", flattenFirewallShaperTrafficShaperMaximumBandwidth(o["maximum-bandwidth"], d, "maximum_bandwidth", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-bandwidth"]) {
			return fmt.Errorf("Error reading maximum_bandwidth: %v", err)
		}
	}

	if err = d.Set("bandwidth_unit", flattenFirewallShaperTrafficShaperBandwidthUnit(o["bandwidth-unit"], d, "bandwidth_unit", sv)); err != nil {
		if !fortiAPIPatch(o["bandwidth-unit"]) {
			return fmt.Errorf("Error reading bandwidth_unit: %v", err)
		}
	}

	if err = d.Set("priority", flattenFirewallShaperTrafficShaperPriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("per_policy", flattenFirewallShaperTrafficShaperPerPolicy(o["per-policy"], d, "per_policy", sv)); err != nil {
		if !fortiAPIPatch(o["per-policy"]) {
			return fmt.Errorf("Error reading per_policy: %v", err)
		}
	}

	if err = d.Set("diffserv", flattenFirewallShaperTrafficShaperDiffserv(o["diffserv"], d, "diffserv", sv)); err != nil {
		if !fortiAPIPatch(o["diffserv"]) {
			return fmt.Errorf("Error reading diffserv: %v", err)
		}
	}

	if err = d.Set("diffservcode", flattenFirewallShaperTrafficShaperDiffservcode(o["diffservcode"], d, "diffservcode", sv)); err != nil {
		if !fortiAPIPatch(o["diffservcode"]) {
			return fmt.Errorf("Error reading diffservcode: %v", err)
		}
	}

	if err = d.Set("dscp_marking_method", flattenFirewallShaperTrafficShaperDscpMarkingMethod(o["dscp-marking-method"], d, "dscp_marking_method", sv)); err != nil {
		if !fortiAPIPatch(o["dscp-marking-method"]) {
			return fmt.Errorf("Error reading dscp_marking_method: %v", err)
		}
	}

	if err = d.Set("exceed_bandwidth", flattenFirewallShaperTrafficShaperExceedBandwidth(o["exceed-bandwidth"], d, "exceed_bandwidth", sv)); err != nil {
		if !fortiAPIPatch(o["exceed-bandwidth"]) {
			return fmt.Errorf("Error reading exceed_bandwidth: %v", err)
		}
	}

	if err = d.Set("exceed_dscp", flattenFirewallShaperTrafficShaperExceedDscp(o["exceed-dscp"], d, "exceed_dscp", sv)); err != nil {
		if !fortiAPIPatch(o["exceed-dscp"]) {
			return fmt.Errorf("Error reading exceed_dscp: %v", err)
		}
	}

	if err = d.Set("maximum_dscp", flattenFirewallShaperTrafficShaperMaximumDscp(o["maximum-dscp"], d, "maximum_dscp", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-dscp"]) {
			return fmt.Errorf("Error reading maximum_dscp: %v", err)
		}
	}

	if err = d.Set("cos_marking", flattenFirewallShaperTrafficShaperCosMarking(o["cos-marking"], d, "cos_marking", sv)); err != nil {
		if !fortiAPIPatch(o["cos-marking"]) {
			return fmt.Errorf("Error reading cos_marking: %v", err)
		}
	}

	if err = d.Set("cos_marking_method", flattenFirewallShaperTrafficShaperCosMarkingMethod(o["cos-marking-method"], d, "cos_marking_method", sv)); err != nil {
		if !fortiAPIPatch(o["cos-marking-method"]) {
			return fmt.Errorf("Error reading cos_marking_method: %v", err)
		}
	}

	if err = d.Set("cos", flattenFirewallShaperTrafficShaperCos(o["cos"], d, "cos", sv)); err != nil {
		if !fortiAPIPatch(o["cos"]) {
			return fmt.Errorf("Error reading cos: %v", err)
		}
	}

	if err = d.Set("exceed_cos", flattenFirewallShaperTrafficShaperExceedCos(o["exceed-cos"], d, "exceed_cos", sv)); err != nil {
		if !fortiAPIPatch(o["exceed-cos"]) {
			return fmt.Errorf("Error reading exceed_cos: %v", err)
		}
	}

	if err = d.Set("maximum_cos", flattenFirewallShaperTrafficShaperMaximumCos(o["maximum-cos"], d, "maximum_cos", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-cos"]) {
			return fmt.Errorf("Error reading maximum_cos: %v", err)
		}
	}

	if err = d.Set("overhead", flattenFirewallShaperTrafficShaperOverhead(o["overhead"], d, "overhead", sv)); err != nil {
		if !fortiAPIPatch(o["overhead"]) {
			return fmt.Errorf("Error reading overhead: %v", err)
		}
	}

	if err = d.Set("exceed_class_id", flattenFirewallShaperTrafficShaperExceedClassId(o["exceed-class-id"], d, "exceed_class_id", sv)); err != nil {
		if !fortiAPIPatch(o["exceed-class-id"]) {
			return fmt.Errorf("Error reading exceed_class_id: %v", err)
		}
	}

	return nil
}

func flattenFirewallShaperTrafficShaperFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallShaperTrafficShaperName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperGuaranteedBandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperMaximumBandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperBandwidthUnit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperPerPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperDiffserv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperDiffservcode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperDscpMarkingMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperExceedBandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperExceedDscp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperMaximumDscp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperCosMarking(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperCosMarkingMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperCos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperExceedCos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperMaximumCos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperOverhead(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperExceedClassId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallShaperTrafficShaper(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallShaperTrafficShaperName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("guaranteed_bandwidth"); ok {
		t, err := expandFirewallShaperTrafficShaperGuaranteedBandwidth(d, v, "guaranteed_bandwidth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guaranteed-bandwidth"] = t
		}
	}

	if v, ok := d.GetOkExists("maximum_bandwidth"); ok {
		t, err := expandFirewallShaperTrafficShaperMaximumBandwidth(d, v, "maximum_bandwidth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_unit"); ok {
		t, err := expandFirewallShaperTrafficShaperBandwidthUnit(d, v, "bandwidth_unit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-unit"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok {
		t, err := expandFirewallShaperTrafficShaperPriority(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("per_policy"); ok {
		t, err := expandFirewallShaperTrafficShaperPerPolicy(d, v, "per_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-policy"] = t
		}
	}

	if v, ok := d.GetOk("diffserv"); ok {
		t, err := expandFirewallShaperTrafficShaperDiffserv(d, v, "diffserv", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode"); ok {
		t, err := expandFirewallShaperTrafficShaperDiffservcode(d, v, "diffservcode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode"] = t
		}
	}

	if v, ok := d.GetOk("dscp_marking_method"); ok {
		t, err := expandFirewallShaperTrafficShaperDscpMarkingMethod(d, v, "dscp_marking_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-marking-method"] = t
		}
	}

	if v, ok := d.GetOkExists("exceed_bandwidth"); ok {
		t, err := expandFirewallShaperTrafficShaperExceedBandwidth(d, v, "exceed_bandwidth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exceed-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("exceed_dscp"); ok {
		t, err := expandFirewallShaperTrafficShaperExceedDscp(d, v, "exceed_dscp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exceed-dscp"] = t
		}
	}

	if v, ok := d.GetOk("maximum_dscp"); ok {
		t, err := expandFirewallShaperTrafficShaperMaximumDscp(d, v, "maximum_dscp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-dscp"] = t
		}
	}

	if v, ok := d.GetOk("cos_marking"); ok {
		t, err := expandFirewallShaperTrafficShaperCosMarking(d, v, "cos_marking", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos-marking"] = t
		}
	}

	if v, ok := d.GetOk("cos_marking_method"); ok {
		t, err := expandFirewallShaperTrafficShaperCosMarkingMethod(d, v, "cos_marking_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos-marking-method"] = t
		}
	}

	if v, ok := d.GetOk("cos"); ok {
		t, err := expandFirewallShaperTrafficShaperCos(d, v, "cos", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos"] = t
		}
	}

	if v, ok := d.GetOk("exceed_cos"); ok {
		t, err := expandFirewallShaperTrafficShaperExceedCos(d, v, "exceed_cos", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exceed-cos"] = t
		}
	}

	if v, ok := d.GetOk("maximum_cos"); ok {
		t, err := expandFirewallShaperTrafficShaperMaximumCos(d, v, "maximum_cos", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-cos"] = t
		}
	}

	if v, ok := d.GetOkExists("overhead"); ok {
		t, err := expandFirewallShaperTrafficShaperOverhead(d, v, "overhead", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overhead"] = t
		}
	}

	if v, ok := d.GetOkExists("exceed_class_id"); ok {
		t, err := expandFirewallShaperTrafficShaperExceedClassId(d, v, "exceed_class_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exceed-class-id"] = t
		}
	}

	return &obj, nil
}
