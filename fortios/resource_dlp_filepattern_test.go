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

func TestAccFortiOSDlpFilepattern_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSDlpFilepattern_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSDlpFilepatternConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSDlpFilepatternExists("fortios_dlp_filepattern.trname"),
					resource.TestCheckResourceAttr("fortios_dlp_filepattern.trname", "fosid", "9"),
					resource.TestCheckResourceAttr("fortios_dlp_filepattern.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSDlpFilepatternExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found DlpFilepattern: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No DlpFilepattern is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadDlpFilepattern(i)

		if err != nil {
			return fmt.Errorf("Error reading DlpFilepattern: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating DlpFilepattern: %s", n)
		}

		return nil
	}
}

func testAccCheckDlpFilepatternDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_dlp_filepattern" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadDlpFilepattern(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error DlpFilepattern %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSDlpFilepatternConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_dlp_filepattern" "trname" {
  fosid = 9
  name  = "%[1]s"
}
`, name)
}
