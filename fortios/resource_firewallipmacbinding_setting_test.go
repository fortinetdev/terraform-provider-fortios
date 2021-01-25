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

func TestAccFortiOSFirewallIpmacbindingSetting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallIpmacbindingSetting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallIpmacbindingSettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallIpmacbindingSettingExists("fortios_firewallipmacbinding_setting.trname"),
					resource.TestCheckResourceAttr("fortios_firewallipmacbinding_setting.trname", "bindthroughfw", "disable"),
					resource.TestCheckResourceAttr("fortios_firewallipmacbinding_setting.trname", "bindtofw", "disable"),
					resource.TestCheckResourceAttr("fortios_firewallipmacbinding_setting.trname", "undefinedhost", "block"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallIpmacbindingSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallIpmacbindingSetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallIpmacbindingSetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallIpmacbindingSetting(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallIpmacbindingSetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallIpmacbindingSetting: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallIpmacbindingSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewallipmacbinding_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallIpmacbindingSetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallIpmacbindingSetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallIpmacbindingSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewallipmacbinding_setting" "trname" {
  bindthroughfw = "disable"
  bindtofw      = "disable"
  undefinedhost = "block"
}
`)
}
