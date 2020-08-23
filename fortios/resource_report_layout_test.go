
// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios
import (
    "fmt"
	"log"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSReportLayout_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSReportLayout_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSReportLayoutConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSReportLayoutExists("fortios_report_layout.trname"),
                    resource.TestCheckResourceAttr("fortios_report_layout.trname", "cutoff_option", "run-time"),
                    resource.TestCheckResourceAttr("fortios_report_layout.trname", "cutoff_time", "00:00"),
                    resource.TestCheckResourceAttr("fortios_report_layout.trname", "day", "sunday"),
                    resource.TestCheckResourceAttr("fortios_report_layout.trname", "email_send", "disable"),
                    resource.TestCheckResourceAttr("fortios_report_layout.trname", "format", "pdf"),
                    resource.TestCheckResourceAttr("fortios_report_layout.trname", "max_pdf_report", "31"),
                    resource.TestCheckResourceAttr("fortios_report_layout.trname", "name", rname),
                    resource.TestCheckResourceAttr("fortios_report_layout.trname", "options", "include-table-of-content view-chart-as-heading"),
                    resource.TestCheckResourceAttr("fortios_report_layout.trname", "schedule_type", "daily"),
                    resource.TestCheckResourceAttr("fortios_report_layout.trname", "style_theme", "default-report"),
                    resource.TestCheckResourceAttr("fortios_report_layout.trname", "time", "00:00"),
                    resource.TestCheckResourceAttr("fortios_report_layout.trname", "title", "FortiGate System Analysis Report"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSReportLayoutExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found ReportLayout: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ReportLayout is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadReportLayout(i)

		if err != nil {
			return fmt.Errorf("Error reading ReportLayout: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating ReportLayout: %s", n)
		}

		return nil
	}
}

func testAccCheckReportLayoutDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_report_layout" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadReportLayout(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error ReportLayout %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSReportLayoutConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_report_layout" "trname" {
  cutoff_option  = "run-time"
  cutoff_time    = "00:00"
  day            = "sunday"
  email_send     = "disable"
  format         = "pdf"
  max_pdf_report = 31
  name           = "%[1]s"
  options        = "include-table-of-content view-chart-as-heading"
  schedule_type  = "daily"
  style_theme    = "default-report"
  time           = "00:00"
  title          = "FortiGate System Analysis Report"
}
`, name)
}
