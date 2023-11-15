// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Speed test schedule for each interface.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSpeedTestSchedule() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSpeedTestScheduleCreate,
		Read:   resourceSystemSpeedTestScheduleRead,
		Update: resourceSystemSpeedTestScheduleUpdate,
		Delete: resourceSystemSpeedTestScheduleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedules": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dynamic_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_inbandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_outbandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_inbandwidth_maximum": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),
				Optional:     true,
				Computed:     true,
			},
			"update_inbandwidth_minimum": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),
				Optional:     true,
				Computed:     true,
			},
			"update_outbandwidth_maximum": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),
				Optional:     true,
				Computed:     true,
			},
			"update_outbandwidth_minimum": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),
				Optional:     true,
				Computed:     true,
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

func resourceSystemSpeedTestScheduleCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemSpeedTestSchedule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSpeedTestSchedule resource while getting object: %v", err)
	}

	o, err := c.CreateSystemSpeedTestSchedule(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemSpeedTestSchedule resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSpeedTestSchedule")
	}

	return resourceSystemSpeedTestScheduleRead(d, m)
}

func resourceSystemSpeedTestScheduleUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemSpeedTestSchedule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSpeedTestSchedule resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSpeedTestSchedule(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemSpeedTestSchedule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSpeedTestSchedule")
	}

	return resourceSystemSpeedTestScheduleRead(d, m)
}

func resourceSystemSpeedTestScheduleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemSpeedTestSchedule(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSpeedTestSchedule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSpeedTestScheduleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemSpeedTestSchedule(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemSpeedTestSchedule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSpeedTestSchedule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSpeedTestSchedule resource from API: %v", err)
	}
	return nil
}

func flattenSystemSpeedTestScheduleInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleDiffserv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleServerName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleSchedules(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenSystemSpeedTestScheduleSchedulesName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSpeedTestScheduleSchedulesName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleDynamicServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleUpdateInbandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleUpdateOutbandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleUpdateInbandwidthMaximum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleUpdateInbandwidthMinimum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleUpdateOutbandwidthMaximum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSpeedTestScheduleUpdateOutbandwidthMinimum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSpeedTestSchedule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("interface", flattenSystemSpeedTestScheduleInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSpeedTestScheduleStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("diffserv", flattenSystemSpeedTestScheduleDiffserv(o["diffserv"], d, "diffserv", sv)); err != nil {
		if !fortiAPIPatch(o["diffserv"]) {
			return fmt.Errorf("Error reading diffserv: %v", err)
		}
	}

	if err = d.Set("server_name", flattenSystemSpeedTestScheduleServerName(o["server-name"], d, "server_name", sv)); err != nil {
		if !fortiAPIPatch(o["server-name"]) {
			return fmt.Errorf("Error reading server_name: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemSpeedTestScheduleMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("schedules", flattenSystemSpeedTestScheduleSchedules(o["schedules"], d, "schedules", sv)); err != nil {
			if !fortiAPIPatch(o["schedules"]) {
				return fmt.Errorf("Error reading schedules: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("schedules"); ok {
			if err = d.Set("schedules", flattenSystemSpeedTestScheduleSchedules(o["schedules"], d, "schedules", sv)); err != nil {
				if !fortiAPIPatch(o["schedules"]) {
					return fmt.Errorf("Error reading schedules: %v", err)
				}
			}
		}
	}

	if err = d.Set("dynamic_server", flattenSystemSpeedTestScheduleDynamicServer(o["dynamic-server"], d, "dynamic_server", sv)); err != nil {
		if !fortiAPIPatch(o["dynamic-server"]) {
			return fmt.Errorf("Error reading dynamic_server: %v", err)
		}
	}

	if err = d.Set("update_inbandwidth", flattenSystemSpeedTestScheduleUpdateInbandwidth(o["update-inbandwidth"], d, "update_inbandwidth", sv)); err != nil {
		if !fortiAPIPatch(o["update-inbandwidth"]) {
			return fmt.Errorf("Error reading update_inbandwidth: %v", err)
		}
	}

	if err = d.Set("update_outbandwidth", flattenSystemSpeedTestScheduleUpdateOutbandwidth(o["update-outbandwidth"], d, "update_outbandwidth", sv)); err != nil {
		if !fortiAPIPatch(o["update-outbandwidth"]) {
			return fmt.Errorf("Error reading update_outbandwidth: %v", err)
		}
	}

	if err = d.Set("update_inbandwidth_maximum", flattenSystemSpeedTestScheduleUpdateInbandwidthMaximum(o["update-inbandwidth-maximum"], d, "update_inbandwidth_maximum", sv)); err != nil {
		if !fortiAPIPatch(o["update-inbandwidth-maximum"]) {
			return fmt.Errorf("Error reading update_inbandwidth_maximum: %v", err)
		}
	}

	if err = d.Set("update_inbandwidth_minimum", flattenSystemSpeedTestScheduleUpdateInbandwidthMinimum(o["update-inbandwidth-minimum"], d, "update_inbandwidth_minimum", sv)); err != nil {
		if !fortiAPIPatch(o["update-inbandwidth-minimum"]) {
			return fmt.Errorf("Error reading update_inbandwidth_minimum: %v", err)
		}
	}

	if err = d.Set("update_outbandwidth_maximum", flattenSystemSpeedTestScheduleUpdateOutbandwidthMaximum(o["update-outbandwidth-maximum"], d, "update_outbandwidth_maximum", sv)); err != nil {
		if !fortiAPIPatch(o["update-outbandwidth-maximum"]) {
			return fmt.Errorf("Error reading update_outbandwidth_maximum: %v", err)
		}
	}

	if err = d.Set("update_outbandwidth_minimum", flattenSystemSpeedTestScheduleUpdateOutbandwidthMinimum(o["update-outbandwidth-minimum"], d, "update_outbandwidth_minimum", sv)); err != nil {
		if !fortiAPIPatch(o["update-outbandwidth-minimum"]) {
			return fmt.Errorf("Error reading update_outbandwidth_minimum: %v", err)
		}
	}

	return nil
}

func flattenSystemSpeedTestScheduleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSpeedTestScheduleInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleDiffserv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleServerName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleSchedules(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemSpeedTestScheduleSchedulesName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSpeedTestScheduleSchedulesName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleDynamicServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleUpdateInbandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleUpdateOutbandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleUpdateInbandwidthMaximum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleUpdateInbandwidthMinimum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleUpdateOutbandwidthMaximum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestScheduleUpdateOutbandwidthMinimum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSpeedTestSchedule(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemSpeedTestScheduleInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemSpeedTestScheduleStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("diffserv"); ok {
		t, err := expandSystemSpeedTestScheduleDiffserv(d, v, "diffserv", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv"] = t
		}
	}

	if v, ok := d.GetOk("server_name"); ok {
		t, err := expandSystemSpeedTestScheduleServerName(d, v, "server_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-name"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandSystemSpeedTestScheduleMode(d, v, "mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("schedules"); ok || d.HasChange("schedules") {
		t, err := expandSystemSpeedTestScheduleSchedules(d, v, "schedules", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedules"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_server"); ok {
		t, err := expandSystemSpeedTestScheduleDynamicServer(d, v, "dynamic_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-server"] = t
		}
	}

	if v, ok := d.GetOk("update_inbandwidth"); ok {
		t, err := expandSystemSpeedTestScheduleUpdateInbandwidth(d, v, "update_inbandwidth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-inbandwidth"] = t
		}
	}

	if v, ok := d.GetOk("update_outbandwidth"); ok {
		t, err := expandSystemSpeedTestScheduleUpdateOutbandwidth(d, v, "update_outbandwidth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-outbandwidth"] = t
		}
	}

	if v, ok := d.GetOkExists("update_inbandwidth_maximum"); ok {
		t, err := expandSystemSpeedTestScheduleUpdateInbandwidthMaximum(d, v, "update_inbandwidth_maximum", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-inbandwidth-maximum"] = t
		}
	}

	if v, ok := d.GetOkExists("update_inbandwidth_minimum"); ok {
		t, err := expandSystemSpeedTestScheduleUpdateInbandwidthMinimum(d, v, "update_inbandwidth_minimum", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-inbandwidth-minimum"] = t
		}
	}

	if v, ok := d.GetOkExists("update_outbandwidth_maximum"); ok {
		t, err := expandSystemSpeedTestScheduleUpdateOutbandwidthMaximum(d, v, "update_outbandwidth_maximum", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-outbandwidth-maximum"] = t
		}
	}

	if v, ok := d.GetOkExists("update_outbandwidth_minimum"); ok {
		t, err := expandSystemSpeedTestScheduleUpdateOutbandwidthMinimum(d, v, "update_outbandwidth_minimum", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-outbandwidth-minimum"] = t
		}
	}

	return &obj, nil
}
