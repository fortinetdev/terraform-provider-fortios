// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Diameter filter profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDiameterFilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceDiameterFilterProfileCreate,
		Read:   resourceDiameterFilterProfileRead,
		Update: resourceDiameterFilterProfileUpdate,
		Delete: resourceDiameterFilterProfileDelete,

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
				Optional:     true,
				Computed:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"monitor_all_messages": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"track_requests_answers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"missing_request_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol_version_invalid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"message_length_invalid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_error_flag_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cmd_flags_reserve_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"command_code_invalid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"command_code_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceDiameterFilterProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectDiameterFilterProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating DiameterFilterProfile resource while getting object: %v", err)
	}

	o, err := c.CreateDiameterFilterProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating DiameterFilterProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DiameterFilterProfile")
	}

	return resourceDiameterFilterProfileRead(d, m)
}

func resourceDiameterFilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectDiameterFilterProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating DiameterFilterProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateDiameterFilterProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating DiameterFilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DiameterFilterProfile")
	}

	return resourceDiameterFilterProfileRead(d, m)
}

func resourceDiameterFilterProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteDiameterFilterProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting DiameterFilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDiameterFilterProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadDiameterFilterProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading DiameterFilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDiameterFilterProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading DiameterFilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenDiameterFilterProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDiameterFilterProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDiameterFilterProfileMonitorAllMessages(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDiameterFilterProfileLogPacket(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDiameterFilterProfileTrackRequestsAnswers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDiameterFilterProfileMissingRequestAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDiameterFilterProfileProtocolVersionInvalid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDiameterFilterProfileMessageLengthInvalid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDiameterFilterProfileRequestErrorFlagSet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDiameterFilterProfileCmdFlagsReserveSet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDiameterFilterProfileCommandCodeInvalid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDiameterFilterProfileCommandCodeRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectDiameterFilterProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenDiameterFilterProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenDiameterFilterProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("monitor_all_messages", flattenDiameterFilterProfileMonitorAllMessages(o["monitor-all-messages"], d, "monitor_all_messages", sv)); err != nil {
		if !fortiAPIPatch(o["monitor-all-messages"]) {
			return fmt.Errorf("Error reading monitor_all_messages: %v", err)
		}
	}

	if err = d.Set("log_packet", flattenDiameterFilterProfileLogPacket(o["log-packet"], d, "log_packet", sv)); err != nil {
		if !fortiAPIPatch(o["log-packet"]) {
			return fmt.Errorf("Error reading log_packet: %v", err)
		}
	}

	if err = d.Set("track_requests_answers", flattenDiameterFilterProfileTrackRequestsAnswers(o["track-requests-answers"], d, "track_requests_answers", sv)); err != nil {
		if !fortiAPIPatch(o["track-requests-answers"]) {
			return fmt.Errorf("Error reading track_requests_answers: %v", err)
		}
	}

	if err = d.Set("missing_request_action", flattenDiameterFilterProfileMissingRequestAction(o["missing-request-action"], d, "missing_request_action", sv)); err != nil {
		if !fortiAPIPatch(o["missing-request-action"]) {
			return fmt.Errorf("Error reading missing_request_action: %v", err)
		}
	}

	if err = d.Set("protocol_version_invalid", flattenDiameterFilterProfileProtocolVersionInvalid(o["protocol-version-invalid"], d, "protocol_version_invalid", sv)); err != nil {
		if !fortiAPIPatch(o["protocol-version-invalid"]) {
			return fmt.Errorf("Error reading protocol_version_invalid: %v", err)
		}
	}

	if err = d.Set("message_length_invalid", flattenDiameterFilterProfileMessageLengthInvalid(o["message-length-invalid"], d, "message_length_invalid", sv)); err != nil {
		if !fortiAPIPatch(o["message-length-invalid"]) {
			return fmt.Errorf("Error reading message_length_invalid: %v", err)
		}
	}

	if err = d.Set("request_error_flag_set", flattenDiameterFilterProfileRequestErrorFlagSet(o["request-error-flag-set"], d, "request_error_flag_set", sv)); err != nil {
		if !fortiAPIPatch(o["request-error-flag-set"]) {
			return fmt.Errorf("Error reading request_error_flag_set: %v", err)
		}
	}

	if err = d.Set("cmd_flags_reserve_set", flattenDiameterFilterProfileCmdFlagsReserveSet(o["cmd-flags-reserve-set"], d, "cmd_flags_reserve_set", sv)); err != nil {
		if !fortiAPIPatch(o["cmd-flags-reserve-set"]) {
			return fmt.Errorf("Error reading cmd_flags_reserve_set: %v", err)
		}
	}

	if err = d.Set("command_code_invalid", flattenDiameterFilterProfileCommandCodeInvalid(o["command-code-invalid"], d, "command_code_invalid", sv)); err != nil {
		if !fortiAPIPatch(o["command-code-invalid"]) {
			return fmt.Errorf("Error reading command_code_invalid: %v", err)
		}
	}

	if err = d.Set("command_code_range", flattenDiameterFilterProfileCommandCodeRange(o["command-code-range"], d, "command_code_range", sv)); err != nil {
		if !fortiAPIPatch(o["command-code-range"]) {
			return fmt.Errorf("Error reading command_code_range: %v", err)
		}
	}

	return nil
}

func flattenDiameterFilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandDiameterFilterProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileMonitorAllMessages(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileLogPacket(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileTrackRequestsAnswers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileMissingRequestAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileProtocolVersionInvalid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileMessageLengthInvalid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileRequestErrorFlagSet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileCmdFlagsReserveSet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileCommandCodeInvalid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDiameterFilterProfileCommandCodeRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectDiameterFilterProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandDiameterFilterProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandDiameterFilterProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("monitor_all_messages"); ok {
		t, err := expandDiameterFilterProfileMonitorAllMessages(d, v, "monitor_all_messages", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-all-messages"] = t
		}
	}

	if v, ok := d.GetOk("log_packet"); ok {
		t, err := expandDiameterFilterProfileLogPacket(d, v, "log_packet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-packet"] = t
		}
	}

	if v, ok := d.GetOk("track_requests_answers"); ok {
		t, err := expandDiameterFilterProfileTrackRequestsAnswers(d, v, "track_requests_answers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["track-requests-answers"] = t
		}
	}

	if v, ok := d.GetOk("missing_request_action"); ok {
		t, err := expandDiameterFilterProfileMissingRequestAction(d, v, "missing_request_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["missing-request-action"] = t
		}
	}

	if v, ok := d.GetOk("protocol_version_invalid"); ok {
		t, err := expandDiameterFilterProfileProtocolVersionInvalid(d, v, "protocol_version_invalid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol-version-invalid"] = t
		}
	}

	if v, ok := d.GetOk("message_length_invalid"); ok {
		t, err := expandDiameterFilterProfileMessageLengthInvalid(d, v, "message_length_invalid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-length-invalid"] = t
		}
	}

	if v, ok := d.GetOk("request_error_flag_set"); ok {
		t, err := expandDiameterFilterProfileRequestErrorFlagSet(d, v, "request_error_flag_set", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-error-flag-set"] = t
		}
	}

	if v, ok := d.GetOk("cmd_flags_reserve_set"); ok {
		t, err := expandDiameterFilterProfileCmdFlagsReserveSet(d, v, "cmd_flags_reserve_set", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmd-flags-reserve-set"] = t
		}
	}

	if v, ok := d.GetOk("command_code_invalid"); ok {
		t, err := expandDiameterFilterProfileCommandCodeInvalid(d, v, "command_code_invalid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command-code-invalid"] = t
		}
	}

	if v, ok := d.GetOk("command_code_range"); ok {
		t, err := expandDiameterFilterProfileCommandCodeRange(d, v, "command_code_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command-code-range"] = t
		}
	}

	return &obj, nil
}
