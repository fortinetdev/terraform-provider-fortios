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
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "associated_interface", "port2"),
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "color", "3"),
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "end_ip", "255.255.255.0"),
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "start_ip", "22.1.1.0"),
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "subnet", "22.1.1.0 255.255.255.0"),
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "type", "ipmask"),
					resource.TestCheckResourceAttr("fortios_firewall_address.trname", "visibility", "enable"),
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
  allow_routing        = "disable"
  associated_interface = "port2"
  color                = 3
  end_ip               = "255.255.255.0"
  name                 = "%[1]s"
  start_ip             = "22.1.1.0"
  subnet               = "22.1.1.0 255.255.255.0"
  type                 = "ipmask"
  visibility           = "enable"
}
`, name)
}
