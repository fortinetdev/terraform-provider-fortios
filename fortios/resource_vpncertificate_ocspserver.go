// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: OCSP server configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceVpnCertificateOcspServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnCertificateOcspServerCreate,
		Read:   resourceVpnCertificateOcspServerRead,
		Update: resourceVpnCertificateOcspServerUpdate,
		Delete: resourceVpnCertificateOcspServerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"secondary_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"secondary_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"unavail_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVpnCertificateOcspServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnCertificateOcspServer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateOcspServer resource while getting object: %v", err)
	}

	o, err := c.CreateVpnCertificateOcspServer(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateOcspServer resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnCertificateOcspServer")
	}

	return resourceVpnCertificateOcspServerRead(d, m)
}

func resourceVpnCertificateOcspServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnCertificateOcspServer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateOcspServer resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnCertificateOcspServer(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateOcspServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnCertificateOcspServer")
	}

	return resourceVpnCertificateOcspServerRead(d, m)
}

func resourceVpnCertificateOcspServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteVpnCertificateOcspServer(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting VpnCertificateOcspServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateOcspServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadVpnCertificateOcspServer(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateOcspServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnCertificateOcspServer(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateOcspServer resource from API: %v", err)
	}
	return nil
}

func flattenVpnCertificateOcspServerName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateOcspServerUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateOcspServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateOcspServerSecondaryUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateOcspServerSecondaryCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateOcspServerUnavailAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateOcspServerSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnCertificateOcspServer(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenVpnCertificateOcspServerName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("url", flattenVpnCertificateOcspServerUrl(o["url"], d, "url", sv)); err != nil {
		if !fortiAPIPatch(o["url"]) {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	if err = d.Set("cert", flattenVpnCertificateOcspServerCert(o["cert"], d, "cert", sv)); err != nil {
		if !fortiAPIPatch(o["cert"]) {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("secondary_url", flattenVpnCertificateOcspServerSecondaryUrl(o["secondary-url"], d, "secondary_url", sv)); err != nil {
		if !fortiAPIPatch(o["secondary-url"]) {
			return fmt.Errorf("Error reading secondary_url: %v", err)
		}
	}

	if err = d.Set("secondary_cert", flattenVpnCertificateOcspServerSecondaryCert(o["secondary-cert"], d, "secondary_cert", sv)); err != nil {
		if !fortiAPIPatch(o["secondary-cert"]) {
			return fmt.Errorf("Error reading secondary_cert: %v", err)
		}
	}

	if err = d.Set("unavail_action", flattenVpnCertificateOcspServerUnavailAction(o["unavail-action"], d, "unavail_action", sv)); err != nil {
		if !fortiAPIPatch(o["unavail-action"]) {
			return fmt.Errorf("Error reading unavail_action: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenVpnCertificateOcspServerSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	return nil
}

func flattenVpnCertificateOcspServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnCertificateOcspServerName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateOcspServerUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateOcspServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateOcspServerSecondaryUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateOcspServerSecondaryCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateOcspServerUnavailAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateOcspServerSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnCertificateOcspServer(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandVpnCertificateOcspServerName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok {

		t, err := expandVpnCertificateOcspServerUrl(d, v, "url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	if v, ok := d.GetOk("cert"); ok {

		t, err := expandVpnCertificateOcspServerCert(d, v, "cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert"] = t
		}
	}

	if v, ok := d.GetOk("secondary_url"); ok {

		t, err := expandVpnCertificateOcspServerSecondaryUrl(d, v, "secondary_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-url"] = t
		}
	}

	if v, ok := d.GetOk("secondary_cert"); ok {

		t, err := expandVpnCertificateOcspServerSecondaryCert(d, v, "secondary_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-cert"] = t
		}
	}

	if v, ok := d.GetOk("unavail_action"); ok {

		t, err := expandVpnCertificateOcspServerUnavailAction(d, v, "unavail_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unavail-action"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {

		t, err := expandVpnCertificateOcspServerSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	return &obj, nil
}
