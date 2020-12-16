// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure interrupt affinity.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemAffinityInterrupt() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAffinityInterruptCreate,
		Read:   resourceSystemAffinityInterruptRead,
		Update: resourceSystemAffinityInterruptUpdate,
		Delete: resourceSystemAffinityInterruptDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Required: true,
			},
			"interrupt": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Required:     true,
			},
			"affinity_cpumask": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Required:     true,
			},
		},
	}
}

func resourceSystemAffinityInterruptCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAffinityInterrupt(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAffinityInterrupt resource while getting object: %v", err)
	}

	o, err := c.CreateSystemAffinityInterrupt(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemAffinityInterrupt resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemAffinityInterrupt")
	}

	return resourceSystemAffinityInterruptRead(d, m)
}

func resourceSystemAffinityInterruptUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAffinityInterrupt(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAffinityInterrupt resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAffinityInterrupt(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemAffinityInterrupt resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemAffinityInterrupt")
	}

	return resourceSystemAffinityInterruptRead(d, m)
}

func resourceSystemAffinityInterruptDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemAffinityInterrupt(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAffinityInterrupt resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAffinityInterruptRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemAffinityInterrupt(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemAffinityInterrupt resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAffinityInterrupt(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAffinityInterrupt resource from API: %v", err)
	}
	return nil
}

func flattenSystemAffinityInterruptId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAffinityInterruptInterrupt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAffinityInterruptAffinityCpumask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAffinityInterrupt(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemAffinityInterruptId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interrupt", flattenSystemAffinityInterruptInterrupt(o["interrupt"], d, "interrupt")); err != nil {
		if !fortiAPIPatch(o["interrupt"]) {
			return fmt.Errorf("Error reading interrupt: %v", err)
		}
	}

	if err = d.Set("affinity_cpumask", flattenSystemAffinityInterruptAffinityCpumask(o["affinity-cpumask"], d, "affinity_cpumask")); err != nil {
		if !fortiAPIPatch(o["affinity-cpumask"]) {
			return fmt.Errorf("Error reading affinity_cpumask: %v", err)
		}
	}

	return nil
}

func flattenSystemAffinityInterruptFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAffinityInterruptId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAffinityInterruptInterrupt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAffinityInterruptAffinityCpumask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAffinityInterrupt(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandSystemAffinityInterruptId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interrupt"); ok {
		t, err := expandSystemAffinityInterruptInterrupt(d, v, "interrupt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interrupt"] = t
		}
	}

	if v, ok := d.GetOk("affinity_cpumask"); ok {
		t, err := expandSystemAffinityInterruptAffinityCpumask(d, v, "affinity_cpumask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["affinity-cpumask"] = t
		}
	}

	return &obj, nil
}
