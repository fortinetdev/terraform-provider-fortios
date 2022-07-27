// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure access profiles for system administrators.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAccprofile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAccprofileCreate,
		Read:   resourceSystemAccprofileRead,
		Update: resourceSystemAccprofileUpdate,
		Delete: resourceSystemAccprofileDelete,

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
				Required:     true,
			},
			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"secfabgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftviewgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sysgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"loggrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpngrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"utmgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wanoptgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netgrp_permission": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cfg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"packet_capture": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_cfg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"sysgrp_permission": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"upd": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cfg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mnt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fwgrp_permission": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"schedule": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"others": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"loggrp_permission": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"data_access": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"report_access": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"threat_weight": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"utmgrp_permission": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"antivirus": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ips": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"webfilter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"emailfilter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"spamfilter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"data_loss_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"file_filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"application_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"icap": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"voip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"waf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dnsfilter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"endpoint_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"videofilter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"admintimeout_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admintimeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: intBetweenWithZero(1, 480),
				Optional:     true,
				Computed:     true,
			},
			"system_diagnostics": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAccprofileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemAccprofile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemAccprofile resource while getting object: %v", err)
	}

	o, err := c.CreateSystemAccprofile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemAccprofile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAccprofile")
	}

	return resourceSystemAccprofileRead(d, m)
}

func resourceSystemAccprofileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemAccprofile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAccprofile resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAccprofile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemAccprofile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAccprofile")
	}

	return resourceSystemAccprofileRead(d, m)
}

func resourceSystemAccprofileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemAccprofile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAccprofile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAccprofileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemAccprofile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemAccprofile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAccprofile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAccprofile resource from API: %v", err)
	}
	return nil
}

func flattenSystemAccprofileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileScope(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileSecfabgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileFtviewgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileAuthgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileSysgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileNetgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileLoggrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileFwgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileVpngrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileWanoptgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileWifi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileNetgrpPermission(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cfg"
	if _, ok := i["cfg"]; ok {

		result["cfg"] = flattenSystemAccprofileNetgrpPermissionCfg(i["cfg"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "packet_capture"
	if _, ok := i["packet-capture"]; ok {

		result["packet_capture"] = flattenSystemAccprofileNetgrpPermissionPacketCapture(i["packet-capture"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "route_cfg"
	if _, ok := i["route-cfg"]; ok {

		result["route_cfg"] = flattenSystemAccprofileNetgrpPermissionRouteCfg(i["route-cfg"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemAccprofileNetgrpPermissionCfg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileNetgrpPermissionPacketCapture(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileNetgrpPermissionRouteCfg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileSysgrpPermission(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "admin"
	if _, ok := i["admin"]; ok {

		result["admin"] = flattenSystemAccprofileSysgrpPermissionAdmin(i["admin"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "upd"
	if _, ok := i["upd"]; ok {

		result["upd"] = flattenSystemAccprofileSysgrpPermissionUpd(i["upd"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cfg"
	if _, ok := i["cfg"]; ok {

		result["cfg"] = flattenSystemAccprofileSysgrpPermissionCfg(i["cfg"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mnt"
	if _, ok := i["mnt"]; ok {

		result["mnt"] = flattenSystemAccprofileSysgrpPermissionMnt(i["mnt"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemAccprofileSysgrpPermissionAdmin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileSysgrpPermissionUpd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileSysgrpPermissionCfg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileSysgrpPermissionMnt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileFwgrpPermission(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "policy"
	if _, ok := i["policy"]; ok {

		result["policy"] = flattenSystemAccprofileFwgrpPermissionPolicy(i["policy"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "address"
	if _, ok := i["address"]; ok {

		result["address"] = flattenSystemAccprofileFwgrpPermissionAddress(i["address"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "service"
	if _, ok := i["service"]; ok {

		result["service"] = flattenSystemAccprofileFwgrpPermissionService(i["service"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "schedule"
	if _, ok := i["schedule"]; ok {

		result["schedule"] = flattenSystemAccprofileFwgrpPermissionSchedule(i["schedule"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "others"
	if _, ok := i["others"]; ok {

		result["others"] = flattenSystemAccprofileFwgrpPermissionOthers(i["others"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemAccprofileFwgrpPermissionPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileFwgrpPermissionAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileFwgrpPermissionService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileFwgrpPermissionSchedule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileFwgrpPermissionOthers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileLoggrpPermission(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "config"
	if _, ok := i["config"]; ok {

		result["config"] = flattenSystemAccprofileLoggrpPermissionConfig(i["config"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "data_access"
	if _, ok := i["data-access"]; ok {

		result["data_access"] = flattenSystemAccprofileLoggrpPermissionDataAccess(i["data-access"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "report_access"
	if _, ok := i["report-access"]; ok {

		result["report_access"] = flattenSystemAccprofileLoggrpPermissionReportAccess(i["report-access"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "threat_weight"
	if _, ok := i["threat-weight"]; ok {

		result["threat_weight"] = flattenSystemAccprofileLoggrpPermissionThreatWeight(i["threat-weight"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemAccprofileLoggrpPermissionConfig(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileLoggrpPermissionDataAccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileLoggrpPermissionReportAccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileLoggrpPermissionThreatWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermission(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "antivirus"
	if _, ok := i["antivirus"]; ok {

		result["antivirus"] = flattenSystemAccprofileUtmgrpPermissionAntivirus(i["antivirus"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ips"
	if _, ok := i["ips"]; ok {

		result["ips"] = flattenSystemAccprofileUtmgrpPermissionIps(i["ips"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "webfilter"
	if _, ok := i["webfilter"]; ok {

		result["webfilter"] = flattenSystemAccprofileUtmgrpPermissionWebfilter(i["webfilter"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "emailfilter"
	if _, ok := i["emailfilter"]; ok {

		result["emailfilter"] = flattenSystemAccprofileUtmgrpPermissionEmailfilter(i["emailfilter"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "spamfilter"
	if _, ok := i["spamfilter"]; ok {

		result["spamfilter"] = flattenSystemAccprofileUtmgrpPermissionSpamfilter(i["spamfilter"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "data_loss_prevention"
	if _, ok := i["data-loss-prevention"]; ok {

		result["data_loss_prevention"] = flattenSystemAccprofileUtmgrpPermissionDataLossPrevention(i["data-loss-prevention"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "file_filter"
	if _, ok := i["file-filter"]; ok {

		result["file_filter"] = flattenSystemAccprofileUtmgrpPermissionFileFilter(i["file-filter"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "application_control"
	if _, ok := i["application-control"]; ok {

		result["application_control"] = flattenSystemAccprofileUtmgrpPermissionApplicationControl(i["application-control"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "icap"
	if _, ok := i["icap"]; ok {

		result["icap"] = flattenSystemAccprofileUtmgrpPermissionIcap(i["icap"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "voip"
	if _, ok := i["voip"]; ok {

		result["voip"] = flattenSystemAccprofileUtmgrpPermissionVoip(i["voip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "waf"
	if _, ok := i["waf"]; ok {

		result["waf"] = flattenSystemAccprofileUtmgrpPermissionWaf(i["waf"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dnsfilter"
	if _, ok := i["dnsfilter"]; ok {

		result["dnsfilter"] = flattenSystemAccprofileUtmgrpPermissionDnsfilter(i["dnsfilter"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "endpoint_control"
	if _, ok := i["endpoint-control"]; ok {

		result["endpoint_control"] = flattenSystemAccprofileUtmgrpPermissionEndpointControl(i["endpoint-control"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "videofilter"
	if _, ok := i["videofilter"]; ok {

		result["videofilter"] = flattenSystemAccprofileUtmgrpPermissionVideofilter(i["videofilter"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemAccprofileUtmgrpPermissionAntivirus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionIps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionWebfilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionEmailfilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionSpamfilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionDataLossPrevention(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionFileFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionApplicationControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionIcap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionVoip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionWaf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionDnsfilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionEndpointControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionVideofilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileAdmintimeoutOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileAdmintimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAccprofileSystemDiagnostics(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAccprofile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemAccprofileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("scope", flattenSystemAccprofileScope(o["scope"], d, "scope", sv)); err != nil {
		if !fortiAPIPatch(o["scope"]) {
			return fmt.Errorf("Error reading scope: %v", err)
		}
	}

	if err = d.Set("comments", flattenSystemAccprofileComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("secfabgrp", flattenSystemAccprofileSecfabgrp(o["secfabgrp"], d, "secfabgrp", sv)); err != nil {
		if !fortiAPIPatch(o["secfabgrp"]) {
			return fmt.Errorf("Error reading secfabgrp: %v", err)
		}
	}

	if err = d.Set("ftviewgrp", flattenSystemAccprofileFtviewgrp(o["ftviewgrp"], d, "ftviewgrp", sv)); err != nil {
		if !fortiAPIPatch(o["ftviewgrp"]) {
			return fmt.Errorf("Error reading ftviewgrp: %v", err)
		}
	}

	if err = d.Set("authgrp", flattenSystemAccprofileAuthgrp(o["authgrp"], d, "authgrp", sv)); err != nil {
		if !fortiAPIPatch(o["authgrp"]) {
			return fmt.Errorf("Error reading authgrp: %v", err)
		}
	}

	if err = d.Set("sysgrp", flattenSystemAccprofileSysgrp(o["sysgrp"], d, "sysgrp", sv)); err != nil {
		if !fortiAPIPatch(o["sysgrp"]) {
			return fmt.Errorf("Error reading sysgrp: %v", err)
		}
	}

	if err = d.Set("netgrp", flattenSystemAccprofileNetgrp(o["netgrp"], d, "netgrp", sv)); err != nil {
		if !fortiAPIPatch(o["netgrp"]) {
			return fmt.Errorf("Error reading netgrp: %v", err)
		}
	}

	if err = d.Set("loggrp", flattenSystemAccprofileLoggrp(o["loggrp"], d, "loggrp", sv)); err != nil {
		if !fortiAPIPatch(o["loggrp"]) {
			return fmt.Errorf("Error reading loggrp: %v", err)
		}
	}

	if err = d.Set("fwgrp", flattenSystemAccprofileFwgrp(o["fwgrp"], d, "fwgrp", sv)); err != nil {
		if !fortiAPIPatch(o["fwgrp"]) {
			return fmt.Errorf("Error reading fwgrp: %v", err)
		}
	}

	if err = d.Set("vpngrp", flattenSystemAccprofileVpngrp(o["vpngrp"], d, "vpngrp", sv)); err != nil {
		if !fortiAPIPatch(o["vpngrp"]) {
			return fmt.Errorf("Error reading vpngrp: %v", err)
		}
	}

	if err = d.Set("utmgrp", flattenSystemAccprofileUtmgrp(o["utmgrp"], d, "utmgrp", sv)); err != nil {
		if !fortiAPIPatch(o["utmgrp"]) {
			return fmt.Errorf("Error reading utmgrp: %v", err)
		}
	}

	if err = d.Set("wanoptgrp", flattenSystemAccprofileWanoptgrp(o["wanoptgrp"], d, "wanoptgrp", sv)); err != nil {
		if !fortiAPIPatch(o["wanoptgrp"]) {
			return fmt.Errorf("Error reading wanoptgrp: %v", err)
		}
	}

	if err = d.Set("wifi", flattenSystemAccprofileWifi(o["wifi"], d, "wifi", sv)); err != nil {
		if !fortiAPIPatch(o["wifi"]) {
			return fmt.Errorf("Error reading wifi: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("netgrp_permission", flattenSystemAccprofileNetgrpPermission(o["netgrp-permission"], d, "netgrp_permission", sv)); err != nil {
			if !fortiAPIPatch(o["netgrp-permission"]) {
				return fmt.Errorf("Error reading netgrp_permission: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("netgrp_permission"); ok {
			if err = d.Set("netgrp_permission", flattenSystemAccprofileNetgrpPermission(o["netgrp-permission"], d, "netgrp_permission", sv)); err != nil {
				if !fortiAPIPatch(o["netgrp-permission"]) {
					return fmt.Errorf("Error reading netgrp_permission: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("sysgrp_permission", flattenSystemAccprofileSysgrpPermission(o["sysgrp-permission"], d, "sysgrp_permission", sv)); err != nil {
			if !fortiAPIPatch(o["sysgrp-permission"]) {
				return fmt.Errorf("Error reading sysgrp_permission: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sysgrp_permission"); ok {
			if err = d.Set("sysgrp_permission", flattenSystemAccprofileSysgrpPermission(o["sysgrp-permission"], d, "sysgrp_permission", sv)); err != nil {
				if !fortiAPIPatch(o["sysgrp-permission"]) {
					return fmt.Errorf("Error reading sysgrp_permission: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("fwgrp_permission", flattenSystemAccprofileFwgrpPermission(o["fwgrp-permission"], d, "fwgrp_permission", sv)); err != nil {
			if !fortiAPIPatch(o["fwgrp-permission"]) {
				return fmt.Errorf("Error reading fwgrp_permission: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fwgrp_permission"); ok {
			if err = d.Set("fwgrp_permission", flattenSystemAccprofileFwgrpPermission(o["fwgrp-permission"], d, "fwgrp_permission", sv)); err != nil {
				if !fortiAPIPatch(o["fwgrp-permission"]) {
					return fmt.Errorf("Error reading fwgrp_permission: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("loggrp_permission", flattenSystemAccprofileLoggrpPermission(o["loggrp-permission"], d, "loggrp_permission", sv)); err != nil {
			if !fortiAPIPatch(o["loggrp-permission"]) {
				return fmt.Errorf("Error reading loggrp_permission: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("loggrp_permission"); ok {
			if err = d.Set("loggrp_permission", flattenSystemAccprofileLoggrpPermission(o["loggrp-permission"], d, "loggrp_permission", sv)); err != nil {
				if !fortiAPIPatch(o["loggrp-permission"]) {
					return fmt.Errorf("Error reading loggrp_permission: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("utmgrp_permission", flattenSystemAccprofileUtmgrpPermission(o["utmgrp-permission"], d, "utmgrp_permission", sv)); err != nil {
			if !fortiAPIPatch(o["utmgrp-permission"]) {
				return fmt.Errorf("Error reading utmgrp_permission: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("utmgrp_permission"); ok {
			if err = d.Set("utmgrp_permission", flattenSystemAccprofileUtmgrpPermission(o["utmgrp-permission"], d, "utmgrp_permission", sv)); err != nil {
				if !fortiAPIPatch(o["utmgrp-permission"]) {
					return fmt.Errorf("Error reading utmgrp_permission: %v", err)
				}
			}
		}
	}

	if err = d.Set("admintimeout_override", flattenSystemAccprofileAdmintimeoutOverride(o["admintimeout-override"], d, "admintimeout_override", sv)); err != nil {
		if !fortiAPIPatch(o["admintimeout-override"]) {
			return fmt.Errorf("Error reading admintimeout_override: %v", err)
		}
	}

	if err = d.Set("admintimeout", flattenSystemAccprofileAdmintimeout(o["admintimeout"], d, "admintimeout", sv)); err != nil {
		if !fortiAPIPatch(o["admintimeout"]) {
			return fmt.Errorf("Error reading admintimeout: %v", err)
		}
	}

	if err = d.Set("system_diagnostics", flattenSystemAccprofileSystemDiagnostics(o["system-diagnostics"], d, "system_diagnostics", sv)); err != nil {
		if !fortiAPIPatch(o["system-diagnostics"]) {
			return fmt.Errorf("Error reading system_diagnostics: %v", err)
		}
	}

	return nil
}

func flattenSystemAccprofileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemAccprofileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileScope(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSecfabgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileFtviewgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileAuthgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSysgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileNetgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileLoggrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileFwgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileVpngrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileWanoptgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileWifi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileNetgrpPermission(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cfg"
	if _, ok := d.GetOk(pre_append); ok {

		result["cfg"], _ = expandSystemAccprofileNetgrpPermissionCfg(d, i["cfg"], pre_append, sv)
	}
	pre_append = pre + ".0." + "packet_capture"
	if _, ok := d.GetOk(pre_append); ok {

		result["packet-capture"], _ = expandSystemAccprofileNetgrpPermissionPacketCapture(d, i["packet_capture"], pre_append, sv)
	}
	pre_append = pre + ".0." + "route_cfg"
	if _, ok := d.GetOk(pre_append); ok {

		result["route-cfg"], _ = expandSystemAccprofileNetgrpPermissionRouteCfg(d, i["route_cfg"], pre_append, sv)
	}

	return result, nil
}

func expandSystemAccprofileNetgrpPermissionCfg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileNetgrpPermissionPacketCapture(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileNetgrpPermissionRouteCfg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSysgrpPermission(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "admin"
	if _, ok := d.GetOk(pre_append); ok {

		result["admin"], _ = expandSystemAccprofileSysgrpPermissionAdmin(d, i["admin"], pre_append, sv)
	}
	pre_append = pre + ".0." + "upd"
	if _, ok := d.GetOk(pre_append); ok {

		result["upd"], _ = expandSystemAccprofileSysgrpPermissionUpd(d, i["upd"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cfg"
	if _, ok := d.GetOk(pre_append); ok {

		result["cfg"], _ = expandSystemAccprofileSysgrpPermissionCfg(d, i["cfg"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mnt"
	if _, ok := d.GetOk(pre_append); ok {

		result["mnt"], _ = expandSystemAccprofileSysgrpPermissionMnt(d, i["mnt"], pre_append, sv)
	}

	return result, nil
}

func expandSystemAccprofileSysgrpPermissionAdmin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSysgrpPermissionUpd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSysgrpPermissionCfg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSysgrpPermissionMnt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileFwgrpPermission(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "policy"
	if _, ok := d.GetOk(pre_append); ok {

		result["policy"], _ = expandSystemAccprofileFwgrpPermissionPolicy(d, i["policy"], pre_append, sv)
	}
	pre_append = pre + ".0." + "address"
	if _, ok := d.GetOk(pre_append); ok {

		result["address"], _ = expandSystemAccprofileFwgrpPermissionAddress(d, i["address"], pre_append, sv)
	}
	pre_append = pre + ".0." + "service"
	if _, ok := d.GetOk(pre_append); ok {

		result["service"], _ = expandSystemAccprofileFwgrpPermissionService(d, i["service"], pre_append, sv)
	}
	pre_append = pre + ".0." + "schedule"
	if _, ok := d.GetOk(pre_append); ok {

		result["schedule"], _ = expandSystemAccprofileFwgrpPermissionSchedule(d, i["schedule"], pre_append, sv)
	}
	pre_append = pre + ".0." + "others"
	if _, ok := d.GetOk(pre_append); ok {

		result["others"], _ = expandSystemAccprofileFwgrpPermissionOthers(d, i["others"], pre_append, sv)
	}

	return result, nil
}

func expandSystemAccprofileFwgrpPermissionPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileFwgrpPermissionAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileFwgrpPermissionService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileFwgrpPermissionSchedule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileFwgrpPermissionOthers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileLoggrpPermission(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "config"
	if _, ok := d.GetOk(pre_append); ok {

		result["config"], _ = expandSystemAccprofileLoggrpPermissionConfig(d, i["config"], pre_append, sv)
	}
	pre_append = pre + ".0." + "data_access"
	if _, ok := d.GetOk(pre_append); ok {

		result["data-access"], _ = expandSystemAccprofileLoggrpPermissionDataAccess(d, i["data_access"], pre_append, sv)
	}
	pre_append = pre + ".0." + "report_access"
	if _, ok := d.GetOk(pre_append); ok {

		result["report-access"], _ = expandSystemAccprofileLoggrpPermissionReportAccess(d, i["report_access"], pre_append, sv)
	}
	pre_append = pre + ".0." + "threat_weight"
	if _, ok := d.GetOk(pre_append); ok {

		result["threat-weight"], _ = expandSystemAccprofileLoggrpPermissionThreatWeight(d, i["threat_weight"], pre_append, sv)
	}

	return result, nil
}

func expandSystemAccprofileLoggrpPermissionConfig(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileLoggrpPermissionDataAccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileLoggrpPermissionReportAccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileLoggrpPermissionThreatWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermission(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "antivirus"
	if _, ok := d.GetOk(pre_append); ok {

		result["antivirus"], _ = expandSystemAccprofileUtmgrpPermissionAntivirus(d, i["antivirus"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ips"
	if _, ok := d.GetOk(pre_append); ok {

		result["ips"], _ = expandSystemAccprofileUtmgrpPermissionIps(d, i["ips"], pre_append, sv)
	}
	pre_append = pre + ".0." + "webfilter"
	if _, ok := d.GetOk(pre_append); ok {

		result["webfilter"], _ = expandSystemAccprofileUtmgrpPermissionWebfilter(d, i["webfilter"], pre_append, sv)
	}
	pre_append = pre + ".0." + "emailfilter"
	if _, ok := d.GetOk(pre_append); ok {

		result["emailfilter"], _ = expandSystemAccprofileUtmgrpPermissionEmailfilter(d, i["emailfilter"], pre_append, sv)
	}
	pre_append = pre + ".0." + "spamfilter"
	if _, ok := d.GetOk(pre_append); ok {

		result["spamfilter"], _ = expandSystemAccprofileUtmgrpPermissionSpamfilter(d, i["spamfilter"], pre_append, sv)
	}
	pre_append = pre + ".0." + "data_loss_prevention"
	if _, ok := d.GetOk(pre_append); ok {

		result["data-loss-prevention"], _ = expandSystemAccprofileUtmgrpPermissionDataLossPrevention(d, i["data_loss_prevention"], pre_append, sv)
	}
	pre_append = pre + ".0." + "file_filter"
	if _, ok := d.GetOk(pre_append); ok {

		result["file-filter"], _ = expandSystemAccprofileUtmgrpPermissionFileFilter(d, i["file_filter"], pre_append, sv)
	}
	pre_append = pre + ".0." + "application_control"
	if _, ok := d.GetOk(pre_append); ok {

		result["application-control"], _ = expandSystemAccprofileUtmgrpPermissionApplicationControl(d, i["application_control"], pre_append, sv)
	}
	pre_append = pre + ".0." + "icap"
	if _, ok := d.GetOk(pre_append); ok {

		result["icap"], _ = expandSystemAccprofileUtmgrpPermissionIcap(d, i["icap"], pre_append, sv)
	}
	pre_append = pre + ".0." + "voip"
	if _, ok := d.GetOk(pre_append); ok {

		result["voip"], _ = expandSystemAccprofileUtmgrpPermissionVoip(d, i["voip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "waf"
	if _, ok := d.GetOk(pre_append); ok {

		result["waf"], _ = expandSystemAccprofileUtmgrpPermissionWaf(d, i["waf"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dnsfilter"
	if _, ok := d.GetOk(pre_append); ok {

		result["dnsfilter"], _ = expandSystemAccprofileUtmgrpPermissionDnsfilter(d, i["dnsfilter"], pre_append, sv)
	}
	pre_append = pre + ".0." + "endpoint_control"
	if _, ok := d.GetOk(pre_append); ok {

		result["endpoint-control"], _ = expandSystemAccprofileUtmgrpPermissionEndpointControl(d, i["endpoint_control"], pre_append, sv)
	}
	pre_append = pre + ".0." + "videofilter"
	if _, ok := d.GetOk(pre_append); ok {

		result["videofilter"], _ = expandSystemAccprofileUtmgrpPermissionVideofilter(d, i["videofilter"], pre_append, sv)
	}

	return result, nil
}

func expandSystemAccprofileUtmgrpPermissionAntivirus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionIps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionWebfilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionEmailfilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionSpamfilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionDataLossPrevention(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionFileFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionApplicationControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionIcap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionVoip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionWaf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionDnsfilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionEndpointControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionVideofilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileAdmintimeoutOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileAdmintimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSystemDiagnostics(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAccprofile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemAccprofileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("scope"); ok {

		t, err := expandSystemAccprofileScope(d, v, "scope", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scope"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandSystemAccprofileComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("secfabgrp"); ok {

		t, err := expandSystemAccprofileSecfabgrp(d, v, "secfabgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secfabgrp"] = t
		}
	}

	if v, ok := d.GetOk("ftviewgrp"); ok {

		t, err := expandSystemAccprofileFtviewgrp(d, v, "ftviewgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftviewgrp"] = t
		}
	}

	if v, ok := d.GetOk("authgrp"); ok {

		t, err := expandSystemAccprofileAuthgrp(d, v, "authgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authgrp"] = t
		}
	}

	if v, ok := d.GetOk("sysgrp"); ok {

		t, err := expandSystemAccprofileSysgrp(d, v, "sysgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sysgrp"] = t
		}
	}

	if v, ok := d.GetOk("netgrp"); ok {

		t, err := expandSystemAccprofileNetgrp(d, v, "netgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netgrp"] = t
		}
	}

	if v, ok := d.GetOk("loggrp"); ok {

		t, err := expandSystemAccprofileLoggrp(d, v, "loggrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loggrp"] = t
		}
	}

	if v, ok := d.GetOk("fwgrp"); ok {

		t, err := expandSystemAccprofileFwgrp(d, v, "fwgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwgrp"] = t
		}
	}

	if v, ok := d.GetOk("vpngrp"); ok {

		t, err := expandSystemAccprofileVpngrp(d, v, "vpngrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpngrp"] = t
		}
	}

	if v, ok := d.GetOk("utmgrp"); ok {

		t, err := expandSystemAccprofileUtmgrp(d, v, "utmgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utmgrp"] = t
		}
	}

	if v, ok := d.GetOk("wanoptgrp"); ok {

		t, err := expandSystemAccprofileWanoptgrp(d, v, "wanoptgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanoptgrp"] = t
		}
	}

	if v, ok := d.GetOk("wifi"); ok {

		t, err := expandSystemAccprofileWifi(d, v, "wifi", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi"] = t
		}
	}

	if v, ok := d.GetOk("netgrp_permission"); ok {

		t, err := expandSystemAccprofileNetgrpPermission(d, v, "netgrp_permission", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netgrp-permission"] = t
		}
	}

	if v, ok := d.GetOk("sysgrp_permission"); ok {

		t, err := expandSystemAccprofileSysgrpPermission(d, v, "sysgrp_permission", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sysgrp-permission"] = t
		}
	}

	if v, ok := d.GetOk("fwgrp_permission"); ok {

		t, err := expandSystemAccprofileFwgrpPermission(d, v, "fwgrp_permission", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwgrp-permission"] = t
		}
	}

	if v, ok := d.GetOk("loggrp_permission"); ok {

		t, err := expandSystemAccprofileLoggrpPermission(d, v, "loggrp_permission", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loggrp-permission"] = t
		}
	}

	if v, ok := d.GetOk("utmgrp_permission"); ok {

		t, err := expandSystemAccprofileUtmgrpPermission(d, v, "utmgrp_permission", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utmgrp-permission"] = t
		}
	}

	if v, ok := d.GetOk("admintimeout_override"); ok {

		t, err := expandSystemAccprofileAdmintimeoutOverride(d, v, "admintimeout_override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admintimeout-override"] = t
		}
	}

	if v, ok := d.GetOkExists("admintimeout"); ok {

		t, err := expandSystemAccprofileAdmintimeout(d, v, "admintimeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admintimeout"] = t
		}
	}

	if v, ok := d.GetOk("system_diagnostics"); ok {

		t, err := expandSystemAccprofileSystemDiagnostics(d, v, "system_diagnostics", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-diagnostics"] = t
		}
	}

	return &obj, nil
}
