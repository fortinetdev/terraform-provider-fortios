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

func TestAccFortiOSWirelessControllerInterController_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWirelessControllerInterController_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWirelessControllerInterControllerConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWirelessControllerInterControllerExists("fortios_wirelesscontroller_intercontroller.trname"),
					resource.TestCheckResourceAttr("fortios_wirelesscontroller_intercontroller.trname", "fast_failover_max", "10"),
					resource.TestCheckResourceAttr("fortios_wirelesscontroller_intercontroller.trname", "fast_failover_wait", "10"),
					resource.TestCheckResourceAttr("fortios_wirelesscontroller_intercontroller.trname", "inter_controller_key", "ENC XXXX"),
					resource.TestCheckResourceAttr("fortios_wirelesscontroller_intercontroller.trname", "inter_controller_mode", "disable"),
					resource.TestCheckResourceAttr("fortios_wirelesscontroller_intercontroller.trname", "inter_controller_pri", "primary"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWirelessControllerInterControllerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WirelessControllerInterController: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WirelessControllerInterController is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWirelessControllerInterController(i)

		if err != nil {
			return fmt.Errorf("Error reading WirelessControllerInterController: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WirelessControllerInterController: %s", n)
		}

		return nil
	}
}

func testAccCheckWirelessControllerInterControllerDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_wirelesscontroller_intercontroller" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWirelessControllerInterController(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WirelessControllerInterController %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWirelessControllerInterControllerConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_wirelesscontroller_intercontroller" "trname" {
  fast_failover_max     = 10
  fast_failover_wait    = 10
  inter_controller_key  = "ENC XXXX"
  inter_controller_mode = "disable"
  inter_controller_pri  = "primary"
}
`)
}
