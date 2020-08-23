// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSLogFortianalyzerOverrideSetting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogFortianalyzerOverrideSetting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogFortianalyzerOverrideSettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogFortianalyzerOverrideSettingExists("fortios_logfortianalyzer_overridesetting.trname"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridesetting.trname", "__change_ip", "0"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridesetting.trname", "conn_timeout", "10"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridesetting.trname", "enc_algorithm", "high"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridesetting.trname", "faz_type", "4"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridesetting.trname", "hmac_algorithm", "sha256"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridesetting.trname", "ips_archive", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridesetting.trname", "monitor_failure_retry_period", "5"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridesetting.trname", "monitor_keepalive_period", "5"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridesetting.trname", "override", "disable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridesetting.trname", "reliable", "disable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridesetting.trname", "ssl_min_proto_version", "default"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridesetting.trname", "status", "disable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridesetting.trname", "upload_interval", "daily"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridesetting.trname", "upload_option", "5-minute"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridesetting.trname", "upload_time", "00:59"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridesetting.trname", "use_management_vdom", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogFortianalyzerOverrideSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogFortianalyzerOverrideSetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogFortianalyzerOverrideSetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogFortianalyzerOverrideSetting(i)

		if err != nil {
			return fmt.Errorf("Error reading LogFortianalyzerOverrideSetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogFortianalyzerOverrideSetting: %s", n)
		}

		return nil
	}
}

func testAccCheckLogFortianalyzerOverrideSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_logfortianalyzer_overridesetting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogFortianalyzerOverrideSetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogFortianalyzerOverrideSetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogFortianalyzerOverrideSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_logfortianalyzer_overridesetting" "trname" {
  __change_ip                  = 0
  conn_timeout                 = 10
  enc_algorithm                = "high"
  faz_type                     = 4
  hmac_algorithm               = "sha256"
  ips_archive                  = "enable"
  monitor_failure_retry_period = 5
  monitor_keepalive_period     = 5
  override                     = "disable"
  reliable                     = "disable"
  ssl_min_proto_version        = "default"
  status                       = "disable"
  upload_interval              = "daily"
  upload_option                = "5-minute"
  upload_time                  = "00:59"
  use_management_vdom          = "disable"
}
`)
}
