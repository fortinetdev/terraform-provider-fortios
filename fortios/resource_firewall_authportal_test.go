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

func TestAccFortiOSFirewallAuthPortal_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallAuthPortal_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallAuthPortalConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallAuthPortalExists("fortios_firewall_authportal.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_authportal.trname", "portal_addr", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_firewall_authportal.trname", "groups.0.name", "Guest-group"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallAuthPortalExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallAuthPortal: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallAuthPortal is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallAuthPortal(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallAuthPortal: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallAuthPortal: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallAuthPortalDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_authportal" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallAuthPortal(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallAuthPortal %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallAuthPortalConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_authportal" "trname" {
  portal_addr = "1.1.1.1"

  groups {
    name = "Guest-group"
  }
}
`)
}
