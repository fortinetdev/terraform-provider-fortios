// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure the FortiGate switch controller to send custom commands to managed FortiSwitch devices.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerCustomCommand() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerCustomCommandCreate,
		Read:   resourceSwitchControllerCustomCommandRead,
		Update: resourceSwitchControllerCustomCommandUpdate,
		Delete: resourceSwitchControllerCustomCommandDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"command_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"command": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 4095),
				Required:     true,
			},
		},
	}
}

func resourceSwitchControllerCustomCommandCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerCustomCommand(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerCustomCommand resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerCustomCommand(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerCustomCommand resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerCustomCommand")
	}

	return resourceSwitchControllerCustomCommandRead(d, m)
}

func resourceSwitchControllerCustomCommandUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerCustomCommand(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerCustomCommand resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerCustomCommand(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerCustomCommand resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerCustomCommand")
	}

	return resourceSwitchControllerCustomCommandRead(d, m)
}

func resourceSwitchControllerCustomCommandDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerCustomCommand(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerCustomCommand resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerCustomCommandRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerCustomCommand(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerCustomCommand resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerCustomCommand(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerCustomCommand resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerCustomCommandCommandName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerCustomCommandDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerCustomCommandCommand(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerCustomCommand(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("command_name", flattenSwitchControllerCustomCommandCommandName(o["command-name"], d, "command_name", sv)); err != nil {
		if !fortiAPIPatch(o["command-name"]) {
			return fmt.Errorf("Error reading command_name: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerCustomCommandDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("command", flattenSwitchControllerCustomCommandCommand(o["command"], d, "command", sv)); err != nil {
		if !fortiAPIPatch(o["command"]) {
			return fmt.Errorf("Error reading command: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerCustomCommandFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerCustomCommandCommandName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerCustomCommandDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerCustomCommandCommand(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerCustomCommand(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("command_name"); ok {

		t, err := expandSwitchControllerCustomCommandCommandName(d, v, "command_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command-name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSwitchControllerCustomCommandDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("command"); ok {

		t, err := expandSwitchControllerCustomCommandCommand(d, v, "command", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command"] = t
		}
	}

	return &obj, nil
}
