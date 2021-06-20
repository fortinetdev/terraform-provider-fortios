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

func TestAccFortiOSSwitchControllerCustomCommand_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSwitchControllerCustomCommand_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSwitchControllerCustomCommandConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSwitchControllerCustomCommandExists("fortios_switchcontroller_customcommand.trname"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_customcommand.trname", "command", "ls"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_customcommand.trname", "command_name", "1"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSwitchControllerCustomCommandExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SwitchControllerCustomCommand: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SwitchControllerCustomCommand is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerCustomCommand(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading SwitchControllerCustomCommand: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SwitchControllerCustomCommand: %s", n)
		}

		return nil
	}
}

func testAccCheckSwitchControllerCustomCommandDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_switchcontroller_customcommand" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerCustomCommand(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SwitchControllerCustomCommand %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSwitchControllerCustomCommandConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_switchcontroller_customcommand" "trname" {
  command      = "ls"
  command_name = "1"
}
`)
}
