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

func TestAccFortiOSSystemPppoeInterface_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemPppoeInterface_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemPppoeInterfaceConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemPppoeInterfaceExists("fortios_system_pppoeinterface.trname"),
					resource.TestCheckResourceAttr("fortios_system_pppoeinterface.trname", "auth_type", "auto"),
					resource.TestCheckResourceAttr("fortios_system_pppoeinterface.trname", "device", "port3"),
					resource.TestCheckResourceAttr("fortios_system_pppoeinterface.trname", "dial_on_demand", "disable"),
					resource.TestCheckResourceAttr("fortios_system_pppoeinterface.trname", "disc_retry_timeout", "1"),
					resource.TestCheckResourceAttr("fortios_system_pppoeinterface.trname", "idle_timeout", "0"),
					resource.TestCheckResourceAttr("fortios_system_pppoeinterface.trname", "ipunnumbered", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_system_pppoeinterface.trname", "ipv6", "disable"),
					resource.TestCheckResourceAttr("fortios_system_pppoeinterface.trname", "lcp_echo_interval", "5"),
					resource.TestCheckResourceAttr("fortios_system_pppoeinterface.trname", "lcp_max_echo_fails", "3"),
					resource.TestCheckResourceAttr("fortios_system_pppoeinterface.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_system_pppoeinterface.trname", "padt_retry_timeout", "1"),
					resource.TestCheckResourceAttr("fortios_system_pppoeinterface.trname", "pppoe_unnumbered_negotiate", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemPppoeInterfaceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemPppoeInterface: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemPppoeInterface is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemPppoeInterface(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemPppoeInterface: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemPppoeInterface: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemPppoeInterfaceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_pppoeinterface" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemPppoeInterface(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemPppoeInterface %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemPppoeInterfaceConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_pppoeinterface" "trname" {
  auth_type                  = "auto"
  device                     = "port3"
  dial_on_demand             = "disable"
  disc_retry_timeout         = 1
  idle_timeout               = 0
  ipunnumbered               = "0.0.0.0"
  ipv6                       = "disable"
  lcp_echo_interval          = 5
  lcp_max_echo_fails         = 3
  name                       = "%[1]s"
  padt_retry_timeout         = 1
  pppoe_unnumbered_negotiate = "enable"
}
`, name)
}
