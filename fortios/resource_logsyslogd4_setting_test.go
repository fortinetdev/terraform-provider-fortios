
// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios
import (
    "fmt"
	"log"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSLogSyslogd4Setting_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSLogSyslogd4Setting_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSLogSyslogd4SettingConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSLogSyslogd4SettingExists("fortios_logsyslogd4_setting.trname"),
                    resource.TestCheckResourceAttr("fortios_logsyslogd4_setting.trname", "enc_algorithm", "disable"),
                    resource.TestCheckResourceAttr("fortios_logsyslogd4_setting.trname", "facility", "local7"),
                    resource.TestCheckResourceAttr("fortios_logsyslogd4_setting.trname", "format", "default"),
                    resource.TestCheckResourceAttr("fortios_logsyslogd4_setting.trname", "mode", "udp"),
                    resource.TestCheckResourceAttr("fortios_logsyslogd4_setting.trname", "port", "514"),
                    resource.TestCheckResourceAttr("fortios_logsyslogd4_setting.trname", "ssl_min_proto_version", "default"),
                    resource.TestCheckResourceAttr("fortios_logsyslogd4_setting.trname", "status", "disable"),
                    resource.TestCheckResourceAttr("fortios_logsyslogd4_setting.trname", "syslog_type", "4"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSLogSyslogd4SettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogSyslogd4Setting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogSyslogd4Setting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogSyslogd4Setting(i)

		if err != nil {
			return fmt.Errorf("Error reading LogSyslogd4Setting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogSyslogd4Setting: %s", n)
		}

		return nil
	}
}

func testAccCheckLogSyslogd4SettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_logsyslogd4_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogSyslogd4Setting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogSyslogd4Setting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogSyslogd4SettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_logsyslogd4_setting" "trname" {
  enc_algorithm         = "disable"
  facility              = "local7"
  format                = "default"
  mode                  = "udp"
  port                  = 514
  ssl_min_proto_version = "default"
  status                = "disable"
  syslog_type           = 4
}
`)
}
