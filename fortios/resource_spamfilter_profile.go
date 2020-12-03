// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure AntiSpam profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSpamfilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpamfilterProfileCreate,
		Read:   resourceSpamfilterProfileRead,
		Update: resourceSpamfilterProfileUpdate,
		Delete: resourceSpamfilterProfileDelete,

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
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"flow_based": &schema.Schema{
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
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log": &schema.Schema{
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
		},
	}
}

func resourceSpamfilterProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSpamfilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating SpamfilterProfile resource while getting object: %v", err)
	}

	o, err := c.CreateSpamfilterProfile(obj)

	if err != nil {
		return fmt.Errorf("Error creating SpamfilterProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SpamfilterProfile")
	}

	return resourceSpamfilterProfileRead(d, m)
}

func resourceSpamfilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSpamfilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating SpamfilterProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateSpamfilterProfile(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SpamfilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SpamfilterProfile")
	}

	return resourceSpamfilterProfileRead(d, m)
}

func resourceSpamfilterProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSpamfilterProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SpamfilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSpamfilterProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSpamfilterProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SpamfilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSpamfilterProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SpamfilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenSpamfilterProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileFlowBased(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileSpamLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileSpamLogFortiguardResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileSpamFiltering(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileImap(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenSpamfilterProfileImapLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenSpamfilterProfileImapAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_type"
	if _, ok := i["tag-type"]; ok {
		result["tag_type"] = flattenSpamfilterProfileImapTagType(i["tag-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_msg"
	if _, ok := i["tag-msg"]; ok {
		result["tag_msg"] = flattenSpamfilterProfileImapTagMsg(i["tag-msg"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSpamfilterProfileImapLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileImapAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileImapTagType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileImapTagMsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfilePop3(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenSpamfilterProfilePop3Log(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenSpamfilterProfilePop3Action(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_type"
	if _, ok := i["tag-type"]; ok {
		result["tag_type"] = flattenSpamfilterProfilePop3TagType(i["tag-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_msg"
	if _, ok := i["tag-msg"]; ok {
		result["tag_msg"] = flattenSpamfilterProfilePop3TagMsg(i["tag-msg"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSpamfilterProfilePop3Log(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfilePop3Action(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfilePop3TagType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfilePop3TagMsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileSmtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenSpamfilterProfileSmtpLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenSpamfilterProfileSmtpAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_type"
	if _, ok := i["tag-type"]; ok {
		result["tag_type"] = flattenSpamfilterProfileSmtpTagType(i["tag-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag_msg"
	if _, ok := i["tag-msg"]; ok {
		result["tag_msg"] = flattenSpamfilterProfileSmtpTagMsg(i["tag-msg"], d, pre_append)
	}

	pre_append = pre + ".0." + "hdrip"
	if _, ok := i["hdrip"]; ok {
		result["hdrip"] = flattenSpamfilterProfileSmtpHdrip(i["hdrip"], d, pre_append)
	}

	pre_append = pre + ".0." + "local_override"
	if _, ok := i["local-override"]; ok {
		result["local_override"] = flattenSpamfilterProfileSmtpLocalOverride(i["local-override"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSpamfilterProfileSmtpLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileSmtpAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileSmtpTagType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileSmtpTagMsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileSmtpHdrip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileSmtpLocalOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileMapi(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenSpamfilterProfileMapiLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenSpamfilterProfileMapiAction(i["action"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSpamfilterProfileMapiLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileMapiAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileMsnHotmail(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenSpamfilterProfileMsnHotmailLog(i["log"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSpamfilterProfileMsnHotmailLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileYahooMail(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenSpamfilterProfileYahooMailLog(i["log"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSpamfilterProfileYahooMailLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileGmail(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenSpamfilterProfileGmailLog(i["log"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSpamfilterProfileGmailLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileSpamBwordThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileSpamBwordTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileSpamBwlTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileSpamMheaderTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileSpamRblTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSpamfilterProfileSpamIptrustTable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSpamfilterProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSpamfilterProfileName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenSpamfilterProfileComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("flow_based", flattenSpamfilterProfileFlowBased(o["flow-based"], d, "flow_based")); err != nil {
		if !fortiAPIPatch(o["flow-based"]) {
			return fmt.Errorf("Error reading flow_based: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenSpamfilterProfileReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if !fortiAPIPatch(o["replacemsg-group"]) {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("spam_log", flattenSpamfilterProfileSpamLog(o["spam-log"], d, "spam_log")); err != nil {
		if !fortiAPIPatch(o["spam-log"]) {
			return fmt.Errorf("Error reading spam_log: %v", err)
		}
	}

	if err = d.Set("spam_log_fortiguard_response", flattenSpamfilterProfileSpamLogFortiguardResponse(o["spam-log-fortiguard-response"], d, "spam_log_fortiguard_response")); err != nil {
		if !fortiAPIPatch(o["spam-log-fortiguard-response"]) {
			return fmt.Errorf("Error reading spam_log_fortiguard_response: %v", err)
		}
	}

	if err = d.Set("spam_filtering", flattenSpamfilterProfileSpamFiltering(o["spam-filtering"], d, "spam_filtering")); err != nil {
		if !fortiAPIPatch(o["spam-filtering"]) {
			return fmt.Errorf("Error reading spam_filtering: %v", err)
		}
	}

	if err = d.Set("external", flattenSpamfilterProfileExternal(o["external"], d, "external")); err != nil {
		if !fortiAPIPatch(o["external"]) {
			return fmt.Errorf("Error reading external: %v", err)
		}
	}

	if err = d.Set("options", flattenSpamfilterProfileOptions(o["options"], d, "options")); err != nil {
		if !fortiAPIPatch(o["options"]) {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("imap", flattenSpamfilterProfileImap(o["imap"], d, "imap")); err != nil {
			if !fortiAPIPatch(o["imap"]) {
				return fmt.Errorf("Error reading imap: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("imap"); ok {
			if err = d.Set("imap", flattenSpamfilterProfileImap(o["imap"], d, "imap")); err != nil {
				if !fortiAPIPatch(o["imap"]) {
					return fmt.Errorf("Error reading imap: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("pop3", flattenSpamfilterProfilePop3(o["pop3"], d, "pop3")); err != nil {
			if !fortiAPIPatch(o["pop3"]) {
				return fmt.Errorf("Error reading pop3: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pop3"); ok {
			if err = d.Set("pop3", flattenSpamfilterProfilePop3(o["pop3"], d, "pop3")); err != nil {
				if !fortiAPIPatch(o["pop3"]) {
					return fmt.Errorf("Error reading pop3: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("smtp", flattenSpamfilterProfileSmtp(o["smtp"], d, "smtp")); err != nil {
			if !fortiAPIPatch(o["smtp"]) {
				return fmt.Errorf("Error reading smtp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("smtp"); ok {
			if err = d.Set("smtp", flattenSpamfilterProfileSmtp(o["smtp"], d, "smtp")); err != nil {
				if !fortiAPIPatch(o["smtp"]) {
					return fmt.Errorf("Error reading smtp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mapi", flattenSpamfilterProfileMapi(o["mapi"], d, "mapi")); err != nil {
			if !fortiAPIPatch(o["mapi"]) {
				return fmt.Errorf("Error reading mapi: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mapi"); ok {
			if err = d.Set("mapi", flattenSpamfilterProfileMapi(o["mapi"], d, "mapi")); err != nil {
				if !fortiAPIPatch(o["mapi"]) {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("msn_hotmail", flattenSpamfilterProfileMsnHotmail(o["msn-hotmail"], d, "msn_hotmail")); err != nil {
			if !fortiAPIPatch(o["msn-hotmail"]) {
				return fmt.Errorf("Error reading msn_hotmail: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("msn_hotmail"); ok {
			if err = d.Set("msn_hotmail", flattenSpamfilterProfileMsnHotmail(o["msn-hotmail"], d, "msn_hotmail")); err != nil {
				if !fortiAPIPatch(o["msn-hotmail"]) {
					return fmt.Errorf("Error reading msn_hotmail: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("yahoo_mail", flattenSpamfilterProfileYahooMail(o["yahoo-mail"], d, "yahoo_mail")); err != nil {
			if !fortiAPIPatch(o["yahoo-mail"]) {
				return fmt.Errorf("Error reading yahoo_mail: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("yahoo_mail"); ok {
			if err = d.Set("yahoo_mail", flattenSpamfilterProfileYahooMail(o["yahoo-mail"], d, "yahoo_mail")); err != nil {
				if !fortiAPIPatch(o["yahoo-mail"]) {
					return fmt.Errorf("Error reading yahoo_mail: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("gmail", flattenSpamfilterProfileGmail(o["gmail"], d, "gmail")); err != nil {
			if !fortiAPIPatch(o["gmail"]) {
				return fmt.Errorf("Error reading gmail: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("gmail"); ok {
			if err = d.Set("gmail", flattenSpamfilterProfileGmail(o["gmail"], d, "gmail")); err != nil {
				if !fortiAPIPatch(o["gmail"]) {
					return fmt.Errorf("Error reading gmail: %v", err)
				}
			}
		}
	}

	if err = d.Set("spam_bword_threshold", flattenSpamfilterProfileSpamBwordThreshold(o["spam-bword-threshold"], d, "spam_bword_threshold")); err != nil {
		if !fortiAPIPatch(o["spam-bword-threshold"]) {
			return fmt.Errorf("Error reading spam_bword_threshold: %v", err)
		}
	}

	if err = d.Set("spam_bword_table", flattenSpamfilterProfileSpamBwordTable(o["spam-bword-table"], d, "spam_bword_table")); err != nil {
		if !fortiAPIPatch(o["spam-bword-table"]) {
			return fmt.Errorf("Error reading spam_bword_table: %v", err)
		}
	}

	if err = d.Set("spam_bwl_table", flattenSpamfilterProfileSpamBwlTable(o["spam-bwl-table"], d, "spam_bwl_table")); err != nil {
		if !fortiAPIPatch(o["spam-bwl-table"]) {
			return fmt.Errorf("Error reading spam_bwl_table: %v", err)
		}
	}

	if err = d.Set("spam_mheader_table", flattenSpamfilterProfileSpamMheaderTable(o["spam-mheader-table"], d, "spam_mheader_table")); err != nil {
		if !fortiAPIPatch(o["spam-mheader-table"]) {
			return fmt.Errorf("Error reading spam_mheader_table: %v", err)
		}
	}

	if err = d.Set("spam_rbl_table", flattenSpamfilterProfileSpamRblTable(o["spam-rbl-table"], d, "spam_rbl_table")); err != nil {
		if !fortiAPIPatch(o["spam-rbl-table"]) {
			return fmt.Errorf("Error reading spam_rbl_table: %v", err)
		}
	}

	if err = d.Set("spam_iptrust_table", flattenSpamfilterProfileSpamIptrustTable(o["spam-iptrust-table"], d, "spam_iptrust_table")); err != nil {
		if !fortiAPIPatch(o["spam-iptrust-table"]) {
			return fmt.Errorf("Error reading spam_iptrust_table: %v", err)
		}
	}

	return nil
}

func flattenSpamfilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSpamfilterProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileFlowBased(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileSpamLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileSpamLogFortiguardResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileSpamFiltering(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileExternal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileImap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandSpamfilterProfileImapLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandSpamfilterProfileImapAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "tag_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag-type"], _ = expandSpamfilterProfileImapTagType(d, i["tag_type"], pre_append)
	}
	pre_append = pre + ".0." + "tag_msg"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag-msg"], _ = expandSpamfilterProfileImapTagMsg(d, i["tag_msg"], pre_append)
	}

	return result, nil
}

func expandSpamfilterProfileImapLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileImapAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileImapTagType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileImapTagMsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfilePop3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandSpamfilterProfilePop3Log(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandSpamfilterProfilePop3Action(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "tag_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag-type"], _ = expandSpamfilterProfilePop3TagType(d, i["tag_type"], pre_append)
	}
	pre_append = pre + ".0." + "tag_msg"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag-msg"], _ = expandSpamfilterProfilePop3TagMsg(d, i["tag_msg"], pre_append)
	}

	return result, nil
}

func expandSpamfilterProfilePop3Log(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfilePop3Action(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfilePop3TagType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfilePop3TagMsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileSmtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandSpamfilterProfileSmtpLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandSpamfilterProfileSmtpAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "tag_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag-type"], _ = expandSpamfilterProfileSmtpTagType(d, i["tag_type"], pre_append)
	}
	pre_append = pre + ".0." + "tag_msg"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag-msg"], _ = expandSpamfilterProfileSmtpTagMsg(d, i["tag_msg"], pre_append)
	}
	pre_append = pre + ".0." + "hdrip"
	if _, ok := d.GetOk(pre_append); ok {
		result["hdrip"], _ = expandSpamfilterProfileSmtpHdrip(d, i["hdrip"], pre_append)
	}
	pre_append = pre + ".0." + "local_override"
	if _, ok := d.GetOk(pre_append); ok {
		result["local-override"], _ = expandSpamfilterProfileSmtpLocalOverride(d, i["local_override"], pre_append)
	}

	return result, nil
}

func expandSpamfilterProfileSmtpLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileSmtpAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileSmtpTagType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileSmtpTagMsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileSmtpHdrip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileSmtpLocalOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileMapi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandSpamfilterProfileMapiLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandSpamfilterProfileMapiAction(d, i["action"], pre_append)
	}

	return result, nil
}

func expandSpamfilterProfileMapiLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileMapiAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileMsnHotmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandSpamfilterProfileMsnHotmailLog(d, i["log"], pre_append)
	}

	return result, nil
}

func expandSpamfilterProfileMsnHotmailLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileYahooMail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandSpamfilterProfileYahooMailLog(d, i["log"], pre_append)
	}

	return result, nil
}

func expandSpamfilterProfileYahooMailLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileGmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandSpamfilterProfileGmailLog(d, i["log"], pre_append)
	}

	return result, nil
}

func expandSpamfilterProfileGmailLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileSpamBwordThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileSpamBwordTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileSpamBwlTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileSpamMheaderTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileSpamRblTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterProfileSpamIptrustTable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSpamfilterProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSpamfilterProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandSpamfilterProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("flow_based"); ok {
		t, err := expandSpamfilterProfileFlowBased(d, v, "flow_based")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flow-based"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok {
		t, err := expandSpamfilterProfileReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("spam_log"); ok {
		t, err := expandSpamfilterProfileSpamLog(d, v, "spam_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-log"] = t
		}
	}

	if v, ok := d.GetOk("spam_log_fortiguard_response"); ok {
		t, err := expandSpamfilterProfileSpamLogFortiguardResponse(d, v, "spam_log_fortiguard_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-log-fortiguard-response"] = t
		}
	}

	if v, ok := d.GetOk("spam_filtering"); ok {
		t, err := expandSpamfilterProfileSpamFiltering(d, v, "spam_filtering")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-filtering"] = t
		}
	}

	if v, ok := d.GetOk("external"); ok {
		t, err := expandSpamfilterProfileExternal(d, v, "external")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok {
		t, err := expandSpamfilterProfileOptions(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("imap"); ok {
		t, err := expandSpamfilterProfileImap(d, v, "imap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imap"] = t
		}
	}

	if v, ok := d.GetOk("pop3"); ok {
		t, err := expandSpamfilterProfilePop3(d, v, "pop3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3"] = t
		}
	}

	if v, ok := d.GetOk("smtp"); ok {
		t, err := expandSpamfilterProfileSmtp(d, v, "smtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtp"] = t
		}
	}

	if v, ok := d.GetOk("mapi"); ok {
		t, err := expandSpamfilterProfileMapi(d, v, "mapi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapi"] = t
		}
	}

	if v, ok := d.GetOk("msn_hotmail"); ok {
		t, err := expandSpamfilterProfileMsnHotmail(d, v, "msn_hotmail")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msn-hotmail"] = t
		}
	}

	if v, ok := d.GetOk("yahoo_mail"); ok {
		t, err := expandSpamfilterProfileYahooMail(d, v, "yahoo_mail")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["yahoo-mail"] = t
		}
	}

	if v, ok := d.GetOk("gmail"); ok {
		t, err := expandSpamfilterProfileGmail(d, v, "gmail")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gmail"] = t
		}
	}

	if v, ok := d.GetOkExists("spam_bword_threshold"); ok {
		t, err := expandSpamfilterProfileSpamBwordThreshold(d, v, "spam_bword_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-bword-threshold"] = t
		}
	}

	if v, ok := d.GetOkExists("spam_bword_table"); ok {
		t, err := expandSpamfilterProfileSpamBwordTable(d, v, "spam_bword_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-bword-table"] = t
		}
	}

	if v, ok := d.GetOkExists("spam_bwl_table"); ok {
		t, err := expandSpamfilterProfileSpamBwlTable(d, v, "spam_bwl_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-bwl-table"] = t
		}
	}

	if v, ok := d.GetOkExists("spam_mheader_table"); ok {
		t, err := expandSpamfilterProfileSpamMheaderTable(d, v, "spam_mheader_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-mheader-table"] = t
		}
	}

	if v, ok := d.GetOkExists("spam_rbl_table"); ok {
		t, err := expandSpamfilterProfileSpamRblTable(d, v, "spam_rbl_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-rbl-table"] = t
		}
	}

	if v, ok := d.GetOkExists("spam_iptrust_table"); ok {
		t, err := expandSpamfilterProfileSpamIptrustTable(d, v, "spam_iptrust_table")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam-iptrust-table"] = t
		}
	}

	return &obj, nil
}
