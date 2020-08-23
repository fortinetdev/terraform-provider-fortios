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

func TestAccFortiOSReportStyle_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSReportStyle_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSReportStyleConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSReportStyleExists("fortios_report_style.trname"),
					resource.TestCheckResourceAttr("fortios_report_style.trname", "border_bottom", "\" none \""),
					resource.TestCheckResourceAttr("fortios_report_style.trname", "border_left", "\" none \""),
					resource.TestCheckResourceAttr("fortios_report_style.trname", "border_right", "\" none \""),
					resource.TestCheckResourceAttr("fortios_report_style.trname", "border_top", "\" none \""),
					resource.TestCheckResourceAttr("fortios_report_style.trname", "column_span", "none"),
					resource.TestCheckResourceAttr("fortios_report_style.trname", "font_style", "normal"),
					resource.TestCheckResourceAttr("fortios_report_style.trname", "font_weight", "normal"),
					resource.TestCheckResourceAttr("fortios_report_style.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_report_style.trname", "options", "font text color"),
				),
			},
		},
	})
}

func testAccCheckFortiOSReportStyleExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found ReportStyle: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ReportStyle is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadReportStyle(i)

		if err != nil {
			return fmt.Errorf("Error reading ReportStyle: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating ReportStyle: %s", n)
		}

		return nil
	}
}

func testAccCheckReportStyleDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_report_style" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadReportStyle(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error ReportStyle %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSReportStyleConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_report_style" "trname" {
  border_bottom = "\" none \""
  border_left   = "\" none \""
  border_right  = "\" none \""
  border_top    = "\" none \""
  column_span   = "none"
  font_style    = "normal"
  font_weight   = "normal"
  name          = "%[1]s"
  options       = "font text color"
}
`, name)
}
