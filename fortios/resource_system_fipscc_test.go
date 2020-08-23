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

func TestAccFortiOSSystemFipsCc_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemFipsCc_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemFipsCcConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemFipsCcExists("fortios_system_fipscc.trname"),
					resource.TestCheckResourceAttr("fortios_system_fipscc.trname", "entropy_token", "enable"),
					resource.TestCheckResourceAttr("fortios_system_fipscc.trname", "key_generation_self_test", "disable"),
					resource.TestCheckResourceAttr("fortios_system_fipscc.trname", "self_test_period", "1440"),
					resource.TestCheckResourceAttr("fortios_system_fipscc.trname", "status", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemFipsCcExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemFipsCc: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemFipsCc is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemFipsCc(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemFipsCc: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemFipsCc: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemFipsCcDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_fipscc" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemFipsCc(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemFipsCc %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemFipsCcConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_fipscc" "trname" {
  entropy_token            = "enable"
  key_generation_self_test = "disable"
  self_test_period         = 1440
  status                   = "disable"
}




`)
}
