// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Web filter profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterProfileCreate,
		Read:   resourceWebfilterProfileRead,
		Update: resourceWebfilterProfileUpdate,
		Delete: resourceWebfilterProfileDelete,

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
				Required:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"feature_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"inspection_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"file_filter": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"scan_archive_contents": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"entries": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"filter": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
										Computed:     true,
									},
									"comment": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
									},
									"protocol": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"direction": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"password_protected": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"file_type": &schema.Schema{
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 39),
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
			"https_replacemsg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_flow_log_encoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ovrd_perm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"post_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ovrd_cookie": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ovrd_scope": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"profile_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ovrd_dur_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ovrd_dur": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"profile_attribute": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ovrd_user_group": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"profile": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
					},
				},
			},
			"web": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bword_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"bword_table": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"urlfilter_table": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"content_header_list": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"allowlist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"blacklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"whitelist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"safe_search": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"youtube_restrict": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vimeo_restrict": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"log_search": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keyword_match": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pattern": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
					},
				},
			},
			"youtube_channel_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"youtube_channel_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"channel_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"comment": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
					},
				},
			},
			"ftgd_wf": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"options": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"exempt_quota": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ovrd": &schema.Schema{
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
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),
										Optional:     true,
										Computed:     true,
									},
									"category": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),
										Optional:     true,
										Computed:     true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"warn_duration": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"auth_usr_grp": &schema.Schema{
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 79),
													Optional:     true,
													Computed:     true,
												},
											},
										},
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"override_replacemsg": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 28),
										Optional:     true,
										Computed:     true,
									},
									"warning_prompt": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"warning_duration_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"quota": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"category": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"unit": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"duration": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"override_replacemsg": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 28),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"max_quota_timeout": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 86400),
							Optional:     true,
							Computed:     true,
						},
						"rate_image_urls": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate_javascript_urls": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate_css_urls": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate_crl_urls": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"antiphish": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"domain_controller": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"ldap": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"default_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"check_uri": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"check_basic_auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"check_username_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_body_len": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"inspection_entries": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
										Computed:     true,
									},
									"fortiguard_category": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"custom_patterns": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pattern": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
										Computed:     true,
									},
									"category": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"wisp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wisp_servers": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"wisp_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_all_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_content_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_activex_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_command_block_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_cookie_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_applet_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_jscript_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_js_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_vbs_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_unknown_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_referer_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_filter_cookie_removal_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_url_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_invalid_domain_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_ftgd_err_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_ftgd_quota_usage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_extended_all_action_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_antiphishing_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceWebfilterProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWebfilterProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWebfilterProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WebfilterProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebfilterProfile")
	}

	return resourceWebfilterProfileRead(d, m)
}

func resourceWebfilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWebfilterProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWebfilterProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebfilterProfile")
	}

	return resourceWebfilterProfileRead(d, m)
}

func resourceWebfilterProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWebfilterProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWebfilterProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFeatureSet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileInspectionMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFileFilter(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWebfilterProfileFileFilterStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWebfilterProfileFileFilterLog(i["log"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "scan_archive_contents"
	if _, ok := i["scan-archive-contents"]; ok {
		result["scan_archive_contents"] = flattenWebfilterProfileFileFilterScanArchiveContents(i["scan-archive-contents"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "entries"
	if _, ok := i["entries"]; ok {
		result["entries"] = flattenWebfilterProfileFileFilterEntries(i["entries"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWebfilterProfileFileFilterStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFileFilterLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFileFilterScanArchiveContents(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFileFilterEntries(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if cur_v, ok := i["filter"]; ok {
			tmp["filter"] = flattenWebfilterProfileFileFilterEntriesFilter(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if cur_v, ok := i["comment"]; ok {
			tmp["comment"] = flattenWebfilterProfileFileFilterEntriesComment(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if cur_v, ok := i["protocol"]; ok {
			tmp["protocol"] = flattenWebfilterProfileFileFilterEntriesProtocol(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if cur_v, ok := i["action"]; ok {
			tmp["action"] = flattenWebfilterProfileFileFilterEntriesAction(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if cur_v, ok := i["direction"]; ok {
			tmp["direction"] = flattenWebfilterProfileFileFilterEntriesDirection(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_protected"
		if cur_v, ok := i["password-protected"]; ok {
			tmp["password_protected"] = flattenWebfilterProfileFileFilterEntriesPasswordProtected(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if cur_v, ok := i["file-type"]; ok {
			tmp["file_type"] = flattenWebfilterProfileFileFilterEntriesFileType(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "filter", d)
	return result
}

func flattenWebfilterProfileFileFilterEntriesFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFileFilterEntriesComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFileFilterEntriesProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFileFilterEntriesAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFileFilterEntriesDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFileFilterEntriesPasswordProtected(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFileFilterEntriesFileType(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenWebfilterProfileFileFilterEntriesFileTypeName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWebfilterProfileFileFilterEntriesFileTypeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileHttpsReplacemsg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebFlowLogEncoding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileOvrdPerm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfilePostAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileOverride(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ovrd_cookie"
	if _, ok := i["ovrd-cookie"]; ok {
		result["ovrd_cookie"] = flattenWebfilterProfileOverrideOvrdCookie(i["ovrd-cookie"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ovrd_scope"
	if _, ok := i["ovrd-scope"]; ok {
		result["ovrd_scope"] = flattenWebfilterProfileOverrideOvrdScope(i["ovrd-scope"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "profile_type"
	if _, ok := i["profile-type"]; ok {
		result["profile_type"] = flattenWebfilterProfileOverrideProfileType(i["profile-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ovrd_dur_mode"
	if _, ok := i["ovrd-dur-mode"]; ok {
		result["ovrd_dur_mode"] = flattenWebfilterProfileOverrideOvrdDurMode(i["ovrd-dur-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ovrd_dur"
	if _, ok := i["ovrd-dur"]; ok {
		result["ovrd_dur"] = flattenWebfilterProfileOverrideOvrdDur(i["ovrd-dur"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "profile_attribute"
	if _, ok := i["profile-attribute"]; ok {
		result["profile_attribute"] = flattenWebfilterProfileOverrideProfileAttribute(i["profile-attribute"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ovrd_user_group"
	if _, ok := i["ovrd-user-group"]; ok {
		result["ovrd_user_group"] = flattenWebfilterProfileOverrideOvrdUserGroup(i["ovrd-user-group"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "profile"
	if _, ok := i["profile"]; ok {
		result["profile"] = flattenWebfilterProfileOverrideProfile(i["profile"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWebfilterProfileOverrideOvrdCookie(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideOvrdScope(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideProfileType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideOvrdDurMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideOvrdDur(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideProfileAttribute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideOvrdUserGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenWebfilterProfileOverrideOvrdUserGroupName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWebfilterProfileOverrideOvrdUserGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideProfile(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenWebfilterProfileOverrideProfileName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWebfilterProfileOverrideProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWeb(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bword_threshold"
	if _, ok := i["bword-threshold"]; ok {
		result["bword_threshold"] = flattenWebfilterProfileWebBwordThreshold(i["bword-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bword_table"
	if _, ok := i["bword-table"]; ok {
		result["bword_table"] = flattenWebfilterProfileWebBwordTable(i["bword-table"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "urlfilter_table"
	if _, ok := i["urlfilter-table"]; ok {
		result["urlfilter_table"] = flattenWebfilterProfileWebUrlfilterTable(i["urlfilter-table"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "content_header_list"
	if _, ok := i["content-header-list"]; ok {
		result["content_header_list"] = flattenWebfilterProfileWebContentHeaderList(i["content-header-list"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "blocklist"
	if _, ok := i["blocklist"]; ok {
		result["blocklist"] = flattenWebfilterProfileWebBlocklist(i["blocklist"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "allowlist"
	if _, ok := i["allowlist"]; ok {
		result["allowlist"] = flattenWebfilterProfileWebAllowlist(i["allowlist"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "blacklist"
	if _, ok := i["blacklist"]; ok {
		result["blacklist"] = flattenWebfilterProfileWebBlacklist(i["blacklist"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "whitelist"
	if _, ok := i["whitelist"]; ok {
		result["whitelist"] = flattenWebfilterProfileWebWhitelist(i["whitelist"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "safe_search"
	if _, ok := i["safe-search"]; ok {
		result["safe_search"] = flattenWebfilterProfileWebSafeSearch(i["safe-search"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "youtube_restrict"
	if _, ok := i["youtube-restrict"]; ok {
		result["youtube_restrict"] = flattenWebfilterProfileWebYoutubeRestrict(i["youtube-restrict"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vimeo_restrict"
	if _, ok := i["vimeo-restrict"]; ok {
		result["vimeo_restrict"] = flattenWebfilterProfileWebVimeoRestrict(i["vimeo-restrict"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "log_search"
	if _, ok := i["log-search"]; ok {
		result["log_search"] = flattenWebfilterProfileWebLogSearch(i["log-search"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keyword_match"
	if _, ok := i["keyword-match"]; ok {
		result["keyword_match"] = flattenWebfilterProfileWebKeywordMatch(i["keyword-match"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWebfilterProfileWebBwordThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebBwordTable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebUrlfilterTable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebContentHeaderList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebBlocklist(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebAllowlist(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebBlacklist(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebWhitelist(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebSafeSearch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebYoutubeRestrict(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebVimeoRestrict(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebLogSearch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebKeywordMatch(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if cur_v, ok := i["pattern"]; ok {
			tmp["pattern"] = flattenWebfilterProfileWebKeywordMatchPattern(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "pattern", d)
	return result
}

func flattenWebfilterProfileWebKeywordMatchPattern(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileYoutubeChannelStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileYoutubeChannelFilter(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenWebfilterProfileYoutubeChannelFilterId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "channel_id"
		if cur_v, ok := i["channel-id"]; ok {
			tmp["channel_id"] = flattenWebfilterProfileYoutubeChannelFilterChannelId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if cur_v, ok := i["comment"]; ok {
			tmp["comment"] = flattenWebfilterProfileYoutubeChannelFilterComment(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWebfilterProfileYoutubeChannelFilterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileYoutubeChannelFilterChannelId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileYoutubeChannelFilterComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWf(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenWebfilterProfileFtgdWfOptions(i["options"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "exempt_quota"
	if _, ok := i["exempt-quota"]; ok {
		result["exempt_quota"] = flattenWebfilterProfileFtgdWfExemptQuota(i["exempt-quota"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ovrd"
	if _, ok := i["ovrd"]; ok {
		result["ovrd"] = flattenWebfilterProfileFtgdWfOvrd(i["ovrd"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "filters"
	if _, ok := i["filters"]; ok {
		result["filters"] = flattenWebfilterProfileFtgdWfFilters(i["filters"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "quota"
	if _, ok := i["quota"]; ok {
		result["quota"] = flattenWebfilterProfileFtgdWfQuota(i["quota"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_quota_timeout"
	if _, ok := i["max-quota-timeout"]; ok {
		result["max_quota_timeout"] = flattenWebfilterProfileFtgdWfMaxQuotaTimeout(i["max-quota-timeout"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "rate_image_urls"
	if _, ok := i["rate-image-urls"]; ok {
		result["rate_image_urls"] = flattenWebfilterProfileFtgdWfRateImageUrls(i["rate-image-urls"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "rate_javascript_urls"
	if _, ok := i["rate-javascript-urls"]; ok {
		result["rate_javascript_urls"] = flattenWebfilterProfileFtgdWfRateJavascriptUrls(i["rate-javascript-urls"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "rate_css_urls"
	if _, ok := i["rate-css-urls"]; ok {
		result["rate_css_urls"] = flattenWebfilterProfileFtgdWfRateCssUrls(i["rate-css-urls"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "rate_crl_urls"
	if _, ok := i["rate-crl-urls"]; ok {
		result["rate_crl_urls"] = flattenWebfilterProfileFtgdWfRateCrlUrls(i["rate-crl-urls"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWebfilterProfileFtgdWfOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfExemptQuota(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfOvrd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFilters(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenWebfilterProfileFtgdWfFiltersId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if cur_v, ok := i["category"]; ok {
			tmp["category"] = flattenWebfilterProfileFtgdWfFiltersCategory(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if cur_v, ok := i["action"]; ok {
			tmp["action"] = flattenWebfilterProfileFtgdWfFiltersAction(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warn_duration"
		if cur_v, ok := i["warn-duration"]; ok {
			tmp["warn_duration"] = flattenWebfilterProfileFtgdWfFiltersWarnDuration(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_usr_grp"
		if cur_v, ok := i["auth-usr-grp"]; ok {
			tmp["auth_usr_grp"] = flattenWebfilterProfileFtgdWfFiltersAuthUsrGrp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if cur_v, ok := i["log"]; ok {
			tmp["log"] = flattenWebfilterProfileFtgdWfFiltersLog(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if cur_v, ok := i["override-replacemsg"]; ok {
			tmp["override_replacemsg"] = flattenWebfilterProfileFtgdWfFiltersOverrideReplacemsg(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_prompt"
		if cur_v, ok := i["warning-prompt"]; ok {
			tmp["warning_prompt"] = flattenWebfilterProfileFtgdWfFiltersWarningPrompt(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_duration_type"
		if cur_v, ok := i["warning-duration-type"]; ok {
			tmp["warning_duration_type"] = flattenWebfilterProfileFtgdWfFiltersWarningDurationType(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWebfilterProfileFtgdWfFiltersId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersWarnDuration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersAuthUsrGrp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenWebfilterProfileFtgdWfFiltersAuthUsrGrpName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWebfilterProfileFtgdWfFiltersAuthUsrGrpName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersOverrideReplacemsg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersWarningPrompt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersWarningDurationType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuota(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenWebfilterProfileFtgdWfQuotaId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if cur_v, ok := i["category"]; ok {
			tmp["category"] = flattenWebfilterProfileFtgdWfQuotaCategory(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if cur_v, ok := i["type"]; ok {
			tmp["type"] = flattenWebfilterProfileFtgdWfQuotaType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unit"
		if cur_v, ok := i["unit"]; ok {
			tmp["unit"] = flattenWebfilterProfileFtgdWfQuotaUnit(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if cur_v, ok := i["value"]; ok {
			tmp["value"] = flattenWebfilterProfileFtgdWfQuotaValue(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "duration"
		if cur_v, ok := i["duration"]; ok {
			tmp["duration"] = flattenWebfilterProfileFtgdWfQuotaDuration(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if cur_v, ok := i["override-replacemsg"]; ok {
			tmp["override_replacemsg"] = flattenWebfilterProfileFtgdWfQuotaOverrideReplacemsg(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWebfilterProfileFtgdWfQuotaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaUnit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaDuration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaOverrideReplacemsg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfMaxQuotaTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRateImageUrls(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRateJavascriptUrls(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRateCssUrls(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRateCrlUrls(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphish(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWebfilterProfileAntiphishStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "domain_controller"
	if _, ok := i["domain-controller"]; ok {
		result["domain_controller"] = flattenWebfilterProfileAntiphishDomainController(i["domain-controller"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ldap"
	if _, ok := i["ldap"]; ok {
		result["ldap"] = flattenWebfilterProfileAntiphishLdap(i["ldap"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "default_action"
	if _, ok := i["default-action"]; ok {
		result["default_action"] = flattenWebfilterProfileAntiphishDefaultAction(i["default-action"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "check_uri"
	if _, ok := i["check-uri"]; ok {
		result["check_uri"] = flattenWebfilterProfileAntiphishCheckUri(i["check-uri"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "check_basic_auth"
	if _, ok := i["check-basic-auth"]; ok {
		result["check_basic_auth"] = flattenWebfilterProfileAntiphishCheckBasicAuth(i["check-basic-auth"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "check_username_only"
	if _, ok := i["check-username-only"]; ok {
		result["check_username_only"] = flattenWebfilterProfileAntiphishCheckUsernameOnly(i["check-username-only"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_body_len"
	if _, ok := i["max-body-len"]; ok {
		result["max_body_len"] = flattenWebfilterProfileAntiphishMaxBodyLen(i["max-body-len"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "inspection_entries"
	if _, ok := i["inspection-entries"]; ok {
		result["inspection_entries"] = flattenWebfilterProfileAntiphishInspectionEntries(i["inspection-entries"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "custom_patterns"
	if _, ok := i["custom-patterns"]; ok {
		result["custom_patterns"] = flattenWebfilterProfileAntiphishCustomPatterns(i["custom-patterns"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "authentication"
	if _, ok := i["authentication"]; ok {
		result["authentication"] = flattenWebfilterProfileAntiphishAuthentication(i["authentication"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWebfilterProfileAntiphishStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishDomainController(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishLdap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishDefaultAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishCheckUri(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishCheckBasicAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishCheckUsernameOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishMaxBodyLen(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishInspectionEntries(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenWebfilterProfileAntiphishInspectionEntriesName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard_category"
		if cur_v, ok := i["fortiguard-category"]; ok {
			tmp["fortiguard_category"] = flattenWebfilterProfileAntiphishInspectionEntriesFortiguardCategory(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if cur_v, ok := i["action"]; ok {
			tmp["action"] = flattenWebfilterProfileAntiphishInspectionEntriesAction(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWebfilterProfileAntiphishInspectionEntriesName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishInspectionEntriesFortiguardCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishInspectionEntriesAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishCustomPatterns(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if cur_v, ok := i["pattern"]; ok {
			tmp["pattern"] = flattenWebfilterProfileAntiphishCustomPatternsPattern(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if cur_v, ok := i["category"]; ok {
			tmp["category"] = flattenWebfilterProfileAntiphishCustomPatternsCategory(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if cur_v, ok := i["type"]; ok {
			tmp["type"] = flattenWebfilterProfileAntiphishCustomPatternsType(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "pattern", d)
	return result
}

func flattenWebfilterProfileAntiphishCustomPatternsPattern(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishCustomPatternsCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishCustomPatternsType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileAntiphishAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWisp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWispServers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenWebfilterProfileWispServersName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWebfilterProfileWispServersName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWispAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileLogAllUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebContentLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterActivexLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterCommandBlockLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterCookieLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterAppletLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterJscriptLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterJsLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterVbsLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterUnknownLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterRefererLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterCookieRemovalLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebUrlLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebInvalidDomainLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebFtgdErrLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebFtgdQuotaUsage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileExtendedLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebExtendedAllActionLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterProfileWebAntiphishingLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWebfilterProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenWebfilterProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWebfilterProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("feature_set", flattenWebfilterProfileFeatureSet(o["feature-set"], d, "feature_set", sv)); err != nil {
		if !fortiAPIPatch(o["feature-set"]) {
			return fmt.Errorf("Error reading feature_set: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenWebfilterProfileReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group", sv)); err != nil {
		if !fortiAPIPatch(o["replacemsg-group"]) {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("inspection_mode", flattenWebfilterProfileInspectionMode(o["inspection-mode"], d, "inspection_mode", sv)); err != nil {
		if !fortiAPIPatch(o["inspection-mode"]) {
			return fmt.Errorf("Error reading inspection_mode: %v", err)
		}
	}

	if err = d.Set("options", flattenWebfilterProfileOptions(o["options"], d, "options", sv)); err != nil {
		if !fortiAPIPatch(o["options"]) {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("file_filter", flattenWebfilterProfileFileFilter(o["file-filter"], d, "file_filter", sv)); err != nil {
			if !fortiAPIPatch(o["file-filter"]) {
				return fmt.Errorf("Error reading file_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("file_filter"); ok {
			if err = d.Set("file_filter", flattenWebfilterProfileFileFilter(o["file-filter"], d, "file_filter", sv)); err != nil {
				if !fortiAPIPatch(o["file-filter"]) {
					return fmt.Errorf("Error reading file_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("https_replacemsg", flattenWebfilterProfileHttpsReplacemsg(o["https-replacemsg"], d, "https_replacemsg", sv)); err != nil {
		if !fortiAPIPatch(o["https-replacemsg"]) {
			return fmt.Errorf("Error reading https_replacemsg: %v", err)
		}
	}

	if err = d.Set("web_flow_log_encoding", flattenWebfilterProfileWebFlowLogEncoding(o["web-flow-log-encoding"], d, "web_flow_log_encoding", sv)); err != nil {
		if !fortiAPIPatch(o["web-flow-log-encoding"]) {
			return fmt.Errorf("Error reading web_flow_log_encoding: %v", err)
		}
	}

	if err = d.Set("ovrd_perm", flattenWebfilterProfileOvrdPerm(o["ovrd-perm"], d, "ovrd_perm", sv)); err != nil {
		if !fortiAPIPatch(o["ovrd-perm"]) {
			return fmt.Errorf("Error reading ovrd_perm: %v", err)
		}
	}

	if err = d.Set("post_action", flattenWebfilterProfilePostAction(o["post-action"], d, "post_action", sv)); err != nil {
		if !fortiAPIPatch(o["post-action"]) {
			return fmt.Errorf("Error reading post_action: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("override", flattenWebfilterProfileOverride(o["override"], d, "override", sv)); err != nil {
			if !fortiAPIPatch(o["override"]) {
				return fmt.Errorf("Error reading override: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("override"); ok {
			if err = d.Set("override", flattenWebfilterProfileOverride(o["override"], d, "override", sv)); err != nil {
				if !fortiAPIPatch(o["override"]) {
					return fmt.Errorf("Error reading override: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("web", flattenWebfilterProfileWeb(o["web"], d, "web", sv)); err != nil {
			if !fortiAPIPatch(o["web"]) {
				return fmt.Errorf("Error reading web: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("web"); ok {
			if err = d.Set("web", flattenWebfilterProfileWeb(o["web"], d, "web", sv)); err != nil {
				if !fortiAPIPatch(o["web"]) {
					return fmt.Errorf("Error reading web: %v", err)
				}
			}
		}
	}

	if err = d.Set("youtube_channel_status", flattenWebfilterProfileYoutubeChannelStatus(o["youtube-channel-status"], d, "youtube_channel_status", sv)); err != nil {
		if !fortiAPIPatch(o["youtube-channel-status"]) {
			return fmt.Errorf("Error reading youtube_channel_status: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("youtube_channel_filter", flattenWebfilterProfileYoutubeChannelFilter(o["youtube-channel-filter"], d, "youtube_channel_filter", sv)); err != nil {
			if !fortiAPIPatch(o["youtube-channel-filter"]) {
				return fmt.Errorf("Error reading youtube_channel_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("youtube_channel_filter"); ok {
			if err = d.Set("youtube_channel_filter", flattenWebfilterProfileYoutubeChannelFilter(o["youtube-channel-filter"], d, "youtube_channel_filter", sv)); err != nil {
				if !fortiAPIPatch(o["youtube-channel-filter"]) {
					return fmt.Errorf("Error reading youtube_channel_filter: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("ftgd_wf", flattenWebfilterProfileFtgdWf(o["ftgd-wf"], d, "ftgd_wf", sv)); err != nil {
			if !fortiAPIPatch(o["ftgd-wf"]) {
				return fmt.Errorf("Error reading ftgd_wf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftgd_wf"); ok {
			if err = d.Set("ftgd_wf", flattenWebfilterProfileFtgdWf(o["ftgd-wf"], d, "ftgd_wf", sv)); err != nil {
				if !fortiAPIPatch(o["ftgd-wf"]) {
					return fmt.Errorf("Error reading ftgd_wf: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("antiphish", flattenWebfilterProfileAntiphish(o["antiphish"], d, "antiphish", sv)); err != nil {
			if !fortiAPIPatch(o["antiphish"]) {
				return fmt.Errorf("Error reading antiphish: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("antiphish"); ok {
			if err = d.Set("antiphish", flattenWebfilterProfileAntiphish(o["antiphish"], d, "antiphish", sv)); err != nil {
				if !fortiAPIPatch(o["antiphish"]) {
					return fmt.Errorf("Error reading antiphish: %v", err)
				}
			}
		}
	}

	if err = d.Set("wisp", flattenWebfilterProfileWisp(o["wisp"], d, "wisp", sv)); err != nil {
		if !fortiAPIPatch(o["wisp"]) {
			return fmt.Errorf("Error reading wisp: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("wisp_servers", flattenWebfilterProfileWispServers(o["wisp-servers"], d, "wisp_servers", sv)); err != nil {
			if !fortiAPIPatch(o["wisp-servers"]) {
				return fmt.Errorf("Error reading wisp_servers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("wisp_servers"); ok {
			if err = d.Set("wisp_servers", flattenWebfilterProfileWispServers(o["wisp-servers"], d, "wisp_servers", sv)); err != nil {
				if !fortiAPIPatch(o["wisp-servers"]) {
					return fmt.Errorf("Error reading wisp_servers: %v", err)
				}
			}
		}
	}

	if err = d.Set("wisp_algorithm", flattenWebfilterProfileWispAlgorithm(o["wisp-algorithm"], d, "wisp_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["wisp-algorithm"]) {
			return fmt.Errorf("Error reading wisp_algorithm: %v", err)
		}
	}

	if err = d.Set("log_all_url", flattenWebfilterProfileLogAllUrl(o["log-all-url"], d, "log_all_url", sv)); err != nil {
		if !fortiAPIPatch(o["log-all-url"]) {
			return fmt.Errorf("Error reading log_all_url: %v", err)
		}
	}

	if err = d.Set("web_content_log", flattenWebfilterProfileWebContentLog(o["web-content-log"], d, "web_content_log", sv)); err != nil {
		if !fortiAPIPatch(o["web-content-log"]) {
			return fmt.Errorf("Error reading web_content_log: %v", err)
		}
	}

	if err = d.Set("web_filter_activex_log", flattenWebfilterProfileWebFilterActivexLog(o["web-filter-activex-log"], d, "web_filter_activex_log", sv)); err != nil {
		if !fortiAPIPatch(o["web-filter-activex-log"]) {
			return fmt.Errorf("Error reading web_filter_activex_log: %v", err)
		}
	}

	if err = d.Set("web_filter_command_block_log", flattenWebfilterProfileWebFilterCommandBlockLog(o["web-filter-command-block-log"], d, "web_filter_command_block_log", sv)); err != nil {
		if !fortiAPIPatch(o["web-filter-command-block-log"]) {
			return fmt.Errorf("Error reading web_filter_command_block_log: %v", err)
		}
	}

	if err = d.Set("web_filter_cookie_log", flattenWebfilterProfileWebFilterCookieLog(o["web-filter-cookie-log"], d, "web_filter_cookie_log", sv)); err != nil {
		if !fortiAPIPatch(o["web-filter-cookie-log"]) {
			return fmt.Errorf("Error reading web_filter_cookie_log: %v", err)
		}
	}

	if err = d.Set("web_filter_applet_log", flattenWebfilterProfileWebFilterAppletLog(o["web-filter-applet-log"], d, "web_filter_applet_log", sv)); err != nil {
		if !fortiAPIPatch(o["web-filter-applet-log"]) {
			return fmt.Errorf("Error reading web_filter_applet_log: %v", err)
		}
	}

	if err = d.Set("web_filter_jscript_log", flattenWebfilterProfileWebFilterJscriptLog(o["web-filter-jscript-log"], d, "web_filter_jscript_log", sv)); err != nil {
		if !fortiAPIPatch(o["web-filter-jscript-log"]) {
			return fmt.Errorf("Error reading web_filter_jscript_log: %v", err)
		}
	}

	if err = d.Set("web_filter_js_log", flattenWebfilterProfileWebFilterJsLog(o["web-filter-js-log"], d, "web_filter_js_log", sv)); err != nil {
		if !fortiAPIPatch(o["web-filter-js-log"]) {
			return fmt.Errorf("Error reading web_filter_js_log: %v", err)
		}
	}

	if err = d.Set("web_filter_vbs_log", flattenWebfilterProfileWebFilterVbsLog(o["web-filter-vbs-log"], d, "web_filter_vbs_log", sv)); err != nil {
		if !fortiAPIPatch(o["web-filter-vbs-log"]) {
			return fmt.Errorf("Error reading web_filter_vbs_log: %v", err)
		}
	}

	if err = d.Set("web_filter_unknown_log", flattenWebfilterProfileWebFilterUnknownLog(o["web-filter-unknown-log"], d, "web_filter_unknown_log", sv)); err != nil {
		if !fortiAPIPatch(o["web-filter-unknown-log"]) {
			return fmt.Errorf("Error reading web_filter_unknown_log: %v", err)
		}
	}

	if err = d.Set("web_filter_referer_log", flattenWebfilterProfileWebFilterRefererLog(o["web-filter-referer-log"], d, "web_filter_referer_log", sv)); err != nil {
		if !fortiAPIPatch(o["web-filter-referer-log"]) {
			return fmt.Errorf("Error reading web_filter_referer_log: %v", err)
		}
	}

	if err = d.Set("web_filter_cookie_removal_log", flattenWebfilterProfileWebFilterCookieRemovalLog(o["web-filter-cookie-removal-log"], d, "web_filter_cookie_removal_log", sv)); err != nil {
		if !fortiAPIPatch(o["web-filter-cookie-removal-log"]) {
			return fmt.Errorf("Error reading web_filter_cookie_removal_log: %v", err)
		}
	}

	if err = d.Set("web_url_log", flattenWebfilterProfileWebUrlLog(o["web-url-log"], d, "web_url_log", sv)); err != nil {
		if !fortiAPIPatch(o["web-url-log"]) {
			return fmt.Errorf("Error reading web_url_log: %v", err)
		}
	}

	if err = d.Set("web_invalid_domain_log", flattenWebfilterProfileWebInvalidDomainLog(o["web-invalid-domain-log"], d, "web_invalid_domain_log", sv)); err != nil {
		if !fortiAPIPatch(o["web-invalid-domain-log"]) {
			return fmt.Errorf("Error reading web_invalid_domain_log: %v", err)
		}
	}

	if err = d.Set("web_ftgd_err_log", flattenWebfilterProfileWebFtgdErrLog(o["web-ftgd-err-log"], d, "web_ftgd_err_log", sv)); err != nil {
		if !fortiAPIPatch(o["web-ftgd-err-log"]) {
			return fmt.Errorf("Error reading web_ftgd_err_log: %v", err)
		}
	}

	if err = d.Set("web_ftgd_quota_usage", flattenWebfilterProfileWebFtgdQuotaUsage(o["web-ftgd-quota-usage"], d, "web_ftgd_quota_usage", sv)); err != nil {
		if !fortiAPIPatch(o["web-ftgd-quota-usage"]) {
			return fmt.Errorf("Error reading web_ftgd_quota_usage: %v", err)
		}
	}

	if err = d.Set("extended_log", flattenWebfilterProfileExtendedLog(o["extended-log"], d, "extended_log", sv)); err != nil {
		if !fortiAPIPatch(o["extended-log"]) {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("web_extended_all_action_log", flattenWebfilterProfileWebExtendedAllActionLog(o["web-extended-all-action-log"], d, "web_extended_all_action_log", sv)); err != nil {
		if !fortiAPIPatch(o["web-extended-all-action-log"]) {
			return fmt.Errorf("Error reading web_extended_all_action_log: %v", err)
		}
	}

	if err = d.Set("web_antiphishing_log", flattenWebfilterProfileWebAntiphishingLog(o["web-antiphishing-log"], d, "web_antiphishing_log", sv)); err != nil {
		if !fortiAPIPatch(o["web-antiphishing-log"]) {
			return fmt.Errorf("Error reading web_antiphishing_log: %v", err)
		}
	}

	return nil
}

func flattenWebfilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWebfilterProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFeatureSet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileInspectionMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFileFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWebfilterProfileFileFilterStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandWebfilterProfileFileFilterLog(d, i["log"], pre_append, sv)
	}
	pre_append = pre + ".0." + "scan_archive_contents"
	if _, ok := d.GetOk(pre_append); ok {
		result["scan-archive-contents"], _ = expandWebfilterProfileFileFilterScanArchiveContents(d, i["scan_archive_contents"], pre_append, sv)
	}
	pre_append = pre + ".0." + "entries"
	if _, ok := d.GetOk(pre_append); ok {
		result["entries"], _ = expandWebfilterProfileFileFilterEntries(d, i["entries"], pre_append, sv)
	} else {
		result["entries"] = make([]string, 0)
	}

	return result, nil
}

func expandWebfilterProfileFileFilterStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFileFilterLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFileFilterScanArchiveContents(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFileFilterEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter"], _ = expandWebfilterProfileFileFilterEntriesFilter(d, i["filter"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comment"], _ = expandWebfilterProfileFileFilterEntriesComment(d, i["comment"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandWebfilterProfileFileFilterEntriesProtocol(d, i["protocol"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandWebfilterProfileFileFilterEntriesAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["direction"], _ = expandWebfilterProfileFileFilterEntriesDirection(d, i["direction"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_protected"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["password-protected"], _ = expandWebfilterProfileFileFilterEntriesPasswordProtected(d, i["password_protected"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["file-type"], _ = expandWebfilterProfileFileFilterEntriesFileType(d, i["file_type"], pre_append, sv)
		} else {
			tmp["file-type"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileFileFilterEntriesFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFileFilterEntriesComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFileFilterEntriesProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFileFilterEntriesAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFileFilterEntriesDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFileFilterEntriesPasswordProtected(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFileFilterEntriesFileType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandWebfilterProfileFileFilterEntriesFileTypeName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileFileFilterEntriesFileTypeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileHttpsReplacemsg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFlowLogEncoding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOvrdPerm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfilePostAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ovrd_cookie"
	if _, ok := d.GetOk(pre_append); ok {
		result["ovrd-cookie"], _ = expandWebfilterProfileOverrideOvrdCookie(d, i["ovrd_cookie"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ovrd_scope"
	if _, ok := d.GetOk(pre_append); ok {
		result["ovrd-scope"], _ = expandWebfilterProfileOverrideOvrdScope(d, i["ovrd_scope"], pre_append, sv)
	}
	pre_append = pre + ".0." + "profile_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["profile-type"], _ = expandWebfilterProfileOverrideProfileType(d, i["profile_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ovrd_dur_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["ovrd-dur-mode"], _ = expandWebfilterProfileOverrideOvrdDurMode(d, i["ovrd_dur_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ovrd_dur"
	if _, ok := d.GetOk(pre_append); ok {
		result["ovrd-dur"], _ = expandWebfilterProfileOverrideOvrdDur(d, i["ovrd_dur"], pre_append, sv)
	}
	pre_append = pre + ".0." + "profile_attribute"
	if _, ok := d.GetOk(pre_append); ok {
		result["profile-attribute"], _ = expandWebfilterProfileOverrideProfileAttribute(d, i["profile_attribute"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ovrd_user_group"
	if _, ok := d.GetOk(pre_append); ok {
		result["ovrd-user-group"], _ = expandWebfilterProfileOverrideOvrdUserGroup(d, i["ovrd_user_group"], pre_append, sv)
	} else {
		result["ovrd-user-group"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["profile"], _ = expandWebfilterProfileOverrideProfile(d, i["profile"], pre_append, sv)
	} else {
		result["profile"] = make([]string, 0)
	}

	return result, nil
}

func expandWebfilterProfileOverrideOvrdCookie(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideOvrdScope(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideProfileType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideOvrdDurMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideOvrdDur(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideProfileAttribute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideOvrdUserGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandWebfilterProfileOverrideOvrdUserGroupName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileOverrideOvrdUserGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandWebfilterProfileOverrideProfileName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileOverrideProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWeb(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bword_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["bword-threshold"], _ = expandWebfilterProfileWebBwordThreshold(d, i["bword_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bword_table"
	if _, ok := d.GetOk(pre_append); ok {
		result["bword-table"], _ = expandWebfilterProfileWebBwordTable(d, i["bword_table"], pre_append, sv)
	}
	pre_append = pre + ".0." + "urlfilter_table"
	if _, ok := d.GetOk(pre_append); ok {
		result["urlfilter-table"], _ = expandWebfilterProfileWebUrlfilterTable(d, i["urlfilter_table"], pre_append, sv)
	}
	pre_append = pre + ".0." + "content_header_list"
	if _, ok := d.GetOk(pre_append); ok {
		result["content-header-list"], _ = expandWebfilterProfileWebContentHeaderList(d, i["content_header_list"], pre_append, sv)
	}
	pre_append = pre + ".0." + "blocklist"
	if _, ok := d.GetOk(pre_append); ok {
		result["blocklist"], _ = expandWebfilterProfileWebBlocklist(d, i["blocklist"], pre_append, sv)
	}
	pre_append = pre + ".0." + "allowlist"
	if _, ok := d.GetOk(pre_append); ok {
		result["allowlist"], _ = expandWebfilterProfileWebAllowlist(d, i["allowlist"], pre_append, sv)
	}
	pre_append = pre + ".0." + "blacklist"
	if _, ok := d.GetOk(pre_append); ok {
		result["blacklist"], _ = expandWebfilterProfileWebBlacklist(d, i["blacklist"], pre_append, sv)
	}
	pre_append = pre + ".0." + "whitelist"
	if _, ok := d.GetOk(pre_append); ok {
		result["whitelist"], _ = expandWebfilterProfileWebWhitelist(d, i["whitelist"], pre_append, sv)
	}
	pre_append = pre + ".0." + "safe_search"
	if _, ok := d.GetOk(pre_append); ok {
		result["safe-search"], _ = expandWebfilterProfileWebSafeSearch(d, i["safe_search"], pre_append, sv)
	}
	pre_append = pre + ".0." + "youtube_restrict"
	if _, ok := d.GetOk(pre_append); ok {
		result["youtube-restrict"], _ = expandWebfilterProfileWebYoutubeRestrict(d, i["youtube_restrict"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vimeo_restrict"
	if _, ok := d.GetOk(pre_append); ok {
		result["vimeo-restrict"], _ = expandWebfilterProfileWebVimeoRestrict(d, i["vimeo_restrict"], pre_append, sv)
	}
	pre_append = pre + ".0." + "log_search"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-search"], _ = expandWebfilterProfileWebLogSearch(d, i["log_search"], pre_append, sv)
	}
	pre_append = pre + ".0." + "keyword_match"
	if _, ok := d.GetOk(pre_append); ok {
		result["keyword-match"], _ = expandWebfilterProfileWebKeywordMatch(d, i["keyword_match"], pre_append, sv)
	} else {
		result["keyword-match"] = make([]string, 0)
	}

	return result, nil
}

func expandWebfilterProfileWebBwordThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebBwordTable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebUrlfilterTable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebContentHeaderList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebBlocklist(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebAllowlist(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebBlacklist(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebWhitelist(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebSafeSearch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebYoutubeRestrict(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebVimeoRestrict(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebLogSearch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebKeywordMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["pattern"], _ = expandWebfilterProfileWebKeywordMatchPattern(d, i["pattern"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileWebKeywordMatchPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileYoutubeChannelStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileYoutubeChannelFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandWebfilterProfileYoutubeChannelFilterId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "channel_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["channel-id"], _ = expandWebfilterProfileYoutubeChannelFilterChannelId(d, i["channel_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comment"], _ = expandWebfilterProfileYoutubeChannelFilterComment(d, i["comment"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileYoutubeChannelFilterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileYoutubeChannelFilterChannelId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileYoutubeChannelFilterComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandWebfilterProfileFtgdWfOptions(d, i["options"], pre_append, sv)
	}
	pre_append = pre + ".0." + "exempt_quota"
	if _, ok := d.GetOk(pre_append); ok {
		result["exempt-quota"], _ = expandWebfilterProfileFtgdWfExemptQuota(d, i["exempt_quota"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ovrd"
	if _, ok := d.GetOk(pre_append); ok {
		result["ovrd"], _ = expandWebfilterProfileFtgdWfOvrd(d, i["ovrd"], pre_append, sv)
	}
	pre_append = pre + ".0." + "filters"
	if _, ok := d.GetOk(pre_append); ok {
		result["filters"], _ = expandWebfilterProfileFtgdWfFilters(d, i["filters"], pre_append, sv)
	} else {
		result["filters"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "quota"
	if _, ok := d.GetOk(pre_append); ok {
		result["quota"], _ = expandWebfilterProfileFtgdWfQuota(d, i["quota"], pre_append, sv)
	} else {
		result["quota"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "max_quota_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-quota-timeout"], _ = expandWebfilterProfileFtgdWfMaxQuotaTimeout(d, i["max_quota_timeout"], pre_append, sv)
	}
	pre_append = pre + ".0." + "rate_image_urls"
	if _, ok := d.GetOk(pre_append); ok {
		result["rate-image-urls"], _ = expandWebfilterProfileFtgdWfRateImageUrls(d, i["rate_image_urls"], pre_append, sv)
	}
	pre_append = pre + ".0." + "rate_javascript_urls"
	if _, ok := d.GetOk(pre_append); ok {
		result["rate-javascript-urls"], _ = expandWebfilterProfileFtgdWfRateJavascriptUrls(d, i["rate_javascript_urls"], pre_append, sv)
	}
	pre_append = pre + ".0." + "rate_css_urls"
	if _, ok := d.GetOk(pre_append); ok {
		result["rate-css-urls"], _ = expandWebfilterProfileFtgdWfRateCssUrls(d, i["rate_css_urls"], pre_append, sv)
	}
	pre_append = pre + ".0." + "rate_crl_urls"
	if _, ok := d.GetOk(pre_append); ok {
		result["rate-crl-urls"], _ = expandWebfilterProfileFtgdWfRateCrlUrls(d, i["rate_crl_urls"], pre_append, sv)
	}

	return result, nil
}

func expandWebfilterProfileFtgdWfOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfExemptQuota(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfOvrd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFilters(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandWebfilterProfileFtgdWfFiltersId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category"], _ = expandWebfilterProfileFtgdWfFiltersCategory(d, i["category"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandWebfilterProfileFtgdWfFiltersAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warn_duration"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["warn-duration"], _ = expandWebfilterProfileFtgdWfFiltersWarnDuration(d, i["warn_duration"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_usr_grp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-usr-grp"], _ = expandWebfilterProfileFtgdWfFiltersAuthUsrGrp(d, i["auth_usr_grp"], pre_append, sv)
		} else {
			tmp["auth-usr-grp"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log"], _ = expandWebfilterProfileFtgdWfFiltersLog(d, i["log"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["override-replacemsg"], _ = expandWebfilterProfileFtgdWfFiltersOverrideReplacemsg(d, i["override_replacemsg"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_prompt"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["warning-prompt"], _ = expandWebfilterProfileFtgdWfFiltersWarningPrompt(d, i["warning_prompt"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_duration_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["warning-duration-type"], _ = expandWebfilterProfileFtgdWfFiltersWarningDurationType(d, i["warning_duration_type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileFtgdWfFiltersId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersWarnDuration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersAuthUsrGrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandWebfilterProfileFtgdWfFiltersAuthUsrGrpName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileFtgdWfFiltersAuthUsrGrpName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersOverrideReplacemsg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersWarningPrompt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersWarningDurationType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuota(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandWebfilterProfileFtgdWfQuotaId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category"], _ = expandWebfilterProfileFtgdWfQuotaCategory(d, i["category"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandWebfilterProfileFtgdWfQuotaType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unit"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["unit"], _ = expandWebfilterProfileFtgdWfQuotaUnit(d, i["unit"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandWebfilterProfileFtgdWfQuotaValue(d, i["value"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "duration"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["duration"], _ = expandWebfilterProfileFtgdWfQuotaDuration(d, i["duration"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["override-replacemsg"], _ = expandWebfilterProfileFtgdWfQuotaOverrideReplacemsg(d, i["override_replacemsg"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileFtgdWfQuotaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaUnit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaDuration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaOverrideReplacemsg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfMaxQuotaTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRateImageUrls(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRateJavascriptUrls(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRateCssUrls(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRateCrlUrls(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphish(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWebfilterProfileAntiphishStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "domain_controller"
	if _, ok := d.GetOk(pre_append); ok {
		result["domain-controller"], _ = expandWebfilterProfileAntiphishDomainController(d, i["domain_controller"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ldap"
	if _, ok := d.GetOk(pre_append); ok {
		result["ldap"], _ = expandWebfilterProfileAntiphishLdap(d, i["ldap"], pre_append, sv)
	}
	pre_append = pre + ".0." + "default_action"
	if _, ok := d.GetOk(pre_append); ok {
		result["default-action"], _ = expandWebfilterProfileAntiphishDefaultAction(d, i["default_action"], pre_append, sv)
	}
	pre_append = pre + ".0." + "check_uri"
	if _, ok := d.GetOk(pre_append); ok {
		result["check-uri"], _ = expandWebfilterProfileAntiphishCheckUri(d, i["check_uri"], pre_append, sv)
	}
	pre_append = pre + ".0." + "check_basic_auth"
	if _, ok := d.GetOk(pre_append); ok {
		result["check-basic-auth"], _ = expandWebfilterProfileAntiphishCheckBasicAuth(d, i["check_basic_auth"], pre_append, sv)
	}
	pre_append = pre + ".0." + "check_username_only"
	if _, ok := d.GetOk(pre_append); ok {
		result["check-username-only"], _ = expandWebfilterProfileAntiphishCheckUsernameOnly(d, i["check_username_only"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_body_len"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-body-len"], _ = expandWebfilterProfileAntiphishMaxBodyLen(d, i["max_body_len"], pre_append, sv)
	}
	pre_append = pre + ".0." + "inspection_entries"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspection-entries"], _ = expandWebfilterProfileAntiphishInspectionEntries(d, i["inspection_entries"], pre_append, sv)
	} else {
		result["inspection-entries"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "custom_patterns"
	if _, ok := d.GetOk(pre_append); ok {
		result["custom-patterns"], _ = expandWebfilterProfileAntiphishCustomPatterns(d, i["custom_patterns"], pre_append, sv)
	} else {
		result["custom-patterns"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "authentication"
	if _, ok := d.GetOk(pre_append); ok {
		result["authentication"], _ = expandWebfilterProfileAntiphishAuthentication(d, i["authentication"], pre_append, sv)
	}

	return result, nil
}

func expandWebfilterProfileAntiphishStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishDomainController(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishLdap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishDefaultAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishCheckUri(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishCheckBasicAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishCheckUsernameOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishMaxBodyLen(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishInspectionEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandWebfilterProfileAntiphishInspectionEntriesName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard_category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fortiguard-category"], _ = expandWebfilterProfileAntiphishInspectionEntriesFortiguardCategory(d, i["fortiguard_category"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandWebfilterProfileAntiphishInspectionEntriesAction(d, i["action"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileAntiphishInspectionEntriesName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishInspectionEntriesFortiguardCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishInspectionEntriesAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishCustomPatterns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pattern"], _ = expandWebfilterProfileAntiphishCustomPatternsPattern(d, i["pattern"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category"], _ = expandWebfilterProfileAntiphishCustomPatternsCategory(d, i["category"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandWebfilterProfileAntiphishCustomPatternsType(d, i["type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileAntiphishCustomPatternsPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishCustomPatternsCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishCustomPatternsType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileAntiphishAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWisp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWispServers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandWebfilterProfileWispServersName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileWispServersName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWispAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileLogAllUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebContentLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterActivexLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterCommandBlockLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterCookieLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterAppletLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterJscriptLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterJsLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterVbsLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterUnknownLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterRefererLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterCookieRemovalLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebUrlLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebInvalidDomainLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFtgdErrLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFtgdQuotaUsage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileExtendedLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebExtendedAllActionLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebAntiphishingLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWebfilterProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWebfilterProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("feature_set"); ok {
		t, err := expandWebfilterProfileFeatureSet(d, v, "feature_set", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["feature-set"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok {
		t, err := expandWebfilterProfileReplacemsgGroup(d, v, "replacemsg_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("inspection_mode"); ok {
		t, err := expandWebfilterProfileInspectionMode(d, v, "inspection_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inspection-mode"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok {
		t, err := expandWebfilterProfileOptions(d, v, "options", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("file_filter"); ok {
		t, err := expandWebfilterProfileFileFilter(d, v, "file_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-filter"] = t
		}
	}

	if v, ok := d.GetOk("https_replacemsg"); ok {
		t, err := expandWebfilterProfileHttpsReplacemsg(d, v, "https_replacemsg", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-replacemsg"] = t
		}
	}

	if v, ok := d.GetOk("web_flow_log_encoding"); ok {
		t, err := expandWebfilterProfileWebFlowLogEncoding(d, v, "web_flow_log_encoding", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-flow-log-encoding"] = t
		}
	}

	if v, ok := d.GetOk("ovrd_perm"); ok {
		t, err := expandWebfilterProfileOvrdPerm(d, v, "ovrd_perm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-perm"] = t
		}
	}

	if v, ok := d.GetOk("post_action"); ok {
		t, err := expandWebfilterProfilePostAction(d, v, "post_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["post-action"] = t
		}
	}

	if v, ok := d.GetOk("override"); ok {
		t, err := expandWebfilterProfileOverride(d, v, "override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override"] = t
		}
	}

	if v, ok := d.GetOk("web"); ok {
		t, err := expandWebfilterProfileWeb(d, v, "web", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web"] = t
		}
	}

	if v, ok := d.GetOk("youtube_channel_status"); ok {
		t, err := expandWebfilterProfileYoutubeChannelStatus(d, v, "youtube_channel_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube-channel-status"] = t
		}
	}

	if v, ok := d.GetOk("youtube_channel_filter"); ok || d.HasChange("youtube_channel_filter") {
		t, err := expandWebfilterProfileYoutubeChannelFilter(d, v, "youtube_channel_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube-channel-filter"] = t
		}
	}

	if v, ok := d.GetOk("ftgd_wf"); ok {
		t, err := expandWebfilterProfileFtgdWf(d, v, "ftgd_wf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftgd-wf"] = t
		}
	}

	if v, ok := d.GetOk("antiphish"); ok {
		t, err := expandWebfilterProfileAntiphish(d, v, "antiphish", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antiphish"] = t
		}
	}

	if v, ok := d.GetOk("wisp"); ok {
		t, err := expandWebfilterProfileWisp(d, v, "wisp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wisp"] = t
		}
	}

	if v, ok := d.GetOk("wisp_servers"); ok || d.HasChange("wisp_servers") {
		t, err := expandWebfilterProfileWispServers(d, v, "wisp_servers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wisp-servers"] = t
		}
	}

	if v, ok := d.GetOk("wisp_algorithm"); ok {
		t, err := expandWebfilterProfileWispAlgorithm(d, v, "wisp_algorithm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wisp-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("log_all_url"); ok {
		t, err := expandWebfilterProfileLogAllUrl(d, v, "log_all_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-all-url"] = t
		}
	}

	if v, ok := d.GetOk("web_content_log"); ok {
		t, err := expandWebfilterProfileWebContentLog(d, v, "web_content_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-content-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_activex_log"); ok {
		t, err := expandWebfilterProfileWebFilterActivexLog(d, v, "web_filter_activex_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-activex-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_command_block_log"); ok {
		t, err := expandWebfilterProfileWebFilterCommandBlockLog(d, v, "web_filter_command_block_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-command-block-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_cookie_log"); ok {
		t, err := expandWebfilterProfileWebFilterCookieLog(d, v, "web_filter_cookie_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-cookie-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_applet_log"); ok {
		t, err := expandWebfilterProfileWebFilterAppletLog(d, v, "web_filter_applet_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-applet-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_jscript_log"); ok {
		t, err := expandWebfilterProfileWebFilterJscriptLog(d, v, "web_filter_jscript_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-jscript-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_js_log"); ok {
		t, err := expandWebfilterProfileWebFilterJsLog(d, v, "web_filter_js_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-js-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_vbs_log"); ok {
		t, err := expandWebfilterProfileWebFilterVbsLog(d, v, "web_filter_vbs_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-vbs-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_unknown_log"); ok {
		t, err := expandWebfilterProfileWebFilterUnknownLog(d, v, "web_filter_unknown_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-unknown-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_referer_log"); ok {
		t, err := expandWebfilterProfileWebFilterRefererLog(d, v, "web_filter_referer_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-referer-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_cookie_removal_log"); ok {
		t, err := expandWebfilterProfileWebFilterCookieRemovalLog(d, v, "web_filter_cookie_removal_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-cookie-removal-log"] = t
		}
	}

	if v, ok := d.GetOk("web_url_log"); ok {
		t, err := expandWebfilterProfileWebUrlLog(d, v, "web_url_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-url-log"] = t
		}
	}

	if v, ok := d.GetOk("web_invalid_domain_log"); ok {
		t, err := expandWebfilterProfileWebInvalidDomainLog(d, v, "web_invalid_domain_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-invalid-domain-log"] = t
		}
	}

	if v, ok := d.GetOk("web_ftgd_err_log"); ok {
		t, err := expandWebfilterProfileWebFtgdErrLog(d, v, "web_ftgd_err_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-ftgd-err-log"] = t
		}
	}

	if v, ok := d.GetOk("web_ftgd_quota_usage"); ok {
		t, err := expandWebfilterProfileWebFtgdQuotaUsage(d, v, "web_ftgd_quota_usage", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-ftgd-quota-usage"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok {
		t, err := expandWebfilterProfileExtendedLog(d, v, "extended_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("web_extended_all_action_log"); ok {
		t, err := expandWebfilterProfileWebExtendedAllActionLog(d, v, "web_extended_all_action_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-extended-all-action-log"] = t
		}
	}

	if v, ok := d.GetOk("web_antiphishing_log"); ok {
		t, err := expandWebfilterProfileWebAntiphishingLog(d, v, "web_antiphishing_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-antiphishing-log"] = t
		}
	}

	return &obj, nil
}
