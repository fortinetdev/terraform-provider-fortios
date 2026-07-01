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

func resourceRouterMulticast6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticast6Update,
		Read:   resourceRouterMulticast6Read,
		Update: resourceRouterMulticast6Update,
		Delete: resourceRouterMulticast6Delete,

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
			"multicast_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_pmtu": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
						"hello_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"hello_holdtime": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
						},
						"rp_candidate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rp_candidate_group": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"rp_candidate_priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"rp_candidate_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16383),
							Optional:     true,
							Computed:     true,
						},
						"static_group": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
					},
				},
			},
			"pim_sm_global": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bsr_candidate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bsr_interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
						"bsr_priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"bsr_hash": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 128),
							Optional:     true,
							Computed:     true,
						},
						"bsr_allow_quick_refresh": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cisco_crp_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"register_rate_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"cisco_ignore_rp_set_priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"spt_threshold": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"spt_threshold_group": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"pim_use_sdwan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rp_address": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"ip6_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"group": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
								},
							},
						},
					},
				},
			},
			"pim_sm_global_vrf": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vrf": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 511),
							Optional:     true,
							Computed:     true,
						},
						"bsr_candidate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bsr_interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
						"bsr_priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"bsr_hash": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 128),
							Optional:     true,
							Computed:     true,
						},
						"bsr_allow_quick_refresh": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cisco_crp_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rp_address": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"ip6_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"group": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
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
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceRouterMulticast6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterMulticast6(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6 resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterMulticast6(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterMulticast6")
	}

	return resourceRouterMulticast6Read(d, m)
}

