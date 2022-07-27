// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure quarantine options.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
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
			"drop_machine_learning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"store_machine_learning": &schema.Schema{
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectAntivirusQuarantine(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusQuarantine resource while getting object: %v", err)
	}

	o, err := c.UpdateAntivirusQuarantine(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectAntivirusQuarantine(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating AntivirusQuarantine resource while getting object: %v", err)
	}

	_, err = c.UpdateAntivirusQuarantine(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing AntivirusQuarantine resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAntivirusQuarantineRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadAntivirusQuarantine(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusQuarantine resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAntivirusQuarantine(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusQuarantine resource from API: %v", err)
	}
	return nil
}

func flattenAntivirusQuarantineAgelimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusQuarantineMaxfilesize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusQuarantineQuarantineQuota(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusQuarantineDropInfected(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusQuarantineStoreInfected(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusQuarantineDropBlocked(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusQuarantineStoreBlocked(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusQuarantineDropMachineLearning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusQuarantineStoreMachineLearning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusQuarantineDropHeuristic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusQuarantineStoreHeuristic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusQuarantineLowspace(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusQuarantineDestination(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectAntivirusQuarantine(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("agelimit", flattenAntivirusQuarantineAgelimit(o["agelimit"], d, "agelimit", sv)); err != nil {
		if !fortiAPIPatch(o["agelimit"]) {
			return fmt.Errorf("Error reading agelimit: %v", err)
		}
	}

	if err = d.Set("maxfilesize", flattenAntivirusQuarantineMaxfilesize(o["maxfilesize"], d, "maxfilesize", sv)); err != nil {
		if !fortiAPIPatch(o["maxfilesize"]) {
			return fmt.Errorf("Error reading maxfilesize: %v", err)
		}
	}

	if err = d.Set("quarantine_quota", flattenAntivirusQuarantineQuarantineQuota(o["quarantine-quota"], d, "quarantine_quota", sv)); err != nil {
		if !fortiAPIPatch(o["quarantine-quota"]) {
			return fmt.Errorf("Error reading quarantine_quota: %v", err)
		}
	}

	if err = d.Set("drop_infected", flattenAntivirusQuarantineDropInfected(o["drop-infected"], d, "drop_infected", sv)); err != nil {
		if !fortiAPIPatch(o["drop-infected"]) {
			return fmt.Errorf("Error reading drop_infected: %v", err)
		}
	}

	if err = d.Set("store_infected", flattenAntivirusQuarantineStoreInfected(o["store-infected"], d, "store_infected", sv)); err != nil {
		if !fortiAPIPatch(o["store-infected"]) {
			return fmt.Errorf("Error reading store_infected: %v", err)
		}
	}

	if err = d.Set("drop_blocked", flattenAntivirusQuarantineDropBlocked(o["drop-blocked"], d, "drop_blocked", sv)); err != nil {
		if !fortiAPIPatch(o["drop-blocked"]) {
			return fmt.Errorf("Error reading drop_blocked: %v", err)
		}
	}

	if err = d.Set("store_blocked", flattenAntivirusQuarantineStoreBlocked(o["store-blocked"], d, "store_blocked", sv)); err != nil {
		if !fortiAPIPatch(o["store-blocked"]) {
			return fmt.Errorf("Error reading store_blocked: %v", err)
		}
	}

	if err = d.Set("drop_machine_learning", flattenAntivirusQuarantineDropMachineLearning(o["drop-machine-learning"], d, "drop_machine_learning", sv)); err != nil {
		if !fortiAPIPatch(o["drop-machine-learning"]) {
			return fmt.Errorf("Error reading drop_machine_learning: %v", err)
		}
	}

	if err = d.Set("store_machine_learning", flattenAntivirusQuarantineStoreMachineLearning(o["store-machine-learning"], d, "store_machine_learning", sv)); err != nil {
		if !fortiAPIPatch(o["store-machine-learning"]) {
			return fmt.Errorf("Error reading store_machine_learning: %v", err)
		}
	}

	if err = d.Set("drop_heuristic", flattenAntivirusQuarantineDropHeuristic(o["drop-heuristic"], d, "drop_heuristic", sv)); err != nil {
		if !fortiAPIPatch(o["drop-heuristic"]) {
			return fmt.Errorf("Error reading drop_heuristic: %v", err)
		}
	}

	if err = d.Set("store_heuristic", flattenAntivirusQuarantineStoreHeuristic(o["store-heuristic"], d, "store_heuristic", sv)); err != nil {
		if !fortiAPIPatch(o["store-heuristic"]) {
			return fmt.Errorf("Error reading store_heuristic: %v", err)
		}
	}

	if err = d.Set("lowspace", flattenAntivirusQuarantineLowspace(o["lowspace"], d, "lowspace", sv)); err != nil {
		if !fortiAPIPatch(o["lowspace"]) {
			return fmt.Errorf("Error reading lowspace: %v", err)
		}
	}

	if err = d.Set("destination", flattenAntivirusQuarantineDestination(o["destination"], d, "destination", sv)); err != nil {
		if !fortiAPIPatch(o["destination"]) {
			return fmt.Errorf("Error reading destination: %v", err)
		}
	}

	return nil
}

func flattenAntivirusQuarantineFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandAntivirusQuarantineAgelimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineMaxfilesize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineQuarantineQuota(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineDropInfected(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineStoreInfected(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineDropBlocked(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineStoreBlocked(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineDropMachineLearning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineStoreMachineLearning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineDropHeuristic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineStoreHeuristic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineLowspace(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusQuarantineDestination(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectAntivirusQuarantine(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("agelimit"); ok {
		if setArgNil {
			obj["agelimit"] = nil
		} else {

			t, err := expandAntivirusQuarantineAgelimit(d, v, "agelimit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["agelimit"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("maxfilesize"); ok {
		if setArgNil {
			obj["maxfilesize"] = nil
		} else {

			t, err := expandAntivirusQuarantineMaxfilesize(d, v, "maxfilesize", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["maxfilesize"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("quarantine_quota"); ok {
		if setArgNil {
			obj["quarantine-quota"] = nil
		} else {

			t, err := expandAntivirusQuarantineQuarantineQuota(d, v, "quarantine_quota", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["quarantine-quota"] = t
			}
		}
	}

	if v, ok := d.GetOk("drop_infected"); ok {
		if setArgNil {
			obj["drop-infected"] = nil
		} else {

			t, err := expandAntivirusQuarantineDropInfected(d, v, "drop_infected", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["drop-infected"] = t
			}
		}
	}

	if v, ok := d.GetOk("store_infected"); ok {
		if setArgNil {
			obj["store-infected"] = nil
		} else {

			t, err := expandAntivirusQuarantineStoreInfected(d, v, "store_infected", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["store-infected"] = t
			}
		}
	}

	if v, ok := d.GetOk("drop_blocked"); ok {
		if setArgNil {
			obj["drop-blocked"] = nil
		} else {

			t, err := expandAntivirusQuarantineDropBlocked(d, v, "drop_blocked", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["drop-blocked"] = t
			}
		}
	}

	if v, ok := d.GetOk("store_blocked"); ok {
		if setArgNil {
			obj["store-blocked"] = nil
		} else {

			t, err := expandAntivirusQuarantineStoreBlocked(d, v, "store_blocked", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["store-blocked"] = t
			}
		}
	}

	if v, ok := d.GetOk("drop_machine_learning"); ok {
		if setArgNil {
			obj["drop-machine-learning"] = nil
		} else {

			t, err := expandAntivirusQuarantineDropMachineLearning(d, v, "drop_machine_learning", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["drop-machine-learning"] = t
			}
		}
	}

	if v, ok := d.GetOk("store_machine_learning"); ok {
		if setArgNil {
			obj["store-machine-learning"] = nil
		} else {

			t, err := expandAntivirusQuarantineStoreMachineLearning(d, v, "store_machine_learning", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["store-machine-learning"] = t
			}
		}
	}

	if v, ok := d.GetOk("drop_heuristic"); ok {
		if setArgNil {
			obj["drop-heuristic"] = nil
		} else {

			t, err := expandAntivirusQuarantineDropHeuristic(d, v, "drop_heuristic", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["drop-heuristic"] = t
			}
		}
	}

	if v, ok := d.GetOk("store_heuristic"); ok {
		if setArgNil {
			obj["store-heuristic"] = nil
		} else {

			t, err := expandAntivirusQuarantineStoreHeuristic(d, v, "store_heuristic", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["store-heuristic"] = t
			}
		}
	}

	if v, ok := d.GetOk("lowspace"); ok {
		if setArgNil {
			obj["lowspace"] = nil
		} else {

			t, err := expandAntivirusQuarantineLowspace(d, v, "lowspace", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["lowspace"] = t
			}
		}
	}

	if v, ok := d.GetOk("destination"); ok {
		if setArgNil {
			obj["destination"] = nil
		} else {

			t, err := expandAntivirusQuarantineDestination(d, v, "destination", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["destination"] = t
			}
		}
	}

	return &obj, nil
}
