// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Hidden table for datasource.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWafMainClass() *schema.Resource {
	return &schema.Resource{
		Create: resourceWafMainClassCreate,
		Read:   resourceWafMainClassRead,
		Update: resourceWafMainClassUpdate,
		Delete: resourceWafMainClassDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceWafMainClassCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWafMainClass(d)
	if err != nil {
		return fmt.Errorf("Error creating WafMainClass resource while getting object: %v", err)
	}

	o, err := c.CreateWafMainClass(obj)

	if err != nil {
		return fmt.Errorf("Error creating WafMainClass resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("WafMainClass")
	}

	return resourceWafMainClassRead(d, m)
}

func resourceWafMainClassUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWafMainClass(d)
	if err != nil {
		return fmt.Errorf("Error updating WafMainClass resource while getting object: %v", err)
	}

	o, err := c.UpdateWafMainClass(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WafMainClass resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("WafMainClass")
	}

	return resourceWafMainClassRead(d, m)
}

func resourceWafMainClassDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWafMainClass(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WafMainClass resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafMainClassRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWafMainClass(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WafMainClass resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafMainClass(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WafMainClass resource from API: %v", err)
	}
	return nil
}

func flattenWafMainClassName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafMainClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWafMainClass(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWafMainClassName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenWafMainClassId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenWafMainClassFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWafMainClassName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafMainClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWafMainClass(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWafMainClassName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandWafMainClassId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
