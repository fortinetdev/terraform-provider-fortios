// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: SAML server entry configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceUserSaml() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserSamlCreate,
		Read:   resourceUserSamlRead,
		Update: resourceUserSamlUpdate,
		Delete: resourceUserSamlDelete,

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
			"cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"entity_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Required:     true,
			},
			"single_sign_on_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Required:     true,
			},
			"single_logout_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"idp_entity_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Required:     true,
			},
			"idp_single_sign_on_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Required:     true,
			},
			"idp_single_logout_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"idp_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"user_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"group_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"digest_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"limit_relaystate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clock_tolerance": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 300),
				Optional:     true,
				Computed:     true,
			},
			"adfs_claim": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_claim_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_claim_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceUserSamlCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserSaml(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserSaml resource while getting object: %v", err)
	}

	o, err := c.CreateUserSaml(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating UserSaml resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserSaml")
	}

	return resourceUserSamlRead(d, m)
}

func resourceUserSamlUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserSaml(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserSaml resource while getting object: %v", err)
	}

	o, err := c.UpdateUserSaml(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserSaml resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserSaml")
	}

	return resourceUserSamlRead(d, m)
}

func resourceUserSamlDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUserSaml(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UserSaml resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserSamlRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadUserSaml(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserSaml resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserSaml(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserSaml resource from API: %v", err)
	}
	return nil
}

func flattenUserSamlName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlEntityId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlIdpEntityId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlIdpSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlIdpSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlIdpCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlUserName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlDigestMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlLimitRelaystate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlClockTolerance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlAdfsClaim(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlUserClaimType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlGroupClaimType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserSaml(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenUserSamlName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("cert", flattenUserSamlCert(o["cert"], d, "cert", sv)); err != nil {
		if !fortiAPIPatch(o["cert"]) {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("entity_id", flattenUserSamlEntityId(o["entity-id"], d, "entity_id", sv)); err != nil {
		if !fortiAPIPatch(o["entity-id"]) {
			return fmt.Errorf("Error reading entity_id: %v", err)
		}
	}

	if err = d.Set("single_sign_on_url", flattenUserSamlSingleSignOnUrl(o["single-sign-on-url"], d, "single_sign_on_url", sv)); err != nil {
		if !fortiAPIPatch(o["single-sign-on-url"]) {
			return fmt.Errorf("Error reading single_sign_on_url: %v", err)
		}
	}

	if err = d.Set("single_logout_url", flattenUserSamlSingleLogoutUrl(o["single-logout-url"], d, "single_logout_url", sv)); err != nil {
		if !fortiAPIPatch(o["single-logout-url"]) {
			return fmt.Errorf("Error reading single_logout_url: %v", err)
		}
	}

	if err = d.Set("idp_entity_id", flattenUserSamlIdpEntityId(o["idp-entity-id"], d, "idp_entity_id", sv)); err != nil {
		if !fortiAPIPatch(o["idp-entity-id"]) {
			return fmt.Errorf("Error reading idp_entity_id: %v", err)
		}
	}

	if err = d.Set("idp_single_sign_on_url", flattenUserSamlIdpSingleSignOnUrl(o["idp-single-sign-on-url"], d, "idp_single_sign_on_url", sv)); err != nil {
		if !fortiAPIPatch(o["idp-single-sign-on-url"]) {
			return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
		}
	}

	if err = d.Set("idp_single_logout_url", flattenUserSamlIdpSingleLogoutUrl(o["idp-single-logout-url"], d, "idp_single_logout_url", sv)); err != nil {
		if !fortiAPIPatch(o["idp-single-logout-url"]) {
			return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
		}
	}

	if err = d.Set("idp_cert", flattenUserSamlIdpCert(o["idp-cert"], d, "idp_cert", sv)); err != nil {
		if !fortiAPIPatch(o["idp-cert"]) {
			return fmt.Errorf("Error reading idp_cert: %v", err)
		}
	}

	if err = d.Set("user_name", flattenUserSamlUserName(o["user-name"], d, "user_name", sv)); err != nil {
		if !fortiAPIPatch(o["user-name"]) {
			return fmt.Errorf("Error reading user_name: %v", err)
		}
	}

	if err = d.Set("group_name", flattenUserSamlGroupName(o["group-name"], d, "group_name", sv)); err != nil {
		if !fortiAPIPatch(o["group-name"]) {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("digest_method", flattenUserSamlDigestMethod(o["digest-method"], d, "digest_method", sv)); err != nil {
		if !fortiAPIPatch(o["digest-method"]) {
			return fmt.Errorf("Error reading digest_method: %v", err)
		}
	}

	if err = d.Set("limit_relaystate", flattenUserSamlLimitRelaystate(o["limit-relaystate"], d, "limit_relaystate", sv)); err != nil {
		if !fortiAPIPatch(o["limit-relaystate"]) {
			return fmt.Errorf("Error reading limit_relaystate: %v", err)
		}
	}

	if err = d.Set("clock_tolerance", flattenUserSamlClockTolerance(o["clock-tolerance"], d, "clock_tolerance", sv)); err != nil {
		if !fortiAPIPatch(o["clock-tolerance"]) {
			return fmt.Errorf("Error reading clock_tolerance: %v", err)
		}
	}

	if err = d.Set("adfs_claim", flattenUserSamlAdfsClaim(o["adfs-claim"], d, "adfs_claim", sv)); err != nil {
		if !fortiAPIPatch(o["adfs-claim"]) {
			return fmt.Errorf("Error reading adfs_claim: %v", err)
		}
	}

	if err = d.Set("user_claim_type", flattenUserSamlUserClaimType(o["user-claim-type"], d, "user_claim_type", sv)); err != nil {
		if !fortiAPIPatch(o["user-claim-type"]) {
			return fmt.Errorf("Error reading user_claim_type: %v", err)
		}
	}

	if err = d.Set("group_claim_type", flattenUserSamlGroupClaimType(o["group-claim-type"], d, "group_claim_type", sv)); err != nil {
		if !fortiAPIPatch(o["group-claim-type"]) {
			return fmt.Errorf("Error reading group_claim_type: %v", err)
		}
	}

	return nil
}

func flattenUserSamlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserSamlName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlEntityId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlIdpEntityId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlIdpSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlIdpSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlIdpCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlUserName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDigestMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlLimitRelaystate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlClockTolerance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlAdfsClaim(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlUserClaimType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlGroupClaimType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserSaml(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandUserSamlName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("cert"); ok {

		t, err := expandUserSamlCert(d, v, "cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert"] = t
		}
	}

	if v, ok := d.GetOk("entity_id"); ok {

		t, err := expandUserSamlEntityId(d, v, "entity_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entity-id"] = t
		}
	}

	if v, ok := d.GetOk("single_sign_on_url"); ok {

		t, err := expandUserSamlSingleSignOnUrl(d, v, "single_sign_on_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["single-sign-on-url"] = t
		}
	}

	if v, ok := d.GetOk("single_logout_url"); ok {

		t, err := expandUserSamlSingleLogoutUrl(d, v, "single_logout_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["single-logout-url"] = t
		}
	}

	if v, ok := d.GetOk("idp_entity_id"); ok {

		t, err := expandUserSamlIdpEntityId(d, v, "idp_entity_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-entity-id"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_sign_on_url"); ok {

		t, err := expandUserSamlIdpSingleSignOnUrl(d, v, "idp_single_sign_on_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-sign-on-url"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_logout_url"); ok {

		t, err := expandUserSamlIdpSingleLogoutUrl(d, v, "idp_single_logout_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-logout-url"] = t
		}
	}

	if v, ok := d.GetOk("idp_cert"); ok {

		t, err := expandUserSamlIdpCert(d, v, "idp_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-cert"] = t
		}
	}

	if v, ok := d.GetOk("user_name"); ok {

		t, err := expandUserSamlUserName(d, v, "user_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-name"] = t
		}
	}

	if v, ok := d.GetOk("group_name"); ok {

		t, err := expandUserSamlGroupName(d, v, "group_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}

	if v, ok := d.GetOk("digest_method"); ok {

		t, err := expandUserSamlDigestMethod(d, v, "digest_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["digest-method"] = t
		}
	}

	if v, ok := d.GetOk("limit_relaystate"); ok {

		t, err := expandUserSamlLimitRelaystate(d, v, "limit_relaystate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["limit-relaystate"] = t
		}
	}

	if v, ok := d.GetOkExists("clock_tolerance"); ok {

		t, err := expandUserSamlClockTolerance(d, v, "clock_tolerance", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clock-tolerance"] = t
		}
	}

	if v, ok := d.GetOk("adfs_claim"); ok {

		t, err := expandUserSamlAdfsClaim(d, v, "adfs_claim", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adfs-claim"] = t
		}
	}

	if v, ok := d.GetOk("user_claim_type"); ok {

		t, err := expandUserSamlUserClaimType(d, v, "user_claim_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-claim-type"] = t
		}
	}

	if v, ok := d.GetOk("group_claim_type"); ok {

		t, err := expandUserSamlGroupClaimType(d, v, "group_claim_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-claim-type"] = t
		}
	}

	return &obj, nil
}
