// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure virtual pool.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerVirtualPortPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerVirtualPortPoolCreate,
		Read:   resourceSwitchControllerVirtualPortPoolRead,
		Update: resourceSwitchControllerVirtualPortPoolUpdate,
		Delete: resourceSwitchControllerVirtualPortPoolDelete,

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
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSwitchControllerVirtualPortPoolCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerVirtualPortPool(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerVirtualPortPool resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerVirtualPortPool(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerVirtualPortPool resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerVirtualPortPool")
	}

	return resourceSwitchControllerVirtualPortPoolRead(d, m)
}

func resourceSwitchControllerVirtualPortPoolUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerVirtualPortPool(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerVirtualPortPool resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerVirtualPortPool(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerVirtualPortPool resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerVirtualPortPool")
	}

	return resourceSwitchControllerVirtualPortPoolRead(d, m)
}

func resourceSwitchControllerVirtualPortPoolDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchControllerVirtualPortPool(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerVirtualPortPool resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerVirtualPortPoolRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchControllerVirtualPortPool(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerVirtualPortPool resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerVirtualPortPool(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerVirtualPortPool resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerVirtualPortPoolName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerVirtualPortPoolDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerVirtualPortPool(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerVirtualPortPoolName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerVirtualPortPoolDescription(o["description"], d, "description")); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerVirtualPortPoolFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerVirtualPortPoolName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVirtualPortPoolDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerVirtualPortPool(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerVirtualPortPoolName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandSwitchControllerVirtualPortPoolDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	return &obj, nil
}
