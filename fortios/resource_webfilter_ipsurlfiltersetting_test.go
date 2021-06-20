// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSWebfilterIpsUrlfilterSetting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWebfilterIpsUrlfilterSetting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWebfilterIpsUrlfilterSettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWebfilterIpsUrlfilterSettingExists("fortios_webfilter_ipsurlfiltersetting.trname"),
					resource.TestCheckResourceAttr("fortios_webfilter_ipsurlfiltersetting.trname", "distance", "1"),
					resource.TestCheckResourceAttr("fortios_webfilter_ipsurlfiltersetting.trname", "gateway", "0.0.0.0"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWebfilterIpsUrlfilterSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WebfilterIpsUrlfilterSetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebfilterIpsUrlfilterSetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWebfilterIpsUrlfilterSetting(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading WebfilterIpsUrlfilterSetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WebfilterIpsUrlfilterSetting: %s", n)
		}

		return nil
	}
}

func testAccCheckWebfilterIpsUrlfilterSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_webfilter_ipsurlfiltersetting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWebfilterIpsUrlfilterSetting(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WebfilterIpsUrlfilterSetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWebfilterIpsUrlfilterSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_webfilter_ipsurlfiltersetting" "trname" {
  distance = 1
  gateway  = "0.0.0.0"
}
`)
}
