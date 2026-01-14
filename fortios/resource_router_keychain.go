// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure key-chain.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterKeyChain() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterKeyChainCreate,
		Read:   resourceRouterKeyChainRead,
		Update: resourceRouterKeyChainUpdate,
		Delete: resourceRouterKeyChainDelete,

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
				Required:     true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"accept_lifetime": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"send_lifetime": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"key_string": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 60),
							Optional:     true,
							Sensitive:    true,
						},
						"algorithm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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

func resourceRouterKeyChainCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterKeyChain(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterKeyChain resource while getting object: %v", err)
	}

	o, err := c.CreateRouterKeyChain(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating RouterKeyChain resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterKeyChain")
	}

	return resourceRouterKeyChainRead(d, m)
}

func resourceRouterKeyChainUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterKeyChain(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterKeyChain resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterKeyChain(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating RouterKeyChain resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterKeyChain")
	}

	return resourceRouterKeyChainRead(d, m)
}

func resourceRouterKeyChainDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteRouterKeyChain(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting RouterKeyChain resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterKeyChainRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterKeyChain(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterKeyChain resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterKeyChain(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterKeyChain resource from API: %v", err)
	}
	return nil
}

func flattenRouterKeyChainName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterKeyChainKey(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenRouterKeyChainKeyId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["accept-lifetime"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_lifetime"
			}
			tmp["accept_lifetime"] = flattenRouterKeyChainKeyAcceptLifetime(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["send-lifetime"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "send_lifetime"
			}
			tmp["send_lifetime"] = flattenRouterKeyChainKeySendLifetime(cur_v, d, pre_append, sv)
		}

		if _, ok := i["key-string"]; ok {
			if tf_exist {
				pre_append := pre + "." + strconv.Itoa(con) + "." + "key_string"
				c := d.Get(pre_append).(string)
				if c != "" {
					tmp["key_string"] = c
				}
			}
		}

		if cur_v, ok := i["algorithm"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "algorithm"
			}
			tmp["algorithm"] = flattenRouterKeyChainKeyAlgorithm(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterKeyChainKeyId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterKeyChainKeyAcceptLifetime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterKeyChainKeySendLifetime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterKeyChainKeyAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterKeyChain(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenRouterKeyChainName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("key", flattenRouterKeyChainKey(o["key"], d, "key", sv)); err != nil {
			if !fortiAPIPatch(o["key"]) {
				return fmt.Errorf("Error reading key: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("key"); ok {
			if err = d.Set("key", flattenRouterKeyChainKey(o["key"], d, "key", sv)); err != nil {
				if !fortiAPIPatch(o["key"]) {
					return fmt.Errorf("Error reading key: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterKeyChainFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterKeyChainName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterKeyChainKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandRouterKeyChainKeyId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_lifetime"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["accept-lifetime"], _ = expandRouterKeyChainKeyAcceptLifetime(d, i["accept_lifetime"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["accept-lifetime"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_lifetime"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["send-lifetime"], _ = expandRouterKeyChainKeySendLifetime(d, i["send_lifetime"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["send-lifetime"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["key-string"], _ = expandRouterKeyChainKeyKeyString(d, i["key_string"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["key-string"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "algorithm"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["algorithm"], _ = expandRouterKeyChainKeyAlgorithm(d, i["algorithm"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterKeyChainKeyId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return convintf2i(v), nil
}

func expandRouterKeyChainKeyAcceptLifetime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterKeyChainKeySendLifetime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterKeyChainKeyKeyString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterKeyChainKeyAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterKeyChain(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandRouterKeyChainName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok || d.HasChange("key") {
		t, err := expandRouterKeyChainKey(d, v, "key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	return &obj, nil
}
