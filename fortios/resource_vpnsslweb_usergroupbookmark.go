// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure SSL VPN user group bookmark.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceVpnSslWebUserGroupBookmark() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslWebUserGroupBookmarkCreate,
		Read:   resourceVpnSslWebUserGroupBookmarkRead,
		Update: resourceVpnSslWebUserGroupBookmarkUpdate,
		Delete: resourceVpnSslWebUserGroupBookmarkDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 64),
				ForceNew:     true,
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
						"domain": &schema.Schema{
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
						"keyboard_layout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"send_preconnection_id": &schema.Schema{
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
						"restricted_admin": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"width": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"height": &schema.Schema{
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

func resourceVpnSslWebUserGroupBookmarkCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnSslWebUserGroupBookmark(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebUserGroupBookmark resource while getting object: %v", err)
	}

	o, err := c.CreateVpnSslWebUserGroupBookmark(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebUserGroupBookmark resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnSslWebUserGroupBookmark")
	}

	return resourceVpnSslWebUserGroupBookmarkRead(d, m)
}

func resourceVpnSslWebUserGroupBookmarkUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnSslWebUserGroupBookmark(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebUserGroupBookmark resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnSslWebUserGroupBookmark(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebUserGroupBookmark resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnSslWebUserGroupBookmark")
	}

	return resourceVpnSslWebUserGroupBookmarkRead(d, m)
}

func resourceVpnSslWebUserGroupBookmarkDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteVpnSslWebUserGroupBookmark(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting VpnSslWebUserGroupBookmark resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebUserGroupBookmarkRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadVpnSslWebUserGroupBookmark(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebUserGroupBookmark resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslWebUserGroupBookmark(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebUserGroupBookmark resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslWebUserGroupBookmarkName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarks(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenVpnSslWebUserGroupBookmarkBookmarksName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if _, ok := i["apptype"]; ok {

			tmp["apptype"] = flattenVpnSslWebUserGroupBookmarkBookmarksApptype(i["apptype"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := i["url"]; ok {

			tmp["url"] = flattenVpnSslWebUserGroupBookmarkBookmarksUrl(i["url"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := i["host"]; ok {

			tmp["host"] = flattenVpnSslWebUserGroupBookmarkBookmarksHost(i["host"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if _, ok := i["folder"]; ok {

			tmp["folder"] = flattenVpnSslWebUserGroupBookmarkBookmarksFolder(i["folder"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {

			tmp["domain"] = flattenVpnSslWebUserGroupBookmarkBookmarksDomain(i["domain"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_params"
		if _, ok := i["additional-params"]; ok {

			tmp["additional_params"] = flattenVpnSslWebUserGroupBookmarkBookmarksAdditionalParams(i["additional-params"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listening_port"
		if _, ok := i["listening-port"]; ok {

			tmp["listening_port"] = flattenVpnSslWebUserGroupBookmarkBookmarksListeningPort(i["listening-port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_port"
		if _, ok := i["remote-port"]; ok {

			tmp["remote_port"] = flattenVpnSslWebUserGroupBookmarkBookmarksRemotePort(i["remote-port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "show_status_window"
		if _, ok := i["show-status-window"]; ok {

			tmp["show_status_window"] = flattenVpnSslWebUserGroupBookmarkBookmarksShowStatusWindow(i["show-status-window"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenVpnSslWebUserGroupBookmarkBookmarksDescription(i["description"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if _, ok := i["keyboard-layout"]; ok {

			tmp["keyboard_layout"] = flattenVpnSslWebUserGroupBookmarkBookmarksKeyboardLayout(i["keyboard-layout"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_layout"
		if _, ok := i["server-layout"]; ok {

			tmp["server_layout"] = flattenVpnSslWebUserGroupBookmarkBookmarksServerLayout(i["server-layout"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := i["security"]; ok {

			tmp["security"] = flattenVpnSslWebUserGroupBookmarkBookmarksSecurity(i["security"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if _, ok := i["send-preconnection-id"]; ok {

			tmp["send_preconnection_id"] = flattenVpnSslWebUserGroupBookmarkBookmarksSendPreconnectionId(i["send-preconnection-id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if _, ok := i["preconnection-id"]; ok {

			tmp["preconnection_id"] = flattenVpnSslWebUserGroupBookmarkBookmarksPreconnectionId(i["preconnection-id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if _, ok := i["preconnection-blob"]; ok {

			tmp["preconnection_blob"] = flattenVpnSslWebUserGroupBookmarkBookmarksPreconnectionBlob(i["preconnection-blob"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if _, ok := i["load-balancing-info"]; ok {

			tmp["load_balancing_info"] = flattenVpnSslWebUserGroupBookmarkBookmarksLoadBalancingInfo(i["load-balancing-info"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if _, ok := i["restricted-admin"]; ok {

			tmp["restricted_admin"] = flattenVpnSslWebUserGroupBookmarkBookmarksRestrictedAdmin(i["restricted-admin"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = flattenVpnSslWebUserGroupBookmarkBookmarksPort(i["port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if _, ok := i["logon-user"]; ok {

			tmp["logon_user"] = flattenVpnSslWebUserGroupBookmarkBookmarksLogonUser(i["logon-user"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_password"
		if _, ok := i["logon-password"]; ok {

			tmp["logon_password"] = flattenVpnSslWebUserGroupBookmarkBookmarksLogonPassword(i["logon-password"], d, pre_append, sv)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["logon_password"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color_depth"
		if _, ok := i["color-depth"]; ok {

			tmp["color_depth"] = flattenVpnSslWebUserGroupBookmarkBookmarksColorDepth(i["color-depth"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if _, ok := i["sso"]; ok {

			tmp["sso"] = flattenVpnSslWebUserGroupBookmarkBookmarksSso(i["sso"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "form_data"
		if _, ok := i["form-data"]; ok {

			tmp["form_data"] = flattenVpnSslWebUserGroupBookmarkBookmarksFormData(i["form-data"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential"
		if _, ok := i["sso-credential"]; ok {

			tmp["sso_credential"] = flattenVpnSslWebUserGroupBookmarkBookmarksSsoCredential(i["sso-credential"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_username"
		if _, ok := i["sso-username"]; ok {

			tmp["sso_username"] = flattenVpnSslWebUserGroupBookmarkBookmarksSsoUsername(i["sso-username"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_password"
		if _, ok := i["sso-password"]; ok {

			tmp["sso_password"] = flattenVpnSslWebUserGroupBookmarkBookmarksSsoPassword(i["sso-password"], d, pre_append, sv)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["sso_password"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential_sent_once"
		if _, ok := i["sso-credential-sent-once"]; ok {

			tmp["sso_credential_sent_once"] = flattenVpnSslWebUserGroupBookmarkBookmarksSsoCredentialSentOnce(i["sso-credential-sent-once"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := i["width"]; ok {

			tmp["width"] = flattenVpnSslWebUserGroupBookmarkBookmarksWidth(i["width"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := i["height"]; ok {

			tmp["height"] = flattenVpnSslWebUserGroupBookmarkBookmarksHeight(i["height"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslWebUserGroupBookmarkBookmarksName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksApptype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksFolder(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksAdditionalParams(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksListeningPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksRemotePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksShowStatusWindow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksKeyboardLayout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksServerLayout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksSecurity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksSendPreconnectionId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksPreconnectionId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksPreconnectionBlob(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksLoadBalancingInfo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksRestrictedAdmin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksLogonUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksLogonPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksColorDepth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksSso(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksFormData(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenVpnSslWebUserGroupBookmarkBookmarksFormDataName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {

			tmp["value"] = flattenVpnSslWebUserGroupBookmarkBookmarksFormDataValue(i["value"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenVpnSslWebUserGroupBookmarkBookmarksFormDataName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksFormDataValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksSsoCredential(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksSsoUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksSsoPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksSsoCredentialSentOnce(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksWidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserGroupBookmarkBookmarksHeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnSslWebUserGroupBookmark(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenVpnSslWebUserGroupBookmarkName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("bookmarks", flattenVpnSslWebUserGroupBookmarkBookmarks(o["bookmarks"], d, "bookmarks", sv)); err != nil {
			if !fortiAPIPatch(o["bookmarks"]) {
				return fmt.Errorf("Error reading bookmarks: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("bookmarks"); ok {
			if err = d.Set("bookmarks", flattenVpnSslWebUserGroupBookmarkBookmarks(o["bookmarks"], d, "bookmarks", sv)); err != nil {
				if !fortiAPIPatch(o["bookmarks"]) {
					return fmt.Errorf("Error reading bookmarks: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenVpnSslWebUserGroupBookmarkFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnSslWebUserGroupBookmarkName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarks(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandVpnSslWebUserGroupBookmarkBookmarksName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["apptype"], _ = expandVpnSslWebUserGroupBookmarkBookmarksApptype(d, i["apptype"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["url"], _ = expandVpnSslWebUserGroupBookmarkBookmarksUrl(d, i["url"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["host"], _ = expandVpnSslWebUserGroupBookmarkBookmarksHost(d, i["host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["folder"], _ = expandVpnSslWebUserGroupBookmarkBookmarksFolder(d, i["folder"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["domain"], _ = expandVpnSslWebUserGroupBookmarkBookmarksDomain(d, i["domain"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_params"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["additional-params"], _ = expandVpnSslWebUserGroupBookmarkBookmarksAdditionalParams(d, i["additional_params"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listening_port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["listening-port"], _ = expandVpnSslWebUserGroupBookmarkBookmarksListeningPort(d, i["listening_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["remote-port"], _ = expandVpnSslWebUserGroupBookmarkBookmarksRemotePort(d, i["remote_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "show_status_window"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["show-status-window"], _ = expandVpnSslWebUserGroupBookmarkBookmarksShowStatusWindow(d, i["show_status_window"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["description"], _ = expandVpnSslWebUserGroupBookmarkBookmarksDescription(d, i["description"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["keyboard-layout"], _ = expandVpnSslWebUserGroupBookmarkBookmarksKeyboardLayout(d, i["keyboard_layout"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_layout"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["server-layout"], _ = expandVpnSslWebUserGroupBookmarkBookmarksServerLayout(d, i["server_layout"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["security"], _ = expandVpnSslWebUserGroupBookmarkBookmarksSecurity(d, i["security"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["send-preconnection-id"], _ = expandVpnSslWebUserGroupBookmarkBookmarksSendPreconnectionId(d, i["send_preconnection_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["preconnection-id"], _ = expandVpnSslWebUserGroupBookmarkBookmarksPreconnectionId(d, i["preconnection_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["preconnection-blob"], _ = expandVpnSslWebUserGroupBookmarkBookmarksPreconnectionBlob(d, i["preconnection_blob"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["load-balancing-info"], _ = expandVpnSslWebUserGroupBookmarkBookmarksLoadBalancingInfo(d, i["load_balancing_info"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["restricted-admin"], _ = expandVpnSslWebUserGroupBookmarkBookmarksRestrictedAdmin(d, i["restricted_admin"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["port"], _ = expandVpnSslWebUserGroupBookmarkBookmarksPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["logon-user"], _ = expandVpnSslWebUserGroupBookmarkBookmarksLogonUser(d, i["logon_user"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_password"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["logon-password"], _ = expandVpnSslWebUserGroupBookmarkBookmarksLogonPassword(d, i["logon_password"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color_depth"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["color-depth"], _ = expandVpnSslWebUserGroupBookmarkBookmarksColorDepth(d, i["color_depth"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["sso"], _ = expandVpnSslWebUserGroupBookmarkBookmarksSso(d, i["sso"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "form_data"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["form-data"], _ = expandVpnSslWebUserGroupBookmarkBookmarksFormData(d, i["form_data"], pre_append, sv)
		} else {
			tmp["form-data"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["sso-credential"], _ = expandVpnSslWebUserGroupBookmarkBookmarksSsoCredential(d, i["sso_credential"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_username"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["sso-username"], _ = expandVpnSslWebUserGroupBookmarkBookmarksSsoUsername(d, i["sso_username"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_password"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["sso-password"], _ = expandVpnSslWebUserGroupBookmarkBookmarksSsoPassword(d, i["sso_password"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential_sent_once"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["sso-credential-sent-once"], _ = expandVpnSslWebUserGroupBookmarkBookmarksSsoCredentialSentOnce(d, i["sso_credential_sent_once"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["width"], _ = expandVpnSslWebUserGroupBookmarkBookmarksWidth(d, i["width"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["height"], _ = expandVpnSslWebUserGroupBookmarkBookmarksHeight(d, i["height"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksApptype(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksFolder(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksAdditionalParams(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksListeningPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksRemotePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksShowStatusWindow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksKeyboardLayout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksServerLayout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSecurity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSendPreconnectionId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksPreconnectionId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksPreconnectionBlob(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksLoadBalancingInfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksRestrictedAdmin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksLogonUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksLogonPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksColorDepth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSso(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksFormData(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandVpnSslWebUserGroupBookmarkBookmarksFormDataName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["value"], _ = expandVpnSslWebUserGroupBookmarkBookmarksFormDataValue(d, i["value"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksFormDataName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksFormDataValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSsoCredential(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSsoUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSsoPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksSsoCredentialSentOnce(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksWidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserGroupBookmarkBookmarksHeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslWebUserGroupBookmark(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandVpnSslWebUserGroupBookmarkName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("bookmarks"); ok || d.HasChange("bookmarks") {

		t, err := expandVpnSslWebUserGroupBookmarkBookmarks(d, v, "bookmarks", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bookmarks"] = t
		}
	}

	return &obj, nil
}
