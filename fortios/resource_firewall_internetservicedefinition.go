// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Internet Service definition.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallInternetServiceDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInternetServiceDefinitionCreate,
		Read:   resourceFirewallInternetServiceDefinitionRead,
		Update: resourceFirewallInternetServiceDefinitionUpdate,
		Delete: resourceFirewallInternetServiceDefinitionDelete,

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
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"entry": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"category_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"protocol": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
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
										ValidateFunc: validation.IntBetween(1, 65535),
										Optional:     true,
										Computed:     true,
									},
									"end_port": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceFirewallInternetServiceDefinitionCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallInternetServiceDefinition(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceDefinition resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallInternetServiceDefinition(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceDefinition resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallInternetServiceDefinition")
	}

	return resourceFirewallInternetServiceDefinitionRead(d, m)
}

func resourceFirewallInternetServiceDefinitionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallInternetServiceDefinition(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceDefinition resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallInternetServiceDefinition(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceDefinition resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallInternetServiceDefinition")
	}

	return resourceFirewallInternetServiceDefinitionRead(d, m)
}

func resourceFirewallInternetServiceDefinitionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallInternetServiceDefinition(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInternetServiceDefinition resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceDefinitionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallInternetServiceDefinition(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceDefinition resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInternetServiceDefinition(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceDefinition resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInternetServiceDefinitionId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceDefinitionEntry(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := i["seq-num"]; ok {

			tmp["seq_num"] = flattenFirewallInternetServiceDefinitionEntrySeqNum(i["seq-num"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category_id"
		if _, ok := i["category-id"]; ok {

			tmp["category_id"] = flattenFirewallInternetServiceDefinitionEntryCategoryId(i["category-id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenFirewallInternetServiceDefinitionEntryName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {

			tmp["protocol"] = flattenFirewallInternetServiceDefinitionEntryProtocol(i["protocol"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_range"
		if _, ok := i["port-range"]; ok {

			tmp["port_range"] = flattenFirewallInternetServiceDefinitionEntryPortRange(i["port-range"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = flattenFirewallInternetServiceDefinitionEntryPort(i["port"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "seq_num", d)
	return result
}

func flattenFirewallInternetServiceDefinitionEntrySeqNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceDefinitionEntryCategoryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceDefinitionEntryName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceDefinitionEntryProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceDefinitionEntryPortRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenFirewallInternetServiceDefinitionEntryPortRangeId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := i["start-port"]; ok {

			tmp["start_port"] = flattenFirewallInternetServiceDefinitionEntryPortRangeStartPort(i["start-port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := i["end-port"]; ok {

			tmp["end_port"] = flattenFirewallInternetServiceDefinitionEntryPortRangeEndPort(i["end-port"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallInternetServiceDefinitionEntryPortRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceDefinitionEntryPortRangeStartPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceDefinitionEntryPortRangeEndPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceDefinitionEntryPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallInternetServiceDefinition(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenFirewallInternetServiceDefinitionId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entry", flattenFirewallInternetServiceDefinitionEntry(o["entry"], d, "entry", sv)); err != nil {
			if !fortiAPIPatch(o["entry"]) {
				return fmt.Errorf("Error reading entry: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entry"); ok {
			if err = d.Set("entry", flattenFirewallInternetServiceDefinitionEntry(o["entry"], d, "entry", sv)); err != nil {
				if !fortiAPIPatch(o["entry"]) {
					return fmt.Errorf("Error reading entry: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFirewallInternetServiceDefinitionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallInternetServiceDefinitionId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceDefinitionEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["seq-num"], _ = expandFirewallInternetServiceDefinitionEntrySeqNum(d, i["seq_num"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["category-id"], _ = expandFirewallInternetServiceDefinitionEntryCategoryId(d, i["category_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandFirewallInternetServiceDefinitionEntryName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["protocol"], _ = expandFirewallInternetServiceDefinitionEntryProtocol(d, i["protocol"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_range"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["port-range"], _ = expandFirewallInternetServiceDefinitionEntryPortRange(d, i["port_range"], pre_append, sv)
		} else {
			tmp["port-range"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["port"], _ = expandFirewallInternetServiceDefinitionEntryPort(d, i["port"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceDefinitionEntrySeqNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceDefinitionEntryCategoryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceDefinitionEntryName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceDefinitionEntryProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceDefinitionEntryPortRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandFirewallInternetServiceDefinitionEntryPortRangeId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["start-port"], _ = expandFirewallInternetServiceDefinitionEntryPortRangeStartPort(d, i["start_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["end-port"], _ = expandFirewallInternetServiceDefinitionEntryPortRangeEndPort(d, i["end_port"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceDefinitionEntryPortRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceDefinitionEntryPortRangeStartPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceDefinitionEntryPortRangeEndPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceDefinitionEntryPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInternetServiceDefinition(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandFirewallInternetServiceDefinitionId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("entry"); ok {

		t, err := expandFirewallInternetServiceDefinitionEntry(d, v, "entry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entry"] = t
		}
	}

	return &obj, nil
}
