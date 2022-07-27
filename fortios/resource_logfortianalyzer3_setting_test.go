// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"log"
	"testing"
)

func TestAccFortiOSLogFortianalyzer3Setting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogFortianalyzer3Setting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogFortianalyzer3SettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogFortianalyzer3SettingExists("fortios_logfortianalyzer3_setting.trname"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer3_setting.trname", "__change_ip", "0"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer3_setting.trname", "conn_timeout", "10"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer3_setting.trname", "enc_algorithm", "high"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer3_setting.trname", "faz_type", "3"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer3_setting.trname", "hmac_algorithm", "sha256"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer3_setting.trname", "ips_archive", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer3_setting.trname", "mgmt_name", "FGh_Log3"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer3_setting.trname", "monitor_failure_retry_period", "5"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer3_setting.trname", "monitor_keepalive_period", "5"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer3_setting.trname", "reliable", "disable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer3_setting.trname", "ssl_min_proto_version", "default"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer3_setting.trname", "status", "disable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer3_setting.trname", "upload_interval", "daily"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer3_setting.trname", "upload_option", "5-minute"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer3_setting.trname", "upload_time", "00:59"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogFortianalyzer3SettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogFortianalyzer3Setting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogFortianalyzer3Setting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogFortianalyzer3Setting(i)

		if err != nil {
			return fmt.Errorf("Error reading LogFortianalyzer3Setting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogFortianalyzer3Setting: %s", n)
		}

		return nil
	}
}

func testAccCheckLogFortianalyzer3SettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_logfortianalyzer3_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogFortianalyzer3Setting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogFortianalyzer3Setting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogFortianalyzer3SettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_logfortianalyzer3_setting" "trname" {
  __change_ip                  = 0
  conn_timeout                 = 10
  enc_algorithm                = "high"
  faz_type                     = 3
  hmac_algorithm               = "sha256"
  ips_archive                  = "enable"
  mgmt_name                    = "FGh_Log3"
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
