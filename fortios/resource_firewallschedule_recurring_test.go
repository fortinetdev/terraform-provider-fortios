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

func TestAccFortiOSFirewallScheduleRecurring_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallScheduleRecurring_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallScheduleRecurringConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallScheduleRecurringExists("fortios_firewallschedule_recurring.trname"),
					resource.TestCheckResourceAttr("fortios_firewallschedule_recurring.trname", "color", "0"),
					resource.TestCheckResourceAttr("fortios_firewallschedule_recurring.trname", "day", "sunday"),
					resource.TestCheckResourceAttr("fortios_firewallschedule_recurring.trname", "end", "00:00"),
					resource.TestCheckResourceAttr("fortios_firewallschedule_recurring.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewallschedule_recurring.trname", "start", "00:00"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallScheduleRecurringExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallScheduleRecurring: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallScheduleRecurring is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallScheduleRecurring(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallScheduleRecurring: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallScheduleRecurring: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallScheduleRecurringDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewallschedule_recurring" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallScheduleRecurring(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallScheduleRecurring %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallScheduleRecurringConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewallschedule_recurring" "trname" {
  color = 0
  day   = "sunday"
  end   = "00:00"
  name  = "%[1]s"
  start = "00:00"
}
`, name)
}
