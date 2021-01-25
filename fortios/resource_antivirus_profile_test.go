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

func TestAccFortiOSAntivirusProfile_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSAntivirusProfile_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSAntivirusProfileConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSAntivirusProfileExists("fortios_antivirus_profile.trname"),
					resource.TestCheckResourceAttr("fortios_antivirus_profile.trname", "analytics_bl_filetype", "0"),
					resource.TestCheckResourceAttr("fortios_antivirus_profile.trname", "analytics_db", "disable"),
					resource.TestCheckResourceAttr("fortios_antivirus_profile.trname", "analytics_max_upload", "10"),
					resource.TestCheckResourceAttr("fortios_antivirus_profile.trname", "analytics_wl_filetype", "0"),
					resource.TestCheckResourceAttr("fortios_antivirus_profile.trname", "av_block_log", "enable"),
					resource.TestCheckResourceAttr("fortios_antivirus_profile.trname", "av_virus_log", "enable"),
					resource.TestCheckResourceAttr("fortios_antivirus_profile.trname", "extended_log", "disable"),
					resource.TestCheckResourceAttr("fortios_antivirus_profile.trname", "ftgd_analytics", "disable"),
					resource.TestCheckResourceAttr("fortios_antivirus_profile.trname", "inspection_mode", "flow-based"),
					resource.TestCheckResourceAttr("fortios_antivirus_profile.trname", "mobile_malware_db", "enable"),
					resource.TestCheckResourceAttr("fortios_antivirus_profile.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_antivirus_profile.trname", "scan_mode", "quick"),
				),
			},
		},
	})
}

func testAccCheckFortiOSAntivirusProfileExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found AntivirusProfile: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No AntivirusProfile is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadAntivirusProfile(i)

		if err != nil {
			return fmt.Errorf("Error reading AntivirusProfile: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating AntivirusProfile: %s", n)
		}

		return nil
	}
}

func testAccCheckAntivirusProfileDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_antivirus_profile" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadAntivirusProfile(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error AntivirusProfile %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSAntivirusProfileConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_antivirus_profile" "trname" {
  analytics_bl_filetype = 0
  analytics_db          = "disable"
  analytics_max_upload  = 10
  analytics_wl_filetype = 0
  av_block_log          = "enable"
  av_virus_log          = "enable"
  extended_log          = "disable"
  ftgd_analytics        = "disable"
  inspection_mode       = "flow-based"
  mobile_malware_db     = "enable"
  name                  = "%[1]s"
  scan_mode             = "quick"
}
`, name)
}
