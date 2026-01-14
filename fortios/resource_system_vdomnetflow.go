// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure NetFlow per VDOM.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemVdomNetflow() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVdomNetflowUpdate,
		Read:   resourceSystemVdomNetflowRead,
		Update: resourceSystemVdomNetflowUpdate,
		Delete: resourceSystemVdomNetflowDelete,

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
			"vdom_netflow": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"collectors": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 6),
							Optional:     true,
						},
						"collector_ip": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
						"collector_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"source_ip": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
						"source_ip_interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
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
							ValidateFunc: validation.IntBetween(0, 511),
							Optional:     true,
						},
					},
				},
			},
			"collector_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"collector_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceSystemVdomNetflowUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemVdomNetflow(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomNetflow resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemVdomNetflow(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomNetflow resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVdomNetflow")
	}

	return resourceSystemVdomNetflowRead(d, m)
}

func resourceSystemVdomNetflowDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemVdomNetflow(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemVdomNetflow resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemVdomNetflow(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemVdomNetflow resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomNetflowRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemVdomNetflow(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomNetflow resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdomNetflow(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomNetflow resource from API: %v", err)
	}
	return nil
}

func flattenSystemVdomNetflowVdomNetflow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomNetflowCollectors(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemVdomNetflowCollectorsId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["collector-ip"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "collector_ip"
			}
			tmp["collector_ip"] = flattenSystemVdomNetflowCollectorsCollectorIp(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["collector-port"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "collector_port"
			}
			tmp["collector_port"] = flattenSystemVdomNetflowCollectorsCollectorPort(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["source-ip"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
			}
			tmp["source_ip"] = flattenSystemVdomNetflowCollectorsSourceIp(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["source-ip-interface"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip_interface"
			}
			tmp["source_ip_interface"] = flattenSystemVdomNetflowCollectorsSourceIpInterface(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["interface-select-method"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
			}
			tmp["interface_select_method"] = flattenSystemVdomNetflowCollectorsInterfaceSelectMethod(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["interface"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
			}
			tmp["interface"] = flattenSystemVdomNetflowCollectorsInterface(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["vrf-select"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
			}
			tmp["vrf_select"] = flattenSystemVdomNetflowCollectorsVrfSelect(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemVdomNetflowCollectorsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVdomNetflowCollectorsCollectorIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomNetflowCollectorsCollectorPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVdomNetflowCollectorsSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomNetflowCollectorsSourceIpInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomNetflowCollectorsInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomNetflowCollectorsInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomNetflowCollectorsVrfSelect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVdomNetflowCollectorIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomNetflowCollectorPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVdomNetflowSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomNetflowInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomNetflowInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemVdomNetflow(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("vdom_netflow", flattenSystemVdomNetflowVdomNetflow(o["vdom-netflow"], d, "vdom_netflow", sv)); err != nil {
		if !fortiAPIPatch(o["vdom-netflow"]) {
			return fmt.Errorf("Error reading vdom_netflow: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("collectors", flattenSystemVdomNetflowCollectors(o["collectors"], d, "collectors", sv)); err != nil {
			if !fortiAPIPatch(o["collectors"]) {
				return fmt.Errorf("Error reading collectors: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("collectors"); ok {
			if err = d.Set("collectors", flattenSystemVdomNetflowCollectors(o["collectors"], d, "collectors", sv)); err != nil {
				if !fortiAPIPatch(o["collectors"]) {
					return fmt.Errorf("Error reading collectors: %v", err)
				}
			}
		}
	}

	if err = d.Set("collector_ip", flattenSystemVdomNetflowCollectorIp(o["collector-ip"], d, "collector_ip", sv)); err != nil {
		if !fortiAPIPatch(o["collector-ip"]) {
			return fmt.Errorf("Error reading collector_ip: %v", err)
		}
	}

	if err = d.Set("collector_port", flattenSystemVdomNetflowCollectorPort(o["collector-port"], d, "collector_port", sv)); err != nil {
		if !fortiAPIPatch(o["collector-port"]) {
			return fmt.Errorf("Error reading collector_port: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemVdomNetflowSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenSystemVdomNetflowInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemVdomNetflowInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func flattenSystemVdomNetflowFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemVdomNetflowVdomNetflow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowCollectors(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemVdomNetflowCollectorsId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "collector_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["collector-ip"], _ = expandSystemVdomNetflowCollectorsCollectorIp(d, i["collector_ip"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["collector-ip"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "collector_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["collector-port"], _ = expandSystemVdomNetflowCollectorsCollectorPort(d, i["collector_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["source-ip"], _ = expandSystemVdomNetflowCollectorsSourceIp(d, i["source_ip"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["source-ip"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip_interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["source-ip-interface"], _ = expandSystemVdomNetflowCollectorsSourceIpInterface(d, i["source_ip_interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["source-ip-interface"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface-select-method"], _ = expandSystemVdomNetflowCollectorsInterfaceSelectMethod(d, i["interface_select_method"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandSystemVdomNetflowCollectorsInterface(d, i["interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["interface"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrf-select"], _ = expandSystemVdomNetflowCollectorsVrfSelect(d, i["vrf_select"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vrf-select"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVdomNetflowCollectorsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowCollectorsCollectorIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowCollectorsCollectorPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowCollectorsSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowCollectorsSourceIpInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowCollectorsInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowCollectorsInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowCollectorsVrfSelect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowCollectorIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowCollectorPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVdomNetflow(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("vdom_netflow"); ok {
		if setArgNil {
			obj["vdom-netflow"] = nil
		} else {
			t, err := expandSystemVdomNetflowVdomNetflow(d, v, "vdom_netflow", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vdom-netflow"] = t
			}
		}
	}

	if v, ok := d.GetOk("collectors"); ok || d.HasChange("collectors") {
		if setArgNil {
			obj["collectors"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemVdomNetflowCollectors(d, v, "collectors", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["collectors"] = t
			}
		}
	}

	if v, ok := d.GetOk("collector_ip"); ok {
		if setArgNil {
			obj["collector-ip"] = nil
		} else {
			t, err := expandSystemVdomNetflowCollectorIp(d, v, "collector_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["collector-ip"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("collector_port"); ok {
		if setArgNil {
			obj["collector-port"] = nil
		} else {
			t, err := expandSystemVdomNetflowCollectorPort(d, v, "collector_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["collector-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		if setArgNil {
			obj["source-ip"] = nil
		} else {
			t, err := expandSystemVdomNetflowSourceIp(d, v, "source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		if setArgNil {
			obj["interface-select-method"] = nil
		} else {
			t, err := expandSystemVdomNetflowInterfaceSelectMethod(d, v, "interface_select_method", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface-select-method"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		if setArgNil {
			obj["interface"] = nil
		} else {
			t, err := expandSystemVdomNetflowInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	return &obj, nil
}
