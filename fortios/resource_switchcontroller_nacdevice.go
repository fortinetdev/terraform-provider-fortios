// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure/list NAC devices learned on the managed FortiSwitch ports which matches NAC policy.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerNacDevice() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerNacDeviceCreate,
		Read:   resourceSwitchControllerNacDeviceRead,
		Update: resourceSwitchControllerNacDeviceUpdate,
		Delete: resourceSwitchControllerNacDeviceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_known_switch": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"last_known_port": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"matched_nac_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"port_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"mac_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"last_seen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerNacDeviceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerNacDevice(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerNacDevice resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerNacDevice(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerNacDevice resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SwitchControllerNacDevice")
	}

	return resourceSwitchControllerNacDeviceRead(d, m)
}

func resourceSwitchControllerNacDeviceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerNacDevice(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerNacDevice resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerNacDevice(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerNacDevice resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SwitchControllerNacDevice")
	}

	return resourceSwitchControllerNacDeviceRead(d, m)
}

func resourceSwitchControllerNacDeviceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerNacDevice(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerNacDevice resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerNacDeviceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerNacDevice(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerNacDevice resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerNacDevice(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerNacDevice resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerNacDeviceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerNacDeviceDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerNacDeviceStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerNacDeviceMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerNacDeviceLastKnownSwitch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerNacDeviceLastKnownPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerNacDeviceMatchedNacPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerNacDevicePortPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerNacDeviceMacPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerNacDeviceLastSeen(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerNacDevice(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenSwitchControllerNacDeviceId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerNacDeviceDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("status", flattenSwitchControllerNacDeviceStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("mac", flattenSwitchControllerNacDeviceMac(o["mac"], d, "mac", sv)); err != nil {
		if !fortiAPIPatch(o["mac"]) {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("last_known_switch", flattenSwitchControllerNacDeviceLastKnownSwitch(o["last-known-switch"], d, "last_known_switch", sv)); err != nil {
		if !fortiAPIPatch(o["last-known-switch"]) {
			return fmt.Errorf("Error reading last_known_switch: %v", err)
		}
	}

	if err = d.Set("last_known_port", flattenSwitchControllerNacDeviceLastKnownPort(o["last-known-port"], d, "last_known_port", sv)); err != nil {
		if !fortiAPIPatch(o["last-known-port"]) {
			return fmt.Errorf("Error reading last_known_port: %v", err)
		}
	}

	if err = d.Set("matched_nac_policy", flattenSwitchControllerNacDeviceMatchedNacPolicy(o["matched-nac-policy"], d, "matched_nac_policy", sv)); err != nil {
		if !fortiAPIPatch(o["matched-nac-policy"]) {
			return fmt.Errorf("Error reading matched_nac_policy: %v", err)
		}
	}

	if err = d.Set("port_policy", flattenSwitchControllerNacDevicePortPolicy(o["port-policy"], d, "port_policy", sv)); err != nil {
		if !fortiAPIPatch(o["port-policy"]) {
			return fmt.Errorf("Error reading port_policy: %v", err)
		}
	}

	if err = d.Set("mac_policy", flattenSwitchControllerNacDeviceMacPolicy(o["mac-policy"], d, "mac_policy", sv)); err != nil {
		if !fortiAPIPatch(o["mac-policy"]) {
			return fmt.Errorf("Error reading mac_policy: %v", err)
		}
	}

	if err = d.Set("last_seen", flattenSwitchControllerNacDeviceLastSeen(o["last-seen"], d, "last_seen", sv)); err != nil {
		if !fortiAPIPatch(o["last-seen"]) {
			return fmt.Errorf("Error reading last_seen: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerNacDeviceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerNacDeviceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerNacDeviceDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerNacDeviceStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerNacDeviceMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerNacDeviceLastKnownSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerNacDeviceLastKnownPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerNacDeviceMatchedNacPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerNacDevicePortPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerNacDeviceMacPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerNacDeviceLastSeen(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerNacDevice(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandSwitchControllerNacDeviceId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandSwitchControllerNacDeviceDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSwitchControllerNacDeviceStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok {
		t, err := expandSwitchControllerNacDeviceMac(d, v, "mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("last_known_switch"); ok {
		t, err := expandSwitchControllerNacDeviceLastKnownSwitch(d, v, "last_known_switch", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-known-switch"] = t
		}
	}

	if v, ok := d.GetOk("last_known_port"); ok {
		t, err := expandSwitchControllerNacDeviceLastKnownPort(d, v, "last_known_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-known-port"] = t
		}
	}

	if v, ok := d.GetOk("matched_nac_policy"); ok {
		t, err := expandSwitchControllerNacDeviceMatchedNacPolicy(d, v, "matched_nac_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["matched-nac-policy"] = t
		}
	}

	if v, ok := d.GetOk("port_policy"); ok {
		t, err := expandSwitchControllerNacDevicePortPolicy(d, v, "port_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-policy"] = t
		}
	}

	if v, ok := d.GetOk("mac_policy"); ok {
		t, err := expandSwitchControllerNacDeviceMacPolicy(d, v, "mac_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-policy"] = t
		}
	}

	if v, ok := d.GetOkExists("last_seen"); ok {
		t, err := expandSwitchControllerNacDeviceLastSeen(d, v, "last_seen", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-seen"] = t
		}
	}

	return &obj, nil
}
