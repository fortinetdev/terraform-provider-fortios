// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure fortiswitch's admin security-policy.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerSecurityPolicyAdmin() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSecurityPolicyAdminCreate,
		Read:   resourceSwitchControllerSecurityPolicyAdminRead,
		Update: resourceSwitchControllerSecurityPolicyAdminUpdate,
		Delete: resourceSwitchControllerSecurityPolicyAdminDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 31),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"auto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost8": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost9": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusthost10": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost8": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost9": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost10": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerSecurityPolicyAdminCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerSecurityPolicyAdmin(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSecurityPolicyAdmin resource while getting object: %v", err)
	}

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSwitchControllerSecurityPolicyAdmin(mkey, vdomparam)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSwitchControllerSecurityPolicyAdmin(obj, mkey, vdomparam)
			if err != nil {
				return fmt.Errorf("Error updating SwitchControllerSecurityPolicyAdmin resource: %v", err)
			}
		}
	}

	if !existing {
		o, err = c.CreateSwitchControllerSecurityPolicyAdmin(obj, vdomparam)

		if err != nil {
			return fmt.Errorf("Error creating SwitchControllerSecurityPolicyAdmin resource: %v", err)
		}
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerSecurityPolicyAdmin")
	}

	return resourceSwitchControllerSecurityPolicyAdminRead(d, m)
}

func resourceSwitchControllerSecurityPolicyAdminUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerSecurityPolicyAdmin(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSecurityPolicyAdmin resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerSecurityPolicyAdmin(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSecurityPolicyAdmin resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerSecurityPolicyAdmin")
	}

	return resourceSwitchControllerSecurityPolicyAdminRead(d, m)
}

func resourceSwitchControllerSecurityPolicyAdminDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerSecurityPolicyAdmin(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerSecurityPolicyAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSecurityPolicyAdminRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerSecurityPolicyAdmin(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSecurityPolicyAdmin resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSecurityPolicyAdmin(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSecurityPolicyAdmin resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSecurityPolicyAdminName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminAuto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost5(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost7(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost8(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost9(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchControllerSecurityPolicyAdminTrusthost10(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost5(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost7(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost8(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost9(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyAdminIp6Trusthost10(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerSecurityPolicyAdmin(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerSecurityPolicyAdminName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("auto", flattenSwitchControllerSecurityPolicyAdminAuto(o["auto"], d, "auto", sv)); err != nil {
		if !fortiAPIPatch(o["auto"]) {
			return fmt.Errorf("Error reading auto: %v", err)
		}
	}

	if err = d.Set("trusthost1", flattenSwitchControllerSecurityPolicyAdminTrusthost1(o["trusthost1"], d, "trusthost1", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost1"]) {
			return fmt.Errorf("Error reading trusthost1: %v", err)
		}
	}

	if err = d.Set("trusthost2", flattenSwitchControllerSecurityPolicyAdminTrusthost2(o["trusthost2"], d, "trusthost2", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost2"]) {
			return fmt.Errorf("Error reading trusthost2: %v", err)
		}
	}

	if err = d.Set("trusthost3", flattenSwitchControllerSecurityPolicyAdminTrusthost3(o["trusthost3"], d, "trusthost3", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost3"]) {
			return fmt.Errorf("Error reading trusthost3: %v", err)
		}
	}

	if err = d.Set("trusthost4", flattenSwitchControllerSecurityPolicyAdminTrusthost4(o["trusthost4"], d, "trusthost4", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost4"]) {
			return fmt.Errorf("Error reading trusthost4: %v", err)
		}
	}

	if err = d.Set("trusthost5", flattenSwitchControllerSecurityPolicyAdminTrusthost5(o["trusthost5"], d, "trusthost5", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost5"]) {
			return fmt.Errorf("Error reading trusthost5: %v", err)
		}
	}

	if err = d.Set("trusthost6", flattenSwitchControllerSecurityPolicyAdminTrusthost6(o["trusthost6"], d, "trusthost6", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost6"]) {
			return fmt.Errorf("Error reading trusthost6: %v", err)
		}
	}

	if err = d.Set("trusthost7", flattenSwitchControllerSecurityPolicyAdminTrusthost7(o["trusthost7"], d, "trusthost7", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost7"]) {
			return fmt.Errorf("Error reading trusthost7: %v", err)
		}
	}

	if err = d.Set("trusthost8", flattenSwitchControllerSecurityPolicyAdminTrusthost8(o["trusthost8"], d, "trusthost8", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost8"]) {
			return fmt.Errorf("Error reading trusthost8: %v", err)
		}
	}

	if err = d.Set("trusthost9", flattenSwitchControllerSecurityPolicyAdminTrusthost9(o["trusthost9"], d, "trusthost9", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost9"]) {
			return fmt.Errorf("Error reading trusthost9: %v", err)
		}
	}

	if err = d.Set("trusthost10", flattenSwitchControllerSecurityPolicyAdminTrusthost10(o["trusthost10"], d, "trusthost10", sv)); err != nil {
		if !fortiAPIPatch(o["trusthost10"]) {
			return fmt.Errorf("Error reading trusthost10: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost1", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost1(o["ip6-trusthost1"], d, "ip6_trusthost1", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost1"]) {
			return fmt.Errorf("Error reading ip6_trusthost1: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost2", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost2(o["ip6-trusthost2"], d, "ip6_trusthost2", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost2"]) {
			return fmt.Errorf("Error reading ip6_trusthost2: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost3", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost3(o["ip6-trusthost3"], d, "ip6_trusthost3", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost3"]) {
			return fmt.Errorf("Error reading ip6_trusthost3: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost4", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost4(o["ip6-trusthost4"], d, "ip6_trusthost4", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost4"]) {
			return fmt.Errorf("Error reading ip6_trusthost4: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost5", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost5(o["ip6-trusthost5"], d, "ip6_trusthost5", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost5"]) {
			return fmt.Errorf("Error reading ip6_trusthost5: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost6", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost6(o["ip6-trusthost6"], d, "ip6_trusthost6", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost6"]) {
			return fmt.Errorf("Error reading ip6_trusthost6: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost7", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost7(o["ip6-trusthost7"], d, "ip6_trusthost7", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost7"]) {
			return fmt.Errorf("Error reading ip6_trusthost7: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost8", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost8(o["ip6-trusthost8"], d, "ip6_trusthost8", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost8"]) {
			return fmt.Errorf("Error reading ip6_trusthost8: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost9", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost9(o["ip6-trusthost9"], d, "ip6_trusthost9", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost9"]) {
			return fmt.Errorf("Error reading ip6_trusthost9: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost10", flattenSwitchControllerSecurityPolicyAdminIp6Trusthost10(o["ip6-trusthost10"], d, "ip6_trusthost10", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-trusthost10"]) {
			return fmt.Errorf("Error reading ip6_trusthost10: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSecurityPolicyAdminFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerSecurityPolicyAdminName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminAuto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost5(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost7(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost8(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost9(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminTrusthost10(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost5(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost7(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost8(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost9(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyAdminIp6Trusthost10(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerSecurityPolicyAdmin(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("auto"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminAuto(d, v, "auto", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto"] = t
		}
	}

	if v, ok := d.GetOk("trusthost1"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost1(d, v, "trusthost1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost1"] = t
		}
	}

	if v, ok := d.GetOk("trusthost2"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost2(d, v, "trusthost2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost2"] = t
		}
	}

	if v, ok := d.GetOk("trusthost3"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost3(d, v, "trusthost3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost3"] = t
		}
	}

	if v, ok := d.GetOk("trusthost4"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost4(d, v, "trusthost4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost4"] = t
		}
	}

	if v, ok := d.GetOk("trusthost5"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost5(d, v, "trusthost5", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost5"] = t
		}
	}

	if v, ok := d.GetOk("trusthost6"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost6(d, v, "trusthost6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost6"] = t
		}
	}

	if v, ok := d.GetOk("trusthost7"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost7(d, v, "trusthost7", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost7"] = t
		}
	}

	if v, ok := d.GetOk("trusthost8"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost8(d, v, "trusthost8", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost8"] = t
		}
	}

	if v, ok := d.GetOk("trusthost9"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost9(d, v, "trusthost9", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost9"] = t
		}
	}

	if v, ok := d.GetOk("trusthost10"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminTrusthost10(d, v, "trusthost10", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost10"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost1"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost1(d, v, "ip6_trusthost1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost1"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost2"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost2(d, v, "ip6_trusthost2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost2"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost3"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost3(d, v, "ip6_trusthost3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost3"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost4"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost4(d, v, "ip6_trusthost4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost4"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost5"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost5(d, v, "ip6_trusthost5", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost5"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost6"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost6(d, v, "ip6_trusthost6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost6"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost7"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost7(d, v, "ip6_trusthost7", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost7"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost8"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost8(d, v, "ip6_trusthost8", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost8"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost9"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost9(d, v, "ip6_trusthost9", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost9"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost10"); ok {
		t, err := expandSwitchControllerSecurityPolicyAdminIp6Trusthost10(d, v, "ip6_trusthost10", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost10"] = t
		}
	}

	return &obj, nil
}
