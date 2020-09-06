// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure dedicated management.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
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

	obj, err := getObjectSystemDedicatedMgmt(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDedicatedMgmt resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemDedicatedMgmt(obj, mkey)
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

	err := c.DeleteSystemDedicatedMgmt(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDedicatedMgmt resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDedicatedMgmtRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemDedicatedMgmt(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemDedicatedMgmt resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDedicatedMgmt(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDedicatedMgmt resource from API: %v", err)
	}
	return nil
}

func flattenSystemDedicatedMgmtStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDedicatedMgmtInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDedicatedMgmtDefaultGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDedicatedMgmtDhcpServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDedicatedMgmtDhcpNetmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDedicatedMgmtDhcpStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDedicatedMgmtDhcpEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDedicatedMgmt(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", flattenSystemDedicatedMgmtStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemDedicatedMgmtInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("default_gateway", flattenSystemDedicatedMgmtDefaultGateway(o["default-gateway"], d, "default_gateway")); err != nil {
		if !fortiAPIPatch(o["default-gateway"]) {
			return fmt.Errorf("Error reading default_gateway: %v", err)
		}
	}

	if err = d.Set("dhcp_server", flattenSystemDedicatedMgmtDhcpServer(o["dhcp-server"], d, "dhcp_server")); err != nil {
		if !fortiAPIPatch(o["dhcp-server"]) {
			return fmt.Errorf("Error reading dhcp_server: %v", err)
		}
	}

	if err = d.Set("dhcp_netmask", flattenSystemDedicatedMgmtDhcpNetmask(o["dhcp-netmask"], d, "dhcp_netmask")); err != nil {
		if !fortiAPIPatch(o["dhcp-netmask"]) {
			return fmt.Errorf("Error reading dhcp_netmask: %v", err)
		}
	}

	if err = d.Set("dhcp_start_ip", flattenSystemDedicatedMgmtDhcpStartIp(o["dhcp-start-ip"], d, "dhcp_start_ip")); err != nil {
		if !fortiAPIPatch(o["dhcp-start-ip"]) {
			return fmt.Errorf("Error reading dhcp_start_ip: %v", err)
		}
	}

	if err = d.Set("dhcp_end_ip", flattenSystemDedicatedMgmtDhcpEndIp(o["dhcp-end-ip"], d, "dhcp_end_ip")); err != nil {
		if !fortiAPIPatch(o["dhcp-end-ip"]) {
			return fmt.Errorf("Error reading dhcp_end_ip: %v", err)
		}
	}

	return nil
}

func flattenSystemDedicatedMgmtFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDedicatedMgmtStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDedicatedMgmtInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDedicatedMgmtDefaultGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDedicatedMgmtDhcpServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDedicatedMgmtDhcpNetmask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDedicatedMgmtDhcpStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDedicatedMgmtDhcpEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDedicatedMgmt(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemDedicatedMgmtStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemDedicatedMgmtInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("default_gateway"); ok {
		t, err := expandSystemDedicatedMgmtDefaultGateway(d, v, "default_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-gateway"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_server"); ok {
		t, err := expandSystemDedicatedMgmtDhcpServer(d, v, "dhcp_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-server"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_netmask"); ok {
		t, err := expandSystemDedicatedMgmtDhcpNetmask(d, v, "dhcp_netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-netmask"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_start_ip"); ok {
		t, err := expandSystemDedicatedMgmtDhcpStartIp(d, v, "dhcp_start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_end_ip"); ok {
		t, err := expandSystemDedicatedMgmtDhcpEndIp(d, v, "dhcp_end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-end-ip"] = t
		}
	}

	return &obj, nil
}
