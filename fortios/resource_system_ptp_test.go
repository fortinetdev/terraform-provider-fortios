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

func TestAccFortiOSSystemPtp_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemPtp_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemPtpConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemPtpExists("fortios_system_ptp.trname"),
					resource.TestCheckResourceAttr("fortios_system_ptp.trname", "delay_mechanism", "E2E"),
					resource.TestCheckResourceAttr("fortios_system_ptp.trname", "interface", "port3"),
					resource.TestCheckResourceAttr("fortios_system_ptp.trname", "mode", "multicast"),
					resource.TestCheckResourceAttr("fortios_system_ptp.trname", "request_interval", "1"),
					resource.TestCheckResourceAttr("fortios_system_ptp.trname", "status", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemPtpExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemPtp: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemPtp is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemPtp(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemPtp: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemPtp: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemPtpDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_ptp" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemPtp(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemPtp %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemPtpConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_ptp" "trname" {
  delay_mechanism  = "E2E"
  interface        = "port3"
  mode             = "multicast"
  request_interval = 1
  status           = "enable"
}
`)
}
