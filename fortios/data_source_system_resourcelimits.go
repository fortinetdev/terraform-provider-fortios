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

func dataSourceSystemResourceLimits() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemResourceLimitsRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"session": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ipsec_phase1": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ipsec_phase2": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ipsec_phase1_interface": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ipsec_phase2_interface": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dialup_tunnel": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"firewall_policy": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"firewall_address": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"firewall_addrgrp": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"custom_service": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"service_group": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"onetime_schedule": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"recurring_schedule": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"user_group": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"sslvpn": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"proxy": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"log_disk_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemResourceLimitsRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "SystemResourceLimits"

	o, err := c.ReadSystemResourceLimits(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemResourceLimits: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemResourceLimits(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemResourceLimits from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemResourceLimitsSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemResourceLimitsIpsecPhase1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemResourceLimitsIpsecPhase2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemResourceLimitsIpsecPhase1Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemResourceLimitsIpsecPhase2Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemResourceLimitsDialupTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemResourceLimitsFirewallPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemResourceLimitsFirewallAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemResourceLimitsFirewallAddrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemResourceLimitsCustomService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemResourceLimitsServiceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemResourceLimitsOnetimeSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemResourceLimitsRecurringSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemResourceLimitsUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemResourceLimitsUserGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemResourceLimitsSslvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemResourceLimitsProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemResourceLimitsLogDiskQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemResourceLimits(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("session", dataSourceFlattenSystemResourceLimitsSession(o["session"], d, "session")); err != nil {
		if !fortiAPIPatch(o["session"]) {
			return fmt.Errorf("Error reading session: %v", err)
		}
	}

	if err = d.Set("ipsec_phase1", dataSourceFlattenSystemResourceLimitsIpsecPhase1(o["ipsec-phase1"], d, "ipsec_phase1")); err != nil {
		if !fortiAPIPatch(o["ipsec-phase1"]) {
			return fmt.Errorf("Error reading ipsec_phase1: %v", err)
		}
	}

	if err = d.Set("ipsec_phase2", dataSourceFlattenSystemResourceLimitsIpsecPhase2(o["ipsec-phase2"], d, "ipsec_phase2")); err != nil {
		if !fortiAPIPatch(o["ipsec-phase2"]) {
			return fmt.Errorf("Error reading ipsec_phase2: %v", err)
		}
	}

	if err = d.Set("ipsec_phase1_interface", dataSourceFlattenSystemResourceLimitsIpsecPhase1Interface(o["ipsec-phase1-interface"], d, "ipsec_phase1_interface")); err != nil {
		if !fortiAPIPatch(o["ipsec-phase1-interface"]) {
			return fmt.Errorf("Error reading ipsec_phase1_interface: %v", err)
		}
	}

	if err = d.Set("ipsec_phase2_interface", dataSourceFlattenSystemResourceLimitsIpsecPhase2Interface(o["ipsec-phase2-interface"], d, "ipsec_phase2_interface")); err != nil {
		if !fortiAPIPatch(o["ipsec-phase2-interface"]) {
			return fmt.Errorf("Error reading ipsec_phase2_interface: %v", err)
		}
	}

	if err = d.Set("dialup_tunnel", dataSourceFlattenSystemResourceLimitsDialupTunnel(o["dialup-tunnel"], d, "dialup_tunnel")); err != nil {
		if !fortiAPIPatch(o["dialup-tunnel"]) {
			return fmt.Errorf("Error reading dialup_tunnel: %v", err)
		}
	}

	if err = d.Set("firewall_policy", dataSourceFlattenSystemResourceLimitsFirewallPolicy(o["firewall-policy"], d, "firewall_policy")); err != nil {
		if !fortiAPIPatch(o["firewall-policy"]) {
			return fmt.Errorf("Error reading firewall_policy: %v", err)
		}
	}

	if err = d.Set("firewall_address", dataSourceFlattenSystemResourceLimitsFirewallAddress(o["firewall-address"], d, "firewall_address")); err != nil {
		if !fortiAPIPatch(o["firewall-address"]) {
			return fmt.Errorf("Error reading firewall_address: %v", err)
		}
	}

	if err = d.Set("firewall_addrgrp", dataSourceFlattenSystemResourceLimitsFirewallAddrgrp(o["firewall-addrgrp"], d, "firewall_addrgrp")); err != nil {
		if !fortiAPIPatch(o["firewall-addrgrp"]) {
			return fmt.Errorf("Error reading firewall_addrgrp: %v", err)
		}
	}

	if err = d.Set("custom_service", dataSourceFlattenSystemResourceLimitsCustomService(o["custom-service"], d, "custom_service")); err != nil {
		if !fortiAPIPatch(o["custom-service"]) {
			return fmt.Errorf("Error reading custom_service: %v", err)
		}
	}

	if err = d.Set("service_group", dataSourceFlattenSystemResourceLimitsServiceGroup(o["service-group"], d, "service_group")); err != nil {
		if !fortiAPIPatch(o["service-group"]) {
			return fmt.Errorf("Error reading service_group: %v", err)
		}
	}

	if err = d.Set("onetime_schedule", dataSourceFlattenSystemResourceLimitsOnetimeSchedule(o["onetime-schedule"], d, "onetime_schedule")); err != nil {
		if !fortiAPIPatch(o["onetime-schedule"]) {
			return fmt.Errorf("Error reading onetime_schedule: %v", err)
		}
	}

	if err = d.Set("recurring_schedule", dataSourceFlattenSystemResourceLimitsRecurringSchedule(o["recurring-schedule"], d, "recurring_schedule")); err != nil {
		if !fortiAPIPatch(o["recurring-schedule"]) {
			return fmt.Errorf("Error reading recurring_schedule: %v", err)
		}
	}

	if err = d.Set("user", dataSourceFlattenSystemResourceLimitsUser(o["user"], d, "user")); err != nil {
		if !fortiAPIPatch(o["user"]) {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("user_group", dataSourceFlattenSystemResourceLimitsUserGroup(o["user-group"], d, "user_group")); err != nil {
		if !fortiAPIPatch(o["user-group"]) {
			return fmt.Errorf("Error reading user_group: %v", err)
		}
	}

	if err = d.Set("sslvpn", dataSourceFlattenSystemResourceLimitsSslvpn(o["sslvpn"], d, "sslvpn")); err != nil {
		if !fortiAPIPatch(o["sslvpn"]) {
			return fmt.Errorf("Error reading sslvpn: %v", err)
		}
	}

	if err = d.Set("proxy", dataSourceFlattenSystemResourceLimitsProxy(o["proxy"], d, "proxy")); err != nil {
		if !fortiAPIPatch(o["proxy"]) {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	if err = d.Set("log_disk_quota", dataSourceFlattenSystemResourceLimitsLogDiskQuota(o["log-disk-quota"], d, "log_disk_quota")); err != nil {
		if !fortiAPIPatch(o["log-disk-quota"]) {
			return fmt.Errorf("Error reading log_disk_quota: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemResourceLimitsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
