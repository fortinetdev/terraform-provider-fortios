// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure VDOM links.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemVdomLink() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVdomLinkCreate,
		Read:   resourceSystemVdomLinkRead,
		Update: resourceSystemVdomLinkUpdate,
		Delete: resourceSystemVdomLinkDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 11),
				Required:     true,
				ForceNew:     true,
			},
			"vcluster": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemVdomLinkCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemVdomLink(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdomLink resource while getting object: %v", err)
	}

	o, err := c.CreateSystemVdomLink(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemVdomLink resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVdomLink")
	}

	return resourceSystemVdomLinkRead(d, m)
}

func resourceSystemVdomLinkUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemVdomLink(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomLink resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemVdomLink(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomLink resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVdomLink")
	}

	return resourceSystemVdomLinkRead(d, m)
}

func resourceSystemVdomLinkDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemVdomLink(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVdomLink resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomLinkRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemVdomLink(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomLink resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdomLink(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomLink resource from API: %v", err)
	}
	return nil
}

func flattenSystemVdomLinkName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomLinkVcluster(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomLinkType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemVdomLink(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemVdomLinkName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vcluster", flattenSystemVdomLinkVcluster(o["vcluster"], d, "vcluster")); err != nil {
		if !fortiAPIPatch(o["vcluster"]) {
			return fmt.Errorf("Error reading vcluster: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemVdomLinkType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenSystemVdomLinkFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemVdomLinkName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomLinkVcluster(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomLinkType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVdomLink(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemVdomLinkName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vcluster"); ok {
		t, err := expandSystemVdomLinkVcluster(d, v, "vcluster")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcluster"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemVdomLinkType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
