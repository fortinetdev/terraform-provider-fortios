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

func TestAccFortiOSSystemVirtualWanLink_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemVirtualWanLink_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemVirtualWanLinkConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemVirtualWanLinkExists("fortios_system_virtualwanlink.trname"),
					resource.TestCheckResourceAttr("fortios_system_virtualwanlink.trname", "fail_detect", "disable"),
					resource.TestCheckResourceAttr("fortios_system_virtualwanlink.trname", "load_balance_mode", "source-ip-based"),
					resource.TestCheckResourceAttr("fortios_system_virtualwanlink.trname", "status", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemVirtualWanLinkExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemVirtualWanLink: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemVirtualWanLink is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemVirtualWanLink(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading SystemVirtualWanLink: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemVirtualWanLink: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemVirtualWanLinkDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_virtualwanlink" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemVirtualWanLink(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemVirtualWanLink %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemVirtualWanLinkConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_virtualwanlink" "trname" {
  fail_detect       = "disable"
  load_balance_mode = "source-ip-based"
  status            = "disable"
}
`)
}
