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

func TestAccFortiOSWebfilterFortiguard_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWebfilterFortiguard_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWebfilterFortiguardConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWebfilterFortiguardExists("fortios_webfilter_fortiguard.trname"),
					resource.TestCheckResourceAttr("fortios_webfilter_fortiguard.trname", "cache_mem_percent", "2"),
					resource.TestCheckResourceAttr("fortios_webfilter_fortiguard.trname", "cache_mode", "ttl"),
					resource.TestCheckResourceAttr("fortios_webfilter_fortiguard.trname", "cache_prefix_match", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_fortiguard.trname", "close_ports", "disable"),
					resource.TestCheckResourceAttr("fortios_webfilter_fortiguard.trname", "ovrd_auth_https", "enable"),
					resource.TestCheckResourceAttr("fortios_webfilter_fortiguard.trname", "ovrd_auth_port", "8008"),
					resource.TestCheckResourceAttr("fortios_webfilter_fortiguard.trname", "ovrd_auth_port_http", "8008"),
					resource.TestCheckResourceAttr("fortios_webfilter_fortiguard.trname", "ovrd_auth_port_https", "8010"),
					resource.TestCheckResourceAttr("fortios_webfilter_fortiguard.trname", "ovrd_auth_port_https_flow", "8015"),
					resource.TestCheckResourceAttr("fortios_webfilter_fortiguard.trname", "ovrd_auth_port_warning", "8020"),
					resource.TestCheckResourceAttr("fortios_webfilter_fortiguard.trname", "warn_auth_https", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWebfilterFortiguardExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WebfilterFortiguard: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebfilterFortiguard is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWebfilterFortiguard(i)

		if err != nil {
			return fmt.Errorf("Error reading WebfilterFortiguard: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WebfilterFortiguard: %s", n)
		}

		return nil
	}
}

func testAccCheckWebfilterFortiguardDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_webfilter_fortiguard" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWebfilterFortiguard(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WebfilterFortiguard %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWebfilterFortiguardConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_webfilter_fortiguard" "trname" {
  cache_mem_percent         = 2
  cache_mode                = "ttl"
  cache_prefix_match        = "enable"
  close_ports               = "disable"
  ovrd_auth_https           = "enable"
  ovrd_auth_port            = 8008
  ovrd_auth_port_http       = 8008
  ovrd_auth_port_https      = 8010
  ovrd_auth_port_https_flow = 8015
  ovrd_auth_port_warning    = 8020
  warn_auth_https           = "enable"
}
`)
}
