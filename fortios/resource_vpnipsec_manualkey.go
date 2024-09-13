// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPsec manual keys.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnIpsecManualkey() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnIpsecManualkeyCreate,
		Read:   resourceVpnIpsecManualkeyRead,
		Update: resourceVpnIpsecManualkeyUpdate,
		Delete: resourceVpnIpsecManualkeyDelete,

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
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Required:     true,
			},
			"remote_gw": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"local_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"authkey": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"enckey": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"localspi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remotespi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"npu_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVpnIpsecManualkeyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnIpsecManualkey(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecManualkey resource while getting object: %v", err)
	}

	o, err := c.CreateVpnIpsecManualkey(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecManualkey resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecManualkey")
	}

	return resourceVpnIpsecManualkeyRead(d, m)
}

func resourceVpnIpsecManualkeyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnIpsecManualkey(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecManualkey resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnIpsecManualkey(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecManualkey resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecManualkey")
	}

	return resourceVpnIpsecManualkeyRead(d, m)
}

func resourceVpnIpsecManualkeyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteVpnIpsecManualkey(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting VpnIpsecManualkey resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecManualkeyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnIpsecManualkey(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecManualkey resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnIpsecManualkey(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecManualkey resource from API: %v", err)
	}
	return nil
}

func flattenVpnIpsecManualkeyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyRemoteGw(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyLocalGw(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyEncryption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyLocalspi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyRemotespi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyNpuOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnIpsecManualkey(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenVpnIpsecManualkeyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("interface", flattenVpnIpsecManualkeyInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("remote_gw", flattenVpnIpsecManualkeyRemoteGw(o["remote-gw"], d, "remote_gw", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw"]) {
			return fmt.Errorf("Error reading remote_gw: %v", err)
		}
	}

	if err = d.Set("local_gw", flattenVpnIpsecManualkeyLocalGw(o["local-gw"], d, "local_gw", sv)); err != nil {
		if !fortiAPIPatch(o["local-gw"]) {
			return fmt.Errorf("Error reading local_gw: %v", err)
		}
	}

	if err = d.Set("authentication", flattenVpnIpsecManualkeyAuthentication(o["authentication"], d, "authentication", sv)); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("encryption", flattenVpnIpsecManualkeyEncryption(o["encryption"], d, "encryption", sv)); err != nil {
		if !fortiAPIPatch(o["encryption"]) {
			return fmt.Errorf("Error reading encryption: %v", err)
		}
	}

	if err = d.Set("localspi", flattenVpnIpsecManualkeyLocalspi(o["localspi"], d, "localspi", sv)); err != nil {
		if !fortiAPIPatch(o["localspi"]) {
			return fmt.Errorf("Error reading localspi: %v", err)
		}
	}

	if err = d.Set("remotespi", flattenVpnIpsecManualkeyRemotespi(o["remotespi"], d, "remotespi", sv)); err != nil {
		if !fortiAPIPatch(o["remotespi"]) {
			return fmt.Errorf("Error reading remotespi: %v", err)
		}
	}

	if err = d.Set("npu_offload", flattenVpnIpsecManualkeyNpuOffload(o["npu-offload"], d, "npu_offload", sv)); err != nil {
		if !fortiAPIPatch(o["npu-offload"]) {
			return fmt.Errorf("Error reading npu_offload: %v", err)
		}
	}

	return nil
}

func flattenVpnIpsecManualkeyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnIpsecManualkeyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyRemoteGw(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyLocalGw(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyEncryption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyAuthkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyEnckey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyLocalspi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyRemotespi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyNpuOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnIpsecManualkey(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandVpnIpsecManualkeyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandVpnIpsecManualkeyInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOk("remote_gw"); ok {
		t, err := expandVpnIpsecManualkeyRemoteGw(d, v, "remote_gw", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw"] = t
		}
	}

	if v, ok := d.GetOk("local_gw"); ok {
		t, err := expandVpnIpsecManualkeyLocalGw(d, v, "local_gw", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-gw"] = t
		}
	}

	if v, ok := d.GetOk("authentication"); ok {
		t, err := expandVpnIpsecManualkeyAuthentication(d, v, "authentication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("encryption"); ok {
		t, err := expandVpnIpsecManualkeyEncryption(d, v, "encryption", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encryption"] = t
		}
	}

	if v, ok := d.GetOk("authkey"); ok {
		t, err := expandVpnIpsecManualkeyAuthkey(d, v, "authkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authkey"] = t
		}
	} else if d.HasChange("authkey") {
		obj["authkey"] = nil
	}

	if v, ok := d.GetOk("enckey"); ok {
		t, err := expandVpnIpsecManualkeyEnckey(d, v, "enckey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enckey"] = t
		}
	} else if d.HasChange("enckey") {
		obj["enckey"] = nil
	}

	if v, ok := d.GetOk("localspi"); ok {
		t, err := expandVpnIpsecManualkeyLocalspi(d, v, "localspi", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["localspi"] = t
		}
	} else if d.HasChange("localspi") {
		obj["localspi"] = nil
	}

	if v, ok := d.GetOk("remotespi"); ok {
		t, err := expandVpnIpsecManualkeyRemotespi(d, v, "remotespi", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remotespi"] = t
		}
	} else if d.HasChange("remotespi") {
		obj["remotespi"] = nil
	}

	if v, ok := d.GetOk("npu_offload"); ok {
		t, err := expandVpnIpsecManualkeyNpuOffload(d, v, "npu_offload", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-offload"] = t
		}
	}

	return &obj, nil
}
