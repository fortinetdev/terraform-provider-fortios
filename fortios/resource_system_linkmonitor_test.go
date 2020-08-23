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

func TestAccFortiOSSystemLinkMonitor_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemLinkMonitor_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemLinkMonitorConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemLinkMonitorExists("fortios_system_linkmonitor.trname"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "addr_mode", "ipv4"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "failtime", "5"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "gateway_ip", "2.2.2.2"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "gateway_ip6", "::"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "ha_priority", "1"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "http_agent", "Chrome/ Safari/"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "http_get", "/"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "interval", "1"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "packet_size", "64"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "port", "80"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "protocol", "ping"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "recoverytime", "5"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "security_mode", "none"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "source_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "source_ip6", "::"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "srcintf", "port4"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "update_cascade_interface", "enable"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "update_static_route", "enable"),
					resource.TestCheckResourceAttr("fortios_system_linkmonitor.trname", "server.0.address", "3.3.3.3"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemLinkMonitorExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemLinkMonitor: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemLinkMonitor is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemLinkMonitor(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemLinkMonitor: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemLinkMonitor: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemLinkMonitorDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_linkmonitor" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemLinkMonitor(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemLinkMonitor %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemLinkMonitorConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_linkmonitor" "trname" {
  addr_mode                = "ipv4"
  failtime                 = 5
  gateway_ip               = "2.2.2.2"
  gateway_ip6              = "::"
  ha_priority              = 1
  http_agent               = "Chrome/ Safari/"
  http_get                 = "/"
  interval                 = 1
  name                     = "%[1]s"
  packet_size              = 64
  port                     = 80
  protocol                 = "ping"
  recoverytime             = 5
  security_mode            = "none"
  source_ip                = "0.0.0.0"
  source_ip6               = "::"
  srcintf                  = "port4"
  status                   = "enable"
  update_cascade_interface = "enable"
  update_static_route      = "enable"
  server {
    address = "3.3.3.3"
  }
}
`, name)
}
