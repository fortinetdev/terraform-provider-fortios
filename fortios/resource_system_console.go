// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure console.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemConsole() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemConsoleUpdate,
		Read:   resourceSystemConsoleRead,
		Update: resourceSystemConsoleUpdate,
		Delete: resourceSystemConsoleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"baudrate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"output": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"login": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemConsoleUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemConsole(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemConsole resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemConsole(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemConsole resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemConsole")
	}

	return resourceSystemConsoleRead(d, m)
}

func resourceSystemConsoleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemConsole(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemConsole resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemConsoleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemConsole(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemConsole resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemConsole(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemConsole resource from API: %v", err)
	}
	return nil
}

func flattenSystemConsoleMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemConsoleBaudrate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemConsoleOutput(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemConsoleLogin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemConsole(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("mode", flattenSystemConsoleMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("baudrate", flattenSystemConsoleBaudrate(o["baudrate"], d, "baudrate")); err != nil {
		if !fortiAPIPatch(o["baudrate"]) {
			return fmt.Errorf("Error reading baudrate: %v", err)
		}
	}

	if err = d.Set("output", flattenSystemConsoleOutput(o["output"], d, "output")); err != nil {
		if !fortiAPIPatch(o["output"]) {
			return fmt.Errorf("Error reading output: %v", err)
		}
	}

	if err = d.Set("login", flattenSystemConsoleLogin(o["login"], d, "login")); err != nil {
		if !fortiAPIPatch(o["login"]) {
			return fmt.Errorf("Error reading login: %v", err)
		}
	}

	return nil
}

func flattenSystemConsoleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemConsoleMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemConsoleBaudrate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemConsoleOutput(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemConsoleLogin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemConsole(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandSystemConsoleMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("baudrate"); ok {
		t, err := expandSystemConsoleBaudrate(d, v, "baudrate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["baudrate"] = t
		}
	}

	if v, ok := d.GetOk("output"); ok {
		t, err := expandSystemConsoleOutput(d, v, "output")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["output"] = t
		}
	}

	if v, ok := d.GetOk("login"); ok {
		t, err := expandSystemConsoleLogin(d, v, "login")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login"] = t
		}
	}

	return &obj, nil
}
