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

func TestAccFortiOSRouterSetting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSRouterSetting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSRouterSettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSRouterSettingExists("fortios_router_setting.trname"),
					resource.TestCheckResourceAttr("fortios_router_setting.trname", "hostname", "s1"),
				),
			},
		},
	})
}

func testAccCheckFortiOSRouterSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterSetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterSetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterSetting(i)

		if err != nil {
			return fmt.Errorf("Error reading RouterSetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterSetting: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterSetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterSetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_setting" "trname" {
  hostname = "s1"
}
`)
}
