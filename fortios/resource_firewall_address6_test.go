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

func TestAccFortiOSFirewallAddress6_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallAddress6_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallAddress6Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallAddress6Exists("fortios_firewall_address6.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_address6.trname", "cache_ttl", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_address6.trname", "color", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_address6.trname", "end_ip", "::"),
					resource.TestCheckResourceAttr("fortios_firewall_address6.trname", "host", "fdff:ffff::"),
					resource.TestCheckResourceAttr("fortios_firewall_address6.trname", "host_type", "any"),
					resource.TestCheckResourceAttr("fortios_firewall_address6.trname", "ip6", "fdff:ffff::/120"),
					resource.TestCheckResourceAttr("fortios_firewall_address6.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_address6.trname", "start_ip", "fdff:ffff::"),
					resource.TestCheckResourceAttr("fortios_firewall_address6.trname", "type", "ipprefix"),
					resource.TestCheckResourceAttr("fortios_firewall_address6.trname", "visibility", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallAddress6Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallAddress6: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallAddress6 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallAddress6(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallAddress6: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallAddress6: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallAddress6Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_address6" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallAddress6(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallAddress6 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallAddress6Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_address6" "trname" {
  cache_ttl  = 0
  color      = 0
  end_ip     = "::"
  host       = "fdff:ffff::"
  host_type  = "any"
  ip6        = "fdff:ffff::/120"
  name       = "%[1]s"
  start_ip   = "fdff:ffff::"
  type       = "ipprefix"
  visibility = "enable"
}
`, name)
}
