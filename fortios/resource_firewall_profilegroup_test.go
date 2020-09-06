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

func TestAccFortiOSFirewallProfileGroup_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallProfileGroup_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallProfileGroupConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallProfileGroupExists("fortios_firewall_profilegroup.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_profilegroup.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_profilegroup.trname", "profile_protocol_options", "default"),
					resource.TestCheckResourceAttr("fortios_firewall_profilegroup.trname", "ssl_ssh_profile", "deep-inspection"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallProfileGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallProfileGroup: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallProfileGroup is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallProfileGroup(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallProfileGroup: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallProfileGroup: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallProfileGroupDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_profilegroup" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallProfileGroup(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallProfileGroup %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallProfileGroupConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_profilegroup" "trname" {
  name                     = "%[1]s"
  profile_protocol_options = "default"
  ssl_ssh_profile          = "deep-inspection"
}
`, name)
}
