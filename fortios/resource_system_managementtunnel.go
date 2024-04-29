// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Management tunnel configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemManagementTunnel() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemManagementTunnelUpdate,
		Read:   resourceSystemManagementTunnelRead,
		Update: resourceSystemManagementTunnelUpdate,
		Delete: resourceSystemManagementTunnelDelete,

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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_config_restore": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_push_configuration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_push_firmware": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_collect_statistics": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorized_manager_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemManagementTunnelUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemManagementTunnel(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemManagementTunnel resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemManagementTunnel(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemManagementTunnel resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemManagementTunnel")
	}

	return resourceSystemManagementTunnelRead(d, m)
}

func resourceSystemManagementTunnelDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemManagementTunnel(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemManagementTunnel resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemManagementTunnel(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemManagementTunnel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemManagementTunnelRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemManagementTunnel(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemManagementTunnel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemManagementTunnel(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemManagementTunnel resource from API: %v", err)
	}
	return nil
}

func flattenSystemManagementTunnelStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemManagementTunnelAllowConfigRestore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemManagementTunnelAllowPushConfiguration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemManagementTunnelAllowPushFirmware(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemManagementTunnelAllowCollectStatistics(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemManagementTunnelAuthorizedManagerOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemManagementTunnelSerialNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemManagementTunnel(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemManagementTunnelStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("allow_config_restore", flattenSystemManagementTunnelAllowConfigRestore(o["allow-config-restore"], d, "allow_config_restore", sv)); err != nil {
		if !fortiAPIPatch(o["allow-config-restore"]) {
			return fmt.Errorf("Error reading allow_config_restore: %v", err)
		}
	}

	if err = d.Set("allow_push_configuration", flattenSystemManagementTunnelAllowPushConfiguration(o["allow-push-configuration"], d, "allow_push_configuration", sv)); err != nil {
		if !fortiAPIPatch(o["allow-push-configuration"]) {
			return fmt.Errorf("Error reading allow_push_configuration: %v", err)
		}
	}

	if err = d.Set("allow_push_firmware", flattenSystemManagementTunnelAllowPushFirmware(o["allow-push-firmware"], d, "allow_push_firmware", sv)); err != nil {
		if !fortiAPIPatch(o["allow-push-firmware"]) {
			return fmt.Errorf("Error reading allow_push_firmware: %v", err)
		}
	}

	if err = d.Set("allow_collect_statistics", flattenSystemManagementTunnelAllowCollectStatistics(o["allow-collect-statistics"], d, "allow_collect_statistics", sv)); err != nil {
		if !fortiAPIPatch(o["allow-collect-statistics"]) {
			return fmt.Errorf("Error reading allow_collect_statistics: %v", err)
		}
	}

	if err = d.Set("authorized_manager_only", flattenSystemManagementTunnelAuthorizedManagerOnly(o["authorized-manager-only"], d, "authorized_manager_only", sv)); err != nil {
		if !fortiAPIPatch(o["authorized-manager-only"]) {
			return fmt.Errorf("Error reading authorized_manager_only: %v", err)
		}
	}

	if err = d.Set("serial_number", flattenSystemManagementTunnelSerialNumber(o["serial-number"], d, "serial_number", sv)); err != nil {
		if !fortiAPIPatch(o["serial-number"]) {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	return nil
}

func flattenSystemManagementTunnelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemManagementTunnelStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemManagementTunnelAllowConfigRestore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemManagementTunnelAllowPushConfiguration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemManagementTunnelAllowPushFirmware(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemManagementTunnelAllowCollectStatistics(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemManagementTunnelAuthorizedManagerOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemManagementTunnelSerialNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemManagementTunnel(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSystemManagementTunnelStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("allow_config_restore"); ok {
		if setArgNil {
			obj["allow-config-restore"] = nil
		} else {
			t, err := expandSystemManagementTunnelAllowConfigRestore(d, v, "allow_config_restore", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["allow-config-restore"] = t
			}
		}
	}

	if v, ok := d.GetOk("allow_push_configuration"); ok {
		if setArgNil {
			obj["allow-push-configuration"] = nil
		} else {
			t, err := expandSystemManagementTunnelAllowPushConfiguration(d, v, "allow_push_configuration", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["allow-push-configuration"] = t
			}
		}
	}

	if v, ok := d.GetOk("allow_push_firmware"); ok {
		if setArgNil {
			obj["allow-push-firmware"] = nil
		} else {
			t, err := expandSystemManagementTunnelAllowPushFirmware(d, v, "allow_push_firmware", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["allow-push-firmware"] = t
			}
		}
	}

	if v, ok := d.GetOk("allow_collect_statistics"); ok {
		if setArgNil {
			obj["allow-collect-statistics"] = nil
		} else {
			t, err := expandSystemManagementTunnelAllowCollectStatistics(d, v, "allow_collect_statistics", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["allow-collect-statistics"] = t
			}
		}
	}

	if v, ok := d.GetOk("authorized_manager_only"); ok {
		if setArgNil {
			obj["authorized-manager-only"] = nil
		} else {
			t, err := expandSystemManagementTunnelAuthorizedManagerOnly(d, v, "authorized_manager_only", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["authorized-manager-only"] = t
			}
		}
	}

	if v, ok := d.GetOk("serial_number"); ok {
		if setArgNil {
			obj["serial-number"] = nil
		} else {
			t, err := expandSystemManagementTunnelSerialNumber(d, v, "serial_number", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["serial-number"] = t
			}
		}
	}

	return &obj, nil
}
