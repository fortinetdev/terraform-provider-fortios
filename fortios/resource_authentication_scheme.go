// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Authentication Schemes.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAuthenticationScheme() *schema.Resource {
	return &schema.Resource{
		Create: resourceAuthenticationSchemeCreate,
		Read:   resourceAuthenticationSchemeRead,
		Update: resourceAuthenticationSchemeUpdate,
		Delete: resourceAuthenticationSchemeDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"negotiate_ntlm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kerberos_keytab": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"domain_controller": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"saml_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"saml_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 1200),
				Optional:     true,
				Computed:     true,
			},
			"fsso_agent_for_ntlm": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"require_tfa": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fsso_guest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_database": &schema.Schema{
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
			"ssh_ca": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"external_idp": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
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

func resourceAuthenticationSchemeCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectAuthenticationScheme(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating AuthenticationScheme resource while getting object: %v", err)
	}

	o, err := c.CreateAuthenticationScheme(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating AuthenticationScheme resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("AuthenticationScheme")
	}

	return resourceAuthenticationSchemeRead(d, m)
}

func resourceAuthenticationSchemeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectAuthenticationScheme(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating AuthenticationScheme resource while getting object: %v", err)
	}

	o, err := c.UpdateAuthenticationScheme(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating AuthenticationScheme resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("AuthenticationScheme")
	}

	return resourceAuthenticationSchemeRead(d, m)
}

func resourceAuthenticationSchemeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteAuthenticationScheme(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting AuthenticationScheme resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAuthenticationSchemeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadAuthenticationScheme(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading AuthenticationScheme resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAuthenticationScheme(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading AuthenticationScheme resource from API: %v", err)
	}
	return nil
}

func flattenAuthenticationSchemeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSchemeMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSchemeNegotiateNtlm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSchemeKerberosKeytab(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSchemeDomainController(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSchemeSamlServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSchemeSamlTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenAuthenticationSchemeFssoAgentForNtlm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSchemeRequireTfa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSchemeFssoGuest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSchemeUserCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSchemeUserDatabase(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenAuthenticationSchemeUserDatabaseName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenAuthenticationSchemeUserDatabaseName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSchemeSshCa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSchemeExternalIdp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectAuthenticationScheme(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenAuthenticationSchemeName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("method", flattenAuthenticationSchemeMethod(o["method"], d, "method", sv)); err != nil {
		if !fortiAPIPatch(o["method"]) {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("negotiate_ntlm", flattenAuthenticationSchemeNegotiateNtlm(o["negotiate-ntlm"], d, "negotiate_ntlm", sv)); err != nil {
		if !fortiAPIPatch(o["negotiate-ntlm"]) {
			return fmt.Errorf("Error reading negotiate_ntlm: %v", err)
		}
	}

	if err = d.Set("kerberos_keytab", flattenAuthenticationSchemeKerberosKeytab(o["kerberos-keytab"], d, "kerberos_keytab", sv)); err != nil {
		if !fortiAPIPatch(o["kerberos-keytab"]) {
			return fmt.Errorf("Error reading kerberos_keytab: %v", err)
		}
	}

	if err = d.Set("domain_controller", flattenAuthenticationSchemeDomainController(o["domain-controller"], d, "domain_controller", sv)); err != nil {
		if !fortiAPIPatch(o["domain-controller"]) {
			return fmt.Errorf("Error reading domain_controller: %v", err)
		}
	}

	if err = d.Set("saml_server", flattenAuthenticationSchemeSamlServer(o["saml-server"], d, "saml_server", sv)); err != nil {
		if !fortiAPIPatch(o["saml-server"]) {
			return fmt.Errorf("Error reading saml_server: %v", err)
		}
	}

	if err = d.Set("saml_timeout", flattenAuthenticationSchemeSamlTimeout(o["saml-timeout"], d, "saml_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["saml-timeout"]) {
			return fmt.Errorf("Error reading saml_timeout: %v", err)
		}
	}

	if err = d.Set("fsso_agent_for_ntlm", flattenAuthenticationSchemeFssoAgentForNtlm(o["fsso-agent-for-ntlm"], d, "fsso_agent_for_ntlm", sv)); err != nil {
		if !fortiAPIPatch(o["fsso-agent-for-ntlm"]) {
			return fmt.Errorf("Error reading fsso_agent_for_ntlm: %v", err)
		}
	}

	if err = d.Set("require_tfa", flattenAuthenticationSchemeRequireTfa(o["require-tfa"], d, "require_tfa", sv)); err != nil {
		if !fortiAPIPatch(o["require-tfa"]) {
			return fmt.Errorf("Error reading require_tfa: %v", err)
		}
	}

	if err = d.Set("fsso_guest", flattenAuthenticationSchemeFssoGuest(o["fsso-guest"], d, "fsso_guest", sv)); err != nil {
		if !fortiAPIPatch(o["fsso-guest"]) {
			return fmt.Errorf("Error reading fsso_guest: %v", err)
		}
	}

	if err = d.Set("user_cert", flattenAuthenticationSchemeUserCert(o["user-cert"], d, "user_cert", sv)); err != nil {
		if !fortiAPIPatch(o["user-cert"]) {
			return fmt.Errorf("Error reading user_cert: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("user_database", flattenAuthenticationSchemeUserDatabase(o["user-database"], d, "user_database", sv)); err != nil {
			if !fortiAPIPatch(o["user-database"]) {
				return fmt.Errorf("Error reading user_database: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("user_database"); ok {
			if err = d.Set("user_database", flattenAuthenticationSchemeUserDatabase(o["user-database"], d, "user_database", sv)); err != nil {
				if !fortiAPIPatch(o["user-database"]) {
					return fmt.Errorf("Error reading user_database: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssh_ca", flattenAuthenticationSchemeSshCa(o["ssh-ca"], d, "ssh_ca", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-ca"]) {
			return fmt.Errorf("Error reading ssh_ca: %v", err)
		}
	}

	if err = d.Set("external_idp", flattenAuthenticationSchemeExternalIdp(o["external-idp"], d, "external_idp", sv)); err != nil {
		if !fortiAPIPatch(o["external-idp"]) {
			return fmt.Errorf("Error reading external_idp: %v", err)
		}
	}

	return nil
}

func flattenAuthenticationSchemeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandAuthenticationSchemeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeNegotiateNtlm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeKerberosKeytab(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeDomainController(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeSamlServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeSamlTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeFssoAgentForNtlm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeRequireTfa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeFssoGuest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeUserCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeUserDatabase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandAuthenticationSchemeUserDatabaseName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandAuthenticationSchemeUserDatabaseName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeSshCa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeExternalIdp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectAuthenticationScheme(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandAuthenticationSchemeName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok {
		t, err := expandAuthenticationSchemeMethod(d, v, "method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	} else if d.HasChange("method") {
		obj["method"] = nil
	}

	if v, ok := d.GetOk("negotiate_ntlm"); ok {
		t, err := expandAuthenticationSchemeNegotiateNtlm(d, v, "negotiate_ntlm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["negotiate-ntlm"] = t
		}
	}

	if v, ok := d.GetOk("kerberos_keytab"); ok {
		t, err := expandAuthenticationSchemeKerberosKeytab(d, v, "kerberos_keytab", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["kerberos-keytab"] = t
		}
	} else if d.HasChange("kerberos_keytab") {
		obj["kerberos-keytab"] = nil
	}

	if v, ok := d.GetOk("domain_controller"); ok {
		t, err := expandAuthenticationSchemeDomainController(d, v, "domain_controller", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-controller"] = t
		}
	} else if d.HasChange("domain_controller") {
		obj["domain-controller"] = nil
	}

	if v, ok := d.GetOk("saml_server"); ok {
		t, err := expandAuthenticationSchemeSamlServer(d, v, "saml_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saml-server"] = t
		}
	} else if d.HasChange("saml_server") {
		obj["saml-server"] = nil
	}

	if v, ok := d.GetOk("saml_timeout"); ok {
		t, err := expandAuthenticationSchemeSamlTimeout(d, v, "saml_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saml-timeout"] = t
		}
	}

	if v, ok := d.GetOk("fsso_agent_for_ntlm"); ok {
		t, err := expandAuthenticationSchemeFssoAgentForNtlm(d, v, "fsso_agent_for_ntlm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-agent-for-ntlm"] = t
		}
	} else if d.HasChange("fsso_agent_for_ntlm") {
		obj["fsso-agent-for-ntlm"] = nil
	}

	if v, ok := d.GetOk("require_tfa"); ok {
		t, err := expandAuthenticationSchemeRequireTfa(d, v, "require_tfa", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["require-tfa"] = t
		}
	}

	if v, ok := d.GetOk("fsso_guest"); ok {
		t, err := expandAuthenticationSchemeFssoGuest(d, v, "fsso_guest", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-guest"] = t
		}
	}

	if v, ok := d.GetOk("user_cert"); ok {
		t, err := expandAuthenticationSchemeUserCert(d, v, "user_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-cert"] = t
		}
	}

	if v, ok := d.GetOk("user_database"); ok || d.HasChange("user_database") {
		t, err := expandAuthenticationSchemeUserDatabase(d, v, "user_database", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-database"] = t
		}
	}

	if v, ok := d.GetOk("ssh_ca"); ok {
		t, err := expandAuthenticationSchemeSshCa(d, v, "ssh_ca", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-ca"] = t
		}
	} else if d.HasChange("ssh_ca") {
		obj["ssh-ca"] = nil
	}

	if v, ok := d.GetOk("external_idp"); ok {
		t, err := expandAuthenticationSchemeExternalIdp(d, v, "external_idp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-idp"] = t
		}
	} else if d.HasChange("external_idp") {
		obj["external-idp"] = nil
	}

	return &obj, nil
}
