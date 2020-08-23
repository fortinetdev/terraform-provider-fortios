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

func TestAccFortiOSWebfilterFtgdLocalCat_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWebfilterFtgdLocalCat_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWebfilterFtgdLocalCatConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWebfilterFtgdLocalCatExists("fortios_webfilter_ftgdlocalcat.trname"),
					resource.TestCheckResourceAttr("fortios_webfilter_ftgdlocalcat.trname", "desc", "s1"),
					resource.TestCheckResourceAttr("fortios_webfilter_ftgdlocalcat.trname", "fosid", "188"),
					resource.TestCheckResourceAttr("fortios_webfilter_ftgdlocalcat.trname", "status", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWebfilterFtgdLocalCatExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WebfilterFtgdLocalCat: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebfilterFtgdLocalCat is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWebfilterFtgdLocalCat(i)

		if err != nil {
			return fmt.Errorf("Error reading WebfilterFtgdLocalCat: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WebfilterFtgdLocalCat: %s", n)
		}

		return nil
	}
}

func testAccCheckWebfilterFtgdLocalCatDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_webfilter_ftgdlocalcat" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWebfilterFtgdLocalCat(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WebfilterFtgdLocalCat %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWebfilterFtgdLocalCatConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_webfilter_ftgdlocalcat" "trname" {
  desc   = "s1"
  fosid  = 188
  status = "enable"
}
`)
}
