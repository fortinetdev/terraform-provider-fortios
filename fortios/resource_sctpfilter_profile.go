// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure SCTP filter profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSctpFilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSctpFilterProfileCreate,
		Read:   resourceSctpFilterProfileRead,
		Update: resourceSctpFilterProfileUpdate,
		Delete: resourceSctpFilterProfileDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"ppid_filters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ppid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"comment": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
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

func resourceSctpFilterProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSctpFilterProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SctpFilterProfile resource while getting object: %v", err)
	}

	o, err := c.CreateSctpFilterProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SctpFilterProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SctpFilterProfile")
	}

	return resourceSctpFilterProfileRead(d, m)
}

func resourceSctpFilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSctpFilterProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SctpFilterProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateSctpFilterProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SctpFilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SctpFilterProfile")
	}

	return resourceSctpFilterProfileRead(d, m)
}

func resourceSctpFilterProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSctpFilterProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SctpFilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSctpFilterProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSctpFilterProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SctpFilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSctpFilterProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SctpFilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenSctpFilterProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSctpFilterProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSctpFilterProfilePpidFilters(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSctpFilterProfilePpidFiltersId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ppid"
		if _, ok := i["ppid"]; ok {

			tmp["ppid"] = flattenSctpFilterProfilePpidFiltersPpid(i["ppid"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {

			tmp["action"] = flattenSctpFilterProfilePpidFiltersAction(i["action"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {

			tmp["comment"] = flattenSctpFilterProfilePpidFiltersComment(i["comment"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "id", d)
	return result
}

func flattenSctpFilterProfilePpidFiltersId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSctpFilterProfilePpidFiltersPpid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSctpFilterProfilePpidFiltersAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSctpFilterProfilePpidFiltersComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSctpFilterProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSctpFilterProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenSctpFilterProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ppid_filters", flattenSctpFilterProfilePpidFilters(o["ppid-filters"], d, "ppid_filters", sv)); err != nil {
			if !fortiAPIPatch(o["ppid-filters"]) {
				return fmt.Errorf("Error reading ppid_filters: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ppid_filters"); ok {
			if err = d.Set("ppid_filters", flattenSctpFilterProfilePpidFilters(o["ppid-filters"], d, "ppid_filters", sv)); err != nil {
				if !fortiAPIPatch(o["ppid-filters"]) {
					return fmt.Errorf("Error reading ppid_filters: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSctpFilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSctpFilterProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSctpFilterProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSctpFilterProfilePpidFilters(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandSctpFilterProfilePpidFiltersId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ppid"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ppid"], _ = expandSctpFilterProfilePpidFiltersPpid(d, i["ppid"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["action"], _ = expandSctpFilterProfilePpidFiltersAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["comment"], _ = expandSctpFilterProfilePpidFiltersComment(d, i["comment"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSctpFilterProfilePpidFiltersId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSctpFilterProfilePpidFiltersPpid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSctpFilterProfilePpidFiltersAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSctpFilterProfilePpidFiltersComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSctpFilterProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSctpFilterProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandSctpFilterProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("ppid_filters"); ok {

		t, err := expandSctpFilterProfilePpidFilters(d, v, "ppid_filters", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppid-filters"] = t
		}
	}

	return &obj, nil
}
