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

func TestAccFortiOSSystemResourceLimits_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemResourceLimits_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemResourceLimitsConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemResourceLimitsExists("fortios_system_resourcelimits.trname"),
					resource.TestCheckResourceAttr("fortios_system_resourcelimits.trname", "custom_service", "0"),
					resource.TestCheckResourceAttr("fortios_system_resourcelimits.trname", "dialup_tunnel", "0"),
					resource.TestCheckResourceAttr("fortios_system_resourcelimits.trname", "firewall_address", "41024"),
					resource.TestCheckResourceAttr("fortios_system_resourcelimits.trname", "firewall_addrgrp", "10692"),
					resource.TestCheckResourceAttr("fortios_system_resourcelimits.trname", "firewall_policy", "41024"),
					resource.TestCheckResourceAttr("fortios_system_resourcelimits.trname", "ipsec_phase1", "2000"),
					resource.TestCheckResourceAttr("fortios_system_resourcelimits.trname", "ipsec_phase1_interface", "0"),
					resource.TestCheckResourceAttr("fortios_system_resourcelimits.trname", "ipsec_phase2", "2000"),
					resource.TestCheckResourceAttr("fortios_system_resourcelimits.trname", "ipsec_phase2_interface", "0"),
					resource.TestCheckResourceAttr("fortios_system_resourcelimits.trname", "log_disk_quota", "30235"),
					resource.TestCheckResourceAttr("fortios_system_resourcelimits.trname", "onetime_schedule", "0"),
					resource.TestCheckResourceAttr("fortios_system_resourcelimits.trname", "proxy", "64000"),
					resource.TestCheckResourceAttr("fortios_system_resourcelimits.trname", "recurring_schedule", "0"),
					resource.TestCheckResourceAttr("fortios_system_resourcelimits.trname", "service_group", "0"),
					resource.TestCheckResourceAttr("fortios_system_resourcelimits.trname", "session", "0"),
					resource.TestCheckResourceAttr("fortios_system_resourcelimits.trname", "sslvpn", "0"),
					resource.TestCheckResourceAttr("fortios_system_resourcelimits.trname", "user", "0"),
					resource.TestCheckResourceAttr("fortios_system_resourcelimits.trname", "user_group", "0"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemResourceLimitsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemResourceLimits: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemResourceLimits is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemResourceLimits(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemResourceLimits: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemResourceLimits: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemResourceLimitsDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_resourcelimits" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemResourceLimits(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemResourceLimits %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemResourceLimitsConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_resourcelimits" "trname" {
  custom_service         = 0
  dialup_tunnel          = 0
  firewall_address       = 41024
  firewall_addrgrp       = 10692
  firewall_policy        = 41024
  ipsec_phase1           = 2000
  ipsec_phase1_interface = 0
  ipsec_phase2           = 2000
  ipsec_phase2_interface = 0
  log_disk_quota         = 30235
  onetime_schedule       = 0
  proxy                  = 64000
  recurring_schedule     = 0
  service_group          = 0
  session                = 0
  sslvpn                 = 0
  user                   = 0
  user_group             = 0
}
`)
}
