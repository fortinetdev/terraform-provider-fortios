// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPS NGFW policy-mode VDOM settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemNgfwSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNgfwSettingsUpdate,
		Read:   resourceSystemNgfwSettingsRead,
		Update: resourceSystemNgfwSettingsUpdate,
		Delete: resourceSystemNgfwSettingsDelete,

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
			"match_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1800),
				Optional:     true,
				Computed:     true,
			},
			"tcp_match_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1800),
				Optional:     true,
				Computed:     true,
			},
			"tcp_halfopen_match_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 300),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemNgfwSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemNgfwSettings(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemNgfwSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemNgfwSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemNgfwSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemNgfwSettings")
	}

	return resourceSystemNgfwSettingsRead(d, m)
}

func resourceSystemNgfwSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemNgfwSettings(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemNgfwSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemNgfwSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemNgfwSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNgfwSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemNgfwSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemNgfwSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNgfwSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemNgfwSettings resource from API: %v", err)
	}
	return nil
}

func flattenSystemNgfwSettingsMatchTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNgfwSettingsTcpMatchTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNgfwSettingsTcpHalfopenMatchTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectSystemNgfwSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("match_timeout", flattenSystemNgfwSettingsMatchTimeout(o["match-timeout"], d, "match_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["match-timeout"]) {
			return fmt.Errorf("Error reading match_timeout: %v", err)
		}
	}

	if err = d.Set("tcp_match_timeout", flattenSystemNgfwSettingsTcpMatchTimeout(o["tcp-match-timeout"], d, "tcp_match_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-match-timeout"]) {
			return fmt.Errorf("Error reading tcp_match_timeout: %v", err)
		}
	}

	if err = d.Set("tcp_halfopen_match_timeout", flattenSystemNgfwSettingsTcpHalfopenMatchTimeout(o["tcp-halfopen-match-timeout"], d, "tcp_halfopen_match_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-halfopen-match-timeout"]) {
			return fmt.Errorf("Error reading tcp_halfopen_match_timeout: %v", err)
		}
	}

	return nil
}

func flattenSystemNgfwSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemNgfwSettingsMatchTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNgfwSettingsTcpMatchTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNgfwSettingsTcpHalfopenMatchTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNgfwSettings(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("match_timeout"); ok {
		if setArgNil {
			obj["match-timeout"] = nil
		} else {
			t, err := expandSystemNgfwSettingsMatchTimeout(d, v, "match_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["match-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("tcp_match_timeout"); ok {
		if setArgNil {
			obj["tcp-match-timeout"] = nil
		} else {
			t, err := expandSystemNgfwSettingsTcpMatchTimeout(d, v, "tcp_match_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tcp-match-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("tcp_halfopen_match_timeout"); ok {
		if setArgNil {
			obj["tcp-halfopen-match-timeout"] = nil
		} else {
			t, err := expandSystemNgfwSettingsTcpHalfopenMatchTimeout(d, v, "tcp_halfopen_match_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tcp-halfopen-match-timeout"] = t
			}
		}
	}

	return &obj, nil
}
