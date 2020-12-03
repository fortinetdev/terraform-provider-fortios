// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure custom application signatures.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceApplicationCustom() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationCustomCreate,
		Read:   resourceApplicationCustomRead,
		Update: resourceApplicationCustomUpdate,
		Delete: resourceApplicationCustomDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"tag": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required:     true,
				ForceNew:     true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"signature": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
				Computed:     true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"technology": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"behavior": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceApplicationCustomCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectApplicationCustom(d)
	if err != nil {
		return fmt.Errorf("Error creating ApplicationCustom resource while getting object: %v", err)
	}

	o, err := c.CreateApplicationCustom(obj)

	if err != nil {
		return fmt.Errorf("Error creating ApplicationCustom resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ApplicationCustom")
	}

	return resourceApplicationCustomRead(d, m)
}

func resourceApplicationCustomUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectApplicationCustom(d)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationCustom resource while getting object: %v", err)
	}

	o, err := c.UpdateApplicationCustom(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationCustom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ApplicationCustom")
	}

	return resourceApplicationCustomRead(d, m)
}

func resourceApplicationCustomDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteApplicationCustom(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting ApplicationCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceApplicationCustomRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadApplicationCustom(mkey)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationCustom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectApplicationCustom(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationCustom resource from API: %v", err)
	}
	return nil
}

func flattenApplicationCustomTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationCustomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationCustomId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationCustomComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationCustomSignature(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationCustomCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationCustomProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationCustomTechnology(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationCustomBehavior(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationCustomVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectApplicationCustom(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("tag", flattenApplicationCustomTag(o["tag"], d, "tag")); err != nil {
		if !fortiAPIPatch(o["tag"]) {
			return fmt.Errorf("Error reading tag: %v", err)
		}
	}

	if err = d.Set("name", flattenApplicationCustomName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenApplicationCustomId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("comment", flattenApplicationCustomComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("signature", flattenApplicationCustomSignature(o["signature"], d, "signature")); err != nil {
		if !fortiAPIPatch(o["signature"]) {
			return fmt.Errorf("Error reading signature: %v", err)
		}
	}

	if err = d.Set("category", flattenApplicationCustomCategory(o["category"], d, "category")); err != nil {
		if !fortiAPIPatch(o["category"]) {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("protocol", flattenApplicationCustomProtocol(o["protocol"], d, "protocol")); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("technology", flattenApplicationCustomTechnology(o["technology"], d, "technology")); err != nil {
		if !fortiAPIPatch(o["technology"]) {
			return fmt.Errorf("Error reading technology: %v", err)
		}
	}

	if err = d.Set("behavior", flattenApplicationCustomBehavior(o["behavior"], d, "behavior")); err != nil {
		if !fortiAPIPatch(o["behavior"]) {
			return fmt.Errorf("Error reading behavior: %v", err)
		}
	}

	if err = d.Set("vendor", flattenApplicationCustomVendor(o["vendor"], d, "vendor")); err != nil {
		if !fortiAPIPatch(o["vendor"]) {
			return fmt.Errorf("Error reading vendor: %v", err)
		}
	}

	return nil
}

func flattenApplicationCustomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandApplicationCustomTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationCustomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationCustomId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationCustomComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationCustomSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationCustomCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationCustomProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationCustomTechnology(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationCustomBehavior(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationCustomVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectApplicationCustom(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("tag"); ok {
		t, err := expandApplicationCustomTag(d, v, "tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandApplicationCustomName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandApplicationCustomId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandApplicationCustomComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("signature"); ok {
		t, err := expandApplicationCustomSignature(d, v, "signature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signature"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok {
		t, err := expandApplicationCustomCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandApplicationCustomProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("technology"); ok {
		t, err := expandApplicationCustomTechnology(d, v, "technology")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["technology"] = t
		}
	}

	if v, ok := d.GetOk("behavior"); ok {
		t, err := expandApplicationCustomBehavior(d, v, "behavior")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["behavior"] = t
		}
	}

	if v, ok := d.GetOk("vendor"); ok {
		t, err := expandApplicationCustomVendor(d, v, "vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor"] = t
		}
	}

	return &obj, nil
}
