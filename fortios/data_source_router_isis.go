// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IS-IS.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceRouterIsis() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterIsisRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"is_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"adv_passive_only": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"adv_passive_only6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_mode_l1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_mode_l2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_password_l1": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"auth_password_l2": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"auth_keychain_l1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_keychain_l2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_sendonly_l1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_sendonly_l2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ignore_lsp_errors": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lsp_gen_interval_l1": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lsp_gen_interval_l2": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lsp_refresh_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_lsp_lifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"spf_interval_exp_l1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"spf_interval_exp_l2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dynamic_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"adjacency_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"adjacency_check6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"overload_bit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"overload_bit_suppress": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"overload_bit_on_startup": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"default_originate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_originate6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"metric_style": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute_l1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute_l1_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute_l2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute_l2_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute6_l1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute6_l1_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute6_l2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute6_l2_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"isis_net": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"net": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"isis_interface": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"circuit_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"csnp_interval_l1": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"csnp_interval_l2": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"hello_interval_l1": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"hello_interval_l2": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"hello_multiplier_l1": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"hello_multiplier_l2": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"hello_padding": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"lsp_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"lsp_retransmit_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"metric_l1": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"metric_l2": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"wide_metric_l1": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"wide_metric_l2": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"auth_password_l1": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"auth_password_l2": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"auth_keychain_l1": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"auth_keychain_l2": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"auth_send_only_l1": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"auth_send_only_l2": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"auth_mode_l1": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"auth_mode_l2": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"priority_l1": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"priority_l2": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"mesh_group": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"mesh_group_id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"summary_address": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"summary_address6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"redistribute": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"metric": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"metric_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"routemap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"redistribute6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"metric": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"metric_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"routemap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterIsisRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "RouterIsis"

	o, err := c.ReadRouterIsis(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing RouterIsis: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterIsis(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterIsis from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterIsisIsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAdvPassiveOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAdvPassiveOnly6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAuthModeL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAuthModeL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAuthPasswordL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAuthPasswordL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAuthKeychainL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAuthKeychainL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAuthSendonlyL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAuthSendonlyL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIgnoreLspErrors(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisLspGenIntervalL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisLspGenIntervalL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisLspRefreshInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisMaxLspLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisSpfIntervalExpL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisSpfIntervalExpL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisDynamicHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAdjacencyCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisAdjacencyCheck6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisOverloadBit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisOverloadBitSuppress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisOverloadBitOnStartup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisDefaultOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisDefaultOriginate6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisMetricStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistributeL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistributeL1List(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistributeL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistributeL2List(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistribute6L1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistribute6L1List(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistribute6L2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistribute6L2List(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisNet(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterIsisIsisNetId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "net"
		if _, ok := i["net"]; ok {
			tmp["net"] = dataSourceFlattenRouterIsisIsisNetNet(i["net"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterIsisIsisNetId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisNetNet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterIsisIsisInterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenRouterIsisIsisInterfaceStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status6"
		if _, ok := i["status6"]; ok {
			tmp["status6"] = dataSourceFlattenRouterIsisIsisInterfaceStatus6(i["status6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_type"
		if _, ok := i["network-type"]; ok {
			tmp["network_type"] = dataSourceFlattenRouterIsisIsisInterfaceNetworkType(i["network-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_type"
		if _, ok := i["circuit-type"]; ok {
			tmp["circuit_type"] = dataSourceFlattenRouterIsisIsisInterfaceCircuitType(i["circuit-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "csnp_interval_l1"
		if _, ok := i["csnp-interval-l1"]; ok {
			tmp["csnp_interval_l1"] = dataSourceFlattenRouterIsisIsisInterfaceCsnpIntervalL1(i["csnp-interval-l1"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "csnp_interval_l2"
		if _, ok := i["csnp-interval-l2"]; ok {
			tmp["csnp_interval_l2"] = dataSourceFlattenRouterIsisIsisInterfaceCsnpIntervalL2(i["csnp-interval-l2"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval_l1"
		if _, ok := i["hello-interval-l1"]; ok {
			tmp["hello_interval_l1"] = dataSourceFlattenRouterIsisIsisInterfaceHelloIntervalL1(i["hello-interval-l1"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval_l2"
		if _, ok := i["hello-interval-l2"]; ok {
			tmp["hello_interval_l2"] = dataSourceFlattenRouterIsisIsisInterfaceHelloIntervalL2(i["hello-interval-l2"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier_l1"
		if _, ok := i["hello-multiplier-l1"]; ok {
			tmp["hello_multiplier_l1"] = dataSourceFlattenRouterIsisIsisInterfaceHelloMultiplierL1(i["hello-multiplier-l1"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier_l2"
		if _, ok := i["hello-multiplier-l2"]; ok {
			tmp["hello_multiplier_l2"] = dataSourceFlattenRouterIsisIsisInterfaceHelloMultiplierL2(i["hello-multiplier-l2"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_padding"
		if _, ok := i["hello-padding"]; ok {
			tmp["hello_padding"] = dataSourceFlattenRouterIsisIsisInterfaceHelloPadding(i["hello-padding"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lsp_interval"
		if _, ok := i["lsp-interval"]; ok {
			tmp["lsp_interval"] = dataSourceFlattenRouterIsisIsisInterfaceLspInterval(i["lsp-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lsp_retransmit_interval"
		if _, ok := i["lsp-retransmit-interval"]; ok {
			tmp["lsp_retransmit_interval"] = dataSourceFlattenRouterIsisIsisInterfaceLspRetransmitInterval(i["lsp-retransmit-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_l1"
		if _, ok := i["metric-l1"]; ok {
			tmp["metric_l1"] = dataSourceFlattenRouterIsisIsisInterfaceMetricL1(i["metric-l1"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_l2"
		if _, ok := i["metric-l2"]; ok {
			tmp["metric_l2"] = dataSourceFlattenRouterIsisIsisInterfaceMetricL2(i["metric-l2"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wide_metric_l1"
		if _, ok := i["wide-metric-l1"]; ok {
			tmp["wide_metric_l1"] = dataSourceFlattenRouterIsisIsisInterfaceWideMetricL1(i["wide-metric-l1"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wide_metric_l2"
		if _, ok := i["wide-metric-l2"]; ok {
			tmp["wide_metric_l2"] = dataSourceFlattenRouterIsisIsisInterfaceWideMetricL2(i["wide-metric-l2"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_keychain_l1"
		if _, ok := i["auth-keychain-l1"]; ok {
			tmp["auth_keychain_l1"] = dataSourceFlattenRouterIsisIsisInterfaceAuthKeychainL1(i["auth-keychain-l1"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_keychain_l2"
		if _, ok := i["auth-keychain-l2"]; ok {
			tmp["auth_keychain_l2"] = dataSourceFlattenRouterIsisIsisInterfaceAuthKeychainL2(i["auth-keychain-l2"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_send_only_l1"
		if _, ok := i["auth-send-only-l1"]; ok {
			tmp["auth_send_only_l1"] = dataSourceFlattenRouterIsisIsisInterfaceAuthSendOnlyL1(i["auth-send-only-l1"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_send_only_l2"
		if _, ok := i["auth-send-only-l2"]; ok {
			tmp["auth_send_only_l2"] = dataSourceFlattenRouterIsisIsisInterfaceAuthSendOnlyL2(i["auth-send-only-l2"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode_l1"
		if _, ok := i["auth-mode-l1"]; ok {
			tmp["auth_mode_l1"] = dataSourceFlattenRouterIsisIsisInterfaceAuthModeL1(i["auth-mode-l1"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode_l2"
		if _, ok := i["auth-mode-l2"]; ok {
			tmp["auth_mode_l2"] = dataSourceFlattenRouterIsisIsisInterfaceAuthModeL2(i["auth-mode-l2"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_l1"
		if _, ok := i["priority-l1"]; ok {
			tmp["priority_l1"] = dataSourceFlattenRouterIsisIsisInterfacePriorityL1(i["priority-l1"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_l2"
		if _, ok := i["priority-l2"]; ok {
			tmp["priority_l2"] = dataSourceFlattenRouterIsisIsisInterfacePriorityL2(i["priority-l2"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mesh_group"
		if _, ok := i["mesh-group"]; ok {
			tmp["mesh_group"] = dataSourceFlattenRouterIsisIsisInterfaceMeshGroup(i["mesh-group"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mesh_group_id"
		if _, ok := i["mesh-group-id"]; ok {
			tmp["mesh_group_id"] = dataSourceFlattenRouterIsisIsisInterfaceMeshGroupId(i["mesh-group-id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterIsisIsisInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceStatus6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceNetworkType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceCircuitType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceCsnpIntervalL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceCsnpIntervalL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceHelloIntervalL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceHelloIntervalL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceHelloMultiplierL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceHelloMultiplierL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceHelloPadding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceLspInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceLspRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceMetricL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceMetricL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceWideMetricL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceWideMetricL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceAuthPasswordL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceAuthPasswordL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceAuthKeychainL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceAuthKeychainL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceAuthSendOnlyL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceAuthSendOnlyL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceAuthModeL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceAuthModeL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfacePriorityL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfacePriorityL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceMeshGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisIsisInterfaceMeshGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisSummaryAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterIsisSummaryAddressId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = dataSourceFlattenRouterIsisSummaryAddressPrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {
			tmp["level"] = dataSourceFlattenRouterIsisSummaryAddressLevel(i["level"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterIsisSummaryAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisSummaryAddressPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterIsisSummaryAddressLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisSummaryAddress6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterIsisSummaryAddress6Id(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			tmp["prefix6"] = dataSourceFlattenRouterIsisSummaryAddress6Prefix6(i["prefix6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {
			tmp["level"] = dataSourceFlattenRouterIsisSummaryAddress6Level(i["level"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterIsisSummaryAddress6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisSummaryAddress6Prefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisSummaryAddress6Level(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			tmp["protocol"] = dataSourceFlattenRouterIsisRedistributeProtocol(i["protocol"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenRouterIsisRedistributeStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {
			tmp["metric"] = dataSourceFlattenRouterIsisRedistributeMetric(i["metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_type"
		if _, ok := i["metric-type"]; ok {
			tmp["metric_type"] = dataSourceFlattenRouterIsisRedistributeMetricType(i["metric-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {
			tmp["level"] = dataSourceFlattenRouterIsisRedistributeLevel(i["level"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {
			tmp["routemap"] = dataSourceFlattenRouterIsisRedistributeRoutemap(i["routemap"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterIsisRedistributeProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistributeMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistributeLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistribute6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			tmp["protocol"] = dataSourceFlattenRouterIsisRedistribute6Protocol(i["protocol"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenRouterIsisRedistribute6Status(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {
			tmp["metric"] = dataSourceFlattenRouterIsisRedistribute6Metric(i["metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_type"
		if _, ok := i["metric-type"]; ok {
			tmp["metric_type"] = dataSourceFlattenRouterIsisRedistribute6MetricType(i["metric-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {
			tmp["level"] = dataSourceFlattenRouterIsisRedistribute6Level(i["level"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {
			tmp["routemap"] = dataSourceFlattenRouterIsisRedistribute6Routemap(i["routemap"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterIsisRedistribute6Protocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistribute6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistribute6Metric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistribute6MetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistribute6Level(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterIsisRedistribute6Routemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterIsis(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("is_type", dataSourceFlattenRouterIsisIsType(o["is-type"], d, "is_type")); err != nil {
		if !fortiAPIPatch(o["is-type"]) {
			return fmt.Errorf("Error reading is_type: %v", err)
		}
	}

	if err = d.Set("adv_passive_only", dataSourceFlattenRouterIsisAdvPassiveOnly(o["adv-passive-only"], d, "adv_passive_only")); err != nil {
		if !fortiAPIPatch(o["adv-passive-only"]) {
			return fmt.Errorf("Error reading adv_passive_only: %v", err)
		}
	}

	if err = d.Set("adv_passive_only6", dataSourceFlattenRouterIsisAdvPassiveOnly6(o["adv-passive-only6"], d, "adv_passive_only6")); err != nil {
		if !fortiAPIPatch(o["adv-passive-only6"]) {
			return fmt.Errorf("Error reading adv_passive_only6: %v", err)
		}
	}

	if err = d.Set("auth_mode_l1", dataSourceFlattenRouterIsisAuthModeL1(o["auth-mode-l1"], d, "auth_mode_l1")); err != nil {
		if !fortiAPIPatch(o["auth-mode-l1"]) {
			return fmt.Errorf("Error reading auth_mode_l1: %v", err)
		}
	}

	if err = d.Set("auth_mode_l2", dataSourceFlattenRouterIsisAuthModeL2(o["auth-mode-l2"], d, "auth_mode_l2")); err != nil {
		if !fortiAPIPatch(o["auth-mode-l2"]) {
			return fmt.Errorf("Error reading auth_mode_l2: %v", err)
		}
	}

	if err = d.Set("auth_keychain_l1", dataSourceFlattenRouterIsisAuthKeychainL1(o["auth-keychain-l1"], d, "auth_keychain_l1")); err != nil {
		if !fortiAPIPatch(o["auth-keychain-l1"]) {
			return fmt.Errorf("Error reading auth_keychain_l1: %v", err)
		}
	}

	if err = d.Set("auth_keychain_l2", dataSourceFlattenRouterIsisAuthKeychainL2(o["auth-keychain-l2"], d, "auth_keychain_l2")); err != nil {
		if !fortiAPIPatch(o["auth-keychain-l2"]) {
			return fmt.Errorf("Error reading auth_keychain_l2: %v", err)
		}
	}

	if err = d.Set("auth_sendonly_l1", dataSourceFlattenRouterIsisAuthSendonlyL1(o["auth-sendonly-l1"], d, "auth_sendonly_l1")); err != nil {
		if !fortiAPIPatch(o["auth-sendonly-l1"]) {
			return fmt.Errorf("Error reading auth_sendonly_l1: %v", err)
		}
	}

	if err = d.Set("auth_sendonly_l2", dataSourceFlattenRouterIsisAuthSendonlyL2(o["auth-sendonly-l2"], d, "auth_sendonly_l2")); err != nil {
		if !fortiAPIPatch(o["auth-sendonly-l2"]) {
			return fmt.Errorf("Error reading auth_sendonly_l2: %v", err)
		}
	}

	if err = d.Set("ignore_lsp_errors", dataSourceFlattenRouterIsisIgnoreLspErrors(o["ignore-lsp-errors"], d, "ignore_lsp_errors")); err != nil {
		if !fortiAPIPatch(o["ignore-lsp-errors"]) {
			return fmt.Errorf("Error reading ignore_lsp_errors: %v", err)
		}
	}

	if err = d.Set("lsp_gen_interval_l1", dataSourceFlattenRouterIsisLspGenIntervalL1(o["lsp-gen-interval-l1"], d, "lsp_gen_interval_l1")); err != nil {
		if !fortiAPIPatch(o["lsp-gen-interval-l1"]) {
			return fmt.Errorf("Error reading lsp_gen_interval_l1: %v", err)
		}
	}

	if err = d.Set("lsp_gen_interval_l2", dataSourceFlattenRouterIsisLspGenIntervalL2(o["lsp-gen-interval-l2"], d, "lsp_gen_interval_l2")); err != nil {
		if !fortiAPIPatch(o["lsp-gen-interval-l2"]) {
			return fmt.Errorf("Error reading lsp_gen_interval_l2: %v", err)
		}
	}

	if err = d.Set("lsp_refresh_interval", dataSourceFlattenRouterIsisLspRefreshInterval(o["lsp-refresh-interval"], d, "lsp_refresh_interval")); err != nil {
		if !fortiAPIPatch(o["lsp-refresh-interval"]) {
			return fmt.Errorf("Error reading lsp_refresh_interval: %v", err)
		}
	}

	if err = d.Set("max_lsp_lifetime", dataSourceFlattenRouterIsisMaxLspLifetime(o["max-lsp-lifetime"], d, "max_lsp_lifetime")); err != nil {
		if !fortiAPIPatch(o["max-lsp-lifetime"]) {
			return fmt.Errorf("Error reading max_lsp_lifetime: %v", err)
		}
	}

	if err = d.Set("spf_interval_exp_l1", dataSourceFlattenRouterIsisSpfIntervalExpL1(o["spf-interval-exp-l1"], d, "spf_interval_exp_l1")); err != nil {
		if !fortiAPIPatch(o["spf-interval-exp-l1"]) {
			return fmt.Errorf("Error reading spf_interval_exp_l1: %v", err)
		}
	}

	if err = d.Set("spf_interval_exp_l2", dataSourceFlattenRouterIsisSpfIntervalExpL2(o["spf-interval-exp-l2"], d, "spf_interval_exp_l2")); err != nil {
		if !fortiAPIPatch(o["spf-interval-exp-l2"]) {
			return fmt.Errorf("Error reading spf_interval_exp_l2: %v", err)
		}
	}

	if err = d.Set("dynamic_hostname", dataSourceFlattenRouterIsisDynamicHostname(o["dynamic-hostname"], d, "dynamic_hostname")); err != nil {
		if !fortiAPIPatch(o["dynamic-hostname"]) {
			return fmt.Errorf("Error reading dynamic_hostname: %v", err)
		}
	}

	if err = d.Set("adjacency_check", dataSourceFlattenRouterIsisAdjacencyCheck(o["adjacency-check"], d, "adjacency_check")); err != nil {
		if !fortiAPIPatch(o["adjacency-check"]) {
			return fmt.Errorf("Error reading adjacency_check: %v", err)
		}
	}

	if err = d.Set("adjacency_check6", dataSourceFlattenRouterIsisAdjacencyCheck6(o["adjacency-check6"], d, "adjacency_check6")); err != nil {
		if !fortiAPIPatch(o["adjacency-check6"]) {
			return fmt.Errorf("Error reading adjacency_check6: %v", err)
		}
	}

	if err = d.Set("overload_bit", dataSourceFlattenRouterIsisOverloadBit(o["overload-bit"], d, "overload_bit")); err != nil {
		if !fortiAPIPatch(o["overload-bit"]) {
			return fmt.Errorf("Error reading overload_bit: %v", err)
		}
	}

	if err = d.Set("overload_bit_suppress", dataSourceFlattenRouterIsisOverloadBitSuppress(o["overload-bit-suppress"], d, "overload_bit_suppress")); err != nil {
		if !fortiAPIPatch(o["overload-bit-suppress"]) {
			return fmt.Errorf("Error reading overload_bit_suppress: %v", err)
		}
	}

	if err = d.Set("overload_bit_on_startup", dataSourceFlattenRouterIsisOverloadBitOnStartup(o["overload-bit-on-startup"], d, "overload_bit_on_startup")); err != nil {
		if !fortiAPIPatch(o["overload-bit-on-startup"]) {
			return fmt.Errorf("Error reading overload_bit_on_startup: %v", err)
		}
	}

	if err = d.Set("default_originate", dataSourceFlattenRouterIsisDefaultOriginate(o["default-originate"], d, "default_originate")); err != nil {
		if !fortiAPIPatch(o["default-originate"]) {
			return fmt.Errorf("Error reading default_originate: %v", err)
		}
	}

	if err = d.Set("default_originate6", dataSourceFlattenRouterIsisDefaultOriginate6(o["default-originate6"], d, "default_originate6")); err != nil {
		if !fortiAPIPatch(o["default-originate6"]) {
			return fmt.Errorf("Error reading default_originate6: %v", err)
		}
	}

	if err = d.Set("metric_style", dataSourceFlattenRouterIsisMetricStyle(o["metric-style"], d, "metric_style")); err != nil {
		if !fortiAPIPatch(o["metric-style"]) {
			return fmt.Errorf("Error reading metric_style: %v", err)
		}
	}

	if err = d.Set("redistribute_l1", dataSourceFlattenRouterIsisRedistributeL1(o["redistribute-l1"], d, "redistribute_l1")); err != nil {
		if !fortiAPIPatch(o["redistribute-l1"]) {
			return fmt.Errorf("Error reading redistribute_l1: %v", err)
		}
	}

	if err = d.Set("redistribute_l1_list", dataSourceFlattenRouterIsisRedistributeL1List(o["redistribute-l1-list"], d, "redistribute_l1_list")); err != nil {
		if !fortiAPIPatch(o["redistribute-l1-list"]) {
			return fmt.Errorf("Error reading redistribute_l1_list: %v", err)
		}
	}

	if err = d.Set("redistribute_l2", dataSourceFlattenRouterIsisRedistributeL2(o["redistribute-l2"], d, "redistribute_l2")); err != nil {
		if !fortiAPIPatch(o["redistribute-l2"]) {
			return fmt.Errorf("Error reading redistribute_l2: %v", err)
		}
	}

	if err = d.Set("redistribute_l2_list", dataSourceFlattenRouterIsisRedistributeL2List(o["redistribute-l2-list"], d, "redistribute_l2_list")); err != nil {
		if !fortiAPIPatch(o["redistribute-l2-list"]) {
			return fmt.Errorf("Error reading redistribute_l2_list: %v", err)
		}
	}

	if err = d.Set("redistribute6_l1", dataSourceFlattenRouterIsisRedistribute6L1(o["redistribute6-l1"], d, "redistribute6_l1")); err != nil {
		if !fortiAPIPatch(o["redistribute6-l1"]) {
			return fmt.Errorf("Error reading redistribute6_l1: %v", err)
		}
	}

	if err = d.Set("redistribute6_l1_list", dataSourceFlattenRouterIsisRedistribute6L1List(o["redistribute6-l1-list"], d, "redistribute6_l1_list")); err != nil {
		if !fortiAPIPatch(o["redistribute6-l1-list"]) {
			return fmt.Errorf("Error reading redistribute6_l1_list: %v", err)
		}
	}

	if err = d.Set("redistribute6_l2", dataSourceFlattenRouterIsisRedistribute6L2(o["redistribute6-l2"], d, "redistribute6_l2")); err != nil {
		if !fortiAPIPatch(o["redistribute6-l2"]) {
			return fmt.Errorf("Error reading redistribute6_l2: %v", err)
		}
	}

	if err = d.Set("redistribute6_l2_list", dataSourceFlattenRouterIsisRedistribute6L2List(o["redistribute6-l2-list"], d, "redistribute6_l2_list")); err != nil {
		if !fortiAPIPatch(o["redistribute6-l2-list"]) {
			return fmt.Errorf("Error reading redistribute6_l2_list: %v", err)
		}
	}

	if err = d.Set("isis_net", dataSourceFlattenRouterIsisIsisNet(o["isis-net"], d, "isis_net")); err != nil {
		if !fortiAPIPatch(o["isis-net"]) {
			return fmt.Errorf("Error reading isis_net: %v", err)
		}
	}

	if err = d.Set("isis_interface", dataSourceFlattenRouterIsisIsisInterface(o["isis-interface"], d, "isis_interface")); err != nil {
		if !fortiAPIPatch(o["isis-interface"]) {
			return fmt.Errorf("Error reading isis_interface: %v", err)
		}
	}

	if err = d.Set("summary_address", dataSourceFlattenRouterIsisSummaryAddress(o["summary-address"], d, "summary_address")); err != nil {
		if !fortiAPIPatch(o["summary-address"]) {
			return fmt.Errorf("Error reading summary_address: %v", err)
		}
	}

	if err = d.Set("summary_address6", dataSourceFlattenRouterIsisSummaryAddress6(o["summary-address6"], d, "summary_address6")); err != nil {
		if !fortiAPIPatch(o["summary-address6"]) {
			return fmt.Errorf("Error reading summary_address6: %v", err)
		}
	}

	if err = d.Set("redistribute", dataSourceFlattenRouterIsisRedistribute(o["redistribute"], d, "redistribute")); err != nil {
		if !fortiAPIPatch(o["redistribute"]) {
			return fmt.Errorf("Error reading redistribute: %v", err)
		}
	}

	if err = d.Set("redistribute6", dataSourceFlattenRouterIsisRedistribute6(o["redistribute6"], d, "redistribute6")); err != nil {
		if !fortiAPIPatch(o["redistribute6"]) {
			return fmt.Errorf("Error reading redistribute6: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterIsisFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
