// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPS custom signature.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceIpsCustom() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpsCustomCreate,
		Read:   resourceIpsCustomRead,
		Update: resourceIpsCustomUpdate,
		Delete: resourceIpsCustomDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"tag": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"signature": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
				Computed:     true,
			},
			"sig_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"rule_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"location": &schema.Schema{
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
			"protocol": &schema.Schema{
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
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceIpsCustomCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectIpsCustom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating IpsCustom resource while getting object: %v", err)
	}

	o, err := c.CreateIpsCustom(obj)

	if err != nil {
		return fmt.Errorf("Error creating IpsCustom resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("IpsCustom")
	}

	return resourceIpsCustomRead(d, m)
}

func resourceIpsCustomUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectIpsCustom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating IpsCustom resource while getting object: %v", err)
	}

	o, err := c.UpdateIpsCustom(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating IpsCustom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("IpsCustom")
	}

	return resourceIpsCustomRead(d, m)
}

func resourceIpsCustomDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteIpsCustom(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting IpsCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsCustomRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadIpsCustom(mkey)
	if err != nil {
		return fmt.Errorf("Error reading IpsCustom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIpsCustom(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading IpsCustom resource from API: %v", err)
	}
	return nil
}

func flattenIpsCustomTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsCustomSignature(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsCustomSigName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsCustomRuleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsCustomSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsCustomLocation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsCustomOs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsCustomApplication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsCustomProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsCustomStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsCustomLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsCustomLogPacket(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsCustomAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsCustomComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectIpsCustom(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("tag", flattenIpsCustomTag(o["tag"], d, "tag", sv)); err != nil {
		if !fortiAPIPatch(o["tag"]) {
			return fmt.Errorf("Error reading tag: %v", err)
		}
	}

	if err = d.Set("signature", flattenIpsCustomSignature(o["signature"], d, "signature", sv)); err != nil {
		if !fortiAPIPatch(o["signature"]) {
			return fmt.Errorf("Error reading signature: %v", err)
		}
	}

	if err = d.Set("sig_name", flattenIpsCustomSigName(o["sig-name"], d, "sig_name", sv)); err != nil {
		if !fortiAPIPatch(o["sig-name"]) {
			return fmt.Errorf("Error reading sig_name: %v", err)
		}
	}

	if err = d.Set("rule_id", flattenIpsCustomRuleId(o["rule-id"], d, "rule_id", sv)); err != nil {
		if !fortiAPIPatch(o["rule-id"]) {
			return fmt.Errorf("Error reading rule_id: %v", err)
		}
	}

	if err = d.Set("severity", flattenIpsCustomSeverity(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("location", flattenIpsCustomLocation(o["location"], d, "location", sv)); err != nil {
		if !fortiAPIPatch(o["location"]) {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("os", flattenIpsCustomOs(o["os"], d, "os", sv)); err != nil {
		if !fortiAPIPatch(o["os"]) {
			return fmt.Errorf("Error reading os: %v", err)
		}
	}

	if err = d.Set("application", flattenIpsCustomApplication(o["application"], d, "application", sv)); err != nil {
		if !fortiAPIPatch(o["application"]) {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("protocol", flattenIpsCustomProtocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("status", flattenIpsCustomStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("log", flattenIpsCustomLog(o["log"], d, "log", sv)); err != nil {
		if !fortiAPIPatch(o["log"]) {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("log_packet", flattenIpsCustomLogPacket(o["log-packet"], d, "log_packet", sv)); err != nil {
		if !fortiAPIPatch(o["log-packet"]) {
			return fmt.Errorf("Error reading log_packet: %v", err)
		}
	}

	if err = d.Set("action", flattenIpsCustomAction(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("comment", flattenIpsCustomComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	return nil
}

func flattenIpsCustomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandIpsCustomTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomSignature(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomSigName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomRuleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomLocation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomOs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomLogPacket(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsCustomComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectIpsCustom(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("tag"); ok {

		t, err := expandIpsCustomTag(d, v, "tag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag"] = t
		}
	}

	if v, ok := d.GetOk("signature"); ok {

		t, err := expandIpsCustomSignature(d, v, "signature", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signature"] = t
		}
	}

	if v, ok := d.GetOk("sig_name"); ok {

		t, err := expandIpsCustomSigName(d, v, "sig_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sig-name"] = t
		}
	}

	if v, ok := d.GetOkExists("rule_id"); ok {

		t, err := expandIpsCustomRuleId(d, v, "rule_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule-id"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok {

		t, err := expandIpsCustomSeverity(d, v, "severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("location"); ok {

		t, err := expandIpsCustomLocation(d, v, "location", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location"] = t
		}
	}

	if v, ok := d.GetOk("os"); ok {

		t, err := expandIpsCustomOs(d, v, "os", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok {

		t, err := expandIpsCustomApplication(d, v, "application", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok {

		t, err := expandIpsCustomProtocol(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandIpsCustomStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok {

		t, err := expandIpsCustomLog(d, v, "log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("log_packet"); ok {

		t, err := expandIpsCustomLogPacket(d, v, "log_packet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-packet"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {

		t, err := expandIpsCustomAction(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandIpsCustomComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	return &obj, nil
}
