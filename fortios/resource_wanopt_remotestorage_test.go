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

func TestAccFortiOSWanoptRemoteStorage_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWanoptRemoteStorage_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWanoptRemoteStorageConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWanoptRemoteStorageExists("fortios_wanopt_remotestorage.trname"),
					resource.TestCheckResourceAttr("fortios_wanopt_remotestorage.trname", "remote_cache_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_wanopt_remotestorage.trname", "status", "disable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWanoptRemoteStorageExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WanoptRemoteStorage: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WanoptRemoteStorage is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWanoptRemoteStorage(i)

		if err != nil {
			return fmt.Errorf("Error reading WanoptRemoteStorage: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WanoptRemoteStorage: %s", n)
		}

		return nil
	}
}

func testAccCheckWanoptRemoteStorageDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_wanopt_remotestorage" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWanoptRemoteStorage(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WanoptRemoteStorage %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWanoptRemoteStorageConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_wanopt_remotestorage" "trname" {
  remote_cache_ip = "0.0.0.0"
  status          = "disable"
}
`)
}
