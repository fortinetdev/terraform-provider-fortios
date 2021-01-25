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

func TestAccFortiOSSystemNetworkVisibility_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemNetworkVisibility_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemNetworkVisibilityConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemNetworkVisibilityExists("fortios_system_networkvisibility.trname"),
					resource.TestCheckResourceAttr("fortios_system_networkvisibility.trname", "destination_hostname_visibility", "enable"),
					resource.TestCheckResourceAttr("fortios_system_networkvisibility.trname", "destination_location", "enable"),
					resource.TestCheckResourceAttr("fortios_system_networkvisibility.trname", "destination_visibility", "enable"),
					resource.TestCheckResourceAttr("fortios_system_networkvisibility.trname", "hostname_limit", "5000"),
					resource.TestCheckResourceAttr("fortios_system_networkvisibility.trname", "hostname_ttl", "86400"),
					resource.TestCheckResourceAttr("fortios_system_networkvisibility.trname", "source_location", "enable"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemNetworkVisibilityExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemNetworkVisibility: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemNetworkVisibility is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemNetworkVisibility(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemNetworkVisibility: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemNetworkVisibility: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemNetworkVisibilityDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_networkvisibility" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemNetworkVisibility(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemNetworkVisibility %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemNetworkVisibilityConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_networkvisibility" "trname" {
  destination_hostname_visibility = "enable"
  destination_location            = "enable"
  destination_visibility          = "enable"
  hostname_limit                  = 5000
  hostname_ttl                    = 86400
  source_location                 = "enable"
}
`)
}
