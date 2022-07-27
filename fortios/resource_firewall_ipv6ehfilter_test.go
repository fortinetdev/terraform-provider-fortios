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

func TestAccFortiOSFirewallIpv6EhFilter_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallIpv6EhFilter_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallIpv6EhFilterConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallIpv6EhFilterExists("fortios_firewall_ipv6ehfilter.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_ipv6ehfilter.trname", "auth", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_ipv6ehfilter.trname", "dest_opt", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_ipv6ehfilter.trname", "fragment", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_ipv6ehfilter.trname", "hop_opt", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_ipv6ehfilter.trname", "no_next", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_ipv6ehfilter.trname", "routing", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallIpv6EhFilterExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallIpv6EhFilter: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallIpv6EhFilter is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallIpv6EhFilter(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallIpv6EhFilter: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallIpv6EhFilter: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallIpv6EhFilterDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_ipv6ehfilter" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallIpv6EhFilter(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallIpv6EhFilter %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallIpv6EhFilterConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_ipv6ehfilter" "trname" {
  auth     = "disable"
  dest_opt = "disable"
  fragment = "disable"
  hop_opt  = "disable"
  no_next  = "disable"
  routing  = "enable"
}
`)
}
