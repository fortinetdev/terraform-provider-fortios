// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure quarantine options.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAntivirusQuarantine() *schema.Resource {
	return &schema.Resource{
		Create: resourceAntivirusQuarantineUpdate,
		Read:   resourceAntivirusQuarantineRead,
		Update: resourceAntivirusQuarantineUpdate,
		Delete: resourceAntivirusQuarantineDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"agelimit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 479),
				Optional:     true,
				Computed:     true,
			},
			"maxfilesize": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 500),
				Optional:     true,
				Computed:     true,
			},
			"quarantine_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"drop_infected": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"store_infected": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"drop_blocked": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"store_blocked": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"drop_heuristic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"store_heuristic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lowspace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destination": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceAntivirusQuarantineUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectAntivirusQuarantine(d)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusQuarantine resource while getting object: %v", err)
	}

	o, err := c.UpdateAntivirusQuarantine(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusQuarantine resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("AntivirusQuarantine")
	}

	return resourceAntivirusQuarantineRead(d, m)
}

func resourceAntivirusQuarantineDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteAntivirusQuarantine(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting AntivirusQuarantine resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAntivirusQuarantineRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadAntivirusQuarantine(mkey)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusQuarantine resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAntivirusQuarantine(d, o)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusQuarantine resource from API: %v", err)
	}
	return nil
}

func flattenAntivirusQuarantineAgelimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusQuarantineMaxfilesize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusQuarantineQuarantineQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusQuarantineDropInfected(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusQuarantineStoreInfected(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusQuarantineDropBlocked(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusQuarantineStoreBlocked(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusQuarantineDropHeuristic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusQuarantineStoreHeuristic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusQuarantineLowspace(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusQuarantineDestination(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectAntivirusQuarantine(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("agelimit", flattenAntivirusQuarantineAgelimit(o["agelimit"], d, "agelimit")); err != nil {
		if !fortiAPIPatch(o["agelimit"]) {
			return fmt.Errorf("Error reading agelimit: %v", err)
		}
	}

	if err = d.Set("maxfilesize", flattenAntivirusQuarantineMaxfilesize(o["maxfilesize"], d, "maxfilesize")); err != nil {
		if !fortiAPIPatch(o["maxfilesize"]) {
			return fmt.Errorf("Error reading maxfilesize: %v", err)
		}
	}

	if err = d.Set("quarantine_quota", flattenAntivirusQuarantineQuarantineQuota(o["quarantine-quota"], d, "quarantine_quota")); err != nil {
		if !fortiAPIPatch(o["quarantine-quota"]) {
			return fmt.Errorf("Error reading quarantine_quota: %v", err)
		}
	}

	if err = d.Set("drop_infected", flattenAntivirusQuarantineDropInfected(o["drop-infected"], d, "drop_infected")); err != nil {
		if !fortiAPIPatch(o["drop-infected"]) {
			return fmt.Errorf("Error reading drop_infected: %v", err)
		}
	}

	if err = d.Set("store_infected", flattenAntivirusQuarantineStoreInfected(o["store-infected"], d, "store_infected")); err != nil {
		if !fortiAPIPatch(o["store-infected"]) {
			return fmt.Errorf("Error reading store_infected: %v", err)
		}
	}

	if err = d.Set("drop_blocked", flattenAntivirusQuarantineDropBlocked(o["drop-blocked"], d, "drop_blocked")); err != nil {
		if !fortiAPIPatch(o["drop-blocked"]) {
			return fmt.Errorf("Error reading drop_blocked: %v", err)
		}
	}

	if err = d.Set("store_blocked", flattenAntivirusQuarantineStoreBlocked(o["store-blocked"], d, "store_blocked")); err != nil {
		if !fortiAPIPatch(o["store-blocked"]) {
			return fmt.Errorf("Error reading store_blocked: %v", err)
		}
	}

	if err = d.Set("drop_heuristic", flattenAntivirusQuarantineDropHeuristic(o["drop-heuristic"], d, "drop_heuristic")); err != nil {
		if !fortiAPIPatch(o["drop-heuristic"]) {
			return fmt.Errorf("Error reading drop_heuristic: %v", err)
		}
	}

	if err = d.Set("store_heuristic", flattenAntivirusQuarantineStoreHeuristic(o["store-heuristic"], d, "store_heuristic")); err != nil {
		if !fortiAPIPatch(o["store-heuristic"]) {
			return fmt.Errorf("Error reading store_heuristic: %v", err)
		}
	}

	if err = d.Set("lowspace", flattenAntivirusQuarantineLowspace(o["lowspace"], d, "lowspace")); err != nil {
		if !fortiAPIPatch(o["lowspace"]) {
			return fmt.Errorf("Error reading lowspace: %v", err)
		}
	}

	if err = d.Set("destination", flattenAntivirusQuarantineDestination(o["destination"], d, "destination")); err != nil {
		if !fortiAPIPatch(o["destination"]) {
			return fmt.Errorf("Error reading destination: %v", err)
		}
	}

	return nil
}

func flattenAntivirusQuarantineFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandAntivirusQuarantineAgelimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineMaxfilesize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineQuarantineQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineDropInfected(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineStoreInfected(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineDropBlocked(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineStoreBlocked(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineDropHeuristic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineStoreHeuristic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineLowspace(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineDestination(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectAntivirusQuarantine(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("agelimit"); ok {
		t, err := expandAntivirusQuarantineAgelimit(d, v, "agelimit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["agelimit"] = t
		}
	}

	if v, ok := d.GetOkExists("maxfilesize"); ok {
		t, err := expandAntivirusQuarantineMaxfilesize(d, v, "maxfilesize")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maxfilesize"] = t
		}
	}

	if v, ok := d.GetOkExists("quarantine_quota"); ok {
		t, err := expandAntivirusQuarantineQuarantineQuota(d, v, "quarantine_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-quota"] = t
		}
	}

	if v, ok := d.GetOk("drop_infected"); ok {
		t, err := expandAntivirusQuarantineDropInfected(d, v, "drop_infected")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drop-infected"] = t
		}
	}

	if v, ok := d.GetOk("store_infected"); ok {
		t, err := expandAntivirusQuarantineStoreInfected(d, v, "store_infected")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["store-infected"] = t
		}
	}

	if v, ok := d.GetOk("drop_blocked"); ok {
		t, err := expandAntivirusQuarantineDropBlocked(d, v, "drop_blocked")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drop-blocked"] = t
		}
	}

	if v, ok := d.GetOk("store_blocked"); ok {
		t, err := expandAntivirusQuarantineStoreBlocked(d, v, "store_blocked")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["store-blocked"] = t
		}
	}

	if v, ok := d.GetOk("drop_heuristic"); ok {
		t, err := expandAntivirusQuarantineDropHeuristic(d, v, "drop_heuristic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drop-heuristic"] = t
		}
	}

	if v, ok := d.GetOk("store_heuristic"); ok {
		t, err := expandAntivirusQuarantineStoreHeuristic(d, v, "store_heuristic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["store-heuristic"] = t
		}
	}

	if v, ok := d.GetOk("lowspace"); ok {
		t, err := expandAntivirusQuarantineLowspace(d, v, "lowspace")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lowspace"] = t
		}
	}

	if v, ok := d.GetOk("destination"); ok {
		t, err := expandAntivirusQuarantineDestination(d, v, "destination")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination"] = t
		}
	}

	return &obj, nil
}
