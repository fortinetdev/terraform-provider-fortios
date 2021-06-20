// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSSystemFortiguard_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemFortiguard_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemFortiguardConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemFortiguardExists("fortios_system_fortiguard.trname"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "antispam_cache", "enable"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "antispam_cache_mpercent", "2"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "antispam_cache_ttl", "1800"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "antispam_expiration", "1618617600"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "antispam_force_off", "disable"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "antispam_license", "1"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "antispam_timeout", "7"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "auto_join_forticloud", "enable"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "ddns_server_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "ddns_server_port", "443"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "load_balance_servers", "1"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "outbreak_prevention_cache", "enable"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "outbreak_prevention_cache_mpercent", "2"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "outbreak_prevention_cache_ttl", "300"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "outbreak_prevention_expiration", "1618617600"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "outbreak_prevention_force_off", "disable"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "outbreak_prevention_license", "1"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "outbreak_prevention_timeout", "7"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "port", "8888"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "sdns_server_ip", "\"208.91.112.220\" "),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "sdns_server_port", "53"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "source_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "source_ip6", "::"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "update_server_location", "usa"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "webfilter_cache", "enable"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "webfilter_cache_ttl", "3600"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "webfilter_expiration", "1618617600"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "webfilter_force_off", "disable"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "webfilter_license", "1"),
					resource.TestCheckResourceAttr("fortios_system_fortiguard.trname", "webfilter_timeout", "15"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemFortiguardExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemFortiguard: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemFortiguard is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemFortiguard(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading SystemFortiguard: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemFortiguard: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemFortiguardDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_fortiguard" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemFortiguard(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemFortiguard %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemFortiguardConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_fortiguard" "trname" {
  antispam_cache                     = "enable"
  antispam_cache_mpercent            = 2
  antispam_cache_ttl                 = 1800
  antispam_expiration                = 1618617600
  antispam_force_off                 = "disable"
  antispam_license                   = 1
  antispam_timeout                   = 7
  auto_join_forticloud               = "enable"
  ddns_server_ip                     = "0.0.0.0"
  ddns_server_port                   = 443
  load_balance_servers               = 1
  outbreak_prevention_cache          = "enable"
  outbreak_prevention_cache_mpercent = 2
  outbreak_prevention_cache_ttl      = 300
  outbreak_prevention_expiration     = 1618617600
  outbreak_prevention_force_off      = "disable"
  outbreak_prevention_license        = 1
  outbreak_prevention_timeout        = 7
  port                               = "8888"
  sdns_server_ip                     = "\"208.91.112.220\" "
  sdns_server_port                   = 53
  source_ip                          = "0.0.0.0"
  source_ip6                         = "::"
  update_server_location             = "usa"
  webfilter_cache                    = "enable"
  webfilter_cache_ttl                = 3600
  webfilter_expiration               = 1618617600
  webfilter_force_off                = "disable"
  webfilter_license                  = 1
  webfilter_timeout                  = 15
}
`)
}
