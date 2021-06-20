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

func TestAccFortiOSWebfilterIpsUrlfilterSetting6_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWebfilterIpsUrlfilterSetting6_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWebfilterIpsUrlfilterSetting6Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWebfilterIpsUrlfilterSetting6Exists("fortios_webfilter_ipsurlfiltersetting6.trname"),
					resource.TestCheckResourceAttr("fortios_webfilter_ipsurlfiltersetting6.trname", "distance", "1"),
					resource.TestCheckResourceAttr("fortios_webfilter_ipsurlfiltersetting6.trname", "gateway6", "::"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWebfilterIpsUrlfilterSetting6Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WebfilterIpsUrlfilterSetting6: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebfilterIpsUrlfilterSetting6 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWebfilterIpsUrlfilterSetting6(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading WebfilterIpsUrlfilterSetting6: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WebfilterIpsUrlfilterSetting6: %s", n)
		}

		return nil
	}
}

func testAccCheckWebfilterIpsUrlfilterSetting6Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_webfilter_ipsurlfiltersetting6" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWebfilterIpsUrlfilterSetting6(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WebfilterIpsUrlfilterSetting6 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWebfilterIpsUrlfilterSetting6Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_webfilter_ipsurlfiltersetting6" "trname" {
  distance = 1
  gateway6 = "::"
}
`)
}
