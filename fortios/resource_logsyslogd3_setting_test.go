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

func TestAccFortiOSLogSyslogd3Setting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogSyslogd3Setting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogSyslogd3SettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogSyslogd3SettingExists("fortios_logsyslogd3_setting.trname"),
					resource.TestCheckResourceAttr("fortios_logsyslogd3_setting.trname", "enc_algorithm", "disable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd3_setting.trname", "facility", "local7"),
					resource.TestCheckResourceAttr("fortios_logsyslogd3_setting.trname", "format", "default"),
					resource.TestCheckResourceAttr("fortios_logsyslogd3_setting.trname", "mode", "udp"),
					resource.TestCheckResourceAttr("fortios_logsyslogd3_setting.trname", "port", "514"),
					resource.TestCheckResourceAttr("fortios_logsyslogd3_setting.trname", "ssl_min_proto_version", "default"),
					resource.TestCheckResourceAttr("fortios_logsyslogd3_setting.trname", "status", "disable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd3_setting.trname", "syslog_type", "3"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogSyslogd3SettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogSyslogd3Setting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogSyslogd3Setting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogSyslogd3Setting(i)

		if err != nil {
			return fmt.Errorf("Error reading LogSyslogd3Setting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogSyslogd3Setting: %s", n)
		}

		return nil
	}
}

func testAccCheckLogSyslogd3SettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_logsyslogd3_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogSyslogd3Setting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogSyslogd3Setting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogSyslogd3SettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_logsyslogd3_setting" "trname" {
  enc_algorithm         = "disable"
  facility              = "local7"
  format                = "default"
  mode                  = "udp"
  port                  = 514
  ssl_min_proto_version = "default"
  status                = "disable"
  syslog_type           = 3
}
`)
}
