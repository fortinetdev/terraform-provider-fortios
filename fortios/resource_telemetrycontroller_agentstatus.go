// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: FortiTelemetry controller agent status.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceTelemetryControllerAgentStatus() *schema.Resource {
	return &schema.Resource{
		Create: resourceTelemetryControllerAgentStatusUpdate,
		Read:   resourceTelemetryControllerAgentStatusRead,
		Update: resourceTelemetryControllerAgentStatusUpdate,
		Delete: resourceTelemetryControllerAgentStatusDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"agent_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceTelemetryControllerAgentStatusUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectTelemetryControllerAgentStatus(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating TelemetryControllerAgentStatus resource while getting object: %v", err)
	}

	o, err := c.UpdateTelemetryControllerAgentStatus(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating TelemetryControllerAgentStatus resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("TelemetryControllerAgentStatus")
	}

	return resourceTelemetryControllerAgentStatusRead(d, m)
}

func resourceTelemetryControllerAgentStatusDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectTelemetryControllerAgentStatus(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating TelemetryControllerAgentStatus resource while getting object: %v", err)
	}

	_, err = c.UpdateTelemetryControllerAgentStatus(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing TelemetryControllerAgentStatus resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceTelemetryControllerAgentStatusRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadTelemetryControllerAgentStatus(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading TelemetryControllerAgentStatus resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectTelemetryControllerAgentStatus(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading TelemetryControllerAgentStatus resource from API: %v", err)
	}
	return nil
}

func flattenTelemetryControllerAgentStatusAgentId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectTelemetryControllerAgentStatus(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("agent_id", flattenTelemetryControllerAgentStatusAgentId(o["<agent-id>"], d, "agent_id", sv)); err != nil {
		if !fortiAPIPatch(o["<agent-id>"]) {
			return fmt.Errorf("Error reading agent_id: %v", err)
		}
	}

	return nil
}

func flattenTelemetryControllerAgentStatusFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandTelemetryControllerAgentStatusAgentId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectTelemetryControllerAgentStatus(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("agent_id"); ok {
		if setArgNil {
			obj["<agent-id>"] = nil
		} else {
			t, err := expandTelemetryControllerAgentStatusAgentId(d, v, "agent_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["<agent-id>"] = t
			}
		}
	} else if d.HasChange("agent_id") {
		obj["<agent-id>"] = nil
	}

	return &obj, nil
}
