// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure virtual domain.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemVdom() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVdomCreate,
		Read:   resourceSystemVdomRead,
		Update: resourceSystemVdomUpdate,
		Delete: resourceSystemVdomDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"short_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 11),
				Optional:     true,
				Computed:     true,
			},
			"vcluster_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"temporary": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemVdomCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemVdom(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdom resource while getting object: %v", err)
	}

	o, err := c.CreateSystemVdom(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemVdom resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVdom")
	}

	return resourceSystemVdomRead(d, m)
}

func resourceSystemVdomUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemVdom(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdom resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemVdom(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVdom")
	}

	return resourceSystemVdomRead(d, m)
}

func resourceSystemVdomDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemVdom(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVdom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemVdom(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdom(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdom resource from API: %v", err)
	}
	return nil
}

func flattenSystemVdomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomShortName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomVclusterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomTemporary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemVdom(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemVdomName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("short_name", flattenSystemVdomShortName(o["short-name"], d, "short_name")); err != nil {
		if !fortiAPIPatch(o["short-name"]) {
			return fmt.Errorf("Error reading short_name: %v", err)
		}
	}

	if err = d.Set("vcluster_id", flattenSystemVdomVclusterId(o["vcluster-id"], d, "vcluster_id")); err != nil {
		if !fortiAPIPatch(o["vcluster-id"]) {
			return fmt.Errorf("Error reading vcluster_id: %v", err)
		}
	}

	if err = d.Set("temporary", flattenSystemVdomTemporary(o["temporary"], d, "temporary")); err != nil {
		if !fortiAPIPatch(o["temporary"]) {
			return fmt.Errorf("Error reading temporary: %v", err)
		}
	}

	return nil
}

func flattenSystemVdomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemVdomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomShortName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomVclusterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomTemporary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVdom(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemVdomName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("short_name"); ok {
		t, err := expandSystemVdomShortName(d, v, "short_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["short-name"] = t
		}
	}

	if v, ok := d.GetOkExists("vcluster_id"); ok {
		t, err := expandSystemVdomVclusterId(d, v, "vcluster_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcluster-id"] = t
		}
	}

	if v, ok := d.GetOkExists("temporary"); ok {
		t, err := expandSystemVdomTemporary(d, v, "temporary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["temporary"] = t
		}
	}

	return &obj, nil
}
