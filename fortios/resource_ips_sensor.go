// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPS sensor.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceIpsSensor() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpsSensorCreate,
		Read:   resourceIpsSensorRead,
		Update: resourceIpsSensorUpdate,
		Delete: resourceIpsSensorDelete,

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
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"replacemsg_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"block_malicious_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scan_botnet_connections": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"rule": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"location": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"os": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"application": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"default_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"default_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cve": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cve_entry": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 19),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"vuln_type": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"last_modified": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_packet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_attack_context": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"rate_duration": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"rate_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"exempt_ip": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"src_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dst_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine_expiry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional:     true,
							Computed:     true,
						},
						"location": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"os": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"application": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_packet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine_expiry": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"quarantine_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"override": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_packet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine_expiry": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"quarantine_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"exempt_ip": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"src_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dst_ip": &schema.Schema{
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

func resourceIpsSensorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectIpsSensor(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating IpsSensor resource while getting object: %v", err)
	}

	o, err := c.CreateIpsSensor(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating IpsSensor resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("IpsSensor")
	}

	return resourceIpsSensorRead(d, m)
}

func resourceIpsSensorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectIpsSensor(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating IpsSensor resource while getting object: %v", err)
	}

	o, err := c.UpdateIpsSensor(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating IpsSensor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("IpsSensor")
	}

	return resourceIpsSensorRead(d, m)
}

func resourceIpsSensorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteIpsSensor(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting IpsSensor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsSensorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadIpsSensor(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading IpsSensor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIpsSensor(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading IpsSensor resource from API: %v", err)
	}
	return nil
}

func flattenIpsSensorName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorBlockMaliciousUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorExtendedLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntries(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenIpsSensorEntriesId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule"
		if _, ok := i["rule"]; ok {

			tmp["rule"] = flattenIpsSensorEntriesRule(i["rule"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "location"
		if _, ok := i["location"]; ok {

			tmp["location"] = flattenIpsSensorEntriesLocation(i["location"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {

			tmp["severity"] = flattenIpsSensorEntriesSeverity(i["severity"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {

			tmp["protocol"] = flattenIpsSensorEntriesProtocol(i["protocol"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os"
		if _, ok := i["os"]; ok {

			tmp["os"] = flattenIpsSensorEntriesOs(i["os"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := i["application"]; ok {

			tmp["application"] = flattenIpsSensorEntriesApplication(i["application"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_action"
		if _, ok := i["default-action"]; ok {

			tmp["default_action"] = flattenIpsSensorEntriesDefaultAction(i["default-action"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_status"
		if _, ok := i["default-status"]; ok {

			tmp["default_status"] = flattenIpsSensorEntriesDefaultStatus(i["default-status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cve"
		if _, ok := i["cve"]; ok {

			tmp["cve"] = flattenIpsSensorEntriesCve(i["cve"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vuln_type"
		if _, ok := i["vuln-type"]; ok {

			tmp["vuln_type"] = flattenIpsSensorEntriesVulnType(i["vuln-type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "last_modified"
		if _, ok := i["last-modified"]; ok {

			tmp["last_modified"] = flattenIpsSensorEntriesLastModified(i["last-modified"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenIpsSensorEntriesStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {

			tmp["log"] = flattenIpsSensorEntriesLog(i["log"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := i["log-packet"]; ok {

			tmp["log_packet"] = flattenIpsSensorEntriesLogPacket(i["log-packet"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_attack_context"
		if _, ok := i["log-attack-context"]; ok {

			tmp["log_attack_context"] = flattenIpsSensorEntriesLogAttackContext(i["log-attack-context"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {

			tmp["action"] = flattenIpsSensorEntriesAction(i["action"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_count"
		if _, ok := i["rate-count"]; ok {

			tmp["rate_count"] = flattenIpsSensorEntriesRateCount(i["rate-count"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_duration"
		if _, ok := i["rate-duration"]; ok {

			tmp["rate_duration"] = flattenIpsSensorEntriesRateDuration(i["rate-duration"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_mode"
		if _, ok := i["rate-mode"]; ok {

			tmp["rate_mode"] = flattenIpsSensorEntriesRateMode(i["rate-mode"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_track"
		if _, ok := i["rate-track"]; ok {

			tmp["rate_track"] = flattenIpsSensorEntriesRateTrack(i["rate-track"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt_ip"
		if _, ok := i["exempt-ip"]; ok {

			tmp["exempt_ip"] = flattenIpsSensorEntriesExemptIp(i["exempt-ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {

			tmp["quarantine"] = flattenIpsSensorEntriesQuarantine(i["quarantine"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := i["quarantine-expiry"]; ok {

			tmp["quarantine_expiry"] = flattenIpsSensorEntriesQuarantineExpiry(i["quarantine-expiry"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := i["quarantine-log"]; ok {

			tmp["quarantine_log"] = flattenIpsSensorEntriesQuarantineLog(i["quarantine-log"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenIpsSensorEntriesId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesRule(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenIpsSensorEntriesRuleId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenIpsSensorEntriesRuleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesLocation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	var rst_v interface{} = strings.TrimRight(v.(string), " ")
	return rst_v
}

func flattenIpsSensorEntriesProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesOs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesApplication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesDefaultAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesDefaultStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesCve(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cve_entry"
		if _, ok := i["cve-entry"]; ok {

			tmp["cve_entry"] = flattenIpsSensorEntriesCveCveEntry(i["cve-entry"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenIpsSensorEntriesCveCveEntry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesVulnType(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenIpsSensorEntriesVulnTypeId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenIpsSensorEntriesVulnTypeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesLastModified(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesLogPacket(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesLogAttackContext(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesRateCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesRateDuration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesRateMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesRateTrack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesExemptIp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenIpsSensorEntriesExemptIpId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := i["src-ip"]; ok {

			tmp["src_ip"] = flattenIpsSensorEntriesExemptIpSrcIp(i["src-ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_ip"
		if _, ok := i["dst-ip"]; ok {

			tmp["dst_ip"] = flattenIpsSensorEntriesExemptIpDstIp(i["dst-ip"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenIpsSensorEntriesExemptIpId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesExemptIpSrcIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenIpsSensorEntriesExemptIpDstIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenIpsSensorEntriesQuarantine(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorEntriesQuarantineLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorFilter(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenIpsSensorFilterName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "location"
		if _, ok := i["location"]; ok {

			tmp["location"] = flattenIpsSensorFilterLocation(i["location"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {

			tmp["severity"] = flattenIpsSensorFilterSeverity(i["severity"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {

			tmp["protocol"] = flattenIpsSensorFilterProtocol(i["protocol"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os"
		if _, ok := i["os"]; ok {

			tmp["os"] = flattenIpsSensorFilterOs(i["os"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := i["application"]; ok {

			tmp["application"] = flattenIpsSensorFilterApplication(i["application"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenIpsSensorFilterStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {

			tmp["log"] = flattenIpsSensorFilterLog(i["log"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := i["log-packet"]; ok {

			tmp["log_packet"] = flattenIpsSensorFilterLogPacket(i["log-packet"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {

			tmp["action"] = flattenIpsSensorFilterAction(i["action"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {

			tmp["quarantine"] = flattenIpsSensorFilterQuarantine(i["quarantine"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := i["quarantine-expiry"]; ok {

			tmp["quarantine_expiry"] = flattenIpsSensorFilterQuarantineExpiry(i["quarantine-expiry"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := i["quarantine-log"]; ok {

			tmp["quarantine_log"] = flattenIpsSensorFilterQuarantineLog(i["quarantine-log"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenIpsSensorFilterName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorFilterLocation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorFilterSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	var rst_v interface{} = strings.TrimRight(v.(string), " ")
	return rst_v
}

func flattenIpsSensorFilterProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorFilterOs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorFilterApplication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorFilterStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorFilterLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorFilterLogPacket(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorFilterAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorFilterQuarantine(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorFilterQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorFilterQuarantineLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorOverride(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_id"
		if _, ok := i["rule-id"]; ok {

			tmp["rule_id"] = flattenIpsSensorOverrideRuleId(i["rule-id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenIpsSensorOverrideStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {

			tmp["log"] = flattenIpsSensorOverrideLog(i["log"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := i["log-packet"]; ok {

			tmp["log_packet"] = flattenIpsSensorOverrideLogPacket(i["log-packet"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {

			tmp["action"] = flattenIpsSensorOverrideAction(i["action"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {

			tmp["quarantine"] = flattenIpsSensorOverrideQuarantine(i["quarantine"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := i["quarantine-expiry"]; ok {

			tmp["quarantine_expiry"] = flattenIpsSensorOverrideQuarantineExpiry(i["quarantine-expiry"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := i["quarantine-log"]; ok {

			tmp["quarantine_log"] = flattenIpsSensorOverrideQuarantineLog(i["quarantine-log"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt_ip"
		if _, ok := i["exempt-ip"]; ok {

			tmp["exempt_ip"] = flattenIpsSensorOverrideExemptIp(i["exempt-ip"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "rule_id", d)
	return result
}

func flattenIpsSensorOverrideRuleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorOverrideStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorOverrideLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorOverrideLogPacket(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorOverrideAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorOverrideQuarantine(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorOverrideQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorOverrideQuarantineLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorOverrideExemptIp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenIpsSensorOverrideExemptIpId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := i["src-ip"]; ok {

			tmp["src_ip"] = flattenIpsSensorOverrideExemptIpSrcIp(i["src-ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_ip"
		if _, ok := i["dst-ip"]; ok {

			tmp["dst_ip"] = flattenIpsSensorOverrideExemptIpDstIp(i["dst-ip"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenIpsSensorOverrideExemptIpId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSensorOverrideExemptIpSrcIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenIpsSensorOverrideExemptIpDstIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func refreshObjectIpsSensor(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenIpsSensorName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenIpsSensorComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenIpsSensorReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group", sv)); err != nil {
		if !fortiAPIPatch(o["replacemsg-group"]) {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("block_malicious_url", flattenIpsSensorBlockMaliciousUrl(o["block-malicious-url"], d, "block_malicious_url", sv)); err != nil {
		if !fortiAPIPatch(o["block-malicious-url"]) {
			return fmt.Errorf("Error reading block_malicious_url: %v", err)
		}
	}

	if err = d.Set("scan_botnet_connections", flattenIpsSensorScanBotnetConnections(o["scan-botnet-connections"], d, "scan_botnet_connections", sv)); err != nil {
		if !fortiAPIPatch(o["scan-botnet-connections"]) {
			return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
		}
	}

	if err = d.Set("extended_log", flattenIpsSensorExtendedLog(o["extended-log"], d, "extended_log", sv)); err != nil {
		if !fortiAPIPatch(o["extended-log"]) {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenIpsSensorEntries(o["entries"], d, "entries", sv)); err != nil {
			if !fortiAPIPatch(o["entries"]) {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenIpsSensorEntries(o["entries"], d, "entries", sv)); err != nil {
				if !fortiAPIPatch(o["entries"]) {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("filter", flattenIpsSensorFilter(o["filter"], d, "filter", sv)); err != nil {
			if !fortiAPIPatch(o["filter"]) {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("filter"); ok {
			if err = d.Set("filter", flattenIpsSensorFilter(o["filter"], d, "filter", sv)); err != nil {
				if !fortiAPIPatch(o["filter"]) {
					return fmt.Errorf("Error reading filter: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("override", flattenIpsSensorOverride(o["override"], d, "override", sv)); err != nil {
			if !fortiAPIPatch(o["override"]) {
				return fmt.Errorf("Error reading override: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("override"); ok {
			if err = d.Set("override", flattenIpsSensorOverride(o["override"], d, "override", sv)); err != nil {
				if !fortiAPIPatch(o["override"]) {
					return fmt.Errorf("Error reading override: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenIpsSensorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandIpsSensorName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorBlockMaliciousUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorScanBotnetConnections(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorExtendedLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandIpsSensorEntriesId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["rule"], _ = expandIpsSensorEntriesRule(d, i["rule"], pre_append, sv)
		} else {
			tmp["rule"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "location"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["location"], _ = expandIpsSensorEntriesLocation(d, i["location"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["severity"], _ = expandIpsSensorEntriesSeverity(d, i["severity"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["protocol"], _ = expandIpsSensorEntriesProtocol(d, i["protocol"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["os"], _ = expandIpsSensorEntriesOs(d, i["os"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["application"], _ = expandIpsSensorEntriesApplication(d, i["application"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_action"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["default-action"], _ = expandIpsSensorEntriesDefaultAction(d, i["default_action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["default-status"], _ = expandIpsSensorEntriesDefaultStatus(d, i["default_status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cve"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["cve"], _ = expandIpsSensorEntriesCve(d, i["cve"], pre_append, sv)
		} else {
			tmp["cve"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vuln_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["vuln-type"], _ = expandIpsSensorEntriesVulnType(d, i["vuln_type"], pre_append, sv)
		} else {
			tmp["vuln-type"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "last_modified"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["last-modified"], _ = expandIpsSensorEntriesLastModified(d, i["last_modified"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandIpsSensorEntriesStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["log"], _ = expandIpsSensorEntriesLog(d, i["log"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["log-packet"], _ = expandIpsSensorEntriesLogPacket(d, i["log_packet"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_attack_context"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["log-attack-context"], _ = expandIpsSensorEntriesLogAttackContext(d, i["log_attack_context"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["action"], _ = expandIpsSensorEntriesAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_count"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["rate-count"], _ = expandIpsSensorEntriesRateCount(d, i["rate_count"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_duration"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["rate-duration"], _ = expandIpsSensorEntriesRateDuration(d, i["rate_duration"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_mode"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["rate-mode"], _ = expandIpsSensorEntriesRateMode(d, i["rate_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_track"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["rate-track"], _ = expandIpsSensorEntriesRateTrack(d, i["rate_track"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt_ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["exempt-ip"], _ = expandIpsSensorEntriesExemptIp(d, i["exempt_ip"], pre_append, sv)
		} else {
			tmp["exempt-ip"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["quarantine"], _ = expandIpsSensorEntriesQuarantine(d, i["quarantine"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["quarantine-expiry"], _ = expandIpsSensorEntriesQuarantineExpiry(d, i["quarantine_expiry"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["quarantine-log"], _ = expandIpsSensorEntriesQuarantineLog(d, i["quarantine_log"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandIpsSensorEntriesId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandIpsSensorEntriesRuleId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandIpsSensorEntriesRuleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesLocation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesOs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesDefaultAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesDefaultStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesCve(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cve_entry"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["cve-entry"], _ = expandIpsSensorEntriesCveCveEntry(d, i["cve_entry"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandIpsSensorEntriesCveCveEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesVulnType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandIpsSensorEntriesVulnTypeId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandIpsSensorEntriesVulnTypeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesLastModified(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesLogPacket(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesLogAttackContext(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesRateCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesRateDuration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesRateMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesRateTrack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesExemptIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandIpsSensorEntriesExemptIpId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["src-ip"], _ = expandIpsSensorEntriesExemptIpSrcIp(d, i["src_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dst-ip"], _ = expandIpsSensorEntriesExemptIpDstIp(d, i["dst_ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandIpsSensorEntriesExemptIpId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesExemptIpSrcIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesExemptIpDstIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesQuarantine(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesQuarantineExpiry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorEntriesQuarantineLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandIpsSensorFilterName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "location"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["location"], _ = expandIpsSensorFilterLocation(d, i["location"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["severity"], _ = expandIpsSensorFilterSeverity(d, i["severity"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["protocol"], _ = expandIpsSensorFilterProtocol(d, i["protocol"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["os"], _ = expandIpsSensorFilterOs(d, i["os"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["application"], _ = expandIpsSensorFilterApplication(d, i["application"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandIpsSensorFilterStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["log"], _ = expandIpsSensorFilterLog(d, i["log"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["log-packet"], _ = expandIpsSensorFilterLogPacket(d, i["log_packet"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["action"], _ = expandIpsSensorFilterAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["quarantine"], _ = expandIpsSensorFilterQuarantine(d, i["quarantine"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["quarantine-expiry"], _ = expandIpsSensorFilterQuarantineExpiry(d, i["quarantine_expiry"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["quarantine-log"], _ = expandIpsSensorFilterQuarantineLog(d, i["quarantine_log"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandIpsSensorFilterName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorFilterLocation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorFilterSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorFilterProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorFilterOs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorFilterApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorFilterStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorFilterLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorFilterLogPacket(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorFilterAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorFilterQuarantine(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorFilterQuarantineExpiry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorFilterQuarantineLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["rule-id"], _ = expandIpsSensorOverrideRuleId(d, i["rule_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandIpsSensorOverrideStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["log"], _ = expandIpsSensorOverrideLog(d, i["log"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["log-packet"], _ = expandIpsSensorOverrideLogPacket(d, i["log_packet"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["action"], _ = expandIpsSensorOverrideAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["quarantine"], _ = expandIpsSensorOverrideQuarantine(d, i["quarantine"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["quarantine-expiry"], _ = expandIpsSensorOverrideQuarantineExpiry(d, i["quarantine_expiry"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["quarantine-log"], _ = expandIpsSensorOverrideQuarantineLog(d, i["quarantine_log"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt_ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["exempt-ip"], _ = expandIpsSensorOverrideExemptIp(d, i["exempt_ip"], pre_append, sv)
		} else {
			tmp["exempt-ip"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandIpsSensorOverrideRuleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorOverrideStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorOverrideLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorOverrideLogPacket(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorOverrideAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorOverrideQuarantine(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorOverrideQuarantineExpiry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorOverrideQuarantineLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorOverrideExemptIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandIpsSensorOverrideExemptIpId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["src-ip"], _ = expandIpsSensorOverrideExemptIpSrcIp(d, i["src_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dst-ip"], _ = expandIpsSensorOverrideExemptIpDstIp(d, i["dst_ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandIpsSensorOverrideExemptIpId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorOverrideExemptIpSrcIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSensorOverrideExemptIpDstIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectIpsSensor(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandIpsSensorName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandIpsSensorComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok {

		t, err := expandIpsSensorReplacemsgGroup(d, v, "replacemsg_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("block_malicious_url"); ok {

		t, err := expandIpsSensorBlockMaliciousUrl(d, v, "block_malicious_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-malicious-url"] = t
		}
	}

	if v, ok := d.GetOk("scan_botnet_connections"); ok {

		t, err := expandIpsSensorScanBotnetConnections(d, v, "scan_botnet_connections", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-botnet-connections"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok {

		t, err := expandIpsSensorExtendedLog(d, v, "extended_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok {

		t, err := expandIpsSensorEntries(d, v, "entries", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok {

		t, err := expandIpsSensorFilter(d, v, "filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("override"); ok {

		t, err := expandIpsSensorOverride(d, v, "override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override"] = t
		}
	}

	return &obj, nil
}
