// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure group of Internet Service.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallInternetServiceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInternetServiceGroupCreate,
		Read:   resourceFirewallInternetServiceGroupRead,
		Update: resourceFirewallInternetServiceGroupUpdate,
		Delete: resourceFirewallInternetServiceGroupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required:     true,
				ForceNew:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceFirewallInternetServiceGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallInternetServiceGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceGroup resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallInternetServiceGroup(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceGroup resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallInternetServiceGroup")
	}

	return resourceFirewallInternetServiceGroupRead(d, m)
}

func resourceFirewallInternetServiceGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallInternetServiceGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceGroup resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallInternetServiceGroup(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallInternetServiceGroup")
	}

	return resourceFirewallInternetServiceGroupRead(d, m)
}

func resourceFirewallInternetServiceGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallInternetServiceGroup(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInternetServiceGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallInternetServiceGroup(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInternetServiceGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceGroup resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInternetServiceGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceGroupComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceGroupDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceGroupMember(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallInternetServiceGroupMemberId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallInternetServiceGroupMemberId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallInternetServiceGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenFirewallInternetServiceGroupName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallInternetServiceGroupComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("direction", flattenFirewallInternetServiceGroupDirection(o["direction"], d, "direction")); err != nil {
		if !fortiAPIPatch(o["direction"]) {
			return fmt.Errorf("Error reading direction: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("member", flattenFirewallInternetServiceGroupMember(o["member"], d, "member")); err != nil {
			if !fortiAPIPatch(o["member"]) {
				return fmt.Errorf("Error reading member: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("member"); ok {
			if err = d.Set("member", flattenFirewallInternetServiceGroupMember(o["member"], d, "member")); err != nil {
				if !fortiAPIPatch(o["member"]) {
					return fmt.Errorf("Error reading member: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFirewallInternetServiceGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallInternetServiceGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceGroupComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceGroupDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceGroupMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallInternetServiceGroupMemberId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceGroupMemberId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInternetServiceGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallInternetServiceGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandFirewallInternetServiceGroupComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("direction"); ok {
		t, err := expandFirewallInternetServiceGroupDirection(d, v, "direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["direction"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok {
		t, err := expandFirewallInternetServiceGroupMember(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	return &obj, nil
}
