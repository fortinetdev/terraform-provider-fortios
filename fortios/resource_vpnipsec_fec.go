// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Forward Error Correction (FEC) mapping profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnIpsecFec() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnIpsecFecCreate,
		Read:   resourceVpnIpsecFecRead,
		Update: resourceVpnIpsecFecUpdate,
		Delete: resourceVpnIpsecFecDelete,

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
			"mappings": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seqno": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"base": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 20),
							Optional:     true,
							Computed:     true,
						},
						"redundant": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 5),
							Optional:     true,
							Computed:     true,
						},
						"packet_loss_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),
							Optional:     true,
							Computed:     true,
						},
						"latency_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"bandwidth_up_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"bandwidth_down_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"bandwidth_bi_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceVpnIpsecFecCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnIpsecFec(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecFec resource while getting object: %v", err)
	}

	o, err := c.CreateVpnIpsecFec(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecFec resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecFec")
	}

	return resourceVpnIpsecFecRead(d, m)
}

func resourceVpnIpsecFecUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnIpsecFec(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecFec resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnIpsecFec(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecFec resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecFec")
	}

	return resourceVpnIpsecFecRead(d, m)
}

func resourceVpnIpsecFecDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteVpnIpsecFec(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting VpnIpsecFec resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecFecRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadVpnIpsecFec(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecFec resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnIpsecFec(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecFec resource from API: %v", err)
	}
	return nil
}

func flattenVpnIpsecFecName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecFecMappings(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seqno"
		if _, ok := i["seqno"]; ok {

			tmp["seqno"] = flattenVpnIpsecFecMappingsSeqno(i["seqno"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base"
		if _, ok := i["base"]; ok {

			tmp["base"] = flattenVpnIpsecFecMappingsBase(i["base"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redundant"
		if _, ok := i["redundant"]; ok {

			tmp["redundant"] = flattenVpnIpsecFecMappingsRedundant(i["redundant"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_threshold"
		if _, ok := i["packet-loss-threshold"]; ok {

			tmp["packet_loss_threshold"] = flattenVpnIpsecFecMappingsPacketLossThreshold(i["packet-loss-threshold"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := i["latency-threshold"]; ok {

			tmp["latency_threshold"] = flattenVpnIpsecFecMappingsLatencyThreshold(i["latency-threshold"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_up_threshold"
		if _, ok := i["bandwidth-up-threshold"]; ok {

			tmp["bandwidth_up_threshold"] = flattenVpnIpsecFecMappingsBandwidthUpThreshold(i["bandwidth-up-threshold"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_down_threshold"
		if _, ok := i["bandwidth-down-threshold"]; ok {

			tmp["bandwidth_down_threshold"] = flattenVpnIpsecFecMappingsBandwidthDownThreshold(i["bandwidth-down-threshold"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_bi_threshold"
		if _, ok := i["bandwidth-bi-threshold"]; ok {

			tmp["bandwidth_bi_threshold"] = flattenVpnIpsecFecMappingsBandwidthBiThreshold(i["bandwidth-bi-threshold"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "seqno", d)
	return result
}

func flattenVpnIpsecFecMappingsSeqno(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsBase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsRedundant(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsPacketLossThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsLatencyThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsBandwidthUpThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsBandwidthDownThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecFecMappingsBandwidthBiThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnIpsecFec(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenVpnIpsecFecName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("mappings", flattenVpnIpsecFecMappings(o["mappings"], d, "mappings", sv)); err != nil {
			if !fortiAPIPatch(o["mappings"]) {
				return fmt.Errorf("Error reading mappings: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mappings"); ok {
			if err = d.Set("mappings", flattenVpnIpsecFecMappings(o["mappings"], d, "mappings", sv)); err != nil {
				if !fortiAPIPatch(o["mappings"]) {
					return fmt.Errorf("Error reading mappings: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenVpnIpsecFecFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnIpsecFecName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappings(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seqno"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["seqno"], _ = expandVpnIpsecFecMappingsSeqno(d, i["seqno"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["base"], _ = expandVpnIpsecFecMappingsBase(d, i["base"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "redundant"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["redundant"], _ = expandVpnIpsecFecMappingsRedundant(d, i["redundant"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_threshold"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["packet-loss-threshold"], _ = expandVpnIpsecFecMappingsPacketLossThreshold(d, i["packet_loss_threshold"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["latency-threshold"], _ = expandVpnIpsecFecMappingsLatencyThreshold(d, i["latency_threshold"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_up_threshold"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["bandwidth-up-threshold"], _ = expandVpnIpsecFecMappingsBandwidthUpThreshold(d, i["bandwidth_up_threshold"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_down_threshold"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["bandwidth-down-threshold"], _ = expandVpnIpsecFecMappingsBandwidthDownThreshold(d, i["bandwidth_down_threshold"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_bi_threshold"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["bandwidth-bi-threshold"], _ = expandVpnIpsecFecMappingsBandwidthBiThreshold(d, i["bandwidth_bi_threshold"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnIpsecFecMappingsSeqno(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsBase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsRedundant(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsPacketLossThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsLatencyThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsBandwidthUpThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsBandwidthDownThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecFecMappingsBandwidthBiThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnIpsecFec(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandVpnIpsecFecName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("mappings"); ok || d.HasChange("mappings") {

		t, err := expandVpnIpsecFecMappings(d, v, "mappings", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappings"] = t
		}
	}

	return &obj, nil
}
