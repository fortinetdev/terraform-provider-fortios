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

func TestAccFortiOSVpnPptp_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSVpnPptp_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSVpnPptpConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSVpnPptpExists("fortios_vpn_pptp.trname"),
					resource.TestCheckResourceAttr("fortios_vpn_pptp.trname", "eip", "1.1.1.22"),
					resource.TestCheckResourceAttr("fortios_vpn_pptp.trname", "ip_mode", "range"),
					resource.TestCheckResourceAttr("fortios_vpn_pptp.trname", "local_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_vpn_pptp.trname", "sip", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_vpn_pptp.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_vpn_pptp.trname", "usrgrp", "Guest-group"),
				),
			},
		},
	})
}

func testAccCheckFortiOSVpnPptpExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found VpnPptp: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VpnPptp is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadVpnPptp(i)

		if err != nil {
			return fmt.Errorf("Error reading VpnPptp: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating VpnPptp: %s", n)
		}

		return nil
	}
}

func testAccCheckVpnPptpDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_vpn_pptp" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadVpnPptp(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error VpnPptp %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSVpnPptpConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_vpn_pptp" "trname" {
  eip      = "1.1.1.22"
  ip_mode  = "range"
  local_ip = "0.0.0.0"
  sip      = "1.1.1.1"
  status   = "enable"
  usrgrp   = "Guest-group"
}
`)
}
