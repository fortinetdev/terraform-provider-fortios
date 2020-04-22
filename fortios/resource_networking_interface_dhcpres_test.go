package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiOSNetworkingInterfaceDHCPIpRes(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSNetworkingInterfaceDHCPRes,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSNetworkingInterfaceDHCPResExists("fortios_networking_interface_dhcp_ipres.test1"),
					resource.TestCheckResourceAttr("fortios_networking_interface_dhcp_ipres.test1", "interface", "TestIntf"),
					resource.TestCheckResourceAttr("fortios_networking_interface_dhcp_ipres.test1", "ip", "10.10.10.12"),
					resource.TestCheckResourceAttr("fortios_networking_interface_dhcp_ipres.test1", "mac", "82:05:27:89:50:01"),
					resource.TestCheckResourceAttr("fortios_networking_interface_dhcp_ipres.test1", "description", "Terraform Test"),
				),
			},
		},
	})
}

// TODO Network interface & associated DHCP config needs to be created for test case to work.
const testAccFortiOSNetworkingInterfaceDHCPRes = `
resource "fortios_networking_interface_dhcp_ipres" "test1" {
	interface = "TestIntf"
	ip = "10.10.10.12"
	mac = "82:05:27:89:50:01"
	description = "Terraform Test"
    vdom = "root"
}
`

func testAccCheckFortiOSNetworkingInterfaceDHCPResExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found DHCP Reservation resource named: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No DHCP Reservation is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client
		i := rs.Primary.ID
		attr := rs.Primary.Attributes

		skey, err := c.NetworkingInterfaceDHCPSrvByName(attr["interface"])
		if err != nil {
			return fmt.Errorf("error retrieving DHCP Server from interface: %s", err)
		}

		o, err := c.ReadNetworkingInterfaceDHCPIpRes(skey, i)

		if err != nil {
			return fmt.Errorf("Error reading DHCP IP Reservation: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating DHCP IP Reservation: %s", n)
		}

		return nil
	}
}
