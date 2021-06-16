// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure explicit FTP proxy settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFtpProxyExplicit() *schema.Resource {
	return &schema.Resource{
		Create: resourceFtpProxyExplicitUpdate,
		Read:   resourceFtpProxyExplicitRead,
		Update: resourceFtpProxyExplicitUpdate,
		Delete: resourceFtpProxyExplicitDelete,

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
			"incoming_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"incoming_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"outgoing_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sec_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ssl_dh_bits": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFtpProxyExplicitUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFtpProxyExplicit(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FtpProxyExplicit resource while getting object: %v", err)
	}

	o, err := c.UpdateFtpProxyExplicit(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FtpProxyExplicit resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FtpProxyExplicit")
	}

	return resourceFtpProxyExplicitRead(d, m)
}

func resourceFtpProxyExplicitDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFtpProxyExplicit(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FtpProxyExplicit resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFtpProxyExplicitRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFtpProxyExplicit(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FtpProxyExplicit resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFtpProxyExplicit(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FtpProxyExplicit resource from API: %v", err)
	}
	return nil
}

func flattenFtpProxyExplicitStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFtpProxyExplicitIncomingPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFtpProxyExplicitIncomingIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFtpProxyExplicitOutgoingIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFtpProxyExplicitSecDefaultAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFtpProxyExplicitSsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFtpProxyExplicitSslCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFtpProxyExplicitSslDhBits(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFtpProxyExplicitSslAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFtpProxyExplicit(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenFtpProxyExplicitStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("incoming_port", flattenFtpProxyExplicitIncomingPort(o["incoming-port"], d, "incoming_port", sv)); err != nil {
		if !fortiAPIPatch(o["incoming-port"]) {
			return fmt.Errorf("Error reading incoming_port: %v", err)
		}
	}

	if err = d.Set("incoming_ip", flattenFtpProxyExplicitIncomingIp(o["incoming-ip"], d, "incoming_ip", sv)); err != nil {
		if !fortiAPIPatch(o["incoming-ip"]) {
			return fmt.Errorf("Error reading incoming_ip: %v", err)
		}
	}

	if err = d.Set("outgoing_ip", flattenFtpProxyExplicitOutgoingIp(o["outgoing-ip"], d, "outgoing_ip", sv)); err != nil {
		if !fortiAPIPatch(o["outgoing-ip"]) {
			return fmt.Errorf("Error reading outgoing_ip: %v", err)
		}
	}

	if err = d.Set("sec_default_action", flattenFtpProxyExplicitSecDefaultAction(o["sec-default-action"], d, "sec_default_action", sv)); err != nil {
		if !fortiAPIPatch(o["sec-default-action"]) {
			return fmt.Errorf("Error reading sec_default_action: %v", err)
		}
	}

	if err = d.Set("ssl", flattenFtpProxyExplicitSsl(o["ssl"], d, "ssl", sv)); err != nil {
		if !fortiAPIPatch(o["ssl"]) {
			return fmt.Errorf("Error reading ssl: %v", err)
		}
	}

	if err = d.Set("ssl_cert", flattenFtpProxyExplicitSslCert(o["ssl-cert"], d, "ssl_cert", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-cert"]) {
			return fmt.Errorf("Error reading ssl_cert: %v", err)
		}
	}

	if err = d.Set("ssl_dh_bits", flattenFtpProxyExplicitSslDhBits(o["ssl-dh-bits"], d, "ssl_dh_bits", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-dh-bits"]) {
			return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenFtpProxyExplicitSslAlgorithm(o["ssl-algorithm"], d, "ssl_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-algorithm"]) {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	return nil
}

func flattenFtpProxyExplicitFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFtpProxyExplicitStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFtpProxyExplicitIncomingPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFtpProxyExplicitIncomingIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFtpProxyExplicitOutgoingIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFtpProxyExplicitSecDefaultAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFtpProxyExplicitSsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFtpProxyExplicitSslCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFtpProxyExplicitSslDhBits(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFtpProxyExplicitSslAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFtpProxyExplicit(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {

		t, err := expandFtpProxyExplicitStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("incoming_port"); ok {

		t, err := expandFtpProxyExplicitIncomingPort(d, v, "incoming_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incoming-port"] = t
		}
	}

	if v, ok := d.GetOk("incoming_ip"); ok {

		t, err := expandFtpProxyExplicitIncomingIp(d, v, "incoming_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incoming-ip"] = t
		}
	}

	if v, ok := d.GetOk("outgoing_ip"); ok {

		t, err := expandFtpProxyExplicitOutgoingIp(d, v, "outgoing_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outgoing-ip"] = t
		}
	}

	if v, ok := d.GetOk("sec_default_action"); ok {

		t, err := expandFtpProxyExplicitSecDefaultAction(d, v, "sec_default_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sec-default-action"] = t
		}
	}

	if v, ok := d.GetOk("ssl"); ok {

		t, err := expandFtpProxyExplicitSsl(d, v, "ssl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cert"); ok {

		t, err := expandFtpProxyExplicitSslCert(d, v, "ssl_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cert"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dh_bits"); ok {

		t, err := expandFtpProxyExplicitSslDhBits(d, v, "ssl_dh_bits", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-dh-bits"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok {

		t, err := expandFtpProxyExplicitSslAlgorithm(d, v, "ssl_algorithm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-algorithm"] = t
		}
	}

	return &obj, nil
}
