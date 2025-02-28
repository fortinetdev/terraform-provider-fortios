// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure ztna web-portal bookmark.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaWebPortalBookmark() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaWebPortalBookmarkCreate,
		Read:   resourceZtnaWebPortalBookmarkRead,
		Update: resourceZtnaWebPortalBookmarkUpdate,
		Delete: resourceZtnaWebPortalBookmarkDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"users": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"groups": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"bookmarks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"apptype": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
						},
						"host": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
						},
						"folder": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
						},
						"domain": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
						},
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
						},
						"keyboard_layout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"security": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"send_preconnection_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"preconnection_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"preconnection_blob": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),
							Optional:     true,
						},
						"load_balancing_info": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),
							Optional:     true,
						},
						"restricted_admin": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"logon_user": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"logon_password": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
							Sensitive:    true,
						},
						"color_depth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sso": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"width": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"height": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"vnc_keyboard_layout": &schema.Schema{
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
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceZtnaWebPortalBookmarkCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectZtnaWebPortalBookmark(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebPortalBookmark resource while getting object: %v", err)
	}

	o, err := c.CreateZtnaWebPortalBookmark(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebPortalBookmark resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ZtnaWebPortalBookmark")
	}

	return resourceZtnaWebPortalBookmarkRead(d, m)
}

func resourceZtnaWebPortalBookmarkUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectZtnaWebPortalBookmark(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebPortalBookmark resource while getting object: %v", err)
	}

	o, err := c.UpdateZtnaWebPortalBookmark(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebPortalBookmark resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ZtnaWebPortalBookmark")
	}

	return resourceZtnaWebPortalBookmarkRead(d, m)
}

func resourceZtnaWebPortalBookmarkDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteZtnaWebPortalBookmark(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaWebPortalBookmark resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaWebPortalBookmarkRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadZtnaWebPortalBookmark(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebPortalBookmark resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaWebPortalBookmark(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebPortalBookmark resource from API: %v", err)
	}
	return nil
}

func flattenZtnaWebPortalBookmarkName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkUsers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenZtnaWebPortalBookmarkUsersName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenZtnaWebPortalBookmarkUsersName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkGroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenZtnaWebPortalBookmarkGroupsName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenZtnaWebPortalBookmarkGroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarks(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenZtnaWebPortalBookmarkBookmarksName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if cur_v, ok := i["apptype"]; ok {
			tmp["apptype"] = flattenZtnaWebPortalBookmarkBookmarksApptype(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if cur_v, ok := i["url"]; ok {
			tmp["url"] = flattenZtnaWebPortalBookmarkBookmarksUrl(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if cur_v, ok := i["host"]; ok {
			tmp["host"] = flattenZtnaWebPortalBookmarkBookmarksHost(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if cur_v, ok := i["folder"]; ok {
			tmp["folder"] = flattenZtnaWebPortalBookmarkBookmarksFolder(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if cur_v, ok := i["domain"]; ok {
			tmp["domain"] = flattenZtnaWebPortalBookmarkBookmarksDomain(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if cur_v, ok := i["description"]; ok {
			tmp["description"] = flattenZtnaWebPortalBookmarkBookmarksDescription(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if cur_v, ok := i["keyboard-layout"]; ok {
			tmp["keyboard_layout"] = flattenZtnaWebPortalBookmarkBookmarksKeyboardLayout(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if cur_v, ok := i["security"]; ok {
			tmp["security"] = flattenZtnaWebPortalBookmarkBookmarksSecurity(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if cur_v, ok := i["send-preconnection-id"]; ok {
			tmp["send_preconnection_id"] = flattenZtnaWebPortalBookmarkBookmarksSendPreconnectionId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if cur_v, ok := i["preconnection-id"]; ok {
			tmp["preconnection_id"] = flattenZtnaWebPortalBookmarkBookmarksPreconnectionId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if cur_v, ok := i["preconnection-blob"]; ok {
			tmp["preconnection_blob"] = flattenZtnaWebPortalBookmarkBookmarksPreconnectionBlob(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if cur_v, ok := i["load-balancing-info"]; ok {
			tmp["load_balancing_info"] = flattenZtnaWebPortalBookmarkBookmarksLoadBalancingInfo(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if cur_v, ok := i["restricted-admin"]; ok {
			tmp["restricted_admin"] = flattenZtnaWebPortalBookmarkBookmarksRestrictedAdmin(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if cur_v, ok := i["port"]; ok {
			tmp["port"] = flattenZtnaWebPortalBookmarkBookmarksPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if cur_v, ok := i["logon-user"]; ok {
			tmp["logon_user"] = flattenZtnaWebPortalBookmarkBookmarksLogonUser(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_password"
		if _, ok := i["logon-password"]; ok {
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["logon_password"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color_depth"
		if cur_v, ok := i["color-depth"]; ok {
			tmp["color_depth"] = flattenZtnaWebPortalBookmarkBookmarksColorDepth(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if cur_v, ok := i["sso"]; ok {
			tmp["sso"] = flattenZtnaWebPortalBookmarkBookmarksSso(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if cur_v, ok := i["width"]; ok {
			tmp["width"] = flattenZtnaWebPortalBookmarkBookmarksWidth(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if cur_v, ok := i["height"]; ok {
			tmp["height"] = flattenZtnaWebPortalBookmarkBookmarksHeight(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vnc_keyboard_layout"
		if cur_v, ok := i["vnc-keyboard-layout"]; ok {
			tmp["vnc_keyboard_layout"] = flattenZtnaWebPortalBookmarkBookmarksVncKeyboardLayout(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenZtnaWebPortalBookmarkBookmarksName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksApptype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksFolder(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksKeyboardLayout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksSecurity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksSendPreconnectionId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksPreconnectionId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebPortalBookmarkBookmarksPreconnectionBlob(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksLoadBalancingInfo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksRestrictedAdmin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebPortalBookmarkBookmarksLogonUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksColorDepth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksSso(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalBookmarkBookmarksWidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebPortalBookmarkBookmarksHeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebPortalBookmarkBookmarksVncKeyboardLayout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectZtnaWebPortalBookmark(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenZtnaWebPortalBookmarkName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("users", flattenZtnaWebPortalBookmarkUsers(o["users"], d, "users", sv)); err != nil {
			if !fortiAPIPatch(o["users"]) {
				return fmt.Errorf("Error reading users: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("users"); ok {
			if err = d.Set("users", flattenZtnaWebPortalBookmarkUsers(o["users"], d, "users", sv)); err != nil {
				if !fortiAPIPatch(o["users"]) {
					return fmt.Errorf("Error reading users: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("groups", flattenZtnaWebPortalBookmarkGroups(o["groups"], d, "groups", sv)); err != nil {
			if !fortiAPIPatch(o["groups"]) {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("groups"); ok {
			if err = d.Set("groups", flattenZtnaWebPortalBookmarkGroups(o["groups"], d, "groups", sv)); err != nil {
				if !fortiAPIPatch(o["groups"]) {
					return fmt.Errorf("Error reading groups: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("bookmarks", flattenZtnaWebPortalBookmarkBookmarks(o["bookmarks"], d, "bookmarks", sv)); err != nil {
			if !fortiAPIPatch(o["bookmarks"]) {
				return fmt.Errorf("Error reading bookmarks: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("bookmarks"); ok {
			if err = d.Set("bookmarks", flattenZtnaWebPortalBookmarkBookmarks(o["bookmarks"], d, "bookmarks", sv)); err != nil {
				if !fortiAPIPatch(o["bookmarks"]) {
					return fmt.Errorf("Error reading bookmarks: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenZtnaWebPortalBookmarkFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandZtnaWebPortalBookmarkName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandZtnaWebPortalBookmarkUsersName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandZtnaWebPortalBookmarkUsersName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandZtnaWebPortalBookmarkGroupsName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandZtnaWebPortalBookmarkGroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarks(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandZtnaWebPortalBookmarkBookmarksName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["apptype"], _ = expandZtnaWebPortalBookmarkBookmarksApptype(d, i["apptype"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["url"], _ = expandZtnaWebPortalBookmarkBookmarksUrl(d, i["url"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["url"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["host"], _ = expandZtnaWebPortalBookmarkBookmarksHost(d, i["host"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["host"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["folder"], _ = expandZtnaWebPortalBookmarkBookmarksFolder(d, i["folder"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["folder"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["domain"], _ = expandZtnaWebPortalBookmarkBookmarksDomain(d, i["domain"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["domain"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandZtnaWebPortalBookmarkBookmarksDescription(d, i["description"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["description"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["keyboard-layout"], _ = expandZtnaWebPortalBookmarkBookmarksKeyboardLayout(d, i["keyboard_layout"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["security"], _ = expandZtnaWebPortalBookmarkBookmarksSecurity(d, i["security"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["send-preconnection-id"], _ = expandZtnaWebPortalBookmarkBookmarksSendPreconnectionId(d, i["send_preconnection_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["preconnection-id"], _ = expandZtnaWebPortalBookmarkBookmarksPreconnectionId(d, i["preconnection_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["preconnection-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["preconnection-blob"], _ = expandZtnaWebPortalBookmarkBookmarksPreconnectionBlob(d, i["preconnection_blob"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["preconnection-blob"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["load-balancing-info"], _ = expandZtnaWebPortalBookmarkBookmarksLoadBalancingInfo(d, i["load_balancing_info"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["load-balancing-info"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["restricted-admin"], _ = expandZtnaWebPortalBookmarkBookmarksRestrictedAdmin(d, i["restricted_admin"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandZtnaWebPortalBookmarkBookmarksPort(d, i["port"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["port"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["logon-user"], _ = expandZtnaWebPortalBookmarkBookmarksLogonUser(d, i["logon_user"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["logon-user"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_password"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["logon-password"], _ = expandZtnaWebPortalBookmarkBookmarksLogonPassword(d, i["logon_password"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["logon-password"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color_depth"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["color-depth"], _ = expandZtnaWebPortalBookmarkBookmarksColorDepth(d, i["color_depth"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sso"], _ = expandZtnaWebPortalBookmarkBookmarksSso(d, i["sso"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["width"], _ = expandZtnaWebPortalBookmarkBookmarksWidth(d, i["width"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["width"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["height"], _ = expandZtnaWebPortalBookmarkBookmarksHeight(d, i["height"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["height"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vnc_keyboard_layout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vnc-keyboard-layout"], _ = expandZtnaWebPortalBookmarkBookmarksVncKeyboardLayout(d, i["vnc_keyboard_layout"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandZtnaWebPortalBookmarkBookmarksName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksApptype(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksFolder(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksKeyboardLayout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksSecurity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksSendPreconnectionId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksPreconnectionId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksPreconnectionBlob(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksLoadBalancingInfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksRestrictedAdmin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksLogonUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksLogonPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksColorDepth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksSso(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksWidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksHeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalBookmarkBookmarksVncKeyboardLayout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectZtnaWebPortalBookmark(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandZtnaWebPortalBookmarkName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok || d.HasChange("users") {
		t, err := expandZtnaWebPortalBookmarkUsers(d, v, "users", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		t, err := expandZtnaWebPortalBookmarkGroups(d, v, "groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("bookmarks"); ok || d.HasChange("bookmarks") {
		t, err := expandZtnaWebPortalBookmarkBookmarks(d, v, "bookmarks", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bookmarks"] = t
		}
	}

	return &obj, nil
}
