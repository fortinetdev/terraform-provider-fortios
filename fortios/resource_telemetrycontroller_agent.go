// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiTelemetry agents managed by a FortiGate unit.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceTelemetryControllerAgent() *schema.Resource {
	return &schema.Resource{
		Create: resourceTelemetryControllerAgentCreate,
		Read:   resourceTelemetryControllerAgentRead,
		Update: resourceTelemetryControllerAgentUpdate,
		Delete: resourceTelemetryControllerAgentDelete,

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
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"agent_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"alias": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"authz": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"agent_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
		},
	}
}

func resourceTelemetryControllerAgentCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectTelemetryControllerAgent(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating TelemetryControllerAgent resource while getting object: %v", err)
	}

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("agent_id")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadTelemetryControllerAgent(mkey, vdomparam)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateTelemetryControllerAgent(obj, mkey, vdomparam)
			if err != nil {
				return fmt.Errorf("Error updating TelemetryControllerAgent resource: %v", err)
			}
		}
	}

	if !existing {
		o, err = c.CreateTelemetryControllerAgent(obj, vdomparam)

		if err != nil {
			return fmt.Errorf("Error creating TelemetryControllerAgent resource: %v", err)
		}
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("TelemetryControllerAgent")
	}

	return resourceTelemetryControllerAgentRead(d, m)
}

func resourceTelemetryControllerAgentUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectTelemetryControllerAgent(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating TelemetryControllerAgent resource while getting object: %v", err)
	}

	o, err := c.UpdateTelemetryControllerAgent(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating TelemetryControllerAgent resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("TelemetryControllerAgent")
	}

	return resourceTelemetryControllerAgentRead(d, m)
}

func resourceTelemetryControllerAgentDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteTelemetryControllerAgent(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting TelemetryControllerAgent resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceTelemetryControllerAgentRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadTelemetryControllerAgent(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading TelemetryControllerAgent resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectTelemetryControllerAgent(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading TelemetryControllerAgent resource from API: %v", err)
	}
	return nil
}

func flattenTelemetryControllerAgentAgentId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenTelemetryControllerAgentComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenTelemetryControllerAgentAlias(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenTelemetryControllerAgentAuthz(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenTelemetryControllerAgentAgentProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectTelemetryControllerAgent(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("agent_id", flattenTelemetryControllerAgentAgentId(o["agent-id"], d, "agent_id", sv)); err != nil {
		if !fortiAPIPatch(o["agent-id"]) {
			return fmt.Errorf("Error reading agent_id: %v", err)
		}
	}

	if err = d.Set("comment", flattenTelemetryControllerAgentComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("alias", flattenTelemetryControllerAgentAlias(o["alias"], d, "alias", sv)); err != nil {
		if !fortiAPIPatch(o["alias"]) {
			return fmt.Errorf("Error reading alias: %v", err)
		}
	}

	if err = d.Set("authz", flattenTelemetryControllerAgentAuthz(o["authz"], d, "authz", sv)); err != nil {
		if !fortiAPIPatch(o["authz"]) {
			return fmt.Errorf("Error reading authz: %v", err)
		}
	}

	if err = d.Set("agent_profile", flattenTelemetryControllerAgentAgentProfile(o["agent-profile"], d, "agent_profile", sv)); err != nil {
		if !fortiAPIPatch(o["agent-profile"]) {
			return fmt.Errorf("Error reading agent_profile: %v", err)
		}
	}

	return nil
}

func flattenTelemetryControllerAgentFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandTelemetryControllerAgentAgentId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerAgentComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerAgentAlias(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerAgentAuthz(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerAgentAgentProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectTelemetryControllerAgent(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("agent_id"); ok {
		t, err := expandTelemetryControllerAgentAgentId(d, v, "agent_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["agent-id"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandTelemetryControllerAgentComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("alias"); ok {
		t, err := expandTelemetryControllerAgentAlias(d, v, "alias", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alias"] = t
		}
	} else if d.HasChange("alias") {
		obj["alias"] = nil
	}

	if v, ok := d.GetOk("authz"); ok {
		t, err := expandTelemetryControllerAgentAuthz(d, v, "authz", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authz"] = t
		}
	}

	if v, ok := d.GetOk("agent_profile"); ok {
		t, err := expandTelemetryControllerAgentAgentProfile(d, v, "agent_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["agent-profile"] = t
		}
	} else if d.HasChange("agent_profile") {
		obj["agent-profile"] = nil
	}

	return &obj, nil
}
