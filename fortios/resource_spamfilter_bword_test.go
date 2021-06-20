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

func TestAccFortiOSSpamfilterBword_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSpamfilterBword_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSpamfilterBwordConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSpamfilterBwordExists("fortios_spamfilter_bword.trname"),
					resource.TestCheckResourceAttr("fortios_spamfilter_bword.trname", "comment", "test"),
					resource.TestCheckResourceAttr("fortios_spamfilter_bword.trname", "fosid", "1"),
					resource.TestCheckResourceAttr("fortios_spamfilter_bword.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_spamfilter_bword.trname", "entries.0.action", "clear"),
					resource.TestCheckResourceAttr("fortios_spamfilter_bword.trname", "entries.0.language", "western"),
					resource.TestCheckResourceAttr("fortios_spamfilter_bword.trname", "entries.0.pattern", "test*patten"),
					resource.TestCheckResourceAttr("fortios_spamfilter_bword.trname", "entries.0.pattern_type", "wildcard"),
					resource.TestCheckResourceAttr("fortios_spamfilter_bword.trname", "entries.0.score", "10"),
					resource.TestCheckResourceAttr("fortios_spamfilter_bword.trname", "entries.0.status", "enable"),
					resource.TestCheckResourceAttr("fortios_spamfilter_bword.trname", "entries.0.where", "subject"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSpamfilterBwordExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SpamfilterBword: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SpamfilterBword is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSpamfilterBword(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading SpamfilterBword: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SpamfilterBword: %s", n)
		}

		return nil
	}
}

func testAccCheckSpamfilterBwordDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_spamfilter_bword" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSpamfilterBword(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SpamfilterBword %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSpamfilterBwordConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_spamfilter_bword" "trname" {
  comment = "test"
  fosid   = 1
  name    = "%[1]s"

  entries {
    action       = "clear"
    language     = "western"
    pattern      = "test*patten"
    pattern_type = "wildcard"
    score        = 10
    status       = "enable"
    where        = "subject"
  }
}
`, name)
}
