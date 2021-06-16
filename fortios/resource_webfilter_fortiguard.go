// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiGuard Web Filter service.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWebfilterFortiguard() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterFortiguardUpdate,
		Read:   resourceWebfilterFortiguardRead,
		Update: resourceWebfilterFortiguardUpdate,
		Delete: resourceWebfilterFortiguardDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"cache_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cache_prefix_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cache_mem_percent": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 15),
				Optional:     true,
				Computed:     true,
			},
			"ovrd_auth_port_http": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"ovrd_auth_port_https": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"ovrd_auth_port_https_flow": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"ovrd_auth_port_warning": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"ovrd_auth_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"warn_auth_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"close_ports": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_packet_size_limit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(576, 10000),
				Optional:     true,
				Computed:     true,
			},
			"ovrd_auth_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceWebfilterFortiguardUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWebfilterFortiguard(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterFortiguard resource while getting object: %v", err)
	}

	o, err := c.UpdateWebfilterFortiguard(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterFortiguard resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebfilterFortiguard")
	}

	return resourceWebfilterFortiguardRead(d, m)
}

func resourceWebfilterFortiguardDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWebfilterFortiguard(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterFortiguard resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterFortiguardRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWebfilterFortiguard(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterFortiguard resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterFortiguard(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterFortiguard resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterFortiguardCacheMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterFortiguardCachePrefixMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterFortiguardCacheMemPercent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterFortiguardOvrdAuthPortHttp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterFortiguardOvrdAuthPortHttps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterFortiguardOvrdAuthPortHttpsFlow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterFortiguardOvrdAuthPortWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterFortiguardOvrdAuthHttps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterFortiguardWarnAuthHttps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterFortiguardClosePorts(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterFortiguardRequestPacketSizeLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterFortiguardOvrdAuthPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWebfilterFortiguard(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("cache_mode", flattenWebfilterFortiguardCacheMode(o["cache-mode"], d, "cache_mode", sv)); err != nil {
		if !fortiAPIPatch(o["cache-mode"]) {
			return fmt.Errorf("Error reading cache_mode: %v", err)
		}
	}

	if err = d.Set("cache_prefix_match", flattenWebfilterFortiguardCachePrefixMatch(o["cache-prefix-match"], d, "cache_prefix_match", sv)); err != nil {
		if !fortiAPIPatch(o["cache-prefix-match"]) {
			return fmt.Errorf("Error reading cache_prefix_match: %v", err)
		}
	}

	if err = d.Set("cache_mem_percent", flattenWebfilterFortiguardCacheMemPercent(o["cache-mem-percent"], d, "cache_mem_percent", sv)); err != nil {
		if !fortiAPIPatch(o["cache-mem-percent"]) {
			return fmt.Errorf("Error reading cache_mem_percent: %v", err)
		}
	}

	if err = d.Set("ovrd_auth_port_http", flattenWebfilterFortiguardOvrdAuthPortHttp(o["ovrd-auth-port-http"], d, "ovrd_auth_port_http", sv)); err != nil {
		if !fortiAPIPatch(o["ovrd-auth-port-http"]) {
			return fmt.Errorf("Error reading ovrd_auth_port_http: %v", err)
		}
	}

	if err = d.Set("ovrd_auth_port_https", flattenWebfilterFortiguardOvrdAuthPortHttps(o["ovrd-auth-port-https"], d, "ovrd_auth_port_https", sv)); err != nil {
		if !fortiAPIPatch(o["ovrd-auth-port-https"]) {
			return fmt.Errorf("Error reading ovrd_auth_port_https: %v", err)
		}
	}

	if err = d.Set("ovrd_auth_port_https_flow", flattenWebfilterFortiguardOvrdAuthPortHttpsFlow(o["ovrd-auth-port-https-flow"], d, "ovrd_auth_port_https_flow", sv)); err != nil {
		if !fortiAPIPatch(o["ovrd-auth-port-https-flow"]) {
			return fmt.Errorf("Error reading ovrd_auth_port_https_flow: %v", err)
		}
	}

	if err = d.Set("ovrd_auth_port_warning", flattenWebfilterFortiguardOvrdAuthPortWarning(o["ovrd-auth-port-warning"], d, "ovrd_auth_port_warning", sv)); err != nil {
		if !fortiAPIPatch(o["ovrd-auth-port-warning"]) {
			return fmt.Errorf("Error reading ovrd_auth_port_warning: %v", err)
		}
	}

	if err = d.Set("ovrd_auth_https", flattenWebfilterFortiguardOvrdAuthHttps(o["ovrd-auth-https"], d, "ovrd_auth_https", sv)); err != nil {
		if !fortiAPIPatch(o["ovrd-auth-https"]) {
			return fmt.Errorf("Error reading ovrd_auth_https: %v", err)
		}
	}

	if err = d.Set("warn_auth_https", flattenWebfilterFortiguardWarnAuthHttps(o["warn-auth-https"], d, "warn_auth_https", sv)); err != nil {
		if !fortiAPIPatch(o["warn-auth-https"]) {
			return fmt.Errorf("Error reading warn_auth_https: %v", err)
		}
	}

	if err = d.Set("close_ports", flattenWebfilterFortiguardClosePorts(o["close-ports"], d, "close_ports", sv)); err != nil {
		if !fortiAPIPatch(o["close-ports"]) {
			return fmt.Errorf("Error reading close_ports: %v", err)
		}
	}

	if err = d.Set("request_packet_size_limit", flattenWebfilterFortiguardRequestPacketSizeLimit(o["request-packet-size-limit"], d, "request_packet_size_limit", sv)); err != nil {
		if !fortiAPIPatch(o["request-packet-size-limit"]) {
			return fmt.Errorf("Error reading request_packet_size_limit: %v", err)
		}
	}

	if err = d.Set("ovrd_auth_port", flattenWebfilterFortiguardOvrdAuthPort(o["ovrd-auth-port"], d, "ovrd_auth_port", sv)); err != nil {
		if !fortiAPIPatch(o["ovrd-auth-port"]) {
			return fmt.Errorf("Error reading ovrd_auth_port: %v", err)
		}
	}

	return nil
}

func flattenWebfilterFortiguardFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWebfilterFortiguardCacheMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardCachePrefixMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardCacheMemPercent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardOvrdAuthPortHttp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardOvrdAuthPortHttps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardOvrdAuthPortHttpsFlow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardOvrdAuthPortWarning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardOvrdAuthHttps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardWarnAuthHttps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardClosePorts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardRequestPacketSizeLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFortiguardOvrdAuthPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterFortiguard(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cache_mode"); ok {

		t, err := expandWebfilterFortiguardCacheMode(d, v, "cache_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-mode"] = t
		}
	}

	if v, ok := d.GetOk("cache_prefix_match"); ok {

		t, err := expandWebfilterFortiguardCachePrefixMatch(d, v, "cache_prefix_match", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-prefix-match"] = t
		}
	}

	if v, ok := d.GetOk("cache_mem_percent"); ok {

		t, err := expandWebfilterFortiguardCacheMemPercent(d, v, "cache_mem_percent", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-mem-percent"] = t
		}
	}

	if v, ok := d.GetOkExists("ovrd_auth_port_http"); ok {

		t, err := expandWebfilterFortiguardOvrdAuthPortHttp(d, v, "ovrd_auth_port_http", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-auth-port-http"] = t
		}
	}

	if v, ok := d.GetOkExists("ovrd_auth_port_https"); ok {

		t, err := expandWebfilterFortiguardOvrdAuthPortHttps(d, v, "ovrd_auth_port_https", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-auth-port-https"] = t
		}
	}

	if v, ok := d.GetOkExists("ovrd_auth_port_https_flow"); ok {

		t, err := expandWebfilterFortiguardOvrdAuthPortHttpsFlow(d, v, "ovrd_auth_port_https_flow", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-auth-port-https-flow"] = t
		}
	}

	if v, ok := d.GetOkExists("ovrd_auth_port_warning"); ok {

		t, err := expandWebfilterFortiguardOvrdAuthPortWarning(d, v, "ovrd_auth_port_warning", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-auth-port-warning"] = t
		}
	}

	if v, ok := d.GetOk("ovrd_auth_https"); ok {

		t, err := expandWebfilterFortiguardOvrdAuthHttps(d, v, "ovrd_auth_https", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-auth-https"] = t
		}
	}

	if v, ok := d.GetOk("warn_auth_https"); ok {

		t, err := expandWebfilterFortiguardWarnAuthHttps(d, v, "warn_auth_https", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warn-auth-https"] = t
		}
	}

	if v, ok := d.GetOk("close_ports"); ok {

		t, err := expandWebfilterFortiguardClosePorts(d, v, "close_ports", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["close-ports"] = t
		}
	}

	if v, ok := d.GetOk("request_packet_size_limit"); ok {

		t, err := expandWebfilterFortiguardRequestPacketSizeLimit(d, v, "request_packet_size_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-packet-size-limit"] = t
		}
	}

	if v, ok := d.GetOkExists("ovrd_auth_port"); ok {

		t, err := expandWebfilterFortiguardOvrdAuthPort(d, v, "ovrd_auth_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-auth-port"] = t
		}
	}

	return &obj, nil
}
