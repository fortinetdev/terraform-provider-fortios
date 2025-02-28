// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure VideoFilter profile.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVideofilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceVideofilterProfileCreate,
		Read:   resourceVideofilterProfileRead,
		Update: resourceVideofilterProfileUpdate,
		Delete: resourceVideofilterProfileDelete,

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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"filters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"comment": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keyword": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"category": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),
							Optional:     true,
						},
						"channel": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"youtube_channel_filter": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fortiguard_category": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"category_id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"youtube": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vimeo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dailymotion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
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

func resourceVideofilterProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVideofilterProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating VideofilterProfile resource while getting object: %v", err)
	}

	o, err := c.CreateVideofilterProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating VideofilterProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VideofilterProfile")
	}

	return resourceVideofilterProfileRead(d, m)
}

func resourceVideofilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVideofilterProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VideofilterProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateVideofilterProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating VideofilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VideofilterProfile")
	}

	return resourceVideofilterProfileRead(d, m)
}

func resourceVideofilterProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteVideofilterProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting VideofilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVideofilterProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVideofilterProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VideofilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVideofilterProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VideofilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenVideofilterProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVideofilterProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVideofilterProfileFilters(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenVideofilterProfileFiltersId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if cur_v, ok := i["comment"]; ok {
			tmp["comment"] = flattenVideofilterProfileFiltersComment(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if cur_v, ok := i["type"]; ok {
			tmp["type"] = flattenVideofilterProfileFiltersType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyword"
		if cur_v, ok := i["keyword"]; ok {
			tmp["keyword"] = flattenVideofilterProfileFiltersKeyword(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if cur_v, ok := i["category"]; ok {
			tmp["category"] = flattenVideofilterProfileFiltersCategory(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "channel"
		if cur_v, ok := i["channel"]; ok {
			tmp["channel"] = flattenVideofilterProfileFiltersChannel(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if cur_v, ok := i["action"]; ok {
			tmp["action"] = flattenVideofilterProfileFiltersAction(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if cur_v, ok := i["log"]; ok {
			tmp["log"] = flattenVideofilterProfileFiltersLog(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenVideofilterProfileFiltersId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVideofilterProfileFiltersComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVideofilterProfileFiltersType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVideofilterProfileFiltersKeyword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVideofilterProfileFiltersCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVideofilterProfileFiltersChannel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVideofilterProfileFiltersAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVideofilterProfileFiltersLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVideofilterProfileDefaultAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVideofilterProfileLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVideofilterProfileYoutubeChannelFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVideofilterProfileFortiguardCategory(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "filters"
	if _, ok := i["filters"]; ok {
		result["filters"] = flattenVideofilterProfileFortiguardCategoryFilters(i["filters"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenVideofilterProfileFortiguardCategoryFilters(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenVideofilterProfileFortiguardCategoryFiltersId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if cur_v, ok := i["action"]; ok {
			tmp["action"] = flattenVideofilterProfileFortiguardCategoryFiltersAction(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category_id"
		if cur_v, ok := i["category-id"]; ok {
			tmp["category_id"] = flattenVideofilterProfileFortiguardCategoryFiltersCategoryId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if cur_v, ok := i["log"]; ok {
			tmp["log"] = flattenVideofilterProfileFortiguardCategoryFiltersLog(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenVideofilterProfileFortiguardCategoryFiltersId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVideofilterProfileFortiguardCategoryFiltersAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVideofilterProfileFortiguardCategoryFiltersCategoryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVideofilterProfileFortiguardCategoryFiltersLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVideofilterProfileYoutube(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVideofilterProfileVimeo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVideofilterProfileDailymotion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVideofilterProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVideofilterProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenVideofilterProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenVideofilterProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("filters", flattenVideofilterProfileFilters(o["filters"], d, "filters", sv)); err != nil {
			if !fortiAPIPatch(o["filters"]) {
				return fmt.Errorf("Error reading filters: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("filters"); ok {
			if err = d.Set("filters", flattenVideofilterProfileFilters(o["filters"], d, "filters", sv)); err != nil {
				if !fortiAPIPatch(o["filters"]) {
					return fmt.Errorf("Error reading filters: %v", err)
				}
			}
		}
	}

	if err = d.Set("default_action", flattenVideofilterProfileDefaultAction(o["default-action"], d, "default_action", sv)); err != nil {
		if !fortiAPIPatch(o["default-action"]) {
			return fmt.Errorf("Error reading default_action: %v", err)
		}
	}

	if err = d.Set("log", flattenVideofilterProfileLog(o["log"], d, "log", sv)); err != nil {
		if !fortiAPIPatch(o["log"]) {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("youtube_channel_filter", flattenVideofilterProfileYoutubeChannelFilter(o["youtube-channel-filter"], d, "youtube_channel_filter", sv)); err != nil {
		if !fortiAPIPatch(o["youtube-channel-filter"]) {
			return fmt.Errorf("Error reading youtube_channel_filter: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("fortiguard_category", flattenVideofilterProfileFortiguardCategory(o["fortiguard-category"], d, "fortiguard_category", sv)); err != nil {
			if !fortiAPIPatch(o["fortiguard-category"]) {
				return fmt.Errorf("Error reading fortiguard_category: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fortiguard_category"); ok {
			if err = d.Set("fortiguard_category", flattenVideofilterProfileFortiguardCategory(o["fortiguard-category"], d, "fortiguard_category", sv)); err != nil {
				if !fortiAPIPatch(o["fortiguard-category"]) {
					return fmt.Errorf("Error reading fortiguard_category: %v", err)
				}
			}
		}
	}

	if err = d.Set("youtube", flattenVideofilterProfileYoutube(o["youtube"], d, "youtube", sv)); err != nil {
		if !fortiAPIPatch(o["youtube"]) {
			return fmt.Errorf("Error reading youtube: %v", err)
		}
	}

	if err = d.Set("vimeo", flattenVideofilterProfileVimeo(o["vimeo"], d, "vimeo", sv)); err != nil {
		if !fortiAPIPatch(o["vimeo"]) {
			return fmt.Errorf("Error reading vimeo: %v", err)
		}
	}

	if err = d.Set("dailymotion", flattenVideofilterProfileDailymotion(o["dailymotion"], d, "dailymotion", sv)); err != nil {
		if !fortiAPIPatch(o["dailymotion"]) {
			return fmt.Errorf("Error reading dailymotion: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenVideofilterProfileReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group", sv)); err != nil {
		if !fortiAPIPatch(o["replacemsg-group"]) {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	return nil
}

func flattenVideofilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVideofilterProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFilters(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandVideofilterProfileFiltersId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comment"], _ = expandVideofilterProfileFiltersComment(d, i["comment"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["comment"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandVideofilterProfileFiltersType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyword"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["keyword"], _ = expandVideofilterProfileFiltersKeyword(d, i["keyword"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["keyword"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category"], _ = expandVideofilterProfileFiltersCategory(d, i["category"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["category"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "channel"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["channel"], _ = expandVideofilterProfileFiltersChannel(d, i["channel"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["channel"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandVideofilterProfileFiltersAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log"], _ = expandVideofilterProfileFiltersLog(d, i["log"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVideofilterProfileFiltersId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFiltersComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFiltersType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFiltersKeyword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFiltersCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFiltersChannel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFiltersAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFiltersLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileDefaultAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileYoutubeChannelFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFortiguardCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "filters"
	if _, ok := d.GetOk(pre_append); ok {
		result["filters"], _ = expandVideofilterProfileFortiguardCategoryFilters(d, i["filters"], pre_append, sv)
	} else {
		result["filters"] = make([]string, 0)
	}

	return result, nil
}

func expandVideofilterProfileFortiguardCategoryFilters(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandVideofilterProfileFortiguardCategoryFiltersId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandVideofilterProfileFortiguardCategoryFiltersAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category-id"], _ = expandVideofilterProfileFortiguardCategoryFiltersCategoryId(d, i["category_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["category-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log"], _ = expandVideofilterProfileFortiguardCategoryFiltersLog(d, i["log"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVideofilterProfileFortiguardCategoryFiltersId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFortiguardCategoryFiltersAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFortiguardCategoryFiltersCategoryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileFortiguardCategoryFiltersLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileYoutube(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileVimeo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileDailymotion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVideofilterProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandVideofilterProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandVideofilterProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("filters"); ok || d.HasChange("filters") {
		t, err := expandVideofilterProfileFilters(d, v, "filters", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filters"] = t
		}
	}

	if v, ok := d.GetOk("default_action"); ok {
		t, err := expandVideofilterProfileDefaultAction(d, v, "default_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-action"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok {
		t, err := expandVideofilterProfileLog(d, v, "log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOkExists("youtube_channel_filter"); ok {
		t, err := expandVideofilterProfileYoutubeChannelFilter(d, v, "youtube_channel_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube-channel-filter"] = t
		}
	} else if d.HasChange("youtube_channel_filter") {
		obj["youtube-channel-filter"] = nil
	}

	if v, ok := d.GetOk("fortiguard_category"); ok {
		t, err := expandVideofilterProfileFortiguardCategory(d, v, "fortiguard_category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-category"] = t
		}
	}

	if v, ok := d.GetOk("youtube"); ok {
		t, err := expandVideofilterProfileYoutube(d, v, "youtube", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube"] = t
		}
	}

	if v, ok := d.GetOk("vimeo"); ok {
		t, err := expandVideofilterProfileVimeo(d, v, "vimeo", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vimeo"] = t
		}
	}

	if v, ok := d.GetOk("dailymotion"); ok {
		t, err := expandVideofilterProfileDailymotion(d, v, "dailymotion", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dailymotion"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok {
		t, err := expandVideofilterProfileReplacemsgGroup(d, v, "replacemsg_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	} else if d.HasChange("replacemsg_group") {
		obj["replacemsg-group"] = nil
	}

	return &obj, nil
}
