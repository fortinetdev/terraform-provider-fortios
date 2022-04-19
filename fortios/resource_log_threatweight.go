// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure threat weight settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceLogThreatWeight() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogThreatWeightUpdate,
		Read:   resourceLogThreatWeightRead,
		Update: resourceLogThreatWeightUpdate,
		Delete: resourceLogThreatWeightDelete,

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
			"level": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"low": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),
							Optional:     true,
							Computed:     true,
						},
						"medium": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),
							Optional:     true,
							Computed:     true,
						},
						"high": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),
							Optional:     true,
							Computed:     true,
						},
						"critical": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"blocked_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"failed_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url_block_detected": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"botnet_connection_detected": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malware": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"virus_infected": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"file_blocked": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"command_blocked": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"oversized": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"virus_scan_error": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"switch_proto": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mimefragmented": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"virus_file_type_executable": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"virus_outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"botnet_connection": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"content_disarm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malware_list": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ems_threat_feed": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fsa_malicious": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fsa_high_risk": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fsa_medium_risk": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"info_severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"low_severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"medium_severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"high_severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"critical_severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"web": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"category": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"geolocation": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"country": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 2),
							Optional:     true,
							Computed:     true,
						},
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"application": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"category": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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

func resourceLogThreatWeightUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogThreatWeight(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogThreatWeight resource while getting object: %v", err)
	}

	o, err := c.UpdateLogThreatWeight(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogThreatWeight resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogThreatWeight")
	}

	return resourceLogThreatWeightRead(d, m)
}

func resourceLogThreatWeightDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogThreatWeight(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogThreatWeight resource while getting object: %v", err)
	}

	_, err = c.UpdateLogThreatWeight(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing LogThreatWeight resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogThreatWeightRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadLogThreatWeight(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogThreatWeight resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogThreatWeight(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogThreatWeight resource from API: %v", err)
	}
	return nil
}

func flattenLogThreatWeightStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightLevel(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "low"
	if _, ok := i["low"]; ok {

		result["low"] = flattenLogThreatWeightLevelLow(i["low"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "medium"
	if _, ok := i["medium"]; ok {

		result["medium"] = flattenLogThreatWeightLevelMedium(i["medium"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "high"
	if _, ok := i["high"]; ok {

		result["high"] = flattenLogThreatWeightLevelHigh(i["high"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "critical"
	if _, ok := i["critical"]; ok {

		result["critical"] = flattenLogThreatWeightLevelCritical(i["critical"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenLogThreatWeightLevelLow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightLevelMedium(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightLevelHigh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightLevelCritical(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightBlockedConnection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightFailedConnection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightUrlBlockDetected(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightBotnetConnectionDetected(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightMalware(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "virus_infected"
	if _, ok := i["virus-infected"]; ok {

		result["virus_infected"] = flattenLogThreatWeightMalwareVirusInfected(i["virus-infected"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {

		result["fortiai"] = flattenLogThreatWeightMalwareFortiai(i["fortiai"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "file_blocked"
	if _, ok := i["file-blocked"]; ok {

		result["file_blocked"] = flattenLogThreatWeightMalwareFileBlocked(i["file-blocked"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "command_blocked"
	if _, ok := i["command-blocked"]; ok {

		result["command_blocked"] = flattenLogThreatWeightMalwareCommandBlocked(i["command-blocked"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "oversized"
	if _, ok := i["oversized"]; ok {

		result["oversized"] = flattenLogThreatWeightMalwareOversized(i["oversized"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "virus_scan_error"
	if _, ok := i["virus-scan-error"]; ok {

		result["virus_scan_error"] = flattenLogThreatWeightMalwareVirusScanError(i["virus-scan-error"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "switch_proto"
	if _, ok := i["switch-proto"]; ok {

		result["switch_proto"] = flattenLogThreatWeightMalwareSwitchProto(i["switch-proto"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mimefragmented"
	if _, ok := i["mimefragmented"]; ok {

		result["mimefragmented"] = flattenLogThreatWeightMalwareMimefragmented(i["mimefragmented"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "virus_file_type_executable"
	if _, ok := i["virus-file-type-executable"]; ok {

		result["virus_file_type_executable"] = flattenLogThreatWeightMalwareVirusFileTypeExecutable(i["virus-file-type-executable"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "virus_outbreak_prevention"
	if _, ok := i["virus-outbreak-prevention"]; ok {

		result["virus_outbreak_prevention"] = flattenLogThreatWeightMalwareVirusOutbreakPrevention(i["virus-outbreak-prevention"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "botnet_connection"
	if _, ok := i["botnet-connection"]; ok {

		result["botnet_connection"] = flattenLogThreatWeightMalwareBotnetConnection(i["botnet-connection"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "content_disarm"
	if _, ok := i["content-disarm"]; ok {

		result["content_disarm"] = flattenLogThreatWeightMalwareContentDisarm(i["content-disarm"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malware_list"
	if _, ok := i["malware-list"]; ok {

		result["malware_list"] = flattenLogThreatWeightMalwareMalwareList(i["malware-list"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ems_threat_feed"
	if _, ok := i["ems-threat-feed"]; ok {

		result["ems_threat_feed"] = flattenLogThreatWeightMalwareEmsThreatFeed(i["ems-threat-feed"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fsa_malicious"
	if _, ok := i["fsa-malicious"]; ok {

		result["fsa_malicious"] = flattenLogThreatWeightMalwareFsaMalicious(i["fsa-malicious"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fsa_high_risk"
	if _, ok := i["fsa-high-risk"]; ok {

		result["fsa_high_risk"] = flattenLogThreatWeightMalwareFsaHighRisk(i["fsa-high-risk"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fsa_medium_risk"
	if _, ok := i["fsa-medium-risk"]; ok {

		result["fsa_medium_risk"] = flattenLogThreatWeightMalwareFsaMediumRisk(i["fsa-medium-risk"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenLogThreatWeightMalwareVirusInfected(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareFortiai(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareFileBlocked(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareCommandBlocked(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareOversized(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareVirusScanError(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareSwitchProto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareMimefragmented(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareVirusFileTypeExecutable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareVirusOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareBotnetConnection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareContentDisarm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareMalwareList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareEmsThreatFeed(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareFsaMalicious(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareFsaHighRisk(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightMalwareFsaMediumRisk(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightIps(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "info_severity"
	if _, ok := i["info-severity"]; ok {

		result["info_severity"] = flattenLogThreatWeightIpsInfoSeverity(i["info-severity"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "low_severity"
	if _, ok := i["low-severity"]; ok {

		result["low_severity"] = flattenLogThreatWeightIpsLowSeverity(i["low-severity"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "medium_severity"
	if _, ok := i["medium-severity"]; ok {

		result["medium_severity"] = flattenLogThreatWeightIpsMediumSeverity(i["medium-severity"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "high_severity"
	if _, ok := i["high-severity"]; ok {

		result["high_severity"] = flattenLogThreatWeightIpsHighSeverity(i["high-severity"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "critical_severity"
	if _, ok := i["critical-severity"]; ok {

		result["critical_severity"] = flattenLogThreatWeightIpsCriticalSeverity(i["critical-severity"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenLogThreatWeightIpsInfoSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightIpsLowSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightIpsMediumSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightIpsHighSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightIpsCriticalSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightWeb(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenLogThreatWeightWebId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {

			tmp["category"] = flattenLogThreatWeightWebCategory(i["category"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {

			tmp["level"] = flattenLogThreatWeightWebLevel(i["level"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenLogThreatWeightWebId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightWebCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightWebLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightGeolocation(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenLogThreatWeightGeolocationId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "country"
		if _, ok := i["country"]; ok {

			tmp["country"] = flattenLogThreatWeightGeolocationCountry(i["country"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {

			tmp["level"] = flattenLogThreatWeightGeolocationLevel(i["level"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenLogThreatWeightGeolocationId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightGeolocationCountry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightGeolocationLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightApplication(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenLogThreatWeightApplicationId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {

			tmp["category"] = flattenLogThreatWeightApplicationCategory(i["category"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {

			tmp["level"] = flattenLogThreatWeightApplicationLevel(i["level"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenLogThreatWeightApplicationId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightApplicationCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogThreatWeightApplicationLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogThreatWeight(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenLogThreatWeightStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("level", flattenLogThreatWeightLevel(o["level"], d, "level", sv)); err != nil {
			if !fortiAPIPatch(o["level"]) {
				return fmt.Errorf("Error reading level: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("level"); ok {
			if err = d.Set("level", flattenLogThreatWeightLevel(o["level"], d, "level", sv)); err != nil {
				if !fortiAPIPatch(o["level"]) {
					return fmt.Errorf("Error reading level: %v", err)
				}
			}
		}
	}

	if err = d.Set("blocked_connection", flattenLogThreatWeightBlockedConnection(o["blocked-connection"], d, "blocked_connection", sv)); err != nil {
		if !fortiAPIPatch(o["blocked-connection"]) {
			return fmt.Errorf("Error reading blocked_connection: %v", err)
		}
	}

	if err = d.Set("failed_connection", flattenLogThreatWeightFailedConnection(o["failed-connection"], d, "failed_connection", sv)); err != nil {
		if !fortiAPIPatch(o["failed-connection"]) {
			return fmt.Errorf("Error reading failed_connection: %v", err)
		}
	}

	if err = d.Set("url_block_detected", flattenLogThreatWeightUrlBlockDetected(o["url-block-detected"], d, "url_block_detected", sv)); err != nil {
		if !fortiAPIPatch(o["url-block-detected"]) {
			return fmt.Errorf("Error reading url_block_detected: %v", err)
		}
	}

	if err = d.Set("botnet_connection_detected", flattenLogThreatWeightBotnetConnectionDetected(o["botnet-connection-detected"], d, "botnet_connection_detected", sv)); err != nil {
		if !fortiAPIPatch(o["botnet-connection-detected"]) {
			return fmt.Errorf("Error reading botnet_connection_detected: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("malware", flattenLogThreatWeightMalware(o["malware"], d, "malware", sv)); err != nil {
			if !fortiAPIPatch(o["malware"]) {
				return fmt.Errorf("Error reading malware: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("malware"); ok {
			if err = d.Set("malware", flattenLogThreatWeightMalware(o["malware"], d, "malware", sv)); err != nil {
				if !fortiAPIPatch(o["malware"]) {
					return fmt.Errorf("Error reading malware: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ips", flattenLogThreatWeightIps(o["ips"], d, "ips", sv)); err != nil {
			if !fortiAPIPatch(o["ips"]) {
				return fmt.Errorf("Error reading ips: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ips"); ok {
			if err = d.Set("ips", flattenLogThreatWeightIps(o["ips"], d, "ips", sv)); err != nil {
				if !fortiAPIPatch(o["ips"]) {
					return fmt.Errorf("Error reading ips: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("web", flattenLogThreatWeightWeb(o["web"], d, "web", sv)); err != nil {
			if !fortiAPIPatch(o["web"]) {
				return fmt.Errorf("Error reading web: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("web"); ok {
			if err = d.Set("web", flattenLogThreatWeightWeb(o["web"], d, "web", sv)); err != nil {
				if !fortiAPIPatch(o["web"]) {
					return fmt.Errorf("Error reading web: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("geolocation", flattenLogThreatWeightGeolocation(o["geolocation"], d, "geolocation", sv)); err != nil {
			if !fortiAPIPatch(o["geolocation"]) {
				return fmt.Errorf("Error reading geolocation: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("geolocation"); ok {
			if err = d.Set("geolocation", flattenLogThreatWeightGeolocation(o["geolocation"], d, "geolocation", sv)); err != nil {
				if !fortiAPIPatch(o["geolocation"]) {
					return fmt.Errorf("Error reading geolocation: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("application", flattenLogThreatWeightApplication(o["application"], d, "application", sv)); err != nil {
			if !fortiAPIPatch(o["application"]) {
				return fmt.Errorf("Error reading application: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("application"); ok {
			if err = d.Set("application", flattenLogThreatWeightApplication(o["application"], d, "application", sv)); err != nil {
				if !fortiAPIPatch(o["application"]) {
					return fmt.Errorf("Error reading application: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenLogThreatWeightFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogThreatWeightStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightLevel(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "low"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["low"] = nil
		} else {

			result["low"], _ = expandLogThreatWeightLevelLow(d, i["low"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "medium"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["medium"] = nil
		} else {

			result["medium"], _ = expandLogThreatWeightLevelMedium(d, i["medium"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "high"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["high"] = nil
		} else {

			result["high"], _ = expandLogThreatWeightLevelHigh(d, i["high"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "critical"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["critical"] = nil
		} else {

			result["critical"], _ = expandLogThreatWeightLevelCritical(d, i["critical"], pre_append, sv)
		}
	}

	return result, nil
}

func expandLogThreatWeightLevelLow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightLevelMedium(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightLevelHigh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightLevelCritical(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightBlockedConnection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightFailedConnection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightUrlBlockDetected(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightBotnetConnectionDetected(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalware(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "virus_infected"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["virus-infected"] = nil
		} else {

			result["virus-infected"], _ = expandLogThreatWeightMalwareVirusInfected(d, i["virus_infected"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["fortiai"] = nil
		} else {

			result["fortiai"], _ = expandLogThreatWeightMalwareFortiai(d, i["fortiai"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "file_blocked"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["file-blocked"] = nil
		} else {

			result["file-blocked"], _ = expandLogThreatWeightMalwareFileBlocked(d, i["file_blocked"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "command_blocked"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["command-blocked"] = nil
		} else {

			result["command-blocked"], _ = expandLogThreatWeightMalwareCommandBlocked(d, i["command_blocked"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "oversized"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["oversized"] = nil
		} else {

			result["oversized"], _ = expandLogThreatWeightMalwareOversized(d, i["oversized"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "virus_scan_error"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["virus-scan-error"] = nil
		} else {

			result["virus-scan-error"], _ = expandLogThreatWeightMalwareVirusScanError(d, i["virus_scan_error"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "switch_proto"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["switch-proto"] = nil
		} else {

			result["switch-proto"], _ = expandLogThreatWeightMalwareSwitchProto(d, i["switch_proto"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "mimefragmented"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mimefragmented"] = nil
		} else {

			result["mimefragmented"], _ = expandLogThreatWeightMalwareMimefragmented(d, i["mimefragmented"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "virus_file_type_executable"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["virus-file-type-executable"] = nil
		} else {

			result["virus-file-type-executable"], _ = expandLogThreatWeightMalwareVirusFileTypeExecutable(d, i["virus_file_type_executable"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "virus_outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["virus-outbreak-prevention"] = nil
		} else {

			result["virus-outbreak-prevention"], _ = expandLogThreatWeightMalwareVirusOutbreakPrevention(d, i["virus_outbreak_prevention"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "botnet_connection"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["botnet-connection"] = nil
		} else {

			result["botnet-connection"], _ = expandLogThreatWeightMalwareBotnetConnection(d, i["botnet_connection"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "content_disarm"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["content-disarm"] = nil
		} else {

			result["content-disarm"], _ = expandLogThreatWeightMalwareContentDisarm(d, i["content_disarm"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "malware_list"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["malware-list"] = nil
		} else {

			result["malware-list"], _ = expandLogThreatWeightMalwareMalwareList(d, i["malware_list"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ems_threat_feed"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ems-threat-feed"] = nil
		} else {

			result["ems-threat-feed"], _ = expandLogThreatWeightMalwareEmsThreatFeed(d, i["ems_threat_feed"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "fsa_malicious"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["fsa-malicious"] = nil
		} else {

			result["fsa-malicious"], _ = expandLogThreatWeightMalwareFsaMalicious(d, i["fsa_malicious"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "fsa_high_risk"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["fsa-high-risk"] = nil
		} else {

			result["fsa-high-risk"], _ = expandLogThreatWeightMalwareFsaHighRisk(d, i["fsa_high_risk"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "fsa_medium_risk"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["fsa-medium-risk"] = nil
		} else {

			result["fsa-medium-risk"], _ = expandLogThreatWeightMalwareFsaMediumRisk(d, i["fsa_medium_risk"], pre_append, sv)
		}
	}

	return result, nil
}

func expandLogThreatWeightMalwareVirusInfected(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareFortiai(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareFileBlocked(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareCommandBlocked(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareOversized(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareVirusScanError(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareSwitchProto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareMimefragmented(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareVirusFileTypeExecutable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareVirusOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareBotnetConnection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareContentDisarm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareMalwareList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareEmsThreatFeed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareFsaMalicious(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareFsaHighRisk(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightMalwareFsaMediumRisk(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightIps(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "info_severity"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["info-severity"] = nil
		} else {

			result["info-severity"], _ = expandLogThreatWeightIpsInfoSeverity(d, i["info_severity"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "low_severity"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["low-severity"] = nil
		} else {

			result["low-severity"], _ = expandLogThreatWeightIpsLowSeverity(d, i["low_severity"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "medium_severity"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["medium-severity"] = nil
		} else {

			result["medium-severity"], _ = expandLogThreatWeightIpsMediumSeverity(d, i["medium_severity"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "high_severity"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["high-severity"] = nil
		} else {

			result["high-severity"], _ = expandLogThreatWeightIpsHighSeverity(d, i["high_severity"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "critical_severity"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["critical-severity"] = nil
		} else {

			result["critical-severity"], _ = expandLogThreatWeightIpsCriticalSeverity(d, i["critical_severity"], pre_append, sv)
		}
	}

	return result, nil
}

func expandLogThreatWeightIpsInfoSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightIpsLowSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightIpsMediumSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightIpsHighSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightIpsCriticalSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightWeb(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandLogThreatWeightWebId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["category"], _ = expandLogThreatWeightWebCategory(d, i["category"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["level"], _ = expandLogThreatWeightWebLevel(d, i["level"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLogThreatWeightWebId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightWebCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightWebLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightGeolocation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandLogThreatWeightGeolocationId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "country"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["country"], _ = expandLogThreatWeightGeolocationCountry(d, i["country"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["level"], _ = expandLogThreatWeightGeolocationLevel(d, i["level"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLogThreatWeightGeolocationId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightGeolocationCountry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightGeolocationLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandLogThreatWeightApplicationId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["category"], _ = expandLogThreatWeightApplicationCategory(d, i["category"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["level"], _ = expandLogThreatWeightApplicationLevel(d, i["level"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLogThreatWeightApplicationId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightApplicationCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogThreatWeightApplicationLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogThreatWeight(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandLogThreatWeightStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("level"); ok {

		t, err := expandLogThreatWeightLevel(d, v, "level", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["level"] = t
		}
	}

	if v, ok := d.GetOk("blocked_connection"); ok {
		if setArgNil {
			obj["blocked-connection"] = nil
		} else {

			t, err := expandLogThreatWeightBlockedConnection(d, v, "blocked_connection", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["blocked-connection"] = t
			}
		}
	}

	if v, ok := d.GetOk("failed_connection"); ok {
		if setArgNil {
			obj["failed-connection"] = nil
		} else {

			t, err := expandLogThreatWeightFailedConnection(d, v, "failed_connection", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["failed-connection"] = t
			}
		}
	}

	if v, ok := d.GetOk("url_block_detected"); ok {
		if setArgNil {
			obj["url-block-detected"] = nil
		} else {

			t, err := expandLogThreatWeightUrlBlockDetected(d, v, "url_block_detected", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["url-block-detected"] = t
			}
		}
	}

	if v, ok := d.GetOk("botnet_connection_detected"); ok {
		if setArgNil {
			obj["botnet-connection-detected"] = nil
		} else {

			t, err := expandLogThreatWeightBotnetConnectionDetected(d, v, "botnet_connection_detected", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["botnet-connection-detected"] = t
			}
		}
	}

	if v, ok := d.GetOk("malware"); ok {

		t, err := expandLogThreatWeightMalware(d, v, "malware", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malware"] = t
		}
	}

	if v, ok := d.GetOk("ips"); ok {

		t, err := expandLogThreatWeightIps(d, v, "ips", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips"] = t
		}
	}

	if v, ok := d.GetOk("web"); ok {
		if setArgNil {
			obj["web"] = make([]struct{}, 0)
		} else {

			t, err := expandLogThreatWeightWeb(d, v, "web", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["web"] = t
			}
		}
	}

	if v, ok := d.GetOk("geolocation"); ok {
		if setArgNil {
			obj["geolocation"] = make([]struct{}, 0)
		} else {

			t, err := expandLogThreatWeightGeolocation(d, v, "geolocation", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["geolocation"] = t
			}
		}
	}

	if v, ok := d.GetOk("application"); ok {
		if setArgNil {
			obj["application"] = make([]struct{}, 0)
		} else {

			t, err := expandLogThreatWeightApplication(d, v, "application", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["application"] = t
			}
		}
	}

	return &obj, nil
}
