// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSLogSetting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogSetting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogSettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogSettingExists("fortios_log_setting.trname"),
					resource.TestCheckResourceAttr("fortios_log_setting.trname", "brief_traffic_format", "disable"),
					resource.TestCheckResourceAttr("fortios_log_setting.trname", "daemon_log", "disable"),
					resource.TestCheckResourceAttr("fortios_log_setting.trname", "expolicy_implicit_log", "disable"),
					resource.TestCheckResourceAttr("fortios_log_setting.trname", "faz_override", "disable"),
					resource.TestCheckResourceAttr("fortios_log_setting.trname", "fwpolicy6_implicit_log", "disable"),
					resource.TestCheckResourceAttr("fortios_log_setting.trname", "fwpolicy_implicit_log", "disable"),
					resource.TestCheckResourceAttr("fortios_log_setting.trname", "local_in_allow", "disable"),
					resource.TestCheckResourceAttr("fortios_log_setting.trname", "local_in_deny_broadcast", "disable"),
					resource.TestCheckResourceAttr("fortios_log_setting.trname", "local_in_deny_unicast", "disable"),
					resource.TestCheckResourceAttr("fortios_log_setting.trname", "local_out", "disable"),
					resource.TestCheckResourceAttr("fortios_log_setting.trname", "log_invalid_packet", "disable"),
					resource.TestCheckResourceAttr("fortios_log_setting.trname", "log_policy_comment", "disable"),
					resource.TestCheckResourceAttr("fortios_log_setting.trname", "log_policy_name", "disable"),
					resource.TestCheckResourceAttr("fortios_log_setting.trname", "log_user_in_upper", "disable"),
					resource.TestCheckResourceAttr("fortios_log_setting.trname", "neighbor_event", "disable"),
					resource.TestCheckResourceAttr("fortios_log_setting.trname", "resolve_ip", "disable"),
					resource.TestCheckResourceAttr("fortios_log_setting.trname", "resolve_port", "enable"),
					resource.TestCheckResourceAttr("fortios_log_setting.trname", "syslog_override", "disable"),
					resource.TestCheckResourceAttr("fortios_log_setting.trname", "user_anonymize", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogSetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogSetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogSetting(i)

		if err != nil {
			return fmt.Errorf("Error reading LogSetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogSetting: %s", n)
		}

		return nil
	}
}

func testAccCheckLogSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_log_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogSetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogSetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_log_setting" "trname" {
  brief_traffic_format    = "disable"
  daemon_log              = "disable"
  expolicy_implicit_log   = "disable"
  faz_override            = "disable"
  fwpolicy6_implicit_log  = "disable"
  fwpolicy_implicit_log   = "disable"
  local_in_allow          = "disable"
  local_in_deny_broadcast = "disable"
  local_in_deny_unicast   = "disable"
  local_out               = "disable"
  log_invalid_packet      = "disable"
  log_policy_comment      = "disable"
  log_policy_name         = "disable"
  log_user_in_upper       = "disable"
  neighbor_event          = "disable"
  resolve_ip              = "disable"
  resolve_port            = "enable"
  syslog_override         = "disable"
  user_anonymize          = "disable"
}
`)
}
