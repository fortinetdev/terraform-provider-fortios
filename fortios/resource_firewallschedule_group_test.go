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

func TestAccFortiOSFirewallScheduleGroup_basic(t *testing.T) {
	rname := acctest.RandString(8)
	var0 := "var0" + rname
	log.Printf(var0)
	log.Printf("TestAccFortiOSFirewallScheduleGroup_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallScheduleGroupConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallScheduleGroupExists("fortios_firewallschedule_group.trname"),
					resource.TestCheckResourceAttr("fortios_firewallschedule_group.trname", "color", "0"),
					resource.TestCheckResourceAttr("fortios_firewallschedule_group.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewallschedule_group.trname", "member.0.name", var0),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallScheduleGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallScheduleGroup: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallScheduleGroup is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallScheduleGroup(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallScheduleGroup: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallScheduleGroup: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallScheduleGroupDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewallschedule_group" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallScheduleGroup(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallScheduleGroup %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallScheduleGroupConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewallschedule_recurring" "trname1" {
  color = 0
  day   = "sunday"
  end   = "00:00"
  name  = "var0%[1]s"
  start = "00:00"
}

resource "fortios_firewallschedule_group" "trname" {
  color = 0
  name  = "%[1]s"

  member {
    name = fortios_firewallschedule_recurring.trname1.name
  }
}
`, name)
}
