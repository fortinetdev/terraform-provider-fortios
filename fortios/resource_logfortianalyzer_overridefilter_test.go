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

func TestAccFortiOSLogFortianalyzerOverrideFilter_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogFortianalyzerOverrideFilter_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogFortianalyzerOverrideFilterConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogFortianalyzerOverrideFilterExists("fortios_logfortianalyzer_overridefilter.trname"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridefilter.trname", "anomaly", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridefilter.trname", "dlp_archive", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridefilter.trname", "dns", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridefilter.trname", "filter_type", "include"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridefilter.trname", "forward_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridefilter.trname", "gtp", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridefilter.trname", "local_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridefilter.trname", "multicast_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridefilter.trname", "severity", "information"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridefilter.trname", "sniffer_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridefilter.trname", "ssh", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortianalyzer_overridefilter.trname", "voip", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogFortianalyzerOverrideFilterExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogFortianalyzerOverrideFilter: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogFortianalyzerOverrideFilter is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogFortianalyzerOverrideFilter(i)

		if err != nil {
			return fmt.Errorf("Error reading LogFortianalyzerOverrideFilter: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogFortianalyzerOverrideFilter: %s", n)
		}

		return nil
	}
}

func testAccCheckLogFortianalyzerOverrideFilterDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_logfortianalyzer_overridefilter" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogFortianalyzerOverrideFilter(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogFortianalyzerOverrideFilter %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogFortianalyzerOverrideFilterConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_logfortianalyzer_overridefilter" "trname" {
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
