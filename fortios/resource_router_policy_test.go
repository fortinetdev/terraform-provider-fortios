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

func TestAccFortiOSRouterPolicy_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSRouterPolicy_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSRouterPolicyConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSRouterPolicyExists("fortios_router_policy.trname"),
					resource.TestCheckResourceAttr("fortios_router_policy.trname", "action", "permit"),
					resource.TestCheckResourceAttr("fortios_router_policy.trname", "dst_negate", "disable"),
					resource.TestCheckResourceAttr("fortios_router_policy.trname", "end_port", "25"),
					resource.TestCheckResourceAttr("fortios_router_policy.trname", "end_source_port", "65535"),
					resource.TestCheckResourceAttr("fortios_router_policy.trname", "gateway", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_router_policy.trname", "output_device", "port2"),
					resource.TestCheckResourceAttr("fortios_router_policy.trname", "protocol", "6"),
					resource.TestCheckResourceAttr("fortios_router_policy.trname", "seq_num", "1"),
					resource.TestCheckResourceAttr("fortios_router_policy.trname", "src_negate", "disable"),
					resource.TestCheckResourceAttr("fortios_router_policy.trname", "start_port", "25"),
					resource.TestCheckResourceAttr("fortios_router_policy.trname", "start_source_port", "0"),
					resource.TestCheckResourceAttr("fortios_router_policy.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_router_policy.trname", "tos", "0x00"),
					resource.TestCheckResourceAttr("fortios_router_policy.trname", "tos_mask", "0x00"),
					resource.TestCheckResourceAttr("fortios_router_policy.trname", "input_device.0.name", "port1"),
				),
			},
		},
	})
}

func testAccCheckFortiOSRouterPolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterPolicy: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterPolicy is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterPolicy(i)

		if err != nil {
			return fmt.Errorf("Error reading RouterPolicy: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterPolicy: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterPolicyDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_policy" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterPolicy(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterPolicy %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterPolicyConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_policy" "trname" {
  action            = "permit"
  dst_negate        = "disable"
  end_port          = 25
  end_source_port   = 65535
  gateway           = "0.0.0.0"
  output_device     = "port2"
  protocol          = 6
  seq_num           = 1
  src_negate        = "disable"
  start_port        = 25
  start_source_port = 0
  status            = "enable"
  tos               = "0x00"
  tos_mask          = "0x00"

  input_device {
    name = "port1"
  }
}
`)
}
