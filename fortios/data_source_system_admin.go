// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure admin users.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemAdmin() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemAdminRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"wildcard": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_auth": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"peer_auth": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusthost1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusthost2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusthost3": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusthost4": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusthost5": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusthost6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusthost7": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusthost8": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusthost9": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusthost10": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_trusthost1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_trusthost2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_trusthost3": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_trusthost4": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_trusthost5": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_trusthost6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_trusthost7": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_trusthost8": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_trusthost9": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_trusthost10": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"accprofile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_remove_admin_session": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hidden": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ssh_public_key1": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"ssh_public_key2": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"ssh_public_key3": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"ssh_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"accprofile_override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"radius_vdom_override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"password_expire": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"force_password_change": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_dashboard": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"scope": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"layout_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"columns": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"widget": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"x_pos": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"y_pos": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"width": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"height": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"interface": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"region": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"industry": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"fabric_device": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"report_by": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"timeframe": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"sort_by": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"visualization": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"filters": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"key": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"two_factor": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"two_factor_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"two_factor_notification": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fortitoken": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"email_to": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sms_server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sms_custom_server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sms_phone": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"guest_auth": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"guest_usergroups": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"guest_lang": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"history0": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"history1": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"login_time": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"usr_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_login": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_failed_login": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"gui_global_menu_favorites": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"gui_vdom_menu_favorites": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"gui_new_feature_acknowledge": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemAdminRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemAdmin: type error")
	}

	o, err := c.ReadSystemAdmin(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemAdmin: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemAdmin(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemAdmin from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemAdminName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminWildcard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminRemoteAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminRemoteGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminPeerAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminPeerGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminTrusthost1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminTrusthost2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminTrusthost3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminTrusthost4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminTrusthost5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminTrusthost6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminTrusthost7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminTrusthost8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminTrusthost9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminTrusthost10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemAdminIp6Trusthost1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminIp6Trusthost2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminIp6Trusthost3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminIp6Trusthost4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminIp6Trusthost5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminIp6Trusthost6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminIp6Trusthost7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminIp6Trusthost8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminIp6Trusthost9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminIp6Trusthost10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminAccprofile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminAllowRemoveAdminSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminHidden(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminVdom(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemAdminVdomName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemAdminVdomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminSshPublicKey1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminSshPublicKey2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminSshPublicKey3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminSshCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminAccprofileOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminRadiusVdomOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminPasswordExpire(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminForcePasswordChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboard(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemAdminGuiDashboardId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenSystemAdminGuiDashboardName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scope"
		if _, ok := i["scope"]; ok {
			tmp["scope"] = dataSourceFlattenSystemAdminGuiDashboardScope(i["scope"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "layout_type"
		if _, ok := i["layout-type"]; ok {
			tmp["layout_type"] = dataSourceFlattenSystemAdminGuiDashboardLayoutType(i["layout-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "columns"
		if _, ok := i["columns"]; ok {
			tmp["columns"] = dataSourceFlattenSystemAdminGuiDashboardColumns(i["columns"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "widget"
		if _, ok := i["widget"]; ok {
			tmp["widget"] = dataSourceFlattenSystemAdminGuiDashboardWidget(i["widget"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemAdminGuiDashboardId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardScope(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardLayoutType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardColumns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardWidget(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemAdminGuiDashboardWidgetId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = dataSourceFlattenSystemAdminGuiDashboardWidgetType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "x_pos"
		if _, ok := i["x-pos"]; ok {
			tmp["x_pos"] = dataSourceFlattenSystemAdminGuiDashboardWidgetXPos(i["x-pos"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "y_pos"
		if _, ok := i["y-pos"]; ok {
			tmp["y_pos"] = dataSourceFlattenSystemAdminGuiDashboardWidgetYPos(i["y-pos"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := i["width"]; ok {
			tmp["width"] = dataSourceFlattenSystemAdminGuiDashboardWidgetWidth(i["width"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := i["height"]; ok {
			tmp["height"] = dataSourceFlattenSystemAdminGuiDashboardWidgetHeight(i["height"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenSystemAdminGuiDashboardWidgetInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "region"
		if _, ok := i["region"]; ok {
			tmp["region"] = dataSourceFlattenSystemAdminGuiDashboardWidgetRegion(i["region"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "industry"
		if _, ok := i["industry"]; ok {
			tmp["industry"] = dataSourceFlattenSystemAdminGuiDashboardWidgetIndustry(i["industry"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_device"
		if _, ok := i["fabric-device"]; ok {
			tmp["fabric_device"] = dataSourceFlattenSystemAdminGuiDashboardWidgetFabricDevice(i["fabric-device"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "title"
		if _, ok := i["title"]; ok {
			tmp["title"] = dataSourceFlattenSystemAdminGuiDashboardWidgetTitle(i["title"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "report_by"
		if _, ok := i["report-by"]; ok {
			tmp["report_by"] = dataSourceFlattenSystemAdminGuiDashboardWidgetReportBy(i["report-by"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "timeframe"
		if _, ok := i["timeframe"]; ok {
			tmp["timeframe"] = dataSourceFlattenSystemAdminGuiDashboardWidgetTimeframe(i["timeframe"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sort_by"
		if _, ok := i["sort-by"]; ok {
			tmp["sort_by"] = dataSourceFlattenSystemAdminGuiDashboardWidgetSortBy(i["sort-by"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "visualization"
		if _, ok := i["visualization"]; ok {
			tmp["visualization"] = dataSourceFlattenSystemAdminGuiDashboardWidgetVisualization(i["visualization"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filters"
		if _, ok := i["filters"]; ok {
			tmp["filters"] = dataSourceFlattenSystemAdminGuiDashboardWidgetFilters(i["filters"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemAdminGuiDashboardWidgetId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardWidgetType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardWidgetXPos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardWidgetYPos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardWidgetWidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardWidgetHeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardWidgetInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardWidgetRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardWidgetIndustry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardWidgetFabricDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardWidgetTitle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardWidgetReportBy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardWidgetTimeframe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardWidgetSortBy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardWidgetVisualization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardWidgetFilters(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemAdminGuiDashboardWidgetFiltersId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := i["key"]; ok {
			tmp["key"] = dataSourceFlattenSystemAdminGuiDashboardWidgetFiltersKey(i["key"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			tmp["value"] = dataSourceFlattenSystemAdminGuiDashboardWidgetFiltersValue(i["value"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemAdminGuiDashboardWidgetFiltersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardWidgetFiltersKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiDashboardWidgetFiltersValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminTwoFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminTwoFactorAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminTwoFactorNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminFortitoken(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminEmailTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminSmsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminSmsCustomServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminSmsPhone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuestAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuestUsergroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemAdminGuestUsergroupsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemAdminGuestUsergroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuestLang(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminHistory0(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminHistory1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminLoginTime(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "usr_name"
		if _, ok := i["usr-name"]; ok {
			tmp["usr_name"] = dataSourceFlattenSystemAdminLoginTimeUsrName(i["usr-name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "last_login"
		if _, ok := i["last-login"]; ok {
			tmp["last_login"] = dataSourceFlattenSystemAdminLoginTimeLastLogin(i["last-login"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "last_failed_login"
		if _, ok := i["last-failed-login"]; ok {
			tmp["last_failed_login"] = dataSourceFlattenSystemAdminLoginTimeLastFailedLogin(i["last-failed-login"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemAdminLoginTimeUsrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminLoginTimeLastLogin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminLoginTimeLastFailedLogin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiGlobalMenuFavorites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemAdminGuiGlobalMenuFavoritesId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemAdminGuiGlobalMenuFavoritesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiVdomMenuFavorites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemAdminGuiVdomMenuFavoritesId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemAdminGuiVdomMenuFavoritesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAdminGuiNewFeatureAcknowledge(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemAdminGuiNewFeatureAcknowledgeId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemAdminGuiNewFeatureAcknowledgeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemAdmin(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemAdminName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("wildcard", dataSourceFlattenSystemAdminWildcard(o["wildcard"], d, "wildcard")); err != nil {
		if !fortiAPIPatch(o["wildcard"]) {
			return fmt.Errorf("Error reading wildcard: %v", err)
		}
	}

	if err = d.Set("remote_auth", dataSourceFlattenSystemAdminRemoteAuth(o["remote-auth"], d, "remote_auth")); err != nil {
		if !fortiAPIPatch(o["remote-auth"]) {
			return fmt.Errorf("Error reading remote_auth: %v", err)
		}
	}

	if err = d.Set("remote_group", dataSourceFlattenSystemAdminRemoteGroup(o["remote-group"], d, "remote_group")); err != nil {
		if !fortiAPIPatch(o["remote-group"]) {
			return fmt.Errorf("Error reading remote_group: %v", err)
		}
	}

	if err = d.Set("peer_auth", dataSourceFlattenSystemAdminPeerAuth(o["peer-auth"], d, "peer_auth")); err != nil {
		if !fortiAPIPatch(o["peer-auth"]) {
			return fmt.Errorf("Error reading peer_auth: %v", err)
		}
	}

	if err = d.Set("peer_group", dataSourceFlattenSystemAdminPeerGroup(o["peer-group"], d, "peer_group")); err != nil {
		if !fortiAPIPatch(o["peer-group"]) {
			return fmt.Errorf("Error reading peer_group: %v", err)
		}
	}

	if err = d.Set("trusthost1", dataSourceFlattenSystemAdminTrusthost1(o["trusthost1"], d, "trusthost1")); err != nil {
		if !fortiAPIPatch(o["trusthost1"]) {
			return fmt.Errorf("Error reading trusthost1: %v", err)
		}
	}

	if err = d.Set("trusthost2", dataSourceFlattenSystemAdminTrusthost2(o["trusthost2"], d, "trusthost2")); err != nil {
		if !fortiAPIPatch(o["trusthost2"]) {
			return fmt.Errorf("Error reading trusthost2: %v", err)
		}
	}

	if err = d.Set("trusthost3", dataSourceFlattenSystemAdminTrusthost3(o["trusthost3"], d, "trusthost3")); err != nil {
		if !fortiAPIPatch(o["trusthost3"]) {
			return fmt.Errorf("Error reading trusthost3: %v", err)
		}
	}

	if err = d.Set("trusthost4", dataSourceFlattenSystemAdminTrusthost4(o["trusthost4"], d, "trusthost4")); err != nil {
		if !fortiAPIPatch(o["trusthost4"]) {
			return fmt.Errorf("Error reading trusthost4: %v", err)
		}
	}

	if err = d.Set("trusthost5", dataSourceFlattenSystemAdminTrusthost5(o["trusthost5"], d, "trusthost5")); err != nil {
		if !fortiAPIPatch(o["trusthost5"]) {
			return fmt.Errorf("Error reading trusthost5: %v", err)
		}
	}

	if err = d.Set("trusthost6", dataSourceFlattenSystemAdminTrusthost6(o["trusthost6"], d, "trusthost6")); err != nil {
		if !fortiAPIPatch(o["trusthost6"]) {
			return fmt.Errorf("Error reading trusthost6: %v", err)
		}
	}

	if err = d.Set("trusthost7", dataSourceFlattenSystemAdminTrusthost7(o["trusthost7"], d, "trusthost7")); err != nil {
		if !fortiAPIPatch(o["trusthost7"]) {
			return fmt.Errorf("Error reading trusthost7: %v", err)
		}
	}

	if err = d.Set("trusthost8", dataSourceFlattenSystemAdminTrusthost8(o["trusthost8"], d, "trusthost8")); err != nil {
		if !fortiAPIPatch(o["trusthost8"]) {
			return fmt.Errorf("Error reading trusthost8: %v", err)
		}
	}

	if err = d.Set("trusthost9", dataSourceFlattenSystemAdminTrusthost9(o["trusthost9"], d, "trusthost9")); err != nil {
		if !fortiAPIPatch(o["trusthost9"]) {
			return fmt.Errorf("Error reading trusthost9: %v", err)
		}
	}

	if err = d.Set("trusthost10", dataSourceFlattenSystemAdminTrusthost10(o["trusthost10"], d, "trusthost10")); err != nil {
		if !fortiAPIPatch(o["trusthost10"]) {
			return fmt.Errorf("Error reading trusthost10: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost1", dataSourceFlattenSystemAdminIp6Trusthost1(o["ip6-trusthost1"], d, "ip6_trusthost1")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost1"]) {
			return fmt.Errorf("Error reading ip6_trusthost1: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost2", dataSourceFlattenSystemAdminIp6Trusthost2(o["ip6-trusthost2"], d, "ip6_trusthost2")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost2"]) {
			return fmt.Errorf("Error reading ip6_trusthost2: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost3", dataSourceFlattenSystemAdminIp6Trusthost3(o["ip6-trusthost3"], d, "ip6_trusthost3")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost3"]) {
			return fmt.Errorf("Error reading ip6_trusthost3: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost4", dataSourceFlattenSystemAdminIp6Trusthost4(o["ip6-trusthost4"], d, "ip6_trusthost4")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost4"]) {
			return fmt.Errorf("Error reading ip6_trusthost4: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost5", dataSourceFlattenSystemAdminIp6Trusthost5(o["ip6-trusthost5"], d, "ip6_trusthost5")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost5"]) {
			return fmt.Errorf("Error reading ip6_trusthost5: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost6", dataSourceFlattenSystemAdminIp6Trusthost6(o["ip6-trusthost6"], d, "ip6_trusthost6")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost6"]) {
			return fmt.Errorf("Error reading ip6_trusthost6: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost7", dataSourceFlattenSystemAdminIp6Trusthost7(o["ip6-trusthost7"], d, "ip6_trusthost7")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost7"]) {
			return fmt.Errorf("Error reading ip6_trusthost7: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost8", dataSourceFlattenSystemAdminIp6Trusthost8(o["ip6-trusthost8"], d, "ip6_trusthost8")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost8"]) {
			return fmt.Errorf("Error reading ip6_trusthost8: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost9", dataSourceFlattenSystemAdminIp6Trusthost9(o["ip6-trusthost9"], d, "ip6_trusthost9")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost9"]) {
			return fmt.Errorf("Error reading ip6_trusthost9: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost10", dataSourceFlattenSystemAdminIp6Trusthost10(o["ip6-trusthost10"], d, "ip6_trusthost10")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost10"]) {
			return fmt.Errorf("Error reading ip6_trusthost10: %v", err)
		}
	}

	if err = d.Set("accprofile", dataSourceFlattenSystemAdminAccprofile(o["accprofile"], d, "accprofile")); err != nil {
		if !fortiAPIPatch(o["accprofile"]) {
			return fmt.Errorf("Error reading accprofile: %v", err)
		}
	}

	if err = d.Set("allow_remove_admin_session", dataSourceFlattenSystemAdminAllowRemoveAdminSession(o["allow-remove-admin-session"], d, "allow_remove_admin_session")); err != nil {
		if !fortiAPIPatch(o["allow-remove-admin-session"]) {
			return fmt.Errorf("Error reading allow_remove_admin_session: %v", err)
		}
	}

	if err = d.Set("comments", dataSourceFlattenSystemAdminComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("hidden", dataSourceFlattenSystemAdminHidden(o["hidden"], d, "hidden")); err != nil {
		if !fortiAPIPatch(o["hidden"]) {
			return fmt.Errorf("Error reading hidden: %v", err)
		}
	}

	if err = d.Set("vdom", dataSourceFlattenSystemAdminVdom(o["vdom"], d, "vdom")); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("ssh_certificate", dataSourceFlattenSystemAdminSshCertificate(o["ssh-certificate"], d, "ssh_certificate")); err != nil {
		if !fortiAPIPatch(o["ssh-certificate"]) {
			return fmt.Errorf("Error reading ssh_certificate: %v", err)
		}
	}

	if err = d.Set("schedule", dataSourceFlattenSystemAdminSchedule(o["schedule"], d, "schedule")); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("accprofile_override", dataSourceFlattenSystemAdminAccprofileOverride(o["accprofile-override"], d, "accprofile_override")); err != nil {
		if !fortiAPIPatch(o["accprofile-override"]) {
			return fmt.Errorf("Error reading accprofile_override: %v", err)
		}
	}

	if err = d.Set("radius_vdom_override", dataSourceFlattenSystemAdminRadiusVdomOverride(o["radius-vdom-override"], d, "radius_vdom_override")); err != nil {
		if !fortiAPIPatch(o["radius-vdom-override"]) {
			return fmt.Errorf("Error reading radius_vdom_override: %v", err)
		}
	}

	if err = d.Set("password_expire", dataSourceFlattenSystemAdminPasswordExpire(o["password-expire"], d, "password_expire")); err != nil {
		if !fortiAPIPatch(o["password-expire"]) {
			return fmt.Errorf("Error reading password_expire: %v", err)
		}
	}

	if err = d.Set("force_password_change", dataSourceFlattenSystemAdminForcePasswordChange(o["force-password-change"], d, "force_password_change")); err != nil {
		if !fortiAPIPatch(o["force-password-change"]) {
			return fmt.Errorf("Error reading force_password_change: %v", err)
		}
	}

	if err = d.Set("gui_dashboard", dataSourceFlattenSystemAdminGuiDashboard(o["gui-dashboard"], d, "gui_dashboard")); err != nil {
		if !fortiAPIPatch(o["gui-dashboard"]) {
			return fmt.Errorf("Error reading gui_dashboard: %v", err)
		}
	}

	if err = d.Set("two_factor", dataSourceFlattenSystemAdminTwoFactor(o["two-factor"], d, "two_factor")); err != nil {
		if !fortiAPIPatch(o["two-factor"]) {
			return fmt.Errorf("Error reading two_factor: %v", err)
		}
	}

	if err = d.Set("two_factor_authentication", dataSourceFlattenSystemAdminTwoFactorAuthentication(o["two-factor-authentication"], d, "two_factor_authentication")); err != nil {
		if !fortiAPIPatch(o["two-factor-authentication"]) {
			return fmt.Errorf("Error reading two_factor_authentication: %v", err)
		}
	}

	if err = d.Set("two_factor_notification", dataSourceFlattenSystemAdminTwoFactorNotification(o["two-factor-notification"], d, "two_factor_notification")); err != nil {
		if !fortiAPIPatch(o["two-factor-notification"]) {
			return fmt.Errorf("Error reading two_factor_notification: %v", err)
		}
	}

	if err = d.Set("fortitoken", dataSourceFlattenSystemAdminFortitoken(o["fortitoken"], d, "fortitoken")); err != nil {
		if !fortiAPIPatch(o["fortitoken"]) {
			return fmt.Errorf("Error reading fortitoken: %v", err)
		}
	}

	if err = d.Set("email_to", dataSourceFlattenSystemAdminEmailTo(o["email-to"], d, "email_to")); err != nil {
		if !fortiAPIPatch(o["email-to"]) {
			return fmt.Errorf("Error reading email_to: %v", err)
		}
	}

	if err = d.Set("sms_server", dataSourceFlattenSystemAdminSmsServer(o["sms-server"], d, "sms_server")); err != nil {
		if !fortiAPIPatch(o["sms-server"]) {
			return fmt.Errorf("Error reading sms_server: %v", err)
		}
	}

	if err = d.Set("sms_custom_server", dataSourceFlattenSystemAdminSmsCustomServer(o["sms-custom-server"], d, "sms_custom_server")); err != nil {
		if !fortiAPIPatch(o["sms-custom-server"]) {
			return fmt.Errorf("Error reading sms_custom_server: %v", err)
		}
	}

	if err = d.Set("sms_phone", dataSourceFlattenSystemAdminSmsPhone(o["sms-phone"], d, "sms_phone")); err != nil {
		if !fortiAPIPatch(o["sms-phone"]) {
			return fmt.Errorf("Error reading sms_phone: %v", err)
		}
	}

	if err = d.Set("guest_auth", dataSourceFlattenSystemAdminGuestAuth(o["guest-auth"], d, "guest_auth")); err != nil {
		if !fortiAPIPatch(o["guest-auth"]) {
			return fmt.Errorf("Error reading guest_auth: %v", err)
		}
	}

	if err = d.Set("guest_usergroups", dataSourceFlattenSystemAdminGuestUsergroups(o["guest-usergroups"], d, "guest_usergroups")); err != nil {
		if !fortiAPIPatch(o["guest-usergroups"]) {
			return fmt.Errorf("Error reading guest_usergroups: %v", err)
		}
	}

	if err = d.Set("guest_lang", dataSourceFlattenSystemAdminGuestLang(o["guest-lang"], d, "guest_lang")); err != nil {
		if !fortiAPIPatch(o["guest-lang"]) {
			return fmt.Errorf("Error reading guest_lang: %v", err)
		}
	}

	if err = d.Set("login_time", dataSourceFlattenSystemAdminLoginTime(o["login-time"], d, "login_time")); err != nil {
		if !fortiAPIPatch(o["login-time"]) {
			return fmt.Errorf("Error reading login_time: %v", err)
		}
	}

	if err = d.Set("gui_global_menu_favorites", dataSourceFlattenSystemAdminGuiGlobalMenuFavorites(o["gui-global-menu-favorites"], d, "gui_global_menu_favorites")); err != nil {
		if !fortiAPIPatch(o["gui-global-menu-favorites"]) {
			return fmt.Errorf("Error reading gui_global_menu_favorites: %v", err)
		}
	}

	if err = d.Set("gui_vdom_menu_favorites", dataSourceFlattenSystemAdminGuiVdomMenuFavorites(o["gui-vdom-menu-favorites"], d, "gui_vdom_menu_favorites")); err != nil {
		if !fortiAPIPatch(o["gui-vdom-menu-favorites"]) {
			return fmt.Errorf("Error reading gui_vdom_menu_favorites: %v", err)
		}
	}

	if err = d.Set("gui_new_feature_acknowledge", dataSourceFlattenSystemAdminGuiNewFeatureAcknowledge(o["gui-new-feature-acknowledge"], d, "gui_new_feature_acknowledge")); err != nil {
		if !fortiAPIPatch(o["gui-new-feature-acknowledge"]) {
			return fmt.Errorf("Error reading gui_new_feature_acknowledge: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemAdminFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
