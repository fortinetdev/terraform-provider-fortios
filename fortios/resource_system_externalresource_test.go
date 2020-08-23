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

func TestAccFortiOSSystemExternalResource_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemExternalResource_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemExternalResourceConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemExternalResourceExists("fortios_system_externalresource.trname"),
					resource.TestCheckResourceAttr("fortios_system_externalresource.trname", "category", "199"),
					resource.TestCheckResourceAttr("fortios_system_externalresource.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_system_externalresource.trname", "refresh_rate", "5"),
					resource.TestCheckResourceAttr("fortios_system_externalresource.trname", "resource", "https://tmpxxxxx.com"),
					resource.TestCheckResourceAttr("fortios_system_externalresource.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_system_externalresource.trname", "type", "category"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemExternalResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemExternalResource: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemExternalResource is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemExternalResource(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemExternalResource: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemExternalResource: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemExternalResourceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_externalresource" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemExternalResource(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemExternalResource %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemExternalResourceConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_externalresource" "trname" {
  category     = 199
  name         = "%[1]s"
  refresh_rate = 5
  resource     = "https://tmpxxxxx.com"
  status       = "enable"
  type         = "category"
}

`, name)
}
