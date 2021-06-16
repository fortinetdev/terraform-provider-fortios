// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure per-IP traffic shaper.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallShaperPerIpShaper() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallShaperPerIpShaperCreate,
		Read:   resourceFirewallShaperPerIpShaperRead,
		Update: resourceFirewallShaperPerIpShaperUpdate,
		Delete: resourceFirewallShaperPerIpShaperDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"max_bandwidth": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),
				Optional:     true,
				Computed:     true,
			},
			"bandwidth_unit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_concurrent_session": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2097000),
				Optional:     true,
				Computed:     true,
			},
			"max_concurrent_tcp_session": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2097000),
				Optional:     true,
				Computed:     true,
			},
			"max_concurrent_udp_session": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2097000),
				Optional:     true,
				Computed:     true,
			},
			"diffserv_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv_reverse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode_rev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallShaperPerIpShaperCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallShaperPerIpShaper(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallShaperPerIpShaper resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallShaperPerIpShaper(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallShaperPerIpShaper resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallShaperPerIpShaper")
	}

	return resourceFirewallShaperPerIpShaperRead(d, m)
}

func resourceFirewallShaperPerIpShaperUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallShaperPerIpShaper(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallShaperPerIpShaper resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallShaperPerIpShaper(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallShaperPerIpShaper resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallShaperPerIpShaper")
	}

	return resourceFirewallShaperPerIpShaperRead(d, m)
}

func resourceFirewallShaperPerIpShaperDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallShaperPerIpShaper(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallShaperPerIpShaper resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallShaperPerIpShaperRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallShaperPerIpShaper(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallShaperPerIpShaper resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallShaperPerIpShaper(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallShaperPerIpShaper resource from API: %v", err)
	}
	return nil
}

func flattenFirewallShaperPerIpShaperName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperPerIpShaperMaxBandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperPerIpShaperBandwidthUnit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperPerIpShaperMaxConcurrentSession(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperPerIpShaperMaxConcurrentTcpSession(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperPerIpShaperMaxConcurrentUdpSession(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperPerIpShaperDiffservForward(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperPerIpShaperDiffservReverse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperPerIpShaperDiffservcodeForward(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShaperPerIpShaperDiffservcodeRev(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallShaperPerIpShaper(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenFirewallShaperPerIpShaperName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("max_bandwidth", flattenFirewallShaperPerIpShaperMaxBandwidth(o["max-bandwidth"], d, "max_bandwidth", sv)); err != nil {
		if !fortiAPIPatch(o["max-bandwidth"]) {
			return fmt.Errorf("Error reading max_bandwidth: %v", err)
		}
	}

	if err = d.Set("bandwidth_unit", flattenFirewallShaperPerIpShaperBandwidthUnit(o["bandwidth-unit"], d, "bandwidth_unit", sv)); err != nil {
		if !fortiAPIPatch(o["bandwidth-unit"]) {
			return fmt.Errorf("Error reading bandwidth_unit: %v", err)
		}
	}

	if err = d.Set("max_concurrent_session", flattenFirewallShaperPerIpShaperMaxConcurrentSession(o["max-concurrent-session"], d, "max_concurrent_session", sv)); err != nil {
		if !fortiAPIPatch(o["max-concurrent-session"]) {
			return fmt.Errorf("Error reading max_concurrent_session: %v", err)
		}
	}

	if err = d.Set("max_concurrent_tcp_session", flattenFirewallShaperPerIpShaperMaxConcurrentTcpSession(o["max-concurrent-tcp-session"], d, "max_concurrent_tcp_session", sv)); err != nil {
		if !fortiAPIPatch(o["max-concurrent-tcp-session"]) {
			return fmt.Errorf("Error reading max_concurrent_tcp_session: %v", err)
		}
	}

	if err = d.Set("max_concurrent_udp_session", flattenFirewallShaperPerIpShaperMaxConcurrentUdpSession(o["max-concurrent-udp-session"], d, "max_concurrent_udp_session", sv)); err != nil {
		if !fortiAPIPatch(o["max-concurrent-udp-session"]) {
			return fmt.Errorf("Error reading max_concurrent_udp_session: %v", err)
		}
	}

	if err = d.Set("diffserv_forward", flattenFirewallShaperPerIpShaperDiffservForward(o["diffserv-forward"], d, "diffserv_forward", sv)); err != nil {
		if !fortiAPIPatch(o["diffserv-forward"]) {
			return fmt.Errorf("Error reading diffserv_forward: %v", err)
		}
	}

	if err = d.Set("diffserv_reverse", flattenFirewallShaperPerIpShaperDiffservReverse(o["diffserv-reverse"], d, "diffserv_reverse", sv)); err != nil {
		if !fortiAPIPatch(o["diffserv-reverse"]) {
			return fmt.Errorf("Error reading diffserv_reverse: %v", err)
		}
	}

	if err = d.Set("diffservcode_forward", flattenFirewallShaperPerIpShaperDiffservcodeForward(o["diffservcode-forward"], d, "diffservcode_forward", sv)); err != nil {
		if !fortiAPIPatch(o["diffservcode-forward"]) {
			return fmt.Errorf("Error reading diffservcode_forward: %v", err)
		}
	}

	if err = d.Set("diffservcode_rev", flattenFirewallShaperPerIpShaperDiffservcodeRev(o["diffservcode-rev"], d, "diffservcode_rev", sv)); err != nil {
		if !fortiAPIPatch(o["diffservcode-rev"]) {
			return fmt.Errorf("Error reading diffservcode_rev: %v", err)
		}
	}

	return nil
}

func flattenFirewallShaperPerIpShaperFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallShaperPerIpShaperName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperPerIpShaperMaxBandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperPerIpShaperBandwidthUnit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperPerIpShaperMaxConcurrentSession(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperPerIpShaperMaxConcurrentTcpSession(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperPerIpShaperMaxConcurrentUdpSession(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperPerIpShaperDiffservForward(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperPerIpShaperDiffservReverse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperPerIpShaperDiffservcodeForward(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperPerIpShaperDiffservcodeRev(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallShaperPerIpShaper(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallShaperPerIpShaperName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("max_bandwidth"); ok {

		t, err := expandFirewallShaperPerIpShaperMaxBandwidth(d, v, "max_bandwidth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_unit"); ok {

		t, err := expandFirewallShaperPerIpShaperBandwidthUnit(d, v, "bandwidth_unit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-unit"] = t
		}
	}

	if v, ok := d.GetOkExists("max_concurrent_session"); ok {

		t, err := expandFirewallShaperPerIpShaperMaxConcurrentSession(d, v, "max_concurrent_session", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-concurrent-session"] = t
		}
	}

	if v, ok := d.GetOkExists("max_concurrent_tcp_session"); ok {

		t, err := expandFirewallShaperPerIpShaperMaxConcurrentTcpSession(d, v, "max_concurrent_tcp_session", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-concurrent-tcp-session"] = t
		}
	}

	if v, ok := d.GetOkExists("max_concurrent_udp_session"); ok {

		t, err := expandFirewallShaperPerIpShaperMaxConcurrentUdpSession(d, v, "max_concurrent_udp_session", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-concurrent-udp-session"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_forward"); ok {

		t, err := expandFirewallShaperPerIpShaperDiffservForward(d, v, "diffserv_forward", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_reverse"); ok {

		t, err := expandFirewallShaperPerIpShaperDiffservReverse(d, v, "diffserv_reverse", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-reverse"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_forward"); ok {

		t, err := expandFirewallShaperPerIpShaperDiffservcodeForward(d, v, "diffservcode_forward", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_rev"); ok {

		t, err := expandFirewallShaperPerIpShaperDiffservcodeRev(d, v, "diffservcode_rev", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-rev"] = t
		}
	}

	return &obj, nil
}
