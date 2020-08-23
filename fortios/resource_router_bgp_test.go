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

func TestAccFortiOSRouterBgp_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSRouterBgp_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSRouterBgpConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSRouterBgpExists("fortios_router_bgp.trname"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "additional_path_select", "2"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "additional_path_select6", "2"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "always_compare_med", "disable"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "as", "0"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "client_to_client_reflection", "enable"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "cluster_id", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "dampening", "disable"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "dampening_max_suppress_time", "60"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "dampening_reachability_half_life", "15"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "dampening_reuse", "750"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "dampening_suppress", "2000"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "dampening_unreachability_half_life", "15"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "default_local_preference", "100"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "deterministic_med", "disable"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "distance_external", "20"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "distance_internal", "200"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "distance_local", "200"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "graceful_restart_time", "120"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "graceful_stalepath_time", "360"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "graceful_update_delay", "120"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "holdtime_timer", "180"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "ibgp_multipath", "disable"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "ignore_optional_capability", "enable"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "keepalive_timer", "60"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "log_neighbour_changes", "enable"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "network_import_check", "enable"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "scan_time", "60"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "synchronization", "disable"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute.0.name", "connected"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute.0.status", "disable"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute.1.name", "rip"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute.1.status", "disable"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute.2.name", "ospf"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute.2.status", "disable"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute.3.name", "static"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute.3.status", "disable"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute.4.name", "isis"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute.4.status", "disable"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute6.0.name", "connected"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute6.0.status", "disable"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute6.1.name", "rip"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute6.1.status", "disable"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute6.2.name", "ospf"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute6.2.status", "disable"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute6.3.name", "static"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute6.3.status", "disable"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute6.4.name", "isis"),
					resource.TestCheckResourceAttr("fortios_router_bgp.trname", "redistribute6.4.status", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSRouterBgpExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterBgp: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterBgp is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterBgp(i)

		if err != nil {
			return fmt.Errorf("Error reading RouterBgp: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterBgp: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterBgpDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_bgp" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterBgp(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterBgp %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterBgpConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_bgp" "trname" {
  additional_path_select             = 2
  additional_path_select6            = 2
  always_compare_med                 = "disable"
  as                                 = 0
  client_to_client_reflection        = "enable"
  cluster_id                         = "0.0.0.0"
  dampening                          = "disable"
  dampening_max_suppress_time        = 60
  dampening_reachability_half_life   = 15
  dampening_reuse                    = 750
  dampening_suppress                 = 2000
  dampening_unreachability_half_life = 15
  default_local_preference           = 100
  deterministic_med                  = "disable"
  distance_external                  = 20
  distance_internal                  = 200
  distance_local                     = 200
  graceful_restart_time              = 120
  graceful_stalepath_time            = 360
  graceful_update_delay              = 120
  holdtime_timer                     = 180
  ibgp_multipath                     = "disable"
  ignore_optional_capability         = "enable"
  keepalive_timer                    = 60
  log_neighbour_changes              = "enable"
  network_import_check               = "enable"
  scan_time                          = 60
  synchronization                    = "disable"

  redistribute {
    name   = "connected"
    status = "disable"
  }
  redistribute {
    name   = "rip"
    status = "disable"
  }
  redistribute {
    name   = "ospf"
    status = "disable"
  }
  redistribute {
    name   = "static"
    status = "disable"
  }
  redistribute {
    name   = "isis"
    status = "disable"
  }

  redistribute6 {
    name   = "connected"
    status = "disable"
  }
  redistribute6 {
    name   = "rip"
    status = "disable"
  }
  redistribute6 {
    name   = "ospf"
    status = "disable"
  }
  redistribute6 {
    name   = "static"
    status = "disable"
  }
  redistribute6 {
    name   = "isis"
    status = "disable"
  }
}
`)
}
