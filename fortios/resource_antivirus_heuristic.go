// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure global heuristic options.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAntivirusHeuristic() *schema.Resource {
	return &schema.Resource{
		Create: resourceAntivirusHeuristicUpdate,
		Read:   resourceAntivirusHeuristicRead,
		Update: resourceAntivirusHeuristicUpdate,
		Delete: resourceAntivirusHeuristicDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceAntivirusHeuristicUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectAntivirusHeuristic(d)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusHeuristic resource while getting object: %v", err)
	}

	o, err := c.UpdateAntivirusHeuristic(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusHeuristic resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("AntivirusHeuristic")
	}

	return resourceAntivirusHeuristicRead(d, m)
}

func resourceAntivirusHeuristicDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteAntivirusHeuristic(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting AntivirusHeuristic resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAntivirusHeuristicRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadAntivirusHeuristic(mkey)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusHeuristic resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAntivirusHeuristic(d, o)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusHeuristic resource from API: %v", err)
	}
	return nil
}

func flattenAntivirusHeuristicMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectAntivirusHeuristic(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("mode", flattenAntivirusHeuristicMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	return nil
}

func flattenAntivirusHeuristicFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandAntivirusHeuristicMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectAntivirusHeuristic(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandAntivirusHeuristicMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	return &obj, nil
}
