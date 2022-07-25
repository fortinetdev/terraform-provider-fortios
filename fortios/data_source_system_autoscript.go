// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure auto script.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemAutoScript() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemAutoScriptRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"repeat": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"start": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"script": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"output_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemAutoScriptRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemAutoScript: type error")
	}

	o, err := c.ReadSystemAutoScript(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemAutoScript: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemAutoScript(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemAutoScript from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemAutoScriptName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoScriptInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoScriptRepeat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoScriptStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoScriptScript(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoScriptOutputSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoScriptTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemAutoScript(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemAutoScriptName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("interval", dataSourceFlattenSystemAutoScriptInterval(o["interval"], d, "interval")); err != nil {
		if !fortiAPIPatch(o["interval"]) {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("repeat", dataSourceFlattenSystemAutoScriptRepeat(o["repeat"], d, "repeat")); err != nil {
		if !fortiAPIPatch(o["repeat"]) {
			return fmt.Errorf("Error reading repeat: %v", err)
		}
	}

	if err = d.Set("start", dataSourceFlattenSystemAutoScriptStart(o["start"], d, "start")); err != nil {
		if !fortiAPIPatch(o["start"]) {
			return fmt.Errorf("Error reading start: %v", err)
		}
	}

	if err = d.Set("script", dataSourceFlattenSystemAutoScriptScript(o["script"], d, "script")); err != nil {
		if !fortiAPIPatch(o["script"]) {
			return fmt.Errorf("Error reading script: %v", err)
		}
	}

	if err = d.Set("output_size", dataSourceFlattenSystemAutoScriptOutputSize(o["output-size"], d, "output_size")); err != nil {
		if !fortiAPIPatch(o["output-size"]) {
			return fmt.Errorf("Error reading output_size: %v", err)
		}
	}

	if err = d.Set("timeout", dataSourceFlattenSystemAutoScriptTimeout(o["timeout"], d, "timeout")); err != nil {
		if !fortiAPIPatch(o["timeout"]) {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemAutoScriptFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
