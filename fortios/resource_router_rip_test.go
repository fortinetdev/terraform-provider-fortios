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

func TestAccFortiOSRouterRip_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSRouterRip_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSRouterRipConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSRouterRipExists("fortios_router_rip.trname"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "default_information_originate", "disable"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "default_metric", "1"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "garbage_timer", "120"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "max_out_metric", "0"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "recv_buffer_size", "655360"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "timeout_timer", "180"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "update_timer", "30"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "version", "2"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "redistribute.0.metric", "10"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "redistribute.0.name", "connected"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "redistribute.0.status", "disable"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "redistribute.1.metric", "10"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "redistribute.1.name", "static"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "redistribute.1.status", "disable"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "redistribute.2.metric", "10"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "redistribute.2.name", "ospf"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "redistribute.2.status", "disable"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "redistribute.3.metric", "10"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "redistribute.3.name", "bgp"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "redistribute.3.status", "disable"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "redistribute.4.metric", "10"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "redistribute.4.name", "isis"),
					resource.TestCheckResourceAttr("fortios_router_rip.trname", "redistribute.4.status", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSRouterRipExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterRip: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterRip is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterRip(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading RouterRip: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterRip: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterRipDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_rip" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterRip(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterRip %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterRipConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_rip" "trname" {
  default_information_originate = "disable"
  default_metric                = 1
  garbage_timer                 = 120
  max_out_metric                = 0
  recv_buffer_size              = 655360
  timeout_timer                 = 180
  update_timer                  = 30
  version                       = "2"

  redistribute {
    metric = 10
    name   = "connected"
    status = "disable"
  }
  redistribute {
    metric = 10
    name   = "static"
    status = "disable"
  }
  redistribute {
    metric = 10
    name   = "ospf"
    status = "disable"
  }
  redistribute {
    metric = 10
    name   = "bgp"
    status = "disable"
  }
  redistribute {
    metric = 10
    name   = "isis"
    status = "disable"
  }
}
`)
}
