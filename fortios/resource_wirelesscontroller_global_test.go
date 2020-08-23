
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

func TestAccFortiOSWirelessControllerGlobal_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSWirelessControllerGlobal_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSWirelessControllerGlobalConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSWirelessControllerGlobalExists("fortios_wirelesscontroller_global.trname"),
                    resource.TestCheckResourceAttr("fortios_wirelesscontroller_global.trname", "ap_log_server", "disable"),
                    resource.TestCheckResourceAttr("fortios_wirelesscontroller_global.trname", "ap_log_server_ip", "0.0.0.0"),
                    resource.TestCheckResourceAttr("fortios_wirelesscontroller_global.trname", "ap_log_server_port", "0"),
                    resource.TestCheckResourceAttr("fortios_wirelesscontroller_global.trname", "control_message_offload", "ebp-frame aeroscout-tag ap-list sta-list sta-cap-list stats aeroscout-mu"),
                    resource.TestCheckResourceAttr("fortios_wirelesscontroller_global.trname", "data_ethernet_ii", "disable"),
                    resource.TestCheckResourceAttr("fortios_wirelesscontroller_global.trname", "discovery_mc_addr", "224.0.1.140"),
                    resource.TestCheckResourceAttr("fortios_wirelesscontroller_global.trname", "fiapp_eth_type", "5252"),
                    resource.TestCheckResourceAttr("fortios_wirelesscontroller_global.trname", "image_download", "enable"),
                    resource.TestCheckResourceAttr("fortios_wirelesscontroller_global.trname", "ipsec_base_ip", "169.254.0.1"),
                    resource.TestCheckResourceAttr("fortios_wirelesscontroller_global.trname", "link_aggregation", "disable"),
                    resource.TestCheckResourceAttr("fortios_wirelesscontroller_global.trname", "max_clients", "0"),
                    resource.TestCheckResourceAttr("fortios_wirelesscontroller_global.trname", "max_retransmit", "3"),
                    resource.TestCheckResourceAttr("fortios_wirelesscontroller_global.trname", "mesh_eth_type", "8755"),
                    resource.TestCheckResourceAttr("fortios_wirelesscontroller_global.trname", "rogue_scan_mac_adjacency", "7"),
                    resource.TestCheckResourceAttr("fortios_wirelesscontroller_global.trname", "wtp_share", "disable"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSWirelessControllerGlobalExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WirelessControllerGlobal: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WirelessControllerGlobal is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWirelessControllerGlobal(i)

		if err != nil {
			return fmt.Errorf("Error reading WirelessControllerGlobal: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WirelessControllerGlobal: %s", n)
		}

		return nil
	}
}

func testAccCheckWirelessControllerGlobalDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_wirelesscontroller_global" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWirelessControllerGlobal(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WirelessControllerGlobal %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWirelessControllerGlobalConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_wirelesscontroller_global" "trname" {
  ap_log_server            = "disable"
  ap_log_server_ip         = "0.0.0.0"
  ap_log_server_port       = 0
  control_message_offload  = "ebp-frame aeroscout-tag ap-list sta-list sta-cap-list stats aeroscout-mu"
  data_ethernet_ii         = "disable"
  discovery_mc_addr        = "224.0.1.140"
  fiapp_eth_type           = 5252
  image_download           = "enable"
  ipsec_base_ip            = "169.254.0.1"
  link_aggregation         = "disable"
  max_clients              = 0
  max_retransmit           = 3
  mesh_eth_type            = 8755
  rogue_scan_mac_adjacency = 7
  wtp_share                = "disable"
}
`)
}
