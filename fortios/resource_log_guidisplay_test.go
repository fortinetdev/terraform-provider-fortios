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

func TestAccFortiOSLogGuiDisplay_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogGuiDisplay_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogGuiDisplayConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogGuiDisplayExists("fortios_log_guidisplay.trname"),
					resource.TestCheckResourceAttr("fortios_log_guidisplay.trname", "fortiview_unscanned_apps", "disable"),
					resource.TestCheckResourceAttr("fortios_log_guidisplay.trname", "resolve_apps", "enable"),
					resource.TestCheckResourceAttr("fortios_log_guidisplay.trname", "resolve_hosts", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogGuiDisplayExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogGuiDisplay: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogGuiDisplay is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogGuiDisplay(i)

		if err != nil {
			return fmt.Errorf("Error reading LogGuiDisplay: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogGuiDisplay: %s", n)
		}

		return nil
	}
}

func testAccCheckLogGuiDisplayDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_log_guidisplay" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogGuiDisplay(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogGuiDisplay %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogGuiDisplayConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_log_guidisplay" "trname" {
  fortiview_unscanned_apps = "disable"
  resolve_apps             = "enable"
  resolve_hosts            = "enable"
}
`)
}
