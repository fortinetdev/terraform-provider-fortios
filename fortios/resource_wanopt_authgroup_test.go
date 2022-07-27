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

func TestAccFortiOSWanoptAuthGroup_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWanoptAuthGroup_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWanoptAuthGroupConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWanoptAuthGroupExists("fortios_wanopt_authgroup.trname"),
					resource.TestCheckResourceAttr("fortios_wanopt_authgroup.trname", "auth_method", "cert"),
					resource.TestCheckResourceAttr("fortios_wanopt_authgroup.trname", "cert", "Fortinet_CA_SSL"),
					resource.TestCheckResourceAttr("fortios_wanopt_authgroup.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_wanopt_authgroup.trname", "peer_accept", "any"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWanoptAuthGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WanoptAuthGroup: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WanoptAuthGroup is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWanoptAuthGroup(i)

		if err != nil {
			return fmt.Errorf("Error reading WanoptAuthGroup: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WanoptAuthGroup: %s", n)
		}

		return nil
	}
}

func testAccCheckWanoptAuthGroupDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_wanopt_authgroup" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWanoptAuthGroup(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WanoptAuthGroup %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWanoptAuthGroupConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_wanopt_authgroup" "trname" {
  auth_method = "cert"
  cert        = "Fortinet_CA_SSL"
  name        = "%[1]s"
  peer_accept = "any"
}
`, name)
}
