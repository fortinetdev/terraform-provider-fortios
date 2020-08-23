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

func TestAccFortiOSSystemAutoupdatePushUpdate_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemAutoupdatePushUpdate_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemAutoupdatePushUpdateConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemAutoupdatePushUpdateExists("fortios_systemautoupdate_pushupdate.trname"),
					resource.TestCheckResourceAttr("fortios_systemautoupdate_pushupdate.trname", "address", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_systemautoupdate_pushupdate.trname", "override", "disable"),
					resource.TestCheckResourceAttr("fortios_systemautoupdate_pushupdate.trname", "port", "9443"),
					resource.TestCheckResourceAttr("fortios_systemautoupdate_pushupdate.trname", "status", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemAutoupdatePushUpdateExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemAutoupdatePushUpdate: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemAutoupdatePushUpdate is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemAutoupdatePushUpdate(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemAutoupdatePushUpdate: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemAutoupdatePushUpdate: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemAutoupdatePushUpdateDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_systemautoupdate_pushupdate" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemAutoupdatePushUpdate(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemAutoupdatePushUpdate %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemAutoupdatePushUpdateConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_systemautoupdate_pushupdate" "trname" {
  address  = "0.0.0.0"
  override = "disable"
  port     = 9443
  status   = "disable"
}
`)
}
