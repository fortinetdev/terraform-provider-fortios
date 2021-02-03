// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure protocol options.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceFirewallProfileProtocolOptions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallProfileProtocolOptionsRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"feature_set": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"oversize_log": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"switching_protocols_log": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"comfort_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"comfort_amount": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"range_block": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"http_policy": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"strip_x_forwarded_for": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"post_lang": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"fortinet_bar": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"fortinet_bar_port": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"streaming_content_bypass": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"switching_protocols": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"unknown_http_version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tunnel_non_http": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"stream_based_uncompressed_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"block_page_status_code": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"retry_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tcp_window_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tcp_window_minimum": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tcp_window_maximum": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tcp_window_size": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ssl_offloaded": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ftp": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"comfort_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"comfort_amount": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"stream_based_uncompressed_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tcp_window_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tcp_window_minimum": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tcp_window_maximum": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tcp_window_size": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ssl_offloaded": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"imap": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ssl_offloaded": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"mapi": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"pop3": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ssl_offloaded": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"smtp": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"server_busy": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ssl_offloaded": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"nntp": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ssh": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"options": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"comfort_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"comfort_amount": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"stream_based_uncompressed_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tcp_window_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tcp_window_minimum": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tcp_window_maximum": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tcp_window_size": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ssl_offloaded": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"dns": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"cifs": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tcp_window_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tcp_window_minimum": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tcp_window_maximum": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tcp_window_size": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"server_credential_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"domain_controller": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"server_keytab": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"principal": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"keytab": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"mail_signature": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"signature": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"rpc_over_http": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceFirewallProfileProtocolOptionsRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing FirewallProfileProtocolOptions: type error")
	}

	o, err := c.ReadFirewallProfileProtocolOptions(mkey)
	if err != nil {
		return fmt.Errorf("Error describing FirewallProfileProtocolOptions: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectFirewallProfileProtocolOptions(d, o)
	if err != nil {
		return fmt.Errorf("Error describing FirewallProfileProtocolOptions from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenFirewallProfileProtocolOptionsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsFeatureSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsOversizeLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSwitchingProtocolsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "comfort_interval"
	if _, ok := i["comfort-interval"]; ok {
		result["comfort_interval"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpComfortInterval(i["comfort-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "comfort_amount"
	if _, ok := i["comfort-amount"]; ok {
		result["comfort_amount"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpComfortAmount(i["comfort-amount"], d, pre_append)
	}

	pre_append = pre + ".0." + "range_block"
	if _, ok := i["range-block"]; ok {
		result["range_block"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpRangeBlock(i["range-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "http_policy"
	if _, ok := i["http-policy"]; ok {
		result["http_policy"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpHttpPolicy(i["http-policy"], d, pre_append)
	}

	pre_append = pre + ".0." + "strip_x_forwarded_for"
	if _, ok := i["strip-x-forwarded-for"]; ok {
		result["strip_x_forwarded_for"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpStripXForwardedFor(i["strip-x-forwarded-for"], d, pre_append)
	}

	pre_append = pre + ".0." + "post_lang"
	if _, ok := i["post-lang"]; ok {
		result["post_lang"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpPostLang(i["post-lang"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortinet_bar"
	if _, ok := i["fortinet-bar"]; ok {
		result["fortinet_bar"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpFortinetBar(i["fortinet-bar"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortinet_bar_port"
	if _, ok := i["fortinet-bar-port"]; ok {
		result["fortinet_bar_port"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpFortinetBarPort(i["fortinet-bar-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "streaming_content_bypass"
	if _, ok := i["streaming-content-bypass"]; ok {
		result["streaming_content_bypass"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpStreamingContentBypass(i["streaming-content-bypass"], d, pre_append)
	}

	pre_append = pre + ".0." + "switching_protocols"
	if _, ok := i["switching-protocols"]; ok {
		result["switching_protocols"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpSwitchingProtocols(i["switching-protocols"], d, pre_append)
	}

	pre_append = pre + ".0." + "unknown_http_version"
	if _, ok := i["unknown-http-version"]; ok {
		result["unknown_http_version"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpUnknownHttpVersion(i["unknown-http-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "tunnel_non_http"
	if _, ok := i["tunnel-non-http"]; ok {
		result["tunnel_non_http"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpTunnelNonHttp(i["tunnel-non-http"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "stream_based_uncompressed_limit"
	if _, ok := i["stream-based-uncompressed-limit"]; ok {
		result["stream_based_uncompressed_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpStreamBasedUncompressedLimit(i["stream-based-uncompressed-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_page_status_code"
	if _, ok := i["block-page-status-code"]; ok {
		result["block_page_status_code"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpBlockPageStatusCode(i["block-page-status-code"], d, pre_append)
	}

	pre_append = pre + ".0." + "retry_count"
	if _, ok := i["retry-count"]; ok {
		result["retry_count"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpRetryCount(i["retry-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := i["tcp-window-type"]; ok {
		result["tcp_window_type"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpTcpWindowType(i["tcp-window-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := i["tcp-window-minimum"]; ok {
		result["tcp_window_minimum"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpTcpWindowMinimum(i["tcp-window-minimum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := i["tcp-window-maximum"]; ok {
		result["tcp_window_maximum"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpTcpWindowMaximum(i["tcp-window-maximum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := i["tcp-window-size"]; ok {
		result["tcp_window_size"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpTcpWindowSize(i["tcp-window-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := i["ssl-offloaded"]; ok {
		result["ssl_offloaded"] = dataSourceFlattenFirewallProfileProtocolOptionsHttpSslOffloaded(i["ssl-offloaded"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpComfortInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpComfortAmount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpRangeBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpHttpPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpStripXForwardedFor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpPostLang(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpFortinetBar(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpFortinetBarPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpStreamingContentBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpSwitchingProtocols(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpUnknownHttpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpTunnelNonHttp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpStreamBasedUncompressedLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpBlockPageStatusCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpRetryCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpTcpWindowType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpTcpWindowMinimum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpTcpWindowMaximum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpTcpWindowSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsHttpSslOffloaded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsFtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = dataSourceFlattenFirewallProfileProtocolOptionsFtpPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = dataSourceFlattenFirewallProfileProtocolOptionsFtpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = dataSourceFlattenFirewallProfileProtocolOptionsFtpInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = dataSourceFlattenFirewallProfileProtocolOptionsFtpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "comfort_interval"
	if _, ok := i["comfort-interval"]; ok {
		result["comfort_interval"] = dataSourceFlattenFirewallProfileProtocolOptionsFtpComfortInterval(i["comfort-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "comfort_amount"
	if _, ok := i["comfort-amount"]; ok {
		result["comfort_amount"] = dataSourceFlattenFirewallProfileProtocolOptionsFtpComfortAmount(i["comfort-amount"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsFtpOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsFtpUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsFtpUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "stream_based_uncompressed_limit"
	if _, ok := i["stream-based-uncompressed-limit"]; ok {
		result["stream_based_uncompressed_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsFtpStreamBasedUncompressedLimit(i["stream-based-uncompressed-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = dataSourceFlattenFirewallProfileProtocolOptionsFtpScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := i["tcp-window-type"]; ok {
		result["tcp_window_type"] = dataSourceFlattenFirewallProfileProtocolOptionsFtpTcpWindowType(i["tcp-window-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := i["tcp-window-minimum"]; ok {
		result["tcp_window_minimum"] = dataSourceFlattenFirewallProfileProtocolOptionsFtpTcpWindowMinimum(i["tcp-window-minimum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := i["tcp-window-maximum"]; ok {
		result["tcp_window_maximum"] = dataSourceFlattenFirewallProfileProtocolOptionsFtpTcpWindowMaximum(i["tcp-window-maximum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := i["tcp-window-size"]; ok {
		result["tcp_window_size"] = dataSourceFlattenFirewallProfileProtocolOptionsFtpTcpWindowSize(i["tcp-window-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := i["ssl-offloaded"]; ok {
		result["ssl_offloaded"] = dataSourceFlattenFirewallProfileProtocolOptionsFtpSslOffloaded(i["ssl-offloaded"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenFirewallProfileProtocolOptionsFtpPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func dataSourceFlattenFirewallProfileProtocolOptionsFtpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsFtpInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsFtpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsFtpComfortInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsFtpComfortAmount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsFtpOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsFtpUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsFtpUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsFtpStreamBasedUncompressedLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsFtpScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsFtpTcpWindowType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsFtpTcpWindowMinimum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsFtpTcpWindowMaximum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsFtpTcpWindowSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsFtpSslOffloaded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsImap(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = dataSourceFlattenFirewallProfileProtocolOptionsImapPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = dataSourceFlattenFirewallProfileProtocolOptionsImapStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = dataSourceFlattenFirewallProfileProtocolOptionsImapInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = dataSourceFlattenFirewallProfileProtocolOptionsImapProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = dataSourceFlattenFirewallProfileProtocolOptionsImapOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsImapOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsImapUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsImapUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = dataSourceFlattenFirewallProfileProtocolOptionsImapScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := i["ssl-offloaded"]; ok {
		result["ssl_offloaded"] = dataSourceFlattenFirewallProfileProtocolOptionsImapSslOffloaded(i["ssl-offloaded"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenFirewallProfileProtocolOptionsImapPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func dataSourceFlattenFirewallProfileProtocolOptionsImapStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsImapInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsImapProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsImapOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsImapOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsImapUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsImapUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsImapScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsImapSslOffloaded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsMapi(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = dataSourceFlattenFirewallProfileProtocolOptionsMapiPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = dataSourceFlattenFirewallProfileProtocolOptionsMapiStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = dataSourceFlattenFirewallProfileProtocolOptionsMapiOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsMapiOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsMapiUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsMapiUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = dataSourceFlattenFirewallProfileProtocolOptionsMapiScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenFirewallProfileProtocolOptionsMapiPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func dataSourceFlattenFirewallProfileProtocolOptionsMapiStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsMapiOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsMapiOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsMapiUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsMapiUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsMapiScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsPop3(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = dataSourceFlattenFirewallProfileProtocolOptionsPop3Ports(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = dataSourceFlattenFirewallProfileProtocolOptionsPop3Status(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = dataSourceFlattenFirewallProfileProtocolOptionsPop3InspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = dataSourceFlattenFirewallProfileProtocolOptionsPop3ProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = dataSourceFlattenFirewallProfileProtocolOptionsPop3Options(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsPop3OversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsPop3UncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsPop3UncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = dataSourceFlattenFirewallProfileProtocolOptionsPop3ScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := i["ssl-offloaded"]; ok {
		result["ssl_offloaded"] = dataSourceFlattenFirewallProfileProtocolOptionsPop3SslOffloaded(i["ssl-offloaded"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenFirewallProfileProtocolOptionsPop3Ports(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func dataSourceFlattenFirewallProfileProtocolOptionsPop3Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsPop3InspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsPop3ProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsPop3Options(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsPop3OversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsPop3UncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsPop3UncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsPop3ScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsPop3SslOffloaded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSmtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = dataSourceFlattenFirewallProfileProtocolOptionsSmtpPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = dataSourceFlattenFirewallProfileProtocolOptionsSmtpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = dataSourceFlattenFirewallProfileProtocolOptionsSmtpInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = dataSourceFlattenFirewallProfileProtocolOptionsSmtpProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = dataSourceFlattenFirewallProfileProtocolOptionsSmtpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsSmtpOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsSmtpUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsSmtpUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = dataSourceFlattenFirewallProfileProtocolOptionsSmtpScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_busy"
	if _, ok := i["server-busy"]; ok {
		result["server_busy"] = dataSourceFlattenFirewallProfileProtocolOptionsSmtpServerBusy(i["server-busy"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := i["ssl-offloaded"]; ok {
		result["ssl_offloaded"] = dataSourceFlattenFirewallProfileProtocolOptionsSmtpSslOffloaded(i["ssl-offloaded"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenFirewallProfileProtocolOptionsSmtpPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func dataSourceFlattenFirewallProfileProtocolOptionsSmtpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSmtpInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSmtpProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSmtpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSmtpOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSmtpUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSmtpUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSmtpScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSmtpServerBusy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSmtpSslOffloaded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsNntp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = dataSourceFlattenFirewallProfileProtocolOptionsNntpPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = dataSourceFlattenFirewallProfileProtocolOptionsNntpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = dataSourceFlattenFirewallProfileProtocolOptionsNntpInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = dataSourceFlattenFirewallProfileProtocolOptionsNntpProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = dataSourceFlattenFirewallProfileProtocolOptionsNntpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsNntpOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsNntpUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsNntpUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = dataSourceFlattenFirewallProfileProtocolOptionsNntpScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenFirewallProfileProtocolOptionsNntpPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func dataSourceFlattenFirewallProfileProtocolOptionsNntpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsNntpInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsNntpProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsNntpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsNntpOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsNntpUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsNntpUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsNntpScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSsh(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = dataSourceFlattenFirewallProfileProtocolOptionsSshOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "comfort_interval"
	if _, ok := i["comfort-interval"]; ok {
		result["comfort_interval"] = dataSourceFlattenFirewallProfileProtocolOptionsSshComfortInterval(i["comfort-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "comfort_amount"
	if _, ok := i["comfort-amount"]; ok {
		result["comfort_amount"] = dataSourceFlattenFirewallProfileProtocolOptionsSshComfortAmount(i["comfort-amount"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsSshOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsSshUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsSshUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "stream_based_uncompressed_limit"
	if _, ok := i["stream-based-uncompressed-limit"]; ok {
		result["stream_based_uncompressed_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsSshStreamBasedUncompressedLimit(i["stream-based-uncompressed-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = dataSourceFlattenFirewallProfileProtocolOptionsSshScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := i["tcp-window-type"]; ok {
		result["tcp_window_type"] = dataSourceFlattenFirewallProfileProtocolOptionsSshTcpWindowType(i["tcp-window-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := i["tcp-window-minimum"]; ok {
		result["tcp_window_minimum"] = dataSourceFlattenFirewallProfileProtocolOptionsSshTcpWindowMinimum(i["tcp-window-minimum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := i["tcp-window-maximum"]; ok {
		result["tcp_window_maximum"] = dataSourceFlattenFirewallProfileProtocolOptionsSshTcpWindowMaximum(i["tcp-window-maximum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := i["tcp-window-size"]; ok {
		result["tcp_window_size"] = dataSourceFlattenFirewallProfileProtocolOptionsSshTcpWindowSize(i["tcp-window-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_offloaded"
	if _, ok := i["ssl-offloaded"]; ok {
		result["ssl_offloaded"] = dataSourceFlattenFirewallProfileProtocolOptionsSshSslOffloaded(i["ssl-offloaded"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenFirewallProfileProtocolOptionsSshOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSshComfortInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSshComfortAmount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSshOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSshUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSshUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSshStreamBasedUncompressedLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSshScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSshTcpWindowType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSshTcpWindowMinimum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSshTcpWindowMaximum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSshTcpWindowSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsSshSslOffloaded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsDns(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = dataSourceFlattenFirewallProfileProtocolOptionsDnsPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = dataSourceFlattenFirewallProfileProtocolOptionsDnsStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenFirewallProfileProtocolOptionsDnsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func dataSourceFlattenFirewallProfileProtocolOptionsDnsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsCifs(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = dataSourceFlattenFirewallProfileProtocolOptionsCifsPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = dataSourceFlattenFirewallProfileProtocolOptionsCifsStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = dataSourceFlattenFirewallProfileProtocolOptionsCifsOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsCifsOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsCifsUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = dataSourceFlattenFirewallProfileProtocolOptionsCifsUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = dataSourceFlattenFirewallProfileProtocolOptionsCifsScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_type"
	if _, ok := i["tcp-window-type"]; ok {
		result["tcp_window_type"] = dataSourceFlattenFirewallProfileProtocolOptionsCifsTcpWindowType(i["tcp-window-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_minimum"
	if _, ok := i["tcp-window-minimum"]; ok {
		result["tcp_window_minimum"] = dataSourceFlattenFirewallProfileProtocolOptionsCifsTcpWindowMinimum(i["tcp-window-minimum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_maximum"
	if _, ok := i["tcp-window-maximum"]; ok {
		result["tcp_window_maximum"] = dataSourceFlattenFirewallProfileProtocolOptionsCifsTcpWindowMaximum(i["tcp-window-maximum"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_window_size"
	if _, ok := i["tcp-window-size"]; ok {
		result["tcp_window_size"] = dataSourceFlattenFirewallProfileProtocolOptionsCifsTcpWindowSize(i["tcp-window-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_credential_type"
	if _, ok := i["server-credential-type"]; ok {
		result["server_credential_type"] = dataSourceFlattenFirewallProfileProtocolOptionsCifsServerCredentialType(i["server-credential-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "domain_controller"
	if _, ok := i["domain-controller"]; ok {
		result["domain_controller"] = dataSourceFlattenFirewallProfileProtocolOptionsCifsDomainController(i["domain-controller"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_keytab"
	if _, ok := i["server-keytab"]; ok {
		result["server_keytab"] = dataSourceFlattenFirewallProfileProtocolOptionsCifsServerKeytab(i["server-keytab"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenFirewallProfileProtocolOptionsCifsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func dataSourceFlattenFirewallProfileProtocolOptionsCifsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsCifsOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsCifsOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsCifsUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsCifsUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsCifsScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsCifsTcpWindowType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsCifsTcpWindowMinimum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsCifsTcpWindowMaximum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsCifsTcpWindowSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsCifsServerCredentialType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsCifsDomainController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsCifsServerKeytab(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "principal"
		if _, ok := i["principal"]; ok {
			tmp["principal"] = dataSourceFlattenFirewallProfileProtocolOptionsCifsServerKeytabPrincipal(i["principal"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keytab"
		if _, ok := i["keytab"]; ok {
			tmp["keytab"] = dataSourceFlattenFirewallProfileProtocolOptionsCifsServerKeytabKeytab(i["keytab"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallProfileProtocolOptionsCifsServerKeytabPrincipal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsCifsServerKeytabKeytab(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsMailSignature(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = dataSourceFlattenFirewallProfileProtocolOptionsMailSignatureStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "signature"
	if _, ok := i["signature"]; ok {
		result["signature"] = dataSourceFlattenFirewallProfileProtocolOptionsMailSignatureSignature(i["signature"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenFirewallProfileProtocolOptionsMailSignatureStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsMailSignatureSignature(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProfileProtocolOptionsRpcOverHttp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectFirewallProfileProtocolOptions(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenFirewallProfileProtocolOptionsName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", dataSourceFlattenFirewallProfileProtocolOptionsComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("feature_set", dataSourceFlattenFirewallProfileProtocolOptionsFeatureSet(o["feature-set"], d, "feature_set")); err != nil {
		if !fortiAPIPatch(o["feature-set"]) {
			return fmt.Errorf("Error reading feature_set: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", dataSourceFlattenFirewallProfileProtocolOptionsReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if !fortiAPIPatch(o["replacemsg-group"]) {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("oversize_log", dataSourceFlattenFirewallProfileProtocolOptionsOversizeLog(o["oversize-log"], d, "oversize_log")); err != nil {
		if !fortiAPIPatch(o["oversize-log"]) {
			return fmt.Errorf("Error reading oversize_log: %v", err)
		}
	}

	if err = d.Set("switching_protocols_log", dataSourceFlattenFirewallProfileProtocolOptionsSwitchingProtocolsLog(o["switching-protocols-log"], d, "switching_protocols_log")); err != nil {
		if !fortiAPIPatch(o["switching-protocols-log"]) {
			return fmt.Errorf("Error reading switching_protocols_log: %v", err)
		}
	}

	if err = d.Set("http", dataSourceFlattenFirewallProfileProtocolOptionsHttp(o["http"], d, "http")); err != nil {
		if !fortiAPIPatch(o["http"]) {
			return fmt.Errorf("Error reading http: %v", err)
		}
	}

	if err = d.Set("ftp", dataSourceFlattenFirewallProfileProtocolOptionsFtp(o["ftp"], d, "ftp")); err != nil {
		if !fortiAPIPatch(o["ftp"]) {
			return fmt.Errorf("Error reading ftp: %v", err)
		}
	}

	if err = d.Set("imap", dataSourceFlattenFirewallProfileProtocolOptionsImap(o["imap"], d, "imap")); err != nil {
		if !fortiAPIPatch(o["imap"]) {
			return fmt.Errorf("Error reading imap: %v", err)
		}
	}

	if err = d.Set("mapi", dataSourceFlattenFirewallProfileProtocolOptionsMapi(o["mapi"], d, "mapi")); err != nil {
		if !fortiAPIPatch(o["mapi"]) {
			return fmt.Errorf("Error reading mapi: %v", err)
		}
	}

	if err = d.Set("pop3", dataSourceFlattenFirewallProfileProtocolOptionsPop3(o["pop3"], d, "pop3")); err != nil {
		if !fortiAPIPatch(o["pop3"]) {
			return fmt.Errorf("Error reading pop3: %v", err)
		}
	}

	if err = d.Set("smtp", dataSourceFlattenFirewallProfileProtocolOptionsSmtp(o["smtp"], d, "smtp")); err != nil {
		if !fortiAPIPatch(o["smtp"]) {
			return fmt.Errorf("Error reading smtp: %v", err)
		}
	}

	if err = d.Set("nntp", dataSourceFlattenFirewallProfileProtocolOptionsNntp(o["nntp"], d, "nntp")); err != nil {
		if !fortiAPIPatch(o["nntp"]) {
			return fmt.Errorf("Error reading nntp: %v", err)
		}
	}

	if err = d.Set("ssh", dataSourceFlattenFirewallProfileProtocolOptionsSsh(o["ssh"], d, "ssh")); err != nil {
		if !fortiAPIPatch(o["ssh"]) {
			return fmt.Errorf("Error reading ssh: %v", err)
		}
	}

	if err = d.Set("dns", dataSourceFlattenFirewallProfileProtocolOptionsDns(o["dns"], d, "dns")); err != nil {
		if !fortiAPIPatch(o["dns"]) {
			return fmt.Errorf("Error reading dns: %v", err)
		}
	}

	if err = d.Set("cifs", dataSourceFlattenFirewallProfileProtocolOptionsCifs(o["cifs"], d, "cifs")); err != nil {
		if !fortiAPIPatch(o["cifs"]) {
			return fmt.Errorf("Error reading cifs: %v", err)
		}
	}

	if err = d.Set("mail_signature", dataSourceFlattenFirewallProfileProtocolOptionsMailSignature(o["mail-signature"], d, "mail_signature")); err != nil {
		if !fortiAPIPatch(o["mail-signature"]) {
			return fmt.Errorf("Error reading mail_signature: %v", err)
		}
	}

	if err = d.Set("rpc_over_http", dataSourceFlattenFirewallProfileProtocolOptionsRpcOverHttp(o["rpc-over-http"], d, "rpc_over_http")); err != nil {
		if !fortiAPIPatch(o["rpc-over-http"]) {
			return fmt.Errorf("Error reading rpc_over_http: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenFirewallProfileProtocolOptionsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
