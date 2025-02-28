// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure TACACS+ server entries.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserTacacs() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserTacacsCreate,
		Read:   resourceUserTacacsRead,
		Update: resourceUserTacacsUpdate,
		Delete: resourceUserTacacsDelete,

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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"secondary_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"tertiary_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"secondary_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"tertiary_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"status_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 600),
				Optional:     true,
				Computed:     true,
			},
			"authen_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"vrf_select": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 511),
				Optional:     true,
			},
		},
	}
}

func resourceUserTacacsCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserTacacs(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserTacacs resource while getting object: %v", err)
	}

	o, err := c.CreateUserTacacs(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating UserTacacs resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserTacacs")
	}

	return resourceUserTacacsRead(d, m)
}

func resourceUserTacacsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserTacacs(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserTacacs resource while getting object: %v", err)
	}

	o, err := c.UpdateUserTacacs(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserTacacs resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserTacacs")
	}

	return resourceUserTacacsRead(d, m)
}

func resourceUserTacacsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUserTacacs(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UserTacacs resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserTacacsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserTacacs(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserTacacs resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserTacacs(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserTacacs resource from API: %v", err)
	}
	return nil
}

func flattenUserTacacsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserTacacsServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserTacacsSecondaryServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserTacacsTertiaryServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserTacacsPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserTacacsStatusTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserTacacsAuthenType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserTacacsAuthorization(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserTacacsSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserTacacsInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserTacacsInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserTacacsVrfSelect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectUserTacacs(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenUserTacacsName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server", flattenUserTacacsServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("secondary_server", flattenUserTacacsSecondaryServer(o["secondary-server"], d, "secondary_server", sv)); err != nil {
		if !fortiAPIPatch(o["secondary-server"]) {
			return fmt.Errorf("Error reading secondary_server: %v", err)
		}
	}

	if err = d.Set("tertiary_server", flattenUserTacacsTertiaryServer(o["tertiary-server"], d, "tertiary_server", sv)); err != nil {
		if !fortiAPIPatch(o["tertiary-server"]) {
			return fmt.Errorf("Error reading tertiary_server: %v", err)
		}
	}

	if err = d.Set("port", flattenUserTacacsPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("status_ttl", flattenUserTacacsStatusTtl(o["status-ttl"], d, "status_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["status-ttl"]) {
			return fmt.Errorf("Error reading status_ttl: %v", err)
		}
	}

	if err = d.Set("authen_type", flattenUserTacacsAuthenType(o["authen-type"], d, "authen_type", sv)); err != nil {
		if !fortiAPIPatch(o["authen-type"]) {
			return fmt.Errorf("Error reading authen_type: %v", err)
		}
	}

	if err = d.Set("authorization", flattenUserTacacsAuthorization(o["authorization"], d, "authorization", sv)); err != nil {
		if !fortiAPIPatch(o["authorization"]) {
			return fmt.Errorf("Error reading authorization: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenUserTacacsSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenUserTacacsInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenUserTacacsInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenUserTacacsVrfSelect(o["vrf-select"], d, "vrf_select", sv)); err != nil {
		if !fortiAPIPatch(o["vrf-select"]) {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func flattenUserTacacsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserTacacsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsSecondaryServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsTertiaryServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsSecondaryKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsTertiaryKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsStatusTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsAuthenType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsAuthorization(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsVrfSelect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserTacacs(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserTacacsName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandUserTacacsServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	} else if d.HasChange("server") {
		obj["server"] = nil
	}

	if v, ok := d.GetOk("secondary_server"); ok {
		t, err := expandUserTacacsSecondaryServer(d, v, "secondary_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-server"] = t
		}
	} else if d.HasChange("secondary_server") {
		obj["secondary-server"] = nil
	}

	if v, ok := d.GetOk("tertiary_server"); ok {
		t, err := expandUserTacacsTertiaryServer(d, v, "tertiary_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-server"] = t
		}
	} else if d.HasChange("tertiary_server") {
		obj["tertiary-server"] = nil
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandUserTacacsPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok {
		t, err := expandUserTacacsKey(d, v, "key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	} else if d.HasChange("key") {
		obj["key"] = nil
	}

	if v, ok := d.GetOk("secondary_key"); ok {
		t, err := expandUserTacacsSecondaryKey(d, v, "secondary_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-key"] = t
		}
	} else if d.HasChange("secondary_key") {
		obj["secondary-key"] = nil
	}

	if v, ok := d.GetOk("tertiary_key"); ok {
		t, err := expandUserTacacsTertiaryKey(d, v, "tertiary_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-key"] = t
		}
	} else if d.HasChange("tertiary_key") {
		obj["tertiary-key"] = nil
	}

	if v, ok := d.GetOkExists("status_ttl"); ok {
		t, err := expandUserTacacsStatusTtl(d, v, "status_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status-ttl"] = t
		}
	}

	if v, ok := d.GetOk("authen_type"); ok {
		t, err := expandUserTacacsAuthenType(d, v, "authen_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authen-type"] = t
		}
	}

	if v, ok := d.GetOk("authorization"); ok {
		t, err := expandUserTacacsAuthorization(d, v, "authorization", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorization"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandUserTacacsSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	} else if d.HasChange("source_ip") {
		obj["source-ip"] = nil
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		t, err := expandUserTacacsInterfaceSelectMethod(d, v, "interface_select_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandUserTacacsInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOkExists("vrf_select"); ok {
		t, err := expandUserTacacsVrfSelect(d, v, "vrf_select", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	} else if d.HasChange("vrf_select") {
		obj["vrf-select"] = nil
	}

	return &obj, nil
}
