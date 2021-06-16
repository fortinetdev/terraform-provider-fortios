// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure firewall IP-translation.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallIpTranslation() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallIpTranslationCreate,
		Read:   resourceFirewallIpTranslationRead,
		Update: resourceFirewallIpTranslationUpdate,
		Delete: resourceFirewallIpTranslationDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"transid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"startip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"endip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"map_startip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFirewallIpTranslationCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallIpTranslation(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallIpTranslation resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallIpTranslation(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallIpTranslation resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallIpTranslation")
	}

	return resourceFirewallIpTranslationRead(d, m)
}

func resourceFirewallIpTranslationUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallIpTranslation(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIpTranslation resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallIpTranslation(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIpTranslation resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallIpTranslation")
	}

	return resourceFirewallIpTranslationRead(d, m)
}

func resourceFirewallIpTranslationDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallIpTranslation(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallIpTranslation resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallIpTranslationRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallIpTranslation(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIpTranslation resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallIpTranslation(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIpTranslation resource from API: %v", err)
	}
	return nil
}

func flattenFirewallIpTranslationTransid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIpTranslationType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIpTranslationStartip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIpTranslationEndip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIpTranslationMapStartip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallIpTranslation(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("transid", flattenFirewallIpTranslationTransid(o["transid"], d, "transid", sv)); err != nil {
		if !fortiAPIPatch(o["transid"]) {
			return fmt.Errorf("Error reading transid: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallIpTranslationType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("startip", flattenFirewallIpTranslationStartip(o["startip"], d, "startip", sv)); err != nil {
		if !fortiAPIPatch(o["startip"]) {
			return fmt.Errorf("Error reading startip: %v", err)
		}
	}

	if err = d.Set("endip", flattenFirewallIpTranslationEndip(o["endip"], d, "endip", sv)); err != nil {
		if !fortiAPIPatch(o["endip"]) {
			return fmt.Errorf("Error reading endip: %v", err)
		}
	}

	if err = d.Set("map_startip", flattenFirewallIpTranslationMapStartip(o["map-startip"], d, "map_startip", sv)); err != nil {
		if !fortiAPIPatch(o["map-startip"]) {
			return fmt.Errorf("Error reading map_startip: %v", err)
		}
	}

	return nil
}

func flattenFirewallIpTranslationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallIpTranslationTransid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpTranslationType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpTranslationStartip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpTranslationEndip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpTranslationMapStartip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallIpTranslation(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("transid"); ok {

		t, err := expandFirewallIpTranslationTransid(d, v, "transid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transid"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {

		t, err := expandFirewallIpTranslationType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("startip"); ok {

		t, err := expandFirewallIpTranslationStartip(d, v, "startip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["startip"] = t
		}
	}

	if v, ok := d.GetOk("endip"); ok {

		t, err := expandFirewallIpTranslationEndip(d, v, "endip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endip"] = t
		}
	}

	if v, ok := d.GetOk("map_startip"); ok {

		t, err := expandFirewallIpTranslationMapStartip(d, v, "map_startip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["map-startip"] = t
		}
	}

	return &obj, nil
}
