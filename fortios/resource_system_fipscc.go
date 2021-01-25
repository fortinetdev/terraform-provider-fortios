// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FIPS-CC mode.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemFipsCc() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFipsCcUpdate,
		Read:   resourceSystemFipsCcRead,
		Update: resourceSystemFipsCcUpdate,
		Delete: resourceSystemFipsCcDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entropy_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"self_test_period": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1440),
				Optional:     true,
				Computed:     true,
			},
			"key_generation_self_test": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemFipsCcUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemFipsCc(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemFipsCc resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemFipsCc(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemFipsCc resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemFipsCc")
	}

	return resourceSystemFipsCcRead(d, m)
}

func resourceSystemFipsCcDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemFipsCc(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFipsCc resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFipsCcRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemFipsCc(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemFipsCc resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFipsCc(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFipsCc resource from API: %v", err)
	}
	return nil
}

func flattenSystemFipsCcStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFipsCcEntropyToken(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFipsCcSelfTestPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFipsCcKeyGenerationSelfTest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemFipsCc(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", flattenSystemFipsCcStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("entropy_token", flattenSystemFipsCcEntropyToken(o["entropy-token"], d, "entropy_token")); err != nil {
		if !fortiAPIPatch(o["entropy-token"]) {
			return fmt.Errorf("Error reading entropy_token: %v", err)
		}
	}

	if err = d.Set("self_test_period", flattenSystemFipsCcSelfTestPeriod(o["self-test-period"], d, "self_test_period")); err != nil {
		if !fortiAPIPatch(o["self-test-period"]) {
			return fmt.Errorf("Error reading self_test_period: %v", err)
		}
	}

	if err = d.Set("key_generation_self_test", flattenSystemFipsCcKeyGenerationSelfTest(o["key-generation-self-test"], d, "key_generation_self_test")); err != nil {
		if !fortiAPIPatch(o["key-generation-self-test"]) {
			return fmt.Errorf("Error reading key_generation_self_test: %v", err)
		}
	}

	return nil
}

func flattenSystemFipsCcFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemFipsCcStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFipsCcEntropyToken(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFipsCcSelfTestPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFipsCcKeyGenerationSelfTest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFipsCc(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemFipsCcStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("entropy_token"); ok {
		t, err := expandSystemFipsCcEntropyToken(d, v, "entropy_token")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entropy-token"] = t
		}
	}

	if v, ok := d.GetOk("self_test_period"); ok {
		t, err := expandSystemFipsCcSelfTestPeriod(d, v, "self_test_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["self-test-period"] = t
		}
	}

	if v, ok := d.GetOk("key_generation_self_test"); ok {
		t, err := expandSystemFipsCcKeyGenerationSelfTest(d, v, "key_generation_self_test")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-generation-self-test"] = t
		}
	}

	return &obj, nil
}
