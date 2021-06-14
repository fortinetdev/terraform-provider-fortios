// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure device categories.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceUserDeviceCategory() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserDeviceCategoryCreate,
		Read:   resourceUserDeviceCategoryRead,
		Update: resourceUserDeviceCategoryUpdate,
		Delete: resourceUserDeviceCategoryDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Required:     true,
			},
			"desc": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
		},
	}
}

func resourceUserDeviceCategoryCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserDeviceCategory(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserDeviceCategory resource while getting object: %v", err)
	}

	o, err := c.CreateUserDeviceCategory(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating UserDeviceCategory resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserDeviceCategory")
	}

	return resourceUserDeviceCategoryRead(d, m)
}

func resourceUserDeviceCategoryUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserDeviceCategory(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserDeviceCategory resource while getting object: %v", err)
	}

	o, err := c.UpdateUserDeviceCategory(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserDeviceCategory resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserDeviceCategory")
	}

	return resourceUserDeviceCategoryRead(d, m)
}

func resourceUserDeviceCategoryDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUserDeviceCategory(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UserDeviceCategory resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserDeviceCategoryRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadUserDeviceCategory(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserDeviceCategory resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserDeviceCategory(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserDeviceCategory resource from API: %v", err)
	}
	return nil
}

func flattenUserDeviceCategoryName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDeviceCategoryDesc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDeviceCategoryComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserDeviceCategory(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenUserDeviceCategoryName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("desc", flattenUserDeviceCategoryDesc(o["desc"], d, "desc", sv)); err != nil {
		if !fortiAPIPatch(o["desc"]) {
			return fmt.Errorf("Error reading desc: %v", err)
		}
	}

	if err = d.Set("comment", flattenUserDeviceCategoryComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	return nil
}

func flattenUserDeviceCategoryFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserDeviceCategoryName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDeviceCategoryDesc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDeviceCategoryComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserDeviceCategory(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandUserDeviceCategoryName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("desc"); ok {

		t, err := expandUserDeviceCategoryDesc(d, v, "desc", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["desc"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandUserDeviceCategoryComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	return &obj, nil
}
