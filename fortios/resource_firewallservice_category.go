// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure service categories.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallServiceCategory() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallServiceCategoryCreate,
		Read:   resourceFirewallServiceCategoryRead,
		Update: resourceFirewallServiceCategoryUpdate,
		Delete: resourceFirewallServiceCategoryDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
		},
	}
}

func resourceFirewallServiceCategoryCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallServiceCategory(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallServiceCategory resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallServiceCategory(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallServiceCategory resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallServiceCategory")
	}

	return resourceFirewallServiceCategoryRead(d, m)
}

func resourceFirewallServiceCategoryUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallServiceCategory(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallServiceCategory resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallServiceCategory(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallServiceCategory resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallServiceCategory")
	}

	return resourceFirewallServiceCategoryRead(d, m)
}

func resourceFirewallServiceCategoryDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallServiceCategory(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallServiceCategory resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallServiceCategoryRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallServiceCategory(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallServiceCategory resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallServiceCategory(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallServiceCategory resource from API: %v", err)
	}
	return nil
}

func flattenFirewallServiceCategoryName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCategoryComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallServiceCategory(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenFirewallServiceCategoryName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallServiceCategoryComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	return nil
}

func flattenFirewallServiceCategoryFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallServiceCategoryName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCategoryComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallServiceCategory(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallServiceCategoryName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandFirewallServiceCategoryComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	return &obj, nil
}
