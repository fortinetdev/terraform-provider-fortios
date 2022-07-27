// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure additional port mappings for Internet Services.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallInternetServiceAppend() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInternetServiceAppendUpdate,
		Read:   resourceFirewallInternetServiceAppendRead,
		Update: resourceFirewallInternetServiceAppendUpdate,
		Delete: resourceFirewallInternetServiceAppendDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"match_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"append_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceFirewallInternetServiceAppendUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallInternetServiceAppend(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceAppend resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallInternetServiceAppend(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceAppend resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallInternetServiceAppend")
	}

	return resourceFirewallInternetServiceAppendRead(d, m)
}

func resourceFirewallInternetServiceAppendDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallInternetServiceAppend(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceAppend resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallInternetServiceAppend(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing FirewallInternetServiceAppend resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceAppendRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallInternetServiceAppend(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceAppend resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInternetServiceAppend(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceAppend resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInternetServiceAppendMatchPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceAppendAppendPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallInternetServiceAppend(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("match_port", flattenFirewallInternetServiceAppendMatchPort(o["match-port"], d, "match_port", sv)); err != nil {
		if !fortiAPIPatch(o["match-port"]) {
			return fmt.Errorf("Error reading match_port: %v", err)
		}
	}

	if err = d.Set("append_port", flattenFirewallInternetServiceAppendAppendPort(o["append-port"], d, "append_port", sv)); err != nil {
		if !fortiAPIPatch(o["append-port"]) {
			return fmt.Errorf("Error reading append_port: %v", err)
		}
	}

	return nil
}

func flattenFirewallInternetServiceAppendFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallInternetServiceAppendMatchPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceAppendAppendPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInternetServiceAppend(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("match_port"); ok {
		if setArgNil {
			obj["match-port"] = nil
		} else {

			t, err := expandFirewallInternetServiceAppendMatchPort(d, v, "match_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["match-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("append_port"); ok {
		if setArgNil {
			obj["append-port"] = nil
		} else {

			t, err := expandFirewallInternetServiceAppendAppendPort(d, v, "append_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["append-port"] = t
			}
		}
	}

	return &obj, nil
}
