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

func TestAccFortiOSSystemManagementTunnel_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemManagementTunnel_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemManagementTunnelConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemManagementTunnelExists("fortios_system_managementtunnel.trname"),
					resource.TestCheckResourceAttr("fortios_system_managementtunnel.trname", "allow_collect_statistics", "enable"),
					resource.TestCheckResourceAttr("fortios_system_managementtunnel.trname", "allow_config_restore", "enable"),
					resource.TestCheckResourceAttr("fortios_system_managementtunnel.trname", "allow_push_configuration", "enable"),
					resource.TestCheckResourceAttr("fortios_system_managementtunnel.trname", "allow_push_firmware", "enable"),
					resource.TestCheckResourceAttr("fortios_system_managementtunnel.trname", "authorized_manager_only", "enable"),
					resource.TestCheckResourceAttr("fortios_system_managementtunnel.trname", "status", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemManagementTunnelExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemManagementTunnel: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemManagementTunnel is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemManagementTunnel(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemManagementTunnel: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemManagementTunnel: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemManagementTunnelDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_managementtunnel" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemManagementTunnel(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemManagementTunnel %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemManagementTunnelConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_managementtunnel" "trname" {
  allow_collect_statistics = "enable"
  allow_config_restore     = "enable"
  allow_push_configuration = "enable"
  allow_push_firmware      = "enable"
  authorized_manager_only  = "enable"
  status                   = "enable"
}
`)
}
