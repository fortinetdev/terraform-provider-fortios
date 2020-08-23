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

func TestAccFortiOSLogFortianalyzer2Setting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogFortianalyzer2Setting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogFortianalyzer2SettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogFortianalyzer2SettingExists("fortios_logfortianalyzer2_setting.trname"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_setting.trname", "__change_ip", "0"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_setting.trname", "conn_timeout", "10"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_setting.trname", "enc_algorithm", "high"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_setting.trname", "faz_type", "2"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_setting.trname", "hmac_algorithm", "sha256"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_setting.trname", "ips_archive", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_setting.trname", "mgmt_name", "FGh_Log2"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_setting.trname", "monitor_failure_retry_period", "5"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_setting.trname", "monitor_keepalive_period", "5"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_setting.trname", "reliable", "disable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_setting.trname", "ssl_min_proto_version", "default"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_setting.trname", "status", "disable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_setting.trname", "upload_interval", "daily"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_setting.trname", "upload_option", "5-minute"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_setting.trname", "upload_time", "00:59"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogFortianalyzer2SettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogFortianalyzer2Setting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogFortianalyzer2Setting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogFortianalyzer2Setting(i)

		if err != nil {
			return fmt.Errorf("Error reading LogFortianalyzer2Setting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogFortianalyzer2Setting: %s", n)
		}

		return nil
	}
}

func testAccCheckLogFortianalyzer2SettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_logfortianalyzer2_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogFortianalyzer2Setting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogFortianalyzer2Setting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogFortianalyzer2SettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_logfortianalyzer2_setting" "trname" {
  __change_ip                  = 0
  conn_timeout                 = 10
  enc_algorithm                = "high"
  faz_type                     = 2
  hmac_algorithm               = "sha256"
  ips_archive                  = "enable"
  mgmt_name                    = "FGh_Log2"
  monitor_failure_retry_period = 5
  monitor_keepalive_period     = 5
  reliable                     = "disable"
  ssl_min_proto_version        = "default"
  status                       = "disable"
  upload_interval              = "daily"
  upload_option                = "5-minute"
  upload_time                  = "00:59"
}
`)
}
