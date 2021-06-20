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

func TestAccFortiOSSystemZone_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemZone_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemZoneConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemZoneExists("fortios_system_zone.trname"),
					resource.TestCheckResourceAttr("fortios_system_zone.trname", "intrazone", "allow"),
					resource.TestCheckResourceAttr("fortios_system_zone.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemZoneExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemZone: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemZone is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemZone(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading SystemZone: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemZone: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemZoneDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_zone" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemZone(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemZone %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemZoneConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_zone" "trname" {
  intrazone = "allow"
  name      = "%[1]s"
}
`, name)
}
