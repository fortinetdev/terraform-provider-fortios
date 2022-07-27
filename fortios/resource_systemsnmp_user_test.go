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

func TestAccFortiOSSystemSnmpUser_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemSnmpUser_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemSnmpUserConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemSnmpUserExists("fortios_systemsnmp_user.trname"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_user.trname", "auth_proto", "sha"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_user.trname", "events", "cpu-high mem-low log-full intf-ip vpn-tun-up vpn-tun-down ha-switch ha-hb-failure ips-signature ips-anomaly av-virus av-oversize av-pattern av-fragmented fm-if-change bgp-established bgp-backward-transition ha-member-up ha-member-down ent-conf-change av-conserve av-bypass av-oversize-passed av-oversize-blocked ips-pkg-update ips-fail-open faz-disconnect wc-ap-up wc-ap-down fswctl-session-up fswctl-session-down load-balance-real-server-down per-cpu-high"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_user.trname", "ha_direct", "disable"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_user.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_systemsnmp_user.trname", "priv_proto", "aes"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_user.trname", "queries", "disable"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_user.trname", "query_port", "161"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_user.trname", "security_level", "no-auth-no-priv"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_user.trname", "source_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_user.trname", "source_ipv6", "::"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_user.trname", "status", "disable"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_user.trname", "trap_lport", "162"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_user.trname", "trap_rport", "162"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_user.trname", "trap_status", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemSnmpUserExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemSnmpUser: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemSnmpUser is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemSnmpUser(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemSnmpUser: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemSnmpUser: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemSnmpUserDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_systemsnmp_user" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemSnmpUser(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemSnmpUser %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemSnmpUserConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_systemsnmp_user" "trname" {
  auth_proto     = "sha"
  events         = "cpu-high mem-low log-full intf-ip vpn-tun-up vpn-tun-down ha-switch ha-hb-failure ips-signature ips-anomaly av-virus av-oversize av-pattern av-fragmented fm-if-change bgp-established bgp-backward-transition ha-member-up ha-member-down ent-conf-change av-conserve av-bypass av-oversize-passed av-oversize-blocked ips-pkg-update ips-fail-open faz-disconnect wc-ap-up wc-ap-down fswctl-session-up fswctl-session-down load-balance-real-server-down per-cpu-high"
  ha_direct      = "disable"
  name           = "%[1]s"
  priv_proto     = "aes"
  queries        = "disable"
  query_port     = 161
  security_level = "no-auth-no-priv"
  source_ip      = "0.0.0.0"
  source_ipv6    = "::"
  status         = "disable"
  trap_lport     = 162
  trap_rport     = 162
  trap_status    = "enable"
}
`, name)
}
