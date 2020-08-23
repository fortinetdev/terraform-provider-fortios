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

func TestAccFortiOSFirewallAddress_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallAddress_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallAddressConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallAddressExists("fortios_firewall_address.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "allow_routing", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "cache_ttl", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "color", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "end_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "fqdn", "www.ms.com"),
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "start_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "subnet", "0.0.0.0 0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "type", "fqdn"),
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "visibility", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "wildcard", "0.0.0.0 0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "wildcard_fqdn", "www.ms.com"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallAddressExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallAddress: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallAddress is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallAddress(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallAddress: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallAddress: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallAddressDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_address" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallAddress(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallAddress %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallAddressConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_address" "trname" {
  allow_routing = "disable"
  cache_ttl     = 0
  color         = 0
  end_ip        = "0.0.0.0"
  fqdn          = "www.ms.com"
  name          = "%[1]s"
  start_ip      = "0.0.0.0"
  subnet        = "0.0.0.0 0.0.0.0"
  type          = "fqdn"
  visibility    = "enable"
  wildcard      = "0.0.0.0 0.0.0.0"
  wildcard_fqdn = "www.ms.com"
}
`, name)
}
