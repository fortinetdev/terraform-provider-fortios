// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure AntiSpam MIME header.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceEmailfilterMheader() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailfilterMheaderCreate,
		Read:   resourceEmailfilterMheaderRead,
		Update: resourceEmailfilterMheaderUpdate,
		Delete: resourceEmailfilterMheaderDelete,

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
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fieldname": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"fieldbody": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
						"pattern_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"action": &schema.Schema{
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
		},
	}
}

func resourceEmailfilterMheaderCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectEmailfilterMheader(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating EmailfilterMheader resource while getting object: %v", err)
	}

	o, err := c.CreateEmailfilterMheader(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating EmailfilterMheader resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("EmailfilterMheader")
	}

	return resourceEmailfilterMheaderRead(d, m)
}

func resourceEmailfilterMheaderUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectEmailfilterMheader(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterMheader resource while getting object: %v", err)
	}

	o, err := c.UpdateEmailfilterMheader(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterMheader resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("EmailfilterMheader")
	}

	return resourceEmailfilterMheaderRead(d, m)
}

func resourceEmailfilterMheaderDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteEmailfilterMheader(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting EmailfilterMheader resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterMheaderRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadEmailfilterMheader(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterMheader resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEmailfilterMheader(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterMheader resource from API: %v", err)
	}
	return nil
}

func flattenEmailfilterMheaderId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterMheaderName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterMheaderComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterMheaderEntries(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenEmailfilterMheaderEntriesStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenEmailfilterMheaderEntriesId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldname"
		if _, ok := i["fieldname"]; ok {

			tmp["fieldname"] = flattenEmailfilterMheaderEntriesFieldname(i["fieldname"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldbody"
		if _, ok := i["fieldbody"]; ok {

			tmp["fieldbody"] = flattenEmailfilterMheaderEntriesFieldbody(i["fieldbody"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern_type"
		if _, ok := i["pattern-type"]; ok {

			tmp["pattern_type"] = flattenEmailfilterMheaderEntriesPatternType(i["pattern-type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {

			tmp["action"] = flattenEmailfilterMheaderEntriesAction(i["action"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenEmailfilterMheaderEntriesStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterMheaderEntriesId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterMheaderEntriesFieldname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterMheaderEntriesFieldbody(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterMheaderEntriesPatternType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterMheaderEntriesAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectEmailfilterMheader(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenEmailfilterMheaderId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenEmailfilterMheaderName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenEmailfilterMheaderComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenEmailfilterMheaderEntries(o["entries"], d, "entries", sv)); err != nil {
			if !fortiAPIPatch(o["entries"]) {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenEmailfilterMheaderEntries(o["entries"], d, "entries", sv)); err != nil {
				if !fortiAPIPatch(o["entries"]) {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenEmailfilterMheaderFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandEmailfilterMheaderId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandEmailfilterMheaderEntriesStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandEmailfilterMheaderEntriesId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldname"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["fieldname"], _ = expandEmailfilterMheaderEntriesFieldname(d, i["fieldname"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldbody"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["fieldbody"], _ = expandEmailfilterMheaderEntriesFieldbody(d, i["fieldbody"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["pattern-type"], _ = expandEmailfilterMheaderEntriesPatternType(d, i["pattern_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["action"], _ = expandEmailfilterMheaderEntriesAction(d, i["action"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandEmailfilterMheaderEntriesStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderEntriesId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderEntriesFieldname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderEntriesFieldbody(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderEntriesPatternType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterMheaderEntriesAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectEmailfilterMheader(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandEmailfilterMheaderId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandEmailfilterMheaderName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandEmailfilterMheaderComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok {

		t, err := expandEmailfilterMheaderEntries(d, v, "entries", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	return &obj, nil
}
