// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure custom languages.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemCustomLanguage() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCustomLanguageCreate,
		Read:   resourceSystemCustomLanguageRead,
		Update: resourceSystemCustomLanguageUpdate,
		Delete: resourceSystemCustomLanguageDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"filename": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required:     true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
		},
	}
}

func resourceSystemCustomLanguageCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemCustomLanguage(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemCustomLanguage resource while getting object: %v", err)
	}

	o, err := c.CreateSystemCustomLanguage(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemCustomLanguage resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemCustomLanguage")
	}

	return resourceSystemCustomLanguageRead(d, m)
}

func resourceSystemCustomLanguageUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemCustomLanguage(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemCustomLanguage resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemCustomLanguage(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemCustomLanguage resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemCustomLanguage")
	}

	return resourceSystemCustomLanguageRead(d, m)
}

func resourceSystemCustomLanguageDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemCustomLanguage(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCustomLanguage resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCustomLanguageRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemCustomLanguage(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemCustomLanguage resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCustomLanguage(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemCustomLanguage resource from API: %v", err)
	}
	return nil
}

func flattenSystemCustomLanguageName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCustomLanguageFilename(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCustomLanguageComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemCustomLanguage(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemCustomLanguageName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("filename", flattenSystemCustomLanguageFilename(o["filename"], d, "filename")); err != nil {
		if !fortiAPIPatch(o["filename"]) {
			return fmt.Errorf("Error reading filename: %v", err)
		}
	}

	if err = d.Set("comments", flattenSystemCustomLanguageComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	return nil
}

func flattenSystemCustomLanguageFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemCustomLanguageName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCustomLanguageFilename(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCustomLanguageComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCustomLanguage(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemCustomLanguageName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("filename"); ok {
		t, err := expandSystemCustomLanguageFilename(d, v, "filename")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filename"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandSystemCustomLanguageComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	return &obj, nil
}
