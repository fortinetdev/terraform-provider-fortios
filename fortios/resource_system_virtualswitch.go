// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure virtual hardware switch interfaces.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemVirtualSwitch() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVirtualSwitchCreate,
		Read:   resourceSystemVirtualSwitchRead,
		Update: resourceSystemVirtualSwitchUpdate,
		Delete: resourceSystemVirtualSwitchDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 15),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"physical_switch": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"alias": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 25),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"span": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"span_source_port": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"span_dest_port": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"span_direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemVirtualSwitchCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemVirtualSwitch(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemVirtualSwitch resource while getting object: %v", err)
	}

	o, err := c.CreateSystemVirtualSwitch(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemVirtualSwitch resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVirtualSwitch")
	}

	return resourceSystemVirtualSwitchRead(d, m)
}

func resourceSystemVirtualSwitchUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemVirtualSwitch(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemVirtualSwitch resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemVirtualSwitch(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemVirtualSwitch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVirtualSwitch")
	}

	return resourceSystemVirtualSwitchRead(d, m)
}

func resourceSystemVirtualSwitchDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemVirtualSwitch(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVirtualSwitch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVirtualSwitchRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemVirtualSwitch(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemVirtualSwitch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVirtualSwitch(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemVirtualSwitch resource from API: %v", err)
	}
	return nil
}

func flattenSystemVirtualSwitchName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualSwitchPhysicalSwitch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualSwitchVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualSwitchPort(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemVirtualSwitchPortName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alias"
		if _, ok := i["alias"]; ok {

			tmp["alias"] = flattenSystemVirtualSwitchPortAlias(i["alias"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemVirtualSwitchPortName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualSwitchPortAlias(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualSwitchSpan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualSwitchSpanSourcePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualSwitchSpanDestPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualSwitchSpanDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemVirtualSwitch(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemVirtualSwitchName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("physical_switch", flattenSystemVirtualSwitchPhysicalSwitch(o["physical-switch"], d, "physical_switch", sv)); err != nil {
		if !fortiAPIPatch(o["physical-switch"]) {
			return fmt.Errorf("Error reading physical_switch: %v", err)
		}
	}

	if err = d.Set("vlan", flattenSystemVirtualSwitchVlan(o["vlan"], d, "vlan", sv)); err != nil {
		if !fortiAPIPatch(o["vlan"]) {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("port", flattenSystemVirtualSwitchPort(o["port"], d, "port", sv)); err != nil {
			if !fortiAPIPatch(o["port"]) {
				return fmt.Errorf("Error reading port: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("port"); ok {
			if err = d.Set("port", flattenSystemVirtualSwitchPort(o["port"], d, "port", sv)); err != nil {
				if !fortiAPIPatch(o["port"]) {
					return fmt.Errorf("Error reading port: %v", err)
				}
			}
		}
	}

	if err = d.Set("span", flattenSystemVirtualSwitchSpan(o["span"], d, "span", sv)); err != nil {
		if !fortiAPIPatch(o["span"]) {
			return fmt.Errorf("Error reading span: %v", err)
		}
	}

	if err = d.Set("span_source_port", flattenSystemVirtualSwitchSpanSourcePort(o["span-source-port"], d, "span_source_port", sv)); err != nil {
		if !fortiAPIPatch(o["span-source-port"]) {
			return fmt.Errorf("Error reading span_source_port: %v", err)
		}
	}

	if err = d.Set("span_dest_port", flattenSystemVirtualSwitchSpanDestPort(o["span-dest-port"], d, "span_dest_port", sv)); err != nil {
		if !fortiAPIPatch(o["span-dest-port"]) {
			return fmt.Errorf("Error reading span_dest_port: %v", err)
		}
	}

	if err = d.Set("span_direction", flattenSystemVirtualSwitchSpanDirection(o["span-direction"], d, "span_direction", sv)); err != nil {
		if !fortiAPIPatch(o["span-direction"]) {
			return fmt.Errorf("Error reading span_direction: %v", err)
		}
	}

	return nil
}

func flattenSystemVirtualSwitchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemVirtualSwitchName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualSwitchPhysicalSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualSwitchVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualSwitchPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemVirtualSwitchPortName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alias"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["alias"], _ = expandSystemVirtualSwitchPortAlias(d, i["alias"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualSwitchPortName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualSwitchPortAlias(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualSwitchSpan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualSwitchSpanSourcePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualSwitchSpanDestPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualSwitchSpanDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVirtualSwitch(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemVirtualSwitchName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("physical_switch"); ok {

		t, err := expandSystemVirtualSwitchPhysicalSwitch(d, v, "physical_switch", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["physical-switch"] = t
		}
	}

	if v, ok := d.GetOkExists("vlan"); ok {

		t, err := expandSystemVirtualSwitchVlan(d, v, "vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {

		t, err := expandSystemVirtualSwitchPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("span"); ok {

		t, err := expandSystemVirtualSwitchSpan(d, v, "span", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["span"] = t
		}
	}

	if v, ok := d.GetOk("span_source_port"); ok {

		t, err := expandSystemVirtualSwitchSpanSourcePort(d, v, "span_source_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["span-source-port"] = t
		}
	}

	if v, ok := d.GetOk("span_dest_port"); ok {

		t, err := expandSystemVirtualSwitchSpanDestPort(d, v, "span_dest_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["span-dest-port"] = t
		}
	}

	if v, ok := d.GetOk("span_direction"); ok {

		t, err := expandSystemVirtualSwitchSpanDirection(d, v, "span_direction", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["span-direction"] = t
		}
	}

	return &obj, nil
}
