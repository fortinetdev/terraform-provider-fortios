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

func TestAccFortiOSFirewallAddrgrp6_basic(t *testing.T) {
	rname := acctest.RandString(8)
	var0 := "var0" + rname
	log.Printf(var0)
	log.Printf("TestAccFortiOSFirewallAddrgrp6_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallAddrgrp6Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallAddrgrp6Exists("fortios_firewall_addrgrp6.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_addrgrp6.trname", "color", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_addrgrp6.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_addrgrp6.trname", "visibility", "enable"),
					resource.TestCheckResourceAttr("fortios_firewall_addrgrp6.trname", "member.0.name", var0),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallAddrgrp6Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallAddrgrp6: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallAddrgrp6 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallAddrgrp6(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallAddrgrp6: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallAddrgrp6: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallAddrgrp6Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_addrgrp6" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallAddrgrp6(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallAddrgrp6 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallAddrgrp6Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_address6" "trname1" {
  cache_ttl  = 0
  color      = 0
  end_ip     = "::"
  host       = ""
  host_type  = "any"
  ip6        = "fdff:ffff::/120"
  name       = "var0%[1]s"
  start_ip   = ""
  type       = "ipprefix"
  visibility = "enable"
}

resource "fortios_firewall_addrgrp6" "trname" {
  color      = 0
  name       = "%[1]s"
  visibility = "enable"

  member {
    name = fortios_firewall_address6.trname1.name
  }
}
`, name)
}
