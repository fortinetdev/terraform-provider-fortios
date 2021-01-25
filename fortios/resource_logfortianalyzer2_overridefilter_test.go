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

func TestAccFortiOSLogFortianalyzer2OverrideFilter_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogFortianalyzer2OverrideFilter_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogFortianalyzer2OverrideFilterConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogFortianalyzer2OverrideFilterExists("fortios_logfortianalyzer2_overridefilter.trname"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_overridefilter.trname", "anomaly", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_overridefilter.trname", "dlp_archive", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_overridefilter.trname", "dns", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_overridefilter.trname", "filter_type", "include"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_overridefilter.trname", "forward_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_overridefilter.trname", "gtp", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_overridefilter.trname", "local_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_overridefilter.trname", "multicast_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_overridefilter.trname", "severity", "information"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_overridefilter.trname", "sniffer_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_overridefilter.trname", "ssh", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer2_overridefilter.trname", "voip", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogFortianalyzer2OverrideFilterExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogFortianalyzer2OverrideFilter: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogFortianalyzer2OverrideFilter is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogFortianalyzer2OverrideFilter(i)

		if err != nil {
			return fmt.Errorf("Error reading LogFortianalyzer2OverrideFilter: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogFortianalyzer2OverrideFilter: %s", n)
		}

		return nil
	}
}

func testAccCheckLogFortianalyzer2OverrideFilterDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_logfortianalyzer2_overridefilter" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogFortianalyzer2OverrideFilter(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogFortianalyzer2OverrideFilter %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogFortianalyzer2OverrideFilterConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_logfortianalyzer2_overridefilter" "trname" {
  anomaly           = "enable"
  dlp_archive       = "enable"
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
