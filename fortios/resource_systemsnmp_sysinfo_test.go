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

func TestAccFortiOSSystemSnmpSysinfo_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemSnmpSysinfo_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemSnmpSysinfoConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemSnmpSysinfoExists("fortios_systemsnmp_sysinfo.trname"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_sysinfo.trname", "status", "disable"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_sysinfo.trname", "trap_high_cpu_threshold", "80"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_sysinfo.trname", "trap_log_full_threshold", "90"),
					resource.TestCheckResourceAttr("fortios_systemsnmp_sysinfo.trname", "trap_low_memory_threshold", "80"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemSnmpSysinfoExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemSnmpSysinfo: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemSnmpSysinfo is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemSnmpSysinfo(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemSnmpSysinfo: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemSnmpSysinfo: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemSnmpSysinfoDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_systemsnmp_sysinfo" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemSnmpSysinfo(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemSnmpSysinfo %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemSnmpSysinfoConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_systemsnmp_sysinfo" "trname" {
  status                    = "disable"
  trap_high_cpu_threshold   = 80
  trap_log_full_threshold   = 90
  trap_low_memory_threshold = 80
}

`)
}
