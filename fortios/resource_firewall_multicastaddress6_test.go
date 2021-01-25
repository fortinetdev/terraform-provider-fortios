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

func TestAccFortiOSFirewallMulticastAddress6_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallMulticastAddress6_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallMulticastAddress6Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallMulticastAddress6Exists("fortios_firewall_multicastaddress6.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_multicastaddress6.trname", "color", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_multicastaddress6.trname", "ip6", "ff02::1:ff0e:8c6c/128"),
					resource.TestCheckResourceAttr("fortios_firewall_multicastaddress6.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_multicastaddress6.trname", "visibility", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallMulticastAddress6Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallMulticastAddress6: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallMulticastAddress6 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallMulticastAddress6(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallMulticastAddress6: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallMulticastAddress6: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallMulticastAddress6Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_multicastaddress6" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallMulticastAddress6(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallMulticastAddress6 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallMulticastAddress6Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_multicastaddress6" "trname" {
  color      = 0
  ip6        = "ff02::1:ff0e:8c6c/128"
  name       = "%[1]s"
  visibility = "enable"
}
`, name)
}
