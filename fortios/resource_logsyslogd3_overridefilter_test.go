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

func TestAccFortiOSLogSyslogd3OverrideFilter_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogSyslogd3OverrideFilter_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogSyslogd3OverrideFilterConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogSyslogd3OverrideFilterExists("fortios_logsyslogd3_overridefilter.trname"),
					resource.TestCheckResourceAttr("fortios_logsyslogd3_overridefilter.trname", "anomaly", "enable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd3_overridefilter.trname", "dns", "enable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd3_overridefilter.trname", "filter_type", "include"),
					resource.TestCheckResourceAttr("fortios_logsyslogd3_overridefilter.trname", "forward_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd3_overridefilter.trname", "gtp", "enable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd3_overridefilter.trname", "local_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd3_overridefilter.trname", "multicast_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd3_overridefilter.trname", "severity", "information"),
					resource.TestCheckResourceAttr("fortios_logsyslogd3_overridefilter.trname", "sniffer_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd3_overridefilter.trname", "ssh", "enable"),
					resource.TestCheckResourceAttr("fortios_logsyslogd3_overridefilter.trname", "voip", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogSyslogd3OverrideFilterExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogSyslogd3OverrideFilter: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogSyslogd3OverrideFilter is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogSyslogd3OverrideFilter(i)

		if err != nil {
			return fmt.Errorf("Error reading LogSyslogd3OverrideFilter: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogSyslogd3OverrideFilter: %s", n)
		}

		return nil
	}
}

func testAccCheckLogSyslogd3OverrideFilterDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_logsyslogd3_overridefilter" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogSyslogd3OverrideFilter(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogSyslogd3OverrideFilter %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogSyslogd3OverrideFilterConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_logsyslogd3_overridefilter" "trname" {
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
