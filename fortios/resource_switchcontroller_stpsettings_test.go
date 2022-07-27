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

func TestAccFortiOSSwitchControllerStpSettings_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSwitchControllerStpSettings_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSwitchControllerStpSettingsConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSwitchControllerStpSettingsExists("fortios_switchcontroller_stpsettings.trname"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_stpsettings.trname", "forward_time", "15"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_stpsettings.trname", "hello_time", "2"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_stpsettings.trname", "max_age", "20"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_stpsettings.trname", "max_hops", "20"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_stpsettings.trname", "pending_timer", "4"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_stpsettings.trname", "revision", "0"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_stpsettings.trname", "status", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSwitchControllerStpSettingsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SwitchControllerStpSettings: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SwitchControllerStpSettings is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerStpSettings(i)

		if err != nil {
			return fmt.Errorf("Error reading SwitchControllerStpSettings: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SwitchControllerStpSettings: %s", n)
		}

		return nil
	}
}

func testAccCheckSwitchControllerStpSettingsDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_switchcontroller_stpsettings" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerStpSettings(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SwitchControllerStpSettings %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSwitchControllerStpSettingsConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_switchcontroller_stpsettings" "trname" {
  forward_time  = 15
  hello_time    = 2
  max_age       = 20
  max_hops      = 20
  pending_timer = 4
  revision      = 0
  status        = "enable"
}
`)
}
