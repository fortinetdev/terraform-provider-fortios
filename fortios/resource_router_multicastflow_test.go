// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSRouterMulticastFlow_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSRouterMulticastFlow_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSRouterMulticastFlowConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSRouterMulticastFlowExists("fortios_router_multicastflow.trname"),
					resource.TestCheckResourceAttr("fortios_router_multicastflow.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_router_multicastflow.trname", "flows.0.group_addr", "224.252.0.0"),
					resource.TestCheckResourceAttr("fortios_router_multicastflow.trname", "flows.0.source_addr", "224.112.0.0"),
				),
			},
		},
	})
}

func testAccCheckFortiOSRouterMulticastFlowExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterMulticastFlow: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterMulticastFlow is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterMulticastFlow(i)

		if err != nil {
			return fmt.Errorf("Error reading RouterMulticastFlow: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterMulticastFlow: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterMulticastFlowDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_multicastflow" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterMulticastFlow(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterMulticastFlow %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterMulticastFlowConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_multicastflow" "trname" {
  name = "%[1]s"

  flows {
    group_addr  = "224.252.0.0"
    source_addr = "224.112.0.0"
  }
}
`, name)
}
