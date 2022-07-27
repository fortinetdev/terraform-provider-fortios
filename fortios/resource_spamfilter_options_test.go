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

func TestAccFortiOSSpamfilterOptions_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSpamfilterOptions_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSpamfilterOptionsConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSpamfilterOptionsExists("fortios_spamfilter_options.trname"),
					resource.TestCheckResourceAttr("fortios_spamfilter_options.trname", "dns_timeout", "7"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSpamfilterOptionsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SpamfilterOptions: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SpamfilterOptions is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSpamfilterOptions(i)

		if err != nil {
			return fmt.Errorf("Error reading SpamfilterOptions: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SpamfilterOptions: %s", n)
		}

		return nil
	}
}

func testAccCheckSpamfilterOptionsDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_spamfilter_options" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSpamfilterOptions(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SpamfilterOptions %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSpamfilterOptionsConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_spamfilter_options" "trname" {
  dns_timeout = 7
}
`)
}
