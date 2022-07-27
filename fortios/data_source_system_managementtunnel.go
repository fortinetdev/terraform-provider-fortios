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

func dataSourceSystemManagementTunnel() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemManagementTunnelRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_config_restore": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_push_configuration": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_push_firmware": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_collect_statistics": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"authorized_manager_only": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemManagementTunnelRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "SystemManagementTunnel"

	o, err := c.ReadSystemManagementTunnel(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemManagementTunnel: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemManagementTunnel(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemManagementTunnel from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemManagementTunnelStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemManagementTunnelAllowConfigRestore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemManagementTunnelAllowPushConfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemManagementTunnelAllowPushFirmware(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemManagementTunnelAllowCollectStatistics(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemManagementTunnelAuthorizedManagerOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemManagementTunnelSerialNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemManagementTunnel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", dataSourceFlattenSystemManagementTunnelStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("allow_config_restore", dataSourceFlattenSystemManagementTunnelAllowConfigRestore(o["allow-config-restore"], d, "allow_config_restore")); err != nil {
		if !fortiAPIPatch(o["allow-config-restore"]) {
			return fmt.Errorf("Error reading allow_config_restore: %v", err)
		}
	}

	if err = d.Set("allow_push_configuration", dataSourceFlattenSystemManagementTunnelAllowPushConfiguration(o["allow-push-configuration"], d, "allow_push_configuration")); err != nil {
		if !fortiAPIPatch(o["allow-push-configuration"]) {
			return fmt.Errorf("Error reading allow_push_configuration: %v", err)
		}
	}

	if err = d.Set("allow_push_firmware", dataSourceFlattenSystemManagementTunnelAllowPushFirmware(o["allow-push-firmware"], d, "allow_push_firmware")); err != nil {
		if !fortiAPIPatch(o["allow-push-firmware"]) {
			return fmt.Errorf("Error reading allow_push_firmware: %v", err)
		}
	}

	if err = d.Set("allow_collect_statistics", dataSourceFlattenSystemManagementTunnelAllowCollectStatistics(o["allow-collect-statistics"], d, "allow_collect_statistics")); err != nil {
		if !fortiAPIPatch(o["allow-collect-statistics"]) {
			return fmt.Errorf("Error reading allow_collect_statistics: %v", err)
		}
	}

	if err = d.Set("authorized_manager_only", dataSourceFlattenSystemManagementTunnelAuthorizedManagerOnly(o["authorized-manager-only"], d, "authorized_manager_only")); err != nil {
		if !fortiAPIPatch(o["authorized-manager-only"]) {
			return fmt.Errorf("Error reading authorized_manager_only: %v", err)
		}
	}

	if err = d.Set("serial_number", dataSourceFlattenSystemManagementTunnelSerialNumber(o["serial-number"], d, "serial_number")); err != nil {
		if !fortiAPIPatch(o["serial-number"]) {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemManagementTunnelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
