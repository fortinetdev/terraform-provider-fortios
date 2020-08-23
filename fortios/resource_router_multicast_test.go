
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

func TestAccFortiOSRouterMulticast_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSRouterMulticast_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSRouterMulticastConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSRouterMulticastExists("fortios_router_multicast.trname"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "multicast_routing", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "route_limit", "2147483647"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "route_threshold", "2147483647"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "pim_sm_global.0.bsr_allow_quick_refresh", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "pim_sm_global.0.bsr_candidate", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "pim_sm_global.0.bsr_hash", "10"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "pim_sm_global.0.bsr_priority", "0"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "pim_sm_global.0.cisco_crp_prefix", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "pim_sm_global.0.cisco_ignore_rp_set_priority", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "pim_sm_global.0.cisco_register_checksum", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "pim_sm_global.0.join_prune_holdtime", "210"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "pim_sm_global.0.message_interval", "60"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "pim_sm_global.0.null_register_retries", "1"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "pim_sm_global.0.register_rate_limit", "0"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "pim_sm_global.0.register_rp_reachability", "enable"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "pim_sm_global.0.register_source", "disable"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "pim_sm_global.0.register_source_ip", "0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "pim_sm_global.0.register_supression", "60"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "pim_sm_global.0.rp_register_keepalive", "185"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "pim_sm_global.0.spt_threshold", "enable"),
                    resource.TestCheckResourceAttr("fortios_router_multicast.trname", "pim_sm_global.0.ssm", "disable"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSRouterMulticastExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found RouterMulticast: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RouterMulticast is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadRouterMulticast(i)

		if err != nil {
			return fmt.Errorf("Error reading RouterMulticast: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating RouterMulticast: %s", n)
		}

		return nil
	}
}

func testAccCheckRouterMulticastDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_router_multicast" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadRouterMulticast(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error RouterMulticast %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSRouterMulticastConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_router_multicast" "trname" {
  multicast_routing = "disable"
  route_limit       = 2147483647
  route_threshold   = 2147483647

  pim_sm_global {
    bsr_allow_quick_refresh      = "disable"
    bsr_candidate                = "disable"
    bsr_hash                     = 10
    bsr_priority                 = 0
    cisco_crp_prefix             = "disable"
    cisco_ignore_rp_set_priority = "disable"
    cisco_register_checksum      = "disable"
    join_prune_holdtime          = 210
    message_interval             = 60
    null_register_retries        = 1
    register_rate_limit          = 0
    register_rp_reachability     = "enable"
    register_source              = "disable"
    register_source_ip           = "0.0.0.0"
    register_supression          = 60
    rp_register_keepalive        = 185
    spt_threshold                = "enable"
    ssm                          = "disable"
  }
}
`)
}
