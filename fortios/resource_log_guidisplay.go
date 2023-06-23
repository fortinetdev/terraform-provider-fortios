// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure how log messages are displayed on the GUI.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogGuiDisplay() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogGuiDisplayUpdate,
		Read:   resourceLogGuiDisplayRead,
		Update: resourceLogGuiDisplayUpdate,
		Delete: resourceLogGuiDisplayDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"resolve_hosts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resolve_apps": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiview_unscanned_apps": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogGuiDisplayUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogGuiDisplay(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogGuiDisplay resource while getting object: %v", err)
	}

	o, err := c.UpdateLogGuiDisplay(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogGuiDisplay resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogGuiDisplay")
	}

	return resourceLogGuiDisplayRead(d, m)
}

func resourceLogGuiDisplayDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogGuiDisplay(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogGuiDisplay resource while getting object: %v", err)
	}

	_, err = c.UpdateLogGuiDisplay(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing LogGuiDisplay resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogGuiDisplayRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadLogGuiDisplay(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogGuiDisplay resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogGuiDisplay(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogGuiDisplay resource from API: %v", err)
	}
	return nil
}

func flattenLogGuiDisplayResolveHosts(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogGuiDisplayResolveApps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogGuiDisplayFortiviewUnscannedApps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogGuiDisplay(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("resolve_hosts", flattenLogGuiDisplayResolveHosts(o["resolve-hosts"], d, "resolve_hosts", sv)); err != nil {
		if !fortiAPIPatch(o["resolve-hosts"]) {
			return fmt.Errorf("Error reading resolve_hosts: %v", err)
		}
	}

	if err = d.Set("resolve_apps", flattenLogGuiDisplayResolveApps(o["resolve-apps"], d, "resolve_apps", sv)); err != nil {
		if !fortiAPIPatch(o["resolve-apps"]) {
			return fmt.Errorf("Error reading resolve_apps: %v", err)
		}
	}

	if err = d.Set("fortiview_unscanned_apps", flattenLogGuiDisplayFortiviewUnscannedApps(o["fortiview-unscanned-apps"], d, "fortiview_unscanned_apps", sv)); err != nil {
		if !fortiAPIPatch(o["fortiview-unscanned-apps"]) {
			return fmt.Errorf("Error reading fortiview_unscanned_apps: %v", err)
		}
	}

	return nil
}

func flattenLogGuiDisplayFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogGuiDisplayResolveHosts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogGuiDisplayResolveApps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogGuiDisplayFortiviewUnscannedApps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogGuiDisplay(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("resolve_hosts"); ok {
		if setArgNil {
			obj["resolve-hosts"] = nil
		} else {
			t, err := expandLogGuiDisplayResolveHosts(d, v, "resolve_hosts", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["resolve-hosts"] = t
			}
		}
	}

	if v, ok := d.GetOk("resolve_apps"); ok {
		if setArgNil {
			obj["resolve-apps"] = nil
		} else {
			t, err := expandLogGuiDisplayResolveApps(d, v, "resolve_apps", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["resolve-apps"] = t
			}
		}
	}

	if v, ok := d.GetOk("fortiview_unscanned_apps"); ok {
		if setArgNil {
			obj["fortiview-unscanned-apps"] = nil
		} else {
			t, err := expandLogGuiDisplayFortiviewUnscannedApps(d, v, "fortiview_unscanned_apps", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fortiview-unscanned-apps"] = t
			}
		}
	}

	return &obj, nil
}
