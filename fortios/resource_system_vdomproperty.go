// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure VDOM property.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemVdomProperty() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVdomPropertyCreate,
		Read:   resourceSystemVdomPropertyRead,
		Update: resourceSystemVdomPropertyUpdate,
		Delete: resourceSystemVdomPropertyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"snmp_index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_phase1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_phase2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_phase1_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_phase2_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dialup_tunnel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"firewall_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"firewall_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"firewall_addrgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"custom_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"onetime_schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"recurring_schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslvpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_disk_quota": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemVdomPropertyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemVdomProperty(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdomProperty resource while getting object: %v", err)
	}

	o, err := c.CreateSystemVdomProperty(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemVdomProperty resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVdomProperty")
	}

	return resourceSystemVdomPropertyRead(d, m)
}

func resourceSystemVdomPropertyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemVdomProperty(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomProperty resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemVdomProperty(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomProperty resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVdomProperty")
	}

	return resourceSystemVdomPropertyRead(d, m)
}

func resourceSystemVdomPropertyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemVdomProperty(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVdomProperty resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomPropertyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemVdomProperty(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomProperty resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdomProperty(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomProperty resource from API: %v", err)
	}
	return nil
}

func flattenSystemVdomPropertyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertyDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertySnmpIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertySession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertyIpsecPhase1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertyIpsecPhase2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertyIpsecPhase1Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertyIpsecPhase2Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertyDialupTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertyFirewallPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertyFirewallAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertyFirewallAddrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertyCustomService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertyServiceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertyOnetimeSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertyRecurringSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertyUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertyUserGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertySslvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertyProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomPropertyLogDiskQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemVdomProperty(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemVdomPropertyName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemVdomPropertyDescription(o["description"], d, "description")); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("snmp_index", flattenSystemVdomPropertySnmpIndex(o["snmp-index"], d, "snmp_index")); err != nil {
		if !fortiAPIPatch(o["snmp-index"]) {
			return fmt.Errorf("Error reading snmp_index: %v", err)
		}
	}

	if err = d.Set("session", flattenSystemVdomPropertySession(o["session"], d, "session")); err != nil {
		if !fortiAPIPatch(o["session"]) {
			return fmt.Errorf("Error reading session: %v", err)
		}
	}

	if err = d.Set("ipsec_phase1", flattenSystemVdomPropertyIpsecPhase1(o["ipsec-phase1"], d, "ipsec_phase1")); err != nil {
		if !fortiAPIPatch(o["ipsec-phase1"]) {
			return fmt.Errorf("Error reading ipsec_phase1: %v", err)
		}
	}

	if err = d.Set("ipsec_phase2", flattenSystemVdomPropertyIpsecPhase2(o["ipsec-phase2"], d, "ipsec_phase2")); err != nil {
		if !fortiAPIPatch(o["ipsec-phase2"]) {
			return fmt.Errorf("Error reading ipsec_phase2: %v", err)
		}
	}

	if err = d.Set("ipsec_phase1_interface", flattenSystemVdomPropertyIpsecPhase1Interface(o["ipsec-phase1-interface"], d, "ipsec_phase1_interface")); err != nil {
		if !fortiAPIPatch(o["ipsec-phase1-interface"]) {
			return fmt.Errorf("Error reading ipsec_phase1_interface: %v", err)
		}
	}

	if err = d.Set("ipsec_phase2_interface", flattenSystemVdomPropertyIpsecPhase2Interface(o["ipsec-phase2-interface"], d, "ipsec_phase2_interface")); err != nil {
		if !fortiAPIPatch(o["ipsec-phase2-interface"]) {
			return fmt.Errorf("Error reading ipsec_phase2_interface: %v", err)
		}
	}

	if err = d.Set("dialup_tunnel", flattenSystemVdomPropertyDialupTunnel(o["dialup-tunnel"], d, "dialup_tunnel")); err != nil {
		if !fortiAPIPatch(o["dialup-tunnel"]) {
			return fmt.Errorf("Error reading dialup_tunnel: %v", err)
		}
	}

	if err = d.Set("firewall_policy", flattenSystemVdomPropertyFirewallPolicy(o["firewall-policy"], d, "firewall_policy")); err != nil {
		if !fortiAPIPatch(o["firewall-policy"]) {
			return fmt.Errorf("Error reading firewall_policy: %v", err)
		}
	}

	if err = d.Set("firewall_address", flattenSystemVdomPropertyFirewallAddress(o["firewall-address"], d, "firewall_address")); err != nil {
		if !fortiAPIPatch(o["firewall-address"]) {
			return fmt.Errorf("Error reading firewall_address: %v", err)
		}
	}

	if err = d.Set("firewall_addrgrp", flattenSystemVdomPropertyFirewallAddrgrp(o["firewall-addrgrp"], d, "firewall_addrgrp")); err != nil {
		if !fortiAPIPatch(o["firewall-addrgrp"]) {
			return fmt.Errorf("Error reading firewall_addrgrp: %v", err)
		}
	}

	if err = d.Set("custom_service", flattenSystemVdomPropertyCustomService(o["custom-service"], d, "custom_service")); err != nil {
		if !fortiAPIPatch(o["custom-service"]) {
			return fmt.Errorf("Error reading custom_service: %v", err)
		}
	}

	if err = d.Set("service_group", flattenSystemVdomPropertyServiceGroup(o["service-group"], d, "service_group")); err != nil {
		if !fortiAPIPatch(o["service-group"]) {
			return fmt.Errorf("Error reading service_group: %v", err)
		}
	}

	if err = d.Set("onetime_schedule", flattenSystemVdomPropertyOnetimeSchedule(o["onetime-schedule"], d, "onetime_schedule")); err != nil {
		if !fortiAPIPatch(o["onetime-schedule"]) {
			return fmt.Errorf("Error reading onetime_schedule: %v", err)
		}
	}

	if err = d.Set("recurring_schedule", flattenSystemVdomPropertyRecurringSchedule(o["recurring-schedule"], d, "recurring_schedule")); err != nil {
		if !fortiAPIPatch(o["recurring-schedule"]) {
			return fmt.Errorf("Error reading recurring_schedule: %v", err)
		}
	}

	if err = d.Set("user", flattenSystemVdomPropertyUser(o["user"], d, "user")); err != nil {
		if !fortiAPIPatch(o["user"]) {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("user_group", flattenSystemVdomPropertyUserGroup(o["user-group"], d, "user_group")); err != nil {
		if !fortiAPIPatch(o["user-group"]) {
			return fmt.Errorf("Error reading user_group: %v", err)
		}
	}

	if err = d.Set("sslvpn", flattenSystemVdomPropertySslvpn(o["sslvpn"], d, "sslvpn")); err != nil {
		if !fortiAPIPatch(o["sslvpn"]) {
			return fmt.Errorf("Error reading sslvpn: %v", err)
		}
	}

	if err = d.Set("proxy", flattenSystemVdomPropertyProxy(o["proxy"], d, "proxy")); err != nil {
		if !fortiAPIPatch(o["proxy"]) {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	if err = d.Set("log_disk_quota", flattenSystemVdomPropertyLogDiskQuota(o["log-disk-quota"], d, "log_disk_quota")); err != nil {
		if !fortiAPIPatch(o["log-disk-quota"]) {
			return fmt.Errorf("Error reading log_disk_quota: %v", err)
		}
	}

	return nil
}

func flattenSystemVdomPropertyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemVdomPropertyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertyDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertySnmpIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertySession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertyIpsecPhase1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertyIpsecPhase2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertyIpsecPhase1Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertyIpsecPhase2Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertyDialupTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertyFirewallPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertyFirewallAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertyFirewallAddrgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertyCustomService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertyServiceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertyOnetimeSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertyRecurringSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertyUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertyUserGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertySslvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertyProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomPropertyLogDiskQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVdomProperty(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemVdomPropertyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandSystemVdomPropertyDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("snmp_index"); ok {
		t, err := expandSystemVdomPropertySnmpIndex(d, v, "snmp_index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["snmp-index"] = t
		}
	}

	if v, ok := d.GetOk("session"); ok {
		t, err := expandSystemVdomPropertySession(d, v, "session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_phase1"); ok {
		t, err := expandSystemVdomPropertyIpsecPhase1(d, v, "ipsec_phase1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-phase1"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_phase2"); ok {
		t, err := expandSystemVdomPropertyIpsecPhase2(d, v, "ipsec_phase2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-phase2"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_phase1_interface"); ok {
		t, err := expandSystemVdomPropertyIpsecPhase1Interface(d, v, "ipsec_phase1_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-phase1-interface"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_phase2_interface"); ok {
		t, err := expandSystemVdomPropertyIpsecPhase2Interface(d, v, "ipsec_phase2_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-phase2-interface"] = t
		}
	}

	if v, ok := d.GetOk("dialup_tunnel"); ok {
		t, err := expandSystemVdomPropertyDialupTunnel(d, v, "dialup_tunnel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dialup-tunnel"] = t
		}
	}

	if v, ok := d.GetOk("firewall_policy"); ok {
		t, err := expandSystemVdomPropertyFirewallPolicy(d, v, "firewall_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-policy"] = t
		}
	}

	if v, ok := d.GetOk("firewall_address"); ok {
		t, err := expandSystemVdomPropertyFirewallAddress(d, v, "firewall_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-address"] = t
		}
	}

	if v, ok := d.GetOk("firewall_addrgrp"); ok {
		t, err := expandSystemVdomPropertyFirewallAddrgrp(d, v, "firewall_addrgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-addrgrp"] = t
		}
	}

	if v, ok := d.GetOk("custom_service"); ok {
		t, err := expandSystemVdomPropertyCustomService(d, v, "custom_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-service"] = t
		}
	}

	if v, ok := d.GetOk("service_group"); ok {
		t, err := expandSystemVdomPropertyServiceGroup(d, v, "service_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-group"] = t
		}
	}

	if v, ok := d.GetOk("onetime_schedule"); ok {
		t, err := expandSystemVdomPropertyOnetimeSchedule(d, v, "onetime_schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["onetime-schedule"] = t
		}
	}

	if v, ok := d.GetOk("recurring_schedule"); ok {
		t, err := expandSystemVdomPropertyRecurringSchedule(d, v, "recurring_schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recurring-schedule"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok {
		t, err := expandSystemVdomPropertyUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOk("user_group"); ok {
		t, err := expandSystemVdomPropertyUserGroup(d, v, "user_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-group"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn"); ok {
		t, err := expandSystemVdomPropertySslvpn(d, v, "sslvpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn"] = t
		}
	}

	if v, ok := d.GetOk("proxy"); ok {
		t, err := expandSystemVdomPropertyProxy(d, v, "proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy"] = t
		}
	}

	if v, ok := d.GetOk("log_disk_quota"); ok {
		t, err := expandSystemVdomPropertyLogDiskQuota(d, v, "log_disk_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-disk-quota"] = t
		}
	}

	return &obj, nil
}
