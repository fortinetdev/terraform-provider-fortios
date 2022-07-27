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

func TestAccFortiOSFirewallVipgrp46_basic(t *testing.T) {
	rname := acctest.RandString(8)
	var0 := "var0" + rname
	log.Printf(var0)
	log.Printf("TestAccFortiOSFirewallVipgrp46_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallVipgrp46Config(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallVipgrp46Exists("fortios_firewall_vipgrp46.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_vipgrp46.trname", "color", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_vipgrp46.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_vipgrp46.trname", "member.0.name", var0),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallVipgrp46Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallVipgrp46: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallVipgrp46 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallVipgrp46(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallVipgrp46: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallVipgrp46: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallVipgrp46Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_vipgrp46" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallVipgrp46(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallVipgrp46 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallVipgrp46Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_vip46" "trname1" {
  arp_reply   = "enable"
  color       = 0
  extip       = "10.202.1.100"
  extport     = "0-65535"
  fosid       = 0
  ldb_method  = "static"
  mappedip    = "2001:1:1:2::100"
  mappedport  = "0-65535"
  name        = "var0%[1]s"
  portforward = "disable"
  protocol    = "tcp"
  type        = "static-nat"
}

resource "fortios_firewall_vipgrp46" "trname" {
  color = 0
  name  = "%[1]s"

  member {
    name = fortios_firewall_vip46.trname1.name
  }
}
`, name)
}
