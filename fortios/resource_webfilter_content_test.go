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

func TestAccFortiOSWebfilterContent_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWebfilterContent_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWebfilterContentConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWebfilterContentExists("fortios_webfilter_content.trname"),
					resource.TestCheckResourceAttr("fortios_webfilter_content.trname", "fosid", "1"),
					resource.TestCheckResourceAttr("fortios_webfilter_content.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSWebfilterContentExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WebfilterContent: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebfilterContent is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWebfilterContent(i)

		if err != nil {
			return fmt.Errorf("Error reading WebfilterContent: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WebfilterContent: %s", n)
		}

		return nil
	}
}

func testAccCheckWebfilterContentDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_webfilter_content" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWebfilterContent(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WebfilterContent %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWebfilterContentConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_webfilter_content" "trname" {
  fosid = 1
  name  = "%[1]s"
}
`, name)
}
