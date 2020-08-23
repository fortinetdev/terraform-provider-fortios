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

func TestAccFortiOSSystemAutoInstall_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemAutoInstall_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemAutoInstallConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemAutoInstallExists("fortios_system_autoinstall.trname"),
					resource.TestCheckResourceAttr("fortios_system_autoinstall.trname", "auto_install_config", "enable"),
					resource.TestCheckResourceAttr("fortios_system_autoinstall.trname", "auto_install_image", "enable"),
					resource.TestCheckResourceAttr("fortios_system_autoinstall.trname", "default_config_file", "fgt_system.conf"),
					resource.TestCheckResourceAttr("fortios_system_autoinstall.trname", "default_image_file", "image.out"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemAutoInstallExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemAutoInstall: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemAutoInstall is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemAutoInstall(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemAutoInstall: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemAutoInstall: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemAutoInstallDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_autoinstall" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemAutoInstall(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemAutoInstall %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemAutoInstallConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_autoinstall" "trname" {
  auto_install_config = "enable"
  auto_install_image  = "enable"
  default_config_file = "fgt_system.conf"
  default_image_file  = "image.out"
}
`)
}
