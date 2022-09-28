// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure VoIP profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVoipProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceVoipProfileCreate,
		Read:   resourceVoipProfileRead,
		Update: resourceVoipProfileUpdate,
		Delete: resourceVoipProfileDelete,

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
			"feature_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"sip": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rtp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nat_port_range": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"open_register_pinhole": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"open_contact_pinhole": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"strict_register": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"register_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"register_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"invite_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"invite_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_dialogs": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_line_length": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(78, 4096),
							Optional:     true,
							Computed:     true,
						},
						"block_long_lines": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_unknown": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"call_keepalive": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10080),
							Optional:     true,
							Computed:     true,
						},
						"block_ack": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_bye": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_cancel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_info": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_invite": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_message": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_notify": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_options": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_prack": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_publish": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_refer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_register": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_subscribe": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_update": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"register_contact_trace": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"open_via_pinhole": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"open_record_route_pinhole": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rfc2543_branch": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_violations": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_call_summary": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nat_trace": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"subscribe_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"subscribe_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"message_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"message_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"notify_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"notify_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"refer_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"refer_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"update_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"update_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"options_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ack_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ack_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prack_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"prack_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"info_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"info_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"publish_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"publish_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bye_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"bye_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cancel_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"cancel_rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"preserve_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"no_sdp_fixup": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"contact_fixup": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_idle_dialogs": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"block_geo_red_options": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"hosted_nat_traversal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"hnt_restrict_source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_body_length": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"unknown_header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_request_line": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_via": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_from": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_to": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_call_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_cseq": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_rack": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_rseq": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_contact": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_record_route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_expires": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_content_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_content_length": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_max_forwards": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_allow": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_p_asserted_identity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_no_require": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_no_proxy_require": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_o": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_s": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_i": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_c": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_b": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_z": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_k": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_a": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_t": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_r": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"malformed_header_sdp_m": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"provisional_invite_expiry_time": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(10, 3600),
							Optional:     true,
							Computed:     true,
						},
						"ips_rtp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_send_empty_frags": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_client_renegotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_algorithm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_pfs": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_min_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_max_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_client_certificate": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"ssl_server_certificate": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"ssl_auth_client": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"ssl_auth_server": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"sccp": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"block_mcast": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"verify_header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_call_summary": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_violations": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_calls": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"msrp": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_violations": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_msg_size": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"max_msg_size_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceVoipProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVoipProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating VoipProfile resource while getting object: %v", err)
	}

	o, err := c.CreateVoipProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating VoipProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VoipProfile")
	}

	return resourceVoipProfileRead(d, m)
}

func resourceVoipProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVoipProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VoipProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateVoipProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating VoipProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VoipProfile")
	}

	return resourceVoipProfileRead(d, m)
}

func resourceVoipProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteVoipProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting VoipProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVoipProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadVoipProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VoipProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVoipProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VoipProfile resource from API: %v", err)
	}
	return nil
}

func flattenVoipProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileFeatureSet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSip(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {

		result["status"] = flattenVoipProfileSipStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "rtp"
	if _, ok := i["rtp"]; ok {

		result["rtp"] = flattenVoipProfileSipRtp(i["rtp"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "nat_port_range"
	if _, ok := i["nat-port-range"]; ok {

		result["nat_port_range"] = flattenVoipProfileSipNatPortRange(i["nat-port-range"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "open_register_pinhole"
	if _, ok := i["open-register-pinhole"]; ok {

		result["open_register_pinhole"] = flattenVoipProfileSipOpenRegisterPinhole(i["open-register-pinhole"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "open_contact_pinhole"
	if _, ok := i["open-contact-pinhole"]; ok {

		result["open_contact_pinhole"] = flattenVoipProfileSipOpenContactPinhole(i["open-contact-pinhole"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "strict_register"
	if _, ok := i["strict-register"]; ok {

		result["strict_register"] = flattenVoipProfileSipStrictRegister(i["strict-register"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "register_rate"
	if _, ok := i["register-rate"]; ok {

		result["register_rate"] = flattenVoipProfileSipRegisterRate(i["register-rate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "register_rate_track"
	if _, ok := i["register-rate-track"]; ok {

		result["register_rate_track"] = flattenVoipProfileSipRegisterRateTrack(i["register-rate-track"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "invite_rate"
	if _, ok := i["invite-rate"]; ok {

		result["invite_rate"] = flattenVoipProfileSipInviteRate(i["invite-rate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "invite_rate_track"
	if _, ok := i["invite-rate-track"]; ok {

		result["invite_rate_track"] = flattenVoipProfileSipInviteRateTrack(i["invite-rate-track"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_dialogs"
	if _, ok := i["max-dialogs"]; ok {

		result["max_dialogs"] = flattenVoipProfileSipMaxDialogs(i["max-dialogs"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_line_length"
	if _, ok := i["max-line-length"]; ok {

		result["max_line_length"] = flattenVoipProfileSipMaxLineLength(i["max-line-length"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block_long_lines"
	if _, ok := i["block-long-lines"]; ok {

		result["block_long_lines"] = flattenVoipProfileSipBlockLongLines(i["block-long-lines"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block_unknown"
	if _, ok := i["block-unknown"]; ok {

		result["block_unknown"] = flattenVoipProfileSipBlockUnknown(i["block-unknown"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "call_keepalive"
	if _, ok := i["call-keepalive"]; ok {

		result["call_keepalive"] = flattenVoipProfileSipCallKeepalive(i["call-keepalive"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block_ack"
	if _, ok := i["block-ack"]; ok {

		result["block_ack"] = flattenVoipProfileSipBlockAck(i["block-ack"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block_bye"
	if _, ok := i["block-bye"]; ok {

		result["block_bye"] = flattenVoipProfileSipBlockBye(i["block-bye"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block_cancel"
	if _, ok := i["block-cancel"]; ok {

		result["block_cancel"] = flattenVoipProfileSipBlockCancel(i["block-cancel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block_info"
	if _, ok := i["block-info"]; ok {

		result["block_info"] = flattenVoipProfileSipBlockInfo(i["block-info"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block_invite"
	if _, ok := i["block-invite"]; ok {

		result["block_invite"] = flattenVoipProfileSipBlockInvite(i["block-invite"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block_message"
	if _, ok := i["block-message"]; ok {

		result["block_message"] = flattenVoipProfileSipBlockMessage(i["block-message"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block_notify"
	if _, ok := i["block-notify"]; ok {

		result["block_notify"] = flattenVoipProfileSipBlockNotify(i["block-notify"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block_options"
	if _, ok := i["block-options"]; ok {

		result["block_options"] = flattenVoipProfileSipBlockOptions(i["block-options"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block_prack"
	if _, ok := i["block-prack"]; ok {

		result["block_prack"] = flattenVoipProfileSipBlockPrack(i["block-prack"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block_publish"
	if _, ok := i["block-publish"]; ok {

		result["block_publish"] = flattenVoipProfileSipBlockPublish(i["block-publish"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block_refer"
	if _, ok := i["block-refer"]; ok {

		result["block_refer"] = flattenVoipProfileSipBlockRefer(i["block-refer"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block_register"
	if _, ok := i["block-register"]; ok {

		result["block_register"] = flattenVoipProfileSipBlockRegister(i["block-register"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block_subscribe"
	if _, ok := i["block-subscribe"]; ok {

		result["block_subscribe"] = flattenVoipProfileSipBlockSubscribe(i["block-subscribe"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block_update"
	if _, ok := i["block-update"]; ok {

		result["block_update"] = flattenVoipProfileSipBlockUpdate(i["block-update"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "register_contact_trace"
	if _, ok := i["register-contact-trace"]; ok {

		result["register_contact_trace"] = flattenVoipProfileSipRegisterContactTrace(i["register-contact-trace"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "open_via_pinhole"
	if _, ok := i["open-via-pinhole"]; ok {

		result["open_via_pinhole"] = flattenVoipProfileSipOpenViaPinhole(i["open-via-pinhole"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "open_record_route_pinhole"
	if _, ok := i["open-record-route-pinhole"]; ok {

		result["open_record_route_pinhole"] = flattenVoipProfileSipOpenRecordRoutePinhole(i["open-record-route-pinhole"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "rfc2543_branch"
	if _, ok := i["rfc2543-branch"]; ok {

		result["rfc2543_branch"] = flattenVoipProfileSipRfc2543Branch(i["rfc2543-branch"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "log_violations"
	if _, ok := i["log-violations"]; ok {

		result["log_violations"] = flattenVoipProfileSipLogViolations(i["log-violations"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "log_call_summary"
	if _, ok := i["log-call-summary"]; ok {

		result["log_call_summary"] = flattenVoipProfileSipLogCallSummary(i["log-call-summary"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "nat_trace"
	if _, ok := i["nat-trace"]; ok {

		result["nat_trace"] = flattenVoipProfileSipNatTrace(i["nat-trace"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "subscribe_rate"
	if _, ok := i["subscribe-rate"]; ok {

		result["subscribe_rate"] = flattenVoipProfileSipSubscribeRate(i["subscribe-rate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "subscribe_rate_track"
	if _, ok := i["subscribe-rate-track"]; ok {

		result["subscribe_rate_track"] = flattenVoipProfileSipSubscribeRateTrack(i["subscribe-rate-track"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "message_rate"
	if _, ok := i["message-rate"]; ok {

		result["message_rate"] = flattenVoipProfileSipMessageRate(i["message-rate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "message_rate_track"
	if _, ok := i["message-rate-track"]; ok {

		result["message_rate_track"] = flattenVoipProfileSipMessageRateTrack(i["message-rate-track"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "notify_rate"
	if _, ok := i["notify-rate"]; ok {

		result["notify_rate"] = flattenVoipProfileSipNotifyRate(i["notify-rate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "notify_rate_track"
	if _, ok := i["notify-rate-track"]; ok {

		result["notify_rate_track"] = flattenVoipProfileSipNotifyRateTrack(i["notify-rate-track"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "refer_rate"
	if _, ok := i["refer-rate"]; ok {

		result["refer_rate"] = flattenVoipProfileSipReferRate(i["refer-rate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "refer_rate_track"
	if _, ok := i["refer-rate-track"]; ok {

		result["refer_rate_track"] = flattenVoipProfileSipReferRateTrack(i["refer-rate-track"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "update_rate"
	if _, ok := i["update-rate"]; ok {

		result["update_rate"] = flattenVoipProfileSipUpdateRate(i["update-rate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "update_rate_track"
	if _, ok := i["update-rate-track"]; ok {

		result["update_rate_track"] = flattenVoipProfileSipUpdateRateTrack(i["update-rate-track"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "options_rate"
	if _, ok := i["options-rate"]; ok {

		result["options_rate"] = flattenVoipProfileSipOptionsRate(i["options-rate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "options_rate_track"
	if _, ok := i["options-rate-track"]; ok {

		result["options_rate_track"] = flattenVoipProfileSipOptionsRateTrack(i["options-rate-track"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ack_rate"
	if _, ok := i["ack-rate"]; ok {

		result["ack_rate"] = flattenVoipProfileSipAckRate(i["ack-rate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ack_rate_track"
	if _, ok := i["ack-rate-track"]; ok {

		result["ack_rate_track"] = flattenVoipProfileSipAckRateTrack(i["ack-rate-track"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "prack_rate"
	if _, ok := i["prack-rate"]; ok {

		result["prack_rate"] = flattenVoipProfileSipPrackRate(i["prack-rate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "prack_rate_track"
	if _, ok := i["prack-rate-track"]; ok {

		result["prack_rate_track"] = flattenVoipProfileSipPrackRateTrack(i["prack-rate-track"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "info_rate"
	if _, ok := i["info-rate"]; ok {

		result["info_rate"] = flattenVoipProfileSipInfoRate(i["info-rate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "info_rate_track"
	if _, ok := i["info-rate-track"]; ok {

		result["info_rate_track"] = flattenVoipProfileSipInfoRateTrack(i["info-rate-track"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "publish_rate"
	if _, ok := i["publish-rate"]; ok {

		result["publish_rate"] = flattenVoipProfileSipPublishRate(i["publish-rate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "publish_rate_track"
	if _, ok := i["publish-rate-track"]; ok {

		result["publish_rate_track"] = flattenVoipProfileSipPublishRateTrack(i["publish-rate-track"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bye_rate"
	if _, ok := i["bye-rate"]; ok {

		result["bye_rate"] = flattenVoipProfileSipByeRate(i["bye-rate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bye_rate_track"
	if _, ok := i["bye-rate-track"]; ok {

		result["bye_rate_track"] = flattenVoipProfileSipByeRateTrack(i["bye-rate-track"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cancel_rate"
	if _, ok := i["cancel-rate"]; ok {

		result["cancel_rate"] = flattenVoipProfileSipCancelRate(i["cancel-rate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cancel_rate_track"
	if _, ok := i["cancel-rate-track"]; ok {

		result["cancel_rate_track"] = flattenVoipProfileSipCancelRateTrack(i["cancel-rate-track"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "preserve_override"
	if _, ok := i["preserve-override"]; ok {

		result["preserve_override"] = flattenVoipProfileSipPreserveOverride(i["preserve-override"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "no_sdp_fixup"
	if _, ok := i["no-sdp-fixup"]; ok {

		result["no_sdp_fixup"] = flattenVoipProfileSipNoSdpFixup(i["no-sdp-fixup"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "contact_fixup"
	if _, ok := i["contact-fixup"]; ok {

		result["contact_fixup"] = flattenVoipProfileSipContactFixup(i["contact-fixup"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_idle_dialogs"
	if _, ok := i["max-idle-dialogs"]; ok {

		result["max_idle_dialogs"] = flattenVoipProfileSipMaxIdleDialogs(i["max-idle-dialogs"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block_geo_red_options"
	if _, ok := i["block-geo-red-options"]; ok {

		result["block_geo_red_options"] = flattenVoipProfileSipBlockGeoRedOptions(i["block-geo-red-options"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "hosted_nat_traversal"
	if _, ok := i["hosted-nat-traversal"]; ok {

		result["hosted_nat_traversal"] = flattenVoipProfileSipHostedNatTraversal(i["hosted-nat-traversal"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "hnt_restrict_source_ip"
	if _, ok := i["hnt-restrict-source-ip"]; ok {

		result["hnt_restrict_source_ip"] = flattenVoipProfileSipHntRestrictSourceIp(i["hnt-restrict-source-ip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_body_length"
	if _, ok := i["max-body-length"]; ok {

		result["max_body_length"] = flattenVoipProfileSipMaxBodyLength(i["max-body-length"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unknown_header"
	if _, ok := i["unknown-header"]; ok {

		result["unknown_header"] = flattenVoipProfileSipUnknownHeader(i["unknown-header"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_request_line"
	if _, ok := i["malformed-request-line"]; ok {

		result["malformed_request_line"] = flattenVoipProfileSipMalformedRequestLine(i["malformed-request-line"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_via"
	if _, ok := i["malformed-header-via"]; ok {

		result["malformed_header_via"] = flattenVoipProfileSipMalformedHeaderVia(i["malformed-header-via"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_from"
	if _, ok := i["malformed-header-from"]; ok {

		result["malformed_header_from"] = flattenVoipProfileSipMalformedHeaderFrom(i["malformed-header-from"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_to"
	if _, ok := i["malformed-header-to"]; ok {

		result["malformed_header_to"] = flattenVoipProfileSipMalformedHeaderTo(i["malformed-header-to"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_call_id"
	if _, ok := i["malformed-header-call-id"]; ok {

		result["malformed_header_call_id"] = flattenVoipProfileSipMalformedHeaderCallId(i["malformed-header-call-id"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_cseq"
	if _, ok := i["malformed-header-cseq"]; ok {

		result["malformed_header_cseq"] = flattenVoipProfileSipMalformedHeaderCseq(i["malformed-header-cseq"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_rack"
	if _, ok := i["malformed-header-rack"]; ok {

		result["malformed_header_rack"] = flattenVoipProfileSipMalformedHeaderRack(i["malformed-header-rack"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_rseq"
	if _, ok := i["malformed-header-rseq"]; ok {

		result["malformed_header_rseq"] = flattenVoipProfileSipMalformedHeaderRseq(i["malformed-header-rseq"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_contact"
	if _, ok := i["malformed-header-contact"]; ok {

		result["malformed_header_contact"] = flattenVoipProfileSipMalformedHeaderContact(i["malformed-header-contact"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_record_route"
	if _, ok := i["malformed-header-record-route"]; ok {

		result["malformed_header_record_route"] = flattenVoipProfileSipMalformedHeaderRecordRoute(i["malformed-header-record-route"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_route"
	if _, ok := i["malformed-header-route"]; ok {

		result["malformed_header_route"] = flattenVoipProfileSipMalformedHeaderRoute(i["malformed-header-route"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_expires"
	if _, ok := i["malformed-header-expires"]; ok {

		result["malformed_header_expires"] = flattenVoipProfileSipMalformedHeaderExpires(i["malformed-header-expires"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_content_type"
	if _, ok := i["malformed-header-content-type"]; ok {

		result["malformed_header_content_type"] = flattenVoipProfileSipMalformedHeaderContentType(i["malformed-header-content-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_content_length"
	if _, ok := i["malformed-header-content-length"]; ok {

		result["malformed_header_content_length"] = flattenVoipProfileSipMalformedHeaderContentLength(i["malformed-header-content-length"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_max_forwards"
	if _, ok := i["malformed-header-max-forwards"]; ok {

		result["malformed_header_max_forwards"] = flattenVoipProfileSipMalformedHeaderMaxForwards(i["malformed-header-max-forwards"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_allow"
	if _, ok := i["malformed-header-allow"]; ok {

		result["malformed_header_allow"] = flattenVoipProfileSipMalformedHeaderAllow(i["malformed-header-allow"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_p_asserted_identity"
	if _, ok := i["malformed-header-p-asserted-identity"]; ok {

		result["malformed_header_p_asserted_identity"] = flattenVoipProfileSipMalformedHeaderPAssertedIdentity(i["malformed-header-p-asserted-identity"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_no_require"
	if _, ok := i["malformed-header-no-require"]; ok {

		result["malformed_header_no_require"] = flattenVoipProfileSipMalformedHeaderNoRequire(i["malformed-header-no-require"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_no_proxy_require"
	if _, ok := i["malformed-header-no-proxy-require"]; ok {

		result["malformed_header_no_proxy_require"] = flattenVoipProfileSipMalformedHeaderNoProxyRequire(i["malformed-header-no-proxy-require"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_v"
	if _, ok := i["malformed-header-sdp-v"]; ok {

		result["malformed_header_sdp_v"] = flattenVoipProfileSipMalformedHeaderSdpV(i["malformed-header-sdp-v"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_o"
	if _, ok := i["malformed-header-sdp-o"]; ok {

		result["malformed_header_sdp_o"] = flattenVoipProfileSipMalformedHeaderSdpO(i["malformed-header-sdp-o"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_s"
	if _, ok := i["malformed-header-sdp-s"]; ok {

		result["malformed_header_sdp_s"] = flattenVoipProfileSipMalformedHeaderSdpS(i["malformed-header-sdp-s"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_i"
	if _, ok := i["malformed-header-sdp-i"]; ok {

		result["malformed_header_sdp_i"] = flattenVoipProfileSipMalformedHeaderSdpI(i["malformed-header-sdp-i"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_c"
	if _, ok := i["malformed-header-sdp-c"]; ok {

		result["malformed_header_sdp_c"] = flattenVoipProfileSipMalformedHeaderSdpC(i["malformed-header-sdp-c"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_b"
	if _, ok := i["malformed-header-sdp-b"]; ok {

		result["malformed_header_sdp_b"] = flattenVoipProfileSipMalformedHeaderSdpB(i["malformed-header-sdp-b"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_z"
	if _, ok := i["malformed-header-sdp-z"]; ok {

		result["malformed_header_sdp_z"] = flattenVoipProfileSipMalformedHeaderSdpZ(i["malformed-header-sdp-z"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_k"
	if _, ok := i["malformed-header-sdp-k"]; ok {

		result["malformed_header_sdp_k"] = flattenVoipProfileSipMalformedHeaderSdpK(i["malformed-header-sdp-k"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_a"
	if _, ok := i["malformed-header-sdp-a"]; ok {

		result["malformed_header_sdp_a"] = flattenVoipProfileSipMalformedHeaderSdpA(i["malformed-header-sdp-a"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_t"
	if _, ok := i["malformed-header-sdp-t"]; ok {

		result["malformed_header_sdp_t"] = flattenVoipProfileSipMalformedHeaderSdpT(i["malformed-header-sdp-t"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_r"
	if _, ok := i["malformed-header-sdp-r"]; ok {

		result["malformed_header_sdp_r"] = flattenVoipProfileSipMalformedHeaderSdpR(i["malformed-header-sdp-r"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "malformed_header_sdp_m"
	if _, ok := i["malformed-header-sdp-m"]; ok {

		result["malformed_header_sdp_m"] = flattenVoipProfileSipMalformedHeaderSdpM(i["malformed-header-sdp-m"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "provisional_invite_expiry_time"
	if _, ok := i["provisional-invite-expiry-time"]; ok {

		result["provisional_invite_expiry_time"] = flattenVoipProfileSipProvisionalInviteExpiryTime(i["provisional-invite-expiry-time"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ips_rtp"
	if _, ok := i["ips-rtp"]; ok {

		result["ips_rtp"] = flattenVoipProfileSipIpsRtp(i["ips-rtp"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ssl_mode"
	if _, ok := i["ssl-mode"]; ok {

		result["ssl_mode"] = flattenVoipProfileSipSslMode(i["ssl-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ssl_send_empty_frags"
	if _, ok := i["ssl-send-empty-frags"]; ok {

		result["ssl_send_empty_frags"] = flattenVoipProfileSipSslSendEmptyFrags(i["ssl-send-empty-frags"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ssl_client_renegotiation"
	if _, ok := i["ssl-client-renegotiation"]; ok {

		result["ssl_client_renegotiation"] = flattenVoipProfileSipSslClientRenegotiation(i["ssl-client-renegotiation"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ssl_algorithm"
	if _, ok := i["ssl-algorithm"]; ok {

		result["ssl_algorithm"] = flattenVoipProfileSipSslAlgorithm(i["ssl-algorithm"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ssl_pfs"
	if _, ok := i["ssl-pfs"]; ok {

		result["ssl_pfs"] = flattenVoipProfileSipSslPfs(i["ssl-pfs"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ssl_min_version"
	if _, ok := i["ssl-min-version"]; ok {

		result["ssl_min_version"] = flattenVoipProfileSipSslMinVersion(i["ssl-min-version"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ssl_max_version"
	if _, ok := i["ssl-max-version"]; ok {

		result["ssl_max_version"] = flattenVoipProfileSipSslMaxVersion(i["ssl-max-version"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ssl_client_certificate"
	if _, ok := i["ssl-client-certificate"]; ok {

		result["ssl_client_certificate"] = flattenVoipProfileSipSslClientCertificate(i["ssl-client-certificate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ssl_server_certificate"
	if _, ok := i["ssl-server-certificate"]; ok {

		result["ssl_server_certificate"] = flattenVoipProfileSipSslServerCertificate(i["ssl-server-certificate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ssl_auth_client"
	if _, ok := i["ssl-auth-client"]; ok {

		result["ssl_auth_client"] = flattenVoipProfileSipSslAuthClient(i["ssl-auth-client"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ssl_auth_server"
	if _, ok := i["ssl-auth-server"]; ok {

		result["ssl_auth_server"] = flattenVoipProfileSipSslAuthServer(i["ssl-auth-server"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenVoipProfileSipStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipRtp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipNatPortRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipOpenRegisterPinhole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipOpenContactPinhole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipStrictRegister(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipRegisterRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipRegisterRateTrack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipInviteRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipInviteRateTrack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMaxDialogs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMaxLineLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipBlockLongLines(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipBlockUnknown(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipCallKeepalive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipBlockAck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipBlockBye(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipBlockCancel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipBlockInfo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipBlockInvite(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipBlockMessage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipBlockNotify(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipBlockOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipBlockPrack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipBlockPublish(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipBlockRefer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipBlockRegister(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipBlockSubscribe(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipBlockUpdate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipRegisterContactTrace(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipOpenViaPinhole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipOpenRecordRoutePinhole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipRfc2543Branch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipLogViolations(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipLogCallSummary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipNatTrace(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipSubscribeRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipSubscribeRateTrack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMessageRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMessageRateTrack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipNotifyRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipNotifyRateTrack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipReferRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipReferRateTrack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipUpdateRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipUpdateRateTrack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipOptionsRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipOptionsRateTrack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipAckRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipAckRateTrack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipPrackRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipPrackRateTrack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipInfoRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipInfoRateTrack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipPublishRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipPublishRateTrack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipByeRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipByeRateTrack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipCancelRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipCancelRateTrack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipPreserveOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipNoSdpFixup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipContactFixup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMaxIdleDialogs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipBlockGeoRedOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipHostedNatTraversal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipHntRestrictSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMaxBodyLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipUnknownHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedRequestLine(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderVia(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderFrom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderTo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderCallId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderCseq(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderRack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderRseq(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderContact(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderRecordRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderExpires(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderContentType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderContentLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderMaxForwards(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderAllow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderPAssertedIdentity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderNoRequire(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderNoProxyRequire(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpV(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpO(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpS(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpI(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpC(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpB(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpZ(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpK(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpA(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpT(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpR(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipMalformedHeaderSdpM(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipProvisionalInviteExpiryTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipIpsRtp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipSslMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipSslSendEmptyFrags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipSslClientRenegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipSslAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipSslPfs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipSslMinVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipSslMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipSslClientCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipSslServerCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipSslAuthClient(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSipSslAuthServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSccp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {

		result["status"] = flattenVoipProfileSccpStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block_mcast"
	if _, ok := i["block-mcast"]; ok {

		result["block_mcast"] = flattenVoipProfileSccpBlockMcast(i["block-mcast"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "verify_header"
	if _, ok := i["verify-header"]; ok {

		result["verify_header"] = flattenVoipProfileSccpVerifyHeader(i["verify-header"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "log_call_summary"
	if _, ok := i["log-call-summary"]; ok {

		result["log_call_summary"] = flattenVoipProfileSccpLogCallSummary(i["log-call-summary"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "log_violations"
	if _, ok := i["log-violations"]; ok {

		result["log_violations"] = flattenVoipProfileSccpLogViolations(i["log-violations"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_calls"
	if _, ok := i["max-calls"]; ok {

		result["max_calls"] = flattenVoipProfileSccpMaxCalls(i["max-calls"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenVoipProfileSccpStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSccpBlockMcast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSccpVerifyHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSccpLogCallSummary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSccpLogViolations(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileSccpMaxCalls(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileMsrp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {

		result["status"] = flattenVoipProfileMsrpStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "log_violations"
	if _, ok := i["log-violations"]; ok {

		result["log_violations"] = flattenVoipProfileMsrpLogViolations(i["log-violations"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_msg_size"
	if _, ok := i["max-msg-size"]; ok {

		result["max_msg_size"] = flattenVoipProfileMsrpMaxMsgSize(i["max-msg-size"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_msg_size_action"
	if _, ok := i["max-msg-size-action"]; ok {

		result["max_msg_size_action"] = flattenVoipProfileMsrpMaxMsgSizeAction(i["max-msg-size-action"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenVoipProfileMsrpStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileMsrpLogViolations(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileMsrpMaxMsgSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVoipProfileMsrpMaxMsgSizeAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVoipProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenVoipProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("feature_set", flattenVoipProfileFeatureSet(o["feature-set"], d, "feature_set", sv)); err != nil {
		if !fortiAPIPatch(o["feature-set"]) {
			return fmt.Errorf("Error reading feature_set: %v", err)
		}
	}

	if err = d.Set("comment", flattenVoipProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sip", flattenVoipProfileSip(o["sip"], d, "sip", sv)); err != nil {
			if !fortiAPIPatch(o["sip"]) {
				return fmt.Errorf("Error reading sip: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sip"); ok {
			if err = d.Set("sip", flattenVoipProfileSip(o["sip"], d, "sip", sv)); err != nil {
				if !fortiAPIPatch(o["sip"]) {
					return fmt.Errorf("Error reading sip: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("sccp", flattenVoipProfileSccp(o["sccp"], d, "sccp", sv)); err != nil {
			if !fortiAPIPatch(o["sccp"]) {
				return fmt.Errorf("Error reading sccp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sccp"); ok {
			if err = d.Set("sccp", flattenVoipProfileSccp(o["sccp"], d, "sccp", sv)); err != nil {
				if !fortiAPIPatch(o["sccp"]) {
					return fmt.Errorf("Error reading sccp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("msrp", flattenVoipProfileMsrp(o["msrp"], d, "msrp", sv)); err != nil {
			if !fortiAPIPatch(o["msrp"]) {
				return fmt.Errorf("Error reading msrp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("msrp"); ok {
			if err = d.Set("msrp", flattenVoipProfileMsrp(o["msrp"], d, "msrp", sv)); err != nil {
				if !fortiAPIPatch(o["msrp"]) {
					return fmt.Errorf("Error reading msrp: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenVoipProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVoipProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileFeatureSet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {

		result["status"], _ = expandVoipProfileSipStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "rtp"
	if _, ok := d.GetOk(pre_append); ok {

		result["rtp"], _ = expandVoipProfileSipRtp(d, i["rtp"], pre_append, sv)
	}
	pre_append = pre + ".0." + "nat_port_range"
	if _, ok := d.GetOk(pre_append); ok {

		result["nat-port-range"], _ = expandVoipProfileSipNatPortRange(d, i["nat_port_range"], pre_append, sv)
	}
	pre_append = pre + ".0." + "open_register_pinhole"
	if _, ok := d.GetOk(pre_append); ok {

		result["open-register-pinhole"], _ = expandVoipProfileSipOpenRegisterPinhole(d, i["open_register_pinhole"], pre_append, sv)
	}
	pre_append = pre + ".0." + "open_contact_pinhole"
	if _, ok := d.GetOk(pre_append); ok {

		result["open-contact-pinhole"], _ = expandVoipProfileSipOpenContactPinhole(d, i["open_contact_pinhole"], pre_append, sv)
	}
	pre_append = pre + ".0." + "strict_register"
	if _, ok := d.GetOk(pre_append); ok {

		result["strict-register"], _ = expandVoipProfileSipStrictRegister(d, i["strict_register"], pre_append, sv)
	}
	pre_append = pre + ".0." + "register_rate"
	if _, ok := d.GetOk(pre_append); ok {

		result["register-rate"], _ = expandVoipProfileSipRegisterRate(d, i["register_rate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "register_rate_track"
	if _, ok := d.GetOk(pre_append); ok {

		result["register-rate-track"], _ = expandVoipProfileSipRegisterRateTrack(d, i["register_rate_track"], pre_append, sv)
	}
	pre_append = pre + ".0." + "invite_rate"
	if _, ok := d.GetOk(pre_append); ok {

		result["invite-rate"], _ = expandVoipProfileSipInviteRate(d, i["invite_rate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "invite_rate_track"
	if _, ok := d.GetOk(pre_append); ok {

		result["invite-rate-track"], _ = expandVoipProfileSipInviteRateTrack(d, i["invite_rate_track"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_dialogs"
	if _, ok := d.GetOk(pre_append); ok {

		result["max-dialogs"], _ = expandVoipProfileSipMaxDialogs(d, i["max_dialogs"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_line_length"
	if _, ok := d.GetOk(pre_append); ok {

		result["max-line-length"], _ = expandVoipProfileSipMaxLineLength(d, i["max_line_length"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block_long_lines"
	if _, ok := d.GetOk(pre_append); ok {

		result["block-long-lines"], _ = expandVoipProfileSipBlockLongLines(d, i["block_long_lines"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block_unknown"
	if _, ok := d.GetOk(pre_append); ok {

		result["block-unknown"], _ = expandVoipProfileSipBlockUnknown(d, i["block_unknown"], pre_append, sv)
	}
	pre_append = pre + ".0." + "call_keepalive"
	if _, ok := d.GetOk(pre_append); ok {

		result["call-keepalive"], _ = expandVoipProfileSipCallKeepalive(d, i["call_keepalive"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block_ack"
	if _, ok := d.GetOk(pre_append); ok {

		result["block-ack"], _ = expandVoipProfileSipBlockAck(d, i["block_ack"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block_bye"
	if _, ok := d.GetOk(pre_append); ok {

		result["block-bye"], _ = expandVoipProfileSipBlockBye(d, i["block_bye"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block_cancel"
	if _, ok := d.GetOk(pre_append); ok {

		result["block-cancel"], _ = expandVoipProfileSipBlockCancel(d, i["block_cancel"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block_info"
	if _, ok := d.GetOk(pre_append); ok {

		result["block-info"], _ = expandVoipProfileSipBlockInfo(d, i["block_info"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block_invite"
	if _, ok := d.GetOk(pre_append); ok {

		result["block-invite"], _ = expandVoipProfileSipBlockInvite(d, i["block_invite"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block_message"
	if _, ok := d.GetOk(pre_append); ok {

		result["block-message"], _ = expandVoipProfileSipBlockMessage(d, i["block_message"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block_notify"
	if _, ok := d.GetOk(pre_append); ok {

		result["block-notify"], _ = expandVoipProfileSipBlockNotify(d, i["block_notify"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block_options"
	if _, ok := d.GetOk(pre_append); ok {

		result["block-options"], _ = expandVoipProfileSipBlockOptions(d, i["block_options"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block_prack"
	if _, ok := d.GetOk(pre_append); ok {

		result["block-prack"], _ = expandVoipProfileSipBlockPrack(d, i["block_prack"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block_publish"
	if _, ok := d.GetOk(pre_append); ok {

		result["block-publish"], _ = expandVoipProfileSipBlockPublish(d, i["block_publish"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block_refer"
	if _, ok := d.GetOk(pre_append); ok {

		result["block-refer"], _ = expandVoipProfileSipBlockRefer(d, i["block_refer"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block_register"
	if _, ok := d.GetOk(pre_append); ok {

		result["block-register"], _ = expandVoipProfileSipBlockRegister(d, i["block_register"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block_subscribe"
	if _, ok := d.GetOk(pre_append); ok {

		result["block-subscribe"], _ = expandVoipProfileSipBlockSubscribe(d, i["block_subscribe"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block_update"
	if _, ok := d.GetOk(pre_append); ok {

		result["block-update"], _ = expandVoipProfileSipBlockUpdate(d, i["block_update"], pre_append, sv)
	}
	pre_append = pre + ".0." + "register_contact_trace"
	if _, ok := d.GetOk(pre_append); ok {

		result["register-contact-trace"], _ = expandVoipProfileSipRegisterContactTrace(d, i["register_contact_trace"], pre_append, sv)
	}
	pre_append = pre + ".0." + "open_via_pinhole"
	if _, ok := d.GetOk(pre_append); ok {

		result["open-via-pinhole"], _ = expandVoipProfileSipOpenViaPinhole(d, i["open_via_pinhole"], pre_append, sv)
	}
	pre_append = pre + ".0." + "open_record_route_pinhole"
	if _, ok := d.GetOk(pre_append); ok {

		result["open-record-route-pinhole"], _ = expandVoipProfileSipOpenRecordRoutePinhole(d, i["open_record_route_pinhole"], pre_append, sv)
	}
	pre_append = pre + ".0." + "rfc2543_branch"
	if _, ok := d.GetOk(pre_append); ok {

		result["rfc2543-branch"], _ = expandVoipProfileSipRfc2543Branch(d, i["rfc2543_branch"], pre_append, sv)
	}
	pre_append = pre + ".0." + "log_violations"
	if _, ok := d.GetOk(pre_append); ok {

		result["log-violations"], _ = expandVoipProfileSipLogViolations(d, i["log_violations"], pre_append, sv)
	}
	pre_append = pre + ".0." + "log_call_summary"
	if _, ok := d.GetOk(pre_append); ok {

		result["log-call-summary"], _ = expandVoipProfileSipLogCallSummary(d, i["log_call_summary"], pre_append, sv)
	}
	pre_append = pre + ".0." + "nat_trace"
	if _, ok := d.GetOk(pre_append); ok {

		result["nat-trace"], _ = expandVoipProfileSipNatTrace(d, i["nat_trace"], pre_append, sv)
	}
	pre_append = pre + ".0." + "subscribe_rate"
	if _, ok := d.GetOk(pre_append); ok {

		result["subscribe-rate"], _ = expandVoipProfileSipSubscribeRate(d, i["subscribe_rate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "subscribe_rate_track"
	if _, ok := d.GetOk(pre_append); ok {

		result["subscribe-rate-track"], _ = expandVoipProfileSipSubscribeRateTrack(d, i["subscribe_rate_track"], pre_append, sv)
	}
	pre_append = pre + ".0." + "message_rate"
	if _, ok := d.GetOk(pre_append); ok {

		result["message-rate"], _ = expandVoipProfileSipMessageRate(d, i["message_rate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "message_rate_track"
	if _, ok := d.GetOk(pre_append); ok {

		result["message-rate-track"], _ = expandVoipProfileSipMessageRateTrack(d, i["message_rate_track"], pre_append, sv)
	}
	pre_append = pre + ".0." + "notify_rate"
	if _, ok := d.GetOk(pre_append); ok {

		result["notify-rate"], _ = expandVoipProfileSipNotifyRate(d, i["notify_rate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "notify_rate_track"
	if _, ok := d.GetOk(pre_append); ok {

		result["notify-rate-track"], _ = expandVoipProfileSipNotifyRateTrack(d, i["notify_rate_track"], pre_append, sv)
	}
	pre_append = pre + ".0." + "refer_rate"
	if _, ok := d.GetOk(pre_append); ok {

		result["refer-rate"], _ = expandVoipProfileSipReferRate(d, i["refer_rate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "refer_rate_track"
	if _, ok := d.GetOk(pre_append); ok {

		result["refer-rate-track"], _ = expandVoipProfileSipReferRateTrack(d, i["refer_rate_track"], pre_append, sv)
	}
	pre_append = pre + ".0." + "update_rate"
	if _, ok := d.GetOk(pre_append); ok {

		result["update-rate"], _ = expandVoipProfileSipUpdateRate(d, i["update_rate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "update_rate_track"
	if _, ok := d.GetOk(pre_append); ok {

		result["update-rate-track"], _ = expandVoipProfileSipUpdateRateTrack(d, i["update_rate_track"], pre_append, sv)
	}
	pre_append = pre + ".0." + "options_rate"
	if _, ok := d.GetOk(pre_append); ok {

		result["options-rate"], _ = expandVoipProfileSipOptionsRate(d, i["options_rate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "options_rate_track"
	if _, ok := d.GetOk(pre_append); ok {

		result["options-rate-track"], _ = expandVoipProfileSipOptionsRateTrack(d, i["options_rate_track"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ack_rate"
	if _, ok := d.GetOk(pre_append); ok {

		result["ack-rate"], _ = expandVoipProfileSipAckRate(d, i["ack_rate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ack_rate_track"
	if _, ok := d.GetOk(pre_append); ok {

		result["ack-rate-track"], _ = expandVoipProfileSipAckRateTrack(d, i["ack_rate_track"], pre_append, sv)
	}
	pre_append = pre + ".0." + "prack_rate"
	if _, ok := d.GetOk(pre_append); ok {

		result["prack-rate"], _ = expandVoipProfileSipPrackRate(d, i["prack_rate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "prack_rate_track"
	if _, ok := d.GetOk(pre_append); ok {

		result["prack-rate-track"], _ = expandVoipProfileSipPrackRateTrack(d, i["prack_rate_track"], pre_append, sv)
	}
	pre_append = pre + ".0." + "info_rate"
	if _, ok := d.GetOk(pre_append); ok {

		result["info-rate"], _ = expandVoipProfileSipInfoRate(d, i["info_rate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "info_rate_track"
	if _, ok := d.GetOk(pre_append); ok {

		result["info-rate-track"], _ = expandVoipProfileSipInfoRateTrack(d, i["info_rate_track"], pre_append, sv)
	}
	pre_append = pre + ".0." + "publish_rate"
	if _, ok := d.GetOk(pre_append); ok {

		result["publish-rate"], _ = expandVoipProfileSipPublishRate(d, i["publish_rate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "publish_rate_track"
	if _, ok := d.GetOk(pre_append); ok {

		result["publish-rate-track"], _ = expandVoipProfileSipPublishRateTrack(d, i["publish_rate_track"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bye_rate"
	if _, ok := d.GetOk(pre_append); ok {

		result["bye-rate"], _ = expandVoipProfileSipByeRate(d, i["bye_rate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bye_rate_track"
	if _, ok := d.GetOk(pre_append); ok {

		result["bye-rate-track"], _ = expandVoipProfileSipByeRateTrack(d, i["bye_rate_track"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cancel_rate"
	if _, ok := d.GetOk(pre_append); ok {

		result["cancel-rate"], _ = expandVoipProfileSipCancelRate(d, i["cancel_rate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cancel_rate_track"
	if _, ok := d.GetOk(pre_append); ok {

		result["cancel-rate-track"], _ = expandVoipProfileSipCancelRateTrack(d, i["cancel_rate_track"], pre_append, sv)
	}
	pre_append = pre + ".0." + "preserve_override"
	if _, ok := d.GetOk(pre_append); ok {

		result["preserve-override"], _ = expandVoipProfileSipPreserveOverride(d, i["preserve_override"], pre_append, sv)
	}
	pre_append = pre + ".0." + "no_sdp_fixup"
	if _, ok := d.GetOk(pre_append); ok {

		result["no-sdp-fixup"], _ = expandVoipProfileSipNoSdpFixup(d, i["no_sdp_fixup"], pre_append, sv)
	}
	pre_append = pre + ".0." + "contact_fixup"
	if _, ok := d.GetOk(pre_append); ok {

		result["contact-fixup"], _ = expandVoipProfileSipContactFixup(d, i["contact_fixup"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_idle_dialogs"
	if _, ok := d.GetOk(pre_append); ok {

		result["max-idle-dialogs"], _ = expandVoipProfileSipMaxIdleDialogs(d, i["max_idle_dialogs"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block_geo_red_options"
	if _, ok := d.GetOk(pre_append); ok {

		result["block-geo-red-options"], _ = expandVoipProfileSipBlockGeoRedOptions(d, i["block_geo_red_options"], pre_append, sv)
	}
	pre_append = pre + ".0." + "hosted_nat_traversal"
	if _, ok := d.GetOk(pre_append); ok {

		result["hosted-nat-traversal"], _ = expandVoipProfileSipHostedNatTraversal(d, i["hosted_nat_traversal"], pre_append, sv)
	}
	pre_append = pre + ".0." + "hnt_restrict_source_ip"
	if _, ok := d.GetOk(pre_append); ok {

		result["hnt-restrict-source-ip"], _ = expandVoipProfileSipHntRestrictSourceIp(d, i["hnt_restrict_source_ip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_body_length"
	if _, ok := d.GetOk(pre_append); ok {

		result["max-body-length"], _ = expandVoipProfileSipMaxBodyLength(d, i["max_body_length"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unknown_header"
	if _, ok := d.GetOk(pre_append); ok {

		result["unknown-header"], _ = expandVoipProfileSipUnknownHeader(d, i["unknown_header"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_request_line"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-request-line"], _ = expandVoipProfileSipMalformedRequestLine(d, i["malformed_request_line"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_via"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-via"], _ = expandVoipProfileSipMalformedHeaderVia(d, i["malformed_header_via"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_from"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-from"], _ = expandVoipProfileSipMalformedHeaderFrom(d, i["malformed_header_from"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_to"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-to"], _ = expandVoipProfileSipMalformedHeaderTo(d, i["malformed_header_to"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_call_id"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-call-id"], _ = expandVoipProfileSipMalformedHeaderCallId(d, i["malformed_header_call_id"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_cseq"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-cseq"], _ = expandVoipProfileSipMalformedHeaderCseq(d, i["malformed_header_cseq"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_rack"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-rack"], _ = expandVoipProfileSipMalformedHeaderRack(d, i["malformed_header_rack"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_rseq"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-rseq"], _ = expandVoipProfileSipMalformedHeaderRseq(d, i["malformed_header_rseq"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_contact"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-contact"], _ = expandVoipProfileSipMalformedHeaderContact(d, i["malformed_header_contact"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_record_route"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-record-route"], _ = expandVoipProfileSipMalformedHeaderRecordRoute(d, i["malformed_header_record_route"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_route"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-route"], _ = expandVoipProfileSipMalformedHeaderRoute(d, i["malformed_header_route"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_expires"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-expires"], _ = expandVoipProfileSipMalformedHeaderExpires(d, i["malformed_header_expires"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_content_type"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-content-type"], _ = expandVoipProfileSipMalformedHeaderContentType(d, i["malformed_header_content_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_content_length"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-content-length"], _ = expandVoipProfileSipMalformedHeaderContentLength(d, i["malformed_header_content_length"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_max_forwards"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-max-forwards"], _ = expandVoipProfileSipMalformedHeaderMaxForwards(d, i["malformed_header_max_forwards"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_allow"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-allow"], _ = expandVoipProfileSipMalformedHeaderAllow(d, i["malformed_header_allow"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_p_asserted_identity"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-p-asserted-identity"], _ = expandVoipProfileSipMalformedHeaderPAssertedIdentity(d, i["malformed_header_p_asserted_identity"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_no_require"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-no-require"], _ = expandVoipProfileSipMalformedHeaderNoRequire(d, i["malformed_header_no_require"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_no_proxy_require"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-no-proxy-require"], _ = expandVoipProfileSipMalformedHeaderNoProxyRequire(d, i["malformed_header_no_proxy_require"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_v"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-sdp-v"], _ = expandVoipProfileSipMalformedHeaderSdpV(d, i["malformed_header_sdp_v"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_o"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-sdp-o"], _ = expandVoipProfileSipMalformedHeaderSdpO(d, i["malformed_header_sdp_o"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_s"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-sdp-s"], _ = expandVoipProfileSipMalformedHeaderSdpS(d, i["malformed_header_sdp_s"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_i"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-sdp-i"], _ = expandVoipProfileSipMalformedHeaderSdpI(d, i["malformed_header_sdp_i"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_c"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-sdp-c"], _ = expandVoipProfileSipMalformedHeaderSdpC(d, i["malformed_header_sdp_c"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_b"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-sdp-b"], _ = expandVoipProfileSipMalformedHeaderSdpB(d, i["malformed_header_sdp_b"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_z"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-sdp-z"], _ = expandVoipProfileSipMalformedHeaderSdpZ(d, i["malformed_header_sdp_z"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_k"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-sdp-k"], _ = expandVoipProfileSipMalformedHeaderSdpK(d, i["malformed_header_sdp_k"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_a"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-sdp-a"], _ = expandVoipProfileSipMalformedHeaderSdpA(d, i["malformed_header_sdp_a"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_t"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-sdp-t"], _ = expandVoipProfileSipMalformedHeaderSdpT(d, i["malformed_header_sdp_t"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_r"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-sdp-r"], _ = expandVoipProfileSipMalformedHeaderSdpR(d, i["malformed_header_sdp_r"], pre_append, sv)
	}
	pre_append = pre + ".0." + "malformed_header_sdp_m"
	if _, ok := d.GetOk(pre_append); ok {

		result["malformed-header-sdp-m"], _ = expandVoipProfileSipMalformedHeaderSdpM(d, i["malformed_header_sdp_m"], pre_append, sv)
	}
	pre_append = pre + ".0." + "provisional_invite_expiry_time"
	if _, ok := d.GetOk(pre_append); ok {

		result["provisional-invite-expiry-time"], _ = expandVoipProfileSipProvisionalInviteExpiryTime(d, i["provisional_invite_expiry_time"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ips_rtp"
	if _, ok := d.GetOk(pre_append); ok {

		result["ips-rtp"], _ = expandVoipProfileSipIpsRtp(d, i["ips_rtp"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ssl_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["ssl-mode"], _ = expandVoipProfileSipSslMode(d, i["ssl_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ssl_send_empty_frags"
	if _, ok := d.GetOk(pre_append); ok {

		result["ssl-send-empty-frags"], _ = expandVoipProfileSipSslSendEmptyFrags(d, i["ssl_send_empty_frags"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ssl_client_renegotiation"
	if _, ok := d.GetOk(pre_append); ok {

		result["ssl-client-renegotiation"], _ = expandVoipProfileSipSslClientRenegotiation(d, i["ssl_client_renegotiation"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ssl_algorithm"
	if _, ok := d.GetOk(pre_append); ok {

		result["ssl-algorithm"], _ = expandVoipProfileSipSslAlgorithm(d, i["ssl_algorithm"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ssl_pfs"
	if _, ok := d.GetOk(pre_append); ok {

		result["ssl-pfs"], _ = expandVoipProfileSipSslPfs(d, i["ssl_pfs"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ssl_min_version"
	if _, ok := d.GetOk(pre_append); ok {

		result["ssl-min-version"], _ = expandVoipProfileSipSslMinVersion(d, i["ssl_min_version"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ssl_max_version"
	if _, ok := d.GetOk(pre_append); ok {

		result["ssl-max-version"], _ = expandVoipProfileSipSslMaxVersion(d, i["ssl_max_version"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ssl_client_certificate"
	if _, ok := d.GetOk(pre_append); ok {

		result["ssl-client-certificate"], _ = expandVoipProfileSipSslClientCertificate(d, i["ssl_client_certificate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ssl_server_certificate"
	if _, ok := d.GetOk(pre_append); ok {

		result["ssl-server-certificate"], _ = expandVoipProfileSipSslServerCertificate(d, i["ssl_server_certificate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ssl_auth_client"
	if _, ok := d.GetOk(pre_append); ok {

		result["ssl-auth-client"], _ = expandVoipProfileSipSslAuthClient(d, i["ssl_auth_client"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ssl_auth_server"
	if _, ok := d.GetOk(pre_append); ok {

		result["ssl-auth-server"], _ = expandVoipProfileSipSslAuthServer(d, i["ssl_auth_server"], pre_append, sv)
	}

	return result, nil
}

func expandVoipProfileSipStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipRtp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipNatPortRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipOpenRegisterPinhole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipOpenContactPinhole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipStrictRegister(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipRegisterRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipRegisterRateTrack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipInviteRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipInviteRateTrack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMaxDialogs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMaxLineLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockLongLines(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockUnknown(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipCallKeepalive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockAck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockBye(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockCancel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockInfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockInvite(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockMessage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockNotify(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockPrack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockPublish(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockRefer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockRegister(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockSubscribe(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockUpdate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipRegisterContactTrace(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipOpenViaPinhole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipOpenRecordRoutePinhole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipRfc2543Branch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipLogViolations(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipLogCallSummary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipNatTrace(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSubscribeRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSubscribeRateTrack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMessageRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMessageRateTrack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipNotifyRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipNotifyRateTrack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipReferRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipReferRateTrack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipUpdateRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipUpdateRateTrack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipOptionsRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipOptionsRateTrack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipAckRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipAckRateTrack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipPrackRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipPrackRateTrack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipInfoRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipInfoRateTrack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipPublishRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipPublishRateTrack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipByeRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipByeRateTrack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipCancelRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipCancelRateTrack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipPreserveOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipNoSdpFixup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipContactFixup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMaxIdleDialogs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipBlockGeoRedOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipHostedNatTraversal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipHntRestrictSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMaxBodyLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipUnknownHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedRequestLine(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderVia(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderFrom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderTo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderCallId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderCseq(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderRack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderRseq(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderContact(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderRecordRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderExpires(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderContentType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderContentLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderMaxForwards(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderAllow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderPAssertedIdentity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderNoRequire(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderNoProxyRequire(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpV(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpO(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpS(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpI(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpC(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpB(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpZ(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpK(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpA(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpT(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpR(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipMalformedHeaderSdpM(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipProvisionalInviteExpiryTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipIpsRtp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSslMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSslSendEmptyFrags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSslClientRenegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSslAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSslPfs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSslMinVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSslMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSslClientCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSslServerCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSslAuthClient(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSipSslAuthServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSccp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {

		result["status"], _ = expandVoipProfileSccpStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block_mcast"
	if _, ok := d.GetOk(pre_append); ok {

		result["block-mcast"], _ = expandVoipProfileSccpBlockMcast(d, i["block_mcast"], pre_append, sv)
	}
	pre_append = pre + ".0." + "verify_header"
	if _, ok := d.GetOk(pre_append); ok {

		result["verify-header"], _ = expandVoipProfileSccpVerifyHeader(d, i["verify_header"], pre_append, sv)
	}
	pre_append = pre + ".0." + "log_call_summary"
	if _, ok := d.GetOk(pre_append); ok {

		result["log-call-summary"], _ = expandVoipProfileSccpLogCallSummary(d, i["log_call_summary"], pre_append, sv)
	}
	pre_append = pre + ".0." + "log_violations"
	if _, ok := d.GetOk(pre_append); ok {

		result["log-violations"], _ = expandVoipProfileSccpLogViolations(d, i["log_violations"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_calls"
	if _, ok := d.GetOk(pre_append); ok {

		result["max-calls"], _ = expandVoipProfileSccpMaxCalls(d, i["max_calls"], pre_append, sv)
	}

	return result, nil
}

func expandVoipProfileSccpStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSccpBlockMcast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSccpVerifyHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSccpLogCallSummary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSccpLogViolations(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSccpMaxCalls(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileMsrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {

		result["status"], _ = expandVoipProfileMsrpStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "log_violations"
	if _, ok := d.GetOk(pre_append); ok {

		result["log-violations"], _ = expandVoipProfileMsrpLogViolations(d, i["log_violations"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_msg_size"
	if _, ok := d.GetOk(pre_append); ok {

		result["max-msg-size"], _ = expandVoipProfileMsrpMaxMsgSize(d, i["max_msg_size"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_msg_size_action"
	if _, ok := d.GetOk(pre_append); ok {

		result["max-msg-size-action"], _ = expandVoipProfileMsrpMaxMsgSizeAction(d, i["max_msg_size_action"], pre_append, sv)
	}

	return result, nil
}

func expandVoipProfileMsrpStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileMsrpLogViolations(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileMsrpMaxMsgSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileMsrpMaxMsgSizeAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVoipProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandVoipProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("feature_set"); ok {

		t, err := expandVoipProfileFeatureSet(d, v, "feature_set", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["feature-set"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandVoipProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("sip"); ok {

		t, err := expandVoipProfileSip(d, v, "sip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip"] = t
		}
	}

	if v, ok := d.GetOk("sccp"); ok {

		t, err := expandVoipProfileSccp(d, v, "sccp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sccp"] = t
		}
	}

	if v, ok := d.GetOk("msrp"); ok {

		t, err := expandVoipProfileMsrp(d, v, "msrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msrp"] = t
		}
	}

	return &obj, nil
}
