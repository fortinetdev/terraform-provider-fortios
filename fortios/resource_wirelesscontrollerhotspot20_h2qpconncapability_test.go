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

func TestAccFortiOSWirelessControllerHotspot20H2QpConnCapability_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWirelessControllerHotspot20H2QpConnCapability_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWirelessControllerHotspot20H2QpConnCapabilityConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWirelessControllerHotspot20H2QpConnCapabilityExists("fortios_wirelesscontrollerhotspot20_h2qpconncapability.trname"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpconncapability.trname", "esp_port", "unknown"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpconncapability.trname", "ftp_port", "unknown"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpconncapability.trname", "http_port", "unknown"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpconncapability.trname", "icmp_port", "closed"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpconncapability.trname", "ikev2_port", "unknown"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpconncapability.trname", "ikev2_xx_port", "unknown"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpconncapability.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpconncapability.trname", "pptp_vpn_port", "unknown"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpconncapability.trname", "ssh_port", "unknown"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpconncapability.trname", "tls_port", "unknown"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpconncapability.trname", "voip_tcp_port", "unknown"),
					resource.TestCheckResourceAttr("fortios_wirelesscontrollerhotspot20_h2qpconncapability.trname", "voip_udp_port", "unknown"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWirelessControllerHotspot20H2QpConnCapabilityExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WirelessControllerHotspot20H2QpConnCapability: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WirelessControllerHotspot20H2QpConnCapability is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWirelessControllerHotspot20H2QpConnCapability(i)

		if err != nil {
			return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpConnCapability: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpConnCapability: %s", n)
		}

		return nil
	}
}

func testAccCheckWirelessControllerHotspot20H2QpConnCapabilityDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_wirelesscontrollerhotspot20_h2qpconncapability" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWirelessControllerHotspot20H2QpConnCapability(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WirelessControllerHotspot20H2QpConnCapability %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWirelessControllerHotspot20H2QpConnCapabilityConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_wirelesscontrollerhotspot20_h2qpconncapability" "trname" {
  esp_port      = "unknown"
  ftp_port      = "unknown"
  http_port     = "unknown"
  icmp_port     = "closed"
  ikev2_port    = "unknown"
  ikev2_xx_port = "unknown"
  name          = "%[1]s"
  pptp_vpn_port = "unknown"
  ssh_port      = "unknown"
  tls_port      = "unknown"
  voip_tcp_port = "unknown"
  voip_udp_port = "unknown"
}
`, name)
}
