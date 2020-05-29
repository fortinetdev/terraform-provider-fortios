package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSSystemSettingDNS_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// CheckDestroy: testAccCheckSystemSettingDNSDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemSettingDNSConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemSettingDNSExists("fortios_system_setting_dns.test1"),
					resource.TestCheckResourceAttr("fortios_system_setting_dns.test1", "primary", "208.91.112.53"),
					resource.TestCheckResourceAttr("fortios_system_setting_dns.test1", "secondary", "208.91.112.22"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemSettingDNSExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found System Setting DNS: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No System Setting DNS is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemSettingDNS(i)

		if err != nil {
			return fmt.Errorf("Error reading System Setting DNS: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating System Setting DNS: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemSettingDNSDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_setting_dns" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemSettingDNS(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error System Setting DNS %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

const testAccFortiOSSystemSettingDNSConfig = `
resource "fortios_system_setting_dns" "test1" { 
	primary = "208.91.112.53"
	secondary = "208.91.112.22"
}
`
