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

func TestAccFortiOSFirewallProxyAddress_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallProxyAddress_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallProxyAddressConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallProxyAddressExists("fortios_firewall_proxyaddress.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_proxyaddress.trname", "case_sensitivity", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_proxyaddress.trname", "color", "2"),
					resource.TestCheckResourceAttr("fortios_firewall_proxyaddress.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_proxyaddress.trname", "referrer", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_proxyaddress.trname", "type", "host-regex"),
					resource.TestCheckResourceAttr("fortios_firewall_proxyaddress.trname", "visibility", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallProxyAddressExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallProxyAddress: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallProxyAddress is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallProxyAddress(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallProxyAddress: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallProxyAddress: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallProxyAddressDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_proxyaddress" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallProxyAddress(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallProxyAddress %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallProxyAddressConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_proxyaddress" "trname" {
  case_sensitivity = "disable"
  color            = 2
  name             = "%[1]s"
  referrer         = "enable"
  type             = "host-regex"
  visibility       = "enable"
}
`, name)
}
