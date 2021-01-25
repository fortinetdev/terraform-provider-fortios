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

func TestAccFortiOSLogFortiguardSetting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSLogFortiguardSetting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSLogFortiguardSettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSLogFortiguardSettingExists("fortios_logfortiguard_setting.trname"),
					resource.TestCheckResourceAttr("fortios_logfortiguard_setting.trname", "enc_algorithm", "high"),
					resource.TestCheckResourceAttr("fortios_logfortiguard_setting.trname", "source_ip", "0.0.0.0"),
					resource.TestCheckResourceAttr("fortios_logfortiguard_setting.trname", "ssl_min_proto_version", "default"),
					resource.TestCheckResourceAttr("fortios_logfortiguard_setting.trname", "status", "disable"),
					resource.TestCheckResourceAttr("fortios_logfortiguard_setting.trname", "upload_interval", "daily"),
					resource.TestCheckResourceAttr("fortios_logfortiguard_setting.trname", "upload_option", "5-minute"),
					resource.TestCheckResourceAttr("fortios_logfortiguard_setting.trname", "upload_time", "00:00"),
				),
			},
		},
	})
}

func testAccCheckFortiOSLogFortiguardSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found LogFortiguardSetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LogFortiguardSetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadLogFortiguardSetting(i)

		if err != nil {
			return fmt.Errorf("Error reading LogFortiguardSetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating LogFortiguardSetting: %s", n)
		}

		return nil
	}
}

func testAccCheckLogFortiguardSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_logfortiguard_setting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadLogFortiguardSetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error LogFortiguardSetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSLogFortiguardSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_logfortiguard_setting" "trname" {
  enc_algorithm         = "high"
  source_ip             = "0.0.0.0"
  ssl_min_proto_version = "default"
  status                = "disable"
  upload_interval       = "daily"
  upload_option         = "5-minute"
  upload_time           = "00:00"
}
`)
}
