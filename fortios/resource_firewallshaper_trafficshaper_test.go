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

func TestAccFortiOSFirewallShaperTrafficShaper_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallShaperTrafficShaper_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallShaperTrafficShaperConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallShaperTrafficShaperExists("fortios_firewallshaper_trafficshaper.trname"),
					resource.TestCheckResourceAttr("fortios_firewallshaper_trafficshaper.trname", "bandwidth_unit", "kbps"),
					resource.TestCheckResourceAttr("fortios_firewallshaper_trafficshaper.trname", "diffserv", "disable"),
					resource.TestCheckResourceAttr("fortios_firewallshaper_trafficshaper.trname", "diffservcode", "000000"),
					resource.TestCheckResourceAttr("fortios_firewallshaper_trafficshaper.trname", "guaranteed_bandwidth", "0"),
					resource.TestCheckResourceAttr("fortios_firewallshaper_trafficshaper.trname", "maximum_bandwidth", "1024"),
					resource.TestCheckResourceAttr("fortios_firewallshaper_trafficshaper.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewallshaper_trafficshaper.trname", "per_policy", "disable"),
					resource.TestCheckResourceAttr("fortios_firewallshaper_trafficshaper.trname", "priority", "low"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallShaperTrafficShaperExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallShaperTrafficShaper: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallShaperTrafficShaper is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallShaperTrafficShaper(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallShaperTrafficShaper: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallShaperTrafficShaper: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallShaperTrafficShaperDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewallshaper_trafficshaper" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallShaperTrafficShaper(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallShaperTrafficShaper %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallShaperTrafficShaperConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewallshaper_trafficshaper" "trname" {
  bandwidth_unit       = "kbps"
  diffserv             = "disable"
  diffservcode         = "000000"
  guaranteed_bandwidth = 0
  maximum_bandwidth    = 1024
  name                 = "%[1]s"
  per_policy           = "disable"
  priority             = "low"
}
`, name)
}
