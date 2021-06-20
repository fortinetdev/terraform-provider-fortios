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

func TestAccFortiOSSpamfilterMheader_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSpamfilterMheader_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSpamfilterMheaderConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSpamfilterMheaderExists("fortios_spamfilter_mheader.trname"),
					resource.TestCheckResourceAttr("fortios_spamfilter_mheader.trname", "comment", "test"),
					resource.TestCheckResourceAttr("fortios_spamfilter_mheader.trname", "fosid", "1"),
					resource.TestCheckResourceAttr("fortios_spamfilter_mheader.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_spamfilter_mheader.trname", "entries.0.action", "spam"),
					resource.TestCheckResourceAttr("fortios_spamfilter_mheader.trname", "entries.0.fieldbody", "scstest"),
					resource.TestCheckResourceAttr("fortios_spamfilter_mheader.trname", "entries.0.fieldname", "EIWEtest"),
					resource.TestCheckResourceAttr("fortios_spamfilter_mheader.trname", "entries.0.id", "1"),
					resource.TestCheckResourceAttr("fortios_spamfilter_mheader.trname", "entries.0.pattern_type", "wildcard"),
					resource.TestCheckResourceAttr("fortios_spamfilter_mheader.trname", "entries.0.status", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSpamfilterMheaderExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SpamfilterMheader: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SpamfilterMheader is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSpamfilterMheader(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading SpamfilterMheader: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SpamfilterMheader: %s", n)
		}

		return nil
	}
}

func testAccCheckSpamfilterMheaderDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_spamfilter_mheader" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSpamfilterMheader(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SpamfilterMheader %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSpamfilterMheaderConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_spamfilter_mheader" "trname" {
  comment = "test"
  fosid   = 1
  name    = "%[1]s"

  entries {
    action       = "spam"
    fieldbody    = "scstest"
    fieldname    = "EIWEtest"
    id           = 1
    pattern_type = "wildcard"
    status       = "enable"
  }
}
`, name)
}
