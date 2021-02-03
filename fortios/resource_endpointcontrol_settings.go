// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure endpoint control settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceEndpointControlSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceEndpointControlSettingsUpdate,
		Read:   resourceEndpointControlSettingsRead,
		Update: resourceEndpointControlSettingsUpdate,
		Delete: resourceEndpointControlSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"forticlient_reg_key_enforce": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_reg_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"forticlient_reg_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 180),
				Optional:     true,
				Computed:     true,
			},
			"download_custom_link": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"download_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_offline_grace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_offline_grace_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 600),
				Optional:     true,
				Computed:     true,
			},
			"forticlient_keepalive_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(20, 300),
				Optional:     true,
				Computed:     true,
			},
			"forticlient_sys_update_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 1440),
				Optional:     true,
				Computed:     true,
			},
			"forticlient_avdb_update_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 24),
				Optional:     true,
				Computed:     true,
			},
			"forticlient_warning_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 24),
				Optional:     true,
				Computed:     true,
			},
			"forticlient_user_avatar": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_disconnect_unsupported_client": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_dereg_unsupported_client": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_ems_rest_api_call_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(500, 30000),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceEndpointControlSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectEndpointControlSettings(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateEndpointControlSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("EndpointControlSettings")
	}

	return resourceEndpointControlSettingsRead(d, m)
}

func resourceEndpointControlSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteEndpointControlSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting EndpointControlSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEndpointControlSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadEndpointControlSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEndpointControlSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlSettings resource from API: %v", err)
	}
	return nil
}

