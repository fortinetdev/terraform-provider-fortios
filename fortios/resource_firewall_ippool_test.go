// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSFirewallIppool_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallIppool_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallIppoolConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallIppoolExists("fortios_firewall_ippool.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_ippool.trname", "arp_reply", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_ippool.trname", "block_size", "128"),
					resource.TestCheckResourceAttr("fortios_firewall_ippool.trname", "endip", "1.0.0.20"),
					resource.TestCheckResourceAttr("fortios_firewall_ippool.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_ippool.trname", "num_blocks_per_user", "8"),
					resource.TestCheckResourceAttr("fortios_firewall_ippool.trname", "pba_timeout", "30"),
					resource.TestCheckResourceAttr("fortios_firewall_ippool.trname", "permit_any_host", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_ippool.trname", "source_endip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_firewall_ippool.trname", "source_startip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_firewall_ippool.trname", "startip", "1.0.0.0"),
					resource.TestCheckResourceAttr("fortios_firewall_ippool.trname", "type", "overload"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallIppoolExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallIppool: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallIppool is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallIppool(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading FirewallIppool: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallIppool: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallIppoolDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_ippool" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallIppool(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallIppool %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallIppoolConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_ippool" "trname" {
  arp_reply           = "enable"
  block_size          = 128
  endip               = "1.0.0.20"
  name                = "%[1]s"
  num_blocks_per_user = 8
  pba_timeout         = 30
  permit_any_host     = "disable"
  source_endip        = "0.0.0.0"
  source_startip      = "0.0.0.0"
  startip             = "1.0.0.0"
  type                = "overload"
}
`, name)
}
