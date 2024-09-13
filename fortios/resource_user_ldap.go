// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure LDAP server entries.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserLdap() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserLdapCreate,
		Read:   resourceUserLdapRead,
		Update: resourceUserLdapUpdate,
		Delete: resourceUserLdapDelete,

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
				Required:     true,
			},
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required:     true,
			},
			"secondary_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"tertiary_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"status_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 600),
				Optional:     true,
				Computed:     true,
			},
			"server_identity_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip_interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"source_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
			},
			"cnid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 20),
				Optional:     true,
				Computed:     true,
			},
			"dn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Required:     true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"two_factor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"two_factor_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"two_factor_notification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"two_factor_filter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),
				Optional:     true,
			},
			"username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"group_member_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_search_base": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
			},
			"group_object_filter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),
				Optional:     true,
				Computed:     true,
			},
			"group_filter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),
				Optional:     true,
			},
			"secure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ca_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"password_expiry_warning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password_renewal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"member_attr": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"account_key_processing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"account_key_cert_field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"account_key_upn_san": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"account_key_filter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),
				Optional:     true,
				Computed:     true,
			},
			"search_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_cert_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"obtain_user_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_info_exchange_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"antiphish": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password_attr": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceUserLdapCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserLdap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserLdap resource while getting object: %v", err)
	}

	o, err := c.CreateUserLdap(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating UserLdap resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserLdap")
	}

	return resourceUserLdapRead(d, m)
}

func resourceUserLdapUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserLdap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserLdap resource while getting object: %v", err)
	}

	o, err := c.UpdateUserLdap(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserLdap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserLdap")
	}

	return resourceUserLdapRead(d, m)
}

func resourceUserLdapDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUserLdap(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UserLdap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserLdapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserLdap(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserLdap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserLdap(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserLdap resource from API: %v", err)
	}
	return nil
}

func flattenUserLdapName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapSecondaryServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapTertiaryServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapStatusTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserLdapServerIdentityCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapSourceIpInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapSourcePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserLdapCnid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapDn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapTwoFactor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapTwoFactorAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapTwoFactorNotification(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapTwoFactorFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapGroupMemberCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapGroupSearchBase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapGroupObjectFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapGroupFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapSecure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapCaCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserLdapPasswordExpiryWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapPasswordRenewal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapMemberAttr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapAccountKeyProcessing(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapAccountKeyCertField(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapAccountKeyUpnSan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapAccountKeyFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapSearchType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapClientCertAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapClientCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapObtainUserInfo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapUserInfoExchangeServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapAntiphish(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLdapPasswordAttr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserLdap(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenUserLdapName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server", flattenUserLdapServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("secondary_server", flattenUserLdapSecondaryServer(o["secondary-server"], d, "secondary_server", sv)); err != nil {
		if !fortiAPIPatch(o["secondary-server"]) {
			return fmt.Errorf("Error reading secondary_server: %v", err)
		}
	}

	if err = d.Set("tertiary_server", flattenUserLdapTertiaryServer(o["tertiary-server"], d, "tertiary_server", sv)); err != nil {
		if !fortiAPIPatch(o["tertiary-server"]) {
			return fmt.Errorf("Error reading tertiary_server: %v", err)
		}
	}

	if err = d.Set("status_ttl", flattenUserLdapStatusTtl(o["status-ttl"], d, "status_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["status-ttl"]) {
			return fmt.Errorf("Error reading status_ttl: %v", err)
		}
	}

	if err = d.Set("server_identity_check", flattenUserLdapServerIdentityCheck(o["server-identity-check"], d, "server_identity_check", sv)); err != nil {
		if !fortiAPIPatch(o["server-identity-check"]) {
			return fmt.Errorf("Error reading server_identity_check: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenUserLdapSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip_interface", flattenUserLdapSourceIpInterface(o["source-ip-interface"], d, "source_ip_interface", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip-interface"]) {
			return fmt.Errorf("Error reading source_ip_interface: %v", err)
		}
	}

	if err = d.Set("source_port", flattenUserLdapSourcePort(o["source-port"], d, "source_port", sv)); err != nil {
		if !fortiAPIPatch(o["source-port"]) {
			return fmt.Errorf("Error reading source_port: %v", err)
		}
	}

	if err = d.Set("cnid", flattenUserLdapCnid(o["cnid"], d, "cnid", sv)); err != nil {
		if !fortiAPIPatch(o["cnid"]) {
			return fmt.Errorf("Error reading cnid: %v", err)
		}
	}

	if err = d.Set("dn", flattenUserLdapDn(o["dn"], d, "dn", sv)); err != nil {
		if !fortiAPIPatch(o["dn"]) {
			return fmt.Errorf("Error reading dn: %v", err)
		}
	}

	if err = d.Set("type", flattenUserLdapType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("two_factor", flattenUserLdapTwoFactor(o["two-factor"], d, "two_factor", sv)); err != nil {
		if !fortiAPIPatch(o["two-factor"]) {
			return fmt.Errorf("Error reading two_factor: %v", err)
		}
	}

	if err = d.Set("two_factor_authentication", flattenUserLdapTwoFactorAuthentication(o["two-factor-authentication"], d, "two_factor_authentication", sv)); err != nil {
		if !fortiAPIPatch(o["two-factor-authentication"]) {
			return fmt.Errorf("Error reading two_factor_authentication: %v", err)
		}
	}

	if err = d.Set("two_factor_notification", flattenUserLdapTwoFactorNotification(o["two-factor-notification"], d, "two_factor_notification", sv)); err != nil {
		if !fortiAPIPatch(o["two-factor-notification"]) {
			return fmt.Errorf("Error reading two_factor_notification: %v", err)
		}
	}

	if err = d.Set("two_factor_filter", flattenUserLdapTwoFactorFilter(o["two-factor-filter"], d, "two_factor_filter", sv)); err != nil {
		if !fortiAPIPatch(o["two-factor-filter"]) {
			return fmt.Errorf("Error reading two_factor_filter: %v", err)
		}
	}

	if err = d.Set("username", flattenUserLdapUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("group_member_check", flattenUserLdapGroupMemberCheck(o["group-member-check"], d, "group_member_check", sv)); err != nil {
		if !fortiAPIPatch(o["group-member-check"]) {
			return fmt.Errorf("Error reading group_member_check: %v", err)
		}
	}

	if err = d.Set("group_search_base", flattenUserLdapGroupSearchBase(o["group-search-base"], d, "group_search_base", sv)); err != nil {
		if !fortiAPIPatch(o["group-search-base"]) {
			return fmt.Errorf("Error reading group_search_base: %v", err)
		}
	}

	if err = d.Set("group_object_filter", flattenUserLdapGroupObjectFilter(o["group-object-filter"], d, "group_object_filter", sv)); err != nil {
		if !fortiAPIPatch(o["group-object-filter"]) {
			return fmt.Errorf("Error reading group_object_filter: %v", err)
		}
	}

	if err = d.Set("group_filter", flattenUserLdapGroupFilter(o["group-filter"], d, "group_filter", sv)); err != nil {
		if !fortiAPIPatch(o["group-filter"]) {
			return fmt.Errorf("Error reading group_filter: %v", err)
		}
	}

	if err = d.Set("secure", flattenUserLdapSecure(o["secure"], d, "secure", sv)); err != nil {
		if !fortiAPIPatch(o["secure"]) {
			return fmt.Errorf("Error reading secure: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenUserLdapSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-version"]) {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("ca_cert", flattenUserLdapCaCert(o["ca-cert"], d, "ca_cert", sv)); err != nil {
		if !fortiAPIPatch(o["ca-cert"]) {
			return fmt.Errorf("Error reading ca_cert: %v", err)
		}
	}

	if err = d.Set("port", flattenUserLdapPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("password_expiry_warning", flattenUserLdapPasswordExpiryWarning(o["password-expiry-warning"], d, "password_expiry_warning", sv)); err != nil {
		if !fortiAPIPatch(o["password-expiry-warning"]) {
			return fmt.Errorf("Error reading password_expiry_warning: %v", err)
		}
	}

	if err = d.Set("password_renewal", flattenUserLdapPasswordRenewal(o["password-renewal"], d, "password_renewal", sv)); err != nil {
		if !fortiAPIPatch(o["password-renewal"]) {
			return fmt.Errorf("Error reading password_renewal: %v", err)
		}
	}

	if err = d.Set("member_attr", flattenUserLdapMemberAttr(o["member-attr"], d, "member_attr", sv)); err != nil {
		if !fortiAPIPatch(o["member-attr"]) {
			return fmt.Errorf("Error reading member_attr: %v", err)
		}
	}

	if err = d.Set("account_key_processing", flattenUserLdapAccountKeyProcessing(o["account-key-processing"], d, "account_key_processing", sv)); err != nil {
		if !fortiAPIPatch(o["account-key-processing"]) {
			return fmt.Errorf("Error reading account_key_processing: %v", err)
		}
	}

	if err = d.Set("account_key_cert_field", flattenUserLdapAccountKeyCertField(o["account-key-cert-field"], d, "account_key_cert_field", sv)); err != nil {
		if !fortiAPIPatch(o["account-key-cert-field"]) {
			return fmt.Errorf("Error reading account_key_cert_field: %v", err)
		}
	}

	if err = d.Set("account_key_upn_san", flattenUserLdapAccountKeyUpnSan(o["account-key-upn-san"], d, "account_key_upn_san", sv)); err != nil {
		if !fortiAPIPatch(o["account-key-upn-san"]) {
			return fmt.Errorf("Error reading account_key_upn_san: %v", err)
		}
	}

	if err = d.Set("account_key_filter", flattenUserLdapAccountKeyFilter(o["account-key-filter"], d, "account_key_filter", sv)); err != nil {
		if !fortiAPIPatch(o["account-key-filter"]) {
			return fmt.Errorf("Error reading account_key_filter: %v", err)
		}
	}

	if err = d.Set("search_type", flattenUserLdapSearchType(o["search-type"], d, "search_type", sv)); err != nil {
		if !fortiAPIPatch(o["search-type"]) {
			return fmt.Errorf("Error reading search_type: %v", err)
		}
	}

	if err = d.Set("client_cert_auth", flattenUserLdapClientCertAuth(o["client-cert-auth"], d, "client_cert_auth", sv)); err != nil {
		if !fortiAPIPatch(o["client-cert-auth"]) {
			return fmt.Errorf("Error reading client_cert_auth: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenUserLdapClientCert(o["client-cert"], d, "client_cert", sv)); err != nil {
		if !fortiAPIPatch(o["client-cert"]) {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("obtain_user_info", flattenUserLdapObtainUserInfo(o["obtain-user-info"], d, "obtain_user_info", sv)); err != nil {
		if !fortiAPIPatch(o["obtain-user-info"]) {
			return fmt.Errorf("Error reading obtain_user_info: %v", err)
		}
	}

	if err = d.Set("user_info_exchange_server", flattenUserLdapUserInfoExchangeServer(o["user-info-exchange-server"], d, "user_info_exchange_server", sv)); err != nil {
		if !fortiAPIPatch(o["user-info-exchange-server"]) {
			return fmt.Errorf("Error reading user_info_exchange_server: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenUserLdapInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenUserLdapInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("antiphish", flattenUserLdapAntiphish(o["antiphish"], d, "antiphish", sv)); err != nil {
		if !fortiAPIPatch(o["antiphish"]) {
			return fmt.Errorf("Error reading antiphish: %v", err)
		}
	}

	if err = d.Set("password_attr", flattenUserLdapPasswordAttr(o["password-attr"], d, "password_attr", sv)); err != nil {
		if !fortiAPIPatch(o["password-attr"]) {
			return fmt.Errorf("Error reading password_attr: %v", err)
		}
	}

	return nil
}

func flattenUserLdapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserLdapName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSecondaryServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapTertiaryServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapStatusTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapServerIdentityCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSourceIpInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSourcePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapCnid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapDn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapTwoFactor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapTwoFactorAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapTwoFactorNotification(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapTwoFactorFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapGroupMemberCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapGroupSearchBase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapGroupObjectFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapGroupFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSecure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapCaCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPasswordExpiryWarning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPasswordRenewal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapMemberAttr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapAccountKeyProcessing(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapAccountKeyCertField(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapAccountKeyUpnSan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapAccountKeyFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapSearchType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapClientCertAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapClientCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapObtainUserInfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapUserInfoExchangeServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapAntiphish(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLdapPasswordAttr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserLdap(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserLdapName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandUserLdapServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	} else if d.HasChange("server") {
		obj["server"] = nil
	}

	if v, ok := d.GetOk("secondary_server"); ok {
		t, err := expandUserLdapSecondaryServer(d, v, "secondary_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-server"] = t
		}
	} else if d.HasChange("secondary_server") {
		obj["secondary-server"] = nil
	}

	if v, ok := d.GetOk("tertiary_server"); ok {
		t, err := expandUserLdapTertiaryServer(d, v, "tertiary_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-server"] = t
		}
	} else if d.HasChange("tertiary_server") {
		obj["tertiary-server"] = nil
	}

	if v, ok := d.GetOkExists("status_ttl"); ok {
		t, err := expandUserLdapStatusTtl(d, v, "status_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status-ttl"] = t
		}
	}

	if v, ok := d.GetOk("server_identity_check"); ok {
		t, err := expandUserLdapServerIdentityCheck(d, v, "server_identity_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-identity-check"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandUserLdapSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip_interface"); ok {
		t, err := expandUserLdapSourceIpInterface(d, v, "source_ip_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-interface"] = t
		}
	} else if d.HasChange("source_ip_interface") {
		obj["source-ip-interface"] = nil
	}

	if v, ok := d.GetOkExists("source_port"); ok {
		t, err := expandUserLdapSourcePort(d, v, "source_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-port"] = t
		}
	} else if d.HasChange("source_port") {
		obj["source-port"] = nil
	}

	if v, ok := d.GetOk("cnid"); ok {
		t, err := expandUserLdapCnid(d, v, "cnid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cnid"] = t
		}
	}

	if v, ok := d.GetOk("dn"); ok {
		t, err := expandUserLdapDn(d, v, "dn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dn"] = t
		}
	} else if d.HasChange("dn") {
		obj["dn"] = nil
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandUserLdapType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("two_factor"); ok {
		t, err := expandUserLdapTwoFactor(d, v, "two_factor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_authentication"); ok {
		t, err := expandUserLdapTwoFactorAuthentication(d, v, "two_factor_authentication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-authentication"] = t
		}
	} else if d.HasChange("two_factor_authentication") {
		obj["two-factor-authentication"] = nil
	}

	if v, ok := d.GetOk("two_factor_notification"); ok {
		t, err := expandUserLdapTwoFactorNotification(d, v, "two_factor_notification", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-notification"] = t
		}
	} else if d.HasChange("two_factor_notification") {
		obj["two-factor-notification"] = nil
	}

	if v, ok := d.GetOk("two_factor_filter"); ok {
		t, err := expandUserLdapTwoFactorFilter(d, v, "two_factor_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-filter"] = t
		}
	} else if d.HasChange("two_factor_filter") {
		obj["two-factor-filter"] = nil
	}

	if v, ok := d.GetOk("username"); ok {
		t, err := expandUserLdapUsername(d, v, "username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	} else if d.HasChange("username") {
		obj["username"] = nil
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandUserLdapPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	} else if d.HasChange("password") {
		obj["password"] = nil
	}

	if v, ok := d.GetOk("group_member_check"); ok {
		t, err := expandUserLdapGroupMemberCheck(d, v, "group_member_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-member-check"] = t
		}
	}

	if v, ok := d.GetOk("group_search_base"); ok {
		t, err := expandUserLdapGroupSearchBase(d, v, "group_search_base", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-search-base"] = t
		}
	} else if d.HasChange("group_search_base") {
		obj["group-search-base"] = nil
	}

	if v, ok := d.GetOk("group_object_filter"); ok {
		t, err := expandUserLdapGroupObjectFilter(d, v, "group_object_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-object-filter"] = t
		}
	}

	if v, ok := d.GetOk("group_filter"); ok {
		t, err := expandUserLdapGroupFilter(d, v, "group_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-filter"] = t
		}
	} else if d.HasChange("group_filter") {
		obj["group-filter"] = nil
	}

	if v, ok := d.GetOk("secure"); ok {
		t, err := expandUserLdapSecure(d, v, "secure", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok {
		t, err := expandUserLdapSslMinProtoVersion(d, v, "ssl_min_proto_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("ca_cert"); ok {
		t, err := expandUserLdapCaCert(d, v, "ca_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-cert"] = t
		}
	} else if d.HasChange("ca_cert") {
		obj["ca-cert"] = nil
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandUserLdapPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("password_expiry_warning"); ok {
		t, err := expandUserLdapPasswordExpiryWarning(d, v, "password_expiry_warning", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-expiry-warning"] = t
		}
	}

	if v, ok := d.GetOk("password_renewal"); ok {
		t, err := expandUserLdapPasswordRenewal(d, v, "password_renewal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-renewal"] = t
		}
	}

	if v, ok := d.GetOk("member_attr"); ok {
		t, err := expandUserLdapMemberAttr(d, v, "member_attr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member-attr"] = t
		}
	}

	if v, ok := d.GetOk("account_key_processing"); ok {
		t, err := expandUserLdapAccountKeyProcessing(d, v, "account_key_processing", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-processing"] = t
		}
	}

	if v, ok := d.GetOk("account_key_cert_field"); ok {
		t, err := expandUserLdapAccountKeyCertField(d, v, "account_key_cert_field", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-cert-field"] = t
		}
	}

	if v, ok := d.GetOk("account_key_upn_san"); ok {
		t, err := expandUserLdapAccountKeyUpnSan(d, v, "account_key_upn_san", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-upn-san"] = t
		}
	}

	if v, ok := d.GetOk("account_key_filter"); ok {
		t, err := expandUserLdapAccountKeyFilter(d, v, "account_key_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-filter"] = t
		}
	}

	if v, ok := d.GetOk("search_type"); ok {
		t, err := expandUserLdapSearchType(d, v, "search_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["search-type"] = t
		}
	} else if d.HasChange("search_type") {
		obj["search-type"] = nil
	}

	if v, ok := d.GetOk("client_cert_auth"); ok {
		t, err := expandUserLdapClientCertAuth(d, v, "client_cert_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert-auth"] = t
		}
	}

	if v, ok := d.GetOk("client_cert"); ok {
		t, err := expandUserLdapClientCert(d, v, "client_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	} else if d.HasChange("client_cert") {
		obj["client-cert"] = nil
	}

	if v, ok := d.GetOk("obtain_user_info"); ok {
		t, err := expandUserLdapObtainUserInfo(d, v, "obtain_user_info", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obtain-user-info"] = t
		}
	}

	if v, ok := d.GetOk("user_info_exchange_server"); ok {
		t, err := expandUserLdapUserInfoExchangeServer(d, v, "user_info_exchange_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-info-exchange-server"] = t
		}
	} else if d.HasChange("user_info_exchange_server") {
		obj["user-info-exchange-server"] = nil
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		t, err := expandUserLdapInterfaceSelectMethod(d, v, "interface_select_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandUserLdapInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOk("antiphish"); ok {
		t, err := expandUserLdapAntiphish(d, v, "antiphish", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antiphish"] = t
		}
	}

	if v, ok := d.GetOk("password_attr"); ok {
		t, err := expandUserLdapPasswordAttr(d, v, "password_attr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-attr"] = t
		}
	}

	return &obj, nil
}
