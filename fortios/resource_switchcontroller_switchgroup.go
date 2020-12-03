// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch switch groups.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerSwitchGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSwitchGroupCreate,
		Read:   resourceSwitchControllerSwitchGroupRead,
		Update: resourceSwitchControllerSwitchGroupUpdate,
		Delete: resourceSwitchControllerSwitchGroupDelete,

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
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"members": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceSwitchControllerSwitchGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerSwitchGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSwitchGroup resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerSwitchGroup(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSwitchGroup resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerSwitchGroup")
	}

	return resourceSwitchControllerSwitchGroupRead(d, m)
}

func resourceSwitchControllerSwitchGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerSwitchGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSwitchGroup resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerSwitchGroup(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSwitchGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerSwitchGroup")
	}

	return resourceSwitchControllerSwitchGroupRead(d, m)
}

func resourceSwitchControllerSwitchGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchControllerSwitchGroup(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerSwitchGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSwitchGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchControllerSwitchGroup(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSwitchGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSwitchGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSwitchGroup resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSwitchGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSwitchGroupDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSwitchGroupMembers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenSwitchControllerSwitchGroupMembersName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSwitchControllerSwitchGroupMembersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerSwitchGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerSwitchGroupName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerSwitchGroupDescription(o["description"], d, "description")); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("members", flattenSwitchControllerSwitchGroupMembers(o["members"], d, "members")); err != nil {
			if !fortiAPIPatch(o["members"]) {
				return fmt.Errorf("Error reading members: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("members"); ok {
			if err = d.Set("members", flattenSwitchControllerSwitchGroupMembers(o["members"], d, "members")); err != nil {
				if !fortiAPIPatch(o["members"]) {
					return fmt.Errorf("Error reading members: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerSwitchGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerSwitchGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSwitchGroupDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSwitchGroupMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSwitchControllerSwitchGroupMembersName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerSwitchGroupMembersName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerSwitchGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerSwitchGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandSwitchControllerSwitchGroupDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("members"); ok {
		t, err := expandSwitchControllerSwitchGroupMembers(d, v, "members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["members"] = t
		}
	}

	return &obj, nil
}
