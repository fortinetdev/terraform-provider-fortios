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

func TestAccFortiOSLogSyslogdSetting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogSyslogdSetting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogSyslogdSettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogSyslogdSettingExists("fortios_logsyslogd_setting.trname"),
					resource.TestCheckResourceAttr("fortios_logsyslogd_setting.trname", "enc_algorithm", "disable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd_setting.trname", "facility", "local7"),
					resource.TestCheckResourceAttr("fortios_logsyslogd_setting.trname", "format", "default"),
					resource.TestCheckResourceAttr("fortios_logsyslogd_setting.trname", "mode", "udp"),
					resource.TestCheckResourceAttr("fortios_logsyslogd_setting.trname", "port", "514"),
					resource.TestCheckResourceAttr("fortios_logsyslogd_setting.trname", "ssl_min_proto_version", "default"),
					resource.TestCheckResourceAttr("fortios_logsyslogd_setting.trname", "status", "disable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd_setting.trname", "syslog_type", "1"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogSyslogdSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogSyslogdSetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogSyslogdSetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogSyslogdSetting(i)

		if err != nil {
			return fmt.Errorf("Error reading LogSyslogdSetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogSyslogdSetting: %s", n)
		}

		return nil
	}
}

func testAccCheckLogSyslogdSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_logsyslogd_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogSyslogdSetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogSyslogdSetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogSyslogdSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_logsyslogd_setting" "trname" {
  enc_algorithm         = "disable"
  facility              = "local7"
  format                = "default"
  mode                  = "udp"
  port                  = 514
  ssl_min_proto_version = "default"
  status                = "disable"
  syslog_type           = 1
}
`)
}
