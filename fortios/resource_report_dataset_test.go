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

func TestAccFortiOSReportDataset_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSReportDataset_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSReportDatasetConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSReportDatasetExists("fortios_report_dataset.trname"),
					resource.TestCheckResourceAttr("fortios_report_dataset.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_report_dataset.trname", "policy", "0"),
					resource.TestCheckResourceAttr("fortios_report_dataset.trname", "query", "select * from testdb"),
				),
			},
		},
	})
}

func testAccCheckFortiOSReportDatasetExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found ReportDataset: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ReportDataset is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadReportDataset(i)

		if err != nil {
			return fmt.Errorf("Error reading ReportDataset: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating ReportDataset: %s", n)
		}

		return nil
	}
}

func testAccCheckReportDatasetDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_report_dataset" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadReportDataset(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error ReportDataset %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSReportDatasetConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_report_dataset" "trname" {
  name   = "%[1]s"
  policy = 0
  query  = "select * from testdb"
}
`, name)
}
