// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FSSO active directory servers for polling mode.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceUserFssoPolling() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserFssoPollingCreate,
		Read:   resourceUserFssoPollingRead,
		Update: resourceUserFssoPollingUpdate,
		Delete: resourceUserFssoPollingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required:     true,
			},
			"default_domain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"user": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"ldap_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"logon_history": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 48),
				Optional:     true,
				Computed:     true,
			},
			"polling_frequency": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),
				Optional:     true,
				Computed:     true,
			},
			"adgrp": &schema.Schema{
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
			"smbv1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"smb_ntlmv1_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceUserFssoPollingCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserFssoPolling(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserFssoPolling resource while getting object: %v", err)
	}

	o, err := c.CreateUserFssoPolling(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating UserFssoPolling resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("UserFssoPolling")
	}

	return resourceUserFssoPollingRead(d, m)
}

func resourceUserFssoPollingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserFssoPolling(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserFssoPolling resource while getting object: %v", err)
	}

	o, err := c.UpdateUserFssoPolling(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserFssoPolling resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("UserFssoPolling")
	}

	return resourceUserFssoPollingRead(d, m)
}

func resourceUserFssoPollingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUserFssoPolling(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UserFssoPolling resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserFssoPollingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadUserFssoPolling(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserFssoPolling resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserFssoPolling(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserFssoPolling resource from API: %v", err)
	}
	return nil
}

func flattenUserFssoPollingId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoPollingStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoPollingServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoPollingDefaultDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoPollingPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoPollingUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoPollingPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoPollingLdapServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoPollingLogonHistory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoPollingPollingFrequency(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoPollingAdgrp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenUserFssoPollingAdgrpName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenUserFssoPollingAdgrpName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoPollingSmbv1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoPollingSmbNtlmv1Auth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserFssoPolling(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenUserFssoPollingId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("status", flattenUserFssoPollingStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("server", flattenUserFssoPollingServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("default_domain", flattenUserFssoPollingDefaultDomain(o["default-domain"], d, "default_domain", sv)); err != nil {
		if !fortiAPIPatch(o["default-domain"]) {
			return fmt.Errorf("Error reading default_domain: %v", err)
		}
	}

	if err = d.Set("port", flattenUserFssoPollingPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("user", flattenUserFssoPollingUser(o["user"], d, "user", sv)); err != nil {
		if !fortiAPIPatch(o["user"]) {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenUserFssoPollingLdapServer(o["ldap-server"], d, "ldap_server", sv)); err != nil {
		if !fortiAPIPatch(o["ldap-server"]) {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("logon_history", flattenUserFssoPollingLogonHistory(o["logon-history"], d, "logon_history", sv)); err != nil {
		if !fortiAPIPatch(o["logon-history"]) {
			return fmt.Errorf("Error reading logon_history: %v", err)
		}
	}

	if err = d.Set("polling_frequency", flattenUserFssoPollingPollingFrequency(o["polling-frequency"], d, "polling_frequency", sv)); err != nil {
		if !fortiAPIPatch(o["polling-frequency"]) {
			return fmt.Errorf("Error reading polling_frequency: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("adgrp", flattenUserFssoPollingAdgrp(o["adgrp"], d, "adgrp", sv)); err != nil {
			if !fortiAPIPatch(o["adgrp"]) {
				return fmt.Errorf("Error reading adgrp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("adgrp"); ok {
			if err = d.Set("adgrp", flattenUserFssoPollingAdgrp(o["adgrp"], d, "adgrp", sv)); err != nil {
				if !fortiAPIPatch(o["adgrp"]) {
					return fmt.Errorf("Error reading adgrp: %v", err)
				}
			}
		}
	}

	if err = d.Set("smbv1", flattenUserFssoPollingSmbv1(o["smbv1"], d, "smbv1", sv)); err != nil {
		if !fortiAPIPatch(o["smbv1"]) {
			return fmt.Errorf("Error reading smbv1: %v", err)
		}
	}

	if err = d.Set("smb_ntlmv1_auth", flattenUserFssoPollingSmbNtlmv1Auth(o["smb-ntlmv1-auth"], d, "smb_ntlmv1_auth", sv)); err != nil {
		if !fortiAPIPatch(o["smb-ntlmv1-auth"]) {
			return fmt.Errorf("Error reading smb_ntlmv1_auth: %v", err)
		}
	}

	return nil
}

func flattenUserFssoPollingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserFssoPollingId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingDefaultDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingLdapServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingLogonHistory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingPollingFrequency(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingAdgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandUserFssoPollingAdgrpName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserFssoPollingAdgrpName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingSmbv1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPollingSmbNtlmv1Auth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserFssoPolling(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandUserFssoPollingId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandUserFssoPollingStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {

		t, err := expandUserFssoPollingServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("default_domain"); ok {

		t, err := expandUserFssoPollingDefaultDomain(d, v, "default_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-domain"] = t
		}
	}

	if v, ok := d.GetOkExists("port"); ok {

		t, err := expandUserFssoPollingPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok {

		t, err := expandUserFssoPollingUser(d, v, "user", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {

		t, err := expandUserFssoPollingPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok {

		t, err := expandUserFssoPollingLdapServer(d, v, "ldap_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOkExists("logon_history"); ok {

		t, err := expandUserFssoPollingLogonHistory(d, v, "logon_history", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logon-history"] = t
		}
	}

	if v, ok := d.GetOk("polling_frequency"); ok {

		t, err := expandUserFssoPollingPollingFrequency(d, v, "polling_frequency", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polling-frequency"] = t
		}
	}

	if v, ok := d.GetOk("adgrp"); ok || d.HasChange("adgrp") {

		t, err := expandUserFssoPollingAdgrp(d, v, "adgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adgrp"] = t
		}
	}

	if v, ok := d.GetOk("smbv1"); ok {

		t, err := expandUserFssoPollingSmbv1(d, v, "smbv1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smbv1"] = t
		}
	}

	if v, ok := d.GetOk("smb_ntlmv1_auth"); ok {

		t, err := expandUserFssoPollingSmbNtlmv1Auth(d, v, "smb_ntlmv1_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smb-ntlmv1-auth"] = t
		}
	}

	return &obj, nil
}
