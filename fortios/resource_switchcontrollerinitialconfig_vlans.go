// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure initial template for auto-generated VLAN interfaces.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerInitialConfigVlans() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerInitialConfigVlansUpdate,
		Read:   resourceSwitchControllerInitialConfigVlansRead,
		Update: resourceSwitchControllerInitialConfigVlansUpdate,
		Delete: resourceSwitchControllerInitialConfigVlansDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"default_vlan": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"quarantine": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"rspan": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"voice": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"video": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"nac": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"nac_segment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSwitchControllerInitialConfigVlansUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerInitialConfigVlans(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerInitialConfigVlans resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerInitialConfigVlans(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerInitialConfigVlans resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerInitialConfigVlans")
	}

	return resourceSwitchControllerInitialConfigVlansRead(d, m)
}

func resourceSwitchControllerInitialConfigVlansDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerInitialConfigVlans(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerInitialConfigVlans resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerInitialConfigVlansRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerInitialConfigVlans(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerInitialConfigVlans resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerInitialConfigVlans(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerInitialConfigVlans resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerInitialConfigVlansDefaultVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerInitialConfigVlansQuarantine(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerInitialConfigVlansRspan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerInitialConfigVlansVoice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerInitialConfigVlansVideo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerInitialConfigVlansNac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerInitialConfigVlansNacSegment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerInitialConfigVlans(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("default_vlan", flattenSwitchControllerInitialConfigVlansDefaultVlan(o["default-vlan"], d, "default_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["default-vlan"]) {
			return fmt.Errorf("Error reading default_vlan: %v", err)
		}
	}

	if err = d.Set("quarantine", flattenSwitchControllerInitialConfigVlansQuarantine(o["quarantine"], d, "quarantine", sv)); err != nil {
		if !fortiAPIPatch(o["quarantine"]) {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if err = d.Set("rspan", flattenSwitchControllerInitialConfigVlansRspan(o["rspan"], d, "rspan", sv)); err != nil {
		if !fortiAPIPatch(o["rspan"]) {
			return fmt.Errorf("Error reading rspan: %v", err)
		}
	}

	if err = d.Set("voice", flattenSwitchControllerInitialConfigVlansVoice(o["voice"], d, "voice", sv)); err != nil {
		if !fortiAPIPatch(o["voice"]) {
			return fmt.Errorf("Error reading voice: %v", err)
		}
	}

	if err = d.Set("video", flattenSwitchControllerInitialConfigVlansVideo(o["video"], d, "video", sv)); err != nil {
		if !fortiAPIPatch(o["video"]) {
			return fmt.Errorf("Error reading video: %v", err)
		}
	}

	if err = d.Set("nac", flattenSwitchControllerInitialConfigVlansNac(o["nac"], d, "nac", sv)); err != nil {
		if !fortiAPIPatch(o["nac"]) {
			return fmt.Errorf("Error reading nac: %v", err)
		}
	}

	if err = d.Set("nac_segment", flattenSwitchControllerInitialConfigVlansNacSegment(o["nac-segment"], d, "nac_segment", sv)); err != nil {
		if !fortiAPIPatch(o["nac-segment"]) {
			return fmt.Errorf("Error reading nac_segment: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerInitialConfigVlansFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerInitialConfigVlansDefaultVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerInitialConfigVlansQuarantine(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerInitialConfigVlansRspan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerInitialConfigVlansVoice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerInitialConfigVlansVideo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerInitialConfigVlansNac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerInitialConfigVlansNacSegment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerInitialConfigVlans(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default_vlan"); ok {

		t, err := expandSwitchControllerInitialConfigVlansDefaultVlan(d, v, "default_vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-vlan"] = t
		}
	}

	if v, ok := d.GetOk("quarantine"); ok {

		t, err := expandSwitchControllerInitialConfigVlansQuarantine(d, v, "quarantine", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	if v, ok := d.GetOk("rspan"); ok {

		t, err := expandSwitchControllerInitialConfigVlansRspan(d, v, "rspan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rspan"] = t
		}
	}

	if v, ok := d.GetOk("voice"); ok {

		t, err := expandSwitchControllerInitialConfigVlansVoice(d, v, "voice", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voice"] = t
		}
	}

	if v, ok := d.GetOk("video"); ok {

		t, err := expandSwitchControllerInitialConfigVlansVideo(d, v, "video", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["video"] = t
		}
	}

	if v, ok := d.GetOk("nac"); ok {

		t, err := expandSwitchControllerInitialConfigVlansNac(d, v, "nac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac"] = t
		}
	}

	if v, ok := d.GetOk("nac_segment"); ok {

		t, err := expandSwitchControllerInitialConfigVlansNacSegment(d, v, "nac_segment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-segment"] = t
		}
	}

	return &obj, nil
}
