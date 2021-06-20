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

func TestAccFortiOSDlpFpSensitivity_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSDlpFpSensitivity_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSDlpFpSensitivityConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSDlpFpSensitivityExists("fortios_dlp_fpsensitivity.trname"),
					resource.TestCheckResourceAttr("fortios_dlp_fpsensitivity.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSDlpFpSensitivityExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found DlpFpSensitivity: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No DlpFpSensitivity is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadDlpFpSensitivity(i, "root")

		if err != nil {
			return fmt.Errorf("Error reading DlpFpSensitivity: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating DlpFpSensitivity: %s", n)
		}

		return nil
	}
}

func testAccCheckDlpFpSensitivityDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_dlp_fpsensitivity" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadDlpFpSensitivity(i, "root")

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error DlpFpSensitivity %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSDlpFpSensitivityConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_dlp_fpsensitivity" "trname" {
  name = "%[1]s"
}
`, name)
}
