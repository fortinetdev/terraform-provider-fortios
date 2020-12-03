// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure switch object tags.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerSwitchInterfaceTag() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSwitchInterfaceTagCreate,
		Read:   resourceSwitchControllerSwitchInterfaceTagRead,
		Update: resourceSwitchControllerSwitchInterfaceTagUpdate,
		Delete: resourceSwitchControllerSwitchInterfaceTagDelete,

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
		},
	}
}

func resourceSwitchControllerSwitchInterfaceTagCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerSwitchInterfaceTag(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSwitchInterfaceTag resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerSwitchInterfaceTag(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSwitchInterfaceTag resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerSwitchInterfaceTag")
	}

	return resourceSwitchControllerSwitchInterfaceTagRead(d, m)
}

func resourceSwitchControllerSwitchInterfaceTagUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerSwitchInterfaceTag(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSwitchInterfaceTag resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerSwitchInterfaceTag(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSwitchInterfaceTag resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerSwitchInterfaceTag")
	}

	return resourceSwitchControllerSwitchInterfaceTagRead(d, m)
}

func resourceSwitchControllerSwitchInterfaceTagDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchControllerSwitchInterfaceTag(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerSwitchInterfaceTag resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSwitchInterfaceTagRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchControllerSwitchInterfaceTag(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSwitchInterfaceTag resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSwitchInterfaceTag(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSwitchInterfaceTag resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSwitchInterfaceTagName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerSwitchInterfaceTag(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerSwitchInterfaceTagName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSwitchInterfaceTagFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerSwitchInterfaceTagName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerSwitchInterfaceTag(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerSwitchInterfaceTagName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
