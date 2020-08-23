// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure external resource.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemExternalResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemExternalResourceCreate,
		Read:   resourceSystemExternalResourceRead,
		Update: resourceSystemExternalResourceUpdate,
		Delete: resourceSystemExternalResourceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required: true,
			},
			"status": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"category": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(192, 221),
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional: true,
			},
			"comments": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional: true,
			},
			"resource": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Required: true,
			},
			"refresh_rate": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 43200),
				Required: true,
			},
		},
	}
}

func resourceSystemExternalResourceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemExternalResource(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemExternalResource resource while getting object: %v", err)
	}

	o, err := c.CreateSystemExternalResource(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemExternalResource resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemExternalResource")
	}

	return resourceSystemExternalResourceRead(d, m)
}

func resourceSystemExternalResourceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemExternalResource(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemExternalResource resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemExternalResource(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemExternalResource resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemExternalResource")
	}

	return resourceSystemExternalResourceRead(d, m)
}

func resourceSystemExternalResourceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemExternalResource(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemExternalResource resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemExternalResourceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemExternalResource(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemExternalResource resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemExternalResource(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemExternalResource resource from API: %v", err)
	}
	return nil
}


func flattenSystemExternalResourceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourcePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceResource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemExternalResourceRefreshRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectSystemExternalResource(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("name", flattenSystemExternalResourceName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemExternalResourceStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemExternalResourceType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("category", flattenSystemExternalResourceCategory(o["category"], d, "category")); err != nil {
		if !fortiAPIPatch(o["category"]) {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("username", flattenSystemExternalResourceUsername(o["username"], d, "username")); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("password", flattenSystemExternalResourcePassword(o["password"], d, "password")); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("comments", flattenSystemExternalResourceComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("resource", flattenSystemExternalResourceResource(o["resource"], d, "resource")); err != nil {
		if !fortiAPIPatch(o["resource"]) {
			return fmt.Errorf("Error reading resource: %v", err)
		}
	}

	if err = d.Set("refresh_rate", flattenSystemExternalResourceRefreshRate(o["refresh-rate"], d, "refresh_rate")); err != nil {
		if !fortiAPIPatch(o["refresh-rate"]) {
			return fmt.Errorf("Error reading refresh_rate: %v", err)
		}
	}


	return nil
}

func flattenSystemExternalResourceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandSystemExternalResourceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourcePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceResource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemExternalResourceRefreshRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectSystemExternalResource(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemExternalResourceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemExternalResourceStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemExternalResourceType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok {
		t, err := expandSystemExternalResourceCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {
		t, err := expandSystemExternalResourceUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandSystemExternalResourcePassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandSystemExternalResourceComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("resource"); ok {
		t, err := expandSystemExternalResourceResource(d, v, "resource")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resource"] = t
		}
	}

	if v, ok := d.GetOk("refresh_rate"); ok {
		t, err := expandSystemExternalResourceRefreshRate(d, v, "refresh_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["refresh-rate"] = t
		}
	}


	return &obj, nil
}

