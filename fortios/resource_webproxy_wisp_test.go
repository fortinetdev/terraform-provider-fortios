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

func TestAccFortiOSWebProxyWisp_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWebProxyWisp_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWebProxyWispConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWebProxyWispExists("fortios_webproxy_wisp.trname"),
					resource.TestCheckResourceAttr("fortios_webproxy_wisp.trname", "max_connections", "64"),
					resource.TestCheckResourceAttr("fortios_webproxy_wisp.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_webproxy_wisp.trname", "outgoing_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_webproxy_wisp.trname", "server_ip", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_webproxy_wisp.trname", "server_port", "15868"),
					resource.TestCheckResourceAttr("fortios_webproxy_wisp.trname", "timeout", "5"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWebProxyWispExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WebProxyWisp: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebProxyWisp is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWebProxyWisp(i)

		if err != nil {
			return fmt.Errorf("Error reading WebProxyWisp: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WebProxyWisp: %s", n)
		}

		return nil
	}
}

func testAccCheckWebProxyWispDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_webproxy_wisp" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWebProxyWisp(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WebProxyWisp %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWebProxyWispConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_webproxy_wisp" "trname" {
  max_connections = 64
  name            = "%[1]s"
  outgoing_ip     = "0.0.0.0"
  server_ip       = "1.1.1.1"
  server_port     = 15868
  timeout         = 5
}
`, name)
}
