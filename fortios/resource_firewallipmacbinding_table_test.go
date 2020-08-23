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

func TestAccFortiOSFirewallIpmacbindingTable_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallIpmacbindingTable_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallIpmacbindingTableConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallIpmacbindingTableExists("fortios_firewallipmacbinding_table.trname"),
					resource.TestCheckResourceAttr("fortios_firewallipmacbinding_table.trname", "ip", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_firewallipmacbinding_table.trname", "mac", "00:01:6c:06:a6:29"),
					resource.TestCheckResourceAttr("fortios_firewallipmacbinding_table.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewallipmacbinding_table.trname", "seq_num", "1"),
					resource.TestCheckResourceAttr("fortios_firewallipmacbinding_table.trname", "status", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallIpmacbindingTableExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallIpmacbindingTable: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallIpmacbindingTable is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallIpmacbindingTable(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallIpmacbindingTable: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallIpmacbindingTable: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallIpmacbindingTableDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewallipmacbinding_table" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallIpmacbindingTable(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallIpmacbindingTable %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallIpmacbindingTableConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewallipmacbinding_table" "trname" {
  ip      = "1.1.1.1"
  mac     = "00:01:6c:06:a6:29"
  name    = "%[1]s"
  seq_num = 1
  status  = "disable"
}
`, name)
}
