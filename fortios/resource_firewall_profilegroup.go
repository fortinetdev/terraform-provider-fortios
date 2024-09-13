// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure profile groups.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallProfileGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallProfileGroupCreate,
		Read:   resourceFirewallProfileGroupRead,
		Update: resourceFirewallProfileGroupUpdate,
		Delete: resourceFirewallProfileGroupDelete,

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
				Required:     true,
			},
			"av_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"webfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"dnsfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"emailfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"dlp_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"spamfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"dlp_sensor": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"file_filter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"ips_sensor": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"application_list": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"voip_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"ips_voip_filter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"sctp_filter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"diameter_filter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"virtual_patch_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"icap_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"cifs_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"videofilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"waf_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"ssh_filter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"casb_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"profile_protocol_options": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ssl_ssh_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceFirewallProfileGroupCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallProfileGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallProfileGroup resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallProfileGroup(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallProfileGroup resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallProfileGroup")
	}

	return resourceFirewallProfileGroupRead(d, m)
}

func resourceFirewallProfileGroupUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallProfileGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileGroup resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallProfileGroup(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallProfileGroup")
	}

	return resourceFirewallProfileGroupRead(d, m)
}

func resourceFirewallProfileGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallProfileGroup(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallProfileGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallProfileGroupRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallProfileGroup(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProfileGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallProfileGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProfileGroup resource from API: %v", err)
	}
	return nil
}

func flattenFirewallProfileGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupAvProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupWebfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupDnsfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupEmailfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupDlpProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupSpamfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupDlpSensor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupFileFilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupIpsSensor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupApplicationList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupVoipProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupIpsVoipFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupSctpFilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupDiameterFilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupVirtualPatchProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupIcapProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupCifsProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupVideofilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupWafProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupSshFilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupCasbProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupProfileProtocolOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProfileGroupSslSshProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallProfileGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenFirewallProfileGroupName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenFirewallProfileGroupAvProfile(o["av-profile"], d, "av_profile", sv)); err != nil {
		if !fortiAPIPatch(o["av-profile"]) {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenFirewallProfileGroupWebfilterProfile(o["webfilter-profile"], d, "webfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-profile"]) {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("dnsfilter_profile", flattenFirewallProfileGroupDnsfilterProfile(o["dnsfilter-profile"], d, "dnsfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["dnsfilter-profile"]) {
			return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile", flattenFirewallProfileGroupEmailfilterProfile(o["emailfilter-profile"], d, "emailfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["emailfilter-profile"]) {
			return fmt.Errorf("Error reading emailfilter_profile: %v", err)
		}
	}

	if err = d.Set("dlp_profile", flattenFirewallProfileGroupDlpProfile(o["dlp-profile"], d, "dlp_profile", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-profile"]) {
			return fmt.Errorf("Error reading dlp_profile: %v", err)
		}
	}

	if err = d.Set("spamfilter_profile", flattenFirewallProfileGroupSpamfilterProfile(o["spamfilter-profile"], d, "spamfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["spamfilter-profile"]) {
			return fmt.Errorf("Error reading spamfilter_profile: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenFirewallProfileGroupDlpSensor(o["dlp-sensor"], d, "dlp_sensor", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-sensor"]) {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("file_filter_profile", flattenFirewallProfileGroupFileFilterProfile(o["file-filter-profile"], d, "file_filter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["file-filter-profile"]) {
			return fmt.Errorf("Error reading file_filter_profile: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenFirewallProfileGroupIpsSensor(o["ips-sensor"], d, "ips_sensor", sv)); err != nil {
		if !fortiAPIPatch(o["ips-sensor"]) {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("application_list", flattenFirewallProfileGroupApplicationList(o["application-list"], d, "application_list", sv)); err != nil {
		if !fortiAPIPatch(o["application-list"]) {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("voip_profile", flattenFirewallProfileGroupVoipProfile(o["voip-profile"], d, "voip_profile", sv)); err != nil {
		if !fortiAPIPatch(o["voip-profile"]) {
			return fmt.Errorf("Error reading voip_profile: %v", err)
		}
	}

	if err = d.Set("ips_voip_filter", flattenFirewallProfileGroupIpsVoipFilter(o["ips-voip-filter"], d, "ips_voip_filter", sv)); err != nil {
		if !fortiAPIPatch(o["ips-voip-filter"]) {
			return fmt.Errorf("Error reading ips_voip_filter: %v", err)
		}
	}

	if err = d.Set("sctp_filter_profile", flattenFirewallProfileGroupSctpFilterProfile(o["sctp-filter-profile"], d, "sctp_filter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["sctp-filter-profile"]) {
			return fmt.Errorf("Error reading sctp_filter_profile: %v", err)
		}
	}

	if err = d.Set("diameter_filter_profile", flattenFirewallProfileGroupDiameterFilterProfile(o["diameter-filter-profile"], d, "diameter_filter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["diameter-filter-profile"]) {
			return fmt.Errorf("Error reading diameter_filter_profile: %v", err)
		}
	}

	if err = d.Set("virtual_patch_profile", flattenFirewallProfileGroupVirtualPatchProfile(o["virtual-patch-profile"], d, "virtual_patch_profile", sv)); err != nil {
		if !fortiAPIPatch(o["virtual-patch-profile"]) {
			return fmt.Errorf("Error reading virtual_patch_profile: %v", err)
		}
	}

	if err = d.Set("icap_profile", flattenFirewallProfileGroupIcapProfile(o["icap-profile"], d, "icap_profile", sv)); err != nil {
		if !fortiAPIPatch(o["icap-profile"]) {
			return fmt.Errorf("Error reading icap_profile: %v", err)
		}
	}

	if err = d.Set("cifs_profile", flattenFirewallProfileGroupCifsProfile(o["cifs-profile"], d, "cifs_profile", sv)); err != nil {
		if !fortiAPIPatch(o["cifs-profile"]) {
			return fmt.Errorf("Error reading cifs_profile: %v", err)
		}
	}

	if err = d.Set("videofilter_profile", flattenFirewallProfileGroupVideofilterProfile(o["videofilter-profile"], d, "videofilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["videofilter-profile"]) {
			return fmt.Errorf("Error reading videofilter_profile: %v", err)
		}
	}

	if err = d.Set("waf_profile", flattenFirewallProfileGroupWafProfile(o["waf-profile"], d, "waf_profile", sv)); err != nil {
		if !fortiAPIPatch(o["waf-profile"]) {
			return fmt.Errorf("Error reading waf_profile: %v", err)
		}
	}

	if err = d.Set("ssh_filter_profile", flattenFirewallProfileGroupSshFilterProfile(o["ssh-filter-profile"], d, "ssh_filter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-filter-profile"]) {
			return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
		}
	}

	if err = d.Set("casb_profile", flattenFirewallProfileGroupCasbProfile(o["casb-profile"], d, "casb_profile", sv)); err != nil {
		if !fortiAPIPatch(o["casb-profile"]) {
			return fmt.Errorf("Error reading casb_profile: %v", err)
		}
	}

	if err = d.Set("profile_protocol_options", flattenFirewallProfileGroupProfileProtocolOptions(o["profile-protocol-options"], d, "profile_protocol_options", sv)); err != nil {
		if !fortiAPIPatch(o["profile-protocol-options"]) {
			return fmt.Errorf("Error reading profile_protocol_options: %v", err)
		}
	}

	if err = d.Set("ssl_ssh_profile", flattenFirewallProfileGroupSslSshProfile(o["ssl-ssh-profile"], d, "ssl_ssh_profile", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-ssh-profile"]) {
			return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
		}
	}

	return nil
}

func flattenFirewallProfileGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallProfileGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupAvProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupWebfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupDnsfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupEmailfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupDlpProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupSpamfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupDlpSensor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupFileFilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupIpsSensor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupApplicationList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupVoipProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupIpsVoipFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupSctpFilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupDiameterFilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupVirtualPatchProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupIcapProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupCifsProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupVideofilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupWafProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupSshFilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupCasbProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupProfileProtocolOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupSslSshProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallProfileGroup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallProfileGroupName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	} else if d.HasChange("name") {
		obj["name"] = nil
	}

	if v, ok := d.GetOk("av_profile"); ok {
		t, err := expandFirewallProfileGroupAvProfile(d, v, "av_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	} else if d.HasChange("av_profile") {
		obj["av-profile"] = nil
	}

	if v, ok := d.GetOk("webfilter_profile"); ok {
		t, err := expandFirewallProfileGroupWebfilterProfile(d, v, "webfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	} else if d.HasChange("webfilter_profile") {
		obj["webfilter-profile"] = nil
	}

	if v, ok := d.GetOk("dnsfilter_profile"); ok {
		t, err := expandFirewallProfileGroupDnsfilterProfile(d, v, "dnsfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnsfilter-profile"] = t
		}
	} else if d.HasChange("dnsfilter_profile") {
		obj["dnsfilter-profile"] = nil
	}

	if v, ok := d.GetOk("emailfilter_profile"); ok {
		t, err := expandFirewallProfileGroupEmailfilterProfile(d, v, "emailfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile"] = t
		}
	} else if d.HasChange("emailfilter_profile") {
		obj["emailfilter-profile"] = nil
	}

	if v, ok := d.GetOk("dlp_profile"); ok {
		t, err := expandFirewallProfileGroupDlpProfile(d, v, "dlp_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-profile"] = t
		}
	} else if d.HasChange("dlp_profile") {
		obj["dlp-profile"] = nil
	}

	if v, ok := d.GetOk("spamfilter_profile"); ok {
		t, err := expandFirewallProfileGroupSpamfilterProfile(d, v, "spamfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spamfilter-profile"] = t
		}
	} else if d.HasChange("spamfilter_profile") {
		obj["spamfilter-profile"] = nil
	}

	if v, ok := d.GetOk("dlp_sensor"); ok {
		t, err := expandFirewallProfileGroupDlpSensor(d, v, "dlp_sensor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	} else if d.HasChange("dlp_sensor") {
		obj["dlp-sensor"] = nil
	}

	if v, ok := d.GetOk("file_filter_profile"); ok {
		t, err := expandFirewallProfileGroupFileFilterProfile(d, v, "file_filter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-filter-profile"] = t
		}
	} else if d.HasChange("file_filter_profile") {
		obj["file-filter-profile"] = nil
	}

	if v, ok := d.GetOk("ips_sensor"); ok {
		t, err := expandFirewallProfileGroupIpsSensor(d, v, "ips_sensor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	} else if d.HasChange("ips_sensor") {
		obj["ips-sensor"] = nil
	}

	if v, ok := d.GetOk("application_list"); ok {
		t, err := expandFirewallProfileGroupApplicationList(d, v, "application_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	} else if d.HasChange("application_list") {
		obj["application-list"] = nil
	}

	if v, ok := d.GetOk("voip_profile"); ok {
		t, err := expandFirewallProfileGroupVoipProfile(d, v, "voip_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip-profile"] = t
		}
	} else if d.HasChange("voip_profile") {
		obj["voip-profile"] = nil
	}

	if v, ok := d.GetOk("ips_voip_filter"); ok {
		t, err := expandFirewallProfileGroupIpsVoipFilter(d, v, "ips_voip_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-voip-filter"] = t
		}
	} else if d.HasChange("ips_voip_filter") {
		obj["ips-voip-filter"] = nil
	}

	if v, ok := d.GetOk("sctp_filter_profile"); ok {
		t, err := expandFirewallProfileGroupSctpFilterProfile(d, v, "sctp_filter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-filter-profile"] = t
		}
	} else if d.HasChange("sctp_filter_profile") {
		obj["sctp-filter-profile"] = nil
	}

	if v, ok := d.GetOk("diameter_filter_profile"); ok {
		t, err := expandFirewallProfileGroupDiameterFilterProfile(d, v, "diameter_filter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diameter-filter-profile"] = t
		}
	} else if d.HasChange("diameter_filter_profile") {
		obj["diameter-filter-profile"] = nil
	}

	if v, ok := d.GetOk("virtual_patch_profile"); ok {
		t, err := expandFirewallProfileGroupVirtualPatchProfile(d, v, "virtual_patch_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-patch-profile"] = t
		}
	} else if d.HasChange("virtual_patch_profile") {
		obj["virtual-patch-profile"] = nil
	}

	if v, ok := d.GetOk("icap_profile"); ok {
		t, err := expandFirewallProfileGroupIcapProfile(d, v, "icap_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-profile"] = t
		}
	} else if d.HasChange("icap_profile") {
		obj["icap-profile"] = nil
	}

	if v, ok := d.GetOk("cifs_profile"); ok {
		t, err := expandFirewallProfileGroupCifsProfile(d, v, "cifs_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs-profile"] = t
		}
	} else if d.HasChange("cifs_profile") {
		obj["cifs-profile"] = nil
	}

	if v, ok := d.GetOk("videofilter_profile"); ok {
		t, err := expandFirewallProfileGroupVideofilterProfile(d, v, "videofilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["videofilter-profile"] = t
		}
	} else if d.HasChange("videofilter_profile") {
		obj["videofilter-profile"] = nil
	}

	if v, ok := d.GetOk("waf_profile"); ok {
		t, err := expandFirewallProfileGroupWafProfile(d, v, "waf_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["waf-profile"] = t
		}
	} else if d.HasChange("waf_profile") {
		obj["waf-profile"] = nil
	}

	if v, ok := d.GetOk("ssh_filter_profile"); ok {
		t, err := expandFirewallProfileGroupSshFilterProfile(d, v, "ssh_filter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-filter-profile"] = t
		}
	} else if d.HasChange("ssh_filter_profile") {
		obj["ssh-filter-profile"] = nil
	}

	if v, ok := d.GetOk("casb_profile"); ok {
		t, err := expandFirewallProfileGroupCasbProfile(d, v, "casb_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["casb-profile"] = t
		}
	} else if d.HasChange("casb_profile") {
		obj["casb-profile"] = nil
	}

	if v, ok := d.GetOk("profile_protocol_options"); ok {
		t, err := expandFirewallProfileGroupProfileProtocolOptions(d, v, "profile_protocol_options", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-protocol-options"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ssh_profile"); ok {
		t, err := expandFirewallProfileGroupSslSshProfile(d, v, "ssl_ssh_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-ssh-profile"] = t
		}
	}

	return &obj, nil
}
