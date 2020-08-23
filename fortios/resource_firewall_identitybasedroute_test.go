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

func TestAccFortiOSFirewallIdentityBasedRoute_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallIdentityBasedRoute_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallIdentityBasedRouteConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallIdentityBasedRouteExists("fortios_firewall_identitybasedroute.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_identitybasedroute.trname", "comments", "test"),
					resource.TestCheckResourceAttr("fortios_firewall_identitybasedroute.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallIdentityBasedRouteExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallIdentityBasedRoute: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallIdentityBasedRoute is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallIdentityBasedRoute(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallIdentityBasedRoute: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallIdentityBasedRoute: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallIdentityBasedRouteDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_identitybasedroute" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallIdentityBasedRoute(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallIdentityBasedRoute %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallIdentityBasedRouteConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_identitybasedroute" "trname" {
  comments = "test"
  name     = "%[1]s"
}

`, name)
}
