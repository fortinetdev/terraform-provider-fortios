// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure NPU attributes.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemNpu() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNpuUpdate,
		Read:   resourceSystemNpuRead,
		Update: resourceSystemNpuUpdate,
		Delete: resourceSystemNpuDelete,

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
			"dedicated_management_cpu": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dedicated_lacp_queue": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dedicated_management_affinity": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"dos_options": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"npu_dos_meter_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"npu_dos_tpe_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"policy_offload_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"napi_break_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
			},
			"hpe": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all_protocol": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 32000000),
							Optional:     true,
							Computed:     true,
						},
						"tcpsyn_max": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),
							Optional:     true,
							Computed:     true,
						},
						"tcpsyn_ack_max": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),
							Optional:     true,
							Computed:     true,
						},
						"tcpfin_rst_max": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),
							Optional:     true,
							Computed:     true,
						},
						"tcp_max": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),
							Optional:     true,
							Computed:     true,
						},
						"udp_max": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),
							Optional:     true,
							Computed:     true,
						},
						"icmp_max": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),
							Optional:     true,
							Computed:     true,
						},
						"sctp_max": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),
							Optional:     true,
							Computed:     true,
						},
						"esp_max": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),
							Optional:     true,
							Computed:     true,
						},
						"ip_frag_max": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),
							Optional:     true,
							Computed:     true,
						},
						"ip_others_max": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),
							Optional:     true,
							Computed:     true,
						},
						"arp_max": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),
							Optional:     true,
							Computed:     true,
						},
						"l2_others_max": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),
							Optional:     true,
							Computed:     true,
						},
						"high_priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1000, 32000000),
							Optional:     true,
							Computed:     true,
						},
						"enable_shaper": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fastpath": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capwap_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vxlan_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vxlan_mac_flapping_guard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_qos_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_ipsec_mcs_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"shaping_stats": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mcs_host_packet_tpe_shaping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gtp_support": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"per_session_accounting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_acct_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 600),
				Optional:     true,
				Computed:     true,
			},
			"max_session_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 1000),
				Optional:     true,
				Computed:     true,
			},
			"fp_anomaly": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_syn_fin": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_fin_noack": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_fin_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_no_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_syn_data": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_winnuke": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_land": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"udp_land": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"icmp_land": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"icmp_frag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_land": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_proto_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_unknopt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_optrr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_optssrr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_optlsrr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_optstream": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_optsecurity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_opttimestamp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_csum_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_csum_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"udp_csum_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"udplite_csum_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"icmp_csum_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gre_csum_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sctp_csum_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_land": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_proto_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_unknopt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_saddr_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_daddr_err": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_optralert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_optjumbo": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_opttunnel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_opthomeaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_optnsap": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_optendpid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_optinvld": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ip_reassembly": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"min_timeout": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 600000000),
							Optional:     true,
							Computed:     true,
						},
						"max_timeout": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 600000000),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"hash_tbl_spread": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_lookup_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_fragment_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"htx_icmp_csum_chk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"htab_msg_queue": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"htab_dedi_queue_nr": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2),
				Optional:     true,
				Computed:     true,
			},
			"dsw_dts_profile": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"profile_id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 32),
							Optional:     true,
						},
						"min_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(32, 2048),
							Optional:     true,
						},
						"step": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 64),
							Optional:     true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dsw_queue_dts_profile": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"iport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"oport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"profile_id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 32),
							Optional:     true,
						},
						"queue_select": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4095),
							Optional:     true,
						},
					},
				},
			},
			"npu_tcam": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"oid": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4095),
							Optional:     true,
						},
						"vid": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4095),
							Optional:     true,
						},
						"data": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gen_buf_cnt": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 3),
										Optional:     true,
									},
									"gen_pri": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 7),
										Optional:     true,
									},
									"gen_pri_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"gen_iv": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"gen_tv": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"gen_pkt_ctrl": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"gen_l3_flags": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),
										Optional:     true,
									},
									"gen_l4_flags": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),
										Optional:     true,
									},
									"vdid": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"tp": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4095),
										Optional:     true,
									},
									"tgt_updt": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"smac_change": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ext_tag": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tgt_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tvid": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4095),
										Optional:     true,
									},
									"tgt_cfi": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tgt_prio": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 7),
										Optional:     true,
									},
									"sp": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4095),
										Optional:     true,
									},
									"src_updt": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"slink": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 7),
										Optional:     true,
									},
									"svid": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4095),
										Optional:     true,
									},
									"src_cfi": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"src_prio": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 7),
										Optional:     true,
									},
									"srcmac": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dstmac": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ethertype": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ipver": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),
										Optional:     true,
									},
									"ihl": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),
										Optional:     true,
									},
									"ip4_id": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"srcip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dstip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ip6_fl": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1048575),
										Optional:     true,
									},
									"srcipv6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dstipv6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ttl": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),
										Optional:     true,
									},
									"protocol": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),
										Optional:     true,
									},
									"tos": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),
										Optional:     true,
									},
									"frag_off": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 31),
										Optional:     true,
									},
									"mf": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"df": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"srcport": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"dstport": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"tcp_fin": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_syn": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_rst": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_push": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_ack": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_urg": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_ece": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_cwr": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"l4_wd8": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"l4_wd9": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"l4_wd10": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"l4_wd11": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
								},
							},
						},
						"mask": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gen_buf_cnt": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 3),
										Optional:     true,
									},
									"gen_pri": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 7),
										Optional:     true,
									},
									"gen_pri_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"gen_iv": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"gen_tv": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"gen_pkt_ctrl": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"gen_l3_flags": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),
										Optional:     true,
									},
									"gen_l4_flags": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),
										Optional:     true,
									},
									"vdid": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"tp": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4095),
										Optional:     true,
									},
									"tgt_updt": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"smac_change": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ext_tag": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tgt_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tvid": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4095),
										Optional:     true,
									},
									"tgt_cfi": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tgt_prio": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 7),
										Optional:     true,
									},
									"sp": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4095),
										Optional:     true,
									},
									"src_updt": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"slink": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 7),
										Optional:     true,
									},
									"svid": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4095),
										Optional:     true,
									},
									"src_cfi": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"src_prio": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 7),
										Optional:     true,
									},
									"srcmac": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dstmac": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ethertype": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ipver": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),
										Optional:     true,
									},
									"ihl": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),
										Optional:     true,
									},
									"ip4_id": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"srcip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dstip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ip6_fl": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1048575),
										Optional:     true,
									},
									"srcipv6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dstipv6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ttl": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),
										Optional:     true,
									},
									"protocol": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),
										Optional:     true,
									},
									"tos": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),
										Optional:     true,
									},
									"frag_off": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 31),
										Optional:     true,
									},
									"mf": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"df": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"srcport": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"dstport": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"tcp_fin": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_syn": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_rst": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_push": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_ack": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_urg": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_ece": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_cwr": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"l4_wd8": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"l4_wd9": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"l4_wd10": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"l4_wd11": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
								},
							},
						},
						"mir_act": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vlif": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 16777215),
										Optional:     true,
									},
								},
							},
						},
						"pri_act": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"priority": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),
										Optional:     true,
									},
									"weight": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),
										Optional:     true,
									},
								},
							},
						},
						"sact": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fwd_lif_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"fwd_lif": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4095),
										Optional:     true,
									},
									"fwd_tvid_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"fwd_tvid": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4095),
										Optional:     true,
									},
									"df_lif_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"df_lif": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4095),
										Optional:     true,
									},
									"act_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"act": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 3),
										Optional:     true,
									},
									"pleen_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"pleen": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"icpen_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"icpen": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"vdm_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vdm": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"learn_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"learn": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"rfsh_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"rfsh": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"fwd_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"fwd": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"x_mode_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"x_mode": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 3),
										Optional:     true,
									},
									"promis_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"promis": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"bmproc_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"bmproc": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"mac_id_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mac_id": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"dosen_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dosen": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"dfr_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dfr": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"m_srh_ctrl_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"m_srh_ctrl": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"tpe_id_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tpe_id": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 16383),
										Optional:     true,
									},
									"vdom_id_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vdom_id": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 16383),
										Optional:     true,
									},
									"mss_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mss": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 16383),
										Optional:     true,
									},
									"tp_smchk_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tp_smchk": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"etype_pid_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"etype_pid": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),
										Optional:     true,
									},
									"frag_proc_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"frag_proc": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"espff_proc_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"espff_proc": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"prio_pid_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"prio_pid": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 7),
										Optional:     true,
									},
									"igmp_mld_snp_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"igmp_mld_snp": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"smac_skip_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"smac_skip": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"dmac_skip_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dmac_skip": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
								},
							},
						},
						"tact": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"act_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"act": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 3),
										Optional:     true,
									},
									"mtuv4_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mtuv4": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 16383),
										Optional:     true,
									},
									"mtuv6_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mtuv6": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 16383),
										Optional:     true,
									},
									"mac_id_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mac_id": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"slif_act_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"slif_act": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 3),
										Optional:     true,
									},
									"tlif_act_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tlif_act": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 3),
										Optional:     true,
									},
									"tgtv_act_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tgtv_act": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"tpeid_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tpeid": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 16383),
										Optional:     true,
									},
									"v6fe_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"v6fe": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"xlt_vid_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"xlt_vid": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4095),
										Optional:     true,
									},
									"xlt_lif_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"xlt_lif": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4095),
										Optional:     true,
									},
									"mss_t_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mss_t": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 16383),
										Optional:     true,
									},
									"lnkid_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"lnkid": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),
										Optional:     true,
									},
									"sublnkid_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sublnkid": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 511),
										Optional:     true,
									},
									"fmtuv4_s_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"fmtuv4_s": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"fmtuv6_s_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"fmtuv6_s": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"vep_en_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vep_en": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1),
										Optional:     true,
									},
									"vep_slid_v": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vep_slid": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 3),
										Optional:     true,
									},
								},
							},
						},
					},
				},
			},
			"np_queues": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"profile": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),
										Optional:     true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"weight": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),
										Optional:     true,
										Computed:     true,
									},
									"cos0": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"cos1": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"cos2": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"cos3": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"cos4": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"cos5": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"cos6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"cos7": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp0": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp1": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp2": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp3": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp4": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp5": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp7": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp8": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp9": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp10": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp11": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp12": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp13": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp14": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp15": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp16": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp17": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp18": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp19": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp20": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp21": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp22": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp23": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp24": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp25": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp26": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp27": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp28": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp29": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp30": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp31": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp32": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp33": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp34": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp35": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp36": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp37": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp38": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp39": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp40": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp41": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp42": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp43": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp44": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp45": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp46": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp47": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp48": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp49": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp50": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp51": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp52": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp53": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp54": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp55": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp56": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp57": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp58": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp59": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp60": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp61": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp62": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dscp63": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"ethernet_type": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"queue": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 11),
										Optional:     true,
									},
									"weight": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"ip_protocol": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
									"protocol": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),
										Optional:     true,
									},
									"queue": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 11),
										Optional:     true,
									},
									"weight": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"ip_service": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
									"protocol": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),
										Optional:     true,
									},
									"sport": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"dport": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"queue": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 11),
										Optional:     true,
									},
									"weight": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 15),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"scheduler": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
									"mode": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"custom_etype_lookup": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"sw_np_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sw_np_rate_unit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sw_np_rate_burst": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"double_level_mcast_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mcast_denied_ses_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_enc_subengine_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipsec_dec_subengine_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"np6_cps_optimization_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sw_np_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strip_esp_padding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strip_clear_text_padding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_inbound_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sse_backpressure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rdp_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_over_vlink": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uesp_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mcast_session_accounting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_mtu_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_denied_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qtm_buf_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_ob_np_sel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_receive_unit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(64, 10000),
				Optional:     true,
				Computed:     true,
			},
			"lag_hash_gre": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_mse_oft": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike_port": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
						},
					},
				},
			},
			"priority_protocol": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bgp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"slbc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bfd": &schema.Schema{
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
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemNpuUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemNpu(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemNpu resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemNpu(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemNpu resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemNpu")
	}

	return resourceSystemNpuRead(d, m)
}

func resourceSystemNpuDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemNpu(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemNpu resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemNpu(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemNpu resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNpuRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemNpu(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemNpu resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNpu(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemNpu resource from API: %v", err)
	}
	return nil
}

func flattenSystemNpuDedicatedManagementCpu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuDedicatedLacpQueue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuDedicatedManagementAffinity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuDosOptions(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "npu_dos_meter_mode"
	if _, ok := i["npu-dos-meter-mode"]; ok {
		result["npu_dos_meter_mode"] = flattenSystemNpuDosOptionsNpuDosMeterMode(i["npu-dos-meter-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "npu_dos_tpe_mode"
	if _, ok := i["npu-dos-tpe-mode"]; ok {
		result["npu_dos_tpe_mode"] = flattenSystemNpuDosOptionsNpuDosTpeMode(i["npu-dos-tpe-mode"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemNpuDosOptionsNpuDosMeterMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuDosOptionsNpuDosTpeMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuPolicyOffloadLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNapiBreakInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuHpe(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "all_protocol"
	if _, ok := i["all-protocol"]; ok {
		result["all_protocol"] = flattenSystemNpuHpeAllProtocol(i["all-protocol"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcpsyn_max"
	if _, ok := i["tcpsyn-max"]; ok {
		result["tcpsyn_max"] = flattenSystemNpuHpeTcpsynMax(i["tcpsyn-max"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcpsyn_ack_max"
	if _, ok := i["tcpsyn-ack-max"]; ok {
		result["tcpsyn_ack_max"] = flattenSystemNpuHpeTcpsynAckMax(i["tcpsyn-ack-max"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcpfin_rst_max"
	if _, ok := i["tcpfin-rst-max"]; ok {
		result["tcpfin_rst_max"] = flattenSystemNpuHpeTcpfinRstMax(i["tcpfin-rst-max"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_max"
	if _, ok := i["tcp-max"]; ok {
		result["tcp_max"] = flattenSystemNpuHpeTcpMax(i["tcp-max"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "udp_max"
	if _, ok := i["udp-max"]; ok {
		result["udp_max"] = flattenSystemNpuHpeUdpMax(i["udp-max"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "icmp_max"
	if _, ok := i["icmp-max"]; ok {
		result["icmp_max"] = flattenSystemNpuHpeIcmpMax(i["icmp-max"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sctp_max"
	if _, ok := i["sctp-max"]; ok {
		result["sctp_max"] = flattenSystemNpuHpeSctpMax(i["sctp-max"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "esp_max"
	if _, ok := i["esp-max"]; ok {
		result["esp_max"] = flattenSystemNpuHpeEspMax(i["esp-max"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ip_frag_max"
	if _, ok := i["ip-frag-max"]; ok {
		result["ip_frag_max"] = flattenSystemNpuHpeIpFragMax(i["ip-frag-max"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ip_others_max"
	if _, ok := i["ip-others-max"]; ok {
		result["ip_others_max"] = flattenSystemNpuHpeIpOthersMax(i["ip-others-max"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "arp_max"
	if _, ok := i["arp-max"]; ok {
		result["arp_max"] = flattenSystemNpuHpeArpMax(i["arp-max"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "l2_others_max"
	if _, ok := i["l2-others-max"]; ok {
		result["l2_others_max"] = flattenSystemNpuHpeL2OthersMax(i["l2-others-max"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "high_priority"
	if _, ok := i["high-priority"]; ok {
		result["high_priority"] = flattenSystemNpuHpeHighPriority(i["high-priority"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "enable_shaper"
	if _, ok := i["enable-shaper"]; ok {
		result["enable_shaper"] = flattenSystemNpuHpeEnableShaper(i["enable-shaper"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemNpuHpeAllProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuHpeTcpsynMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuHpeTcpsynAckMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuHpeTcpfinRstMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuHpeTcpMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuHpeUdpMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuHpeIcmpMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuHpeSctpMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuHpeEspMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuHpeIpFragMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuHpeIpOthersMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuHpeArpMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuHpeL2OthersMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuHpeHighPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuHpeEnableShaper(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFastpath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuCapwapOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuVxlanOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuVxlanMacFlappingGuard(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuDefaultQosType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuDefaultIpsecMcsType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuShapingStats(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuMcsHostPacketTpeShaping(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuGtpSupport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuPerSessionAccounting(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuSessionAcctInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuMaxSessionTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuFpAnomaly(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "tcp_syn_fin"
	if _, ok := i["tcp-syn-fin"]; ok {
		result["tcp_syn_fin"] = flattenSystemNpuFpAnomalyTcpSynFin(i["tcp-syn-fin"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_fin_noack"
	if _, ok := i["tcp-fin-noack"]; ok {
		result["tcp_fin_noack"] = flattenSystemNpuFpAnomalyTcpFinNoack(i["tcp-fin-noack"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_fin_only"
	if _, ok := i["tcp-fin-only"]; ok {
		result["tcp_fin_only"] = flattenSystemNpuFpAnomalyTcpFinOnly(i["tcp-fin-only"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_no_flag"
	if _, ok := i["tcp-no-flag"]; ok {
		result["tcp_no_flag"] = flattenSystemNpuFpAnomalyTcpNoFlag(i["tcp-no-flag"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_syn_data"
	if _, ok := i["tcp-syn-data"]; ok {
		result["tcp_syn_data"] = flattenSystemNpuFpAnomalyTcpSynData(i["tcp-syn-data"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_winnuke"
	if _, ok := i["tcp-winnuke"]; ok {
		result["tcp_winnuke"] = flattenSystemNpuFpAnomalyTcpWinnuke(i["tcp-winnuke"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_land"
	if _, ok := i["tcp-land"]; ok {
		result["tcp_land"] = flattenSystemNpuFpAnomalyTcpLand(i["tcp-land"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "udp_land"
	if _, ok := i["udp-land"]; ok {
		result["udp_land"] = flattenSystemNpuFpAnomalyUdpLand(i["udp-land"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "icmp_land"
	if _, ok := i["icmp-land"]; ok {
		result["icmp_land"] = flattenSystemNpuFpAnomalyIcmpLand(i["icmp-land"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "icmp_frag"
	if _, ok := i["icmp-frag"]; ok {
		result["icmp_frag"] = flattenSystemNpuFpAnomalyIcmpFrag(i["icmp-frag"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv4_land"
	if _, ok := i["ipv4-land"]; ok {
		result["ipv4_land"] = flattenSystemNpuFpAnomalyIpv4Land(i["ipv4-land"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv4_proto_err"
	if _, ok := i["ipv4-proto-err"]; ok {
		result["ipv4_proto_err"] = flattenSystemNpuFpAnomalyIpv4ProtoErr(i["ipv4-proto-err"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv4_unknopt"
	if _, ok := i["ipv4-unknopt"]; ok {
		result["ipv4_unknopt"] = flattenSystemNpuFpAnomalyIpv4Unknopt(i["ipv4-unknopt"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv4_optrr"
	if _, ok := i["ipv4-optrr"]; ok {
		result["ipv4_optrr"] = flattenSystemNpuFpAnomalyIpv4Optrr(i["ipv4-optrr"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv4_optssrr"
	if _, ok := i["ipv4-optssrr"]; ok {
		result["ipv4_optssrr"] = flattenSystemNpuFpAnomalyIpv4Optssrr(i["ipv4-optssrr"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv4_optlsrr"
	if _, ok := i["ipv4-optlsrr"]; ok {
		result["ipv4_optlsrr"] = flattenSystemNpuFpAnomalyIpv4Optlsrr(i["ipv4-optlsrr"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv4_optstream"
	if _, ok := i["ipv4-optstream"]; ok {
		result["ipv4_optstream"] = flattenSystemNpuFpAnomalyIpv4Optstream(i["ipv4-optstream"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv4_optsecurity"
	if _, ok := i["ipv4-optsecurity"]; ok {
		result["ipv4_optsecurity"] = flattenSystemNpuFpAnomalyIpv4Optsecurity(i["ipv4-optsecurity"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv4_opttimestamp"
	if _, ok := i["ipv4-opttimestamp"]; ok {
		result["ipv4_opttimestamp"] = flattenSystemNpuFpAnomalyIpv4Opttimestamp(i["ipv4-opttimestamp"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv4_csum_err"
	if _, ok := i["ipv4-csum-err"]; ok {
		result["ipv4_csum_err"] = flattenSystemNpuFpAnomalyIpv4CsumErr(i["ipv4-csum-err"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_csum_err"
	if _, ok := i["tcp-csum-err"]; ok {
		result["tcp_csum_err"] = flattenSystemNpuFpAnomalyTcpCsumErr(i["tcp-csum-err"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "udp_csum_err"
	if _, ok := i["udp-csum-err"]; ok {
		result["udp_csum_err"] = flattenSystemNpuFpAnomalyUdpCsumErr(i["udp-csum-err"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "udplite_csum_err"
	if _, ok := i["udplite-csum-err"]; ok {
		result["udplite_csum_err"] = flattenSystemNpuFpAnomalyUdpliteCsumErr(i["udplite-csum-err"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "icmp_csum_err"
	if _, ok := i["icmp-csum-err"]; ok {
		result["icmp_csum_err"] = flattenSystemNpuFpAnomalyIcmpCsumErr(i["icmp-csum-err"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gre_csum_err"
	if _, ok := i["gre-csum-err"]; ok {
		result["gre_csum_err"] = flattenSystemNpuFpAnomalyGreCsumErr(i["gre-csum-err"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sctp_csum_err"
	if _, ok := i["sctp-csum-err"]; ok {
		result["sctp_csum_err"] = flattenSystemNpuFpAnomalySctpCsumErr(i["sctp-csum-err"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv6_land"
	if _, ok := i["ipv6-land"]; ok {
		result["ipv6_land"] = flattenSystemNpuFpAnomalyIpv6Land(i["ipv6-land"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv6_proto_err"
	if _, ok := i["ipv6-proto-err"]; ok {
		result["ipv6_proto_err"] = flattenSystemNpuFpAnomalyIpv6ProtoErr(i["ipv6-proto-err"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv6_unknopt"
	if _, ok := i["ipv6-unknopt"]; ok {
		result["ipv6_unknopt"] = flattenSystemNpuFpAnomalyIpv6Unknopt(i["ipv6-unknopt"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv6_saddr_err"
	if _, ok := i["ipv6-saddr-err"]; ok {
		result["ipv6_saddr_err"] = flattenSystemNpuFpAnomalyIpv6SaddrErr(i["ipv6-saddr-err"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv6_daddr_err"
	if _, ok := i["ipv6-daddr-err"]; ok {
		result["ipv6_daddr_err"] = flattenSystemNpuFpAnomalyIpv6DaddrErr(i["ipv6-daddr-err"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv6_optralert"
	if _, ok := i["ipv6-optralert"]; ok {
		result["ipv6_optralert"] = flattenSystemNpuFpAnomalyIpv6Optralert(i["ipv6-optralert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv6_optjumbo"
	if _, ok := i["ipv6-optjumbo"]; ok {
		result["ipv6_optjumbo"] = flattenSystemNpuFpAnomalyIpv6Optjumbo(i["ipv6-optjumbo"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv6_opttunnel"
	if _, ok := i["ipv6-opttunnel"]; ok {
		result["ipv6_opttunnel"] = flattenSystemNpuFpAnomalyIpv6Opttunnel(i["ipv6-opttunnel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv6_opthomeaddr"
	if _, ok := i["ipv6-opthomeaddr"]; ok {
		result["ipv6_opthomeaddr"] = flattenSystemNpuFpAnomalyIpv6Opthomeaddr(i["ipv6-opthomeaddr"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv6_optnsap"
	if _, ok := i["ipv6-optnsap"]; ok {
		result["ipv6_optnsap"] = flattenSystemNpuFpAnomalyIpv6Optnsap(i["ipv6-optnsap"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv6_optendpid"
	if _, ok := i["ipv6-optendpid"]; ok {
		result["ipv6_optendpid"] = flattenSystemNpuFpAnomalyIpv6Optendpid(i["ipv6-optendpid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipv6_optinvld"
	if _, ok := i["ipv6-optinvld"]; ok {
		result["ipv6_optinvld"] = flattenSystemNpuFpAnomalyIpv6Optinvld(i["ipv6-optinvld"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemNpuFpAnomalyTcpSynFin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyTcpFinNoack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyTcpFinOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyTcpNoFlag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyTcpSynData(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyTcpWinnuke(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyTcpLand(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyUdpLand(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIcmpLand(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIcmpFrag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv4Land(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv4ProtoErr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv4Unknopt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv4Optrr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv4Optssrr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv4Optlsrr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv4Optstream(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv4Optsecurity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv4Opttimestamp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv4CsumErr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyTcpCsumErr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyUdpCsumErr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyUdpliteCsumErr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIcmpCsumErr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyGreCsumErr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalySctpCsumErr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv6Land(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv6ProtoErr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv6Unknopt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv6SaddrErr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv6DaddrErr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv6Optralert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv6Optjumbo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv6Opttunnel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv6Opthomeaddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv6Optnsap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv6Optendpid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFpAnomalyIpv6Optinvld(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuIpReassembly(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "min_timeout"
	if _, ok := i["min-timeout"]; ok {
		result["min_timeout"] = flattenSystemNpuIpReassemblyMinTimeout(i["min-timeout"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_timeout"
	if _, ok := i["max-timeout"]; ok {
		result["max_timeout"] = flattenSystemNpuIpReassemblyMaxTimeout(i["max-timeout"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemNpuIpReassemblyStatus(i["status"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemNpuIpReassemblyMinTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuIpReassemblyMaxTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuIpReassemblyStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuHashTblSpread(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuVlanLookupCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuIpFragmentOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuHtxIcmpCsumChk(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuHtabMsgQueue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuHtabDediQueueNr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuDswDtsProfile(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "profile-id", "profile_id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["profile-id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "profile_id"
			}
			tmp["profile_id"] = flattenSystemNpuDswDtsProfileProfileId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["min-limit"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "min_limit"
			}
			tmp["min_limit"] = flattenSystemNpuDswDtsProfileMinLimit(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["step"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "step"
			}
			tmp["step"] = flattenSystemNpuDswDtsProfileStep(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["action"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
			}
			tmp["action"] = flattenSystemNpuDswDtsProfileAction(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "profile_id", d)
	return result
}

func flattenSystemNpuDswDtsProfileProfileId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuDswDtsProfileMinLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuDswDtsProfileStep(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuDswDtsProfileAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuDswQueueDtsProfile(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemNpuDswQueueDtsProfileName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["iport"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "iport"
			}
			tmp["iport"] = flattenSystemNpuDswQueueDtsProfileIport(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["oport"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "oport"
			}
			tmp["oport"] = flattenSystemNpuDswQueueDtsProfileOport(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["profile-id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "profile_id"
			}
			tmp["profile_id"] = flattenSystemNpuDswQueueDtsProfileProfileId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["queue-select"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "queue_select"
			}
			tmp["queue_select"] = flattenSystemNpuDswQueueDtsProfileQueueSelect(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemNpuDswQueueDtsProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuDswQueueDtsProfileIport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuDswQueueDtsProfileOport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuDswQueueDtsProfileProfileId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuDswQueueDtsProfileQueueSelect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcam(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemNpuNpuTcamName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
			}
			tmp["type"] = flattenSystemNpuNpuTcamType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["oid"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "oid"
			}
			tmp["oid"] = flattenSystemNpuNpuTcamOid(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["vid"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vid"
			}
			tmp["vid"] = flattenSystemNpuNpuTcamVid(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["data"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "data"
			}
			tmp["data"] = flattenSystemNpuNpuTcamData(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["mask"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "mask"
			}
			tmp["mask"] = flattenSystemNpuNpuTcamMask(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["mir-act"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "mir_act"
			}
			tmp["mir_act"] = flattenSystemNpuNpuTcamMirAct(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["pri-act"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "pri_act"
			}
			tmp["pri_act"] = flattenSystemNpuNpuTcamPriAct(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["sact"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "sact"
			}
			tmp["sact"] = flattenSystemNpuNpuTcamSact(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["tact"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "tact"
			}
			tmp["tact"] = flattenSystemNpuNpuTcamTact(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemNpuNpuTcamName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamOid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamVid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamData(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "gen_buf_cnt"
	if _, ok := i["gen-buf-cnt"]; ok {
		result["gen_buf_cnt"] = flattenSystemNpuNpuTcamDataGenBufCnt(i["gen-buf-cnt"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gen_pri"
	if _, ok := i["gen-pri"]; ok {
		result["gen_pri"] = flattenSystemNpuNpuTcamDataGenPri(i["gen-pri"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gen_pri_v"
	if _, ok := i["gen-pri-v"]; ok {
		result["gen_pri_v"] = flattenSystemNpuNpuTcamDataGenPriV(i["gen-pri-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gen_iv"
	if _, ok := i["gen-iv"]; ok {
		result["gen_iv"] = flattenSystemNpuNpuTcamDataGenIv(i["gen-iv"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gen_tv"
	if _, ok := i["gen-tv"]; ok {
		result["gen_tv"] = flattenSystemNpuNpuTcamDataGenTv(i["gen-tv"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gen_pkt_ctrl"
	if _, ok := i["gen-pkt-ctrl"]; ok {
		result["gen_pkt_ctrl"] = flattenSystemNpuNpuTcamDataGenPktCtrl(i["gen-pkt-ctrl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gen_l3_flags"
	if _, ok := i["gen-l3-flags"]; ok {
		result["gen_l3_flags"] = flattenSystemNpuNpuTcamDataGenL3Flags(i["gen-l3-flags"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gen_l4_flags"
	if _, ok := i["gen-l4-flags"]; ok {
		result["gen_l4_flags"] = flattenSystemNpuNpuTcamDataGenL4Flags(i["gen-l4-flags"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vdid"
	if _, ok := i["vdid"]; ok {
		result["vdid"] = flattenSystemNpuNpuTcamDataVdid(i["vdid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tp"
	if _, ok := i["tp"]; ok {
		result["tp"] = flattenSystemNpuNpuTcamDataTp(i["tp"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tgt_updt"
	if _, ok := i["tgt-updt"]; ok {
		result["tgt_updt"] = flattenSystemNpuNpuTcamDataTgtUpdt(i["tgt-updt"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "smac_change"
	if _, ok := i["smac-change"]; ok {
		result["smac_change"] = flattenSystemNpuNpuTcamDataSmacChange(i["smac-change"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ext_tag"
	if _, ok := i["ext-tag"]; ok {
		result["ext_tag"] = flattenSystemNpuNpuTcamDataExtTag(i["ext-tag"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tgt_v"
	if _, ok := i["tgt-v"]; ok {
		result["tgt_v"] = flattenSystemNpuNpuTcamDataTgtV(i["tgt-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tvid"
	if _, ok := i["tvid"]; ok {
		result["tvid"] = flattenSystemNpuNpuTcamDataTvid(i["tvid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tgt_cfi"
	if _, ok := i["tgt-cfi"]; ok {
		result["tgt_cfi"] = flattenSystemNpuNpuTcamDataTgtCfi(i["tgt-cfi"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tgt_prio"
	if _, ok := i["tgt-prio"]; ok {
		result["tgt_prio"] = flattenSystemNpuNpuTcamDataTgtPrio(i["tgt-prio"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sp"
	if _, ok := i["sp"]; ok {
		result["sp"] = flattenSystemNpuNpuTcamDataSp(i["sp"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "src_updt"
	if _, ok := i["src-updt"]; ok {
		result["src_updt"] = flattenSystemNpuNpuTcamDataSrcUpdt(i["src-updt"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "slink"
	if _, ok := i["slink"]; ok {
		result["slink"] = flattenSystemNpuNpuTcamDataSlink(i["slink"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "svid"
	if _, ok := i["svid"]; ok {
		result["svid"] = flattenSystemNpuNpuTcamDataSvid(i["svid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "src_cfi"
	if _, ok := i["src-cfi"]; ok {
		result["src_cfi"] = flattenSystemNpuNpuTcamDataSrcCfi(i["src-cfi"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "src_prio"
	if _, ok := i["src-prio"]; ok {
		result["src_prio"] = flattenSystemNpuNpuTcamDataSrcPrio(i["src-prio"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "srcmac"
	if _, ok := i["srcmac"]; ok {
		result["srcmac"] = flattenSystemNpuNpuTcamDataSrcmac(i["srcmac"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dstmac"
	if _, ok := i["dstmac"]; ok {
		result["dstmac"] = flattenSystemNpuNpuTcamDataDstmac(i["dstmac"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ethertype"
	if _, ok := i["ethertype"]; ok {
		result["ethertype"] = flattenSystemNpuNpuTcamDataEthertype(i["ethertype"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipver"
	if _, ok := i["ipver"]; ok {
		result["ipver"] = flattenSystemNpuNpuTcamDataIpver(i["ipver"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ihl"
	if _, ok := i["ihl"]; ok {
		result["ihl"] = flattenSystemNpuNpuTcamDataIhl(i["ihl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ip4_id"
	if _, ok := i["ip4-id"]; ok {
		result["ip4_id"] = flattenSystemNpuNpuTcamDataIp4Id(i["ip4-id"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "srcip"
	if _, ok := i["srcip"]; ok {
		result["srcip"] = flattenSystemNpuNpuTcamDataSrcip(i["srcip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dstip"
	if _, ok := i["dstip"]; ok {
		result["dstip"] = flattenSystemNpuNpuTcamDataDstip(i["dstip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ip6_fl"
	if _, ok := i["ip6-fl"]; ok {
		result["ip6_fl"] = flattenSystemNpuNpuTcamDataIp6Fl(i["ip6-fl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "srcipv6"
	if _, ok := i["srcipv6"]; ok {
		result["srcipv6"] = flattenSystemNpuNpuTcamDataSrcipv6(i["srcipv6"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dstipv6"
	if _, ok := i["dstipv6"]; ok {
		result["dstipv6"] = flattenSystemNpuNpuTcamDataDstipv6(i["dstipv6"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ttl"
	if _, ok := i["ttl"]; ok {
		result["ttl"] = flattenSystemNpuNpuTcamDataTtl(i["ttl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "protocol"
	if _, ok := i["protocol"]; ok {
		result["protocol"] = flattenSystemNpuNpuTcamDataProtocol(i["protocol"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tos"
	if _, ok := i["tos"]; ok {
		result["tos"] = flattenSystemNpuNpuTcamDataTos(i["tos"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "frag_off"
	if _, ok := i["frag-off"]; ok {
		result["frag_off"] = flattenSystemNpuNpuTcamDataFragOff(i["frag-off"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mf"
	if _, ok := i["mf"]; ok {
		result["mf"] = flattenSystemNpuNpuTcamDataMf(i["mf"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "df"
	if _, ok := i["df"]; ok {
		result["df"] = flattenSystemNpuNpuTcamDataDf(i["df"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "srcport"
	if _, ok := i["srcport"]; ok {
		result["srcport"] = flattenSystemNpuNpuTcamDataSrcport(i["srcport"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dstport"
	if _, ok := i["dstport"]; ok {
		result["dstport"] = flattenSystemNpuNpuTcamDataDstport(i["dstport"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_fin"
	if _, ok := i["tcp-fin"]; ok {
		result["tcp_fin"] = flattenSystemNpuNpuTcamDataTcpFin(i["tcp-fin"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_syn"
	if _, ok := i["tcp-syn"]; ok {
		result["tcp_syn"] = flattenSystemNpuNpuTcamDataTcpSyn(i["tcp-syn"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_rst"
	if _, ok := i["tcp-rst"]; ok {
		result["tcp_rst"] = flattenSystemNpuNpuTcamDataTcpRst(i["tcp-rst"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_push"
	if _, ok := i["tcp-push"]; ok {
		result["tcp_push"] = flattenSystemNpuNpuTcamDataTcpPush(i["tcp-push"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_ack"
	if _, ok := i["tcp-ack"]; ok {
		result["tcp_ack"] = flattenSystemNpuNpuTcamDataTcpAck(i["tcp-ack"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_urg"
	if _, ok := i["tcp-urg"]; ok {
		result["tcp_urg"] = flattenSystemNpuNpuTcamDataTcpUrg(i["tcp-urg"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_ece"
	if _, ok := i["tcp-ece"]; ok {
		result["tcp_ece"] = flattenSystemNpuNpuTcamDataTcpEce(i["tcp-ece"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_cwr"
	if _, ok := i["tcp-cwr"]; ok {
		result["tcp_cwr"] = flattenSystemNpuNpuTcamDataTcpCwr(i["tcp-cwr"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "l4_wd8"
	if _, ok := i["l4-wd8"]; ok {
		result["l4_wd8"] = flattenSystemNpuNpuTcamDataL4Wd8(i["l4-wd8"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "l4_wd9"
	if _, ok := i["l4-wd9"]; ok {
		result["l4_wd9"] = flattenSystemNpuNpuTcamDataL4Wd9(i["l4-wd9"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "l4_wd10"
	if _, ok := i["l4-wd10"]; ok {
		result["l4_wd10"] = flattenSystemNpuNpuTcamDataL4Wd10(i["l4-wd10"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "l4_wd11"
	if _, ok := i["l4-wd11"]; ok {
		result["l4_wd11"] = flattenSystemNpuNpuTcamDataL4Wd11(i["l4-wd11"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemNpuNpuTcamDataGenBufCnt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataGenPri(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataGenPriV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataGenIv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataGenTv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataGenPktCtrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataGenL3Flags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataGenL4Flags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataVdid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataTp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataTgtUpdt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataSmacChange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataExtTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataTgtV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataTvid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataTgtCfi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataTgtPrio(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataSrcUpdt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataSlink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataSvid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataSrcCfi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataSrcPrio(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataSrcmac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataDstmac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataEthertype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataIpver(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataIhl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataIp4Id(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataSrcip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataDstip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataIp6Fl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataSrcipv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataDstipv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataTos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataFragOff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataMf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataDf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataSrcport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataDstport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataTcpFin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataTcpSyn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataTcpRst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataTcpPush(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataTcpAck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataTcpUrg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataTcpEce(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataTcpCwr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamDataL4Wd8(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataL4Wd9(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataL4Wd10(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamDataL4Wd11(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMask(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "gen_buf_cnt"
	if _, ok := i["gen-buf-cnt"]; ok {
		result["gen_buf_cnt"] = flattenSystemNpuNpuTcamMaskGenBufCnt(i["gen-buf-cnt"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gen_pri"
	if _, ok := i["gen-pri"]; ok {
		result["gen_pri"] = flattenSystemNpuNpuTcamMaskGenPri(i["gen-pri"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gen_pri_v"
	if _, ok := i["gen-pri-v"]; ok {
		result["gen_pri_v"] = flattenSystemNpuNpuTcamMaskGenPriV(i["gen-pri-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gen_iv"
	if _, ok := i["gen-iv"]; ok {
		result["gen_iv"] = flattenSystemNpuNpuTcamMaskGenIv(i["gen-iv"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gen_tv"
	if _, ok := i["gen-tv"]; ok {
		result["gen_tv"] = flattenSystemNpuNpuTcamMaskGenTv(i["gen-tv"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gen_pkt_ctrl"
	if _, ok := i["gen-pkt-ctrl"]; ok {
		result["gen_pkt_ctrl"] = flattenSystemNpuNpuTcamMaskGenPktCtrl(i["gen-pkt-ctrl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gen_l3_flags"
	if _, ok := i["gen-l3-flags"]; ok {
		result["gen_l3_flags"] = flattenSystemNpuNpuTcamMaskGenL3Flags(i["gen-l3-flags"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gen_l4_flags"
	if _, ok := i["gen-l4-flags"]; ok {
		result["gen_l4_flags"] = flattenSystemNpuNpuTcamMaskGenL4Flags(i["gen-l4-flags"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vdid"
	if _, ok := i["vdid"]; ok {
		result["vdid"] = flattenSystemNpuNpuTcamMaskVdid(i["vdid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tp"
	if _, ok := i["tp"]; ok {
		result["tp"] = flattenSystemNpuNpuTcamMaskTp(i["tp"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tgt_updt"
	if _, ok := i["tgt-updt"]; ok {
		result["tgt_updt"] = flattenSystemNpuNpuTcamMaskTgtUpdt(i["tgt-updt"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "smac_change"
	if _, ok := i["smac-change"]; ok {
		result["smac_change"] = flattenSystemNpuNpuTcamMaskSmacChange(i["smac-change"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ext_tag"
	if _, ok := i["ext-tag"]; ok {
		result["ext_tag"] = flattenSystemNpuNpuTcamMaskExtTag(i["ext-tag"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tgt_v"
	if _, ok := i["tgt-v"]; ok {
		result["tgt_v"] = flattenSystemNpuNpuTcamMaskTgtV(i["tgt-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tvid"
	if _, ok := i["tvid"]; ok {
		result["tvid"] = flattenSystemNpuNpuTcamMaskTvid(i["tvid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tgt_cfi"
	if _, ok := i["tgt-cfi"]; ok {
		result["tgt_cfi"] = flattenSystemNpuNpuTcamMaskTgtCfi(i["tgt-cfi"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tgt_prio"
	if _, ok := i["tgt-prio"]; ok {
		result["tgt_prio"] = flattenSystemNpuNpuTcamMaskTgtPrio(i["tgt-prio"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sp"
	if _, ok := i["sp"]; ok {
		result["sp"] = flattenSystemNpuNpuTcamMaskSp(i["sp"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "src_updt"
	if _, ok := i["src-updt"]; ok {
		result["src_updt"] = flattenSystemNpuNpuTcamMaskSrcUpdt(i["src-updt"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "slink"
	if _, ok := i["slink"]; ok {
		result["slink"] = flattenSystemNpuNpuTcamMaskSlink(i["slink"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "svid"
	if _, ok := i["svid"]; ok {
		result["svid"] = flattenSystemNpuNpuTcamMaskSvid(i["svid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "src_cfi"
	if _, ok := i["src-cfi"]; ok {
		result["src_cfi"] = flattenSystemNpuNpuTcamMaskSrcCfi(i["src-cfi"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "src_prio"
	if _, ok := i["src-prio"]; ok {
		result["src_prio"] = flattenSystemNpuNpuTcamMaskSrcPrio(i["src-prio"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "srcmac"
	if _, ok := i["srcmac"]; ok {
		result["srcmac"] = flattenSystemNpuNpuTcamMaskSrcmac(i["srcmac"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dstmac"
	if _, ok := i["dstmac"]; ok {
		result["dstmac"] = flattenSystemNpuNpuTcamMaskDstmac(i["dstmac"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ethertype"
	if _, ok := i["ethertype"]; ok {
		result["ethertype"] = flattenSystemNpuNpuTcamMaskEthertype(i["ethertype"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipver"
	if _, ok := i["ipver"]; ok {
		result["ipver"] = flattenSystemNpuNpuTcamMaskIpver(i["ipver"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ihl"
	if _, ok := i["ihl"]; ok {
		result["ihl"] = flattenSystemNpuNpuTcamMaskIhl(i["ihl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ip4_id"
	if _, ok := i["ip4-id"]; ok {
		result["ip4_id"] = flattenSystemNpuNpuTcamMaskIp4Id(i["ip4-id"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "srcip"
	if _, ok := i["srcip"]; ok {
		result["srcip"] = flattenSystemNpuNpuTcamMaskSrcip(i["srcip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dstip"
	if _, ok := i["dstip"]; ok {
		result["dstip"] = flattenSystemNpuNpuTcamMaskDstip(i["dstip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ip6_fl"
	if _, ok := i["ip6-fl"]; ok {
		result["ip6_fl"] = flattenSystemNpuNpuTcamMaskIp6Fl(i["ip6-fl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "srcipv6"
	if _, ok := i["srcipv6"]; ok {
		result["srcipv6"] = flattenSystemNpuNpuTcamMaskSrcipv6(i["srcipv6"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dstipv6"
	if _, ok := i["dstipv6"]; ok {
		result["dstipv6"] = flattenSystemNpuNpuTcamMaskDstipv6(i["dstipv6"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ttl"
	if _, ok := i["ttl"]; ok {
		result["ttl"] = flattenSystemNpuNpuTcamMaskTtl(i["ttl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "protocol"
	if _, ok := i["protocol"]; ok {
		result["protocol"] = flattenSystemNpuNpuTcamMaskProtocol(i["protocol"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tos"
	if _, ok := i["tos"]; ok {
		result["tos"] = flattenSystemNpuNpuTcamMaskTos(i["tos"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "frag_off"
	if _, ok := i["frag-off"]; ok {
		result["frag_off"] = flattenSystemNpuNpuTcamMaskFragOff(i["frag-off"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mf"
	if _, ok := i["mf"]; ok {
		result["mf"] = flattenSystemNpuNpuTcamMaskMf(i["mf"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "df"
	if _, ok := i["df"]; ok {
		result["df"] = flattenSystemNpuNpuTcamMaskDf(i["df"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "srcport"
	if _, ok := i["srcport"]; ok {
		result["srcport"] = flattenSystemNpuNpuTcamMaskSrcport(i["srcport"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dstport"
	if _, ok := i["dstport"]; ok {
		result["dstport"] = flattenSystemNpuNpuTcamMaskDstport(i["dstport"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_fin"
	if _, ok := i["tcp-fin"]; ok {
		result["tcp_fin"] = flattenSystemNpuNpuTcamMaskTcpFin(i["tcp-fin"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_syn"
	if _, ok := i["tcp-syn"]; ok {
		result["tcp_syn"] = flattenSystemNpuNpuTcamMaskTcpSyn(i["tcp-syn"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_rst"
	if _, ok := i["tcp-rst"]; ok {
		result["tcp_rst"] = flattenSystemNpuNpuTcamMaskTcpRst(i["tcp-rst"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_push"
	if _, ok := i["tcp-push"]; ok {
		result["tcp_push"] = flattenSystemNpuNpuTcamMaskTcpPush(i["tcp-push"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_ack"
	if _, ok := i["tcp-ack"]; ok {
		result["tcp_ack"] = flattenSystemNpuNpuTcamMaskTcpAck(i["tcp-ack"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_urg"
	if _, ok := i["tcp-urg"]; ok {
		result["tcp_urg"] = flattenSystemNpuNpuTcamMaskTcpUrg(i["tcp-urg"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_ece"
	if _, ok := i["tcp-ece"]; ok {
		result["tcp_ece"] = flattenSystemNpuNpuTcamMaskTcpEce(i["tcp-ece"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tcp_cwr"
	if _, ok := i["tcp-cwr"]; ok {
		result["tcp_cwr"] = flattenSystemNpuNpuTcamMaskTcpCwr(i["tcp-cwr"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "l4_wd8"
	if _, ok := i["l4-wd8"]; ok {
		result["l4_wd8"] = flattenSystemNpuNpuTcamMaskL4Wd8(i["l4-wd8"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "l4_wd9"
	if _, ok := i["l4-wd9"]; ok {
		result["l4_wd9"] = flattenSystemNpuNpuTcamMaskL4Wd9(i["l4-wd9"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "l4_wd10"
	if _, ok := i["l4-wd10"]; ok {
		result["l4_wd10"] = flattenSystemNpuNpuTcamMaskL4Wd10(i["l4-wd10"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "l4_wd11"
	if _, ok := i["l4-wd11"]; ok {
		result["l4_wd11"] = flattenSystemNpuNpuTcamMaskL4Wd11(i["l4-wd11"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemNpuNpuTcamMaskGenBufCnt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskGenPri(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskGenPriV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskGenIv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskGenTv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskGenPktCtrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskGenL3Flags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskGenL4Flags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskVdid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskTp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskTgtUpdt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskSmacChange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskExtTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskTgtV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskTvid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskTgtCfi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskTgtPrio(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskSrcUpdt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskSlink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskSvid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskSrcCfi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskSrcPrio(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskSrcmac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskDstmac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskEthertype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskIpver(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskIhl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskIp4Id(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskSrcip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskDstip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskIp6Fl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskSrcipv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskDstipv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskTos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskFragOff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskMf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskDf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskSrcport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskDstport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskTcpFin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskTcpSyn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskTcpRst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskTcpPush(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskTcpAck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskTcpUrg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskTcpEce(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskTcpCwr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamMaskL4Wd8(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskL4Wd9(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskL4Wd10(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMaskL4Wd11(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamMirAct(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "vlif"
	if _, ok := i["vlif"]; ok {
		result["vlif"] = flattenSystemNpuNpuTcamMirActVlif(i["vlif"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemNpuNpuTcamMirActVlif(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamPriAct(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemNpuNpuTcamPriActPriority(i["priority"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "weight"
	if _, ok := i["weight"]; ok {
		result["weight"] = flattenSystemNpuNpuTcamPriActWeight(i["weight"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemNpuNpuTcamPriActPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamPriActWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSact(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "fwd_lif_v"
	if _, ok := i["fwd-lif-v"]; ok {
		result["fwd_lif_v"] = flattenSystemNpuNpuTcamSactFwdLifV(i["fwd-lif-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fwd_lif"
	if _, ok := i["fwd-lif"]; ok {
		result["fwd_lif"] = flattenSystemNpuNpuTcamSactFwdLif(i["fwd-lif"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fwd_tvid_v"
	if _, ok := i["fwd-tvid-v"]; ok {
		result["fwd_tvid_v"] = flattenSystemNpuNpuTcamSactFwdTvidV(i["fwd-tvid-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fwd_tvid"
	if _, ok := i["fwd-tvid"]; ok {
		result["fwd_tvid"] = flattenSystemNpuNpuTcamSactFwdTvid(i["fwd-tvid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "df_lif_v"
	if _, ok := i["df-lif-v"]; ok {
		result["df_lif_v"] = flattenSystemNpuNpuTcamSactDfLifV(i["df-lif-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "df_lif"
	if _, ok := i["df-lif"]; ok {
		result["df_lif"] = flattenSystemNpuNpuTcamSactDfLif(i["df-lif"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "act_v"
	if _, ok := i["act-v"]; ok {
		result["act_v"] = flattenSystemNpuNpuTcamSactActV(i["act-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "act"
	if _, ok := i["act"]; ok {
		result["act"] = flattenSystemNpuNpuTcamSactAct(i["act"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "pleen_v"
	if _, ok := i["pleen-v"]; ok {
		result["pleen_v"] = flattenSystemNpuNpuTcamSactPleenV(i["pleen-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "pleen"
	if _, ok := i["pleen"]; ok {
		result["pleen"] = flattenSystemNpuNpuTcamSactPleen(i["pleen"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "icpen_v"
	if _, ok := i["icpen-v"]; ok {
		result["icpen_v"] = flattenSystemNpuNpuTcamSactIcpenV(i["icpen-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "icpen"
	if _, ok := i["icpen"]; ok {
		result["icpen"] = flattenSystemNpuNpuTcamSactIcpen(i["icpen"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vdm_v"
	if _, ok := i["vdm-v"]; ok {
		result["vdm_v"] = flattenSystemNpuNpuTcamSactVdmV(i["vdm-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vdm"
	if _, ok := i["vdm"]; ok {
		result["vdm"] = flattenSystemNpuNpuTcamSactVdm(i["vdm"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "learn_v"
	if _, ok := i["learn-v"]; ok {
		result["learn_v"] = flattenSystemNpuNpuTcamSactLearnV(i["learn-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "learn"
	if _, ok := i["learn"]; ok {
		result["learn"] = flattenSystemNpuNpuTcamSactLearn(i["learn"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "rfsh_v"
	if _, ok := i["rfsh-v"]; ok {
		result["rfsh_v"] = flattenSystemNpuNpuTcamSactRfshV(i["rfsh-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "rfsh"
	if _, ok := i["rfsh"]; ok {
		result["rfsh"] = flattenSystemNpuNpuTcamSactRfsh(i["rfsh"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fwd_v"
	if _, ok := i["fwd-v"]; ok {
		result["fwd_v"] = flattenSystemNpuNpuTcamSactFwdV(i["fwd-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fwd"
	if _, ok := i["fwd"]; ok {
		result["fwd"] = flattenSystemNpuNpuTcamSactFwd(i["fwd"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "x_mode_v"
	if _, ok := i["x-mode-v"]; ok {
		result["x_mode_v"] = flattenSystemNpuNpuTcamSactXModeV(i["x-mode-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "x_mode"
	if _, ok := i["x-mode"]; ok {
		result["x_mode"] = flattenSystemNpuNpuTcamSactXMode(i["x-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "promis_v"
	if _, ok := i["promis-v"]; ok {
		result["promis_v"] = flattenSystemNpuNpuTcamSactPromisV(i["promis-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "promis"
	if _, ok := i["promis"]; ok {
		result["promis"] = flattenSystemNpuNpuTcamSactPromis(i["promis"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bmproc_v"
	if _, ok := i["bmproc-v"]; ok {
		result["bmproc_v"] = flattenSystemNpuNpuTcamSactBmprocV(i["bmproc-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bmproc"
	if _, ok := i["bmproc"]; ok {
		result["bmproc"] = flattenSystemNpuNpuTcamSactBmproc(i["bmproc"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mac_id_v"
	if _, ok := i["mac-id-v"]; ok {
		result["mac_id_v"] = flattenSystemNpuNpuTcamSactMacIdV(i["mac-id-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mac_id"
	if _, ok := i["mac-id"]; ok {
		result["mac_id"] = flattenSystemNpuNpuTcamSactMacId(i["mac-id"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dosen_v"
	if _, ok := i["dosen-v"]; ok {
		result["dosen_v"] = flattenSystemNpuNpuTcamSactDosenV(i["dosen-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dosen"
	if _, ok := i["dosen"]; ok {
		result["dosen"] = flattenSystemNpuNpuTcamSactDosen(i["dosen"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dfr_v"
	if _, ok := i["dfr-v"]; ok {
		result["dfr_v"] = flattenSystemNpuNpuTcamSactDfrV(i["dfr-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dfr"
	if _, ok := i["dfr"]; ok {
		result["dfr"] = flattenSystemNpuNpuTcamSactDfr(i["dfr"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "m_srh_ctrl_v"
	if _, ok := i["m-srh-ctrl-v"]; ok {
		result["m_srh_ctrl_v"] = flattenSystemNpuNpuTcamSactMSrhCtrlV(i["m-srh-ctrl-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "m_srh_ctrl"
	if _, ok := i["m-srh-ctrl"]; ok {
		result["m_srh_ctrl"] = flattenSystemNpuNpuTcamSactMSrhCtrl(i["m-srh-ctrl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tpe_id_v"
	if _, ok := i["tpe-id-v"]; ok {
		result["tpe_id_v"] = flattenSystemNpuNpuTcamSactTpeIdV(i["tpe-id-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tpe_id"
	if _, ok := i["tpe-id"]; ok {
		result["tpe_id"] = flattenSystemNpuNpuTcamSactTpeId(i["tpe-id"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vdom_id_v"
	if _, ok := i["vdom-id-v"]; ok {
		result["vdom_id_v"] = flattenSystemNpuNpuTcamSactVdomIdV(i["vdom-id-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vdom_id"
	if _, ok := i["vdom-id"]; ok {
		result["vdom_id"] = flattenSystemNpuNpuTcamSactVdomId(i["vdom-id"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mss_v"
	if _, ok := i["mss-v"]; ok {
		result["mss_v"] = flattenSystemNpuNpuTcamSactMssV(i["mss-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mss"
	if _, ok := i["mss"]; ok {
		result["mss"] = flattenSystemNpuNpuTcamSactMss(i["mss"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tp_smchk_v"
	if _, ok := i["tp-smchk-v"]; ok {
		result["tp_smchk_v"] = flattenSystemNpuNpuTcamSactTpSmchkV(i["tp-smchk-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tp_smchk"
	if _, ok := i["tp_smchk"]; ok {
		result["tp_smchk"] = flattenSystemNpuNpuTcamSactTpSmchk(i["tp_smchk"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "etype_pid_v"
	if _, ok := i["etype-pid-v"]; ok {
		result["etype_pid_v"] = flattenSystemNpuNpuTcamSactEtypePidV(i["etype-pid-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "etype_pid"
	if _, ok := i["etype-pid"]; ok {
		result["etype_pid"] = flattenSystemNpuNpuTcamSactEtypePid(i["etype-pid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "frag_proc_v"
	if _, ok := i["frag-proc-v"]; ok {
		result["frag_proc_v"] = flattenSystemNpuNpuTcamSactFragProcV(i["frag-proc-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "frag_proc"
	if _, ok := i["frag-proc"]; ok {
		result["frag_proc"] = flattenSystemNpuNpuTcamSactFragProc(i["frag-proc"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "espff_proc_v"
	if _, ok := i["espff-proc-v"]; ok {
		result["espff_proc_v"] = flattenSystemNpuNpuTcamSactEspffProcV(i["espff-proc-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "espff_proc"
	if _, ok := i["espff-proc"]; ok {
		result["espff_proc"] = flattenSystemNpuNpuTcamSactEspffProc(i["espff-proc"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "prio_pid_v"
	if _, ok := i["prio-pid-v"]; ok {
		result["prio_pid_v"] = flattenSystemNpuNpuTcamSactPrioPidV(i["prio-pid-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "prio_pid"
	if _, ok := i["prio-pid"]; ok {
		result["prio_pid"] = flattenSystemNpuNpuTcamSactPrioPid(i["prio-pid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "igmp_mld_snp_v"
	if _, ok := i["igmp-mld-snp-v"]; ok {
		result["igmp_mld_snp_v"] = flattenSystemNpuNpuTcamSactIgmpMldSnpV(i["igmp-mld-snp-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "igmp_mld_snp"
	if _, ok := i["igmp-mld-snp"]; ok {
		result["igmp_mld_snp"] = flattenSystemNpuNpuTcamSactIgmpMldSnp(i["igmp-mld-snp"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "smac_skip_v"
	if _, ok := i["smac-skip-v"]; ok {
		result["smac_skip_v"] = flattenSystemNpuNpuTcamSactSmacSkipV(i["smac-skip-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "smac_skip"
	if _, ok := i["smac-skip"]; ok {
		result["smac_skip"] = flattenSystemNpuNpuTcamSactSmacSkip(i["smac-skip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dmac_skip_v"
	if _, ok := i["dmac-skip-v"]; ok {
		result["dmac_skip_v"] = flattenSystemNpuNpuTcamSactDmacSkipV(i["dmac-skip-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dmac_skip"
	if _, ok := i["dmac-skip"]; ok {
		result["dmac_skip"] = flattenSystemNpuNpuTcamSactDmacSkip(i["dmac-skip"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemNpuNpuTcamSactFwdLifV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactFwdLif(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactFwdTvidV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactFwdTvid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactDfLifV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactDfLif(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactActV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactAct(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactPleenV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactPleen(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactIcpenV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactIcpen(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactVdmV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactVdm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactLearnV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactLearn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactRfshV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactRfsh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactFwdV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactFwd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactXModeV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactXMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactPromisV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactPromis(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactBmprocV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactBmproc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactMacIdV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactMacId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactDosenV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactDosen(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactDfrV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactDfr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactMSrhCtrlV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactMSrhCtrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactTpeIdV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactTpeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactVdomIdV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactVdomId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactMssV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactMss(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactTpSmchkV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactTpSmchk(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactEtypePidV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactEtypePid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactFragProcV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactFragProc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactEspffProcV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactEspffProc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactPrioPidV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactPrioPid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactIgmpMldSnpV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactIgmpMldSnp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactSmacSkipV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactSmacSkip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamSactDmacSkipV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamSactDmacSkip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamTact(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "act_v"
	if _, ok := i["act-v"]; ok {
		result["act_v"] = flattenSystemNpuNpuTcamTactActV(i["act-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "act"
	if _, ok := i["act"]; ok {
		result["act"] = flattenSystemNpuNpuTcamTactAct(i["act"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mtuv4_v"
	if _, ok := i["mtuv4-v"]; ok {
		result["mtuv4_v"] = flattenSystemNpuNpuTcamTactMtuv4V(i["mtuv4-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mtuv4"
	if _, ok := i["mtuv4"]; ok {
		result["mtuv4"] = flattenSystemNpuNpuTcamTactMtuv4(i["mtuv4"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mtuv6_v"
	if _, ok := i["mtuv6-v"]; ok {
		result["mtuv6_v"] = flattenSystemNpuNpuTcamTactMtuv6V(i["mtuv6-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mtuv6"
	if _, ok := i["mtuv6"]; ok {
		result["mtuv6"] = flattenSystemNpuNpuTcamTactMtuv6(i["mtuv6"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mac_id_v"
	if _, ok := i["mac-id-v"]; ok {
		result["mac_id_v"] = flattenSystemNpuNpuTcamTactMacIdV(i["mac-id-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mac_id"
	if _, ok := i["mac-id"]; ok {
		result["mac_id"] = flattenSystemNpuNpuTcamTactMacId(i["mac-id"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "slif_act_v"
	if _, ok := i["slif-act-v"]; ok {
		result["slif_act_v"] = flattenSystemNpuNpuTcamTactSlifActV(i["slif-act-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "slif_act"
	if _, ok := i["slif-act"]; ok {
		result["slif_act"] = flattenSystemNpuNpuTcamTactSlifAct(i["slif-act"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tlif_act_v"
	if _, ok := i["tlif-act-v"]; ok {
		result["tlif_act_v"] = flattenSystemNpuNpuTcamTactTlifActV(i["tlif-act-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tlif_act"
	if _, ok := i["tlif-act"]; ok {
		result["tlif_act"] = flattenSystemNpuNpuTcamTactTlifAct(i["tlif-act"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tgtv_act_v"
	if _, ok := i["tgtv-act-v"]; ok {
		result["tgtv_act_v"] = flattenSystemNpuNpuTcamTactTgtvActV(i["tgtv-act-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tgtv_act"
	if _, ok := i["tgtv-act"]; ok {
		result["tgtv_act"] = flattenSystemNpuNpuTcamTactTgtvAct(i["tgtv-act"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tpeid_v"
	if _, ok := i["tpeid-v"]; ok {
		result["tpeid_v"] = flattenSystemNpuNpuTcamTactTpeidV(i["tpeid-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tpeid"
	if _, ok := i["tpeid"]; ok {
		result["tpeid"] = flattenSystemNpuNpuTcamTactTpeid(i["tpeid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "v6fe_v"
	if _, ok := i["v6fe-v"]; ok {
		result["v6fe_v"] = flattenSystemNpuNpuTcamTactV6FeV(i["v6fe-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "v6fe"
	if _, ok := i["v6fe"]; ok {
		result["v6fe"] = flattenSystemNpuNpuTcamTactV6Fe(i["v6fe"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "xlt_vid_v"
	if _, ok := i["xlt-vid-v"]; ok {
		result["xlt_vid_v"] = flattenSystemNpuNpuTcamTactXltVidV(i["xlt-vid-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "xlt_vid"
	if _, ok := i["xlt-vid"]; ok {
		result["xlt_vid"] = flattenSystemNpuNpuTcamTactXltVid(i["xlt-vid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "xlt_lif_v"
	if _, ok := i["xlt-lif-v"]; ok {
		result["xlt_lif_v"] = flattenSystemNpuNpuTcamTactXltLifV(i["xlt-lif-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "xlt_lif"
	if _, ok := i["xlt-lif"]; ok {
		result["xlt_lif"] = flattenSystemNpuNpuTcamTactXltLif(i["xlt-lif"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mss_t_v"
	if _, ok := i["mss-t-v"]; ok {
		result["mss_t_v"] = flattenSystemNpuNpuTcamTactMssTV(i["mss-t-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mss_t"
	if _, ok := i["mss-t"]; ok {
		result["mss_t"] = flattenSystemNpuNpuTcamTactMssT(i["mss-t"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "lnkid_v"
	if _, ok := i["lnkid-v"]; ok {
		result["lnkid_v"] = flattenSystemNpuNpuTcamTactLnkidV(i["lnkid-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "lnkid"
	if _, ok := i["lnkid"]; ok {
		result["lnkid"] = flattenSystemNpuNpuTcamTactLnkid(i["lnkid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sublnkid_v"
	if _, ok := i["sublnkid-v"]; ok {
		result["sublnkid_v"] = flattenSystemNpuNpuTcamTactSublnkidV(i["sublnkid-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sublnkid"
	if _, ok := i["sublnkid"]; ok {
		result["sublnkid"] = flattenSystemNpuNpuTcamTactSublnkid(i["sublnkid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fmtuv4_s_v"
	if _, ok := i["fmtuv4-s-v"]; ok {
		result["fmtuv4_s_v"] = flattenSystemNpuNpuTcamTactFmtuv4SV(i["fmtuv4-s-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fmtuv4_s"
	if _, ok := i["fmtuv4-s"]; ok {
		result["fmtuv4_s"] = flattenSystemNpuNpuTcamTactFmtuv4S(i["fmtuv4-s"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fmtuv6_s_v"
	if _, ok := i["fmtuv6-s-v"]; ok {
		result["fmtuv6_s_v"] = flattenSystemNpuNpuTcamTactFmtuv6SV(i["fmtuv6-s-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fmtuv6_s"
	if _, ok := i["fmtuv6-s"]; ok {
		result["fmtuv6_s"] = flattenSystemNpuNpuTcamTactFmtuv6S(i["fmtuv6-s"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vep_en_v"
	if _, ok := i["vep-en-v"]; ok {
		result["vep_en_v"] = flattenSystemNpuNpuTcamTactVepEnV(i["vep-en-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vep_en"
	if _, ok := i["vep_en"]; ok {
		result["vep_en"] = flattenSystemNpuNpuTcamTactVepEn(i["vep_en"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vep_slid_v"
	if _, ok := i["vep-slid-v"]; ok {
		result["vep_slid_v"] = flattenSystemNpuNpuTcamTactVepSlidV(i["vep-slid-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vep_slid"
	if _, ok := i["vep-slid"]; ok {
		result["vep_slid"] = flattenSystemNpuNpuTcamTactVepSlid(i["vep-slid"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemNpuNpuTcamTactActV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamTactAct(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamTactMtuv4V(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamTactMtuv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamTactMtuv6V(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamTactMtuv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamTactMacIdV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamTactMacId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamTactSlifActV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamTactSlifAct(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamTactTlifActV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamTactTlifAct(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamTactTgtvActV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamTactTgtvAct(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamTactTpeidV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamTactTpeid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamTactV6FeV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamTactV6Fe(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamTactXltVidV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamTactXltVid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamTactXltLifV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamTactXltLif(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamTactMssTV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamTactMssT(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamTactLnkidV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamTactLnkid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamTactSublnkidV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamTactSublnkid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamTactFmtuv4SV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamTactFmtuv4S(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamTactFmtuv6SV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamTactFmtuv6S(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamTactVepEnV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamTactVepEn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpuTcamTactVepSlidV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpuTcamTactVepSlid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpQueues(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "profile"
	if _, ok := i["profile"]; ok {
		result["profile"] = flattenSystemNpuNpQueuesProfile(i["profile"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ethernet_type"
	if _, ok := i["ethernet-type"]; ok {
		result["ethernet_type"] = flattenSystemNpuNpQueuesEthernetType(i["ethernet-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ip_protocol"
	if _, ok := i["ip-protocol"]; ok {
		result["ip_protocol"] = flattenSystemNpuNpQueuesIpProtocol(i["ip-protocol"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ip_service"
	if _, ok := i["ip-service"]; ok {
		result["ip_service"] = flattenSystemNpuNpQueuesIpService(i["ip-service"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "scheduler"
	if _, ok := i["scheduler"]; ok {
		result["scheduler"] = flattenSystemNpuNpQueuesScheduler(i["scheduler"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "custom_etype_lookup"
	if _, ok := i["custom-etype-lookup"]; ok {
		result["custom_etype_lookup"] = flattenSystemNpuNpQueuesCustomEtypeLookup(i["custom-etype-lookup"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemNpuNpQueuesProfile(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemNpuNpQueuesProfileId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
			}
			tmp["type"] = flattenSystemNpuNpQueuesProfileType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["weight"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
			}
			tmp["weight"] = flattenSystemNpuNpQueuesProfileWeight(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["cos0"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "cos0"
			}
			tmp["cos0"] = flattenSystemNpuNpQueuesProfileCos0(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["cos1"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "cos1"
			}
			tmp["cos1"] = flattenSystemNpuNpQueuesProfileCos1(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["cos2"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "cos2"
			}
			tmp["cos2"] = flattenSystemNpuNpQueuesProfileCos2(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["cos3"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "cos3"
			}
			tmp["cos3"] = flattenSystemNpuNpQueuesProfileCos3(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["cos4"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "cos4"
			}
			tmp["cos4"] = flattenSystemNpuNpQueuesProfileCos4(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["cos5"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "cos5"
			}
			tmp["cos5"] = flattenSystemNpuNpQueuesProfileCos5(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["cos6"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "cos6"
			}
			tmp["cos6"] = flattenSystemNpuNpQueuesProfileCos6(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["cos7"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "cos7"
			}
			tmp["cos7"] = flattenSystemNpuNpQueuesProfileCos7(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp0"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp0"
			}
			tmp["dscp0"] = flattenSystemNpuNpQueuesProfileDscp0(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp1"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp1"
			}
			tmp["dscp1"] = flattenSystemNpuNpQueuesProfileDscp1(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp2"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp2"
			}
			tmp["dscp2"] = flattenSystemNpuNpQueuesProfileDscp2(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp3"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp3"
			}
			tmp["dscp3"] = flattenSystemNpuNpQueuesProfileDscp3(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp4"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp4"
			}
			tmp["dscp4"] = flattenSystemNpuNpQueuesProfileDscp4(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp5"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp5"
			}
			tmp["dscp5"] = flattenSystemNpuNpQueuesProfileDscp5(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp6"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp6"
			}
			tmp["dscp6"] = flattenSystemNpuNpQueuesProfileDscp6(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp7"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp7"
			}
			tmp["dscp7"] = flattenSystemNpuNpQueuesProfileDscp7(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp8"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp8"
			}
			tmp["dscp8"] = flattenSystemNpuNpQueuesProfileDscp8(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp9"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp9"
			}
			tmp["dscp9"] = flattenSystemNpuNpQueuesProfileDscp9(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp10"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp10"
			}
			tmp["dscp10"] = flattenSystemNpuNpQueuesProfileDscp10(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp11"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp11"
			}
			tmp["dscp11"] = flattenSystemNpuNpQueuesProfileDscp11(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp12"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp12"
			}
			tmp["dscp12"] = flattenSystemNpuNpQueuesProfileDscp12(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp13"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp13"
			}
			tmp["dscp13"] = flattenSystemNpuNpQueuesProfileDscp13(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp14"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp14"
			}
			tmp["dscp14"] = flattenSystemNpuNpQueuesProfileDscp14(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp15"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp15"
			}
			tmp["dscp15"] = flattenSystemNpuNpQueuesProfileDscp15(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp16"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp16"
			}
			tmp["dscp16"] = flattenSystemNpuNpQueuesProfileDscp16(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp17"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp17"
			}
			tmp["dscp17"] = flattenSystemNpuNpQueuesProfileDscp17(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp18"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp18"
			}
			tmp["dscp18"] = flattenSystemNpuNpQueuesProfileDscp18(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp19"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp19"
			}
			tmp["dscp19"] = flattenSystemNpuNpQueuesProfileDscp19(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp20"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp20"
			}
			tmp["dscp20"] = flattenSystemNpuNpQueuesProfileDscp20(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp21"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp21"
			}
			tmp["dscp21"] = flattenSystemNpuNpQueuesProfileDscp21(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp22"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp22"
			}
			tmp["dscp22"] = flattenSystemNpuNpQueuesProfileDscp22(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp23"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp23"
			}
			tmp["dscp23"] = flattenSystemNpuNpQueuesProfileDscp23(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp24"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp24"
			}
			tmp["dscp24"] = flattenSystemNpuNpQueuesProfileDscp24(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp25"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp25"
			}
			tmp["dscp25"] = flattenSystemNpuNpQueuesProfileDscp25(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp26"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp26"
			}
			tmp["dscp26"] = flattenSystemNpuNpQueuesProfileDscp26(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp27"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp27"
			}
			tmp["dscp27"] = flattenSystemNpuNpQueuesProfileDscp27(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp28"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp28"
			}
			tmp["dscp28"] = flattenSystemNpuNpQueuesProfileDscp28(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp29"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp29"
			}
			tmp["dscp29"] = flattenSystemNpuNpQueuesProfileDscp29(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp30"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp30"
			}
			tmp["dscp30"] = flattenSystemNpuNpQueuesProfileDscp30(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp31"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp31"
			}
			tmp["dscp31"] = flattenSystemNpuNpQueuesProfileDscp31(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp32"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp32"
			}
			tmp["dscp32"] = flattenSystemNpuNpQueuesProfileDscp32(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp33"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp33"
			}
			tmp["dscp33"] = flattenSystemNpuNpQueuesProfileDscp33(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp34"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp34"
			}
			tmp["dscp34"] = flattenSystemNpuNpQueuesProfileDscp34(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp35"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp35"
			}
			tmp["dscp35"] = flattenSystemNpuNpQueuesProfileDscp35(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp36"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp36"
			}
			tmp["dscp36"] = flattenSystemNpuNpQueuesProfileDscp36(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp37"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp37"
			}
			tmp["dscp37"] = flattenSystemNpuNpQueuesProfileDscp37(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp38"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp38"
			}
			tmp["dscp38"] = flattenSystemNpuNpQueuesProfileDscp38(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp39"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp39"
			}
			tmp["dscp39"] = flattenSystemNpuNpQueuesProfileDscp39(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp40"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp40"
			}
			tmp["dscp40"] = flattenSystemNpuNpQueuesProfileDscp40(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp41"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp41"
			}
			tmp["dscp41"] = flattenSystemNpuNpQueuesProfileDscp41(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp42"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp42"
			}
			tmp["dscp42"] = flattenSystemNpuNpQueuesProfileDscp42(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp43"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp43"
			}
			tmp["dscp43"] = flattenSystemNpuNpQueuesProfileDscp43(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp44"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp44"
			}
			tmp["dscp44"] = flattenSystemNpuNpQueuesProfileDscp44(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp45"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp45"
			}
			tmp["dscp45"] = flattenSystemNpuNpQueuesProfileDscp45(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp46"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp46"
			}
			tmp["dscp46"] = flattenSystemNpuNpQueuesProfileDscp46(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp47"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp47"
			}
			tmp["dscp47"] = flattenSystemNpuNpQueuesProfileDscp47(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp48"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp48"
			}
			tmp["dscp48"] = flattenSystemNpuNpQueuesProfileDscp48(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp49"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp49"
			}
			tmp["dscp49"] = flattenSystemNpuNpQueuesProfileDscp49(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp50"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp50"
			}
			tmp["dscp50"] = flattenSystemNpuNpQueuesProfileDscp50(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp51"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp51"
			}
			tmp["dscp51"] = flattenSystemNpuNpQueuesProfileDscp51(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp52"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp52"
			}
			tmp["dscp52"] = flattenSystemNpuNpQueuesProfileDscp52(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp53"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp53"
			}
			tmp["dscp53"] = flattenSystemNpuNpQueuesProfileDscp53(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp54"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp54"
			}
			tmp["dscp54"] = flattenSystemNpuNpQueuesProfileDscp54(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp55"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp55"
			}
			tmp["dscp55"] = flattenSystemNpuNpQueuesProfileDscp55(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp56"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp56"
			}
			tmp["dscp56"] = flattenSystemNpuNpQueuesProfileDscp56(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp57"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp57"
			}
			tmp["dscp57"] = flattenSystemNpuNpQueuesProfileDscp57(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp58"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp58"
			}
			tmp["dscp58"] = flattenSystemNpuNpQueuesProfileDscp58(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp59"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp59"
			}
			tmp["dscp59"] = flattenSystemNpuNpQueuesProfileDscp59(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp60"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp60"
			}
			tmp["dscp60"] = flattenSystemNpuNpQueuesProfileDscp60(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp61"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp61"
			}
			tmp["dscp61"] = flattenSystemNpuNpQueuesProfileDscp61(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp62"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp62"
			}
			tmp["dscp62"] = flattenSystemNpuNpQueuesProfileDscp62(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp63"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp63"
			}
			tmp["dscp63"] = flattenSystemNpuNpQueuesProfileDscp63(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemNpuNpQueuesProfileId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpQueuesProfileType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpQueuesProfileCos0(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileCos1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileCos2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileCos3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileCos4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileCos5(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileCos6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileCos7(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp0(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp5(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp7(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp8(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp9(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp10(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp11(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp12(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp13(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp14(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp15(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp16(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp17(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp18(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp19(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp20(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp21(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp22(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp23(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp24(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp25(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp26(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp27(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp28(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp29(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp30(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp31(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp32(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp33(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp34(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp35(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp36(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp37(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp38(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp39(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp40(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp41(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp42(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp43(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp44(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp45(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp46(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp47(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp48(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp49(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp50(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp51(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp52(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp53(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp54(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp55(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp56(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp57(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp58(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp59(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp60(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp61(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp62(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp63(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesEthernetType(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemNpuNpQueuesEthernetTypeName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
			}
			tmp["type"] = flattenSystemNpuNpQueuesEthernetTypeType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["queue"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
			}
			tmp["queue"] = flattenSystemNpuNpQueuesEthernetTypeQueue(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["weight"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
			}
			tmp["weight"] = flattenSystemNpuNpQueuesEthernetTypeWeight(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemNpuNpQueuesEthernetTypeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesEthernetTypeType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesEthernetTypeQueue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpQueuesEthernetTypeWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpQueuesIpProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemNpuNpQueuesIpProtocolName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["protocol"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
			}
			tmp["protocol"] = flattenSystemNpuNpQueuesIpProtocolProtocol(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["queue"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
			}
			tmp["queue"] = flattenSystemNpuNpQueuesIpProtocolQueue(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["weight"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
			}
			tmp["weight"] = flattenSystemNpuNpQueuesIpProtocolWeight(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemNpuNpQueuesIpProtocolName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesIpProtocolProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpQueuesIpProtocolQueue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpQueuesIpProtocolWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpQueuesIpService(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemNpuNpQueuesIpServiceName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["protocol"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
			}
			tmp["protocol"] = flattenSystemNpuNpQueuesIpServiceProtocol(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["sport"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "sport"
			}
			tmp["sport"] = flattenSystemNpuNpQueuesIpServiceSport(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dport"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dport"
			}
			tmp["dport"] = flattenSystemNpuNpQueuesIpServiceDport(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["queue"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
			}
			tmp["queue"] = flattenSystemNpuNpQueuesIpServiceQueue(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["weight"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
			}
			tmp["weight"] = flattenSystemNpuNpQueuesIpServiceWeight(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemNpuNpQueuesIpServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesIpServiceProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpQueuesIpServiceSport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpQueuesIpServiceDport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpQueuesIpServiceQueue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpQueuesIpServiceWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuNpQueuesScheduler(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemNpuNpQueuesSchedulerName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["mode"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
			}
			tmp["mode"] = flattenSystemNpuNpQueuesSchedulerMode(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemNpuNpQueuesSchedulerName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesSchedulerMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesCustomEtypeLookup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuSwNpRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuSwNpRateUnit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuSwNpRateBurst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuDoubleLevelMcastOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuMcastDeniedSesOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuIpsecEncSubengineMask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuIpsecDecSubengineMask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNp6CpsOptimizationMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuSwNpBandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuStripEspPadding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuStripClearTextPadding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuIpsecInboundCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuSseBackpressure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuRdpOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuIpsecOverVlink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuUespOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuMcastSessionAccounting(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuIpsecMtuOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuSessionDeniedOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuQtmBufMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuIpsecObNpSel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuMaxReceiveUnit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuLagHashGre(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuUseMseOft(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuIkePort(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "port", "port")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["port"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
			}
			tmp["port"] = flattenSystemNpuIkePortPort(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "port", d)
	return result
}

func flattenSystemNpuIkePortPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNpuPriorityProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bgp"
	if _, ok := i["bgp"]; ok {
		result["bgp"] = flattenSystemNpuPriorityProtocolBgp(i["bgp"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "slbc"
	if _, ok := i["slbc"]; ok {
		result["slbc"] = flattenSystemNpuPriorityProtocolSlbc(i["slbc"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bfd"
	if _, ok := i["bfd"]; ok {
		result["bfd"] = flattenSystemNpuPriorityProtocolBfd(i["bfd"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemNpuPriorityProtocolBgp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuPriorityProtocolSlbc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuPriorityProtocolBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemNpu(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("dedicated_management_cpu", flattenSystemNpuDedicatedManagementCpu(o["dedicated-management-cpu"], d, "dedicated_management_cpu", sv)); err != nil {
		if !fortiAPIPatch(o["dedicated-management-cpu"]) {
			return fmt.Errorf("Error reading dedicated_management_cpu: %v", err)
		}
	}

	if err = d.Set("dedicated_lacp_queue", flattenSystemNpuDedicatedLacpQueue(o["dedicated-lacp-queue"], d, "dedicated_lacp_queue", sv)); err != nil {
		if !fortiAPIPatch(o["dedicated-lacp-queue"]) {
			return fmt.Errorf("Error reading dedicated_lacp_queue: %v", err)
		}
	}

	if err = d.Set("dedicated_management_affinity", flattenSystemNpuDedicatedManagementAffinity(o["dedicated-management-affinity"], d, "dedicated_management_affinity", sv)); err != nil {
		if !fortiAPIPatch(o["dedicated-management-affinity"]) {
			return fmt.Errorf("Error reading dedicated_management_affinity: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("dos_options", flattenSystemNpuDosOptions(o["dos-options"], d, "dos_options", sv)); err != nil {
			if !fortiAPIPatch(o["dos-options"]) {
				return fmt.Errorf("Error reading dos_options: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dos_options"); ok {
			if err = d.Set("dos_options", flattenSystemNpuDosOptions(o["dos-options"], d, "dos_options", sv)); err != nil {
				if !fortiAPIPatch(o["dos-options"]) {
					return fmt.Errorf("Error reading dos_options: %v", err)
				}
			}
		}
	}

	if err = d.Set("policy_offload_level", flattenSystemNpuPolicyOffloadLevel(o["policy-offload-level"], d, "policy_offload_level", sv)); err != nil {
		if !fortiAPIPatch(o["policy-offload-level"]) {
			return fmt.Errorf("Error reading policy_offload_level: %v", err)
		}
	}

	if err = d.Set("napi_break_interval", flattenSystemNpuNapiBreakInterval(o["napi-break-interval"], d, "napi_break_interval", sv)); err != nil {
		if !fortiAPIPatch(o["napi-break-interval"]) {
			return fmt.Errorf("Error reading napi_break_interval: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("hpe", flattenSystemNpuHpe(o["hpe"], d, "hpe", sv)); err != nil {
			if !fortiAPIPatch(o["hpe"]) {
				return fmt.Errorf("Error reading hpe: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("hpe"); ok {
			if err = d.Set("hpe", flattenSystemNpuHpe(o["hpe"], d, "hpe", sv)); err != nil {
				if !fortiAPIPatch(o["hpe"]) {
					return fmt.Errorf("Error reading hpe: %v", err)
				}
			}
		}
	}

	if err = d.Set("fastpath", flattenSystemNpuFastpath(o["fastpath"], d, "fastpath", sv)); err != nil {
		if !fortiAPIPatch(o["fastpath"]) {
			return fmt.Errorf("Error reading fastpath: %v", err)
		}
	}

	if err = d.Set("capwap_offload", flattenSystemNpuCapwapOffload(o["capwap-offload"], d, "capwap_offload", sv)); err != nil {
		if !fortiAPIPatch(o["capwap-offload"]) {
			return fmt.Errorf("Error reading capwap_offload: %v", err)
		}
	}

	if err = d.Set("vxlan_offload", flattenSystemNpuVxlanOffload(o["vxlan-offload"], d, "vxlan_offload", sv)); err != nil {
		if !fortiAPIPatch(o["vxlan-offload"]) {
			return fmt.Errorf("Error reading vxlan_offload: %v", err)
		}
	}

	if err = d.Set("vxlan_mac_flapping_guard", flattenSystemNpuVxlanMacFlappingGuard(o["vxlan-mac-flapping-guard"], d, "vxlan_mac_flapping_guard", sv)); err != nil {
		if !fortiAPIPatch(o["vxlan-mac-flapping-guard"]) {
			return fmt.Errorf("Error reading vxlan_mac_flapping_guard: %v", err)
		}
	}

	if err = d.Set("default_qos_type", flattenSystemNpuDefaultQosType(o["default-qos-type"], d, "default_qos_type", sv)); err != nil {
		if !fortiAPIPatch(o["default-qos-type"]) {
			return fmt.Errorf("Error reading default_qos_type: %v", err)
		}
	}

	if err = d.Set("default_ipsec_mcs_type", flattenSystemNpuDefaultIpsecMcsType(o["default-ipsec-mcs-type"], d, "default_ipsec_mcs_type", sv)); err != nil {
		if !fortiAPIPatch(o["default-ipsec-mcs-type"]) {
			return fmt.Errorf("Error reading default_ipsec_mcs_type: %v", err)
		}
	}

	if err = d.Set("shaping_stats", flattenSystemNpuShapingStats(o["shaping-stats"], d, "shaping_stats", sv)); err != nil {
		if !fortiAPIPatch(o["shaping-stats"]) {
			return fmt.Errorf("Error reading shaping_stats: %v", err)
		}
	}

	if err = d.Set("mcs_host_packet_tpe_shaping", flattenSystemNpuMcsHostPacketTpeShaping(o["mcs-host-packet-tpe-shaping"], d, "mcs_host_packet_tpe_shaping", sv)); err != nil {
		if !fortiAPIPatch(o["mcs-host-packet-tpe-shaping"]) {
			return fmt.Errorf("Error reading mcs_host_packet_tpe_shaping: %v", err)
		}
	}

	if err = d.Set("gtp_support", flattenSystemNpuGtpSupport(o["gtp-support"], d, "gtp_support", sv)); err != nil {
		if !fortiAPIPatch(o["gtp-support"]) {
			return fmt.Errorf("Error reading gtp_support: %v", err)
		}
	}

	if err = d.Set("per_session_accounting", flattenSystemNpuPerSessionAccounting(o["per-session-accounting"], d, "per_session_accounting", sv)); err != nil {
		if !fortiAPIPatch(o["per-session-accounting"]) {
			return fmt.Errorf("Error reading per_session_accounting: %v", err)
		}
	}

	if err = d.Set("session_acct_interval", flattenSystemNpuSessionAcctInterval(o["session-acct-interval"], d, "session_acct_interval", sv)); err != nil {
		if !fortiAPIPatch(o["session-acct-interval"]) {
			return fmt.Errorf("Error reading session_acct_interval: %v", err)
		}
	}

	if err = d.Set("max_session_timeout", flattenSystemNpuMaxSessionTimeout(o["max-session-timeout"], d, "max_session_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["max-session-timeout"]) {
			return fmt.Errorf("Error reading max_session_timeout: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("fp_anomaly", flattenSystemNpuFpAnomaly(o["fp-anomaly"], d, "fp_anomaly", sv)); err != nil {
			if !fortiAPIPatch(o["fp-anomaly"]) {
				return fmt.Errorf("Error reading fp_anomaly: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fp_anomaly"); ok {
			if err = d.Set("fp_anomaly", flattenSystemNpuFpAnomaly(o["fp-anomaly"], d, "fp_anomaly", sv)); err != nil {
				if !fortiAPIPatch(o["fp-anomaly"]) {
					return fmt.Errorf("Error reading fp_anomaly: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("ip_reassembly", flattenSystemNpuIpReassembly(o["ip-reassembly"], d, "ip_reassembly", sv)); err != nil {
			if !fortiAPIPatch(o["ip-reassembly"]) {
				return fmt.Errorf("Error reading ip_reassembly: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_reassembly"); ok {
			if err = d.Set("ip_reassembly", flattenSystemNpuIpReassembly(o["ip-reassembly"], d, "ip_reassembly", sv)); err != nil {
				if !fortiAPIPatch(o["ip-reassembly"]) {
					return fmt.Errorf("Error reading ip_reassembly: %v", err)
				}
			}
		}
	}

	if err = d.Set("hash_tbl_spread", flattenSystemNpuHashTblSpread(o["hash-tbl-spread"], d, "hash_tbl_spread", sv)); err != nil {
		if !fortiAPIPatch(o["hash-tbl-spread"]) {
			return fmt.Errorf("Error reading hash_tbl_spread: %v", err)
		}
	}

	if err = d.Set("vlan_lookup_cache", flattenSystemNpuVlanLookupCache(o["vlan-lookup-cache"], d, "vlan_lookup_cache", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-lookup-cache"]) {
			return fmt.Errorf("Error reading vlan_lookup_cache: %v", err)
		}
	}

	if err = d.Set("ip_fragment_offload", flattenSystemNpuIpFragmentOffload(o["ip-fragment-offload"], d, "ip_fragment_offload", sv)); err != nil {
		if !fortiAPIPatch(o["ip-fragment-offload"]) {
			return fmt.Errorf("Error reading ip_fragment_offload: %v", err)
		}
	}

	if err = d.Set("htx_icmp_csum_chk", flattenSystemNpuHtxIcmpCsumChk(o["htx-icmp-csum-chk"], d, "htx_icmp_csum_chk", sv)); err != nil {
		if !fortiAPIPatch(o["htx-icmp-csum-chk"]) {
			return fmt.Errorf("Error reading htx_icmp_csum_chk: %v", err)
		}
	}

	if err = d.Set("htab_msg_queue", flattenSystemNpuHtabMsgQueue(o["htab-msg-queue"], d, "htab_msg_queue", sv)); err != nil {
		if !fortiAPIPatch(o["htab-msg-queue"]) {
			return fmt.Errorf("Error reading htab_msg_queue: %v", err)
		}
	}

	if err = d.Set("htab_dedi_queue_nr", flattenSystemNpuHtabDediQueueNr(o["htab-dedi-queue-nr"], d, "htab_dedi_queue_nr", sv)); err != nil {
		if !fortiAPIPatch(o["htab-dedi-queue-nr"]) {
			return fmt.Errorf("Error reading htab_dedi_queue_nr: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("dsw_dts_profile", flattenSystemNpuDswDtsProfile(o["dsw-dts-profile"], d, "dsw_dts_profile", sv)); err != nil {
			if !fortiAPIPatch(o["dsw-dts-profile"]) {
				return fmt.Errorf("Error reading dsw_dts_profile: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dsw_dts_profile"); ok {
			if err = d.Set("dsw_dts_profile", flattenSystemNpuDswDtsProfile(o["dsw-dts-profile"], d, "dsw_dts_profile", sv)); err != nil {
				if !fortiAPIPatch(o["dsw-dts-profile"]) {
					return fmt.Errorf("Error reading dsw_dts_profile: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dsw_queue_dts_profile", flattenSystemNpuDswQueueDtsProfile(o["dsw-queue-dts-profile"], d, "dsw_queue_dts_profile", sv)); err != nil {
			if !fortiAPIPatch(o["dsw-queue-dts-profile"]) {
				return fmt.Errorf("Error reading dsw_queue_dts_profile: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dsw_queue_dts_profile"); ok {
			if err = d.Set("dsw_queue_dts_profile", flattenSystemNpuDswQueueDtsProfile(o["dsw-queue-dts-profile"], d, "dsw_queue_dts_profile", sv)); err != nil {
				if !fortiAPIPatch(o["dsw-queue-dts-profile"]) {
					return fmt.Errorf("Error reading dsw_queue_dts_profile: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("npu_tcam", flattenSystemNpuNpuTcam(o["npu-tcam"], d, "npu_tcam", sv)); err != nil {
			if !fortiAPIPatch(o["npu-tcam"]) {
				return fmt.Errorf("Error reading npu_tcam: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("npu_tcam"); ok {
			if err = d.Set("npu_tcam", flattenSystemNpuNpuTcam(o["npu-tcam"], d, "npu_tcam", sv)); err != nil {
				if !fortiAPIPatch(o["npu-tcam"]) {
					return fmt.Errorf("Error reading npu_tcam: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("np_queues", flattenSystemNpuNpQueues(o["np-queues"], d, "np_queues", sv)); err != nil {
			if !fortiAPIPatch(o["np-queues"]) {
				return fmt.Errorf("Error reading np_queues: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("np_queues"); ok {
			if err = d.Set("np_queues", flattenSystemNpuNpQueues(o["np-queues"], d, "np_queues", sv)); err != nil {
				if !fortiAPIPatch(o["np-queues"]) {
					return fmt.Errorf("Error reading np_queues: %v", err)
				}
			}
		}
	}

	if err = d.Set("sw_np_rate", flattenSystemNpuSwNpRate(o["sw-np-rate"], d, "sw_np_rate", sv)); err != nil {
		if !fortiAPIPatch(o["sw-np-rate"]) {
			return fmt.Errorf("Error reading sw_np_rate: %v", err)
		}
	}

	if err = d.Set("sw_np_rate_unit", flattenSystemNpuSwNpRateUnit(o["sw-np-rate-unit"], d, "sw_np_rate_unit", sv)); err != nil {
		if !fortiAPIPatch(o["sw-np-rate-unit"]) {
			return fmt.Errorf("Error reading sw_np_rate_unit: %v", err)
		}
	}

	if err = d.Set("sw_np_rate_burst", flattenSystemNpuSwNpRateBurst(o["sw-np-rate-burst"], d, "sw_np_rate_burst", sv)); err != nil {
		if !fortiAPIPatch(o["sw-np-rate-burst"]) {
			return fmt.Errorf("Error reading sw_np_rate_burst: %v", err)
		}
	}

	if err = d.Set("double_level_mcast_offload", flattenSystemNpuDoubleLevelMcastOffload(o["double-level-mcast-offload"], d, "double_level_mcast_offload", sv)); err != nil {
		if !fortiAPIPatch(o["double-level-mcast-offload"]) {
			return fmt.Errorf("Error reading double_level_mcast_offload: %v", err)
		}
	}

	if err = d.Set("mcast_denied_ses_offload", flattenSystemNpuMcastDeniedSesOffload(o["mcast-denied-ses-offload"], d, "mcast_denied_ses_offload", sv)); err != nil {
		if !fortiAPIPatch(o["mcast-denied-ses-offload"]) {
			return fmt.Errorf("Error reading mcast_denied_ses_offload: %v", err)
		}
	}

	if err = d.Set("ipsec_enc_subengine_mask", flattenSystemNpuIpsecEncSubengineMask(o["ipsec-enc-subengine-mask"], d, "ipsec_enc_subengine_mask", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-enc-subengine-mask"]) {
			return fmt.Errorf("Error reading ipsec_enc_subengine_mask: %v", err)
		}
	}

	if err = d.Set("ipsec_dec_subengine_mask", flattenSystemNpuIpsecDecSubengineMask(o["ipsec-dec-subengine-mask"], d, "ipsec_dec_subengine_mask", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-dec-subengine-mask"]) {
			return fmt.Errorf("Error reading ipsec_dec_subengine_mask: %v", err)
		}
	}

	if err = d.Set("np6_cps_optimization_mode", flattenSystemNpuNp6CpsOptimizationMode(o["np6-cps-optimization-mode"], d, "np6_cps_optimization_mode", sv)); err != nil {
		if !fortiAPIPatch(o["np6-cps-optimization-mode"]) {
			return fmt.Errorf("Error reading np6_cps_optimization_mode: %v", err)
		}
	}

	if err = d.Set("sw_np_bandwidth", flattenSystemNpuSwNpBandwidth(o["sw-np-bandwidth"], d, "sw_np_bandwidth", sv)); err != nil {
		if !fortiAPIPatch(o["sw-np-bandwidth"]) {
			return fmt.Errorf("Error reading sw_np_bandwidth: %v", err)
		}
	}

	if err = d.Set("strip_esp_padding", flattenSystemNpuStripEspPadding(o["strip-esp-padding"], d, "strip_esp_padding", sv)); err != nil {
		if !fortiAPIPatch(o["strip-esp-padding"]) {
			return fmt.Errorf("Error reading strip_esp_padding: %v", err)
		}
	}

	if err = d.Set("strip_clear_text_padding", flattenSystemNpuStripClearTextPadding(o["strip-clear-text-padding"], d, "strip_clear_text_padding", sv)); err != nil {
		if !fortiAPIPatch(o["strip-clear-text-padding"]) {
			return fmt.Errorf("Error reading strip_clear_text_padding: %v", err)
		}
	}

	if err = d.Set("ipsec_inbound_cache", flattenSystemNpuIpsecInboundCache(o["ipsec-inbound-cache"], d, "ipsec_inbound_cache", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-inbound-cache"]) {
			return fmt.Errorf("Error reading ipsec_inbound_cache: %v", err)
		}
	}

	if err = d.Set("sse_backpressure", flattenSystemNpuSseBackpressure(o["sse-backpressure"], d, "sse_backpressure", sv)); err != nil {
		if !fortiAPIPatch(o["sse-backpressure"]) {
			return fmt.Errorf("Error reading sse_backpressure: %v", err)
		}
	}

	if err = d.Set("rdp_offload", flattenSystemNpuRdpOffload(o["rdp-offload"], d, "rdp_offload", sv)); err != nil {
		if !fortiAPIPatch(o["rdp-offload"]) {
			return fmt.Errorf("Error reading rdp_offload: %v", err)
		}
	}

	if err = d.Set("ipsec_over_vlink", flattenSystemNpuIpsecOverVlink(o["ipsec-over-vlink"], d, "ipsec_over_vlink", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-over-vlink"]) {
			return fmt.Errorf("Error reading ipsec_over_vlink: %v", err)
		}
	}

	if err = d.Set("uesp_offload", flattenSystemNpuUespOffload(o["uesp-offload"], d, "uesp_offload", sv)); err != nil {
		if !fortiAPIPatch(o["uesp-offload"]) {
			return fmt.Errorf("Error reading uesp_offload: %v", err)
		}
	}

	if err = d.Set("mcast_session_accounting", flattenSystemNpuMcastSessionAccounting(o["mcast-session-accounting"], d, "mcast_session_accounting", sv)); err != nil {
		if !fortiAPIPatch(o["mcast-session-accounting"]) {
			return fmt.Errorf("Error reading mcast_session_accounting: %v", err)
		}
	}

	if err = d.Set("ipsec_mtu_override", flattenSystemNpuIpsecMtuOverride(o["ipsec-mtu-override"], d, "ipsec_mtu_override", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-mtu-override"]) {
			return fmt.Errorf("Error reading ipsec_mtu_override: %v", err)
		}
	}

	if err = d.Set("session_denied_offload", flattenSystemNpuSessionDeniedOffload(o["session-denied-offload"], d, "session_denied_offload", sv)); err != nil {
		if !fortiAPIPatch(o["session-denied-offload"]) {
			return fmt.Errorf("Error reading session_denied_offload: %v", err)
		}
	}

	if err = d.Set("qtm_buf_mode", flattenSystemNpuQtmBufMode(o["qtm-buf-mode"], d, "qtm_buf_mode", sv)); err != nil {
		if !fortiAPIPatch(o["qtm-buf-mode"]) {
			return fmt.Errorf("Error reading qtm_buf_mode: %v", err)
		}
	}

	if err = d.Set("ipsec_ob_np_sel", flattenSystemNpuIpsecObNpSel(o["ipsec-ob-np-sel"], d, "ipsec_ob_np_sel", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-ob-np-sel"]) {
			return fmt.Errorf("Error reading ipsec_ob_np_sel: %v", err)
		}
	}

	if err = d.Set("max_receive_unit", flattenSystemNpuMaxReceiveUnit(o["max-receive-unit"], d, "max_receive_unit", sv)); err != nil {
		if !fortiAPIPatch(o["max-receive-unit"]) {
			return fmt.Errorf("Error reading max_receive_unit: %v", err)
		}
	}

	if err = d.Set("lag_hash_gre", flattenSystemNpuLagHashGre(o["lag-hash-gre"], d, "lag_hash_gre", sv)); err != nil {
		if !fortiAPIPatch(o["lag-hash-gre"]) {
			return fmt.Errorf("Error reading lag_hash_gre: %v", err)
		}
	}

	if err = d.Set("use_mse_oft", flattenSystemNpuUseMseOft(o["use-mse-oft"], d, "use_mse_oft", sv)); err != nil {
		if !fortiAPIPatch(o["use-mse-oft"]) {
			return fmt.Errorf("Error reading use_mse_oft: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ike_port", flattenSystemNpuIkePort(o["ike-port"], d, "ike_port", sv)); err != nil {
			if !fortiAPIPatch(o["ike-port"]) {
				return fmt.Errorf("Error reading ike_port: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ike_port"); ok {
			if err = d.Set("ike_port", flattenSystemNpuIkePort(o["ike-port"], d, "ike_port", sv)); err != nil {
				if !fortiAPIPatch(o["ike-port"]) {
					return fmt.Errorf("Error reading ike_port: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("priority_protocol", flattenSystemNpuPriorityProtocol(o["priority-protocol"], d, "priority_protocol", sv)); err != nil {
			if !fortiAPIPatch(o["priority-protocol"]) {
				return fmt.Errorf("Error reading priority_protocol: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("priority_protocol"); ok {
			if err = d.Set("priority_protocol", flattenSystemNpuPriorityProtocol(o["priority-protocol"], d, "priority_protocol", sv)); err != nil {
				if !fortiAPIPatch(o["priority-protocol"]) {
					return fmt.Errorf("Error reading priority_protocol: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemNpuFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemNpuDedicatedManagementCpu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDedicatedLacpQueue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDedicatedManagementAffinity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDosOptions(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "npu_dos_meter_mode"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["npu-dos-meter-mode"] = nil
		} else {
			result["npu-dos-meter-mode"], _ = expandSystemNpuDosOptionsNpuDosMeterMode(d, i["npu_dos_meter_mode"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "npu_dos_tpe_mode"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["npu-dos-tpe-mode"] = nil
		} else {
			result["npu-dos-tpe-mode"], _ = expandSystemNpuDosOptionsNpuDosTpeMode(d, i["npu_dos_tpe_mode"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemNpuDosOptionsNpuDosMeterMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDosOptionsNpuDosTpeMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuPolicyOffloadLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNapiBreakInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuHpe(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "all_protocol"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["all-protocol"] = nil
		} else {
			result["all-protocol"], _ = expandSystemNpuHpeAllProtocol(d, i["all_protocol"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "tcpsyn_max"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["tcpsyn-max"] = nil
		} else {
			result["tcpsyn-max"], _ = expandSystemNpuHpeTcpsynMax(d, i["tcpsyn_max"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "tcpsyn_ack_max"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["tcpsyn-ack-max"] = nil
		} else {
			result["tcpsyn-ack-max"], _ = expandSystemNpuHpeTcpsynAckMax(d, i["tcpsyn_ack_max"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "tcpfin_rst_max"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["tcpfin-rst-max"] = nil
		} else {
			result["tcpfin-rst-max"], _ = expandSystemNpuHpeTcpfinRstMax(d, i["tcpfin_rst_max"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "tcp_max"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["tcp-max"] = nil
		} else {
			result["tcp-max"], _ = expandSystemNpuHpeTcpMax(d, i["tcp_max"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "udp_max"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["udp-max"] = nil
		} else {
			result["udp-max"], _ = expandSystemNpuHpeUdpMax(d, i["udp_max"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "icmp_max"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["icmp-max"] = nil
		} else {
			result["icmp-max"], _ = expandSystemNpuHpeIcmpMax(d, i["icmp_max"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "sctp_max"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["sctp-max"] = nil
		} else {
			result["sctp-max"], _ = expandSystemNpuHpeSctpMax(d, i["sctp_max"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "esp_max"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["esp-max"] = nil
		} else {
			result["esp-max"], _ = expandSystemNpuHpeEspMax(d, i["esp_max"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ip_frag_max"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ip-frag-max"] = nil
		} else {
			result["ip-frag-max"], _ = expandSystemNpuHpeIpFragMax(d, i["ip_frag_max"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ip_others_max"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ip-others-max"] = nil
		} else {
			result["ip-others-max"], _ = expandSystemNpuHpeIpOthersMax(d, i["ip_others_max"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "arp_max"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["arp-max"] = nil
		} else {
			result["arp-max"], _ = expandSystemNpuHpeArpMax(d, i["arp_max"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "l2_others_max"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["l2-others-max"] = nil
		} else {
			result["l2-others-max"], _ = expandSystemNpuHpeL2OthersMax(d, i["l2_others_max"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "high_priority"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["high-priority"] = nil
		} else {
			result["high-priority"], _ = expandSystemNpuHpeHighPriority(d, i["high_priority"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "enable_shaper"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["enable-shaper"] = nil
		} else {
			result["enable-shaper"], _ = expandSystemNpuHpeEnableShaper(d, i["enable_shaper"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemNpuHpeAllProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuHpeTcpsynMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuHpeTcpsynAckMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuHpeTcpfinRstMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuHpeTcpMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuHpeUdpMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuHpeIcmpMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuHpeSctpMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuHpeEspMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuHpeIpFragMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuHpeIpOthersMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuHpeArpMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuHpeL2OthersMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuHpeHighPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuHpeEnableShaper(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFastpath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuCapwapOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuVxlanOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuVxlanMacFlappingGuard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDefaultQosType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDefaultIpsecMcsType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuShapingStats(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuMcsHostPacketTpeShaping(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuGtpSupport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuPerSessionAccounting(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuSessionAcctInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuMaxSessionTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomaly(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "tcp_syn_fin"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["tcp-syn-fin"] = nil
		} else {
			result["tcp-syn-fin"], _ = expandSystemNpuFpAnomalyTcpSynFin(d, i["tcp_syn_fin"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "tcp_fin_noack"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["tcp-fin-noack"] = nil
		} else {
			result["tcp-fin-noack"], _ = expandSystemNpuFpAnomalyTcpFinNoack(d, i["tcp_fin_noack"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "tcp_fin_only"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["tcp-fin-only"] = nil
		} else {
			result["tcp-fin-only"], _ = expandSystemNpuFpAnomalyTcpFinOnly(d, i["tcp_fin_only"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "tcp_no_flag"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["tcp-no-flag"] = nil
		} else {
			result["tcp-no-flag"], _ = expandSystemNpuFpAnomalyTcpNoFlag(d, i["tcp_no_flag"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "tcp_syn_data"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["tcp-syn-data"] = nil
		} else {
			result["tcp-syn-data"], _ = expandSystemNpuFpAnomalyTcpSynData(d, i["tcp_syn_data"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "tcp_winnuke"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["tcp-winnuke"] = nil
		} else {
			result["tcp-winnuke"], _ = expandSystemNpuFpAnomalyTcpWinnuke(d, i["tcp_winnuke"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "tcp_land"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["tcp-land"] = nil
		} else {
			result["tcp-land"], _ = expandSystemNpuFpAnomalyTcpLand(d, i["tcp_land"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "udp_land"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["udp-land"] = nil
		} else {
			result["udp-land"], _ = expandSystemNpuFpAnomalyUdpLand(d, i["udp_land"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "icmp_land"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["icmp-land"] = nil
		} else {
			result["icmp-land"], _ = expandSystemNpuFpAnomalyIcmpLand(d, i["icmp_land"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "icmp_frag"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["icmp-frag"] = nil
		} else {
			result["icmp-frag"], _ = expandSystemNpuFpAnomalyIcmpFrag(d, i["icmp_frag"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv4_land"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv4-land"] = nil
		} else {
			result["ipv4-land"], _ = expandSystemNpuFpAnomalyIpv4Land(d, i["ipv4_land"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv4_proto_err"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv4-proto-err"] = nil
		} else {
			result["ipv4-proto-err"], _ = expandSystemNpuFpAnomalyIpv4ProtoErr(d, i["ipv4_proto_err"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv4_unknopt"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv4-unknopt"] = nil
		} else {
			result["ipv4-unknopt"], _ = expandSystemNpuFpAnomalyIpv4Unknopt(d, i["ipv4_unknopt"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv4_optrr"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv4-optrr"] = nil
		} else {
			result["ipv4-optrr"], _ = expandSystemNpuFpAnomalyIpv4Optrr(d, i["ipv4_optrr"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv4_optssrr"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv4-optssrr"] = nil
		} else {
			result["ipv4-optssrr"], _ = expandSystemNpuFpAnomalyIpv4Optssrr(d, i["ipv4_optssrr"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv4_optlsrr"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv4-optlsrr"] = nil
		} else {
			result["ipv4-optlsrr"], _ = expandSystemNpuFpAnomalyIpv4Optlsrr(d, i["ipv4_optlsrr"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv4_optstream"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv4-optstream"] = nil
		} else {
			result["ipv4-optstream"], _ = expandSystemNpuFpAnomalyIpv4Optstream(d, i["ipv4_optstream"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv4_optsecurity"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv4-optsecurity"] = nil
		} else {
			result["ipv4-optsecurity"], _ = expandSystemNpuFpAnomalyIpv4Optsecurity(d, i["ipv4_optsecurity"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv4_opttimestamp"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv4-opttimestamp"] = nil
		} else {
			result["ipv4-opttimestamp"], _ = expandSystemNpuFpAnomalyIpv4Opttimestamp(d, i["ipv4_opttimestamp"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv4_csum_err"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv4-csum-err"] = nil
		} else {
			result["ipv4-csum-err"], _ = expandSystemNpuFpAnomalyIpv4CsumErr(d, i["ipv4_csum_err"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "tcp_csum_err"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["tcp-csum-err"] = nil
		} else {
			result["tcp-csum-err"], _ = expandSystemNpuFpAnomalyTcpCsumErr(d, i["tcp_csum_err"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "udp_csum_err"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["udp-csum-err"] = nil
		} else {
			result["udp-csum-err"], _ = expandSystemNpuFpAnomalyUdpCsumErr(d, i["udp_csum_err"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "udplite_csum_err"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["udplite-csum-err"] = nil
		} else {
			result["udplite-csum-err"], _ = expandSystemNpuFpAnomalyUdpliteCsumErr(d, i["udplite_csum_err"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "icmp_csum_err"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["icmp-csum-err"] = nil
		} else {
			result["icmp-csum-err"], _ = expandSystemNpuFpAnomalyIcmpCsumErr(d, i["icmp_csum_err"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "gre_csum_err"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["gre-csum-err"] = nil
		} else {
			result["gre-csum-err"], _ = expandSystemNpuFpAnomalyGreCsumErr(d, i["gre_csum_err"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "sctp_csum_err"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["sctp-csum-err"] = nil
		} else {
			result["sctp-csum-err"], _ = expandSystemNpuFpAnomalySctpCsumErr(d, i["sctp_csum_err"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv6_land"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv6-land"] = nil
		} else {
			result["ipv6-land"], _ = expandSystemNpuFpAnomalyIpv6Land(d, i["ipv6_land"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv6_proto_err"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv6-proto-err"] = nil
		} else {
			result["ipv6-proto-err"], _ = expandSystemNpuFpAnomalyIpv6ProtoErr(d, i["ipv6_proto_err"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv6_unknopt"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv6-unknopt"] = nil
		} else {
			result["ipv6-unknopt"], _ = expandSystemNpuFpAnomalyIpv6Unknopt(d, i["ipv6_unknopt"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv6_saddr_err"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv6-saddr-err"] = nil
		} else {
			result["ipv6-saddr-err"], _ = expandSystemNpuFpAnomalyIpv6SaddrErr(d, i["ipv6_saddr_err"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv6_daddr_err"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv6-daddr-err"] = nil
		} else {
			result["ipv6-daddr-err"], _ = expandSystemNpuFpAnomalyIpv6DaddrErr(d, i["ipv6_daddr_err"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv6_optralert"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv6-optralert"] = nil
		} else {
			result["ipv6-optralert"], _ = expandSystemNpuFpAnomalyIpv6Optralert(d, i["ipv6_optralert"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv6_optjumbo"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv6-optjumbo"] = nil
		} else {
			result["ipv6-optjumbo"], _ = expandSystemNpuFpAnomalyIpv6Optjumbo(d, i["ipv6_optjumbo"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv6_opttunnel"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv6-opttunnel"] = nil
		} else {
			result["ipv6-opttunnel"], _ = expandSystemNpuFpAnomalyIpv6Opttunnel(d, i["ipv6_opttunnel"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv6_opthomeaddr"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv6-opthomeaddr"] = nil
		} else {
			result["ipv6-opthomeaddr"], _ = expandSystemNpuFpAnomalyIpv6Opthomeaddr(d, i["ipv6_opthomeaddr"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv6_optnsap"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv6-optnsap"] = nil
		} else {
			result["ipv6-optnsap"], _ = expandSystemNpuFpAnomalyIpv6Optnsap(d, i["ipv6_optnsap"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv6_optendpid"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv6-optendpid"] = nil
		} else {
			result["ipv6-optendpid"], _ = expandSystemNpuFpAnomalyIpv6Optendpid(d, i["ipv6_optendpid"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ipv6_optinvld"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ipv6-optinvld"] = nil
		} else {
			result["ipv6-optinvld"], _ = expandSystemNpuFpAnomalyIpv6Optinvld(d, i["ipv6_optinvld"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemNpuFpAnomalyTcpSynFin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyTcpFinNoack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyTcpFinOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyTcpNoFlag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyTcpSynData(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyTcpWinnuke(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyTcpLand(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyUdpLand(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIcmpLand(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIcmpFrag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv4Land(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv4ProtoErr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv4Unknopt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv4Optrr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv4Optssrr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv4Optlsrr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv4Optstream(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv4Optsecurity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv4Opttimestamp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv4CsumErr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyTcpCsumErr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyUdpCsumErr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyUdpliteCsumErr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIcmpCsumErr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyGreCsumErr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalySctpCsumErr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv6Land(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv6ProtoErr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv6Unknopt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv6SaddrErr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv6DaddrErr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv6Optralert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv6Optjumbo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv6Opttunnel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv6Opthomeaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv6Optnsap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv6Optendpid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFpAnomalyIpv6Optinvld(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuIpReassembly(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "min_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["min-timeout"] = nil
		} else {
			result["min-timeout"], _ = expandSystemNpuIpReassemblyMinTimeout(d, i["min_timeout"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "max_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["max-timeout"] = nil
		} else {
			result["max-timeout"], _ = expandSystemNpuIpReassemblyMaxTimeout(d, i["max_timeout"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["status"] = nil
		} else {
			result["status"], _ = expandSystemNpuIpReassemblyStatus(d, i["status"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemNpuIpReassemblyMinTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuIpReassemblyMaxTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuIpReassemblyStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuHashTblSpread(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuVlanLookupCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuIpFragmentOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuHtxIcmpCsumChk(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuHtabMsgQueue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuHtabDediQueueNr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDswDtsProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "profile_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["profile-id"], _ = expandSystemNpuDswDtsProfileProfileId(d, i["profile_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["profile-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_limit"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["min-limit"], _ = expandSystemNpuDswDtsProfileMinLimit(d, i["min_limit"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["min-limit"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "step"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["step"], _ = expandSystemNpuDswDtsProfileStep(d, i["step"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["step"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandSystemNpuDswDtsProfileAction(d, i["action"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemNpuDswDtsProfileProfileId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDswDtsProfileMinLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDswDtsProfileStep(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDswDtsProfileAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDswQueueDtsProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemNpuDswQueueDtsProfileName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "iport"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["iport"], _ = expandSystemNpuDswQueueDtsProfileIport(d, i["iport"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oport"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["oport"], _ = expandSystemNpuDswQueueDtsProfileOport(d, i["oport"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "profile_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["profile-id"], _ = expandSystemNpuDswQueueDtsProfileProfileId(d, i["profile_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["profile-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue_select"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["queue-select"], _ = expandSystemNpuDswQueueDtsProfileQueueSelect(d, i["queue_select"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["queue-select"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemNpuDswQueueDtsProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDswQueueDtsProfileIport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDswQueueDtsProfileOport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDswQueueDtsProfileProfileId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDswQueueDtsProfileQueueSelect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcam(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemNpuNpuTcamName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandSystemNpuNpuTcamType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oid"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["oid"], _ = expandSystemNpuNpuTcamOid(d, i["oid"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["oid"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vid"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vid"], _ = expandSystemNpuNpuTcamVid(d, i["vid"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vid"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "data"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["data"], _ = expandSystemNpuNpuTcamData(d, i["data"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["data"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mask"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mask"], _ = expandSystemNpuNpuTcamMask(d, i["mask"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["mask"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mir_act"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mir-act"], _ = expandSystemNpuNpuTcamMirAct(d, i["mir_act"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["mir-act"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pri_act"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pri-act"], _ = expandSystemNpuNpuTcamPriAct(d, i["pri_act"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["pri-act"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sact"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sact"], _ = expandSystemNpuNpuTcamSact(d, i["sact"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["sact"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tact"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tact"], _ = expandSystemNpuNpuTcamTact(d, i["tact"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["tact"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemNpuNpuTcamName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamOid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamVid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamData(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "gen_buf_cnt"
	if _, ok := d.GetOk(pre_append); ok {
		result["gen-buf-cnt"], _ = expandSystemNpuNpuTcamDataGenBufCnt(d, i["gen_buf_cnt"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["gen-buf-cnt"] = nil
	}
	pre_append = pre + ".0." + "gen_pri"
	if _, ok := d.GetOk(pre_append); ok {
		result["gen-pri"], _ = expandSystemNpuNpuTcamDataGenPri(d, i["gen_pri"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["gen-pri"] = nil
	}
	pre_append = pre + ".0." + "gen_pri_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["gen-pri-v"], _ = expandSystemNpuNpuTcamDataGenPriV(d, i["gen_pri_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "gen_iv"
	if _, ok := d.GetOk(pre_append); ok {
		result["gen-iv"], _ = expandSystemNpuNpuTcamDataGenIv(d, i["gen_iv"], pre_append, sv)
	}
	pre_append = pre + ".0." + "gen_tv"
	if _, ok := d.GetOk(pre_append); ok {
		result["gen-tv"], _ = expandSystemNpuNpuTcamDataGenTv(d, i["gen_tv"], pre_append, sv)
	}
	pre_append = pre + ".0." + "gen_pkt_ctrl"
	if _, ok := d.GetOk(pre_append); ok {
		result["gen-pkt-ctrl"], _ = expandSystemNpuNpuTcamDataGenPktCtrl(d, i["gen_pkt_ctrl"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["gen-pkt-ctrl"] = nil
	}
	pre_append = pre + ".0." + "gen_l3_flags"
	if _, ok := d.GetOk(pre_append); ok {
		result["gen-l3-flags"], _ = expandSystemNpuNpuTcamDataGenL3Flags(d, i["gen_l3_flags"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["gen-l3-flags"] = nil
	}
	pre_append = pre + ".0." + "gen_l4_flags"
	if _, ok := d.GetOk(pre_append); ok {
		result["gen-l4-flags"], _ = expandSystemNpuNpuTcamDataGenL4Flags(d, i["gen_l4_flags"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["gen-l4-flags"] = nil
	}
	pre_append = pre + ".0." + "vdid"
	if _, ok := d.GetOk(pre_append); ok {
		result["vdid"], _ = expandSystemNpuNpuTcamDataVdid(d, i["vdid"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["vdid"] = nil
	}
	pre_append = pre + ".0." + "tp"
	if _, ok := d.GetOk(pre_append); ok {
		result["tp"], _ = expandSystemNpuNpuTcamDataTp(d, i["tp"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["tp"] = nil
	}
	pre_append = pre + ".0." + "tgt_updt"
	if _, ok := d.GetOk(pre_append); ok {
		result["tgt-updt"], _ = expandSystemNpuNpuTcamDataTgtUpdt(d, i["tgt_updt"], pre_append, sv)
	}
	pre_append = pre + ".0." + "smac_change"
	if _, ok := d.GetOk(pre_append); ok {
		result["smac-change"], _ = expandSystemNpuNpuTcamDataSmacChange(d, i["smac_change"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ext_tag"
	if _, ok := d.GetOk(pre_append); ok {
		result["ext-tag"], _ = expandSystemNpuNpuTcamDataExtTag(d, i["ext_tag"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tgt_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["tgt-v"], _ = expandSystemNpuNpuTcamDataTgtV(d, i["tgt_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tvid"
	if _, ok := d.GetOk(pre_append); ok {
		result["tvid"], _ = expandSystemNpuNpuTcamDataTvid(d, i["tvid"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["tvid"] = nil
	}
	pre_append = pre + ".0." + "tgt_cfi"
	if _, ok := d.GetOk(pre_append); ok {
		result["tgt-cfi"], _ = expandSystemNpuNpuTcamDataTgtCfi(d, i["tgt_cfi"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tgt_prio"
	if _, ok := d.GetOk(pre_append); ok {
		result["tgt-prio"], _ = expandSystemNpuNpuTcamDataTgtPrio(d, i["tgt_prio"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["tgt-prio"] = nil
	}
	pre_append = pre + ".0." + "sp"
	if _, ok := d.GetOk(pre_append); ok {
		result["sp"], _ = expandSystemNpuNpuTcamDataSp(d, i["sp"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["sp"] = nil
	}
	pre_append = pre + ".0." + "src_updt"
	if _, ok := d.GetOk(pre_append); ok {
		result["src-updt"], _ = expandSystemNpuNpuTcamDataSrcUpdt(d, i["src_updt"], pre_append, sv)
	}
	pre_append = pre + ".0." + "slink"
	if _, ok := d.GetOk(pre_append); ok {
		result["slink"], _ = expandSystemNpuNpuTcamDataSlink(d, i["slink"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["slink"] = nil
	}
	pre_append = pre + ".0." + "svid"
	if _, ok := d.GetOk(pre_append); ok {
		result["svid"], _ = expandSystemNpuNpuTcamDataSvid(d, i["svid"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["svid"] = nil
	}
	pre_append = pre + ".0." + "src_cfi"
	if _, ok := d.GetOk(pre_append); ok {
		result["src-cfi"], _ = expandSystemNpuNpuTcamDataSrcCfi(d, i["src_cfi"], pre_append, sv)
	}
	pre_append = pre + ".0." + "src_prio"
	if _, ok := d.GetOk(pre_append); ok {
		result["src-prio"], _ = expandSystemNpuNpuTcamDataSrcPrio(d, i["src_prio"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["src-prio"] = nil
	}
	pre_append = pre + ".0." + "srcmac"
	if _, ok := d.GetOk(pre_append); ok {
		result["srcmac"], _ = expandSystemNpuNpuTcamDataSrcmac(d, i["srcmac"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dstmac"
	if _, ok := d.GetOk(pre_append); ok {
		result["dstmac"], _ = expandSystemNpuNpuTcamDataDstmac(d, i["dstmac"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ethertype"
	if _, ok := d.GetOk(pre_append); ok {
		result["ethertype"], _ = expandSystemNpuNpuTcamDataEthertype(d, i["ethertype"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ipver"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipver"], _ = expandSystemNpuNpuTcamDataIpver(d, i["ipver"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["ipver"] = nil
	}
	pre_append = pre + ".0." + "ihl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ihl"], _ = expandSystemNpuNpuTcamDataIhl(d, i["ihl"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["ihl"] = nil
	}
	pre_append = pre + ".0." + "ip4_id"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip4-id"], _ = expandSystemNpuNpuTcamDataIp4Id(d, i["ip4_id"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["ip4-id"] = nil
	}
	pre_append = pre + ".0." + "srcip"
	if _, ok := d.GetOk(pre_append); ok {
		result["srcip"], _ = expandSystemNpuNpuTcamDataSrcip(d, i["srcip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dstip"
	if _, ok := d.GetOk(pre_append); ok {
		result["dstip"], _ = expandSystemNpuNpuTcamDataDstip(d, i["dstip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ip6_fl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-fl"], _ = expandSystemNpuNpuTcamDataIp6Fl(d, i["ip6_fl"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["ip6-fl"] = nil
	}
	pre_append = pre + ".0." + "srcipv6"
	if _, ok := d.GetOk(pre_append); ok {
		result["srcipv6"], _ = expandSystemNpuNpuTcamDataSrcipv6(d, i["srcipv6"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dstipv6"
	if _, ok := d.GetOk(pre_append); ok {
		result["dstipv6"], _ = expandSystemNpuNpuTcamDataDstipv6(d, i["dstipv6"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ttl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ttl"], _ = expandSystemNpuNpuTcamDataTtl(d, i["ttl"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["ttl"] = nil
	}
	pre_append = pre + ".0." + "protocol"
	if _, ok := d.GetOk(pre_append); ok {
		result["protocol"], _ = expandSystemNpuNpuTcamDataProtocol(d, i["protocol"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["protocol"] = nil
	}
	pre_append = pre + ".0." + "tos"
	if _, ok := d.GetOk(pre_append); ok {
		result["tos"], _ = expandSystemNpuNpuTcamDataTos(d, i["tos"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["tos"] = nil
	}
	pre_append = pre + ".0." + "frag_off"
	if _, ok := d.GetOk(pre_append); ok {
		result["frag-off"], _ = expandSystemNpuNpuTcamDataFragOff(d, i["frag_off"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["frag-off"] = nil
	}
	pre_append = pre + ".0." + "mf"
	if _, ok := d.GetOk(pre_append); ok {
		result["mf"], _ = expandSystemNpuNpuTcamDataMf(d, i["mf"], pre_append, sv)
	}
	pre_append = pre + ".0." + "df"
	if _, ok := d.GetOk(pre_append); ok {
		result["df"], _ = expandSystemNpuNpuTcamDataDf(d, i["df"], pre_append, sv)
	}
	pre_append = pre + ".0." + "srcport"
	if _, ok := d.GetOk(pre_append); ok {
		result["srcport"], _ = expandSystemNpuNpuTcamDataSrcport(d, i["srcport"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["srcport"] = nil
	}
	pre_append = pre + ".0." + "dstport"
	if _, ok := d.GetOk(pre_append); ok {
		result["dstport"], _ = expandSystemNpuNpuTcamDataDstport(d, i["dstport"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["dstport"] = nil
	}
	pre_append = pre + ".0." + "tcp_fin"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-fin"], _ = expandSystemNpuNpuTcamDataTcpFin(d, i["tcp_fin"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tcp_syn"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-syn"], _ = expandSystemNpuNpuTcamDataTcpSyn(d, i["tcp_syn"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tcp_rst"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-rst"], _ = expandSystemNpuNpuTcamDataTcpRst(d, i["tcp_rst"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tcp_push"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-push"], _ = expandSystemNpuNpuTcamDataTcpPush(d, i["tcp_push"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tcp_ack"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-ack"], _ = expandSystemNpuNpuTcamDataTcpAck(d, i["tcp_ack"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tcp_urg"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-urg"], _ = expandSystemNpuNpuTcamDataTcpUrg(d, i["tcp_urg"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tcp_ece"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-ece"], _ = expandSystemNpuNpuTcamDataTcpEce(d, i["tcp_ece"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tcp_cwr"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-cwr"], _ = expandSystemNpuNpuTcamDataTcpCwr(d, i["tcp_cwr"], pre_append, sv)
	}
	pre_append = pre + ".0." + "l4_wd8"
	if _, ok := d.GetOk(pre_append); ok {
		result["l4-wd8"], _ = expandSystemNpuNpuTcamDataL4Wd8(d, i["l4_wd8"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["l4-wd8"] = nil
	}
	pre_append = pre + ".0." + "l4_wd9"
	if _, ok := d.GetOk(pre_append); ok {
		result["l4-wd9"], _ = expandSystemNpuNpuTcamDataL4Wd9(d, i["l4_wd9"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["l4-wd9"] = nil
	}
	pre_append = pre + ".0." + "l4_wd10"
	if _, ok := d.GetOk(pre_append); ok {
		result["l4-wd10"], _ = expandSystemNpuNpuTcamDataL4Wd10(d, i["l4_wd10"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["l4-wd10"] = nil
	}
	pre_append = pre + ".0." + "l4_wd11"
	if _, ok := d.GetOk(pre_append); ok {
		result["l4-wd11"], _ = expandSystemNpuNpuTcamDataL4Wd11(d, i["l4_wd11"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["l4-wd11"] = nil
	}

	return result, nil
}

func expandSystemNpuNpuTcamDataGenBufCnt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataGenPri(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataGenPriV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataGenIv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataGenTv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataGenPktCtrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataGenL3Flags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataGenL4Flags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataVdid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataTp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataTgtUpdt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataSmacChange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataExtTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataTgtV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataTvid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataTgtCfi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataTgtPrio(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataSrcUpdt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataSlink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataSvid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataSrcCfi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataSrcPrio(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataSrcmac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataDstmac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataEthertype(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataIpver(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataIhl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataIp4Id(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataSrcip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataDstip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataIp6Fl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataSrcipv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataDstipv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataTos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataFragOff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataMf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataDf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataSrcport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataDstport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataTcpFin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataTcpSyn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataTcpRst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataTcpPush(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataTcpAck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataTcpUrg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataTcpEce(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataTcpCwr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataL4Wd8(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataL4Wd9(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataL4Wd10(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamDataL4Wd11(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "gen_buf_cnt"
	if _, ok := d.GetOk(pre_append); ok {
		result["gen-buf-cnt"], _ = expandSystemNpuNpuTcamMaskGenBufCnt(d, i["gen_buf_cnt"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["gen-buf-cnt"] = nil
	}
	pre_append = pre + ".0." + "gen_pri"
	if _, ok := d.GetOk(pre_append); ok {
		result["gen-pri"], _ = expandSystemNpuNpuTcamMaskGenPri(d, i["gen_pri"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["gen-pri"] = nil
	}
	pre_append = pre + ".0." + "gen_pri_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["gen-pri-v"], _ = expandSystemNpuNpuTcamMaskGenPriV(d, i["gen_pri_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "gen_iv"
	if _, ok := d.GetOk(pre_append); ok {
		result["gen-iv"], _ = expandSystemNpuNpuTcamMaskGenIv(d, i["gen_iv"], pre_append, sv)
	}
	pre_append = pre + ".0." + "gen_tv"
	if _, ok := d.GetOk(pre_append); ok {
		result["gen-tv"], _ = expandSystemNpuNpuTcamMaskGenTv(d, i["gen_tv"], pre_append, sv)
	}
	pre_append = pre + ".0." + "gen_pkt_ctrl"
	if _, ok := d.GetOk(pre_append); ok {
		result["gen-pkt-ctrl"], _ = expandSystemNpuNpuTcamMaskGenPktCtrl(d, i["gen_pkt_ctrl"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["gen-pkt-ctrl"] = nil
	}
	pre_append = pre + ".0." + "gen_l3_flags"
	if _, ok := d.GetOk(pre_append); ok {
		result["gen-l3-flags"], _ = expandSystemNpuNpuTcamMaskGenL3Flags(d, i["gen_l3_flags"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["gen-l3-flags"] = nil
	}
	pre_append = pre + ".0." + "gen_l4_flags"
	if _, ok := d.GetOk(pre_append); ok {
		result["gen-l4-flags"], _ = expandSystemNpuNpuTcamMaskGenL4Flags(d, i["gen_l4_flags"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["gen-l4-flags"] = nil
	}
	pre_append = pre + ".0." + "vdid"
	if _, ok := d.GetOk(pre_append); ok {
		result["vdid"], _ = expandSystemNpuNpuTcamMaskVdid(d, i["vdid"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["vdid"] = nil
	}
	pre_append = pre + ".0." + "tp"
	if _, ok := d.GetOk(pre_append); ok {
		result["tp"], _ = expandSystemNpuNpuTcamMaskTp(d, i["tp"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["tp"] = nil
	}
	pre_append = pre + ".0." + "tgt_updt"
	if _, ok := d.GetOk(pre_append); ok {
		result["tgt-updt"], _ = expandSystemNpuNpuTcamMaskTgtUpdt(d, i["tgt_updt"], pre_append, sv)
	}
	pre_append = pre + ".0." + "smac_change"
	if _, ok := d.GetOk(pre_append); ok {
		result["smac-change"], _ = expandSystemNpuNpuTcamMaskSmacChange(d, i["smac_change"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ext_tag"
	if _, ok := d.GetOk(pre_append); ok {
		result["ext-tag"], _ = expandSystemNpuNpuTcamMaskExtTag(d, i["ext_tag"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tgt_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["tgt-v"], _ = expandSystemNpuNpuTcamMaskTgtV(d, i["tgt_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tvid"
	if _, ok := d.GetOk(pre_append); ok {
		result["tvid"], _ = expandSystemNpuNpuTcamMaskTvid(d, i["tvid"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["tvid"] = nil
	}
	pre_append = pre + ".0." + "tgt_cfi"
	if _, ok := d.GetOk(pre_append); ok {
		result["tgt-cfi"], _ = expandSystemNpuNpuTcamMaskTgtCfi(d, i["tgt_cfi"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tgt_prio"
	if _, ok := d.GetOk(pre_append); ok {
		result["tgt-prio"], _ = expandSystemNpuNpuTcamMaskTgtPrio(d, i["tgt_prio"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["tgt-prio"] = nil
	}
	pre_append = pre + ".0." + "sp"
	if _, ok := d.GetOk(pre_append); ok {
		result["sp"], _ = expandSystemNpuNpuTcamMaskSp(d, i["sp"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["sp"] = nil
	}
	pre_append = pre + ".0." + "src_updt"
	if _, ok := d.GetOk(pre_append); ok {
		result["src-updt"], _ = expandSystemNpuNpuTcamMaskSrcUpdt(d, i["src_updt"], pre_append, sv)
	}
	pre_append = pre + ".0." + "slink"
	if _, ok := d.GetOk(pre_append); ok {
		result["slink"], _ = expandSystemNpuNpuTcamMaskSlink(d, i["slink"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["slink"] = nil
	}
	pre_append = pre + ".0." + "svid"
	if _, ok := d.GetOk(pre_append); ok {
		result["svid"], _ = expandSystemNpuNpuTcamMaskSvid(d, i["svid"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["svid"] = nil
	}
	pre_append = pre + ".0." + "src_cfi"
	if _, ok := d.GetOk(pre_append); ok {
		result["src-cfi"], _ = expandSystemNpuNpuTcamMaskSrcCfi(d, i["src_cfi"], pre_append, sv)
	}
	pre_append = pre + ".0." + "src_prio"
	if _, ok := d.GetOk(pre_append); ok {
		result["src-prio"], _ = expandSystemNpuNpuTcamMaskSrcPrio(d, i["src_prio"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["src-prio"] = nil
	}
	pre_append = pre + ".0." + "srcmac"
	if _, ok := d.GetOk(pre_append); ok {
		result["srcmac"], _ = expandSystemNpuNpuTcamMaskSrcmac(d, i["srcmac"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dstmac"
	if _, ok := d.GetOk(pre_append); ok {
		result["dstmac"], _ = expandSystemNpuNpuTcamMaskDstmac(d, i["dstmac"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ethertype"
	if _, ok := d.GetOk(pre_append); ok {
		result["ethertype"], _ = expandSystemNpuNpuTcamMaskEthertype(d, i["ethertype"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ipver"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipver"], _ = expandSystemNpuNpuTcamMaskIpver(d, i["ipver"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["ipver"] = nil
	}
	pre_append = pre + ".0." + "ihl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ihl"], _ = expandSystemNpuNpuTcamMaskIhl(d, i["ihl"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["ihl"] = nil
	}
	pre_append = pre + ".0." + "ip4_id"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip4-id"], _ = expandSystemNpuNpuTcamMaskIp4Id(d, i["ip4_id"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["ip4-id"] = nil
	}
	pre_append = pre + ".0." + "srcip"
	if _, ok := d.GetOk(pre_append); ok {
		result["srcip"], _ = expandSystemNpuNpuTcamMaskSrcip(d, i["srcip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dstip"
	if _, ok := d.GetOk(pre_append); ok {
		result["dstip"], _ = expandSystemNpuNpuTcamMaskDstip(d, i["dstip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ip6_fl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-fl"], _ = expandSystemNpuNpuTcamMaskIp6Fl(d, i["ip6_fl"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["ip6-fl"] = nil
	}
	pre_append = pre + ".0." + "srcipv6"
	if _, ok := d.GetOk(pre_append); ok {
		result["srcipv6"], _ = expandSystemNpuNpuTcamMaskSrcipv6(d, i["srcipv6"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dstipv6"
	if _, ok := d.GetOk(pre_append); ok {
		result["dstipv6"], _ = expandSystemNpuNpuTcamMaskDstipv6(d, i["dstipv6"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ttl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ttl"], _ = expandSystemNpuNpuTcamMaskTtl(d, i["ttl"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["ttl"] = nil
	}
	pre_append = pre + ".0." + "protocol"
	if _, ok := d.GetOk(pre_append); ok {
		result["protocol"], _ = expandSystemNpuNpuTcamMaskProtocol(d, i["protocol"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["protocol"] = nil
	}
	pre_append = pre + ".0." + "tos"
	if _, ok := d.GetOk(pre_append); ok {
		result["tos"], _ = expandSystemNpuNpuTcamMaskTos(d, i["tos"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["tos"] = nil
	}
	pre_append = pre + ".0." + "frag_off"
	if _, ok := d.GetOk(pre_append); ok {
		result["frag-off"], _ = expandSystemNpuNpuTcamMaskFragOff(d, i["frag_off"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["frag-off"] = nil
	}
	pre_append = pre + ".0." + "mf"
	if _, ok := d.GetOk(pre_append); ok {
		result["mf"], _ = expandSystemNpuNpuTcamMaskMf(d, i["mf"], pre_append, sv)
	}
	pre_append = pre + ".0." + "df"
	if _, ok := d.GetOk(pre_append); ok {
		result["df"], _ = expandSystemNpuNpuTcamMaskDf(d, i["df"], pre_append, sv)
	}
	pre_append = pre + ".0." + "srcport"
	if _, ok := d.GetOk(pre_append); ok {
		result["srcport"], _ = expandSystemNpuNpuTcamMaskSrcport(d, i["srcport"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["srcport"] = nil
	}
	pre_append = pre + ".0." + "dstport"
	if _, ok := d.GetOk(pre_append); ok {
		result["dstport"], _ = expandSystemNpuNpuTcamMaskDstport(d, i["dstport"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["dstport"] = nil
	}
	pre_append = pre + ".0." + "tcp_fin"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-fin"], _ = expandSystemNpuNpuTcamMaskTcpFin(d, i["tcp_fin"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tcp_syn"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-syn"], _ = expandSystemNpuNpuTcamMaskTcpSyn(d, i["tcp_syn"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tcp_rst"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-rst"], _ = expandSystemNpuNpuTcamMaskTcpRst(d, i["tcp_rst"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tcp_push"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-push"], _ = expandSystemNpuNpuTcamMaskTcpPush(d, i["tcp_push"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tcp_ack"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-ack"], _ = expandSystemNpuNpuTcamMaskTcpAck(d, i["tcp_ack"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tcp_urg"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-urg"], _ = expandSystemNpuNpuTcamMaskTcpUrg(d, i["tcp_urg"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tcp_ece"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-ece"], _ = expandSystemNpuNpuTcamMaskTcpEce(d, i["tcp_ece"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tcp_cwr"
	if _, ok := d.GetOk(pre_append); ok {
		result["tcp-cwr"], _ = expandSystemNpuNpuTcamMaskTcpCwr(d, i["tcp_cwr"], pre_append, sv)
	}
	pre_append = pre + ".0." + "l4_wd8"
	if _, ok := d.GetOk(pre_append); ok {
		result["l4-wd8"], _ = expandSystemNpuNpuTcamMaskL4Wd8(d, i["l4_wd8"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["l4-wd8"] = nil
	}
	pre_append = pre + ".0." + "l4_wd9"
	if _, ok := d.GetOk(pre_append); ok {
		result["l4-wd9"], _ = expandSystemNpuNpuTcamMaskL4Wd9(d, i["l4_wd9"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["l4-wd9"] = nil
	}
	pre_append = pre + ".0." + "l4_wd10"
	if _, ok := d.GetOk(pre_append); ok {
		result["l4-wd10"], _ = expandSystemNpuNpuTcamMaskL4Wd10(d, i["l4_wd10"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["l4-wd10"] = nil
	}
	pre_append = pre + ".0." + "l4_wd11"
	if _, ok := d.GetOk(pre_append); ok {
		result["l4-wd11"], _ = expandSystemNpuNpuTcamMaskL4Wd11(d, i["l4_wd11"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["l4-wd11"] = nil
	}

	return result, nil
}

func expandSystemNpuNpuTcamMaskGenBufCnt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskGenPri(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskGenPriV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskGenIv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskGenTv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskGenPktCtrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskGenL3Flags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskGenL4Flags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskVdid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskTp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskTgtUpdt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskSmacChange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskExtTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskTgtV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskTvid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskTgtCfi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskTgtPrio(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskSrcUpdt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskSlink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskSvid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskSrcCfi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskSrcPrio(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskSrcmac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskDstmac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskEthertype(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskIpver(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskIhl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskIp4Id(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskSrcip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskDstip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskIp6Fl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskSrcipv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskDstipv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskTos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskFragOff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskMf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskDf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskSrcport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskDstport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskTcpFin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskTcpSyn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskTcpRst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskTcpPush(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskTcpAck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskTcpUrg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskTcpEce(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskTcpCwr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskL4Wd8(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskL4Wd9(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskL4Wd10(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMaskL4Wd11(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamMirAct(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "vlif"
	if _, ok := d.GetOk(pre_append); ok {
		result["vlif"], _ = expandSystemNpuNpuTcamMirActVlif(d, i["vlif"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["vlif"] = nil
	}

	return result, nil
}

func expandSystemNpuNpuTcamMirActVlif(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamPriAct(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok {
		result["priority"], _ = expandSystemNpuNpuTcamPriActPriority(d, i["priority"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["priority"] = nil
	}
	pre_append = pre + ".0." + "weight"
	if _, ok := d.GetOk(pre_append); ok {
		result["weight"], _ = expandSystemNpuNpuTcamPriActWeight(d, i["weight"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["weight"] = nil
	}

	return result, nil
}

func expandSystemNpuNpuTcamPriActPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamPriActWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSact(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "fwd_lif_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["fwd-lif-v"], _ = expandSystemNpuNpuTcamSactFwdLifV(d, i["fwd_lif_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "fwd_lif"
	if _, ok := d.GetOk(pre_append); ok {
		result["fwd-lif"], _ = expandSystemNpuNpuTcamSactFwdLif(d, i["fwd_lif"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["fwd-lif"] = nil
	}
	pre_append = pre + ".0." + "fwd_tvid_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["fwd-tvid-v"], _ = expandSystemNpuNpuTcamSactFwdTvidV(d, i["fwd_tvid_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "fwd_tvid"
	if _, ok := d.GetOk(pre_append); ok {
		result["fwd-tvid"], _ = expandSystemNpuNpuTcamSactFwdTvid(d, i["fwd_tvid"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["fwd-tvid"] = nil
	}
	pre_append = pre + ".0." + "df_lif_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["df-lif-v"], _ = expandSystemNpuNpuTcamSactDfLifV(d, i["df_lif_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "df_lif"
	if _, ok := d.GetOk(pre_append); ok {
		result["df-lif"], _ = expandSystemNpuNpuTcamSactDfLif(d, i["df_lif"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["df-lif"] = nil
	}
	pre_append = pre + ".0." + "act_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["act-v"], _ = expandSystemNpuNpuTcamSactActV(d, i["act_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "act"
	if _, ok := d.GetOk(pre_append); ok {
		result["act"], _ = expandSystemNpuNpuTcamSactAct(d, i["act"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["act"] = nil
	}
	pre_append = pre + ".0." + "pleen_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["pleen-v"], _ = expandSystemNpuNpuTcamSactPleenV(d, i["pleen_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "pleen"
	if _, ok := d.GetOk(pre_append); ok {
		result["pleen"], _ = expandSystemNpuNpuTcamSactPleen(d, i["pleen"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["pleen"] = nil
	}
	pre_append = pre + ".0." + "icpen_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["icpen-v"], _ = expandSystemNpuNpuTcamSactIcpenV(d, i["icpen_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "icpen"
	if _, ok := d.GetOk(pre_append); ok {
		result["icpen"], _ = expandSystemNpuNpuTcamSactIcpen(d, i["icpen"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["icpen"] = nil
	}
	pre_append = pre + ".0." + "vdm_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["vdm-v"], _ = expandSystemNpuNpuTcamSactVdmV(d, i["vdm_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vdm"
	if _, ok := d.GetOk(pre_append); ok {
		result["vdm"], _ = expandSystemNpuNpuTcamSactVdm(d, i["vdm"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["vdm"] = nil
	}
	pre_append = pre + ".0." + "learn_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["learn-v"], _ = expandSystemNpuNpuTcamSactLearnV(d, i["learn_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "learn"
	if _, ok := d.GetOk(pre_append); ok {
		result["learn"], _ = expandSystemNpuNpuTcamSactLearn(d, i["learn"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["learn"] = nil
	}
	pre_append = pre + ".0." + "rfsh_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["rfsh-v"], _ = expandSystemNpuNpuTcamSactRfshV(d, i["rfsh_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "rfsh"
	if _, ok := d.GetOk(pre_append); ok {
		result["rfsh"], _ = expandSystemNpuNpuTcamSactRfsh(d, i["rfsh"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["rfsh"] = nil
	}
	pre_append = pre + ".0." + "fwd_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["fwd-v"], _ = expandSystemNpuNpuTcamSactFwdV(d, i["fwd_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "fwd"
	if _, ok := d.GetOk(pre_append); ok {
		result["fwd"], _ = expandSystemNpuNpuTcamSactFwd(d, i["fwd"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["fwd"] = nil
	}
	pre_append = pre + ".0." + "x_mode_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["x-mode-v"], _ = expandSystemNpuNpuTcamSactXModeV(d, i["x_mode_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "x_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["x-mode"], _ = expandSystemNpuNpuTcamSactXMode(d, i["x_mode"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["x-mode"] = nil
	}
	pre_append = pre + ".0." + "promis_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["promis-v"], _ = expandSystemNpuNpuTcamSactPromisV(d, i["promis_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "promis"
	if _, ok := d.GetOk(pre_append); ok {
		result["promis"], _ = expandSystemNpuNpuTcamSactPromis(d, i["promis"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["promis"] = nil
	}
	pre_append = pre + ".0." + "bmproc_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["bmproc-v"], _ = expandSystemNpuNpuTcamSactBmprocV(d, i["bmproc_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bmproc"
	if _, ok := d.GetOk(pre_append); ok {
		result["bmproc"], _ = expandSystemNpuNpuTcamSactBmproc(d, i["bmproc"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["bmproc"] = nil
	}
	pre_append = pre + ".0." + "mac_id_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["mac-id-v"], _ = expandSystemNpuNpuTcamSactMacIdV(d, i["mac_id_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mac_id"
	if _, ok := d.GetOk(pre_append); ok {
		result["mac-id"], _ = expandSystemNpuNpuTcamSactMacId(d, i["mac_id"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["mac-id"] = nil
	}
	pre_append = pre + ".0." + "dosen_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["dosen-v"], _ = expandSystemNpuNpuTcamSactDosenV(d, i["dosen_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dosen"
	if _, ok := d.GetOk(pre_append); ok {
		result["dosen"], _ = expandSystemNpuNpuTcamSactDosen(d, i["dosen"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["dosen"] = nil
	}
	pre_append = pre + ".0." + "dfr_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["dfr-v"], _ = expandSystemNpuNpuTcamSactDfrV(d, i["dfr_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dfr"
	if _, ok := d.GetOk(pre_append); ok {
		result["dfr"], _ = expandSystemNpuNpuTcamSactDfr(d, i["dfr"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["dfr"] = nil
	}
	pre_append = pre + ".0." + "m_srh_ctrl_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["m-srh-ctrl-v"], _ = expandSystemNpuNpuTcamSactMSrhCtrlV(d, i["m_srh_ctrl_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "m_srh_ctrl"
	if _, ok := d.GetOk(pre_append); ok {
		result["m-srh-ctrl"], _ = expandSystemNpuNpuTcamSactMSrhCtrl(d, i["m_srh_ctrl"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["m-srh-ctrl"] = nil
	}
	pre_append = pre + ".0." + "tpe_id_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["tpe-id-v"], _ = expandSystemNpuNpuTcamSactTpeIdV(d, i["tpe_id_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tpe_id"
	if _, ok := d.GetOk(pre_append); ok {
		result["tpe-id"], _ = expandSystemNpuNpuTcamSactTpeId(d, i["tpe_id"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["tpe-id"] = nil
	}
	pre_append = pre + ".0." + "vdom_id_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["vdom-id-v"], _ = expandSystemNpuNpuTcamSactVdomIdV(d, i["vdom_id_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vdom_id"
	if _, ok := d.GetOk(pre_append); ok {
		result["vdom-id"], _ = expandSystemNpuNpuTcamSactVdomId(d, i["vdom_id"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["vdom-id"] = nil
	}
	pre_append = pre + ".0." + "mss_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["mss-v"], _ = expandSystemNpuNpuTcamSactMssV(d, i["mss_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mss"
	if _, ok := d.GetOk(pre_append); ok {
		result["mss"], _ = expandSystemNpuNpuTcamSactMss(d, i["mss"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["mss"] = nil
	}
	pre_append = pre + ".0." + "tp_smchk_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["tp-smchk-v"], _ = expandSystemNpuNpuTcamSactTpSmchkV(d, i["tp_smchk_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tp_smchk"
	if _, ok := d.GetOk(pre_append); ok {
		result["tp_smchk"], _ = expandSystemNpuNpuTcamSactTpSmchk(d, i["tp_smchk"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["tp_smchk"] = nil
	}
	pre_append = pre + ".0." + "etype_pid_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["etype-pid-v"], _ = expandSystemNpuNpuTcamSactEtypePidV(d, i["etype_pid_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "etype_pid"
	if _, ok := d.GetOk(pre_append); ok {
		result["etype-pid"], _ = expandSystemNpuNpuTcamSactEtypePid(d, i["etype_pid"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["etype-pid"] = nil
	}
	pre_append = pre + ".0." + "frag_proc_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["frag-proc-v"], _ = expandSystemNpuNpuTcamSactFragProcV(d, i["frag_proc_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "frag_proc"
	if _, ok := d.GetOk(pre_append); ok {
		result["frag-proc"], _ = expandSystemNpuNpuTcamSactFragProc(d, i["frag_proc"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["frag-proc"] = nil
	}
	pre_append = pre + ".0." + "espff_proc_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["espff-proc-v"], _ = expandSystemNpuNpuTcamSactEspffProcV(d, i["espff_proc_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "espff_proc"
	if _, ok := d.GetOk(pre_append); ok {
		result["espff-proc"], _ = expandSystemNpuNpuTcamSactEspffProc(d, i["espff_proc"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["espff-proc"] = nil
	}
	pre_append = pre + ".0." + "prio_pid_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["prio-pid-v"], _ = expandSystemNpuNpuTcamSactPrioPidV(d, i["prio_pid_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "prio_pid"
	if _, ok := d.GetOk(pre_append); ok {
		result["prio-pid"], _ = expandSystemNpuNpuTcamSactPrioPid(d, i["prio_pid"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["prio-pid"] = nil
	}
	pre_append = pre + ".0." + "igmp_mld_snp_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["igmp-mld-snp-v"], _ = expandSystemNpuNpuTcamSactIgmpMldSnpV(d, i["igmp_mld_snp_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "igmp_mld_snp"
	if _, ok := d.GetOk(pre_append); ok {
		result["igmp-mld-snp"], _ = expandSystemNpuNpuTcamSactIgmpMldSnp(d, i["igmp_mld_snp"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["igmp-mld-snp"] = nil
	}
	pre_append = pre + ".0." + "smac_skip_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["smac-skip-v"], _ = expandSystemNpuNpuTcamSactSmacSkipV(d, i["smac_skip_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "smac_skip"
	if _, ok := d.GetOk(pre_append); ok {
		result["smac-skip"], _ = expandSystemNpuNpuTcamSactSmacSkip(d, i["smac_skip"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["smac-skip"] = nil
	}
	pre_append = pre + ".0." + "dmac_skip_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["dmac-skip-v"], _ = expandSystemNpuNpuTcamSactDmacSkipV(d, i["dmac_skip_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dmac_skip"
	if _, ok := d.GetOk(pre_append); ok {
		result["dmac-skip"], _ = expandSystemNpuNpuTcamSactDmacSkip(d, i["dmac_skip"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["dmac-skip"] = nil
	}

	return result, nil
}

func expandSystemNpuNpuTcamSactFwdLifV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactFwdLif(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactFwdTvidV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactFwdTvid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactDfLifV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactDfLif(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactActV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactAct(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactPleenV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactPleen(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactIcpenV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactIcpen(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactVdmV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactVdm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactLearnV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactLearn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactRfshV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactRfsh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactFwdV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactFwd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactXModeV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactXMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactPromisV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactPromis(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactBmprocV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactBmproc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactMacIdV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactMacId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactDosenV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactDosen(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactDfrV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactDfr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactMSrhCtrlV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactMSrhCtrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactTpeIdV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactTpeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactVdomIdV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactVdomId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactMssV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactMss(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactTpSmchkV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactTpSmchk(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactEtypePidV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactEtypePid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactFragProcV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactFragProc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactEspffProcV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactEspffProc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactPrioPidV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactPrioPid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactIgmpMldSnpV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactIgmpMldSnp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactSmacSkipV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactSmacSkip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactDmacSkipV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamSactDmacSkip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTact(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "act_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["act-v"], _ = expandSystemNpuNpuTcamTactActV(d, i["act_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "act"
	if _, ok := d.GetOk(pre_append); ok {
		result["act"], _ = expandSystemNpuNpuTcamTactAct(d, i["act"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["act"] = nil
	}
	pre_append = pre + ".0." + "mtuv4_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["mtuv4-v"], _ = expandSystemNpuNpuTcamTactMtuv4V(d, i["mtuv4_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mtuv4"
	if _, ok := d.GetOk(pre_append); ok {
		result["mtuv4"], _ = expandSystemNpuNpuTcamTactMtuv4(d, i["mtuv4"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["mtuv4"] = nil
	}
	pre_append = pre + ".0." + "mtuv6_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["mtuv6-v"], _ = expandSystemNpuNpuTcamTactMtuv6V(d, i["mtuv6_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mtuv6"
	if _, ok := d.GetOk(pre_append); ok {
		result["mtuv6"], _ = expandSystemNpuNpuTcamTactMtuv6(d, i["mtuv6"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["mtuv6"] = nil
	}
	pre_append = pre + ".0." + "mac_id_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["mac-id-v"], _ = expandSystemNpuNpuTcamTactMacIdV(d, i["mac_id_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mac_id"
	if _, ok := d.GetOk(pre_append); ok {
		result["mac-id"], _ = expandSystemNpuNpuTcamTactMacId(d, i["mac_id"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["mac-id"] = nil
	}
	pre_append = pre + ".0." + "slif_act_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["slif-act-v"], _ = expandSystemNpuNpuTcamTactSlifActV(d, i["slif_act_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "slif_act"
	if _, ok := d.GetOk(pre_append); ok {
		result["slif-act"], _ = expandSystemNpuNpuTcamTactSlifAct(d, i["slif_act"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["slif-act"] = nil
	}
	pre_append = pre + ".0." + "tlif_act_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["tlif-act-v"], _ = expandSystemNpuNpuTcamTactTlifActV(d, i["tlif_act_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tlif_act"
	if _, ok := d.GetOk(pre_append); ok {
		result["tlif-act"], _ = expandSystemNpuNpuTcamTactTlifAct(d, i["tlif_act"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["tlif-act"] = nil
	}
	pre_append = pre + ".0." + "tgtv_act_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["tgtv-act-v"], _ = expandSystemNpuNpuTcamTactTgtvActV(d, i["tgtv_act_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tgtv_act"
	if _, ok := d.GetOk(pre_append); ok {
		result["tgtv-act"], _ = expandSystemNpuNpuTcamTactTgtvAct(d, i["tgtv_act"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["tgtv-act"] = nil
	}
	pre_append = pre + ".0." + "tpeid_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["tpeid-v"], _ = expandSystemNpuNpuTcamTactTpeidV(d, i["tpeid_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tpeid"
	if _, ok := d.GetOk(pre_append); ok {
		result["tpeid"], _ = expandSystemNpuNpuTcamTactTpeid(d, i["tpeid"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["tpeid"] = nil
	}
	pre_append = pre + ".0." + "v6fe_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["v6fe-v"], _ = expandSystemNpuNpuTcamTactV6FeV(d, i["v6fe_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "v6fe"
	if _, ok := d.GetOk(pre_append); ok {
		result["v6fe"], _ = expandSystemNpuNpuTcamTactV6Fe(d, i["v6fe"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["v6fe"] = nil
	}
	pre_append = pre + ".0." + "xlt_vid_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["xlt-vid-v"], _ = expandSystemNpuNpuTcamTactXltVidV(d, i["xlt_vid_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "xlt_vid"
	if _, ok := d.GetOk(pre_append); ok {
		result["xlt-vid"], _ = expandSystemNpuNpuTcamTactXltVid(d, i["xlt_vid"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["xlt-vid"] = nil
	}
	pre_append = pre + ".0." + "xlt_lif_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["xlt-lif-v"], _ = expandSystemNpuNpuTcamTactXltLifV(d, i["xlt_lif_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "xlt_lif"
	if _, ok := d.GetOk(pre_append); ok {
		result["xlt-lif"], _ = expandSystemNpuNpuTcamTactXltLif(d, i["xlt_lif"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["xlt-lif"] = nil
	}
	pre_append = pre + ".0." + "mss_t_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["mss-t-v"], _ = expandSystemNpuNpuTcamTactMssTV(d, i["mss_t_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mss_t"
	if _, ok := d.GetOk(pre_append); ok {
		result["mss-t"], _ = expandSystemNpuNpuTcamTactMssT(d, i["mss_t"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["mss-t"] = nil
	}
	pre_append = pre + ".0." + "lnkid_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["lnkid-v"], _ = expandSystemNpuNpuTcamTactLnkidV(d, i["lnkid_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "lnkid"
	if _, ok := d.GetOk(pre_append); ok {
		result["lnkid"], _ = expandSystemNpuNpuTcamTactLnkid(d, i["lnkid"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["lnkid"] = nil
	}
	pre_append = pre + ".0." + "sublnkid_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["sublnkid-v"], _ = expandSystemNpuNpuTcamTactSublnkidV(d, i["sublnkid_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sublnkid"
	if _, ok := d.GetOk(pre_append); ok {
		result["sublnkid"], _ = expandSystemNpuNpuTcamTactSublnkid(d, i["sublnkid"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["sublnkid"] = nil
	}
	pre_append = pre + ".0." + "fmtuv4_s_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["fmtuv4-s-v"], _ = expandSystemNpuNpuTcamTactFmtuv4SV(d, i["fmtuv4_s_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "fmtuv4_s"
	if _, ok := d.GetOk(pre_append); ok {
		result["fmtuv4-s"], _ = expandSystemNpuNpuTcamTactFmtuv4S(d, i["fmtuv4_s"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["fmtuv4-s"] = nil
	}
	pre_append = pre + ".0." + "fmtuv6_s_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["fmtuv6-s-v"], _ = expandSystemNpuNpuTcamTactFmtuv6SV(d, i["fmtuv6_s_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "fmtuv6_s"
	if _, ok := d.GetOk(pre_append); ok {
		result["fmtuv6-s"], _ = expandSystemNpuNpuTcamTactFmtuv6S(d, i["fmtuv6_s"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["fmtuv6-s"] = nil
	}
	pre_append = pre + ".0." + "vep_en_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["vep-en-v"], _ = expandSystemNpuNpuTcamTactVepEnV(d, i["vep_en_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vep_en"
	if _, ok := d.GetOk(pre_append); ok {
		result["vep_en"], _ = expandSystemNpuNpuTcamTactVepEn(d, i["vep_en"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["vep_en"] = nil
	}
	pre_append = pre + ".0." + "vep_slid_v"
	if _, ok := d.GetOk(pre_append); ok {
		result["vep-slid-v"], _ = expandSystemNpuNpuTcamTactVepSlidV(d, i["vep_slid_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vep_slid"
	if _, ok := d.GetOk(pre_append); ok {
		result["vep-slid"], _ = expandSystemNpuNpuTcamTactVepSlid(d, i["vep_slid"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["vep-slid"] = nil
	}

	return result, nil
}

func expandSystemNpuNpuTcamTactActV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactAct(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactMtuv4V(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactMtuv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactMtuv6V(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactMtuv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactMacIdV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactMacId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactSlifActV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactSlifAct(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactTlifActV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactTlifAct(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactTgtvActV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactTgtvAct(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactTpeidV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactTpeid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactV6FeV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactV6Fe(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactXltVidV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactXltVid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactXltLifV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactXltLif(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactMssTV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactMssT(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactLnkidV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactLnkid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactSublnkidV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactSublnkid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactFmtuv4SV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactFmtuv4S(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactFmtuv6SV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactFmtuv6S(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactVepEnV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactVepEn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactVepSlidV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpuTcamTactVepSlid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueues(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "profile"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["profile"] = make([]struct{}, 0)
		} else {
			result["profile"], _ = expandSystemNpuNpQueuesProfile(d, i["profile"], pre_append, sv)
		}
	} else {
		result["profile"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ethernet_type"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ethernet-type"] = make([]struct{}, 0)
		} else {
			result["ethernet-type"], _ = expandSystemNpuNpQueuesEthernetType(d, i["ethernet_type"], pre_append, sv)
		}
	} else {
		result["ethernet-type"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ip_protocol"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ip-protocol"] = make([]struct{}, 0)
		} else {
			result["ip-protocol"], _ = expandSystemNpuNpQueuesIpProtocol(d, i["ip_protocol"], pre_append, sv)
		}
	} else {
		result["ip-protocol"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ip_service"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ip-service"] = make([]struct{}, 0)
		} else {
			result["ip-service"], _ = expandSystemNpuNpQueuesIpService(d, i["ip_service"], pre_append, sv)
		}
	} else {
		result["ip-service"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "scheduler"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["scheduler"] = make([]struct{}, 0)
		} else {
			result["scheduler"], _ = expandSystemNpuNpQueuesScheduler(d, i["scheduler"], pre_append, sv)
		}
	} else {
		result["scheduler"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "custom_etype_lookup"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["custom-etype-lookup"] = nil
		} else {
			result["custom-etype-lookup"], _ = expandSystemNpuNpQueuesCustomEtypeLookup(d, i["custom_etype_lookup"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemNpuNpQueuesProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemNpuNpQueuesProfileId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandSystemNpuNpQueuesProfileType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandSystemNpuNpQueuesProfileWeight(d, i["weight"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos0"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cos0"], _ = expandSystemNpuNpQueuesProfileCos0(d, i["cos0"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos1"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cos1"], _ = expandSystemNpuNpQueuesProfileCos1(d, i["cos1"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos2"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cos2"], _ = expandSystemNpuNpQueuesProfileCos2(d, i["cos2"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos3"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cos3"], _ = expandSystemNpuNpQueuesProfileCos3(d, i["cos3"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cos4"], _ = expandSystemNpuNpQueuesProfileCos4(d, i["cos4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos5"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cos5"], _ = expandSystemNpuNpQueuesProfileCos5(d, i["cos5"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cos6"], _ = expandSystemNpuNpQueuesProfileCos6(d, i["cos6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos7"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cos7"], _ = expandSystemNpuNpQueuesProfileCos7(d, i["cos7"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp0"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp0"], _ = expandSystemNpuNpQueuesProfileDscp0(d, i["dscp0"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp1"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp1"], _ = expandSystemNpuNpQueuesProfileDscp1(d, i["dscp1"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp2"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp2"], _ = expandSystemNpuNpQueuesProfileDscp2(d, i["dscp2"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp3"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp3"], _ = expandSystemNpuNpQueuesProfileDscp3(d, i["dscp3"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp4"], _ = expandSystemNpuNpQueuesProfileDscp4(d, i["dscp4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp5"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp5"], _ = expandSystemNpuNpQueuesProfileDscp5(d, i["dscp5"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp6"], _ = expandSystemNpuNpQueuesProfileDscp6(d, i["dscp6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp7"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp7"], _ = expandSystemNpuNpQueuesProfileDscp7(d, i["dscp7"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp8"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp8"], _ = expandSystemNpuNpQueuesProfileDscp8(d, i["dscp8"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp9"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp9"], _ = expandSystemNpuNpQueuesProfileDscp9(d, i["dscp9"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp10"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp10"], _ = expandSystemNpuNpQueuesProfileDscp10(d, i["dscp10"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp11"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp11"], _ = expandSystemNpuNpQueuesProfileDscp11(d, i["dscp11"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp12"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp12"], _ = expandSystemNpuNpQueuesProfileDscp12(d, i["dscp12"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp13"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp13"], _ = expandSystemNpuNpQueuesProfileDscp13(d, i["dscp13"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp14"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp14"], _ = expandSystemNpuNpQueuesProfileDscp14(d, i["dscp14"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp15"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp15"], _ = expandSystemNpuNpQueuesProfileDscp15(d, i["dscp15"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp16"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp16"], _ = expandSystemNpuNpQueuesProfileDscp16(d, i["dscp16"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp17"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp17"], _ = expandSystemNpuNpQueuesProfileDscp17(d, i["dscp17"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp18"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp18"], _ = expandSystemNpuNpQueuesProfileDscp18(d, i["dscp18"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp19"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp19"], _ = expandSystemNpuNpQueuesProfileDscp19(d, i["dscp19"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp20"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp20"], _ = expandSystemNpuNpQueuesProfileDscp20(d, i["dscp20"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp21"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp21"], _ = expandSystemNpuNpQueuesProfileDscp21(d, i["dscp21"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp22"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp22"], _ = expandSystemNpuNpQueuesProfileDscp22(d, i["dscp22"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp23"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp23"], _ = expandSystemNpuNpQueuesProfileDscp23(d, i["dscp23"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp24"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp24"], _ = expandSystemNpuNpQueuesProfileDscp24(d, i["dscp24"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp25"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp25"], _ = expandSystemNpuNpQueuesProfileDscp25(d, i["dscp25"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp26"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp26"], _ = expandSystemNpuNpQueuesProfileDscp26(d, i["dscp26"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp27"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp27"], _ = expandSystemNpuNpQueuesProfileDscp27(d, i["dscp27"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp28"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp28"], _ = expandSystemNpuNpQueuesProfileDscp28(d, i["dscp28"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp29"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp29"], _ = expandSystemNpuNpQueuesProfileDscp29(d, i["dscp29"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp30"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp30"], _ = expandSystemNpuNpQueuesProfileDscp30(d, i["dscp30"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp31"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp31"], _ = expandSystemNpuNpQueuesProfileDscp31(d, i["dscp31"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp32"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp32"], _ = expandSystemNpuNpQueuesProfileDscp32(d, i["dscp32"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp33"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp33"], _ = expandSystemNpuNpQueuesProfileDscp33(d, i["dscp33"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp34"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp34"], _ = expandSystemNpuNpQueuesProfileDscp34(d, i["dscp34"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp35"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp35"], _ = expandSystemNpuNpQueuesProfileDscp35(d, i["dscp35"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp36"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp36"], _ = expandSystemNpuNpQueuesProfileDscp36(d, i["dscp36"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp37"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp37"], _ = expandSystemNpuNpQueuesProfileDscp37(d, i["dscp37"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp38"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp38"], _ = expandSystemNpuNpQueuesProfileDscp38(d, i["dscp38"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp39"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp39"], _ = expandSystemNpuNpQueuesProfileDscp39(d, i["dscp39"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp40"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp40"], _ = expandSystemNpuNpQueuesProfileDscp40(d, i["dscp40"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp41"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp41"], _ = expandSystemNpuNpQueuesProfileDscp41(d, i["dscp41"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp42"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp42"], _ = expandSystemNpuNpQueuesProfileDscp42(d, i["dscp42"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp43"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp43"], _ = expandSystemNpuNpQueuesProfileDscp43(d, i["dscp43"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp44"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp44"], _ = expandSystemNpuNpQueuesProfileDscp44(d, i["dscp44"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp45"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp45"], _ = expandSystemNpuNpQueuesProfileDscp45(d, i["dscp45"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp46"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp46"], _ = expandSystemNpuNpQueuesProfileDscp46(d, i["dscp46"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp47"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp47"], _ = expandSystemNpuNpQueuesProfileDscp47(d, i["dscp47"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp48"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp48"], _ = expandSystemNpuNpQueuesProfileDscp48(d, i["dscp48"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp49"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp49"], _ = expandSystemNpuNpQueuesProfileDscp49(d, i["dscp49"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp50"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp50"], _ = expandSystemNpuNpQueuesProfileDscp50(d, i["dscp50"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp51"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp51"], _ = expandSystemNpuNpQueuesProfileDscp51(d, i["dscp51"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp52"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp52"], _ = expandSystemNpuNpQueuesProfileDscp52(d, i["dscp52"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp53"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp53"], _ = expandSystemNpuNpQueuesProfileDscp53(d, i["dscp53"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp54"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp54"], _ = expandSystemNpuNpQueuesProfileDscp54(d, i["dscp54"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp55"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp55"], _ = expandSystemNpuNpQueuesProfileDscp55(d, i["dscp55"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp56"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp56"], _ = expandSystemNpuNpQueuesProfileDscp56(d, i["dscp56"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp57"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp57"], _ = expandSystemNpuNpQueuesProfileDscp57(d, i["dscp57"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp58"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp58"], _ = expandSystemNpuNpQueuesProfileDscp58(d, i["dscp58"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp59"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp59"], _ = expandSystemNpuNpQueuesProfileDscp59(d, i["dscp59"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp60"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp60"], _ = expandSystemNpuNpQueuesProfileDscp60(d, i["dscp60"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp61"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp61"], _ = expandSystemNpuNpQueuesProfileDscp61(d, i["dscp61"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp62"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp62"], _ = expandSystemNpuNpQueuesProfileDscp62(d, i["dscp62"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp63"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp63"], _ = expandSystemNpuNpQueuesProfileDscp63(d, i["dscp63"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemNpuNpQueuesProfileId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileCos0(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileCos1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileCos2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileCos3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileCos4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileCos5(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileCos6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileCos7(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp0(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp5(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp7(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp8(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp9(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp10(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp11(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp12(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp13(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp14(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp15(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp16(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp17(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp18(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp19(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp20(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp21(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp22(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp23(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp24(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp25(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp26(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp27(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp28(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp29(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp30(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp31(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp32(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp33(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp34(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp35(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp36(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp37(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp38(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp39(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp40(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp41(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp42(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp43(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp44(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp45(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp46(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp47(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp48(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp49(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp50(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp51(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp52(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp53(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp54(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp55(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp56(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp57(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp58(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp59(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp60(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp61(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp62(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp63(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesEthernetType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemNpuNpQueuesEthernetTypeName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandSystemNpuNpQueuesEthernetTypeType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["queue"], _ = expandSystemNpuNpQueuesEthernetTypeQueue(d, i["queue"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["queue"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandSystemNpuNpQueuesEthernetTypeWeight(d, i["weight"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemNpuNpQueuesEthernetTypeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesEthernetTypeType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesEthernetTypeQueue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesEthernetTypeWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemNpuNpQueuesIpProtocolName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandSystemNpuNpQueuesIpProtocolProtocol(d, i["protocol"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["protocol"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["queue"], _ = expandSystemNpuNpQueuesIpProtocolQueue(d, i["queue"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["queue"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandSystemNpuNpQueuesIpProtocolWeight(d, i["weight"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemNpuNpQueuesIpProtocolName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpProtocolProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpProtocolQueue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpProtocolWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemNpuNpQueuesIpServiceName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandSystemNpuNpQueuesIpServiceProtocol(d, i["protocol"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["protocol"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sport"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sport"], _ = expandSystemNpuNpQueuesIpServiceSport(d, i["sport"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["sport"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dport"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dport"], _ = expandSystemNpuNpQueuesIpServiceDport(d, i["dport"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dport"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["queue"], _ = expandSystemNpuNpQueuesIpServiceQueue(d, i["queue"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["queue"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandSystemNpuNpQueuesIpServiceWeight(d, i["weight"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemNpuNpQueuesIpServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpServiceProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpServiceSport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpServiceDport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpServiceQueue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpServiceWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesScheduler(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemNpuNpQueuesSchedulerName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mode"], _ = expandSystemNpuNpQueuesSchedulerMode(d, i["mode"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemNpuNpQueuesSchedulerName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesSchedulerMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesCustomEtypeLookup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuSwNpRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuSwNpRateUnit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuSwNpRateBurst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDoubleLevelMcastOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuMcastDeniedSesOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuIpsecEncSubengineMask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuIpsecDecSubengineMask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNp6CpsOptimizationMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuSwNpBandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuStripEspPadding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuStripClearTextPadding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuIpsecInboundCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuSseBackpressure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuRdpOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuIpsecOverVlink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuUespOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuMcastSessionAccounting(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuIpsecMtuOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuSessionDeniedOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuQtmBufMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuIpsecObNpSel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuMaxReceiveUnit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuLagHashGre(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuUseMseOft(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuIkePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["port"], _ = expandSystemNpuIkePortPort(d, i["port"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemNpuIkePortPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuPriorityProtocol(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bgp"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["bgp"] = nil
		} else {
			result["bgp"], _ = expandSystemNpuPriorityProtocolBgp(d, i["bgp"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "slbc"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["slbc"] = nil
		} else {
			result["slbc"], _ = expandSystemNpuPriorityProtocolSlbc(d, i["slbc"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "bfd"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["bfd"] = nil
		} else {
			result["bfd"], _ = expandSystemNpuPriorityProtocolBfd(d, i["bfd"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemNpuPriorityProtocolBgp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuPriorityProtocolSlbc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuPriorityProtocolBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNpu(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dedicated_management_cpu"); ok {
		if setArgNil {
			obj["dedicated-management-cpu"] = nil
		} else {
			t, err := expandSystemNpuDedicatedManagementCpu(d, v, "dedicated_management_cpu", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dedicated-management-cpu"] = t
			}
		}
	}

	if v, ok := d.GetOk("dedicated_lacp_queue"); ok {
		if setArgNil {
			obj["dedicated-lacp-queue"] = nil
		} else {
			t, err := expandSystemNpuDedicatedLacpQueue(d, v, "dedicated_lacp_queue", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dedicated-lacp-queue"] = t
			}
		}
	}

	if v, ok := d.GetOk("dedicated_management_affinity"); ok {
		if setArgNil {
			obj["dedicated-management-affinity"] = nil
		} else {
			t, err := expandSystemNpuDedicatedManagementAffinity(d, v, "dedicated_management_affinity", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dedicated-management-affinity"] = t
			}
		}
	}

	if v, ok := d.GetOk("dos_options"); ok {
		t, err := expandSystemNpuDosOptions(d, v, "dos_options", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dos-options"] = t
		}
	}

	if v, ok := d.GetOk("policy_offload_level"); ok {
		if setArgNil {
			obj["policy-offload-level"] = nil
		} else {
			t, err := expandSystemNpuPolicyOffloadLevel(d, v, "policy_offload_level", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["policy-offload-level"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("napi_break_interval"); ok {
		if setArgNil {
			obj["napi-break-interval"] = nil
		} else {
			t, err := expandSystemNpuNapiBreakInterval(d, v, "napi_break_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["napi-break-interval"] = t
			}
		}
	} else if d.HasChange("napi_break_interval") {
		obj["napi-break-interval"] = nil
	}

	if v, ok := d.GetOk("hpe"); ok {
		t, err := expandSystemNpuHpe(d, v, "hpe", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hpe"] = t
		}
	}

	if v, ok := d.GetOk("fastpath"); ok {
		if setArgNil {
			obj["fastpath"] = nil
		} else {
			t, err := expandSystemNpuFastpath(d, v, "fastpath", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fastpath"] = t
			}
		}
	}

	if v, ok := d.GetOk("capwap_offload"); ok {
		if setArgNil {
			obj["capwap-offload"] = nil
		} else {
			t, err := expandSystemNpuCapwapOffload(d, v, "capwap_offload", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["capwap-offload"] = t
			}
		}
	}

	if v, ok := d.GetOk("vxlan_offload"); ok {
		if setArgNil {
			obj["vxlan-offload"] = nil
		} else {
			t, err := expandSystemNpuVxlanOffload(d, v, "vxlan_offload", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vxlan-offload"] = t
			}
		}
	}

	if v, ok := d.GetOk("vxlan_mac_flapping_guard"); ok {
		if setArgNil {
			obj["vxlan-mac-flapping-guard"] = nil
		} else {
			t, err := expandSystemNpuVxlanMacFlappingGuard(d, v, "vxlan_mac_flapping_guard", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vxlan-mac-flapping-guard"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_qos_type"); ok {
		if setArgNil {
			obj["default-qos-type"] = nil
		} else {
			t, err := expandSystemNpuDefaultQosType(d, v, "default_qos_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-qos-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_ipsec_mcs_type"); ok {
		if setArgNil {
			obj["default-ipsec-mcs-type"] = nil
		} else {
			t, err := expandSystemNpuDefaultIpsecMcsType(d, v, "default_ipsec_mcs_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-ipsec-mcs-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("shaping_stats"); ok {
		if setArgNil {
			obj["shaping-stats"] = nil
		} else {
			t, err := expandSystemNpuShapingStats(d, v, "shaping_stats", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["shaping-stats"] = t
			}
		}
	}

	if v, ok := d.GetOk("mcs_host_packet_tpe_shaping"); ok {
		if setArgNil {
			obj["mcs-host-packet-tpe-shaping"] = nil
		} else {
			t, err := expandSystemNpuMcsHostPacketTpeShaping(d, v, "mcs_host_packet_tpe_shaping", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mcs-host-packet-tpe-shaping"] = t
			}
		}
	}

	if v, ok := d.GetOk("gtp_support"); ok {
		if setArgNil {
			obj["gtp-support"] = nil
		} else {
			t, err := expandSystemNpuGtpSupport(d, v, "gtp_support", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["gtp-support"] = t
			}
		}
	}

	if v, ok := d.GetOk("per_session_accounting"); ok {
		if setArgNil {
			obj["per-session-accounting"] = nil
		} else {
			t, err := expandSystemNpuPerSessionAccounting(d, v, "per_session_accounting", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["per-session-accounting"] = t
			}
		}
	}

	if v, ok := d.GetOk("session_acct_interval"); ok {
		if setArgNil {
			obj["session-acct-interval"] = nil
		} else {
			t, err := expandSystemNpuSessionAcctInterval(d, v, "session_acct_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["session-acct-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("max_session_timeout"); ok {
		if setArgNil {
			obj["max-session-timeout"] = nil
		} else {
			t, err := expandSystemNpuMaxSessionTimeout(d, v, "max_session_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-session-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("fp_anomaly"); ok {
		t, err := expandSystemNpuFpAnomaly(d, v, "fp_anomaly", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fp-anomaly"] = t
		}
	}

	if v, ok := d.GetOk("ip_reassembly"); ok {
		t, err := expandSystemNpuIpReassembly(d, v, "ip_reassembly", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-reassembly"] = t
		}
	}

	if v, ok := d.GetOk("hash_tbl_spread"); ok {
		if setArgNil {
			obj["hash-tbl-spread"] = nil
		} else {
			t, err := expandSystemNpuHashTblSpread(d, v, "hash_tbl_spread", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hash-tbl-spread"] = t
			}
		}
	}

	if v, ok := d.GetOk("vlan_lookup_cache"); ok {
		if setArgNil {
			obj["vlan-lookup-cache"] = nil
		} else {
			t, err := expandSystemNpuVlanLookupCache(d, v, "vlan_lookup_cache", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vlan-lookup-cache"] = t
			}
		}
	}

	if v, ok := d.GetOk("ip_fragment_offload"); ok {
		if setArgNil {
			obj["ip-fragment-offload"] = nil
		} else {
			t, err := expandSystemNpuIpFragmentOffload(d, v, "ip_fragment_offload", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip-fragment-offload"] = t
			}
		}
	}

	if v, ok := d.GetOk("htx_icmp_csum_chk"); ok {
		if setArgNil {
			obj["htx-icmp-csum-chk"] = nil
		} else {
			t, err := expandSystemNpuHtxIcmpCsumChk(d, v, "htx_icmp_csum_chk", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["htx-icmp-csum-chk"] = t
			}
		}
	}

	if v, ok := d.GetOk("htab_msg_queue"); ok {
		if setArgNil {
			obj["htab-msg-queue"] = nil
		} else {
			t, err := expandSystemNpuHtabMsgQueue(d, v, "htab_msg_queue", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["htab-msg-queue"] = t
			}
		}
	}

	if v, ok := d.GetOk("htab_dedi_queue_nr"); ok {
		if setArgNil {
			obj["htab-dedi-queue-nr"] = nil
		} else {
			t, err := expandSystemNpuHtabDediQueueNr(d, v, "htab_dedi_queue_nr", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["htab-dedi-queue-nr"] = t
			}
		}
	}

	if v, ok := d.GetOk("dsw_dts_profile"); ok || d.HasChange("dsw_dts_profile") {
		if setArgNil {
			obj["dsw-dts-profile"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemNpuDswDtsProfile(d, v, "dsw_dts_profile", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dsw-dts-profile"] = t
			}
		}
	}

	if v, ok := d.GetOk("dsw_queue_dts_profile"); ok || d.HasChange("dsw_queue_dts_profile") {
		if setArgNil {
			obj["dsw-queue-dts-profile"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemNpuDswQueueDtsProfile(d, v, "dsw_queue_dts_profile", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dsw-queue-dts-profile"] = t
			}
		}
	}

	if v, ok := d.GetOk("npu_tcam"); ok || d.HasChange("npu_tcam") {
		if setArgNil {
			obj["npu-tcam"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemNpuNpuTcam(d, v, "npu_tcam", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["npu-tcam"] = t
			}
		}
	}

	if v, ok := d.GetOk("np_queues"); ok {
		t, err := expandSystemNpuNpQueues(d, v, "np_queues", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["np-queues"] = t
		}
	}

	if v, ok := d.GetOkExists("sw_np_rate"); ok {
		if setArgNil {
			obj["sw-np-rate"] = nil
		} else {
			t, err := expandSystemNpuSwNpRate(d, v, "sw_np_rate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sw-np-rate"] = t
			}
		}
	} else if d.HasChange("sw_np_rate") {
		obj["sw-np-rate"] = nil
	}

	if v, ok := d.GetOk("sw_np_rate_unit"); ok {
		if setArgNil {
			obj["sw-np-rate-unit"] = nil
		} else {
			t, err := expandSystemNpuSwNpRateUnit(d, v, "sw_np_rate_unit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sw-np-rate-unit"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("sw_np_rate_burst"); ok {
		if setArgNil {
			obj["sw-np-rate-burst"] = nil
		} else {
			t, err := expandSystemNpuSwNpRateBurst(d, v, "sw_np_rate_burst", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sw-np-rate-burst"] = t
			}
		}
	}

	if v, ok := d.GetOk("double_level_mcast_offload"); ok {
		if setArgNil {
			obj["double-level-mcast-offload"] = nil
		} else {
			t, err := expandSystemNpuDoubleLevelMcastOffload(d, v, "double_level_mcast_offload", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["double-level-mcast-offload"] = t
			}
		}
	}

	if v, ok := d.GetOk("mcast_denied_ses_offload"); ok {
		if setArgNil {
			obj["mcast-denied-ses-offload"] = nil
		} else {
			t, err := expandSystemNpuMcastDeniedSesOffload(d, v, "mcast_denied_ses_offload", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mcast-denied-ses-offload"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipsec_enc_subengine_mask"); ok {
		if setArgNil {
			obj["ipsec-enc-subengine-mask"] = nil
		} else {
			t, err := expandSystemNpuIpsecEncSubengineMask(d, v, "ipsec_enc_subengine_mask", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipsec-enc-subengine-mask"] = t
			}
		}
	} else if d.HasChange("ipsec_enc_subengine_mask") {
		obj["ipsec-enc-subengine-mask"] = nil
	}

	if v, ok := d.GetOk("ipsec_dec_subengine_mask"); ok {
		if setArgNil {
			obj["ipsec-dec-subengine-mask"] = nil
		} else {
			t, err := expandSystemNpuIpsecDecSubengineMask(d, v, "ipsec_dec_subengine_mask", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipsec-dec-subengine-mask"] = t
			}
		}
	} else if d.HasChange("ipsec_dec_subengine_mask") {
		obj["ipsec-dec-subengine-mask"] = nil
	}

	if v, ok := d.GetOk("np6_cps_optimization_mode"); ok {
		if setArgNil {
			obj["np6-cps-optimization-mode"] = nil
		} else {
			t, err := expandSystemNpuNp6CpsOptimizationMode(d, v, "np6_cps_optimization_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["np6-cps-optimization-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("sw_np_bandwidth"); ok {
		if setArgNil {
			obj["sw-np-bandwidth"] = nil
		} else {
			t, err := expandSystemNpuSwNpBandwidth(d, v, "sw_np_bandwidth", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sw-np-bandwidth"] = t
			}
		}
	}

	if v, ok := d.GetOk("strip_esp_padding"); ok {
		if setArgNil {
			obj["strip-esp-padding"] = nil
		} else {
			t, err := expandSystemNpuStripEspPadding(d, v, "strip_esp_padding", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["strip-esp-padding"] = t
			}
		}
	}

	if v, ok := d.GetOk("strip_clear_text_padding"); ok {
		if setArgNil {
			obj["strip-clear-text-padding"] = nil
		} else {
			t, err := expandSystemNpuStripClearTextPadding(d, v, "strip_clear_text_padding", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["strip-clear-text-padding"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipsec_inbound_cache"); ok {
		if setArgNil {
			obj["ipsec-inbound-cache"] = nil
		} else {
			t, err := expandSystemNpuIpsecInboundCache(d, v, "ipsec_inbound_cache", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipsec-inbound-cache"] = t
			}
		}
	}

	if v, ok := d.GetOk("sse_backpressure"); ok {
		if setArgNil {
			obj["sse-backpressure"] = nil
		} else {
			t, err := expandSystemNpuSseBackpressure(d, v, "sse_backpressure", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sse-backpressure"] = t
			}
		}
	}

	if v, ok := d.GetOk("rdp_offload"); ok {
		if setArgNil {
			obj["rdp-offload"] = nil
		} else {
			t, err := expandSystemNpuRdpOffload(d, v, "rdp_offload", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["rdp-offload"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipsec_over_vlink"); ok {
		if setArgNil {
			obj["ipsec-over-vlink"] = nil
		} else {
			t, err := expandSystemNpuIpsecOverVlink(d, v, "ipsec_over_vlink", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipsec-over-vlink"] = t
			}
		}
	}

	if v, ok := d.GetOk("uesp_offload"); ok {
		if setArgNil {
			obj["uesp-offload"] = nil
		} else {
			t, err := expandSystemNpuUespOffload(d, v, "uesp_offload", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uesp-offload"] = t
			}
		}
	}

	if v, ok := d.GetOk("mcast_session_accounting"); ok {
		if setArgNil {
			obj["mcast-session-accounting"] = nil
		} else {
			t, err := expandSystemNpuMcastSessionAccounting(d, v, "mcast_session_accounting", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mcast-session-accounting"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipsec_mtu_override"); ok {
		if setArgNil {
			obj["ipsec-mtu-override"] = nil
		} else {
			t, err := expandSystemNpuIpsecMtuOverride(d, v, "ipsec_mtu_override", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipsec-mtu-override"] = t
			}
		}
	}

	if v, ok := d.GetOk("session_denied_offload"); ok {
		if setArgNil {
			obj["session-denied-offload"] = nil
		} else {
			t, err := expandSystemNpuSessionDeniedOffload(d, v, "session_denied_offload", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["session-denied-offload"] = t
			}
		}
	}

	if v, ok := d.GetOk("qtm_buf_mode"); ok {
		if setArgNil {
			obj["qtm-buf-mode"] = nil
		} else {
			t, err := expandSystemNpuQtmBufMode(d, v, "qtm_buf_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["qtm-buf-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipsec_ob_np_sel"); ok {
		if setArgNil {
			obj["ipsec-ob-np-sel"] = nil
		} else {
			t, err := expandSystemNpuIpsecObNpSel(d, v, "ipsec_ob_np_sel", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipsec-ob-np-sel"] = t
			}
		}
	}

	if v, ok := d.GetOk("max_receive_unit"); ok {
		if setArgNil {
			obj["max-receive-unit"] = nil
		} else {
			t, err := expandSystemNpuMaxReceiveUnit(d, v, "max_receive_unit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-receive-unit"] = t
			}
		}
	}

	if v, ok := d.GetOk("lag_hash_gre"); ok {
		if setArgNil {
			obj["lag-hash-gre"] = nil
		} else {
			t, err := expandSystemNpuLagHashGre(d, v, "lag_hash_gre", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["lag-hash-gre"] = t
			}
		}
	}

	if v, ok := d.GetOk("use_mse_oft"); ok {
		if setArgNil {
			obj["use-mse-oft"] = nil
		} else {
			t, err := expandSystemNpuUseMseOft(d, v, "use_mse_oft", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["use-mse-oft"] = t
			}
		}
	}

	if v, ok := d.GetOk("ike_port"); ok || d.HasChange("ike_port") {
		if setArgNil {
			obj["ike-port"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemNpuIkePort(d, v, "ike_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ike-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("priority_protocol"); ok {
		t, err := expandSystemNpuPriorityProtocol(d, v, "priority_protocol", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-protocol"] = t
		}
	}

	return &obj, nil
}
