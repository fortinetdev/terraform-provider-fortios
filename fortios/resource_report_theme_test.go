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

func TestAccFortiOSReportTheme_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSReportTheme_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSReportThemeConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSReportThemeExists("fortios_report_theme.trname"),
					resource.TestCheckResourceAttr("fortios_report_theme.trname", "column_count", "1"),
					resource.TestCheckResourceAttr("fortios_report_theme.trname", "graph_chart_style", "PS"),
					resource.TestCheckResourceAttr("fortios_report_theme.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_report_theme.trname", "page_orient", "portrait"),
				),
			},
		},
	})
}

func testAccCheckFortiOSReportThemeExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found ReportTheme: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ReportTheme is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadReportTheme(i)

		if err != nil {
			return fmt.Errorf("Error reading ReportTheme: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating ReportTheme: %s", n)
		}

		return nil
	}
}

func testAccCheckReportThemeDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_report_theme" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadReportTheme(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error ReportTheme %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSReportThemeConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_report_theme" "trname" {
  column_count      = "1"
  graph_chart_style = "PS"
  name              = "%[1]s"
  page_orient       = "portrait"
}

`, name)
}
