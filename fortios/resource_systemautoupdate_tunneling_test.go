
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

func TestAccFortiOSSystemAutoupdateTunneling_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSystemAutoupdateTunneling_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSystemAutoupdateTunnelingConfig(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSystemAutoupdateTunnelingExists("fortios_systemautoupdate_tunneling.trname"),
                    resource.TestCheckResourceAttr("fortios_systemautoupdate_tunneling.trname", "port", "0"),
                    resource.TestCheckResourceAttr("fortios_systemautoupdate_tunneling.trname", "status", "disable"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSystemAutoupdateTunnelingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemAutoupdateTunneling: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemAutoupdateTunneling is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemAutoupdateTunneling(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemAutoupdateTunneling: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemAutoupdateTunneling: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemAutoupdateTunnelingDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_systemautoupdate_tunneling" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemAutoupdateTunneling(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemAutoupdateTunneling %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemAutoupdateTunnelingConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_systemautoupdate_tunneling" "trname" {
  port   = 0
  status = "disable"
}
`)
}
