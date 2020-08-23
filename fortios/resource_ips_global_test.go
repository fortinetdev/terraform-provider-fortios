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

func TestAccFortiOSIpsGlobal_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSIpsGlobal_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSIpsGlobalConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSIpsGlobalExists("fortios_ips_global.trname"),
					resource.TestCheckResourceAttr("fortios_ips_global.trname", "anomaly_mode", "continuous"),
					resource.TestCheckResourceAttr("fortios_ips_global.trname", "database", "regular"),
					resource.TestCheckResourceAttr("fortios_ips_global.trname", "deep_app_insp_db_limit", "0"),
					resource.TestCheckResourceAttr("fortios_ips_global.trname", "deep_app_insp_timeout", "0"),
					resource.TestCheckResourceAttr("fortios_ips_global.trname", "engine_count", "0"),
					resource.TestCheckResourceAttr("fortios_ips_global.trname", "exclude_signatures", "industrial"),
					resource.TestCheckResourceAttr("fortios_ips_global.trname", "fail_open", "disable"),
					resource.TestCheckResourceAttr("fortios_ips_global.trname", "intelligent_mode", "enable"),
					resource.TestCheckResourceAttr("fortios_ips_global.trname", "session_limit_mode", "heuristic"),
					resource.TestCheckResourceAttr("fortios_ips_global.trname", "socket_size", "0"),
					resource.TestCheckResourceAttr("fortios_ips_global.trname", "sync_session_ttl", "enable"),
					resource.TestCheckResourceAttr("fortios_ips_global.trname", "traffic_submit", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSIpsGlobalExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found IpsGlobal: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No IpsGlobal is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadIpsGlobal(i)

		if err != nil {
			return fmt.Errorf("Error reading IpsGlobal: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating IpsGlobal: %s", n)
		}

		return nil
	}
}

func testAccCheckIpsGlobalDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_ips_global" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadIpsGlobal(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error IpsGlobal %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSIpsGlobalConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_ips_global" "trname" {
  anomaly_mode           = "continuous"
  database               = "regular"
  deep_app_insp_db_limit = 0
  deep_app_insp_timeout  = 0
  engine_count           = 0
  exclude_signatures     = "industrial"
  fail_open              = "disable"
  intelligent_mode       = "enable"
  session_limit_mode     = "heuristic"
  socket_size            = 0
  sync_session_ttl       = "enable"
  traffic_submit         = "disable"
}

`)
}
