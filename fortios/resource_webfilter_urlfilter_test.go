
// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios
import (
    "fmt"
	"log"
    "testing"
    "github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccFortiOSWebfilterUrlfilter_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSWebfilterUrlfilter_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSWebfilterUrlfilterConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSWebfilterUrlfilterExists("fortios_webfilter_urlfilter.trname"),
                    resource.TestCheckResourceAttr("fortios_webfilter_urlfilter.trname", "fosid", "1"),
                    resource.TestCheckResourceAttr("fortios_webfilter_urlfilter.trname", "ip_addr_block", "enable"),
                    resource.TestCheckResourceAttr("fortios_webfilter_urlfilter.trname", "name", rname),
                    resource.TestCheckResourceAttr("fortios_webfilter_urlfilter.trname", "one_arm_ips_urlfilter", "enable"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSWebfilterUrlfilterExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found WebfilterUrlfilter: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WebfilterUrlfilter is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadWebfilterUrlfilter(i)

		if err != nil {
			return fmt.Errorf("Error reading WebfilterUrlfilter: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating WebfilterUrlfilter: %s", n)
		}

		return nil
	}
}

func testAccCheckWebfilterUrlfilterDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_webfilter_urlfilter" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadWebfilterUrlfilter(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error WebfilterUrlfilter %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSWebfilterUrlfilterConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_webfilter_urlfilter" "trname" {
  fosid                 = 1
  ip_addr_block         = "enable"
  name                  = "%[1]s"
  one_arm_ips_urlfilter = "enable"
}

`, name)
}
