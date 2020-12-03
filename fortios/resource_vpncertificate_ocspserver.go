// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: OCSP server configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"

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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
				ForceNew:     true,
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

	obj, err := getObjectVpnCertificateOcspServer(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateOcspServer resource while getting object: %v", err)
	}

	o, err := c.CreateVpnCertificateOcspServer(obj)

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

	obj, err := getObjectVpnCertificateOcspServer(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateOcspServer resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnCertificateOcspServer(obj, mkey)
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

	err := c.DeleteVpnCertificateOcspServer(mkey)
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

	o, err := c.ReadVpnCertificateOcspServer(mkey)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateOcspServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnCertificateOcspServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateOcspServer resource from API: %v", err)
	}
	return nil
}

func flattenVpnCertificateOcspServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateOcspServerUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateOcspServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateOcspServerSecondaryUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateOcspServerSecondaryCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateOcspServerUnavailAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateOcspServerSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnCertificateOcspServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenVpnCertificateOcspServerName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("url", flattenVpnCertificateOcspServerUrl(o["url"], d, "url")); err != nil {
		if !fortiAPIPatch(o["url"]) {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	if err = d.Set("cert", flattenVpnCertificateOcspServerCert(o["cert"], d, "cert")); err != nil {
		if !fortiAPIPatch(o["cert"]) {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("secondary_url", flattenVpnCertificateOcspServerSecondaryUrl(o["secondary-url"], d, "secondary_url")); err != nil {
		if !fortiAPIPatch(o["secondary-url"]) {
			return fmt.Errorf("Error reading secondary_url: %v", err)
		}
	}

	if err = d.Set("secondary_cert", flattenVpnCertificateOcspServerSecondaryCert(o["secondary-cert"], d, "secondary_cert")); err != nil {
		if !fortiAPIPatch(o["secondary-cert"]) {
			return fmt.Errorf("Error reading secondary_cert: %v", err)
		}
	}

	if err = d.Set("unavail_action", flattenVpnCertificateOcspServerUnavailAction(o["unavail-action"], d, "unavail_action")); err != nil {
		if !fortiAPIPatch(o["unavail-action"]) {
			return fmt.Errorf("Error reading unavail_action: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenVpnCertificateOcspServerSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	return nil
}

func flattenVpnCertificateOcspServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnCertificateOcspServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateOcspServerUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateOcspServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateOcspServerSecondaryUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateOcspServerSecondaryCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateOcspServerUnavailAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateOcspServerSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnCertificateOcspServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandVpnCertificateOcspServerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok {
		t, err := expandVpnCertificateOcspServerUrl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	if v, ok := d.GetOk("cert"); ok {
		t, err := expandVpnCertificateOcspServerCert(d, v, "cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert"] = t
		}
	}

	if v, ok := d.GetOk("secondary_url"); ok {
		t, err := expandVpnCertificateOcspServerSecondaryUrl(d, v, "secondary_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-url"] = t
		}
	}

	if v, ok := d.GetOk("secondary_cert"); ok {
		t, err := expandVpnCertificateOcspServerSecondaryCert(d, v, "secondary_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-cert"] = t
		}
	}

	if v, ok := d.GetOk("unavail_action"); ok {
		t, err := expandVpnCertificateOcspServerUnavailAction(d, v, "unavail_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unavail-action"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandVpnCertificateOcspServerSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	return &obj, nil
}
