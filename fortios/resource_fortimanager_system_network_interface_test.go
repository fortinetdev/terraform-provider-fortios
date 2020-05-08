package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiManagerSystemNetworkInterface(t *testing.T) {
	name := "port3"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFortiManager(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerSystemNetworkInterfaceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerSystemNetworkInterfaceExists("fortios_fmg_system_network_interface.test1"),
					resource.TestCheckResourceAttr("fortios_fmg_system_network_interface.test1", "name", name),
					resource.TestCheckResourceAttr("fortios_fmg_system_network_interface.test1", "ip", "1.1.1.3 255.255.255.0"),
					resource.TestCheckResourceAttr("fortios_fmg_system_network_interface.test1", "status", "up"),
				),
			},
		},
	})
}

func testAccCheckFortiManagerSystemNetworkInterfaceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found system network interface: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No system network interface is set")
		}

		c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

		i := rs.Primary.ID
		o, err := c.ReadSystemNetworkInterface(i)

		if err != nil {
			return fmt.Errorf("Error reading system network interface: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating system network interface: %s", n)
		}

		return nil
	}
}

func testAccFortiManagerSystemNetworkInterfaceConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_fmg_system_network_interface" "test1" {
    name = "%s"
    ip = "1.1.1.3 255.255.255.0"
    status = "up"
    allow_access = ["ping"]
    service_access = ["webfilter"]
}
`, name)
}
