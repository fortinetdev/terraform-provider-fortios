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

func TestAccFortiOSWebfilterOverride_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWebfilterOverride_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWebfilterOverrideConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWebfilterOverrideExists("fortios_webfilter_override.trname"),
					resource.TestCheckResourceAttr("fortios_webfilter_override.trname", "expires", "2021/03/16 19:18:25"),
					resource.TestCheckResourceAttr("fortios_webfilter_override.trname", "fosid", "1"),
					resource.TestCheckResourceAttr("fortios_webfilter_override.trname", "ip", "69.101.119.0"),
					resource.TestCheckResourceAttr("fortios_webfilter_override.trname", "ip6", "4565:7700::"),
					resource.TestCheckResourceAttr("fortios_webfilter_override.trname", "new_profile", "monitor-all"),
					resource.TestCheckResourceAttr("fortios_webfilter_override.trname", "old_profile", "default"),
					resource.TestCheckResourceAttr("fortios_webfilter_override.trname", "scope", "user"),
					resource.TestCheckResourceAttr("fortios_webfilter_override.trname", "status", "disable"),
					resource.TestCheckResourceAttr("fortios_webfilter_override.trname", "user", "Eew"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWebfilterOverrideExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WebfilterOverride: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebfilterOverride is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWebfilterOverride(i)

		if err != nil {
			return fmt.Errorf("Error reading WebfilterOverride: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WebfilterOverride: %s", n)
		}

		return nil
	}
}

func testAccCheckWebfilterOverrideDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_webfilter_override" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWebfilterOverride(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WebfilterOverride %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWebfilterOverrideConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_webfilter_override" "trname" {
  expires     = "2021/03/16 19:18:25"
  fosid       = 1
  ip          = "69.101.119.0"
  ip6         = "4565:7700::"
  new_profile = "monitor-all"
  old_profile = "default"
  scope       = "user"
  status      = "disable"
  user        = "Eew"

}
`)
}
