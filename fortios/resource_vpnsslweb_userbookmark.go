// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure SSL VPN user bookmark.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 101),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"custom_lang": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
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
						"additional_params": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
						},
						"listening_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"remote_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"show_status_window": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						"form_data": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
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

func resourceVpnSslWebUserBookmarkCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnSslWebUserBookmark(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebUserBookmark resource while getting object: %v", err)
	}

	o, err := c.CreateVpnSslWebUserBookmark(obj, vdomparam)

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

	obj, err := getObjectVpnSslWebUserBookmark(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebUserBookmark resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnSslWebUserBookmark(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteVpnSslWebUserBookmark(mkey, vdomparam)
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

	o, err := c.ReadVpnSslWebUserBookmark(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebUserBookmark resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslWebUserBookmark(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebUserBookmark resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslWebUserBookmarkName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkCustomLang(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarks(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenVpnSslWebUserBookmarkBookmarksName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if cur_v, ok := i["apptype"]; ok {
			tmp["apptype"] = flattenVpnSslWebUserBookmarkBookmarksApptype(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if cur_v, ok := i["url"]; ok {
			tmp["url"] = flattenVpnSslWebUserBookmarkBookmarksUrl(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if cur_v, ok := i["host"]; ok {
			tmp["host"] = flattenVpnSslWebUserBookmarkBookmarksHost(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if cur_v, ok := i["folder"]; ok {
			tmp["folder"] = flattenVpnSslWebUserBookmarkBookmarksFolder(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if cur_v, ok := i["domain"]; ok {
			tmp["domain"] = flattenVpnSslWebUserBookmarkBookmarksDomain(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_params"
		if cur_v, ok := i["additional-params"]; ok {
			tmp["additional_params"] = flattenVpnSslWebUserBookmarkBookmarksAdditionalParams(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listening_port"
		if cur_v, ok := i["listening-port"]; ok {
			tmp["listening_port"] = flattenVpnSslWebUserBookmarkBookmarksListeningPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_port"
		if cur_v, ok := i["remote-port"]; ok {
			tmp["remote_port"] = flattenVpnSslWebUserBookmarkBookmarksRemotePort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "show_status_window"
		if cur_v, ok := i["show-status-window"]; ok {
			tmp["show_status_window"] = flattenVpnSslWebUserBookmarkBookmarksShowStatusWindow(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if cur_v, ok := i["description"]; ok {
			tmp["description"] = flattenVpnSslWebUserBookmarkBookmarksDescription(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if cur_v, ok := i["keyboard-layout"]; ok {
			tmp["keyboard_layout"] = flattenVpnSslWebUserBookmarkBookmarksKeyboardLayout(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_layout"
		if cur_v, ok := i["server-layout"]; ok {
			tmp["server_layout"] = flattenVpnSslWebUserBookmarkBookmarksServerLayout(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if cur_v, ok := i["security"]; ok {
			tmp["security"] = flattenVpnSslWebUserBookmarkBookmarksSecurity(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if cur_v, ok := i["send-preconnection-id"]; ok {
			tmp["send_preconnection_id"] = flattenVpnSslWebUserBookmarkBookmarksSendPreconnectionId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if cur_v, ok := i["preconnection-id"]; ok {
			tmp["preconnection_id"] = flattenVpnSslWebUserBookmarkBookmarksPreconnectionId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if cur_v, ok := i["preconnection-blob"]; ok {
			tmp["preconnection_blob"] = flattenVpnSslWebUserBookmarkBookmarksPreconnectionBlob(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if cur_v, ok := i["load-balancing-info"]; ok {
			tmp["load_balancing_info"] = flattenVpnSslWebUserBookmarkBookmarksLoadBalancingInfo(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if cur_v, ok := i["restricted-admin"]; ok {
			tmp["restricted_admin"] = flattenVpnSslWebUserBookmarkBookmarksRestrictedAdmin(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if cur_v, ok := i["port"]; ok {
			tmp["port"] = flattenVpnSslWebUserBookmarkBookmarksPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if cur_v, ok := i["logon-user"]; ok {
			tmp["logon_user"] = flattenVpnSslWebUserBookmarkBookmarksLogonUser(cur_v, d, pre_append, sv)
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
			tmp["color_depth"] = flattenVpnSslWebUserBookmarkBookmarksColorDepth(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if cur_v, ok := i["sso"]; ok {
			tmp["sso"] = flattenVpnSslWebUserBookmarkBookmarksSso(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "form_data"
		if cur_v, ok := i["form-data"]; ok {
			tmp["form_data"] = flattenVpnSslWebUserBookmarkBookmarksFormData(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential"
		if cur_v, ok := i["sso-credential"]; ok {
			tmp["sso_credential"] = flattenVpnSslWebUserBookmarkBookmarksSsoCredential(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_username"
		if cur_v, ok := i["sso-username"]; ok {
			tmp["sso_username"] = flattenVpnSslWebUserBookmarkBookmarksSsoUsername(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_password"
		if _, ok := i["sso-password"]; ok {
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["sso_password"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential_sent_once"
		if cur_v, ok := i["sso-credential-sent-once"]; ok {
			tmp["sso_credential_sent_once"] = flattenVpnSslWebUserBookmarkBookmarksSsoCredentialSentOnce(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if cur_v, ok := i["width"]; ok {
			tmp["width"] = flattenVpnSslWebUserBookmarkBookmarksWidth(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if cur_v, ok := i["height"]; ok {
			tmp["height"] = flattenVpnSslWebUserBookmarkBookmarksHeight(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vnc_keyboard_layout"
		if cur_v, ok := i["vnc-keyboard-layout"]; ok {
			tmp["vnc_keyboard_layout"] = flattenVpnSslWebUserBookmarkBookmarksVncKeyboardLayout(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslWebUserBookmarkBookmarksName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksApptype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksFolder(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksAdditionalParams(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksListeningPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnSslWebUserBookmarkBookmarksRemotePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnSslWebUserBookmarkBookmarksShowStatusWindow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksKeyboardLayout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksServerLayout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSecurity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSendPreconnectionId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksPreconnectionId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnSslWebUserBookmarkBookmarksPreconnectionBlob(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksLoadBalancingInfo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksRestrictedAdmin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnSslWebUserBookmarkBookmarksLogonUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksColorDepth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSso(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksFormData(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenVpnSslWebUserBookmarkBookmarksFormDataName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if cur_v, ok := i["value"]; ok {
			tmp["value"] = flattenVpnSslWebUserBookmarkBookmarksFormDataValue(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslWebUserBookmarkBookmarksFormDataName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksFormDataValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSsoCredential(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSsoUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksSsoCredentialSentOnce(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebUserBookmarkBookmarksWidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnSslWebUserBookmarkBookmarksHeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnSslWebUserBookmarkBookmarksVncKeyboardLayout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnSslWebUserBookmark(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenVpnSslWebUserBookmarkName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("custom_lang", flattenVpnSslWebUserBookmarkCustomLang(o["custom-lang"], d, "custom_lang", sv)); err != nil {
		if !fortiAPIPatch(o["custom-lang"]) {
			return fmt.Errorf("Error reading custom_lang: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("bookmarks", flattenVpnSslWebUserBookmarkBookmarks(o["bookmarks"], d, "bookmarks", sv)); err != nil {
			if !fortiAPIPatch(o["bookmarks"]) {
				return fmt.Errorf("Error reading bookmarks: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("bookmarks"); ok {
			if err = d.Set("bookmarks", flattenVpnSslWebUserBookmarkBookmarks(o["bookmarks"], d, "bookmarks", sv)); err != nil {
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
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnSslWebUserBookmarkName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkCustomLang(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarks(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandVpnSslWebUserBookmarkBookmarksName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["apptype"], _ = expandVpnSslWebUserBookmarkBookmarksApptype(d, i["apptype"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["url"], _ = expandVpnSslWebUserBookmarkBookmarksUrl(d, i["url"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["url"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["host"], _ = expandVpnSslWebUserBookmarkBookmarksHost(d, i["host"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["host"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["folder"], _ = expandVpnSslWebUserBookmarkBookmarksFolder(d, i["folder"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["folder"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["domain"], _ = expandVpnSslWebUserBookmarkBookmarksDomain(d, i["domain"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["domain"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_params"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["additional-params"], _ = expandVpnSslWebUserBookmarkBookmarksAdditionalParams(d, i["additional_params"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["additional-params"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listening_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["listening-port"], _ = expandVpnSslWebUserBookmarkBookmarksListeningPort(d, i["listening_port"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["listening-port"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remote-port"], _ = expandVpnSslWebUserBookmarkBookmarksRemotePort(d, i["remote_port"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["remote-port"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "show_status_window"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["show-status-window"], _ = expandVpnSslWebUserBookmarkBookmarksShowStatusWindow(d, i["show_status_window"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["show-status-window"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandVpnSslWebUserBookmarkBookmarksDescription(d, i["description"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["description"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["keyboard-layout"], _ = expandVpnSslWebUserBookmarkBookmarksKeyboardLayout(d, i["keyboard_layout"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_layout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["server-layout"], _ = expandVpnSslWebUserBookmarkBookmarksServerLayout(d, i["server_layout"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["server-layout"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["security"], _ = expandVpnSslWebUserBookmarkBookmarksSecurity(d, i["security"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["send-preconnection-id"], _ = expandVpnSslWebUserBookmarkBookmarksSendPreconnectionId(d, i["send_preconnection_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["preconnection-id"], _ = expandVpnSslWebUserBookmarkBookmarksPreconnectionId(d, i["preconnection_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["preconnection-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["preconnection-blob"], _ = expandVpnSslWebUserBookmarkBookmarksPreconnectionBlob(d, i["preconnection_blob"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["preconnection-blob"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["load-balancing-info"], _ = expandVpnSslWebUserBookmarkBookmarksLoadBalancingInfo(d, i["load_balancing_info"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["load-balancing-info"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["restricted-admin"], _ = expandVpnSslWebUserBookmarkBookmarksRestrictedAdmin(d, i["restricted_admin"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandVpnSslWebUserBookmarkBookmarksPort(d, i["port"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["port"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["logon-user"], _ = expandVpnSslWebUserBookmarkBookmarksLogonUser(d, i["logon_user"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["logon-user"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_password"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["logon-password"], _ = expandVpnSslWebUserBookmarkBookmarksLogonPassword(d, i["logon_password"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["logon-password"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color_depth"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["color-depth"], _ = expandVpnSslWebUserBookmarkBookmarksColorDepth(d, i["color_depth"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sso"], _ = expandVpnSslWebUserBookmarkBookmarksSso(d, i["sso"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "form_data"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["form-data"], _ = expandVpnSslWebUserBookmarkBookmarksFormData(d, i["form_data"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["form-data"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sso-credential"], _ = expandVpnSslWebUserBookmarkBookmarksSsoCredential(d, i["sso_credential"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_username"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sso-username"], _ = expandVpnSslWebUserBookmarkBookmarksSsoUsername(d, i["sso_username"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["sso-username"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_password"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sso-password"], _ = expandVpnSslWebUserBookmarkBookmarksSsoPassword(d, i["sso_password"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["sso-password"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential_sent_once"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sso-credential-sent-once"], _ = expandVpnSslWebUserBookmarkBookmarksSsoCredentialSentOnce(d, i["sso_credential_sent_once"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["width"], _ = expandVpnSslWebUserBookmarkBookmarksWidth(d, i["width"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["height"], _ = expandVpnSslWebUserBookmarkBookmarksHeight(d, i["height"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vnc_keyboard_layout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vnc-keyboard-layout"], _ = expandVpnSslWebUserBookmarkBookmarksVncKeyboardLayout(d, i["vnc_keyboard_layout"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebUserBookmarkBookmarksName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksApptype(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksFolder(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksAdditionalParams(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksListeningPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksRemotePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksShowStatusWindow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksKeyboardLayout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksServerLayout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSecurity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSendPreconnectionId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksPreconnectionId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksPreconnectionBlob(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksLoadBalancingInfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksRestrictedAdmin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksLogonUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksLogonPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksColorDepth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSso(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksFormData(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandVpnSslWebUserBookmarkBookmarksFormDataName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandVpnSslWebUserBookmarkBookmarksFormDataValue(d, i["value"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["value"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebUserBookmarkBookmarksFormDataName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksFormDataValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSsoCredential(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSsoUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSsoPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksSsoCredentialSentOnce(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksWidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksHeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebUserBookmarkBookmarksVncKeyboardLayout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslWebUserBookmark(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandVpnSslWebUserBookmarkName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("custom_lang"); ok {
		t, err := expandVpnSslWebUserBookmarkCustomLang(d, v, "custom_lang", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-lang"] = t
		}
	} else if d.HasChange("custom_lang") {
		obj["custom-lang"] = nil
	}

	if v, ok := d.GetOk("bookmarks"); ok || d.HasChange("bookmarks") {
		t, err := expandVpnSslWebUserBookmarkBookmarks(d, v, "bookmarks", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bookmarks"] = t
		}
	}

	return &obj, nil
}
