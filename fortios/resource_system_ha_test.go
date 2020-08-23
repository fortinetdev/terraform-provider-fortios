
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

func TestAccFortiOSSystemHa_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSystemHa_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSystemHaConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSystemHaExists("fortios_system_ha.trname"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "cpu_threshold", "5 0 0"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "encryption", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "ftp_proxy_threshold", "5 0 0"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "gratuitous_arps", "enable"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "group_id", "0"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "ha_direct", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "ha_eth_type", "8890"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "ha_mgmt_status", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "ha_uptime_diff_margin", "300"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "hb_interval", "2"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "hb_lost_threshold", "20"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "hc_eth_type", "8891"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "hello_holddown", "20"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "http_proxy_threshold", "5 0 0"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "imap_proxy_threshold", "5 0 0"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "inter_cluster_session_sync", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "l2ep_eth_type", "8893"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "link_failed_signal", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "load_balance_all", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "memory_compatible_mode", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "memory_threshold", "5 0 0"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "mode", "standalone"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "multicast_ttl", "600"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "nntp_proxy_threshold", "5 0 0"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "override", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "override_wait_time", "0"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "weight", "40 "),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "secondary_vcluster.0.override", "enable"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "secondary_vcluster.0.override_wait_time", "0"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "secondary_vcluster.0.pingserver_failover_threshold", "0"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "secondary_vcluster.0.pingserver_slave_force_reset", "enable"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "secondary_vcluster.0.priority", "128"),
                    resource.TestCheckResourceAttr("fortios_system_ha.trname", "secondary_vcluster.0.vcluster_id", "1"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSystemHaExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemHa: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemHa is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemHa(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemHa: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemHa: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemHaDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_ha" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemHa(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemHa %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemHaConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_ha" "trname" {
  cpu_threshold              = "5 0 0"
  encryption                 = "disable"
  ftp_proxy_threshold        = "5 0 0"
  gratuitous_arps            = "enable"
  group_id                   = 0
  ha_direct                  = "disable"
  ha_eth_type                = "8890"
  ha_mgmt_status             = "disable"
  ha_uptime_diff_margin      = 300
  hb_interval                = 2
  hb_lost_threshold          = 20
  hc_eth_type                = "8891"
  hello_holddown             = 20
  http_proxy_threshold       = "5 0 0"
  imap_proxy_threshold       = "5 0 0"
  inter_cluster_session_sync = "disable"
  l2ep_eth_type              = "8893"
  link_failed_signal         = "disable"
  load_balance_all           = "disable"
  memory_compatible_mode     = "disable"
  memory_threshold           = "5 0 0"
  mode                       = "standalone"
  multicast_ttl              = 600
  nntp_proxy_threshold       = "5 0 0"
  override                   = "disable"
  override_wait_time         = 0
  weight                     = "40 "
  secondary_vcluster {
    override                      = "enable"
    override_wait_time            = 0
    pingserver_failover_threshold = 0
    pingserver_slave_force_reset  = "enable"
    priority                      = 128
    vcluster_id                   = 1
  }
}
`)
}
