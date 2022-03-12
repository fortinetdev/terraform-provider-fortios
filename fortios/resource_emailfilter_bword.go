// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure AntiSpam banned word list.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceEmailfilterBword() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailfilterBwordCreate,
		Read:   resourceEmailfilterBwordRead,
		Update: resourceEmailfilterBwordUpdate,
		Delete: resourceEmailfilterBwordDelete,

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
						"pattern": &schema.Schema{
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
						"where": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"language": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"score": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 99999),
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

func resourceEmailfilterBwordCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectEmailfilterBword(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating EmailfilterBword resource while getting object: %v", err)
	}

	o, err := c.CreateEmailfilterBword(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating EmailfilterBword resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("EmailfilterBword")
	}

	return resourceEmailfilterBwordRead(d, m)
}

func resourceEmailfilterBwordUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectEmailfilterBword(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterBword resource while getting object: %v", err)
	}

	o, err := c.UpdateEmailfilterBword(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterBword resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("EmailfilterBword")
	}

	return resourceEmailfilterBwordRead(d, m)
}

func resourceEmailfilterBwordDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteEmailfilterBword(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting EmailfilterBword resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterBwordRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadEmailfilterBword(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterBword resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEmailfilterBword(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterBword resource from API: %v", err)
	}
	return nil
}

func flattenEmailfilterBwordId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterBwordName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterBwordComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterBwordEntries(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
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

			tmp["status"] = flattenEmailfilterBwordEntriesStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenEmailfilterBwordEntriesId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {

			tmp["pattern"] = flattenEmailfilterBwordEntriesPattern(i["pattern"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern_type"
		if _, ok := i["pattern-type"]; ok {

			tmp["pattern_type"] = flattenEmailfilterBwordEntriesPatternType(i["pattern-type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {

			tmp["action"] = flattenEmailfilterBwordEntriesAction(i["action"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "where"
		if _, ok := i["where"]; ok {

			tmp["where"] = flattenEmailfilterBwordEntriesWhere(i["where"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "language"
		if _, ok := i["language"]; ok {

			tmp["language"] = flattenEmailfilterBwordEntriesLanguage(i["language"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "score"
		if _, ok := i["score"]; ok {

			tmp["score"] = flattenEmailfilterBwordEntriesScore(i["score"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "id", d)
	return result
}

func flattenEmailfilterBwordEntriesStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterBwordEntriesId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterBwordEntriesPattern(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterBwordEntriesPatternType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterBwordEntriesAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterBwordEntriesWhere(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterBwordEntriesLanguage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterBwordEntriesScore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectEmailfilterBword(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenEmailfilterBwordId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenEmailfilterBwordName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenEmailfilterBwordComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenEmailfilterBwordEntries(o["entries"], d, "entries", sv)); err != nil {
			if !fortiAPIPatch(o["entries"]) {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenEmailfilterBwordEntries(o["entries"], d, "entries", sv)); err != nil {
				if !fortiAPIPatch(o["entries"]) {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenEmailfilterBwordFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandEmailfilterBwordId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBwordName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBwordComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBwordEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["status"], _ = expandEmailfilterBwordEntriesStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandEmailfilterBwordEntriesId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["pattern"], _ = expandEmailfilterBwordEntriesPattern(d, i["pattern"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["pattern-type"], _ = expandEmailfilterBwordEntriesPatternType(d, i["pattern_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["action"], _ = expandEmailfilterBwordEntriesAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "where"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["where"], _ = expandEmailfilterBwordEntriesWhere(d, i["where"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "language"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["language"], _ = expandEmailfilterBwordEntriesLanguage(d, i["language"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "score"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["score"], _ = expandEmailfilterBwordEntriesScore(d, i["score"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandEmailfilterBwordEntriesStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBwordEntriesId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBwordEntriesPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBwordEntriesPatternType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBwordEntriesAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBwordEntriesWhere(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBwordEntriesLanguage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterBwordEntriesScore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectEmailfilterBword(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandEmailfilterBwordId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandEmailfilterBwordName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandEmailfilterBwordComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok {

		t, err := expandEmailfilterBwordEntries(d, v, "entries", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	return &obj, nil
}
