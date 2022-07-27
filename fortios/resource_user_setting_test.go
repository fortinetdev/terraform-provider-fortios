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

func TestAccFortiOSUserSetting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSUserSetting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSUserSettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSUserSettingExists("fortios_user_setting.trname"),
					resource.TestCheckResourceAttr("fortios_user_setting.trname", "auth_blackout_time", "0"),
					resource.TestCheckResourceAttr("fortios_user_setting.trname", "auth_cert", "Fortinet_Factory"),
					resource.TestCheckResourceAttr("fortios_user_setting.trname", "auth_http_basic", "disable"),
					resource.TestCheckResourceAttr("fortios_user_setting.trname", "auth_invalid_max", "5"),
					resource.TestCheckResourceAttr("fortios_user_setting.trname", "auth_lockout_duration", "0"),
					resource.TestCheckResourceAttr("fortios_user_setting.trname", "auth_lockout_threshold", "3"),
					resource.TestCheckResourceAttr("fortios_user_setting.trname", "auth_on_demand", "implicitly"),
					resource.TestCheckResourceAttr("fortios_user_setting.trname", "auth_portal_timeout", "3"),
					resource.TestCheckResourceAttr("fortios_user_setting.trname", "auth_secure_http", "disable"),
					resource.TestCheckResourceAttr("fortios_user_setting.trname", "auth_src_mac", "enable"),
					resource.TestCheckResourceAttr("fortios_user_setting.trname", "auth_ssl_allow_renegotiation", "disable"),
					resource.TestCheckResourceAttr("fortios_user_setting.trname", "auth_timeout", "5"),
					resource.TestCheckResourceAttr("fortios_user_setting.trname", "auth_timeout_type", "idle-timeout"),
					resource.TestCheckResourceAttr("fortios_user_setting.trname", "auth_type", "http https ftp telnet"),
					resource.TestCheckResourceAttr("fortios_user_setting.trname", "radius_ses_timeout_act", "hard-timeout"),
				),
			},
		},
	})
}

func testAccCheckFortiOSUserSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found UserSetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No UserSetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadUserSetting(i)

		if err != nil {
			return fmt.Errorf("Error reading UserSetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating UserSetting: %s", n)
		}

		return nil
	}
}

func testAccCheckUserSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_user_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadUserSetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error UserSetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSUserSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_user_setting" "trname" {
  auth_blackout_time           = 0
  auth_cert                    = "Fortinet_Factory"
  auth_http_basic              = "disable"
  auth_invalid_max             = 5
  auth_lockout_duration        = 0
  auth_lockout_threshold       = 3
  auth_on_demand               = "implicitly"
  auth_portal_timeout          = 3
  auth_secure_http             = "disable"
  auth_src_mac                 = "enable"
  auth_ssl_allow_renegotiation = "disable"
  auth_timeout                 = 5
  auth_timeout_type            = "idle-timeout"
  auth_type                    = "http https ftp telnet"
  radius_ses_timeout_act       = "hard-timeout"
}
`)
}
