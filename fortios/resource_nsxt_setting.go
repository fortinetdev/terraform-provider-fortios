// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure NSX-T setting.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceNsxtSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtSettingUpdate,
		Read:   resourceNsxtSettingRead,
		Update: resourceNsxtSettingUpdate,
		Delete: resourceNsxtSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"liveness": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceNsxtSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectNsxtSetting(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating NsxtSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateNsxtSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating NsxtSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("NsxtSetting")
	}

	return resourceNsxtSettingRead(d, m)
}

func resourceNsxtSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteNsxtSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting NsxtSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceNsxtSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadNsxtSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading NsxtSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectNsxtSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading NsxtSetting resource from API: %v", err)
	}
	return nil
}

func flattenNsxtSettingLiveness(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenNsxtSettingService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectNsxtSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("liveness", flattenNsxtSettingLiveness(o["liveness"], d, "liveness", sv)); err != nil {
		if !fortiAPIPatch(o["liveness"]) {
			return fmt.Errorf("Error reading liveness: %v", err)
		}
	}

	if err = d.Set("service", flattenNsxtSettingService(o["service"], d, "service", sv)); err != nil {
		if !fortiAPIPatch(o["service"]) {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	return nil
}

func flattenNsxtSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandNsxtSettingLiveness(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandNsxtSettingService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectNsxtSetting(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("liveness"); ok {

		t, err := expandNsxtSettingLiveness(d, v, "liveness", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["liveness"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok {

		t, err := expandNsxtSettingService(d, v, "service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	return &obj, nil
}
