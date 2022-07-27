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

func TestAccFortiOSFirewallVipgrp_basic(t *testing.T) {
	rname := acctest.RandString(8)
	var0 := "var0" + rname
	log.Printf(var0)
	log.Printf("TestAccFortiOSFirewallVipgrp_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSFirewallVipgrpConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSFirewallVipgrpExists("fortios_firewall_vipgrp.trname"),
					resource.TestCheckResourceAttr("fortios_firewall_vipgrp.trname", "color", "0"),
					resource.TestCheckResourceAttr("fortios_firewall_vipgrp.trname", "interface", "any"),
					resource.TestCheckResourceAttr("fortios_firewall_vipgrp.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_firewall_vipgrp.trname", "member.0.name", var0),
				),
			},
		},
	})
}

func testAccCheckFortiOSFirewallVipgrpExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found FirewallVipgrp: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No FirewallVipgrp is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadFirewallVipgrp(i)

		if err != nil {
			return fmt.Errorf("Error reading FirewallVipgrp: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating FirewallVipgrp: %s", n)
		}

		return nil
	}
}

func testAccCheckFirewallVipgrpDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_firewall_vipgrp" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadFirewallVipgrp(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error FirewallVipgrp %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSFirewallVipgrpConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_firewall_vip" "trname1" {
  extintf = "any"
  extport = "0-65535"
  extip   = "2.0.0.1-2.0.0.4"
  name    = "var0%[1]s"
  mappedip {
    range = "3.0.0.0-3.0.0.3"
  }
}

resource "fortios_firewall_vipgrp" "trname" {
  color     = 0
  interface = "any"
  name      = "%[1]s"

  member {
    name = fortios_firewall_vip.trname1.name
  }
}
`, name)
}
