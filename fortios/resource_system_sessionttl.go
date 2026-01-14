// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure global session TTL timers for this FortiGate.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSessionTtl() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSessionTtlUpdate,
		Read:   resourceSystemSessionTtlRead,
		Update: resourceSystemSessionTtlUpdate,
		Delete: resourceSystemSessionTtlDelete,

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
			"default": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"protocol": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"start_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"end_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"timeout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"refresh_direction": &schema.Schema{
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

func resourceSystemSessionTtlUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSessionTtl(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSessionTtl resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSessionTtl(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemSessionTtl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSessionTtl")
	}

	return resourceSystemSessionTtlRead(d, m)
}

func resourceSystemSessionTtlDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemSessionTtl(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemSessionTtl resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSessionTtl(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemSessionTtl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSessionTtlRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSessionTtl(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemSessionTtl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSessionTtl(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSessionTtl resource from API: %v", err)
	}
	return nil
}

func flattenSystemSessionTtlDefault(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSessionTtlPort(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemSessionTtlPortId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["protocol"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
			}
			tmp["protocol"] = flattenSystemSessionTtlPortProtocol(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["start-port"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
			}
			tmp["start_port"] = flattenSystemSessionTtlPortStartPort(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["end-port"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
			}
			tmp["end_port"] = flattenSystemSessionTtlPortEndPort(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["timeout"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "timeout"
			}
			tmp["timeout"] = flattenSystemSessionTtlPortTimeout(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["refresh-direction"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "refresh_direction"
			}
			tmp["refresh_direction"] = flattenSystemSessionTtlPortRefreshDirection(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemSessionTtlPortId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSessionTtlPortProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSessionTtlPortStartPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSessionTtlPortEndPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSessionTtlPortTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSessionTtlPortRefreshDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSessionTtl(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("default", flattenSystemSessionTtlDefault(o["default"], d, "default", sv)); err != nil {
		if !fortiAPIPatch(o["default"]) {
			return fmt.Errorf("Error reading default: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("port", flattenSystemSessionTtlPort(o["port"], d, "port", sv)); err != nil {
			if !fortiAPIPatch(o["port"]) {
				return fmt.Errorf("Error reading port: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("port"); ok {
			if err = d.Set("port", flattenSystemSessionTtlPort(o["port"], d, "port", sv)); err != nil {
				if !fortiAPIPatch(o["port"]) {
					return fmt.Errorf("Error reading port: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemSessionTtlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSessionTtlDefault(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSessionTtlPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemSessionTtlPortId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandSystemSessionTtlPortProtocol(d, i["protocol"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["protocol"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-port"], _ = expandSystemSessionTtlPortStartPort(d, i["start_port"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["start-port"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-port"], _ = expandSystemSessionTtlPortEndPort(d, i["end_port"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["end-port"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "timeout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["timeout"], _ = expandSystemSessionTtlPortTimeout(d, i["timeout"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["timeout"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "refresh_direction"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["refresh-direction"], _ = expandSystemSessionTtlPortRefreshDirection(d, i["refresh_direction"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSessionTtlPortId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSessionTtlPortProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSessionTtlPortStartPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSessionTtlPortEndPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSessionTtlPortTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSessionTtlPortRefreshDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSessionTtl(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default"); ok {
		if setArgNil {
			obj["default"] = nil
		} else {
			t, err := expandSystemSessionTtlDefault(d, v, "default", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default"] = t
			}
		}
	} else if d.HasChange("default") {
		obj["default"] = nil
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		if setArgNil {
			obj["port"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemSessionTtlPort(d, v, "port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["port"] = t
			}
		}
	}

	return &obj, nil
}
