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

func TestAccFortiOSWanoptSettings_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWanoptSettings_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWanoptSettingsConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWanoptSettingsExists("fortios_wanopt_settings.trname"),
					resource.TestCheckResourceAttr("fortios_wanopt_settings.trname", "auto_detect_algorithm", "simple"),
					resource.TestCheckResourceAttr("fortios_wanopt_settings.trname", "host_id", "default-id"),
					resource.TestCheckResourceAttr("fortios_wanopt_settings.trname", "tunnel_ssl_algorithm", "high"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWanoptSettingsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WanoptSettings: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WanoptSettings is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWanoptSettings(i)

		if err != nil {
			return fmt.Errorf("Error reading WanoptSettings: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WanoptSettings: %s", n)
		}

		return nil
	}
}

func testAccCheckWanoptSettingsDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_wanopt_settings" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWanoptSettings(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WanoptSettings %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWanoptSettingsConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_wanopt_settings" "trname" {
  auto_detect_algorithm = "simple"
  host_id               = "default-id"
  tunnel_ssl_algorithm  = "high"
}
`)
}
