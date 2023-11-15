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
			},
			"fosid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"rd": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
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
							Computed:     true,
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
							Computed:     true,
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemEvpn(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemEvpn resource while getting object: %v", err)
	}

	o, err := c.CreateSystemEvpn(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemEvpn resource: %v", err)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_target"
		if _, ok := i["route-target"]; ok {
			tmp["route_target"] = flattenSystemEvpnImportRtRouteTarget(i["route-target"], d, pre_append, sv)
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_target"
		if _, ok := i["route-target"]; ok {
			tmp["route_target"] = flattenSystemEvpnExportRtRouteTarget(i["route-target"], d, pre_append, sv)
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

	if v, ok := d.GetOk("rd"); ok {
		t, err := expandSystemEvpnRd(d, v, "rd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rd"] = t
		}
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

	return &obj, nil
}
