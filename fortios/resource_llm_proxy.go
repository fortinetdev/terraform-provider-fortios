// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure LLM Proxy.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLlmProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceLlmProxyUpdate,
		Read:   resourceLlmProxyRead,
		Update: resourceLlmProxyUpdate,
		Delete: resourceLlmProxyDelete,

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
				Optional:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"incoming_http_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"incoming_https_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"incoming_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"incoming_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"ssl_certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"ssl_min_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srv_pool_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"srv_pool_max_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"srv_pool_max_concurrent_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceLlmProxyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLlmProxy(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LlmProxy resource while getting object: %v", err)
	}

	o, err := c.UpdateLlmProxy(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LlmProxy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LlmProxy")
	}

	return resourceLlmProxyRead(d, m)
}

func resourceLlmProxyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLlmProxy(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LlmProxy resource while getting object: %v", err)
	}

	_, err = c.UpdateLlmProxy(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing LlmProxy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLlmProxyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLlmProxy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LlmProxy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLlmProxy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LlmProxy resource from API: %v", err)
	}
	return nil
}

func flattenLlmProxyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmProxyStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmProxyProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmProxyIpv6Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmProxyInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenLlmProxyInterfaceName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenLlmProxyInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmProxyIncomingHttpPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenLlmProxyIncomingHttpsPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenLlmProxyIncomingIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmProxyIncomingIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmProxyHostname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmProxySslCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenLlmProxySslCertificateName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenLlmProxySslCertificateName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmProxySslMinVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmProxySslMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmProxySrvPoolTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenLlmProxySrvPoolMaxRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenLlmProxySrvPoolMaxConcurrentRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectLlmProxy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenLlmProxyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenLlmProxyStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("protocol", flattenLlmProxyProtocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("ipv6_status", flattenLlmProxyIpv6Status(o["ipv6-status"], d, "ipv6_status", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-status"]) {
			return fmt.Errorf("Error reading ipv6_status: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("interface", flattenLlmProxyInterface(o["interface"], d, "interface", sv)); err != nil {
			if !fortiAPIPatch(o["interface"]) {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("interface"); ok {
			if err = d.Set("interface", flattenLlmProxyInterface(o["interface"], d, "interface", sv)); err != nil {
				if !fortiAPIPatch(o["interface"]) {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			}
		}
	}

	if err = d.Set("incoming_http_port", flattenLlmProxyIncomingHttpPort(o["incoming-http-port"], d, "incoming_http_port", sv)); err != nil {
		if !fortiAPIPatch(o["incoming-http-port"]) {
			return fmt.Errorf("Error reading incoming_http_port: %v", err)
		}
	}

	if err = d.Set("incoming_https_port", flattenLlmProxyIncomingHttpsPort(o["incoming-https-port"], d, "incoming_https_port", sv)); err != nil {
		if !fortiAPIPatch(o["incoming-https-port"]) {
			return fmt.Errorf("Error reading incoming_https_port: %v", err)
		}
	}

	if err = d.Set("incoming_ip", flattenLlmProxyIncomingIp(o["incoming-ip"], d, "incoming_ip", sv)); err != nil {
		if !fortiAPIPatch(o["incoming-ip"]) {
			return fmt.Errorf("Error reading incoming_ip: %v", err)
		}
	}

	if err = d.Set("incoming_ip6", flattenLlmProxyIncomingIp6(o["incoming-ip6"], d, "incoming_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["incoming-ip6"]) {
			return fmt.Errorf("Error reading incoming_ip6: %v", err)
		}
	}

	if err = d.Set("hostname", flattenLlmProxyHostname(o["hostname"], d, "hostname", sv)); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ssl_certificate", flattenLlmProxySslCertificate(o["ssl-certificate"], d, "ssl_certificate", sv)); err != nil {
			if !fortiAPIPatch(o["ssl-certificate"]) {
				return fmt.Errorf("Error reading ssl_certificate: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_certificate"); ok {
			if err = d.Set("ssl_certificate", flattenLlmProxySslCertificate(o["ssl-certificate"], d, "ssl_certificate", sv)); err != nil {
				if !fortiAPIPatch(o["ssl-certificate"]) {
					return fmt.Errorf("Error reading ssl_certificate: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_min_version", flattenLlmProxySslMinVersion(o["ssl-min-version"], d, "ssl_min_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-min-version"]) {
			return fmt.Errorf("Error reading ssl_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenLlmProxySslMaxVersion(o["ssl-max-version"], d, "ssl_max_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-max-version"]) {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("srv_pool_ttl", flattenLlmProxySrvPoolTtl(o["srv-pool-ttl"], d, "srv_pool_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["srv-pool-ttl"]) {
			return fmt.Errorf("Error reading srv_pool_ttl: %v", err)
		}
	}

	if err = d.Set("srv_pool_max_request", flattenLlmProxySrvPoolMaxRequest(o["srv-pool-max-request"], d, "srv_pool_max_request", sv)); err != nil {
		if !fortiAPIPatch(o["srv-pool-max-request"]) {
			return fmt.Errorf("Error reading srv_pool_max_request: %v", err)
		}
	}

	if err = d.Set("srv_pool_max_concurrent_request", flattenLlmProxySrvPoolMaxConcurrentRequest(o["srv-pool-max-concurrent-request"], d, "srv_pool_max_concurrent_request", sv)); err != nil {
		if !fortiAPIPatch(o["srv-pool-max-concurrent-request"]) {
			return fmt.Errorf("Error reading srv_pool_max_concurrent_request: %v", err)
		}
	}

	return nil
}

func flattenLlmProxyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLlmProxyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProxyStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProxyProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProxyIpv6Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProxyInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandLlmProxyInterfaceName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLlmProxyInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProxyIncomingHttpPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProxyIncomingHttpsPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProxyIncomingIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProxyIncomingIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProxyHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProxySslCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandLlmProxySslCertificateName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLlmProxySslCertificateName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProxySslMinVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProxySslMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProxySrvPoolTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProxySrvPoolMaxRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProxySrvPoolMaxConcurrentRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLlmProxy(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		if setArgNil {
			obj["name"] = nil
		} else {
			t, err := expandLlmProxyName(d, v, "name", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["name"] = t
			}
		}
	} else if d.HasChange("name") {
		obj["name"] = nil
	}

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandLlmProxyStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("protocol"); ok {
		if setArgNil {
			obj["protocol"] = nil
		} else {
			t, err := expandLlmProxyProtocol(d, v, "protocol", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["protocol"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipv6_status"); ok {
		if setArgNil {
			obj["ipv6-status"] = nil
		} else {
			t, err := expandLlmProxyIpv6Status(d, v, "ipv6_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipv6-status"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		if setArgNil {
			obj["interface"] = make([]struct{}, 0)
		} else {
			t, err := expandLlmProxyInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("incoming_http_port"); ok {
		if setArgNil {
			obj["incoming-http-port"] = nil
		} else {
			t, err := expandLlmProxyIncomingHttpPort(d, v, "incoming_http_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["incoming-http-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("incoming_https_port"); ok {
		if setArgNil {
			obj["incoming-https-port"] = nil
		} else {
			t, err := expandLlmProxyIncomingHttpsPort(d, v, "incoming_https_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["incoming-https-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("incoming_ip"); ok {
		if setArgNil {
			obj["incoming-ip"] = nil
		} else {
			t, err := expandLlmProxyIncomingIp(d, v, "incoming_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["incoming-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("incoming_ip6"); ok {
		if setArgNil {
			obj["incoming-ip6"] = nil
		} else {
			t, err := expandLlmProxyIncomingIp6(d, v, "incoming_ip6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["incoming-ip6"] = t
			}
		}
	}

	if v, ok := d.GetOk("hostname"); ok {
		if setArgNil {
			obj["hostname"] = nil
		} else {
			t, err := expandLlmProxyHostname(d, v, "hostname", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hostname"] = t
			}
		}
	} else if d.HasChange("hostname") {
		obj["hostname"] = nil
	}

	if v, ok := d.GetOk("ssl_certificate"); ok || d.HasChange("ssl_certificate") {
		if setArgNil {
			obj["ssl-certificate"] = make([]struct{}, 0)
		} else {
			t, err := expandLlmProxySslCertificate(d, v, "ssl_certificate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssl-certificate"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssl_min_version"); ok {
		if setArgNil {
			obj["ssl-min-version"] = nil
		} else {
			t, err := expandLlmProxySslMinVersion(d, v, "ssl_min_version", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssl-min-version"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok {
		if setArgNil {
			obj["ssl-max-version"] = nil
		} else {
			t, err := expandLlmProxySslMaxVersion(d, v, "ssl_max_version", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssl-max-version"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("srv_pool_ttl"); ok {
		if setArgNil {
			obj["srv-pool-ttl"] = nil
		} else {
			t, err := expandLlmProxySrvPoolTtl(d, v, "srv_pool_ttl", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["srv-pool-ttl"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("srv_pool_max_request"); ok {
		if setArgNil {
			obj["srv-pool-max-request"] = nil
		} else {
			t, err := expandLlmProxySrvPoolMaxRequest(d, v, "srv_pool_max_request", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["srv-pool-max-request"] = t
			}
		}
	} else if d.HasChange("srv_pool_max_request") {
		obj["srv-pool-max-request"] = nil
	}

	if v, ok := d.GetOkExists("srv_pool_max_concurrent_request"); ok {
		if setArgNil {
			obj["srv-pool-max-concurrent-request"] = nil
		} else {
			t, err := expandLlmProxySrvPoolMaxConcurrentRequest(d, v, "srv_pool_max_concurrent_request", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["srv-pool-max-concurrent-request"] = t
			}
		}
	} else if d.HasChange("srv_pool_max_concurrent_request") {
		obj["srv-pool-max-concurrent-request"] = nil
	}

	return &obj, nil
}
