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

func TestAccFortiOSAntivirusSettings_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSAntivirusSettings_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSAntivirusSettingsConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSAntivirusSettingsExists("fortios_antivirus_settings.trname"),
					resource.TestCheckResourceAttr("fortios_antivirus_settings.trname", "default_db", "extended"),
					resource.TestCheckResourceAttr("fortios_antivirus_settings.trname", "grayware", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSAntivirusSettingsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found AntivirusSettings: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No AntivirusSettings is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadAntivirusSettings(i)

		if err != nil {
			return fmt.Errorf("Error reading AntivirusSettings: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating AntivirusSettings: %s", n)
		}

		return nil
	}
}

func testAccCheckAntivirusSettingsDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_antivirus_settings" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadAntivirusSettings(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error AntivirusSettings %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSAntivirusSettingsConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_antivirus_settings" "trname" {
  default_db = "extended"
  grayware   = "enable"
}
`)
}
