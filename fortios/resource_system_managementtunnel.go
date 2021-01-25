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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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

	obj, err := getObjectSystemManagementTunnel(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemManagementTunnel resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemManagementTunnel(obj, mkey)
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

	err := c.DeleteSystemManagementTunnel(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemManagementTunnel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemManagementTunnelRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemManagementTunnel(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemManagementTunnel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemManagementTunnel(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemManagementTunnel resource from API: %v", err)
	}
	return nil
}

func flattenSystemManagementTunnelStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemManagementTunnelAllowConfigRestore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemManagementTunnelAllowPushConfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemManagementTunnelAllowPushFirmware(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemManagementTunnelAllowCollectStatistics(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemManagementTunnelAuthorizedManagerOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemManagementTunnelSerialNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemManagementTunnel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", flattenSystemManagementTunnelStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("allow_config_restore", flattenSystemManagementTunnelAllowConfigRestore(o["allow-config-restore"], d, "allow_config_restore")); err != nil {
		if !fortiAPIPatch(o["allow-config-restore"]) {
			return fmt.Errorf("Error reading allow_config_restore: %v", err)
		}
	}

	if err = d.Set("allow_push_configuration", flattenSystemManagementTunnelAllowPushConfiguration(o["allow-push-configuration"], d, "allow_push_configuration")); err != nil {
		if !fortiAPIPatch(o["allow-push-configuration"]) {
			return fmt.Errorf("Error reading allow_push_configuration: %v", err)
		}
	}

	if err = d.Set("allow_push_firmware", flattenSystemManagementTunnelAllowPushFirmware(o["allow-push-firmware"], d, "allow_push_firmware")); err != nil {
		if !fortiAPIPatch(o["allow-push-firmware"]) {
			return fmt.Errorf("Error reading allow_push_firmware: %v", err)
		}
	}

	if err = d.Set("allow_collect_statistics", flattenSystemManagementTunnelAllowCollectStatistics(o["allow-collect-statistics"], d, "allow_collect_statistics")); err != nil {
		if !fortiAPIPatch(o["allow-collect-statistics"]) {
			return fmt.Errorf("Error reading allow_collect_statistics: %v", err)
		}
	}

	if err = d.Set("authorized_manager_only", flattenSystemManagementTunnelAuthorizedManagerOnly(o["authorized-manager-only"], d, "authorized_manager_only")); err != nil {
		if !fortiAPIPatch(o["authorized-manager-only"]) {
			return fmt.Errorf("Error reading authorized_manager_only: %v", err)
		}
	}

	if err = d.Set("serial_number", flattenSystemManagementTunnelSerialNumber(o["serial-number"], d, "serial_number")); err != nil {
		if !fortiAPIPatch(o["serial-number"]) {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	return nil
}

func flattenSystemManagementTunnelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemManagementTunnelStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemManagementTunnelAllowConfigRestore(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemManagementTunnelAllowPushConfiguration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemManagementTunnelAllowPushFirmware(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemManagementTunnelAllowCollectStatistics(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemManagementTunnelAuthorizedManagerOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemManagementTunnelSerialNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemManagementTunnel(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemManagementTunnelStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("allow_config_restore"); ok {
		t, err := expandSystemManagementTunnelAllowConfigRestore(d, v, "allow_config_restore")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-config-restore"] = t
		}
	}

	if v, ok := d.GetOk("allow_push_configuration"); ok {
		t, err := expandSystemManagementTunnelAllowPushConfiguration(d, v, "allow_push_configuration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-push-configuration"] = t
		}
	}

	if v, ok := d.GetOk("allow_push_firmware"); ok {
		t, err := expandSystemManagementTunnelAllowPushFirmware(d, v, "allow_push_firmware")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-push-firmware"] = t
		}
	}

	if v, ok := d.GetOk("allow_collect_statistics"); ok {
		t, err := expandSystemManagementTunnelAllowCollectStatistics(d, v, "allow_collect_statistics")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-collect-statistics"] = t
		}
	}

	if v, ok := d.GetOk("authorized_manager_only"); ok {
		t, err := expandSystemManagementTunnelAuthorizedManagerOnly(d, v, "authorized_manager_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorized-manager-only"] = t
		}
	}

	if v, ok := d.GetOk("serial_number"); ok {
		t, err := expandSystemManagementTunnelSerialNumber(d, v, "serial_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial-number"] = t
		}
	}

	return &obj, nil
}
