
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

func TestAccFortiOSSystemNat64_basic(t *testing.T) {
    rname := acctest.RandString(8)
    log.Printf("TestAccFortiOSSystemNat64_basic %s", rname)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccFortiOSSystemNat64Config(rname),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckFortiOSSystemNat64Exists("fortios_system_nat64.trname"),
                    resource.TestCheckResourceAttr("fortios_system_nat64.trname", "always_synthesize_aaaa_record", "enable"),
                    resource.TestCheckResourceAttr("fortios_system_nat64.trname", "generate_ipv6_fragment_header", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_nat64.trname", "nat46_force_ipv4_packet_forwarding", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_nat64.trname", "nat64_prefix", "2001:1:2:3::/96"),
                    resource.TestCheckResourceAttr("fortios_system_nat64.trname", "secondary_prefix_status", "disable"),
                    resource.TestCheckResourceAttr("fortios_system_nat64.trname", "status", "disable"),
                ),
            },
        },
    })
}

func testAccCheckFortiOSSystemNat64Exists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemNat64: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemNat64 is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemNat64(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemNat64: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemNat64: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemNat64Destroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_nat64" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemNat64(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemNat64 %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemNat64Config(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_nat64" "trname" {
  always_synthesize_aaaa_record      = "enable"
  generate_ipv6_fragment_header      = "disable"
  nat46_force_ipv4_packet_forwarding = "disable"
  nat64_prefix                       = "2001:1:2:3::/96"
  secondary_prefix_status            = "disable"
  status                             = "disable"
}
`)
}
