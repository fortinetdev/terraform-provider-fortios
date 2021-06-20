// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSRouterStatic6_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSRouterStatic6_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSRouterStatic6Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSRouterStatic6Exists("fortios_router_static6.trname"),
					resource.TestCheckResourceAttr("fortios_router_static6.trname", "bfd", "disable"),
					resource.TestCheckResourceAttr("fortios_router_static6.trname", "blackhole", "disable"),
					resource.TestCheckResourceAttr("fortios_router_static6.trname", "device", "port3"),
					resource.TestCheckResourceAttr("fortios_router_static6.trname", "devindex", "5"),
					resource.TestCheckResourceAttr("fortios_router_static6.trname", "distance", "10"),
					resource.TestCheckResourceAttr("fortios_router_static6.trname", "dst", "2001:db8::/32"),
					resource.TestCheckResourceAttr("fortios_router_static6.trname", "gateway", "::"),
					resource.TestCheckResourceAttr("fortios_router_static6.trname", "priority", "32"),
					resource.TestCheckResourceAttr("fortios_router_static6.trname", "seq_num", "1"),
					resource.TestCheckResourceAttr("fortios_router_static6.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_router_static6.trname", "virtual_wan_link", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSRouterStatic6Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterStatic6: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterStatic6 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterStatic6(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading RouterStatic6: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterStatic6: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterStatic6Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_static6" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterStatic6(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterStatic6 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterStatic6Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_static6" "trname" {
  bfd              = "disable"
  blackhole        = "disable"
  device           = "port3"
  devindex         = 5
  distance         = 10
  dst              = "2001:db8::/32"
  gateway          = "::"
  priority         = 32
  seq_num          = 1
  status           = "enable"
  virtual_wan_link = "disable"
}
`)
}
