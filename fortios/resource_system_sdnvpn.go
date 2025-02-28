// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure public cloud VPN service.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSdnVpn() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdnVpnCreate,
		Read:   resourceSystemSdnVpnRead,
		Update: resourceSystemSdnVpnUpdate,
		Delete: resourceSystemSdnVpnDelete,

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
			"sdn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"remote_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"routing_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vgw_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"tgw_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"subnet_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"bgp_as": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cgw_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat_traversal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"internal_interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"local_cidr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_cidr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cgw_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"psksecret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"status": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
			},
			"code": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
			},
		},
	}
}

func resourceSystemSdnVpnCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSdnVpn(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdnVpn resource while getting object: %v", err)
	}

	o, err := c.CreateSystemSdnVpn(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemSdnVpn resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSdnVpn")
	}

	return resourceSystemSdnVpnRead(d, m)
}

func resourceSystemSdnVpnUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSdnVpn(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnVpn resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSdnVpn(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnVpn resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSdnVpn")
	}

	return resourceSystemSdnVpnRead(d, m)
}

func resourceSystemSdnVpnDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemSdnVpn(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSdnVpn resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdnVpnRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSdnVpn(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdnVpn resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdnVpn(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdnVpn resource from API: %v", err)
	}
	return nil
}

func flattenSystemSdnVpnName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnVpnSdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnVpnRemoteType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnVpnRoutingType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnVpnVgwId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnVpnTgwId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnVpnSubnetId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnVpnBgpAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSdnVpnCgwGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnVpnNatTraversal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnVpnTunnelInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnVpnInternalInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnVpnLocalCidr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemSdnVpnRemoteCidr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemSdnVpnCgwName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnVpnPsksecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnVpnType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSdnVpnStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSdnVpnCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectSystemSdnVpn(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemSdnVpnName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("sdn", flattenSystemSdnVpnSdn(o["sdn"], d, "sdn", sv)); err != nil {
		if !fortiAPIPatch(o["sdn"]) {
			return fmt.Errorf("Error reading sdn: %v", err)
		}
	}

	if err = d.Set("remote_type", flattenSystemSdnVpnRemoteType(o["remote-type"], d, "remote_type", sv)); err != nil {
		if !fortiAPIPatch(o["remote-type"]) {
			return fmt.Errorf("Error reading remote_type: %v", err)
		}
	}

	if err = d.Set("routing_type", flattenSystemSdnVpnRoutingType(o["routing-type"], d, "routing_type", sv)); err != nil {
		if !fortiAPIPatch(o["routing-type"]) {
			return fmt.Errorf("Error reading routing_type: %v", err)
		}
	}

	if err = d.Set("vgw_id", flattenSystemSdnVpnVgwId(o["vgw-id"], d, "vgw_id", sv)); err != nil {
		if !fortiAPIPatch(o["vgw-id"]) {
			return fmt.Errorf("Error reading vgw_id: %v", err)
		}
	}

	if err = d.Set("tgw_id", flattenSystemSdnVpnTgwId(o["tgw-id"], d, "tgw_id", sv)); err != nil {
		if !fortiAPIPatch(o["tgw-id"]) {
			return fmt.Errorf("Error reading tgw_id: %v", err)
		}
	}

	if err = d.Set("subnet_id", flattenSystemSdnVpnSubnetId(o["subnet-id"], d, "subnet_id", sv)); err != nil {
		if !fortiAPIPatch(o["subnet-id"]) {
			return fmt.Errorf("Error reading subnet_id: %v", err)
		}
	}

	if err = d.Set("bgp_as", flattenSystemSdnVpnBgpAs(o["bgp-as"], d, "bgp_as", sv)); err != nil {
		if !fortiAPIPatch(o["bgp-as"]) {
			return fmt.Errorf("Error reading bgp_as: %v", err)
		}
	}

	if err = d.Set("cgw_gateway", flattenSystemSdnVpnCgwGateway(o["cgw-gateway"], d, "cgw_gateway", sv)); err != nil {
		if !fortiAPIPatch(o["cgw-gateway"]) {
			return fmt.Errorf("Error reading cgw_gateway: %v", err)
		}
	}

	if err = d.Set("nat_traversal", flattenSystemSdnVpnNatTraversal(o["nat-traversal"], d, "nat_traversal", sv)); err != nil {
		if !fortiAPIPatch(o["nat-traversal"]) {
			return fmt.Errorf("Error reading nat_traversal: %v", err)
		}
	}

	if err = d.Set("tunnel_interface", flattenSystemSdnVpnTunnelInterface(o["tunnel-interface"], d, "tunnel_interface", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel-interface"]) {
			return fmt.Errorf("Error reading tunnel_interface: %v", err)
		}
	}

	if err = d.Set("internal_interface", flattenSystemSdnVpnInternalInterface(o["internal-interface"], d, "internal_interface", sv)); err != nil {
		if !fortiAPIPatch(o["internal-interface"]) {
			return fmt.Errorf("Error reading internal_interface: %v", err)
		}
	}

	if err = d.Set("local_cidr", flattenSystemSdnVpnLocalCidr(o["local-cidr"], d, "local_cidr", sv)); err != nil {
		if !fortiAPIPatch(o["local-cidr"]) {
			return fmt.Errorf("Error reading local_cidr: %v", err)
		}
	}

	if err = d.Set("remote_cidr", flattenSystemSdnVpnRemoteCidr(o["remote-cidr"], d, "remote_cidr", sv)); err != nil {
		if !fortiAPIPatch(o["remote-cidr"]) {
			return fmt.Errorf("Error reading remote_cidr: %v", err)
		}
	}

	if err = d.Set("cgw_name", flattenSystemSdnVpnCgwName(o["cgw-name"], d, "cgw_name", sv)); err != nil {
		if !fortiAPIPatch(o["cgw-name"]) {
			return fmt.Errorf("Error reading cgw_name: %v", err)
		}
	}

	if err = d.Set("psksecret", flattenSystemSdnVpnPsksecret(o["psksecret"], d, "psksecret", sv)); err != nil {
		if !fortiAPIPatch(o["psksecret"]) {
			return fmt.Errorf("Error reading psksecret: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemSdnVpnType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSdnVpnStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("code", flattenSystemSdnVpnCode(o["code"], d, "code", sv)); err != nil {
		if !fortiAPIPatch(o["code"]) {
			return fmt.Errorf("Error reading code: %v", err)
		}
	}

	return nil
}

func flattenSystemSdnVpnFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSdnVpnName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnSdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnRemoteType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnRoutingType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnVgwId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnTgwId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnSubnetId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnBgpAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnCgwGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnNatTraversal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnTunnelInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnInternalInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnLocalCidr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnRemoteCidr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnCgwName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnPsksecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSdnVpn(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemSdnVpnName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("sdn"); ok {
		t, err := expandSystemSdnVpnSdn(d, v, "sdn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn"] = t
		}
	} else if d.HasChange("sdn") {
		obj["sdn"] = nil
	}

	if v, ok := d.GetOk("remote_type"); ok {
		t, err := expandSystemSdnVpnRemoteType(d, v, "remote_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-type"] = t
		}
	}

	if v, ok := d.GetOk("routing_type"); ok {
		t, err := expandSystemSdnVpnRoutingType(d, v, "routing_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["routing-type"] = t
		}
	}

	if v, ok := d.GetOk("vgw_id"); ok {
		t, err := expandSystemSdnVpnVgwId(d, v, "vgw_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vgw-id"] = t
		}
	} else if d.HasChange("vgw_id") {
		obj["vgw-id"] = nil
	}

	if v, ok := d.GetOk("tgw_id"); ok {
		t, err := expandSystemSdnVpnTgwId(d, v, "tgw_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tgw-id"] = t
		}
	} else if d.HasChange("tgw_id") {
		obj["tgw-id"] = nil
	}

	if v, ok := d.GetOk("subnet_id"); ok {
		t, err := expandSystemSdnVpnSubnetId(d, v, "subnet_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet-id"] = t
		}
	} else if d.HasChange("subnet_id") {
		obj["subnet-id"] = nil
	}

	if v, ok := d.GetOk("bgp_as"); ok {
		t, err := expandSystemSdnVpnBgpAs(d, v, "bgp_as", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bgp-as"] = t
		}
	}

	if v, ok := d.GetOk("cgw_gateway"); ok {
		t, err := expandSystemSdnVpnCgwGateway(d, v, "cgw_gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgw-gateway"] = t
		}
	}

	if v, ok := d.GetOk("nat_traversal"); ok {
		t, err := expandSystemSdnVpnNatTraversal(d, v, "nat_traversal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat-traversal"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_interface"); ok {
		t, err := expandSystemSdnVpnTunnelInterface(d, v, "tunnel_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-interface"] = t
		}
	} else if d.HasChange("tunnel_interface") {
		obj["tunnel-interface"] = nil
	}

	if v, ok := d.GetOk("internal_interface"); ok {
		t, err := expandSystemSdnVpnInternalInterface(d, v, "internal_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal-interface"] = t
		}
	} else if d.HasChange("internal_interface") {
		obj["internal-interface"] = nil
	}

	if v, ok := d.GetOk("local_cidr"); ok {
		t, err := expandSystemSdnVpnLocalCidr(d, v, "local_cidr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-cidr"] = t
		}
	}

	if v, ok := d.GetOk("remote_cidr"); ok {
		t, err := expandSystemSdnVpnRemoteCidr(d, v, "remote_cidr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-cidr"] = t
		}
	}

	if v, ok := d.GetOk("cgw_name"); ok {
		t, err := expandSystemSdnVpnCgwName(d, v, "cgw_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgw-name"] = t
		}
	} else if d.HasChange("cgw_name") {
		obj["cgw-name"] = nil
	}

	if v, ok := d.GetOk("psksecret"); ok {
		t, err := expandSystemSdnVpnPsksecret(d, v, "psksecret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psksecret"] = t
		}
	} else if d.HasChange("psksecret") {
		obj["psksecret"] = nil
	}

	if v, ok := d.GetOkExists("type"); ok {
		t, err := expandSystemSdnVpnType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOkExists("status"); ok {
		t, err := expandSystemSdnVpnStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	} else if d.HasChange("status") {
		obj["status"] = nil
	}

	if v, ok := d.GetOkExists("code"); ok {
		t, err := expandSystemSdnVpnCode(d, v, "code", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["code"] = t
		}
	} else if d.HasChange("code") {
		obj["code"] = nil
	}

	return &obj, nil
}
