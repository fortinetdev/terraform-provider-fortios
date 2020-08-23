
// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios
import (
    "fmt"
	"log"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSFirewallVipgrp64_basic(t *testing.T) {
    rname := acctest.RandString(8)
    var0 := "var0" + rname
    log.Printf(var0)
    log.Printf("TestAccFortiOSFirewallVipgrp64_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSFirewallVipgrp64Config(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSFirewallVipgrp64Exists("fortios_firewall_vipgrp64.trname"),
                    resource.TestCheckResourceAttr("fortios_firewall_vipgrp64.trname", "color", "0"),
                    resource.TestCheckResourceAttr("fortios_firewall_vipgrp64.trname", "name", rname),
                    resource.TestCheckResourceAttr("fortios_firewall_vipgrp64.trname", "member.0.name", var0),
                ),
            },
        },
    })
}

func testAccCheckFortiOSFirewallVipgrp64Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallVipgrp64: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallVipgrp64 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallVipgrp64(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallVipgrp64: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallVipgrp64: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallVipgrp64Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_vipgrp64" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallVipgrp64(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallVipgrp64 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallVipgrp64Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_vip64" "trname1" {
  arp_reply   = "enable"
  color       = 0
  extip       = "2001:db8:99:503::22"
  extport     = "0-65535"
  fosid       = 0
  ldb_method  = "static"
  mappedip    = "1.1.3.1"
  mappedport  = "0-65535"
  name        = "var0%[1]s"
  portforward = "disable"
  protocol    = "tcp"
  type        = "static-nat"
}

resource "fortios_firewall_vipgrp64" "trname" {
  color = 0
  name  = "%[1]s"

  member {
    name = fortios_firewall_vip64.trname1.name
  }
}




`, name)
}
