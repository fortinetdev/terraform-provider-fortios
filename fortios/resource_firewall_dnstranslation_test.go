
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

func TestAccFortiOSFirewallDnstranslation_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSFirewallDnstranslation_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSFirewallDnstranslationConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSFirewallDnstranslationExists("fortios_firewall_dnstranslation.trname"),
                    resource.TestCheckResourceAttr("fortios_firewall_dnstranslation.trname", "dst", "2.2.2.2"),
                    resource.TestCheckResourceAttr("fortios_firewall_dnstranslation.trname", "fosid", "1"),
                    resource.TestCheckResourceAttr("fortios_firewall_dnstranslation.trname", "netmask", "255.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_firewall_dnstranslation.trname", "src", "1.1.1.1"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSFirewallDnstranslationExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallDnstranslation: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallDnstranslation is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallDnstranslation(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallDnstranslation: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallDnstranslation: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallDnstranslationDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_dnstranslation" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallDnstranslation(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallDnstranslation %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallDnstranslationConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_dnstranslation" "trname" {
  dst     = "2.2.2.2"
  fosid   = 1
  netmask = "255.0.0.0"
  src     = "1.1.1.1"
}
`)
}
