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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEndpointControlFctemsOverride() *schema.Resource {
	return &schema.Resource{
		Create: resourceEndpointControlFctemsOverrideCreate,
		Read:   resourceEndpointControlFctemsOverrideRead,
		Update: resourceEndpointControlFctemsOverrideUpdate,
		Delete: resourceEndpointControlFctemsOverrideDelete,

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
			"ems_id": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 7),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"dirty_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortinetone_cloud_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cloud_authentication_access_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 20),
				Optional:     true,
			},
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"https_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"serial_number": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 16),
				Optional:     true,
			},
			"tenant_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),
				Optional:     true,
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
			"pull_malware_hash": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cloud_server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capabilities": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"call_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 180),
				Optional:     true,
				Computed:     true,
			},
			"out_of_sync_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 3600),
				Optional:     true,
				Computed:     true,
			},
			"send_tags_to_all_vdoms": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"websocket_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preserve_ssl_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"trust_ca_cn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"verifying_ca": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
		},
	}
}

func resourceEndpointControlFctemsOverrideCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectEndpointControlFctemsOverride(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating EndpointControlFctemsOverride resource while getting object: %v", err)
	}

	o, err := c.CreateEndpointControlFctemsOverride(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating EndpointControlFctemsOverride resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("EndpointControlFctemsOverride")
	}

	return resourceEndpointControlFctemsOverrideRead(d, m)
}

func resourceEndpointControlFctemsOverrideUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectEndpointControlFctemsOverride(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlFctemsOverride resource while getting object: %v", err)
	}

	o, err := c.UpdateEndpointControlFctemsOverride(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlFctemsOverride resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("EndpointControlFctemsOverride")
	}

	return resourceEndpointControlFctemsOverrideRead(d, m)
}

func resourceEndpointControlFctemsOverrideDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteEndpointControlFctemsOverride(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting EndpointControlFctemsOverride resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEndpointControlFctemsOverrideRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadEndpointControlFctemsOverride(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlFctemsOverride resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEndpointControlFctemsOverride(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlFctemsOverride resource from API: %v", err)
	}
	return nil
}

func flattenEndpointControlFctemsOverrideEmsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenEndpointControlFctemsOverrideStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverrideName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverrideDirtyReason(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverrideFortinetoneCloudAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverrideCloudAuthenticationAccessKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverrideServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverrideHttpsPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenEndpointControlFctemsOverrideSerialNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverrideTenantId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverrideSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverridePullSysinfo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverridePullVulnerabilities(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverridePullAvatars(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverridePullTags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverridePullMalwareHash(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverrideCloudServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverrideCapabilities(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverrideCallTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenEndpointControlFctemsOverrideOutOfSyncThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenEndpointControlFctemsOverrideSendTagsToAllVdoms(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverrideWebsocketOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverridePreserveSslSession(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverrideInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverrideInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverrideTrustCaCn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlFctemsOverrideVerifyingCa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectEndpointControlFctemsOverride(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ems_id", flattenEndpointControlFctemsOverrideEmsId(o["ems-id"], d, "ems_id", sv)); err != nil {
		if !fortiAPIPatch(o["ems-id"]) {
			return fmt.Errorf("Error reading ems_id: %v", err)
		}
	}

	if err = d.Set("status", flattenEndpointControlFctemsOverrideStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("name", flattenEndpointControlFctemsOverrideName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("dirty_reason", flattenEndpointControlFctemsOverrideDirtyReason(o["dirty-reason"], d, "dirty_reason", sv)); err != nil {
		if !fortiAPIPatch(o["dirty-reason"]) {
			return fmt.Errorf("Error reading dirty_reason: %v", err)
		}
	}

	if err = d.Set("fortinetone_cloud_authentication", flattenEndpointControlFctemsOverrideFortinetoneCloudAuthentication(o["fortinetone-cloud-authentication"], d, "fortinetone_cloud_authentication", sv)); err != nil {
		if !fortiAPIPatch(o["fortinetone-cloud-authentication"]) {
			return fmt.Errorf("Error reading fortinetone_cloud_authentication: %v", err)
		}
	}

	if err = d.Set("cloud_authentication_access_key", flattenEndpointControlFctemsOverrideCloudAuthenticationAccessKey(o["cloud-authentication-access-key"], d, "cloud_authentication_access_key", sv)); err != nil {
		if !fortiAPIPatch(o["cloud-authentication-access-key"]) {
			return fmt.Errorf("Error reading cloud_authentication_access_key: %v", err)
		}
	}

	if err = d.Set("server", flattenEndpointControlFctemsOverrideServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("https_port", flattenEndpointControlFctemsOverrideHttpsPort(o["https-port"], d, "https_port", sv)); err != nil {
		if !fortiAPIPatch(o["https-port"]) {
			return fmt.Errorf("Error reading https_port: %v", err)
		}
	}

	if err = d.Set("serial_number", flattenEndpointControlFctemsOverrideSerialNumber(o["serial-number"], d, "serial_number", sv)); err != nil {
		if !fortiAPIPatch(o["serial-number"]) {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	if err = d.Set("tenant_id", flattenEndpointControlFctemsOverrideTenantId(o["tenant-id"], d, "tenant_id", sv)); err != nil {
		if !fortiAPIPatch(o["tenant-id"]) {
			return fmt.Errorf("Error reading tenant_id: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenEndpointControlFctemsOverrideSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("pull_sysinfo", flattenEndpointControlFctemsOverridePullSysinfo(o["pull-sysinfo"], d, "pull_sysinfo", sv)); err != nil {
		if !fortiAPIPatch(o["pull-sysinfo"]) {
			return fmt.Errorf("Error reading pull_sysinfo: %v", err)
		}
	}

	if err = d.Set("pull_vulnerabilities", flattenEndpointControlFctemsOverridePullVulnerabilities(o["pull-vulnerabilities"], d, "pull_vulnerabilities", sv)); err != nil {
		if !fortiAPIPatch(o["pull-vulnerabilities"]) {
			return fmt.Errorf("Error reading pull_vulnerabilities: %v", err)
		}
	}

	if err = d.Set("pull_avatars", flattenEndpointControlFctemsOverridePullAvatars(o["pull-avatars"], d, "pull_avatars", sv)); err != nil {
		if !fortiAPIPatch(o["pull-avatars"]) {
			return fmt.Errorf("Error reading pull_avatars: %v", err)
		}
	}

	if err = d.Set("pull_tags", flattenEndpointControlFctemsOverridePullTags(o["pull-tags"], d, "pull_tags", sv)); err != nil {
		if !fortiAPIPatch(o["pull-tags"]) {
			return fmt.Errorf("Error reading pull_tags: %v", err)
		}
	}

	if err = d.Set("pull_malware_hash", flattenEndpointControlFctemsOverridePullMalwareHash(o["pull-malware-hash"], d, "pull_malware_hash", sv)); err != nil {
		if !fortiAPIPatch(o["pull-malware-hash"]) {
			return fmt.Errorf("Error reading pull_malware_hash: %v", err)
		}
	}

	if err = d.Set("cloud_server_type", flattenEndpointControlFctemsOverrideCloudServerType(o["cloud-server-type"], d, "cloud_server_type", sv)); err != nil {
		if !fortiAPIPatch(o["cloud-server-type"]) {
			return fmt.Errorf("Error reading cloud_server_type: %v", err)
		}
	}

	if err = d.Set("capabilities", flattenEndpointControlFctemsOverrideCapabilities(o["capabilities"], d, "capabilities", sv)); err != nil {
		if !fortiAPIPatch(o["capabilities"]) {
			return fmt.Errorf("Error reading capabilities: %v", err)
		}
	}

	if err = d.Set("call_timeout", flattenEndpointControlFctemsOverrideCallTimeout(o["call-timeout"], d, "call_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["call-timeout"]) {
			return fmt.Errorf("Error reading call_timeout: %v", err)
		}
	}

	if err = d.Set("out_of_sync_threshold", flattenEndpointControlFctemsOverrideOutOfSyncThreshold(o["out-of-sync-threshold"], d, "out_of_sync_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["out-of-sync-threshold"]) {
			return fmt.Errorf("Error reading out_of_sync_threshold: %v", err)
		}
	}

	if err = d.Set("send_tags_to_all_vdoms", flattenEndpointControlFctemsOverrideSendTagsToAllVdoms(o["send-tags-to-all-vdoms"], d, "send_tags_to_all_vdoms", sv)); err != nil {
		if !fortiAPIPatch(o["send-tags-to-all-vdoms"]) {
			return fmt.Errorf("Error reading send_tags_to_all_vdoms: %v", err)
		}
	}

	if err = d.Set("websocket_override", flattenEndpointControlFctemsOverrideWebsocketOverride(o["websocket-override"], d, "websocket_override", sv)); err != nil {
		if !fortiAPIPatch(o["websocket-override"]) {
			return fmt.Errorf("Error reading websocket_override: %v", err)
		}
	}

	if err = d.Set("preserve_ssl_session", flattenEndpointControlFctemsOverridePreserveSslSession(o["preserve-ssl-session"], d, "preserve_ssl_session", sv)); err != nil {
		if !fortiAPIPatch(o["preserve-ssl-session"]) {
			return fmt.Errorf("Error reading preserve_ssl_session: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenEndpointControlFctemsOverrideInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenEndpointControlFctemsOverrideInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("trust_ca_cn", flattenEndpointControlFctemsOverrideTrustCaCn(o["trust-ca-cn"], d, "trust_ca_cn", sv)); err != nil {
		if !fortiAPIPatch(o["trust-ca-cn"]) {
			return fmt.Errorf("Error reading trust_ca_cn: %v", err)
		}
	}

	if err = d.Set("verifying_ca", flattenEndpointControlFctemsOverrideVerifyingCa(o["verifying-ca"], d, "verifying_ca", sv)); err != nil {
		if !fortiAPIPatch(o["verifying-ca"]) {
			return fmt.Errorf("Error reading verifying_ca: %v", err)
		}
	}

	return nil
}

func flattenEndpointControlFctemsOverrideFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandEndpointControlFctemsOverrideEmsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideDirtyReason(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideFortinetoneCloudAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideCloudAuthenticationAccessKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideHttpsPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideSerialNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideTenantId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverridePullSysinfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverridePullVulnerabilities(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverridePullAvatars(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverridePullTags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverridePullMalwareHash(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideCloudServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideCapabilities(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideCallTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideOutOfSyncThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideSendTagsToAllVdoms(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideWebsocketOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverridePreserveSslSession(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideTrustCaCn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOverrideVerifyingCa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectEndpointControlFctemsOverride(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ems_id"); ok {
		t, err := expandEndpointControlFctemsOverrideEmsId(d, v, "ems_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ems-id"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandEndpointControlFctemsOverrideStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandEndpointControlFctemsOverrideName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	} else if d.HasChange("name") {
		obj["name"] = nil
	}

	if v, ok := d.GetOk("dirty_reason"); ok {
		t, err := expandEndpointControlFctemsOverrideDirtyReason(d, v, "dirty_reason", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dirty-reason"] = t
		}
	}

	if v, ok := d.GetOk("fortinetone_cloud_authentication"); ok {
		t, err := expandEndpointControlFctemsOverrideFortinetoneCloudAuthentication(d, v, "fortinetone_cloud_authentication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortinetone-cloud-authentication"] = t
		}
	}

	if v, ok := d.GetOk("cloud_authentication_access_key"); ok {
		t, err := expandEndpointControlFctemsOverrideCloudAuthenticationAccessKey(d, v, "cloud_authentication_access_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cloud-authentication-access-key"] = t
		}
	} else if d.HasChange("cloud_authentication_access_key") {
		obj["cloud-authentication-access-key"] = nil
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandEndpointControlFctemsOverrideServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	} else if d.HasChange("server") {
		obj["server"] = nil
	}

	if v, ok := d.GetOk("https_port"); ok {
		t, err := expandEndpointControlFctemsOverrideHttpsPort(d, v, "https_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-port"] = t
		}
	}

	if v, ok := d.GetOk("serial_number"); ok {
		t, err := expandEndpointControlFctemsOverrideSerialNumber(d, v, "serial_number", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial-number"] = t
		}
	} else if d.HasChange("serial_number") {
		obj["serial-number"] = nil
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		t, err := expandEndpointControlFctemsOverrideTenantId(d, v, "tenant_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant-id"] = t
		}
	} else if d.HasChange("tenant_id") {
		obj["tenant-id"] = nil
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandEndpointControlFctemsOverrideSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("pull_sysinfo"); ok {
		t, err := expandEndpointControlFctemsOverridePullSysinfo(d, v, "pull_sysinfo", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pull-sysinfo"] = t
		}
	}

	if v, ok := d.GetOk("pull_vulnerabilities"); ok {
		t, err := expandEndpointControlFctemsOverridePullVulnerabilities(d, v, "pull_vulnerabilities", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pull-vulnerabilities"] = t
		}
	}

	if v, ok := d.GetOk("pull_avatars"); ok {
		t, err := expandEndpointControlFctemsOverridePullAvatars(d, v, "pull_avatars", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pull-avatars"] = t
		}
	}

	if v, ok := d.GetOk("pull_tags"); ok {
		t, err := expandEndpointControlFctemsOverridePullTags(d, v, "pull_tags", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pull-tags"] = t
		}
	}

	if v, ok := d.GetOk("pull_malware_hash"); ok {
		t, err := expandEndpointControlFctemsOverridePullMalwareHash(d, v, "pull_malware_hash", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pull-malware-hash"] = t
		}
	}

	if v, ok := d.GetOk("cloud_server_type"); ok {
		t, err := expandEndpointControlFctemsOverrideCloudServerType(d, v, "cloud_server_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cloud-server-type"] = t
		}
	}

	if v, ok := d.GetOk("capabilities"); ok {
		t, err := expandEndpointControlFctemsOverrideCapabilities(d, v, "capabilities", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capabilities"] = t
		}
	} else if d.HasChange("capabilities") {
		obj["capabilities"] = nil
	}

	if v, ok := d.GetOk("call_timeout"); ok {
		t, err := expandEndpointControlFctemsOverrideCallTimeout(d, v, "call_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-timeout"] = t
		}
	}

	if v, ok := d.GetOk("out_of_sync_threshold"); ok {
		t, err := expandEndpointControlFctemsOverrideOutOfSyncThreshold(d, v, "out_of_sync_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["out-of-sync-threshold"] = t
		}
	}

	if v, ok := d.GetOk("send_tags_to_all_vdoms"); ok {
		t, err := expandEndpointControlFctemsOverrideSendTagsToAllVdoms(d, v, "send_tags_to_all_vdoms", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-tags-to-all-vdoms"] = t
		}
	}

	if v, ok := d.GetOk("websocket_override"); ok {
		t, err := expandEndpointControlFctemsOverrideWebsocketOverride(d, v, "websocket_override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["websocket-override"] = t
		}
	}

	if v, ok := d.GetOk("preserve_ssl_session"); ok {
		t, err := expandEndpointControlFctemsOverridePreserveSslSession(d, v, "preserve_ssl_session", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preserve-ssl-session"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		t, err := expandEndpointControlFctemsOverrideInterfaceSelectMethod(d, v, "interface_select_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandEndpointControlFctemsOverrideInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOk("trust_ca_cn"); ok {
		t, err := expandEndpointControlFctemsOverrideTrustCaCn(d, v, "trust_ca_cn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ca-cn"] = t
		}
	}

	if v, ok := d.GetOk("verifying_ca"); ok {
		t, err := expandEndpointControlFctemsOverrideVerifyingCa(d, v, "verifying_ca", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verifying-ca"] = t
		}
	} else if d.HasChange("verifying_ca") {
		obj["verifying-ca"] = nil
	}

	return &obj, nil
}
