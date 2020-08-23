// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure SSL servers.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallSslServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSslServerCreate,
		Read:   resourceFirewallSslServerRead,
		Update: resourceFirewallSslServerUpdate,
		Delete: resourceFirewallSslServerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Required:     true,
			},
			"ssl_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"add_header_x_forwarded_proto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mapped_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Required:     true,
			},
			"ssl_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
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
			"ssl_client_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_min_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_send_empty_frags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url_rewrite": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallSslServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallSslServer(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallSslServer resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallSslServer(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallSslServer resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallSslServer")
	}

	return resourceFirewallSslServerRead(d, m)
}

func resourceFirewallSslServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallSslServer(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslServer resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallSslServer(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallSslServer")
	}

	return resourceFirewallSslServerRead(d, m)
}

func resourceFirewallSslServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallSslServer(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallSslServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSslServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallSslServer(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSslServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSslServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSslServer resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSslServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslServerSslMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslServerAddHeaderXForwardedProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslServerMappedPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslServerSslCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslServerSslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslServerSslAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslServerSslClientRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslServerSslMinVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslServerSslMaxVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslServerSslSendEmptyFrags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslServerUrlRewrite(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallSslServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenFirewallSslServerName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ip", flattenFirewallSslServerIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("port", flattenFirewallSslServerPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("ssl_mode", flattenFirewallSslServerSslMode(o["ssl-mode"], d, "ssl_mode")); err != nil {
		if !fortiAPIPatch(o["ssl-mode"]) {
			return fmt.Errorf("Error reading ssl_mode: %v", err)
		}
	}

	if err = d.Set("add_header_x_forwarded_proto", flattenFirewallSslServerAddHeaderXForwardedProto(o["add-header-x-forwarded-proto"], d, "add_header_x_forwarded_proto")); err != nil {
		if !fortiAPIPatch(o["add-header-x-forwarded-proto"]) {
			return fmt.Errorf("Error reading add_header_x_forwarded_proto: %v", err)
		}
	}

	if err = d.Set("mapped_port", flattenFirewallSslServerMappedPort(o["mapped-port"], d, "mapped_port")); err != nil {
		if !fortiAPIPatch(o["mapped-port"]) {
			return fmt.Errorf("Error reading mapped_port: %v", err)
		}
	}

	if err = d.Set("ssl_cert", flattenFirewallSslServerSslCert(o["ssl-cert"], d, "ssl_cert")); err != nil {
		if !fortiAPIPatch(o["ssl-cert"]) {
			return fmt.Errorf("Error reading ssl_cert: %v", err)
		}
	}

	if err = d.Set("ssl_dh_bits", flattenFirewallSslServerSslDhBits(o["ssl-dh-bits"], d, "ssl_dh_bits")); err != nil {
		if !fortiAPIPatch(o["ssl-dh-bits"]) {
			return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenFirewallSslServerSslAlgorithm(o["ssl-algorithm"], d, "ssl_algorithm")); err != nil {
		if !fortiAPIPatch(o["ssl-algorithm"]) {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_client_renegotiation", flattenFirewallSslServerSslClientRenegotiation(o["ssl-client-renegotiation"], d, "ssl_client_renegotiation")); err != nil {
		if !fortiAPIPatch(o["ssl-client-renegotiation"]) {
			return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_min_version", flattenFirewallSslServerSslMinVersion(o["ssl-min-version"], d, "ssl_min_version")); err != nil {
		if !fortiAPIPatch(o["ssl-min-version"]) {
			return fmt.Errorf("Error reading ssl_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenFirewallSslServerSslMaxVersion(o["ssl-max-version"], d, "ssl_max_version")); err != nil {
		if !fortiAPIPatch(o["ssl-max-version"]) {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_send_empty_frags", flattenFirewallSslServerSslSendEmptyFrags(o["ssl-send-empty-frags"], d, "ssl_send_empty_frags")); err != nil {
		if !fortiAPIPatch(o["ssl-send-empty-frags"]) {
			return fmt.Errorf("Error reading ssl_send_empty_frags: %v", err)
		}
	}

	if err = d.Set("url_rewrite", flattenFirewallSslServerUrlRewrite(o["url-rewrite"], d, "url_rewrite")); err != nil {
		if !fortiAPIPatch(o["url-rewrite"]) {
			return fmt.Errorf("Error reading url_rewrite: %v", err)
		}
	}

	return nil
}

func flattenFirewallSslServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallSslServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslServerSslMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslServerAddHeaderXForwardedProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslServerMappedPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslServerSslCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslServerSslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslServerSslAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslServerSslClientRenegotiation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslServerSslMinVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslServerSslMaxVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslServerSslSendEmptyFrags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslServerUrlRewrite(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallSslServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallSslServerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandFirewallSslServerIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandFirewallSslServerPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mode"); ok {
		t, err := expandFirewallSslServerSslMode(d, v, "ssl_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mode"] = t
		}
	}

	if v, ok := d.GetOk("add_header_x_forwarded_proto"); ok {
		t, err := expandFirewallSslServerAddHeaderXForwardedProto(d, v, "add_header_x_forwarded_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-header-x-forwarded-proto"] = t
		}
	}

	if v, ok := d.GetOk("mapped_port"); ok {
		t, err := expandFirewallSslServerMappedPort(d, v, "mapped_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapped-port"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cert"); ok {
		t, err := expandFirewallSslServerSslCert(d, v, "ssl_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cert"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dh_bits"); ok {
		t, err := expandFirewallSslServerSslDhBits(d, v, "ssl_dh_bits")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-dh-bits"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok {
		t, err := expandFirewallSslServerSslAlgorithm(d, v, "ssl_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_renegotiation"); ok {
		t, err := expandFirewallSslServerSslClientRenegotiation(d, v, "ssl_client_renegotiation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_version"); ok {
		t, err := expandFirewallSslServerSslMinVersion(d, v, "ssl_min_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok {
		t, err := expandFirewallSslServerSslMaxVersion(d, v, "ssl_max_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_send_empty_frags"); ok {
		t, err := expandFirewallSslServerSslSendEmptyFrags(d, v, "ssl_send_empty_frags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-send-empty-frags"] = t
		}
	}

	if v, ok := d.GetOk("url_rewrite"); ok {
		t, err := expandFirewallSslServerUrlRewrite(d, v, "url_rewrite")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-rewrite"] = t
		}
	}

	return &obj, nil
}
