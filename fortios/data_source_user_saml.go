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

func dataSourceUserSaml() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUserSamlRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"single_sign_on_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"single_logout_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"idp_entity_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"idp_single_sign_on_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"idp_single_logout_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"idp_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"scim_client": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"scim_user_attr_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"scim_group_attr_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"digest_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"require_signed_resp_and_asrt": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"limit_relaystate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"clock_tolerance": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"auth_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"adfs_claim": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_claim_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_claim_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"reauth": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceUserSamlRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing UserSaml: type error")
	}

	o, err := c.ReadUserSaml(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing UserSaml: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectUserSaml(d, o)
	if err != nil {
		return fmt.Errorf("Error describing UserSaml from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenUserSamlName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlIdpEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlIdpSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlIdpSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlIdpCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlScimClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlScimUserAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlScimGroupAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlDigestMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlRequireSignedRespAndAsrt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlLimitRelaystate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlClockTolerance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlAuthUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlAdfsClaim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlUserClaimType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlGroupClaimType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenUserSamlReauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectUserSaml(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenUserSamlName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("cert", dataSourceFlattenUserSamlCert(o["cert"], d, "cert")); err != nil {
		if !fortiAPIPatch(o["cert"]) {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("entity_id", dataSourceFlattenUserSamlEntityId(o["entity-id"], d, "entity_id")); err != nil {
		if !fortiAPIPatch(o["entity-id"]) {
			return fmt.Errorf("Error reading entity_id: %v", err)
		}
	}

	if err = d.Set("single_sign_on_url", dataSourceFlattenUserSamlSingleSignOnUrl(o["single-sign-on-url"], d, "single_sign_on_url")); err != nil {
		if !fortiAPIPatch(o["single-sign-on-url"]) {
			return fmt.Errorf("Error reading single_sign_on_url: %v", err)
		}
	}

	if err = d.Set("single_logout_url", dataSourceFlattenUserSamlSingleLogoutUrl(o["single-logout-url"], d, "single_logout_url")); err != nil {
		if !fortiAPIPatch(o["single-logout-url"]) {
			return fmt.Errorf("Error reading single_logout_url: %v", err)
		}
	}

	if err = d.Set("idp_entity_id", dataSourceFlattenUserSamlIdpEntityId(o["idp-entity-id"], d, "idp_entity_id")); err != nil {
		if !fortiAPIPatch(o["idp-entity-id"]) {
			return fmt.Errorf("Error reading idp_entity_id: %v", err)
		}
	}

	if err = d.Set("idp_single_sign_on_url", dataSourceFlattenUserSamlIdpSingleSignOnUrl(o["idp-single-sign-on-url"], d, "idp_single_sign_on_url")); err != nil {
		if !fortiAPIPatch(o["idp-single-sign-on-url"]) {
			return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
		}
	}

	if err = d.Set("idp_single_logout_url", dataSourceFlattenUserSamlIdpSingleLogoutUrl(o["idp-single-logout-url"], d, "idp_single_logout_url")); err != nil {
		if !fortiAPIPatch(o["idp-single-logout-url"]) {
			return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
		}
	}

	if err = d.Set("idp_cert", dataSourceFlattenUserSamlIdpCert(o["idp-cert"], d, "idp_cert")); err != nil {
		if !fortiAPIPatch(o["idp-cert"]) {
			return fmt.Errorf("Error reading idp_cert: %v", err)
		}
	}

	if err = d.Set("scim_client", dataSourceFlattenUserSamlScimClient(o["scim-client"], d, "scim_client")); err != nil {
		if !fortiAPIPatch(o["scim-client"]) {
			return fmt.Errorf("Error reading scim_client: %v", err)
		}
	}

	if err = d.Set("scim_user_attr_type", dataSourceFlattenUserSamlScimUserAttrType(o["scim-user-attr-type"], d, "scim_user_attr_type")); err != nil {
		if !fortiAPIPatch(o["scim-user-attr-type"]) {
			return fmt.Errorf("Error reading scim_user_attr_type: %v", err)
		}
	}

	if err = d.Set("scim_group_attr_type", dataSourceFlattenUserSamlScimGroupAttrType(o["scim-group-attr-type"], d, "scim_group_attr_type")); err != nil {
		if !fortiAPIPatch(o["scim-group-attr-type"]) {
			return fmt.Errorf("Error reading scim_group_attr_type: %v", err)
		}
	}

	if err = d.Set("user_name", dataSourceFlattenUserSamlUserName(o["user-name"], d, "user_name")); err != nil {
		if !fortiAPIPatch(o["user-name"]) {
			return fmt.Errorf("Error reading user_name: %v", err)
		}
	}

	if err = d.Set("group_name", dataSourceFlattenUserSamlGroupName(o["group-name"], d, "group_name")); err != nil {
		if !fortiAPIPatch(o["group-name"]) {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("digest_method", dataSourceFlattenUserSamlDigestMethod(o["digest-method"], d, "digest_method")); err != nil {
		if !fortiAPIPatch(o["digest-method"]) {
			return fmt.Errorf("Error reading digest_method: %v", err)
		}
	}

	if err = d.Set("require_signed_resp_and_asrt", dataSourceFlattenUserSamlRequireSignedRespAndAsrt(o["require-signed-resp-and-asrt"], d, "require_signed_resp_and_asrt")); err != nil {
		if !fortiAPIPatch(o["require-signed-resp-and-asrt"]) {
			return fmt.Errorf("Error reading require_signed_resp_and_asrt: %v", err)
		}
	}

	if err = d.Set("limit_relaystate", dataSourceFlattenUserSamlLimitRelaystate(o["limit-relaystate"], d, "limit_relaystate")); err != nil {
		if !fortiAPIPatch(o["limit-relaystate"]) {
			return fmt.Errorf("Error reading limit_relaystate: %v", err)
		}
	}

	if err = d.Set("clock_tolerance", dataSourceFlattenUserSamlClockTolerance(o["clock-tolerance"], d, "clock_tolerance")); err != nil {
		if !fortiAPIPatch(o["clock-tolerance"]) {
			return fmt.Errorf("Error reading clock_tolerance: %v", err)
		}
	}

	if err = d.Set("auth_url", dataSourceFlattenUserSamlAuthUrl(o["auth-url"], d, "auth_url")); err != nil {
		if !fortiAPIPatch(o["auth-url"]) {
			return fmt.Errorf("Error reading auth_url: %v", err)
		}
	}

	if err = d.Set("adfs_claim", dataSourceFlattenUserSamlAdfsClaim(o["adfs-claim"], d, "adfs_claim")); err != nil {
		if !fortiAPIPatch(o["adfs-claim"]) {
			return fmt.Errorf("Error reading adfs_claim: %v", err)
		}
	}

	if err = d.Set("user_claim_type", dataSourceFlattenUserSamlUserClaimType(o["user-claim-type"], d, "user_claim_type")); err != nil {
		if !fortiAPIPatch(o["user-claim-type"]) {
			return fmt.Errorf("Error reading user_claim_type: %v", err)
		}
	}

	if err = d.Set("group_claim_type", dataSourceFlattenUserSamlGroupClaimType(o["group-claim-type"], d, "group_claim_type")); err != nil {
		if !fortiAPIPatch(o["group-claim-type"]) {
			return fmt.Errorf("Error reading group_claim_type: %v", err)
		}
	}

	if err = d.Set("reauth", dataSourceFlattenUserSamlReauth(o["reauth"], d, "reauth")); err != nil {
		if !fortiAPIPatch(o["reauth"]) {
			return fmt.Errorf("Error reading reauth: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenUserSamlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
