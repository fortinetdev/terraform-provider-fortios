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

func TestAccFortiOSReportChart_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSReportChart_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSReportChartConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSReportChartExists("fortios_report_chart.trname"),
					resource.TestCheckResourceAttr("fortios_report_chart.trname", "category", "misc"),
					resource.TestCheckResourceAttr("fortios_report_chart.trname", "comments", "test report chart"),
					resource.TestCheckResourceAttr("fortios_report_chart.trname", "dataset", "s1"),
					resource.TestCheckResourceAttr("fortios_report_chart.trname", "dimension", "3D"),
					resource.TestCheckResourceAttr("fortios_report_chart.trname", "favorite", "no"),
					resource.TestCheckResourceAttr("fortios_report_chart.trname", "graph_type", "none"),
					resource.TestCheckResourceAttr("fortios_report_chart.trname", "legend", "enable"),
					resource.TestCheckResourceAttr("fortios_report_chart.trname", "legend_font_size", "0"),
					resource.TestCheckResourceAttr("fortios_report_chart.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_report_chart.trname", "period", "last24h"),
					resource.TestCheckResourceAttr("fortios_report_chart.trname", "policy", "0"),
					resource.TestCheckResourceAttr("fortios_report_chart.trname", "style", "auto"),
					resource.TestCheckResourceAttr("fortios_report_chart.trname", "title_font_size", "0"),
					resource.TestCheckResourceAttr("fortios_report_chart.trname", "type", "graph"),
				),
			},
		},
	})
}

func testAccCheckFortiOSReportChartExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found ReportChart: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ReportChart is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadReportChart(i)

		if err != nil {
			return fmt.Errorf("Error reading ReportChart: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating ReportChart: %s", n)
		}

		return nil
	}
}

func testAccCheckReportChartDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_report_chart" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadReportChart(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error ReportChart %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSReportChartConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_report_chart" "trname" {
  category         = "misc"
  comments         = "test report chart"
  dataset          = "s1"
  dimension        = "3D"
  favorite         = "no"
  graph_type       = "none"
  legend           = "enable"
  legend_font_size = 0
  name             = "%[1]s"
  period           = "last24h"
  policy           = 0
  style            = "auto"
  title_font_size  = 0
  type             = "graph"
}
`, name)
}
