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

func TestAccFortiOSWebfilterContentHeader_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWebfilterContentHeader_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWebfilterContentHeaderConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWebfilterContentHeaderExists("fortios_webfilter_contentheader.trname"),
					resource.TestCheckResourceAttr("fortios_webfilter_contentheader.trname", "fosid", "1"),
					resource.TestCheckResourceAttr("fortios_webfilter_contentheader.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSWebfilterContentHeaderExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WebfilterContentHeader: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebfilterContentHeader is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWebfilterContentHeader(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading WebfilterContentHeader: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WebfilterContentHeader: %s", n)
		}

		return nil
	}
}

func testAccCheckWebfilterContentHeaderDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_webfilter_contentheader" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWebfilterContentHeader(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WebfilterContentHeader %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWebfilterContentHeaderConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_webfilter_contentheader" "trname" {
  fosid = 1
  name  = "%[1]s"
}
`, name)
}
