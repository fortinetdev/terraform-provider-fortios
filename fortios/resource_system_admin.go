// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure admin users.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemAdmin() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAdminCreate,
		Read:   resourceSystemAdminRead,
		Update: resourceSystemAdminUpdate,
		Delete: resourceSystemAdminDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
				ForceNew:     true,
			},
			"wildcard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"peer_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peer_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"trusthost1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost8": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost9": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost10": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost8": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost9": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost10": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"accprofile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"allow_remove_admin_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"hidden": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"ssh_public_key1": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"ssh_public_key2": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"ssh_public_key3": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"ssh_certificate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"schedule": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"accprofile_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_vdom_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password_expire": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"force_password_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_dashboard": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"scope": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"layout_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"columns": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 20),
							Optional:     true,
							Computed:     true,
						},
						"widget": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"x_pos": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1000),
										Optional:     true,
										Computed:     true,
									},
									"y_pos": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1000),
										Optional:     true,
										Computed:     true,
									},
									"width": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 50),
										Optional:     true,
										Computed:     true,
									},
									"height": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 50),
										Optional:     true,
										Computed:     true,
									},
									"interface": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),
										Optional:     true,
										Computed:     true,
									},
									"region": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"industry": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"fabric_device": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
										Computed:     true,
									},
									"title": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),
										Optional:     true,
										Computed:     true,
									},
									"report_by": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"timeframe": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sort_by": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),
										Optional:     true,
										Computed:     true,
									},
									"visualization": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"filters": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"key": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 127),
													Optional:     true,
													Computed:     true,
												},
												"value": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 127),
													Optional:     true,
													Computed:     true,
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
				Optional: true,
				Computed: true,
			},
			"fortitoken": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 16),
				Optional:     true,
				Computed:     true,
			},
			"email_to": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"sms_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sms_custom_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"sms_phone": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"guest_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"guest_usergroups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"guest_lang": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"history0": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"history1": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"login_time": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"usr_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"last_login": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"last_failed_login": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"gui_global_menu_favorites": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"gui_vdom_menu_favorites": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"gui_new_feature_acknowledge": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemAdminCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAdmin(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdmin resource while getting object: %v", err)
	}

	o, err := c.CreateSystemAdmin(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemAdmin resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAdmin")
	}

	return resourceSystemAdminRead(d, m)
}

func resourceSystemAdminUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAdmin(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdmin resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAdmin(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdmin resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAdmin")
	}

	return resourceSystemAdminRead(d, m)
}

func resourceSystemAdminDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemAdmin(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAdminRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemAdmin(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdmin resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAdmin(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdmin resource from API: %v", err)
	}
	return nil
}

func flattenSystemAdminName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminWildcard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminRemoteAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminRemoteGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminPeerAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminPeerGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminTrusthost1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminTrusthost2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminTrusthost3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminTrusthost4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminTrusthost5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminTrusthost6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminTrusthost7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminTrusthost8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminTrusthost9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminTrusthost10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemAdminIp6Trusthost1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminAccprofile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminAllowRemoveAdminSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminHidden(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminVdom(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemAdminVdomName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemAdminVdomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSshPublicKey1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSshPublicKey2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSshPublicKey3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSshCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminAccprofileOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminRadiusVdomOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminPasswordExpire(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminForcePasswordChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboard(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemAdminGuiDashboardId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenSystemAdminGuiDashboardName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scope"
		if _, ok := i["scope"]; ok {
			tmp["scope"] = flattenSystemAdminGuiDashboardScope(i["scope"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "layout_type"
		if _, ok := i["layout-type"]; ok {
			tmp["layout_type"] = flattenSystemAdminGuiDashboardLayoutType(i["layout-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "columns"
		if _, ok := i["columns"]; ok {
			tmp["columns"] = flattenSystemAdminGuiDashboardColumns(i["columns"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "widget"
		if _, ok := i["widget"]; ok {
			tmp["widget"] = flattenSystemAdminGuiDashboardWidget(i["widget"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemAdminGuiDashboardId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardScope(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardLayoutType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardColumns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardWidget(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemAdminGuiDashboardWidgetId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = flattenSystemAdminGuiDashboardWidgetType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "x_pos"
		if _, ok := i["x-pos"]; ok {
			tmp["x_pos"] = flattenSystemAdminGuiDashboardWidgetXPos(i["x-pos"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "y_pos"
		if _, ok := i["y-pos"]; ok {
			tmp["y_pos"] = flattenSystemAdminGuiDashboardWidgetYPos(i["y-pos"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := i["width"]; ok {
			tmp["width"] = flattenSystemAdminGuiDashboardWidgetWidth(i["width"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := i["height"]; ok {
			tmp["height"] = flattenSystemAdminGuiDashboardWidgetHeight(i["height"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = flattenSystemAdminGuiDashboardWidgetInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "region"
		if _, ok := i["region"]; ok {
			tmp["region"] = flattenSystemAdminGuiDashboardWidgetRegion(i["region"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "industry"
		if _, ok := i["industry"]; ok {
			tmp["industry"] = flattenSystemAdminGuiDashboardWidgetIndustry(i["industry"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_device"
		if _, ok := i["fabric-device"]; ok {
			tmp["fabric_device"] = flattenSystemAdminGuiDashboardWidgetFabricDevice(i["fabric-device"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "title"
		if _, ok := i["title"]; ok {
			tmp["title"] = flattenSystemAdminGuiDashboardWidgetTitle(i["title"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "report_by"
		if _, ok := i["report-by"]; ok {
			tmp["report_by"] = flattenSystemAdminGuiDashboardWidgetReportBy(i["report-by"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "timeframe"
		if _, ok := i["timeframe"]; ok {
			tmp["timeframe"] = flattenSystemAdminGuiDashboardWidgetTimeframe(i["timeframe"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sort_by"
		if _, ok := i["sort-by"]; ok {
			tmp["sort_by"] = flattenSystemAdminGuiDashboardWidgetSortBy(i["sort-by"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "visualization"
		if _, ok := i["visualization"]; ok {
			tmp["visualization"] = flattenSystemAdminGuiDashboardWidgetVisualization(i["visualization"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filters"
		if _, ok := i["filters"]; ok {
			tmp["filters"] = flattenSystemAdminGuiDashboardWidgetFilters(i["filters"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemAdminGuiDashboardWidgetId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardWidgetType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardWidgetXPos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardWidgetYPos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardWidgetWidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardWidgetHeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardWidgetInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardWidgetRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardWidgetIndustry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardWidgetFabricDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardWidgetTitle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardWidgetReportBy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardWidgetTimeframe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardWidgetSortBy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardWidgetVisualization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardWidgetFilters(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemAdminGuiDashboardWidgetFiltersId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := i["key"]; ok {
			tmp["key"] = flattenSystemAdminGuiDashboardWidgetFiltersKey(i["key"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			tmp["value"] = flattenSystemAdminGuiDashboardWidgetFiltersValue(i["value"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemAdminGuiDashboardWidgetFiltersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardWidgetFiltersKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDashboardWidgetFiltersValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminTwoFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminFortitoken(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminEmailTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSmsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSmsCustomServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSmsPhone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuestAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuestUsergroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemAdminGuestUsergroupsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemAdminGuestUsergroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuestLang(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminHistory0(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminHistory1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLoginTime(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["usr_name"] = flattenSystemAdminLoginTimeUsrName(i["usr-name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "last_login"
		if _, ok := i["last-login"]; ok {
			tmp["last_login"] = flattenSystemAdminLoginTimeLastLogin(i["last-login"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "last_failed_login"
		if _, ok := i["last-failed-login"]; ok {
			tmp["last_failed_login"] = flattenSystemAdminLoginTimeLastFailedLogin(i["last-failed-login"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemAdminLoginTimeUsrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLoginTimeLastLogin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminLoginTimeLastFailedLogin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiGlobalMenuFavorites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemAdminGuiGlobalMenuFavoritesId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemAdminGuiGlobalMenuFavoritesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiVdomMenuFavorites(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemAdminGuiVdomMenuFavoritesId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemAdminGuiVdomMenuFavoritesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiNewFeatureAcknowledge(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemAdminGuiNewFeatureAcknowledgeId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemAdminGuiNewFeatureAcknowledgeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAdmin(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemAdminName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("wildcard", flattenSystemAdminWildcard(o["wildcard"], d, "wildcard")); err != nil {
		if !fortiAPIPatch(o["wildcard"]) {
			return fmt.Errorf("Error reading wildcard: %v", err)
		}
	}

	if err = d.Set("remote_auth", flattenSystemAdminRemoteAuth(o["remote-auth"], d, "remote_auth")); err != nil {
		if !fortiAPIPatch(o["remote-auth"]) {
			return fmt.Errorf("Error reading remote_auth: %v", err)
		}
	}

	if err = d.Set("remote_group", flattenSystemAdminRemoteGroup(o["remote-group"], d, "remote_group")); err != nil {
		if !fortiAPIPatch(o["remote-group"]) {
			return fmt.Errorf("Error reading remote_group: %v", err)
		}
	}

	if err = d.Set("peer_auth", flattenSystemAdminPeerAuth(o["peer-auth"], d, "peer_auth")); err != nil {
		if !fortiAPIPatch(o["peer-auth"]) {
			return fmt.Errorf("Error reading peer_auth: %v", err)
		}
	}

	if err = d.Set("peer_group", flattenSystemAdminPeerGroup(o["peer-group"], d, "peer_group")); err != nil {
		if !fortiAPIPatch(o["peer-group"]) {
			return fmt.Errorf("Error reading peer_group: %v", err)
		}
	}

	if err = d.Set("trusthost1", flattenSystemAdminTrusthost1(o["trusthost1"], d, "trusthost1")); err != nil {
		if !fortiAPIPatch(o["trusthost1"]) {
			return fmt.Errorf("Error reading trusthost1: %v", err)
		}
	}

	if err = d.Set("trusthost2", flattenSystemAdminTrusthost2(o["trusthost2"], d, "trusthost2")); err != nil {
		if !fortiAPIPatch(o["trusthost2"]) {
			return fmt.Errorf("Error reading trusthost2: %v", err)
		}
	}

	if err = d.Set("trusthost3", flattenSystemAdminTrusthost3(o["trusthost3"], d, "trusthost3")); err != nil {
		if !fortiAPIPatch(o["trusthost3"]) {
			return fmt.Errorf("Error reading trusthost3: %v", err)
		}
	}

	if err = d.Set("trusthost4", flattenSystemAdminTrusthost4(o["trusthost4"], d, "trusthost4")); err != nil {
		if !fortiAPIPatch(o["trusthost4"]) {
			return fmt.Errorf("Error reading trusthost4: %v", err)
		}
	}

	if err = d.Set("trusthost5", flattenSystemAdminTrusthost5(o["trusthost5"], d, "trusthost5")); err != nil {
		if !fortiAPIPatch(o["trusthost5"]) {
			return fmt.Errorf("Error reading trusthost5: %v", err)
		}
	}

	if err = d.Set("trusthost6", flattenSystemAdminTrusthost6(o["trusthost6"], d, "trusthost6")); err != nil {
		if !fortiAPIPatch(o["trusthost6"]) {
			return fmt.Errorf("Error reading trusthost6: %v", err)
		}
	}

	if err = d.Set("trusthost7", flattenSystemAdminTrusthost7(o["trusthost7"], d, "trusthost7")); err != nil {
		if !fortiAPIPatch(o["trusthost7"]) {
			return fmt.Errorf("Error reading trusthost7: %v", err)
		}
	}

	if err = d.Set("trusthost8", flattenSystemAdminTrusthost8(o["trusthost8"], d, "trusthost8")); err != nil {
		if !fortiAPIPatch(o["trusthost8"]) {
			return fmt.Errorf("Error reading trusthost8: %v", err)
		}
	}

	if err = d.Set("trusthost9", flattenSystemAdminTrusthost9(o["trusthost9"], d, "trusthost9")); err != nil {
		if !fortiAPIPatch(o["trusthost9"]) {
			return fmt.Errorf("Error reading trusthost9: %v", err)
		}
	}

	if err = d.Set("trusthost10", flattenSystemAdminTrusthost10(o["trusthost10"], d, "trusthost10")); err != nil {
		if !fortiAPIPatch(o["trusthost10"]) {
			return fmt.Errorf("Error reading trusthost10: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost1", flattenSystemAdminIp6Trusthost1(o["ip6-trusthost1"], d, "ip6_trusthost1")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost1"]) {
			return fmt.Errorf("Error reading ip6_trusthost1: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost2", flattenSystemAdminIp6Trusthost2(o["ip6-trusthost2"], d, "ip6_trusthost2")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost2"]) {
			return fmt.Errorf("Error reading ip6_trusthost2: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost3", flattenSystemAdminIp6Trusthost3(o["ip6-trusthost3"], d, "ip6_trusthost3")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost3"]) {
			return fmt.Errorf("Error reading ip6_trusthost3: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost4", flattenSystemAdminIp6Trusthost4(o["ip6-trusthost4"], d, "ip6_trusthost4")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost4"]) {
			return fmt.Errorf("Error reading ip6_trusthost4: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost5", flattenSystemAdminIp6Trusthost5(o["ip6-trusthost5"], d, "ip6_trusthost5")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost5"]) {
			return fmt.Errorf("Error reading ip6_trusthost5: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost6", flattenSystemAdminIp6Trusthost6(o["ip6-trusthost6"], d, "ip6_trusthost6")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost6"]) {
			return fmt.Errorf("Error reading ip6_trusthost6: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost7", flattenSystemAdminIp6Trusthost7(o["ip6-trusthost7"], d, "ip6_trusthost7")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost7"]) {
			return fmt.Errorf("Error reading ip6_trusthost7: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost8", flattenSystemAdminIp6Trusthost8(o["ip6-trusthost8"], d, "ip6_trusthost8")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost8"]) {
			return fmt.Errorf("Error reading ip6_trusthost8: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost9", flattenSystemAdminIp6Trusthost9(o["ip6-trusthost9"], d, "ip6_trusthost9")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost9"]) {
			return fmt.Errorf("Error reading ip6_trusthost9: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost10", flattenSystemAdminIp6Trusthost10(o["ip6-trusthost10"], d, "ip6_trusthost10")); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost10"]) {
			return fmt.Errorf("Error reading ip6_trusthost10: %v", err)
		}
	}

	if err = d.Set("accprofile", flattenSystemAdminAccprofile(o["accprofile"], d, "accprofile")); err != nil {
		if !fortiAPIPatch(o["accprofile"]) {
			return fmt.Errorf("Error reading accprofile: %v", err)
		}
	}

	if err = d.Set("allow_remove_admin_session", flattenSystemAdminAllowRemoveAdminSession(o["allow-remove-admin-session"], d, "allow_remove_admin_session")); err != nil {
		if !fortiAPIPatch(o["allow-remove-admin-session"]) {
			return fmt.Errorf("Error reading allow_remove_admin_session: %v", err)
		}
	}

	if err = d.Set("comments", flattenSystemAdminComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("hidden", flattenSystemAdminHidden(o["hidden"], d, "hidden")); err != nil {
		if !fortiAPIPatch(o["hidden"]) {
			return fmt.Errorf("Error reading hidden: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vdom", flattenSystemAdminVdom(o["vdom"], d, "vdom")); err != nil {
			if !fortiAPIPatch(o["vdom"]) {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vdom"); ok {
			if err = d.Set("vdom", flattenSystemAdminVdom(o["vdom"], d, "vdom")); err != nil {
				if !fortiAPIPatch(o["vdom"]) {
					return fmt.Errorf("Error reading vdom: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssh_certificate", flattenSystemAdminSshCertificate(o["ssh-certificate"], d, "ssh_certificate")); err != nil {
		if !fortiAPIPatch(o["ssh-certificate"]) {
			return fmt.Errorf("Error reading ssh_certificate: %v", err)
		}
	}

	if err = d.Set("schedule", flattenSystemAdminSchedule(o["schedule"], d, "schedule")); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("accprofile_override", flattenSystemAdminAccprofileOverride(o["accprofile-override"], d, "accprofile_override")); err != nil {
		if !fortiAPIPatch(o["accprofile-override"]) {
			return fmt.Errorf("Error reading accprofile_override: %v", err)
		}
	}

	if err = d.Set("radius_vdom_override", flattenSystemAdminRadiusVdomOverride(o["radius-vdom-override"], d, "radius_vdom_override")); err != nil {
		if !fortiAPIPatch(o["radius-vdom-override"]) {
			return fmt.Errorf("Error reading radius_vdom_override: %v", err)
		}
	}

	if err = d.Set("password_expire", flattenSystemAdminPasswordExpire(o["password-expire"], d, "password_expire")); err != nil {
		if !fortiAPIPatch(o["password-expire"]) {
			return fmt.Errorf("Error reading password_expire: %v", err)
		}
	}

	if err = d.Set("force_password_change", flattenSystemAdminForcePasswordChange(o["force-password-change"], d, "force_password_change")); err != nil {
		if !fortiAPIPatch(o["force-password-change"]) {
			return fmt.Errorf("Error reading force_password_change: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("gui_dashboard", flattenSystemAdminGuiDashboard(o["gui-dashboard"], d, "gui_dashboard")); err != nil {
			if !fortiAPIPatch(o["gui-dashboard"]) {
				return fmt.Errorf("Error reading gui_dashboard: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("gui_dashboard"); ok {
			if err = d.Set("gui_dashboard", flattenSystemAdminGuiDashboard(o["gui-dashboard"], d, "gui_dashboard")); err != nil {
				if !fortiAPIPatch(o["gui-dashboard"]) {
					return fmt.Errorf("Error reading gui_dashboard: %v", err)
				}
			}
		}
	}

	if err = d.Set("two_factor", flattenSystemAdminTwoFactor(o["two-factor"], d, "two_factor")); err != nil {
		if !fortiAPIPatch(o["two-factor"]) {
			return fmt.Errorf("Error reading two_factor: %v", err)
		}
	}

	if err = d.Set("fortitoken", flattenSystemAdminFortitoken(o["fortitoken"], d, "fortitoken")); err != nil {
		if !fortiAPIPatch(o["fortitoken"]) {
			return fmt.Errorf("Error reading fortitoken: %v", err)
		}
	}

	if err = d.Set("email_to", flattenSystemAdminEmailTo(o["email-to"], d, "email_to")); err != nil {
		if !fortiAPIPatch(o["email-to"]) {
			return fmt.Errorf("Error reading email_to: %v", err)
		}
	}

	if err = d.Set("sms_server", flattenSystemAdminSmsServer(o["sms-server"], d, "sms_server")); err != nil {
		if !fortiAPIPatch(o["sms-server"]) {
			return fmt.Errorf("Error reading sms_server: %v", err)
		}
	}

	if err = d.Set("sms_custom_server", flattenSystemAdminSmsCustomServer(o["sms-custom-server"], d, "sms_custom_server")); err != nil {
		if !fortiAPIPatch(o["sms-custom-server"]) {
			return fmt.Errorf("Error reading sms_custom_server: %v", err)
		}
	}

	if err = d.Set("sms_phone", flattenSystemAdminSmsPhone(o["sms-phone"], d, "sms_phone")); err != nil {
		if !fortiAPIPatch(o["sms-phone"]) {
			return fmt.Errorf("Error reading sms_phone: %v", err)
		}
	}

	if err = d.Set("guest_auth", flattenSystemAdminGuestAuth(o["guest-auth"], d, "guest_auth")); err != nil {
		if !fortiAPIPatch(o["guest-auth"]) {
			return fmt.Errorf("Error reading guest_auth: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("guest_usergroups", flattenSystemAdminGuestUsergroups(o["guest-usergroups"], d, "guest_usergroups")); err != nil {
			if !fortiAPIPatch(o["guest-usergroups"]) {
				return fmt.Errorf("Error reading guest_usergroups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("guest_usergroups"); ok {
			if err = d.Set("guest_usergroups", flattenSystemAdminGuestUsergroups(o["guest-usergroups"], d, "guest_usergroups")); err != nil {
				if !fortiAPIPatch(o["guest-usergroups"]) {
					return fmt.Errorf("Error reading guest_usergroups: %v", err)
				}
			}
		}
	}

	if err = d.Set("guest_lang", flattenSystemAdminGuestLang(o["guest-lang"], d, "guest_lang")); err != nil {
		if !fortiAPIPatch(o["guest-lang"]) {
			return fmt.Errorf("Error reading guest_lang: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("login_time", flattenSystemAdminLoginTime(o["login-time"], d, "login_time")); err != nil {
			if !fortiAPIPatch(o["login-time"]) {
				return fmt.Errorf("Error reading login_time: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("login_time"); ok {
			if err = d.Set("login_time", flattenSystemAdminLoginTime(o["login-time"], d, "login_time")); err != nil {
				if !fortiAPIPatch(o["login-time"]) {
					return fmt.Errorf("Error reading login_time: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("gui_global_menu_favorites", flattenSystemAdminGuiGlobalMenuFavorites(o["gui-global-menu-favorites"], d, "gui_global_menu_favorites")); err != nil {
			if !fortiAPIPatch(o["gui-global-menu-favorites"]) {
				return fmt.Errorf("Error reading gui_global_menu_favorites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("gui_global_menu_favorites"); ok {
			if err = d.Set("gui_global_menu_favorites", flattenSystemAdminGuiGlobalMenuFavorites(o["gui-global-menu-favorites"], d, "gui_global_menu_favorites")); err != nil {
				if !fortiAPIPatch(o["gui-global-menu-favorites"]) {
					return fmt.Errorf("Error reading gui_global_menu_favorites: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("gui_vdom_menu_favorites", flattenSystemAdminGuiVdomMenuFavorites(o["gui-vdom-menu-favorites"], d, "gui_vdom_menu_favorites")); err != nil {
			if !fortiAPIPatch(o["gui-vdom-menu-favorites"]) {
				return fmt.Errorf("Error reading gui_vdom_menu_favorites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("gui_vdom_menu_favorites"); ok {
			if err = d.Set("gui_vdom_menu_favorites", flattenSystemAdminGuiVdomMenuFavorites(o["gui-vdom-menu-favorites"], d, "gui_vdom_menu_favorites")); err != nil {
				if !fortiAPIPatch(o["gui-vdom-menu-favorites"]) {
					return fmt.Errorf("Error reading gui_vdom_menu_favorites: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("gui_new_feature_acknowledge", flattenSystemAdminGuiNewFeatureAcknowledge(o["gui-new-feature-acknowledge"], d, "gui_new_feature_acknowledge")); err != nil {
			if !fortiAPIPatch(o["gui-new-feature-acknowledge"]) {
				return fmt.Errorf("Error reading gui_new_feature_acknowledge: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("gui_new_feature_acknowledge"); ok {
			if err = d.Set("gui_new_feature_acknowledge", flattenSystemAdminGuiNewFeatureAcknowledge(o["gui-new-feature-acknowledge"], d, "gui_new_feature_acknowledge")); err != nil {
				if !fortiAPIPatch(o["gui-new-feature-acknowledge"]) {
					return fmt.Errorf("Error reading gui_new_feature_acknowledge: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemAdminFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAdminName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminWildcard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminRemoteAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminRemoteGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminPeerAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminPeerGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost9(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost10(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost9(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost10(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminAccprofile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminAllowRemoveAdminSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminHidden(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemAdminVdomName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAdminVdomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSshPublicKey1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSshPublicKey2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSshPublicKey3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSshCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminAccprofileOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminRadiusVdomOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminPasswordExpire(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminForcePasswordChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemAdminGuiDashboardId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSystemAdminGuiDashboardName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scope"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["scope"], _ = expandSystemAdminGuiDashboardScope(d, i["scope"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "layout_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["layout-type"], _ = expandSystemAdminGuiDashboardLayoutType(d, i["layout_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "columns"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["columns"], _ = expandSystemAdminGuiDashboardColumns(d, i["columns"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "widget"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["widget"], _ = expandSystemAdminGuiDashboardWidget(d, i["widget"], pre_append)
		} else {
			tmp["widget"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAdminGuiDashboardId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardLayoutType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardColumns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardWidget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemAdminGuiDashboardWidgetId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandSystemAdminGuiDashboardWidgetType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "x_pos"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["x-pos"], _ = expandSystemAdminGuiDashboardWidgetXPos(d, i["x_pos"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "y_pos"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["y-pos"], _ = expandSystemAdminGuiDashboardWidgetYPos(d, i["y_pos"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["width"], _ = expandSystemAdminGuiDashboardWidgetWidth(d, i["width"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["height"], _ = expandSystemAdminGuiDashboardWidgetHeight(d, i["height"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandSystemAdminGuiDashboardWidgetInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "region"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["region"], _ = expandSystemAdminGuiDashboardWidgetRegion(d, i["region"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "industry"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["industry"], _ = expandSystemAdminGuiDashboardWidgetIndustry(d, i["industry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fabric_device"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fabric-device"], _ = expandSystemAdminGuiDashboardWidgetFabricDevice(d, i["fabric_device"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "title"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["title"], _ = expandSystemAdminGuiDashboardWidgetTitle(d, i["title"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "report_by"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["report-by"], _ = expandSystemAdminGuiDashboardWidgetReportBy(d, i["report_by"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "timeframe"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["timeframe"], _ = expandSystemAdminGuiDashboardWidgetTimeframe(d, i["timeframe"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sort_by"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sort-by"], _ = expandSystemAdminGuiDashboardWidgetSortBy(d, i["sort_by"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "visualization"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["visualization"], _ = expandSystemAdminGuiDashboardWidgetVisualization(d, i["visualization"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filters"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filters"], _ = expandSystemAdminGuiDashboardWidgetFilters(d, i["filters"], pre_append)
		} else {
			tmp["filters"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAdminGuiDashboardWidgetId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardWidgetType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardWidgetXPos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardWidgetYPos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardWidgetWidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardWidgetHeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardWidgetInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardWidgetRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardWidgetIndustry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardWidgetFabricDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardWidgetTitle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardWidgetReportBy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardWidgetTimeframe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardWidgetSortBy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardWidgetVisualization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardWidgetFilters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemAdminGuiDashboardWidgetFiltersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["key"], _ = expandSystemAdminGuiDashboardWidgetFiltersKey(d, i["key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandSystemAdminGuiDashboardWidgetFiltersValue(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAdminGuiDashboardWidgetFiltersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardWidgetFiltersKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDashboardWidgetFiltersValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTwoFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminFortitoken(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminEmailTo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSmsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSmsCustomServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSmsPhone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuestAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuestUsergroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemAdminGuestUsergroupsName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAdminGuestUsergroupsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuestLang(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminHistory0(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminHistory1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLoginTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "usr_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["usr-name"], _ = expandSystemAdminLoginTimeUsrName(d, i["usr_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "last_login"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["last-login"], _ = expandSystemAdminLoginTimeLastLogin(d, i["last_login"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "last_failed_login"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["last-failed-login"], _ = expandSystemAdminLoginTimeLastFailedLogin(d, i["last_failed_login"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAdminLoginTimeUsrName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLoginTimeLastLogin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminLoginTimeLastFailedLogin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiGlobalMenuFavorites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemAdminGuiGlobalMenuFavoritesId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAdminGuiGlobalMenuFavoritesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiVdomMenuFavorites(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemAdminGuiVdomMenuFavoritesId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAdminGuiVdomMenuFavoritesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiNewFeatureAcknowledge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemAdminGuiNewFeatureAcknowledgeId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAdminGuiNewFeatureAcknowledgeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAdmin(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemAdminName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("wildcard"); ok {
		t, err := expandSystemAdminWildcard(d, v, "wildcard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard"] = t
		}
	}

	if v, ok := d.GetOk("remote_auth"); ok {
		t, err := expandSystemAdminRemoteAuth(d, v, "remote_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-auth"] = t
		}
	}

	if v, ok := d.GetOk("remote_group"); ok {
		t, err := expandSystemAdminRemoteGroup(d, v, "remote_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-group"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandSystemAdminPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("peer_auth"); ok {
		t, err := expandSystemAdminPeerAuth(d, v, "peer_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-auth"] = t
		}
	}

	if v, ok := d.GetOk("peer_group"); ok {
		t, err := expandSystemAdminPeerGroup(d, v, "peer_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-group"] = t
		}
	}

	if v, ok := d.GetOk("trusthost1"); ok {
		t, err := expandSystemAdminTrusthost1(d, v, "trusthost1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost1"] = t
		}
	}

	if v, ok := d.GetOk("trusthost2"); ok {
		t, err := expandSystemAdminTrusthost2(d, v, "trusthost2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost2"] = t
		}
	}

	if v, ok := d.GetOk("trusthost3"); ok {
		t, err := expandSystemAdminTrusthost3(d, v, "trusthost3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost3"] = t
		}
	}

	if v, ok := d.GetOk("trusthost4"); ok {
		t, err := expandSystemAdminTrusthost4(d, v, "trusthost4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost4"] = t
		}
	}

	if v, ok := d.GetOk("trusthost5"); ok {
		t, err := expandSystemAdminTrusthost5(d, v, "trusthost5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost5"] = t
		}
	}

	if v, ok := d.GetOk("trusthost6"); ok {
		t, err := expandSystemAdminTrusthost6(d, v, "trusthost6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost6"] = t
		}
	}

	if v, ok := d.GetOk("trusthost7"); ok {
		t, err := expandSystemAdminTrusthost7(d, v, "trusthost7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost7"] = t
		}
	}

	if v, ok := d.GetOk("trusthost8"); ok {
		t, err := expandSystemAdminTrusthost8(d, v, "trusthost8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost8"] = t
		}
	}

	if v, ok := d.GetOk("trusthost9"); ok {
		t, err := expandSystemAdminTrusthost9(d, v, "trusthost9")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost9"] = t
		}
	}

	if v, ok := d.GetOk("trusthost10"); ok {
		t, err := expandSystemAdminTrusthost10(d, v, "trusthost10")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost10"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost1"); ok {
		t, err := expandSystemAdminIp6Trusthost1(d, v, "ip6_trusthost1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost1"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost2"); ok {
		t, err := expandSystemAdminIp6Trusthost2(d, v, "ip6_trusthost2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost2"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost3"); ok {
		t, err := expandSystemAdminIp6Trusthost3(d, v, "ip6_trusthost3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost3"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost4"); ok {
		t, err := expandSystemAdminIp6Trusthost4(d, v, "ip6_trusthost4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost4"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost5"); ok {
		t, err := expandSystemAdminIp6Trusthost5(d, v, "ip6_trusthost5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost5"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost6"); ok {
		t, err := expandSystemAdminIp6Trusthost6(d, v, "ip6_trusthost6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost6"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost7"); ok {
		t, err := expandSystemAdminIp6Trusthost7(d, v, "ip6_trusthost7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost7"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost8"); ok {
		t, err := expandSystemAdminIp6Trusthost8(d, v, "ip6_trusthost8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost8"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost9"); ok {
		t, err := expandSystemAdminIp6Trusthost9(d, v, "ip6_trusthost9")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost9"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost10"); ok {
		t, err := expandSystemAdminIp6Trusthost10(d, v, "ip6_trusthost10")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost10"] = t
		}
	}

	if v, ok := d.GetOk("accprofile"); ok {
		t, err := expandSystemAdminAccprofile(d, v, "accprofile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accprofile"] = t
		}
	}

	if v, ok := d.GetOk("allow_remove_admin_session"); ok {
		t, err := expandSystemAdminAllowRemoveAdminSession(d, v, "allow_remove_admin_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-remove-admin-session"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandSystemAdminComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("hidden"); ok {
		t, err := expandSystemAdminHidden(d, v, "hidden")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hidden"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		t, err := expandSystemAdminVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("ssh_public_key1"); ok {
		t, err := expandSystemAdminSshPublicKey1(d, v, "ssh_public_key1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-public-key1"] = t
		}
	}

	if v, ok := d.GetOk("ssh_public_key2"); ok {
		t, err := expandSystemAdminSshPublicKey2(d, v, "ssh_public_key2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-public-key2"] = t
		}
	}

	if v, ok := d.GetOk("ssh_public_key3"); ok {
		t, err := expandSystemAdminSshPublicKey3(d, v, "ssh_public_key3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-public-key3"] = t
		}
	}

	if v, ok := d.GetOk("ssh_certificate"); ok {
		t, err := expandSystemAdminSshCertificate(d, v, "ssh_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-certificate"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {
		t, err := expandSystemAdminSchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("accprofile_override"); ok {
		t, err := expandSystemAdminAccprofileOverride(d, v, "accprofile_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accprofile-override"] = t
		}
	}

	if v, ok := d.GetOk("radius_vdom_override"); ok {
		t, err := expandSystemAdminRadiusVdomOverride(d, v, "radius_vdom_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-vdom-override"] = t
		}
	}

	if v, ok := d.GetOk("password_expire"); ok {
		t, err := expandSystemAdminPasswordExpire(d, v, "password_expire")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-expire"] = t
		}
	}

	if v, ok := d.GetOk("force_password_change"); ok {
		t, err := expandSystemAdminForcePasswordChange(d, v, "force_password_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["force-password-change"] = t
		}
	}

	if v, ok := d.GetOk("gui_dashboard"); ok {
		t, err := expandSystemAdminGuiDashboard(d, v, "gui_dashboard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-dashboard"] = t
		}
	}

	if v, ok := d.GetOk("two_factor"); ok {
		t, err := expandSystemAdminTwoFactor(d, v, "two_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor"] = t
		}
	}

	if v, ok := d.GetOk("fortitoken"); ok {
		t, err := expandSystemAdminFortitoken(d, v, "fortitoken")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortitoken"] = t
		}
	}

	if v, ok := d.GetOk("email_to"); ok {
		t, err := expandSystemAdminEmailTo(d, v, "email_to")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-to"] = t
		}
	}

	if v, ok := d.GetOk("sms_server"); ok {
		t, err := expandSystemAdminSmsServer(d, v, "sms_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-server"] = t
		}
	}

	if v, ok := d.GetOk("sms_custom_server"); ok {
		t, err := expandSystemAdminSmsCustomServer(d, v, "sms_custom_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-custom-server"] = t
		}
	}

	if v, ok := d.GetOk("sms_phone"); ok {
		t, err := expandSystemAdminSmsPhone(d, v, "sms_phone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-phone"] = t
		}
	}

	if v, ok := d.GetOk("guest_auth"); ok {
		t, err := expandSystemAdminGuestAuth(d, v, "guest_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest-auth"] = t
		}
	}

	if v, ok := d.GetOk("guest_usergroups"); ok {
		t, err := expandSystemAdminGuestUsergroups(d, v, "guest_usergroups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest-usergroups"] = t
		}
	}

	if v, ok := d.GetOk("guest_lang"); ok {
		t, err := expandSystemAdminGuestLang(d, v, "guest_lang")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest-lang"] = t
		}
	}

	if v, ok := d.GetOk("history0"); ok {
		t, err := expandSystemAdminHistory0(d, v, "history0")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history0"] = t
		}
	}

	if v, ok := d.GetOk("history1"); ok {
		t, err := expandSystemAdminHistory1(d, v, "history1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history1"] = t
		}
	}

	if v, ok := d.GetOk("login_time"); ok {
		t, err := expandSystemAdminLoginTime(d, v, "login_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-time"] = t
		}
	}

	if v, ok := d.GetOk("gui_global_menu_favorites"); ok {
		t, err := expandSystemAdminGuiGlobalMenuFavorites(d, v, "gui_global_menu_favorites")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-global-menu-favorites"] = t
		}
	}

	if v, ok := d.GetOk("gui_vdom_menu_favorites"); ok {
		t, err := expandSystemAdminGuiVdomMenuFavorites(d, v, "gui_vdom_menu_favorites")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-vdom-menu-favorites"] = t
		}
	}

	if v, ok := d.GetOk("gui_new_feature_acknowledge"); ok {
		t, err := expandSystemAdminGuiNewFeatureAcknowledge(d, v, "gui_new_feature_acknowledge")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-new-feature-acknowledge"] = t
		}
	}

	return &obj, nil
}
