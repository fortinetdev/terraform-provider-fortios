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

func TestAccFortiOSRouterPrefixList_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSRouterPrefixList_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSRouterPrefixListConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSRouterPrefixListExists("fortios_router_prefixlist.trname"),
					resource.TestCheckResourceAttr("fortios_router_prefixlist.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSRouterPrefixListExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterPrefixList: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterPrefixList is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterPrefixList(i)

		if err != nil {
			return fmt.Errorf("Error reading RouterPrefixList: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterPrefixList: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterPrefixListDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_prefixlist" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterPrefixList(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterPrefixList %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterPrefixListConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_prefixlist" "trname" {
  name = "%[1]s"
}
`, name)
}
