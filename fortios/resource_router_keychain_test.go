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

func TestAccFortiOSRouterKeyChain_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSRouterKeyChain_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSRouterKeyChainConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSRouterKeyChainExists("fortios_router_keychain.trname"),
					resource.TestCheckResourceAttr("fortios_router_keychain.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_router_keychain.trname", "key.0.accept_lifetime", "04:00:00 01 01 2008 04:00:00 01 01 2022"),
					resource.TestCheckResourceAttr("fortios_router_keychain.trname", "key.0.key_string", "ewiwn3i23232s212"),
					resource.TestCheckResourceAttr("fortios_router_keychain.trname", "key.0.send_lifetime", "04:00:00 01 01 2008 04:00:00 01 01 2022"),
				),
			},
		},
	})
}

func testAccCheckFortiOSRouterKeyChainExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterKeyChain: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterKeyChain is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterKeyChain(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading RouterKeyChain: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterKeyChain: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterKeyChainDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_keychain" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterKeyChain(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterKeyChain %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterKeyChainConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_keychain" "trname" {
  name = "%[1]s"

  key {
    accept_lifetime = "04:00:00 01 01 2008 04:00:00 01 01 2022"
    key_string      = "ewiwn3i23232s212"
    send_lifetime   = "04:00:00 01 01 2008 04:00:00 01 01 2022"
  }
}
`, name)
}
