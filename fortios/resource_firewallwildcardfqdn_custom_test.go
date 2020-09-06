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

func TestAccFortiOSFirewallWildcardFqdnCustom_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallWildcardFqdnCustom_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallWildcardFqdnCustomConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallWildcardFqdnCustomExists("fortios_firewallwildcardfqdn_custom.trname"),
					resource.TestCheckResourceAttr("fortios_firewallwildcardfqdn_custom.trname", "color", "0"),
					resource.TestCheckResourceAttr("fortios_firewallwildcardfqdn_custom.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewallwildcardfqdn_custom.trname", "visibility", "enable"),
					resource.TestCheckResourceAttr("fortios_firewallwildcardfqdn_custom.trname", "wildcard_fqdn", "*.go.google.com"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallWildcardFqdnCustomExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallWildcardFqdnCustom: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallWildcardFqdnCustom is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallWildcardFqdnCustom(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallWildcardFqdnCustom: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallWildcardFqdnCustom: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallWildcardFqdnCustomDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewallwildcardfqdn_custom" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallWildcardFqdnCustom(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallWildcardFqdnCustom %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallWildcardFqdnCustomConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewallwildcardfqdn_custom" "trname" {
  color         = 0
  name          = "%[1]s"
  visibility    = "enable"
  wildcard_fqdn = "*.go.google.com"
}
`, name)
}
