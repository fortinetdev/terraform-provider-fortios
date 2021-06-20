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

func TestAccFortiOSFirewallServiceCategory_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallServiceCategory_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallServiceCategoryConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallServiceCategoryExists("fortios_firewallservice_category.trname"),
					resource.TestCheckResourceAttr("fortios_firewallservice_category.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallServiceCategoryExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallServiceCategory: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallServiceCategory is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallServiceCategory(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading FirewallServiceCategory: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallServiceCategory: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallServiceCategoryDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewallservice_category" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallServiceCategory(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallServiceCategory %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallServiceCategoryConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewallservice_category" "trname" {
  name = "%[1]s"
}
`, name)
}
