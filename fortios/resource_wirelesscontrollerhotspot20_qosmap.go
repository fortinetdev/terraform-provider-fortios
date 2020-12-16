// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure QoS map set.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerHotspot20QosMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20QosMapCreate,
		Read:   resourceWirelessControllerHotspot20QosMapRead,
		Update: resourceWirelessControllerHotspot20QosMapUpdate,
		Delete: resourceWirelessControllerHotspot20QosMapDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"dscp_except": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"index": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 21),
							Optional:     true,
							Computed:     true,
						},
						"dscp": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"up": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dscp_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"index": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 8),
							Optional:     true,
							Computed:     true,
						},
						"up": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"low": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"high": &schema.Schema{
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

func resourceWirelessControllerHotspot20QosMapCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerHotspot20QosMap(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20QosMap resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerHotspot20QosMap(obj)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20QosMap resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20QosMap")
	}

	return resourceWirelessControllerHotspot20QosMapRead(d, m)
}

func resourceWirelessControllerHotspot20QosMapUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerHotspot20QosMap(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20QosMap resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerHotspot20QosMap(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20QosMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20QosMap")
	}

	return resourceWirelessControllerHotspot20QosMapRead(d, m)
}

func resourceWirelessControllerHotspot20QosMapDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWirelessControllerHotspot20QosMap(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20QosMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20QosMapRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWirelessControllerHotspot20QosMap(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20QosMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20QosMap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20QosMap resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20QosMapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20QosMapDscpExcept(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			tmp["index"] = flattenWirelessControllerHotspot20QosMapDscpExceptIndex(i["index"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp"
		if _, ok := i["dscp"]; ok {
			tmp["dscp"] = flattenWirelessControllerHotspot20QosMapDscpExceptDscp(i["dscp"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "up"
		if _, ok := i["up"]; ok {
			tmp["up"] = flattenWirelessControllerHotspot20QosMapDscpExceptUp(i["up"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerHotspot20QosMapDscpExceptIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20QosMapDscpExceptDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20QosMapDscpExceptUp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20QosMapDscpRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			tmp["index"] = flattenWirelessControllerHotspot20QosMapDscpRangeIndex(i["index"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "up"
		if _, ok := i["up"]; ok {
			tmp["up"] = flattenWirelessControllerHotspot20QosMapDscpRangeUp(i["up"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "low"
		if _, ok := i["low"]; ok {
			tmp["low"] = flattenWirelessControllerHotspot20QosMapDscpRangeLow(i["low"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "high"
		if _, ok := i["high"]; ok {
			tmp["high"] = flattenWirelessControllerHotspot20QosMapDscpRangeHigh(i["high"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerHotspot20QosMapDscpRangeIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20QosMapDscpRangeUp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20QosMapDscpRangeLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20QosMapDscpRangeHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20QosMap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerHotspot20QosMapName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dscp_except", flattenWirelessControllerHotspot20QosMapDscpExcept(o["dscp-except"], d, "dscp_except")); err != nil {
			if !fortiAPIPatch(o["dscp-except"]) {
				return fmt.Errorf("Error reading dscp_except: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dscp_except"); ok {
			if err = d.Set("dscp_except", flattenWirelessControllerHotspot20QosMapDscpExcept(o["dscp-except"], d, "dscp_except")); err != nil {
				if !fortiAPIPatch(o["dscp-except"]) {
					return fmt.Errorf("Error reading dscp_except: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dscp_range", flattenWirelessControllerHotspot20QosMapDscpRange(o["dscp-range"], d, "dscp_range")); err != nil {
			if !fortiAPIPatch(o["dscp-range"]) {
				return fmt.Errorf("Error reading dscp_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dscp_range"); ok {
			if err = d.Set("dscp_range", flattenWirelessControllerHotspot20QosMapDscpRange(o["dscp-range"], d, "dscp_range")); err != nil {
				if !fortiAPIPatch(o["dscp-range"]) {
					return fmt.Errorf("Error reading dscp_range: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20QosMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20QosMapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20QosMapDscpExcept(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["index"], _ = expandWirelessControllerHotspot20QosMapDscpExceptIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp"], _ = expandWirelessControllerHotspot20QosMapDscpExceptDscp(d, i["dscp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "up"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["up"], _ = expandWirelessControllerHotspot20QosMapDscpExceptUp(d, i["up"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20QosMapDscpExceptIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20QosMapDscpExceptDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20QosMapDscpExceptUp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20QosMapDscpRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["index"], _ = expandWirelessControllerHotspot20QosMapDscpRangeIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "up"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["up"], _ = expandWirelessControllerHotspot20QosMapDscpRangeUp(d, i["up"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "low"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["low"], _ = expandWirelessControllerHotspot20QosMapDscpRangeLow(d, i["low"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "high"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["high"], _ = expandWirelessControllerHotspot20QosMapDscpRangeHigh(d, i["high"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20QosMapDscpRangeIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20QosMapDscpRangeUp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20QosMapDscpRangeLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20QosMapDscpRangeHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20QosMap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerHotspot20QosMapName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("dscp_except"); ok {
		t, err := expandWirelessControllerHotspot20QosMapDscpExcept(d, v, "dscp_except")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-except"] = t
		}
	}

	if v, ok := d.GetOk("dscp_range"); ok {
		t, err := expandWirelessControllerHotspot20QosMapDscpRange(d, v, "dscp_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-range"] = t
		}
	}

	return &obj, nil
}
