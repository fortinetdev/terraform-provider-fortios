// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure system-wide switch controller settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerSystem() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSystemUpdate,
		Read:   resourceSwitchControllerSystemRead,
		Update: resourceSwitchControllerSystemUpdate,
		Delete: resourceSwitchControllerSystemDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"parallel_process_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parallel_process": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSwitchControllerSystemUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerSystem(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSystem resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerSystem(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSystem resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerSystem")
	}

	return resourceSwitchControllerSystemRead(d, m)
}

func resourceSwitchControllerSystemDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchControllerSystem(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerSystem resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSystemRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchControllerSystem(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSystem resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSystem(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSystem resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSystemParallelProcessOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSystemParallelProcess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerSystem(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("parallel_process_override", flattenSwitchControllerSystemParallelProcessOverride(o["parallel-process-override"], d, "parallel_process_override")); err != nil {
		if !fortiAPIPatch(o["parallel-process-override"]) {
			return fmt.Errorf("Error reading parallel_process_override: %v", err)
		}
	}

	if err = d.Set("parallel_process", flattenSwitchControllerSystemParallelProcess(o["parallel-process"], d, "parallel_process")); err != nil {
		if !fortiAPIPatch(o["parallel-process"]) {
			return fmt.Errorf("Error reading parallel_process: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSystemFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerSystemParallelProcessOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemParallelProcess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerSystem(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("parallel_process_override"); ok {
		t, err := expandSwitchControllerSystemParallelProcessOverride(d, v, "parallel_process_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parallel-process-override"] = t
		}
	}

	if v, ok := d.GetOk("parallel_process"); ok {
		t, err := expandSwitchControllerSystemParallelProcess(d, v, "parallel_process")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parallel-process"] = t
		}
	}

	return &obj, nil
}
