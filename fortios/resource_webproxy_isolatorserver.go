// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure forward-server addresses.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebProxyIsolatorServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyIsolatorServerCreate,
		Read:   resourceWebProxyIsolatorServerRead,
		Update: resourceWebProxyIsolatorServerUpdate,
		Delete: resourceWebProxyIsolatorServerDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fqdn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"vrf_select": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: intBetweenWithMinus1(0, 511),
				Optional:     true,
				Computed:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
		},
	}
}

func resourceWebProxyIsolatorServerCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyIsolatorServer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyIsolatorServer resource while getting object: %v", err)
	}

	o, err := c.CreateWebProxyIsolatorServer(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WebProxyIsolatorServer resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebProxyIsolatorServer")
	}

	return resourceWebProxyIsolatorServerRead(d, m)
}

func resourceWebProxyIsolatorServerUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyIsolatorServer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyIsolatorServer resource while getting object: %v", err)
	}

	o, err := c.UpdateWebProxyIsolatorServer(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyIsolatorServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebProxyIsolatorServer")
	}

	return resourceWebProxyIsolatorServerRead(d, m)
}

func resourceWebProxyIsolatorServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWebProxyIsolatorServer(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyIsolatorServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyIsolatorServerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebProxyIsolatorServer(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyIsolatorServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyIsolatorServer(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyIsolatorServer resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyIsolatorServerName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyIsolatorServerAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyIsolatorServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyIsolatorServerIpv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyIsolatorServerFqdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyIsolatorServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWebProxyIsolatorServerInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyIsolatorServerInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyIsolatorServerVrfSelect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if vf, ok := v.(float64); ok && (vf < 0 || vf == 4294967295) {
		var vi interface{}
		vi = -1
		return vi
	}
	return v
}

func flattenWebProxyIsolatorServerComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWebProxyIsolatorServer(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWebProxyIsolatorServerName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("addr_type", flattenWebProxyIsolatorServerAddrType(o["addr-type"], d, "addr_type", sv)); err != nil {
		if !fortiAPIPatch(o["addr-type"]) {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("ip", flattenWebProxyIsolatorServerIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ipv6", flattenWebProxyIsolatorServerIpv6(o["ipv6"], d, "ipv6", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6"]) {
			return fmt.Errorf("Error reading ipv6: %v", err)
		}
	}

	if err = d.Set("fqdn", flattenWebProxyIsolatorServerFqdn(o["fqdn"], d, "fqdn", sv)); err != nil {
		if !fortiAPIPatch(o["fqdn"]) {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("port", flattenWebProxyIsolatorServerPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenWebProxyIsolatorServerInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenWebProxyIsolatorServerInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenWebProxyIsolatorServerVrfSelect(o["vrf-select"], d, "vrf_select", sv)); err != nil {
		if !fortiAPIPatch(o["vrf-select"]) {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	if err = d.Set("comment", flattenWebProxyIsolatorServerComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	return nil
}

func flattenWebProxyIsolatorServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWebProxyIsolatorServerName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyIsolatorServerAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyIsolatorServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyIsolatorServerIpv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyIsolatorServerFqdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyIsolatorServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyIsolatorServerInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyIsolatorServerInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyIsolatorServerVrfSelect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	if vi, ok := v.(int); ok && vi < 0 {
		return nil, nil
	}
	return v, nil
}

func expandWebProxyIsolatorServerComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyIsolatorServer(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWebProxyIsolatorServerName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("addr_type"); ok {
		t, err := expandWebProxyIsolatorServerAddrType(d, v, "addr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandWebProxyIsolatorServerIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv6"); ok {
		t, err := expandWebProxyIsolatorServerIpv6(d, v, "ipv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6"] = t
		}
	}

	if v, ok := d.GetOk("fqdn"); ok {
		t, err := expandWebProxyIsolatorServerFqdn(d, v, "fqdn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn"] = t
		}
	} else if d.HasChange("fqdn") {
		obj["fqdn"] = nil
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandWebProxyIsolatorServerPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		t, err := expandWebProxyIsolatorServerInterfaceSelectMethod(d, v, "interface_select_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandWebProxyIsolatorServerInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOkExists("vrf_select"); ok {
		t, err := expandWebProxyIsolatorServerVrfSelect(d, v, "vrf_select", sv)
		if err != nil {
			return &obj, err
		}
		obj["vrf-select"] = t
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWebProxyIsolatorServerComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	return &obj, nil
}
