
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

func TestAccFortiOSRouterOspf6_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSRouterOspf6_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSRouterOspf6Config(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSRouterOspf6Exists("fortios_router_ospf6.trname"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "abr_type", "standard"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "auto_cost_ref_bandwidth", "1000"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "bfd", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "default_information_metric", "10"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "default_information_metric_type", "2"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "default_information_originate", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "default_metric", "10"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "log_neighbour_changes", "enable"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "router_id", "0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "spf_timers", "5 10"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.0.metric", "0"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.0.metric_type", "2"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.0.name", "connected"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.0.status", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.1.metric", "0"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.1.metric_type", "2"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.1.name", "static"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.1.status", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.2.metric", "0"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.2.metric_type", "2"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.2.name", "rip"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.2.status", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.3.metric", "0"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.3.metric_type", "2"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.3.name", "bgp"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.3.status", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.4.metric", "0"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.4.metric_type", "2"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.4.name", "isis"),
                    resource.TestCheckResourceAttr("fortios_router_ospf6.trname", "redistribute.4.status", "disable"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSRouterOspf6Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterOspf6: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterOspf6 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterOspf6(i)

		if err != nil {
			return fmt.Errorf("Error reading RouterOspf6: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterOspf6: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterOspf6Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_ospf6" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterOspf6(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterOspf6 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterOspf6Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_ospf6" "trname" {
  abr_type                        = "standard"
  auto_cost_ref_bandwidth         = 1000
  bfd                             = "disable"
  default_information_metric      = 10
  default_information_metric_type = "2"
  default_information_originate   = "disable"
  default_metric                  = 10
  log_neighbour_changes           = "enable"
  router_id                       = "0.0.0.0"
  spf_timers                      = "5 10"

  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "connected"
    status      = "disable"
  }
  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "static"
    status      = "disable"
  }
  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "rip"
    status      = "disable"
  }
  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "bgp"
    status      = "disable"
  }
  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "isis"
    status      = "disable"
  }
}

`)
}
