// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch multiple spanning tree protocol (MSTP) instances.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerStpInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerStpInstanceCreate,
		Read:   resourceSwitchControllerStpInstanceRead,
		Update: resourceSwitchControllerStpInstanceUpdate,
		Delete: resourceSwitchControllerStpInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"vlan_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
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

func resourceSwitchControllerStpInstanceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerStpInstance(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerStpInstance resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerStpInstance(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerStpInstance resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerStpInstance")
	}

	return resourceSwitchControllerStpInstanceRead(d, m)
}

func resourceSwitchControllerStpInstanceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerStpInstance(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerStpInstance resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerStpInstance(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerStpInstance resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerStpInstance")
	}

	return resourceSwitchControllerStpInstanceRead(d, m)
}

func resourceSwitchControllerStpInstanceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerStpInstance(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerStpInstance resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerStpInstanceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerStpInstance(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerStpInstance resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerStpInstance(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerStpInstance resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerStpInstanceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerStpInstanceVlanRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := i["vlan-name"]; ok {

			tmp["vlan_name"] = flattenSwitchControllerStpInstanceVlanRangeVlanName(i["vlan-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vlan_name", d)
	return result
}

func flattenSwitchControllerStpInstanceVlanRangeVlanName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerStpInstance(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenSwitchControllerStpInstanceId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vlan_range", flattenSwitchControllerStpInstanceVlanRange(o["vlan-range"], d, "vlan_range", sv)); err != nil {
			if !fortiAPIPatch(o["vlan-range"]) {
				return fmt.Errorf("Error reading vlan_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vlan_range"); ok {
			if err = d.Set("vlan_range", flattenSwitchControllerStpInstanceVlanRange(o["vlan-range"], d, "vlan_range", sv)); err != nil {
				if !fortiAPIPatch(o["vlan-range"]) {
					return fmt.Errorf("Error reading vlan_range: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerStpInstanceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerStpInstanceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpInstanceVlanRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["vlan-name"], _ = expandSwitchControllerStpInstanceVlanRangeVlanName(d, i["vlan_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerStpInstanceVlanRangeVlanName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerStpInstance(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok {

		t, err := expandSwitchControllerStpInstanceId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("vlan_range"); ok {

		t, err := expandSwitchControllerStpInstanceVlanRange(d, v, "vlan_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-range"] = t
		}
	}

	return &obj, nil
}
