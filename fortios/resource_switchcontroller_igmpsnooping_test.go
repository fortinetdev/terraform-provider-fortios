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

func TestAccFortiOSSwitchControllerIgmpSnooping_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSwitchControllerIgmpSnooping_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSwitchControllerIgmpSnoopingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSwitchControllerIgmpSnoopingExists("fortios_switchcontroller_igmpsnooping.trname"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_igmpsnooping.trname", "aging_time", "300"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_igmpsnooping.trname", "flood_unknown_multicast", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSwitchControllerIgmpSnoopingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SwitchControllerIgmpSnooping: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SwitchControllerIgmpSnooping is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerIgmpSnooping(i)

		if err != nil {
			return fmt.Errorf("Error reading SwitchControllerIgmpSnooping: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SwitchControllerIgmpSnooping: %s", n)
		}

		return nil
	}
}

func testAccCheckSwitchControllerIgmpSnoopingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_switchcontroller_igmpsnooping" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerIgmpSnooping(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SwitchControllerIgmpSnooping %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSwitchControllerIgmpSnoopingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_switchcontroller_igmpsnooping" "trname" {
  aging_time              = 300
  flood_unknown_multicast = "disable"
}
`)
}
