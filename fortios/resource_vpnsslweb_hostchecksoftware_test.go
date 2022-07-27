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

func TestAccFortiOSVpnSslWebHostCheckSoftware_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSVpnSslWebHostCheckSoftware_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSVpnSslWebHostCheckSoftwareConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSVpnSslWebHostCheckSoftwareExists("fortios_vpnsslweb_hostchecksoftware.trname"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_hostchecksoftware.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_hostchecksoftware.trname", "os_type", "windows"),
					resource.TestCheckResourceAttr("fortios_vpnsslweb_hostchecksoftware.trname", "type", "fw"),
				),
			},
		},
	})
}

func testAccCheckFortiOSVpnSslWebHostCheckSoftwareExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found VpnSslWebHostCheckSoftware: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VpnSslWebHostCheckSoftware is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadVpnSslWebHostCheckSoftware(i)

		if err != nil {
			return fmt.Errorf("Error reading VpnSslWebHostCheckSoftware: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating VpnSslWebHostCheckSoftware: %s", n)
		}

		return nil
	}
}

func testAccCheckVpnSslWebHostCheckSoftwareDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_vpnsslweb_hostchecksoftware" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadVpnSslWebHostCheckSoftware(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error VpnSslWebHostCheckSoftware %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSVpnSslWebHostCheckSoftwareConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_vpnsslweb_hostchecksoftware" "trname" {
  name    = "%[1]s"
  os_type = "windows"
  type    = "fw"
}
`, name)
}
