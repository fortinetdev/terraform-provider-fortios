// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure dedicated management.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDedicatedMgmt() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDedicatedMgmtUpdate,
		Read:   resourceSystemDedicatedMgmtRead,
		Update: resourceSystemDedicatedMgmtUpdate,
		Delete: resourceSystemDedicatedMgmtDelete,

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
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"default_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemDedicatedMgmtUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDedicatedMgmt(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemDedicatedMgmt resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemDedicatedMgmt(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemDedicatedMgmt resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemDedicatedMgmt")
	}

	return resourceSystemDedicatedMgmtRead(d, m)
}

func resourceSystemDedicatedMgmtDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemDedicatedMgmt(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemDedicatedMgmt resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDedicatedMgmt(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemDedicatedMgmt resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDedicatedMgmtRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemDedicatedMgmt(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemDedicatedMgmt resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDedicatedMgmt(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemDedicatedMgmt resource from API: %v", err)
	}
	return nil
}

func flattenSystemDedicatedMgmtStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDedicatedMgmtInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDedicatedMgmtDefaultGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDedicatedMgmtDhcpServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDedicatedMgmtDhcpNetmask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDedicatedMgmtDhcpStartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDedicatedMgmtDhcpEndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemDedicatedMgmt(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemDedicatedMgmtStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemDedicatedMgmtInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("default_gateway", flattenSystemDedicatedMgmtDefaultGateway(o["default-gateway"], d, "default_gateway", sv)); err != nil {
		if !fortiAPIPatch(o["default-gateway"]) {
			return fmt.Errorf("Error reading default_gateway: %v", err)
		}
	}

	if err = d.Set("dhcp_server", flattenSystemDedicatedMgmtDhcpServer(o["dhcp-server"], d, "dhcp_server", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-server"]) {
			return fmt.Errorf("Error reading dhcp_server: %v", err)
		}
	}

	if err = d.Set("dhcp_netmask", flattenSystemDedicatedMgmtDhcpNetmask(o["dhcp-netmask"], d, "dhcp_netmask", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-netmask"]) {
			return fmt.Errorf("Error reading dhcp_netmask: %v", err)
		}
	}

	if err = d.Set("dhcp_start_ip", flattenSystemDedicatedMgmtDhcpStartIp(o["dhcp-start-ip"], d, "dhcp_start_ip", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-start-ip"]) {
			return fmt.Errorf("Error reading dhcp_start_ip: %v", err)
		}
	}

	if err = d.Set("dhcp_end_ip", flattenSystemDedicatedMgmtDhcpEndIp(o["dhcp-end-ip"], d, "dhcp_end_ip", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-end-ip"]) {
			return fmt.Errorf("Error reading dhcp_end_ip: %v", err)
		}
	}

	return nil
}

func flattenSystemDedicatedMgmtFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemDedicatedMgmtStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDedicatedMgmtInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDedicatedMgmtDefaultGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDedicatedMgmtDhcpServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDedicatedMgmtDhcpNetmask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDedicatedMgmtDhcpStartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDedicatedMgmtDhcpEndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDedicatedMgmt(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSystemDedicatedMgmtStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		if setArgNil {
			obj["interface"] = nil
		} else {
			t, err := expandSystemDedicatedMgmtInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOk("default_gateway"); ok {
		if setArgNil {
			obj["default-gateway"] = nil
		} else {
			t, err := expandSystemDedicatedMgmtDefaultGateway(d, v, "default_gateway", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-gateway"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp_server"); ok {
		if setArgNil {
			obj["dhcp-server"] = nil
		} else {
			t, err := expandSystemDedicatedMgmtDhcpServer(d, v, "dhcp_server", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp-server"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp_netmask"); ok {
		if setArgNil {
			obj["dhcp-netmask"] = nil
		} else {
			t, err := expandSystemDedicatedMgmtDhcpNetmask(d, v, "dhcp_netmask", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp-netmask"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp_start_ip"); ok {
		if setArgNil {
			obj["dhcp-start-ip"] = nil
		} else {
			t, err := expandSystemDedicatedMgmtDhcpStartIp(d, v, "dhcp_start_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp-start-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp_end_ip"); ok {
		if setArgNil {
			obj["dhcp-end-ip"] = nil
		} else {
			t, err := expandSystemDedicatedMgmtDhcpEndIp(d, v, "dhcp_end_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp-end-ip"] = t
			}
		}
	}

	return &obj, nil
}
