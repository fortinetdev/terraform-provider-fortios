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

func TestAccFortiOSSystemGlobal_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemGlobal_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemGlobalConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemGlobalExists("fortios_system_global.trname"),
					resource.TestCheckResourceAttr("fortios_system_global.trname", "admin_sport", "443"),
					resource.TestCheckResourceAttr("fortios_system_global.trname", "alias", "FGVM02TM20003062"),
					resource.TestCheckResourceAttr("fortios_system_global.trname", "hostname", "ste11"),
					resource.TestCheckResourceAttr("fortios_system_global.trname", "timezone", "04"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemGlobalExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemGlobal: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemGlobal is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemGlobal(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemGlobal: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemGlobal: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemGlobalDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_global" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemGlobal(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemGlobal %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemGlobalConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_global" "trname" {
  admin_sport = 443
  alias       = "FGVM02TM20003062"
  hostname    = "ste11"
  timezone    = "04"
}
`)
}
