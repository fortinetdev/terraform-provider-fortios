// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure wireless access gateway (WAG) profiles used for tunnels on AP.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerWagProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerWagProfileCreate,
		Read:   resourceWirelessControllerWagProfileRead,
		Update: resourceWirelessControllerWagProfileUpdate,
		Delete: resourceWirelessControllerWagProfileDelete,

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
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"tunnel_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wag_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wag_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"ping_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"ping_number": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"return_packet_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"dhcp_ip_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerWagProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerWagProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWagProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerWagProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWagProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerWagProfile")
	}

	return resourceWirelessControllerWagProfileRead(d, m)
}

func resourceWirelessControllerWagProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerWagProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWagProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerWagProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWagProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerWagProfile")
	}

	return resourceWirelessControllerWagProfileRead(d, m)
}

func resourceWirelessControllerWagProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerWagProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerWagProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWagProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerWagProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWagProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerWagProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWagProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerWagProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWagProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWagProfileTunnelType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWagProfileWagIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWagProfileWagPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWagProfilePingInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWagProfilePingNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWagProfileReturnPacketTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWagProfileDhcpIpAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerWagProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerWagProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerWagProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("tunnel_type", flattenWirelessControllerWagProfileTunnelType(o["tunnel-type"], d, "tunnel_type", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel-type"]) {
			return fmt.Errorf("Error reading tunnel_type: %v", err)
		}
	}

	if err = d.Set("wag_ip", flattenWirelessControllerWagProfileWagIp(o["wag-ip"], d, "wag_ip", sv)); err != nil {
		if !fortiAPIPatch(o["wag-ip"]) {
			return fmt.Errorf("Error reading wag_ip: %v", err)
		}
	}

	if err = d.Set("wag_port", flattenWirelessControllerWagProfileWagPort(o["wag-port"], d, "wag_port", sv)); err != nil {
		if !fortiAPIPatch(o["wag-port"]) {
			return fmt.Errorf("Error reading wag_port: %v", err)
		}
	}

	if err = d.Set("ping_interval", flattenWirelessControllerWagProfilePingInterval(o["ping-interval"], d, "ping_interval", sv)); err != nil {
		if !fortiAPIPatch(o["ping-interval"]) {
			return fmt.Errorf("Error reading ping_interval: %v", err)
		}
	}

	if err = d.Set("ping_number", flattenWirelessControllerWagProfilePingNumber(o["ping-number"], d, "ping_number", sv)); err != nil {
		if !fortiAPIPatch(o["ping-number"]) {
			return fmt.Errorf("Error reading ping_number: %v", err)
		}
	}

	if err = d.Set("return_packet_timeout", flattenWirelessControllerWagProfileReturnPacketTimeout(o["return-packet-timeout"], d, "return_packet_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["return-packet-timeout"]) {
			return fmt.Errorf("Error reading return_packet_timeout: %v", err)
		}
	}

	if err = d.Set("dhcp_ip_addr", flattenWirelessControllerWagProfileDhcpIpAddr(o["dhcp-ip-addr"], d, "dhcp_ip_addr", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-ip-addr"]) {
			return fmt.Errorf("Error reading dhcp_ip_addr: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerWagProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerWagProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWagProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWagProfileTunnelType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWagProfileWagIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWagProfileWagPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWagProfilePingInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWagProfilePingNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWagProfileReturnPacketTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWagProfileDhcpIpAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerWagProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerWagProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWirelessControllerWagProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("tunnel_type"); ok {
		t, err := expandWirelessControllerWagProfileTunnelType(d, v, "tunnel_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-type"] = t
		}
	}

	if v, ok := d.GetOk("wag_ip"); ok {
		t, err := expandWirelessControllerWagProfileWagIp(d, v, "wag_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wag-ip"] = t
		}
	}

	if v, ok := d.GetOkExists("wag_port"); ok {
		t, err := expandWirelessControllerWagProfileWagPort(d, v, "wag_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wag-port"] = t
		}
	}

	if v, ok := d.GetOk("ping_interval"); ok {
		t, err := expandWirelessControllerWagProfilePingInterval(d, v, "ping_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ping-interval"] = t
		}
	}

	if v, ok := d.GetOk("ping_number"); ok {
		t, err := expandWirelessControllerWagProfilePingNumber(d, v, "ping_number", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ping-number"] = t
		}
	}

	if v, ok := d.GetOk("return_packet_timeout"); ok {
		t, err := expandWirelessControllerWagProfileReturnPacketTimeout(d, v, "return_packet_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["return-packet-timeout"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_ip_addr"); ok {
		t, err := expandWirelessControllerWagProfileDhcpIpAddr(d, v, "dhcp_ip_addr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-ip-addr"] = t
		}
	}

	return &obj, nil
}
