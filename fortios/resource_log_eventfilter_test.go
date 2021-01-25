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

func TestAccFortiOSLogEventfilter_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogEventfilter_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogEventfilterConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogEventfilterExists("fortios_log_eventfilter.trname"),
					resource.TestCheckResourceAttr("fortios_log_eventfilter.trname", "compliance_check", "enable"),
					resource.TestCheckResourceAttr("fortios_log_eventfilter.trname", "endpoint", "enable"),
					resource.TestCheckResourceAttr("fortios_log_eventfilter.trname", "event", "enable"),
					resource.TestCheckResourceAttr("fortios_log_eventfilter.trname", "ha", "enable"),
					resource.TestCheckResourceAttr("fortios_log_eventfilter.trname", "router", "enable"),
					resource.TestCheckResourceAttr("fortios_log_eventfilter.trname", "security_rating", "enable"),
					resource.TestCheckResourceAttr("fortios_log_eventfilter.trname", "system", "enable"),
					resource.TestCheckResourceAttr("fortios_log_eventfilter.trname", "user", "enable"),
					resource.TestCheckResourceAttr("fortios_log_eventfilter.trname", "vpn", "enable"),
					resource.TestCheckResourceAttr("fortios_log_eventfilter.trname", "wan_opt", "enable"),
					resource.TestCheckResourceAttr("fortios_log_eventfilter.trname", "wireless_activity", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogEventfilterExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogEventfilter: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogEventfilter is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogEventfilter(i)

		if err != nil {
			return fmt.Errorf("Error reading LogEventfilter: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogEventfilter: %s", n)
		}

		return nil
	}
}

func testAccCheckLogEventfilterDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_log_eventfilter" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogEventfilter(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogEventfilter %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogEventfilterConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_log_eventfilter" "trname" {
  compliance_check  = "enable"
  endpoint          = "enable"
  event             = "enable"
  ha                = "enable"
  router            = "enable"
  security_rating   = "enable"
  system            = "enable"
  user              = "enable"
  vpn               = "enable"
  wan_opt           = "enable"
  wireless_activity = "enable"
}
`)
}
