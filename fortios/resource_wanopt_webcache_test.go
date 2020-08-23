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

func TestAccFortiOSWanoptWebcache_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWanoptWebcache_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWanoptWebcacheConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWanoptWebcacheExists("fortios_wanopt_webcache.trname"),
					resource.TestCheckResourceAttr("fortios_wanopt_webcache.trname", "always_revalidate", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_webcache.trname", "cache_by_default", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_webcache.trname", "cache_cookie", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_webcache.trname", "cache_expired", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_webcache.trname", "default_ttl", "1440"),
					resource.TestCheckResourceAttr("fortios_wanopt_webcache.trname", "external", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_webcache.trname", "fresh_factor", "100"),
					resource.TestCheckResourceAttr("fortios_wanopt_webcache.trname", "host_validate", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_webcache.trname", "ignore_conditional", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_webcache.trname", "ignore_ie_reload", "enable"),
					resource.TestCheckResourceAttr("fortios_wanopt_webcache.trname", "ignore_ims", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_webcache.trname", "ignore_pnc", "disable"),
					resource.TestCheckResourceAttr("fortios_wanopt_webcache.trname", "max_object_size", "512000"),
					resource.TestCheckResourceAttr("fortios_wanopt_webcache.trname", "max_ttl", "7200"),
					resource.TestCheckResourceAttr("fortios_wanopt_webcache.trname", "min_ttl", "5"),
					resource.TestCheckResourceAttr("fortios_wanopt_webcache.trname", "neg_resp_time", "0"),
					resource.TestCheckResourceAttr("fortios_wanopt_webcache.trname", "reval_pnc", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWanoptWebcacheExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WanoptWebcache: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WanoptWebcache is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWanoptWebcache(i)

		if err != nil {
			return fmt.Errorf("Error reading WanoptWebcache: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WanoptWebcache: %s", n)
		}

		return nil
	}
}

func testAccCheckWanoptWebcacheDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_wanopt_webcache" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWanoptWebcache(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WanoptWebcache %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWanoptWebcacheConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_wanopt_webcache" "trname" {
  always_revalidate  = "disable"
  cache_by_default   = "disable"
  cache_cookie       = "disable"
  cache_expired      = "disable"
  default_ttl        = 1440
  external           = "disable"
  fresh_factor       = 100
  host_validate      = "disable"
  ignore_conditional = "disable"
  ignore_ie_reload   = "enable"
  ignore_ims         = "disable"
  ignore_pnc         = "disable"
  max_object_size    = 512000
  max_ttl            = 7200
  min_ttl            = 5
  neg_resp_time      = 0
  reval_pnc          = "disable"
}

`)
}
