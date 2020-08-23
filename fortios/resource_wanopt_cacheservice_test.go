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

func TestAccFortiOSWanoptCacheService_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWanoptCacheService_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWanoptCacheServiceConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWanoptCacheServiceExists("fortios_wanopt_cacheservice.trname"),
					resource.TestCheckResourceAttr("fortios_wanopt_cacheservice.trname", "acceptable_connections", "any"),
					resource.TestCheckResourceAttr("fortios_wanopt_cacheservice.trname", "collaboration", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_cacheservice.trname", "device_id", "default_dev_id"),
					resource.TestCheckResourceAttr("fortios_wanopt_cacheservice.trname", "prefer_scenario", "balance"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWanoptCacheServiceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WanoptCacheService: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WanoptCacheService is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWanoptCacheService(i)

		if err != nil {
			return fmt.Errorf("Error reading WanoptCacheService: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WanoptCacheService: %s", n)
		}

		return nil
	}
}

func testAccCheckWanoptCacheServiceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_wanopt_cacheservice" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWanoptCacheService(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WanoptCacheService %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWanoptCacheServiceConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_wanopt_cacheservice" "trname" {
  acceptable_connections = "any"
  collaboration          = "disable"
  device_id              = "default_dev_id"
  prefer_scenario        = "balance"
}

`)
}
