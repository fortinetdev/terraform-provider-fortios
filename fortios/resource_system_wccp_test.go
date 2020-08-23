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

func TestAccFortiOSSystemWccp_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemWccp_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemWccpConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemWccpExists("fortios_system_wccp.trname"),
					resource.TestCheckResourceAttr("fortios_system_wccp.trname", "assignment_bucket_format", "cisco-implementation"),
					resource.TestCheckResourceAttr("fortios_system_wccp.trname", "assignment_dstaddr_mask", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_system_wccp.trname", "assignment_method", "HASH"),
					resource.TestCheckResourceAttr("fortios_system_wccp.trname", "assignment_srcaddr_mask", "0.0.23.65"),
					resource.TestCheckResourceAttr("fortios_system_wccp.trname", "assignment_weight", "0"),
					resource.TestCheckResourceAttr("fortios_system_wccp.trname", "authentication", "disable"),
					resource.TestCheckResourceAttr("fortios_system_wccp.trname", "cache_engine_method", "GRE"),
					resource.TestCheckResourceAttr("fortios_system_wccp.trname", "cache_id", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_system_wccp.trname", "forward_method", "GRE"),
					resource.TestCheckResourceAttr("fortios_system_wccp.trname", "group_address", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_system_wccp.trname", "primary_hash", "dst-ip"),
					resource.TestCheckResourceAttr("fortios_system_wccp.trname", "priority", "0"),
					resource.TestCheckResourceAttr("fortios_system_wccp.trname", "protocol", "0"),
					resource.TestCheckResourceAttr("fortios_system_wccp.trname", "return_method", "GRE"),
					resource.TestCheckResourceAttr("fortios_system_wccp.trname", "router_id", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_system_wccp.trname", "router_list", "\"1.0.0.0\" "),
					resource.TestCheckResourceAttr("fortios_system_wccp.trname", "server_type", "forward"),
					resource.TestCheckResourceAttr("fortios_system_wccp.trname", "service_id", "1"),
					resource.TestCheckResourceAttr("fortios_system_wccp.trname", "service_type", "auto"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemWccpExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemWccp: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemWccp is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemWccp(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemWccp: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemWccp: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemWccpDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_wccp" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemWccp(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemWccp %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemWccpConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_wccp" "trname" {
  assignment_bucket_format = "cisco-implementation"
  assignment_dstaddr_mask  = "0.0.0.0"
  assignment_method        = "HASH"
  assignment_srcaddr_mask  = "0.0.23.65"
  assignment_weight        = 0
  authentication           = "disable"
  cache_engine_method      = "GRE"
  cache_id                 = "1.1.1.1"
  forward_method           = "GRE"
  group_address            = "0.0.0.0"
  primary_hash             = "dst-ip"
  priority                 = 0
  protocol                 = 0
  return_method            = "GRE"
  router_id                = "1.1.1.1"
  router_list              = "\"1.0.0.0\" "
  server_type              = "forward"
  service_id               = "1"
  service_type             = "auto"
}
`)
}
