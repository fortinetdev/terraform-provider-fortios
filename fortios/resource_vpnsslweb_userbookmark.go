// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure SSL VPN user bookmark.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceVpnSslWebUserBookmark() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslWebUserBookmarkCreate,
		Read:   resourceVpnSslWebUserBookmarkRead,
		Update: resourceVpnSslWebUserBookmarkUpdate,
		Delete: resourceVpnSslWebUserBookmarkDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 101),
				Required:     true,
				ForceNew:     true,
			},
			"custom_lang": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
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
							Computed:     true,
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
						"additional_params": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
						},
						"listening_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"remote_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"show_status_window": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
						},
						"server_layout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"security": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"preconnection_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
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
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
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
						"sso": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"form_data": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
										Computed:     true,
									},
									"value": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
									},
								},
							},
						},
						"sso_credential": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sso_username": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"sso_password": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
							Sensitive:    true,
						},
						"sso_credential_sent_once": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceVpnSslWebUserBookmarkCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVpnSslWebUserBookmark(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebUserBookmark resource while getting object: %v", err)
	}

	o, err := c.CreateVpnSslWebUserBookmark(obj)

	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebUserBookmark resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnSslWebUserBookmark")
	}

	return resourceVpnSslWebUserBookmarkRead(d, m)
}

func resourceVpnSslWebUserBookmarkUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVpnSslWebUserBookmark(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebUserBookmark resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnSslWebUserBookmark(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebUserBookmark resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnSslWebUserBookmark")
	}

	return resourceVpnSslWebUserBookmarkRead(d, m)
}

func resourceVpnSslWebUserBookmarkDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteVpnSslWebUserBookmark(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting VpnSslWebUserBookmark resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebUserBookmarkRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadVpnSslWebUserBookmark(mkey)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebUserBookmark resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslWebUserBookmark(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebUserBookmark resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslWebUserBookmarkName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkCustomLang(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarks(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenVpnSslWebUserBookmarkBookmarksName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if _, ok := i["apptype"]; ok {
			tmp["apptype"] = flattenVpnSslWebUserBookmarkBookmarksApptype(i["apptype"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := i["url"]; ok {
			tmp["url"] = flattenVpnSslWebUserBookmarkBookmarksUrl(i["url"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := i["host"]; ok {
			tmp["host"] = flattenVpnSslWebUserBookmarkBookmarksHost(i["host"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if _, ok := i["folder"]; ok {
			tmp["folder"] = flattenVpnSslWebUserBookmarkBookmarksFolder(i["folder"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_params"
		if _, ok := i["additional-params"]; ok {
			tmp["additional_params"] = flattenVpnSslWebUserBookmarkBookmarksAdditionalParams(i["additional-params"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listening_port"
		if _, ok := i["listening-port"]; ok {
			tmp["listening_port"] = flattenVpnSslWebUserBookmarkBookmarksListeningPort(i["listening-port"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_port"
		if _, ok := i["remote-port"]; ok {
			tmp["remote_port"] = flattenVpnSslWebUserBookmarkBookmarksRemotePort(i["remote-port"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "show_status_window"
		if _, ok := i["show-status-window"]; ok {
			tmp["show_status_window"] = flattenVpnSslWebUserBookmarkBookmarksShowStatusWindow(i["show-status-window"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			tmp["description"] = flattenVpnSslWebUserBookmarkBookmarksDescription(i["description"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_layout"
		if _, ok := i["server-layout"]; ok {
			tmp["server_layout"] = flattenVpnSslWebUserBookmarkBookmarksServerLayout(i["server-layout"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := i["security"]; ok {
			tmp["security"] = flattenVpnSslWebUserBookmarkBookmarksSecurity(i["security"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if _, ok := i["preconnection-id"]; ok {
			tmp["preconnection_id"] = flattenVpnSslWebUserBookmarkBookmarksPreconnectionId(i["preconnection-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if _, ok := i["preconnection-blob"]; ok {
			tmp["preconnection_blob"] = flattenVpnSslWebUserBookmarkBookmarksPreconnectionBlob(i["preconnection-blob"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if _, ok := i["load-balancing-info"]; ok {
			tmp["load_balancing_info"] = flattenVpnSslWebUserBookmarkBookmarksLoadBalancingInfo(i["load-balancing-info"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			tmp["port"] = flattenVpnSslWebUserBookmarkBookmarksPort(i["port"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if _, ok := i["logon-user"]; ok {
			tmp["logon_user"] = flattenVpnSslWebUserBookmarkBookmarksLogonUser(i["logon-user"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_password"
		if _, ok := i["logon-password"]; ok {
			tmp["logon_password"] = flattenVpnSslWebUserBookmarkBookmarksLogonPassword(i["logon-password"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["logon_password"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if _, ok := i["sso"]; ok {
			tmp["sso"] = flattenVpnSslWebUserBookmarkBookmarksSso(i["sso"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "form_data"
		if _, ok := i["form-data"]; ok {
			tmp["form_data"] = flattenVpnSslWebUserBookmarkBookmarksFormData(i["form-data"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential"
		if _, ok := i["sso-credential"]; ok {
			tmp["sso_credential"] = flattenVpnSslWebUserBookmarkBookmarksSsoCredential(i["sso-credential"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_username"
		if _, ok := i["sso-username"]; ok {
			tmp["sso_username"] = flattenVpnSslWebUserBookmarkBookmarksSsoUsername(i["sso-username"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_password"
		if _, ok := i["sso-password"]; ok {
			tmp["sso_password"] = flattenVpnSslWebUserBookmarkBookmarksSsoPassword(i["sso-password"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["sso_password"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential_sent_once"
		if _, ok := i["sso-credential-sent-once"]; ok {
			tmp["sso_credential_sent_once"] = flattenVpnSslWebUserBookmarkBookmarksSsoCredentialSentOnce(i["sso-credential-sent-once"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenVpnSslWebUserBookmarkBookmarksName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksApptype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksFolder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksAdditionalParams(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksListeningPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksRemotePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksShowStatusWindow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksServerLayout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksPreconnectionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksPreconnectionBlob(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksLoadBalancingInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksLogonUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksLogonPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksFormData(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenVpnSslWebUserBookmarkBookmarksFormDataName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			tmp["value"] = flattenVpnSslWebUserBookmarkBookmarksFormDataValue(i["value"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenVpnSslWebUserBookmarkBookmarksFormDataName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksFormDataValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSsoCredential(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSsoUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSsoPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSsoCredentialSentOnce(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnSslWebUserBookmark(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenVpnSslWebUserBookmarkName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("custom_lang", flattenVpnSslWebUserBookmarkCustomLang(o["custom-lang"], d, "custom_lang")); err != nil {
		if !fortiAPIPatch(o["custom-lang"]) {
			return fmt.Errorf("Error reading custom_lang: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("bookmarks", flattenVpnSslWebUserBookmarkBookmarks(o["bookmarks"], d, "bookmarks")); err != nil {
			if !fortiAPIPatch(o["bookmarks"]) {
				return fmt.Errorf("Error reading bookmarks: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("bookmarks"); ok {
			if err = d.Set("bookmarks", flattenVpnSslWebUserBookmarkBookmarks(o["bookmarks"], d, "bookmarks")); err != nil {
				if !fortiAPIPatch(o["bookmarks"]) {
					return fmt.Errorf("Error reading bookmarks: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenVpnSslWebUserBookmarkFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnSslWebUserBookmarkName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkCustomLang(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarks(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandVpnSslWebUserBookmarkBookmarksName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["apptype"], _ = expandVpnSslWebUserBookmarkBookmarksApptype(d, i["apptype"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["url"], _ = expandVpnSslWebUserBookmarkBookmarksUrl(d, i["url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["host"], _ = expandVpnSslWebUserBookmarkBookmarksHost(d, i["host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["folder"], _ = expandVpnSslWebUserBookmarkBookmarksFolder(d, i["folder"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_params"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["additional-params"], _ = expandVpnSslWebUserBookmarkBookmarksAdditionalParams(d, i["additional_params"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listening_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["listening-port"], _ = expandVpnSslWebUserBookmarkBookmarksListeningPort(d, i["listening_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remote-port"], _ = expandVpnSslWebUserBookmarkBookmarksRemotePort(d, i["remote_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "show_status_window"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["show-status-window"], _ = expandVpnSslWebUserBookmarkBookmarksShowStatusWindow(d, i["show_status_window"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandVpnSslWebUserBookmarkBookmarksDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_layout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["server-layout"], _ = expandVpnSslWebUserBookmarkBookmarksServerLayout(d, i["server_layout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["security"], _ = expandVpnSslWebUserBookmarkBookmarksSecurity(d, i["security"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["preconnection-id"], _ = expandVpnSslWebUserBookmarkBookmarksPreconnectionId(d, i["preconnection_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["preconnection-blob"], _ = expandVpnSslWebUserBookmarkBookmarksPreconnectionBlob(d, i["preconnection_blob"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["load-balancing-info"], _ = expandVpnSslWebUserBookmarkBookmarksLoadBalancingInfo(d, i["load_balancing_info"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandVpnSslWebUserBookmarkBookmarksPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["logon-user"], _ = expandVpnSslWebUserBookmarkBookmarksLogonUser(d, i["logon_user"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_password"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["logon-password"], _ = expandVpnSslWebUserBookmarkBookmarksLogonPassword(d, i["logon_password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sso"], _ = expandVpnSslWebUserBookmarkBookmarksSso(d, i["sso"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "form_data"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["form-data"], _ = expandVpnSslWebUserBookmarkBookmarksFormData(d, i["form_data"], pre_append)
		} else {
			tmp["form-data"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sso-credential"], _ = expandVpnSslWebUserBookmarkBookmarksSsoCredential(d, i["sso_credential"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_username"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sso-username"], _ = expandVpnSslWebUserBookmarkBookmarksSsoUsername(d, i["sso_username"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_password"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sso-password"], _ = expandVpnSslWebUserBookmarkBookmarksSsoPassword(d, i["sso_password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential_sent_once"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sso-credential-sent-once"], _ = expandVpnSslWebUserBookmarkBookmarksSsoCredentialSentOnce(d, i["sso_credential_sent_once"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebUserBookmarkBookmarksName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksApptype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksFolder(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksAdditionalParams(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksListeningPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksRemotePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksShowStatusWindow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksServerLayout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSecurity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksPreconnectionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksPreconnectionBlob(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksLoadBalancingInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksLogonUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksLogonPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksFormData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandVpnSslWebUserBookmarkBookmarksFormDataName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandVpnSslWebUserBookmarkBookmarksFormDataValue(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebUserBookmarkBookmarksFormDataName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksFormDataValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSsoCredential(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSsoUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSsoPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSsoCredentialSentOnce(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslWebUserBookmark(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandVpnSslWebUserBookmarkName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("custom_lang"); ok {
		t, err := expandVpnSslWebUserBookmarkCustomLang(d, v, "custom_lang")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-lang"] = t
		}
	}

	if v, ok := d.GetOk("bookmarks"); ok {
		t, err := expandVpnSslWebUserBookmarkBookmarks(d, v, "bookmarks")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bookmarks"] = t
		}
	}

	return &obj, nil
}
