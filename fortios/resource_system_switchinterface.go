// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure software switch interfaces by grouping physical and WiFi interfaces.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemSwitchInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSwitchInterfaceCreate,
		Read:   resourceSystemSwitchInterfaceRead,
		Update: resourceSystemSwitchInterfaceUpdate,
		Delete: resourceSystemSwitchInterfaceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				ForceNew:     true,
				Required:     true,
			},
			"vdom": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"span_dest_port": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"span_source_port": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"member": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"intra_switch_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 8640000),
				Optional:     true,
				Computed:     true,
			},
			"span": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceSystemSwitchInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemSwitchInterface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSwitchInterface resource while getting object: %v", err)
	}

	o, err := c.CreateSystemSwitchInterface(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemSwitchInterface resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSwitchInterface")
	}

	return resourceSystemSwitchInterfaceRead(d, m)
}

func resourceSystemSwitchInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemSwitchInterface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSwitchInterface resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSwitchInterface(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemSwitchInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSwitchInterface")
	}

	return resourceSystemSwitchInterfaceRead(d, m)
}

func resourceSystemSwitchInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemSwitchInterface(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSwitchInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSwitchInterfaceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemSwitchInterface(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemSwitchInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSwitchInterface(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSwitchInterface resource from API: %v", err)
	}
	return nil
}

func flattenSystemSwitchInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSwitchInterfaceVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSwitchInterfaceSpanDestPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSwitchInterfaceSpanSourcePort(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := i["interface-name"]; ok {

			tmp["interface_name"] = flattenSystemSwitchInterfaceSpanSourcePortInterfaceName(i["interface-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "interface_name", d)
	return result
}

func flattenSystemSwitchInterfaceSpanSourcePortInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSwitchInterfaceMember(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := i["interface-name"]; ok {

			tmp["interface_name"] = flattenSystemSwitchInterfaceMemberInterfaceName(i["interface-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "interface_name", d)
	return result
}

func flattenSystemSwitchInterfaceMemberInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSwitchInterfaceType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSwitchInterfaceIntraSwitchPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSwitchInterfaceMacTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSwitchInterfaceSpan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSwitchInterfaceSpanDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSwitchInterface(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemSwitchInterfaceName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemSwitchInterfaceVdom(o["vdom"], d, "vdom", sv)); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("span_dest_port", flattenSystemSwitchInterfaceSpanDestPort(o["span-dest-port"], d, "span_dest_port", sv)); err != nil {
		if !fortiAPIPatch(o["span-dest-port"]) {
			return fmt.Errorf("Error reading span_dest_port: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("span_source_port", flattenSystemSwitchInterfaceSpanSourcePort(o["span-source-port"], d, "span_source_port", sv)); err != nil {
			if !fortiAPIPatch(o["span-source-port"]) {
				return fmt.Errorf("Error reading span_source_port: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("span_source_port"); ok {
			if err = d.Set("span_source_port", flattenSystemSwitchInterfaceSpanSourcePort(o["span-source-port"], d, "span_source_port", sv)); err != nil {
				if !fortiAPIPatch(o["span-source-port"]) {
					return fmt.Errorf("Error reading span_source_port: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("member", flattenSystemSwitchInterfaceMember(o["member"], d, "member", sv)); err != nil {
			if !fortiAPIPatch(o["member"]) {
				return fmt.Errorf("Error reading member: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("member"); ok {
			if err = d.Set("member", flattenSystemSwitchInterfaceMember(o["member"], d, "member", sv)); err != nil {
				if !fortiAPIPatch(o["member"]) {
					return fmt.Errorf("Error reading member: %v", err)
				}
			}
		}
	}

	if err = d.Set("type", flattenSystemSwitchInterfaceType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("intra_switch_policy", flattenSystemSwitchInterfaceIntraSwitchPolicy(o["intra-switch-policy"], d, "intra_switch_policy", sv)); err != nil {
		if !fortiAPIPatch(o["intra-switch-policy"]) {
			return fmt.Errorf("Error reading intra_switch_policy: %v", err)
		}
	}

	if err = d.Set("mac_ttl", flattenSystemSwitchInterfaceMacTtl(o["mac-ttl"], d, "mac_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["mac-ttl"]) {
			return fmt.Errorf("Error reading mac_ttl: %v", err)
		}
	}

	if err = d.Set("span", flattenSystemSwitchInterfaceSpan(o["span"], d, "span", sv)); err != nil {
		if !fortiAPIPatch(o["span"]) {
			return fmt.Errorf("Error reading span: %v", err)
		}
	}

	if err = d.Set("span_direction", flattenSystemSwitchInterfaceSpanDirection(o["span-direction"], d, "span_direction", sv)); err != nil {
		if !fortiAPIPatch(o["span-direction"]) {
			return fmt.Errorf("Error reading span_direction: %v", err)
		}
	}

	return nil
}

func flattenSystemSwitchInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSwitchInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSwitchInterfaceVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSwitchInterfaceSpanDestPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSwitchInterfaceSpanSourcePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["interface-name"], _ = expandSystemSwitchInterfaceSpanSourcePortInterfaceName(d, i["interface_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSwitchInterfaceSpanSourcePortInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSwitchInterfaceMember(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["interface-name"], _ = expandSystemSwitchInterfaceMemberInterfaceName(d, i["interface_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSwitchInterfaceMemberInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSwitchInterfaceType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSwitchInterfaceIntraSwitchPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSwitchInterfaceMacTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSwitchInterfaceSpan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSwitchInterfaceSpanDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSwitchInterface(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemSwitchInterfaceName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok {

		t, err := expandSystemSwitchInterfaceVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("span_dest_port"); ok {

		t, err := expandSystemSwitchInterfaceSpanDestPort(d, v, "span_dest_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["span-dest-port"] = t
		}
	}

	if v, ok := d.GetOk("span_source_port"); ok {

		t, err := expandSystemSwitchInterfaceSpanSourcePort(d, v, "span_source_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["span-source-port"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok {

		t, err := expandSystemSwitchInterfaceMember(d, v, "member", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {

		t, err := expandSystemSwitchInterfaceType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("intra_switch_policy"); ok {

		t, err := expandSystemSwitchInterfaceIntraSwitchPolicy(d, v, "intra_switch_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intra-switch-policy"] = t
		}
	}

	if v, ok := d.GetOk("mac_ttl"); ok {

		t, err := expandSystemSwitchInterfaceMacTtl(d, v, "mac_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-ttl"] = t
		}
	}

	if v, ok := d.GetOk("span"); ok {

		t, err := expandSystemSwitchInterfaceSpan(d, v, "span", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["span"] = t
		}
	}

	if v, ok := d.GetOk("span_direction"); ok {

		t, err := expandSystemSwitchInterfaceSpanDirection(d, v, "span_direction", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["span-direction"] = t
		}
	}

	return &obj, nil
}
