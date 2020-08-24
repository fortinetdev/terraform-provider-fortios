// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure Web filter profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
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
			"https_replacemsg": &schema.Schema{
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
						"profile": &schema.Schema{
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
					},
				},
			}, "web": &schema.Schema{
				Type:     schema.TypeList,
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
						"log_search": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keyword_match": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pattern": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 64),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
					},
				},
			}, "youtube_channel_status": &schema.Schema{
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
			}, "wisp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wisp_servers": &schema.Schema{
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
		},
	}
}

func resourceWebfilterProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWebfilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWebfilterProfile(obj)

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

	obj, err := getObjectWebfilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWebfilterProfile(obj, mkey)
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

	err := c.DeleteWebfilterProfile(mkey)
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

	o, err := c.ReadWebfilterProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileInspectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileHttpsReplacemsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOvrdPerm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfilePostAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOverride(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ovrd_cookie"
	if _, ok := i["ovrd-cookie"]; ok {
		result["ovrd_cookie"] = flattenWebfilterProfileOverrideOvrdCookie(i["ovrd-cookie"], d, pre_append)
	}

	pre_append = pre + ".0." + "ovrd_scope"
	if _, ok := i["ovrd-scope"]; ok {
		result["ovrd_scope"] = flattenWebfilterProfileOverrideOvrdScope(i["ovrd-scope"], d, pre_append)
	}

	pre_append = pre + ".0." + "profile_type"
	if _, ok := i["profile-type"]; ok {
		result["profile_type"] = flattenWebfilterProfileOverrideProfileType(i["profile-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "ovrd_dur_mode"
	if _, ok := i["ovrd-dur-mode"]; ok {
		result["ovrd_dur_mode"] = flattenWebfilterProfileOverrideOvrdDurMode(i["ovrd-dur-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ovrd_dur"
	if _, ok := i["ovrd-dur"]; ok {
		result["ovrd_dur"] = flattenWebfilterProfileOverrideOvrdDur(i["ovrd-dur"], d, pre_append)
	}

	pre_append = pre + ".0." + "profile_attribute"
	if _, ok := i["profile-attribute"]; ok {
		result["profile_attribute"] = flattenWebfilterProfileOverrideProfileAttribute(i["profile-attribute"], d, pre_append)
	}

	pre_append = pre + ".0." + "ovrd_user_group"
	if _, ok := i["ovrd-user-group"]; ok {
		result["ovrd_user_group"] = flattenWebfilterProfileOverrideOvrdUserGroup(i["ovrd-user-group"], d, pre_append)
	}

	pre_append = pre + ".0." + "profile"
	if _, ok := i["profile"]; ok {
		result["profile"] = flattenWebfilterProfileOverrideProfile(i["profile"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWebfilterProfileOverrideOvrdCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideOvrdScope(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideProfileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideOvrdDurMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideOvrdDur(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideProfileAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideOvrdUserGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenWebfilterProfileOverrideOvrdUserGroupName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWebfilterProfileOverrideOvrdUserGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileOverrideProfile(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenWebfilterProfileOverrideProfileName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWebfilterProfileOverrideProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWeb(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bword_threshold"
	if _, ok := i["bword-threshold"]; ok {
		result["bword_threshold"] = flattenWebfilterProfileWebBwordThreshold(i["bword-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "bword_table"
	if _, ok := i["bword-table"]; ok {
		result["bword_table"] = flattenWebfilterProfileWebBwordTable(i["bword-table"], d, pre_append)
	}

	pre_append = pre + ".0." + "urlfilter_table"
	if _, ok := i["urlfilter-table"]; ok {
		result["urlfilter_table"] = flattenWebfilterProfileWebUrlfilterTable(i["urlfilter-table"], d, pre_append)
	}

	pre_append = pre + ".0." + "content_header_list"
	if _, ok := i["content-header-list"]; ok {
		result["content_header_list"] = flattenWebfilterProfileWebContentHeaderList(i["content-header-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "blacklist"
	if _, ok := i["blacklist"]; ok {
		result["blacklist"] = flattenWebfilterProfileWebBlacklist(i["blacklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "whitelist"
	if _, ok := i["whitelist"]; ok {
		result["whitelist"] = flattenWebfilterProfileWebWhitelist(i["whitelist"], d, pre_append)
	}

	pre_append = pre + ".0." + "safe_search"
	if _, ok := i["safe-search"]; ok {
		result["safe_search"] = flattenWebfilterProfileWebSafeSearch(i["safe-search"], d, pre_append)
	}

	pre_append = pre + ".0." + "youtube_restrict"
	if _, ok := i["youtube-restrict"]; ok {
		result["youtube_restrict"] = flattenWebfilterProfileWebYoutubeRestrict(i["youtube-restrict"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_search"
	if _, ok := i["log-search"]; ok {
		result["log_search"] = flattenWebfilterProfileWebLogSearch(i["log-search"], d, pre_append)
	}

	pre_append = pre + ".0." + "keyword_match"
	if _, ok := i["keyword-match"]; ok {
		result["keyword_match"] = flattenWebfilterProfileWebKeywordMatch(i["keyword-match"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWebfilterProfileWebBwordThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebBwordTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebUrlfilterTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebContentHeaderList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebBlacklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebWhitelist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebSafeSearch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebYoutubeRestrict(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebLogSearch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebKeywordMatch(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			tmp["pattern"] = flattenWebfilterProfileWebKeywordMatchPattern(i["pattern"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWebfilterProfileWebKeywordMatchPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileYoutubeChannelStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileYoutubeChannelFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWebfilterProfileYoutubeChannelFilterId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "channel_id"
		if _, ok := i["channel-id"]; ok {
			tmp["channel_id"] = flattenWebfilterProfileYoutubeChannelFilterChannelId(i["channel-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			tmp["comment"] = flattenWebfilterProfileYoutubeChannelFilterComment(i["comment"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWebfilterProfileYoutubeChannelFilterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileYoutubeChannelFilterChannelId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileYoutubeChannelFilterComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWf(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenWebfilterProfileFtgdWfOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "exempt_quota"
	if _, ok := i["exempt-quota"]; ok {
		result["exempt_quota"] = flattenWebfilterProfileFtgdWfExemptQuota(i["exempt-quota"], d, pre_append)
	}

	pre_append = pre + ".0." + "ovrd"
	if _, ok := i["ovrd"]; ok {
		result["ovrd"] = flattenWebfilterProfileFtgdWfOvrd(i["ovrd"], d, pre_append)
	}

	pre_append = pre + ".0." + "filters"
	if _, ok := i["filters"]; ok {
		result["filters"] = flattenWebfilterProfileFtgdWfFilters(i["filters"], d, pre_append)
	}

	pre_append = pre + ".0." + "quota"
	if _, ok := i["quota"]; ok {
		result["quota"] = flattenWebfilterProfileFtgdWfQuota(i["quota"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_quota_timeout"
	if _, ok := i["max-quota-timeout"]; ok {
		result["max_quota_timeout"] = flattenWebfilterProfileFtgdWfMaxQuotaTimeout(i["max-quota-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "rate_image_urls"
	if _, ok := i["rate-image-urls"]; ok {
		result["rate_image_urls"] = flattenWebfilterProfileFtgdWfRateImageUrls(i["rate-image-urls"], d, pre_append)
	}

	pre_append = pre + ".0." + "rate_javascript_urls"
	if _, ok := i["rate-javascript-urls"]; ok {
		result["rate_javascript_urls"] = flattenWebfilterProfileFtgdWfRateJavascriptUrls(i["rate-javascript-urls"], d, pre_append)
	}

	pre_append = pre + ".0." + "rate_css_urls"
	if _, ok := i["rate-css-urls"]; ok {
		result["rate_css_urls"] = flattenWebfilterProfileFtgdWfRateCssUrls(i["rate-css-urls"], d, pre_append)
	}

	pre_append = pre + ".0." + "rate_crl_urls"
	if _, ok := i["rate-crl-urls"]; ok {
		result["rate_crl_urls"] = flattenWebfilterProfileFtgdWfRateCrlUrls(i["rate-crl-urls"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWebfilterProfileFtgdWfOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfExemptQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfOvrd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFilters(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWebfilterProfileFtgdWfFiltersId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			tmp["category"] = flattenWebfilterProfileFtgdWfFiltersCategory(i["category"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = flattenWebfilterProfileFtgdWfFiltersAction(i["action"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warn_duration"
		if _, ok := i["warn-duration"]; ok {
			tmp["warn_duration"] = flattenWebfilterProfileFtgdWfFiltersWarnDuration(i["warn-duration"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_usr_grp"
		if _, ok := i["auth-usr-grp"]; ok {
			tmp["auth_usr_grp"] = flattenWebfilterProfileFtgdWfFiltersAuthUsrGrp(i["auth-usr-grp"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			tmp["log"] = flattenWebfilterProfileFtgdWfFiltersLog(i["log"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := i["override-replacemsg"]; ok {
			tmp["override_replacemsg"] = flattenWebfilterProfileFtgdWfFiltersOverrideReplacemsg(i["override-replacemsg"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_prompt"
		if _, ok := i["warning-prompt"]; ok {
			tmp["warning_prompt"] = flattenWebfilterProfileFtgdWfFiltersWarningPrompt(i["warning-prompt"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_duration_type"
		if _, ok := i["warning-duration-type"]; ok {
			tmp["warning_duration_type"] = flattenWebfilterProfileFtgdWfFiltersWarningDurationType(i["warning-duration-type"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWebfilterProfileFtgdWfFiltersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersWarnDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersAuthUsrGrp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenWebfilterProfileFtgdWfFiltersAuthUsrGrpName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWebfilterProfileFtgdWfFiltersAuthUsrGrpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersOverrideReplacemsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersWarningPrompt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfFiltersWarningDurationType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuota(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWebfilterProfileFtgdWfQuotaId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			tmp["category"] = flattenWebfilterProfileFtgdWfQuotaCategory(i["category"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = flattenWebfilterProfileFtgdWfQuotaType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unit"
		if _, ok := i["unit"]; ok {
			tmp["unit"] = flattenWebfilterProfileFtgdWfQuotaUnit(i["unit"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			tmp["value"] = flattenWebfilterProfileFtgdWfQuotaValue(i["value"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "duration"
		if _, ok := i["duration"]; ok {
			tmp["duration"] = flattenWebfilterProfileFtgdWfQuotaDuration(i["duration"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := i["override-replacemsg"]; ok {
			tmp["override_replacemsg"] = flattenWebfilterProfileFtgdWfQuotaOverrideReplacemsg(i["override-replacemsg"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWebfilterProfileFtgdWfQuotaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaUnit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfQuotaOverrideReplacemsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfMaxQuotaTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRateImageUrls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRateJavascriptUrls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRateCssUrls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileFtgdWfRateCrlUrls(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWisp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWispServers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenWebfilterProfileWispServersName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWebfilterProfileWispServersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWispAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileLogAllUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebContentLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterActivexLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterCommandBlockLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterCookieLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterAppletLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterJscriptLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterJsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterVbsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterUnknownLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterRefererLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFilterCookieRemovalLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebUrlLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebInvalidDomainLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFtgdErrLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebFtgdQuotaUsage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterProfileWebExtendedAllActionLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWebfilterProfileName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWebfilterProfileComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenWebfilterProfileReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if !fortiAPIPatch(o["replacemsg-group"]) {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("inspection_mode", flattenWebfilterProfileInspectionMode(o["inspection-mode"], d, "inspection_mode")); err != nil {
		if !fortiAPIPatch(o["inspection-mode"]) {
			return fmt.Errorf("Error reading inspection_mode: %v", err)
		}
	}

	if err = d.Set("options", flattenWebfilterProfileOptions(o["options"], d, "options")); err != nil {
		if !fortiAPIPatch(o["options"]) {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("https_replacemsg", flattenWebfilterProfileHttpsReplacemsg(o["https-replacemsg"], d, "https_replacemsg")); err != nil {
		if !fortiAPIPatch(o["https-replacemsg"]) {
			return fmt.Errorf("Error reading https_replacemsg: %v", err)
		}
	}

	if err = d.Set("ovrd_perm", flattenWebfilterProfileOvrdPerm(o["ovrd-perm"], d, "ovrd_perm")); err != nil {
		if !fortiAPIPatch(o["ovrd-perm"]) {
			return fmt.Errorf("Error reading ovrd_perm: %v", err)
		}
	}

	if err = d.Set("post_action", flattenWebfilterProfilePostAction(o["post-action"], d, "post_action")); err != nil {
		if !fortiAPIPatch(o["post-action"]) {
			return fmt.Errorf("Error reading post_action: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("override", flattenWebfilterProfileOverride(o["override"], d, "override")); err != nil {
			if !fortiAPIPatch(o["override"]) {
				return fmt.Errorf("Error reading override: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("override"); ok {
			if err = d.Set("override", flattenWebfilterProfileOverride(o["override"], d, "override")); err != nil {
				if !fortiAPIPatch(o["override"]) {
					return fmt.Errorf("Error reading override: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("web", flattenWebfilterProfileWeb(o["web"], d, "web")); err != nil {
			if !fortiAPIPatch(o["web"]) {
				return fmt.Errorf("Error reading web: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("web"); ok {
			if err = d.Set("web", flattenWebfilterProfileWeb(o["web"], d, "web")); err != nil {
				if !fortiAPIPatch(o["web"]) {
					return fmt.Errorf("Error reading web: %v", err)
				}
			}
		}
	}

	if err = d.Set("youtube_channel_status", flattenWebfilterProfileYoutubeChannelStatus(o["youtube-channel-status"], d, "youtube_channel_status")); err != nil {
		if !fortiAPIPatch(o["youtube-channel-status"]) {
			return fmt.Errorf("Error reading youtube_channel_status: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("youtube_channel_filter", flattenWebfilterProfileYoutubeChannelFilter(o["youtube-channel-filter"], d, "youtube_channel_filter")); err != nil {
			if !fortiAPIPatch(o["youtube-channel-filter"]) {
				return fmt.Errorf("Error reading youtube_channel_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("youtube_channel_filter"); ok {
			if err = d.Set("youtube_channel_filter", flattenWebfilterProfileYoutubeChannelFilter(o["youtube-channel-filter"], d, "youtube_channel_filter")); err != nil {
				if !fortiAPIPatch(o["youtube-channel-filter"]) {
					return fmt.Errorf("Error reading youtube_channel_filter: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ftgd_wf", flattenWebfilterProfileFtgdWf(o["ftgd-wf"], d, "ftgd_wf")); err != nil {
			if !fortiAPIPatch(o["ftgd-wf"]) {
				return fmt.Errorf("Error reading ftgd_wf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftgd_wf"); ok {
			if err = d.Set("ftgd_wf", flattenWebfilterProfileFtgdWf(o["ftgd-wf"], d, "ftgd_wf")); err != nil {
				if !fortiAPIPatch(o["ftgd-wf"]) {
					return fmt.Errorf("Error reading ftgd_wf: %v", err)
				}
			}
		}
	}

	if err = d.Set("wisp", flattenWebfilterProfileWisp(o["wisp"], d, "wisp")); err != nil {
		if !fortiAPIPatch(o["wisp"]) {
			return fmt.Errorf("Error reading wisp: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("wisp_servers", flattenWebfilterProfileWispServers(o["wisp-servers"], d, "wisp_servers")); err != nil {
			if !fortiAPIPatch(o["wisp-servers"]) {
				return fmt.Errorf("Error reading wisp_servers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("wisp_servers"); ok {
			if err = d.Set("wisp_servers", flattenWebfilterProfileWispServers(o["wisp-servers"], d, "wisp_servers")); err != nil {
				if !fortiAPIPatch(o["wisp-servers"]) {
					return fmt.Errorf("Error reading wisp_servers: %v", err)
				}
			}
		}
	}

	if err = d.Set("wisp_algorithm", flattenWebfilterProfileWispAlgorithm(o["wisp-algorithm"], d, "wisp_algorithm")); err != nil {
		if !fortiAPIPatch(o["wisp-algorithm"]) {
			return fmt.Errorf("Error reading wisp_algorithm: %v", err)
		}
	}

	if err = d.Set("log_all_url", flattenWebfilterProfileLogAllUrl(o["log-all-url"], d, "log_all_url")); err != nil {
		if !fortiAPIPatch(o["log-all-url"]) {
			return fmt.Errorf("Error reading log_all_url: %v", err)
		}
	}

	if err = d.Set("web_content_log", flattenWebfilterProfileWebContentLog(o["web-content-log"], d, "web_content_log")); err != nil {
		if !fortiAPIPatch(o["web-content-log"]) {
			return fmt.Errorf("Error reading web_content_log: %v", err)
		}
	}

	if err = d.Set("web_filter_activex_log", flattenWebfilterProfileWebFilterActivexLog(o["web-filter-activex-log"], d, "web_filter_activex_log")); err != nil {
		if !fortiAPIPatch(o["web-filter-activex-log"]) {
			return fmt.Errorf("Error reading web_filter_activex_log: %v", err)
		}
	}

	if err = d.Set("web_filter_command_block_log", flattenWebfilterProfileWebFilterCommandBlockLog(o["web-filter-command-block-log"], d, "web_filter_command_block_log")); err != nil {
		if !fortiAPIPatch(o["web-filter-command-block-log"]) {
			return fmt.Errorf("Error reading web_filter_command_block_log: %v", err)
		}
	}

	if err = d.Set("web_filter_cookie_log", flattenWebfilterProfileWebFilterCookieLog(o["web-filter-cookie-log"], d, "web_filter_cookie_log")); err != nil {
		if !fortiAPIPatch(o["web-filter-cookie-log"]) {
			return fmt.Errorf("Error reading web_filter_cookie_log: %v", err)
		}
	}

	if err = d.Set("web_filter_applet_log", flattenWebfilterProfileWebFilterAppletLog(o["web-filter-applet-log"], d, "web_filter_applet_log")); err != nil {
		if !fortiAPIPatch(o["web-filter-applet-log"]) {
			return fmt.Errorf("Error reading web_filter_applet_log: %v", err)
		}
	}

	if err = d.Set("web_filter_jscript_log", flattenWebfilterProfileWebFilterJscriptLog(o["web-filter-jscript-log"], d, "web_filter_jscript_log")); err != nil {
		if !fortiAPIPatch(o["web-filter-jscript-log"]) {
			return fmt.Errorf("Error reading web_filter_jscript_log: %v", err)
		}
	}

	if err = d.Set("web_filter_js_log", flattenWebfilterProfileWebFilterJsLog(o["web-filter-js-log"], d, "web_filter_js_log")); err != nil {
		if !fortiAPIPatch(o["web-filter-js-log"]) {
			return fmt.Errorf("Error reading web_filter_js_log: %v", err)
		}
	}

	if err = d.Set("web_filter_vbs_log", flattenWebfilterProfileWebFilterVbsLog(o["web-filter-vbs-log"], d, "web_filter_vbs_log")); err != nil {
		if !fortiAPIPatch(o["web-filter-vbs-log"]) {
			return fmt.Errorf("Error reading web_filter_vbs_log: %v", err)
		}
	}

	if err = d.Set("web_filter_unknown_log", flattenWebfilterProfileWebFilterUnknownLog(o["web-filter-unknown-log"], d, "web_filter_unknown_log")); err != nil {
		if !fortiAPIPatch(o["web-filter-unknown-log"]) {
			return fmt.Errorf("Error reading web_filter_unknown_log: %v", err)
		}
	}

	if err = d.Set("web_filter_referer_log", flattenWebfilterProfileWebFilterRefererLog(o["web-filter-referer-log"], d, "web_filter_referer_log")); err != nil {
		if !fortiAPIPatch(o["web-filter-referer-log"]) {
			return fmt.Errorf("Error reading web_filter_referer_log: %v", err)
		}
	}

	if err = d.Set("web_filter_cookie_removal_log", flattenWebfilterProfileWebFilterCookieRemovalLog(o["web-filter-cookie-removal-log"], d, "web_filter_cookie_removal_log")); err != nil {
		if !fortiAPIPatch(o["web-filter-cookie-removal-log"]) {
			return fmt.Errorf("Error reading web_filter_cookie_removal_log: %v", err)
		}
	}

	if err = d.Set("web_url_log", flattenWebfilterProfileWebUrlLog(o["web-url-log"], d, "web_url_log")); err != nil {
		if !fortiAPIPatch(o["web-url-log"]) {
			return fmt.Errorf("Error reading web_url_log: %v", err)
		}
	}

	if err = d.Set("web_invalid_domain_log", flattenWebfilterProfileWebInvalidDomainLog(o["web-invalid-domain-log"], d, "web_invalid_domain_log")); err != nil {
		if !fortiAPIPatch(o["web-invalid-domain-log"]) {
			return fmt.Errorf("Error reading web_invalid_domain_log: %v", err)
		}
	}

	if err = d.Set("web_ftgd_err_log", flattenWebfilterProfileWebFtgdErrLog(o["web-ftgd-err-log"], d, "web_ftgd_err_log")); err != nil {
		if !fortiAPIPatch(o["web-ftgd-err-log"]) {
			return fmt.Errorf("Error reading web_ftgd_err_log: %v", err)
		}
	}

	if err = d.Set("web_ftgd_quota_usage", flattenWebfilterProfileWebFtgdQuotaUsage(o["web-ftgd-quota-usage"], d, "web_ftgd_quota_usage")); err != nil {
		if !fortiAPIPatch(o["web-ftgd-quota-usage"]) {
			return fmt.Errorf("Error reading web_ftgd_quota_usage: %v", err)
		}
	}

	if err = d.Set("extended_log", flattenWebfilterProfileExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if !fortiAPIPatch(o["extended-log"]) {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("web_extended_all_action_log", flattenWebfilterProfileWebExtendedAllActionLog(o["web-extended-all-action-log"], d, "web_extended_all_action_log")); err != nil {
		if !fortiAPIPatch(o["web-extended-all-action-log"]) {
			return fmt.Errorf("Error reading web_extended_all_action_log: %v", err)
		}
	}

	return nil
}

func flattenWebfilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileInspectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileHttpsReplacemsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOvrdPerm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfilePostAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ovrd_cookie"
	if _, ok := d.GetOk(pre_append); ok {
		result["ovrd-cookie"], _ = expandWebfilterProfileOverrideOvrdCookie(d, i["ovrd_cookie"], pre_append)
	}
	pre_append = pre + ".0." + "ovrd_scope"
	if _, ok := d.GetOk(pre_append); ok {
		result["ovrd-scope"], _ = expandWebfilterProfileOverrideOvrdScope(d, i["ovrd_scope"], pre_append)
	}
	pre_append = pre + ".0." + "profile_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["profile-type"], _ = expandWebfilterProfileOverrideProfileType(d, i["profile_type"], pre_append)
	}
	pre_append = pre + ".0." + "ovrd_dur_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["ovrd-dur-mode"], _ = expandWebfilterProfileOverrideOvrdDurMode(d, i["ovrd_dur_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ovrd_dur"
	if _, ok := d.GetOk(pre_append); ok {
		result["ovrd-dur"], _ = expandWebfilterProfileOverrideOvrdDur(d, i["ovrd_dur"], pre_append)
	}
	pre_append = pre + ".0." + "profile_attribute"
	if _, ok := d.GetOk(pre_append); ok {
		result["profile-attribute"], _ = expandWebfilterProfileOverrideProfileAttribute(d, i["profile_attribute"], pre_append)
	}
	pre_append = pre + ".0." + "ovrd_user_group"
	if _, ok := d.GetOk(pre_append); ok {
		result["ovrd-user-group"], _ = expandWebfilterProfileOverrideOvrdUserGroup(d, i["ovrd_user_group"], pre_append)
	} else {
		result["ovrd-user-group"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["profile"], _ = expandWebfilterProfileOverrideProfile(d, i["profile"], pre_append)
	} else {
		result["profile"] = make([]string, 0)
	}

	return result, nil
}

func expandWebfilterProfileOverrideOvrdCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideOvrdScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideProfileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideOvrdDurMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideOvrdDur(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideProfileAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideOvrdUserGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandWebfilterProfileOverrideOvrdUserGroupName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileOverrideOvrdUserGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileOverrideProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandWebfilterProfileOverrideProfileName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileOverrideProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWeb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bword_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["bword-threshold"], _ = expandWebfilterProfileWebBwordThreshold(d, i["bword_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "bword_table"
	if _, ok := d.GetOk(pre_append); ok {
		result["bword-table"], _ = expandWebfilterProfileWebBwordTable(d, i["bword_table"], pre_append)
	}
	pre_append = pre + ".0." + "urlfilter_table"
	if _, ok := d.GetOk(pre_append); ok {
		result["urlfilter-table"], _ = expandWebfilterProfileWebUrlfilterTable(d, i["urlfilter_table"], pre_append)
	}
	pre_append = pre + ".0." + "content_header_list"
	if _, ok := d.GetOk(pre_append); ok {
		result["content-header-list"], _ = expandWebfilterProfileWebContentHeaderList(d, i["content_header_list"], pre_append)
	}
	pre_append = pre + ".0." + "blacklist"
	if _, ok := d.GetOk(pre_append); ok {
		result["blacklist"], _ = expandWebfilterProfileWebBlacklist(d, i["blacklist"], pre_append)
	}
	pre_append = pre + ".0." + "whitelist"
	if _, ok := d.GetOk(pre_append); ok {
		result["whitelist"], _ = expandWebfilterProfileWebWhitelist(d, i["whitelist"], pre_append)
	}
	pre_append = pre + ".0." + "safe_search"
	if _, ok := d.GetOk(pre_append); ok {
		result["safe-search"], _ = expandWebfilterProfileWebSafeSearch(d, i["safe_search"], pre_append)
	}
	pre_append = pre + ".0." + "youtube_restrict"
	if _, ok := d.GetOk(pre_append); ok {
		result["youtube-restrict"], _ = expandWebfilterProfileWebYoutubeRestrict(d, i["youtube_restrict"], pre_append)
	}
	pre_append = pre + ".0." + "log_search"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-search"], _ = expandWebfilterProfileWebLogSearch(d, i["log_search"], pre_append)
	}
	pre_append = pre + ".0." + "keyword_match"
	if _, ok := d.GetOk(pre_append); ok {
		result["keyword-match"], _ = expandWebfilterProfileWebKeywordMatch(d, i["keyword_match"], pre_append)
	} else {
		result["keyword-match"] = make([]string, 0)
	}

	return result, nil
}

func expandWebfilterProfileWebBwordThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebBwordTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebUrlfilterTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebContentHeaderList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebBlacklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebWhitelist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebSafeSearch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebYoutubeRestrict(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebLogSearch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebKeywordMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pattern"], _ = expandWebfilterProfileWebKeywordMatchPattern(d, i["pattern"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileWebKeywordMatchPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileYoutubeChannelStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileYoutubeChannelFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWebfilterProfileYoutubeChannelFilterId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "channel_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["channel-id"], _ = expandWebfilterProfileYoutubeChannelFilterChannelId(d, i["channel_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comment"], _ = expandWebfilterProfileYoutubeChannelFilterComment(d, i["comment"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileYoutubeChannelFilterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileYoutubeChannelFilterChannelId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileYoutubeChannelFilterComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandWebfilterProfileFtgdWfOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "exempt_quota"
	if _, ok := d.GetOk(pre_append); ok {
		result["exempt-quota"], _ = expandWebfilterProfileFtgdWfExemptQuota(d, i["exempt_quota"], pre_append)
	}
	pre_append = pre + ".0." + "ovrd"
	if _, ok := d.GetOk(pre_append); ok {
		result["ovrd"], _ = expandWebfilterProfileFtgdWfOvrd(d, i["ovrd"], pre_append)
	}
	pre_append = pre + ".0." + "filters"
	if _, ok := d.GetOk(pre_append); ok {
		result["filters"], _ = expandWebfilterProfileFtgdWfFilters(d, i["filters"], pre_append)
	} else {
		result["filters"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "quota"
	if _, ok := d.GetOk(pre_append); ok {
		result["quota"], _ = expandWebfilterProfileFtgdWfQuota(d, i["quota"], pre_append)
	} else {
		result["quota"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "max_quota_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-quota-timeout"], _ = expandWebfilterProfileFtgdWfMaxQuotaTimeout(d, i["max_quota_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "rate_image_urls"
	if _, ok := d.GetOk(pre_append); ok {
		result["rate-image-urls"], _ = expandWebfilterProfileFtgdWfRateImageUrls(d, i["rate_image_urls"], pre_append)
	}
	pre_append = pre + ".0." + "rate_javascript_urls"
	if _, ok := d.GetOk(pre_append); ok {
		result["rate-javascript-urls"], _ = expandWebfilterProfileFtgdWfRateJavascriptUrls(d, i["rate_javascript_urls"], pre_append)
	}
	pre_append = pre + ".0." + "rate_css_urls"
	if _, ok := d.GetOk(pre_append); ok {
		result["rate-css-urls"], _ = expandWebfilterProfileFtgdWfRateCssUrls(d, i["rate_css_urls"], pre_append)
	}
	pre_append = pre + ".0." + "rate_crl_urls"
	if _, ok := d.GetOk(pre_append); ok {
		result["rate-crl-urls"], _ = expandWebfilterProfileFtgdWfRateCrlUrls(d, i["rate_crl_urls"], pre_append)
	}

	return result, nil
}

func expandWebfilterProfileFtgdWfOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfExemptQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfOvrd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFilters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWebfilterProfileFtgdWfFiltersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category"], _ = expandWebfilterProfileFtgdWfFiltersCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandWebfilterProfileFtgdWfFiltersAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warn_duration"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["warn-duration"], _ = expandWebfilterProfileFtgdWfFiltersWarnDuration(d, i["warn_duration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_usr_grp"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["auth-usr-grp"], _ = expandWebfilterProfileFtgdWfFiltersAuthUsrGrp(d, i["auth_usr_grp"], pre_append)
		} else {
			tmp["auth-usr-grp"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log"], _ = expandWebfilterProfileFtgdWfFiltersLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["override-replacemsg"], _ = expandWebfilterProfileFtgdWfFiltersOverrideReplacemsg(d, i["override_replacemsg"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_prompt"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["warning-prompt"], _ = expandWebfilterProfileFtgdWfFiltersWarningPrompt(d, i["warning_prompt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "warning_duration_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["warning-duration-type"], _ = expandWebfilterProfileFtgdWfFiltersWarningDurationType(d, i["warning_duration_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileFtgdWfFiltersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersWarnDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersAuthUsrGrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandWebfilterProfileFtgdWfFiltersAuthUsrGrpName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileFtgdWfFiltersAuthUsrGrpName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersOverrideReplacemsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersWarningPrompt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfFiltersWarningDurationType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWebfilterProfileFtgdWfQuotaId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category"], _ = expandWebfilterProfileFtgdWfQuotaCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandWebfilterProfileFtgdWfQuotaType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unit"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["unit"], _ = expandWebfilterProfileFtgdWfQuotaUnit(d, i["unit"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandWebfilterProfileFtgdWfQuotaValue(d, i["value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "duration"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["duration"], _ = expandWebfilterProfileFtgdWfQuotaDuration(d, i["duration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_replacemsg"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["override-replacemsg"], _ = expandWebfilterProfileFtgdWfQuotaOverrideReplacemsg(d, i["override_replacemsg"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileFtgdWfQuotaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaUnit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfQuotaOverrideReplacemsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfMaxQuotaTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRateImageUrls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRateJavascriptUrls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRateCssUrls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileFtgdWfRateCrlUrls(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWisp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWispServers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandWebfilterProfileWispServersName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterProfileWispServersName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWispAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileLogAllUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebContentLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterActivexLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterCommandBlockLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterCookieLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterAppletLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterJscriptLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterJsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterVbsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterUnknownLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterRefererLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFilterCookieRemovalLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebUrlLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebInvalidDomainLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFtgdErrLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebFtgdQuotaUsage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterProfileWebExtendedAllActionLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWebfilterProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWebfilterProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok {
		t, err := expandWebfilterProfileReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("inspection_mode"); ok {
		t, err := expandWebfilterProfileInspectionMode(d, v, "inspection_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inspection-mode"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok {
		t, err := expandWebfilterProfileOptions(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("https_replacemsg"); ok {
		t, err := expandWebfilterProfileHttpsReplacemsg(d, v, "https_replacemsg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-replacemsg"] = t
		}
	}

	if v, ok := d.GetOk("ovrd_perm"); ok {
		t, err := expandWebfilterProfileOvrdPerm(d, v, "ovrd_perm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ovrd-perm"] = t
		}
	}

	if v, ok := d.GetOk("post_action"); ok {
		t, err := expandWebfilterProfilePostAction(d, v, "post_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["post-action"] = t
		}
	}

	if v, ok := d.GetOk("override"); ok {
		t, err := expandWebfilterProfileOverride(d, v, "override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override"] = t
		}
	}

	if v, ok := d.GetOk("web"); ok {
		t, err := expandWebfilterProfileWeb(d, v, "web")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web"] = t
		}
	}

	if v, ok := d.GetOk("youtube_channel_status"); ok {
		t, err := expandWebfilterProfileYoutubeChannelStatus(d, v, "youtube_channel_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube-channel-status"] = t
		}
	}

	if v, ok := d.GetOk("youtube_channel_filter"); ok {
		t, err := expandWebfilterProfileYoutubeChannelFilter(d, v, "youtube_channel_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube-channel-filter"] = t
		}
	}

	if v, ok := d.GetOk("ftgd_wf"); ok {
		t, err := expandWebfilterProfileFtgdWf(d, v, "ftgd_wf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftgd-wf"] = t
		}
	}

	if v, ok := d.GetOk("wisp"); ok {
		t, err := expandWebfilterProfileWisp(d, v, "wisp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wisp"] = t
		}
	}

	if v, ok := d.GetOk("wisp_servers"); ok {
		t, err := expandWebfilterProfileWispServers(d, v, "wisp_servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wisp-servers"] = t
		}
	}

	if v, ok := d.GetOk("wisp_algorithm"); ok {
		t, err := expandWebfilterProfileWispAlgorithm(d, v, "wisp_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wisp-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("log_all_url"); ok {
		t, err := expandWebfilterProfileLogAllUrl(d, v, "log_all_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-all-url"] = t
		}
	}

	if v, ok := d.GetOk("web_content_log"); ok {
		t, err := expandWebfilterProfileWebContentLog(d, v, "web_content_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-content-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_activex_log"); ok {
		t, err := expandWebfilterProfileWebFilterActivexLog(d, v, "web_filter_activex_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-activex-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_command_block_log"); ok {
		t, err := expandWebfilterProfileWebFilterCommandBlockLog(d, v, "web_filter_command_block_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-command-block-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_cookie_log"); ok {
		t, err := expandWebfilterProfileWebFilterCookieLog(d, v, "web_filter_cookie_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-cookie-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_applet_log"); ok {
		t, err := expandWebfilterProfileWebFilterAppletLog(d, v, "web_filter_applet_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-applet-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_jscript_log"); ok {
		t, err := expandWebfilterProfileWebFilterJscriptLog(d, v, "web_filter_jscript_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-jscript-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_js_log"); ok {
		t, err := expandWebfilterProfileWebFilterJsLog(d, v, "web_filter_js_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-js-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_vbs_log"); ok {
		t, err := expandWebfilterProfileWebFilterVbsLog(d, v, "web_filter_vbs_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-vbs-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_unknown_log"); ok {
		t, err := expandWebfilterProfileWebFilterUnknownLog(d, v, "web_filter_unknown_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-unknown-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_referer_log"); ok {
		t, err := expandWebfilterProfileWebFilterRefererLog(d, v, "web_filter_referer_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-referer-log"] = t
		}
	}

	if v, ok := d.GetOk("web_filter_cookie_removal_log"); ok {
		t, err := expandWebfilterProfileWebFilterCookieRemovalLog(d, v, "web_filter_cookie_removal_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-filter-cookie-removal-log"] = t
		}
	}

	if v, ok := d.GetOk("web_url_log"); ok {
		t, err := expandWebfilterProfileWebUrlLog(d, v, "web_url_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-url-log"] = t
		}
	}

	if v, ok := d.GetOk("web_invalid_domain_log"); ok {
		t, err := expandWebfilterProfileWebInvalidDomainLog(d, v, "web_invalid_domain_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-invalid-domain-log"] = t
		}
	}

	if v, ok := d.GetOk("web_ftgd_err_log"); ok {
		t, err := expandWebfilterProfileWebFtgdErrLog(d, v, "web_ftgd_err_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-ftgd-err-log"] = t
		}
	}

	if v, ok := d.GetOk("web_ftgd_quota_usage"); ok {
		t, err := expandWebfilterProfileWebFtgdQuotaUsage(d, v, "web_ftgd_quota_usage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-ftgd-quota-usage"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok {
		t, err := expandWebfilterProfileExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("web_extended_all_action_log"); ok {
		t, err := expandWebfilterProfileWebExtendedAllActionLog(d, v, "web_extended_all_action_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-extended-all-action-log"] = t
		}
	}

	return &obj, nil
}
