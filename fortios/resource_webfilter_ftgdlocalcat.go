// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure FortiGuard Web Filter local categories.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWebfilterFtgdLocalCat() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterFtgdLocalCatCreate,
		Read:   resourceWebfilterFtgdLocalCatRead,
		Update: resourceWebfilterFtgdLocalCatUpdate,
		Delete: resourceWebfilterFtgdLocalCatDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(140, 191),
				Optional:     true,
				Computed:     true,
			},
			"desc": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Required:     true,
				ForceNew:     true,
			},
		},
	}
}

func resourceWebfilterFtgdLocalCatCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWebfilterFtgdLocalCat(d)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterFtgdLocalCat resource while getting object: %v", err)
	}

	o, err := c.CreateWebfilterFtgdLocalCat(obj)

	if err != nil {
		return fmt.Errorf("Error creating WebfilterFtgdLocalCat resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebfilterFtgdLocalCat")
	}

	return resourceWebfilterFtgdLocalCatRead(d, m)
}

func resourceWebfilterFtgdLocalCatUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWebfilterFtgdLocalCat(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterFtgdLocalCat resource while getting object: %v", err)
	}

	o, err := c.UpdateWebfilterFtgdLocalCat(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterFtgdLocalCat resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebfilterFtgdLocalCat")
	}

	return resourceWebfilterFtgdLocalCatRead(d, m)
}

func resourceWebfilterFtgdLocalCatDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWebfilterFtgdLocalCat(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterFtgdLocalCat resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterFtgdLocalCatRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWebfilterFtgdLocalCat(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterFtgdLocalCat resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterFtgdLocalCat(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterFtgdLocalCat resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterFtgdLocalCatStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFtgdLocalCatId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterFtgdLocalCatDesc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterFtgdLocalCat(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", flattenWebfilterFtgdLocalCatStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("fosid", flattenWebfilterFtgdLocalCatId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("desc", flattenWebfilterFtgdLocalCatDesc(o["desc"], d, "desc")); err != nil {
		if !fortiAPIPatch(o["desc"]) {
			return fmt.Errorf("Error reading desc: %v", err)
		}
	}

	return nil
}

func flattenWebfilterFtgdLocalCatFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterFtgdLocalCatStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFtgdLocalCatId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFtgdLocalCatDesc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterFtgdLocalCat(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandWebfilterFtgdLocalCatStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandWebfilterFtgdLocalCatId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("desc"); ok {
		t, err := expandWebfilterFtgdLocalCatDesc(d, v, "desc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["desc"] = t
		}
	}

	return &obj, nil
}
