// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure PPTP.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnPptp() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnPptpUpdate,
		Read:   resourceVpnPptpRead,
		Update: resourceVpnPptpUpdate,
		Delete: resourceVpnPptpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usrgrp": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceVpnPptpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectVpnPptp(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnPptp resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnPptp(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating VpnPptp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnPptp")
	}

	return resourceVpnPptpRead(d, m)
}

func resourceVpnPptpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnPptp(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating VpnPptp resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnPptp(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing VpnPptp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnPptpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadVpnPptp(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VpnPptp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnPptp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnPptp resource from API: %v", err)
	}
	return nil
}

func flattenVpnPptpStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnPptpIpMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnPptpEip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnPptpSip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnPptpLocalIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnPptpUsrgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnPptp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenVpnPptpStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("ip_mode", flattenVpnPptpIpMode(o["ip-mode"], d, "ip_mode", sv)); err != nil {
		if !fortiAPIPatch(o["ip-mode"]) {
			return fmt.Errorf("Error reading ip_mode: %v", err)
		}
	}

	if err = d.Set("eip", flattenVpnPptpEip(o["eip"], d, "eip", sv)); err != nil {
		if !fortiAPIPatch(o["eip"]) {
			return fmt.Errorf("Error reading eip: %v", err)
		}
	}

	if err = d.Set("sip", flattenVpnPptpSip(o["sip"], d, "sip", sv)); err != nil {
		if !fortiAPIPatch(o["sip"]) {
			return fmt.Errorf("Error reading sip: %v", err)
		}
	}

	if err = d.Set("local_ip", flattenVpnPptpLocalIp(o["local-ip"], d, "local_ip", sv)); err != nil {
		if !fortiAPIPatch(o["local-ip"]) {
			return fmt.Errorf("Error reading local_ip: %v", err)
		}
	}

	if err = d.Set("usrgrp", flattenVpnPptpUsrgrp(o["usrgrp"], d, "usrgrp", sv)); err != nil {
		if !fortiAPIPatch(o["usrgrp"]) {
			return fmt.Errorf("Error reading usrgrp: %v", err)
		}
	}

	return nil
}

func flattenVpnPptpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnPptpStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnPptpIpMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnPptpEip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnPptpSip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnPptpLocalIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnPptpUsrgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnPptp(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandVpnPptpStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("ip_mode"); ok {
		if setArgNil {
			obj["ip-mode"] = nil
		} else {
			t, err := expandVpnPptpIpMode(d, v, "ip_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("eip"); ok {
		if setArgNil {
			obj["eip"] = nil
		} else {
			t, err := expandVpnPptpEip(d, v, "eip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["eip"] = t
			}
		}
	}

	if v, ok := d.GetOk("sip"); ok {
		if setArgNil {
			obj["sip"] = nil
		} else {
			t, err := expandVpnPptpSip(d, v, "sip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sip"] = t
			}
		}
	}

	if v, ok := d.GetOk("local_ip"); ok {
		if setArgNil {
			obj["local-ip"] = nil
		} else {
			t, err := expandVpnPptpLocalIp(d, v, "local_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["local-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("usrgrp"); ok {
		if setArgNil {
			obj["usrgrp"] = nil
		} else {
			t, err := expandVpnPptpUsrgrp(d, v, "usrgrp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["usrgrp"] = t
			}
		}
	}

	return &obj, nil
}
