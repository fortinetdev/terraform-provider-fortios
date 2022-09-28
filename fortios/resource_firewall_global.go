// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Global firewall settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallGlobalUpdate,
		Read:   resourceFirewallGlobalRead,
		Update: resourceFirewallGlobalUpdate,
		Delete: resourceFirewallGlobalDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"banned_ip_persistency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallGlobalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallGlobal(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallGlobal resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallGlobal(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallGlobal")
	}

	return resourceFirewallGlobalRead(d, m)
}

func resourceFirewallGlobalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallGlobal(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating FirewallGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallGlobal(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing FirewallGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallGlobalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallGlobal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallGlobal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallGlobal resource from API: %v", err)
	}
	return nil
}

func flattenFirewallGlobalBannedIpPersistency(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallGlobal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("banned_ip_persistency", flattenFirewallGlobalBannedIpPersistency(o["banned-ip-persistency"], d, "banned_ip_persistency", sv)); err != nil {
		if !fortiAPIPatch(o["banned-ip-persistency"]) {
			return fmt.Errorf("Error reading banned_ip_persistency: %v", err)
		}
	}

	return nil
}

func flattenFirewallGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallGlobalBannedIpPersistency(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallGlobal(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("banned_ip_persistency"); ok {
		if setArgNil {
			obj["banned-ip-persistency"] = nil
		} else {

			t, err := expandFirewallGlobalBannedIpPersistency(d, v, "banned_ip_persistency", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["banned-ip-persistency"] = t
			}
		}
	}

	return &obj, nil
}
