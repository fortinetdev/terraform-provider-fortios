// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 multicast.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceRouterMulticast6() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterMulticast6Read,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"multicast_routing": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast_pmtu": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hello_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"hello_holdtime": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rp_candidate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"rp_candidate_group": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"rp_candidate_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rp_candidate_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"static_group": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"pim_sm_global": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bsr_candidate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"bsr_interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"bsr_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"bsr_hash": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"bsr_allow_quick_refresh": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"cisco_crp_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"register_rate_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"cisco_ignore_rp_set_priority": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"spt_threshold": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"spt_threshold_group": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"pim_use_sdwan": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"rp_address": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"ip6_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"group": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"pim_sm_global_vrf": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vrf": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"bsr_candidate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"bsr_interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"bsr_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"bsr_hash": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"bsr_allow_quick_refresh": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"cisco_crp_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"rp_address": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"ip6_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"group": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterMulticast6Read(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "RouterMulticast6"

	o, err := c.ReadRouterMulticast6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing RouterMulticast6: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterMulticast6(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterMulticast6 from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterMulticast6MulticastRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6MulticastPmtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6Interface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenRouterMulticast6InterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = dataSourceFlattenRouterMulticast6InterfaceHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_holdtime"
		if _, ok := i["hello-holdtime"]; ok {
			tmp["hello_holdtime"] = dataSourceFlattenRouterMulticast6InterfaceHelloHoldtime(i["hello-holdtime"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate"
		if _, ok := i["rp-candidate"]; ok {
			tmp["rp_candidate"] = dataSourceFlattenRouterMulticast6InterfaceRpCandidate(i["rp-candidate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_group"
		if _, ok := i["rp-candidate-group"]; ok {
			tmp["rp_candidate_group"] = dataSourceFlattenRouterMulticast6InterfaceRpCandidateGroup(i["rp-candidate-group"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_priority"
		if _, ok := i["rp-candidate-priority"]; ok {
			tmp["rp_candidate_priority"] = dataSourceFlattenRouterMulticast6InterfaceRpCandidatePriority(i["rp-candidate-priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_interval"
		if _, ok := i["rp-candidate-interval"]; ok {
			tmp["rp_candidate_interval"] = dataSourceFlattenRouterMulticast6InterfaceRpCandidateInterval(i["rp-candidate-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "static_group"
		if _, ok := i["static-group"]; ok {
			tmp["static_group"] = dataSourceFlattenRouterMulticast6InterfaceStaticGroup(i["static-group"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterMulticast6InterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6InterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6InterfaceHelloHoldtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6InterfaceRpCandidate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6InterfaceRpCandidateGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6InterfaceRpCandidatePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6InterfaceRpCandidateInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6InterfaceStaticGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobal(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bsr_candidate"
	if _, ok := i["bsr-candidate"]; ok {
		result["bsr_candidate"] = dataSourceFlattenRouterMulticast6PimSmGlobalBsrCandidate(i["bsr-candidate"], d, pre_append)
	}

	pre_append = pre + ".0." + "bsr_interface"
	if _, ok := i["bsr-interface"]; ok {
		result["bsr_interface"] = dataSourceFlattenRouterMulticast6PimSmGlobalBsrInterface(i["bsr-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "bsr_priority"
	if _, ok := i["bsr-priority"]; ok {
		result["bsr_priority"] = dataSourceFlattenRouterMulticast6PimSmGlobalBsrPriority(i["bsr-priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "bsr_hash"
	if _, ok := i["bsr-hash"]; ok {
		result["bsr_hash"] = dataSourceFlattenRouterMulticast6PimSmGlobalBsrHash(i["bsr-hash"], d, pre_append)
	}

	pre_append = pre + ".0." + "bsr_allow_quick_refresh"
	if _, ok := i["bsr-allow-quick-refresh"]; ok {
		result["bsr_allow_quick_refresh"] = dataSourceFlattenRouterMulticast6PimSmGlobalBsrAllowQuickRefresh(i["bsr-allow-quick-refresh"], d, pre_append)
	}

	pre_append = pre + ".0." + "cisco_crp_prefix"
	if _, ok := i["cisco-crp-prefix"]; ok {
		result["cisco_crp_prefix"] = dataSourceFlattenRouterMulticast6PimSmGlobalCiscoCrpPrefix(i["cisco-crp-prefix"], d, pre_append)
	}

	pre_append = pre + ".0." + "register_rate_limit"
	if _, ok := i["register-rate-limit"]; ok {
		result["register_rate_limit"] = dataSourceFlattenRouterMulticast6PimSmGlobalRegisterRateLimit(i["register-rate-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "cisco_ignore_rp_set_priority"
	if _, ok := i["cisco-ignore-rp-set-priority"]; ok {
		result["cisco_ignore_rp_set_priority"] = dataSourceFlattenRouterMulticast6PimSmGlobalCiscoIgnoreRpSetPriority(i["cisco-ignore-rp-set-priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "spt_threshold"
	if _, ok := i["spt-threshold"]; ok {
		result["spt_threshold"] = dataSourceFlattenRouterMulticast6PimSmGlobalSptThreshold(i["spt-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "spt_threshold_group"
	if _, ok := i["spt-threshold-group"]; ok {
		result["spt_threshold_group"] = dataSourceFlattenRouterMulticast6PimSmGlobalSptThresholdGroup(i["spt-threshold-group"], d, pre_append)
	}

	pre_append = pre + ".0." + "pim_use_sdwan"
	if _, ok := i["pim-use-sdwan"]; ok {
		result["pim_use_sdwan"] = dataSourceFlattenRouterMulticast6PimSmGlobalPimUseSdwan(i["pim-use-sdwan"], d, pre_append)
	}

	pre_append = pre + ".0." + "rp_address"
	if _, ok := i["rp-address"]; ok {
		result["rp_address"] = dataSourceFlattenRouterMulticast6PimSmGlobalRpAddress(i["rp-address"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenRouterMulticast6PimSmGlobalBsrCandidate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalBsrInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalBsrPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalBsrHash(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalBsrAllowQuickRefresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalCiscoCrpPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalRegisterRateLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalCiscoIgnoreRpSetPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalSptThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalSptThresholdGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalPimUseSdwan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalRpAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterMulticast6PimSmGlobalRpAddressId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := i["ip6-address"]; ok {
			tmp["ip6_address"] = dataSourceFlattenRouterMulticast6PimSmGlobalRpAddressIp6Address(i["ip6-address"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := i["group"]; ok {
			tmp["group"] = dataSourceFlattenRouterMulticast6PimSmGlobalRpAddressGroup(i["group"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterMulticast6PimSmGlobalRpAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalRpAddressIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalRpAddressGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalVrf(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {
			tmp["vrf"] = dataSourceFlattenRouterMulticast6PimSmGlobalVrfVrf(i["vrf"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_candidate"
		if _, ok := i["bsr-candidate"]; ok {
			tmp["bsr_candidate"] = dataSourceFlattenRouterMulticast6PimSmGlobalVrfBsrCandidate(i["bsr-candidate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_interface"
		if _, ok := i["bsr-interface"]; ok {
			tmp["bsr_interface"] = dataSourceFlattenRouterMulticast6PimSmGlobalVrfBsrInterface(i["bsr-interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_priority"
		if _, ok := i["bsr-priority"]; ok {
			tmp["bsr_priority"] = dataSourceFlattenRouterMulticast6PimSmGlobalVrfBsrPriority(i["bsr-priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_hash"
		if _, ok := i["bsr-hash"]; ok {
			tmp["bsr_hash"] = dataSourceFlattenRouterMulticast6PimSmGlobalVrfBsrHash(i["bsr-hash"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_allow_quick_refresh"
		if _, ok := i["bsr-allow-quick-refresh"]; ok {
			tmp["bsr_allow_quick_refresh"] = dataSourceFlattenRouterMulticast6PimSmGlobalVrfBsrAllowQuickRefresh(i["bsr-allow-quick-refresh"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cisco_crp_prefix"
		if _, ok := i["cisco-crp-prefix"]; ok {
			tmp["cisco_crp_prefix"] = dataSourceFlattenRouterMulticast6PimSmGlobalVrfCiscoCrpPrefix(i["cisco-crp-prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_address"
		if _, ok := i["rp-address"]; ok {
			tmp["rp_address"] = dataSourceFlattenRouterMulticast6PimSmGlobalVrfRpAddress(i["rp-address"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterMulticast6PimSmGlobalVrfVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalVrfBsrCandidate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalVrfBsrInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalVrfBsrPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalVrfBsrHash(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalVrfBsrAllowQuickRefresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalVrfCiscoCrpPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalVrfRpAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterMulticast6PimSmGlobalVrfRpAddressId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := i["ip6-address"]; ok {
			tmp["ip6_address"] = dataSourceFlattenRouterMulticast6PimSmGlobalVrfRpAddressIp6Address(i["ip6-address"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := i["group"]; ok {
			tmp["group"] = dataSourceFlattenRouterMulticast6PimSmGlobalVrfRpAddressGroup(i["group"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterMulticast6PimSmGlobalVrfRpAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalVrfRpAddressIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalVrfRpAddressGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterMulticast6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("multicast_routing", dataSourceFlattenRouterMulticast6MulticastRouting(o["multicast-routing"], d, "multicast_routing")); err != nil {
		if !fortiAPIPatch(o["multicast-routing"]) {
			return fmt.Errorf("Error reading multicast_routing: %v", err)
		}
	}

	if err = d.Set("multicast_pmtu", dataSourceFlattenRouterMulticast6MulticastPmtu(o["multicast-pmtu"], d, "multicast_pmtu")); err != nil {
		if !fortiAPIPatch(o["multicast-pmtu"]) {
			return fmt.Errorf("Error reading multicast_pmtu: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenRouterMulticast6Interface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("pim_sm_global", dataSourceFlattenRouterMulticast6PimSmGlobal(o["pim-sm-global"], d, "pim_sm_global")); err != nil {
		if !fortiAPIPatch(o["pim-sm-global"]) {
			return fmt.Errorf("Error reading pim_sm_global: %v", err)
		}
	}

	if err = d.Set("pim_sm_global_vrf", dataSourceFlattenRouterMulticast6PimSmGlobalVrf(o["pim-sm-global-vrf"], d, "pim_sm_global_vrf")); err != nil {
		if !fortiAPIPatch(o["pim-sm-global-vrf"]) {
			return fmt.Errorf("Error reading pim_sm_global_vrf: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterMulticast6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
