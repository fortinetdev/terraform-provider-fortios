// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure USB LTE/WIMAX devices.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLteModem() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLteModemUpdate,
		Read:   resourceSystemLteModemRead,
		Update: resourceSystemLteModemUpdate,
		Delete: resourceSystemLteModemDelete,

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
				Optional: true,
				Computed: true,
			},
			"extra_init": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"authtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"passwd": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"apn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"modem_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 20),
				Optional:     true,
				Computed:     true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"holddown_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 60),
				Optional:     true,
				Computed:     true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
		},
	}
}

func resourceSystemLteModemUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemLteModem(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemLteModem resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemLteModem(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemLteModem resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemLteModem")
	}

	return resourceSystemLteModemRead(d, m)
}

func resourceSystemLteModemDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemLteModem(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemLteModem resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLteModem(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemLteModem resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLteModemRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemLteModem(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemLteModem resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLteModem(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemLteModem resource from API: %v", err)
	}
	return nil
}

func flattenSystemLteModemStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLteModemExtraInit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLteModemAuthtype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLteModemUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLteModemApn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLteModemModemPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemLteModemMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLteModemHolddownTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemLteModemInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemLteModem(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemLteModemStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("extra_init", flattenSystemLteModemExtraInit(o["extra-init"], d, "extra_init", sv)); err != nil {
		if !fortiAPIPatch(o["extra-init"]) {
			return fmt.Errorf("Error reading extra_init: %v", err)
		}
	}

	if err = d.Set("authtype", flattenSystemLteModemAuthtype(o["authtype"], d, "authtype", sv)); err != nil {
		if !fortiAPIPatch(o["authtype"]) {
			return fmt.Errorf("Error reading authtype: %v", err)
		}
	}

	if err = d.Set("username", flattenSystemLteModemUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("apn", flattenSystemLteModemApn(o["apn"], d, "apn", sv)); err != nil {
		if !fortiAPIPatch(o["apn"]) {
			return fmt.Errorf("Error reading apn: %v", err)
		}
	}

	if err = d.Set("modem_port", flattenSystemLteModemModemPort(o["modem-port"], d, "modem_port", sv)); err != nil {
		if !fortiAPIPatch(o["modem-port"]) {
			return fmt.Errorf("Error reading modem_port: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemLteModemMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("holddown_timer", flattenSystemLteModemHolddownTimer(o["holddown-timer"], d, "holddown_timer", sv)); err != nil {
		if !fortiAPIPatch(o["holddown-timer"]) {
			return fmt.Errorf("Error reading holddown_timer: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemLteModemInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func flattenSystemLteModemFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemLteModemStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLteModemExtraInit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLteModemAuthtype(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLteModemUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLteModemPasswd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLteModemApn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLteModemModemPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLteModemMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLteModemHolddownTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLteModemInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLteModem(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSystemLteModemStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("extra_init"); ok {
		if setArgNil {
			obj["extra-init"] = nil
		} else {
			t, err := expandSystemLteModemExtraInit(d, v, "extra_init", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["extra-init"] = t
			}
		}
	} else if d.HasChange("extra_init") {
		obj["extra-init"] = nil
	}

	if v, ok := d.GetOk("authtype"); ok {
		if setArgNil {
			obj["authtype"] = nil
		} else {
			t, err := expandSystemLteModemAuthtype(d, v, "authtype", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["authtype"] = t
			}
		}
	}

	if v, ok := d.GetOk("username"); ok {
		if setArgNil {
			obj["username"] = nil
		} else {
			t, err := expandSystemLteModemUsername(d, v, "username", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["username"] = t
			}
		}
	} else if d.HasChange("username") {
		obj["username"] = nil
	}

	if v, ok := d.GetOk("passwd"); ok {
		if setArgNil {
			obj["passwd"] = nil
		} else {
			t, err := expandSystemLteModemPasswd(d, v, "passwd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["passwd"] = t
			}
		}
	} else if d.HasChange("passwd") {
		obj["passwd"] = nil
	}

	if v, ok := d.GetOk("apn"); ok {
		if setArgNil {
			obj["apn"] = nil
		} else {
			t, err := expandSystemLteModemApn(d, v, "apn", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["apn"] = t
			}
		}
	} else if d.HasChange("apn") {
		obj["apn"] = nil
	}

	if v, ok := d.GetOkExists("modem_port"); ok {
		if setArgNil {
			obj["modem-port"] = nil
		} else {
			t, err := expandSystemLteModemModemPort(d, v, "modem_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["modem-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		if setArgNil {
			obj["mode"] = nil
		} else {
			t, err := expandSystemLteModemMode(d, v, "mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("holddown_timer"); ok {
		if setArgNil {
			obj["holddown-timer"] = nil
		} else {
			t, err := expandSystemLteModemHolddownTimer(d, v, "holddown_timer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["holddown-timer"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		if setArgNil {
			obj["interface"] = nil
		} else {
			t, err := expandSystemLteModemInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	return &obj, nil
}
