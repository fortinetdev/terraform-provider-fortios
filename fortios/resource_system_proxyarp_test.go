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

func TestAccFortiOSSystemProxyArp_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemProxyArp_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemProxyArpConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemProxyArpExists("fortios_system_proxyarp.trname"),
					resource.TestCheckResourceAttr("fortios_system_proxyarp.trname", "end_ip", "1.1.1.3"),
					resource.TestCheckResourceAttr("fortios_system_proxyarp.trname", "fosid", "1"),
					resource.TestCheckResourceAttr("fortios_system_proxyarp.trname", "interface", "port4"),
					resource.TestCheckResourceAttr("fortios_system_proxyarp.trname", "ip", "1.1.1.1"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemProxyArpExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemProxyArp: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemProxyArp is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemProxyArp(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemProxyArp: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemProxyArp: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemProxyArpDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_proxyarp" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemProxyArp(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemProxyArp %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemProxyArpConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_proxyarp" "trname" {
  end_ip    = "1.1.1.3"
  fosid     = 1
  interface = "port4"
  ip        = "1.1.1.1"
}


`)
}
