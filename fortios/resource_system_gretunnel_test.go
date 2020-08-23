
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

func TestAccFortiOSSystemGreTunnel_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSystemGreTunnel_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSystemGreTunnelConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSystemGreTunnelExists("fortios_system_gretunnel.trname"),
                    resource.TestCheckResourceAttr("fortios_system_gretunnel.trname", "checksum_reception", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_gretunnel.trname", "checksum_transmission", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_gretunnel.trname", "dscp_copying", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_gretunnel.trname", "interface", "port3"),
                    resource.TestCheckResourceAttr("fortios_system_gretunnel.trname", "ip_version", "4"),
                    resource.TestCheckResourceAttr("fortios_system_gretunnel.trname", "keepalive_failtimes", "10"),
                    resource.TestCheckResourceAttr("fortios_system_gretunnel.trname", "keepalive_interval", "0"),
                    resource.TestCheckResourceAttr("fortios_system_gretunnel.trname", "key_inbound", "0"),
                    resource.TestCheckResourceAttr("fortios_system_gretunnel.trname", "key_outbound", "0"),
                    resource.TestCheckResourceAttr("fortios_system_gretunnel.trname", "local_gw", "3.3.3.3"),
                    resource.TestCheckResourceAttr("fortios_system_gretunnel.trname", "local_gw6", "::"),
                    resource.TestCheckResourceAttr("fortios_system_gretunnel.trname", "name", rname),
                    resource.TestCheckResourceAttr("fortios_system_gretunnel.trname", "remote_gw", "1.1.1.1"),
                    resource.TestCheckResourceAttr("fortios_system_gretunnel.trname", "remote_gw6", "::"),
                    resource.TestCheckResourceAttr("fortios_system_gretunnel.trname", "sequence_number_reception", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_gretunnel.trname", "sequence_number_transmission", "enable"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSystemGreTunnelExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemGreTunnel: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemGreTunnel is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemGreTunnel(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemGreTunnel: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemGreTunnel: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemGreTunnelDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_gretunnel" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemGreTunnel(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemGreTunnel %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemGreTunnelConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_gretunnel" "trname" {
  checksum_reception           = "disable"
  checksum_transmission        = "disable"
  dscp_copying                 = "disable"
  interface                    = "port3"
  ip_version                   = "4"
  keepalive_failtimes          = 10
  keepalive_interval           = 0
  key_inbound                  = 0
  key_outbound                 = 0
  local_gw                     = "3.3.3.3"
  local_gw6                    = "::"
  name                         = "%[1]s"
  remote_gw                    = "1.1.1.1"
  remote_gw6                   = "::"
  sequence_number_reception    = "disable"
  sequence_number_transmission = "enable"
}

`, name)
}
