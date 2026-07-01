// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: OpenID Connect server entry configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserOidc() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserOidcCreate,
		Read:   resourceUserOidcRead,
		Update: resourceUserOidcUpdate,
		Delete: resourceUserOidcDelete,

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
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"client_secret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 123),
				Optional:     true,
				Sensitive:    true,
			},
			"discovery_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"authorization_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"token_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"jwks_uri": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"domain_hint": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"issuer": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"verify_issuer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_attr_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_regex": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"ldap_server": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"clock_tolerance": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 300),
				Optional:     true,
				Computed:     true,
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

func resourceUserOidcCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserOidc(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserOidc resource while getting object: %v", err)
	}

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserOidc(mkey, vdomparam)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserOidc(obj, mkey, vdomparam)
			if err != nil {
				return fmt.Errorf("Error updating UserOidc resource: %v", err)
			}
		}
	}

	if !existing {
		o, err = c.CreateUserOidc(obj, vdomparam)

		if err != nil {
			return fmt.Errorf("Error creating UserOidc resource: %v", err)
		}
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserOidc")
	}

	return resourceUserOidcRead(d, m)
}

func resourceUserOidcUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserOidc(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserOidc resource while getting object: %v", err)
	}

	o, err := c.UpdateUserOidc(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserOidc resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserOidc")
	}

	return resourceUserOidcRead(d, m)
}

func resourceUserOidcDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUserOidc(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UserOidc resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserOidcRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserOidc(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserOidc resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserOidc(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserOidc resource from API: %v", err)
	}
	return nil
}

func flattenUserOidcName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserOidcType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserOidcClientId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserOidcDiscoveryUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserOidcAuthorizationUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserOidcTokenUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserOidcJwksUri(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserOidcDomainHint(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserOidcIssuer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserOidcVerifyIssuer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserOidcUserAttrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserOidcUserRegex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserOidcLdapServer(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenUserOidcLdapServerName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenUserOidcLdapServerName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserOidcClockTolerance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectUserOidc(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenUserOidcName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenUserOidcType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("client_id", flattenUserOidcClientId(o["client-id"], d, "client_id", sv)); err != nil {
		if !fortiAPIPatch(o["client-id"]) {
			return fmt.Errorf("Error reading client_id: %v", err)
		}
	}

	if err = d.Set("discovery_url", flattenUserOidcDiscoveryUrl(o["discovery-url"], d, "discovery_url", sv)); err != nil {
		if !fortiAPIPatch(o["discovery-url"]) {
			return fmt.Errorf("Error reading discovery_url: %v", err)
		}
	}

	if err = d.Set("authorization_url", flattenUserOidcAuthorizationUrl(o["authorization-url"], d, "authorization_url", sv)); err != nil {
		if !fortiAPIPatch(o["authorization-url"]) {
			return fmt.Errorf("Error reading authorization_url: %v", err)
		}
	}

	if err = d.Set("token_url", flattenUserOidcTokenUrl(o["token-url"], d, "token_url", sv)); err != nil {
		if !fortiAPIPatch(o["token-url"]) {
			return fmt.Errorf("Error reading token_url: %v", err)
		}
	}

	if err = d.Set("jwks_uri", flattenUserOidcJwksUri(o["jwks-uri"], d, "jwks_uri", sv)); err != nil {
		if !fortiAPIPatch(o["jwks-uri"]) {
			return fmt.Errorf("Error reading jwks_uri: %v", err)
		}
	}

	if err = d.Set("domain_hint", flattenUserOidcDomainHint(o["domain-hint"], d, "domain_hint", sv)); err != nil {
		if !fortiAPIPatch(o["domain-hint"]) {
			return fmt.Errorf("Error reading domain_hint: %v", err)
		}
	}

	if err = d.Set("issuer", flattenUserOidcIssuer(o["issuer"], d, "issuer", sv)); err != nil {
		if !fortiAPIPatch(o["issuer"]) {
			return fmt.Errorf("Error reading issuer: %v", err)
		}
	}

	if err = d.Set("verify_issuer", flattenUserOidcVerifyIssuer(o["verify-issuer"], d, "verify_issuer", sv)); err != nil {
		if !fortiAPIPatch(o["verify-issuer"]) {
			return fmt.Errorf("Error reading verify_issuer: %v", err)
		}
	}

	if err = d.Set("user_attr_name", flattenUserOidcUserAttrName(o["user-attr-name"], d, "user_attr_name", sv)); err != nil {
		if !fortiAPIPatch(o["user-attr-name"]) {
			return fmt.Errorf("Error reading user_attr_name: %v", err)
		}
	}

	if err = d.Set("user_regex", flattenUserOidcUserRegex(o["user-regex"], d, "user_regex", sv)); err != nil {
		if !fortiAPIPatch(o["user-regex"]) {
			return fmt.Errorf("Error reading user_regex: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ldap_server", flattenUserOidcLdapServer(o["ldap-server"], d, "ldap_server", sv)); err != nil {
			if !fortiAPIPatch(o["ldap-server"]) {
				return fmt.Errorf("Error reading ldap_server: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ldap_server"); ok {
			if err = d.Set("ldap_server", flattenUserOidcLdapServer(o["ldap-server"], d, "ldap_server", sv)); err != nil {
				if !fortiAPIPatch(o["ldap-server"]) {
					return fmt.Errorf("Error reading ldap_server: %v", err)
				}
			}
		}
	}

	if err = d.Set("clock_tolerance", flattenUserOidcClockTolerance(o["clock-tolerance"], d, "clock_tolerance", sv)); err != nil {
		if !fortiAPIPatch(o["clock-tolerance"]) {
			return fmt.Errorf("Error reading clock_tolerance: %v", err)
		}
	}

	return nil
}

func flattenUserOidcFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserOidcName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserOidcType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserOidcClientId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserOidcClientSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserOidcDiscoveryUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserOidcAuthorizationUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserOidcTokenUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserOidcJwksUri(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserOidcDomainHint(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserOidcIssuer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserOidcVerifyIssuer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserOidcUserAttrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserOidcUserRegex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserOidcLdapServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandUserOidcLdapServerName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserOidcLdapServerName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserOidcClockTolerance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserOidc(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserOidcName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandUserOidcType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("client_id"); ok {
		t, err := expandUserOidcClientId(d, v, "client_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-id"] = t
		}
	} else if d.HasChange("client_id") {
		obj["client-id"] = nil
	}

	if v, ok := d.GetOk("client_secret"); ok {
		t, err := expandUserOidcClientSecret(d, v, "client_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-secret"] = t
		}
	} else if d.HasChange("client_secret") {
		obj["client-secret"] = nil
	}

	if v, ok := d.GetOk("discovery_url"); ok {
		t, err := expandUserOidcDiscoveryUrl(d, v, "discovery_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["discovery-url"] = t
		}
	} else if d.HasChange("discovery_url") {
		obj["discovery-url"] = nil
	}

	if v, ok := d.GetOk("authorization_url"); ok {
		t, err := expandUserOidcAuthorizationUrl(d, v, "authorization_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorization-url"] = t
		}
	} else if d.HasChange("authorization_url") {
		obj["authorization-url"] = nil
	}

	if v, ok := d.GetOk("token_url"); ok {
		t, err := expandUserOidcTokenUrl(d, v, "token_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["token-url"] = t
		}
	} else if d.HasChange("token_url") {
		obj["token-url"] = nil
	}

	if v, ok := d.GetOk("jwks_uri"); ok {
		t, err := expandUserOidcJwksUri(d, v, "jwks_uri", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jwks-uri"] = t
		}
	} else if d.HasChange("jwks_uri") {
		obj["jwks-uri"] = nil
	}

	if v, ok := d.GetOk("domain_hint"); ok {
		t, err := expandUserOidcDomainHint(d, v, "domain_hint", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-hint"] = t
		}
	} else if d.HasChange("domain_hint") {
		obj["domain-hint"] = nil
	}

	if v, ok := d.GetOk("issuer"); ok {
		t, err := expandUserOidcIssuer(d, v, "issuer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["issuer"] = t
		}
	} else if d.HasChange("issuer") {
		obj["issuer"] = nil
	}

	if v, ok := d.GetOk("verify_issuer"); ok {
		t, err := expandUserOidcVerifyIssuer(d, v, "verify_issuer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-issuer"] = t
		}
	}

	if v, ok := d.GetOk("user_attr_name"); ok {
		t, err := expandUserOidcUserAttrName(d, v, "user_attr_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-attr-name"] = t
		}
	}

	if v, ok := d.GetOk("user_regex"); ok {
		t, err := expandUserOidcUserRegex(d, v, "user_regex", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-regex"] = t
		}
	} else if d.HasChange("user_regex") {
		obj["user-regex"] = nil
	}

	if v, ok := d.GetOk("ldap_server"); ok || d.HasChange("ldap_server") {
		t, err := expandUserOidcLdapServer(d, v, "ldap_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOkExists("clock_tolerance"); ok {
		t, err := expandUserOidcClockTolerance(d, v, "clock_tolerance", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clock-tolerance"] = t
		}
	}

	return &obj, nil
}
