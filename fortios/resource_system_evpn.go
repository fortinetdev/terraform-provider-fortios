// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure EVPN instance.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemEvpn() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemEvpnCreate,
		Read:   resourceSystemEvpnRead,
		Update: resourceSystemEvpnUpdate,
		Delete: resourceSystemEvpnDelete,

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
			"fosid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rd": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"import_rt": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"route_target": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"export_rt": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"route_target": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"ip_local_learning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arp_suppression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adv_default_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"distribute_local_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtual_mac_vrid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
			},
			"l3_instance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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

func resourceSystemEvpnCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemEvpn(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemEvpn resource while getting object: %v", err)
	}

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemEvpn(mkey, vdomparam)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemEvpn(obj, mkey, vdomparam)
			if err != nil {
				return fmt.Errorf("Error updating SystemEvpn resource: %v", err)
			}
		}
	}

	if !existing {
		o, err = c.CreateSystemEvpn(obj, vdomparam)

		if err != nil {
			return fmt.Errorf("Error creating SystemEvpn resource: %v", err)
		}
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemEvpn")
	}

	return resourceSystemEvpnRead(d, m)
}

func resourceSystemEvpnUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemEvpn(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemEvpn resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemEvpn(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemEvpn resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemEvpn")
	}

	return resourceSystemEvpnRead(d, m)
}

func resourceSystemEvpnDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemEvpn(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemEvpn resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemEvpnRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemEvpn(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemEvpn resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemEvpn(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemEvpn resource from API: %v", err)
	}
	return nil
}

func flattenSystemEvpnId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemEvpnType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemEvpnRd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemEvpnImportRt(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "route-target", "route_target")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["route-target"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "route_target"
			}
			tmp["route_target"] = flattenSystemEvpnImportRtRouteTarget(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "route_target", d)
	return result
}

func flattenSystemEvpnImportRtRouteTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemEvpnExportRt(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "route-target", "route_target")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["route-target"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "route_target"
			}
			tmp["route_target"] = flattenSystemEvpnExportRtRouteTarget(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "route_target", d)
	return result
}

func flattenSystemEvpnExportRtRouteTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemEvpnIpLocalLearning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemEvpnArpSuppression(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemEvpnAdvDefaultGw(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemEvpnDistributeLocalRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemEvpnVirtualMacVrid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemEvpnL3Instance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemEvpnInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemEvpn(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("fosid", flattenSystemEvpnId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemEvpnType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("rd", flattenSystemEvpnRd(o["rd"], d, "rd", sv)); err != nil {
		if !fortiAPIPatch(o["rd"]) {
			return fmt.Errorf("Error reading rd: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("import_rt", flattenSystemEvpnImportRt(o["import-rt"], d, "import_rt", sv)); err != nil {
			if !fortiAPIPatch(o["import-rt"]) {
				return fmt.Errorf("Error reading import_rt: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("import_rt"); ok {
			if err = d.Set("import_rt", flattenSystemEvpnImportRt(o["import-rt"], d, "import_rt", sv)); err != nil {
				if !fortiAPIPatch(o["import-rt"]) {
					return fmt.Errorf("Error reading import_rt: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("export_rt", flattenSystemEvpnExportRt(o["export-rt"], d, "export_rt", sv)); err != nil {
			if !fortiAPIPatch(o["export-rt"]) {
				return fmt.Errorf("Error reading export_rt: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("export_rt"); ok {
			if err = d.Set("export_rt", flattenSystemEvpnExportRt(o["export-rt"], d, "export_rt", sv)); err != nil {
				if !fortiAPIPatch(o["export-rt"]) {
					return fmt.Errorf("Error reading export_rt: %v", err)
				}
			}
		}
	}

	if err = d.Set("ip_local_learning", flattenSystemEvpnIpLocalLearning(o["ip-local-learning"], d, "ip_local_learning", sv)); err != nil {
		if !fortiAPIPatch(o["ip-local-learning"]) {
			return fmt.Errorf("Error reading ip_local_learning: %v", err)
		}
	}

	if err = d.Set("arp_suppression", flattenSystemEvpnArpSuppression(o["arp-suppression"], d, "arp_suppression", sv)); err != nil {
		if !fortiAPIPatch(o["arp-suppression"]) {
			return fmt.Errorf("Error reading arp_suppression: %v", err)
		}
	}

	if err = d.Set("adv_default_gw", flattenSystemEvpnAdvDefaultGw(o["adv-default-gw"], d, "adv_default_gw", sv)); err != nil {
		if !fortiAPIPatch(o["adv-default-gw"]) {
			return fmt.Errorf("Error reading adv_default_gw: %v", err)
		}
	}

	if err = d.Set("distribute_local_route", flattenSystemEvpnDistributeLocalRoute(o["distribute-local-route"], d, "distribute_local_route", sv)); err != nil {
		if !fortiAPIPatch(o["distribute-local-route"]) {
			return fmt.Errorf("Error reading distribute_local_route: %v", err)
		}
	}

	if err = d.Set("virtual_mac_vrid", flattenSystemEvpnVirtualMacVrid(o["virtual-mac-vrid"], d, "virtual_mac_vrid", sv)); err != nil {
		if !fortiAPIPatch(o["virtual-mac-vrid"]) {
			return fmt.Errorf("Error reading virtual_mac_vrid: %v", err)
		}
	}

	if err = d.Set("l3_instance", flattenSystemEvpnL3Instance(o["l3-instance"], d, "l3_instance", sv)); err != nil {
		if !fortiAPIPatch(o["l3-instance"]) {
			return fmt.Errorf("Error reading l3_instance: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemEvpnInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func flattenSystemEvpnFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemEvpnId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEvpnType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEvpnRd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEvpnImportRt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["route-target"], _ = expandSystemEvpnImportRtRouteTarget(d, i["route_target"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemEvpnImportRtRouteTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEvpnExportRt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["route-target"], _ = expandSystemEvpnExportRtRouteTarget(d, i["route_target"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemEvpnExportRtRouteTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEvpnIpLocalLearning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEvpnArpSuppression(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEvpnAdvDefaultGw(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEvpnDistributeLocalRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEvpnVirtualMacVrid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEvpnL3Instance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEvpnInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemEvpn(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandSystemEvpnId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemEvpnType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("rd"); ok {
		t, err := expandSystemEvpnRd(d, v, "rd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rd"] = t
		}
	} else if d.HasChange("rd") {
		obj["rd"] = nil
	}

	if v, ok := d.GetOk("import_rt"); ok || d.HasChange("import_rt") {
		t, err := expandSystemEvpnImportRt(d, v, "import_rt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["import-rt"] = t
		}
	}

	if v, ok := d.GetOk("export_rt"); ok || d.HasChange("export_rt") {
		t, err := expandSystemEvpnExportRt(d, v, "export_rt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["export-rt"] = t
		}
	}

	if v, ok := d.GetOk("ip_local_learning"); ok {
		t, err := expandSystemEvpnIpLocalLearning(d, v, "ip_local_learning", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-local-learning"] = t
		}
	}

	if v, ok := d.GetOk("arp_suppression"); ok {
		t, err := expandSystemEvpnArpSuppression(d, v, "arp_suppression", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-suppression"] = t
		}
	}

	if v, ok := d.GetOk("adv_default_gw"); ok {
		t, err := expandSystemEvpnAdvDefaultGw(d, v, "adv_default_gw", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-default-gw"] = t
		}
	}

	if v, ok := d.GetOk("distribute_local_route"); ok {
		t, err := expandSystemEvpnDistributeLocalRoute(d, v, "distribute_local_route", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-local-route"] = t
		}
	}

	if v, ok := d.GetOk("virtual_mac_vrid"); ok {
		t, err := expandSystemEvpnVirtualMacVrid(d, v, "virtual_mac_vrid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-mac-vrid"] = t
		}
	} else if d.HasChange("virtual_mac_vrid") {
		obj["virtual-mac-vrid"] = nil
	}

	if v, ok := d.GetOkExists("l3_instance"); ok {
		t, err := expandSystemEvpnL3Instance(d, v, "l3_instance", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l3-instance"] = t
		}
	} else if d.HasChange("l3_instance") {
		obj["l3-instance"] = nil
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemEvpnInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	return &obj, nil
}
