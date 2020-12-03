// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Extender controller configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceExtenderControllerExtender() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtenderControllerExtenderCreate,
		Read:   resourceExtenderControllerExtenderRead,
		Update: resourceExtenderControllerExtenderUpdate,
		Delete: resourceExtenderControllerExtenderDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),
				Required:     true,
				ForceNew:     true,
			},
			"admin": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ifname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dial_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redundant_intf": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"dial_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"conn_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ext_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"quota_limit_mb": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10485760),
				Optional:     true,
				Computed:     true,
			},
			"billing_start_day": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 28),
				Optional:     true,
				Computed:     true,
			},
			"at_dial_script": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"modem_passwd": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 27),
				Optional:     true,
				Sensitive:    true,
			},
			"initiated_update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"modem_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ppp_username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"ppp_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 27),
				Optional:     true,
				Sensitive:    true,
			},
			"ppp_auth_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ppp_echo_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wimax_carrier": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"wimax_realm": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"wimax_auth_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sim_pin": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 27),
				Optional:     true,
				Sensitive:    true,
			},
			"access_point_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"multi_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"roaming": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cdma_nai": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"aaa_shared_secret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 27),
				Optional:     true,
				Sensitive:    true,
			},
			"ha_shared_secret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 27),
				Optional:     true,
				Sensitive:    true,
			},
			"primary_ha": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"secondary_ha": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"cdma_aaa_spi": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"cdma_ha_spi": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceExtenderControllerExtenderCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectExtenderControllerExtender(d)
	if err != nil {
		return fmt.Errorf("Error creating ExtenderControllerExtender resource while getting object: %v", err)
	}

	o, err := c.CreateExtenderControllerExtender(obj)

	if err != nil {
		return fmt.Errorf("Error creating ExtenderControllerExtender resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtenderControllerExtender")
	}

	return resourceExtenderControllerExtenderRead(d, m)
}

func resourceExtenderControllerExtenderUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectExtenderControllerExtender(d)
	if err != nil {
		return fmt.Errorf("Error updating ExtenderControllerExtender resource while getting object: %v", err)
	}

	o, err := c.UpdateExtenderControllerExtender(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating ExtenderControllerExtender resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtenderControllerExtender")
	}

	return resourceExtenderControllerExtenderRead(d, m)
}

func resourceExtenderControllerExtenderDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteExtenderControllerExtender(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting ExtenderControllerExtender resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtenderControllerExtenderRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadExtenderControllerExtender(mkey)
	if err != nil {
		return fmt.Errorf("Error reading ExtenderControllerExtender resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtenderControllerExtender(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ExtenderControllerExtender resource from API: %v", err)
	}
	return nil
}

