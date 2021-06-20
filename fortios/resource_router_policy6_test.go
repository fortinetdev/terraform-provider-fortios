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

func TestAccFortiOSRouterPolicy6_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSRouterPolicy6_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSRouterPolicy6Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSRouterPolicy6Exists("fortios_router_policy6.trname"),
					resource.TestCheckResourceAttr("fortios_router_policy6.trname", "dst", "::/0"),
					resource.TestCheckResourceAttr("fortios_router_policy6.trname", "end_port", "65535"),
					resource.TestCheckResourceAttr("fortios_router_policy6.trname", "gateway", "::"),
					resource.TestCheckResourceAttr("fortios_router_policy6.trname", "input_device", "port1"),
					resource.TestCheckResourceAttr("fortios_router_policy6.trname", "output_device", "port3"),
					resource.TestCheckResourceAttr("fortios_router_policy6.trname", "protocol", "33"),
					resource.TestCheckResourceAttr("fortios_router_policy6.trname", "seq_num", "1"),
					resource.TestCheckResourceAttr("fortios_router_policy6.trname", "src", "2001:db8:85a3::8a2e:370:7334/128"),
					resource.TestCheckResourceAttr("fortios_router_policy6.trname", "start_port", "1"),
					resource.TestCheckResourceAttr("fortios_router_policy6.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_router_policy6.trname", "tos", "0x00"),
					resource.TestCheckResourceAttr("fortios_router_policy6.trname", "tos_mask", "0x00"),
				),
			},
		},
	})
}

func testAccCheckFortiOSRouterPolicy6Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterPolicy6: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterPolicy6 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterPolicy6(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading RouterPolicy6: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterPolicy6: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterPolicy6Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_policy6" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterPolicy6(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterPolicy6 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterPolicy6Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_policy6" "trname" {
  dst           = "::/0"
  end_port      = 65535
  gateway       = "::"
  input_device  = "port1"
  output_device = "port3"
  protocol      = 33
  seq_num       = 1
  src           = "2001:db8:85a3::8a2e:370:7334/128"
  start_port    = 1
  status        = "enable"
  tos           = "0x00"
  tos_mask      = "0x00"
}
`)
}
