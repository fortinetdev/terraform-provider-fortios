package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSLogFortiAnalyzerSetting_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// CheckDestroy: testAccCheckLogFortiAnalyzerSettingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogFortiAnalyzerSettingConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogFortiAnalyzerSettingExists("fortios_log_fortianalyzer_setting.test1"),
					resource.TestCheckResourceAttr("fortios_log_fortianalyzer_setting.test1", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_log_fortianalyzer_setting.test1", "server", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_log_fortianalyzer_setting.test1", "source_ip", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_log_fortianalyzer_setting.test1", "upload_option", "realtime"),
					resource.TestCheckResourceAttr("fortios_log_fortianalyzer_setting.test1", "reliable", "enable"),
					resource.TestCheckResourceAttr("fortios_log_fortianalyzer_setting.test1", "hmac_algorithm", "sha256"),
					resource.TestCheckResourceAttr("fortios_log_fortianalyzer_setting.test1", "enc_algorithm", "high-medium"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogFortiAnalyzerSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found Log FortiAnalyzer Setting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Log FortiAnalyzer Setting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogFortiAnalyzerSetting(i)

		if err != nil {
			return fmt.Errorf("Error reading Log FortiAnalyzer Setting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating Log FortiAnalyzer Setting: %s", n)
		}

		return nil
	}
}

func testAccCheckLogFortiAnalyzerSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_log_fortianalyzer_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogFortiAnalyzerSetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error Log FortiAnalyzer Setting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

const testAccFortiOSLogFortiAnalyzerSettingConfig = `
resource "fortios_networking_interface_port" "test1" { 
	name = "port4"
	ip = "1.1.1.1 255.255.255.0"
	status = "down"
	allowaccess = "ping https"
	mode = "static"
	type = "physical"
}

resource "fortios_log_fortianalyzer_setting" "test1" {
	status = "enable"
	server = "${replace("${fortios_networking_interface_port.test1.ip}", " 255.255.255.0", "")}"
	source_ip = "${replace("${fortios_networking_interface_port.test1.ip}", " 255.255.255.0", "")}"
	upload_option = "realtime"
	reliable = "enable"
	hmac_algorithm = "sha256"
	enc_algorithm = "high-medium"
}
`
