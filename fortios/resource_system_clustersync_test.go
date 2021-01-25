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

func TestAccFortiOSSystemClusterSync_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemClusterSync_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemClusterSyncConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemClusterSyncExists("fortios_system_clustersync.trname"),
					resource.TestCheckResourceAttr("fortios_system_clustersync.trname", "hb_interval", "3"),
					resource.TestCheckResourceAttr("fortios_system_clustersync.trname", "hb_lost_threshold", "3"),
					resource.TestCheckResourceAttr("fortios_system_clustersync.trname", "peerip", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_system_clustersync.trname", "peervd", "root"),
					resource.TestCheckResourceAttr("fortios_system_clustersync.trname", "slave_add_ike_routes", "enable"),
					resource.TestCheckResourceAttr("fortios_system_clustersync.trname", "sync_id", "1"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemClusterSyncExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemClusterSync: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemClusterSync is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemClusterSync(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemClusterSync: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemClusterSync: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemClusterSyncDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_clustersync" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemClusterSync(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemClusterSync %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemClusterSyncConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_clustersync" "trname" {
  hb_interval          = 3
  hb_lost_threshold    = 3
  peerip               = "1.1.1.1"
  peervd               = "root"
  slave_add_ike_routes = "enable"
  sync_id              = 1
}
`)
}
