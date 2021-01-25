// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure connection capability.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerHotspot20H2QpConnCapability() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20H2QpConnCapabilityCreate,
		Read:   resourceWirelessControllerHotspot20H2QpConnCapabilityRead,
		Update: resourceWirelessControllerHotspot20H2QpConnCapabilityUpdate,
		Delete: resourceWirelessControllerHotspot20H2QpConnCapabilityDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"icmp_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftp_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pptp_vpn_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"voip_tcp_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"voip_udp_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ikev2_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ikev2_xx_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"esp_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20H2QpConnCapabilityCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerHotspot20H2QpConnCapability(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpConnCapability resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerHotspot20H2QpConnCapability(obj)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpConnCapability resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20H2QpConnCapability")
	}

	return resourceWirelessControllerHotspot20H2QpConnCapabilityRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpConnCapabilityUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerHotspot20H2QpConnCapability(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpConnCapability resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerHotspot20H2QpConnCapability(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpConnCapability resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20H2QpConnCapability")
	}

	return resourceWirelessControllerHotspot20H2QpConnCapabilityRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpConnCapabilityDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWirelessControllerHotspot20H2QpConnCapability(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20H2QpConnCapability resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20H2QpConnCapabilityRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWirelessControllerHotspot20H2QpConnCapability(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpConnCapability resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20H2QpConnCapability(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpConnCapability resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityIcmpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityFtpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilitySshPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityHttpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityTlsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityPptpVpnPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityVoipTcpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityVoipUdpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityIkev2Port(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityIkev2XxPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityEspPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20H2QpConnCapability(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerHotspot20H2QpConnCapabilityName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("icmp_port", flattenWirelessControllerHotspot20H2QpConnCapabilityIcmpPort(o["icmp-port"], d, "icmp_port")); err != nil {
		if !fortiAPIPatch(o["icmp-port"]) {
			return fmt.Errorf("Error reading icmp_port: %v", err)
		}
	}

	if err = d.Set("ftp_port", flattenWirelessControllerHotspot20H2QpConnCapabilityFtpPort(o["ftp-port"], d, "ftp_port")); err != nil {
		if !fortiAPIPatch(o["ftp-port"]) {
			return fmt.Errorf("Error reading ftp_port: %v", err)
		}
	}

	if err = d.Set("ssh_port", flattenWirelessControllerHotspot20H2QpConnCapabilitySshPort(o["ssh-port"], d, "ssh_port")); err != nil {
		if !fortiAPIPatch(o["ssh-port"]) {
			return fmt.Errorf("Error reading ssh_port: %v", err)
		}
	}

	if err = d.Set("http_port", flattenWirelessControllerHotspot20H2QpConnCapabilityHttpPort(o["http-port"], d, "http_port")); err != nil {
		if !fortiAPIPatch(o["http-port"]) {
			return fmt.Errorf("Error reading http_port: %v", err)
		}
	}

	if err = d.Set("tls_port", flattenWirelessControllerHotspot20H2QpConnCapabilityTlsPort(o["tls-port"], d, "tls_port")); err != nil {
		if !fortiAPIPatch(o["tls-port"]) {
			return fmt.Errorf("Error reading tls_port: %v", err)
		}
	}

	if err = d.Set("pptp_vpn_port", flattenWirelessControllerHotspot20H2QpConnCapabilityPptpVpnPort(o["pptp-vpn-port"], d, "pptp_vpn_port")); err != nil {
		if !fortiAPIPatch(o["pptp-vpn-port"]) {
			return fmt.Errorf("Error reading pptp_vpn_port: %v", err)
		}
	}

	if err = d.Set("voip_tcp_port", flattenWirelessControllerHotspot20H2QpConnCapabilityVoipTcpPort(o["voip-tcp-port"], d, "voip_tcp_port")); err != nil {
		if !fortiAPIPatch(o["voip-tcp-port"]) {
			return fmt.Errorf("Error reading voip_tcp_port: %v", err)
		}
	}

	if err = d.Set("voip_udp_port", flattenWirelessControllerHotspot20H2QpConnCapabilityVoipUdpPort(o["voip-udp-port"], d, "voip_udp_port")); err != nil {
		if !fortiAPIPatch(o["voip-udp-port"]) {
			return fmt.Errorf("Error reading voip_udp_port: %v", err)
		}
	}

	if err = d.Set("ikev2_port", flattenWirelessControllerHotspot20H2QpConnCapabilityIkev2Port(o["ikev2-port"], d, "ikev2_port")); err != nil {
		if !fortiAPIPatch(o["ikev2-port"]) {
			return fmt.Errorf("Error reading ikev2_port: %v", err)
		}
	}

	if err = d.Set("ikev2_xx_port", flattenWirelessControllerHotspot20H2QpConnCapabilityIkev2XxPort(o["ikev2-xx-port"], d, "ikev2_xx_port")); err != nil {
		if !fortiAPIPatch(o["ikev2-xx-port"]) {
			return fmt.Errorf("Error reading ikev2_xx_port: %v", err)
		}
	}

	if err = d.Set("esp_port", flattenWirelessControllerHotspot20H2QpConnCapabilityEspPort(o["esp-port"], d, "esp_port")); err != nil {
		if !fortiAPIPatch(o["esp-port"]) {
			return fmt.Errorf("Error reading esp_port: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20H2QpConnCapabilityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20H2QpConnCapabilityName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityIcmpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityFtpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilitySshPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityHttpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityTlsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityPptpVpnPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityVoipTcpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityVoipUdpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityIkev2Port(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityIkev2XxPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpConnCapabilityEspPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20H2QpConnCapability(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("icmp_port"); ok {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityIcmpPort(d, v, "icmp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-port"] = t
		}
	}

	if v, ok := d.GetOk("ftp_port"); ok {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityFtpPort(d, v, "ftp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp-port"] = t
		}
	}

	if v, ok := d.GetOk("ssh_port"); ok {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilitySshPort(d, v, "ssh_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-port"] = t
		}
	}

	if v, ok := d.GetOk("http_port"); ok {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityHttpPort(d, v, "http_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-port"] = t
		}
	}

	if v, ok := d.GetOk("tls_port"); ok {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityTlsPort(d, v, "tls_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-port"] = t
		}
	}

	if v, ok := d.GetOk("pptp_vpn_port"); ok {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityPptpVpnPort(d, v, "pptp_vpn_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pptp-vpn-port"] = t
		}
	}

	if v, ok := d.GetOk("voip_tcp_port"); ok {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityVoipTcpPort(d, v, "voip_tcp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip-tcp-port"] = t
		}
	}

	if v, ok := d.GetOk("voip_udp_port"); ok {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityVoipUdpPort(d, v, "voip_udp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip-udp-port"] = t
		}
	}

	if v, ok := d.GetOk("ikev2_port"); ok {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityIkev2Port(d, v, "ikev2_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ikev2-port"] = t
		}
	}

	if v, ok := d.GetOk("ikev2_xx_port"); ok {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityIkev2XxPort(d, v, "ikev2_xx_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ikev2-xx-port"] = t
		}
	}

	if v, ok := d.GetOk("esp_port"); ok {
		t, err := expandWirelessControllerHotspot20H2QpConnCapabilityEspPort(d, v, "esp_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["esp-port"] = t
		}
	}

	return &obj, nil
}
