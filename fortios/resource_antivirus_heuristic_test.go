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

func TestAccFortiOSAntivirusHeuristic_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSAntivirusHeuristic_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSAntivirusHeuristicConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSAntivirusHeuristicExists("fortios_antivirus_heuristic.trname"),
					resource.TestCheckResourceAttr("fortios_antivirus_heuristic.trname", "mode", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSAntivirusHeuristicExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found AntivirusHeuristic: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No AntivirusHeuristic is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadAntivirusHeuristic(i)

		if err != nil {
			return fmt.Errorf("Error reading AntivirusHeuristic: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating AntivirusHeuristic: %s", n)
		}

		return nil
	}
}

func testAccCheckAntivirusHeuristicDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_antivirus_heuristic" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadAntivirusHeuristic(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error AntivirusHeuristic %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSAntivirusHeuristicConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_antivirus_heuristic" "trname" {
  mode = "disable"
}
`)
}
