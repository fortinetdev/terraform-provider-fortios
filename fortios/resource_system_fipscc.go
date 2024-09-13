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
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
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

	obj, err := getObjectSystemFipsCc(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemFipsCc resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemFipsCc(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemFipsCc(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemFipsCc resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFipsCc(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemFipsCc resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFipsCcRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemFipsCc(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemFipsCc resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFipsCc(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemFipsCc resource from API: %v", err)
	}
	return nil
}

func flattenSystemFipsCcStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFipsCcEntropyToken(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFipsCcSelfTestPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemFipsCcKeyGenerationSelfTest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemFipsCc(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemFipsCcStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("entropy_token", flattenSystemFipsCcEntropyToken(o["entropy-token"], d, "entropy_token", sv)); err != nil {
		if !fortiAPIPatch(o["entropy-token"]) {
			return fmt.Errorf("Error reading entropy_token: %v", err)
		}
	}

	if err = d.Set("self_test_period", flattenSystemFipsCcSelfTestPeriod(o["self-test-period"], d, "self_test_period", sv)); err != nil {
		if !fortiAPIPatch(o["self-test-period"]) {
			return fmt.Errorf("Error reading self_test_period: %v", err)
		}
	}

	if err = d.Set("key_generation_self_test", flattenSystemFipsCcKeyGenerationSelfTest(o["key-generation-self-test"], d, "key_generation_self_test", sv)); err != nil {
		if !fortiAPIPatch(o["key-generation-self-test"]) {
			return fmt.Errorf("Error reading key_generation_self_test: %v", err)
		}
	}

	return nil
}

func flattenSystemFipsCcFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemFipsCcStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFipsCcEntropyToken(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFipsCcSelfTestPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFipsCcKeyGenerationSelfTest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFipsCc(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSystemFipsCcStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("entropy_token"); ok {
		if setArgNil {
			obj["entropy-token"] = nil
		} else {
			t, err := expandSystemFipsCcEntropyToken(d, v, "entropy_token", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["entropy-token"] = t
			}
		}
	}

	if v, ok := d.GetOk("self_test_period"); ok {
		if setArgNil {
			obj["self-test-period"] = nil
		} else {
			t, err := expandSystemFipsCcSelfTestPeriod(d, v, "self_test_period", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["self-test-period"] = t
			}
		}
	}

	if v, ok := d.GetOk("key_generation_self_test"); ok {
		if setArgNil {
			obj["key-generation-self-test"] = nil
		} else {
			t, err := expandSystemFipsCcKeyGenerationSelfTest(d, v, "key_generation_self_test", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["key-generation-self-test"] = t
			}
		}
	}

	return &obj, nil
}
