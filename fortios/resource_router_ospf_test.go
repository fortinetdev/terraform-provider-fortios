
// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios
import (
    "fmt"
	"log"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSRouterOspf_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSRouterOspf_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSRouterOspfConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSRouterOspfExists("fortios_router_ospf.trname"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "abr_type", "standard"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "auto_cost_ref_bandwidth", "1000"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "bfd", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "database_overflow", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "database_overflow_max_lsas", "10000"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "database_overflow_time_to_recover", "300"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "default_information_metric", "10"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "default_information_metric_type", "2"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "default_information_originate", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "default_metric", "10"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "distance", "110"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "distance_external", "110"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "distance_inter_area", "110"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "distance_intra_area", "110"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "log_neighbour_changes", "enable"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "restart_mode", "none"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "restart_period", "120"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "rfc1583_compatible", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "router_id", "0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "spf_timers", "5 10"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.0.metric", "0"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.0.metric_type", "2"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.0.name", "connected"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.0.status", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.0.tag", "0"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.1.metric", "0"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.1.metric_type", "2"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.1.name", "static"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.1.status", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.1.tag", "0"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.2.metric", "0"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.2.metric_type", "2"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.2.name", "rip"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.2.status", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.2.tag", "0"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.3.metric", "0"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.3.metric_type", "2"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.3.name", "bgp"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.3.status", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.3.tag", "0"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.4.metric", "0"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.4.metric_type", "2"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.4.name", "isis"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.4.status", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_ospf.trname", "redistribute.4.tag", "0"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSRouterOspfExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterOspf: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterOspf is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterOspf(i)

		if err != nil {
			return fmt.Errorf("Error reading RouterOspf: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterOspf: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterOspfDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_ospf" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterOspf(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterOspf %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterOspfConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_ospf" "trname" {
  abr_type                          = "standard"
  auto_cost_ref_bandwidth           = 1000
  bfd                               = "disable"
  database_overflow                 = "disable"
  database_overflow_max_lsas        = 10000
  database_overflow_time_to_recover = 300
  default_information_metric        = 10
  default_information_metric_type   = "2"
  default_information_originate     = "disable"
  default_metric                    = 10
  distance                          = 110
  distance_external                 = 110
  distance_inter_area               = 110
  distance_intra_area               = 110
  log_neighbour_changes             = "enable"
  restart_mode                      = "none"
  restart_period                    = 120
  rfc1583_compatible                = "disable"
  router_id                         = "0.0.0.0"
  spf_timers                        = "5 10"

  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "connected"
    status      = "disable"
    tag         = 0
  }
  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "static"
    status      = "disable"
    tag         = 0
  }
  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "rip"
    status      = "disable"
    tag         = 0
  }
  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "bgp"
    status      = "disable"
    tag         = 0
  }
  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "isis"
    status      = "disable"
    tag         = 0
  }
}


`)
}
