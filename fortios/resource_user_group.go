// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure user groups.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceUserGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserGroupCreate,
		Read:   resourceUserGroupRead,
		Update: resourceUserGroupUpdate,
		Delete: resourceUserGroupDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"group_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authtimeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 43200),
				Optional:     true,
				Computed:     true,
			},
			"auth_concurrent_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_concurrent_value": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),
				Optional:     true,
				Computed:     true,
			},
			"http_digest_realm": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"sso_attribute_value": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
				Computed:     true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"match": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"server_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"group_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sponsor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"company": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mobile_phone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"expire_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expire": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 31536000),
				Optional:     true,
				Computed:     true,
			},
			"max_accounts": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1024),
				Optional:     true,
				Computed:     true,
			},
			"multiple_guest_add": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"guest": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"user_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"password": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
							Sensitive:    true,
						},
						"mobile_phone": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"sponsor": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"company": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"email": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"expiration": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"comment": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
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

func resourceUserGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserGroup resource while getting object: %v", err)
	}

	o, err := c.CreateUserGroup(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating UserGroup resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserGroup")
	}

	return resourceUserGroupRead(d, m)
}

func resourceUserGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserGroup resource while getting object: %v", err)
	}

	o, err := c.UpdateUserGroup(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserGroup")
	}

	return resourceUserGroupRead(d, m)
}

func resourceUserGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUserGroup(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UserGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadUserGroup(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserGroup resource from API: %v", err)
	}
	return nil
}

func flattenUserGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupGroupType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupAuthtimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupAuthConcurrentOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupAuthConcurrentValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupHttpDigestRealm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupSsoAttributeValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupMember(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenUserGroupMemberName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenUserGroupMemberName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupMatch(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenUserGroupMatchId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_name"
		if _, ok := i["server-name"]; ok {

			tmp["server_name"] = flattenUserGroupMatchServerName(i["server-name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_name"
		if _, ok := i["group-name"]; ok {

			tmp["group_name"] = flattenUserGroupMatchGroupName(i["group-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenUserGroupMatchId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupMatchServerName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupMatchGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupUserId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupUserName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupSponsor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupCompany(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupEmail(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupMobilePhone(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupSmsServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupSmsCustomServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupExpireType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupExpire(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupMaxAccounts(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupMultipleGuestAdd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupGuest(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenUserGroupGuestId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_id"
		if _, ok := i["user-id"]; ok {

			tmp["user_id"] = flattenUserGroupGuestUserId(i["user-id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenUserGroupGuestName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {

			tmp["password"] = flattenUserGroupGuestPassword(i["password"], d, pre_append, sv)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["password"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mobile_phone"
		if _, ok := i["mobile-phone"]; ok {

			tmp["mobile_phone"] = flattenUserGroupGuestMobilePhone(i["mobile-phone"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sponsor"
		if _, ok := i["sponsor"]; ok {

			tmp["sponsor"] = flattenUserGroupGuestSponsor(i["sponsor"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company"
		if _, ok := i["company"]; ok {

			tmp["company"] = flattenUserGroupGuestCompany(i["company"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email"
		if _, ok := i["email"]; ok {

			tmp["email"] = flattenUserGroupGuestEmail(i["email"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiration"
		if _, ok := i["expiration"]; ok {

			tmp["expiration"] = flattenUserGroupGuestExpiration(i["expiration"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {

			tmp["comment"] = flattenUserGroupGuestComment(i["comment"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "user_id", d)
	return result
}

func flattenUserGroupGuestId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupGuestUserId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupGuestName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupGuestPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupGuestMobilePhone(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupGuestSponsor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupGuestCompany(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupGuestEmail(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupGuestExpiration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserGroupGuestComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenUserGroupName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenUserGroupId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("group_type", flattenUserGroupGroupType(o["group-type"], d, "group_type", sv)); err != nil {
		if !fortiAPIPatch(o["group-type"]) {
			return fmt.Errorf("Error reading group_type: %v", err)
		}
	}

	if err = d.Set("authtimeout", flattenUserGroupAuthtimeout(o["authtimeout"], d, "authtimeout", sv)); err != nil {
		if !fortiAPIPatch(o["authtimeout"]) {
			return fmt.Errorf("Error reading authtimeout: %v", err)
		}
	}

	if err = d.Set("auth_concurrent_override", flattenUserGroupAuthConcurrentOverride(o["auth-concurrent-override"], d, "auth_concurrent_override", sv)); err != nil {
		if !fortiAPIPatch(o["auth-concurrent-override"]) {
			return fmt.Errorf("Error reading auth_concurrent_override: %v", err)
		}
	}

	if err = d.Set("auth_concurrent_value", flattenUserGroupAuthConcurrentValue(o["auth-concurrent-value"], d, "auth_concurrent_value", sv)); err != nil {
		if !fortiAPIPatch(o["auth-concurrent-value"]) {
			return fmt.Errorf("Error reading auth_concurrent_value: %v", err)
		}
	}

	if err = d.Set("http_digest_realm", flattenUserGroupHttpDigestRealm(o["http-digest-realm"], d, "http_digest_realm", sv)); err != nil {
		if !fortiAPIPatch(o["http-digest-realm"]) {
			return fmt.Errorf("Error reading http_digest_realm: %v", err)
		}
	}

	if err = d.Set("sso_attribute_value", flattenUserGroupSsoAttributeValue(o["sso-attribute-value"], d, "sso_attribute_value", sv)); err != nil {
		if !fortiAPIPatch(o["sso-attribute-value"]) {
			return fmt.Errorf("Error reading sso_attribute_value: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("member", flattenUserGroupMember(o["member"], d, "member", sv)); err != nil {
			if !fortiAPIPatch(o["member"]) {
				return fmt.Errorf("Error reading member: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("member"); ok {
			if err = d.Set("member", flattenUserGroupMember(o["member"], d, "member", sv)); err != nil {
				if !fortiAPIPatch(o["member"]) {
					return fmt.Errorf("Error reading member: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("match", flattenUserGroupMatch(o["match"], d, "match", sv)); err != nil {
			if !fortiAPIPatch(o["match"]) {
				return fmt.Errorf("Error reading match: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("match"); ok {
			if err = d.Set("match", flattenUserGroupMatch(o["match"], d, "match", sv)); err != nil {
				if !fortiAPIPatch(o["match"]) {
					return fmt.Errorf("Error reading match: %v", err)
				}
			}
		}
	}

	if err = d.Set("user_id", flattenUserGroupUserId(o["user-id"], d, "user_id", sv)); err != nil {
		if !fortiAPIPatch(o["user-id"]) {
			return fmt.Errorf("Error reading user_id: %v", err)
		}
	}

	if err = d.Set("password", flattenUserGroupPassword(o["password"], d, "password", sv)); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("user_name", flattenUserGroupUserName(o["user-name"], d, "user_name", sv)); err != nil {
		if !fortiAPIPatch(o["user-name"]) {
			return fmt.Errorf("Error reading user_name: %v", err)
		}
	}

	if err = d.Set("sponsor", flattenUserGroupSponsor(o["sponsor"], d, "sponsor", sv)); err != nil {
		if !fortiAPIPatch(o["sponsor"]) {
			return fmt.Errorf("Error reading sponsor: %v", err)
		}
	}

	if err = d.Set("company", flattenUserGroupCompany(o["company"], d, "company", sv)); err != nil {
		if !fortiAPIPatch(o["company"]) {
			return fmt.Errorf("Error reading company: %v", err)
		}
	}

	if err = d.Set("email", flattenUserGroupEmail(o["email"], d, "email", sv)); err != nil {
		if !fortiAPIPatch(o["email"]) {
			return fmt.Errorf("Error reading email: %v", err)
		}
	}

	if err = d.Set("mobile_phone", flattenUserGroupMobilePhone(o["mobile-phone"], d, "mobile_phone", sv)); err != nil {
		if !fortiAPIPatch(o["mobile-phone"]) {
			return fmt.Errorf("Error reading mobile_phone: %v", err)
		}
	}

	if err = d.Set("sms_server", flattenUserGroupSmsServer(o["sms-server"], d, "sms_server", sv)); err != nil {
		if !fortiAPIPatch(o["sms-server"]) {
			return fmt.Errorf("Error reading sms_server: %v", err)
		}
	}

	if err = d.Set("sms_custom_server", flattenUserGroupSmsCustomServer(o["sms-custom-server"], d, "sms_custom_server", sv)); err != nil {
		if !fortiAPIPatch(o["sms-custom-server"]) {
			return fmt.Errorf("Error reading sms_custom_server: %v", err)
		}
	}

	if err = d.Set("expire_type", flattenUserGroupExpireType(o["expire-type"], d, "expire_type", sv)); err != nil {
		if !fortiAPIPatch(o["expire-type"]) {
			return fmt.Errorf("Error reading expire_type: %v", err)
		}
	}

	if err = d.Set("expire", flattenUserGroupExpire(o["expire"], d, "expire", sv)); err != nil {
		if !fortiAPIPatch(o["expire"]) {
			return fmt.Errorf("Error reading expire: %v", err)
		}
	}

	if err = d.Set("max_accounts", flattenUserGroupMaxAccounts(o["max-accounts"], d, "max_accounts", sv)); err != nil {
		if !fortiAPIPatch(o["max-accounts"]) {
			return fmt.Errorf("Error reading max_accounts: %v", err)
		}
	}

	if err = d.Set("multiple_guest_add", flattenUserGroupMultipleGuestAdd(o["multiple-guest-add"], d, "multiple_guest_add", sv)); err != nil {
		if !fortiAPIPatch(o["multiple-guest-add"]) {
			return fmt.Errorf("Error reading multiple_guest_add: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("guest", flattenUserGroupGuest(o["guest"], d, "guest", sv)); err != nil {
			if !fortiAPIPatch(o["guest"]) {
				return fmt.Errorf("Error reading guest: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("guest"); ok {
			if err = d.Set("guest", flattenUserGroupGuest(o["guest"], d, "guest", sv)); err != nil {
				if !fortiAPIPatch(o["guest"]) {
					return fmt.Errorf("Error reading guest: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenUserGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGroupType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupAuthtimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupAuthConcurrentOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupAuthConcurrentValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupHttpDigestRealm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupSsoAttributeValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupMember(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandUserGroupMemberName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserGroupMemberName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandUserGroupMatchId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["server-name"], _ = expandUserGroupMatchServerName(d, i["server_name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["group-name"], _ = expandUserGroupMatchGroupName(d, i["group_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserGroupMatchId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupMatchServerName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupMatchGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupUserId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupUserName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupSponsor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupCompany(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupEmail(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupMobilePhone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupSmsServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupSmsCustomServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupExpireType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupExpire(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupMaxAccounts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupMultipleGuestAdd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGuest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandUserGroupGuestId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["user-id"], _ = expandUserGroupGuestUserId(d, i["user_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandUserGroupGuestName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["password"], _ = expandUserGroupGuestPassword(d, i["password"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mobile_phone"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["mobile-phone"], _ = expandUserGroupGuestMobilePhone(d, i["mobile_phone"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sponsor"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["sponsor"], _ = expandUserGroupGuestSponsor(d, i["sponsor"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["company"], _ = expandUserGroupGuestCompany(d, i["company"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "email"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["email"], _ = expandUserGroupGuestEmail(d, i["email"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiration"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["expiration"], _ = expandUserGroupGuestExpiration(d, i["expiration"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["comment"], _ = expandUserGroupGuestComment(d, i["comment"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserGroupGuestId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGuestUserId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGuestName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGuestPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGuestMobilePhone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGuestSponsor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGuestCompany(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGuestEmail(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGuestExpiration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserGroupGuestComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserGroup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandUserGroupName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandUserGroupId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("group_type"); ok {

		t, err := expandUserGroupGroupType(d, v, "group_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-type"] = t
		}
	}

	if v, ok := d.GetOkExists("authtimeout"); ok {

		t, err := expandUserGroupAuthtimeout(d, v, "authtimeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authtimeout"] = t
		}
	}

	if v, ok := d.GetOk("auth_concurrent_override"); ok {

		t, err := expandUserGroupAuthConcurrentOverride(d, v, "auth_concurrent_override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-override"] = t
		}
	}

	if v, ok := d.GetOkExists("auth_concurrent_value"); ok {

		t, err := expandUserGroupAuthConcurrentValue(d, v, "auth_concurrent_value", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-value"] = t
		}
	}

	if v, ok := d.GetOk("http_digest_realm"); ok {

		t, err := expandUserGroupHttpDigestRealm(d, v, "http_digest_realm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-digest-realm"] = t
		}
	}

	if v, ok := d.GetOk("sso_attribute_value"); ok {

		t, err := expandUserGroupSsoAttributeValue(d, v, "sso_attribute_value", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-attribute-value"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok || d.HasChange("member") {

		t, err := expandUserGroupMember(d, v, "member", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("match"); ok || d.HasChange("match") {

		t, err := expandUserGroupMatch(d, v, "match", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match"] = t
		}
	}

	if v, ok := d.GetOk("user_id"); ok {

		t, err := expandUserGroupUserId(d, v, "user_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-id"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {

		t, err := expandUserGroupPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("user_name"); ok {

		t, err := expandUserGroupUserName(d, v, "user_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-name"] = t
		}
	}

	if v, ok := d.GetOk("sponsor"); ok {

		t, err := expandUserGroupSponsor(d, v, "sponsor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sponsor"] = t
		}
	}

	if v, ok := d.GetOk("company"); ok {

		t, err := expandUserGroupCompany(d, v, "company", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["company"] = t
		}
	}

	if v, ok := d.GetOk("email"); ok {

		t, err := expandUserGroupEmail(d, v, "email", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email"] = t
		}
	}

	if v, ok := d.GetOk("mobile_phone"); ok {

		t, err := expandUserGroupMobilePhone(d, v, "mobile_phone", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mobile-phone"] = t
		}
	}

	if v, ok := d.GetOk("sms_server"); ok {

		t, err := expandUserGroupSmsServer(d, v, "sms_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-server"] = t
		}
	}

	if v, ok := d.GetOk("sms_custom_server"); ok {

		t, err := expandUserGroupSmsCustomServer(d, v, "sms_custom_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-custom-server"] = t
		}
	}

	if v, ok := d.GetOk("expire_type"); ok {

		t, err := expandUserGroupExpireType(d, v, "expire_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-type"] = t
		}
	}

	if v, ok := d.GetOk("expire"); ok {

		t, err := expandUserGroupExpire(d, v, "expire", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire"] = t
		}
	}

	if v, ok := d.GetOkExists("max_accounts"); ok {

		t, err := expandUserGroupMaxAccounts(d, v, "max_accounts", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-accounts"] = t
		}
	}

	if v, ok := d.GetOk("multiple_guest_add"); ok {

		t, err := expandUserGroupMultipleGuestAdd(d, v, "multiple_guest_add", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multiple-guest-add"] = t
		}
	}

	if v, ok := d.GetOk("guest"); ok || d.HasChange("guest") {

		t, err := expandUserGroupGuest(d, v, "guest", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest"] = t
		}
	}

	return &obj, nil
}
