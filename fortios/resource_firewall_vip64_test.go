
// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios
import (
    "fmt"
	"log"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSFirewallVip64_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSFirewallVip64_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSFirewallVip64Config(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSFirewallVip64Exists("fortios_firewall_vip64.trname"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip64.trname", "arp_reply", "enable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip64.trname", "color", "0"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip64.trname", "extip", "2001:db8:99:203::22"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip64.trname", "extport", "0-65535"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip64.trname", "fosid", "0"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip64.trname", "ldb_method", "static"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip64.trname", "mappedip", "1.1.1.1"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip64.trname", "mappedport", "0-65535"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip64.trname", "name", rname),
                    resource.TestCheckResourceAttr("fortios_firewall_vip64.trname", "portforward", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip64.trname", "protocol", "tcp"),
                    resource.TestCheckResourceAttr("fortios_firewall_vip64.trname", "type", "static-nat"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSFirewallVip64Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallVip64: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallVip64 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallVip64(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallVip64: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallVip64: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallVip64Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_vip64" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallVip64(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallVip64 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallVip64Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_vip64" "trname" {
  arp_reply   = "enable"
  color       = 0
  extip       = "2001:db8:99:203::22"
  extport     = "0-65535"
  fosid       = 0
  ldb_method  = "static"
  mappedip    = "1.1.1.1"
  mappedport  = "0-65535"
  name        = "%[1]s"
  portforward = "disable"
  protocol    = "tcp"
  type        = "static-nat"
}

`, name)
}
