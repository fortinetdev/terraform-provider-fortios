// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure ICAP servers.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceIcapServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceIcapServerCreate,
		Read:   resourceIcapServerRead,
		Update: resourceIcapServerUpdate,
		Delete: resourceIcapServerDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fqdn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"max_connections": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"secure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"healthcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"healthcheck_service": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceIcapServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectIcapServer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating IcapServer resource while getting object: %v", err)
	}

	o, err := c.CreateIcapServer(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating IcapServer resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("IcapServer")
	}

	return resourceIcapServerRead(d, m)
}

func resourceIcapServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectIcapServer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating IcapServer resource while getting object: %v", err)
	}

	o, err := c.UpdateIcapServer(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating IcapServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("IcapServer")
	}

	return resourceIcapServerRead(d, m)
}

func resourceIcapServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteIcapServer(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting IcapServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIcapServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadIcapServer(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading IcapServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIcapServer(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading IcapServer resource from API: %v", err)
	}
	return nil
}

func flattenIcapServerName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapServerAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapServerIpVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapServerIpAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapServerIp6Address(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapServerFqdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapServerMaxConnections(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapServerSecure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapServerSslCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapServerHealthcheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapServerHealthcheckService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectIcapServer(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenIcapServerName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("addr_type", flattenIcapServerAddrType(o["addr-type"], d, "addr_type", sv)); err != nil {
		if !fortiAPIPatch(o["addr-type"]) {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("ip_version", flattenIcapServerIpVersion(o["ip-version"], d, "ip_version", sv)); err != nil {
		if !fortiAPIPatch(o["ip-version"]) {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if err = d.Set("ip_address", flattenIcapServerIpAddress(o["ip-address"], d, "ip_address", sv)); err != nil {
		if !fortiAPIPatch(o["ip-address"]) {
			return fmt.Errorf("Error reading ip_address: %v", err)
		}
	}

	if err = d.Set("ip6_address", flattenIcapServerIp6Address(o["ip6-address"], d, "ip6_address", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-address"]) {
			return fmt.Errorf("Error reading ip6_address: %v", err)
		}
	}

	if err = d.Set("fqdn", flattenIcapServerFqdn(o["fqdn"], d, "fqdn", sv)); err != nil {
		if !fortiAPIPatch(o["fqdn"]) {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("port", flattenIcapServerPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("max_connections", flattenIcapServerMaxConnections(o["max-connections"], d, "max_connections", sv)); err != nil {
		if !fortiAPIPatch(o["max-connections"]) {
			return fmt.Errorf("Error reading max_connections: %v", err)
		}
	}

	if err = d.Set("secure", flattenIcapServerSecure(o["secure"], d, "secure", sv)); err != nil {
		if !fortiAPIPatch(o["secure"]) {
			return fmt.Errorf("Error reading secure: %v", err)
		}
	}

	if err = d.Set("ssl_cert", flattenIcapServerSslCert(o["ssl-cert"], d, "ssl_cert", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-cert"]) {
			return fmt.Errorf("Error reading ssl_cert: %v", err)
		}
	}

	if err = d.Set("healthcheck", flattenIcapServerHealthcheck(o["healthcheck"], d, "healthcheck", sv)); err != nil {
		if !fortiAPIPatch(o["healthcheck"]) {
			return fmt.Errorf("Error reading healthcheck: %v", err)
		}
	}

	if err = d.Set("healthcheck_service", flattenIcapServerHealthcheckService(o["healthcheck-service"], d, "healthcheck_service", sv)); err != nil {
		if !fortiAPIPatch(o["healthcheck-service"]) {
			return fmt.Errorf("Error reading healthcheck_service: %v", err)
		}
	}

	return nil
}

func flattenIcapServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandIcapServerName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapServerAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapServerIpVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapServerIpAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapServerIp6Address(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapServerFqdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapServerMaxConnections(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapServerSecure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapServerSslCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapServerHealthcheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapServerHealthcheckService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectIcapServer(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandIcapServerName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("addr_type"); ok {

		t, err := expandIcapServerAddrType(d, v, "addr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("ip_version"); ok {

		t, err := expandIcapServerIpVersion(d, v, "ip_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-version"] = t
		}
	}

	if v, ok := d.GetOk("ip_address"); ok {

		t, err := expandIcapServerIpAddress(d, v, "ip_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-address"] = t
		}
	}

	if v, ok := d.GetOk("ip6_address"); ok {

		t, err := expandIcapServerIp6Address(d, v, "ip6_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-address"] = t
		}
	}

	if v, ok := d.GetOk("fqdn"); ok {

		t, err := expandIcapServerFqdn(d, v, "fqdn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {

		t, err := expandIcapServerPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("max_connections"); ok {

		t, err := expandIcapServerMaxConnections(d, v, "max_connections", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-connections"] = t
		}
	}

	if v, ok := d.GetOk("secure"); ok {

		t, err := expandIcapServerSecure(d, v, "secure", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cert"); ok {

		t, err := expandIcapServerSslCert(d, v, "ssl_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cert"] = t
		}
	}

	if v, ok := d.GetOk("healthcheck"); ok {

		t, err := expandIcapServerHealthcheck(d, v, "healthcheck", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["healthcheck"] = t
		}
	}

	if v, ok := d.GetOk("healthcheck_service"); ok {

		t, err := expandIcapServerHealthcheckService(d, v, "healthcheck_service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["healthcheck-service"] = t
		}
	}

	return &obj, nil
}
