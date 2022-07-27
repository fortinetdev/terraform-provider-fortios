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

func TestAccFortiOSWebfilterFtgdLocalRating_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWebfilterFtgdLocalRating_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWebfilterFtgdLocalRatingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWebfilterFtgdLocalRatingExists("fortios_webfilter_ftgdlocalrating.trname"),
					resource.TestCheckResourceAttr("fortios_webfilter_ftgdlocalrating.trname", "rating", "1"),
					resource.TestCheckResourceAttr("fortios_webfilter_ftgdlocalrating.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_ftgdlocalrating.trname", "url", "sgala.com"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWebfilterFtgdLocalRatingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WebfilterFtgdLocalRating: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebfilterFtgdLocalRating is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWebfilterFtgdLocalRating(i)

		if err != nil {
			return fmt.Errorf("Error reading WebfilterFtgdLocalRating: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WebfilterFtgdLocalRating: %s", n)
		}

		return nil
	}
}

func testAccCheckWebfilterFtgdLocalRatingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_webfilter_ftgdlocalrating" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWebfilterFtgdLocalRating(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WebfilterFtgdLocalRating %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWebfilterFtgdLocalRatingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_webfilter_ftgdlocalrating" "trname" {
  rating = "1"
  status = "enable"
  url    = "sgala.com"
}
`)
}
