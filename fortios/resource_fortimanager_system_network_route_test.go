package fortios

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiManagerSystemNetworkRoute(t *testing.T) {
	route_id := acctest.RandIntRange(400, 500)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckFortiManager(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFMGSystemNetworkRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerSystemNetworkRouteConfig(route_id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerSystemNetworkRouteExists("fortios_fmg_system_network_route.test1"),
					resource.TestCheckResourceAttr("fortios_fmg_system_network_route.test1", "route_id", strconv.Itoa(route_id)),
					resource.TestCheckResourceAttr("fortios_fmg_system_network_route.test1", "destination", "1.1.1.0 255.255.255.0"),
					resource.TestCheckResourceAttr("fortios_fmg_system_network_route.test1", "gateway", "192.168.1.2"),
					resource.TestCheckResourceAttr("fortios_fmg_system_network_route.test1", "device", "port3"),
				),
			},
		},
	})
}

func testAccCheckFortiManagerSystemNetworkRouteExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found system network route: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No system network route is set")
		}

		c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

		i := rs.Primary.ID
		o, err := c.ReadSystemNetworkRoute(i)

		if err != nil {
			return fmt.Errorf("Error reading system network route: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating system network route: %s", n)
		}

		return nil
	}
}

func testAccCheckFMGSystemNetworkRouteDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).ClientFortimanager

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_fmg_system_network_route" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemNetworkRoute(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error system network route %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiManagerSystemNetworkRouteConfig(route_id int) string {
	return fmt.Sprintf(`
resource "fortios_fmg_system_network_route" "test1" {
    route_id = %d
    destination = "1.1.1.0 255.255.255.0"
    gateway = "192.168.1.2"
    device = "port3"
}
`, route_id)
}
