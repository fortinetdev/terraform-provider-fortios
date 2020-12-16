// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure Type of Service (ToS) based priority table to set network traffic priorities.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemTosBasedPriority() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemTosBasedPriorityCreate,
		Read:   resourceSystemTosBasedPriorityRead,
		Update: resourceSystemTosBasedPriorityUpdate,
		Delete: resourceSystemTosBasedPriorityDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"tos": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemTosBasedPriorityCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemTosBasedPriority(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemTosBasedPriority resource while getting object: %v", err)
	}

	o, err := c.CreateSystemTosBasedPriority(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemTosBasedPriority resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemTosBasedPriority")
	}

	return resourceSystemTosBasedPriorityRead(d, m)
}

func resourceSystemTosBasedPriorityUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemTosBasedPriority(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemTosBasedPriority resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemTosBasedPriority(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemTosBasedPriority resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemTosBasedPriority")
	}

	return resourceSystemTosBasedPriorityRead(d, m)
}

func resourceSystemTosBasedPriorityDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemTosBasedPriority(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemTosBasedPriority resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemTosBasedPriorityRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemTosBasedPriority(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemTosBasedPriority resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemTosBasedPriority(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemTosBasedPriority resource from API: %v", err)
	}
	return nil
}

func flattenSystemTosBasedPriorityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemTosBasedPriorityTos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemTosBasedPriorityPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemTosBasedPriority(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemTosBasedPriorityId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("tos", flattenSystemTosBasedPriorityTos(o["tos"], d, "tos")); err != nil {
		if !fortiAPIPatch(o["tos"]) {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystemTosBasedPriorityPriority(o["priority"], d, "priority")); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	return nil
}

func flattenSystemTosBasedPriorityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemTosBasedPriorityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemTosBasedPriorityTos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemTosBasedPriorityPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemTosBasedPriority(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandSystemTosBasedPriorityId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOkExists("tos"); ok {
		t, err := expandSystemTosBasedPriorityTos(d, v, "tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok {
		t, err := expandSystemTosBasedPriorityPriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	return &obj, nil
}
