// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSReportSetting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSReportSetting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSReportSettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSReportSettingExists("fortios_report_setting.trname"),
					resource.TestCheckResourceAttr("fortios_report_setting.trname", "fortiview", "enable"),
					resource.TestCheckResourceAttr("fortios_report_setting.trname", "pdf_report", "enable"),
					resource.TestCheckResourceAttr("fortios_report_setting.trname", "report_source", "forward-traffic"),
					resource.TestCheckResourceAttr("fortios_report_setting.trname", "top_n", "1000"),
					resource.TestCheckResourceAttr("fortios_report_setting.trname", "web_browsing_threshold", "3"),
				),
			},
		},
	})
}

func testAccCheckFortiOSReportSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found ReportSetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ReportSetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadReportSetting(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading ReportSetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating ReportSetting: %s", n)
		}

		return nil
	}
}

func testAccCheckReportSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_report_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadReportSetting(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error ReportSetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSReportSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_report_setting" "trname" {
  fortiview              = "enable"
  pdf_report             = "enable"
  report_source          = "forward-traffic"
  top_n                  = 1000
  web_browsing_threshold = 3
}
`)
}
