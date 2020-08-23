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

func TestAccFortiOSLogMemoryGlobalSetting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogMemoryGlobalSetting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogMemoryGlobalSettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogMemoryGlobalSettingExists("fortios_logmemory_globalsetting.trname"),
					resource.TestCheckResourceAttr("fortios_logmemory_globalsetting.trname", "full_final_warning_threshold", "95"),
					resource.TestCheckResourceAttr("fortios_logmemory_globalsetting.trname", "full_first_warning_threshold", "75"),
					resource.TestCheckResourceAttr("fortios_logmemory_globalsetting.trname", "full_second_warning_threshold", "90"),
					resource.TestCheckResourceAttr("fortios_logmemory_globalsetting.trname", "max_size", "163840"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogMemoryGlobalSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogMemoryGlobalSetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogMemoryGlobalSetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogMemoryGlobalSetting(i)

		if err != nil {
			return fmt.Errorf("Error reading LogMemoryGlobalSetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogMemoryGlobalSetting: %s", n)
		}

		return nil
	}
}

func testAccCheckLogMemoryGlobalSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_logmemory_globalsetting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogMemoryGlobalSetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogMemoryGlobalSetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogMemoryGlobalSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_logmemory_globalsetting" "trname" {
  full_final_warning_threshold  = 95
  full_first_warning_threshold  = 75
  full_second_warning_threshold = 90
  max_size                      = 163840
}
`)
}
