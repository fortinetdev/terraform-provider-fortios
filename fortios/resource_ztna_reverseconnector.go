// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure ZTNA Reverse-Connector.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaReverseConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaReverseConnectorCreate,
		Read:   resourceZtnaReverseConnectorRead,
		Update: resourceZtnaReverseConnectorUpdate,
		Delete: resourceZtnaReverseConnectorDelete,

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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
			},
			"health_check_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 600),
				Optional:     true,
				Computed:     true,
			},
			"ssl_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certificate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"trusted_server_ca": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
		},
	}
}

func resourceZtnaReverseConnectorCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectZtnaReverseConnector(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaReverseConnector resource while getting object: %v", err)
	}

	o, err := c.CreateZtnaReverseConnector(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating ZtnaReverseConnector resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ZtnaReverseConnector")
	}

	return resourceZtnaReverseConnectorRead(d, m)
}

func resourceZtnaReverseConnectorUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectZtnaReverseConnector(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaReverseConnector resource while getting object: %v", err)
	}

	o, err := c.UpdateZtnaReverseConnector(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaReverseConnector resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ZtnaReverseConnector")
	}

	return resourceZtnaReverseConnectorRead(d, m)
}

func resourceZtnaReverseConnectorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteZtnaReverseConnector(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaReverseConnector resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaReverseConnectorRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadZtnaReverseConnector(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaReverseConnector resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaReverseConnector(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaReverseConnector resource from API: %v", err)
	}
	return nil
}

func flattenZtnaReverseConnectorName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaReverseConnectorStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaReverseConnectorAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaReverseConnectorPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaReverseConnectorHealthCheckInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaReverseConnectorSslMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaReverseConnectorCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaReverseConnectorTrustedServerCa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectZtnaReverseConnector(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenZtnaReverseConnectorName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenZtnaReverseConnectorStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("address", flattenZtnaReverseConnectorAddress(o["address"], d, "address", sv)); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("port", flattenZtnaReverseConnectorPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("health_check_interval", flattenZtnaReverseConnectorHealthCheckInterval(o["health-check-interval"], d, "health_check_interval", sv)); err != nil {
		if !fortiAPIPatch(o["health-check-interval"]) {
			return fmt.Errorf("Error reading health_check_interval: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenZtnaReverseConnectorSslMaxVersion(o["ssl-max-version"], d, "ssl_max_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-max-version"]) {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("certificate", flattenZtnaReverseConnectorCertificate(o["certificate"], d, "certificate", sv)); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("trusted_server_ca", flattenZtnaReverseConnectorTrustedServerCa(o["trusted-server-ca"], d, "trusted_server_ca", sv)); err != nil {
		if !fortiAPIPatch(o["trusted-server-ca"]) {
			return fmt.Errorf("Error reading trusted_server_ca: %v", err)
		}
	}

	return nil
}

func flattenZtnaReverseConnectorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandZtnaReverseConnectorName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaReverseConnectorStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaReverseConnectorAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaReverseConnectorPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaReverseConnectorHealthCheckInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaReverseConnectorSslMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaReverseConnectorCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaReverseConnectorTrustedServerCa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectZtnaReverseConnector(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandZtnaReverseConnectorName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandZtnaReverseConnectorStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("address"); ok {
		t, err := expandZtnaReverseConnectorAddress(d, v, "address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	} else if d.HasChange("address") {
		obj["address"] = nil
	}

	if v, ok := d.GetOkExists("port"); ok {
		t, err := expandZtnaReverseConnectorPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	} else if d.HasChange("port") {
		obj["port"] = nil
	}

	if v, ok := d.GetOkExists("health_check_interval"); ok {
		t, err := expandZtnaReverseConnectorHealthCheckInterval(d, v, "health_check_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-interval"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok {
		t, err := expandZtnaReverseConnectorSslMaxVersion(d, v, "ssl_max_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok {
		t, err := expandZtnaReverseConnectorCertificate(d, v, "certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	} else if d.HasChange("certificate") {
		obj["certificate"] = nil
	}

	if v, ok := d.GetOk("trusted_server_ca"); ok {
		t, err := expandZtnaReverseConnectorTrustedServerCa(d, v, "trusted_server_ca", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusted-server-ca"] = t
		}
	} else if d.HasChange("trusted_server_ca") {
		obj["trusted-server-ca"] = nil
	}

	return &obj, nil
}
