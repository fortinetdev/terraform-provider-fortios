// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure Authentication Schemes.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
				Computed:     true,
			},
			"domain_controller": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"fsso_agent_for_ntlm": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
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
			"user_database": &schema.Schema{
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
			"ssh_ca": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceAuthenticationSchemeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectAuthenticationScheme(d)
	if err != nil {
		return fmt.Errorf("Error creating AuthenticationScheme resource while getting object: %v", err)
	}

	o, err := c.CreateAuthenticationScheme(obj)

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

	obj, err := getObjectAuthenticationScheme(d)
	if err != nil {
		return fmt.Errorf("Error updating AuthenticationScheme resource while getting object: %v", err)
	}

	o, err := c.UpdateAuthenticationScheme(obj, mkey)
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

	err := c.DeleteAuthenticationScheme(mkey)
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

	o, err := c.ReadAuthenticationScheme(mkey)
	if err != nil {
		return fmt.Errorf("Error reading AuthenticationScheme resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAuthenticationScheme(d, o)
	if err != nil {
		return fmt.Errorf("Error reading AuthenticationScheme resource from API: %v", err)
	}
	return nil
}

func flattenAuthenticationSchemeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSchemeMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSchemeNegotiateNtlm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSchemeKerberosKeytab(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSchemeDomainController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSchemeFssoAgentForNtlm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSchemeRequireTfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSchemeFssoGuest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSchemeUserDatabase(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenAuthenticationSchemeUserDatabaseName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenAuthenticationSchemeUserDatabaseName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSchemeSshCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectAuthenticationScheme(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenAuthenticationSchemeName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("method", flattenAuthenticationSchemeMethod(o["method"], d, "method")); err != nil {
		if !fortiAPIPatch(o["method"]) {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("negotiate_ntlm", flattenAuthenticationSchemeNegotiateNtlm(o["negotiate-ntlm"], d, "negotiate_ntlm")); err != nil {
		if !fortiAPIPatch(o["negotiate-ntlm"]) {
			return fmt.Errorf("Error reading negotiate_ntlm: %v", err)
		}
	}

	if err = d.Set("kerberos_keytab", flattenAuthenticationSchemeKerberosKeytab(o["kerberos-keytab"], d, "kerberos_keytab")); err != nil {
		if !fortiAPIPatch(o["kerberos-keytab"]) {
			return fmt.Errorf("Error reading kerberos_keytab: %v", err)
		}
	}

	if err = d.Set("domain_controller", flattenAuthenticationSchemeDomainController(o["domain-controller"], d, "domain_controller")); err != nil {
		if !fortiAPIPatch(o["domain-controller"]) {
			return fmt.Errorf("Error reading domain_controller: %v", err)
		}
	}

	if err = d.Set("fsso_agent_for_ntlm", flattenAuthenticationSchemeFssoAgentForNtlm(o["fsso-agent-for-ntlm"], d, "fsso_agent_for_ntlm")); err != nil {
		if !fortiAPIPatch(o["fsso-agent-for-ntlm"]) {
			return fmt.Errorf("Error reading fsso_agent_for_ntlm: %v", err)
		}
	}

	if err = d.Set("require_tfa", flattenAuthenticationSchemeRequireTfa(o["require-tfa"], d, "require_tfa")); err != nil {
		if !fortiAPIPatch(o["require-tfa"]) {
			return fmt.Errorf("Error reading require_tfa: %v", err)
		}
	}

	if err = d.Set("fsso_guest", flattenAuthenticationSchemeFssoGuest(o["fsso-guest"], d, "fsso_guest")); err != nil {
		if !fortiAPIPatch(o["fsso-guest"]) {
			return fmt.Errorf("Error reading fsso_guest: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("user_database", flattenAuthenticationSchemeUserDatabase(o["user-database"], d, "user_database")); err != nil {
			if !fortiAPIPatch(o["user-database"]) {
				return fmt.Errorf("Error reading user_database: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("user_database"); ok {
			if err = d.Set("user_database", flattenAuthenticationSchemeUserDatabase(o["user-database"], d, "user_database")); err != nil {
				if !fortiAPIPatch(o["user-database"]) {
					return fmt.Errorf("Error reading user_database: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssh_ca", flattenAuthenticationSchemeSshCa(o["ssh-ca"], d, "ssh_ca")); err != nil {
		if !fortiAPIPatch(o["ssh-ca"]) {
			return fmt.Errorf("Error reading ssh_ca: %v", err)
		}
	}

	return nil
}

func flattenAuthenticationSchemeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandAuthenticationSchemeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeNegotiateNtlm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeKerberosKeytab(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeDomainController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeFssoAgentForNtlm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeRequireTfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeFssoGuest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeUserDatabase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandAuthenticationSchemeUserDatabaseName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandAuthenticationSchemeUserDatabaseName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSchemeSshCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectAuthenticationScheme(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandAuthenticationSchemeName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok {
		t, err := expandAuthenticationSchemeMethod(d, v, "method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("negotiate_ntlm"); ok {
		t, err := expandAuthenticationSchemeNegotiateNtlm(d, v, "negotiate_ntlm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["negotiate-ntlm"] = t
		}
	}

	if v, ok := d.GetOk("kerberos_keytab"); ok {
		t, err := expandAuthenticationSchemeKerberosKeytab(d, v, "kerberos_keytab")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["kerberos-keytab"] = t
		}
	}

	if v, ok := d.GetOk("domain_controller"); ok {
		t, err := expandAuthenticationSchemeDomainController(d, v, "domain_controller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-controller"] = t
		}
	}

	if v, ok := d.GetOk("fsso_agent_for_ntlm"); ok {
		t, err := expandAuthenticationSchemeFssoAgentForNtlm(d, v, "fsso_agent_for_ntlm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-agent-for-ntlm"] = t
		}
	}

	if v, ok := d.GetOk("require_tfa"); ok {
		t, err := expandAuthenticationSchemeRequireTfa(d, v, "require_tfa")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["require-tfa"] = t
		}
	}

	if v, ok := d.GetOk("fsso_guest"); ok {
		t, err := expandAuthenticationSchemeFssoGuest(d, v, "fsso_guest")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-guest"] = t
		}
	}

	if v, ok := d.GetOk("user_database"); ok {
		t, err := expandAuthenticationSchemeUserDatabase(d, v, "user_database")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-database"] = t
		}
	}

	if v, ok := d.GetOk("ssh_ca"); ok {
		t, err := expandAuthenticationSchemeSshCa(d, v, "ssh_ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-ca"] = t
		}
	}

	return &obj, nil
}
