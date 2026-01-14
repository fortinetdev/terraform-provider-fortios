// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure ZTNA traffic forward proxy reverse service.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaTrafficForwardProxyReverseService() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaTrafficForwardProxyReverseServiceUpdate,
		Read:   resourceZtnaTrafficForwardProxyReverseServiceRead,
		Update: resourceZtnaTrafficForwardProxyReverseServiceUpdate,
		Delete: resourceZtnaTrafficForwardProxyReverseServiceDelete,

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
			"remote_servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"address": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
						"health_check_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 600),
							Optional:     true,
							Computed:     true,
						},
						"ssl_max_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"certificate": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"trusted_server_ca": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
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

func resourceZtnaTrafficForwardProxyReverseServiceUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectZtnaTrafficForwardProxyReverseService(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaTrafficForwardProxyReverseService resource while getting object: %v", err)
	}

	o, err := c.UpdateZtnaTrafficForwardProxyReverseService(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaTrafficForwardProxyReverseService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ZtnaTrafficForwardProxyReverseService")
	}

	return resourceZtnaTrafficForwardProxyReverseServiceRead(d, m)
}

func resourceZtnaTrafficForwardProxyReverseServiceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectZtnaTrafficForwardProxyReverseService(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating ZtnaTrafficForwardProxyReverseService resource while getting object: %v", err)
	}

	_, err = c.UpdateZtnaTrafficForwardProxyReverseService(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing ZtnaTrafficForwardProxyReverseService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaTrafficForwardProxyReverseServiceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadZtnaTrafficForwardProxyReverseService(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaTrafficForwardProxyReverseService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaTrafficForwardProxyReverseService(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaTrafficForwardProxyReverseService resource from API: %v", err)
	}
	return nil
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenZtnaTrafficForwardProxyReverseServiceRemoteServersName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["status"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
			}
			tmp["status"] = flattenZtnaTrafficForwardProxyReverseServiceRemoteServersStatus(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["address"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
			}
			tmp["address"] = flattenZtnaTrafficForwardProxyReverseServiceRemoteServersAddress(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["health-check-interval"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_interval"
			}
			tmp["health_check_interval"] = flattenZtnaTrafficForwardProxyReverseServiceRemoteServersHealthCheckInterval(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ssl-max-version"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
			}
			tmp["ssl_max_version"] = flattenZtnaTrafficForwardProxyReverseServiceRemoteServersSslMaxVersion(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["port"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
			}
			tmp["port"] = flattenZtnaTrafficForwardProxyReverseServiceRemoteServersPort(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["certificate"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "certificate"
			}
			tmp["certificate"] = flattenZtnaTrafficForwardProxyReverseServiceRemoteServersCertificate(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["trusted-server-ca"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "trusted_server_ca"
			}
			tmp["trusted_server_ca"] = flattenZtnaTrafficForwardProxyReverseServiceRemoteServersTrustedServerCa(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersHealthCheckInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersSslMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyReverseServiceRemoteServersTrustedServerCa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectZtnaTrafficForwardProxyReverseService(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if b_get_all_tables {
		if err = d.Set("remote_servers", flattenZtnaTrafficForwardProxyReverseServiceRemoteServers(o["remote-servers"], d, "remote_servers", sv)); err != nil {
			if !fortiAPIPatch(o["remote-servers"]) {
				return fmt.Errorf("Error reading remote_servers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("remote_servers"); ok {
			if err = d.Set("remote_servers", flattenZtnaTrafficForwardProxyReverseServiceRemoteServers(o["remote-servers"], d, "remote_servers", sv)); err != nil {
				if !fortiAPIPatch(o["remote-servers"]) {
					return fmt.Errorf("Error reading remote_servers: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenZtnaTrafficForwardProxyReverseServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandZtnaTrafficForwardProxyReverseServiceRemoteServersName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandZtnaTrafficForwardProxyReverseServiceRemoteServersStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["address"], _ = expandZtnaTrafficForwardProxyReverseServiceRemoteServersAddress(d, i["address"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["address"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["health-check-interval"], _ = expandZtnaTrafficForwardProxyReverseServiceRemoteServersHealthCheckInterval(d, i["health_check_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-max-version"], _ = expandZtnaTrafficForwardProxyReverseServiceRemoteServersSslMaxVersion(d, i["ssl_max_version"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandZtnaTrafficForwardProxyReverseServiceRemoteServersPort(d, i["port"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["port"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "certificate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["certificate"], _ = expandZtnaTrafficForwardProxyReverseServiceRemoteServersCertificate(d, i["certificate"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["certificate"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trusted_server_ca"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["trusted-server-ca"], _ = expandZtnaTrafficForwardProxyReverseServiceRemoteServersTrustedServerCa(d, i["trusted_server_ca"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["trusted-server-ca"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersHealthCheckInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersSslMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyReverseServiceRemoteServersTrustedServerCa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectZtnaTrafficForwardProxyReverseService(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("remote_servers"); ok || d.HasChange("remote_servers") {
		if setArgNil {
			obj["remote-servers"] = make([]struct{}, 0)
		} else {
			t, err := expandZtnaTrafficForwardProxyReverseServiceRemoteServers(d, v, "remote_servers", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["remote-servers"] = t
			}
		}
	}

	return &obj, nil
}
