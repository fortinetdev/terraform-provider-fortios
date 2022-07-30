// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure custom Internet Services.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallInternetServiceCustom() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInternetServiceCustomCreate,
		Read:   resourceFirewallInternetServiceCustomRead,
		Update: resourceFirewallInternetServiceCustomUpdate,
		Delete: resourceFirewallInternetServiceCustomDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"reputation": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
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
		},
	}
}

func resourceFirewallInternetServiceCustomCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallInternetServiceCustom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceCustom resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallInternetServiceCustom(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceCustom resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallInternetServiceCustom")
	}

	return resourceFirewallInternetServiceCustomRead(d, m)
}

func resourceFirewallInternetServiceCustomUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallInternetServiceCustom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceCustom resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallInternetServiceCustom(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceCustom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallInternetServiceCustom")
	}

	return resourceFirewallInternetServiceCustomRead(d, m)
}

func resourceFirewallInternetServiceCustomDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallInternetServiceCustom(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInternetServiceCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceCustomRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallInternetServiceCustom(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceCustom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInternetServiceCustom(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceCustom resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInternetServiceCustomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceCustomReputation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceCustomComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceCustomEntry(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenFirewallInternetServiceCustomEntryId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {

			tmp["protocol"] = flattenFirewallInternetServiceCustomEntryProtocol(i["protocol"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_range"
		if _, ok := i["port-range"]; ok {

			tmp["port_range"] = flattenFirewallInternetServiceCustomEntryPortRange(i["port-range"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := i["dst"]; ok {

			tmp["dst"] = flattenFirewallInternetServiceCustomEntryDst(i["dst"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallInternetServiceCustomEntryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceCustomEntryProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceCustomEntryPortRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenFirewallInternetServiceCustomEntryPortRangeId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := i["start-port"]; ok {

			tmp["start_port"] = flattenFirewallInternetServiceCustomEntryPortRangeStartPort(i["start-port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := i["end-port"]; ok {

			tmp["end_port"] = flattenFirewallInternetServiceCustomEntryPortRangeEndPort(i["end-port"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallInternetServiceCustomEntryPortRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceCustomEntryPortRangeStartPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceCustomEntryPortRangeEndPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceCustomEntryDst(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallInternetServiceCustomEntryDstName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallInternetServiceCustomEntryDstName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallInternetServiceCustom(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenFirewallInternetServiceCustomName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("reputation", flattenFirewallInternetServiceCustomReputation(o["reputation"], d, "reputation", sv)); err != nil {
		if !fortiAPIPatch(o["reputation"]) {
			return fmt.Errorf("Error reading reputation: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallInternetServiceCustomComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entry", flattenFirewallInternetServiceCustomEntry(o["entry"], d, "entry", sv)); err != nil {
			if !fortiAPIPatch(o["entry"]) {
				return fmt.Errorf("Error reading entry: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entry"); ok {
			if err = d.Set("entry", flattenFirewallInternetServiceCustomEntry(o["entry"], d, "entry", sv)); err != nil {
				if !fortiAPIPatch(o["entry"]) {
					return fmt.Errorf("Error reading entry: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFirewallInternetServiceCustomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallInternetServiceCustomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceCustomReputation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceCustomComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceCustomEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandFirewallInternetServiceCustomEntryId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["protocol"], _ = expandFirewallInternetServiceCustomEntryProtocol(d, i["protocol"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_range"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["port-range"], _ = expandFirewallInternetServiceCustomEntryPortRange(d, i["port_range"], pre_append, sv)
		} else {
			tmp["port-range"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["dst"], _ = expandFirewallInternetServiceCustomEntryDst(d, i["dst"], pre_append, sv)
		} else {
			tmp["dst"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceCustomEntryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceCustomEntryProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceCustomEntryPortRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandFirewallInternetServiceCustomEntryPortRangeId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["start-port"], _ = expandFirewallInternetServiceCustomEntryPortRangeStartPort(d, i["start_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["end-port"], _ = expandFirewallInternetServiceCustomEntryPortRangeEndPort(d, i["end_port"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceCustomEntryPortRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceCustomEntryPortRangeStartPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceCustomEntryPortRangeEndPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceCustomEntryDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallInternetServiceCustomEntryDstName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceCustomEntryDstName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInternetServiceCustom(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallInternetServiceCustomName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("reputation"); ok {

		t, err := expandFirewallInternetServiceCustomReputation(d, v, "reputation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reputation"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandFirewallInternetServiceCustomComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entry"); ok || d.HasChange("entry") {

		t, err := expandFirewallInternetServiceCustomEntry(d, v, "entry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entry"] = t
		}
	}

	return &obj, nil
}
