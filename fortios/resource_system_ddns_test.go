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

func TestAccFortiOSSystemDdns_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemDdns_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemDdnsConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemDdnsExists("fortios_system_ddns.trname"),
					resource.TestCheckResourceAttr("fortios_system_ddns.trname", "bound_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_system_ddns.trname", "clear_text", "disable"),
					resource.TestCheckResourceAttr("fortios_system_ddns.trname", "ddns_auth", "disable"),
					resource.TestCheckResourceAttr("fortios_system_ddns.trname", "ddns_domain", "www.s.com"),
					resource.TestCheckResourceAttr("fortios_system_ddns.trname", "ddns_password", "ewewcd"),
					resource.TestCheckResourceAttr("fortios_system_ddns.trname", "ddns_server", "tzo.com"),
					resource.TestCheckResourceAttr("fortios_system_ddns.trname", "ddns_server_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_system_ddns.trname", "ddns_ttl", "300"),
					resource.TestCheckResourceAttr("fortios_system_ddns.trname", "ddns_username", "sie2ae"),
					resource.TestCheckResourceAttr("fortios_system_ddns.trname", "ddnsid", "1"),
					resource.TestCheckResourceAttr("fortios_system_ddns.trname", "ssl_certificate", "Fortinet_Factory"),
					resource.TestCheckResourceAttr("fortios_system_ddns.trname", "update_interval", "300"),
					resource.TestCheckResourceAttr("fortios_system_ddns.trname", "use_public_ip", "disable"),
					resource.TestCheckResourceAttr("fortios_system_ddns.trname", "monitor_interface.0.interface_name", "port2"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemDdnsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemDdns: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemDdns is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemDdns(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemDdns: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemDdns: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemDdnsDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_ddns" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemDdns(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemDdns %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemDdnsConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_ddns" "trname" {
  bound_ip        = "0.0.0.0"
  clear_text      = "disable"
  ddns_auth       = "disable"
  ddns_domain     = "www.s.com"
  ddns_password   = "ewewcd"
  ddns_server     = "tzo.com"
  ddns_server_ip  = "0.0.0.0"
  ddns_ttl        = 300
  ddns_username   = "sie2ae"
  ddnsid          = 1
  ssl_certificate = "Fortinet_Factory"
  update_interval = 300
  use_public_ip   = "disable"
  monitor_interface {
    interface_name = "port2"
  }
}
`)
}
