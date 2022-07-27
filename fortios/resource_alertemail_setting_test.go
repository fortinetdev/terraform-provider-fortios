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

func TestAccFortiOSAlertemailSetting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSAlertemailSetting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSAlertemailSettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSAlertemailSettingExists("fortios_alertemail_setting.trname"),
					resource.TestCheckResourceAttr("fortios_alertemail_setting.trname", "admin_login_logs", "disable"),
					resource.TestCheckResourceAttr("fortios_alertemail_setting.trname", "alert_interval", "2"),
					resource.TestCheckResourceAttr("fortios_alertemail_setting.trname", "amc_interface_bypass_mode", "disable"),
					resource.TestCheckResourceAttr("fortios_alertemail_setting.trname", "antivirus_logs", "disable"),
					resource.TestCheckResourceAttr("fortios_alertemail_setting.trname", "configuration_changes_logs", "disable"),
					resource.TestCheckResourceAttr("fortios_alertemail_setting.trname", "critical_interval", "3"),
					resource.TestCheckResourceAttr("fortios_alertemail_setting.trname", "debug_interval", "60"),
					resource.TestCheckResourceAttr("fortios_alertemail_setting.trname", "email_interval", "5"),
					resource.TestCheckResourceAttr("fortios_alertemail_setting.trname", "emergency_interval", "1"),
					resource.TestCheckResourceAttr("fortios_alertemail_setting.trname", "error_interval", "5"),
					resource.TestCheckResourceAttr("fortios_alertemail_setting.trname", "fds_license_expiring_days", "15"),
					resource.TestCheckResourceAttr("fortios_alertemail_setting.trname", "information_interval", "30"),
				),
			},
		},
	})
}

func testAccCheckFortiOSAlertemailSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found AlertemailSetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No AlertemailSetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadAlertemailSetting(i)

		if err != nil {
			return fmt.Errorf("Error reading AlertemailSetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating AlertemailSetting: %s", n)
		}

		return nil
	}
}

func testAccCheckAlertemailSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_alertemail_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadAlertemailSetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error AlertemailSetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSAlertemailSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_alertemail_setting" "trname" {
  admin_login_logs           = "disable"
  alert_interval             = 2
  amc_interface_bypass_mode  = "disable"
  antivirus_logs             = "disable"
  configuration_changes_logs = "disable"
  critical_interval          = 3
  debug_interval             = 60
  email_interval             = 5
  emergency_interval         = 1
  error_interval             = 5
  fds_license_expiring_days  = 15
  information_interval       = 30
}
`)
}
