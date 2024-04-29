// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure AntiSpam options.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEmailfilterOptions() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailfilterOptionsUpdate,
		Read:   resourceEmailfilterOptionsRead,
		Update: resourceEmailfilterOptionsUpdate,
		Delete: resourceEmailfilterOptionsDelete,

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
			"dns_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceEmailfilterOptionsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectEmailfilterOptions(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterOptions resource while getting object: %v", err)
	}

	o, err := c.UpdateEmailfilterOptions(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating EmailfilterOptions resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("EmailfilterOptions")
	}

	return resourceEmailfilterOptionsRead(d, m)
}

func resourceEmailfilterOptionsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectEmailfilterOptions(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating EmailfilterOptions resource while getting object: %v", err)
	}

	_, err = c.UpdateEmailfilterOptions(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing EmailfilterOptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterOptionsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadEmailfilterOptions(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterOptions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEmailfilterOptions(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading EmailfilterOptions resource from API: %v", err)
	}
	return nil
}

func flattenEmailfilterOptionsDnsTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectEmailfilterOptions(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("dns_timeout", flattenEmailfilterOptionsDnsTimeout(o["dns-timeout"], d, "dns_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["dns-timeout"]) {
			return fmt.Errorf("Error reading dns_timeout: %v", err)
		}
	}

	return nil
}

func flattenEmailfilterOptionsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandEmailfilterOptionsDnsTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectEmailfilterOptions(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dns_timeout"); ok {
		if setArgNil {
			obj["dns-timeout"] = nil
		} else {
			t, err := expandEmailfilterOptionsDnsTimeout(d, v, "dns_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dns-timeout"] = t
			}
		}
	}

	return &obj, nil
}
