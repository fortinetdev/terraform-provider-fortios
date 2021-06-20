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

func TestAccFortiOSSystemFm_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemFm_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemFmConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemFmExists("fortios_system_fm.trname"),
					resource.TestCheckResourceAttr("fortios_system_fm.trname", "auto_backup", "disable"),
					resource.TestCheckResourceAttr("fortios_system_fm.trname", "ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_system_fm.trname", "ipsec", "disable"),
					resource.TestCheckResourceAttr("fortios_system_fm.trname", "scheduled_config_restore", "disable"),
					resource.TestCheckResourceAttr("fortios_system_fm.trname", "status", "disable"),
					resource.TestCheckResourceAttr("fortios_system_fm.trname", "vdom", "root"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemFmExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemFm: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemFm is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemFm(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading SystemFm: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemFm: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemFmDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_fm" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemFm(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemFm %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemFmConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_fm" "trname" {
  auto_backup              = "disable"
  ip                       = "0.0.0.0"
  ipsec                    = "disable"
  scheduled_config_restore = "disable"
  status                   = "disable"
  vdom                     = "root"
}
`)
}
