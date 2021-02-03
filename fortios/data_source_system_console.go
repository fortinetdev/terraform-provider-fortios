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

func dataSourceSystemConsole() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemConsoleRead,
		Schema: map[string]*schema.Schema{
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"baudrate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"output": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"login": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemConsoleRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "SystemConsole"

	o, err := c.ReadSystemConsole(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemConsole: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemConsole(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemConsole from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemConsoleMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemConsoleBaudrate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemConsoleOutput(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemConsoleLogin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemConsole(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("mode", dataSourceFlattenSystemConsoleMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("baudrate", dataSourceFlattenSystemConsoleBaudrate(o["baudrate"], d, "baudrate")); err != nil {
		if !fortiAPIPatch(o["baudrate"]) {
			return fmt.Errorf("Error reading baudrate: %v", err)
		}
	}

	if err = d.Set("output", dataSourceFlattenSystemConsoleOutput(o["output"], d, "output")); err != nil {
		if !fortiAPIPatch(o["output"]) {
			return fmt.Errorf("Error reading output: %v", err)
		}
	}

	if err = d.Set("login", dataSourceFlattenSystemConsoleLogin(o["login"], d, "login")); err != nil {
		if !fortiAPIPatch(o["login"]) {
			return fmt.Errorf("Error reading login: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemConsoleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
