// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure LoRaWAN profile.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerLwProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerLwProfileCreate,
		Read:   resourceWirelessControllerLwProfileRead,
		Update: resourceWirelessControllerLwProfileUpdate,
		Delete: resourceWirelessControllerLwProfileDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"lw_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cups_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"cups_server_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
			},
			"cups_api_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"tc_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"tc_server_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
			},
			"tc_api_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
		},
	}
}

func resourceWirelessControllerLwProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerLwProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerLwProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerLwProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerLwProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerLwProfile")
	}

	return resourceWirelessControllerLwProfileRead(d, m)
}

func resourceWirelessControllerLwProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerLwProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerLwProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerLwProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerLwProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerLwProfile")
	}

	return resourceWirelessControllerLwProfileRead(d, m)
}

func resourceWirelessControllerLwProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerLwProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerLwProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerLwProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerLwProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerLwProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerLwProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerLwProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerLwProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerLwProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerLwProfileLwProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerLwProfileCupsServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerLwProfileCupsServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerLwProfileTcServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerLwProfileTcServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectWirelessControllerLwProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerLwProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerLwProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("lw_protocol", flattenWirelessControllerLwProfileLwProtocol(o["lw-protocol"], d, "lw_protocol", sv)); err != nil {
		if !fortiAPIPatch(o["lw-protocol"]) {
			return fmt.Errorf("Error reading lw_protocol: %v", err)
		}
	}

	if err = d.Set("cups_server", flattenWirelessControllerLwProfileCupsServer(o["cups-server"], d, "cups_server", sv)); err != nil {
		if !fortiAPIPatch(o["cups-server"]) {
			return fmt.Errorf("Error reading cups_server: %v", err)
		}
	}

	if err = d.Set("cups_server_port", flattenWirelessControllerLwProfileCupsServerPort(o["cups-server-port"], d, "cups_server_port", sv)); err != nil {
		if !fortiAPIPatch(o["cups-server-port"]) {
			return fmt.Errorf("Error reading cups_server_port: %v", err)
		}
	}

	if err = d.Set("tc_server", flattenWirelessControllerLwProfileTcServer(o["tc-server"], d, "tc_server", sv)); err != nil {
		if !fortiAPIPatch(o["tc-server"]) {
			return fmt.Errorf("Error reading tc_server: %v", err)
		}
	}

	if err = d.Set("tc_server_port", flattenWirelessControllerLwProfileTcServerPort(o["tc-server-port"], d, "tc_server_port", sv)); err != nil {
		if !fortiAPIPatch(o["tc-server-port"]) {
			return fmt.Errorf("Error reading tc_server_port: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerLwProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerLwProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLwProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLwProfileLwProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLwProfileCupsServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLwProfileCupsServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLwProfileCupsApiKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLwProfileTcServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLwProfileTcServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLwProfileTcApiKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerLwProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerLwProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWirelessControllerLwProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("lw_protocol"); ok {
		t, err := expandWirelessControllerLwProfileLwProtocol(d, v, "lw_protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lw-protocol"] = t
		}
	}

	if v, ok := d.GetOk("cups_server"); ok {
		t, err := expandWirelessControllerLwProfileCupsServer(d, v, "cups_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cups-server"] = t
		}
	} else if d.HasChange("cups_server") {
		obj["cups-server"] = nil
	}

	if v, ok := d.GetOkExists("cups_server_port"); ok {
		t, err := expandWirelessControllerLwProfileCupsServerPort(d, v, "cups_server_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cups-server-port"] = t
		}
	} else if d.HasChange("cups_server_port") {
		obj["cups-server-port"] = nil
	}

	if v, ok := d.GetOk("cups_api_key"); ok {
		t, err := expandWirelessControllerLwProfileCupsApiKey(d, v, "cups_api_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cups-api-key"] = t
		}
	} else if d.HasChange("cups_api_key") {
		obj["cups-api-key"] = nil
	}

	if v, ok := d.GetOk("tc_server"); ok {
		t, err := expandWirelessControllerLwProfileTcServer(d, v, "tc_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tc-server"] = t
		}
	} else if d.HasChange("tc_server") {
		obj["tc-server"] = nil
	}

	if v, ok := d.GetOkExists("tc_server_port"); ok {
		t, err := expandWirelessControllerLwProfileTcServerPort(d, v, "tc_server_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tc-server-port"] = t
		}
	} else if d.HasChange("tc_server_port") {
		obj["tc-server-port"] = nil
	}

	if v, ok := d.GetOk("tc_api_key"); ok {
		t, err := expandWirelessControllerLwProfileTcApiKey(d, v, "tc_api_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tc-api-key"] = t
		}
	} else if d.HasChange("tc_api_key") {
		obj["tc-api-key"] = nil
	}

	return &obj, nil
}
