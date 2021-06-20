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

func TestAccFortiOSFirewallVip46_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallVip46_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallVip46Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallVip46Exists("fortios_firewall_vip46.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_vip46.trname", "arp_reply", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_vip46.trname", "color", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_vip46.trname", "extip", "10.202.1.200"),
					resource.TestCheckResourceAttr("fortios_firewall_vip46.trname", "extport", "0-65535"),
					resource.TestCheckResourceAttr("fortios_firewall_vip46.trname", "fosid", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_vip46.trname", "ldb_method", "static"),
					resource.TestCheckResourceAttr("fortios_firewall_vip46.trname", "mappedip", "2001:1:1:2::200"),
					resource.TestCheckResourceAttr("fortios_firewall_vip46.trname", "mappedport", "0-65535"),
					resource.TestCheckResourceAttr("fortios_firewall_vip46.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_vip46.trname", "portforward", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_vip46.trname", "protocol", "tcp"),
					resource.TestCheckResourceAttr("fortios_firewall_vip46.trname", "type", "static-nat"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallVip46Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallVip46: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallVip46 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallVip46(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading FirewallVip46: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallVip46: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallVip46Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_vip46" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallVip46(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallVip46 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallVip46Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_vip46" "trname" {
  arp_reply   = "enable"
  color       = 0
  extip       = "10.202.1.200"
  extport     = "0-65535"
  fosid       = 0
  ldb_method  = "static"
  mappedip    = "2001:1:1:2::200"
  mappedport  = "0-65535"
  name        = "%[1]s"
  portforward = "disable"
  protocol    = "tcp"
  type        = "static-nat"
}
`, name)
}
