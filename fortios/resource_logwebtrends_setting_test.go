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

func TestAccFortiOSLogWebtrendsSetting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogWebtrendsSetting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogWebtrendsSettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogWebtrendsSettingExists("fortios_logwebtrends_setting.trname"),
					resource.TestCheckResourceAttr("fortios_logwebtrends_setting.trname", "status", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogWebtrendsSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogWebtrendsSetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogWebtrendsSetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogWebtrendsSetting(i)

		if err != nil {
			return fmt.Errorf("Error reading LogWebtrendsSetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogWebtrendsSetting: %s", n)
		}

		return nil
	}
}

func testAccCheckLogWebtrendsSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_logwebtrends_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogWebtrendsSetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogWebtrendsSetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogWebtrendsSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_logwebtrends_setting" "trname" {
  status = "disable"
}
`)
}
