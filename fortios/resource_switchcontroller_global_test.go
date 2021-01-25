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

func TestAccFortiOSSwitchControllerGlobal_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSwitchControllerGlobal_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSwitchControllerGlobalConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSwitchControllerGlobalExists("fortios_switchcontroller_global.trname"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_global.trname", "allow_multiple_interfaces", "disable"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_global.trname", "https_image_push", "disable"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_global.trname", "log_mac_limit_violations", "disable"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_global.trname", "mac_aging_interval", "332"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_global.trname", "mac_retention_period", "24"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_global.trname", "mac_violation_timer", "0"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSwitchControllerGlobalExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SwitchControllerGlobal: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SwitchControllerGlobal is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerGlobal(i)

		if err != nil {
			return fmt.Errorf("Error reading SwitchControllerGlobal: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SwitchControllerGlobal: %s", n)
		}

		return nil
	}
}

func testAccCheckSwitchControllerGlobalDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_switchcontroller_global" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerGlobal(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SwitchControllerGlobal %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSwitchControllerGlobalConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_switchcontroller_global" "trname" {
  allow_multiple_interfaces = "disable"
  https_image_push          = "disable"
  log_mac_limit_violations  = "disable"
  mac_aging_interval        = 332
  mac_retention_period      = 24
  mac_violation_timer       = 0
}
`)
}
