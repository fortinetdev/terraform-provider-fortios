// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Global settings for SAML authentication.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemSaml() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSamlUpdate,
		Read:   resourceSystemSamlRead,
		Update: resourceSystemSamlUpdate,
		Delete: resourceSystemSamlDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_login_page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"binding_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"portal_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"entity_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"single_sign_on_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
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
				Optional:     true,
				Computed:     true,
			},
			"idp_single_sign_on_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
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
				Optional:     true,
				Computed:     true,
			},
			"server_address": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"tolerance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"life": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"service_providers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"prefix": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"sp_binding_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sp_cert": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"sp_entity_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"sp_single_sign_on_url": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"sp_single_logout_url": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"sp_portal_url": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"idp_entity_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"idp_single_sign_on_url": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"idp_single_logout_url": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"assertion_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
										Computed:     true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
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

func resourceSystemSamlUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemSaml(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSaml resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSaml(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemSaml resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSaml")
	}

	return resourceSystemSamlRead(d, m)
}

func resourceSystemSamlDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemSaml(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemSaml resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSaml(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemSaml resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSamlRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemSaml(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemSaml resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSaml(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSaml resource from API: %v", err)
	}
	return nil
}

func flattenSystemSamlStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlRole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlDefaultLoginPage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlDefaultProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlBindingProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlPortalUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlEntityId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlIdpEntityId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlIdpSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlIdpSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlIdpCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlServerAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlTolerance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlLife(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlServiceProviders(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSamlServiceProvidersName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {

			tmp["prefix"] = flattenSystemSamlServiceProvidersPrefix(i["prefix"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_binding_protocol"
		if _, ok := i["sp-binding-protocol"]; ok {

			tmp["sp_binding_protocol"] = flattenSystemSamlServiceProvidersSpBindingProtocol(i["sp-binding-protocol"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_cert"
		if _, ok := i["sp-cert"]; ok {

			tmp["sp_cert"] = flattenSystemSamlServiceProvidersSpCert(i["sp-cert"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_entity_id"
		if _, ok := i["sp-entity-id"]; ok {

			tmp["sp_entity_id"] = flattenSystemSamlServiceProvidersSpEntityId(i["sp-entity-id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_single_sign_on_url"
		if _, ok := i["sp-single-sign-on-url"]; ok {

			tmp["sp_single_sign_on_url"] = flattenSystemSamlServiceProvidersSpSingleSignOnUrl(i["sp-single-sign-on-url"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_single_logout_url"
		if _, ok := i["sp-single-logout-url"]; ok {

			tmp["sp_single_logout_url"] = flattenSystemSamlServiceProvidersSpSingleLogoutUrl(i["sp-single-logout-url"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_portal_url"
		if _, ok := i["sp-portal-url"]; ok {

			tmp["sp_portal_url"] = flattenSystemSamlServiceProvidersSpPortalUrl(i["sp-portal-url"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_entity_id"
		if _, ok := i["idp-entity-id"]; ok {

			tmp["idp_entity_id"] = flattenSystemSamlServiceProvidersIdpEntityId(i["idp-entity-id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_sign_on_url"
		if _, ok := i["idp-single-sign-on-url"]; ok {

			tmp["idp_single_sign_on_url"] = flattenSystemSamlServiceProvidersIdpSingleSignOnUrl(i["idp-single-sign-on-url"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_logout_url"
		if _, ok := i["idp-single-logout-url"]; ok {

			tmp["idp_single_logout_url"] = flattenSystemSamlServiceProvidersIdpSingleLogoutUrl(i["idp-single-logout-url"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assertion_attributes"
		if _, ok := i["assertion-attributes"]; ok {

			tmp["assertion_attributes"] = flattenSystemSamlServiceProvidersAssertionAttributes(i["assertion-attributes"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSamlServiceProvidersName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpBindingProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpEntityId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpPortalUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersIdpEntityId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersIdpSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersIdpSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersAssertionAttributes(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSamlServiceProvidersAssertionAttributesName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = flattenSystemSamlServiceProvidersAssertionAttributesType(i["type"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSamlServiceProvidersAssertionAttributesName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersAssertionAttributesType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSaml(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemSamlStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("role", flattenSystemSamlRole(o["role"], d, "role", sv)); err != nil {
		if !fortiAPIPatch(o["role"]) {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("default_login_page", flattenSystemSamlDefaultLoginPage(o["default-login-page"], d, "default_login_page", sv)); err != nil {
		if !fortiAPIPatch(o["default-login-page"]) {
			return fmt.Errorf("Error reading default_login_page: %v", err)
		}
	}

	if err = d.Set("default_profile", flattenSystemSamlDefaultProfile(o["default-profile"], d, "default_profile", sv)); err != nil {
		if !fortiAPIPatch(o["default-profile"]) {
			return fmt.Errorf("Error reading default_profile: %v", err)
		}
	}

	if err = d.Set("cert", flattenSystemSamlCert(o["cert"], d, "cert", sv)); err != nil {
		if !fortiAPIPatch(o["cert"]) {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("binding_protocol", flattenSystemSamlBindingProtocol(o["binding-protocol"], d, "binding_protocol", sv)); err != nil {
		if !fortiAPIPatch(o["binding-protocol"]) {
			return fmt.Errorf("Error reading binding_protocol: %v", err)
		}
	}

	if err = d.Set("portal_url", flattenSystemSamlPortalUrl(o["portal-url"], d, "portal_url", sv)); err != nil {
		if !fortiAPIPatch(o["portal-url"]) {
			return fmt.Errorf("Error reading portal_url: %v", err)
		}
	}

	if err = d.Set("entity_id", flattenSystemSamlEntityId(o["entity-id"], d, "entity_id", sv)); err != nil {
		if !fortiAPIPatch(o["entity-id"]) {
			return fmt.Errorf("Error reading entity_id: %v", err)
		}
	}

	if err = d.Set("single_sign_on_url", flattenSystemSamlSingleSignOnUrl(o["single-sign-on-url"], d, "single_sign_on_url", sv)); err != nil {
		if !fortiAPIPatch(o["single-sign-on-url"]) {
			return fmt.Errorf("Error reading single_sign_on_url: %v", err)
		}
	}

	if err = d.Set("single_logout_url", flattenSystemSamlSingleLogoutUrl(o["single-logout-url"], d, "single_logout_url", sv)); err != nil {
		if !fortiAPIPatch(o["single-logout-url"]) {
			return fmt.Errorf("Error reading single_logout_url: %v", err)
		}
	}

	if err = d.Set("idp_entity_id", flattenSystemSamlIdpEntityId(o["idp-entity-id"], d, "idp_entity_id", sv)); err != nil {
		if !fortiAPIPatch(o["idp-entity-id"]) {
			return fmt.Errorf("Error reading idp_entity_id: %v", err)
		}
	}

	if err = d.Set("idp_single_sign_on_url", flattenSystemSamlIdpSingleSignOnUrl(o["idp-single-sign-on-url"], d, "idp_single_sign_on_url", sv)); err != nil {
		if !fortiAPIPatch(o["idp-single-sign-on-url"]) {
			return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
		}
	}

	if err = d.Set("idp_single_logout_url", flattenSystemSamlIdpSingleLogoutUrl(o["idp-single-logout-url"], d, "idp_single_logout_url", sv)); err != nil {
		if !fortiAPIPatch(o["idp-single-logout-url"]) {
			return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
		}
	}

	if err = d.Set("idp_cert", flattenSystemSamlIdpCert(o["idp-cert"], d, "idp_cert", sv)); err != nil {
		if !fortiAPIPatch(o["idp-cert"]) {
			return fmt.Errorf("Error reading idp_cert: %v", err)
		}
	}

	if err = d.Set("server_address", flattenSystemSamlServerAddress(o["server-address"], d, "server_address", sv)); err != nil {
		if !fortiAPIPatch(o["server-address"]) {
			return fmt.Errorf("Error reading server_address: %v", err)
		}
	}

	if err = d.Set("tolerance", flattenSystemSamlTolerance(o["tolerance"], d, "tolerance", sv)); err != nil {
		if !fortiAPIPatch(o["tolerance"]) {
			return fmt.Errorf("Error reading tolerance: %v", err)
		}
	}

	if err = d.Set("life", flattenSystemSamlLife(o["life"], d, "life", sv)); err != nil {
		if !fortiAPIPatch(o["life"]) {
			return fmt.Errorf("Error reading life: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("service_providers", flattenSystemSamlServiceProviders(o["service-providers"], d, "service_providers", sv)); err != nil {
			if !fortiAPIPatch(o["service-providers"]) {
				return fmt.Errorf("Error reading service_providers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service_providers"); ok {
			if err = d.Set("service_providers", flattenSystemSamlServiceProviders(o["service-providers"], d, "service_providers", sv)); err != nil {
				if !fortiAPIPatch(o["service-providers"]) {
					return fmt.Errorf("Error reading service_providers: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemSamlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSamlStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlRole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlDefaultLoginPage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlDefaultProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlBindingProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlPortalUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlEntityId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlIdpEntityId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlIdpSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlIdpSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlIdpCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServerAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlTolerance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlLife(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProviders(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSamlServiceProvidersName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix"], _ = expandSystemSamlServiceProvidersPrefix(d, i["prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_binding_protocol"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["sp-binding-protocol"], _ = expandSystemSamlServiceProvidersSpBindingProtocol(d, i["sp_binding_protocol"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_cert"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["sp-cert"], _ = expandSystemSamlServiceProvidersSpCert(d, i["sp_cert"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_entity_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["sp-entity-id"], _ = expandSystemSamlServiceProvidersSpEntityId(d, i["sp_entity_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_single_sign_on_url"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["sp-single-sign-on-url"], _ = expandSystemSamlServiceProvidersSpSingleSignOnUrl(d, i["sp_single_sign_on_url"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_single_logout_url"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["sp-single-logout-url"], _ = expandSystemSamlServiceProvidersSpSingleLogoutUrl(d, i["sp_single_logout_url"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_portal_url"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["sp-portal-url"], _ = expandSystemSamlServiceProvidersSpPortalUrl(d, i["sp_portal_url"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_entity_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["idp-entity-id"], _ = expandSystemSamlServiceProvidersIdpEntityId(d, i["idp_entity_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_sign_on_url"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["idp-single-sign-on-url"], _ = expandSystemSamlServiceProvidersIdpSingleSignOnUrl(d, i["idp_single_sign_on_url"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_logout_url"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["idp-single-logout-url"], _ = expandSystemSamlServiceProvidersIdpSingleLogoutUrl(d, i["idp_single_logout_url"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assertion_attributes"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["assertion-attributes"], _ = expandSystemSamlServiceProvidersAssertionAttributes(d, i["assertion_attributes"], pre_append, sv)
		} else {
			tmp["assertion-attributes"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSamlServiceProvidersName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpBindingProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpEntityId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpPortalUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersIdpEntityId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersIdpSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersIdpSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersAssertionAttributes(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSamlServiceProvidersAssertionAttributesName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["type"], _ = expandSystemSamlServiceProvidersAssertionAttributesType(d, i["type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSamlServiceProvidersAssertionAttributesName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersAssertionAttributesType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSaml(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandSystemSamlStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("role"); ok {
		if setArgNil {
			obj["role"] = nil
		} else {

			t, err := expandSystemSamlRole(d, v, "role", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["role"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_login_page"); ok {
		if setArgNil {
			obj["default-login-page"] = nil
		} else {

			t, err := expandSystemSamlDefaultLoginPage(d, v, "default_login_page", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-login-page"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_profile"); ok {
		if setArgNil {
			obj["default-profile"] = nil
		} else {

			t, err := expandSystemSamlDefaultProfile(d, v, "default_profile", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-profile"] = t
			}
		}
	}

	if v, ok := d.GetOk("cert"); ok {
		if setArgNil {
			obj["cert"] = nil
		} else {

			t, err := expandSystemSamlCert(d, v, "cert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cert"] = t
			}
		}
	}

	if v, ok := d.GetOk("binding_protocol"); ok {
		if setArgNil {
			obj["binding-protocol"] = nil
		} else {

			t, err := expandSystemSamlBindingProtocol(d, v, "binding_protocol", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["binding-protocol"] = t
			}
		}
	}

	if v, ok := d.GetOk("portal_url"); ok {
		if setArgNil {
			obj["portal-url"] = nil
		} else {

			t, err := expandSystemSamlPortalUrl(d, v, "portal_url", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["portal-url"] = t
			}
		}
	}

	if v, ok := d.GetOk("entity_id"); ok {
		if setArgNil {
			obj["entity-id"] = nil
		} else {

			t, err := expandSystemSamlEntityId(d, v, "entity_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["entity-id"] = t
			}
		}
	}

	if v, ok := d.GetOk("single_sign_on_url"); ok {
		if setArgNil {
			obj["single-sign-on-url"] = nil
		} else {

			t, err := expandSystemSamlSingleSignOnUrl(d, v, "single_sign_on_url", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["single-sign-on-url"] = t
			}
		}
	}

	if v, ok := d.GetOk("single_logout_url"); ok {
		if setArgNil {
			obj["single-logout-url"] = nil
		} else {

			t, err := expandSystemSamlSingleLogoutUrl(d, v, "single_logout_url", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["single-logout-url"] = t
			}
		}
	}

	if v, ok := d.GetOk("idp_entity_id"); ok {
		if setArgNil {
			obj["idp-entity-id"] = nil
		} else {

			t, err := expandSystemSamlIdpEntityId(d, v, "idp_entity_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["idp-entity-id"] = t
			}
		}
	}

	if v, ok := d.GetOk("idp_single_sign_on_url"); ok {
		if setArgNil {
			obj["idp-single-sign-on-url"] = nil
		} else {

			t, err := expandSystemSamlIdpSingleSignOnUrl(d, v, "idp_single_sign_on_url", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["idp-single-sign-on-url"] = t
			}
		}
	}

	if v, ok := d.GetOk("idp_single_logout_url"); ok {
		if setArgNil {
			obj["idp-single-logout-url"] = nil
		} else {

			t, err := expandSystemSamlIdpSingleLogoutUrl(d, v, "idp_single_logout_url", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["idp-single-logout-url"] = t
			}
		}
	}

	if v, ok := d.GetOk("idp_cert"); ok {
		if setArgNil {
			obj["idp-cert"] = nil
		} else {

			t, err := expandSystemSamlIdpCert(d, v, "idp_cert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["idp-cert"] = t
			}
		}
	}

	if v, ok := d.GetOk("server_address"); ok {
		if setArgNil {
			obj["server-address"] = nil
		} else {

			t, err := expandSystemSamlServerAddress(d, v, "server_address", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server-address"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("tolerance"); ok {
		if setArgNil {
			obj["tolerance"] = nil
		} else {

			t, err := expandSystemSamlTolerance(d, v, "tolerance", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tolerance"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("life"); ok {
		if setArgNil {
			obj["life"] = nil
		} else {

			t, err := expandSystemSamlLife(d, v, "life", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["life"] = t
			}
		}
	}

	if v, ok := d.GetOk("service_providers"); ok || d.HasChange("service_providers") {
		if setArgNil {
			obj["service-providers"] = make([]struct{}, 0)
		} else {

			t, err := expandSystemSamlServiceProviders(d, v, "service_providers", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["service-providers"] = t
			}
		}
	}

	return &obj, nil
}
