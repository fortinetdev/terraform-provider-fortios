// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch storm control.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerStormControl() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerStormControlUpdate,
		Read:   resourceSwitchControllerStormControlRead,
		Update: resourceSwitchControllerStormControlUpdate,
		Delete: resourceSwitchControllerStormControlDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"rate": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10000000),
				Optional:     true,
				Computed:     true,
			},
			"unknown_unicast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_multicast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"broadcast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerStormControlUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerStormControl(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerStormControl resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerStormControl(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerStormControl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerStormControl")
	}

	return resourceSwitchControllerStormControlRead(d, m)
}

func resourceSwitchControllerStormControlDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerStormControl(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerStormControl resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerStormControl(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchControllerStormControl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerStormControlRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerStormControl(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerStormControl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerStormControl(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerStormControl resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerStormControlRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerStormControlUnknownUnicast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerStormControlUnknownMulticast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerStormControlBroadcast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerStormControl(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("rate", flattenSwitchControllerStormControlRate(o["rate"], d, "rate", sv)); err != nil {
		if !fortiAPIPatch(o["rate"]) {
			return fmt.Errorf("Error reading rate: %v", err)
		}
	}

	if err = d.Set("unknown_unicast", flattenSwitchControllerStormControlUnknownUnicast(o["unknown-unicast"], d, "unknown_unicast", sv)); err != nil {
		if !fortiAPIPatch(o["unknown-unicast"]) {
			return fmt.Errorf("Error reading unknown_unicast: %v", err)
		}
	}

	if err = d.Set("unknown_multicast", flattenSwitchControllerStormControlUnknownMulticast(o["unknown-multicast"], d, "unknown_multicast", sv)); err != nil {
		if !fortiAPIPatch(o["unknown-multicast"]) {
			return fmt.Errorf("Error reading unknown_multicast: %v", err)
		}
	}

	if err = d.Set("broadcast", flattenSwitchControllerStormControlBroadcast(o["broadcast"], d, "broadcast", sv)); err != nil {
		if !fortiAPIPatch(o["broadcast"]) {
			return fmt.Errorf("Error reading broadcast: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerStormControlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerStormControlRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStormControlUnknownUnicast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStormControlUnknownMulticast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStormControlBroadcast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerStormControl(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("rate"); ok {
		if setArgNil {
			obj["rate"] = nil
		} else {

			t, err := expandSwitchControllerStormControlRate(d, v, "rate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["rate"] = t
			}
		}
	}

	if v, ok := d.GetOk("unknown_unicast"); ok {
		if setArgNil {
			obj["unknown-unicast"] = nil
		} else {

			t, err := expandSwitchControllerStormControlUnknownUnicast(d, v, "unknown_unicast", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["unknown-unicast"] = t
			}
		}
	}

	if v, ok := d.GetOk("unknown_multicast"); ok {
		if setArgNil {
			obj["unknown-multicast"] = nil
		} else {

			t, err := expandSwitchControllerStormControlUnknownMulticast(d, v, "unknown_multicast", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["unknown-multicast"] = t
			}
		}
	}

	if v, ok := d.GetOk("broadcast"); ok {
		if setArgNil {
			obj["broadcast"] = nil
		} else {

			t, err := expandSwitchControllerStormControlBroadcast(d, v, "broadcast", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["broadcast"] = t
			}
		}
	}

	return &obj, nil
}
