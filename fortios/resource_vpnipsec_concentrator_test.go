// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSVpnIpsecConcentrator_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSVpnIpsecConcentrator_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSVpnIpsecConcentratorConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSVpnIpsecConcentratorExists("fortios_vpnipsec_concentrator.trname"),
					resource.TestCheckResourceAttr("fortios_vpnipsec_concentrator.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_vpnipsec_concentrator.trname", "src_check", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSVpnIpsecConcentratorExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found VpnIpsecConcentrator: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VpnIpsecConcentrator is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadVpnIpsecConcentrator(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading VpnIpsecConcentrator: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating VpnIpsecConcentrator: %s", n)
		}

		return nil
	}
}

func testAccCheckVpnIpsecConcentratorDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_vpnipsec_concentrator" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadVpnIpsecConcentrator(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error VpnIpsecConcentrator %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSVpnIpsecConcentratorConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_vpnipsec_concentrator" "trname" {
  name      = "%[1]s"
  src_check = "disable"
}
`, name)
}
