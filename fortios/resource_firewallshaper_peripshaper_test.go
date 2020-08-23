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

func TestAccFortiOSFirewallShaperPerIpShaper_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallShaperPerIpShaper_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallShaperPerIpShaperConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallShaperPerIpShaperExists("fortios_firewallshaper_peripshaper.trname"),
					resource.TestCheckResourceAttr("fortios_firewallshaper_peripshaper.trname", "bandwidth_unit", "kbps"),
					resource.TestCheckResourceAttr("fortios_firewallshaper_peripshaper.trname", "diffserv_forward", "disable"),
					resource.TestCheckResourceAttr("fortios_firewallshaper_peripshaper.trname", "diffserv_reverse", "disable"),
					resource.TestCheckResourceAttr("fortios_firewallshaper_peripshaper.trname", "diffservcode_forward", "000000"),
					resource.TestCheckResourceAttr("fortios_firewallshaper_peripshaper.trname", "diffservcode_rev", "000000"),
					resource.TestCheckResourceAttr("fortios_firewallshaper_peripshaper.trname", "max_bandwidth", "1024"),
					resource.TestCheckResourceAttr("fortios_firewallshaper_peripshaper.trname", "max_concurrent_session", "33"),
					resource.TestCheckResourceAttr("fortios_firewallshaper_peripshaper.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallShaperPerIpShaperExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallShaperPerIpShaper: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallShaperPerIpShaper is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallShaperPerIpShaper(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallShaperPerIpShaper: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallShaperPerIpShaper: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallShaperPerIpShaperDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewallshaper_peripshaper" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallShaperPerIpShaper(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallShaperPerIpShaper %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallShaperPerIpShaperConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewallshaper_peripshaper" "trname" {
  bandwidth_unit         = "kbps"
  diffserv_forward       = "disable"
  diffserv_reverse       = "disable"
  diffservcode_forward   = "000000"
  diffservcode_rev       = "000000"
  max_bandwidth          = 1024
  max_concurrent_session = 33
  name                   = "%[1]s"
}


`, name)
}
