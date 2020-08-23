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

func TestAccFortiOSSystemCsf_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemCsf_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemCsfConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemCsfExists("fortios_system_csf.trname"),
					resource.TestCheckResourceAttr("fortios_system_csf.trname", "configuration_sync", "default"),
					resource.TestCheckResourceAttr("fortios_system_csf.trname", "management_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_system_csf.trname", "management_port", "33"),
					resource.TestCheckResourceAttr("fortios_system_csf.trname", "status", "disable"),
					resource.TestCheckResourceAttr("fortios_system_csf.trname", "upstream_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_system_csf.trname", "upstream_port", "8013"),
					resource.TestCheckResourceAttr("fortios_system_csf.trname", "group_password", "tmp"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemCsfExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemCsf: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemCsf is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemCsf(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemCsf: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemCsf: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemCsfDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_csf" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemCsf(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemCsf %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemCsfConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_csf" "trname" {
  configuration_sync = "default"
  management_ip      = "0.0.0.0"
  management_port    = 33
  status             = "disable"
  upstream_ip        = "0.0.0.0"
  upstream_port      = 8013
  group_password     = "tmp"
}
`)
}
