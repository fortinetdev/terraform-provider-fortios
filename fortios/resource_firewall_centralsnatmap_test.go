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

func TestAccFortiOSFirewallCentralSnatMap_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallCentralSnatMap_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallCentralSnatMapConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallCentralSnatMapExists("fortios_firewall_centralsnatmap.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_centralsnatmap.trname", "nat", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_centralsnatmap.trname", "nat_port", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_centralsnatmap.trname", "orig_port", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_centralsnatmap.trname", "policyid", "1"),
					resource.TestCheckResourceAttr("fortios_firewall_centralsnatmap.trname", "protocol", "33"),
					resource.TestCheckResourceAttr("fortios_firewall_centralsnatmap.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_centralsnatmap.trname", "dst_addr.0.name", "all"),
					resource.TestCheckResourceAttr("fortios_firewall_centralsnatmap.trname", "dstintf.0.name", "port3"),
					resource.TestCheckResourceAttr("fortios_firewall_centralsnatmap.trname", "orig_addr.0.name", "all"),
					resource.TestCheckResourceAttr("fortios_firewall_centralsnatmap.trname", "srcintf.0.name", "port1"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallCentralSnatMapExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallCentralSnatMap: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallCentralSnatMap is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallCentralSnatMap(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallCentralSnatMap: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallCentralSnatMap: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallCentralSnatMapDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_centralsnatmap" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallCentralSnatMap(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallCentralSnatMap %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallCentralSnatMapConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_centralsnatmap" "trname" {
  nat       = "enable"
  nat_port  = "0"
  orig_port = "0"
  policyid  = 1
  protocol  = 33
  status    = "enable"

  dst_addr {
    name = "all"
  }

  dstintf {
    name = "port3"
  }

  orig_addr {
    name = "all"
  }

  srcintf {
    name = "port1"
  }
}
`)
}
