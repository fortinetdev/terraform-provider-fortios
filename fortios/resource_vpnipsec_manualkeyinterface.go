// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPsec manual keys.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnIpsecManualkeyInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnIpsecManualkeyInterfaceCreate,
		Read:   resourceVpnIpsecManualkeyInterfaceRead,
		Update: resourceVpnIpsecManualkeyInterfaceUpdate,
		Delete: resourceVpnIpsecManualkeyInterfaceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
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
				Required:     true,
			},
			"ip_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"remote_gw6": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"local_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_gw6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_alg": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"enc_alg": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"auth_key": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"enc_key": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"local_spi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_spi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"npu_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVpnIpsecManualkeyInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnIpsecManualkeyInterface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecManualkeyInterface resource while getting object: %v", err)
	}

	o, err := c.CreateVpnIpsecManualkeyInterface(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecManualkeyInterface resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecManualkeyInterface")
	}

	return resourceVpnIpsecManualkeyInterfaceRead(d, m)
}

func resourceVpnIpsecManualkeyInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnIpsecManualkeyInterface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecManualkeyInterface resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnIpsecManualkeyInterface(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecManualkeyInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecManualkeyInterface")
	}

	return resourceVpnIpsecManualkeyInterfaceRead(d, m)
}

func resourceVpnIpsecManualkeyInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteVpnIpsecManualkeyInterface(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting VpnIpsecManualkeyInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecManualkeyInterfaceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadVpnIpsecManualkeyInterface(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecManualkeyInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnIpsecManualkeyInterface(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecManualkeyInterface resource from API: %v", err)
	}
	return nil
}

func flattenVpnIpsecManualkeyInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyInterfaceInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyInterfaceIpVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyInterfaceAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyInterfaceRemoteGw(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyInterfaceRemoteGw6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyInterfaceLocalGw(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyInterfaceLocalGw6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyInterfaceAuthAlg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyInterfaceEncAlg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyInterfaceAuthKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyInterfaceEncKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyInterfaceLocalSpi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyInterfaceRemoteSpi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyInterfaceNpuOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnIpsecManualkeyInterface(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenVpnIpsecManualkeyInterfaceName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("interface", flattenVpnIpsecManualkeyInterfaceInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip_version", flattenVpnIpsecManualkeyInterfaceIpVersion(o["ip-version"], d, "ip_version", sv)); err != nil {
		if !fortiAPIPatch(o["ip-version"]) {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if err = d.Set("addr_type", flattenVpnIpsecManualkeyInterfaceAddrType(o["addr-type"], d, "addr_type", sv)); err != nil {
		if !fortiAPIPatch(o["addr-type"]) {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("remote_gw", flattenVpnIpsecManualkeyInterfaceRemoteGw(o["remote-gw"], d, "remote_gw", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw"]) {
			return fmt.Errorf("Error reading remote_gw: %v", err)
		}
	}

	if err = d.Set("remote_gw6", flattenVpnIpsecManualkeyInterfaceRemoteGw6(o["remote-gw6"], d, "remote_gw6", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw6"]) {
			return fmt.Errorf("Error reading remote_gw6: %v", err)
		}
	}

	if err = d.Set("local_gw", flattenVpnIpsecManualkeyInterfaceLocalGw(o["local-gw"], d, "local_gw", sv)); err != nil {
		if !fortiAPIPatch(o["local-gw"]) {
			return fmt.Errorf("Error reading local_gw: %v", err)
		}
	}

	if err = d.Set("local_gw6", flattenVpnIpsecManualkeyInterfaceLocalGw6(o["local-gw6"], d, "local_gw6", sv)); err != nil {
		if !fortiAPIPatch(o["local-gw6"]) {
			return fmt.Errorf("Error reading local_gw6: %v", err)
		}
	}

	if err = d.Set("auth_alg", flattenVpnIpsecManualkeyInterfaceAuthAlg(o["auth-alg"], d, "auth_alg", sv)); err != nil {
		if !fortiAPIPatch(o["auth-alg"]) {
			return fmt.Errorf("Error reading auth_alg: %v", err)
		}
	}

	if err = d.Set("enc_alg", flattenVpnIpsecManualkeyInterfaceEncAlg(o["enc-alg"], d, "enc_alg", sv)); err != nil {
		if !fortiAPIPatch(o["enc-alg"]) {
			return fmt.Errorf("Error reading enc_alg: %v", err)
		}
	}

	if err = d.Set("local_spi", flattenVpnIpsecManualkeyInterfaceLocalSpi(o["local-spi"], d, "local_spi", sv)); err != nil {
		if !fortiAPIPatch(o["local-spi"]) {
			return fmt.Errorf("Error reading local_spi: %v", err)
		}
	}

	if err = d.Set("remote_spi", flattenVpnIpsecManualkeyInterfaceRemoteSpi(o["remote-spi"], d, "remote_spi", sv)); err != nil {
		if !fortiAPIPatch(o["remote-spi"]) {
			return fmt.Errorf("Error reading remote_spi: %v", err)
		}
	}

	if err = d.Set("npu_offload", flattenVpnIpsecManualkeyInterfaceNpuOffload(o["npu-offload"], d, "npu_offload", sv)); err != nil {
		if !fortiAPIPatch(o["npu-offload"]) {
			return fmt.Errorf("Error reading npu_offload: %v", err)
		}
	}

	return nil
}

func flattenVpnIpsecManualkeyInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnIpsecManualkeyInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyInterfaceInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyInterfaceIpVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyInterfaceAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyInterfaceRemoteGw(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyInterfaceRemoteGw6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyInterfaceLocalGw(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyInterfaceLocalGw6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyInterfaceAuthAlg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyInterfaceEncAlg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyInterfaceAuthKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyInterfaceEncKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyInterfaceLocalSpi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyInterfaceRemoteSpi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyInterfaceNpuOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnIpsecManualkeyInterface(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandVpnIpsecManualkeyInterfaceName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandVpnIpsecManualkeyInterfaceInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip_version"); ok {
		t, err := expandVpnIpsecManualkeyInterfaceIpVersion(d, v, "ip_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-version"] = t
		}
	}

	if v, ok := d.GetOk("addr_type"); ok {
		t, err := expandVpnIpsecManualkeyInterfaceAddrType(d, v, "addr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw"); ok {
		t, err := expandVpnIpsecManualkeyInterfaceRemoteGw(d, v, "remote_gw", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6"); ok {
		t, err := expandVpnIpsecManualkeyInterfaceRemoteGw6(d, v, "remote_gw6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6"] = t
		}
	}

	if v, ok := d.GetOk("local_gw"); ok {
		t, err := expandVpnIpsecManualkeyInterfaceLocalGw(d, v, "local_gw", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-gw"] = t
		}
	}

	if v, ok := d.GetOk("local_gw6"); ok {
		t, err := expandVpnIpsecManualkeyInterfaceLocalGw6(d, v, "local_gw6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-gw6"] = t
		}
	}

	if v, ok := d.GetOk("auth_alg"); ok {
		t, err := expandVpnIpsecManualkeyInterfaceAuthAlg(d, v, "auth_alg", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-alg"] = t
		}
	}

	if v, ok := d.GetOk("enc_alg"); ok {
		t, err := expandVpnIpsecManualkeyInterfaceEncAlg(d, v, "enc_alg", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-alg"] = t
		}
	}

	if v, ok := d.GetOk("auth_key"); ok {
		t, err := expandVpnIpsecManualkeyInterfaceAuthKey(d, v, "auth_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-key"] = t
		}
	}

	if v, ok := d.GetOk("enc_key"); ok {
		t, err := expandVpnIpsecManualkeyInterfaceEncKey(d, v, "enc_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-key"] = t
		}
	}

	if v, ok := d.GetOk("local_spi"); ok {
		t, err := expandVpnIpsecManualkeyInterfaceLocalSpi(d, v, "local_spi", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-spi"] = t
		}
	}

	if v, ok := d.GetOk("remote_spi"); ok {
		t, err := expandVpnIpsecManualkeyInterfaceRemoteSpi(d, v, "remote_spi", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-spi"] = t
		}
	}

	if v, ok := d.GetOk("npu_offload"); ok {
		t, err := expandVpnIpsecManualkeyInterfaceNpuOffload(d, v, "npu_offload", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-offload"] = t
		}
	}

	return &obj, nil
}
