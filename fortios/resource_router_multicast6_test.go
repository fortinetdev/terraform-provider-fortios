// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"log"
	"testing"
)

func TestAccFortiOSRouterMulticast6_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSRouterMulticast6_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSRouterMulticast6Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSRouterMulticast6Exists("fortios_router_multicast6.trname"),
					resource.TestCheckResourceAttr("fortios_router_multicast6.trname", "multicast_pmtu", "disable"),
					resource.TestCheckResourceAttr("fortios_router_multicast6.trname", "multicast_routing", "disable"),
					resource.TestCheckResourceAttr("fortios_router_multicast6.trname", "pim_sm_global.0.register_rate_limit", "0"),
				),
			},
		},
	})
}

func testAccCheckFortiOSRouterMulticast6Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterMulticast6: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterMulticast6 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterMulticast6(i)

		if err != nil {
			return fmt.Errorf("Error reading RouterMulticast6: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterMulticast6: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterMulticast6Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_multicast6" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterMulticast6(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterMulticast6 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterMulticast6Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_multicast6" "trname" {
  multicast_pmtu    = "disable"
  multicast_routing = "disable"

  pim_sm_global {
    register_rate_limit = 0
  }
}
`)
}
