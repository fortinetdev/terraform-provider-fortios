// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"log"
	"testing"
)

func TestAccFortiOSSystemIpv6Tunnel_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemIpv6Tunnel_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemIpv6TunnelConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemIpv6TunnelExists("fortios_system_ipv6tunnel.trname"),
					resource.TestCheckResourceAttr("fortios_system_ipv6tunnel.trname", "destination", "2001:db8:85a3::8a2e:370:7324"),
					resource.TestCheckResourceAttr("fortios_system_ipv6tunnel.trname", "interface", "port3"),
					resource.TestCheckResourceAttr("fortios_system_ipv6tunnel.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_system_ipv6tunnel.trname", "source", "2001:db8:85a3::8a2e:370:7334"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemIpv6TunnelExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemIpv6Tunnel: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemIpv6Tunnel is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemIpv6Tunnel(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemIpv6Tunnel: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemIpv6Tunnel: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemIpv6TunnelDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_ipv6tunnel" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemIpv6Tunnel(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemIpv6Tunnel %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemIpv6TunnelConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_ipv6tunnel" "trname" {
  destination = "2001:db8:85a3::8a2e:370:7324"
  interface   = "port3"
  name        = "%[1]s"
  source      = "2001:db8:85a3::8a2e:370:7334"
}
`, name)
}
