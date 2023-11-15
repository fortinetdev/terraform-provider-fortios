// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Email Filter profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEmailfilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailfilterProfileCreate,
		Read:   resourceEmailfilterProfileRead,
		Update: resourceEmailfilterProfileUpdate,
		Delete: resourceEmailfilterProfileDelete,

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
			"spam_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spam_log_fortiguard_response": &schema.Schema{
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
			"spam_filtering": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"imap": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag_msg": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"pop3": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag_msg": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"smtp": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag_msg": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"hdrip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"local_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"mapi": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
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
			"msn_hotmail": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_all": &schema.Schema{
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
			"yahoo_mail": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_all": &schema.Schema{
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
			"gmail": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_all": &schema.Schema{
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
			"other_webmails": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"spam_bword_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"spam_bword_table": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"spam_bal_table": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"spam_bwl_table": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"spam_mheader_table": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"spam_rbl_table": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"spam_iptrust_table": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceEmailfilterProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectEmailfilterProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating EmailfilterProfile resource while getting object: %v", err)
	}

	o, err := c.CreateEmailfilterProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating EmailfilterProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("EmailfilterProfile")
	}

	return resourceEmailfilterProfileRead(d, m)
}

func resourceEmailfilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectEmailfilterProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateEmailfilterProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("EmailfilterProfile")
	}

	return resourceEmailfilterProfileRead(d, m)
}

func resourceEmailfilterProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteEmailfilterProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting EmailfilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadEmailfilterProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEmailfilterProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenEmailfilterProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileFeatureSet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileSpamLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileSpamLogFortiguardResponse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileFileFilter(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenEmailfilterProfileFileFilterStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenEmailfilterProfileFileFilterLog(i["log"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "scan_archive_contents"
	if _, ok := i["scan-archive-contents"]; ok {
		result["scan_archive_contents"] = flattenEmailfilterProfileFileFilterScanArchiveContents(i["scan-archive-contents"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "entries"
	if _, ok := i["entries"]; ok {
		result["entries"] = flattenEmailfilterProfileFileFilterEntries(i["entries"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenEmailfilterProfileFileFilterStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileFileFilterLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileFileFilterScanArchiveContents(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileFileFilterEntries(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["filter"]; ok {
			tmp["filter"] = flattenEmailfilterProfileFileFilterEntriesFilter(i["filter"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			tmp["comment"] = flattenEmailfilterProfileFileFilterEntriesComment(i["comment"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			tmp["protocol"] = flattenEmailfilterProfileFileFilterEntriesProtocol(i["protocol"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = flattenEmailfilterProfileFileFilterEntriesAction(i["action"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_protected"
		if _, ok := i["password-protected"]; ok {
			tmp["password_protected"] = flattenEmailfilterProfileFileFilterEntriesPasswordProtected(i["password-protected"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := i["file-type"]; ok {
			tmp["file_type"] = flattenEmailfilterProfileFileFilterEntriesFileType(i["file-type"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "filter", d)
	return result
}

func flattenEmailfilterProfileFileFilterEntriesFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileFileFilterEntriesComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileFileFilterEntriesProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileFileFilterEntriesAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileFileFilterEntriesPasswordProtected(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileFileFilterEntriesFileType(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenEmailfilterProfileFileFilterEntriesFileTypeName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenEmailfilterProfileFileFilterEntriesFileTypeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileSpamFiltering(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileExternal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileImap(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenEmailfilterProfileImapLogAll(i["log-all"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenEmailfilterProfileImapLog(i["log"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenEmailfilterProfileImapAction(i["action"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tag_type"
	if _, ok := i["tag-type"]; ok {
		result["tag_type"] = flattenEmailfilterProfileImapTagType(i["tag-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tag_msg"
	if _, ok := i["tag-msg"]; ok {
		result["tag_msg"] = flattenEmailfilterProfileImapTagMsg(i["tag-msg"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenEmailfilterProfileImapLogAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileImapLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileImapAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileImapTagType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileImapTagMsg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfilePop3(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenEmailfilterProfilePop3LogAll(i["log-all"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenEmailfilterProfilePop3Log(i["log"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenEmailfilterProfilePop3Action(i["action"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tag_type"
	if _, ok := i["tag-type"]; ok {
		result["tag_type"] = flattenEmailfilterProfilePop3TagType(i["tag-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tag_msg"
	if _, ok := i["tag-msg"]; ok {
		result["tag_msg"] = flattenEmailfilterProfilePop3TagMsg(i["tag-msg"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenEmailfilterProfilePop3LogAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfilePop3Log(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfilePop3Action(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfilePop3TagType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfilePop3TagMsg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileSmtp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenEmailfilterProfileSmtpLogAll(i["log-all"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenEmailfilterProfileSmtpLog(i["log"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenEmailfilterProfileSmtpAction(i["action"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tag_type"
	if _, ok := i["tag-type"]; ok {
		result["tag_type"] = flattenEmailfilterProfileSmtpTagType(i["tag-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tag_msg"
	if _, ok := i["tag-msg"]; ok {
		result["tag_msg"] = flattenEmailfilterProfileSmtpTagMsg(i["tag-msg"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "hdrip"
	if _, ok := i["hdrip"]; ok {
		result["hdrip"] = flattenEmailfilterProfileSmtpHdrip(i["hdrip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "local_override"
	if _, ok := i["local-override"]; ok {
		result["local_override"] = flattenEmailfilterProfileSmtpLocalOverride(i["local-override"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenEmailfilterProfileSmtpLogAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileSmtpLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileSmtpAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileSmtpTagType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileSmtpTagMsg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileSmtpHdrip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileSmtpLocalOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileMapi(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenEmailfilterProfileMapiLogAll(i["log-all"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenEmailfilterProfileMapiLog(i["log"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenEmailfilterProfileMapiAction(i["action"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenEmailfilterProfileMapiLogAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileMapiLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileMapiAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileMsnHotmail(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenEmailfilterProfileMsnHotmailLogAll(i["log-all"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenEmailfilterProfileMsnHotmailLog(i["log"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenEmailfilterProfileMsnHotmailLogAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileMsnHotmailLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileYahooMail(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenEmailfilterProfileYahooMailLogAll(i["log-all"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenEmailfilterProfileYahooMailLog(i["log"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenEmailfilterProfileYahooMailLogAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileYahooMailLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileGmail(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenEmailfilterProfileGmailLogAll(i["log-all"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenEmailfilterProfileGmailLog(i["log"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenEmailfilterProfileGmailLogAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileGmailLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileOtherWebmails(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := i["log-all"]; ok {
		result["log_all"] = flattenEmailfilterProfileOtherWebmailsLogAll(i["log-all"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenEmailfilterProfileOtherWebmailsLogAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileSpamBwordThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileSpamBwordTable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileSpamBalTable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileSpamBwlTable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileSpamMheaderTable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileSpamRblTable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEmailfilterProfileSpamIptrustTable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectEmailfilterProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenEmailfilterProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenEmailfilterProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("feature_set", flattenEmailfilterProfileFeatureSet(o["feature-set"], d, "feature_set", sv)); err != nil {
		if !fortiAPIPatch(o["feature-set"]) {
			return fmt.Errorf("Error reading feature_set: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenEmailfilterProfileReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group", sv)); err != nil {
		if !fortiAPIPatch(o["replacemsg-group"]) {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("spam_log", flattenEmailfilterProfileSpamLog(o["spam-log"], d, "spam_log", sv)); err != nil {
		if !fortiAPIPatch(o["spam-log"]) {
			return fmt.Errorf("Error reading spam_log: %v", err)
		}
	}

	if err = d.Set("spam_log_fortiguard_response", flattenEmailfilterProfileSpamLogFortiguardResponse(o["spam-log-fortiguard-response"], d, "spam_log_fortiguard_response", sv)); err != nil {
		if !fortiAPIPatch(o["spam-log-fortiguard-response"]) {
			return fmt.Errorf("Error reading spam_log_fortiguard_response: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("file_filter", flattenEmailfilterProfileFileFilter(o["file-filter"], d, "file_filter", sv)); err != nil {
			if !fortiAPIPatch(o["file-filter"]) {
				return fmt.Errorf("Error reading file_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("file_filter"); ok {
			if err = d.Set("file_filter", flattenEmailfilterProfileFileFilter(o["file-filter"], d, "file_filter", sv)); err != nil {
				if !fortiAPIPatch(o["file-filter"]) {
					return fmt.Errorf("Error reading file_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("spam_filtering", flattenEmailfilterProfileSpamFiltering(o["spam-filtering"], d, "spam_filtering", sv)); err != nil {
		if !fortiAPIPatch(o["spam-filtering"]) {
			return fmt.Errorf("Error reading spam_filtering: %v", err)
		}
	}

	if err = d.Set("external", flattenEmailfilterProfileExternal(o["external"], d, "external", sv)); err != nil {
		if !fortiAPIPatch(o["external"]) {
			return fmt.Errorf("Error reading external: %v", err)
		}
	}

	if err = d.Set("options", flattenEmailfilterProfileOptions(o["options"], d, "options", sv)); err != nil {
		if !fortiAPIPatch(o["options"]) {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("imap", flattenEmailfilterProfileImap(o["imap"], d, "imap", sv)); err != nil {
			if !fortiAPIPatch(o["imap"]) {
				return fmt.Errorf("Error reading imap: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("imap"); ok {
			if err = d.Set("imap", flattenEmailfilterProfileImap(o["imap"], d, "imap", sv)); err != nil {
				if !fortiAPIPatch(o["imap"]) {
					return fmt.Errorf("Error reading imap: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("pop3", flattenEmailfilterProfilePop3(o["pop3"], d, "pop3", sv)); err != nil {
			if !fortiAPIPatch(o["pop3"]) {
				return fmt.Errorf("Error reading pop3: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pop3"); ok {
			if err = d.Set("pop3", flattenEmailfilterProfilePop3(o["pop3"], d, "pop3", sv)); err != nil {
				if !fortiAPIPatch(o["pop3"]) {
					return fmt.Errorf("Error reading pop3: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("smtp", flattenEmailfilterProfileSmtp(o["smtp"], d, "smtp", sv)); err != nil {
			if !fortiAPIPatch(o["smtp"]) {
				return fmt.Errorf("Error reading smtp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("smtp"); ok {
			if err = d.Set("smtp", flattenEmailfilterProfileSmtp(o["smtp"], d, "smtp", sv)); err != nil {
				if !fortiAPIPatch(o["smtp"]) {
					return fmt.Errorf("Error reading smtp: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("mapi", flattenEmailfilterProfileMapi(o["mapi"], d, "mapi", sv)); err != nil {
			if !fortiAPIPatch(o["mapi"]) {
				return fmt.Errorf("Error reading mapi: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mapi"); ok {
			if err = d.Set("mapi", flattenEmailfilterProfileMapi(o["mapi"], d, "mapi", sv)); err != nil {
				if !fortiAPIPatch(o["mapi"]) {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("msn_hotmail", flattenEmailfilterProfileMsnHotmail(o["msn-hotmail"], d, "msn_hotmail", sv)); err != nil {
			if !fortiAPIPatch(o["msn-hotmail"]) {
				return fmt.Errorf("Error reading msn_hotmail: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("msn_hotmail"); ok {
			if err = d.Set("msn_hotmail", flattenEmailfilterProfileMsnHotmail(o["msn-hotmail"], d, "msn_hotmail", sv)); err != nil {
				if !fortiAPIPatch(o["msn-hotmail"]) {
					return fmt.Errorf("Error reading msn_hotmail: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("yahoo_mail", flattenEmailfilterProfileYahooMail(o["yahoo-mail"], d, "yahoo_mail", sv)); err != nil {
			if !fortiAPIPatch(o["yahoo-mail"]) {
				return fmt.Errorf("Error reading yahoo_mail: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("yahoo_mail"); ok {
			if err = d.Set("yahoo_mail", flattenEmailfilterProfileYahooMail(o["yahoo-mail"], d, "yahoo_mail", sv)); err != nil {
				if !fortiAPIPatch(o["yahoo-mail"]) {
					return fmt.Errorf("Error reading yahoo_mail: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("gmail", flattenEmailfilterProfileGmail(o["gmail"], d, "gmail", sv)); err != nil {
			if !fortiAPIPatch(o["gmail"]) {
				return fmt.Errorf("Error reading gmail: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("gmail"); ok {
			if err = d.Set("gmail", flattenEmailfilterProfileGmail(o["gmail"], d, "gmail", sv)); err != nil {
				if !fortiAPIPatch(o["gmail"]) {
					return fmt.Errorf("Error reading gmail: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("other_webmails", flattenEmailfilterProfileOtherWebmails(o["other-webmails"], d, "other_webmails", sv)); err != nil {
			if !fortiAPIPatch(o["other-webmails"]) {
				return fmt.Errorf("Error reading other_webmails: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("other_webmails"); ok {
			if err = d.Set("other_webmails", flattenEmailfilterProfileOtherWebmails(o["other-webmails"], d, "other_webmails", sv)); err != nil {
				if !fortiAPIPatch(o["other-webmails"]) {
					return fmt.Errorf("Error reading other_webmails: %v", err)
				}
			}
		}
	}

	if err = d.Set("spam_bword_threshold", flattenEmailfilterProfileSpamBwordThreshold(o["spam-bword-threshold"], d, "spam_bword_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["spam-bword-threshold"]) {
			return fmt.Errorf("Error reading spam_bword_threshold: %v", err)
		}
	}

	if err = d.Set("spam_bword_table", flattenEmailfilterProfileSpamBwordTable(o["spam-bword-table"], d, "spam_bword_table", sv)); err != nil {
		if !fortiAPIPatch(o["spam-bword-table"]) {
			return fmt.Errorf("Error reading spam_bword_table: %v", err)
		}
	}

	if err = d.Set("spam_bal_table", flattenEmailfilterProfileSpamBalTable(o["spam-bal-table"], d, "spam_bal_table", sv)); err != nil {
		if !fortiAPIPatch(o["spam-bal-table"]) {
			return fmt.Errorf("Error reading spam_bal_table: %v", err)
		}
	}

	if err = d.Set("spam_bwl_table", flattenEmailfilterProfileSpamBwlTable(o["spam-bwl-table"], d, "spam_bwl_table", sv)); err != nil {
		if !fortiAPIPatch(o["spam-bwl-table"]) {
			return fmt.Errorf("Error reading spam_bwl_table: %v", err)
		}
	}

	if err = d.Set("spam_mheader_table", flattenEmailfilterProfileSpamMheaderTable(o["spam-mheader-table"], d, "spam_mheader_table", sv)); err != nil {
		if !fortiAPIPatch(o["spam-mheader-table"]) {
			return fmt.Errorf("Error reading spam_mheader_table: %v", err)
		}
	}

	if err = d.Set("spam_rbl_table", flattenEmailfilterProfileSpamRblTable(o["spam-rbl-table"], d, "spam_rbl_table", sv)); err != nil {
		if !fortiAPIPatch(o["spam-rbl-table"]) {
			return fmt.Errorf("Error reading spam_rbl_table: %v", err)
		}
	}

	if err = d.Set("spam_iptrust_table", flattenEmailfilterProfileSpamIptrustTable(o["spam-iptrust-table"], d, "spam_iptrust_table", sv)); err != nil {
		if !fortiAPIPatch(o["spam-iptrust-table"]) {
			return fmt.Errorf("Error reading spam_iptrust_table: %v", err)
		}
	}

	return nil
}

func flattenEmailfilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandEmailfilterProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileFeatureSet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSpamLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSpamLogFortiguardResponse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileFileFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandEmailfilterProfileFileFilterStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandEmailfilterProfileFileFilterLog(d, i["log"], pre_append, sv)
	}
	pre_append = pre + ".0." + "scan_archive_contents"
	if _, ok := d.GetOk(pre_append); ok {
		result["scan-archive-contents"], _ = expandEmailfilterProfileFileFilterScanArchiveContents(d, i["scan_archive_contents"], pre_append, sv)
	}
	pre_append = pre + ".0." + "entries"
	if _, ok := d.GetOk(pre_append); ok {
		result["entries"], _ = expandEmailfilterProfileFileFilterEntries(d, i["entries"], pre_append, sv)
	} else {
		result["entries"] = make([]string, 0)
	}

	return result, nil
}

func expandEmailfilterProfileFileFilterStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileFileFilterLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileFileFilterScanArchiveContents(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileFileFilterEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["filter"], _ = expandEmailfilterProfileFileFilterEntriesFilter(d, i["filter"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comment"], _ = expandEmailfilterProfileFileFilterEntriesComment(d, i["comment"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandEmailfilterProfileFileFilterEntriesProtocol(d, i["protocol"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandEmailfilterProfileFileFilterEntriesAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_protected"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["password-protected"], _ = expandEmailfilterProfileFileFilterEntriesPasswordProtected(d, i["password_protected"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["file-type"], _ = expandEmailfilterProfileFileFilterEntriesFileType(d, i["file_type"], pre_append, sv)
		} else {
			tmp["file-type"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandEmailfilterProfileFileFilterEntriesFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileFileFilterEntriesComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileFileFilterEntriesProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileFileFilterEntriesAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileFileFilterEntriesPasswordProtected(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileFileFilterEntriesFileType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandEmailfilterProfileFileFilterEntriesFileTypeName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandEmailfilterProfileFileFilterEntriesFileTypeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSpamFiltering(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileExternal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileImap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-all"], _ = expandEmailfilterProfileImapLogAll(d, i["log_all"], pre_append, sv)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandEmailfilterProfileImapLog(d, i["log"], pre_append, sv)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandEmailfilterProfileImapAction(d, i["action"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tag_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag-type"], _ = expandEmailfilterProfileImapTagType(d, i["tag_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tag_msg"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag-msg"], _ = expandEmailfilterProfileImapTagMsg(d, i["tag_msg"], pre_append, sv)
	}

	return result, nil
}

func expandEmailfilterProfileImapLogAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileImapLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileImapAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileImapTagType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileImapTagMsg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfilePop3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-all"], _ = expandEmailfilterProfilePop3LogAll(d, i["log_all"], pre_append, sv)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandEmailfilterProfilePop3Log(d, i["log"], pre_append, sv)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandEmailfilterProfilePop3Action(d, i["action"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tag_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag-type"], _ = expandEmailfilterProfilePop3TagType(d, i["tag_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tag_msg"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag-msg"], _ = expandEmailfilterProfilePop3TagMsg(d, i["tag_msg"], pre_append, sv)
	}

	return result, nil
}

func expandEmailfilterProfilePop3LogAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfilePop3Log(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfilePop3Action(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfilePop3TagType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfilePop3TagMsg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSmtp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-all"], _ = expandEmailfilterProfileSmtpLogAll(d, i["log_all"], pre_append, sv)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandEmailfilterProfileSmtpLog(d, i["log"], pre_append, sv)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandEmailfilterProfileSmtpAction(d, i["action"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tag_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag-type"], _ = expandEmailfilterProfileSmtpTagType(d, i["tag_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tag_msg"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag-msg"], _ = expandEmailfilterProfileSmtpTagMsg(d, i["tag_msg"], pre_append, sv)
	}
	pre_append = pre + ".0." + "hdrip"
	if _, ok := d.GetOk(pre_append); ok {
		result["hdrip"], _ = expandEmailfilterProfileSmtpHdrip(d, i["hdrip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "local_override"
	if _, ok := d.GetOk(pre_append); ok {
		result["local-override"], _ = expandEmailfilterProfileSmtpLocalOverride(d, i["local_override"], pre_append, sv)
	}

	return result, nil
}

func expandEmailfilterProfileSmtpLogAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSmtpLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSmtpAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSmtpTagType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSmtpTagMsg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSmtpHdrip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSmtpLocalOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileMapi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-all"], _ = expandEmailfilterProfileMapiLogAll(d, i["log_all"], pre_append, sv)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandEmailfilterProfileMapiLog(d, i["log"], pre_append, sv)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandEmailfilterProfileMapiAction(d, i["action"], pre_append, sv)
	}

	return result, nil
}

func expandEmailfilterProfileMapiLogAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileMapiLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileMapiAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileMsnHotmail(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-all"], _ = expandEmailfilterProfileMsnHotmailLogAll(d, i["log_all"], pre_append, sv)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandEmailfilterProfileMsnHotmailLog(d, i["log"], pre_append, sv)
	}

	return result, nil
}

func expandEmailfilterProfileMsnHotmailLogAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileMsnHotmailLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileYahooMail(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-all"], _ = expandEmailfilterProfileYahooMailLogAll(d, i["log_all"], pre_append, sv)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandEmailfilterProfileYahooMailLog(d, i["log"], pre_append, sv)
	}

	return result, nil
}

func expandEmailfilterProfileYahooMailLogAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileYahooMailLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileGmail(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-all"], _ = expandEmailfilterProfileGmailLogAll(d, i["log_all"], pre_append, sv)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandEmailfilterProfileGmailLog(d, i["log"], pre_append, sv)
	}

	return result, nil
}

func expandEmailfilterProfileGmailLogAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileGmailLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileOtherWebmails(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-all"], _ = expandEmailfilterProfileOtherWebmailsLogAll(d, i["log_all"], pre_append, sv)
	}

	return result, nil
}

func expandEmailfilterProfileOtherWebmailsLogAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSpamBwordThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSpamBwordTable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSpamBalTable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSpamBwlTable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSpamMheaderTable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSpamRblTable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEmailfilterProfileSpamIptrustTable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectEmailfilterProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandEmailfilterProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandEmailfilterProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("feature_set"); ok {
		t, err := expandEmailfilterProfileFeatureSet(d, v, "feature_set", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["feature-set"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok {
		t, err := expandEmailfilterProfileReplacemsgGroup(d, v, "replacemsg_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("spam_log"); ok {
		t, err := expandEmailfilterProfileSpamLog(d, v, "spam_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-log"] = t
		}
	}

	if v, ok := d.GetOk("spam_log_fortiguard_response"); ok {
		t, err := expandEmailfilterProfileSpamLogFortiguardResponse(d, v, "spam_log_fortiguard_response", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-log-fortiguard-response"] = t
		}
	}

	if v, ok := d.GetOk("file_filter"); ok {
		t, err := expandEmailfilterProfileFileFilter(d, v, "file_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-filter"] = t
		}
	}

	if v, ok := d.GetOk("spam_filtering"); ok {
		t, err := expandEmailfilterProfileSpamFiltering(d, v, "spam_filtering", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-filtering"] = t
		}
	}

	if v, ok := d.GetOk("external"); ok {
		t, err := expandEmailfilterProfileExternal(d, v, "external", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok {
		t, err := expandEmailfilterProfileOptions(d, v, "options", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("imap"); ok {
		t, err := expandEmailfilterProfileImap(d, v, "imap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imap"] = t
		}
	}

	if v, ok := d.GetOk("pop3"); ok {
		t, err := expandEmailfilterProfilePop3(d, v, "pop3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3"] = t
		}
	}

	if v, ok := d.GetOk("smtp"); ok {
		t, err := expandEmailfilterProfileSmtp(d, v, "smtp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtp"] = t
		}
	}

	if v, ok := d.GetOk("mapi"); ok {
		t, err := expandEmailfilterProfileMapi(d, v, "mapi", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapi"] = t
		}
	}

	if v, ok := d.GetOk("msn_hotmail"); ok {
		t, err := expandEmailfilterProfileMsnHotmail(d, v, "msn_hotmail", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msn-hotmail"] = t
		}
	}

	if v, ok := d.GetOk("yahoo_mail"); ok {
		t, err := expandEmailfilterProfileYahooMail(d, v, "yahoo_mail", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["yahoo-mail"] = t
		}
	}

	if v, ok := d.GetOk("gmail"); ok {
		t, err := expandEmailfilterProfileGmail(d, v, "gmail", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gmail"] = t
		}
	}

	if v, ok := d.GetOk("other_webmails"); ok {
		t, err := expandEmailfilterProfileOtherWebmails(d, v, "other_webmails", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["other-webmails"] = t
		}
	}

	if v, ok := d.GetOkExists("spam_bword_threshold"); ok {
		t, err := expandEmailfilterProfileSpamBwordThreshold(d, v, "spam_bword_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-bword-threshold"] = t
		}
	}

	if v, ok := d.GetOkExists("spam_bword_table"); ok {
		t, err := expandEmailfilterProfileSpamBwordTable(d, v, "spam_bword_table", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-bword-table"] = t
		}
	}

	if v, ok := d.GetOkExists("spam_bal_table"); ok {
		t, err := expandEmailfilterProfileSpamBalTable(d, v, "spam_bal_table", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-bal-table"] = t
		}
	}

	if v, ok := d.GetOkExists("spam_bwl_table"); ok {
		t, err := expandEmailfilterProfileSpamBwlTable(d, v, "spam_bwl_table", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-bwl-table"] = t
		}
	}

	if v, ok := d.GetOkExists("spam_mheader_table"); ok {
		t, err := expandEmailfilterProfileSpamMheaderTable(d, v, "spam_mheader_table", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-mheader-table"] = t
		}
	}

	if v, ok := d.GetOkExists("spam_rbl_table"); ok {
		t, err := expandEmailfilterProfileSpamRblTable(d, v, "spam_rbl_table", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-rbl-table"] = t
		}
	}

	if v, ok := d.GetOkExists("spam_iptrust_table"); ok {
		t, err := expandEmailfilterProfileSpamIptrustTable(d, v, "spam_iptrust_table", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-iptrust-table"] = t
		}
	}

	return &obj, nil
}