func flattenEndpointControlSettingsForticlientRegKeyEnforce(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlSettingsForticlientRegKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlSettingsForticlientRegTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlSettingsDownloadCustomLink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlSettingsDownloadLocation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlSettingsForticlientOfflineGrace(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlSettingsForticlientOfflineGraceInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlSettingsForticlientKeepaliveInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlSettingsForticlientSysUpdateInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlSettingsForticlientAvdbUpdateInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlSettingsForticlientWarningInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlSettingsForticlientUserAvatar(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlSettingsForticlientDisconnectUnsupportedClient(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlSettingsForticlientDeregUnsupportedClient(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlSettingsForticlientEmsRestApiCallTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectEndpointControlSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("forticlient_reg_key_enforce", flattenEndpointControlSettingsForticlientRegKeyEnforce(o["forticlient-reg-key-enforce"], d, "forticlient_reg_key_enforce", sv)); err != nil {
		if !fortiAPIPatch(o["forticlient-reg-key-enforce"]) {
			return fmt.Errorf("Error reading forticlient_reg_key_enforce: %v", err)
		}
	}

	if err = d.Set("forticlient_reg_timeout", flattenEndpointControlSettingsForticlientRegTimeout(o["forticlient-reg-timeout"], d, "forticlient_reg_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["forticlient-reg-timeout"]) {
			return fmt.Errorf("Error reading forticlient_reg_timeout: %v", err)
		}
	}

	if err = d.Set("download_custom_link", flattenEndpointControlSettingsDownloadCustomLink(o["download-custom-link"], d, "download_custom_link", sv)); err != nil {
		if !fortiAPIPatch(o["download-custom-link"]) {
			return fmt.Errorf("Error reading download_custom_link: %v", err)
		}
	}

	if err = d.Set("download_location", flattenEndpointControlSettingsDownloadLocation(o["download-location"], d, "download_location", sv)); err != nil {
		if !fortiAPIPatch(o["download-location"]) {
			return fmt.Errorf("Error reading download_location: %v", err)
		}
	}

	if err = d.Set("forticlient_offline_grace", flattenEndpointControlSettingsForticlientOfflineGrace(o["forticlient-offline-grace"], d, "forticlient_offline_grace", sv)); err != nil {
		if !fortiAPIPatch(o["forticlient-offline-grace"]) {
			return fmt.Errorf("Error reading forticlient_offline_grace: %v", err)
		}
	}

	if err = d.Set("forticlient_offline_grace_interval", flattenEndpointControlSettingsForticlientOfflineGraceInterval(o["forticlient-offline-grace-interval"], d, "forticlient_offline_grace_interval", sv)); err != nil {
		if !fortiAPIPatch(o["forticlient-offline-grace-interval"]) {
			return fmt.Errorf("Error reading forticlient_offline_grace_interval: %v", err)
		}
	}

	if err = d.Set("forticlient_keepalive_interval", flattenEndpointControlSettingsForticlientKeepaliveInterval(o["forticlient-keepalive-interval"], d, "forticlient_keepalive_interval", sv)); err != nil {
		if !fortiAPIPatch(o["forticlient-keepalive-interval"]) {
			return fmt.Errorf("Error reading forticlient_keepalive_interval: %v", err)
		}
	}

	if err = d.Set("forticlient_sys_update_interval", flattenEndpointControlSettingsForticlientSysUpdateInterval(o["forticlient-sys-update-interval"], d, "forticlient_sys_update_interval", sv)); err != nil {
		if !fortiAPIPatch(o["forticlient-sys-update-interval"]) {
			return fmt.Errorf("Error reading forticlient_sys_update_interval: %v", err)
		}
	}

	if err = d.Set("forticlient_avdb_update_interval", flattenEndpointControlSettingsForticlientAvdbUpdateInterval(o["forticlient-avdb-update-interval"], d, "forticlient_avdb_update_interval", sv)); err != nil {
		if !fortiAPIPatch(o["forticlient-avdb-update-interval"]) {
			return fmt.Errorf("Error reading forticlient_avdb_update_interval: %v", err)
		}
	}

	if err = d.Set("forticlient_warning_interval", flattenEndpointControlSettingsForticlientWarningInterval(o["forticlient-warning-interval"], d, "forticlient_warning_interval", sv)); err != nil {
		if !fortiAPIPatch(o["forticlient-warning-interval"]) {
			return fmt.Errorf("Error reading forticlient_warning_interval: %v", err)
		}
	}

	if err = d.Set("forticlient_user_avatar", flattenEndpointControlSettingsForticlientUserAvatar(o["forticlient-user-avatar"], d, "forticlient_user_avatar", sv)); err != nil {
		if !fortiAPIPatch(o["forticlient-user-avatar"]) {
			return fmt.Errorf("Error reading forticlient_user_avatar: %v", err)
		}
	}

	if err = d.Set("forticlient_disconnect_unsupported_client", flattenEndpointControlSettingsForticlientDisconnectUnsupportedClient(o["forticlient-disconnect-unsupported-client"], d, "forticlient_disconnect_unsupported_client", sv)); err != nil {
		if !fortiAPIPatch(o["forticlient-disconnect-unsupported-client"]) {
			return fmt.Errorf("Error reading forticlient_disconnect_unsupported_client: %v", err)
		}
	}

	if err = d.Set("forticlient_dereg_unsupported_client", flattenEndpointControlSettingsForticlientDeregUnsupportedClient(o["forticlient-dereg-unsupported-client"], d, "forticlient_dereg_unsupported_client", sv)); err != nil {
		if !fortiAPIPatch(o["forticlient-dereg-unsupported-client"]) {
			return fmt.Errorf("Error reading forticlient_dereg_unsupported_client: %v", err)
		}
	}

	if err = d.Set("forticlient_ems_rest_api_call_timeout", flattenEndpointControlSettingsForticlientEmsRestApiCallTimeout(o["forticlient-ems-rest-api-call-timeout"], d, "forticlient_ems_rest_api_call_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["forticlient-ems-rest-api-call-timeout"]) {
			return fmt.Errorf("Error reading forticlient_ems_rest_api_call_timeout: %v", err)
		}
	}

	return nil
}

func flattenEndpointControlSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandEndpointControlSettingsForticlientRegKeyEnforce(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlSettingsForticlientRegKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlSettingsForticlientRegTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlSettingsDownloadCustomLink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlSettingsDownloadLocation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlSettingsForticlientOfflineGrace(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlSettingsForticlientOfflineGraceInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlSettingsForticlientKeepaliveInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlSettingsForticlientSysUpdateInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlSettingsForticlientAvdbUpdateInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlSettingsForticlientWarningInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlSettingsForticlientUserAvatar(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlSettingsForticlientDisconnectUnsupportedClient(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlSettingsForticlientDeregUnsupportedClient(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlSettingsForticlientEmsRestApiCallTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectEndpointControlSettings(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("forticlient_reg_key_enforce"); ok {

		t, err := expandEndpointControlSettingsForticlientRegKeyEnforce(d, v, "forticlient_reg_key_enforce", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-reg-key-enforce"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_reg_key"); ok {

		t, err := expandEndpointControlSettingsForticlientRegKey(d, v, "forticlient_reg_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-reg-key"] = t
		}
	}

	if v, ok := d.GetOkExists("forticlient_reg_timeout"); ok {

		t, err := expandEndpointControlSettingsForticlientRegTimeout(d, v, "forticlient_reg_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-reg-timeout"] = t
		}
	}

	if v, ok := d.GetOk("download_custom_link"); ok {

		t, err := expandEndpointControlSettingsDownloadCustomLink(d, v, "download_custom_link", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["download-custom-link"] = t
		}
	}

	if v, ok := d.GetOk("download_location"); ok {

		t, err := expandEndpointControlSettingsDownloadLocation(d, v, "download_location", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["download-location"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_offline_grace"); ok {

		t, err := expandEndpointControlSettingsForticlientOfflineGrace(d, v, "forticlient_offline_grace", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-offline-grace"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_offline_grace_interval"); ok {

		t, err := expandEndpointControlSettingsForticlientOfflineGraceInterval(d, v, "forticlient_offline_grace_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-offline-grace-interval"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_keepalive_interval"); ok {

		t, err := expandEndpointControlSettingsForticlientKeepaliveInterval(d, v, "forticlient_keepalive_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-keepalive-interval"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_sys_update_interval"); ok {

		t, err := expandEndpointControlSettingsForticlientSysUpdateInterval(d, v, "forticlient_sys_update_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-sys-update-interval"] = t
		}
	}

	if v, ok := d.GetOkExists("forticlient_avdb_update_interval"); ok {

		t, err := expandEndpointControlSettingsForticlientAvdbUpdateInterval(d, v, "forticlient_avdb_update_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-avdb-update-interval"] = t
		}
	}

	if v, ok := d.GetOkExists("forticlient_warning_interval"); ok {

		t, err := expandEndpointControlSettingsForticlientWarningInterval(d, v, "forticlient_warning_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-warning-interval"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_user_avatar"); ok {

		t, err := expandEndpointControlSettingsForticlientUserAvatar(d, v, "forticlient_user_avatar", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-user-avatar"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_disconnect_unsupported_client"); ok {

		t, err := expandEndpointControlSettingsForticlientDisconnectUnsupportedClient(d, v, "forticlient_disconnect_unsupported_client", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-disconnect-unsupported-client"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_dereg_unsupported_client"); ok {

		t, err := expandEndpointControlSettingsForticlientDeregUnsupportedClient(d, v, "forticlient_dereg_unsupported_client", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-dereg-unsupported-client"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_ems_rest_api_call_timeout"); ok {

		t, err := expandEndpointControlSettingsForticlientEmsRestApiCallTimeout(d, v, "forticlient_ems_rest_api_call_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-ems-rest-api-call-timeout"] = t
		}
	}

	return &obj, nil
}
