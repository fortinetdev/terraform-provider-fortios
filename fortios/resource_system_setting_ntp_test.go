package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSSystemSettingNTP_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// CheckDestroy: testAccCheckSystemSettingNTPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemSettingNTPConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemSettingNTPExists("fortios_system_setting_ntp.test1"),
					resource.TestCheckResourceAttr("fortios_system_setting_ntp.test1", "type", "fortiguard"),
					resource.TestCheckResourceAttr("fortios_system_setting_ntp.test1", "ntpsync", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemSettingNTPExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found System Setting NTP: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No System Setting NTP is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemSettingNTP(i)

		if err != nil {
			return fmt.Errorf("Error reading System Setting NTP: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating System Setting NTP: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemSettingNTPDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_setting_ntp" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemSettingNTP(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error System Setting NTP %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

const testAccFortiOSSystemSettingNTPConfig = `
resource "fortios_system_setting_ntp" "test1" {
	type = "fortiguard"
	ntpserver = ["1.1.1.1", "2.2.2.2"]
	ntpsync = "disable"
}
`
