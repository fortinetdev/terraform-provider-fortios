
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

func TestAccFortiOSSpamfilterFortishield_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSpamfilterFortishield_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSpamfilterFortishieldConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSpamfilterFortishieldExists("fortios_spamfilter_fortishield.trname"),
                    resource.TestCheckResourceAttr("fortios_spamfilter_fortishield.trname", "spam_submit_force", "enable"),
                    resource.TestCheckResourceAttr("fortios_spamfilter_fortishield.trname", "spam_submit_srv", "www.nospammer.net"),
                    resource.TestCheckResourceAttr("fortios_spamfilter_fortishield.trname", "spam_submit_txt2htm", "enable"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSpamfilterFortishieldExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SpamfilterFortishield: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SpamfilterFortishield is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSpamfilterFortishield(i)

		if err != nil {
			return fmt.Errorf("Error reading SpamfilterFortishield: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SpamfilterFortishield: %s", n)
		}

		return nil
	}
}

func testAccCheckSpamfilterFortishieldDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_spamfilter_fortishield" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSpamfilterFortishield(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SpamfilterFortishield %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSpamfilterFortishieldConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_spamfilter_fortishield" "trname" {
  spam_submit_force   = "enable"
  spam_submit_srv     = "www.nospammer.net"
  spam_submit_txt2htm = "enable"
}
`)
}
