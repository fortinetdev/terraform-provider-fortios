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

func TestAccFortiOSSystemDnsServer_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemDnsServer_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemDnsServerConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemDnsServerExists("fortios_system_dnsserver.trname"),
					resource.TestCheckResourceAttr("fortios_system_dnsserver.trname", "dnsfilter_profile", "default"),
					resource.TestCheckResourceAttr("fortios_system_dnsserver.trname", "mode", "forward-only"),
					resource.TestCheckResourceAttr("fortios_system_dnsserver.trname", "name", "port3"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemDnsServerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemDnsServer: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemDnsServer is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemDnsServer(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemDnsServer: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemDnsServer: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemDnsServerDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_dnsserver" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemDnsServer(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemDnsServer %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemDnsServerConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_dnsserver" "trname" {
  dnsfilter_profile = "default"
  mode              = "forward-only"
  name              = "port3"
}
`)
}
