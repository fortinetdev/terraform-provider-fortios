package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSLogSyslogSetting_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// CheckDestroy: testAccCheckLogSyslogSettingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogSyslogSettingConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogSyslogSettingExists("fortios_log_syslog_setting.test1"),
					resource.TestCheckResourceAttr("fortios_log_syslog_setting.test1", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_log_syslog_setting.test1", "server", "2.2.2.2"),
					resource.TestCheckResourceAttr("fortios_log_syslog_setting.test1", "mode", "udp"),
					resource.TestCheckResourceAttr("fortios_log_syslog_setting.test1", "port", "514"),
					resource.TestCheckResourceAttr("fortios_log_syslog_setting.test1", "facility", "local7"),
					resource.TestCheckResourceAttr("fortios_log_syslog_setting.test1", "source_ip", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_log_syslog_setting.test1", "format", "csv"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogSyslogSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found Log Syslog Setting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Log Syslog Setting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogSyslogSetting(i)

		if err != nil {
			return fmt.Errorf("Error reading Log Syslog Setting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating Log Syslog Setting: %s", n)
		}

		return nil
	}
}

func testAccCheckLogSyslogSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_log_syslog_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogSyslogSetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error Log Syslog Setting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

const testAccFortiOSLogSyslogSettingConfig = `
resource "fortios_networking_interface_port" "test1" {
	name = "port4"
	ip = "1.1.1.1 255.255.255.0"
	status = "down"
	allowaccess = "ping https"
	mode = "static"
	type = "physical"
}

resource "fortios_log_syslog_setting" "test1" {
	status = "enable"
	server = "2.2.2.2"
	mode = "udp"
	port = "514"
	facility = "local7"
	source_ip = "${replace("${fortios_networking_interface_port.test1.ip}", " 255.255.255.0", "")}"
	format = "csv"
}
`
