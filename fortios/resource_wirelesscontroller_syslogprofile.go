// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Wireless Termination Points (WTP) system log server profile.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerSyslogProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerSyslogProfileCreate,
		Read:   resourceWirelessControllerSyslogProfileRead,
		Update: resourceWirelessControllerSyslogProfileUpdate,
		Delete: resourceWirelessControllerSyslogProfileDelete,

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
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"server_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"server_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_fqdn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerSyslogProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerSyslogProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerSyslogProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerSyslogProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerSyslogProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerSyslogProfile")
	}

	return resourceWirelessControllerSyslogProfileRead(d, m)
}

func resourceWirelessControllerSyslogProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerSyslogProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSyslogProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerSyslogProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSyslogProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerSyslogProfile")
	}

	return resourceWirelessControllerSyslogProfileRead(d, m)
}

func resourceWirelessControllerSyslogProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerSyslogProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerSyslogProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerSyslogProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerSyslogProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSyslogProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerSyslogProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSyslogProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerSyslogProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSyslogProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSyslogProfileServerStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSyslogProfileServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSyslogProfileServerAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSyslogProfileServerFqdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSyslogProfileServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSyslogProfileServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerSyslogProfileServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSyslogProfileLogLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerSyslogProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerSyslogProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerSyslogProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("server_status", flattenWirelessControllerSyslogProfileServerStatus(o["server-status"], d, "server_status", sv)); err != nil {
		if !fortiAPIPatch(o["server-status"]) {
			return fmt.Errorf("Error reading server_status: %v", err)
		}
	}

	if err = d.Set("server", flattenWirelessControllerSyslogProfileServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("server_addr_type", flattenWirelessControllerSyslogProfileServerAddrType(o["server-addr-type"], d, "server_addr_type", sv)); err != nil {
		if !fortiAPIPatch(o["server-addr-type"]) {
			return fmt.Errorf("Error reading server_addr_type: %v", err)
		}
	}

	if err = d.Set("server_fqdn", flattenWirelessControllerSyslogProfileServerFqdn(o["server-fqdn"], d, "server_fqdn", sv)); err != nil {
		if !fortiAPIPatch(o["server-fqdn"]) {
			return fmt.Errorf("Error reading server_fqdn: %v", err)
		}
	}

	if err = d.Set("server_ip", flattenWirelessControllerSyslogProfileServerIp(o["server-ip"], d, "server_ip", sv)); err != nil {
		if !fortiAPIPatch(o["server-ip"]) {
			return fmt.Errorf("Error reading server_ip: %v", err)
		}
	}

	if err = d.Set("server_port", flattenWirelessControllerSyslogProfileServerPort(o["server-port"], d, "server_port", sv)); err != nil {
		if !fortiAPIPatch(o["server-port"]) {
			return fmt.Errorf("Error reading server_port: %v", err)
		}
	}

	if err = d.Set("server_type", flattenWirelessControllerSyslogProfileServerType(o["server-type"], d, "server_type", sv)); err != nil {
		if !fortiAPIPatch(o["server-type"]) {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("log_level", flattenWirelessControllerSyslogProfileLogLevel(o["log-level"], d, "log_level", sv)); err != nil {
		if !fortiAPIPatch(o["log-level"]) {
			return fmt.Errorf("Error reading log_level: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerSyslogProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerSyslogProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSyslogProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSyslogProfileServerStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSyslogProfileServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSyslogProfileServerAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSyslogProfileServerFqdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSyslogProfileServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSyslogProfileServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSyslogProfileServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSyslogProfileLogLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerSyslogProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerSyslogProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWirelessControllerSyslogProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("server_status"); ok {
		t, err := expandWirelessControllerSyslogProfileServerStatus(d, v, "server_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-status"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandWirelessControllerSyslogProfileServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	} else if d.HasChange("server") {
		obj["server"] = nil
	}

	if v, ok := d.GetOk("server_addr_type"); ok {
		t, err := expandWirelessControllerSyslogProfileServerAddrType(d, v, "server_addr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("server_fqdn"); ok {
		t, err := expandWirelessControllerSyslogProfileServerFqdn(d, v, "server_fqdn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-fqdn"] = t
		}
	} else if d.HasChange("server_fqdn") {
		obj["server-fqdn"] = nil
	}

	if v, ok := d.GetOk("server_ip"); ok {
		t, err := expandWirelessControllerSyslogProfileServerIp(d, v, "server_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-ip"] = t
		}
	}

	if v, ok := d.GetOkExists("server_port"); ok {
		t, err := expandWirelessControllerSyslogProfileServerPort(d, v, "server_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-port"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok {
		t, err := expandWirelessControllerSyslogProfileServerType(d, v, "server_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("log_level"); ok {
		t, err := expandWirelessControllerSyslogProfileLogLevel(d, v, "log_level", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-level"] = t
		}
	}

	return &obj, nil
}
