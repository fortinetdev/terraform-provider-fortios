
// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios
import (
    "fmt"
	"log"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSRouterStatic_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSRouterStatic_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSRouterStaticConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSRouterStaticExists("fortios_router_static.trname"),
                    resource.TestCheckResourceAttr("fortios_router_static.trname", "bfd", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_static.trname", "blackhole", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_static.trname", "device", "port4"),
                    resource.TestCheckResourceAttr("fortios_router_static.trname", "distance", "10"),
                    resource.TestCheckResourceAttr("fortios_router_static.trname", "dst", "1.0.0.0 255.240.0.0"),
                    resource.TestCheckResourceAttr("fortios_router_static.trname", "dynamic_gateway", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_static.trname", "gateway", "0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_router_static.trname", "internet_service", "0"),
                    resource.TestCheckResourceAttr("fortios_router_static.trname", "link_monitor_exempt", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_static.trname", "priority", "22"),
                    resource.TestCheckResourceAttr("fortios_router_static.trname", "seq_num", "1"),
                    resource.TestCheckResourceAttr("fortios_router_static.trname", "src", "0.0.0.0 0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_router_static.trname", "status", "enable"),
                    resource.TestCheckResourceAttr("fortios_router_static.trname", "virtual_wan_link", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_static.trname", "vrf", "0"),
                    resource.TestCheckResourceAttr("fortios_router_static.trname", "weight", "2"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSRouterStaticExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterStatic: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterStatic is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterStatic(i)

		if err != nil {
			return fmt.Errorf("Error reading RouterStatic: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterStatic: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterStaticDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_static" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterStatic(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterStatic %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterStaticConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_static" "trname" {
  bfd                 = "disable"
  blackhole           = "disable"
  device              = "port4"
  distance            = 10
  dst                 = "1.0.0.0 255.240.0.0"
  dynamic_gateway     = "disable"
  gateway             = "0.0.0.0"
  internet_service    = 0
  link_monitor_exempt = "disable"
  priority            = 22
  seq_num             = 1
  src                 = "0.0.0.0 0.0.0.0"
  status              = "enable"
  virtual_wan_link    = "disable"
  vrf                 = 0
  weight              = 2
}
`)
}
