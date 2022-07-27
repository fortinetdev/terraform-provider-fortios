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

func TestAccFortiOSSystemAutoScript_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemAutoScript_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemAutoScriptConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemAutoScriptExists("fortios_system_autoscript.trname"),
					resource.TestCheckResourceAttr("fortios_system_autoscript.trname", "interval", "3"),
					resource.TestCheckResourceAttr("fortios_system_autoscript.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_system_autoscript.trname", "output_size", "10"),
					resource.TestCheckResourceAttr("fortios_system_autoscript.trname", "repeat", "2"),
					resource.TestCheckResourceAttr("fortios_system_autoscript.trname", "script", "action"),
					resource.TestCheckResourceAttr("fortios_system_autoscript.trname", "start", "manual"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemAutoScriptExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemAutoScript: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemAutoScript is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemAutoScript(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemAutoScript: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemAutoScript: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemAutoScriptDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_autoscript" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemAutoScript(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemAutoScript %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemAutoScriptConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_autoscript" "trname" {
  interval    = 3
  name        = "%[1]s"
  output_size = 10
  repeat      = 2
  script      = "action"
  start       = "manual"
}
`, name)
}
