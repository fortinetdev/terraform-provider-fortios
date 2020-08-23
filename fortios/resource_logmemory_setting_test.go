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

func TestAccFortiOSLogMemorySetting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogMemorySetting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogMemorySettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogMemorySettingExists("fortios_logmemory_setting.trname"),
					resource.TestCheckResourceAttr("fortios_logmemory_setting.trname", "diskfull", "overwrite"),
					resource.TestCheckResourceAttr("fortios_logmemory_setting.trname", "status", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogMemorySettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogMemorySetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogMemorySetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogMemorySetting(i)

		if err != nil {
			return fmt.Errorf("Error reading LogMemorySetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogMemorySetting: %s", n)
		}

		return nil
	}
}

func testAccCheckLogMemorySettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_logmemory_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogMemorySetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogMemorySetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogMemorySettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_logmemory_setting" "trname" {
  diskfull = "overwrite"
  status   = "disable"
}
`)
}
