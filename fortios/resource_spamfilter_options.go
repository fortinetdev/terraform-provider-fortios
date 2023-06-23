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

func resourceSpamfilterOptions() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpamfilterOptionsUpdate,
		Read:   resourceSpamfilterOptionsRead,
		Update: resourceSpamfilterOptionsUpdate,
		Delete: resourceSpamfilterOptionsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
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

func resourceSpamfilterOptionsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSpamfilterOptions(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SpamfilterOptions resource while getting object: %v", err)
	}

	o, err := c.UpdateSpamfilterOptions(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SpamfilterOptions resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SpamfilterOptions")
	}

	return resourceSpamfilterOptionsRead(d, m)
}

func resourceSpamfilterOptionsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSpamfilterOptions(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SpamfilterOptions resource while getting object: %v", err)
	}

	_, err = c.UpdateSpamfilterOptions(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SpamfilterOptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSpamfilterOptionsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSpamfilterOptions(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SpamfilterOptions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSpamfilterOptions(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SpamfilterOptions resource from API: %v", err)
	}
	return nil
}

func flattenSpamfilterOptionsDnsTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSpamfilterOptions(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("dns_timeout", flattenSpamfilterOptionsDnsTimeout(o["dns-timeout"], d, "dns_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["dns-timeout"]) {
			return fmt.Errorf("Error reading dns_timeout: %v", err)
		}
	}

	return nil
}

func flattenSpamfilterOptionsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSpamfilterOptionsDnsTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSpamfilterOptions(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dns_timeout"); ok {
		if setArgNil {
			obj["dns-timeout"] = nil
		} else {
			t, err := expandSpamfilterOptionsDnsTimeout(d, v, "dns_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dns-timeout"] = t
			}
		}
	}

	return &obj, nil
}
