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

func TestAccFortiOSLogSyslogd4Filter_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogSyslogd4Filter_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogSyslogd4FilterConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogSyslogd4FilterExists("fortios_logsyslogd4_filter.trname"),
					resource.TestCheckResourceAttr("fortios_logsyslogd4_filter.trname", "anomaly", "enable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd4_filter.trname", "dns", "enable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd4_filter.trname", "filter_type", "include"),
					resource.TestCheckResourceAttr("fortios_logsyslogd4_filter.trname", "forward_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd4_filter.trname", "gtp", "enable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd4_filter.trname", "local_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd4_filter.trname", "multicast_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd4_filter.trname", "severity", "information"),
					resource.TestCheckResourceAttr("fortios_logsyslogd4_filter.trname", "sniffer_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd4_filter.trname", "ssh", "enable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd4_filter.trname", "voip", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogSyslogd4FilterExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogSyslogd4Filter: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogSyslogd4Filter is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogSyslogd4Filter(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading LogSyslogd4Filter: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogSyslogd4Filter: %s", n)
		}

		return nil
	}
}

func testAccCheckLogSyslogd4FilterDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_logsyslogd4_filter" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogSyslogd4Filter(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogSyslogd4Filter %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogSyslogd4FilterConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_logsyslogd4_filter" "trname" {
  anomaly           = "enable"
  dns               = "enable"
  filter_type       = "include"
  forward_traffic   = "enable"
  gtp               = "enable"
  local_traffic     = "enable"
  multicast_traffic = "enable"
  severity          = "information"
  sniffer_traffic   = "enable"
  ssh               = "enable"
  voip              = "enable"
}
`)
}
