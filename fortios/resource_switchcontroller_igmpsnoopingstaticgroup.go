// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch IGMP snooping static group settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerIgmpSnoopingStaticGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerIgmpSnoopingStaticGroupCreate,
		Read:   resourceSwitchControllerIgmpSnoopingStaticGroupRead,
		Update: resourceSwitchControllerIgmpSnoopingStaticGroupUpdate,
		Delete: resourceSwitchControllerIgmpSnoopingStaticGroupDelete,

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
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"ignore_reports": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mcast_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"switch_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"ports": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"igmp_port_name": &schema.Schema{
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

func resourceSwitchControllerIgmpSnoopingStaticGroupCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerIgmpSnoopingStaticGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerIgmpSnoopingStaticGroup resource while getting object: %v", err)
	}

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSwitchControllerIgmpSnoopingStaticGroup(mkey, vdomparam)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSwitchControllerIgmpSnoopingStaticGroup(obj, mkey, vdomparam)
			if err != nil {
				return fmt.Errorf("Error updating SwitchControllerIgmpSnoopingStaticGroup resource: %v", err)
			}
		}
	}

	if !existing {
		o, err = c.CreateSwitchControllerIgmpSnoopingStaticGroup(obj, vdomparam)

		if err != nil {
			return fmt.Errorf("Error creating SwitchControllerIgmpSnoopingStaticGroup resource: %v", err)
		}
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerIgmpSnoopingStaticGroup")
	}

	return resourceSwitchControllerIgmpSnoopingStaticGroupRead(d, m)
}

func resourceSwitchControllerIgmpSnoopingStaticGroupUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerIgmpSnoopingStaticGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerIgmpSnoopingStaticGroup resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerIgmpSnoopingStaticGroup(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerIgmpSnoopingStaticGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerIgmpSnoopingStaticGroup")
	}

	return resourceSwitchControllerIgmpSnoopingStaticGroupRead(d, m)
}

func resourceSwitchControllerIgmpSnoopingStaticGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerIgmpSnoopingStaticGroup(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerIgmpSnoopingStaticGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerIgmpSnoopingStaticGroupRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerIgmpSnoopingStaticGroup(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerIgmpSnoopingStaticGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerIgmpSnoopingStaticGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerIgmpSnoopingStaticGroup resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerIgmpSnoopingStaticGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerIgmpSnoopingStaticGroupDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerIgmpSnoopingStaticGroupIgnoreReports(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerIgmpSnoopingStaticGroupMcastAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerIgmpSnoopingStaticGroupVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerIgmpSnoopingStaticGroupPorts(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSwitchControllerIgmpSnoopingStaticGroupPortsId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["switch-id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_id"
			}
			tmp["switch_id"] = flattenSwitchControllerIgmpSnoopingStaticGroupPortsSwitchId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ports"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ports"
			}
			tmp["ports"] = flattenSwitchControllerIgmpSnoopingStaticGroupPortsPorts(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSwitchControllerIgmpSnoopingStaticGroupPortsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSwitchControllerIgmpSnoopingStaticGroupPortsSwitchId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerIgmpSnoopingStaticGroupPortsPorts(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "igmp-port-name", "igmp_port_name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["igmp-port-name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp_port_name"
			}
			tmp["igmp_port_name"] = flattenSwitchControllerIgmpSnoopingStaticGroupPortsPortsIgmpPortName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "igmp_port_name", d)
	return result
}

func flattenSwitchControllerIgmpSnoopingStaticGroupPortsPortsIgmpPortName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerIgmpSnoopingStaticGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenSwitchControllerIgmpSnoopingStaticGroupName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerIgmpSnoopingStaticGroupDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("ignore_reports", flattenSwitchControllerIgmpSnoopingStaticGroupIgnoreReports(o["ignore-reports"], d, "ignore_reports", sv)); err != nil {
		if !fortiAPIPatch(o["ignore-reports"]) {
			return fmt.Errorf("Error reading ignore_reports: %v", err)
		}
	}

	if err = d.Set("mcast_addr", flattenSwitchControllerIgmpSnoopingStaticGroupMcastAddr(o["mcast-addr"], d, "mcast_addr", sv)); err != nil {
		if !fortiAPIPatch(o["mcast-addr"]) {
			return fmt.Errorf("Error reading mcast_addr: %v", err)
		}
	}

	if err = d.Set("vlan", flattenSwitchControllerIgmpSnoopingStaticGroupVlan(o["vlan"], d, "vlan", sv)); err != nil {
		if !fortiAPIPatch(o["vlan"]) {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ports", flattenSwitchControllerIgmpSnoopingStaticGroupPorts(o["ports"], d, "ports", sv)); err != nil {
			if !fortiAPIPatch(o["ports"]) {
				return fmt.Errorf("Error reading ports: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ports"); ok {
			if err = d.Set("ports", flattenSwitchControllerIgmpSnoopingStaticGroupPorts(o["ports"], d, "ports", sv)); err != nil {
				if !fortiAPIPatch(o["ports"]) {
					return fmt.Errorf("Error reading ports: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerIgmpSnoopingStaticGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerIgmpSnoopingStaticGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerIgmpSnoopingStaticGroupDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerIgmpSnoopingStaticGroupIgnoreReports(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerIgmpSnoopingStaticGroupMcastAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerIgmpSnoopingStaticGroupVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerIgmpSnoopingStaticGroupPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSwitchControllerIgmpSnoopingStaticGroupPortsId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["switch-id"], _ = expandSwitchControllerIgmpSnoopingStaticGroupPortsSwitchId(d, i["switch_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["switch-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ports"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ports"], _ = expandSwitchControllerIgmpSnoopingStaticGroupPortsPorts(d, i["ports"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ports"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerIgmpSnoopingStaticGroupPortsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerIgmpSnoopingStaticGroupPortsSwitchId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerIgmpSnoopingStaticGroupPortsPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["igmp-port-name"], _ = expandSwitchControllerIgmpSnoopingStaticGroupPortsPortsIgmpPortName(d, i["igmp_port_name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerIgmpSnoopingStaticGroupPortsPortsIgmpPortName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerIgmpSnoopingStaticGroup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerIgmpSnoopingStaticGroupName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandSwitchControllerIgmpSnoopingStaticGroupDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	} else if d.HasChange("description") {
		obj["description"] = nil
	}

	if v, ok := d.GetOk("ignore_reports"); ok {
		t, err := expandSwitchControllerIgmpSnoopingStaticGroupIgnoreReports(d, v, "ignore_reports", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-reports"] = t
		}
	}

	if v, ok := d.GetOk("mcast_addr"); ok {
		t, err := expandSwitchControllerIgmpSnoopingStaticGroupMcastAddr(d, v, "mcast_addr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mcast-addr"] = t
		}
	}

	if v, ok := d.GetOk("vlan"); ok {
		t, err := expandSwitchControllerIgmpSnoopingStaticGroupVlan(d, v, "vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	} else if d.HasChange("vlan") {
		obj["vlan"] = nil
	}

	if v, ok := d.GetOk("ports"); ok || d.HasChange("ports") {
		t, err := expandSwitchControllerIgmpSnoopingStaticGroupPorts(d, v, "ports", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports"] = t
		}
	}

	return &obj, nil
}
