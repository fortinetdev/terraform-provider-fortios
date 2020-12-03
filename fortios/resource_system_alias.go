// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure alias command.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemAlias() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAliasCreate,
		Read:   resourceSystemAliasRead,
		Update: resourceSystemAliasUpdate,
		Delete: resourceSystemAliasDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
				ForceNew:     true,
			},
			"command": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
		},
	}
}

func resourceSystemAliasCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAlias(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAlias resource while getting object: %v", err)
	}

	o, err := c.CreateSystemAlias(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemAlias resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAlias")
	}

	return resourceSystemAliasRead(d, m)
}

func resourceSystemAliasUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAlias(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAlias resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAlias(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemAlias resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAlias")
	}

	return resourceSystemAliasRead(d, m)
}

func resourceSystemAliasDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemAlias(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAlias resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAliasRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemAlias(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemAlias resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAlias(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAlias resource from API: %v", err)
	}
	return nil
}

func flattenSystemAliasName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAliasCommand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAlias(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemAliasName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("command", flattenSystemAliasCommand(o["command"], d, "command")); err != nil {
		if !fortiAPIPatch(o["command"]) {
			return fmt.Errorf("Error reading command: %v", err)
		}
	}

	return nil
}

func flattenSystemAliasFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAliasName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAliasCommand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAlias(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemAliasName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("command"); ok {
		t, err := expandSystemAliasCommand(d, v, "command")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command"] = t
		}
	}

	return &obj, nil
}
