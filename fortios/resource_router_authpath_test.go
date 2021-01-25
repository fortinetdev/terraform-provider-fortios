// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSRouterAuthPath_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSRouterAuthPath_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSRouterAuthPathConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSRouterAuthPathExists("fortios_router_authpath.trname"),
					resource.TestCheckResourceAttr("fortios_router_authpath.trname", "device", "port3"),
					resource.TestCheckResourceAttr("fortios_router_authpath.trname", "gateway", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_router_authpath.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSRouterAuthPathExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterAuthPath: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterAuthPath is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterAuthPath(i)

		if err != nil {
			return fmt.Errorf("Error reading RouterAuthPath: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterAuthPath: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterAuthPathDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_authpath" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterAuthPath(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterAuthPath %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterAuthPathConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_authpath" "trname" {
  device  = "port3"
  gateway = "1.1.1.1"
  name    = "%[1]s"
}
`, name)
}
