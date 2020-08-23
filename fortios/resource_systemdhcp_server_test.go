
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

func TestAccFortiOSSystemDhcpServer_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSystemDhcpServer_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSystemDhcpServerConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSystemDhcpServerExists("fortios_systemdhcp_server.trname"),
                    resource.TestCheckResourceAttr("fortios_systemdhcp_server.trname", "dns_service", "default"),
                    resource.TestCheckResourceAttr("fortios_systemdhcp_server.trname", "fosid", "1"),
                    resource.TestCheckResourceAttr("fortios_systemdhcp_server.trname", "interface", "port2"),
                    resource.TestCheckResourceAttr("fortios_systemdhcp_server.trname", "netmask", "255.255.255.0"),
                    resource.TestCheckResourceAttr("fortios_systemdhcp_server.trname", "status", "disable"),
                    resource.TestCheckResourceAttr("fortios_systemdhcp_server.trname", "ntp_server1", "192.168.52.22"),
                    resource.TestCheckResourceAttr("fortios_systemdhcp_server.trname", "timezone", "00"),
                    resource.TestCheckResourceAttr("fortios_systemdhcp_server.trname", "ip_range.0.end_ip", "1.1.1.22"),
                    resource.TestCheckResourceAttr("fortios_systemdhcp_server.trname", "ip_range.0.id", "1"),
                    resource.TestCheckResourceAttr("fortios_systemdhcp_server.trname", "ip_range.0.start_ip", "1.1.1.1"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSystemDhcpServerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemDhcpServer: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemDhcpServer is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemDhcpServer(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemDhcpServer: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemDhcpServer: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemDhcpServerDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_systemdhcp_server" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemDhcpServer(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemDhcpServer %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemDhcpServerConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_systemdhcp_server" "trname" {
  dns_service = "default"
  fosid       = 1
  interface   = "port2"
  netmask     = "255.255.255.0"
  status      = "disable"
  ntp_server1 = "192.168.52.22"
  timezone    = "00"

  ip_range {
    end_ip   = "1.1.1.22"
    id       = 1
    start_ip = "1.1.1.1"
  }
}

`)
}
