// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch IGMP snooping global settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerIgmpSnooping() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerIgmpSnoopingUpdate,
		Read:   resourceSwitchControllerIgmpSnoopingRead,
		Update: resourceSwitchControllerIgmpSnoopingUpdate,
		Delete: resourceSwitchControllerIgmpSnoopingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"aging_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(15, 3600),
				Optional:     true,
				Computed:     true,
			},
			"flood_unknown_multicast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerIgmpSnoopingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerIgmpSnooping(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerIgmpSnooping resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerIgmpSnooping(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerIgmpSnooping resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerIgmpSnooping")
	}

	return resourceSwitchControllerIgmpSnoopingRead(d, m)
}

func resourceSwitchControllerIgmpSnoopingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchControllerIgmpSnooping(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerIgmpSnooping resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerIgmpSnoopingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchControllerIgmpSnooping(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerIgmpSnooping resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerIgmpSnooping(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerIgmpSnooping resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerIgmpSnoopingAgingTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerIgmpSnoopingFloodUnknownMulticast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerIgmpSnooping(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("aging_time", flattenSwitchControllerIgmpSnoopingAgingTime(o["aging-time"], d, "aging_time")); err != nil {
		if !fortiAPIPatch(o["aging-time"]) {
			return fmt.Errorf("Error reading aging_time: %v", err)
		}
	}

	if err = d.Set("flood_unknown_multicast", flattenSwitchControllerIgmpSnoopingFloodUnknownMulticast(o["flood-unknown-multicast"], d, "flood_unknown_multicast")); err != nil {
		if !fortiAPIPatch(o["flood-unknown-multicast"]) {
			return fmt.Errorf("Error reading flood_unknown_multicast: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerIgmpSnoopingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerIgmpSnoopingAgingTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerIgmpSnoopingFloodUnknownMulticast(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerIgmpSnooping(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("aging_time"); ok {
		t, err := expandSwitchControllerIgmpSnoopingAgingTime(d, v, "aging_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aging-time"] = t
		}
	}

	if v, ok := d.GetOk("flood_unknown_multicast"); ok {
		t, err := expandSwitchControllerIgmpSnoopingFloodUnknownMulticast(d, v, "flood_unknown_multicast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flood-unknown-multicast"] = t
		}
	}

	return &obj, nil
}
