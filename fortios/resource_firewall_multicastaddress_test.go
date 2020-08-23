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

func TestAccFortiOSFirewallMulticastAddress_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallMulticastAddress_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallMulticastAddressConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallMulticastAddressExists("fortios_firewall_multicastaddress.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_multicastaddress.trname", "color", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_multicastaddress.trname", "end_ip", "224.0.0.22"),
					resource.TestCheckResourceAttr("fortios_firewall_multicastaddress.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_multicastaddress.trname", "start_ip", "224.0.0.11"),
					resource.TestCheckResourceAttr("fortios_firewall_multicastaddress.trname", "subnet", "224.0.0.11 224.0.0.22"),
					resource.TestCheckResourceAttr("fortios_firewall_multicastaddress.trname", "type", "multicastrange"),
					resource.TestCheckResourceAttr("fortios_firewall_multicastaddress.trname", "visibility", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallMulticastAddressExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallMulticastAddress: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallMulticastAddress is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallMulticastAddress(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallMulticastAddress: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallMulticastAddress: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallMulticastAddressDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_multicastaddress" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallMulticastAddress(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallMulticastAddress %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallMulticastAddressConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_multicastaddress" "trname" {
  color      = 0
  end_ip     = "224.0.0.22"
  name       = "%[1]s"
  start_ip   = "224.0.0.11"
  subnet     = "224.0.0.11 224.0.0.22"
  type       = "multicastrange"
  visibility = "enable"
}

`, name)
}
