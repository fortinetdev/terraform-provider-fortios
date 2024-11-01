// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure virtual network enabler tunnels.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemVneInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVneInterfaceCreate,
		Read:   resourceSystemVneInterfaceRead,
		Update: resourceSystemVneInterfaceUpdate,
		Delete: resourceSystemVneInterfaceDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 15),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"ssl_certificate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"bmr_hostname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"auto_asic_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"br": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"update_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
			},
			"http_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
		},
	}
}

func resourceSystemVneInterfaceCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemVneInterface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemVneInterface resource while getting object: %v", err)
	}

	o, err := c.CreateSystemVneInterface(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemVneInterface resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVneInterface")
	}

	return resourceSystemVneInterfaceRead(d, m)
}

func resourceSystemVneInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemVneInterface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemVneInterface resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemVneInterface(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemVneInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVneInterface")
	}

	return resourceSystemVneInterfaceRead(d, m)
}

func resourceSystemVneInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemVneInterface(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVneInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVneInterfaceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemVneInterface(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemVneInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVneInterface(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemVneInterface resource from API: %v", err)
	}
	return nil
}

func flattenSystemVneInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVneInterfaceInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVneInterfaceSslCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVneInterfaceAutoAsicOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVneInterfaceIpv4Address(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemVneInterfaceBr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVneInterfaceUpdateUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVneInterfaceMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVneInterfaceHttpUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemVneInterface(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemVneInterfaceName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemVneInterfaceInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", flattenSystemVneInterfaceSslCertificate(o["ssl-certificate"], d, "ssl_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-certificate"]) {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if err = d.Set("auto_asic_offload", flattenSystemVneInterfaceAutoAsicOffload(o["auto-asic-offload"], d, "auto_asic_offload", sv)); err != nil {
		if !fortiAPIPatch(o["auto-asic-offload"]) {
			return fmt.Errorf("Error reading auto_asic_offload: %v", err)
		}
	}

	if err = d.Set("ipv4_address", flattenSystemVneInterfaceIpv4Address(o["ipv4-address"], d, "ipv4_address", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-address"]) {
			return fmt.Errorf("Error reading ipv4_address: %v", err)
		}
	}

	if err = d.Set("br", flattenSystemVneInterfaceBr(o["br"], d, "br", sv)); err != nil {
		if !fortiAPIPatch(o["br"]) {
			return fmt.Errorf("Error reading br: %v", err)
		}
	}

	if err = d.Set("update_url", flattenSystemVneInterfaceUpdateUrl(o["update-url"], d, "update_url", sv)); err != nil {
		if !fortiAPIPatch(o["update-url"]) {
			return fmt.Errorf("Error reading update_url: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemVneInterfaceMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("http_username", flattenSystemVneInterfaceHttpUsername(o["http-username"], d, "http_username", sv)); err != nil {
		if !fortiAPIPatch(o["http-username"]) {
			return fmt.Errorf("Error reading http_username: %v", err)
		}
	}

	return nil
}

func flattenSystemVneInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemVneInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVneInterfaceInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVneInterfaceSslCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVneInterfaceBmrHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVneInterfaceAutoAsicOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVneInterfaceIpv4Address(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVneInterfaceBr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVneInterfaceUpdateUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVneInterfaceMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVneInterfaceHttpUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVneInterfaceHttpPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVneInterface(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemVneInterfaceName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemVneInterfaceInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOk("ssl_certificate"); ok {
		t, err := expandSystemVneInterfaceSslCertificate(d, v, "ssl_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-certificate"] = t
		}
	}

	if v, ok := d.GetOk("bmr_hostname"); ok {
		t, err := expandSystemVneInterfaceBmrHostname(d, v, "bmr_hostname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bmr-hostname"] = t
		}
	} else if d.HasChange("bmr_hostname") {
		obj["bmr-hostname"] = nil
	}

	if v, ok := d.GetOk("auto_asic_offload"); ok {
		t, err := expandSystemVneInterfaceAutoAsicOffload(d, v, "auto_asic_offload", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-asic-offload"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_address"); ok {
		t, err := expandSystemVneInterfaceIpv4Address(d, v, "ipv4_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-address"] = t
		}
	}

	if v, ok := d.GetOk("br"); ok {
		t, err := expandSystemVneInterfaceBr(d, v, "br", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["br"] = t
		}
	} else if d.HasChange("br") {
		obj["br"] = nil
	}

	if v, ok := d.GetOk("update_url"); ok {
		t, err := expandSystemVneInterfaceUpdateUrl(d, v, "update_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-url"] = t
		}
	} else if d.HasChange("update_url") {
		obj["update-url"] = nil
	}

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandSystemVneInterfaceMode(d, v, "mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("http_username"); ok {
		t, err := expandSystemVneInterfaceHttpUsername(d, v, "http_username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-username"] = t
		}
	} else if d.HasChange("http_username") {
		obj["http-username"] = nil
	}

	if v, ok := d.GetOk("http_password"); ok {
		t, err := expandSystemVneInterfaceHttpPassword(d, v, "http_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-password"] = t
		}
	} else if d.HasChange("http_password") {
		obj["http-password"] = nil
	}

	return &obj, nil
}
