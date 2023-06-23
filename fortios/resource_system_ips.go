// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPS system settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemIps() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemIpsUpdate,
		Read:   resourceSystemIpsRead,
		Update: resourceSystemIpsUpdate,
		Delete: resourceSystemIpsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"signature_hold_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_signature_hold_by_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemIpsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemIps(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemIps resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemIps(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemIps resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemIps")
	}

	return resourceSystemIpsRead(d, m)
}

func resourceSystemIpsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemIps(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemIps resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemIps(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemIps resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemIps(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemIps resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIps(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemIps resource from API: %v", err)
	}
	return nil
}

func flattenSystemIpsSignatureHoldTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIpsOverrideSignatureHoldById(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemIps(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("signature_hold_time", flattenSystemIpsSignatureHoldTime(o["signature-hold-time"], d, "signature_hold_time", sv)); err != nil {
		if !fortiAPIPatch(o["signature-hold-time"]) {
			return fmt.Errorf("Error reading signature_hold_time: %v", err)
		}
	}

	if err = d.Set("override_signature_hold_by_id", flattenSystemIpsOverrideSignatureHoldById(o["override-signature-hold-by-id"], d, "override_signature_hold_by_id", sv)); err != nil {
		if !fortiAPIPatch(o["override-signature-hold-by-id"]) {
			return fmt.Errorf("Error reading override_signature_hold_by_id: %v", err)
		}
	}

	return nil
}

func flattenSystemIpsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemIpsSignatureHoldTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIpsOverrideSignatureHoldById(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIps(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("signature_hold_time"); ok {
		if setArgNil {
			obj["signature-hold-time"] = nil
		} else {
			t, err := expandSystemIpsSignatureHoldTime(d, v, "signature_hold_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["signature-hold-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("override_signature_hold_by_id"); ok {
		if setArgNil {
			obj["override-signature-hold-by-id"] = nil
		} else {
			t, err := expandSystemIpsOverrideSignatureHoldById(d, v, "override_signature_hold_by_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["override-signature-hold-by-id"] = t
			}
		}
	}

	return &obj, nil
}
