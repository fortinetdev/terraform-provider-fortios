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

func TestAccFortiOSSwitchControllerLldpSettings_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSwitchControllerLldpSettings_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSwitchControllerLldpSettingsConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSwitchControllerLldpSettingsExists("fortios_switchcontroller_lldpsettings.trname"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_lldpsettings.trname", "fast_start_interval", "2"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_lldpsettings.trname", "management_interface", "internal"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_lldpsettings.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_lldpsettings.trname", "tx_hold", "4"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_lldpsettings.trname", "tx_interval", "30"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSwitchControllerLldpSettingsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SwitchControllerLldpSettings: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SwitchControllerLldpSettings is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerLldpSettings(i)

		if err != nil {
			return fmt.Errorf("Error reading SwitchControllerLldpSettings: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SwitchControllerLldpSettings: %s", n)
		}

		return nil
	}
}

func testAccCheckSwitchControllerLldpSettingsDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_switchcontroller_lldpsettings" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerLldpSettings(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SwitchControllerLldpSettings %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSwitchControllerLldpSettingsConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_switchcontroller_lldpsettings" "trname" {
  fast_start_interval  = 2
  management_interface = "internal"
  status               = "enable"
  tx_hold              = 4
  tx_interval          = 30
}
`)
}
