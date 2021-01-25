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

func TestAccFortiOSSpamfilterDnsbl_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSpamfilterDnsbl_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSpamfilterDnsblConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSpamfilterDnsblExists("fortios_spamfilter_dnsbl.trname"),
					resource.TestCheckResourceAttr("fortios_spamfilter_dnsbl.trname", "comment", "test"),
					resource.TestCheckResourceAttr("fortios_spamfilter_dnsbl.trname", "fosid", "1"),
					resource.TestCheckResourceAttr("fortios_spamfilter_dnsbl.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_spamfilter_dnsbl.trname", "entries.0.action", "reject"),
					resource.TestCheckResourceAttr("fortios_spamfilter_dnsbl.trname", "entries.0.server", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_spamfilter_dnsbl.trname", "entries.0.status", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSpamfilterDnsblExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SpamfilterDnsbl: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SpamfilterDnsbl is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSpamfilterDnsbl(i)

		if err != nil {
			return fmt.Errorf("Error reading SpamfilterDnsbl: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SpamfilterDnsbl: %s", n)
		}

		return nil
	}
}

func testAccCheckSpamfilterDnsblDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_spamfilter_dnsbl" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSpamfilterDnsbl(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SpamfilterDnsbl %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSpamfilterDnsblConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_spamfilter_dnsbl" "trname" {
  comment = "test"
  fosid   = 1
  name    = "%[1]s"

  entries {
    action = "reject"
    server = "1.1.1.1"
    status = "enable"
  }
}
`, name)
}
