// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure protocol options.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallProfileProtocolOptions() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallProfileProtocolOptionsCreate,
		Read:   resourceFirewallProfileProtocolOptionsRead,
		Update: resourceFirewallProfileProtocolOptionsUpdate,
		Delete: resourceFirewallProfileProtocolOptionsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
				ForceNew:     true,
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
			"oversize_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switching_protocols_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"comfort_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 900),
							Optional:     true,
							Computed:     true,
						},
						"comfort_amount": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 10240),
							Optional:     true,
							Computed:     true,
						},
						"range_block": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"http_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"strip_x_forwarded_for": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"post_lang": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortinet_bar": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortinet_bar_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"streaming_content_bypass": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"switching_protocols": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),
							Optional:     true,
							Computed:     true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 26214),
							Optional:     true,
							Computed:     true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),
							Optional:     true,
							Computed:     true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_page_status_code": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(100, 599),
							Optional:     true,
							Computed:     true,
						},
						"retry_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"ftp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"comfort_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 900),
							Optional:     true,
							Computed:     true,
						},
						"comfort_amount": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 10240),
							Optional:     true,
							Computed:     true,
						},
						"oversize_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),
							Optional:     true,
							Computed:     true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 26214),
							Optional:     true,
							Computed:     true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),
							Optional:     true,
							Computed:     true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"imap": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),
							Optional:     true,
							Computed:     true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 26214),
							Optional:     true,
							Computed:     true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),
							Optional:     true,
							Computed:     true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"mapi": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),
							Optional:     true,
							Computed:     true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 26214),
							Optional:     true,
							Computed:     true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),
							Optional:     true,
							Computed:     true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"pop3": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),
							Optional:     true,
							Computed:     true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 26214),
							Optional:     true,
							Computed:     true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),
							Optional:     true,
							Computed:     true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"smtp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),
							Optional:     true,
							Computed:     true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 26214),
							Optional:     true,
							Computed:     true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),
							Optional:     true,
							Computed:     true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server_busy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"nntp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"oversize_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),
							Optional:     true,
							Computed:     true,
						},
						"uncompressed_oversize_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 26214),
							Optional:     true,
							Computed:     true,
						},
						"uncompressed_nest_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),
							Optional:     true,
							Computed:     true,
						},
						"scan_bzip2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dns": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
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
			"mail_signature": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"signature": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1023),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"rpc_over_http": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallProfileProtocolOptionsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallProfileProtocolOptions(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallProfileProtocolOptions resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallProfileProtocolOptions(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallProfileProtocolOptions resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallProfileProtocolOptions")
	}

	return resourceFirewallProfileProtocolOptionsRead(d, m)
}

func resourceFirewallProfileProtocolOptionsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallProfileProtocolOptions(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileProtocolOptions resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallProfileProtocolOptions(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileProtocolOptions resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallProfileProtocolOptions")
	}

	return resourceFirewallProfileProtocolOptionsRead(d, m)
}

func resourceFirewallProfileProtocolOptionsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallProfileProtocolOptions(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallProfileProtocolOptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallProfileProtocolOptionsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallProfileProtocolOptions(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProfileProtocolOptions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallProfileProtocolOptions(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProfileProtocolOptions resource from API: %v", err)
	}
	return nil
}

func flattenFirewallProfileProtocolOptionsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsOversizeLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSwitchingProtocolsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallProfileProtocolOptionsHttpPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsHttpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenFirewallProfileProtocolOptionsHttpInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenFirewallProfileProtocolOptionsHttpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "comfort_interval"
	if _, ok := i["comfort-interval"]; ok {
		result["comfort_interval"] = flattenFirewallProfileProtocolOptionsHttpComfortInterval(i["comfort-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "comfort_amount"
	if _, ok := i["comfort-amount"]; ok {
		result["comfort_amount"] = flattenFirewallProfileProtocolOptionsHttpComfortAmount(i["comfort-amount"], d, pre_append)
	}

	pre_append = pre + ".0." + "range_block"
	if _, ok := i["range-block"]; ok {
		result["range_block"] = flattenFirewallProfileProtocolOptionsHttpRangeBlock(i["range-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "http_policy"
	if _, ok := i["http-policy"]; ok {
		result["http_policy"] = flattenFirewallProfileProtocolOptionsHttpHttpPolicy(i["http-policy"], d, pre_append)
	}

	pre_append = pre + ".0." + "strip_x_forwarded_for"
	if _, ok := i["strip-x-forwarded-for"]; ok {
		result["strip_x_forwarded_for"] = flattenFirewallProfileProtocolOptionsHttpStripXForwardedFor(i["strip-x-forwarded-for"], d, pre_append)
	}

	pre_append = pre + ".0." + "post_lang"
	if _, ok := i["post-lang"]; ok {
		result["post_lang"] = flattenFirewallProfileProtocolOptionsHttpPostLang(i["post-lang"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortinet_bar"
	if _, ok := i["fortinet-bar"]; ok {
		result["fortinet_bar"] = flattenFirewallProfileProtocolOptionsHttpFortinetBar(i["fortinet-bar"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortinet_bar_port"
	if _, ok := i["fortinet-bar-port"]; ok {
		result["fortinet_bar_port"] = flattenFirewallProfileProtocolOptionsHttpFortinetBarPort(i["fortinet-bar-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "streaming_content_bypass"
	if _, ok := i["streaming-content-bypass"]; ok {
		result["streaming_content_bypass"] = flattenFirewallProfileProtocolOptionsHttpStreamingContentBypass(i["streaming-content-bypass"], d, pre_append)
	}

	pre_append = pre + ".0." + "switching_protocols"
	if _, ok := i["switching-protocols"]; ok {
		result["switching_protocols"] = flattenFirewallProfileProtocolOptionsHttpSwitchingProtocols(i["switching-protocols"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenFirewallProfileProtocolOptionsHttpOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenFirewallProfileProtocolOptionsHttpUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenFirewallProfileProtocolOptionsHttpUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenFirewallProfileProtocolOptionsHttpScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "block_page_status_code"
	if _, ok := i["block-page-status-code"]; ok {
		result["block_page_status_code"] = flattenFirewallProfileProtocolOptionsHttpBlockPageStatusCode(i["block-page-status-code"], d, pre_append)
	}

	pre_append = pre + ".0." + "retry_count"
	if _, ok := i["retry-count"]; ok {
		result["retry_count"] = flattenFirewallProfileProtocolOptionsHttpRetryCount(i["retry-count"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsHttpPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenFirewallProfileProtocolOptionsHttpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpComfortInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpComfortAmount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpRangeBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpHttpPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpStripXForwardedFor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpPostLang(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpFortinetBar(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpFortinetBarPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpStreamingContentBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpSwitchingProtocols(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpBlockPageStatusCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsHttpRetryCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallProfileProtocolOptionsFtpPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsFtpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenFirewallProfileProtocolOptionsFtpInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenFirewallProfileProtocolOptionsFtpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "comfort_interval"
	if _, ok := i["comfort-interval"]; ok {
		result["comfort_interval"] = flattenFirewallProfileProtocolOptionsFtpComfortInterval(i["comfort-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "comfort_amount"
	if _, ok := i["comfort-amount"]; ok {
		result["comfort_amount"] = flattenFirewallProfileProtocolOptionsFtpComfortAmount(i["comfort-amount"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenFirewallProfileProtocolOptionsFtpOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenFirewallProfileProtocolOptionsFtpUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenFirewallProfileProtocolOptionsFtpUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenFirewallProfileProtocolOptionsFtpScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsFtpPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenFirewallProfileProtocolOptionsFtpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpComfortInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpComfortAmount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsFtpScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsImap(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallProfileProtocolOptionsImapPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsImapStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenFirewallProfileProtocolOptionsImapInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenFirewallProfileProtocolOptionsImapOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenFirewallProfileProtocolOptionsImapOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenFirewallProfileProtocolOptionsImapUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenFirewallProfileProtocolOptionsImapUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenFirewallProfileProtocolOptionsImapScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsImapPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenFirewallProfileProtocolOptionsImapStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsImapInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsImapOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsImapOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsImapUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsImapUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsImapScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMapi(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallProfileProtocolOptionsMapiPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsMapiStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenFirewallProfileProtocolOptionsMapiOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenFirewallProfileProtocolOptionsMapiOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenFirewallProfileProtocolOptionsMapiUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenFirewallProfileProtocolOptionsMapiUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenFirewallProfileProtocolOptionsMapiScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsMapiPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenFirewallProfileProtocolOptionsMapiStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMapiOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMapiOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMapiUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMapiUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMapiScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsPop3(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallProfileProtocolOptionsPop3Ports(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsPop3Status(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenFirewallProfileProtocolOptionsPop3InspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenFirewallProfileProtocolOptionsPop3Options(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenFirewallProfileProtocolOptionsPop3OversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenFirewallProfileProtocolOptionsPop3UncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenFirewallProfileProtocolOptionsPop3UncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenFirewallProfileProtocolOptionsPop3ScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsPop3Ports(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenFirewallProfileProtocolOptionsPop3Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsPop3InspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsPop3Options(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsPop3OversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsPop3UncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsPop3UncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsPop3ScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallProfileProtocolOptionsSmtpPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsSmtpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenFirewallProfileProtocolOptionsSmtpInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenFirewallProfileProtocolOptionsSmtpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenFirewallProfileProtocolOptionsSmtpOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenFirewallProfileProtocolOptionsSmtpUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenFirewallProfileProtocolOptionsSmtpUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenFirewallProfileProtocolOptionsSmtpScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_busy"
	if _, ok := i["server-busy"]; ok {
		result["server_busy"] = flattenFirewallProfileProtocolOptionsSmtpServerBusy(i["server-busy"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsSmtpPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenFirewallProfileProtocolOptionsSmtpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsSmtpServerBusy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsNntp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallProfileProtocolOptionsNntpPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsNntpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenFirewallProfileProtocolOptionsNntpInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenFirewallProfileProtocolOptionsNntpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := i["oversize-limit"]; ok {
		result["oversize_limit"] = flattenFirewallProfileProtocolOptionsNntpOversizeLimit(i["oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := i["uncompressed-oversize-limit"]; ok {
		result["uncompressed_oversize_limit"] = flattenFirewallProfileProtocolOptionsNntpUncompressedOversizeLimit(i["uncompressed-oversize-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := i["uncompressed-nest-limit"]; ok {
		result["uncompressed_nest_limit"] = flattenFirewallProfileProtocolOptionsNntpUncompressedNestLimit(i["uncompressed-nest-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := i["scan-bzip2"]; ok {
		result["scan_bzip2"] = flattenFirewallProfileProtocolOptionsNntpScanBzip2(i["scan-bzip2"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsNntpPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenFirewallProfileProtocolOptionsNntpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsNntpInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsNntpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsNntpOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsNntpUncompressedOversizeLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsNntpUncompressedNestLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsNntpScanBzip2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsDns(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallProfileProtocolOptionsDnsPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsDnsStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsDnsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenFirewallProfileProtocolOptionsDnsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMailSignature(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsMailSignatureStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "signature"
	if _, ok := i["signature"]; ok {
		result["signature"] = flattenFirewallProfileProtocolOptionsMailSignatureSignature(i["signature"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsMailSignatureStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMailSignatureSignature(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsRpcOverHttp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallProfileProtocolOptions(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenFirewallProfileProtocolOptionsName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallProfileProtocolOptionsComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenFirewallProfileProtocolOptionsReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if !fortiAPIPatch(o["replacemsg-group"]) {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("oversize_log", flattenFirewallProfileProtocolOptionsOversizeLog(o["oversize-log"], d, "oversize_log")); err != nil {
		if !fortiAPIPatch(o["oversize-log"]) {
			return fmt.Errorf("Error reading oversize_log: %v", err)
		}
	}

	if err = d.Set("switching_protocols_log", flattenFirewallProfileProtocolOptionsSwitchingProtocolsLog(o["switching-protocols-log"], d, "switching_protocols_log")); err != nil {
		if !fortiAPIPatch(o["switching-protocols-log"]) {
			return fmt.Errorf("Error reading switching_protocols_log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("http", flattenFirewallProfileProtocolOptionsHttp(o["http"], d, "http")); err != nil {
			if !fortiAPIPatch(o["http"]) {
				return fmt.Errorf("Error reading http: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("http"); ok {
			if err = d.Set("http", flattenFirewallProfileProtocolOptionsHttp(o["http"], d, "http")); err != nil {
				if !fortiAPIPatch(o["http"]) {
					return fmt.Errorf("Error reading http: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ftp", flattenFirewallProfileProtocolOptionsFtp(o["ftp"], d, "ftp")); err != nil {
			if !fortiAPIPatch(o["ftp"]) {
				return fmt.Errorf("Error reading ftp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftp"); ok {
			if err = d.Set("ftp", flattenFirewallProfileProtocolOptionsFtp(o["ftp"], d, "ftp")); err != nil {
				if !fortiAPIPatch(o["ftp"]) {
					return fmt.Errorf("Error reading ftp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("imap", flattenFirewallProfileProtocolOptionsImap(o["imap"], d, "imap")); err != nil {
			if !fortiAPIPatch(o["imap"]) {
				return fmt.Errorf("Error reading imap: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("imap"); ok {
			if err = d.Set("imap", flattenFirewallProfileProtocolOptionsImap(o["imap"], d, "imap")); err != nil {
				if !fortiAPIPatch(o["imap"]) {
					return fmt.Errorf("Error reading imap: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mapi", flattenFirewallProfileProtocolOptionsMapi(o["mapi"], d, "mapi")); err != nil {
			if !fortiAPIPatch(o["mapi"]) {
				return fmt.Errorf("Error reading mapi: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mapi"); ok {
			if err = d.Set("mapi", flattenFirewallProfileProtocolOptionsMapi(o["mapi"], d, "mapi")); err != nil {
				if !fortiAPIPatch(o["mapi"]) {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("pop3", flattenFirewallProfileProtocolOptionsPop3(o["pop3"], d, "pop3")); err != nil {
			if !fortiAPIPatch(o["pop3"]) {
				return fmt.Errorf("Error reading pop3: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pop3"); ok {
			if err = d.Set("pop3", flattenFirewallProfileProtocolOptionsPop3(o["pop3"], d, "pop3")); err != nil {
				if !fortiAPIPatch(o["pop3"]) {
					return fmt.Errorf("Error reading pop3: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("smtp", flattenFirewallProfileProtocolOptionsSmtp(o["smtp"], d, "smtp")); err != nil {
			if !fortiAPIPatch(o["smtp"]) {
				return fmt.Errorf("Error reading smtp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("smtp"); ok {
			if err = d.Set("smtp", flattenFirewallProfileProtocolOptionsSmtp(o["smtp"], d, "smtp")); err != nil {
				if !fortiAPIPatch(o["smtp"]) {
					return fmt.Errorf("Error reading smtp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("nntp", flattenFirewallProfileProtocolOptionsNntp(o["nntp"], d, "nntp")); err != nil {
			if !fortiAPIPatch(o["nntp"]) {
				return fmt.Errorf("Error reading nntp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nntp"); ok {
			if err = d.Set("nntp", flattenFirewallProfileProtocolOptionsNntp(o["nntp"], d, "nntp")); err != nil {
				if !fortiAPIPatch(o["nntp"]) {
					return fmt.Errorf("Error reading nntp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dns", flattenFirewallProfileProtocolOptionsDns(o["dns"], d, "dns")); err != nil {
			if !fortiAPIPatch(o["dns"]) {
				return fmt.Errorf("Error reading dns: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dns"); ok {
			if err = d.Set("dns", flattenFirewallProfileProtocolOptionsDns(o["dns"], d, "dns")); err != nil {
				if !fortiAPIPatch(o["dns"]) {
					return fmt.Errorf("Error reading dns: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mail_signature", flattenFirewallProfileProtocolOptionsMailSignature(o["mail-signature"], d, "mail_signature")); err != nil {
			if !fortiAPIPatch(o["mail-signature"]) {
				return fmt.Errorf("Error reading mail_signature: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mail_signature"); ok {
			if err = d.Set("mail_signature", flattenFirewallProfileProtocolOptionsMailSignature(o["mail-signature"], d, "mail_signature")); err != nil {
				if !fortiAPIPatch(o["mail-signature"]) {
					return fmt.Errorf("Error reading mail_signature: %v", err)
				}
			}
		}
	}

	if err = d.Set("rpc_over_http", flattenFirewallProfileProtocolOptionsRpcOverHttp(o["rpc-over-http"], d, "rpc_over_http")); err != nil {
		if !fortiAPIPatch(o["rpc-over-http"]) {
			return fmt.Errorf("Error reading rpc_over_http: %v", err)
		}
	}

	return nil
}

func flattenFirewallProfileProtocolOptionsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallProfileProtocolOptionsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsOversizeLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSwitchingProtocolsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallProfileProtocolOptionsHttpPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallProfileProtocolOptionsHttpStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspect-all"], _ = expandFirewallProfileProtocolOptionsHttpInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandFirewallProfileProtocolOptionsHttpOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "comfort_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["comfort-interval"], _ = expandFirewallProfileProtocolOptionsHttpComfortInterval(d, i["comfort_interval"], pre_append)
	}
	pre_append = pre + ".0." + "comfort_amount"
	if _, ok := d.GetOk(pre_append); ok {
		result["comfort-amount"], _ = expandFirewallProfileProtocolOptionsHttpComfortAmount(d, i["comfort_amount"], pre_append)
	}
	pre_append = pre + ".0." + "range_block"
	if _, ok := d.GetOk(pre_append); ok {
		result["range-block"], _ = expandFirewallProfileProtocolOptionsHttpRangeBlock(d, i["range_block"], pre_append)
	}
	pre_append = pre + ".0." + "http_policy"
	if _, ok := d.GetOk(pre_append); ok {
		result["http-policy"], _ = expandFirewallProfileProtocolOptionsHttpHttpPolicy(d, i["http_policy"], pre_append)
	}
	pre_append = pre + ".0." + "strip_x_forwarded_for"
	if _, ok := d.GetOk(pre_append); ok {
		result["strip-x-forwarded-for"], _ = expandFirewallProfileProtocolOptionsHttpStripXForwardedFor(d, i["strip_x_forwarded_for"], pre_append)
	}
	pre_append = pre + ".0." + "post_lang"
	if _, ok := d.GetOk(pre_append); ok {
		result["post-lang"], _ = expandFirewallProfileProtocolOptionsHttpPostLang(d, i["post_lang"], pre_append)
	}
	pre_append = pre + ".0." + "fortinet_bar"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortinet-bar"], _ = expandFirewallProfileProtocolOptionsHttpFortinetBar(d, i["fortinet_bar"], pre_append)
	}
	pre_append = pre + ".0." + "fortinet_bar_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortinet-bar-port"], _ = expandFirewallProfileProtocolOptionsHttpFortinetBarPort(d, i["fortinet_bar_port"], pre_append)
	}
	pre_append = pre + ".0." + "streaming_content_bypass"
	if _, ok := d.GetOk(pre_append); ok {
		result["streaming-content-bypass"], _ = expandFirewallProfileProtocolOptionsHttpStreamingContentBypass(d, i["streaming_content_bypass"], pre_append)
	}
	pre_append = pre + ".0." + "switching_protocols"
	if _, ok := d.GetOk(pre_append); ok {
		result["switching-protocols"], _ = expandFirewallProfileProtocolOptionsHttpSwitchingProtocols(d, i["switching_protocols"], pre_append)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["oversize-limit"], _ = expandFirewallProfileProtocolOptionsHttpOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-oversize-limit"], _ = expandFirewallProfileProtocolOptionsHttpUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-nest-limit"], _ = expandFirewallProfileProtocolOptionsHttpUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok {
		result["scan-bzip2"], _ = expandFirewallProfileProtocolOptionsHttpScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "block_page_status_code"
	if _, ok := d.GetOk(pre_append); ok {
		result["block-page-status-code"], _ = expandFirewallProfileProtocolOptionsHttpBlockPageStatusCode(d, i["block_page_status_code"], pre_append)
	}
	pre_append = pre + ".0." + "retry_count"
	if _, ok := d.GetOk(pre_append); ok {
		result["retry-count"], _ = expandFirewallProfileProtocolOptionsHttpRetryCount(d, i["retry_count"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsHttpPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpComfortInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpComfortAmount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpRangeBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpHttpPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpStripXForwardedFor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpPostLang(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpFortinetBar(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpFortinetBarPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpStreamingContentBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpSwitchingProtocols(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpBlockPageStatusCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsHttpRetryCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallProfileProtocolOptionsFtpPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallProfileProtocolOptionsFtpStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspect-all"], _ = expandFirewallProfileProtocolOptionsFtpInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandFirewallProfileProtocolOptionsFtpOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "comfort_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["comfort-interval"], _ = expandFirewallProfileProtocolOptionsFtpComfortInterval(d, i["comfort_interval"], pre_append)
	}
	pre_append = pre + ".0." + "comfort_amount"
	if _, ok := d.GetOk(pre_append); ok {
		result["comfort-amount"], _ = expandFirewallProfileProtocolOptionsFtpComfortAmount(d, i["comfort_amount"], pre_append)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["oversize-limit"], _ = expandFirewallProfileProtocolOptionsFtpOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-oversize-limit"], _ = expandFirewallProfileProtocolOptionsFtpUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-nest-limit"], _ = expandFirewallProfileProtocolOptionsFtpUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok {
		result["scan-bzip2"], _ = expandFirewallProfileProtocolOptionsFtpScanBzip2(d, i["scan_bzip2"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsFtpPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpComfortInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpComfortAmount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsFtpScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsImap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallProfileProtocolOptionsImapPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallProfileProtocolOptionsImapStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspect-all"], _ = expandFirewallProfileProtocolOptionsImapInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandFirewallProfileProtocolOptionsImapOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["oversize-limit"], _ = expandFirewallProfileProtocolOptionsImapOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-oversize-limit"], _ = expandFirewallProfileProtocolOptionsImapUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-nest-limit"], _ = expandFirewallProfileProtocolOptionsImapUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok {
		result["scan-bzip2"], _ = expandFirewallProfileProtocolOptionsImapScanBzip2(d, i["scan_bzip2"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsImapPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsImapStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsImapInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsImapOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsImapOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsImapUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsImapUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsImapScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMapi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallProfileProtocolOptionsMapiPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallProfileProtocolOptionsMapiStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandFirewallProfileProtocolOptionsMapiOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["oversize-limit"], _ = expandFirewallProfileProtocolOptionsMapiOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-oversize-limit"], _ = expandFirewallProfileProtocolOptionsMapiUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-nest-limit"], _ = expandFirewallProfileProtocolOptionsMapiUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok {
		result["scan-bzip2"], _ = expandFirewallProfileProtocolOptionsMapiScanBzip2(d, i["scan_bzip2"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsMapiPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMapiStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMapiOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMapiOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMapiUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMapiUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMapiScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsPop3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallProfileProtocolOptionsPop3Ports(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallProfileProtocolOptionsPop3Status(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspect-all"], _ = expandFirewallProfileProtocolOptionsPop3InspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandFirewallProfileProtocolOptionsPop3Options(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["oversize-limit"], _ = expandFirewallProfileProtocolOptionsPop3OversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-oversize-limit"], _ = expandFirewallProfileProtocolOptionsPop3UncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-nest-limit"], _ = expandFirewallProfileProtocolOptionsPop3UncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok {
		result["scan-bzip2"], _ = expandFirewallProfileProtocolOptionsPop3ScanBzip2(d, i["scan_bzip2"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsPop3Ports(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsPop3Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsPop3InspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsPop3Options(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsPop3OversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsPop3UncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsPop3UncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsPop3ScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallProfileProtocolOptionsSmtpPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallProfileProtocolOptionsSmtpStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspect-all"], _ = expandFirewallProfileProtocolOptionsSmtpInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandFirewallProfileProtocolOptionsSmtpOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["oversize-limit"], _ = expandFirewallProfileProtocolOptionsSmtpOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-oversize-limit"], _ = expandFirewallProfileProtocolOptionsSmtpUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-nest-limit"], _ = expandFirewallProfileProtocolOptionsSmtpUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok {
		result["scan-bzip2"], _ = expandFirewallProfileProtocolOptionsSmtpScanBzip2(d, i["scan_bzip2"], pre_append)
	}
	pre_append = pre + ".0." + "server_busy"
	if _, ok := d.GetOk(pre_append); ok {
		result["server-busy"], _ = expandFirewallProfileProtocolOptionsSmtpServerBusy(d, i["server_busy"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsSmtpPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsSmtpServerBusy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsNntp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallProfileProtocolOptionsNntpPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallProfileProtocolOptionsNntpStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspect-all"], _ = expandFirewallProfileProtocolOptionsNntpInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandFirewallProfileProtocolOptionsNntpOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["oversize-limit"], _ = expandFirewallProfileProtocolOptionsNntpOversizeLimit(d, i["oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_oversize_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-oversize-limit"], _ = expandFirewallProfileProtocolOptionsNntpUncompressedOversizeLimit(d, i["uncompressed_oversize_limit"], pre_append)
	}
	pre_append = pre + ".0." + "uncompressed_nest_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["uncompressed-nest-limit"], _ = expandFirewallProfileProtocolOptionsNntpUncompressedNestLimit(d, i["uncompressed_nest_limit"], pre_append)
	}
	pre_append = pre + ".0." + "scan_bzip2"
	if _, ok := d.GetOk(pre_append); ok {
		result["scan-bzip2"], _ = expandFirewallProfileProtocolOptionsNntpScanBzip2(d, i["scan_bzip2"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsNntpPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsNntpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsNntpInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsNntpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsNntpOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsNntpUncompressedOversizeLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsNntpUncompressedNestLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsNntpScanBzip2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallProfileProtocolOptionsDnsPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallProfileProtocolOptionsDnsStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsDnsPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsDnsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMailSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallProfileProtocolOptionsMailSignatureStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "signature"
	if _, ok := d.GetOk(pre_append); ok {
		result["signature"], _ = expandFirewallProfileProtocolOptionsMailSignatureSignature(d, i["signature"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsMailSignatureStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMailSignatureSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsRpcOverHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallProfileProtocolOptions(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallProfileProtocolOptionsName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandFirewallProfileProtocolOptionsComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok {
		t, err := expandFirewallProfileProtocolOptionsReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("oversize_log"); ok {
		t, err := expandFirewallProfileProtocolOptionsOversizeLog(d, v, "oversize_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oversize-log"] = t
		}
	}

	if v, ok := d.GetOk("switching_protocols_log"); ok {
		t, err := expandFirewallProfileProtocolOptionsSwitchingProtocolsLog(d, v, "switching_protocols_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switching-protocols-log"] = t
		}
	}

	if v, ok := d.GetOk("http"); ok {
		t, err := expandFirewallProfileProtocolOptionsHttp(d, v, "http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http"] = t
		}
	}

	if v, ok := d.GetOk("ftp"); ok {
		t, err := expandFirewallProfileProtocolOptionsFtp(d, v, "ftp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp"] = t
		}
	}

	if v, ok := d.GetOk("imap"); ok {
		t, err := expandFirewallProfileProtocolOptionsImap(d, v, "imap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imap"] = t
		}
	}

	if v, ok := d.GetOk("mapi"); ok {
		t, err := expandFirewallProfileProtocolOptionsMapi(d, v, "mapi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapi"] = t
		}
	}

	if v, ok := d.GetOk("pop3"); ok {
		t, err := expandFirewallProfileProtocolOptionsPop3(d, v, "pop3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3"] = t
		}
	}

	if v, ok := d.GetOk("smtp"); ok {
		t, err := expandFirewallProfileProtocolOptionsSmtp(d, v, "smtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtp"] = t
		}
	}

	if v, ok := d.GetOk("nntp"); ok {
		t, err := expandFirewallProfileProtocolOptionsNntp(d, v, "nntp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nntp"] = t
		}
	}

	if v, ok := d.GetOk("dns"); ok {
		t, err := expandFirewallProfileProtocolOptionsDns(d, v, "dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns"] = t
		}
	}

	if v, ok := d.GetOk("mail_signature"); ok {
		t, err := expandFirewallProfileProtocolOptionsMailSignature(d, v, "mail_signature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mail-signature"] = t
		}
	}

	if v, ok := d.GetOk("rpc_over_http"); ok {
		t, err := expandFirewallProfileProtocolOptionsRpcOverHttp(d, v, "rpc_over_http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpc-over-http"] = t
		}
	}

	return &obj, nil
}
