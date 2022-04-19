// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure virtual network enabler tunnel.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemVneTunnel() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVneTunnelUpdate,
		Read:   resourceSystemVneTunnelRead,
		Update: resourceSystemVneTunnelUpdate,
		Delete: resourceSystemVneTunnelDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
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
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
				Computed:     true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemVneTunnelUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemVneTunnel(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemVneTunnel resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemVneTunnel(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemVneTunnel resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVneTunnel")
	}

	return resourceSystemVneTunnelRead(d, m)
}

func resourceSystemVneTunnelDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemVneTunnel(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemVneTunnel resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemVneTunnel(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemVneTunnel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVneTunnelRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemVneTunnel(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemVneTunnel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVneTunnel(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemVneTunnel resource from API: %v", err)
	}
	return nil
}

func flattenSystemVneTunnelStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVneTunnelInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVneTunnelSslCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVneTunnelBmrHostname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVneTunnelAutoAsicOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVneTunnelIpv4Address(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVneTunnelBr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVneTunnelUpdateUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVneTunnelMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemVneTunnel(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemVneTunnelStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemVneTunnelInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", flattenSystemVneTunnelSslCertificate(o["ssl-certificate"], d, "ssl_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-certificate"]) {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if err = d.Set("auto_asic_offload", flattenSystemVneTunnelAutoAsicOffload(o["auto-asic-offload"], d, "auto_asic_offload", sv)); err != nil {
		if !fortiAPIPatch(o["auto-asic-offload"]) {
			return fmt.Errorf("Error reading auto_asic_offload: %v", err)
		}
	}

	if err = d.Set("ipv4_address", flattenSystemVneTunnelIpv4Address(o["ipv4-address"], d, "ipv4_address", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-address"]) {
			return fmt.Errorf("Error reading ipv4_address: %v", err)
		}
	}

	if err = d.Set("br", flattenSystemVneTunnelBr(o["br"], d, "br", sv)); err != nil {
		if !fortiAPIPatch(o["br"]) {
			return fmt.Errorf("Error reading br: %v", err)
		}
	}

	if err = d.Set("update_url", flattenSystemVneTunnelUpdateUrl(o["update-url"], d, "update_url", sv)); err != nil {
		if !fortiAPIPatch(o["update-url"]) {
			return fmt.Errorf("Error reading update_url: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemVneTunnelMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	return nil
}

func flattenSystemVneTunnelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemVneTunnelStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVneTunnelInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVneTunnelSslCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVneTunnelBmrHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVneTunnelAutoAsicOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVneTunnelIpv4Address(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVneTunnelBr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVneTunnelUpdateUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVneTunnelMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVneTunnel(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandSystemVneTunnelStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		if setArgNil {
			obj["interface"] = nil
		} else {

			t, err := expandSystemVneTunnelInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssl_certificate"); ok {
		if setArgNil {
			obj["ssl-certificate"] = nil
		} else {

			t, err := expandSystemVneTunnelSslCertificate(d, v, "ssl_certificate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssl-certificate"] = t
			}
		}
	}

	if v, ok := d.GetOk("bmr_hostname"); ok {
		if setArgNil {
			obj["bmr-hostname"] = nil
		} else {

			t, err := expandSystemVneTunnelBmrHostname(d, v, "bmr_hostname", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bmr-hostname"] = t
			}
		}
	}

	if v, ok := d.GetOk("auto_asic_offload"); ok {
		if setArgNil {
			obj["auto-asic-offload"] = nil
		} else {

			t, err := expandSystemVneTunnelAutoAsicOffload(d, v, "auto_asic_offload", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-asic-offload"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipv4_address"); ok {
		if setArgNil {
			obj["ipv4-address"] = nil
		} else {

			t, err := expandSystemVneTunnelIpv4Address(d, v, "ipv4_address", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipv4-address"] = t
			}
		}
	}

	if v, ok := d.GetOk("br"); ok {
		if setArgNil {
			obj["br"] = nil
		} else {

			t, err := expandSystemVneTunnelBr(d, v, "br", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["br"] = t
			}
		}
	}

	if v, ok := d.GetOk("update_url"); ok {
		if setArgNil {
			obj["update-url"] = nil
		} else {

			t, err := expandSystemVneTunnelUpdateUrl(d, v, "update_url", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["update-url"] = t
			}
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		if setArgNil {
			obj["mode"] = nil
		} else {

			t, err := expandSystemVneTunnelMode(d, v, "mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mode"] = t
			}
		}
	}

	return &obj, nil
}
