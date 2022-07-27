// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"log"
	"testing"
)

func TestAccFortiOSDlpSettings_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSDlpSettings_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSDlpSettingsConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSDlpSettingsExists("fortios_dlp_settings.trname"),
					resource.TestCheckResourceAttr("fortios_dlp_settings.trname", "cache_mem_percent", "2"),
					resource.TestCheckResourceAttr("fortios_dlp_settings.trname", "chunk_size", "2800"),
					resource.TestCheckResourceAttr("fortios_dlp_settings.trname", "db_mode", "stop-adding"),
					resource.TestCheckResourceAttr("fortios_dlp_settings.trname", "size", "16"),
				),
			},
		},
	})
}

func testAccCheckFortiOSDlpSettingsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found DlpSettings: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No DlpSettings is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadDlpSettings(i)

		if err != nil {
			return fmt.Errorf("Error reading DlpSettings: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating DlpSettings: %s", n)
		}

		return nil
	}
}

func testAccCheckDlpSettingsDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_dlp_settings" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadDlpSettings(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error DlpSettings %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSDlpSettingsConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_dlp_settings" "trname" {
  cache_mem_percent = 2
  chunk_size        = 2800
  db_mode           = "stop-adding"
  size              = 16
}
`)
}
