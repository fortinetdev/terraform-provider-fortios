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

func TestAccFortiOSFirewallAddrgrp_basic(t *testing.T) {
	rname := acctest.RandString(8)
	var0 := "var0" + rname
	log.Printf(var0)
	log.Printf("TestAccFortiOSFirewallAddrgrp_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallAddrgrpConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallAddrgrpExists("fortios_firewall_addrgrp.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_addrgrp.trname", "allow_routing", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_addrgrp.trname", "color", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_addrgrp.trname", "exclude", "disable"),
					resource.TestCheckResourceAttr("fortios_firewall_addrgrp.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_addrgrp.trname", "visibility", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_addrgrp.trname", "member.0.name", var0),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallAddrgrpExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallAddrgrp: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallAddrgrp is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallAddrgrp(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallAddrgrp: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallAddrgrp: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallAddrgrpDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_addrgrp" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallAddrgrp(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallAddrgrp %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallAddrgrpConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_address" "trname1" {
  allow_routing = "disable"
  cache_ttl     = 0
  color         = 0
  end_ip        = "255.0.0.0"
  name          = "var0%[1]s"
  start_ip      = "12.0.0.0"
  subnet        = "12.0.0.0 255.0.0.0"
  type          = "ipmask"
  visibility    = "enable"
  wildcard      = "12.0.0.0 255.0.0.0"
}

resource "fortios_firewall_addrgrp" "trname" {
  allow_routing = "disable"
  color         = 0
  exclude       = "disable"
  name          = "%[1]s"
  visibility    = "enable"

  member {
    name = fortios_firewall_address.trname1.name
  }
}
`, name)
}
