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

func TestAccFortiOSSpamfilterBwl_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSpamfilterBwl_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSpamfilterBwlConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSpamfilterBwlExists("fortios_spamfilter_bwl.trname"),
					resource.TestCheckResourceAttr("fortios_spamfilter_bwl.trname", "comment", "test"),
					resource.TestCheckResourceAttr("fortios_spamfilter_bwl.trname", "fosid", "1"),
					resource.TestCheckResourceAttr("fortios_spamfilter_bwl.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_spamfilter_bwl.trname", "entries.0.action", "reject"),
					resource.TestCheckResourceAttr("fortios_spamfilter_bwl.trname", "entries.0.addr_type", "ipv4"),
					resource.TestCheckResourceAttr("fortios_spamfilter_bwl.trname", "entries.0.ip4_subnet", "1.1.1.0 255.255.255.0"),
					resource.TestCheckResourceAttr("fortios_spamfilter_bwl.trname", "entries.0.ip6_subnet", "::/128"),
					resource.TestCheckResourceAttr("fortios_spamfilter_bwl.trname", "entries.0.pattern_type", "wildcard"),
					resource.TestCheckResourceAttr("fortios_spamfilter_bwl.trname", "entries.0.status", "enable"),
					resource.TestCheckResourceAttr("fortios_spamfilter_bwl.trname", "entries.0.type", "ip"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSpamfilterBwlExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SpamfilterBwl: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SpamfilterBwl is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSpamfilterBwl(i)

		if err != nil {
			return fmt.Errorf("Error reading SpamfilterBwl: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SpamfilterBwl: %s", n)
		}

		return nil
	}
}

func testAccCheckSpamfilterBwlDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_spamfilter_bwl" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSpamfilterBwl(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SpamfilterBwl %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSpamfilterBwlConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_spamfilter_bwl" "trname" {
  comment = "test"
  fosid   = 1
  name    = "%[1]s"

  entries {
    action       = "reject"
    addr_type    = "ipv4"
    ip4_subnet   = "1.1.1.0 255.255.255.0"
    ip6_subnet   = "::/128"
    pattern_type = "wildcard"
    status       = "enable"
    type         = "ip"
  }
}
`, name)
}
