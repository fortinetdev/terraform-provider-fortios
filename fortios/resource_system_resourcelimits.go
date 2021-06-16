// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure resource limits.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemResourceLimits() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemResourceLimitsUpdate,
		Read:   resourceSystemResourceLimitsRead,
		Update: resourceSystemResourceLimitsUpdate,
		Delete: resourceSystemResourceLimitsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"session": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipsec_phase1": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipsec_phase2": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipsec_phase1_interface": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipsec_phase2_interface": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dialup_tunnel": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"firewall_policy": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"firewall_address": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"firewall_addrgrp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"custom_service": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"service_group": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"onetime_schedule": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"recurring_schedule": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"user_group": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sslvpn": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"proxy": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_disk_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemResourceLimitsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemResourceLimits(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemResourceLimits resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemResourceLimits(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemResourceLimits resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemResourceLimits")
	}

	return resourceSystemResourceLimitsRead(d, m)
}

func resourceSystemResourceLimitsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemResourceLimits(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemResourceLimits resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemResourceLimitsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemResourceLimits(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemResourceLimits resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemResourceLimits(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemResourceLimits resource from API: %v", err)
	}
	return nil
}

func flattenSystemResourceLimitsSession(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemResourceLimitsIpsecPhase1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemResourceLimitsIpsecPhase2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemResourceLimitsIpsecPhase1Interface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemResourceLimitsIpsecPhase2Interface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemResourceLimitsDialupTunnel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemResourceLimitsFirewallPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemResourceLimitsFirewallAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemResourceLimitsFirewallAddrgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemResourceLimitsCustomService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemResourceLimitsServiceGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemResourceLimitsOnetimeSchedule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemResourceLimitsRecurringSchedule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemResourceLimitsUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemResourceLimitsUserGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemResourceLimitsSslvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemResourceLimitsProxy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemResourceLimitsLogDiskQuota(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemResourceLimits(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("session", flattenSystemResourceLimitsSession(o["session"], d, "session", sv)); err != nil {
		if !fortiAPIPatch(o["session"]) {
			return fmt.Errorf("Error reading session: %v", err)
		}
	}

	if err = d.Set("ipsec_phase1", flattenSystemResourceLimitsIpsecPhase1(o["ipsec-phase1"], d, "ipsec_phase1", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-phase1"]) {
			return fmt.Errorf("Error reading ipsec_phase1: %v", err)
		}
	}

	if err = d.Set("ipsec_phase2", flattenSystemResourceLimitsIpsecPhase2(o["ipsec-phase2"], d, "ipsec_phase2", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-phase2"]) {
			return fmt.Errorf("Error reading ipsec_phase2: %v", err)
		}
	}

	if err = d.Set("ipsec_phase1_interface", flattenSystemResourceLimitsIpsecPhase1Interface(o["ipsec-phase1-interface"], d, "ipsec_phase1_interface", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-phase1-interface"]) {
			return fmt.Errorf("Error reading ipsec_phase1_interface: %v", err)
		}
	}

	if err = d.Set("ipsec_phase2_interface", flattenSystemResourceLimitsIpsecPhase2Interface(o["ipsec-phase2-interface"], d, "ipsec_phase2_interface", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-phase2-interface"]) {
			return fmt.Errorf("Error reading ipsec_phase2_interface: %v", err)
		}
	}

	if err = d.Set("dialup_tunnel", flattenSystemResourceLimitsDialupTunnel(o["dialup-tunnel"], d, "dialup_tunnel", sv)); err != nil {
		if !fortiAPIPatch(o["dialup-tunnel"]) {
			return fmt.Errorf("Error reading dialup_tunnel: %v", err)
		}
	}

	if err = d.Set("firewall_policy", flattenSystemResourceLimitsFirewallPolicy(o["firewall-policy"], d, "firewall_policy", sv)); err != nil {
		if !fortiAPIPatch(o["firewall-policy"]) {
			return fmt.Errorf("Error reading firewall_policy: %v", err)
		}
	}

	if err = d.Set("firewall_address", flattenSystemResourceLimitsFirewallAddress(o["firewall-address"], d, "firewall_address", sv)); err != nil {
		if !fortiAPIPatch(o["firewall-address"]) {
			return fmt.Errorf("Error reading firewall_address: %v", err)
		}
	}

	if err = d.Set("firewall_addrgrp", flattenSystemResourceLimitsFirewallAddrgrp(o["firewall-addrgrp"], d, "firewall_addrgrp", sv)); err != nil {
		if !fortiAPIPatch(o["firewall-addrgrp"]) {
			return fmt.Errorf("Error reading firewall_addrgrp: %v", err)
		}
	}

	if err = d.Set("custom_service", flattenSystemResourceLimitsCustomService(o["custom-service"], d, "custom_service", sv)); err != nil {
		if !fortiAPIPatch(o["custom-service"]) {
			return fmt.Errorf("Error reading custom_service: %v", err)
		}
	}

	if err = d.Set("service_group", flattenSystemResourceLimitsServiceGroup(o["service-group"], d, "service_group", sv)); err != nil {
		if !fortiAPIPatch(o["service-group"]) {
			return fmt.Errorf("Error reading service_group: %v", err)
		}
	}

	if err = d.Set("onetime_schedule", flattenSystemResourceLimitsOnetimeSchedule(o["onetime-schedule"], d, "onetime_schedule", sv)); err != nil {
		if !fortiAPIPatch(o["onetime-schedule"]) {
			return fmt.Errorf("Error reading onetime_schedule: %v", err)
		}
	}

	if err = d.Set("recurring_schedule", flattenSystemResourceLimitsRecurringSchedule(o["recurring-schedule"], d, "recurring_schedule", sv)); err != nil {
		if !fortiAPIPatch(o["recurring-schedule"]) {
			return fmt.Errorf("Error reading recurring_schedule: %v", err)
		}
	}

	if err = d.Set("user", flattenSystemResourceLimitsUser(o["user"], d, "user", sv)); err != nil {
		if !fortiAPIPatch(o["user"]) {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("user_group", flattenSystemResourceLimitsUserGroup(o["user-group"], d, "user_group", sv)); err != nil {
		if !fortiAPIPatch(o["user-group"]) {
			return fmt.Errorf("Error reading user_group: %v", err)
		}
	}

	if err = d.Set("sslvpn", flattenSystemResourceLimitsSslvpn(o["sslvpn"], d, "sslvpn", sv)); err != nil {
		if !fortiAPIPatch(o["sslvpn"]) {
			return fmt.Errorf("Error reading sslvpn: %v", err)
		}
	}

	if err = d.Set("proxy", flattenSystemResourceLimitsProxy(o["proxy"], d, "proxy", sv)); err != nil {
		if !fortiAPIPatch(o["proxy"]) {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	if err = d.Set("log_disk_quota", flattenSystemResourceLimitsLogDiskQuota(o["log-disk-quota"], d, "log_disk_quota", sv)); err != nil {
		if !fortiAPIPatch(o["log-disk-quota"]) {
			return fmt.Errorf("Error reading log_disk_quota: %v", err)
		}
	}

	return nil
}

func flattenSystemResourceLimitsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemResourceLimitsSession(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsIpsecPhase1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsIpsecPhase2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsIpsecPhase1Interface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsIpsecPhase2Interface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsDialupTunnel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsFirewallPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsFirewallAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsFirewallAddrgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsCustomService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsServiceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsOnetimeSchedule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsRecurringSchedule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsUserGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsSslvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsProxy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsLogDiskQuota(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemResourceLimits(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("session"); ok {

		t, err := expandSystemResourceLimitsSession(d, v, "session", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session"] = t
		}
	}

	if v, ok := d.GetOkExists("ipsec_phase1"); ok {

		t, err := expandSystemResourceLimitsIpsecPhase1(d, v, "ipsec_phase1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-phase1"] = t
		}
	}

	if v, ok := d.GetOkExists("ipsec_phase2"); ok {

		t, err := expandSystemResourceLimitsIpsecPhase2(d, v, "ipsec_phase2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-phase2"] = t
		}
	}

	if v, ok := d.GetOkExists("ipsec_phase1_interface"); ok {

		t, err := expandSystemResourceLimitsIpsecPhase1Interface(d, v, "ipsec_phase1_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-phase1-interface"] = t
		}
	}

	if v, ok := d.GetOkExists("ipsec_phase2_interface"); ok {

		t, err := expandSystemResourceLimitsIpsecPhase2Interface(d, v, "ipsec_phase2_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-phase2-interface"] = t
		}
	}

	if v, ok := d.GetOkExists("dialup_tunnel"); ok {

		t, err := expandSystemResourceLimitsDialupTunnel(d, v, "dialup_tunnel", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dialup-tunnel"] = t
		}
	}

	if v, ok := d.GetOkExists("firewall_policy"); ok {

		t, err := expandSystemResourceLimitsFirewallPolicy(d, v, "firewall_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-policy"] = t
		}
	}

	if v, ok := d.GetOkExists("firewall_address"); ok {

		t, err := expandSystemResourceLimitsFirewallAddress(d, v, "firewall_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-address"] = t
		}
	}

	if v, ok := d.GetOkExists("firewall_addrgrp"); ok {

		t, err := expandSystemResourceLimitsFirewallAddrgrp(d, v, "firewall_addrgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-addrgrp"] = t
		}
	}

	if v, ok := d.GetOkExists("custom_service"); ok {

		t, err := expandSystemResourceLimitsCustomService(d, v, "custom_service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-service"] = t
		}
	}

	if v, ok := d.GetOkExists("service_group"); ok {

		t, err := expandSystemResourceLimitsServiceGroup(d, v, "service_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-group"] = t
		}
	}

	if v, ok := d.GetOkExists("onetime_schedule"); ok {

		t, err := expandSystemResourceLimitsOnetimeSchedule(d, v, "onetime_schedule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["onetime-schedule"] = t
		}
	}

	if v, ok := d.GetOkExists("recurring_schedule"); ok {

		t, err := expandSystemResourceLimitsRecurringSchedule(d, v, "recurring_schedule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recurring-schedule"] = t
		}
	}

	if v, ok := d.GetOkExists("user"); ok {

		t, err := expandSystemResourceLimitsUser(d, v, "user", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOkExists("user_group"); ok {

		t, err := expandSystemResourceLimitsUserGroup(d, v, "user_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-group"] = t
		}
	}

	if v, ok := d.GetOkExists("sslvpn"); ok {

		t, err := expandSystemResourceLimitsSslvpn(d, v, "sslvpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn"] = t
		}
	}

	if v, ok := d.GetOkExists("proxy"); ok {

		t, err := expandSystemResourceLimitsProxy(d, v, "proxy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy"] = t
		}
	}

	if v, ok := d.GetOkExists("log_disk_quota"); ok {

		t, err := expandSystemResourceLimitsLogDiskQuota(d, v, "log_disk_quota", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-disk-quota"] = t
		}
	}

	return &obj, nil
}
