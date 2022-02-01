// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure physical switches.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemPhysicalSwitch() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemPhysicalSwitchCreate,
		Read:   resourceSystemPhysicalSwitchRead,
		Update: resourceSystemPhysicalSwitchUpdate,
		Delete: resourceSystemPhysicalSwitchDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 15),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"age_enable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"age_val": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemPhysicalSwitchCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemPhysicalSwitch(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemPhysicalSwitch resource while getting object: %v", err)
	}

	o, err := c.CreateSystemPhysicalSwitch(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemPhysicalSwitch resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemPhysicalSwitch")
	}

	return resourceSystemPhysicalSwitchRead(d, m)
}

func resourceSystemPhysicalSwitchUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemPhysicalSwitch(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemPhysicalSwitch resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemPhysicalSwitch(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemPhysicalSwitch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemPhysicalSwitch")
	}

	return resourceSystemPhysicalSwitchRead(d, m)
}

func resourceSystemPhysicalSwitchDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemPhysicalSwitch(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemPhysicalSwitch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPhysicalSwitchRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemPhysicalSwitch(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemPhysicalSwitch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemPhysicalSwitch(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemPhysicalSwitch resource from API: %v", err)
	}
	return nil
}

func flattenSystemPhysicalSwitchName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPhysicalSwitchAgeEnable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPhysicalSwitchAgeVal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemPhysicalSwitch(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemPhysicalSwitchName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("age_enable", flattenSystemPhysicalSwitchAgeEnable(o["age-enable"], d, "age_enable", sv)); err != nil {
		if !fortiAPIPatch(o["age-enable"]) {
			return fmt.Errorf("Error reading age_enable: %v", err)
		}
	}

	if err = d.Set("age_val", flattenSystemPhysicalSwitchAgeVal(o["age-val"], d, "age_val", sv)); err != nil {
		if !fortiAPIPatch(o["age-val"]) {
			return fmt.Errorf("Error reading age_val: %v", err)
		}
	}

	return nil
}

func flattenSystemPhysicalSwitchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemPhysicalSwitchName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPhysicalSwitchAgeEnable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPhysicalSwitchAgeVal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemPhysicalSwitch(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemPhysicalSwitchName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("age_enable"); ok {

		t, err := expandSystemPhysicalSwitchAgeEnable(d, v, "age_enable", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["age-enable"] = t
		}
	}

	if v, ok := d.GetOkExists("age_val"); ok {

		t, err := expandSystemPhysicalSwitchAgeVal(d, v, "age_val", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["age-val"] = t
		}
	}

	return &obj, nil
}
