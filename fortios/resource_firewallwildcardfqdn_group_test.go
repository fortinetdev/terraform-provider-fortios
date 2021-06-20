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

func TestAccFortiOSFirewallWildcardFqdnGroup_basic(t *testing.T) {
	rname := acctest.RandString(8)
	var0 := "var0" + rname
	log.Printf(var0)
	log.Printf("TestAccFortiOSFirewallWildcardFqdnGroup_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallWildcardFqdnGroupConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallWildcardFqdnGroupExists("fortios_firewallwildcardfqdn_group.trname"),
					resource.TestCheckResourceAttr("fortios_firewallwildcardfqdn_group.trname", "color", "0"),
					resource.TestCheckResourceAttr("fortios_firewallwildcardfqdn_group.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewallwildcardfqdn_group.trname", "visibility", "enable"),
					resource.TestCheckResourceAttr("fortios_firewallwildcardfqdn_group.trname", "member.0.name", var0),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallWildcardFqdnGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallWildcardFqdnGroup: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallWildcardFqdnGroup is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallWildcardFqdnGroup(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading FirewallWildcardFqdnGroup: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallWildcardFqdnGroup: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallWildcardFqdnGroupDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewallwildcardfqdn_group" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallWildcardFqdnGroup(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallWildcardFqdnGroup %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallWildcardFqdnGroupConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewallwildcardfqdn_custom" "trname1" {
  color         = 0
  name          = "var0%[1]s"
  visibility    = "enable"
  wildcard_fqdn = "*.ms.com"
}

resource "fortios_firewallwildcardfqdn_group" "trname" {
  color      = 0
  name       = "%[1]s"
  visibility = "enable"

  member {
    name = fortios_firewallwildcardfqdn_custom.trname1.name
  }
}
`, name)
}
