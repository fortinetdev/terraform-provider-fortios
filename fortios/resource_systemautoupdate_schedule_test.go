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

func TestAccFortiOSSystemAutoupdateSchedule_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemAutoupdateSchedule_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemAutoupdateScheduleConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemAutoupdateScheduleExists("fortios_systemautoupdate_schedule.trname"),
					resource.TestCheckResourceAttr("fortios_systemautoupdate_schedule.trname", "day", "Monday"),
					resource.TestCheckResourceAttr("fortios_systemautoupdate_schedule.trname", "frequency", "every"),
					resource.TestCheckResourceAttr("fortios_systemautoupdate_schedule.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_systemautoupdate_schedule.trname", "time", "02:60"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemAutoupdateScheduleExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemAutoupdateSchedule: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemAutoupdateSchedule is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemAutoupdateSchedule(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading SystemAutoupdateSchedule: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemAutoupdateSchedule: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemAutoupdateScheduleDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_systemautoupdate_schedule" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemAutoupdateSchedule(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemAutoupdateSchedule %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemAutoupdateScheduleConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_systemautoupdate_schedule" "trname" {
  day       = "Monday"
  frequency = "every"
  status    = "enable"
  time      = "02:60"
}
`)
}
