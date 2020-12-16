// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure replacement message images.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemReplacemsgImage() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemReplacemsgImageCreate,
		Read:   resourceSystemReplacemsgImageRead,
		Update: resourceSystemReplacemsgImageUpdate,
		Delete: resourceSystemReplacemsgImageDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 23),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"image_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"image_base64": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32768),
				Optional:     true,
			},
		},
	}
}

func resourceSystemReplacemsgImageCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemReplacemsgImage(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemReplacemsgImage resource while getting object: %v", err)
	}

	o, err := c.CreateSystemReplacemsgImage(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemReplacemsgImage resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemReplacemsgImage")
	}

	return resourceSystemReplacemsgImageRead(d, m)
}

func resourceSystemReplacemsgImageUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemReplacemsgImage(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgImage resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemReplacemsgImage(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgImage resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemReplacemsgImage")
	}

	return resourceSystemReplacemsgImageRead(d, m)
}

func resourceSystemReplacemsgImageDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemReplacemsgImage(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemReplacemsgImage resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemReplacemsgImageRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemReplacemsgImage(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemReplacemsgImage resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemReplacemsgImage(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemReplacemsgImage resource from API: %v", err)
	}
	return nil
}

func flattenSystemReplacemsgImageName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgImageImageType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgImageImageBase64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemReplacemsgImage(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemReplacemsgImageName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("image_type", flattenSystemReplacemsgImageImageType(o["image-type"], d, "image_type")); err != nil {
		if !fortiAPIPatch(o["image-type"]) {
			return fmt.Errorf("Error reading image_type: %v", err)
		}
	}

	if err = d.Set("image_base64", flattenSystemReplacemsgImageImageBase64(o["image-base64"], d, "image_base64")); err != nil {
		if !fortiAPIPatch(o["image-base64"]) {
			return fmt.Errorf("Error reading image_base64: %v", err)
		}
	}

	return nil
}

func flattenSystemReplacemsgImageFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemReplacemsgImageName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgImageImageType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgImageImageBase64(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemReplacemsgImage(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemReplacemsgImageName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("image_type"); ok {
		t, err := expandSystemReplacemsgImageImageType(d, v, "image_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-type"] = t
		}
	}

	if v, ok := d.GetOk("image_base64"); ok {
		t, err := expandSystemReplacemsgImageImageBase64(d, v, "image_base64")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-base64"] = t
		}
	}

	return &obj, nil
}
