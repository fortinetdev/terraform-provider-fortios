
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

func TestAccFortiOSRouterRouteMap_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSRouterRouteMap_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSRouterRouteMapConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSRouterRouteMapExists("fortios_router_routemap.trname"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "name", rname),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.action", "deny"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.match_community_exact", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.match_flags", "0"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.match_metric", "0"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.match_origin", "none"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.match_route_type", "No type specified"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.match_tag", "0"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_aggregator_as", "0"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_aggregator_ip", "0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_aspath_action", "prepend"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_atomic_aggregate", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_community_additive", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_dampening_max_suppress", "0"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_dampening_reachability_half_life", "0"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_dampening_reuse", "0"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_dampening_suppress", "0"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_dampening_unreachability_half_life", "0"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_flags", "128"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_ip6_nexthop", "::"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_ip6_nexthop_local", "::"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_ip_nexthop", "0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_local_preference", "0"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_metric", "0"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_metric_type", "No type specified"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_origin", "none"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_originator_id", "0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_route_tag", "0"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_tag", "0"),
                    resource.TestCheckResourceAttr("fortios_router_routemap.trname", "rule.0.set_weight", "21"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSRouterRouteMapExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterRouteMap: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterRouteMap is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterRouteMap(i)

		if err != nil {
			return fmt.Errorf("Error reading RouterRouteMap: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterRouteMap: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterRouteMapDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_routemap" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterRouteMap(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterRouteMap %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterRouteMapConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_routemap" "trname" {
  name = "%[1]s"

  rule {
    action                                 = "deny"
    match_community_exact                  = "disable"
    match_flags                            = 0
    match_metric                           = 0
    match_origin                           = "none"
    match_route_type                       = "No type specified"
    match_tag                              = 0
    set_aggregator_as                      = 0
    set_aggregator_ip                      = "0.0.0.0"
    set_aspath_action                      = "prepend"
    set_atomic_aggregate                   = "disable"
    set_community_additive                 = "disable"
    set_dampening_max_suppress             = 0
    set_dampening_reachability_half_life   = 0
    set_dampening_reuse                    = 0
    set_dampening_suppress                 = 0
    set_dampening_unreachability_half_life = 0
    set_flags                              = 128
    set_ip6_nexthop                        = "::"
    set_ip6_nexthop_local                  = "::"
    set_ip_nexthop                         = "0.0.0.0"
    set_local_preference                   = 0
    set_metric                             = 0
    set_metric_type                        = "No type specified"
    set_origin                             = "none"
    set_originator_id                      = "0.0.0.0"
    set_route_tag                          = 0
    set_tag                                = 0
    set_weight                             = 21
  }
}
`, name)
}
