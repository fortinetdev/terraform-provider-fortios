// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSWafProfile_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWafProfile_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWafProfileConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWafProfileExists("fortios_waf_profile.trname"),
					resource.TestCheckResourceAttr("fortios_waf_profile.trname", "extended_log", "disable"),
					resource.TestCheckResourceAttr("fortios_waf_profile.trname", "external", "disable"),
					resource.TestCheckResourceAttr("fortios_waf_profile.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSWafProfileExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WafProfile: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WafProfile is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWafProfile(i)

		if err != nil {
			return fmt.Errorf("Error reading WafProfile: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WafProfile: %s", n)
		}

		return nil
	}
}

func testAccCheckWafProfileDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_waf_profile" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWafProfile(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WafProfile %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWafProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_waf_profile" "trname" {
  extended_log = "disable"
  external     = "disable"
  name         = "%[1]s"
}
`, name)
}
