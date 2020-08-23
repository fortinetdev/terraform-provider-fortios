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

func TestAccFortiOSSystemVdomException_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemVdomException_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemVdomExceptionConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemVdomExceptionExists("fortios_system_vdomexception.trname"),
					resource.TestCheckResourceAttr("fortios_system_vdomexception.trname", "fosid", "1"),
					resource.TestCheckResourceAttr("fortios_system_vdomexception.trname", "object", "log.fortianalyzer.setting"),
					resource.TestCheckResourceAttr("fortios_system_vdomexception.trname", "oid", "7150"),
					resource.TestCheckResourceAttr("fortios_system_vdomexception.trname", "scope", "all"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemVdomExceptionExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemVdomException: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemVdomException is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemVdomException(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemVdomException: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemVdomException: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemVdomExceptionDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_vdomexception" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemVdomException(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemVdomException %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemVdomExceptionConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_vdomexception" "trname" {
  fosid  = 1
  object = "log.fortianalyzer.setting"
  oid    = 7150
  scope  = "all"
}


`)
}
