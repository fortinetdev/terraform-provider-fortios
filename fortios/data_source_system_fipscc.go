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

func dataSourceSystemFipsCc() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemFipsCcRead,
		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entropy_token": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"self_test_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"key_generation_self_test": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemFipsCcRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "SystemFipsCc"

	o, err := c.ReadSystemFipsCc(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemFipsCc: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemFipsCc(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemFipsCc from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemFipsCcStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFipsCcEntropyToken(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFipsCcSelfTestPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFipsCcKeyGenerationSelfTest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemFipsCc(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", dataSourceFlattenSystemFipsCcStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("entropy_token", dataSourceFlattenSystemFipsCcEntropyToken(o["entropy-token"], d, "entropy_token")); err != nil {
		if !fortiAPIPatch(o["entropy-token"]) {
			return fmt.Errorf("Error reading entropy_token: %v", err)
		}
	}

	if err = d.Set("self_test_period", dataSourceFlattenSystemFipsCcSelfTestPeriod(o["self-test-period"], d, "self_test_period")); err != nil {
		if !fortiAPIPatch(o["self-test-period"]) {
			return fmt.Errorf("Error reading self_test_period: %v", err)
		}
	}

	if err = d.Set("key_generation_self_test", dataSourceFlattenSystemFipsCcKeyGenerationSelfTest(o["key-generation-self-test"], d, "key_generation_self_test")); err != nil {
		if !fortiAPIPatch(o["key-generation-self-test"]) {
			return fmt.Errorf("Error reading key_generation_self_test: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemFipsCcFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
