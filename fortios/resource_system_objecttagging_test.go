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

func TestAccFortiOSSystemObjectTagging_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemObjectTagging_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemObjectTaggingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemObjectTaggingExists("fortios_system_objecttagging.trname"),
					resource.TestCheckResourceAttr("fortios_system_objecttagging.trname", "address", "disable"),
					resource.TestCheckResourceAttr("fortios_system_objecttagging.trname", "category", "s1"),
					resource.TestCheckResourceAttr("fortios_system_objecttagging.trname", "color", "0"),
					resource.TestCheckResourceAttr("fortios_system_objecttagging.trname", "device", "mandatory"),
					resource.TestCheckResourceAttr("fortios_system_objecttagging.trname", "interface", "disable"),
					resource.TestCheckResourceAttr("fortios_system_objecttagging.trname", "multiple", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemObjectTaggingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemObjectTagging: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemObjectTagging is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemObjectTagging(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemObjectTagging: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemObjectTagging: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemObjectTaggingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_objecttagging" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemObjectTagging(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemObjectTagging %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemObjectTaggingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_objecttagging" "trname" {
  address   = "disable"
  category  = "s1"
  color     = 0
  device    = "mandatory"
  interface = "disable"
  multiple  = "enable"
}
`)
}
