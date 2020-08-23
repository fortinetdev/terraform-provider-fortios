
// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios
import (
    "fmt"
	"log"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSLogSyslogd2OverrideFilter_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSLogSyslogd2OverrideFilter_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSLogSyslogd2OverrideFilterConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSLogSyslogd2OverrideFilterExists("fortios_logsyslogd2_overridefilter.trname"),
                    resource.TestCheckResourceAttr("fortios_logsyslogd2_overridefilter.trname", "anomaly", "enable"),
                    resource.TestCheckResourceAttr("fortios_logsyslogd2_overridefilter.trname", "dns", "enable"),
                    resource.TestCheckResourceAttr("fortios_logsyslogd2_overridefilter.trname", "filter_type", "include"),
                    resource.TestCheckResourceAttr("fortios_logsyslogd2_overridefilter.trname", "forward_traffic", "enable"),
                    resource.TestCheckResourceAttr("fortios_logsyslogd2_overridefilter.trname", "gtp", "enable"),
                    resource.TestCheckResourceAttr("fortios_logsyslogd2_overridefilter.trname", "local_traffic", "enable"),
                    resource.TestCheckResourceAttr("fortios_logsyslogd2_overridefilter.trname", "multicast_traffic", "enable"),
                    resource.TestCheckResourceAttr("fortios_logsyslogd2_overridefilter.trname", "severity", "information"),
                    resource.TestCheckResourceAttr("fortios_logsyslogd2_overridefilter.trname", "sniffer_traffic", "enable"),
                    resource.TestCheckResourceAttr("fortios_logsyslogd2_overridefilter.trname", "ssh", "enable"),
                    resource.TestCheckResourceAttr("fortios_logsyslogd2_overridefilter.trname", "voip", "enable"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSLogSyslogd2OverrideFilterExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogSyslogd2OverrideFilter: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogSyslogd2OverrideFilter is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogSyslogd2OverrideFilter(i)

		if err != nil {
			return fmt.Errorf("Error reading LogSyslogd2OverrideFilter: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogSyslogd2OverrideFilter: %s", n)
		}

		return nil
	}
}

func testAccCheckLogSyslogd2OverrideFilterDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_logsyslogd2_overridefilter" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogSyslogd2OverrideFilter(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogSyslogd2OverrideFilter %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogSyslogd2OverrideFilterConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_logsyslogd2_overridefilter" "trname" {
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
