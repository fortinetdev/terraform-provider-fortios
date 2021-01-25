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

func TestAccFortiOSSystemSnmpCommunity_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemSnmpCommunity_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemSnmpCommunityConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemSnmpCommunityExists("fortios_systemsnmp_community.trname"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_community.trname", "events", "cpu-high mem-low log-full intf-ip vpn-tun-up vpn-tun-down ha-switch ha-hb-failure ips-signature ips-anomaly av-virus av-oversize av-pattern av-fragmented fm-if-change bgp-established bgp-backward-transition ha-member-up ha-member-down ent-conf-change av-conserve av-bypass av-oversize-passed av-oversize-blocked ips-pkg-update ips-fail-open faz-disconnect wc-ap-up wc-ap-down fswctl-session-up fswctl-session-down load-balance-real-server-down per-cpu-high"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_community.trname", "fosid", "1"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_community.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_systemsnmp_community.trname", "query_v1_port", "161"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_community.trname", "query_v1_status", "enable"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_community.trname", "query_v2c_port", "161"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_community.trname", "query_v2c_status", "enable"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_community.trname", "status", "enable"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_community.trname", "trap_v1_lport", "162"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_community.trname", "trap_v1_rport", "162"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_community.trname", "trap_v1_status", "enable"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_community.trname", "trap_v2c_lport", "162"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_community.trname", "trap_v2c_rport", "162"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_community.trname", "trap_v2c_status", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemSnmpCommunityExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemSnmpCommunity: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemSnmpCommunity is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemSnmpCommunity(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemSnmpCommunity: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemSnmpCommunity: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemSnmpCommunityDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_systemsnmp_community" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemSnmpCommunity(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemSnmpCommunity %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemSnmpCommunityConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_systemsnmp_community" "trname" {
  events           = "cpu-high mem-low log-full intf-ip vpn-tun-up vpn-tun-down ha-switch ha-hb-failure ips-signature ips-anomaly av-virus av-oversize av-pattern av-fragmented fm-if-change bgp-established bgp-backward-transition ha-member-up ha-member-down ent-conf-change av-conserve av-bypass av-oversize-passed av-oversize-blocked ips-pkg-update ips-fail-open faz-disconnect wc-ap-up wc-ap-down fswctl-session-up fswctl-session-down load-balance-real-server-down per-cpu-high"
  fosid            = 1
  name             = "%[1]s"
  query_v1_port    = 161
  query_v1_status  = "enable"
  query_v2c_port   = 161
  query_v2c_status = "enable"
  status           = "enable"
  trap_v1_lport    = 162
  trap_v1_rport    = 162
  trap_v1_status   = "enable"
  trap_v2c_lport   = 162
  trap_v2c_rport   = 162
  trap_v2c_status  = "enable"
}
`, name)
}
