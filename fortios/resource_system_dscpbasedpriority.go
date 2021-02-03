// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure DSCP based priority table.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemDscpBasedPriority() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDscpBasedPriorityCreate,
		Read:   resourceSystemDscpBasedPriorityRead,
		Update: resourceSystemDscpBasedPriorityUpdate,
		Delete: resourceSystemDscpBasedPriorityDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"ds": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemDscpBasedPriorityCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemDscpBasedPriority(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemDscpBasedPriority resource while getting object: %v", err)
	}

	o, err := c.CreateSystemDscpBasedPriority(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemDscpBasedPriority resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemDscpBasedPriority")
	}

	return resourceSystemDscpBasedPriorityRead(d, m)
}

func resourceSystemDscpBasedPriorityUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemDscpBasedPriority(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemDscpBasedPriority resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemDscpBasedPriority(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemDscpBasedPriority resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemDscpBasedPriority")
	}

	return resourceSystemDscpBasedPriorityRead(d, m)
}

func resourceSystemDscpBasedPriorityDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemDscpBasedPriority(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDscpBasedPriority resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDscpBasedPriorityRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemDscpBasedPriority(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemDscpBasedPriority resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDscpBasedPriority(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemDscpBasedPriority resource from API: %v", err)
	}
	return nil
}

func flattenSystemDscpBasedPriorityId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDscpBasedPriorityDs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDscpBasedPriorityPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemDscpBasedPriority(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenSystemDscpBasedPriorityId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ds", flattenSystemDscpBasedPriorityDs(o["ds"], d, "ds", sv)); err != nil {
		if !fortiAPIPatch(o["ds"]) {
			return fmt.Errorf("Error reading ds: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystemDscpBasedPriorityPriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	return nil
}

func flattenSystemDscpBasedPriorityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemDscpBasedPriorityId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDscpBasedPriorityDs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDscpBasedPriorityPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDscpBasedPriority(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandSystemDscpBasedPriorityId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOkExists("ds"); ok {

		t, err := expandSystemDscpBasedPriorityDs(d, v, "ds", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ds"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok {

		t, err := expandSystemDscpBasedPriorityPriority(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	return &obj, nil
}
