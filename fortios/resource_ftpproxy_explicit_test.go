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

func TestAccFortiOSFtpProxyExplicit_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSFtpProxyExplicit_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFtpProxyExplicitConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFtpProxyExplicitExists("fortios_ftpproxy_explicit.trname"),
					resource.TestCheckResourceAttr("fortios_ftpproxy_explicit.trname", "incoming_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_ftpproxy_explicit.trname", "sec_default_action", "deny"),
					resource.TestCheckResourceAttr("fortios_ftpproxy_explicit.trname", "status", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSFtpProxyExplicitExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FtpProxyExplicit: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FtpProxyExplicit is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFtpProxyExplicit(i)

		if err != nil {
			return fmt.Errorf("Error reading FtpProxyExplicit: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FtpProxyExplicit: %s", n)
		}

		return nil
	}
}

func testAccCheckFtpProxyExplicitDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_ftpproxy_explicit" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFtpProxyExplicit(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FtpProxyExplicit %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFtpProxyExplicitConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_ftpproxy_explicit" "trname" {
  incoming_ip        = "0.0.0.0"
  sec_default_action = "deny"
  status             = "disable"
}
`)
}