func flattenExtenderControllerExtenderId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderAdmin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderIfname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderDialMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderRedial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderRedundantIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderDialStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderConnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderExtName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderQuotaLimitMb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderBillingStartDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderAtDialScript(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModemPasswd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderInitiatedUpdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModemType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderPppUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderPppPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderPppAuthProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderPppEchoRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderWimaxCarrier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderWimaxRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderWimaxAuthProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderSimPin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderAccessPointName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderMultiMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderRoaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderCdmaNai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderAaaSharedSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderHaSharedSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderPrimaryHa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderSecondaryHa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderCdmaAaaSpi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtenderCdmaHaSpi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectExtenderControllerExtender(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenExtenderControllerExtenderId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("admin", flattenExtenderControllerExtenderAdmin(o["admin"], d, "admin")); err != nil {
		if !fortiAPIPatch(o["admin"]) {
			return fmt.Errorf("Error reading admin: %v", err)
		}
	}

	if err = d.Set("ifname", flattenExtenderControllerExtenderIfname(o["ifname"], d, "ifname")); err != nil {
		if !fortiAPIPatch(o["ifname"]) {
			return fmt.Errorf("Error reading ifname: %v", err)
		}
	}

	if err = d.Set("vdom", flattenExtenderControllerExtenderVdom(o["vdom"], d, "vdom")); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("role", flattenExtenderControllerExtenderRole(o["role"], d, "role")); err != nil {
		if !fortiAPIPatch(o["role"]) {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("mode", flattenExtenderControllerExtenderMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("dial_mode", flattenExtenderControllerExtenderDialMode(o["dial-mode"], d, "dial_mode")); err != nil {
		if !fortiAPIPatch(o["dial-mode"]) {
			return fmt.Errorf("Error reading dial_mode: %v", err)
		}
	}

	if err = d.Set("redial", flattenExtenderControllerExtenderRedial(o["redial"], d, "redial")); err != nil {
		if !fortiAPIPatch(o["redial"]) {
			return fmt.Errorf("Error reading redial: %v", err)
		}
	}

	if err = d.Set("redundant_intf", flattenExtenderControllerExtenderRedundantIntf(o["redundant-intf"], d, "redundant_intf")); err != nil {
		if !fortiAPIPatch(o["redundant-intf"]) {
			return fmt.Errorf("Error reading redundant_intf: %v", err)
		}
	}

	if err = d.Set("dial_status", flattenExtenderControllerExtenderDialStatus(o["dial-status"], d, "dial_status")); err != nil {
		if !fortiAPIPatch(o["dial-status"]) {
			return fmt.Errorf("Error reading dial_status: %v", err)
		}
	}

	if err = d.Set("conn_status", flattenExtenderControllerExtenderConnStatus(o["conn-status"], d, "conn_status")); err != nil {
		if !fortiAPIPatch(o["conn-status"]) {
			return fmt.Errorf("Error reading conn_status: %v", err)
		}
	}

	if err = d.Set("ext_name", flattenExtenderControllerExtenderExtName(o["ext-name"], d, "ext_name")); err != nil {
		if !fortiAPIPatch(o["ext-name"]) {
			return fmt.Errorf("Error reading ext_name: %v", err)
		}
	}

	if err = d.Set("description", flattenExtenderControllerExtenderDescription(o["description"], d, "description")); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("quota_limit_mb", flattenExtenderControllerExtenderQuotaLimitMb(o["quota-limit-mb"], d, "quota_limit_mb")); err != nil {
		if !fortiAPIPatch(o["quota-limit-mb"]) {
			return fmt.Errorf("Error reading quota_limit_mb: %v", err)
		}
	}

	if err = d.Set("billing_start_day", flattenExtenderControllerExtenderBillingStartDay(o["billing-start-day"], d, "billing_start_day")); err != nil {
		if !fortiAPIPatch(o["billing-start-day"]) {
			return fmt.Errorf("Error reading billing_start_day: %v", err)
		}
	}

	if err = d.Set("at_dial_script", flattenExtenderControllerExtenderAtDialScript(o["at-dial-script"], d, "at_dial_script")); err != nil {
		if !fortiAPIPatch(o["at-dial-script"]) {
			return fmt.Errorf("Error reading at_dial_script: %v", err)
		}
	}

	if err = d.Set("initiated_update", flattenExtenderControllerExtenderInitiatedUpdate(o["initiated-update"], d, "initiated_update")); err != nil {
		if !fortiAPIPatch(o["initiated-update"]) {
			return fmt.Errorf("Error reading initiated_update: %v", err)
		}
	}

	if err = d.Set("modem_type", flattenExtenderControllerExtenderModemType(o["modem-type"], d, "modem_type")); err != nil {
		if !fortiAPIPatch(o["modem-type"]) {
			return fmt.Errorf("Error reading modem_type: %v", err)
		}
	}

	if err = d.Set("ppp_username", flattenExtenderControllerExtenderPppUsername(o["ppp-username"], d, "ppp_username")); err != nil {
		if !fortiAPIPatch(o["ppp-username"]) {
			return fmt.Errorf("Error reading ppp_username: %v", err)
		}
	}

	if err = d.Set("ppp_auth_protocol", flattenExtenderControllerExtenderPppAuthProtocol(o["ppp-auth-protocol"], d, "ppp_auth_protocol")); err != nil {
		if !fortiAPIPatch(o["ppp-auth-protocol"]) {
			return fmt.Errorf("Error reading ppp_auth_protocol: %v", err)
		}
	}

	if err = d.Set("ppp_echo_request", flattenExtenderControllerExtenderPppEchoRequest(o["ppp-echo-request"], d, "ppp_echo_request")); err != nil {
		if !fortiAPIPatch(o["ppp-echo-request"]) {
			return fmt.Errorf("Error reading ppp_echo_request: %v", err)
		}
	}

	if err = d.Set("wimax_carrier", flattenExtenderControllerExtenderWimaxCarrier(o["wimax-carrier"], d, "wimax_carrier")); err != nil {
		if !fortiAPIPatch(o["wimax-carrier"]) {
			return fmt.Errorf("Error reading wimax_carrier: %v", err)
		}
	}

	if err = d.Set("wimax_realm", flattenExtenderControllerExtenderWimaxRealm(o["wimax-realm"], d, "wimax_realm")); err != nil {
		if !fortiAPIPatch(o["wimax-realm"]) {
			return fmt.Errorf("Error reading wimax_realm: %v", err)
		}
	}

	if err = d.Set("wimax_auth_protocol", flattenExtenderControllerExtenderWimaxAuthProtocol(o["wimax-auth-protocol"], d, "wimax_auth_protocol")); err != nil {
		if !fortiAPIPatch(o["wimax-auth-protocol"]) {
			return fmt.Errorf("Error reading wimax_auth_protocol: %v", err)
		}
	}

	if err = d.Set("access_point_name", flattenExtenderControllerExtenderAccessPointName(o["access-point-name"], d, "access_point_name")); err != nil {
		if !fortiAPIPatch(o["access-point-name"]) {
			return fmt.Errorf("Error reading access_point_name: %v", err)
		}
	}

	if err = d.Set("multi_mode", flattenExtenderControllerExtenderMultiMode(o["multi-mode"], d, "multi_mode")); err != nil {
		if !fortiAPIPatch(o["multi-mode"]) {
			return fmt.Errorf("Error reading multi_mode: %v", err)
		}
	}

	if err = d.Set("roaming", flattenExtenderControllerExtenderRoaming(o["roaming"], d, "roaming")); err != nil {
		if !fortiAPIPatch(o["roaming"]) {
			return fmt.Errorf("Error reading roaming: %v", err)
		}
	}

	if err = d.Set("cdma_nai", flattenExtenderControllerExtenderCdmaNai(o["cdma-nai"], d, "cdma_nai")); err != nil {
		if !fortiAPIPatch(o["cdma-nai"]) {
			return fmt.Errorf("Error reading cdma_nai: %v", err)
		}
	}

	if err = d.Set("primary_ha", flattenExtenderControllerExtenderPrimaryHa(o["primary-ha"], d, "primary_ha")); err != nil {
		if !fortiAPIPatch(o["primary-ha"]) {
			return fmt.Errorf("Error reading primary_ha: %v", err)
		}
	}

	if err = d.Set("secondary_ha", flattenExtenderControllerExtenderSecondaryHa(o["secondary-ha"], d, "secondary_ha")); err != nil {
		if !fortiAPIPatch(o["secondary-ha"]) {
			return fmt.Errorf("Error reading secondary_ha: %v", err)
		}
	}

	if err = d.Set("cdma_aaa_spi", flattenExtenderControllerExtenderCdmaAaaSpi(o["cdma-aaa-spi"], d, "cdma_aaa_spi")); err != nil {
		if !fortiAPIPatch(o["cdma-aaa-spi"]) {
			return fmt.Errorf("Error reading cdma_aaa_spi: %v", err)
		}
	}

	if err = d.Set("cdma_ha_spi", flattenExtenderControllerExtenderCdmaHaSpi(o["cdma-ha-spi"], d, "cdma_ha_spi")); err != nil {
		if !fortiAPIPatch(o["cdma-ha-spi"]) {
			return fmt.Errorf("Error reading cdma_ha_spi: %v", err)
		}
	}

	return nil
}

func flattenExtenderControllerExtenderFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExtenderControllerExtenderId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderAdmin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderIfname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderDialMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderRedial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderRedundantIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderDialStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderConnStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderExtName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderQuotaLimitMb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderBillingStartDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderAtDialScript(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModemPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderInitiatedUpdate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModemType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderPppUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderPppPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderPppAuthProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderPppEchoRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderWimaxCarrier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderWimaxRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderWimaxAuthProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderSimPin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderAccessPointName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderMultiMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderRoaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderCdmaNai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderAaaSharedSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderHaSharedSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderPrimaryHa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderSecondaryHa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderCdmaAaaSpi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderCdmaHaSpi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectExtenderControllerExtender(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandExtenderControllerExtenderId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("admin"); ok {
		t, err := expandExtenderControllerExtenderAdmin(d, v, "admin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin"] = t
		}
	}

	if v, ok := d.GetOk("ifname"); ok {
		t, err := expandExtenderControllerExtenderIfname(d, v, "ifname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ifname"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		t, err := expandExtenderControllerExtenderVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok {
		t, err := expandExtenderControllerExtenderRole(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandExtenderControllerExtenderMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("dial_mode"); ok {
		t, err := expandExtenderControllerExtenderDialMode(d, v, "dial_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dial-mode"] = t
		}
	}

	if v, ok := d.GetOk("redial"); ok {
		t, err := expandExtenderControllerExtenderRedial(d, v, "redial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redial"] = t
		}
	}

	if v, ok := d.GetOk("redundant_intf"); ok {
		t, err := expandExtenderControllerExtenderRedundantIntf(d, v, "redundant_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant-intf"] = t
		}
	}

	if v, ok := d.GetOk("dial_status"); ok {
		t, err := expandExtenderControllerExtenderDialStatus(d, v, "dial_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dial-status"] = t
		}
	}

	if v, ok := d.GetOk("conn_status"); ok {
		t, err := expandExtenderControllerExtenderConnStatus(d, v, "conn_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn-status"] = t
		}
	}

	if v, ok := d.GetOk("ext_name"); ok {
		t, err := expandExtenderControllerExtenderExtName(d, v, "ext_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ext-name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandExtenderControllerExtenderDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("quota_limit_mb"); ok {
		t, err := expandExtenderControllerExtenderQuotaLimitMb(d, v, "quota_limit_mb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quota-limit-mb"] = t
		}
	}

	if v, ok := d.GetOk("billing_start_day"); ok {
		t, err := expandExtenderControllerExtenderBillingStartDay(d, v, "billing_start_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["billing-start-day"] = t
		}
	}

	if v, ok := d.GetOk("at_dial_script"); ok {
		t, err := expandExtenderControllerExtenderAtDialScript(d, v, "at_dial_script")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["at-dial-script"] = t
		}
	}

	if v, ok := d.GetOk("modem_passwd"); ok {
		t, err := expandExtenderControllerExtenderModemPasswd(d, v, "modem_passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem-passwd"] = t
		}
	}

	if v, ok := d.GetOk("initiated_update"); ok {
		t, err := expandExtenderControllerExtenderInitiatedUpdate(d, v, "initiated_update")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["initiated-update"] = t
		}
	}

	if v, ok := d.GetOk("modem_type"); ok {
		t, err := expandExtenderControllerExtenderModemType(d, v, "modem_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem-type"] = t
		}
	}

	if v, ok := d.GetOk("ppp_username"); ok {
		t, err := expandExtenderControllerExtenderPppUsername(d, v, "ppp_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppp-username"] = t
		}
	}

	if v, ok := d.GetOk("ppp_password"); ok {
		t, err := expandExtenderControllerExtenderPppPassword(d, v, "ppp_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppp-password"] = t
		}
	}

	if v, ok := d.GetOk("ppp_auth_protocol"); ok {
		t, err := expandExtenderControllerExtenderPppAuthProtocol(d, v, "ppp_auth_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppp-auth-protocol"] = t
		}
	}

	if v, ok := d.GetOk("ppp_echo_request"); ok {
		t, err := expandExtenderControllerExtenderPppEchoRequest(d, v, "ppp_echo_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppp-echo-request"] = t
		}
	}

	if v, ok := d.GetOk("wimax_carrier"); ok {
		t, err := expandExtenderControllerExtenderWimaxCarrier(d, v, "wimax_carrier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wimax-carrier"] = t
		}
	}

	if v, ok := d.GetOk("wimax_realm"); ok {
		t, err := expandExtenderControllerExtenderWimaxRealm(d, v, "wimax_realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wimax-realm"] = t
		}
	}

	if v, ok := d.GetOk("wimax_auth_protocol"); ok {
		t, err := expandExtenderControllerExtenderWimaxAuthProtocol(d, v, "wimax_auth_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wimax-auth-protocol"] = t
		}
	}

	if v, ok := d.GetOk("sim_pin"); ok {
		t, err := expandExtenderControllerExtenderSimPin(d, v, "sim_pin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim-pin"] = t
		}
	}

	if v, ok := d.GetOk("access_point_name"); ok {
		t, err := expandExtenderControllerExtenderAccessPointName(d, v, "access_point_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-point-name"] = t
		}
	}

	if v, ok := d.GetOk("multi_mode"); ok {
		t, err := expandExtenderControllerExtenderMultiMode(d, v, "multi_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multi-mode"] = t
		}
	}

	if v, ok := d.GetOk("roaming"); ok {
		t, err := expandExtenderControllerExtenderRoaming(d, v, "roaming")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roaming"] = t
		}
	}

	if v, ok := d.GetOk("cdma_nai"); ok {
		t, err := expandExtenderControllerExtenderCdmaNai(d, v, "cdma_nai")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cdma-nai"] = t
		}
	}

	if v, ok := d.GetOk("aaa_shared_secret"); ok {
		t, err := expandExtenderControllerExtenderAaaSharedSecret(d, v, "aaa_shared_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aaa-shared-secret"] = t
		}
	}

	if v, ok := d.GetOk("ha_shared_secret"); ok {
		t, err := expandExtenderControllerExtenderHaSharedSecret(d, v, "ha_shared_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-shared-secret"] = t
		}
	}

	if v, ok := d.GetOk("primary_ha"); ok {
		t, err := expandExtenderControllerExtenderPrimaryHa(d, v, "primary_ha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary-ha"] = t
		}
	}

	if v, ok := d.GetOk("secondary_ha"); ok {
		t, err := expandExtenderControllerExtenderSecondaryHa(d, v, "secondary_ha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-ha"] = t
		}
	}

	if v, ok := d.GetOk("cdma_aaa_spi"); ok {
		t, err := expandExtenderControllerExtenderCdmaAaaSpi(d, v, "cdma_aaa_spi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cdma-aaa-spi"] = t
		}
	}

	if v, ok := d.GetOk("cdma_ha_spi"); ok {
		t, err := expandExtenderControllerExtenderCdmaHaSpi(d, v, "cdma_ha_spi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cdma-ha-spi"] = t
		}
	}

	return &obj, nil
}
