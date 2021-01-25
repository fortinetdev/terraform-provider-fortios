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

func TestAccFortiOSSystemConsole_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemConsole_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemConsoleConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemConsoleExists("fortios_system_console.trname"),
					resource.TestCheckResourceAttr("fortios_system_console.trname", "baudrate", "9600"),
					resource.TestCheckResourceAttr("fortios_system_console.trname", "login", "enable"),
					resource.TestCheckResourceAttr("fortios_system_console.trname", "mode", "line"),
					resource.TestCheckResourceAttr("fortios_system_console.trname", "output", "more"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemConsoleExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemConsole: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemConsole is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemConsole(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemConsole: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemConsole: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemConsoleDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_console" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemConsole(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemConsole %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemConsoleConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_console" "trname" {
  baudrate = "9600"
  login    = "enable"
  mode     = "line"
  output   = "more"
}
`)
}
