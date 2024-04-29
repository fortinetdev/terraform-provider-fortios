// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure GRE tunnel.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemGreTunnel() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemGreTunnelCreate,
		Read:   resourceSystemGreTunnelRead,
		Update: resourceSystemGreTunnelUpdate,
		Delete: resourceSystemGreTunnelDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 15),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"ip_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_gw6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"local_gw": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"use_sdwan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sequence_number_transmission": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sequence_number_reception": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"checksum_transmission": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"checksum_reception": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key_outbound": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"key_inbound": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dscp_copying": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keepalive_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32767),
				Optional:     true,
				Computed:     true,
			},
			"keepalive_failtimes": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemGreTunnelCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemGreTunnel(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemGreTunnel resource while getting object: %v", err)
	}

	o, err := c.CreateSystemGreTunnel(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemGreTunnel resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemGreTunnel")
	}

	return resourceSystemGreTunnelRead(d, m)
}

func resourceSystemGreTunnelUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemGreTunnel(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemGreTunnel resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemGreTunnel(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemGreTunnel resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemGreTunnel")
	}

	return resourceSystemGreTunnelRead(d, m)
}

func resourceSystemGreTunnelDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemGreTunnel(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemGreTunnel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemGreTunnelRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemGreTunnel(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemGreTunnel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemGreTunnel(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemGreTunnel resource from API: %v", err)
	}
	return nil
}

func flattenSystemGreTunnelName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGreTunnelInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGreTunnelIpVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGreTunnelRemoteGw6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGreTunnelLocalGw6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGreTunnelRemoteGw(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGreTunnelLocalGw(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGreTunnelUseSdwan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGreTunnelSequenceNumberTransmission(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGreTunnelSequenceNumberReception(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGreTunnelChecksumTransmission(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGreTunnelChecksumReception(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGreTunnelKeyOutbound(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGreTunnelKeyInbound(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGreTunnelDscpCopying(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGreTunnelDiffservcode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGreTunnelKeepaliveInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGreTunnelKeepaliveFailtimes(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemGreTunnel(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemGreTunnelName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemGreTunnelInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip_version", flattenSystemGreTunnelIpVersion(o["ip-version"], d, "ip_version", sv)); err != nil {
		if !fortiAPIPatch(o["ip-version"]) {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if err = d.Set("remote_gw6", flattenSystemGreTunnelRemoteGw6(o["remote-gw6"], d, "remote_gw6", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw6"]) {
			return fmt.Errorf("Error reading remote_gw6: %v", err)
		}
	}

	if err = d.Set("local_gw6", flattenSystemGreTunnelLocalGw6(o["local-gw6"], d, "local_gw6", sv)); err != nil {
		if !fortiAPIPatch(o["local-gw6"]) {
			return fmt.Errorf("Error reading local_gw6: %v", err)
		}
	}

	if err = d.Set("remote_gw", flattenSystemGreTunnelRemoteGw(o["remote-gw"], d, "remote_gw", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw"]) {
			return fmt.Errorf("Error reading remote_gw: %v", err)
		}
	}

	if err = d.Set("local_gw", flattenSystemGreTunnelLocalGw(o["local-gw"], d, "local_gw", sv)); err != nil {
		if !fortiAPIPatch(o["local-gw"]) {
			return fmt.Errorf("Error reading local_gw: %v", err)
		}
	}

	if err = d.Set("use_sdwan", flattenSystemGreTunnelUseSdwan(o["use-sdwan"], d, "use_sdwan", sv)); err != nil {
		if !fortiAPIPatch(o["use-sdwan"]) {
			return fmt.Errorf("Error reading use_sdwan: %v", err)
		}
	}

	if err = d.Set("sequence_number_transmission", flattenSystemGreTunnelSequenceNumberTransmission(o["sequence-number-transmission"], d, "sequence_number_transmission", sv)); err != nil {
		if !fortiAPIPatch(o["sequence-number-transmission"]) {
			return fmt.Errorf("Error reading sequence_number_transmission: %v", err)
		}
	}

	if err = d.Set("sequence_number_reception", flattenSystemGreTunnelSequenceNumberReception(o["sequence-number-reception"], d, "sequence_number_reception", sv)); err != nil {
		if !fortiAPIPatch(o["sequence-number-reception"]) {
			return fmt.Errorf("Error reading sequence_number_reception: %v", err)
		}
	}

	if err = d.Set("checksum_transmission", flattenSystemGreTunnelChecksumTransmission(o["checksum-transmission"], d, "checksum_transmission", sv)); err != nil {
		if !fortiAPIPatch(o["checksum-transmission"]) {
			return fmt.Errorf("Error reading checksum_transmission: %v", err)
		}
	}

	if err = d.Set("checksum_reception", flattenSystemGreTunnelChecksumReception(o["checksum-reception"], d, "checksum_reception", sv)); err != nil {
		if !fortiAPIPatch(o["checksum-reception"]) {
			return fmt.Errorf("Error reading checksum_reception: %v", err)
		}
	}

	if err = d.Set("key_outbound", flattenSystemGreTunnelKeyOutbound(o["key-outbound"], d, "key_outbound", sv)); err != nil {
		if !fortiAPIPatch(o["key-outbound"]) {
			return fmt.Errorf("Error reading key_outbound: %v", err)
		}
	}

	if err = d.Set("key_inbound", flattenSystemGreTunnelKeyInbound(o["key-inbound"], d, "key_inbound", sv)); err != nil {
		if !fortiAPIPatch(o["key-inbound"]) {
			return fmt.Errorf("Error reading key_inbound: %v", err)
		}
	}

	if err = d.Set("dscp_copying", flattenSystemGreTunnelDscpCopying(o["dscp-copying"], d, "dscp_copying", sv)); err != nil {
		if !fortiAPIPatch(o["dscp-copying"]) {
			return fmt.Errorf("Error reading dscp_copying: %v", err)
		}
	}

	if err = d.Set("diffservcode", flattenSystemGreTunnelDiffservcode(o["diffservcode"], d, "diffservcode", sv)); err != nil {
		if !fortiAPIPatch(o["diffservcode"]) {
			return fmt.Errorf("Error reading diffservcode: %v", err)
		}
	}

	if err = d.Set("keepalive_interval", flattenSystemGreTunnelKeepaliveInterval(o["keepalive-interval"], d, "keepalive_interval", sv)); err != nil {
		if !fortiAPIPatch(o["keepalive-interval"]) {
			return fmt.Errorf("Error reading keepalive_interval: %v", err)
		}
	}

	if err = d.Set("keepalive_failtimes", flattenSystemGreTunnelKeepaliveFailtimes(o["keepalive-failtimes"], d, "keepalive_failtimes", sv)); err != nil {
		if !fortiAPIPatch(o["keepalive-failtimes"]) {
			return fmt.Errorf("Error reading keepalive_failtimes: %v", err)
		}
	}

	return nil
}

func flattenSystemGreTunnelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemGreTunnelName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelIpVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelRemoteGw6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelLocalGw6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelRemoteGw(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelLocalGw(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelUseSdwan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelSequenceNumberTransmission(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelSequenceNumberReception(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelChecksumTransmission(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelChecksumReception(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelKeyOutbound(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelKeyInbound(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelDscpCopying(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelDiffservcode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelKeepaliveInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelKeepaliveFailtimes(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemGreTunnel(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemGreTunnelName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemGreTunnelInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip_version"); ok {
		t, err := expandSystemGreTunnelIpVersion(d, v, "ip_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-version"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6"); ok {
		t, err := expandSystemGreTunnelRemoteGw6(d, v, "remote_gw6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6"] = t
		}
	}

	if v, ok := d.GetOk("local_gw6"); ok {
		t, err := expandSystemGreTunnelLocalGw6(d, v, "local_gw6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-gw6"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw"); ok {
		t, err := expandSystemGreTunnelRemoteGw(d, v, "remote_gw", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw"] = t
		}
	}

	if v, ok := d.GetOk("local_gw"); ok {
		t, err := expandSystemGreTunnelLocalGw(d, v, "local_gw", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-gw"] = t
		}
	}

	if v, ok := d.GetOk("use_sdwan"); ok {
		t, err := expandSystemGreTunnelUseSdwan(d, v, "use_sdwan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-sdwan"] = t
		}
	}

	if v, ok := d.GetOk("sequence_number_transmission"); ok {
		t, err := expandSystemGreTunnelSequenceNumberTransmission(d, v, "sequence_number_transmission", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sequence-number-transmission"] = t
		}
	}

	if v, ok := d.GetOk("sequence_number_reception"); ok {
		t, err := expandSystemGreTunnelSequenceNumberReception(d, v, "sequence_number_reception", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sequence-number-reception"] = t
		}
	}

	if v, ok := d.GetOk("checksum_transmission"); ok {
		t, err := expandSystemGreTunnelChecksumTransmission(d, v, "checksum_transmission", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["checksum-transmission"] = t
		}
	}

	if v, ok := d.GetOk("checksum_reception"); ok {
		t, err := expandSystemGreTunnelChecksumReception(d, v, "checksum_reception", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["checksum-reception"] = t
		}
	}

	if v, ok := d.GetOkExists("key_outbound"); ok {
		t, err := expandSystemGreTunnelKeyOutbound(d, v, "key_outbound", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-outbound"] = t
		}
	}

	if v, ok := d.GetOkExists("key_inbound"); ok {
		t, err := expandSystemGreTunnelKeyInbound(d, v, "key_inbound", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-inbound"] = t
		}
	}

	if v, ok := d.GetOk("dscp_copying"); ok {
		t, err := expandSystemGreTunnelDscpCopying(d, v, "dscp_copying", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-copying"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode"); ok {
		t, err := expandSystemGreTunnelDiffservcode(d, v, "diffservcode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode"] = t
		}
	}

	if v, ok := d.GetOkExists("keepalive_interval"); ok {
		t, err := expandSystemGreTunnelKeepaliveInterval(d, v, "keepalive_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keepalive-interval"] = t
		}
	}

	if v, ok := d.GetOk("keepalive_failtimes"); ok {
		t, err := expandSystemGreTunnelKeepaliveFailtimes(d, v, "keepalive_failtimes", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keepalive-failtimes"] = t
		}
	}

	return &obj, nil
}
