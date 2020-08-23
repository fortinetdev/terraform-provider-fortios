// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure resource limits.

package fortios

import (
	"fmt"
	"log"
	"strconv"

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
			"session": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"ipsec_phase1": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"ipsec_phase2": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"ipsec_phase1_interface": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"ipsec_phase2_interface": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"dialup_tunnel": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"firewall_policy": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"firewall_address": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"firewall_addrgrp": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"custom_service": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"service_group": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"onetime_schedule": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"recurring_schedule": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"user": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"user_group": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"sslvpn": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"proxy": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"log_disk_quota": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemResourceLimitsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemResourceLimits(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemResourceLimits resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemResourceLimits(obj, mkey)
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

	err := c.DeleteSystemResourceLimits(mkey)
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

	o, err := c.ReadSystemResourceLimits(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemResourceLimits resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemResourceLimits(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemResourceLimits resource from API: %v", err)
	}
	return nil
}

func flattenSystemResourceLimitsSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemResourceLimitsIpsecPhase1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemResourceLimitsIpsecPhase2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemResourceLimitsIpsecPhase1Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemResourceLimitsIpsecPhase2Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemResourceLimitsDialupTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemResourceLimitsFirewallPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemResourceLimitsFirewallAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemResourceLimitsFirewallAddrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemResourceLimitsCustomService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemResourceLimitsServiceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemResourceLimitsOnetimeSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemResourceLimitsRecurringSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemResourceLimitsUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemResourceLimitsUserGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemResourceLimitsSslvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemResourceLimitsProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemResourceLimitsLogDiskQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemResourceLimits(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("session", flattenSystemResourceLimitsSession(o["session"], d, "session")); err != nil {
		if !fortiAPIPatch(o["session"]) {
			return fmt.Errorf("Error reading session: %v", err)
		}
	}

	if err = d.Set("ipsec_phase1", flattenSystemResourceLimitsIpsecPhase1(o["ipsec-phase1"], d, "ipsec_phase1")); err != nil {
		if !fortiAPIPatch(o["ipsec-phase1"]) {
			return fmt.Errorf("Error reading ipsec_phase1: %v", err)
		}
	}

	if err = d.Set("ipsec_phase2", flattenSystemResourceLimitsIpsecPhase2(o["ipsec-phase2"], d, "ipsec_phase2")); err != nil {
		if !fortiAPIPatch(o["ipsec-phase2"]) {
			return fmt.Errorf("Error reading ipsec_phase2: %v", err)
		}
	}

	if err = d.Set("ipsec_phase1_interface", flattenSystemResourceLimitsIpsecPhase1Interface(o["ipsec-phase1-interface"], d, "ipsec_phase1_interface")); err != nil {
		if !fortiAPIPatch(o["ipsec-phase1-interface"]) {
			return fmt.Errorf("Error reading ipsec_phase1_interface: %v", err)
		}
	}

	if err = d.Set("ipsec_phase2_interface", flattenSystemResourceLimitsIpsecPhase2Interface(o["ipsec-phase2-interface"], d, "ipsec_phase2_interface")); err != nil {
		if !fortiAPIPatch(o["ipsec-phase2-interface"]) {
			return fmt.Errorf("Error reading ipsec_phase2_interface: %v", err)
		}
	}

	if err = d.Set("dialup_tunnel", flattenSystemResourceLimitsDialupTunnel(o["dialup-tunnel"], d, "dialup_tunnel")); err != nil {
		if !fortiAPIPatch(o["dialup-tunnel"]) {
			return fmt.Errorf("Error reading dialup_tunnel: %v", err)
		}
	}

	if err = d.Set("firewall_policy", flattenSystemResourceLimitsFirewallPolicy(o["firewall-policy"], d, "firewall_policy")); err != nil {
		if !fortiAPIPatch(o["firewall-policy"]) {
			return fmt.Errorf("Error reading firewall_policy: %v", err)
		}
	}

	if err = d.Set("firewall_address", flattenSystemResourceLimitsFirewallAddress(o["firewall-address"], d, "firewall_address")); err != nil {
		if !fortiAPIPatch(o["firewall-address"]) {
			return fmt.Errorf("Error reading firewall_address: %v", err)
		}
	}

	if err = d.Set("firewall_addrgrp", flattenSystemResourceLimitsFirewallAddrgrp(o["firewall-addrgrp"], d, "firewall_addrgrp")); err != nil {
		if !fortiAPIPatch(o["firewall-addrgrp"]) {
			return fmt.Errorf("Error reading firewall_addrgrp: %v", err)
		}
	}

	if err = d.Set("custom_service", flattenSystemResourceLimitsCustomService(o["custom-service"], d, "custom_service")); err != nil {
		if !fortiAPIPatch(o["custom-service"]) {
			return fmt.Errorf("Error reading custom_service: %v", err)
		}
	}

	if err = d.Set("service_group", flattenSystemResourceLimitsServiceGroup(o["service-group"], d, "service_group")); err != nil {
		if !fortiAPIPatch(o["service-group"]) {
			return fmt.Errorf("Error reading service_group: %v", err)
		}
	}

	if err = d.Set("onetime_schedule", flattenSystemResourceLimitsOnetimeSchedule(o["onetime-schedule"], d, "onetime_schedule")); err != nil {
		if !fortiAPIPatch(o["onetime-schedule"]) {
			return fmt.Errorf("Error reading onetime_schedule: %v", err)
		}
	}

	if err = d.Set("recurring_schedule", flattenSystemResourceLimitsRecurringSchedule(o["recurring-schedule"], d, "recurring_schedule")); err != nil {
		if !fortiAPIPatch(o["recurring-schedule"]) {
			return fmt.Errorf("Error reading recurring_schedule: %v", err)
		}
	}

	if err = d.Set("user", flattenSystemResourceLimitsUser(o["user"], d, "user")); err != nil {
		if !fortiAPIPatch(o["user"]) {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("user_group", flattenSystemResourceLimitsUserGroup(o["user-group"], d, "user_group")); err != nil {
		if !fortiAPIPatch(o["user-group"]) {
			return fmt.Errorf("Error reading user_group: %v", err)
		}
	}

	if err = d.Set("sslvpn", flattenSystemResourceLimitsSslvpn(o["sslvpn"], d, "sslvpn")); err != nil {
		if !fortiAPIPatch(o["sslvpn"]) {
			return fmt.Errorf("Error reading sslvpn: %v", err)
		}
	}

	if err = d.Set("proxy", flattenSystemResourceLimitsProxy(o["proxy"], d, "proxy")); err != nil {
		if !fortiAPIPatch(o["proxy"]) {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	if err = d.Set("log_disk_quota", flattenSystemResourceLimitsLogDiskQuota(o["log-disk-quota"], d, "log_disk_quota")); err != nil {
		if !fortiAPIPatch(o["log-disk-quota"]) {
			return fmt.Errorf("Error reading log_disk_quota: %v", err)
		}
	}

	return nil
}

func flattenSystemResourceLimitsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemResourceLimitsSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsIpsecPhase1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsIpsecPhase2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsIpsecPhase1Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsIpsecPhase2Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsDialupTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsFirewallPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsFirewallAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsFirewallAddrgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsCustomService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsServiceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsOnetimeSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsRecurringSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsUserGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsSslvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemResourceLimitsLogDiskQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemResourceLimits(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("session"); ok {
		t, err := expandSystemResourceLimitsSession(d, v, "session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_phase1"); ok {
		t, err := expandSystemResourceLimitsIpsecPhase1(d, v, "ipsec_phase1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-phase1"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_phase2"); ok {
		t, err := expandSystemResourceLimitsIpsecPhase2(d, v, "ipsec_phase2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-phase2"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_phase1_interface"); ok {
		t, err := expandSystemResourceLimitsIpsecPhase1Interface(d, v, "ipsec_phase1_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-phase1-interface"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_phase2_interface"); ok {
		t, err := expandSystemResourceLimitsIpsecPhase2Interface(d, v, "ipsec_phase2_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-phase2-interface"] = t
		}
	}

	if v, ok := d.GetOk("dialup_tunnel"); ok {
		t, err := expandSystemResourceLimitsDialupTunnel(d, v, "dialup_tunnel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dialup-tunnel"] = t
		}
	}

	if v, ok := d.GetOk("firewall_policy"); ok {
		t, err := expandSystemResourceLimitsFirewallPolicy(d, v, "firewall_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-policy"] = t
		}
	}

	if v, ok := d.GetOk("firewall_address"); ok {
		t, err := expandSystemResourceLimitsFirewallAddress(d, v, "firewall_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-address"] = t
		}
	}

	if v, ok := d.GetOk("firewall_addrgrp"); ok {
		t, err := expandSystemResourceLimitsFirewallAddrgrp(d, v, "firewall_addrgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-addrgrp"] = t
		}
	}

	if v, ok := d.GetOk("custom_service"); ok {
		t, err := expandSystemResourceLimitsCustomService(d, v, "custom_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-service"] = t
		}
	}

	if v, ok := d.GetOk("service_group"); ok {
		t, err := expandSystemResourceLimitsServiceGroup(d, v, "service_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-group"] = t
		}
	}

	if v, ok := d.GetOk("onetime_schedule"); ok {
		t, err := expandSystemResourceLimitsOnetimeSchedule(d, v, "onetime_schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["onetime-schedule"] = t
		}
	}

	if v, ok := d.GetOk("recurring_schedule"); ok {
		t, err := expandSystemResourceLimitsRecurringSchedule(d, v, "recurring_schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recurring-schedule"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok {
		t, err := expandSystemResourceLimitsUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOk("user_group"); ok {
		t, err := expandSystemResourceLimitsUserGroup(d, v, "user_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-group"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn"); ok {
		t, err := expandSystemResourceLimitsSslvpn(d, v, "sslvpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn"] = t
		}
	}

	if v, ok := d.GetOk("proxy"); ok {
		t, err := expandSystemResourceLimitsProxy(d, v, "proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy"] = t
		}
	}

	if v, ok := d.GetOk("log_disk_quota"); ok {
		t, err := expandSystemResourceLimitsLogDiskQuota(d, v, "log_disk_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-disk-quota"] = t
		}
	}

	return &obj, nil
}
