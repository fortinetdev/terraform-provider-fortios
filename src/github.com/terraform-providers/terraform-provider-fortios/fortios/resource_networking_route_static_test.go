package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiOSNetworkingRouteStatic_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkingRouteStaticDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSNetworkingRouteStaticConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSNetworkingRouteStaticExists("fortios_networking_route_static.test1"),
					resource.TestCheckResourceAttr("fortios_networking_route_static.test1", "dst", "110.2.2.122 255.255.255.255"),
					resource.TestCheckResourceAttr("fortios_networking_route_static.test1", "gateway", "2.2.2.2"),
					resource.TestCheckResourceAttr("fortios_networking_route_static.test1", "blackhole", "disable"),
					resource.TestCheckResourceAttr("fortios_networking_route_static.test1", "distance", "22"),
					resource.TestCheckResourceAttr("fortios_networking_route_static.test1", "weight", "3"),
					resource.TestCheckResourceAttr("fortios_networking_route_static.test1", "priority", "3"),
					resource.TestCheckResourceAttr("fortios_networking_route_static.test1", "device", "port2"),
					resource.TestCheckResourceAttr("fortios_networking_route_static.test1", "comment", "Terraform test")),
			},
		},
	})
}

func testAccCheckFortiOSNetworkingRouteStaticExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Route Static is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadNetworkingRouteStatic(i)

		if err != nil {
			return fmt.Errorf("Error reading Networking Route Static: %s", err)
		}

		if o == nil {
			return fmt.Errorf("route static %s was not created", n)
		}

		return nil
	}
}

func testAccCheckNetworkingRouteStaticDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_networking_route_static" {
			continue
		}

		i := rs.Primary.ID
		_, err := c.ReadNetworkingRouteStatic(i)

		if err == nil {
			return fmt.Errorf("route static %s still exists", rs.Primary.ID)
		}

		return nil
	}

	return nil
}

const testAccFortiOSNetworkingRouteStaticConfig = `
resource "fortios_networking_route_static" "test1" { 
	dst = "110.2.2.122 255.255.255.255"
	gateway = "2.2.2.2"
	blackhole = "disable"
	distance = "22"
	weight = "3"
	priority = "3"
	device = "port2"
	comment = "Terraform test"
}
`
