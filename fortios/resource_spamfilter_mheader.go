// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure AntiSpam MIME header.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSpamfilterMheader() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpamfilterMheaderCreate,
		Read:   resourceSpamfilterMheaderRead,
		Update: resourceSpamfilterMheaderUpdate,
		Delete: resourceSpamfilterMheaderDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Required: true,
			},
			"name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required: true,
			},
			"comment": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional: true,
			},
			"entries": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"fieldname": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional: true,
							Computed: true,
						},
						"fieldbody": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional: true,
							Computed: true,
						},
						"pattern_type": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"action": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceSpamfilterMheaderCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSpamfilterMheader(d)
	if err != nil {
		return fmt.Errorf("Error creating SpamfilterMheader resource while getting object: %v", err)
	}

	o, err := c.CreateSpamfilterMheader(obj)

	if err != nil {
		return fmt.Errorf("Error creating SpamfilterMheader resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SpamfilterMheader")
	}

	return resourceSpamfilterMheaderRead(d, m)
}

func resourceSpamfilterMheaderUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSpamfilterMheader(d)
	if err != nil {
		return fmt.Errorf("Error updating SpamfilterMheader resource while getting object: %v", err)
	}

	o, err := c.UpdateSpamfilterMheader(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SpamfilterMheader resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SpamfilterMheader")
	}

	return resourceSpamfilterMheaderRead(d, m)
}

func resourceSpamfilterMheaderDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSpamfilterMheader(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SpamfilterMheader resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSpamfilterMheaderRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSpamfilterMheader(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SpamfilterMheader resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSpamfilterMheader(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SpamfilterMheader resource from API: %v", err)
	}
	return nil
}


func flattenSpamfilterMheaderId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterMheaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterMheaderComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterMheaderEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["status"] = flattenSpamfilterMheaderEntriesStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenSpamfilterMheaderEntriesId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldname"
		if _, ok := i["fieldname"]; ok {
			tmp["fieldname"] = flattenSpamfilterMheaderEntriesFieldname(i["fieldname"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldbody"
		if _, ok := i["fieldbody"]; ok {
			tmp["fieldbody"] = flattenSpamfilterMheaderEntriesFieldbody(i["fieldbody"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern_type"
		if _, ok := i["pattern-type"]; ok {
			tmp["pattern_type"] = flattenSpamfilterMheaderEntriesPatternType(i["pattern-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = flattenSpamfilterMheaderEntriesAction(i["action"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSpamfilterMheaderEntriesStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterMheaderEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterMheaderEntriesFieldname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterMheaderEntriesFieldbody(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterMheaderEntriesPatternType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterMheaderEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectSpamfilterMheader(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("fosid", flattenSpamfilterMheaderId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenSpamfilterMheaderName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenSpamfilterMheaderComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("entries", flattenSpamfilterMheaderEntries(o["entries"], d, "entries")); err != nil {
            if !fortiAPIPatch(o["entries"]) {
                return fmt.Errorf("Error reading entries: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("entries"); ok {
            if err = d.Set("entries", flattenSpamfilterMheaderEntries(o["entries"], d, "entries")); err != nil {
                if !fortiAPIPatch(o["entries"]) {
                    return fmt.Errorf("Error reading entries: %v", err)
                }
            }
        }
    }


	return nil
}

func flattenSpamfilterMheaderFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandSpamfilterMheaderId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterMheaderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterMheaderComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterMheaderEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandSpamfilterMheaderEntriesStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandSpamfilterMheaderEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldname"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fieldname"], _ = expandSpamfilterMheaderEntriesFieldname(d, i["fieldname"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fieldbody"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fieldbody"], _ = expandSpamfilterMheaderEntriesFieldbody(d, i["fieldbody"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pattern-type"], _ = expandSpamfilterMheaderEntriesPatternType(d, i["pattern_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandSpamfilterMheaderEntriesAction(d, i["action"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSpamfilterMheaderEntriesStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterMheaderEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterMheaderEntriesFieldname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterMheaderEntriesFieldbody(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterMheaderEntriesPatternType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterMheaderEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectSpamfilterMheader(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandSpamfilterMheaderId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSpamfilterMheaderName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandSpamfilterMheaderComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok {
		t, err := expandSpamfilterMheaderEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}


	return &obj, nil
}

