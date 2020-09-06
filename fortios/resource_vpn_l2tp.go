// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure L2TP.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceVpnL2Tp() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnL2TpUpdate,
		Read:   resourceVpnL2TpRead,
		Update: resourceVpnL2TpUpdate,
		Delete: resourceVpnL2TpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"eip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"usrgrp": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"enforce_ipsec": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceVpnL2TpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVpnL2Tp(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnL2Tp resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnL2Tp(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating VpnL2Tp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnL2Tp")
	}

	return resourceVpnL2TpRead(d, m)
}

func resourceVpnL2TpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteVpnL2Tp(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting VpnL2Tp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnL2TpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadVpnL2Tp(mkey)
	if err != nil {
		return fmt.Errorf("Error reading VpnL2Tp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnL2Tp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnL2Tp resource from API: %v", err)
	}
	return nil
}

func flattenVpnL2TpEip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnL2TpSip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnL2TpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnL2TpUsrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnL2TpEnforceIpsec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnL2Tp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("eip", flattenVpnL2TpEip(o["eip"], d, "eip")); err != nil {
		if !fortiAPIPatch(o["eip"]) {
			return fmt.Errorf("Error reading eip: %v", err)
		}
	}

	if err = d.Set("sip", flattenVpnL2TpSip(o["sip"], d, "sip")); err != nil {
		if !fortiAPIPatch(o["sip"]) {
			return fmt.Errorf("Error reading sip: %v", err)
		}
	}

	if err = d.Set("status", flattenVpnL2TpStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("usrgrp", flattenVpnL2TpUsrgrp(o["usrgrp"], d, "usrgrp")); err != nil {
		if !fortiAPIPatch(o["usrgrp"]) {
			return fmt.Errorf("Error reading usrgrp: %v", err)
		}
	}

	if err = d.Set("enforce_ipsec", flattenVpnL2TpEnforceIpsec(o["enforce-ipsec"], d, "enforce_ipsec")); err != nil {
		if !fortiAPIPatch(o["enforce-ipsec"]) {
			return fmt.Errorf("Error reading enforce_ipsec: %v", err)
		}
	}

	return nil
}

func flattenVpnL2TpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnL2TpEip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpSip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpUsrgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpEnforceIpsec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnL2Tp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("eip"); ok {
		t, err := expandVpnL2TpEip(d, v, "eip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eip"] = t
		}
	}

	if v, ok := d.GetOk("sip"); ok {
		t, err := expandVpnL2TpSip(d, v, "sip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandVpnL2TpStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("usrgrp"); ok {
		t, err := expandVpnL2TpUsrgrp(d, v, "usrgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usrgrp"] = t
		}
	}

	if v, ok := d.GetOk("enforce_ipsec"); ok {
		t, err := expandVpnL2TpEnforceIpsec(d, v, "enforce_ipsec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-ipsec"] = t
		}
	}

	return &obj, nil
}
