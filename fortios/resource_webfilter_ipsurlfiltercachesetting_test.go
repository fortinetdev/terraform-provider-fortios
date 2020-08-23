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

func TestAccFortiOSWebfilterIpsUrlfilterCacheSetting_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSWebfilterIpsUrlfilterCacheSetting_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSWebfilterIpsUrlfilterCacheSettingConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSWebfilterIpsUrlfilterCacheSettingExists("fortios_webfilter_ipsurlfiltercachesetting.trname"),
					resource.TestCheckResourceAttr("fortios_webfilter_ipsurlfiltercachesetting.trname", "dns_retry_interval", "0"),
					resource.TestCheckResourceAttr("fortios_webfilter_ipsurlfiltercachesetting.trname", "extended_ttl", "0"),
				),
			},
		},
	})
}

func testAccCheckFortiOSWebfilterIpsUrlfilterCacheSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WebfilterIpsUrlfilterCacheSetting: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebfilterIpsUrlfilterCacheSetting is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWebfilterIpsUrlfilterCacheSetting(i)

		if err != nil {
			return fmt.Errorf("Error reading WebfilterIpsUrlfilterCacheSetting: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WebfilterIpsUrlfilterCacheSetting: %s", n)
		}

		return nil
	}
}

func testAccCheckWebfilterIpsUrlfilterCacheSettingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_webfilter_ipsurlfiltercachesetting" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWebfilterIpsUrlfilterCacheSetting(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WebfilterIpsUrlfilterCacheSetting %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWebfilterIpsUrlfilterCacheSettingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_webfilter_ipsurlfiltercachesetting" "trname" {
  dns_retry_interval = 0
  extended_ttl       = 0
}
`)
}
