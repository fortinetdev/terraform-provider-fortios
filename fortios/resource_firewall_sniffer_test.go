
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

func TestAccFortiOSFirewallSniffer_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSFirewallSniffer_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSFirewallSnifferConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSFirewallSnifferExists("fortios_firewall_sniffer.trname"),
                    resource.TestCheckResourceAttr("fortios_firewall_sniffer.trname", "application_list_status", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_sniffer.trname", "av_profile_status", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_sniffer.trname", "dlp_sensor_status", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_sniffer.trname", "dsri", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_sniffer.trname", "fosid", "1"),
                    resource.TestCheckResourceAttr("fortios_firewall_sniffer.trname", "interface", "port4"),
                    resource.TestCheckResourceAttr("fortios_firewall_sniffer.trname", "ips_dos_status", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_sniffer.trname", "ips_sensor_status", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_sniffer.trname", "ipv6", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_sniffer.trname", "logtraffic", "utm"),
                    resource.TestCheckResourceAttr("fortios_firewall_sniffer.trname", "max_packet_count", "4000"),
                    resource.TestCheckResourceAttr("fortios_firewall_sniffer.trname", "non_ip", "enable"),
                    resource.TestCheckResourceAttr("fortios_firewall_sniffer.trname", "scan_botnet_connections", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_sniffer.trname", "spamfilter_profile_status", "disable"),
                    resource.TestCheckResourceAttr("fortios_firewall_sniffer.trname", "status", "enable"),
                    resource.TestCheckResourceAttr("fortios_firewall_sniffer.trname", "webfilter_profile_status", "disable"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSFirewallSnifferExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallSniffer: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallSniffer is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallSniffer(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallSniffer: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallSniffer: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallSnifferDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_sniffer" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallSniffer(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallSniffer %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallSnifferConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_sniffer" "trname" {
  application_list_status   = "disable"
  av_profile_status         = "disable"
  dlp_sensor_status         = "disable"
  dsri                      = "disable"
  fosid                     = 1
  interface                 = "port4"
  ips_dos_status            = "disable"
  ips_sensor_status         = "disable"
  ipv6                      = "disable"
  logtraffic                = "utm"
  max_packet_count          = 4000
  non_ip                    = "enable"
  scan_botnet_connections   = "disable"
  spamfilter_profile_status = "disable"
  status                    = "enable"
  webfilter_profile_status  = "disable"
}
`)
}
