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

func TestAccFortiOSIpsSettings_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSIpsSettings_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSIpsSettingsConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSIpsSettingsExists("fortios_ips_settings.trname"),
					resource.TestCheckResourceAttr("fortios_ips_settings.trname", "ips_packet_quota", "0"),
					resource.TestCheckResourceAttr("fortios_ips_settings.trname", "packet_log_history", "1"),
					resource.TestCheckResourceAttr("fortios_ips_settings.trname", "packet_log_memory", "256"),
					resource.TestCheckResourceAttr("fortios_ips_settings.trname", "packet_log_post_attack", "0"),
				),
			},
		},
	})
}

func testAccCheckFortiOSIpsSettingsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found IpsSettings: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No IpsSettings is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadIpsSettings(i)

		if err != nil {
			return fmt.Errorf("Error reading IpsSettings: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating IpsSettings: %s", n)
		}

		return nil
	}
}

func testAccCheckIpsSettingsDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_ips_settings" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadIpsSettings(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error IpsSettings %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSIpsSettingsConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_ips_settings" "trname" {
  ips_packet_quota       = 0
  packet_log_history     = 1
  packet_log_memory      = 256
  packet_log_post_attack = 0
}
`)
}
