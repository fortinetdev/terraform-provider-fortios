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

func TestAccFortiOSSystemSettings_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemSettings_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemSettingsConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemSettingsExists("fortios_system_settings.trname"),
					resource.TestCheckResourceAttr("fortios_system_settings.trname", "allow_linkdown_path", "disable"),
					resource.TestCheckResourceAttr("fortios_system_settings.trname", "gui_webfilter", "enable"),
					resource.TestCheckResourceAttr("fortios_system_settings.trname", "opmode", "nat"),
					resource.TestCheckResourceAttr("fortios_system_settings.trname", "sip_ssl_port", "5061"),
					resource.TestCheckResourceAttr("fortios_system_settings.trname", "status", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemSettingsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemSettings: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemSettings is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemSettings(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemSettings: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemSettings: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemSettingsDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_settings" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemSettings(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemSettings %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemSettingsConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_settings" "trname" {
  allow_linkdown_path = "disable"
  gui_webfilter       = "enable"
  opmode              = "nat"
  sip_ssl_port        = 5061
  status              = "enable"
}
`)
}
