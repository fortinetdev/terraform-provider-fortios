package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiManagerSystemDNS(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckFortiManager(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFMGSystemDNSDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerSystemDNSConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerSystemDNSExists("fortios_fortimanager_system_dns.test1"),
					resource.TestCheckResourceAttr("fortios_fortimanager_system_dns.test1", "primary", "208.91.112.52"),
					resource.TestCheckResourceAttr("fortios_fortimanager_system_dns.test1", "secondary", "208.91.112.54"),
				),
			},
		},
	})
}

func testAccCheckFortiManagerSystemDNSExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found system dns: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No system dns is set")
		}

		c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

		o, err := c.ReadSystemDNS()

		if err != nil {
			return fmt.Errorf("Error reading system dns: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating system dns: %s", n)
		}

		return nil
	}
}

func testAccCheckFMGSystemDNSDestroy(s *terraform.State) error {
	return nil
}

func testAccFortiManagerSystemDNSConfig() string {
	return fmt.Sprintf(`
resource "fortios_fortimanager_system_dns" "test1" {
    primary = "208.91.112.52"
    secondary = "208.91.112.54"
}
`)
}
