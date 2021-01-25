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

func TestAccFortiOSSystemIpv6NeighborCache_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemIpv6NeighborCache_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemIpv6NeighborCacheConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemIpv6NeighborCacheExists("fortios_system_ipv6neighborcache.trname"),
					resource.TestCheckResourceAttr("fortios_system_ipv6neighborcache.trname", "fosid", "1"),
					resource.TestCheckResourceAttr("fortios_system_ipv6neighborcache.trname", "interface", "port2"),
					resource.TestCheckResourceAttr("fortios_system_ipv6neighborcache.trname", "ipv6", "fe80::b11a:5ae3:198:ba1c"),
					resource.TestCheckResourceAttr("fortios_system_ipv6neighborcache.trname", "mac", "00:00:00:00:00:00"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemIpv6NeighborCacheExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemIpv6NeighborCache: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemIpv6NeighborCache is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemIpv6NeighborCache(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemIpv6NeighborCache: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemIpv6NeighborCache: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemIpv6NeighborCacheDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_ipv6neighborcache" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemIpv6NeighborCache(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemIpv6NeighborCache %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemIpv6NeighborCacheConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_ipv6neighborcache" "trname" {
  fosid     = 1
  interface = "port2"
  ipv6      = "fe80::b11a:5ae3:198:ba1c"
  mac       = "00:00:00:00:00:00"
}
`)
}
