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

func TestAccFortiOSSystemIpipTunnel_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemIpipTunnel_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemIpipTunnelConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemIpipTunnelExists("fortios_system_ipiptunnel.trname"),
					resource.TestCheckResourceAttr("fortios_system_ipiptunnel.trname", "interface", "port3"),
					resource.TestCheckResourceAttr("fortios_system_ipiptunnel.trname", "local_gw", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_system_ipiptunnel.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_system_ipiptunnel.trname", "remote_gw", "2.2.2.2"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemIpipTunnelExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemIpipTunnel: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemIpipTunnel is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemIpipTunnel(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemIpipTunnel: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemIpipTunnel: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemIpipTunnelDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_ipiptunnel" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemIpipTunnel(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemIpipTunnel %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemIpipTunnelConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_ipiptunnel" "trname" {
  interface = "port3"
  local_gw  = "1.1.1.1"
  name      = "%[1]s"
  remote_gw = "2.2.2.2"
}
`, name)
}
