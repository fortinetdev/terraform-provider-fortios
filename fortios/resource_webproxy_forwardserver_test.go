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

func TestAccFortiOSWebProxyForwardServer_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWebProxyForwardServer_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWebProxyForwardServerConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWebProxyForwardServerExists("fortios_webproxy_forwardserver.trname"),
					resource.TestCheckResourceAttr("fortios_webproxy_forwardserver.trname", "addr_type", "fqdn"),
					resource.TestCheckResourceAttr("fortios_webproxy_forwardserver.trname", "healthcheck", "disable"),
					resource.TestCheckResourceAttr("fortios_webproxy_forwardserver.trname", "ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_webproxy_forwardserver.trname", "monitor", "http://www.google.com"),
					resource.TestCheckResourceAttr("fortios_webproxy_forwardserver.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_webproxy_forwardserver.trname", "port", "3128"),
					resource.TestCheckResourceAttr("fortios_webproxy_forwardserver.trname", "server_down_option", "block"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWebProxyForwardServerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WebProxyForwardServer: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebProxyForwardServer is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWebProxyForwardServer(i)

		if err != nil {
			return fmt.Errorf("Error reading WebProxyForwardServer: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WebProxyForwardServer: %s", n)
		}

		return nil
	}
}

func testAccCheckWebProxyForwardServerDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_webproxy_forwardserver" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWebProxyForwardServer(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WebProxyForwardServer %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWebProxyForwardServerConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_webproxy_forwardserver" "trname" {
  addr_type          = "fqdn"
  healthcheck        = "disable"
  ip                 = "0.0.0.0"
  monitor            = "http://www.google.com"
  name               = "%[1]s"
  port               = 3128
  server_down_option = "block"
}
`, name)
}
