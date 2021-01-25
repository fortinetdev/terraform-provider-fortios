// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSFirewallSslServer_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFirewallSslServer_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallSslServerConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallSslServerExists("fortios_firewall_sslserver.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_sslserver.trname", "add_header_x_forwarded_proto", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_sslserver.trname", "ip", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_firewall_sslserver.trname", "mapped_port", "2234"),
					resource.TestCheckResourceAttr("fortios_firewall_sslserver.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_sslserver.trname", "port", "32321"),
					resource.TestCheckResourceAttr("fortios_firewall_sslserver.trname", "ssl_algorithm", "high"),
					resource.TestCheckResourceAttr("fortios_firewall_sslserver.trname", "ssl_cert", "Fortinet_CA_SSL"),
					resource.TestCheckResourceAttr("fortios_firewall_sslserver.trname", "ssl_client_renegotiation", "allow"),
					resource.TestCheckResourceAttr("fortios_firewall_sslserver.trname", "ssl_dh_bits", "2048"),
					resource.TestCheckResourceAttr("fortios_firewall_sslserver.trname", "ssl_max_version", "tls-1.2"),
					resource.TestCheckResourceAttr("fortios_firewall_sslserver.trname", "ssl_min_version", "tls-1.1"),
					resource.TestCheckResourceAttr("fortios_firewall_sslserver.trname", "ssl_mode", "half"),
					resource.TestCheckResourceAttr("fortios_firewall_sslserver.trname", "ssl_send_empty_frags", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_sslserver.trname", "url_rewrite", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallSslServerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallSslServer: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallSslServer is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallSslServer(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallSslServer: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallSslServer: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallSslServerDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_sslserver" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallSslServer(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallSslServer %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallSslServerConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_sslserver" "trname" {
  add_header_x_forwarded_proto = "enable"
  ip                           = "1.1.1.1"
  mapped_port                  = 2234
  name                         = "%[1]s"
  port                         = 32321
  ssl_algorithm                = "high"
  ssl_cert                     = "Fortinet_CA_SSL"
  ssl_client_renegotiation     = "allow"
  ssl_dh_bits                  = "2048"
  ssl_max_version              = "tls-1.2"
  ssl_min_version              = "tls-1.1"
  ssl_mode                     = "half"
  ssl_send_empty_frags         = "enable"
  url_rewrite                  = "disable"
}
`, name)
}
