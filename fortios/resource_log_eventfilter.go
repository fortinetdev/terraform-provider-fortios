// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure log event filters.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceLogEventfilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogEventfilterUpdate,
		Read:   resourceLogEventfilterRead,
		Update: resourceLogEventfilterUpdate,
		Delete: resourceLogEventfilterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"event": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"router": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wireless_activity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wan_opt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"compliance_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_rating": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiextender": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connector": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdwan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cifs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogEventfilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogEventfilter(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogEventfilter resource while getting object: %v", err)
	}

	o, err := c.UpdateLogEventfilter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogEventfilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogEventfilter")
	}

	return resourceLogEventfilterRead(d, m)
}

func resourceLogEventfilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteLogEventfilter(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting LogEventfilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogEventfilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadLogEventfilter(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogEventfilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogEventfilter(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogEventfilter resource from API: %v", err)
	}
	return nil
}

func flattenLogEventfilterEvent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterSystem(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterVpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterRouter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterWirelessActivity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterWanOpt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterEndpoint(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterHa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterComplianceCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterSecurityRating(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterFortiextender(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterConnector(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterSdwan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterCifs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogEventfilterSwitchController(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogEventfilter(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("event", flattenLogEventfilterEvent(o["event"], d, "event", sv)); err != nil {
		if !fortiAPIPatch(o["event"]) {
			return fmt.Errorf("Error reading event: %v", err)
		}
	}

	if err = d.Set("system", flattenLogEventfilterSystem(o["system"], d, "system", sv)); err != nil {
		if !fortiAPIPatch(o["system"]) {
			return fmt.Errorf("Error reading system: %v", err)
		}
	}

	if err = d.Set("vpn", flattenLogEventfilterVpn(o["vpn"], d, "vpn", sv)); err != nil {
		if !fortiAPIPatch(o["vpn"]) {
			return fmt.Errorf("Error reading vpn: %v", err)
		}
	}

	if err = d.Set("user", flattenLogEventfilterUser(o["user"], d, "user", sv)); err != nil {
		if !fortiAPIPatch(o["user"]) {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("router", flattenLogEventfilterRouter(o["router"], d, "router", sv)); err != nil {
		if !fortiAPIPatch(o["router"]) {
			return fmt.Errorf("Error reading router: %v", err)
		}
	}

	if err = d.Set("wireless_activity", flattenLogEventfilterWirelessActivity(o["wireless-activity"], d, "wireless_activity", sv)); err != nil {
		if !fortiAPIPatch(o["wireless-activity"]) {
			return fmt.Errorf("Error reading wireless_activity: %v", err)
		}
	}

	if err = d.Set("wan_opt", flattenLogEventfilterWanOpt(o["wan-opt"], d, "wan_opt", sv)); err != nil {
		if !fortiAPIPatch(o["wan-opt"]) {
			return fmt.Errorf("Error reading wan_opt: %v", err)
		}
	}

	if err = d.Set("endpoint", flattenLogEventfilterEndpoint(o["endpoint"], d, "endpoint", sv)); err != nil {
		if !fortiAPIPatch(o["endpoint"]) {
			return fmt.Errorf("Error reading endpoint: %v", err)
		}
	}

	if err = d.Set("ha", flattenLogEventfilterHa(o["ha"], d, "ha", sv)); err != nil {
		if !fortiAPIPatch(o["ha"]) {
			return fmt.Errorf("Error reading ha: %v", err)
		}
	}

	if err = d.Set("compliance_check", flattenLogEventfilterComplianceCheck(o["compliance-check"], d, "compliance_check", sv)); err != nil {
		if !fortiAPIPatch(o["compliance-check"]) {
			return fmt.Errorf("Error reading compliance_check: %v", err)
		}
	}

	if err = d.Set("security_rating", flattenLogEventfilterSecurityRating(o["security-rating"], d, "security_rating", sv)); err != nil {
		if !fortiAPIPatch(o["security-rating"]) {
			return fmt.Errorf("Error reading security_rating: %v", err)
		}
	}

	if err = d.Set("fortiextender", flattenLogEventfilterFortiextender(o["fortiextender"], d, "fortiextender", sv)); err != nil {
		if !fortiAPIPatch(o["fortiextender"]) {
			return fmt.Errorf("Error reading fortiextender: %v", err)
		}
	}

	if err = d.Set("connector", flattenLogEventfilterConnector(o["connector"], d, "connector", sv)); err != nil {
		if !fortiAPIPatch(o["connector"]) {
			return fmt.Errorf("Error reading connector: %v", err)
		}
	}

	if err = d.Set("sdwan", flattenLogEventfilterSdwan(o["sdwan"], d, "sdwan", sv)); err != nil {
		if !fortiAPIPatch(o["sdwan"]) {
			return fmt.Errorf("Error reading sdwan: %v", err)
		}
	}

	if err = d.Set("cifs", flattenLogEventfilterCifs(o["cifs"], d, "cifs", sv)); err != nil {
		if !fortiAPIPatch(o["cifs"]) {
			return fmt.Errorf("Error reading cifs: %v", err)
		}
	}

	if err = d.Set("switch_controller", flattenLogEventfilterSwitchController(o["switch-controller"], d, "switch_controller", sv)); err != nil {
		if !fortiAPIPatch(o["switch-controller"]) {
			return fmt.Errorf("Error reading switch_controller: %v", err)
		}
	}

	return nil
}

func flattenLogEventfilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogEventfilterEvent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterSystem(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterVpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterRouter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterWirelessActivity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterWanOpt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterEndpoint(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterHa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterComplianceCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterSecurityRating(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterFortiextender(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterConnector(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterSdwan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterCifs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterSwitchController(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogEventfilter(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("event"); ok {

		t, err := expandLogEventfilterEvent(d, v, "event", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["event"] = t
		}
	}

	if v, ok := d.GetOk("system"); ok {

		t, err := expandLogEventfilterSystem(d, v, "system", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system"] = t
		}
	}

	if v, ok := d.GetOk("vpn"); ok {

		t, err := expandLogEventfilterVpn(d, v, "vpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok {

		t, err := expandLogEventfilterUser(d, v, "user", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOk("router"); ok {

		t, err := expandLogEventfilterRouter(d, v, "router", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["router"] = t
		}
	}

	if v, ok := d.GetOk("wireless_activity"); ok {

		t, err := expandLogEventfilterWirelessActivity(d, v, "wireless_activity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wireless-activity"] = t
		}
	}

	if v, ok := d.GetOk("wan_opt"); ok {

		t, err := expandLogEventfilterWanOpt(d, v, "wan_opt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-opt"] = t
		}
	}

	if v, ok := d.GetOk("endpoint"); ok {

		t, err := expandLogEventfilterEndpoint(d, v, "endpoint", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endpoint"] = t
		}
	}

	if v, ok := d.GetOk("ha"); ok {

		t, err := expandLogEventfilterHa(d, v, "ha", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha"] = t
		}
	}

	if v, ok := d.GetOk("compliance_check"); ok {

		t, err := expandLogEventfilterComplianceCheck(d, v, "compliance_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compliance-check"] = t
		}
	}

	if v, ok := d.GetOk("security_rating"); ok {

		t, err := expandLogEventfilterSecurityRating(d, v, "security_rating", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-rating"] = t
		}
	}

	if v, ok := d.GetOk("fortiextender"); ok {

		t, err := expandLogEventfilterFortiextender(d, v, "fortiextender", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiextender"] = t
		}
	}

	if v, ok := d.GetOk("connector"); ok {

		t, err := expandLogEventfilterConnector(d, v, "connector", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connector"] = t
		}
	}

	if v, ok := d.GetOk("sdwan"); ok {

		t, err := expandLogEventfilterSdwan(d, v, "sdwan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdwan"] = t
		}
	}

	if v, ok := d.GetOk("cifs"); ok {

		t, err := expandLogEventfilterCifs(d, v, "cifs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller"); ok {

		t, err := expandLogEventfilterSwitchController(d, v, "switch_controller", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller"] = t
		}
	}

	return &obj, nil
}
