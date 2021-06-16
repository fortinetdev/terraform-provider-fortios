// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch QoS 802.1p.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerQosDot1PMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerQosDot1PMapCreate,
		Read:   resourceSwitchControllerQosDot1PMapRead,
		Update: resourceSwitchControllerQosDot1PMapUpdate,
		Delete: resourceSwitchControllerQosDot1PMapDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Required:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"egress_pri_tagging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_0": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerQosDot1PMapCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerQosDot1PMap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosDot1PMap resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerQosDot1PMap(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosDot1PMap resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerQosDot1PMap")
	}

	return resourceSwitchControllerQosDot1PMapRead(d, m)
}

func resourceSwitchControllerQosDot1PMapUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerQosDot1PMap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosDot1PMap resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerQosDot1PMap(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosDot1PMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerQosDot1PMap")
	}

	return resourceSwitchControllerQosDot1PMapRead(d, m)
}

func resourceSwitchControllerQosDot1PMapDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerQosDot1PMap(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerQosDot1PMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerQosDot1PMapRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerQosDot1PMap(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosDot1PMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerQosDot1PMap(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosDot1PMap resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerQosDot1PMapName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapEgressPriTagging(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapPriority0(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapPriority1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapPriority2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapPriority3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapPriority4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapPriority5(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapPriority6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosDot1PMapPriority7(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerQosDot1PMap(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerQosDot1PMapName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerQosDot1PMapDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("egress_pri_tagging", flattenSwitchControllerQosDot1PMapEgressPriTagging(o["egress-pri-tagging"], d, "egress_pri_tagging", sv)); err != nil {
		if !fortiAPIPatch(o["egress-pri-tagging"]) {
			return fmt.Errorf("Error reading egress_pri_tagging: %v", err)
		}
	}

	if err = d.Set("priority_0", flattenSwitchControllerQosDot1PMapPriority0(o["priority-0"], d, "priority_0", sv)); err != nil {
		if !fortiAPIPatch(o["priority-0"]) {
			return fmt.Errorf("Error reading priority_0: %v", err)
		}
	}

	if err = d.Set("priority_1", flattenSwitchControllerQosDot1PMapPriority1(o["priority-1"], d, "priority_1", sv)); err != nil {
		if !fortiAPIPatch(o["priority-1"]) {
			return fmt.Errorf("Error reading priority_1: %v", err)
		}
	}

	if err = d.Set("priority_2", flattenSwitchControllerQosDot1PMapPriority2(o["priority-2"], d, "priority_2", sv)); err != nil {
		if !fortiAPIPatch(o["priority-2"]) {
			return fmt.Errorf("Error reading priority_2: %v", err)
		}
	}

	if err = d.Set("priority_3", flattenSwitchControllerQosDot1PMapPriority3(o["priority-3"], d, "priority_3", sv)); err != nil {
		if !fortiAPIPatch(o["priority-3"]) {
			return fmt.Errorf("Error reading priority_3: %v", err)
		}
	}

	if err = d.Set("priority_4", flattenSwitchControllerQosDot1PMapPriority4(o["priority-4"], d, "priority_4", sv)); err != nil {
		if !fortiAPIPatch(o["priority-4"]) {
			return fmt.Errorf("Error reading priority_4: %v", err)
		}
	}

	if err = d.Set("priority_5", flattenSwitchControllerQosDot1PMapPriority5(o["priority-5"], d, "priority_5", sv)); err != nil {
		if !fortiAPIPatch(o["priority-5"]) {
			return fmt.Errorf("Error reading priority_5: %v", err)
		}
	}

	if err = d.Set("priority_6", flattenSwitchControllerQosDot1PMapPriority6(o["priority-6"], d, "priority_6", sv)); err != nil {
		if !fortiAPIPatch(o["priority-6"]) {
			return fmt.Errorf("Error reading priority_6: %v", err)
		}
	}

	if err = d.Set("priority_7", flattenSwitchControllerQosDot1PMapPriority7(o["priority-7"], d, "priority_7", sv)); err != nil {
		if !fortiAPIPatch(o["priority-7"]) {
			return fmt.Errorf("Error reading priority_7: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerQosDot1PMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerQosDot1PMapName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapEgressPriTagging(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapPriority0(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapPriority1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapPriority2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapPriority3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapPriority4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapPriority5(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapPriority6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosDot1PMapPriority7(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerQosDot1PMap(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchControllerQosDot1PMapName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSwitchControllerQosDot1PMapDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("egress_pri_tagging"); ok {

		t, err := expandSwitchControllerQosDot1PMapEgressPriTagging(d, v, "egress_pri_tagging", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["egress-pri-tagging"] = t
		}
	}

	if v, ok := d.GetOk("priority_0"); ok {

		t, err := expandSwitchControllerQosDot1PMapPriority0(d, v, "priority_0", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-0"] = t
		}
	}

	if v, ok := d.GetOk("priority_1"); ok {

		t, err := expandSwitchControllerQosDot1PMapPriority1(d, v, "priority_1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-1"] = t
		}
	}

	if v, ok := d.GetOk("priority_2"); ok {

		t, err := expandSwitchControllerQosDot1PMapPriority2(d, v, "priority_2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-2"] = t
		}
	}

	if v, ok := d.GetOk("priority_3"); ok {

		t, err := expandSwitchControllerQosDot1PMapPriority3(d, v, "priority_3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-3"] = t
		}
	}

	if v, ok := d.GetOk("priority_4"); ok {

		t, err := expandSwitchControllerQosDot1PMapPriority4(d, v, "priority_4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-4"] = t
		}
	}

	if v, ok := d.GetOk("priority_5"); ok {

		t, err := expandSwitchControllerQosDot1PMapPriority5(d, v, "priority_5", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-5"] = t
		}
	}

	if v, ok := d.GetOk("priority_6"); ok {

		t, err := expandSwitchControllerQosDot1PMapPriority6(d, v, "priority_6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-6"] = t
		}
	}

	if v, ok := d.GetOk("priority_7"); ok {

		t, err := expandSwitchControllerQosDot1PMapPriority7(d, v, "priority_7", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-7"] = t
		}
	}

	return &obj, nil
}
