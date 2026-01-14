// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiGuard Internet Services.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallInternetServiceFortiguard() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInternetServiceFortiguardCreate,
		Read:   resourceFirewallInternetServiceFortiguardRead,
		Update: resourceFirewallInternetServiceFortiguardUpdate,
		Delete: resourceFirewallInternetServiceFortiguardDelete,

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
				Optional:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"entry": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"addr_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"port_range": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"start_port": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
										Computed:     true,
									},
									"end_port": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"dst": &schema.Schema{
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
						"dst6": &schema.Schema{
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

func resourceFirewallInternetServiceFortiguardCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallInternetServiceFortiguard(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceFortiguard resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallInternetServiceFortiguard(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceFortiguard resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallInternetServiceFortiguard")
	}

	return resourceFirewallInternetServiceFortiguardRead(d, m)
}

func resourceFirewallInternetServiceFortiguardUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallInternetServiceFortiguard(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceFortiguard resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallInternetServiceFortiguard(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceFortiguard resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallInternetServiceFortiguard")
	}

	return resourceFirewallInternetServiceFortiguardRead(d, m)
}

func resourceFirewallInternetServiceFortiguardDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallInternetServiceFortiguard(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInternetServiceFortiguard resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceFortiguardRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallInternetServiceFortiguard(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceFortiguard resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInternetServiceFortiguard(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceFortiguard resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInternetServiceFortiguardName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceFortiguardComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceFortiguardEntry(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallInternetServiceFortiguardEntryId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["addr-mode"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
			}
			tmp["addr_mode"] = flattenFirewallInternetServiceFortiguardEntryAddrMode(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["protocol"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
			}
			tmp["protocol"] = flattenFirewallInternetServiceFortiguardEntryProtocol(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["port-range"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "port_range"
			}
			tmp["port_range"] = flattenFirewallInternetServiceFortiguardEntryPortRange(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dst"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
			}
			tmp["dst"] = flattenFirewallInternetServiceFortiguardEntryDst(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dst6"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
			}
			tmp["dst6"] = flattenFirewallInternetServiceFortiguardEntryDst6(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallInternetServiceFortiguardEntryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceFortiguardEntryAddrMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceFortiguardEntryProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceFortiguardEntryPortRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallInternetServiceFortiguardEntryPortRangeId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["start-port"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
			}
			tmp["start_port"] = flattenFirewallInternetServiceFortiguardEntryPortRangeStartPort(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["end-port"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
			}
			tmp["end_port"] = flattenFirewallInternetServiceFortiguardEntryPortRangeEndPort(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallInternetServiceFortiguardEntryPortRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceFortiguardEntryPortRangeStartPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceFortiguardEntryPortRangeEndPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceFortiguardEntryDst(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallInternetServiceFortiguardEntryDstName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallInternetServiceFortiguardEntryDstName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceFortiguardEntryDst6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallInternetServiceFortiguardEntryDst6Name(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallInternetServiceFortiguardEntryDst6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallInternetServiceFortiguard(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenFirewallInternetServiceFortiguardName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallInternetServiceFortiguardComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("entry", flattenFirewallInternetServiceFortiguardEntry(o["entry"], d, "entry", sv)); err != nil {
			if !fortiAPIPatch(o["entry"]) {
				return fmt.Errorf("Error reading entry: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entry"); ok {
			if err = d.Set("entry", flattenFirewallInternetServiceFortiguardEntry(o["entry"], d, "entry", sv)); err != nil {
				if !fortiAPIPatch(o["entry"]) {
					return fmt.Errorf("Error reading entry: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFirewallInternetServiceFortiguardFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallInternetServiceFortiguardName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceFortiguardComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceFortiguardEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallInternetServiceFortiguardEntryId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["addr-mode"], _ = expandFirewallInternetServiceFortiguardEntryAddrMode(d, i["addr_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandFirewallInternetServiceFortiguardEntryProtocol(d, i["protocol"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["protocol"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_range"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port-range"], _ = expandFirewallInternetServiceFortiguardEntryPortRange(d, i["port_range"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["port-range"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst"], _ = expandFirewallInternetServiceFortiguardEntryDst(d, i["dst"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dst"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst6"], _ = expandFirewallInternetServiceFortiguardEntryDst6(d, i["dst6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dst6"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceFortiguardEntryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceFortiguardEntryAddrMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceFortiguardEntryProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceFortiguardEntryPortRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallInternetServiceFortiguardEntryPortRangeId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-port"], _ = expandFirewallInternetServiceFortiguardEntryPortRangeStartPort(d, i["start_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-port"], _ = expandFirewallInternetServiceFortiguardEntryPortRangeEndPort(d, i["end_port"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceFortiguardEntryPortRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceFortiguardEntryPortRangeStartPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceFortiguardEntryPortRangeEndPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceFortiguardEntryDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallInternetServiceFortiguardEntryDstName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceFortiguardEntryDstName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceFortiguardEntryDst6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallInternetServiceFortiguardEntryDst6Name(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceFortiguardEntryDst6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInternetServiceFortiguard(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallInternetServiceFortiguardName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	} else if d.HasChange("name") {
		obj["name"] = nil
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandFirewallInternetServiceFortiguardComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("entry"); ok || d.HasChange("entry") {
		t, err := expandFirewallInternetServiceFortiguardEntry(d, v, "entry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entry"] = t
		}
	}

	return &obj, nil
}
