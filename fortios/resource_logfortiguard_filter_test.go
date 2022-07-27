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

func TestAccFortiOSLogFortiguardFilter_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogFortiguardFilter_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogFortiguardFilterConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogFortiguardFilterExists("fortios_logfortiguard_filter.trname"),
					resource.TestCheckResourceAttr("fortios_logfortiguard_filter.trname", "anomaly", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortiguard_filter.trname", "dlp_archive", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortiguard_filter.trname", "dns", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortiguard_filter.trname", "filter_type", "include"),
					resource.TestCheckResourceAttr("fortios_logfortiguard_filter.trname", "forward_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortiguard_filter.trname", "gtp", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortiguard_filter.trname", "local_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortiguard_filter.trname", "multicast_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortiguard_filter.trname", "severity", "information"),
					resource.TestCheckResourceAttr("fortios_logfortiguard_filter.trname", "sniffer_traffic", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortiguard_filter.trname", "ssh", "enable"),
					resource.TestCheckResourceAttr("fortios_logfortiguard_filter.trname", "voip", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogFortiguardFilterExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogFortiguardFilter: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogFortiguardFilter is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogFortiguardFilter(i)

		if err != nil {
			return fmt.Errorf("Error reading LogFortiguardFilter: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogFortiguardFilter: %s", n)
		}

		return nil
	}
}

func testAccCheckLogFortiguardFilterDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_logfortiguard_filter" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogFortiguardFilter(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogFortiguardFilter %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogFortiguardFilterConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_logfortiguard_filter" "trname" {
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
