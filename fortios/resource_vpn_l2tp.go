// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure L2TP.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"usrgrp": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"enforce_ipsec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lcp_echo_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32767),
				Optional:     true,
				Computed:     true,
			},
			"lcp_max_echo_fails": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32767),
				Optional:     true,
				Computed:     true,
			},
			"hello_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),
				Optional:     true,
				Computed:     true,
			},
			"compress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVpnL2TpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnL2Tp(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnL2Tp resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnL2Tp(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnL2Tp(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating VpnL2Tp resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnL2Tp(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing VpnL2Tp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnL2TpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnL2Tp(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VpnL2Tp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnL2Tp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnL2Tp resource from API: %v", err)
	}
	return nil
}

func flattenVpnL2TpEip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnL2TpSip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnL2TpStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnL2TpUsrgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnL2TpEnforceIpsec(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnL2TpLcpEchoInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnL2TpLcpMaxEchoFails(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnL2TpHelloInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnL2TpCompress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnL2Tp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("eip", flattenVpnL2TpEip(o["eip"], d, "eip", sv)); err != nil {
		if !fortiAPIPatch(o["eip"]) {
			return fmt.Errorf("Error reading eip: %v", err)
		}
	}

	if err = d.Set("sip", flattenVpnL2TpSip(o["sip"], d, "sip", sv)); err != nil {
		if !fortiAPIPatch(o["sip"]) {
			return fmt.Errorf("Error reading sip: %v", err)
		}
	}

	if err = d.Set("status", flattenVpnL2TpStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("usrgrp", flattenVpnL2TpUsrgrp(o["usrgrp"], d, "usrgrp", sv)); err != nil {
		if !fortiAPIPatch(o["usrgrp"]) {
			return fmt.Errorf("Error reading usrgrp: %v", err)
		}
	}

	if err = d.Set("enforce_ipsec", flattenVpnL2TpEnforceIpsec(o["enforce-ipsec"], d, "enforce_ipsec", sv)); err != nil {
		if !fortiAPIPatch(o["enforce-ipsec"]) {
			return fmt.Errorf("Error reading enforce_ipsec: %v", err)
		}
	}

	if err = d.Set("lcp_echo_interval", flattenVpnL2TpLcpEchoInterval(o["lcp-echo-interval"], d, "lcp_echo_interval", sv)); err != nil {
		if !fortiAPIPatch(o["lcp-echo-interval"]) {
			return fmt.Errorf("Error reading lcp_echo_interval: %v", err)
		}
	}

	if err = d.Set("lcp_max_echo_fails", flattenVpnL2TpLcpMaxEchoFails(o["lcp-max-echo-fails"], d, "lcp_max_echo_fails", sv)); err != nil {
		if !fortiAPIPatch(o["lcp-max-echo-fails"]) {
			return fmt.Errorf("Error reading lcp_max_echo_fails: %v", err)
		}
	}

	if err = d.Set("hello_interval", flattenVpnL2TpHelloInterval(o["hello-interval"], d, "hello_interval", sv)); err != nil {
		if !fortiAPIPatch(o["hello-interval"]) {
			return fmt.Errorf("Error reading hello_interval: %v", err)
		}
	}

	if err = d.Set("compress", flattenVpnL2TpCompress(o["compress"], d, "compress", sv)); err != nil {
		if !fortiAPIPatch(o["compress"]) {
			return fmt.Errorf("Error reading compress: %v", err)
		}
	}

	return nil
}

func flattenVpnL2TpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnL2TpEip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpSip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpUsrgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpEnforceIpsec(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpLcpEchoInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpLcpMaxEchoFails(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpHelloInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnL2TpCompress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnL2Tp(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("eip"); ok {
		if setArgNil {
			obj["eip"] = nil
		} else {
			t, err := expandVpnL2TpEip(d, v, "eip", sv)
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
			t, err := expandVpnL2TpSip(d, v, "sip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sip"] = t
			}
		}
	}

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandVpnL2TpStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("usrgrp"); ok {
		if setArgNil {
			obj["usrgrp"] = nil
		} else {
			t, err := expandVpnL2TpUsrgrp(d, v, "usrgrp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["usrgrp"] = t
			}
		}
	}

	if v, ok := d.GetOk("enforce_ipsec"); ok {
		if setArgNil {
			obj["enforce-ipsec"] = nil
		} else {
			t, err := expandVpnL2TpEnforceIpsec(d, v, "enforce_ipsec", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["enforce-ipsec"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("lcp_echo_interval"); ok {
		if setArgNil {
			obj["lcp-echo-interval"] = nil
		} else {
			t, err := expandVpnL2TpLcpEchoInterval(d, v, "lcp_echo_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["lcp-echo-interval"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("lcp_max_echo_fails"); ok {
		if setArgNil {
			obj["lcp-max-echo-fails"] = nil
		} else {
			t, err := expandVpnL2TpLcpMaxEchoFails(d, v, "lcp_max_echo_fails", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["lcp-max-echo-fails"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("hello_interval"); ok {
		if setArgNil {
			obj["hello-interval"] = nil
		} else {
			t, err := expandVpnL2TpHelloInterval(d, v, "hello_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hello-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("compress"); ok {
		if setArgNil {
			obj["compress"] = nil
		} else {
			t, err := expandVpnL2TpCompress(d, v, "compress", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["compress"] = t
			}
		}
	}

	return &obj, nil
}
