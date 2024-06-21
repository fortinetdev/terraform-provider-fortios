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

func dataSourceSystemAccprofile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemAccprofileRead,
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
			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"secfabgrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ftviewgrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"authgrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sysgrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"netgrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"loggrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fwgrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vpngrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"utmgrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wanoptgrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wifi": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"netgrp_permission": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cfg": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"packet_capture": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_cfg": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"sysgrp_permission": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"upd": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"cfg": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"mnt": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"fwgrp_permission": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"service": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"schedule": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"others": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"loggrp_permission": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"data_access": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"report_access": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"threat_weight": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"utmgrp_permission": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"antivirus": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ips": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"webfilter": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"emailfilter": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dlp": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"data_leak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"spamfilter": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"data_loss_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"file_filter": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"application_control": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"icap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"voip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"waf": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dnsfilter": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"endpoint_control": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"videofilter": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"virtual_patch": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"casb": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"admintimeout_override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admintimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cli_diagnose": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cli_get": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cli_show": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cli_exec": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cli_config": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_diagnostics": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_execute_ssh": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_execute_telnet": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemAccprofileRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemAccprofile: type error")
	}

	o, err := c.ReadSystemAccprofile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemAccprofile: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemAccprofile(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemAccprofile from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemAccprofileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileScope(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileSecfabgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileFtviewgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileAuthgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileSysgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileNetgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileLoggrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileFwgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileVpngrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileWanoptgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileWifi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileNetgrpPermission(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cfg"
	if _, ok := i["cfg"]; ok {
		result["cfg"] = dataSourceFlattenSystemAccprofileNetgrpPermissionCfg(i["cfg"], d, pre_append)
	}

	pre_append = pre + ".0." + "packet_capture"
	if _, ok := i["packet-capture"]; ok {
		result["packet_capture"] = dataSourceFlattenSystemAccprofileNetgrpPermissionPacketCapture(i["packet-capture"], d, pre_append)
	}

	pre_append = pre + ".0." + "route_cfg"
	if _, ok := i["route-cfg"]; ok {
		result["route_cfg"] = dataSourceFlattenSystemAccprofileNetgrpPermissionRouteCfg(i["route-cfg"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemAccprofileNetgrpPermissionCfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileNetgrpPermissionPacketCapture(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileNetgrpPermissionRouteCfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileSysgrpPermission(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "admin"
	if _, ok := i["admin"]; ok {
		result["admin"] = dataSourceFlattenSystemAccprofileSysgrpPermissionAdmin(i["admin"], d, pre_append)
	}

	pre_append = pre + ".0." + "upd"
	if _, ok := i["upd"]; ok {
		result["upd"] = dataSourceFlattenSystemAccprofileSysgrpPermissionUpd(i["upd"], d, pre_append)
	}

	pre_append = pre + ".0." + "cfg"
	if _, ok := i["cfg"]; ok {
		result["cfg"] = dataSourceFlattenSystemAccprofileSysgrpPermissionCfg(i["cfg"], d, pre_append)
	}

	pre_append = pre + ".0." + "mnt"
	if _, ok := i["mnt"]; ok {
		result["mnt"] = dataSourceFlattenSystemAccprofileSysgrpPermissionMnt(i["mnt"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemAccprofileSysgrpPermissionAdmin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileSysgrpPermissionUpd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileSysgrpPermissionCfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileSysgrpPermissionMnt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileFwgrpPermission(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "policy"
	if _, ok := i["policy"]; ok {
		result["policy"] = dataSourceFlattenSystemAccprofileFwgrpPermissionPolicy(i["policy"], d, pre_append)
	}

	pre_append = pre + ".0." + "address"
	if _, ok := i["address"]; ok {
		result["address"] = dataSourceFlattenSystemAccprofileFwgrpPermissionAddress(i["address"], d, pre_append)
	}

	pre_append = pre + ".0." + "service"
	if _, ok := i["service"]; ok {
		result["service"] = dataSourceFlattenSystemAccprofileFwgrpPermissionService(i["service"], d, pre_append)
	}

	pre_append = pre + ".0." + "schedule"
	if _, ok := i["schedule"]; ok {
		result["schedule"] = dataSourceFlattenSystemAccprofileFwgrpPermissionSchedule(i["schedule"], d, pre_append)
	}

	pre_append = pre + ".0." + "others"
	if _, ok := i["others"]; ok {
		result["others"] = dataSourceFlattenSystemAccprofileFwgrpPermissionOthers(i["others"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemAccprofileFwgrpPermissionPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileFwgrpPermissionAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileFwgrpPermissionService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileFwgrpPermissionSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileFwgrpPermissionOthers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileLoggrpPermission(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "config"
	if _, ok := i["config"]; ok {
		result["config"] = dataSourceFlattenSystemAccprofileLoggrpPermissionConfig(i["config"], d, pre_append)
	}

	pre_append = pre + ".0." + "data_access"
	if _, ok := i["data-access"]; ok {
		result["data_access"] = dataSourceFlattenSystemAccprofileLoggrpPermissionDataAccess(i["data-access"], d, pre_append)
	}

	pre_append = pre + ".0." + "report_access"
	if _, ok := i["report-access"]; ok {
		result["report_access"] = dataSourceFlattenSystemAccprofileLoggrpPermissionReportAccess(i["report-access"], d, pre_append)
	}

	pre_append = pre + ".0." + "threat_weight"
	if _, ok := i["threat-weight"]; ok {
		result["threat_weight"] = dataSourceFlattenSystemAccprofileLoggrpPermissionThreatWeight(i["threat-weight"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemAccprofileLoggrpPermissionConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileLoggrpPermissionDataAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileLoggrpPermissionReportAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileLoggrpPermissionThreatWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermission(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "antivirus"
	if _, ok := i["antivirus"]; ok {
		result["antivirus"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionAntivirus(i["antivirus"], d, pre_append)
	}

	pre_append = pre + ".0." + "ips"
	if _, ok := i["ips"]; ok {
		result["ips"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionIps(i["ips"], d, pre_append)
	}

	pre_append = pre + ".0." + "webfilter"
	if _, ok := i["webfilter"]; ok {
		result["webfilter"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionWebfilter(i["webfilter"], d, pre_append)
	}

	pre_append = pre + ".0." + "emailfilter"
	if _, ok := i["emailfilter"]; ok {
		result["emailfilter"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionEmailfilter(i["emailfilter"], d, pre_append)
	}

	pre_append = pre + ".0." + "dlp"
	if _, ok := i["dlp"]; ok {
		result["dlp"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionDlp(i["dlp"], d, pre_append)
	}

	pre_append = pre + ".0." + "data_leak_prevention"
	if _, ok := i["data-leak-prevention"]; ok {
		result["data_leak_prevention"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionDataLeakPrevention(i["data-leak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "spamfilter"
	if _, ok := i["spamfilter"]; ok {
		result["spamfilter"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionSpamfilter(i["spamfilter"], d, pre_append)
	}

	pre_append = pre + ".0." + "data_loss_prevention"
	if _, ok := i["data-loss-prevention"]; ok {
		result["data_loss_prevention"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionDataLossPrevention(i["data-loss-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "file_filter"
	if _, ok := i["file-filter"]; ok {
		result["file_filter"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionFileFilter(i["file-filter"], d, pre_append)
	}

	pre_append = pre + ".0." + "application_control"
	if _, ok := i["application-control"]; ok {
		result["application_control"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionApplicationControl(i["application-control"], d, pre_append)
	}

	pre_append = pre + ".0." + "icap"
	if _, ok := i["icap"]; ok {
		result["icap"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionIcap(i["icap"], d, pre_append)
	}

	pre_append = pre + ".0." + "voip"
	if _, ok := i["voip"]; ok {
		result["voip"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionVoip(i["voip"], d, pre_append)
	}

	pre_append = pre + ".0." + "waf"
	if _, ok := i["waf"]; ok {
		result["waf"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionWaf(i["waf"], d, pre_append)
	}

	pre_append = pre + ".0." + "dnsfilter"
	if _, ok := i["dnsfilter"]; ok {
		result["dnsfilter"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionDnsfilter(i["dnsfilter"], d, pre_append)
	}

	pre_append = pre + ".0." + "endpoint_control"
	if _, ok := i["endpoint-control"]; ok {
		result["endpoint_control"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionEndpointControl(i["endpoint-control"], d, pre_append)
	}

	pre_append = pre + ".0." + "videofilter"
	if _, ok := i["videofilter"]; ok {
		result["videofilter"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionVideofilter(i["videofilter"], d, pre_append)
	}

	pre_append = pre + ".0." + "virtual_patch"
	if _, ok := i["virtual-patch"]; ok {
		result["virtual_patch"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionVirtualPatch(i["virtual-patch"], d, pre_append)
	}

	pre_append = pre + ".0." + "casb"
	if _, ok := i["casb"]; ok {
		result["casb"] = dataSourceFlattenSystemAccprofileUtmgrpPermissionCasb(i["casb"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionAntivirus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionIps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionWebfilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionEmailfilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionDlp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionDataLeakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionSpamfilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionDataLossPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionFileFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionApplicationControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionIcap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionVoip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionWaf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionDnsfilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionEndpointControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionVideofilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionVirtualPatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileUtmgrpPermissionCasb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileAdmintimeoutOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileAdmintimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileCliDiagnose(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileCliGet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileCliShow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileCliExec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileCliConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileSystemDiagnostics(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileSystemExecuteSsh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAccprofileSystemExecuteTelnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemAccprofile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemAccprofileName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("scope", dataSourceFlattenSystemAccprofileScope(o["scope"], d, "scope")); err != nil {
		if !fortiAPIPatch(o["scope"]) {
			return fmt.Errorf("Error reading scope: %v", err)
		}
	}

	if err = d.Set("comments", dataSourceFlattenSystemAccprofileComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("secfabgrp", dataSourceFlattenSystemAccprofileSecfabgrp(o["secfabgrp"], d, "secfabgrp")); err != nil {
		if !fortiAPIPatch(o["secfabgrp"]) {
			return fmt.Errorf("Error reading secfabgrp: %v", err)
		}
	}

	if err = d.Set("ftviewgrp", dataSourceFlattenSystemAccprofileFtviewgrp(o["ftviewgrp"], d, "ftviewgrp")); err != nil {
		if !fortiAPIPatch(o["ftviewgrp"]) {
			return fmt.Errorf("Error reading ftviewgrp: %v", err)
		}
	}

	if err = d.Set("authgrp", dataSourceFlattenSystemAccprofileAuthgrp(o["authgrp"], d, "authgrp")); err != nil {
		if !fortiAPIPatch(o["authgrp"]) {
			return fmt.Errorf("Error reading authgrp: %v", err)
		}
	}

	if err = d.Set("sysgrp", dataSourceFlattenSystemAccprofileSysgrp(o["sysgrp"], d, "sysgrp")); err != nil {
		if !fortiAPIPatch(o["sysgrp"]) {
			return fmt.Errorf("Error reading sysgrp: %v", err)
		}
	}

	if err = d.Set("netgrp", dataSourceFlattenSystemAccprofileNetgrp(o["netgrp"], d, "netgrp")); err != nil {
		if !fortiAPIPatch(o["netgrp"]) {
			return fmt.Errorf("Error reading netgrp: %v", err)
		}
	}

	if err = d.Set("loggrp", dataSourceFlattenSystemAccprofileLoggrp(o["loggrp"], d, "loggrp")); err != nil {
		if !fortiAPIPatch(o["loggrp"]) {
			return fmt.Errorf("Error reading loggrp: %v", err)
		}
	}

	if err = d.Set("fwgrp", dataSourceFlattenSystemAccprofileFwgrp(o["fwgrp"], d, "fwgrp")); err != nil {
		if !fortiAPIPatch(o["fwgrp"]) {
			return fmt.Errorf("Error reading fwgrp: %v", err)
		}
	}

	if err = d.Set("vpngrp", dataSourceFlattenSystemAccprofileVpngrp(o["vpngrp"], d, "vpngrp")); err != nil {
		if !fortiAPIPatch(o["vpngrp"]) {
			return fmt.Errorf("Error reading vpngrp: %v", err)
		}
	}

	if err = d.Set("utmgrp", dataSourceFlattenSystemAccprofileUtmgrp(o["utmgrp"], d, "utmgrp")); err != nil {
		if !fortiAPIPatch(o["utmgrp"]) {
			return fmt.Errorf("Error reading utmgrp: %v", err)
		}
	}

	if err = d.Set("wanoptgrp", dataSourceFlattenSystemAccprofileWanoptgrp(o["wanoptgrp"], d, "wanoptgrp")); err != nil {
		if !fortiAPIPatch(o["wanoptgrp"]) {
			return fmt.Errorf("Error reading wanoptgrp: %v", err)
		}
	}

	if err = d.Set("wifi", dataSourceFlattenSystemAccprofileWifi(o["wifi"], d, "wifi")); err != nil {
		if !fortiAPIPatch(o["wifi"]) {
			return fmt.Errorf("Error reading wifi: %v", err)
		}
	}

	if err = d.Set("netgrp_permission", dataSourceFlattenSystemAccprofileNetgrpPermission(o["netgrp-permission"], d, "netgrp_permission")); err != nil {
		if !fortiAPIPatch(o["netgrp-permission"]) {
			return fmt.Errorf("Error reading netgrp_permission: %v", err)
		}
	}

	if err = d.Set("sysgrp_permission", dataSourceFlattenSystemAccprofileSysgrpPermission(o["sysgrp-permission"], d, "sysgrp_permission")); err != nil {
		if !fortiAPIPatch(o["sysgrp-permission"]) {
			return fmt.Errorf("Error reading sysgrp_permission: %v", err)
		}
	}

	if err = d.Set("fwgrp_permission", dataSourceFlattenSystemAccprofileFwgrpPermission(o["fwgrp-permission"], d, "fwgrp_permission")); err != nil {
		if !fortiAPIPatch(o["fwgrp-permission"]) {
			return fmt.Errorf("Error reading fwgrp_permission: %v", err)
		}
	}

	if err = d.Set("loggrp_permission", dataSourceFlattenSystemAccprofileLoggrpPermission(o["loggrp-permission"], d, "loggrp_permission")); err != nil {
		if !fortiAPIPatch(o["loggrp-permission"]) {
			return fmt.Errorf("Error reading loggrp_permission: %v", err)
		}
	}

	if err = d.Set("utmgrp_permission", dataSourceFlattenSystemAccprofileUtmgrpPermission(o["utmgrp-permission"], d, "utmgrp_permission")); err != nil {
		if !fortiAPIPatch(o["utmgrp-permission"]) {
			return fmt.Errorf("Error reading utmgrp_permission: %v", err)
		}
	}

	if err = d.Set("admintimeout_override", dataSourceFlattenSystemAccprofileAdmintimeoutOverride(o["admintimeout-override"], d, "admintimeout_override")); err != nil {
		if !fortiAPIPatch(o["admintimeout-override"]) {
			return fmt.Errorf("Error reading admintimeout_override: %v", err)
		}
	}

	if err = d.Set("admintimeout", dataSourceFlattenSystemAccprofileAdmintimeout(o["admintimeout"], d, "admintimeout")); err != nil {
		if !fortiAPIPatch(o["admintimeout"]) {
			return fmt.Errorf("Error reading admintimeout: %v", err)
		}
	}

	if err = d.Set("cli_diagnose", dataSourceFlattenSystemAccprofileCliDiagnose(o["cli-diagnose"], d, "cli_diagnose")); err != nil {
		if !fortiAPIPatch(o["cli-diagnose"]) {
			return fmt.Errorf("Error reading cli_diagnose: %v", err)
		}
	}

	if err = d.Set("cli_get", dataSourceFlattenSystemAccprofileCliGet(o["cli-get"], d, "cli_get")); err != nil {
		if !fortiAPIPatch(o["cli-get"]) {
			return fmt.Errorf("Error reading cli_get: %v", err)
		}
	}

	if err = d.Set("cli_show", dataSourceFlattenSystemAccprofileCliShow(o["cli-show"], d, "cli_show")); err != nil {
		if !fortiAPIPatch(o["cli-show"]) {
			return fmt.Errorf("Error reading cli_show: %v", err)
		}
	}

	if err = d.Set("cli_exec", dataSourceFlattenSystemAccprofileCliExec(o["cli-exec"], d, "cli_exec")); err != nil {
		if !fortiAPIPatch(o["cli-exec"]) {
			return fmt.Errorf("Error reading cli_exec: %v", err)
		}
	}

	if err = d.Set("cli_config", dataSourceFlattenSystemAccprofileCliConfig(o["cli-config"], d, "cli_config")); err != nil {
		if !fortiAPIPatch(o["cli-config"]) {
			return fmt.Errorf("Error reading cli_config: %v", err)
		}
	}

	if err = d.Set("system_diagnostics", dataSourceFlattenSystemAccprofileSystemDiagnostics(o["system-diagnostics"], d, "system_diagnostics")); err != nil {
		if !fortiAPIPatch(o["system-diagnostics"]) {
			return fmt.Errorf("Error reading system_diagnostics: %v", err)
		}
	}

	if err = d.Set("system_execute_ssh", dataSourceFlattenSystemAccprofileSystemExecuteSsh(o["system-execute-ssh"], d, "system_execute_ssh")); err != nil {
		if !fortiAPIPatch(o["system-execute-ssh"]) {
			return fmt.Errorf("Error reading system_execute_ssh: %v", err)
		}
	}

	if err = d.Set("system_execute_telnet", dataSourceFlattenSystemAccprofileSystemExecuteTelnet(o["system-execute-telnet"], d, "system_execute_telnet")); err != nil {
		if !fortiAPIPatch(o["system-execute-telnet"]) {
			return fmt.Errorf("Error reading system_execute_telnet: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemAccprofileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
