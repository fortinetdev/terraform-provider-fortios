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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_force_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_object_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_provider_address": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
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
			},
			"idp_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"scim_client": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"scim_user_attr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scim_group_attr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"group_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"digest_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"require_signed_resp_and_asrt": &schema.Schema{
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
			"realm": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"user_source": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"auth_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
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
			"reauth": &schema.Schema{
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

	obj, err := getObjectUserSaml(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserSaml resource while getting object: %v", err)
	}

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserSaml(mkey, vdomparam)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserSaml(obj, mkey, vdomparam)
			if err != nil {
				return fmt.Errorf("Error updating UserSaml resource: %v", err)
			}
		}
	}

	if !existing {
		o, err = c.CreateUserSaml(obj, vdomparam)

		if err != nil {
			return fmt.Errorf("Error creating UserSaml resource: %v", err)
		}
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

func flattenUserSamlUuid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlFabricObject(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlFabricForceSync(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlFabricObjectSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlServiceProviderAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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

func flattenUserSamlScimClient(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlScimUserAttrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlScimGroupAttrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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

func flattenUserSamlRequireSignedRespAndAsrt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlLimitRelaystate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlClockTolerance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserSamlRealm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlUserSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSamlAuthUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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

func flattenUserSamlReauth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserSaml(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenUserSamlName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", flattenUserSamlUuid(o["uuid"], d, "uuid", sv)); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenUserSamlFabricObject(o["fabric-object"], d, "fabric_object", sv)); err != nil {
		if !fortiAPIPatch(o["fabric-object"]) {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("fabric_force_sync", flattenUserSamlFabricForceSync(o["fabric-force-sync"], d, "fabric_force_sync", sv)); err != nil {
		if !fortiAPIPatch(o["fabric-force-sync"]) {
			return fmt.Errorf("Error reading fabric_force_sync: %v", err)
		}
	}

	if err = d.Set("fabric_object_source", flattenUserSamlFabricObjectSource(o["fabric-object-source"], d, "fabric_object_source", sv)); err != nil {
		if !fortiAPIPatch(o["fabric-object-source"]) {
			return fmt.Errorf("Error reading fabric_object_source: %v", err)
		}
	}

	if err = d.Set("cert", flattenUserSamlCert(o["cert"], d, "cert", sv)); err != nil {
		if !fortiAPIPatch(o["cert"]) {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("type", flattenUserSamlType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("service_provider_address", flattenUserSamlServiceProviderAddress(o["service-provider-address"], d, "service_provider_address", sv)); err != nil {
		if !fortiAPIPatch(o["service-provider-address"]) {
			return fmt.Errorf("Error reading service_provider_address: %v", err)
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

	if err = d.Set("scim_client", flattenUserSamlScimClient(o["scim-client"], d, "scim_client", sv)); err != nil {
		if !fortiAPIPatch(o["scim-client"]) {
			return fmt.Errorf("Error reading scim_client: %v", err)
		}
	}

	if err = d.Set("scim_user_attr_type", flattenUserSamlScimUserAttrType(o["scim-user-attr-type"], d, "scim_user_attr_type", sv)); err != nil {
		if !fortiAPIPatch(o["scim-user-attr-type"]) {
			return fmt.Errorf("Error reading scim_user_attr_type: %v", err)
		}
	}

	if err = d.Set("scim_group_attr_type", flattenUserSamlScimGroupAttrType(o["scim-group-attr-type"], d, "scim_group_attr_type", sv)); err != nil {
		if !fortiAPIPatch(o["scim-group-attr-type"]) {
			return fmt.Errorf("Error reading scim_group_attr_type: %v", err)
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

	if err = d.Set("require_signed_resp_and_asrt", flattenUserSamlRequireSignedRespAndAsrt(o["require-signed-resp-and-asrt"], d, "require_signed_resp_and_asrt", sv)); err != nil {
		if !fortiAPIPatch(o["require-signed-resp-and-asrt"]) {
			return fmt.Errorf("Error reading require_signed_resp_and_asrt: %v", err)
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

	if err = d.Set("realm", flattenUserSamlRealm(o["realm"], d, "realm", sv)); err != nil {
		if !fortiAPIPatch(o["realm"]) {
			return fmt.Errorf("Error reading realm: %v", err)
		}
	}

	if err = d.Set("user_source", flattenUserSamlUserSource(o["user-source"], d, "user_source", sv)); err != nil {
		if !fortiAPIPatch(o["user-source"]) {
			return fmt.Errorf("Error reading user_source: %v", err)
		}
	}

	if err = d.Set("auth_url", flattenUserSamlAuthUrl(o["auth-url"], d, "auth_url", sv)); err != nil {
		if !fortiAPIPatch(o["auth-url"]) {
			return fmt.Errorf("Error reading auth_url: %v", err)
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

	if err = d.Set("reauth", flattenUserSamlReauth(o["reauth"], d, "reauth", sv)); err != nil {
		if !fortiAPIPatch(o["reauth"]) {
			return fmt.Errorf("Error reading reauth: %v", err)
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

func expandUserSamlUuid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlFabricObject(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlFabricForceSync(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlFabricObjectSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlServiceProviderAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

func expandUserSamlScimClient(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlScimUserAttrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlScimGroupAttrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

func expandUserSamlRequireSignedRespAndAsrt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlLimitRelaystate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlClockTolerance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlRealm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlUserSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSamlAuthUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

func expandUserSamlReauth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandUserSamlUuid(d, v, "uuid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok {
		t, err := expandUserSamlFabricObject(d, v, "fabric_object", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("fabric_force_sync"); ok {
		t, err := expandUserSamlFabricForceSync(d, v, "fabric_force_sync", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-force-sync"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object_source"); ok {
		t, err := expandUserSamlFabricObjectSource(d, v, "fabric_object_source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object-source"] = t
		}
	}

	if v, ok := d.GetOk("cert"); ok {
		t, err := expandUserSamlCert(d, v, "cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert"] = t
		}
	} else if d.HasChange("cert") {
		obj["cert"] = nil
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandUserSamlType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("service_provider_address"); ok {
		t, err := expandUserSamlServiceProviderAddress(d, v, "service_provider_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-provider-address"] = t
		}
	} else if d.HasChange("service_provider_address") {
		obj["service-provider-address"] = nil
	}

	if v, ok := d.GetOk("entity_id"); ok {
		t, err := expandUserSamlEntityId(d, v, "entity_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entity-id"] = t
		}
	} else if d.HasChange("entity_id") {
		obj["entity-id"] = nil
	}

	if v, ok := d.GetOk("single_sign_on_url"); ok {
		t, err := expandUserSamlSingleSignOnUrl(d, v, "single_sign_on_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["single-sign-on-url"] = t
		}
	} else if d.HasChange("single_sign_on_url") {
		obj["single-sign-on-url"] = nil
	}

	if v, ok := d.GetOk("single_logout_url"); ok {
		t, err := expandUserSamlSingleLogoutUrl(d, v, "single_logout_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["single-logout-url"] = t
		}
	} else if d.HasChange("single_logout_url") {
		obj["single-logout-url"] = nil
	}

	if v, ok := d.GetOk("idp_entity_id"); ok {
		t, err := expandUserSamlIdpEntityId(d, v, "idp_entity_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-entity-id"] = t
		}
	} else if d.HasChange("idp_entity_id") {
		obj["idp-entity-id"] = nil
	}

	if v, ok := d.GetOk("idp_single_sign_on_url"); ok {
		t, err := expandUserSamlIdpSingleSignOnUrl(d, v, "idp_single_sign_on_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-sign-on-url"] = t
		}
	} else if d.HasChange("idp_single_sign_on_url") {
		obj["idp-single-sign-on-url"] = nil
	}

	if v, ok := d.GetOk("idp_single_logout_url"); ok {
		t, err := expandUserSamlIdpSingleLogoutUrl(d, v, "idp_single_logout_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-logout-url"] = t
		}
	} else if d.HasChange("idp_single_logout_url") {
		obj["idp-single-logout-url"] = nil
	}

	if v, ok := d.GetOk("idp_cert"); ok {
		t, err := expandUserSamlIdpCert(d, v, "idp_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-cert"] = t
		}
	} else if d.HasChange("idp_cert") {
		obj["idp-cert"] = nil
	}

	if v, ok := d.GetOk("scim_client"); ok {
		t, err := expandUserSamlScimClient(d, v, "scim_client", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scim-client"] = t
		}
	} else if d.HasChange("scim_client") {
		obj["scim-client"] = nil
	}

	if v, ok := d.GetOk("scim_user_attr_type"); ok {
		t, err := expandUserSamlScimUserAttrType(d, v, "scim_user_attr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scim-user-attr-type"] = t
		}
	}

	if v, ok := d.GetOk("scim_group_attr_type"); ok {
		t, err := expandUserSamlScimGroupAttrType(d, v, "scim_group_attr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scim-group-attr-type"] = t
		}
	}

	if v, ok := d.GetOk("user_name"); ok {
		t, err := expandUserSamlUserName(d, v, "user_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-name"] = t
		}
	} else if d.HasChange("user_name") {
		obj["user-name"] = nil
	}

	if v, ok := d.GetOk("group_name"); ok {
		t, err := expandUserSamlGroupName(d, v, "group_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	} else if d.HasChange("group_name") {
		obj["group-name"] = nil
	}

	if v, ok := d.GetOk("digest_method"); ok {
		t, err := expandUserSamlDigestMethod(d, v, "digest_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["digest-method"] = t
		}
	}

	if v, ok := d.GetOk("require_signed_resp_and_asrt"); ok {
		t, err := expandUserSamlRequireSignedRespAndAsrt(d, v, "require_signed_resp_and_asrt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["require-signed-resp-and-asrt"] = t
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

	if v, ok := d.GetOk("realm"); ok {
		t, err := expandUserSamlRealm(d, v, "realm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realm"] = t
		}
	}

	if v, ok := d.GetOk("user_source"); ok {
		t, err := expandUserSamlUserSource(d, v, "user_source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-source"] = t
		}
	} else if d.HasChange("user_source") {
		obj["user-source"] = nil
	}

	if v, ok := d.GetOk("auth_url"); ok {
		t, err := expandUserSamlAuthUrl(d, v, "auth_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-url"] = t
		}
	} else if d.HasChange("auth_url") {
		obj["auth-url"] = nil
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

	if v, ok := d.GetOk("reauth"); ok {
		t, err := expandUserSamlReauth(d, v, "reauth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reauth"] = t
		}
	}

	return &obj, nil
}
