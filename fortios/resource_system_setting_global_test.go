package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiOSSystemSettingGlobal_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// CheckDestroy: testAccCheckSystemSettingGlobalDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemSettingGlobalConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemSettingGlobalExists("fortios_system_setting_global.test1"),
					resource.TestCheckResourceAttr("fortios_system_setting_global.test1", "admintimeout", "65"),
					resource.TestCheckResourceAttr("fortios_system_setting_global.test1", "timezone", "04"),
					resource.TestCheckResourceAttr("fortios_system_setting_global.test1", "hostname", "mytestFortiGate"),
					resource.TestCheckResourceAttr("fortios_system_setting_global.test1", "admin_sport", "443"),
					resource.TestCheckResourceAttr("fortios_system_setting_global.test1", "admin_ssh_port", "22"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemSettingGlobalExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found System Setting Global: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No System Setting Global is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemSettingGlobal(i)

		if err != nil {
			return fmt.Errorf("Error reading System Setting Global: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating System Setting Global: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemSettingGlobalDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_setting_global" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemSettingGlobal(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error System Setting Global %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

const testAccFortiOSSystemSettingGlobalConfig = `
resource "fortios_system_setting_global" "test1" { 
	admintimeout = 65
	timezone = "04"
	hostname = "mytestFortiGate"
	admin_sport = 443
	admin_ssh_port = 22
}
`
