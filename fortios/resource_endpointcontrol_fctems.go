// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiClient Enterprise Management Server (EMS) entries.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceEndpointControlFctems() *schema.Resource {
	return &schema.Resource{
		Create: resourceEndpointControlFctemsCreate,
		Read:   resourceEndpointControlFctemsRead,
		Update: resourceEndpointControlFctemsUpdate,
		Delete: resourceEndpointControlFctemsDelete,

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
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"serial_number": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 16),
				Optional:     true,
				Computed:     true,
			},
			"fortinetone_cloud_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"https_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"admin_username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Computed:     true,
			},
			"admin_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 68),
				Optional:     true,
				Sensitive:    true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pull_sysinfo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pull_vulnerabilities": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pull_avatars": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pull_tags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cloud_server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"call_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(500, 30000),
				Optional:     true,
				Computed:     true,
			},
			"certificate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceEndpointControlFctemsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectEndpointControlFctems(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating EndpointControlFctems resource while getting object: %v", err)
	}

	o, err := c.CreateEndpointControlFctems(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating EndpointControlFctems resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("EndpointControlFctems")
	}

	return resourceEndpointControlFctemsRead(d, m)
}

func resourceEndpointControlFctemsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectEndpointControlFctems(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlFctems resource while getting object: %v", err)
	}

	o, err := c.UpdateEndpointControlFctems(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlFctems resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("EndpointControlFctems")
	}

	return resourceEndpointControlFctemsRead(d, m)
}

func resourceEndpointControlFctemsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteEndpointControlFctems(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting EndpointControlFctems resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEndpointControlFctemsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadEndpointControlFctems(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlFctems resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEndpointControlFctems(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlFctems resource from API: %v", err)
	}
	return nil
}

func flattenEndpointControlFctemsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsSerialNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsFortinetoneCloudAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsHttpsPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsAdminUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsAdminPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsPullSysinfo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsPullVulnerabilities(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsPullAvatars(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsPullTags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsCloudServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsCallTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectEndpointControlFctems(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenEndpointControlFctemsName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server", flattenEndpointControlFctemsServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("serial_number", flattenEndpointControlFctemsSerialNumber(o["serial-number"], d, "serial_number", sv)); err != nil {
		if !fortiAPIPatch(o["serial-number"]) {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	if err = d.Set("fortinetone_cloud_authentication", flattenEndpointControlFctemsFortinetoneCloudAuthentication(o["fortinetone-cloud-authentication"], d, "fortinetone_cloud_authentication", sv)); err != nil {
		if !fortiAPIPatch(o["fortinetone-cloud-authentication"]) {
			return fmt.Errorf("Error reading fortinetone_cloud_authentication: %v", err)
		}
	}

	if err = d.Set("https_port", flattenEndpointControlFctemsHttpsPort(o["https-port"], d, "https_port", sv)); err != nil {
		if !fortiAPIPatch(o["https-port"]) {
			return fmt.Errorf("Error reading https_port: %v", err)
		}
	}

	if err = d.Set("admin_username", flattenEndpointControlFctemsAdminUsername(o["admin-username"], d, "admin_username", sv)); err != nil {
		if !fortiAPIPatch(o["admin-username"]) {
			return fmt.Errorf("Error reading admin_username: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenEndpointControlFctemsSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("pull_sysinfo", flattenEndpointControlFctemsPullSysinfo(o["pull-sysinfo"], d, "pull_sysinfo", sv)); err != nil {
		if !fortiAPIPatch(o["pull-sysinfo"]) {
			return fmt.Errorf("Error reading pull_sysinfo: %v", err)
		}
	}

	if err = d.Set("pull_vulnerabilities", flattenEndpointControlFctemsPullVulnerabilities(o["pull-vulnerabilities"], d, "pull_vulnerabilities", sv)); err != nil {
		if !fortiAPIPatch(o["pull-vulnerabilities"]) {
			return fmt.Errorf("Error reading pull_vulnerabilities: %v", err)
		}
	}

	if err = d.Set("pull_avatars", flattenEndpointControlFctemsPullAvatars(o["pull-avatars"], d, "pull_avatars", sv)); err != nil {
		if !fortiAPIPatch(o["pull-avatars"]) {
			return fmt.Errorf("Error reading pull_avatars: %v", err)
		}
	}

	if err = d.Set("pull_tags", flattenEndpointControlFctemsPullTags(o["pull-tags"], d, "pull_tags", sv)); err != nil {
		if !fortiAPIPatch(o["pull-tags"]) {
			return fmt.Errorf("Error reading pull_tags: %v", err)
		}
	}

	if err = d.Set("cloud_server_type", flattenEndpointControlFctemsCloudServerType(o["cloud-server-type"], d, "cloud_server_type", sv)); err != nil {
		if !fortiAPIPatch(o["cloud-server-type"]) {
			return fmt.Errorf("Error reading cloud_server_type: %v", err)
		}
	}

	if err = d.Set("call_timeout", flattenEndpointControlFctemsCallTimeout(o["call-timeout"], d, "call_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["call-timeout"]) {
			return fmt.Errorf("Error reading call_timeout: %v", err)
		}
	}

	if err = d.Set("certificate", flattenEndpointControlFctemsCertificate(o["certificate"], d, "certificate", sv)); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	return nil
}

func flattenEndpointControlFctemsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandEndpointControlFctemsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsSerialNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsFortinetoneCloudAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsHttpsPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsAdminUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsAdminPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsPullSysinfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsPullVulnerabilities(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsPullAvatars(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsPullTags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsCloudServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsCallTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectEndpointControlFctems(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandEndpointControlFctemsName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {

		t, err := expandEndpointControlFctemsServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("serial_number"); ok {

		t, err := expandEndpointControlFctemsSerialNumber(d, v, "serial_number", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial-number"] = t
		}
	}

	if v, ok := d.GetOk("fortinetone_cloud_authentication"); ok {

		t, err := expandEndpointControlFctemsFortinetoneCloudAuthentication(d, v, "fortinetone_cloud_authentication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortinetone-cloud-authentication"] = t
		}
	}

	if v, ok := d.GetOk("https_port"); ok {

		t, err := expandEndpointControlFctemsHttpsPort(d, v, "https_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-port"] = t
		}
	}

	if v, ok := d.GetOk("admin_username"); ok {

		t, err := expandEndpointControlFctemsAdminUsername(d, v, "admin_username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-username"] = t
		}
	}

	if v, ok := d.GetOk("admin_password"); ok {

		t, err := expandEndpointControlFctemsAdminPassword(d, v, "admin_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-password"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {

		t, err := expandEndpointControlFctemsSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("pull_sysinfo"); ok {

		t, err := expandEndpointControlFctemsPullSysinfo(d, v, "pull_sysinfo", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pull-sysinfo"] = t
		}
	}

	if v, ok := d.GetOk("pull_vulnerabilities"); ok {

		t, err := expandEndpointControlFctemsPullVulnerabilities(d, v, "pull_vulnerabilities", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pull-vulnerabilities"] = t
		}
	}

	if v, ok := d.GetOk("pull_avatars"); ok {

		t, err := expandEndpointControlFctemsPullAvatars(d, v, "pull_avatars", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pull-avatars"] = t
		}
	}

	if v, ok := d.GetOk("pull_tags"); ok {

		t, err := expandEndpointControlFctemsPullTags(d, v, "pull_tags", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pull-tags"] = t
		}
	}

	if v, ok := d.GetOk("cloud_server_type"); ok {

		t, err := expandEndpointControlFctemsCloudServerType(d, v, "cloud_server_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cloud-server-type"] = t
		}
	}

	if v, ok := d.GetOk("call_timeout"); ok {

		t, err := expandEndpointControlFctemsCallTimeout(d, v, "call_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-timeout"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok {

		t, err := expandEndpointControlFctemsCertificate(d, v, "certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	return &obj, nil
}
