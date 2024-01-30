// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch Auto-Config custom QoS policy.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerAutoConfigCustom() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerAutoConfigCustomCreate,
		Read:   resourceSwitchControllerAutoConfigCustomRead,
		Update: resourceSwitchControllerAutoConfigCustomUpdate,
		Delete: resourceSwitchControllerAutoConfigCustomDelete,

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
				Required:     true,
			},
			"switch_binding": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"switch_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 16),
							Optional:     true,
							Computed:     true,
						},
						"policy": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
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
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSwitchControllerAutoConfigCustomCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerAutoConfigCustom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerAutoConfigCustom resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerAutoConfigCustom(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerAutoConfigCustom resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerAutoConfigCustom")
	}

	return resourceSwitchControllerAutoConfigCustomRead(d, m)
}

func resourceSwitchControllerAutoConfigCustomUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerAutoConfigCustom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAutoConfigCustom resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerAutoConfigCustom(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAutoConfigCustom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerAutoConfigCustom")
	}

	return resourceSwitchControllerAutoConfigCustomRead(d, m)
}

func resourceSwitchControllerAutoConfigCustomDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerAutoConfigCustom(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerAutoConfigCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerAutoConfigCustomRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerAutoConfigCustom(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAutoConfigCustom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerAutoConfigCustom(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAutoConfigCustom resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerAutoConfigCustomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerAutoConfigCustomSwitchBinding(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_id"
		if cur_v, ok := i["switch-id"]; ok {
			tmp["switch_id"] = flattenSwitchControllerAutoConfigCustomSwitchBindingSwitchId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policy"
		if cur_v, ok := i["policy"]; ok {
			tmp["policy"] = flattenSwitchControllerAutoConfigCustomSwitchBindingPolicy(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "switch_id", d)
	return result
}

func flattenSwitchControllerAutoConfigCustomSwitchBindingSwitchId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerAutoConfigCustomSwitchBindingPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerAutoConfigCustom(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenSwitchControllerAutoConfigCustomName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("switch_binding", flattenSwitchControllerAutoConfigCustomSwitchBinding(o["switch-binding"], d, "switch_binding", sv)); err != nil {
			if !fortiAPIPatch(o["switch-binding"]) {
				return fmt.Errorf("Error reading switch_binding: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("switch_binding"); ok {
			if err = d.Set("switch_binding", flattenSwitchControllerAutoConfigCustomSwitchBinding(o["switch-binding"], d, "switch_binding", sv)); err != nil {
				if !fortiAPIPatch(o["switch-binding"]) {
					return fmt.Errorf("Error reading switch_binding: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerAutoConfigCustomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerAutoConfigCustomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAutoConfigCustomSwitchBinding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["switch-id"], _ = expandSwitchControllerAutoConfigCustomSwitchBindingSwitchId(d, i["switch_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policy"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["policy"], _ = expandSwitchControllerAutoConfigCustomSwitchBindingPolicy(d, i["policy"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerAutoConfigCustomSwitchBindingSwitchId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAutoConfigCustomSwitchBindingPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerAutoConfigCustom(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerAutoConfigCustomName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("switch_binding"); ok || d.HasChange("switch_binding") {
		t, err := expandSwitchControllerAutoConfigCustomSwitchBinding(d, v, "switch_binding", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-binding"] = t
		}
	}

	return &obj, nil
}