func resourceRouterMulticast6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterMulticast6(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6 resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterMulticast6(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing RouterMulticast6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticast6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterMulticast6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticast6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticast6(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticast6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticast6MulticastRouting(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6MulticastPmtu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6Interface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenRouterMulticast6InterfaceName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["hello-interval"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
			}
			tmp["hello_interval"] = flattenRouterMulticast6InterfaceHelloInterval(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["hello-holdtime"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_holdtime"
			}
			tmp["hello_holdtime"] = flattenRouterMulticast6InterfaceHelloHoldtime(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["rp-candidate"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate"
			}
			tmp["rp_candidate"] = flattenRouterMulticast6InterfaceRpCandidate(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["rp-candidate-group"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_group"
			}
			tmp["rp_candidate_group"] = flattenRouterMulticast6InterfaceRpCandidateGroup(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["rp-candidate-priority"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_priority"
			}
			tmp["rp_candidate_priority"] = flattenRouterMulticast6InterfaceRpCandidatePriority(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["rp-candidate-interval"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_interval"
			}
			tmp["rp_candidate_interval"] = flattenRouterMulticast6InterfaceRpCandidateInterval(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["static-group"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "static_group"
			}
			tmp["static_group"] = flattenRouterMulticast6InterfaceStaticGroup(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterMulticast6InterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6InterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticast6InterfaceHelloHoldtime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticast6InterfaceRpCandidate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6InterfaceRpCandidateGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6InterfaceRpCandidatePriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticast6InterfaceRpCandidateInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticast6InterfaceStaticGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobal(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bsr_candidate"
	if _, ok := i["bsr-candidate"]; ok {
		result["bsr_candidate"] = flattenRouterMulticast6PimSmGlobalBsrCandidate(i["bsr-candidate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bsr_interface"
	if _, ok := i["bsr-interface"]; ok {
		result["bsr_interface"] = flattenRouterMulticast6PimSmGlobalBsrInterface(i["bsr-interface"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bsr_priority"
	if _, ok := i["bsr-priority"]; ok {
		result["bsr_priority"] = flattenRouterMulticast6PimSmGlobalBsrPriority(i["bsr-priority"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bsr_hash"
	if _, ok := i["bsr-hash"]; ok {
		result["bsr_hash"] = flattenRouterMulticast6PimSmGlobalBsrHash(i["bsr-hash"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bsr_allow_quick_refresh"
	if _, ok := i["bsr-allow-quick-refresh"]; ok {
		result["bsr_allow_quick_refresh"] = flattenRouterMulticast6PimSmGlobalBsrAllowQuickRefresh(i["bsr-allow-quick-refresh"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cisco_crp_prefix"
	if _, ok := i["cisco-crp-prefix"]; ok {
		result["cisco_crp_prefix"] = flattenRouterMulticast6PimSmGlobalCiscoCrpPrefix(i["cisco-crp-prefix"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "register_rate_limit"
	if _, ok := i["register-rate-limit"]; ok {
		result["register_rate_limit"] = flattenRouterMulticast6PimSmGlobalRegisterRateLimit(i["register-rate-limit"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cisco_ignore_rp_set_priority"
	if _, ok := i["cisco-ignore-rp-set-priority"]; ok {
		result["cisco_ignore_rp_set_priority"] = flattenRouterMulticast6PimSmGlobalCiscoIgnoreRpSetPriority(i["cisco-ignore-rp-set-priority"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "spt_threshold"
	if _, ok := i["spt-threshold"]; ok {
		result["spt_threshold"] = flattenRouterMulticast6PimSmGlobalSptThreshold(i["spt-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "spt_threshold_group"
	if _, ok := i["spt-threshold-group"]; ok {
		result["spt_threshold_group"] = flattenRouterMulticast6PimSmGlobalSptThresholdGroup(i["spt-threshold-group"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "pim_use_sdwan"
	if _, ok := i["pim-use-sdwan"]; ok {
		result["pim_use_sdwan"] = flattenRouterMulticast6PimSmGlobalPimUseSdwan(i["pim-use-sdwan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "rp_address"
	if _, ok := i["rp-address"]; ok {
		result["rp_address"] = flattenRouterMulticast6PimSmGlobalRpAddress(i["rp-address"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenRouterMulticast6PimSmGlobalBsrCandidate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalBsrInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalBsrPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticast6PimSmGlobalBsrHash(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticast6PimSmGlobalBsrAllowQuickRefresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalCiscoCrpPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalRegisterRateLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticast6PimSmGlobalCiscoIgnoreRpSetPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalSptThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalSptThresholdGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalPimUseSdwan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalRpAddress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenRouterMulticast6PimSmGlobalRpAddressId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ip6-address"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
			}
			tmp["ip6_address"] = flattenRouterMulticast6PimSmGlobalRpAddressIp6Address(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["group"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
			}
			tmp["group"] = flattenRouterMulticast6PimSmGlobalRpAddressGroup(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterMulticast6PimSmGlobalRpAddressId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticast6PimSmGlobalRpAddressIp6Address(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalRpAddressGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrf(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "vrf", "vrf")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["vrf"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
			}
			tmp["vrf"] = flattenRouterMulticast6PimSmGlobalVrfVrf(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["bsr-candidate"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_candidate"
			}
			tmp["bsr_candidate"] = flattenRouterMulticast6PimSmGlobalVrfBsrCandidate(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["bsr-interface"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_interface"
			}
			tmp["bsr_interface"] = flattenRouterMulticast6PimSmGlobalVrfBsrInterface(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["bsr-priority"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_priority"
			}
			tmp["bsr_priority"] = flattenRouterMulticast6PimSmGlobalVrfBsrPriority(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["bsr-hash"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_hash"
			}
			tmp["bsr_hash"] = flattenRouterMulticast6PimSmGlobalVrfBsrHash(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["bsr-allow-quick-refresh"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_allow_quick_refresh"
			}
			tmp["bsr_allow_quick_refresh"] = flattenRouterMulticast6PimSmGlobalVrfBsrAllowQuickRefresh(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["cisco-crp-prefix"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "cisco_crp_prefix"
			}
			tmp["cisco_crp_prefix"] = flattenRouterMulticast6PimSmGlobalVrfCiscoCrpPrefix(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["rp-address"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_address"
			}
			tmp["rp_address"] = flattenRouterMulticast6PimSmGlobalVrfRpAddress(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vrf", d)
	return result
}

func flattenRouterMulticast6PimSmGlobalVrfVrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticast6PimSmGlobalVrfBsrCandidate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfBsrInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfBsrPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticast6PimSmGlobalVrfBsrHash(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticast6PimSmGlobalVrfBsrAllowQuickRefresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfCiscoCrpPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfRpAddress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenRouterMulticast6PimSmGlobalVrfRpAddressId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ip6-address"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
			}
			tmp["ip6_address"] = flattenRouterMulticast6PimSmGlobalVrfRpAddressIp6Address(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["group"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
			}
			tmp["group"] = flattenRouterMulticast6PimSmGlobalVrfRpAddressGroup(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterMulticast6PimSmGlobalVrfRpAddressId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticast6PimSmGlobalVrfRpAddressIp6Address(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalVrfRpAddressGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterMulticast6(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("multicast_routing", flattenRouterMulticast6MulticastRouting(o["multicast-routing"], d, "multicast_routing", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-routing"]) {
			return fmt.Errorf("Error reading multicast_routing: %v", err)
		}
	}

	if err = d.Set("multicast_pmtu", flattenRouterMulticast6MulticastPmtu(o["multicast-pmtu"], d, "multicast_pmtu", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-pmtu"]) {
			return fmt.Errorf("Error reading multicast_pmtu: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("interface", flattenRouterMulticast6Interface(o["interface"], d, "interface", sv)); err != nil {
			if !fortiAPIPatch(o["interface"]) {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("interface"); ok {
			if err = d.Set("interface", flattenRouterMulticast6Interface(o["interface"], d, "interface", sv)); err != nil {
				if !fortiAPIPatch(o["interface"]) {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("pim_sm_global", flattenRouterMulticast6PimSmGlobal(o["pim-sm-global"], d, "pim_sm_global", sv)); err != nil {
			if !fortiAPIPatch(o["pim-sm-global"]) {
				return fmt.Errorf("Error reading pim_sm_global: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pim_sm_global"); ok {
			if err = d.Set("pim_sm_global", flattenRouterMulticast6PimSmGlobal(o["pim-sm-global"], d, "pim_sm_global", sv)); err != nil {
				if !fortiAPIPatch(o["pim-sm-global"]) {
					return fmt.Errorf("Error reading pim_sm_global: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("pim_sm_global_vrf", flattenRouterMulticast6PimSmGlobalVrf(o["pim-sm-global-vrf"], d, "pim_sm_global_vrf", sv)); err != nil {
			if !fortiAPIPatch(o["pim-sm-global-vrf"]) {
				return fmt.Errorf("Error reading pim_sm_global_vrf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pim_sm_global_vrf"); ok {
			if err = d.Set("pim_sm_global_vrf", flattenRouterMulticast6PimSmGlobalVrf(o["pim-sm-global-vrf"], d, "pim_sm_global_vrf", sv)); err != nil {
				if !fortiAPIPatch(o["pim-sm-global-vrf"]) {
					return fmt.Errorf("Error reading pim_sm_global_vrf: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterMulticast6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterMulticast6MulticastRouting(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6MulticastPmtu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6Interface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandRouterMulticast6InterfaceName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hello-interval"], _ = expandRouterMulticast6InterfaceHelloInterval(d, i["hello_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_holdtime"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hello-holdtime"], _ = expandRouterMulticast6InterfaceHelloHoldtime(d, i["hello_holdtime"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["hello-holdtime"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rp-candidate"], _ = expandRouterMulticast6InterfaceRpCandidate(d, i["rp_candidate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rp-candidate-group"], _ = expandRouterMulticast6InterfaceRpCandidateGroup(d, i["rp_candidate_group"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["rp-candidate-group"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rp-candidate-priority"], _ = expandRouterMulticast6InterfaceRpCandidatePriority(d, i["rp_candidate_priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rp-candidate-interval"], _ = expandRouterMulticast6InterfaceRpCandidateInterval(d, i["rp_candidate_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "static_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["static-group"], _ = expandRouterMulticast6InterfaceStaticGroup(d, i["static_group"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["static-group"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterMulticast6InterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6InterfaceHelloInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6InterfaceHelloHoldtime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6InterfaceRpCandidate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6InterfaceRpCandidateGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6InterfaceRpCandidatePriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6InterfaceRpCandidateInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6InterfaceStaticGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobal(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bsr_candidate"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["bsr-candidate"] = nil
		} else {
			result["bsr-candidate"], _ = expandRouterMulticast6PimSmGlobalBsrCandidate(d, i["bsr_candidate"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "bsr_interface"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["bsr-interface"] = nil
		} else {
			result["bsr-interface"], _ = expandRouterMulticast6PimSmGlobalBsrInterface(d, i["bsr_interface"], pre_append, sv)
		}
	} else if d.HasChange(pre_append) {
		result["bsr-interface"] = nil
	}
	pre_append = pre + ".0." + "bsr_priority"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["bsr-priority"] = nil
		} else {
			result["bsr-priority"], _ = expandRouterMulticast6PimSmGlobalBsrPriority(d, i["bsr_priority"], pre_append, sv)
		}
	} else if d.HasChange(pre_append) {
		result["bsr-priority"] = nil
	}
	pre_append = pre + ".0." + "bsr_hash"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["bsr-hash"] = nil
		} else {
			result["bsr-hash"], _ = expandRouterMulticast6PimSmGlobalBsrHash(d, i["bsr_hash"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "bsr_allow_quick_refresh"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["bsr-allow-quick-refresh"] = nil
		} else {
			result["bsr-allow-quick-refresh"], _ = expandRouterMulticast6PimSmGlobalBsrAllowQuickRefresh(d, i["bsr_allow_quick_refresh"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "cisco_crp_prefix"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["cisco-crp-prefix"] = nil
		} else {
			result["cisco-crp-prefix"], _ = expandRouterMulticast6PimSmGlobalCiscoCrpPrefix(d, i["cisco_crp_prefix"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "register_rate_limit"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["register-rate-limit"] = nil
		} else {
			result["register-rate-limit"], _ = expandRouterMulticast6PimSmGlobalRegisterRateLimit(d, i["register_rate_limit"], pre_append, sv)
		}
	} else if d.HasChange(pre_append) {
		result["register-rate-limit"] = nil
	}
	pre_append = pre + ".0." + "cisco_ignore_rp_set_priority"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["cisco-ignore-rp-set-priority"] = nil
		} else {
			result["cisco-ignore-rp-set-priority"], _ = expandRouterMulticast6PimSmGlobalCiscoIgnoreRpSetPriority(d, i["cisco_ignore_rp_set_priority"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "spt_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["spt-threshold"] = nil
		} else {
			result["spt-threshold"], _ = expandRouterMulticast6PimSmGlobalSptThreshold(d, i["spt_threshold"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "spt_threshold_group"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["spt-threshold-group"] = nil
		} else {
			result["spt-threshold-group"], _ = expandRouterMulticast6PimSmGlobalSptThresholdGroup(d, i["spt_threshold_group"], pre_append, sv)
		}
	} else if d.HasChange(pre_append) {
		result["spt-threshold-group"] = nil
	}
	pre_append = pre + ".0." + "pim_use_sdwan"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["pim-use-sdwan"] = nil
		} else {
			result["pim-use-sdwan"], _ = expandRouterMulticast6PimSmGlobalPimUseSdwan(d, i["pim_use_sdwan"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "rp_address"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["rp-address"] = make([]struct{}, 0)
		} else {
			result["rp-address"], _ = expandRouterMulticast6PimSmGlobalRpAddress(d, i["rp_address"], pre_append, sv)
		}
	} else {
		result["rp-address"] = make([]string, 0)
	}

	return result, nil
}

func expandRouterMulticast6PimSmGlobalBsrCandidate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalBsrInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalBsrPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalBsrHash(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalBsrAllowQuickRefresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalCiscoCrpPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalRegisterRateLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalCiscoIgnoreRpSetPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalSptThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalSptThresholdGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalPimUseSdwan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalRpAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandRouterMulticast6PimSmGlobalRpAddressId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip6-address"], _ = expandRouterMulticast6PimSmGlobalRpAddressIp6Address(d, i["ip6_address"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["group"], _ = expandRouterMulticast6PimSmGlobalRpAddressGroup(d, i["group"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["group"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterMulticast6PimSmGlobalRpAddressId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalRpAddressIp6Address(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalRpAddressGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrf"], _ = expandRouterMulticast6PimSmGlobalVrfVrf(d, i["vrf"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_candidate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bsr-candidate"], _ = expandRouterMulticast6PimSmGlobalVrfBsrCandidate(d, i["bsr_candidate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bsr-interface"], _ = expandRouterMulticast6PimSmGlobalVrfBsrInterface(d, i["bsr_interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["bsr-interface"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bsr-priority"], _ = expandRouterMulticast6PimSmGlobalVrfBsrPriority(d, i["bsr_priority"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["bsr-priority"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_hash"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bsr-hash"], _ = expandRouterMulticast6PimSmGlobalVrfBsrHash(d, i["bsr_hash"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_allow_quick_refresh"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bsr-allow-quick-refresh"], _ = expandRouterMulticast6PimSmGlobalVrfBsrAllowQuickRefresh(d, i["bsr_allow_quick_refresh"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cisco_crp_prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cisco-crp-prefix"], _ = expandRouterMulticast6PimSmGlobalVrfCiscoCrpPrefix(d, i["cisco_crp_prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rp-address"], _ = expandRouterMulticast6PimSmGlobalVrfRpAddress(d, i["rp_address"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["rp-address"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterMulticast6PimSmGlobalVrfVrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfBsrCandidate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfBsrInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfBsrPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfBsrHash(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfBsrAllowQuickRefresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfCiscoCrpPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfRpAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandRouterMulticast6PimSmGlobalVrfRpAddressId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip6-address"], _ = expandRouterMulticast6PimSmGlobalVrfRpAddressIp6Address(d, i["ip6_address"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["group"], _ = expandRouterMulticast6PimSmGlobalVrfRpAddressGroup(d, i["group"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["group"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterMulticast6PimSmGlobalVrfRpAddressId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfRpAddressIp6Address(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalVrfRpAddressGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMulticast6(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("multicast_routing"); ok {
		if setArgNil {
			obj["multicast-routing"] = nil
		} else {
			t, err := expandRouterMulticast6MulticastRouting(d, v, "multicast_routing", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["multicast-routing"] = t
			}
		}
	}

	if v, ok := d.GetOk("multicast_pmtu"); ok {
		if setArgNil {
			obj["multicast-pmtu"] = nil
		} else {
			t, err := expandRouterMulticast6MulticastPmtu(d, v, "multicast_pmtu", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["multicast-pmtu"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		if setArgNil {
			obj["interface"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterMulticast6Interface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("pim_sm_global"); ok {
		t, err := expandRouterMulticast6PimSmGlobal(d, v, "pim_sm_global", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pim-sm-global"] = t
		}
	}

	if v, ok := d.GetOk("pim_sm_global_vrf"); ok || d.HasChange("pim_sm_global_vrf") {
		if setArgNil {
			obj["pim-sm-global-vrf"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterMulticast6PimSmGlobalVrf(d, v, "pim_sm_global_vrf", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pim-sm-global-vrf"] = t
			}
		}
	}

	return &obj, nil
}
