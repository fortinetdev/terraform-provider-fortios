// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Create self-explanatory DLP sensitivity levels to be used when setting sensitivity under config fp-doc-source.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDlpFpSensitivity() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpFpSensitivityCreate,
		Read:   resourceDlpFpSensitivityRead,
		Update: resourceDlpFpSensitivityUpdate,
		Delete: resourceDlpFpSensitivityDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceDlpFpSensitivityCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectDlpFpSensitivity(d)
	if err != nil {
		return fmt.Errorf("Error creating DlpFpSensitivity resource while getting object: %v", err)
	}

	o, err := c.CreateDlpFpSensitivity(obj)

	if err != nil {
		return fmt.Errorf("Error creating DlpFpSensitivity resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DlpFpSensitivity")
	}

	return resourceDlpFpSensitivityRead(d, m)
}

func resourceDlpFpSensitivityUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectDlpFpSensitivity(d)
	if err != nil {
		return fmt.Errorf("Error updating DlpFpSensitivity resource while getting object: %v", err)
	}

	o, err := c.UpdateDlpFpSensitivity(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating DlpFpSensitivity resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DlpFpSensitivity")
	}

	return resourceDlpFpSensitivityRead(d, m)
}

func resourceDlpFpSensitivityDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteDlpFpSensitivity(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting DlpFpSensitivity resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpFpSensitivityRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadDlpFpSensitivity(mkey)
	if err != nil {
		return fmt.Errorf("Error reading DlpFpSensitivity resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpFpSensitivity(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpFpSensitivity resource from API: %v", err)
	}
	return nil
}

func flattenDlpFpSensitivityName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDlpFpSensitivity(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenDlpFpSensitivityName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenDlpFpSensitivityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpFpSensitivityName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDlpFpSensitivity(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandDlpFpSensitivityName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
